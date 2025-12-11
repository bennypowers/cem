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
)

var (
	// ErrPoolQueueFull is returned when the work queue is at capacity
	ErrPoolQueueFull = errors.New("transform pool queue is full")
	// ErrPoolClosed is returned when attempting to submit to a closed pool
	ErrPoolClosed = errors.New("transform pool is closed")
)

// Pool manages concurrent transform operations with backpressure
type Pool struct {
	workers   chan struct{} // Semaphore for limiting concurrent workers
	queue     chan task     // Task queue
	closeOnce sync.Once
	closed    chan struct{}
}

// task represents a unit of work
type task struct {
	fn func() error
}

// NewPool creates a worker pool with specified concurrency and queue depth
// maxWorkers: maximum number of concurrent transform operations
// queueDepth: maximum number of queued tasks before rejecting new submissions
func NewPool(maxWorkers, queueDepth int) *Pool {
	p := &Pool{
		workers: make(chan struct{}, maxWorkers),
		queue:   make(chan task, queueDepth),
		closed:  make(chan struct{}),
	}

	// Start worker dispatcher
	go p.dispatch()

	return p
}

// dispatch processes tasks from the queue
func (p *Pool) dispatch() {
	for {
		select {
		case <-p.closed:
			return
		case t := <-p.queue:
			// Acquire worker slot BEFORE spawning goroutine
			// This blocks the dispatcher when all workers are busy,
			// keeping tasks in the queue and enforcing backpressure
			select {
			case p.workers <- struct{}{}:
				// Worker acquired - now spawn goroutine to execute task
				go func(t task) {
					defer func() {
						// Release worker slot
						<-p.workers
					}()

					// Execute task (errors are handled by caller)
					_ = t.fn()
				}(t)
			case <-p.closed:
				// Pool closed while waiting for worker
				return
			}
		}
	}
}

// Submit adds a task to the pool (non-blocking)
// Returns an error if the queue is full or the pool is closed
// Task execution errors are not returned (handle within task function)
func (p *Pool) Submit(fn func() error) error {
	select {
	case <-p.closed:
		return ErrPoolClosed
	default:
	}

	// Create task
	t := task{fn: fn}

	// Try to enqueue (non-blocking)
	select {
	case p.queue <- t:
		// Task queued successfully
		return nil
	case <-p.closed:
		return ErrPoolClosed
	default:
		// Queue full
		return ErrPoolQueueFull
	}
}

// SubmitSync adds a task to the pool and waits for completion
// Returns the result of the task function or a pool error
// This is a synchronous wrapper around Submit for HTTP handler use cases
func (p *Pool) SubmitSync(fn func() error) error {
	// Create channel to receive result
	done := make(chan error, 1)

	// Wrap task to signal completion
	wrappedFn := func() error {
		result := fn()
		done <- result
		return result
	}

	// Submit task (non-blocking queue check)
	if err := p.Submit(wrappedFn); err != nil {
		return err // ErrPoolQueueFull or ErrPoolClosed
	}

	// Wait for task completion
	return <-done
}

// Close stops the pool from accepting new tasks
func (p *Pool) Close() {
	p.closeOnce.Do(func() {
		close(p.closed)
	})
}
