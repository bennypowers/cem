//go:build cemdev

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
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"bennypowers.dev/cem/serve/internal/elements"
	"bennypowers.dev/cem/serve/middleware/routes"
	"github.com/fsnotify/fsnotify"
)

// readInternalModule reads an internal module from disk during development
// This allows live-reload of chrome UI assets without rebuilding
func readInternalModule(path string) ([]byte, error) {
	// Get the source directory relative to this file
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return nil, fmt.Errorf("failed to get caller information")
	}
	serveDir := filepath.Dir(filename)
	templatesDir := filepath.Join(serveDir, "middleware", "routes")

	fullPath := filepath.Join(templatesDir, path)
	return os.ReadFile(fullPath)
}

// setupDevWatcher sets up a file watcher for element TypeScript files
// When changes are detected, it transpiles them and triggers a reload
func setupDevWatcher(server *Server) error {
	// Override the ReadInternalModule function to read from disk
	routes.ReadInternalModule = readInternalModule

	server.logger.Info("Dev mode: reading internal modules from disk")

	// Get the source directory
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return fmt.Errorf("failed to get caller information")
	}
	serveDir := filepath.Dir(filename)
	elementsDir := filepath.Join(serveDir, "middleware", "routes", "templates", "elements")

	// Check if elements directory exists
	if _, err := os.Stat(elementsDir); os.IsNotExist(err) {
		server.logger.Debug("Elements directory does not exist, skipping dev watcher: %s", elementsDir)
		return nil
	}

	// Create fsnotify watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return fmt.Errorf("creating fsnotify watcher: %w", err)
	}

	// Walk the elements directory and watch all subdirectories
	err = filepath.Walk(elementsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			if err := watcher.Add(path); err != nil {
				return fmt.Errorf("watching directory %s: %w", path, err)
			}
			server.logger.Debug("Watching for changes: %s", path)
		}
		return nil
	})
	if err != nil {
		_ = watcher.Close()
		return fmt.Errorf("walking elements directory: %w", err)
	}

	server.logger.Info("Dev mode: watching elements directory for changes")

	// Start the watcher goroutine
	go func() {
		defer watcher.Close()

		// Debounce timer to avoid rebuilding on every file save
		var debounceTimer *time.Timer
		const debounceDelay = 300 * time.Millisecond

		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}

				// Only react to TypeScript file changes
				if !strings.HasSuffix(event.Name, ".ts") {
					continue
				}

				// Ignore test files
				if strings.HasSuffix(event.Name, ".test.ts") {
					continue
				}

				// Only rebuild on Write events
				if event.Op&fsnotify.Write != fsnotify.Write {
					continue
				}

				server.logger.Debug("Element file changed: %s", event.Name)

				// Reset debounce timer
				if debounceTimer != nil {
					debounceTimer.Stop()
				}

				debounceTimer = time.AfterFunc(debounceDelay, func() {
					server.logger.Info("Transpiling elements...")
					if err := elements.TranspileElements(elementsDir); err != nil {
						server.logger.Error("Failed to transpile elements: %v", err)
						// Broadcast error to clients
						if server.wsManager != nil {
							errMsg := fmt.Sprintf(`{"type":"error","title":"Element Transpilation Error","message":"Failed to transpile elements: %v"}`, err)
							_ = server.wsManager.Broadcast([]byte(errMsg))
						}
						return
					}

					server.logger.Info("Elements transpiled successfully, reloading clients...")

					// Trigger reload
					if server.wsManager != nil {
						reloadMsg := []byte(`{"type":"reload"}`)
						_ = server.wsManager.Broadcast(reloadMsg)
					}
				})

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				server.logger.Error("Element watcher error: %v", err)
			}
		}
	}()

	return nil
}
