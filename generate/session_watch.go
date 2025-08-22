/*
Copyright © 2025 Benny Powers

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
package generate

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"

	C "bennypowers.dev/cem/cmd/config"
	"bennypowers.dev/cem/internal/logging"
	M "bennypowers.dev/cem/manifest"
	W "bennypowers.dev/cem/workspace"
	DS "github.com/bmatcuk/doublestar"
	"github.com/fsnotify/fsnotify"
	"github.com/pterm/pterm"
)

// WatchSession manages the long-lived watch mode state
type WatchSession struct {
	ctx             W.WorkspaceContext
	globs           []string
	generateSession *GenerateSession
	debounceTimer   *time.Timer
	mu              sync.Mutex
	cancelCurrent   context.CancelFunc
	lastWrittenHash map[string][32]byte  // file path -> SHA256 hash of content we wrote
	lastWrittenTime map[string]time.Time // file path -> modification time when we wrote it
	pendingChanges  map[string]bool      // files that have changed and are pending processing
	// Demo discovery optimization
	demoFilesChanged bool                  // flag set when demo files change (event-driven)
	lastDemoConfig   C.DemoDiscoveryConfig // stored config copy for change detection
}

// NewWatchSession creates a new watch session with the given workspace context and globs
func NewWatchSession(ctx W.WorkspaceContext, globs []string) (*WatchSession, error) {
	generateSession, err := NewGenerateSession(ctx)
	if err != nil {
		return nil, fmt.Errorf("create generate session: %w", err)
	}

	return &WatchSession{
		ctx:             ctx,
		globs:           globs,
		generateSession: generateSession,
		lastWrittenHash: make(map[string][32]byte),
		lastWrittenTime: make(map[string]time.Time),
		pendingChanges:  make(map[string]bool),
	}, nil
}

// Close releases resources held by the watch session
func (ws *WatchSession) Close() {
	if ws.generateSession != nil {
		ws.generateSession.Close()
	}
}

// RunWatch starts the file watching mode with initial generation
func (ws *WatchSession) RunWatch() error {
	logging.Info("Starting watch mode...")

	// Do initial generation
	start := time.Now()

	if err := ws.generateOnce(context.Background()); err != nil {
		logging.Error("Initial generation failed: %v", err)
		return err
	}
	// Update demo discovery state after initial generation
	ws.updateDemoDiscoveryState()

	duration := time.Since(start)
	logging.Success("Generated in %s", ColorizeDuration(duration).Sprint(duration))

	// Set up file watcher
	watcher, err := ws.setupWatcher()
	if err != nil {
		return fmt.Errorf("setup file watcher: %w", err)
	}
	defer watcher.Close()

	logging.Info("Watching for file changes... (Ctrl+C to stop)")

	// Watch for changes
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return nil
			}
			ws.handleFileChange(event)
		case err, ok := <-watcher.Errors:
			if !ok {
				return nil
			}
			logging.Warning("File watcher error: %v", err)
		}
	}
}

// setupWatcher creates and configures the file system watcher
func (ws *WatchSession) setupWatcher() (*fsnotify.Watcher, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	// Get all directories that might contain files matching our globs
	dirs := make(map[string]bool)
	dirs[ws.ctx.Root()] = true // Always watch the root

	// Expand globs to get current files and their directories
	for _, glob := range ws.globs {
		files, err := ws.ctx.Glob(glob)
		if err != nil {
			continue
		}
		for _, file := range files {
			dir := filepath.Dir(file)
			if !filepath.IsAbs(dir) {
				dir = filepath.Join(ws.ctx.Root(), dir)
			}
			dirs[dir] = true
		}
	}

	// Add demo discovery directories if configured
	cfg, err := ws.ctx.Config()
	if err == nil && cfg.Generate.DemoDiscovery.FileGlob != "" {
		demoFiles, err := ws.ctx.Glob(cfg.Generate.DemoDiscovery.FileGlob)
		if err == nil {
			for _, file := range demoFiles {
				dir := filepath.Dir(file)
				if !filepath.IsAbs(dir) {
					dir = filepath.Join(ws.ctx.Root(), dir)
				}
				dirs[dir] = true
			}
		}
	}

	// Watch all identified directories
	for dir := range dirs {
		if err := watcher.Add(dir); err != nil {
			pterm.Warning.Printf("Failed to watch directory %s: %v\n", dir, err)
		}
	}

	return watcher, nil
}

// handleFileChange processes a file system event with debouncing
func (ws *WatchSession) handleFileChange(event fsnotify.Event) {
	// Only care about write and create events
	if event.Op&fsnotify.Write == 0 && event.Op&fsnotify.Create == 0 {
		return
	}

	// Ignore changes that we just made to prevent infinite loops
	if ws.isOurWrite(event.Name) {
		pterm.Debug.Printf("Ignoring our own write to: %s\n", event.Name)
		return
	}

	// Check if this is a demo file that affects demo discovery
	isDemoFile := ws.matchesDemoGlobs(event.Name)
	if isDemoFile {
		ws.mu.Lock()
		ws.demoFilesChanged = true
		ws.mu.Unlock()
		pterm.Debug.Printf("Demo file changed: %s\n", event.Name)
	}

	// Only process files that match our input globs or are demo files
	if !ws.matchesInputGlobs(event.Name) && !isDemoFile {
		pterm.Debug.Printf("Ignoring file not matching input or demo globs: %s\n", event.Name)
		return
	}

	pterm.Debug.Printf("File change detected: %s (op: %s)\n", event.Name, event.Op)

	ws.mu.Lock()
	defer ws.mu.Unlock()

	// Track this file as changed - convert FS path to module path
	modulePath, err := ws.ctx.FSPathToModule(event.Name)
	if err != nil {
		// If conversion fails, use the FS path as-is
		modulePath = event.Name
	}
	ws.pendingChanges[modulePath] = true

	// Cancel any existing timer
	if ws.debounceTimer != nil {
		ws.debounceTimer.Stop()
	}

	// Start a new debounce timer (100ms delay)
	// Later, we can make this configurable, but for now, it's ok
	// to hard code the value, for simplicity's sake
	debounceDelay := 100 * time.Millisecond
	ws.debounceTimer = time.AfterFunc(debounceDelay, func() {
		ws.processChanges()
	})
}

// isOurWrite checks if the file change was caused by our own write operation
func (ws *WatchSession) isOurWrite(filePath string) bool {
	cleanPath := filepath.Clean(filePath)

	// Fast path: Check filesystem stat first (much faster than hash computation)
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return false // If we can't stat it, assume it's not our write
	}

	ws.mu.Lock()
	lastWriteTime, timeExists := ws.lastWrittenTime[cleanPath]
	lastHash, hashExists := ws.lastWrittenHash[cleanPath]
	ws.mu.Unlock()

	// If we have both time and hash records, use time as fast check
	if timeExists && hashExists {
		// If file modification time matches when we wrote it, likely our write
		// Allow small tolerance for filesystem time precision issues
		timeDiff := fileInfo.ModTime().Sub(lastWriteTime)
		if timeDiff >= 0 && timeDiff < time.Second {
			return true
		}

		// If time doesn't match, it's definitely not our write
		if fileInfo.ModTime().After(lastWriteTime.Add(time.Second)) {
			return false
		}
	}

	// Fallback: Compute hash for definitive answer (expensive but reliable)
	file, err := os.Open(filePath)
	if err != nil {
		return false // If we can't read it, assume it's not our write
	}
	defer file.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return false
	}

	currentHash := [32]byte(hasher.Sum(nil))
	return hashExists && currentHash == lastHash
}

// matchesInputGlobs checks if a file path matches any of our input globs
func (ws *WatchSession) matchesInputGlobs(filePath string) bool {
	// Convert absolute path to relative path from workspace root
	relPath, err := filepath.Rel(ws.ctx.Root(), filePath)
	if err != nil {
		// If we can't get relative path, try with the original path
		relPath = filePath
	}

	// Check against each glob pattern
	for _, glob := range ws.globs {
		matched, err := filepath.Match(glob, relPath)
		if err == nil && matched {
			return true
		}

		// Also try with doublestar for more complex patterns like **/*.ts
		matched, err = DS.PathMatch(glob, relPath)
		if err == nil && matched {
			return true
		}
	}

	return false
}

// matchesDemoGlobs checks if a file path matches the demo discovery glob pattern
func (ws *WatchSession) matchesDemoGlobs(filePath string) bool {
	cfg, err := ws.ctx.Config()
	if err != nil || cfg.Generate.DemoDiscovery.FileGlob == "" {
		return false
	}

	// Convert absolute path to relative path from workspace root
	relPath, err := filepath.Rel(ws.ctx.Root(), filePath)
	if err != nil {
		// If we can't get relative path, try with the original path
		relPath = filePath
	}

	// Check against demo discovery glob pattern
	matched, err := filepath.Match(cfg.Generate.DemoDiscovery.FileGlob, relPath)
	if err == nil && matched {
		return true
	}

	// Also try with doublestar for more complex patterns like **/*.html
	matched, err = DS.PathMatch(cfg.Generate.DemoDiscovery.FileGlob, relPath)
	return err == nil && matched
}

// processChanges handles the actual regeneration after debouncing
func (ws *WatchSession) processChanges() {
	ws.mu.Lock()

	// Cancel any in-progress generation
	if ws.cancelCurrent != nil {
		ws.cancelCurrent()
	}

	// Create new cancellable context for this generation cycle
	ctx, cancel := context.WithCancel(context.Background())
	ws.cancelCurrent = cancel

	// Capture the list of changed files and reset pending changes
	changedFiles := make([]string, 0, len(ws.pendingChanges))
	for file := range ws.pendingChanges {
		changedFiles = append(changedFiles, file)
	}
	ws.pendingChanges = make(map[string]bool)

	ws.mu.Unlock()

	pterm.Info.Println("File changes detected, regenerating...")
	start := time.Now()

	// Try incremental processing first
	if len(changedFiles) > 0 {
		pterm.Debug.Printf("Changed files: %v\n", changedFiles)

		// Check if we can skip demo discovery for this generation cycle
		skipDemoDiscovery := ws.shouldSkipDemoDiscovery()
		if skipDemoDiscovery {
			pterm.Debug.Println("Skipping demo discovery - config and demo files unchanged")
		}

		pkg, err := ws.generateSession.ProcessChangedFilesWithSkip(ctx, changedFiles, skipDemoDiscovery)
		if err != nil {
			if errors.Is(err, context.Canceled) {
				pterm.Warning.Println("Generation cancelled due to new file changes")
			} else {
				pterm.Error.Printf("Generation failed: %v\n", err)
				// Also show individual warnings for dependency issues
				ws.printGenerationWarnings(err)
			}
			return
		}

		// Write the manifest
		manifestStr, err := M.SerializeToString(pkg)
		if err != nil {
			pterm.Error.Printf("Failed to serialize manifest: %v\n", err)
			return
		}

		if err := ws.writeManifest(manifestStr); err != nil {
			pterm.Error.Printf("Failed to write manifest: %v\n", err)
			return
		}

		// Update demo discovery state after successful generation
		if !skipDemoDiscovery {
			ws.updateDemoDiscoveryState()
		}
	} else {
		// Fallback to full rebuild if no specific files tracked
		if err := ws.generateOnce(ctx); err != nil {
			if errors.Is(err, context.Canceled) {
				pterm.Warning.Println("Generation cancelled due to new file changes")
			} else {
				pterm.Error.Printf("Generation failed: %v\n", err)
				// Also show individual warnings for dependency issues
				ws.printGenerationWarnings(err)
			}
			return
		}
		// Update demo discovery state after full rebuild
		ws.updateDemoDiscoveryState()
	}

	duration := time.Since(start)
	pterm.Success.Printf("Regenerated in %s\n", ColorizeDuration(duration).Sprint(duration))
}

// generateOnce performs a single generation cycle, respecting cancellation
func (ws *WatchSession) generateOnce(ctx context.Context) error {
	// Use the reusable generate session for efficient regeneration
	pkg, err := ws.generateSession.GenerateFullManifest(ctx)
	if err != nil {
		return err
	}

	// Serialize to string
	manifestStr, err := M.SerializeToString(pkg)
	if err != nil {
		return fmt.Errorf("serialize manifest: %w", err)
	}

	// Write the manifest
	return ws.writeManifest(manifestStr)
}

// writeManifest writes the manifest to the configured output path
func (ws *WatchSession) writeManifest(manifestStr string) error {
	cfg, err := ws.ctx.Config()
	if err != nil {
		return err
	}

	// Determine output path (copied from main generate logic)
	outputPath := cfg.Generate.Output
	if outputPath == "" {
		pkgjson, err := ws.ctx.PackageJSON()
		if err != nil {
			return err
		}
		if pkgjson != nil && pkgjson.CustomElements != "" {
			outputPath = filepath.Join(ws.ctx.Root(), pkgjson.CustomElements)
		}
	}

	if outputPath == "" {
		// If no output path, just skip writing (or could print to stdout)
		return nil
	}

	// Prepare the content we're about to write
	content := manifestStr + "\n"
	contentBytes := []byte(content)

	writer, err := ws.ctx.OutputWriter(outputPath)
	if err != nil {
		return err
	}

	_, err = writer.Write(contentBytes)
	if err != nil {
		writer.Close()
		return err
	}

	// Close the writer first to ensure all data is flushed
	err = writer.Close()
	if err != nil {
		return err
	}

	// Record the hash and modification time of what we just wrote to prevent infinite loops
	// We do this AFTER closing to ensure the file is fully written
	hash := sha256.Sum256(contentBytes)

	// Get the file's modification time after writing
	fileInfo, err := os.Stat(outputPath)
	var modTime time.Time
	if err == nil {
		modTime = fileInfo.ModTime()
	} else {
		// Fallback to current time if we can't stat the file
		modTime = time.Now()
	}

	cleanPath := filepath.Clean(outputPath)
	ws.mu.Lock()
	ws.lastWrittenHash[cleanPath] = hash
	ws.lastWrittenTime[cleanPath] = modTime
	ws.mu.Unlock()

	return nil
}

// shouldSkipDemoDiscovery checks if demo discovery can be skipped
func (ws *WatchSession) shouldSkipDemoDiscovery() bool {
	cfg, err := ws.ctx.Config()
	if err != nil || cfg.Generate.DemoDiscovery.FileGlob == "" {
		return true // Skip if no demo discovery configured
	}

	ws.mu.Lock()
	// Make copies to avoid holding lock during comparison
	lastConfig := ws.lastDemoConfig
	demoFilesChanged := ws.demoFilesChanged
	// Check if this is the first run
	isFirstRun := ws.lastDemoConfig == (C.DemoDiscoveryConfig{})
	ws.mu.Unlock()

	if isFirstRun {
		return false // Need to run discovery on first run
	}

	// Safe comparison outside the lock (both are value types)
	configChanged := lastConfig != cfg.Generate.DemoDiscovery

	// Skip only if: no config changes AND no demo file changes
	return !configChanged && !demoFilesChanged
}

// updateDemoDiscoveryState updates the tracked state after demo discovery
func (ws *WatchSession) updateDemoDiscoveryState() {
	cfg, err := ws.ctx.Config()
	if err != nil {
		return
	}

	ws.mu.Lock()
	// Store value copy, not pointer - safe for concurrent access
	ws.lastDemoConfig = cfg.Generate.DemoDiscovery
	ws.demoFilesChanged = false // Reset flag
	ws.mu.Unlock()
}

// printGenerationWarnings displays errors as warnings, showing specific error details for watch mode
func (ws *WatchSession) printGenerationWarnings(err error) {
	if err == nil {
		return
	}

	// Split the error into individual components for better display
	errList := ws.flattenErrors(err)
	for _, e := range errList {
		pterm.Warning.Printf("Warning: %v\n", e)
	}
}

// flattenErrors recursively flattens joined errors into a slice
func (ws *WatchSession) flattenErrors(err error) []error {
	if err == nil {
		return nil
	}

	// Try to unwrap as joined error
	if joinedErr, ok := err.(interface{ Unwrap() []error }); ok {
		var result []error
		for _, e := range joinedErr.Unwrap() {
			result = append(result, ws.flattenErrors(e)...)
		}
		return result
	}

	return []error{err}
}
