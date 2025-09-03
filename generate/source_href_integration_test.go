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
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"testing/synctest"

	M "bennypowers.dev/cem/manifest"
	W "bennypowers.dev/cem/workspace"
)

// TestSourceHrefGeneration verifies that source hrefs are generated correctly when configured
func TestSourceHrefGeneration(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
	fixtureDir := filepath.Join("test", "fixtures", "project-source-hrefs")

	ctx := W.NewFileSystemWorkspaceContext(fixtureDir)
	if err := ctx.Init(); err != nil {
		t.Fatalf("Failed to initialize workspace: %v", err)
	}

	cfg, err := ctx.Config()
	if err != nil {
		t.Fatalf("Failed to get config: %v", err)
	}
	cfg.Generate.Files = []string{"src/source-hrefs.ts"}

	manifestJSON, err := Generate(ctx)
	if err != nil {
		t.Fatalf("Failed to generate manifest: %v", err)
	}

	var manifest M.Package
	if err := json.Unmarshal([]byte(*manifestJSON), &manifest); err != nil {
		t.Fatalf("Failed to parse manifest JSON: %v", err)
	}

	if len(manifest.Modules) == 0 {
		t.Fatal("Expected at least one module in manifest")
	}

	module := manifest.Modules[0]
	if len(module.Declarations) == 0 {
		t.Fatal("Expected at least one declaration in module")
	}

	// Verify declarations have source hrefs with correct format
	foundClass := false
	foundFunction := false
	foundVariable := false

	for _, decl := range module.Declarations {
		switch d := decl.(type) {
		case *M.ClassDeclaration:
			if d.Name == "TestClass" {
				foundClass = true
				if d.Source == nil || d.Source.Href == "" {
					t.Errorf("Expected source href for class %s, got nil", d.Name)
				} else {
					expectedPrefix := "https://github.com/example/repo/tree/main/src/source-hrefs.ts#L"
					if !strings.HasPrefix(d.Source.Href, expectedPrefix) {
						t.Errorf("Expected source href to start with %s, got: %s", expectedPrefix, d.Source.Href)
					}
				}
			}
		case *M.FunctionDeclaration:
			if d.Name == "testFunction" {
				foundFunction = true
				if d.Source == nil || d.Source.Href == "" {
					t.Errorf("Expected source href for function %s, got nil", d.Name)
				} else {
					expectedPrefix := "https://github.com/example/repo/tree/main/src/source-hrefs.ts#L"
					if !strings.HasPrefix(d.Source.Href, expectedPrefix) {
						t.Errorf("Expected source href to start with %s, got: %s", expectedPrefix, d.Source.Href)
					}
				}
			}
		case *M.VariableDeclaration:
			if d.Name == "testVariable" {
				foundVariable = true
				if d.Source == nil || d.Source.Href == "" {
					t.Errorf("Expected source href for variable %s, got nil", d.Name)
				} else {
					expectedPrefix := "https://github.com/example/repo/tree/main/src/source-hrefs.ts#L"
					if !strings.HasPrefix(d.Source.Href, expectedPrefix) {
						t.Errorf("Expected source href to start with %s, got: %s", expectedPrefix, d.Source.Href)
					}
				}
			}
		}
	}

	if !foundClass {
		t.Error("Expected to find TestClass declaration")
	}
	if !foundFunction {
		t.Error("Expected to find testFunction declaration")
	}
	if !foundVariable {
		t.Error("Expected to find testVariable declaration")
	}
	}) // End synctest.Test
}

// TestSourceHrefGenerationWithoutConfig verifies no source hrefs when not configured
func TestSourceHrefGenerationWithoutConfig(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
	// Create temporary fixture without sourceControlRootUrl
	tempDir := t.TempDir()
	configDir := filepath.Join(tempDir, ".config")
	srcDir := filepath.Join(tempDir, "src")

	if err := os.MkdirAll(configDir, 0755); err != nil {
		t.Fatalf("Failed to create config dir: %v", err)
	}
	if err := os.MkdirAll(srcDir, 0755); err != nil {
		t.Fatalf("Failed to create src dir: %v", err)
	}

	// Config without sourceControlRootUrl
	configContent := `files:
  - src/*.ts`
	if err := os.WriteFile(filepath.Join(configDir, "cem.yaml"), []byte(configContent), 0644); err != nil {
		t.Fatalf("Failed to write config: %v", err)
	}

	// Simple source file
	sourceContent := `export class TestClass {}`
	if err := os.WriteFile(filepath.Join(srcDir, "test.ts"), []byte(sourceContent), 0644); err != nil {
		t.Fatalf("Failed to write source: %v", err)
	}

	ctx := W.NewFileSystemWorkspaceContext(tempDir)
	if err := ctx.Init(); err != nil {
		t.Fatalf("Failed to initialize workspace: %v", err)
	}

	cfg, err := ctx.Config()
	if err != nil {
		t.Fatalf("Failed to get config: %v", err)
	}
	cfg.Generate.Files = []string{"src/test.ts"}

	manifestJSON, err := Generate(ctx)
	if err != nil {
		t.Fatalf("Failed to generate manifest: %v", err)
	}

	var manifest M.Package
	if err := json.Unmarshal([]byte(*manifestJSON), &manifest); err != nil {
		t.Fatalf("Failed to parse manifest JSON: %v", err)
	}

	// Verify no source hrefs are present
	for _, module := range manifest.Modules {
		for _, decl := range module.Declarations {
			switch d := decl.(type) {
			case *M.ClassDeclaration:
				if d.Source != nil {
					t.Errorf("Expected no source href for class %s when sourceControlRootUrl is not configured, but got: %s", d.Name, d.Source.Href)
				}
			case *M.FunctionDeclaration:
				if d.Source != nil {
					t.Errorf("Expected no source href for function %s when sourceControlRootUrl is not configured, but got: %s", d.Name, d.Source.Href)
				}
			case *M.VariableDeclaration:
				if d.Source != nil {
					t.Errorf("Expected no source href for variable %s when sourceControlRootUrl is not configured, but got: %s", d.Name, d.Source.Href)
				}
			}
		}
	}
	}) // End synctest.Test
}
