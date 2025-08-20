package definition_test

import (
	"os"
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/internal/platform"
	"bennypowers.dev/cem/lsp/methods/textDocument/definition"
	"bennypowers.dev/cem/lsp/testhelpers"
)

func TestDefinition_PathResolutionDebug(t *testing.T) {
	// Test to debug CI path resolution issues
	workspaceRoot := "/home/bennyp/Developer/cem/lsp/methods/textDocument/definition/definition-test-fixtures"

	// Check if workspace root exists
	if _, err := os.Stat(workspaceRoot); err != nil {
		t.Errorf("Workspace root does not exist: %v", err)
		return
	}

	// Check if components directory exists
	componentsDir := filepath.Join(workspaceRoot, "components")
	if _, err := os.Stat(componentsDir); err != nil {
		t.Errorf("Components directory does not exist: %v", err)
		return
	}

	// Check if .ts file exists
	tsFile := filepath.Join(componentsDir, "card-element.ts")
	if _, err := os.Stat(tsFile); err != nil {
		t.Errorf("TypeScript file does not exist: %v", err)
		// List what files ARE in the directory
		entries, listErr := os.ReadDir(componentsDir)
		if listErr == nil {
			t.Logf("Files in components directory:")
			for _, entry := range entries {
				t.Logf("  - %s", entry.Name())
			}
		}
	} else {
		t.Logf("✅ TypeScript file exists: %s", tsFile)
	}

	// Check if .js file exists
	jsFile := filepath.Join(componentsDir, "card-element.js")
	if _, err := os.Stat(jsFile); err != nil {
		t.Logf("JavaScript file does not exist: %v", err)
	} else {
		t.Logf("⚠️  JavaScript file exists: %s", jsFile)
	}

	// Test the actual path resolution using the exported function
	mockDef := &testhelpers.MockElementDefinition{
		ModulePathStr: "components/card-element.js", // This is what the manifest has
	}

	resolved := definition.ResolveSourcePathForTesting(mockDef, workspaceRoot)
	t.Logf("Path resolution result: %s", resolved)

	// Check what the OS filesystem would find
	fs := platform.NewOSFileSystem()
	localPath := filepath.Join(workspaceRoot, "components/card-element.js")
	t.Logf("Checking local path: %s", localPath)

	if _, err := fs.Stat(localPath); err != nil {
		t.Logf("JS file stat failed: %v", err)
	} else {
		t.Logf("JS file stat succeeded")
	}

	// Try .ts version
	tsPath := filepath.Join(workspaceRoot, "components/card-element.ts")
	t.Logf("Checking TS path: %s", tsPath)

	if _, err := fs.Stat(tsPath); err != nil {
		t.Logf("TS file stat failed: %v", err)
	} else {
		t.Logf("TS file stat succeeded")
	}
}
