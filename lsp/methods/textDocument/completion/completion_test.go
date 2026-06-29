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
	"os"
	"path/filepath"
	"strings"
	"testing"

	G "bennypowers.dev/cem/generate"
	"bennypowers.dev/cem/internal/platform"
	testworkspace "bennypowers.dev/cem/internal/platform/testutil/workspace"
	"bennypowers.dev/cem/lsp"
	"bennypowers.dev/cem/lsp/document"
	"bennypowers.dev/cem/lsp/methods/textDocument/completion"
	"bennypowers.dev/cem/lsp/testhelpers"
	M "bennypowers.dev/cem/manifest"
	W "bennypowers.dev/cem/internal/workspace"
)

// TestServerLevelIntegration tests the complete server integration including
// manifest reloading behavior exactly as it would work in real usage
func TestServerLevelIntegration(t *testing.T) {
	fixtureDir := filepath.Join("testdata", "integration", "server-integration")

	t.Run("os", func(t *testing.T) {
		// Create a temporary workspace directory that mimics a real project
		tempDir := t.TempDir()

		// Copy fixture files to temporary directory
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
		osWorkspace := W.NewFileSystemWorkspaceContext(tempDir)
		err = osWorkspace.Init()
		if err != nil {
			t.Fatalf("Failed to initialize workspace: %v", err)
		}

		// Run initial generation to create the manifest
		generateSession, err := G.NewGenerateSession(osWorkspace, platform.NewOSFileSystem())
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
		server, err := lsp.NewServer(osWorkspace, lsp.TransportStdio)
		if err != nil {
			t.Fatalf("Failed to create LSP server: %v", err)
		}
		defer func() { _ = server.Close() }()

		// Initialize the server (this loads manifests and starts file watching)
		err = server.InitializeForTesting()
		if err != nil {
			t.Fatalf("Failed to initialize server: %v", err)
		}

		ctx := buildMockServerContext(t, server.Registry())

		htmlFilePath := filepath.Join(tempDir, "index.html")
		htmlContentBytes, err := os.ReadFile(htmlFilePath)
		if err != nil {
			t.Fatalf("Failed to read HTML file: %v", err)
		}
		htmlContent := string(htmlContentBytes)

		htmlURI := "file://" + htmlFilePath
		dm, dmErr := ctx.DocumentManager()
		if dmErr != nil {
			t.Fatalf("Failed to get document manager: %v", dmErr)
		}
		doc := dm.OpenDocument(htmlURI, htmlContent, 1)
		if doc == nil {
			t.Fatalf("Failed to open HTML document")
		}

		// Test initial completions - should have info, success, warning
		verifyInitialCompletions(t, ctx)

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

		// Verify the file was actually updated on disk
		verifyContent, err := os.ReadFile(tsFilePath)
		if err != nil {
			t.Fatalf("Failed to read back updated file: %v", err)
		}
		if !strings.Contains(string(verifyContent), "error") {
			t.Fatalf("File was not properly updated - missing 'error' in content: %s", string(verifyContent))
		}
		t.Logf("Verified file was updated on disk with 'error' in union type")

		t.Logf("File updated, manually generating updated manifest...")

		// Instead of waiting for the file watcher, manually generate the updated manifest
		genSession, err := G.NewGenerateSession(osWorkspace, platform.NewOSFileSystem())
		if err != nil {
			t.Fatalf("Failed to create generate session: %v", err)
		}
		defer genSession.Close()

		updatedManifest, err := genSession.GenerateFullManifest(context.Background())
		if err != nil {
			t.Fatalf("Failed to generate updated manifest: %v", err)
		}

		ctx.AddManifest(updatedManifest)
		t.Logf("Manually updated registry with new manifest")

		// Test updated completions - should now include 'error'
		verifyUpdatedCompletions(t, ctx)
	})

	t.Run("mapfs", func(t *testing.T) {
		// Create workspace from MapFS -- loads fixture files into MapFS rooted at "/"
		mapWorkspace := testworkspace.NewMapWorkspaceContext(t, fixtureDir)

		// Create dist directory in MapFS
		err := mapWorkspace.FileSystem().MkdirAll("dist", 0755)
		if err != nil {
			t.Fatalf("Failed to create dist directory in MapFS: %v", err)
		}
		err = mapWorkspace.Init()
		if err != nil {
			t.Fatalf("Failed to initialize workspace: %v", err)
		}

		// Run initial generation using the MapFS
		generateSession, err := G.NewGenerateSession(mapWorkspace, mapWorkspace.FileSystem())
		if err != nil {
			t.Fatalf("Failed to create generate session: %v", err)
		}
		defer generateSession.Close()

		pkg, err := generateSession.GenerateFullManifest(context.Background())
		if err != nil {
			t.Fatalf("Failed to generate initial manifest: %v", err)
		}

		// Write initial manifest to MapFS
		manifestStr, err := M.SerializeToString(pkg)
		if err != nil {
			t.Fatalf("Failed to serialize manifest: %v", err)
		}

		err = mapWorkspace.FileSystem().WriteFile("dist/custom-elements.json", []byte(manifestStr), 0644)
		if err != nil {
			t.Fatalf("Failed to write manifest to MapFS: %v", err)
		}

		// Create LSP server with the MapWorkspaceContext
		server, err := lsp.NewServer(mapWorkspace, lsp.TransportStdio)
		if err != nil {
			t.Fatalf("Failed to create LSP server: %v", err)
		}
		defer func() { _ = server.Close() }()

		// Initialize the server (file watching may warn but won't fail)
		err = server.InitializeForTesting()
		if err != nil {
			t.Fatalf("Failed to initialize server: %v", err)
		}

		ctx := buildMockServerContext(t, server.Registry())

		// Read HTML content from MapFS
		htmlContentBytes, err := mapWorkspace.FileSystem().ReadFile("index.html")
		if err != nil {
			t.Fatalf("Failed to read HTML file from MapFS: %v", err)
		}
		htmlContent := string(htmlContentBytes)

		htmlURI := "file:///index.html"
		dm, dmErr := ctx.DocumentManager()
		if dmErr != nil {
			t.Fatalf("Failed to get document manager: %v", dmErr)
		}
		doc := dm.OpenDocument(htmlURI, htmlContent, 1)
		if doc == nil {
			t.Fatalf("Failed to open HTML document")
		}

		// Test initial completions - should have info, success, warning
		verifyInitialCompletions(t, ctx)

		// Now simulate the user editing the TypeScript file to add 'error'
		t.Logf("=== User edits test-alert.ts to add 'error' ===")

		// Read the updated TS content from MapFS (loaded as part of fixtures)
		updatedTSContent, err := mapWorkspace.FileSystem().ReadFile("src/test-alert-updated.ts")
		if err != nil {
			t.Fatalf("Failed to read updated TypeScript file from MapFS: %v", err)
		}

		// Write the updated content over the original file in MapFS
		err = mapWorkspace.FileSystem().WriteFile("src/test-alert.ts", updatedTSContent, 0644)
		if err != nil {
			t.Fatalf("Failed to update TypeScript file in MapFS: %v", err)
		}

		// Verify the file was actually updated in MapFS
		verifyContent, err := mapWorkspace.FileSystem().ReadFile("src/test-alert.ts")
		if err != nil {
			t.Fatalf("Failed to read back updated file from MapFS: %v", err)
		}
		if !strings.Contains(string(verifyContent), "error") {
			t.Fatalf("File was not properly updated - missing 'error' in content: %s", string(verifyContent))
		}
		t.Logf("Verified file was updated in MapFS with 'error' in union type")

		t.Logf("File updated, manually generating updated manifest...")

		// Generate updated manifest using MapFS
		genSession, err := G.NewGenerateSession(mapWorkspace, mapWorkspace.FileSystem())
		if err != nil {
			t.Fatalf("Failed to create generate session: %v", err)
		}
		defer genSession.Close()

		updatedManifest, err := genSession.GenerateFullManifest(context.Background())
		if err != nil {
			t.Fatalf("Failed to generate updated manifest: %v", err)
		}

		ctx.AddManifest(updatedManifest)
		t.Logf("Manually updated registry with new manifest")

		// Test updated completions - should now include 'error'
		verifyUpdatedCompletions(t, ctx)
	})
}

// buildMockServerContext creates a MockServerContext populated from the given registry.
func buildMockServerContext(t *testing.T, registry *lsp.Registry) *testhelpers.MockServerContext {
	t.Helper()

	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create document manager: %v", err)
	}
	t.Cleanup(func() { dm.Close() })

	ctx := testhelpers.NewMockServerContext()
	for _, tagName := range registry.AllTagNames() {
		if element, exists := registry.Element(tagName); exists {
			ctx.AddElement(tagName, element)
		}
		if attrs, exists := registry.Attributes(tagName); exists {
			ctx.AddAttributes(tagName, attrs)
		}
		if slots, exists := registry.Slots(tagName); exists {
			ctx.AddSlots(tagName, slots)
		}
	}
	ctx.SetDocumentManager(dm)
	return ctx
}

// verifyInitialCompletions checks that the initial completions contain info, success, and warning.
func verifyInitialCompletions(t *testing.T, ctx *testhelpers.MockServerContext) {
	t.Helper()

	items := completion.GetAttributeValueCompletions(ctx, "test-alert", "state")
	t.Logf("Initial completions: %v", testhelpers.GetCompletionLabels(items))

	hasInfo := false
	hasSuccess := false
	hasWarning := false
	for _, item := range items {
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
		t.Fatalf("Initial completions missing expected values. Got: %v", testhelpers.GetCompletionLabels(items))
	}
}

// verifyUpdatedCompletions checks that updated completions contain info, success, warning, and error.
func verifyUpdatedCompletions(t *testing.T, ctx *testhelpers.MockServerContext) {
	t.Helper()

	items := completion.GetAttributeValueCompletions(ctx, "test-alert", "state")
	t.Logf("Updated completions: %v", testhelpers.GetCompletionLabels(items))

	hasInfo := false
	hasSuccess := false
	hasWarning := false
	hasError := false
	for _, item := range items {
		switch item.Label {
		case "info":
			hasInfo = true
		case "success":
			hasSuccess = true
		case "warning":
			hasWarning = true
		case "error":
			hasError = true
		}
	}

	if !hasInfo || !hasSuccess || !hasWarning {
		t.Errorf("Updated completions missing original values. Got: %v", testhelpers.GetCompletionLabels(items))
	}

	if !hasError {
		t.Errorf("Updated completions missing 'error' value. Got: %v", testhelpers.GetCompletionLabels(items))

		if attrs, exists := ctx.Attributes("test-alert"); exists {
			if stateAttr, hasState := attrs["state"]; hasState {
				t.Logf("Registry shows state attribute type: %s", stateAttr.Type.Text)
			}
		}
	} else {
		t.Logf("SUCCESS: 'error' found in completions!")
	}
}
