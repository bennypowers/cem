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
	"testing"

	"bennypowers.dev/cem/lsp"
	M "bennypowers.dev/cem/manifest"
	W "bennypowers.dev/cem/workspace"
)

func TestServer_Workspace(t *testing.T) {
	// Create workspace
	workspace := W.NewFileSystemWorkspaceContext("/test/workspace")
	if err := workspace.Init(); err != nil {
		t.Fatalf("Failed to init workspace: %v", err)
	}

	// Create server with workspace
	server, err := lsp.NewServer(workspace, lsp.TransportStdio)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer func() {
		if err := server.Close(); err != nil {
			t.Fatalf("Failed to close server: %v", err)
		}
	}()

	// Test Workspace() method
	retrievedWorkspace := server.Workspace()
	if retrievedWorkspace == nil {
		t.Fatal("Expected workspace, got nil")
	}

	if retrievedWorkspace.Root() != "/test/workspace" {
		t.Errorf("Expected workspace root '/test/workspace', got '%s'", retrievedWorkspace.Root())
	}
}

func TestServer_WorkspaceRoot(t *testing.T) {
	tests := []struct {
		name     string
		root     string
		expected string
	}{
		{
			name:     "NormalPath",
			root:     "/home/user/project",
			expected: "/home/user/project",
		},
		{
			name:     "RootPath",
			root:     "/",
			expected: "/",
		},
		{
			name:     "RelativePath",
			root:     "./project",
			expected: "./project",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			workspace := W.NewFileSystemWorkspaceContext(tt.root)
			if err := workspace.Init(); err != nil {
				t.Fatalf("Failed to init workspace: %v", err)
			}

			server, err := lsp.NewServer(workspace, lsp.TransportStdio)
			if err != nil {
				t.Fatalf("Failed to create server: %v", err)
			}
			defer func() {
				if err := server.Close(); err != nil {
					t.Fatalf("Failed to close server: %v", err)
				}
			}()

			root := server.WorkspaceRoot()
			if root != tt.expected {
				t.Errorf("Expected workspace root '%s', got '%s'", tt.expected, root)
			}
		})
	}
}

func TestServer_Element(t *testing.T) {
	workspace := W.NewFileSystemWorkspaceContext("/test")
	if err := workspace.Init(); err != nil {
		t.Fatalf("Failed to init workspace: %v", err)
	}

	server, err := lsp.NewServer(workspace, lsp.TransportStdio)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer func() { _ = server.Close() }()

	// Load manifest from fixture file
	manifestJSON, err := os.ReadFile("test-fixtures/server-context/test-element.json")
	if err != nil {
		t.Fatalf("Failed to read manifest fixture: %v", err)
	}

	var manifest M.Package
	if err := json.Unmarshal(manifestJSON, &manifest); err != nil {
		t.Fatalf("Failed to unmarshal manifest: %v", err)
	}

	server.AddManifest(&manifest)

	// Test getting an element that exists
	element, exists := server.Element("test-element")
	if !exists {
		t.Fatal("Expected element to exist")
	}

	if element.TagName != "test-element" {
		t.Errorf("Expected tag name 'test-element', got '%s'", element.TagName)
	}

	// Test getting an element that doesn't exist
	_, exists = server.Element("non-existent")
	if exists {
		t.Error("Expected element to not exist")
	}
}

func TestServer_Attributes(t *testing.T) {
	workspace := W.NewFileSystemWorkspaceContext("/test")
	if err := workspace.Init(); err != nil {
		t.Fatalf("Failed to init workspace: %v", err)
	}

	server, err := lsp.NewServer(workspace, lsp.TransportStdio)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer func() { _ = server.Close() }()

	// Load manifest from fixture file
	manifestJSON, err := os.ReadFile("test-fixtures/server-context/element-with-attributes.json")
	if err != nil {
		t.Fatalf("Failed to read manifest fixture: %v", err)
	}

	var manifest M.Package
	if err := json.Unmarshal(manifestJSON, &manifest); err != nil {
		t.Fatalf("Failed to unmarshal manifest: %v", err)
	}

	server.AddManifest(&manifest)

	// Test getting attributes for element that exists
	attrs, exists := server.Attributes("test-element")
	if !exists {
		t.Fatal("Expected attributes to exist")
	}

	if len(attrs) != 2 {
		t.Errorf("Expected 2 attributes, got %d", len(attrs))
	}

	if attrs["variant"] == nil {
		t.Error("Expected variant attribute to exist")
	}

	if attrs["size"] == nil {
		t.Error("Expected size attribute to exist")
	}

	// Test getting attributes for element that doesn't exist
	_, exists = server.Attributes("non-existent")
	if exists {
		t.Error("Expected attributes to not exist")
	}
}

func TestServer_Slots(t *testing.T) {
	workspace := W.NewFileSystemWorkspaceContext("/test")
	if err := workspace.Init(); err != nil {
		t.Fatalf("Failed to init workspace: %v", err)
	}

	server, err := lsp.NewServer(workspace, lsp.TransportStdio)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer func() { _ = server.Close() }()

	// Load manifest from fixture file
	manifestJSON, err := os.ReadFile("test-fixtures/server-context/element-with-slots.json")
	if err != nil {
		t.Fatalf("Failed to read manifest fixture: %v", err)
	}

	var manifest M.Package
	if err := json.Unmarshal(manifestJSON, &manifest); err != nil {
		t.Fatalf("Failed to unmarshal manifest: %v", err)
	}

	server.AddManifest(&manifest)

	// Test getting slots for element that exists
	slots, exists := server.Slots("test-element")
	if !exists {
		t.Fatal("Expected slots to exist")
	}

	if len(slots) != 2 {
		t.Errorf("Expected 2 slots, got %d", len(slots))
	}

	if slots[0].Name != "header" {
		t.Errorf("Expected first slot name 'header', got '%s'", slots[0].Name)
	}

	if slots[1].Name != "footer" {
		t.Errorf("Expected second slot name 'footer', got '%s'", slots[1].Name)
	}

	// Test getting slots for element that doesn't exist
	_, exists = server.Slots("non-existent")
	if exists {
		t.Error("Expected slots to not exist")
	}
}

func TestServer_AddManifest(t *testing.T) {
	workspace := W.NewFileSystemWorkspaceContext("/test")
	if err := workspace.Init(); err != nil {
		t.Fatalf("Failed to init workspace: %v", err)
	}

	server, err := lsp.NewServer(workspace, lsp.TransportStdio)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer func() { _ = server.Close() }()

	// Initially no elements
	initialCount := server.ElementCount()

	// Load manifest from fixture file
	manifestJSON, err := os.ReadFile("test-fixtures/server-context/multiple-elements.json")
	if err != nil {
		t.Fatalf("Failed to read manifest fixture: %v", err)
	}

	var manifest M.Package
	if err := json.Unmarshal(manifestJSON, &manifest); err != nil {
		t.Fatalf("Failed to unmarshal manifest: %v", err)
	}

	server.AddManifest(&manifest)

	// Should have 2 more elements now
	newCount := server.ElementCount()
	if newCount != initialCount+2 {
		t.Errorf("Expected %d elements after adding manifest, got %d", initialCount+2, newCount)
	}

	// Verify elements are accessible
	_, exists := server.Element("element-one")
	if !exists {
		t.Error("Expected element-one to exist after adding manifest")
	}

	_, exists = server.Element("element-two")
	if !exists {
		t.Error("Expected element-two to exist after adding manifest")
	}
}

func TestServer_ElementSource(t *testing.T) {
	tests := []struct {
		name            string
		fixtureFile     string
		tagName         string
		expectedSource  string
		expectedExists  bool
	}{
		{
			name:           "ElementWithModulePath",
			fixtureFile:    "test-fixtures/server-context/package-with-module-path.json",
			tagName:        "test-button",
			expectedSource: "button/button-element.js",
			expectedExists: true,
		},
		{
			name:           "ElementWithSimplePath",
			fixtureFile:    "test-fixtures/server-context/package-only.json",
			tagName:        "test-card",
			expectedSource: "test-card.js",
			expectedExists: true,
		},
		{
			name:           "NonExistentElement",
			fixtureFile:    "test-fixtures/server-context/empty-manifest.json",
			tagName:        "non-existent",
			expectedSource: "",
			expectedExists: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			workspace := W.NewFileSystemWorkspaceContext("/test/workspace")
			if err := workspace.Init(); err != nil {
				t.Fatalf("Failed to init workspace: %v", err)
			}

			server, err := lsp.NewServer(workspace, lsp.TransportStdio)
			if err != nil {
				t.Fatalf("Failed to create server: %v", err)
			}
			defer func() { _ = server.Close() }()

			// Load manifest from fixture file
			manifestJSON, err := os.ReadFile(tt.fixtureFile)
			if err != nil {
				t.Fatalf("Failed to read manifest fixture %s: %v", tt.fixtureFile, err)
			}

			var manifest M.Package
			if err := json.Unmarshal(manifestJSON, &manifest); err != nil {
				t.Fatalf("Failed to unmarshal manifest %s: %v", tt.fixtureFile, err)
			}

			server.AddManifest(&manifest)

			source, exists := server.ElementSource(tt.tagName)

			if exists != tt.expectedExists {
				t.Errorf("Expected exists=%v, got exists=%v", tt.expectedExists, exists)
			}

			if source != tt.expectedSource {
				t.Errorf("Expected source '%s', got '%s'", tt.expectedSource, source)
			}
		})
	}
}

func TestServer_ModuleGraph(t *testing.T) {
	workspace := W.NewFileSystemWorkspaceContext("/test")
	if err := workspace.Init(); err != nil {
		t.Fatalf("Failed to init workspace: %v", err)
	}

	server, err := lsp.NewServer(workspace, lsp.TransportStdio)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer func() { _ = server.Close() }()

	// Get module graph
	graph := server.ModuleGraph()

	// Should not be nil for a properly initialized server
	if graph == nil {
		t.Error("Expected module graph to not be nil")
	}
}

func TestServer_ElementDescription(t *testing.T) {
	workspace := W.NewFileSystemWorkspaceContext("/test")
	if err := workspace.Init(); err != nil {
		t.Fatalf("Failed to init workspace: %v", err)
	}

	server, err := lsp.NewServer(workspace, lsp.TransportStdio)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer func() { _ = server.Close() }()

	// Load manifest from fixture file
	manifestJSON, err := os.ReadFile("test-fixtures/server-context/test-element.json")
	if err != nil {
		t.Fatalf("Failed to read manifest fixture: %v", err)
	}

	var manifest M.Package
	if err := json.Unmarshal(manifestJSON, &manifest); err != nil {
		t.Fatalf("Failed to unmarshal manifest: %v", err)
	}

	server.AddManifest(&manifest)

	// Test getting description (currently returns empty)
	desc, exists := server.ElementDescription("test-element")

	// Element exists, but description is not implemented yet
	// So exists should be false and desc should be empty
	if exists {
		t.Error("Expected exists to be false (description not implemented)")
	}

	if desc != "" {
		t.Errorf("Expected empty description, got '%s'", desc)
	}

	// Test non-existent element
	_, exists = server.ElementDescription("non-existent")
	if exists {
		t.Error("Expected description to not exist for non-existent element")
	}
}
