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

package transform

import (
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestPool_LimitsConcurrentTasks(t *testing.T) {
	// Create pool with max 2 concurrent tasks
	pool := NewPool(2, 10)
	defer pool.Close()

	// Track concurrent executions
	var concurrent int32
	var maxConcurrent int32

	// Function that tracks concurrency
	task := func() error {
		current := atomic.AddInt32(&concurrent, 1)
		defer atomic.AddInt32(&concurrent, -1)

		// Update max if needed
		for {
			max := atomic.LoadInt32(&maxConcurrent)
			if current <= max || atomic.CompareAndSwapInt32(&maxConcurrent, max, current) {
				break
			}
		}

		// Simulate work
		time.Sleep(50 * time.Millisecond)
		return nil
	}

	// Submit 10 tasks
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = pool.Submit(task)
		}()
	}

	wg.Wait()

	// Verify max concurrent was respected
	max := atomic.LoadInt32(&maxConcurrent)
	if max > 2 {
		t.Errorf("Expected max 2 concurrent tasks, got %d", max)
	}
}

func TestPool_QueueOverflowReturnsError(t *testing.T) {
	// Create pool with small queue to test overflow
	pool := NewPool(1, 2)
	defer pool.Close()

	// Submit many blocking tasks rapidly
	// This should eventually hit the queue limit
	blocker := make(chan struct{})
	defer close(blocker)

	var gotError bool
	for i := 0; i < 10; i++ {
		err := pool.Submit(func() error {
			<-blocker
			return nil
		})

		if err == ErrPoolQueueFull {
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
	pool := NewPool(2, 10)
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
	pool := NewPool(2, 10)
	pool.Close()

	err := pool.Submit(func() error {
		return nil
	})

	if err == nil {
		t.Fatal("Expected error when submitting to closed pool")
	}
	if err != ErrPoolClosed {
		t.Errorf("Expected ErrPoolClosed, got %v", err)
	}
}
