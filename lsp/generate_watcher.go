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
package lsp

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	G "bennypowers.dev/cem/generate"
	"bennypowers.dev/cem/lsp/helpers"
	M "bennypowers.dev/cem/manifest"
	"bennypowers.dev/cem/types"
	"github.com/fsnotify/fsnotify"
)

// ManifestUpdateCallback is called when a new manifest is generated
type ManifestUpdateCallback func(pkg *M.Package) error

// InProcessGenerateWatcher manages in-process generate watching for the LSP server
// This replaces the subprocess-based approach to prevent hanging processes
// It generates manifests in-memory only, never writing to files
type InProcessGenerateWatcher struct {
	generateSession *G.GenerateSession
	ctx             context.Context
	cancel          context.CancelFunc
	mu              sync.RWMutex
	done            chan struct{} // Signal to stop file watching
	wg              sync.WaitGroup // Wait for goroutine to exit
	workspace       types.WorkspaceContext
	globs           []string
	callback        ManifestUpdateCallback
	debounceTimer   *time.Timer  // Track debounce timer for cleanup
}

// NewInProcessGenerateWatcher creates a new in-process generate watcher
func NewInProcessGenerateWatcher(
	workspace types.WorkspaceContext,
	globs []string,
	callback ManifestUpdateCallback,
) (*InProcessGenerateWatcher, error) {
	// Create cancellable context for clean shutdown
	ctx, cancel := context.WithCancel(context.Background())

	helpers.SafeDebugLog("Creating InProcessGenerateWatcher for workspace: %s, globs: %v", workspace.Root(), globs)

	// Create the generate session (NOT watch session - we manage watching ourselves)
	generateSession, err := G.NewGenerateSession(workspace)
	if err != nil {
		cancel()
		helpers.SafeDebugLog("Failed to create generate session: %v", err)
		return nil, fmt.Errorf("failed to create generate session: %w", err)
	}

	helpers.SafeDebugLog("Successfully created generate session")

	return &InProcessGenerateWatcher{
		generateSession: generateSession,
		ctx:             ctx,
		cancel:          cancel,
		workspace:       workspace,
		globs:           globs,
		callback:        callback,
	}, nil
}

// Start begins the file watching process in a goroutine
// This will watch for source file changes and regenerate the manifest automatically
// NOTE: This generates manifests in-memory only and calls the callback, never writes files
// Thread-safety: All operations in this function are protected by mu
func (w *InProcessGenerateWatcher) Start() error {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.generateSession == nil {
		return fmt.Errorf("watcher not initialized")
	}

	// Prevent starting the watcher multiple times
	if w.done != nil {
		return fmt.Errorf("generate watcher already running")
	}

	// PROTECTED BY LOCK: Initialize done channel for shutdown signaling
	w.done = make(chan struct{})

	// Do initial generation and call callback
	if err := w.generateAndCallback(); err != nil {
		helpers.SafeDebugLog("Initial generation failed: %v", err)
		return err
	}

	// Start watching in a goroutine to avoid blocking the LSP server
	w.wg.Add(1)
	go func() {
		defer w.wg.Done()
		helpers.SafeDebugLog("Starting in-process generate watcher for workspace: %s", w.workspace.Root())

		// Sleep for grace period BEFORE setting up fsnotify watcher
		// This avoids the initial event flood entirely
		gracePeriod := 2 * time.Second
		helpers.SafeDebugLog("Waiting %v before starting file watcher to avoid initial event flood", gracePeriod)

		select {
		case <-time.After(gracePeriod):
			// Grace period elapsed, proceed with watching
		case <-w.done:
			// Shutdown requested during grace period
			helpers.SafeDebugLog("Generate watcher stopped during grace period")
			return
		case <-w.ctx.Done():
			// Context cancelled during grace period
			helpers.SafeDebugLog("Generate watcher cancelled during grace period")
			return
		}

		helpers.SafeDebugLog("Grace period elapsed, starting file watcher")

		// Set up file watcher
		if err := w.watchFiles(); err != nil {
			helpers.SafeDebugLog("File watching failed: %v", err)
		}
		helpers.SafeDebugLog("Generate watcher goroutine exiting")
	}()

	return nil
}

// watchFiles sets up file watching for source files and calls generateAndCallback on changes
func (w *InProcessGenerateWatcher) watchFiles() error {
	// Get done channel reference under lock to avoid race condition
	w.mu.RLock()
	done := w.done
	w.mu.RUnlock()

	// Create file watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return fmt.Errorf("failed to create file watcher: %w", err)
	}
	defer func() { _ = watcher.Close() }()

	// Walk the workspace to find files, filtering out node_modules
	var filesToWatch []string
	rootDir := w.workspace.Root()

	err = filepath.WalkDir(rootDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return nil // Skip errors
		}

		// ALWAYS skip node_modules
		if d.IsDir() && d.Name() == "node_modules" {
			return filepath.SkipDir
		}

		// Skip hidden directories
		if d.IsDir() && strings.HasPrefix(d.Name(), ".") {
			return filepath.SkipDir
		}

		// Only watch .ts and .js files
		if !d.IsDir() && (strings.HasSuffix(path, ".ts") || strings.HasSuffix(path, ".js")) {
			// Check if this file matches our globs
			if w.shouldProcessFile(path) {
				filesToWatch = append(filesToWatch, path)
			}
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to walk workspace: %w", err)
	}

	// Fail early if there are too many files to watch
	const maxFiles = 1000
	if len(filesToWatch) > maxFiles {
		return fmt.Errorf("refusing to watch %d files (max %d) - adjust your globs or use `cem generate --watch` separately", len(filesToWatch), maxFiles)
	}

	// Watch individual files (not directories) to avoid event floods
	helpers.SafeDebugLog("Watching %d source files for changes", len(filesToWatch))
	for _, file := range filesToWatch {
		if err := watcher.Add(file); err != nil {
			helpers.SafeDebugLog("Failed to watch file %s: %v", file, err)
		}
	}

	// Debouncing to avoid excessive regeneration
	debounceDuration := 500 * time.Millisecond

	// Watch for file changes
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return nil
			}

			// Check if this file matches our globs
			if w.shouldProcessFile(event.Name) {
				helpers.SafeDebugLog("File change detected: %s", event.Name)

				// Reset and store debounce timer
				w.mu.Lock()
				if w.debounceTimer != nil {
					w.debounceTimer.Stop()
				}
				w.debounceTimer = time.AfterFunc(debounceDuration, func() {
					if err := w.generateFullAndCallback(); err != nil {
						helpers.SafeDebugLog("Failed to regenerate manifest: %v", err)
					}
				})
				w.mu.Unlock()
			}

		case err, ok := <-watcher.Errors:
			if !ok {
				return nil
			}
			helpers.SafeDebugLog("File watcher error: %v", err)

		case <-done:
			helpers.SafeDebugLog("Generate watcher stopped due to shutdown signal")
			return nil

		case <-w.ctx.Done():
			helpers.SafeDebugLog("Generate watcher stopped due to context cancellation")
			return nil
		}
	}
}

// shouldProcessFile checks if a file matches our globs
func (w *InProcessGenerateWatcher) shouldProcessFile(filename string) bool {
	// Convert to relative path
	relPath, err := filepath.Rel(w.workspace.Root(), filename)
	if err != nil {
		return false
	}

	// Check against globs
	for _, glob := range w.globs {
		// Simple glob matching - check if file has .ts or .js extension and is in the right directory
		if strings.HasSuffix(relPath, ".ts") || strings.HasSuffix(relPath, ".js") {
			// Basic check if path matches glob pattern
			if strings.Contains(glob, filepath.Dir(relPath)) || strings.Contains(relPath, strings.TrimSuffix(glob, "*.ts")) {
				return true
			}
		}
	}
	return false
}

// generateAndCallback generates the manifest in-memory and calls the callback
// For LSP startup performance, this now generates lazily instead of processing all workspace files
func (w *InProcessGenerateWatcher) generateAndCallback() error {
	helpers.SafeDebugLog("Generating manifest in-memory for LSP...")

	// For LSP performance, we now generate an empty manifest initially
	// and only do full generation when files actually change.
	// This eliminates the 5+ minute startup delay from scanning thousands of files.

	// Try to get existing in-memory manifest from the session
	var pkg *M.Package
	if w.generateSession != nil {
		pkg = w.generateSession.InMemoryManifest()
	}

	// If no existing manifest, create minimal empty manifest instead of full workspace scan
	if pkg == nil {
		helpers.SafeDebugLog("Creating minimal empty manifest for fast LSP startup")
		// Create empty package - this will be populated later when files change
		emptyPkg := M.NewPackage([]M.Module{})
		pkg = &emptyPkg
		helpers.SafeDebugLog("Created empty manifest for lazy initialization")
	} else {
		helpers.SafeDebugLog("Using existing in-memory manifest with %d modules", len(pkg.Modules))
	}

	// Call the callback with the manifest (empty or existing)
	if w.callback != nil {
		if err := w.callback(pkg); err != nil {
			return fmt.Errorf("manifest callback failed: %w", err)
		}
	}

	helpers.SafeDebugLog("Successfully provided manifest to LSP (lazy mode)")
	return nil
}

// generateFullAndCallback performs full manifest generation when files actually change
func (w *InProcessGenerateWatcher) generateFullAndCallback() error {
	helpers.SafeDebugLog("Performing full manifest generation due to file changes...")

	// Create a fresh GenerateSession to avoid any caching issues
	// that might prevent updated file content from being parsed correctly
	freshSession, err := G.NewGenerateSession(w.workspace)
	if err != nil {
		return fmt.Errorf("create fresh generate session: %w", err)
	}
	defer freshSession.Close()

	// Generate manifest in-memory with fresh session (no file writes)
	pkg, err := freshSession.GenerateFullManifest(w.ctx)
	if err != nil {
		return fmt.Errorf("generate manifest: %w", err)
	}

	// Call the callback with the in-memory manifest
	if w.callback != nil {
		if err := w.callback(pkg); err != nil {
			return fmt.Errorf("manifest callback failed: %w", err)
		}
	}

	helpers.SafeDebugLog("Successfully regenerated full manifest due to file changes")
	return nil
}

// Stop cleanly shuts down the file watcher
// Thread-safety: All operations in this function are protected by mu
func (w *InProcessGenerateWatcher) Stop() error {
	w.mu.Lock()
	defer w.mu.Unlock()

	helpers.SafeDebugLog("Stopping in-process generate watcher")

	// PROTECTED BY LOCK: Stop the debounce timer to prevent it from firing after shutdown
	if w.debounceTimer != nil {
		w.debounceTimer.Stop()
		w.debounceTimer = nil
	}

	// PROTECTED BY LOCK: Signal the watchFiles goroutine to stop
	// Closing the channel is thread-safe and will cause watchFiles to exit
	if w.done != nil {
		close(w.done)
		// Do NOT set w.done = nil - this races with goroutine reading it at line 125
		// Closing is sufficient for shutdown signaling
	}

	// Cancel the context to signal shutdown
	if w.cancel != nil {
		w.cancel()
	}

	// Release the lock before waiting to avoid deadlock
	w.mu.Unlock()

	// Wait for goroutine to exit
	w.wg.Wait()
	helpers.SafeDebugLog("Generate watcher goroutine has exited")

	// Reacquire lock for final cleanup
	w.mu.Lock()

	// Now it's safe to set done to nil - goroutine has exited
	w.done = nil

	// Close the generate session
	if w.generateSession != nil {
		w.generateSession.Close()
		w.generateSession = nil
	}

	return nil
}

// IsRunning returns true if the watcher is currently active
func (w *InProcessGenerateWatcher) IsRunning() bool {
	w.mu.RLock()
	defer w.mu.RUnlock()

	// Check if context is still active and session exists
	return w.ctx.Err() == nil && w.generateSession != nil
}
