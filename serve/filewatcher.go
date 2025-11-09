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

package serve

import (
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
)

// fileWatcher implements FileWatcher interface
type fileWatcher struct {
	watcher        *fsnotify.Watcher
	events         chan FileEvent
	debounceWindow time.Duration
	debouncedFiles map[string]time.Time
	debounceTimer  *time.Timer
	mu             sync.Mutex
	logger         Logger
	done           chan struct{}
}

// newFileWatcher creates a new file watcher with debouncing
func newFileWatcher(debounceWindow time.Duration, logger Logger) (FileWatcher, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	fw := &fileWatcher{
		watcher:        watcher,
		events:         make(chan FileEvent, 100),
		debounceWindow: debounceWindow,
		debouncedFiles: make(map[string]time.Time),
		logger:         logger,
		done:           make(chan struct{}),
	}

	// Start event processing loop
	go fw.processEvents()

	return fw, nil
}

// Watch adds a path to watch (recursively if directory)
func (fw *fileWatcher) Watch(path string) error {
	// Add the root path
	if err := fw.watcher.Add(path); err != nil {
		return err
	}

	// Walk subdirectories and add them to the watcher
	return filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip non-directories
		if !info.IsDir() {
			return nil
		}

		// Skip if it's the root path (already added)
		if p == path {
			return nil
		}

		// Skip ignored directories
		name := filepath.Base(p)
		if shouldIgnore(name) {
			return filepath.SkipDir
		}

		// Add this directory to the watcher
		if err := fw.watcher.Add(p); err != nil {
			return err
		}

		return nil
	})
}

// Events returns the channel for receiving debounced file events
func (fw *fileWatcher) Events() <-chan FileEvent {
	return fw.events
}

// Close stops the file watcher
func (fw *fileWatcher) Close() error {
	close(fw.done)
	if fw.watcher != nil {
		return fw.watcher.Close()
	}
	return nil
}

// processEvents processes raw fsnotify events and applies debouncing
func (fw *fileWatcher) processEvents() {
	for {
		select {
		case event, ok := <-fw.watcher.Events:
			if !ok {
				return
			}

			// Ignore .git and node_modules directories
			if shouldIgnore(event.Name) {
				continue
			}

			// Track file change
			fw.mu.Lock()
			fw.debouncedFiles[event.Name] = time.Now()

			// Reset debounce timer
			if fw.debounceTimer != nil {
				fw.debounceTimer.Stop()
			}

			fw.debounceTimer = time.AfterFunc(fw.debounceWindow, func() {
				fw.flushDebouncedEvents()
			})
			fw.mu.Unlock()

			if fw.logger != nil {
				fw.logger.Debug("File changed: %s", event.Name)
			}

		case err, ok := <-fw.watcher.Errors:
			if !ok {
				return
			}
			if fw.logger != nil {
				fw.logger.Error("File watcher error: %v", err)
			}

		case <-fw.done:
			return
		}
	}
}

// flushDebouncedEvents sends accumulated file changes as a single event
func (fw *fileWatcher) flushDebouncedEvents() {
	fw.mu.Lock()
	defer fw.mu.Unlock()

	if len(fw.debouncedFiles) == 0 {
		return
	}

	// Collect all changed files
	files := make([]string, 0, len(fw.debouncedFiles))
	for file := range fw.debouncedFiles {
		files = append(files, file)
	}

	// Clear the map
	fw.debouncedFiles = make(map[string]time.Time)

	// Send batched event
	// For now, just send the first file as a representative
	// (in a real implementation, we'd pass all files)
	if len(files) > 0 {
		event := FileEvent{
			Path:      files[0],
			EventType: "modified",
			Timestamp: time.Now(),
		}

		select {
		case fw.events <- event:
		default:
			// Channel full, drop event
			if fw.logger != nil {
				fw.logger.Debug("Dropped file event (channel full)")
			}
		}
	}

	if fw.logger != nil {
		fw.logger.Info("File changes detected: %d files", len(files))
	}
}

// shouldIgnore checks if a file path should be ignored
func shouldIgnore(path string) bool {
	base := filepath.Base(path)

	// Ignore hidden files and common directories
	ignoredDirs := []string{".git", "node_modules", "dist", "build", ".cache"}
	for _, dir := range ignoredDirs {
		if base == dir {
			return true
		}
	}

	return false
}
