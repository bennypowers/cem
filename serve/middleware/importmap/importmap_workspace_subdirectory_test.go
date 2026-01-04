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

package importmap

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestImportMap_WorkspaceSubdirectoryPaths is a regression test for issue #192
// Verifies that when running from a workspace subdirectory, import map paths
// are generated correctly without invalid "/../../" prefixes
func TestImportMap_WorkspaceSubdirectoryPaths(t *testing.T) {
	// Create workspace root directory structure
	workspaceRoot := t.TempDir()

	// Create workspace root package.json with workspaces
	workspacePackageJSON := `{
  "name": "monorepo",
  "workspaces": [
    "examples/*"
  ]
}`
	err := os.WriteFile(filepath.Join(workspaceRoot, "package.json"), []byte(workspacePackageJSON), 0644)
	if err != nil {
		t.Fatalf("Failed to write workspace package.json: %v", err)
	}

	// Create workspace root node_modules with lit package
	litPath := filepath.Join(workspaceRoot, "node_modules", "lit")
	err = os.MkdirAll(litPath, 0755)
	if err != nil {
		t.Fatalf("Failed to create lit package: %v", err)
	}

	litPackageJSON := `{
  "name": "lit",
  "version": "3.0.0",
  "exports": {
    ".": {
      "import": "./index.js"
    },
    "./decorators.js": {
      "import": "./decorators.js"
    }
  }
}`
	err = os.WriteFile(filepath.Join(litPath, "package.json"), []byte(litPackageJSON), 0644)
	if err != nil {
		t.Fatalf("Failed to write lit package.json: %v", err)
	}

	// Create @lit/reactive-element (scoped package)
	reactiveElementPath := filepath.Join(workspaceRoot, "node_modules", "@lit", "reactive-element")
	err = os.MkdirAll(reactiveElementPath, 0755)
	if err != nil {
		t.Fatalf("Failed to create @lit/reactive-element package: %v", err)
	}

	// Include lit as a dependency of @lit/reactive-element (for transitive dependency testing)
	reactiveElementPackageJSON := `{
  "name": "@lit/reactive-element",
  "version": "2.0.0",
  "dependencies": {
    "lit": "^3.0.0"
  },
  "exports": {
    ".": "./reactive-element.js",
    "./decorators/custom-element.js": "./decorators/custom-element.js"
  }
}`
	err = os.WriteFile(filepath.Join(reactiveElementPath, "package.json"), []byte(reactiveElementPackageJSON), 0644)
	if err != nil {
		t.Fatalf("Failed to write reactive-element package.json: %v", err)
	}

	// Create workspace package subdirectory (examples/kitchen-sink)
	kitchenSinkPath := filepath.Join(workspaceRoot, "examples", "kitchen-sink")
	err = os.MkdirAll(kitchenSinkPath, 0755)
	if err != nil {
		t.Fatalf("Failed to create kitchen-sink directory: %v", err)
	}

	kitchenSinkPackageJSON := `{
  "name": "@examples/kitchen-sink",
  "version": "1.0.0",
  "dependencies": {
    "lit": "^3.0.0",
    "@lit/reactive-element": "^2.0.0"
  }
}`
	err = os.WriteFile(filepath.Join(kitchenSinkPath, "package.json"), []byte(kitchenSinkPackageJSON), 0644)
	if err != nil {
		t.Fatalf("Failed to write kitchen-sink package.json: %v", err)
	}

	// Generate import map from the subdirectory (simulating `cem serve` from examples/kitchen-sink)
	debugs := []string{}
	config := &Config{
		Logger: &testLogger{debugs: &debugs},
	}

	importMap, err := Generate(kitchenSinkPath, config)
	if err != nil {
		t.Fatalf("Generate returned error: %v", err)
	}

	if importMap == nil {
		t.Fatal("Expected import map, got nil")
	}

	// Verify top-level imports do NOT contain invalid "/../../" paths
	// They should use "/node_modules/" paths since workspaceRoot should be detected
	t.Run("TopLevelImports", func(t *testing.T) {
		litImport := importMap.Imports["lit"]
		if litImport == "" {
			t.Fatal("Expected lit import to exist")
		}
		if strings.Contains(litImport, "/../") {
			t.Errorf("lit import contains invalid /../ path traversal: %s", litImport)
		}
		if !strings.HasPrefix(litImport, "/node_modules/") {
			t.Errorf("Expected lit import to start with /node_modules/, got: %s", litImport)
		}
		expectedLitPath := "/node_modules/lit/index.js"
		if litImport != expectedLitPath {
			t.Errorf("Expected lit import to be %s, got: %s", expectedLitPath, litImport)
		}

		reactiveElementImport := importMap.Imports["@lit/reactive-element"]
		if reactiveElementImport == "" {
			t.Fatal("Expected @lit/reactive-element import to exist")
		}
		if strings.Contains(reactiveElementImport, "/../") {
			t.Errorf("@lit/reactive-element import contains invalid /../ path traversal: %s", reactiveElementImport)
		}
		if !strings.HasPrefix(reactiveElementImport, "/node_modules/") {
			t.Errorf("Expected @lit/reactive-element import to start with /node_modules/, got: %s", reactiveElementImport)
		}
	})

	// Verify scopes do NOT contain invalid "/../../" paths
	t.Run("ScopePaths", func(t *testing.T) {
		for scopeKey, scopeImports := range importMap.Scopes {
			// Scope keys should not contain /../
			if strings.Contains(scopeKey, "/../") {
				t.Errorf("Scope key contains invalid /../ path traversal: %s", scopeKey)
			}

			// Scope import values should not contain /../
			for importKey, importValue := range scopeImports {
				if strings.Contains(importValue, "/../") {
					t.Errorf("Scope import value for %s in scope %s contains invalid /../ path traversal: %s",
						importKey, scopeKey, importValue)
				}
				// All scoped imports should use /node_modules/ paths
				if !strings.HasPrefix(importValue, "/node_modules/") {
					t.Errorf("Expected scope import %s in scope %s to start with /node_modules/, got: %s",
						importKey, scopeKey, importValue)
				}
			}
		}
	})

	// Verify workspace root was detected
	t.Run("WorkspaceRootDetection", func(t *testing.T) {
		foundDebugMsg := false
		for _, msg := range debugs {
			if strings.Contains(msg, "Detected workspace subdirectory") {
				foundDebugMsg = true
				break
			}
		}
		if !foundDebugMsg {
			t.Logf("Debug messages: %v", debugs)
			t.Error("Expected debug message about workspace subdirectory detection")
		}
	})
}

// TestImportMap_WorkspaceSubdirectoryTransitiveDeps verifies transitive dependencies
// in scopes are correctly generated when running from a workspace subdirectory
func TestImportMap_WorkspaceSubdirectoryTransitiveDeps(t *testing.T) {
	workspaceRoot := t.TempDir()

	// Create workspace root package.json
	workspacePackageJSON := `{
  "name": "monorepo",
  "workspaces": ["packages/*"]
}`
	err := os.WriteFile(filepath.Join(workspaceRoot, "package.json"), []byte(workspacePackageJSON), 0644)
	if err != nil {
		t.Fatalf("Failed to write workspace package.json: %v", err)
	}

	// Create lit package
	litPath := filepath.Join(workspaceRoot, "node_modules", "lit")
	err = os.MkdirAll(litPath, 0755)
	if err != nil {
		t.Fatalf("Failed to create lit: %v", err)
	}
	err = os.WriteFile(filepath.Join(litPath, "package.json"), []byte(`{
  "name": "lit",
  "exports": { ".": "./index.js" }
}`), 0644)
	if err != nil {
		t.Fatalf("Failed to write lit package.json: %v", err)
	}

	// Create @lit/reactive-element with dependency on lit
	reactiveElementPath := filepath.Join(workspaceRoot, "node_modules", "@lit", "reactive-element")
	err = os.MkdirAll(reactiveElementPath, 0755)
	if err != nil {
		t.Fatalf("Failed to create reactive-element: %v", err)
	}
	err = os.WriteFile(filepath.Join(reactiveElementPath, "package.json"), []byte(`{
  "name": "@lit/reactive-element",
  "dependencies": {
    "lit": "^3.0.0"
  },
  "exports": { ".": "./reactive-element.js" }
}`), 0644)
	if err != nil {
		t.Fatalf("Failed to write reactive-element package.json: %v", err)
	}

	// Create workspace package
	pkgPath := filepath.Join(workspaceRoot, "packages", "my-component")
	err = os.MkdirAll(pkgPath, 0755)
	if err != nil {
		t.Fatalf("Failed to create package: %v", err)
	}
	err = os.WriteFile(filepath.Join(pkgPath, "package.json"), []byte(`{
  "name": "@workspace/my-component",
  "dependencies": {
    "@lit/reactive-element": "^2.0.0"
  }
}`), 0644)
	if err != nil {
		t.Fatalf("Failed to write package.json: %v", err)
	}

	// Generate from the workspace package subdirectory
	importMap, err := Generate(pkgPath, nil)
	if err != nil {
		t.Fatalf("Generate returned error: %v", err)
	}

	// Verify scopes for transitive dependencies exist and have valid paths
	scopeKey := "/node_modules/@lit/reactive-element/"
	if importMap.Scopes[scopeKey] == nil {
		t.Fatalf("Expected scope for @lit/reactive-element")
	}

	litInScope := importMap.Scopes[scopeKey]["lit"]
	if litInScope == "" {
		t.Fatal("Expected lit to be in @lit/reactive-element scope (transitive dependency)")
	}

	// Should NOT contain /../
	if strings.Contains(litInScope, "/../") {
		t.Errorf("Transitive dependency 'lit' in scope contains invalid /../: %s", litInScope)
	}

	// Should use /node_modules/ path
	if !strings.HasPrefix(litInScope, "/node_modules/") {
		t.Errorf("Expected transitive dependency to use /node_modules/ path, got: %s", litInScope)
	}
}
