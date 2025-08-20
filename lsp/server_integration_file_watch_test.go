package lsp_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"

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
		// Create fresh registry for this test
		testRegistry, err := lsp.NewRegistryWithDefaults()
		if err != nil {
			t.Fatalf("Failed to create registry: %v", err)
		}
		err = testRegistry.LoadFromWorkspace(workspace)
		if err != nil {
			t.Fatalf("Failed to load from workspace: %v", err)
		}

		var reloadCount int
		var reloadMu sync.Mutex

		// Set up file watching
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

		// Wait for file system event to be processed
		timeout := time.After(3 * time.Second)
		for {
			select {
			case <-timeout:
				t.Fatal("Timeout waiting for file change to be detected")
			default:
				reloadMu.Lock()
				count := reloadCount
				reloadMu.Unlock()

				if count > 0 {
					// Reload was triggered, verify the registry was updated
					if len(testRegistry.Elements) == 2 {
						if _, exists := testRegistry.Elements["new-element"]; exists {
							return // Success!
						}
					}
				}
				time.Sleep(100 * time.Millisecond)
			}
		}
	})

	t.Run("multiple file changes", func(t *testing.T) {
		// Create fresh registry for this test
		testRegistry, err := lsp.NewRegistryWithDefaults()
		if err != nil {
			t.Fatalf("Failed to create registry: %v", err)
		}
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

		// Make multiple rapid changes by copying fixtures alternately
		fixtures := []string{"initial-manifest.json", "updated-manifest.json"}
		for i := 0; i < 3; i++ {
			fixtureName := fixtures[i%2]
			fixturePath := filepath.Join("test", "fixtures", "file-watch-integration", fixtureName)
			err = testhelpers.CopyFile(fixturePath, manifestPath)
			if err != nil {
				t.Fatalf("Failed to copy manifest fixture %s: %v", fixtureName, err)
			}
			time.Sleep(200 * time.Millisecond)
		}

		// Wait for changes to be processed
		time.Sleep(1 * time.Second)

		reloadMu.Lock()
		count := reloadCount
		reloadMu.Unlock()

		if count == 0 {
			t.Error("Expected at least one reload, got none")
		}
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
		var reloadTriggered bool
		var reloadMu sync.Mutex

		onReload := func() {
			reloadMu.Lock()
			defer reloadMu.Unlock()
			reloadTriggered = true
		}

		err := registry.StartFileWatching(onReload)
		if err != nil {
			t.Fatalf("Failed to start file watching: %v", err)
		}
		defer registry.StopFileWatching()

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

		// Wait for file change detection
		timeout := time.After(2 * time.Second)
		for {
			select {
			case <-timeout:
				t.Fatal("Timeout waiting for package.json change to be detected")
			default:
				reloadMu.Lock()
				triggered := reloadTriggered
				reloadMu.Unlock()

				if triggered {
					return // Success!
				}
				time.Sleep(100 * time.Millisecond)
			}
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
