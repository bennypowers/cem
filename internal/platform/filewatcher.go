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
package platform

import (
	"fmt"
	"sync"

	"github.com/fsnotify/fsnotify"
)

// FileWatcher provides an abstraction over file watching operations.
// This interface enables:
// - Testing with mock file watchers (instant callbacks)
// - Platform-specific implementations
// - Graceful degradation when file watching is unavailable
type FileWatcher interface {
	// Add starts watching the named file or directory
	Add(name string) error

	// Remove stops watching the named file or directory
	Remove(name string) error

	// Close stops the watcher and releases resources
	Close() error

	// Events returns a channel of file system events
	Events() <-chan FileWatchEvent

	// Errors returns a channel of errors
	Errors() <-chan error
}

// FileWatchEvent represents a file system event
type FileWatchEvent struct {
	Name string  // File path
	Op   WatchOp // Operation type
}

// WatchOp describes a set of file operations
type WatchOp uint32

const (
	Create WatchOp = 1 << iota
	Write
	Remove
	Rename
	Chmod
)

func (op WatchOp) String() string {
	var names []string
	if op&Create != 0 {
		names = append(names, "CREATE")
	}
	if op&Write != 0 {
		names = append(names, "WRITE")
	}
	if op&Remove != 0 {
		names = append(names, "REMOVE")
	}
	if op&Rename != 0 {
		names = append(names, "RENAME")
	}
	if op&Chmod != 0 {
		names = append(names, "CHMOD")
	}
	if len(names) == 0 {
		return ""
	}
	return names[0]
}

// FSNotifyFileWatcher implements FileWatcher using fsnotify.
// This is the production implementation.
type FSNotifyFileWatcher struct {
	watcher *fsnotify.Watcher
	events  chan FileWatchEvent
	errors  chan error
	mu      sync.RWMutex
	closed  bool
	done    chan struct{}  // Signal to stop translateEvents goroutine
	wg      sync.WaitGroup // Wait for goroutine to exit
}

// NewFSNotifyFileWatcher creates a new file watcher using fsnotify.
func NewFSNotifyFileWatcher() (*FSNotifyFileWatcher, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, fmt.Errorf("failed to create fsnotify watcher: %w", err)
	}

	fw := &FSNotifyFileWatcher{
		watcher: watcher,
		events:  make(chan FileWatchEvent, 100),
		errors:  make(chan error, 10),
		done:    make(chan struct{}),
	}

	// Start event translation goroutine
	fw.wg.Add(1)
	go func() {
		defer fw.wg.Done()
		fw.translateEvents()
	}()

	return fw, nil
}

func (fw *FSNotifyFileWatcher) Add(name string) error {
	fw.mu.RLock()
	defer fw.mu.RUnlock()

	if fw.closed {
		return fmt.Errorf("file watcher is closed")
	}

	return fw.watcher.Add(name)
}

func (fw *FSNotifyFileWatcher) Remove(name string) error {
	fw.mu.RLock()
	defer fw.mu.RUnlock()

	if fw.closed {
		return fmt.Errorf("file watcher is closed")
	}

	return fw.watcher.Remove(name)
}

func (fw *FSNotifyFileWatcher) Close() error {
	fw.mu.Lock()

	if fw.closed {
		fw.mu.Unlock()
		return nil
	}

	fw.closed = true

	// Signal the translateEvents goroutine to stop
	close(fw.done)

	fw.mu.Unlock()

	// Wait for the goroutine to exit
	fw.wg.Wait()

	// Now close the watcher and channels
	err := fw.watcher.Close()
	close(fw.events)
	close(fw.errors)

	return err
}

func (fw *FSNotifyFileWatcher) Events() <-chan FileWatchEvent {
	return fw.events
}

func (fw *FSNotifyFileWatcher) Errors() <-chan error {
	return fw.errors
}

// translateEvents converts fsnotify events to our abstracted events
func (fw *FSNotifyFileWatcher) translateEvents() {
	// Add an initial blocking operation to prevent CPU spinning
	// Wait a moment for the watcher to be fully initialized
	select {
	case <-fw.done:
		return
	default:
	}

	for {
		select {
		case event, ok := <-fw.watcher.Events:
			if !ok {
				return
			}

			// Convert fsnotify.Op to our WatchOp
			var op WatchOp
			if event.Op&fsnotify.Create != 0 {
				op |= Create
			}
			if event.Op&fsnotify.Write != 0 {
				op |= Write
			}
			if event.Op&fsnotify.Remove != 0 {
				op |= Remove
			}
			if event.Op&fsnotify.Rename != 0 {
				op |= Rename
			}
			if event.Op&fsnotify.Chmod != 0 {
				op |= Chmod
			}

			fw.mu.RLock()
			if !fw.closed {
				select {
				case fw.events <- FileWatchEvent{
					Name: event.Name,
					Op:   op,
				}:
				case <-fw.done:
					fw.mu.RUnlock()
					return
				}
			}
			fw.mu.RUnlock()

		case err, ok := <-fw.watcher.Errors:
			if !ok {
				return
			}

			fw.mu.RLock()
			if !fw.closed {
				select {
				case fw.errors <- err:
				case <-fw.done:
					fw.mu.RUnlock()
					return
				}
			}
			fw.mu.RUnlock()

		case <-fw.done:
			// Shutdown signal received
			return
		}
	}
}

// GenerateWatcher provides an abstraction over generate watching operations.
// This interface enables:
// - Testing with mock generate watchers (instant callbacks)
// - In-process vs subprocess implementations
// - Graceful handling of generate watcher unavailability
type GenerateWatcher interface {
	// Start begins watching for source file changes and generating manifests
	Start() error
	// Stop ceases watching and cleans up resources
	Stop() error
	// IsRunning returns whether the watcher is currently active
	IsRunning() bool
}
