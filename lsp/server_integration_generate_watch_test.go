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
	"os"
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/lsp"
	"bennypowers.dev/cem/lsp/testhelpers"
	W "bennypowers.dev/cem/workspace"
)

func TestGenerateWatcherIntegration(t *testing.T) {
	// Create a temporary directory for test files
	tempDir, err := os.MkdirTemp("", "lsp-generate-watch-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Copy fixture files to create a realistic project structure
	fixtureDir := filepath.Join("test", "fixtures", "generate-watch-test")

	sourceFixture := filepath.Join(fixtureDir, "test-element.ts")
	sourceFile := filepath.Join(tempDir, "test-element.ts")
	err = testhelpers.CopyFile(sourceFixture, sourceFile)
	if err != nil {
		t.Fatalf("Failed to copy source fixture: %v", err)
	}

	packageFixture := filepath.Join(fixtureDir, "package.json")
	packageJSON := filepath.Join(tempDir, "package.json")
	err = testhelpers.CopyFile(packageFixture, packageJSON)
	if err != nil {
		t.Fatalf("Failed to copy package.json fixture: %v", err)
	}

	t.Run("generate watcher starts for local project", func(t *testing.T) {
		// Create workspace and registry
		workspace := W.NewFileSystemWorkspaceContext(tempDir)
		err := workspace.Init()
		if err != nil {
			t.Fatalf("Failed to initialize workspace: %v", err)
		}

		registry, err := lsp.NewRegistryWithDefaults()
		if err != nil {
			t.Fatalf("Failed to create registry: %v", err)
		}

		// Load from workspace (this should detect it as a local project)
		err = registry.LoadFromWorkspace(workspace)
		if err != nil {
			t.Fatalf("Failed to load from workspace: %v", err)
		}

		// Start the generate watcher
		err = registry.StartGenerateWatcher()
		if err != nil {
			t.Fatalf("Failed to start generate watcher: %v", err)
		}
		defer registry.StopGenerateWatcher()

		// No delay needed - Go 1.25 eliminates timing dependencies
		// The watcher should be running (we can't easily test the actual process without complex setup)
		// For now, just verify that StartGenerateWatcher didn't error
		// In a real scenario, we would modify the source file and check if the manifest gets updated
	})

	t.Run("generate watcher stops cleanly", func(t *testing.T) {
		workspace := W.NewFileSystemWorkspaceContext(tempDir)
		err = workspace.Init()
		if err != nil {
			t.Fatalf("Failed to initialize workspace: %v", err)
		}

		registry, err := lsp.NewRegistryWithDefaults()
		if err != nil {
			t.Fatalf("Failed to create registry: %v", err)
		}
		err = registry.LoadFromWorkspace(workspace)
		if err != nil {
			t.Fatalf("Failed to load from workspace: %v", err)
		}

		// Start and then stop the watcher
		err = registry.StartGenerateWatcher()
		if err != nil {
			t.Fatalf("Failed to start generate watcher: %v", err)
		}

		// No delay needed - Go 1.25 eliminates timing dependencies between start and stop

		err = registry.StopGenerateWatcher()
		if err != nil {
			t.Fatalf("Failed to stop generate watcher: %v", err)
		}
	})

	t.Run("generate watcher handles no local workspace", func(t *testing.T) {
		// Create a registry without loading a local workspace
		registry, err := lsp.NewRegistryWithDefaults()
		if err != nil {
			t.Fatalf("Failed to create registry: %v", err)
		}

		// Should not error but also should not start a watcher
		err = registry.StartGenerateWatcher()
		if err != nil {
			t.Errorf("StartGenerateWatcher should handle no local workspace gracefully, got error: %v", err)
		}

		// Stopping should also not error
		err = registry.StopGenerateWatcher()
		if err != nil {
			t.Errorf("StopGenerateWatcher should handle no watcher gracefully, got error: %v", err)
		}
	})

	t.Run("generate watcher prevents double start", func(t *testing.T) {
		workspace := W.NewFileSystemWorkspaceContext(tempDir)
		err = workspace.Init()
		if err != nil {
			t.Fatalf("Failed to initialize workspace: %v", err)
		}

		registry, err := lsp.NewRegistryWithDefaults()
		if err != nil {
			t.Fatalf("Failed to create registry: %v", err)
		}
		err = registry.LoadFromWorkspace(workspace)
		if err != nil {
			t.Fatalf("Failed to load from workspace: %v", err)
		}

		// Start watcher twice
		err1 := registry.StartGenerateWatcher()
		err2 := registry.StartGenerateWatcher()

		if err1 != nil {
			t.Errorf("First StartGenerateWatcher failed: %v", err1)
		}
		if err2 != nil {
			t.Errorf("Second StartGenerateWatcher should not error when already running: %v", err2)
		}

		defer registry.StopGenerateWatcher()
	})
}
