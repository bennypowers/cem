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
package lsp_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"testing/synctest"
	"time"

	"bennypowers.dev/cem/internal/platform"
	"bennypowers.dev/cem/lsp"
	"bennypowers.dev/cem/lsp/testhelpers"
	W "bennypowers.dev/cem/workspace"
)

func TestManifestFileWatchingIntegration(t *testing.T) {
	// Create a temporary directory for test files
	tempDir, err := os.MkdirTemp("", "lsp-file-watch-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Copy initial manifest fixture to temp directory
	fixturePath := filepath.Join("test", "fixtures", "file-watch-integration", "initial-manifest.json")
	manifestPath := filepath.Join(tempDir, "custom-elements.json")

	err = testhelpers.CopyFile(fixturePath, manifestPath)
	if err != nil {
		t.Fatalf("Failed to copy initial manifest fixture: %v", err)
	}

	// Create workspace context
	workspace := W.NewFileSystemWorkspaceContext(tempDir)
	err = workspace.Init() // Initialize workspace to discover manifest paths
	if err != nil {
		t.Fatalf("Failed to initialize workspace: %v", err)
	}
	registry, err := lsp.NewRegistryWithDefaults()
	if err != nil {
		t.Fatalf("Failed to create registry: %v", err)
	}

	t.Run("file watching setup and initial load", func(t *testing.T) {
		// Load from workspace (simulating server startup)
		err := registry.LoadFromWorkspace(workspace)
		if err != nil {
			t.Fatalf("Failed to load from workspace: %v", err)
		}

		// Verify initial state
		if len(registry.Elements) != 1 {
			t.Errorf("Expected 1 element, got %d", len(registry.Elements))
		}
		if _, exists := registry.Elements["test-element"]; !exists {
			t.Error("Expected test-element to be loaded")
		}

		// Verify manifest path is tracked
		if len(registry.ManifestPaths) == 0 {
			t.Error("Expected manifest paths to be tracked")
		}
	})

	t.Run("file watching triggers reload", func(t *testing.T) {
		synctest.Test(t, func(t *testing.T) {
			// Go 1.25: Using synctest with MockFileWatcher for deterministic file watching tests.
			// This eliminates timing dependencies while testing the actual file watcher integration.

			// Create registry with MockFileWatcher for testing
			mockFileWatcher := platform.NewMockFileWatcher()
			testRegistry := lsp.NewRegistry(mockFileWatcher)
			err = testRegistry.LoadFromWorkspace(workspace)
			if err != nil {
				t.Fatalf("Failed to load from workspace: %v", err)
			}

			var reloadCount int
			var reloadMu sync.Mutex

			// Set up file watching with reload callback
			onReload := func() {
				reloadMu.Lock()
				defer reloadMu.Unlock()
				reloadCount++
				// Simulate the server's reload process using public API
				// Create fresh workspace to avoid caching issues
				freshWorkspace := W.NewFileSystemWorkspaceContext(tempDir)
				err := freshWorkspace.Init()
				if err != nil {
					return
				}
				testRegistry.LoadFromWorkspace(freshWorkspace)
			}

			err = testRegistry.StartFileWatching(onReload)
			if err != nil {
				t.Fatalf("Failed to start file watching: %v", err)
			}
			defer testRegistry.StopFileWatching()

			// Copy updated manifest fixture to trigger file change
			updatedFixturePath := filepath.Join("test", "fixtures", "file-watch-integration", "updated-manifest.json")
			err = testhelpers.CopyFile(updatedFixturePath, manifestPath)
			if err != nil {
				t.Fatalf("Failed to copy updated manifest fixture: %v", err)
			}

			// Simulate file change detection via MockFileWatcher - instant with virtual time
			// This tests the actual file watcher event handling but without real filesystem delays

			// Get the absolute path to match what the registry tracks
			absManifestPath, err := filepath.Abs(manifestPath)
			if err != nil {
				t.Fatalf("Failed to get absolute path: %v", err)
			}

			// Trigger event with absolute path
			mockFileWatcher.TriggerEvent(absManifestPath, platform.Write)

			// Yield to allow the watchFiles goroutine to process the event in synctest
			var quickWait sync.WaitGroup
			quickWait.Add(1)
			go func() {
				defer quickWait.Done()
			}()
			quickWait.Wait()

			// Verify reload was triggered - instant with virtual time
			reloadMu.Lock()
			count := reloadCount
			reloadMu.Unlock()

			if count == 0 {
				t.Fatal("Reload should have been triggered by file watcher event")
			}

			// Verify the registry was updated with the new manifest content
			if len(testRegistry.Elements) == 2 {
				if _, exists := testRegistry.Elements["new-element"]; exists {
					return // Success!
				}
			}
			t.Errorf("Registry not properly updated after file change. Elements: %d", len(testRegistry.Elements))
		})
	})

	t.Run("multiple file changes", func(t *testing.T) {
		synctest.Test(t, func(t *testing.T) {
			// Create registry with MockFileWatcher for testing
			mockFileWatcher := platform.NewMockFileWatcher()
			testRegistry := lsp.NewRegistry(mockFileWatcher)
			err = testRegistry.LoadFromWorkspace(workspace)
			if err != nil {
				t.Fatalf("Failed to load from workspace: %v", err)
			}

			var reloadCount int
			var reloadMu sync.Mutex

			onReload := func() {
				reloadMu.Lock()
				defer reloadMu.Unlock()
				reloadCount++
				// Create fresh workspace to avoid caching issues
				freshWorkspace := W.NewFileSystemWorkspaceContext(tempDir)
				err := freshWorkspace.Init()
				if err != nil {
					return
				}
				testRegistry.LoadFromWorkspace(freshWorkspace)
			}

			err = testRegistry.StartFileWatching(onReload)
			if err != nil {
				t.Fatalf("Failed to start file watching: %v", err)
			}
			defer testRegistry.StopFileWatching()

			// Use the MockFileWatcher for triggering events

			// Make multiple rapid changes by copying fixtures alternately
			fixtures := []string{"initial-manifest.json", "updated-manifest.json"}

			// Get absolute path for consistent triggering
			absManifestPath, err := filepath.Abs(manifestPath)
			if err != nil {
				t.Fatalf("Failed to get absolute path: %v", err)
			}

			for i := 0; i < 3; i++ {
				fixtureName := fixtures[i%2]
				fixturePath := filepath.Join("test", "fixtures", "file-watch-integration", fixtureName)
				err = testhelpers.CopyFile(fixturePath, manifestPath)
				if err != nil {
					t.Fatalf("Failed to copy manifest fixture %s: %v", fixtureName, err)
				}
				// Trigger file change event via MockFileWatcher - instant with virtual time
				mockFileWatcher.TriggerEvent(absManifestPath, platform.Write)

				// Yield to allow the watchFiles goroutine to process the event
				var quickWait sync.WaitGroup
				quickWait.Add(1)
				go func() {
					defer quickWait.Done()
				}()
				quickWait.Wait()
			}

			// Verify multiple changes were processed - instant with virtual time
			reloadMu.Lock()
			count := reloadCount
			reloadMu.Unlock()

			if count == 0 {
				t.Error("Expected at least one reload, got none")
			}
			if count != 3 {
				t.Logf("Expected 3 reloads, got %d", count)
			}
		})
	})
}

func TestPackageJSONWatching(t *testing.T) {
	// Create a temporary directory for test files
	tempDir, err := os.MkdirTemp("", "lsp-package-watch-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Copy manifest to the referenced location
	manifestPath := filepath.Join(tempDir, "custom-elements.json")
	fixturePath := filepath.Join("test", "fixtures", "file-watch-integration", "initial-manifest.json")
	err = testhelpers.CopyFile(fixturePath, manifestPath)
	if err != nil {
		t.Fatalf("Failed to copy manifest fixture: %v", err)
	}

	// Create a mock node_modules structure to trigger package.json discovery
	nodeModulesPath := filepath.Join(tempDir, "node_modules", "test-package")
	err = os.MkdirAll(nodeModulesPath, 0755)
	if err != nil {
		t.Fatalf("Failed to create node_modules structure: %v", err)
	}

	// Copy package.json fixture to the mock package
	packagePath := filepath.Join(nodeModulesPath, "package.json")
	packageFixturePath := filepath.Join("test", "fixtures", "file-watch-integration", "package.json")
	err = testhelpers.CopyFile(packageFixturePath, packagePath)
	if err != nil {
		t.Fatalf("Failed to copy package.json fixture: %v", err)
	}

	// Copy the manifest to the location referenced by package.json
	manifestInPackagePath := filepath.Join(nodeModulesPath, "custom-elements.json")
	err = testhelpers.CopyFile(fixturePath, manifestInPackagePath)
	if err != nil {
		t.Fatalf("Failed to copy manifest to package location: %v", err)
	}

	registry, err := lsp.NewRegistryWithDefaults()
	if err != nil {
		t.Fatalf("Failed to create registry: %v", err)
	}

	t.Run("package.json with customElements tracked", func(t *testing.T) {
		// Create workspace and load manifests to trigger package.json discovery
		workspace := W.NewFileSystemWorkspaceContext(tempDir)
		err := workspace.Init()
		if err != nil {
			t.Fatalf("Failed to initialize workspace: %v", err)
		}

		err = registry.LoadFromWorkspace(workspace)
		if err != nil {
			t.Fatalf("Failed to load from workspace: %v", err)
		}

		// Verify package.json is tracked for watching
		found := false
		for _, path := range registry.ManifestPaths {
			if filepath.Base(path) == "package.json" {
				found = true
				break
			}
		}
		if !found {
			t.Error("Expected package.json to be tracked for watching")
		}
	})

	t.Run("package.json changes trigger reload", func(t *testing.T) {
		// Create registry with MockFileWatcher for testing
		mockFileWatcher := platform.NewMockFileWatcher()
		testRegistry := lsp.NewRegistry(mockFileWatcher)

		// Load workspace to populate manifest paths
		workspace := W.NewFileSystemWorkspaceContext(tempDir)
		err = workspace.Init()
		if err != nil {
			t.Fatalf("Failed to initialize workspace: %v", err)
		}
		err = testRegistry.LoadFromWorkspace(workspace)
		if err != nil {
			t.Fatalf("Failed to load from workspace: %v", err)
		}

		// Use a simple mutex approach like the other tests
		var reloadTriggered bool
		var reloadMu sync.Mutex

		onReload := func() {
			reloadMu.Lock()
			defer reloadMu.Unlock()
			reloadTriggered = true
		}

		err = testRegistry.StartFileWatching(onReload)
		if err != nil {
			t.Fatalf("Failed to start file watching: %v", err)
		}
		defer testRegistry.StopFileWatching()

		// Update package.json by reading, modifying, and writing back
		var packageJSON map[string]interface{}
		packageData, err := os.ReadFile(packagePath)
		if err != nil {
			t.Fatalf("Failed to read package.json: %v", err)
		}
		err = json.Unmarshal(packageData, &packageJSON)
		if err != nil {
			t.Fatalf("Failed to parse package.json: %v", err)
		}

		packageJSON["version"] = "1.0.1"
		updatedData, _ := json.MarshalIndent(packageJSON, "", "  ")
		err = os.WriteFile(packagePath, updatedData, 0644)
		if err != nil {
			t.Fatalf("Failed to update package.json: %v", err)
		}

		// Trigger file change event via MockFileWatcher
		// Get absolute path for consistent triggering
		absPackagePath, err := filepath.Abs(packagePath)
		if err != nil {
			t.Fatalf("Failed to get absolute path: %v", err)
		}

		mockFileWatcher.TriggerEvent(absPackagePath, platform.Write)

		// Give time for the event to be processed (using real time)
		time.Sleep(10 * time.Millisecond)

		// Verify reload was triggered
		reloadMu.Lock()
		triggered := reloadTriggered
		reloadMu.Unlock()

		if !triggered {
			t.Fatal("package.json change should have triggered reload")
		}
	})
}

func TestFileWatchingErrorHandling(t *testing.T) {
	registry, err := lsp.NewRegistryWithDefaults()
	if err != nil {
		t.Fatalf("Failed to create registry: %v", err)
	}

	t.Run("start watching without files", func(t *testing.T) {
		// Should succeed even with no files to watch
		err := registry.StartFileWatching(func() {})
		if err != nil {
			t.Errorf("Expected StartFileWatching to succeed with no files, got error: %v", err)
		}
		defer registry.StopFileWatching()
	})

	t.Run("stop watching when not started", func(t *testing.T) {
		// Should not error
		err := registry.StopFileWatching()
		if err != nil {
			t.Errorf("Expected StopFileWatching to succeed when not started, got error: %v", err)
		}
	})

	t.Run("double start watching", func(t *testing.T) {
		err1 := registry.StartFileWatching(func() {})
		err2 := registry.StartFileWatching(func() {})

		if err1 != nil {
			t.Errorf("First StartFileWatching failed: %v", err1)
		}
		if err2 != nil {
			t.Errorf("Second StartFileWatching failed: %v", err2)
		}

		defer registry.StopFileWatching()
	})

	t.Run("watch nonexistent file", func(t *testing.T) {
		// Create a registry and add a path that doesn't exist
		testRegistry, err := lsp.NewRegistryWithDefaults()
		if err != nil {
			t.Fatalf("Failed to create registry: %v", err)
		}
		testRegistry.AddManifestPath("/nonexistent/path/manifest.json")

		// Starting the watcher should not fail due to one bad path
		err = testRegistry.StartFileWatching(func() {})
		if err != nil {
			t.Errorf("Expected StartFileWatching to handle nonexistent files gracefully, got error: %v", err)
		}
		defer testRegistry.StopFileWatching()
	})
}
