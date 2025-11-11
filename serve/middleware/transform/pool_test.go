/*
Copyright © 2025 Benny Powers <web@bennypowers.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

package transform_test

import (
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"bennypowers.dev/cem/serve/middleware/transform"
)

func TestPool_LimitsConcurrentTasks(t *testing.T) {
	// Create pool with max 2 concurrent tasks
	pool := transform.NewPool(2, 10)
	defer pool.Close()

	// Track concurrent executions
	var concurrent int32
	var maxConcurrent int32
	var maxMu sync.Mutex

	// Function that tracks concurrency
	task := func() error {
		current := atomic.AddInt32(&concurrent, 1)
		defer atomic.AddInt32(&concurrent, -1)

		// Update max if needed (using mutex for simplicity in test code)
		maxMu.Lock()
		if current > maxConcurrent {
			maxConcurrent = current
		}
		maxMu.Unlock()

		// Simulate work
		time.Sleep(50 * time.Millisecond)
		return nil
	}

	// Submit 10 tasks
	var wg sync.WaitGroup
	for range 10 {
		wg.Go(func() {
			_ = pool.Submit(task)
		})
	}

	wg.Wait()

	// Verify max concurrent was respected
	maxMu.Lock()
	max := maxConcurrent
	maxMu.Unlock()

	if max > 2 {
		t.Errorf("Expected max 2 concurrent tasks, got %d", max)
	}
}

func TestPool_QueueOverflowReturnsError(t *testing.T) {
	// Create pool with small queue to test overflow
	pool := transform.NewPool(1, 2)
	defer pool.Close()

	// Submit many blocking tasks rapidly
	// This should eventually hit the queue limit
	blocker := make(chan struct{})
	defer close(blocker)

	var gotError bool
	for range 10 {
		err := pool.Submit(func() error {
			<-blocker
			return nil
		})

		if err == transform.ErrPoolQueueFull {
			gotError = true
			break
		} else if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
	}

	if !gotError {
		t.Fatal("Expected ErrPoolQueueFull when submitting many tasks, but never received it")
	}
}

func TestPool_SubmitReturnsNilOnSuccess(t *testing.T) {
	pool := transform.NewPool(2, 10)
	defer pool.Close()

	// Submit should return nil when task is queued successfully
	err := pool.Submit(func() error {
		return errors.New("task error") // Error is not propagated
	})

	if err != nil {
		t.Errorf("Expected nil error from Submit, got %v", err)
	}
}

func TestPool_CloseStopsAcceptingTasks(t *testing.T) {
	pool := transform.NewPool(2, 10)
	pool.Close()

	err := pool.Submit(func() error {
		return nil
	})

	if err == nil {
		t.Fatal("Expected error when submitting to closed pool")
	}
	if err != transform.ErrPoolClosed {
		t.Errorf("Expected ErrPoolClosed, got %v", err)
	}
}

// TestPool_BackpressurePreventsBoundlessGoroutines tests that worker slots are acquired
// before spawning goroutines, preventing unbounded goroutine accumulation
func TestPool_BackpressurePreventsBoundlessGoroutines(t *testing.T) {
	// Create pool with 2 workers and 3 queue slots
	pool := transform.NewPool(2, 3)
	defer pool.Close()

	// Block to control when workers become available
	workerBlocker := make(chan struct{})

	// Track which tasks have started executing
	var executingCount int32
	executionBarrier := make(chan struct{})

	taskFn := func() error {
		// Signal that this task goroutine started
		atomic.AddInt32(&executingCount, 1)
		executionBarrier <- struct{}{}
		// Block until test completes
		<-workerBlocker
		return nil
	}

	// Submit first 2 tasks - should acquire workers immediately
	for i := range 2 {
		if err := pool.Submit(taskFn); err != nil {
			t.Fatalf("Task %d should succeed but got error: %v", i, err)
		}
	}

	// Wait for both workers to be acquired
	<-executionBarrier
	<-executionBarrier

	// Now workers are busy. Submit 3 more tasks to fill queue
	t.Log("Submitting 3 tasks to fill queue (queue depth=3)")
	for i := range 3 {
		if err := pool.Submit(taskFn); err != nil {
			t.Fatalf("Queue task %d should succeed but got error: %v", i, err)
		}
	}

	// Give dispatcher significant time to try processing queue
	// With old bug: dispatcher drains queue, spawns 3 more goroutines waiting for workers
	// With fix: dispatcher blocks on first queued task (can't acquire worker)
	time.Sleep(200 * time.Millisecond)

	// Check how many tasks are executing
	// Expected: Only 2 (the worker count)
	// Bug: 5 (all queued tasks drained and spawned as goroutines)
	count := atomic.LoadInt32(&executingCount)
	t.Logf("After sleep: %d tasks executing", count)

	if count > 2 {
		t.Errorf("Expected only 2 tasks executing (worker count), got %d - unbounded goroutines accumulated!", count)
		t.Logf("Bug confirmed: Dispatcher drained queue and spawned %d goroutines waiting for workers", count)
	}

	// Additional verification: at most one more task can be queued
	// The dispatcher pulls one task and blocks waiting for a worker, so:
	// - 2 workers busy (executing tasks 0, 1)
	// - 1 task held by dispatcher (task 2, waiting for worker)
	// - Queue has 2 tasks (tasks 3, 4)
	// - One more task can be added to queue (task 5)
	// - After that, queue is full (tasks 3, 4, 5)
	//
	// So: 6th task MAY succeed (going to last queue slot), but 7th MUST fail
	t.Log("Attempting to submit additional tasks")
	err := pool.Submit(taskFn)
	t.Logf("6th task result: %v", err)

	// 7th task MUST fail - queue completely full
	err2 := pool.Submit(taskFn)
	t.Logf("7th task result: %v", err2)
	if err2 != transform.ErrPoolQueueFull {
		t.Errorf("Expected 7th task to get ErrPoolQueueFull, got: %v", err2)
		t.Error("Bug: Queue was drained beyond dispatcher holding one task")
	}

	// Cleanup: unblock workers so test can complete
	close(workerBlocker)
}
