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
	"strings"
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
	fileEventTypes map[string]string // Track event type (create/delete/modify) for each file
	debounceTimer  *time.Timer
	mu             sync.Mutex
	logger         Logger
	done           chan struct{}
	wg             sync.WaitGroup
	ignorePatterns []string // Glob patterns to ignore
	watchDir       string   // Root watch directory for relative path resolution
}

// getDefaultIgnorePatterns returns the default list of glob patterns to ignore
func getDefaultIgnorePatterns() []string {
	return []string{
		".git",
		"node_modules",
		"dist",
		"build",
		"_site",
		".cache",
	}
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
		ignorePatterns: getDefaultIgnorePatterns(),
	}

	// Start event processing loop
	fw.wg.Add(1)
	go fw.processEvents()

	return fw, nil
}

// SetIgnorePatterns sets custom ignore patterns (completely overrides defaults)
func (fw *fileWatcher) SetIgnorePatterns(watchDir string, patterns []string) {
	fw.mu.Lock()
	defer fw.mu.Unlock()
	fw.watchDir = watchDir
	// If custom patterns provided, use them exclusively (override defaults)
	// If no patterns provided, defaults remain active
	if len(patterns) > 0 {
		fw.ignorePatterns = patterns
	}
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
		if fw.shouldIgnore(p) {
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
	// Stop the debounce timer first
	fw.mu.Lock()
	if fw.debounceTimer != nil {
		fw.debounceTimer.Stop()
	}
	fw.mu.Unlock()

	// Close the watcher to stop new events
	var err error
	if fw.watcher != nil {
		err = fw.watcher.Close()
	}

	// Signal processEvents goroutine to exit
	close(fw.done)

	// Wait for processEvents to exit
	fw.wg.Wait()

	// Close events channel after processEvents has exited
	close(fw.events)

	return err
}

// processEvents processes raw fsnotify events and applies debouncing
func (fw *fileWatcher) processEvents() {
	defer fw.wg.Done()
	for {
		select {
		case event, ok := <-fw.watcher.Events:
			if !ok {
				return
			}

			// Ignore files and directories based on configured patterns
			if fw.shouldIgnore(event.Name) {
				continue
			}

			// Determine event type
			var eventType string
			if event.Op&fsnotify.Create == fsnotify.Create {
				eventType = "create"
			} else if event.Op&fsnotify.Remove == fsnotify.Remove {
				eventType = "delete"
			} else if event.Op&fsnotify.Write == fsnotify.Write {
				eventType = "modify"
			} else if event.Op&fsnotify.Rename == fsnotify.Rename {
				eventType = "rename"
			} else {
				eventType = "modify" // Default
			}

			// Track file change with event type
			fw.mu.Lock()
			fw.debouncedFiles[event.Name] = time.Now()
			// Store event type for this file (we'll use the last one if multiple events)
			if fw.fileEventTypes == nil {
				fw.fileEventTypes = make(map[string]string)
			}
			fw.fileEventTypes[event.Name] = eventType

			// Reset debounce timer
			if fw.debounceTimer != nil {
				fw.debounceTimer.Stop()
			}

			fw.debounceTimer = time.AfterFunc(fw.debounceWindow, func() {
				fw.flushDebouncedEvents()
			})
			fw.mu.Unlock()

			if fw.logger != nil {
				fw.logger.Debug("File %s: %s", eventType, event.Name)
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

	// Check if watcher is closed
	select {
	case <-fw.done:
		return
	default:
	}

	if len(fw.debouncedFiles) == 0 {
		return
	}

	// Collect all changed files and analyze event types
	files := make([]string, 0, len(fw.debouncedFiles))
	var hasCreates, hasDeletes, hasPackageJSON bool
	var primaryEventType string

	for file := range fw.debouncedFiles {
		files = append(files, file)

		// Check event type for this file
		eventType := fw.fileEventTypes[file]
		ext := filepath.Ext(file)

		// Only count creates/deletes for source files (.ts/.js), not temp files
		isSourceFile := ext == ".ts" || ext == ".js"
		if isSourceFile && eventType == "create" {
			hasCreates = true
		} else if isSourceFile && eventType == "delete" {
			hasDeletes = true
		}

		// Check if package.json was modified
		if filepath.Base(file) == "package.json" {
			hasPackageJSON = true
		}

		// Use first file's event type as primary
		if primaryEventType == "" {
			primaryEventType = eventType
		}
	}

	// Clear the maps
	fw.debouncedFiles = make(map[string]time.Time)
	fw.fileEventTypes = make(map[string]string)

	// Send batched event with all files
	if len(files) > 0 {
		if primaryEventType == "" {
			primaryEventType = "modify"
		}

		event := FileEvent{
			Path:           files[0],
			Paths:          files,
			EventType:      primaryEventType,
			HasCreates:     hasCreates,
			HasDeletes:     hasDeletes,
			HasPackageJSON: hasPackageJSON,
			Timestamp:      time.Now(),
		}

		select {
		case fw.events <- event:
		case <-fw.done:
			// Watcher closed, don't send
			return
		default:
			// Channel full, drop event
			if fw.logger != nil {
				fw.logger.Warning("Dropped file event (channel full) - consider reducing change frequency")
			}
		}
	}
}

// shouldIgnore checks if a file path should be ignored based on patterns
func (fw *fileWatcher) shouldIgnore(path string) bool {
	base := filepath.Base(path)

	// Ignore editor temp files (always ignored regardless of patterns)
	// Vim/Neovim: .swp, .swo, .swn, ~
	// Emacs: #file#, .#file
	if strings.HasPrefix(base, ".") && (strings.HasSuffix(base, ".swp") ||
		strings.HasSuffix(base, ".swo") || strings.HasSuffix(base, ".swn")) {
		return true
	}
	if strings.HasSuffix(base, "~") ||
		(strings.HasPrefix(base, "#") && strings.HasSuffix(base, "#")) ||
		strings.HasPrefix(base, ".#") {
		return true
	}

	// Check against configured ignore patterns
	fw.mu.Lock()
	patterns := fw.ignorePatterns
	watchDir := fw.watchDir
	fw.mu.Unlock()

	// Convert path to relative if we have a watch directory
	relPath := path
	if watchDir != "" {
		if rel, err := filepath.Rel(watchDir, path); err == nil {
			relPath = rel
		}
	}

	// Check each pattern
	for _, pattern := range patterns {
		// Direct basename match (for simple patterns like "node_modules", "_site")
		if base == pattern {
			if fw.logger != nil {
				fw.logger.Debug("Ignoring %s (matches pattern: %s)", relPath, pattern)
			}
			return true
		}

		// Glob pattern match on relative path
		if matched, err := filepath.Match(pattern, relPath); err == nil && matched {
			if fw.logger != nil {
				fw.logger.Debug("Ignoring %s (matches glob pattern: %s)", relPath, pattern)
			}
			return true
		}

		// Glob pattern match with ** wildcard (match any subdirectory)
		// Convert glob pattern to prefix match for ** patterns
		if strings.Contains(pattern, "**") {
			// Simple ** handling: _site/** matches anything under _site/
			prefix := strings.TrimSuffix(pattern, "/**")
			if strings.HasPrefix(relPath, prefix+"/") || relPath == prefix {
				if fw.logger != nil {
					fw.logger.Debug("Ignoring %s (matches recursive pattern: %s)", relPath, pattern)
				}
				return true
			}
		}

		// Check if any parent directory matches the pattern
		dir := filepath.Dir(relPath)
		for dir != "." && dir != "/" {
			dirBase := filepath.Base(dir)
			if dirBase == pattern {
				if fw.logger != nil {
					fw.logger.Debug("Ignoring %s (parent directory %s matches pattern: %s)", relPath, dirBase, pattern)
				}
				return true
			}
			dir = filepath.Dir(dir)
		}
	}

	// Note: We don't filter numeric filenames here because that would incorrectly
	// ignore legitimate directories like "2025". Instead, manifest-based filtering
	// in handleFileChanges() filters out files that aren't in the source list.
	// Neovim's atomic write temp files will be filtered there since they won't
	// be in the manifest.

	return false
}
