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

// TestImportMap_ConfigOverrideImportsOnly verifies config override with imports only
func TestImportMap_ConfigOverrideImportsOnly(t *testing.T) {
	tmpDir := t.TempDir()

	packageJSON := `{
  "name": "test-project",
  "dependencies": {
    "lit": "^3.0.0"
  }
}`
	err := os.WriteFile(filepath.Join(tmpDir, "package.json"), []byte(packageJSON), 0644)
	if err != nil {
		t.Fatalf("Failed to write package.json: %v", err)
	}

	nodeModules := filepath.Join(tmpDir, "node_modules", "lit")
	err = os.MkdirAll(nodeModules, 0755)
	if err != nil {
		t.Fatalf("Failed to create node_modules: %v", err)
	}

	litPackageJSON := `{
  "name": "lit",
  "exports": {
    ".": "./index.js"
  }
}`
	err = os.WriteFile(filepath.Join(nodeModules, "package.json"), []byte(litPackageJSON), 0644)
	if err != nil {
		t.Fatalf("Failed to write lit package.json: %v", err)
	}

	// Config override with imports only (no scopes)
	config := &Config{
		ConfigOverride: &ImportMap{
			Imports: map[string]string{
				"lit":  "https://cdn.jsdelivr.net/npm/lit@3/+esm",
				"lit/": "https://cdn.jsdelivr.net/npm/lit@3/",
			},
		},
	}

	importMap, err := Generate(tmpDir, config)
	if err != nil {
		t.Fatalf("Generate returned error: %v", err)
	}

	if importMap == nil {
		t.Fatal("Expected import map, got nil")
	}

	// Config override should win over auto-generated
	if importMap.Imports["lit"] != "https://cdn.jsdelivr.net/npm/lit@3/+esm" {
		t.Errorf("Expected config override for lit, got: %s", importMap.Imports["lit"])
	}
	if importMap.Imports["lit/"] != "https://cdn.jsdelivr.net/npm/lit@3/" {
		t.Errorf("Expected config override for lit/, got: %s", importMap.Imports["lit/"])
	}
}

// TestImportMap_ConfigOverrideScopesOnly verifies config override with scopes only
func TestImportMap_ConfigOverrideScopesOnly(t *testing.T) {
	tmpDir := t.TempDir()

	packageJSON := `{
  "name": "test-project",
  "dependencies": {
    "lit": "^3.0.0"
  }
}`
	err := os.WriteFile(filepath.Join(tmpDir, "package.json"), []byte(packageJSON), 0644)
	if err != nil {
		t.Fatalf("Failed to write package.json: %v", err)
	}

	nodeModules := filepath.Join(tmpDir, "node_modules", "lit")
	err = os.MkdirAll(nodeModules, 0755)
	if err != nil {
		t.Fatalf("Failed to create node_modules: %v", err)
	}

	litPackageJSON := `{
  "name": "lit",
  "exports": {
    ".": "./index.js"
  }
}`
	err = os.WriteFile(filepath.Join(nodeModules, "package.json"), []byte(litPackageJSON), 0644)
	if err != nil {
		t.Fatalf("Failed to write lit package.json: %v", err)
	}

	// Config override with scopes only (no imports override)
	config := &Config{
		ConfigOverride: &ImportMap{
			Scopes: map[string]map[string]string{
				"/demos/": {
					"lit": "/node_modules/lit/index.js",
				},
			},
		},
	}

	importMap, err := Generate(tmpDir, config)
	if err != nil {
		t.Fatalf("Generate returned error: %v", err)
	}

	if importMap == nil {
		t.Fatal("Expected import map, got nil")
	}

	// Should have auto-generated lit import
	if _, exists := importMap.Imports["lit"]; !exists {
		t.Error("Expected auto-generated lit import")
	}

	// Should have scoped override
	if importMap.Scopes == nil {
		t.Fatal("Expected scopes to be present")
	}
	if demoScope, exists := importMap.Scopes["/demos/"]; !exists {
		t.Error("Expected /demos/ scope")
	} else if demoScope["lit"] != "/node_modules/lit/index.js" {
		t.Errorf("Expected scoped override for lit in /demos/, got: %s", demoScope["lit"])
	}
}

// TestImportMap_ConfigOverrideBothImportsAndScopes verifies config override with both imports and scopes
func TestImportMap_ConfigOverrideBothImportsAndScopes(t *testing.T) {
	tmpDir := t.TempDir()

	packageJSON := `{
  "name": "test-project",
  "dependencies": {
    "lit": "^3.0.0",
    "react": "^18.0.0"
  }
}`
	err := os.WriteFile(filepath.Join(tmpDir, "package.json"), []byte(packageJSON), 0644)
	if err != nil {
		t.Fatalf("Failed to write package.json: %v", err)
	}

	// Create node_modules for lit and react
	for _, pkg := range []string{"lit", "react"} {
		nodeModules := filepath.Join(tmpDir, "node_modules", pkg)
		err = os.MkdirAll(nodeModules, 0755)
		if err != nil {
			t.Fatalf("Failed to create node_modules for %s: %v", pkg, err)
		}

		pkgJSON := `{
  "name": "` + pkg + `",
  "exports": {
    ".": "./index.js"
  }
}`
		err = os.WriteFile(filepath.Join(nodeModules, "package.json"), []byte(pkgJSON), 0644)
		if err != nil {
			t.Fatalf("Failed to write %s package.json: %v", pkg, err)
		}
	}

	// Config override with both imports and scopes
	config := &Config{
		ConfigOverride: &ImportMap{
			Imports: map[string]string{
				"react":  "https://esm.sh/react@18",
				"react/": "https://esm.sh/react@18/",
			},
			Scopes: map[string]map[string]string{
				"/demos/legacy/": {
					"react": "https://esm.sh/react@17",
				},
			},
		},
	}

	importMap, err := Generate(tmpDir, config)
	if err != nil {
		t.Fatalf("Generate returned error: %v", err)
	}

	if importMap == nil {
		t.Fatal("Expected import map, got nil")
	}

	// Should have auto-generated lit
	if _, exists := importMap.Imports["lit"]; !exists {
		t.Error("Expected auto-generated lit import")
	}

	// Config override should override react globally
	if importMap.Imports["react"] != "https://esm.sh/react@18" {
		t.Errorf("Expected global react override, got: %s", importMap.Imports["react"])
	}

	// Should have scoped override for legacy demos
	if importMap.Scopes == nil {
		t.Fatal("Expected scopes to be present")
	}
	if legacyScope, exists := importMap.Scopes["/demos/legacy/"]; !exists {
		t.Error("Expected /demos/legacy/ scope")
	} else if legacyScope["react"] != "https://esm.sh/react@17" {
		t.Errorf("Expected scoped react override in /demos/legacy/, got: %s", legacyScope["react"])
	}
}

// TestImportMap_ConfigOverridePriorityOverFile verifies config override beats file override
func TestImportMap_ConfigOverridePriorityOverFile(t *testing.T) {
	tmpDir := t.TempDir()

	packageJSON := `{
  "name": "test-project",
  "dependencies": {
    "lit": "^3.0.0"
  }
}`
	err := os.WriteFile(filepath.Join(tmpDir, "package.json"), []byte(packageJSON), 0644)
	if err != nil {
		t.Fatalf("Failed to write package.json: %v", err)
	}

	nodeModules := filepath.Join(tmpDir, "node_modules", "lit")
	err = os.MkdirAll(nodeModules, 0755)
	if err != nil {
		t.Fatalf("Failed to create node_modules: %v", err)
	}

	litPackageJSON := `{
  "name": "lit",
  "exports": {
    ".": "./index.js"
  }
}`
	err = os.WriteFile(filepath.Join(nodeModules, "package.json"), []byte(litPackageJSON), 0644)
	if err != nil {
		t.Fatalf("Failed to write lit package.json: %v", err)
	}

	// Create override file
	overrideFile := `{
  "imports": {
    "lit": "https://esm.sh/lit@3"
  }
}`
	overridePath := filepath.Join(tmpDir, "importmap.json")
	err = os.WriteFile(overridePath, []byte(overrideFile), 0644)
	if err != nil {
		t.Fatalf("Failed to write override file: %v", err)
	}

	// Config override should win over file override
	config := &Config{
		InputMapPath: overridePath,
		ConfigOverride: &ImportMap{
			Imports: map[string]string{
				"lit": "https://cdn.jsdelivr.net/npm/lit@3/+esm",
			},
		},
	}

	importMap, err := Generate(tmpDir, config)
	if err != nil {
		t.Fatalf("Generate returned error: %v", err)
	}

	if importMap == nil {
		t.Fatal("Expected import map, got nil")
	}

	// Config override should win (highest priority)
	if importMap.Imports["lit"] != "https://cdn.jsdelivr.net/npm/lit@3/+esm" {
		t.Errorf("Expected config override to win over file override, got: %s", importMap.Imports["lit"])
	}
}

// TestImportMap_ConfigOverrideDeepMergeScopes verifies deep merge of scopes
func TestImportMap_ConfigOverrideDeepMergeScopes(t *testing.T) {
	tmpDir := t.TempDir()

	packageJSON := `{
  "name": "test-project",
  "dependencies": {
    "lit": "^3.0.0",
    "react": "^18.0.0"
  }
}`
	err := os.WriteFile(filepath.Join(tmpDir, "package.json"), []byte(packageJSON), 0644)
	if err != nil {
		t.Fatalf("Failed to write package.json: %v", err)
	}

	// Create node_modules
	for _, pkg := range []string{"lit", "react"} {
		nodeModules := filepath.Join(tmpDir, "node_modules", pkg)
		err = os.MkdirAll(nodeModules, 0755)
		if err != nil {
			t.Fatalf("Failed to create node_modules for %s: %v", pkg, err)
		}

		pkgJSON := `{
  "name": "` + pkg + `",
  "exports": {
    ".": "./index.js"
  }
}`
		err = os.WriteFile(filepath.Join(nodeModules, "package.json"), []byte(pkgJSON), 0644)
		if err != nil {
			t.Fatalf("Failed to write %s package.json: %v", pkg, err)
		}
	}

	// Create override file with one scope entry
	overrideFile := `{
  "scopes": {
    "/demos/": {
      "lit": "/vendor/lit.js"
    }
  }
}`
	overridePath := filepath.Join(tmpDir, "importmap.json")
	err = os.WriteFile(overridePath, []byte(overrideFile), 0644)
	if err != nil {
		t.Fatalf("Failed to write override file: %v", err)
	}

	// Config override adds another entry to the same scope
	config := &Config{
		InputMapPath: overridePath,
		ConfigOverride: &ImportMap{
			Scopes: map[string]map[string]string{
				"/demos/": {
					"react": "/vendor/react.js",
				},
			},
		},
	}

	importMap, err := Generate(tmpDir, config)
	if err != nil {
		t.Fatalf("Generate returned error: %v", err)
	}

	if importMap == nil {
		t.Fatal("Expected import map, got nil")
	}

	// Should have merged scopes (both lit and react in /demos/ scope)
	if importMap.Scopes == nil {
		t.Fatal("Expected scopes to be present")
	}
	demoScope, exists := importMap.Scopes["/demos/"]
	if !exists {
		t.Fatal("Expected /demos/ scope")
	}

	// Should have both entries merged
	if demoScope["lit"] != "/vendor/lit.js" {
		t.Errorf("Expected lit from file override in /demos/ scope, got: %s", demoScope["lit"])
	}
	if demoScope["react"] != "/vendor/react.js" {
		t.Errorf("Expected react from config override in /demos/ scope, got: %s", demoScope["react"])
	}
}

// TestImportMap_ConfigOverrideReplacesFileScope verifies config override replaces conflicting scope entries
func TestImportMap_ConfigOverrideReplacesFileScope(t *testing.T) {
	tmpDir := t.TempDir()

	packageJSON := `{
  "name": "test-project"
}`
	err := os.WriteFile(filepath.Join(tmpDir, "package.json"), []byte(packageJSON), 0644)
	if err != nil {
		t.Fatalf("Failed to write package.json: %v", err)
	}

	// Create override file with scope entry
	overrideFile := `{
  "scopes": {
    "/demos/": {
      "lit": "https://esm.sh/lit@3"
    }
  }
}`
	overridePath := filepath.Join(tmpDir, "importmap.json")
	err = os.WriteFile(overridePath, []byte(overrideFile), 0644)
	if err != nil {
		t.Fatalf("Failed to write override file: %v", err)
	}

	// Config override for same key in same scope (should replace)
	config := &Config{
		InputMapPath: overridePath,
		ConfigOverride: &ImportMap{
			Scopes: map[string]map[string]string{
				"/demos/": {
					"lit": "https://cdn.jsdelivr.net/npm/lit@3/+esm",
				},
			},
		},
	}

	importMap, err := Generate(tmpDir, config)
	if err != nil {
		t.Fatalf("Generate returned error: %v", err)
	}

	if importMap == nil {
		t.Fatal("Expected import map, got nil")
	}

	// Config override should replace file override in scope
	if importMap.Scopes == nil {
		t.Fatal("Expected scopes to be present")
	}
	demoScope, exists := importMap.Scopes["/demos/"]
	if !exists {
		t.Fatal("Expected /demos/ scope")
	}

	if demoScope["lit"] != "https://cdn.jsdelivr.net/npm/lit@3/+esm" {
		t.Errorf("Expected config override to replace file override in scope, got: %s", demoScope["lit"])
	}
}

// TestImportMap_ConfigOverrideEmptyDoesNotCrash verifies nil/empty config override doesn't crash
func TestImportMap_ConfigOverrideEmptyDoesNotCrash(t *testing.T) {
	tmpDir := t.TempDir()

	packageJSON := `{
  "name": "test-project",
  "dependencies": {
    "lit": "^3.0.0"
  }
}`
	err := os.WriteFile(filepath.Join(tmpDir, "package.json"), []byte(packageJSON), 0644)
	if err != nil {
		t.Fatalf("Failed to write package.json: %v", err)
	}

	nodeModules := filepath.Join(tmpDir, "node_modules", "lit")
	err = os.MkdirAll(nodeModules, 0755)
	if err != nil {
		t.Fatalf("Failed to create node_modules: %v", err)
	}

	litPackageJSON := `{
  "name": "lit",
  "exports": {
    ".": "./index.js"
  }
}`
	err = os.WriteFile(filepath.Join(nodeModules, "package.json"), []byte(litPackageJSON), 0644)
	if err != nil {
		t.Fatalf("Failed to write lit package.json: %v", err)
	}

	// Test with nil ConfigOverride
	config := &Config{
		ConfigOverride: nil,
	}

	importMap, err := Generate(tmpDir, config)
	if err != nil {
		t.Fatalf("Generate returned error with nil ConfigOverride: %v", err)
	}

	if importMap == nil {
		t.Fatal("Expected import map, got nil")
	}

	// Should have auto-generated lit
	if _, exists := importMap.Imports["lit"]; !exists {
		t.Error("Expected auto-generated lit import with nil ConfigOverride")
	}

	// Test with empty ConfigOverride
	config2 := &Config{
		ConfigOverride: &ImportMap{
			Imports: map[string]string{},
			Scopes:  map[string]map[string]string{},
		},
	}

	importMap2, err := Generate(tmpDir, config2)
	if err != nil {
		t.Fatalf("Generate returned error with empty ConfigOverride: %v", err)
	}

	if importMap2 == nil {
		t.Fatal("Expected import map, got nil")
	}

	// Should have auto-generated lit
	if _, exists := importMap2.Imports["lit"]; !exists {
		t.Error("Expected auto-generated lit import with empty ConfigOverride")
	}
}

// TestImportMap_ConfigOverrideWithNoPackageJSON verifies config override works without package.json
func TestImportMap_ConfigOverrideWithNoPackageJSON(t *testing.T) {
	tmpDir := t.TempDir()

	// No package.json - only config override
	config := &Config{
		ConfigOverride: &ImportMap{
			Imports: map[string]string{
				"lit":    "https://cdn.jsdelivr.net/npm/lit@3/+esm",
				"react":  "https://esm.sh/react@18",
				"react/": "https://esm.sh/react@18/",
			},
			Scopes: map[string]map[string]string{
				"/demos/": {
					"lit": "/vendor/lit.js",
				},
			},
		},
	}

	importMap, err := Generate(tmpDir, config)
	if err != nil {
		t.Fatalf("Generate returned error: %v", err)
	}

	if importMap == nil {
		t.Fatal("Expected import map, got nil")
	}

	// Should have config overrides only
	if importMap.Imports["lit"] != "https://cdn.jsdelivr.net/npm/lit@3/+esm" {
		t.Errorf("Expected config override for lit, got: %s", importMap.Imports["lit"])
	}
	if importMap.Imports["react"] != "https://esm.sh/react@18" {
		t.Errorf("Expected config override for react, got: %s", importMap.Imports["react"])
	}

	// Should have scopes
	if importMap.Scopes == nil {
		t.Fatal("Expected scopes to be present")
	}
	if demoScope, exists := importMap.Scopes["/demos/"]; !exists {
		t.Error("Expected /demos/ scope")
	} else if demoScope["lit"] != "/vendor/lit.js" {
		t.Errorf("Expected scoped lit override, got: %s", demoScope["lit"])
	}
}

// TestImportMap_ConfigOverrideValidation verifies control characters are rejected
func TestImportMap_ConfigOverrideValidation(t *testing.T) {
	tmpDir := t.TempDir()

	packageJSON := `{
  "name": "test-project"
}`
	err := os.WriteFile(filepath.Join(tmpDir, "package.json"), []byte(packageJSON), 0644)
	if err != nil {
		t.Fatalf("Failed to write package.json: %v", err)
	}

	warnings := []string{}
	config := &Config{
		Logger: &testLogger{warnings: &warnings},
		ConfigOverride: &ImportMap{
			Imports: map[string]string{
				"valid-key":           "/valid/path.js",
				"key\nwith\nnewline":  "/bad/path.js",        // Should be rejected
				"key\x00with\x00null": "/bad/path2.js",       // Should be rejected
				"another-valid":       "https://example.com", // Should be kept
			},
			Scopes: map[string]map[string]string{
				"/valid-scope/": {
					"lit":                   "/vendor/lit.js",
					"react\rwith\rcarriage": "/bad/react.js", // Should be rejected
				},
				"/scope\nwith\nnewline/": { // Scope key should be rejected
					"lit": "/vendor/lit2.js",
				},
			},
		},
	}

	importMap, err := Generate(tmpDir, config)
	if err != nil {
		t.Fatalf("Generate returned error: %v", err)
	}

	if importMap == nil {
		t.Fatal("Expected import map, got nil")
	}

	// Valid keys should be present
	if _, exists := importMap.Imports["valid-key"]; !exists {
		t.Error("Expected valid-key to be present")
	}
	if _, exists := importMap.Imports["another-valid"]; !exists {
		t.Error("Expected another-valid to be present")
	}

	// Invalid keys should be rejected
	if _, exists := importMap.Imports["key\nwith\nnewline"]; exists {
		t.Error("Expected key with newline to be rejected")
	}
	if _, exists := importMap.Imports["key\x00with\x00null"]; exists {
		t.Error("Expected key with null bytes to be rejected")
	}

	// Valid scope should be present with valid import
	if validScope, exists := importMap.Scopes["/valid-scope/"]; !exists {
		t.Error("Expected /valid-scope/ to be present")
	} else {
		if _, exists := validScope["lit"]; !exists {
			t.Error("Expected lit in /valid-scope/")
		}
		if _, exists := validScope["react\rwith\rcarriage"]; exists {
			t.Error("Expected import key with carriage return to be rejected")
		}
	}

	// Invalid scope key should be rejected
	if _, exists := importMap.Scopes["/scope\nwith\nnewline/"]; exists {
		t.Error("Expected scope key with newline to be rejected")
	}

	// Should have logged warnings
	if len(warnings) == 0 {
		t.Error("Expected warnings to be logged for invalid keys")
	}

	// Verify specific warnings were logged
	foundInvalidImport := false
	foundInvalidScope := false
	for _, w := range warnings {
		if strings.Contains(w, "control characters") {
			if strings.Contains(w, "import map key") {
				foundInvalidImport = true
			}
			if strings.Contains(w, "scope key") {
				foundInvalidScope = true
			}
		}
	}
	if !foundInvalidImport {
		t.Errorf("Expected warning about invalid import key, got warnings: %v", warnings)
	}
	if !foundInvalidScope {
		t.Errorf("Expected warning about invalid scope key, got warnings: %v", warnings)
	}
}
