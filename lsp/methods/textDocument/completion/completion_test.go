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
package completion_test

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	G "bennypowers.dev/cem/generate"
	"bennypowers.dev/cem/lsp"
	"bennypowers.dev/cem/lsp/methods/textDocument/completion"
	"bennypowers.dev/cem/lsp/testhelpers"
	M "bennypowers.dev/cem/manifest"
	W "bennypowers.dev/cem/workspace"
)

// copyFixtureFiles copies files from the fixture directory to the target directory,
// preserving directory structure
func copyFixtureFiles(t *testing.T, fixtureDir, targetDir string) {
	t.Helper()

	err := filepath.Walk(fixtureDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip README files
		if strings.HasSuffix(path, "README.md") {
			return nil
		}

		// Get relative path from fixture directory
		relPath, err := filepath.Rel(fixtureDir, path)
		if err != nil {
			return err
		}

		targetPath := filepath.Join(targetDir, relPath)

		if info.IsDir() {
			return os.MkdirAll(targetPath, info.Mode())
		}

		// Copy file
		src, err := os.Open(path)
		if err != nil {
			return err
		}
		defer src.Close()

		err = os.MkdirAll(filepath.Dir(targetPath), 0755)
		if err != nil {
			return err
		}

		dst, err := os.Create(targetPath)
		if err != nil {
			return err
		}
		defer dst.Close()

		_, err = io.Copy(dst, src)
		return err
	})

	if err != nil {
		t.Fatalf("Failed to copy fixture files: %v", err)
	}
}

// TestServerLevelIntegration tests the complete server integration including
// manifest reloading behavior exactly as it would work in real usage
func TestServerLevelIntegration(t *testing.T) {
	// Create a temporary workspace directory that mimics a real project
	tempDir := t.TempDir()

	// Copy fixture files to temporary directory
	fixtureDir := filepath.Join("test-fixtures", "server-integration")
	CopyFixtureFiles(t, fixtureDir, tempDir)

	// Create dist directory
	distDir := filepath.Join(tempDir, "dist")
	err := os.MkdirAll(distDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create dist directory: %v", err)
	}

	// Get paths for later use
	tsFilePath := filepath.Join(tempDir, "src", "test-alert.ts")

	// Create initial manifest by running generate manually
	workspace := W.NewFileSystemWorkspaceContext(tempDir)
	err = workspace.Init()
	if err != nil {
		t.Fatalf("Failed to initialize workspace: %v", err)
	}

	// Run initial generation to create the manifest
	generateSession, err := G.NewGenerateSession(workspace)
	if err != nil {
		t.Fatalf("Failed to create generate session: %v", err)
	}
	defer generateSession.Close()

	pkg, err := generateSession.GenerateFullManifest(context.Background())
	if err != nil {
		t.Fatalf("Failed to generate initial manifest: %v", err)
	}

	// Write initial manifest to file
	manifestStr, err := M.SerializeToString(pkg)
	if err != nil {
		t.Fatalf("Failed to serialize manifest: %v", err)
	}

	manifestPath := filepath.Join(distDir, "custom-elements.json")
	err = os.WriteFile(manifestPath, []byte(manifestStr), 0644)
	if err != nil {
		t.Fatalf("Failed to write manifest: %v", err)
	}

	// Now create a full LSP server instance - this simulates real usage
	server, err := lsp.NewServer(workspace)
	if err != nil {
		t.Fatalf("Failed to create LSP server: %v", err)
	}
	defer server.Close()

	// Initialize the server (this loads manifests and starts file watching)
	err = server.InitializeForTesting()
	if err != nil {
		t.Fatalf("Failed to initialize server: %v", err)
	}

	// Create document manager
	dm, err := lsp.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create document manager: %v", err)
	}
	defer dm.Close()

	// Set up the server context adapter with document manager
	ctx := &testCompletionContextWithDM{
		registry: server.Registry(), // We need to expose this for testing
		docMgr:   dm,
	}

	// Get paths for HTML file (already copied from fixtures)
	htmlFilePath := filepath.Join(tempDir, "index.html")

	// Read the HTML content
	htmlContentBytes, err := os.ReadFile(htmlFilePath)
	if err != nil {
		t.Fatalf("Failed to read HTML file: %v", err)
	}
	htmlContent := string(htmlContentBytes)

	// Open the HTML document in document manager
	htmlURI := "file://" + htmlFilePath
	doc := dm.OpenDocument(htmlURI, htmlContent, 1)
	if doc == nil {
		t.Fatalf("Failed to open HTML document")
	}

	// Test initial completions - should have info, success, warning
	initialItems := completion.GetAttributeValueCompletions(ctx, "test-alert", "state")

	t.Logf("Initial completions: %v", testhelpers.GetCompletionLabels(initialItems))

	// Verify initial completions contain expected values
	hasInfo := false
	hasSuccess := false
	hasWarning := false
	for _, item := range initialItems {
		switch item.Label {
		case "info":
			hasInfo = true
		case "success":
			hasSuccess = true
		case "warning":
			hasWarning = true
		}
	}

	if !hasInfo || !hasSuccess || !hasWarning {
		t.Fatalf("Initial completions missing expected values. Got: %v", testhelpers.GetCompletionLabels(initialItems))
	}

	// Now simulate the user editing the TypeScript file to add 'error'
	t.Logf("=== User edits test-alert.ts to add 'error' ===")

	// Copy the updated version from fixtures
	updatedTSPath := filepath.Join(fixtureDir, "src", "test-alert-updated.ts")
	updatedTSContent, err := os.ReadFile(updatedTSPath)
	if err != nil {
		t.Fatalf("Failed to read updated TypeScript file: %v", err)
	}

	err = os.WriteFile(tsFilePath, updatedTSContent, 0644)
	if err != nil {
		t.Fatalf("Failed to update TypeScript file: %v", err)
	}

	t.Logf("File updated, waiting for generate watcher to process changes...")

	// Wait for the generate watcher to detect changes and regenerate
	time.Sleep(5 * time.Second)

	// Test updated completions - should now include 'error'
	updatedItems := completion.GetAttributeValueCompletions(ctx, "test-alert", "state")

	t.Logf("Updated completions: %v", testhelpers.GetCompletionLabels(updatedItems))

	// Verify updated completions contain 'error'
	hasUpdatedInfo := false
	hasUpdatedSuccess := false
	hasUpdatedWarning := false
	hasError := false
	for _, item := range updatedItems {
		switch item.Label {
		case "info":
			hasUpdatedInfo = true
		case "success":
			hasUpdatedSuccess = true
		case "warning":
			hasUpdatedWarning = true
		case "error":
			hasError = true
		}
	}

	if !hasUpdatedInfo || !hasUpdatedSuccess || !hasUpdatedWarning {
		t.Errorf("Updated completions missing original values. Got: %v", testhelpers.GetCompletionLabels(updatedItems))
	}

	if !hasError {
		t.Errorf("Updated completions missing 'error' value. Got: %v", testhelpers.GetCompletionLabels(updatedItems))

		// Debug: Check what's in the registry
		if attrs, exists := ctx.Attributes("test-alert"); exists {
			if stateAttr, hasState := attrs["state"]; hasState {
				t.Logf("Registry shows state attribute type: %s", stateAttr.Type.Text)
			}
		}

		// Debug: Check the manifest file content
		manifestContent, err := os.ReadFile(manifestPath)
		if err == nil {
			t.Logf("Current manifest content: %s", string(manifestContent))
		}
	} else {
		t.Logf("✅ SUCCESS: 'error' found in completions!")
	}
}
