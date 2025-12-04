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
	"testing"
)

// TestImportMap_SinglePackageWithExportsMap tests export map resolution for a single package repo
func TestImportMap_SinglePackageWithExportsMap(t *testing.T) {
	tmpDir := t.TempDir()

	// Single package with complex exports
	packageJSON := `{
		"name": "my-components",
		"version": "1.0.0",
		"exports": {
			".": "./index.js",
			"./lib/*": "./lib/*",
			"./components/*": "./components/*"
		}
	}`

	if err := os.WriteFile(filepath.Join(tmpDir, "package.json"), []byte(packageJSON), 0644); err != nil {
		t.Fatalf("Failed to write package.json: %v", err)
	}

	importMap, err := Generate(tmpDir, nil)
	if err != nil {
		t.Fatalf("GenerateImportMap failed: %v", err)
	}

	if importMap == nil {
		t.Fatal("Expected import map, got nil")
	}

	// Should have main export
	if mainPath, ok := importMap.Imports["my-components"]; !ok {
		t.Error("Expected 'my-components' in import map")
	} else if mainPath != "/index.js" {
		t.Errorf("Expected '/index.js', got %s", mainPath)
	}

	// Should NOT have broad trailing slash mapping
	if _, exists := importMap.Imports["my-components/"]; exists {
		t.Error("Should not have root 'my-components/' mapping - would incorrectly expose all package subpaths")
	}

	// Should have wildcard pattern mappings for declared subpaths
	if libPath, ok := importMap.Imports["my-components/lib/"]; !ok {
		t.Error("Expected 'my-components/lib/' wildcard mapping")
	} else if libPath != "/lib/" {
		t.Errorf("Expected '/lib/', got %s", libPath)
	}

	if componentsPath, ok := importMap.Imports["my-components/components/"]; !ok {
		t.Error("Expected 'my-components/components/' wildcard mapping")
	} else if componentsPath != "/components/" {
		t.Errorf("Expected '/components/', got %s", componentsPath)
	}

	// Should have main export + 2 wildcard mappings
	if len(importMap.Imports) != 3 {
		t.Errorf("Expected 3 import map entries (main + 2 wildcards), got %d: %v",
			len(importMap.Imports), importMap.Imports)
	}
}

// TestImportMap_WorkspacePackageWithExportsMap tests workspace package export resolution
func TestImportMap_WorkspacePackageWithExportsMap(t *testing.T) {
	tmpDir := t.TempDir()

	// Root package.json with workspaces
	rootPackageJSON := `{
		"name": "monorepo-root",
		"workspaces": ["packages/*"]
	}`

	if err := os.WriteFile(filepath.Join(tmpDir, "package.json"), []byte(rootPackageJSON), 0644); err != nil {
		t.Fatalf("Failed to write package.json: %v", err)
	}

	// Create workspace package with exports
	workspaceDir := filepath.Join(tmpDir, "packages", "components")
	if err := os.MkdirAll(workspaceDir, 0755); err != nil {
		t.Fatalf("Failed to create workspace directory: %v", err)
	}

	workspacePackageJSON := `{
		"name": "@myorg/components",
		"version": "1.0.0",
		"exports": {
			".": "./index.js",
			"./button": "./elements/button.js",
			"./alert": "./elements/alert.js",
			"./elements/*": "./elements/*"
		}
	}`

	if err := os.WriteFile(filepath.Join(workspaceDir, "package.json"), []byte(workspacePackageJSON), 0644); err != nil {
		t.Fatalf("Failed to write workspace package.json: %v", err)
	}

	importMap, err := Generate(tmpDir, nil)
	if err != nil {
		t.Fatalf("GenerateImportMap failed: %v", err)
	}

	if importMap == nil {
		t.Fatal("Expected import map, got nil")
	}

	// Should have main export
	if mainPath, ok := importMap.Imports["@myorg/components"]; !ok {
		t.Error("Expected '@myorg/components' in import map")
	} else if mainPath != "/packages/components/index.js" {
		t.Errorf("Expected '/packages/components/index.js', got %s", mainPath)
	}

	// Should have named exports
	if buttonPath, ok := importMap.Imports["@myorg/components/button"]; !ok {
		t.Error("Expected '@myorg/components/button' in import map")
	} else if buttonPath != "/packages/components/elements/button.js" {
		t.Errorf("Expected '/packages/components/elements/button.js', got %s", buttonPath)
	}

	if alertPath, ok := importMap.Imports["@myorg/components/alert"]; !ok {
		t.Error("Expected '@myorg/components/alert' in import map")
	} else if alertPath != "/packages/components/elements/alert.js" {
		t.Errorf("Expected '/packages/components/elements/alert.js', got %s", alertPath)
	}

	// Should NOT have broad root trailing slash
	if _, exists := importMap.Imports["@myorg/components/"]; exists {
		t.Error("Should not have root '@myorg/components/' - would expose all package subpaths")
	}

	// Should have wildcard pattern mapping for ./elements/*
	if elementsPath, ok := importMap.Imports["@myorg/components/elements/"]; !ok {
		t.Error("Expected '@myorg/components/elements/' wildcard mapping")
	} else if elementsPath != "/packages/components/elements/" {
		t.Errorf("Expected '/packages/components/elements/', got %s", elementsPath)
	}

	// Should have main + 2 named exports + 1 wildcard mapping
	if len(importMap.Imports) != 4 {
		t.Errorf("Expected 4 import map entries, got %d: %v",
			len(importMap.Imports), importMap.Imports)
	}
}

// TestImportMap_NodeModuleDependencyWithExportsMap tests node_modules export resolution
func TestImportMap_NodeModuleDependencyWithExportsMap(t *testing.T) {
	tmpDir := t.TempDir()

	packageJSON := `{
		"name": "test-project",
		"dependencies": {
			"some-lib": "^1.0.0"
		}
	}`

	if err := os.WriteFile(filepath.Join(tmpDir, "package.json"), []byte(packageJSON), 0644); err != nil {
		t.Fatalf("Failed to write package.json: %v", err)
	}

	// Create node_modules dependency with exports
	depDir := filepath.Join(tmpDir, "node_modules", "some-lib")
	if err := os.MkdirAll(depDir, 0755); err != nil {
		t.Fatalf("Failed to create dependency directory: %v", err)
	}

	depPackageJSON := `{
		"name": "some-lib",
		"version": "1.0.0",
		"exports": {
			".": "./dist/index.js",
			"./utils": "./dist/utils.js",
			"./helpers/*": "./dist/helpers/*"
		}
	}`

	if err := os.WriteFile(filepath.Join(depDir, "package.json"), []byte(depPackageJSON), 0644); err != nil {
		t.Fatalf("Failed to write dependency package.json: %v", err)
	}

	importMap, err := Generate(tmpDir, nil)
	if err != nil {
		t.Fatalf("GenerateImportMap failed: %v", err)
	}

	if importMap == nil {
		t.Fatal("Expected import map, got nil")
	}

	// Should have main export
	if mainPath, ok := importMap.Imports["some-lib"]; !ok {
		t.Error("Expected 'some-lib' in import map")
	} else if mainPath != "/node_modules/some-lib/dist/index.js" {
		t.Errorf("Expected '/node_modules/some-lib/dist/index.js', got %s", mainPath)
	}

	// Should have named export
	if utilsPath, ok := importMap.Imports["some-lib/utils"]; !ok {
		t.Error("Expected 'some-lib/utils' in import map")
	} else if utilsPath != "/node_modules/some-lib/dist/utils.js" {
		t.Errorf("Expected '/node_modules/some-lib/dist/utils.js', got %s", utilsPath)
	}

	// Should NOT have broad root trailing slash mapping
	if _, exists := importMap.Imports["some-lib/"]; exists {
		t.Error("Should not have broad root 'some-lib/' mapping - would incorrectly expose all package subpaths")
	}

	// Should have wildcard pattern mapping for ./helpers/*
	if helpersPath, ok := importMap.Imports["some-lib/helpers/"]; !ok {
		t.Error("Expected 'some-lib/helpers/' wildcard mapping")
	} else if helpersPath != "/node_modules/some-lib/dist/helpers/" {
		t.Errorf("Expected '/node_modules/some-lib/dist/helpers/', got %s", helpersPath)
	}

	// Verify we have main + named export + wildcard mapping
	if len(importMap.Imports) != 3 {
		t.Errorf("Expected 3 import map entries, got %d. Entries: %v",
			len(importMap.Imports), importMap.Imports)
	}
}

// TestImportMap_ConditionalExportsWithImportCondition tests conditional export resolution
func TestImportMap_ConditionalExportsWithImportCondition(t *testing.T) {
	tmpDir := t.TempDir()

	packageJSON := `{
		"name": "test-project",
		"dependencies": {
			"dual-package": "^1.0.0"
		}
	}`

	if err := os.WriteFile(filepath.Join(tmpDir, "package.json"), []byte(packageJSON), 0644); err != nil {
		t.Fatalf("Failed to write package.json: %v", err)
	}

	depDir := filepath.Join(tmpDir, "node_modules", "dual-package")
	if err := os.MkdirAll(depDir, 0755); err != nil {
		t.Fatalf("Failed to create dependency directory: %v", err)
	}

	// Package with conditional exports (import vs require)
	depPackageJSON := `{
		"name": "dual-package",
		"version": "1.0.0",
		"exports": {
			".": {
				"import": "./esm/index.js",
				"require": "./cjs/index.js"
			},
			"./utils": {
				"import": "./esm/utils.js",
				"require": "./cjs/utils.js"
			}
		}
	}`

	if err := os.WriteFile(filepath.Join(depDir, "package.json"), []byte(depPackageJSON), 0644); err != nil {
		t.Fatalf("Failed to write dependency package.json: %v", err)
	}

	importMap, err := Generate(tmpDir, nil)
	if err != nil {
		t.Fatalf("GenerateImportMap failed: %v", err)
	}

	if importMap == nil {
		t.Fatal("Expected import map, got nil")
	}

	// Should use "import" condition, not "require"
	if mainPath, ok := importMap.Imports["dual-package"]; !ok {
		t.Error("Expected 'dual-package' in import map")
	} else if mainPath != "/node_modules/dual-package/esm/index.js" {
		t.Errorf("Expected ESM path '/node_modules/dual-package/esm/index.js', got %s", mainPath)
	}

	if utilsPath, ok := importMap.Imports["dual-package/utils"]; !ok {
		t.Error("Expected 'dual-package/utils' in import map")
	} else if utilsPath != "/node_modules/dual-package/esm/utils.js" {
		t.Errorf("Expected ESM path '/node_modules/dual-package/esm/utils.js', got %s", utilsPath)
	}
}

// TestImportMap_WildcardExports tests wildcard pattern handling
func TestImportMap_WildcardExports(t *testing.T) {
	tmpDir := t.TempDir()

	packageJSON := `{
		"name": "test-project",
		"dependencies": {
			"wildcard-lib": "^1.0.0"
		}
	}`

	if err := os.WriteFile(filepath.Join(tmpDir, "package.json"), []byte(packageJSON), 0644); err != nil {
		t.Fatalf("Failed to write package.json: %v", err)
	}

	depDir := filepath.Join(tmpDir, "node_modules", "wildcard-lib")
	if err := os.MkdirAll(depDir, 0755); err != nil {
		t.Fatalf("Failed to create dependency directory: %v", err)
	}

	depPackageJSON := `{
		"name": "wildcard-lib",
		"version": "1.0.0",
		"exports": {
			".": "./index.js",
			"./*": "./src/*",
			"./features/*": "./dist/features/*"
		}
	}`

	if err := os.WriteFile(filepath.Join(depDir, "package.json"), []byte(depPackageJSON), 0644); err != nil {
		t.Fatalf("Failed to write dependency package.json: %v", err)
	}

	importMap, err := Generate(tmpDir, nil)
	if err != nil {
		t.Fatalf("GenerateImportMap failed: %v", err)
	}

	if importMap == nil {
		t.Fatal("Expected import map, got nil")
	}

	// Should have main export
	if mainPath, ok := importMap.Imports["wildcard-lib"]; !ok {
		t.Error("Expected 'wildcard-lib' in import map")
	} else if mainPath != "/node_modules/wildcard-lib/index.js" {
		t.Errorf("Expected '/node_modules/wildcard-lib/index.js', got %s", mainPath)
	}

	// Root wildcard "./*" -> "./src/*" should map trailing slash to target
	// wildcard-lib/foo.js should resolve to /node_modules/wildcard-lib/src/foo.js
	if slashPath, ok := importMap.Imports["wildcard-lib/"]; !ok {
		t.Error("Expected 'wildcard-lib/' in import map")
	} else if slashPath != "/node_modules/wildcard-lib/src/" {
		t.Errorf("Expected '/node_modules/wildcard-lib/src/', got %s", slashPath)
	}
}

// TestImportMap_WildcardWithPrefix tests wildcard patterns with path prefixes like "./*": "./components/*"
func TestImportMap_WildcardWithPrefix(t *testing.T) {
	tmpDir := t.TempDir()

	// Single package with root wildcard that maps to a subdirectory
	packageJSON := `{
		"name": "design-system",
		"version": "1.0.0",
		"exports": {
			".": "./components.js",
			"./*": "./components/*"
		}
	}`

	if err := os.WriteFile(filepath.Join(tmpDir, "package.json"), []byte(packageJSON), 0644); err != nil {
		t.Fatalf("Failed to write package.json: %v", err)
	}

	importMap, err := Generate(tmpDir, nil)
	if err != nil {
		t.Fatalf("GenerateImportMap failed: %v", err)
	}

	if importMap == nil {
		t.Fatal("Expected import map, got nil")
	}

	// Main export
	if mainPath, ok := importMap.Imports["design-system"]; !ok {
		t.Error("Expected 'design-system' in import map")
	} else if mainPath != "/components.js" {
		t.Errorf("Expected '/components.js', got %s", mainPath)
	}

	// Wildcard should map to the prefix path
	// design-system/button/button.js should resolve to /components/button/button.js
	if slashPath, ok := importMap.Imports["design-system/"]; !ok {
		t.Error("Expected 'design-system/' in import map")
	} else if slashPath != "/components/" {
		t.Errorf("Expected '/components/', got %s", slashPath)
	}
}

// TestImportMap_ConditionOnlyExports tests exports that are only a condition map (no subpaths)
func TestImportMap_ConditionOnlyExports(t *testing.T) {
	tmpDir := t.TempDir()

	packageJSON := `{
		"name": "test-project",
		"dependencies": {
			"condition-only": "^1.0.0"
		}
	}`

	if err := os.WriteFile(filepath.Join(tmpDir, "package.json"), []byte(packageJSON), 0644); err != nil {
		t.Fatalf("Failed to write package.json: %v", err)
	}

	// Package with condition-only exports (common pattern for ESM/CJS dual packages)
	depDir := filepath.Join(tmpDir, "node_modules", "condition-only")
	if err := os.MkdirAll(depDir, 0755); err != nil {
		t.Fatalf("Failed to create dependency directory: %v", err)
	}

	depPackageJSON := `{
		"name": "condition-only",
		"version": "1.0.0",
		"exports": {
			"import": "./dist/index.mjs",
			"require": "./dist/index.cjs"
		}
	}`

	if err := os.WriteFile(filepath.Join(depDir, "package.json"), []byte(depPackageJSON), 0644); err != nil {
		t.Fatalf("Failed to write dependency package.json: %v", err)
	}

	importMap, err := Generate(tmpDir, nil)
	if err != nil {
		t.Fatalf("GenerateImportMap failed: %v", err)
	}

	if importMap == nil {
		t.Fatal("Expected import map, got nil")
	}

	// Should map the package name to the import condition, NOT "condition-only/import"
	if mainPath, ok := importMap.Imports["condition-only"]; !ok {
		t.Error("Expected 'condition-only' in import map")
	} else if mainPath != "/node_modules/condition-only/dist/index.mjs" {
		t.Errorf("Expected '/node_modules/condition-only/dist/index.mjs', got %s", mainPath)
	}

	// Should NOT have "condition-only/import" key
	if _, exists := importMap.Imports["condition-only/import"]; exists {
		t.Error("Should not have 'condition-only/import' key - exports is condition-only, not subpath")
	}
}

// TestImportMap_SubpathWildcardPatterns tests that wildcard patterns like "./utils/*" are handled
func TestImportMap_SubpathWildcardPatterns(t *testing.T) {
	tmpDir := t.TempDir()

	packageJSON := `{
		"name": "test-project",
		"dependencies": {
			"utils-lib": "^1.0.0"
		}
	}`

	if err := os.WriteFile(filepath.Join(tmpDir, "package.json"), []byte(packageJSON), 0644); err != nil {
		t.Fatalf("Failed to write package.json: %v", err)
	}

	depDir := filepath.Join(tmpDir, "node_modules", "utils-lib")
	if err := os.MkdirAll(depDir, 0755); err != nil {
		t.Fatalf("Failed to create dependency directory: %v", err)
	}

	depPackageJSON := `{
		"name": "utils-lib",
		"version": "1.0.0",
		"exports": {
			".": "./index.js",
			"./utils/*": "./src/utils/*",
			"./helpers/*": "./dist/helpers/*"
		}
	}`

	if err := os.WriteFile(filepath.Join(depDir, "package.json"), []byte(depPackageJSON), 0644); err != nil {
		t.Fatalf("Failed to write dependency package.json: %v", err)
	}

	importMap, err := Generate(tmpDir, nil)
	if err != nil {
		t.Fatalf("GenerateImportMap failed: %v", err)
	}

	if importMap == nil {
		t.Fatal("Expected import map, got nil")
	}

	// Should have main export
	if mainPath, ok := importMap.Imports["utils-lib"]; !ok {
		t.Error("Expected 'utils-lib' in import map")
	} else if mainPath != "/node_modules/utils-lib/index.js" {
		t.Errorf("Expected '/node_modules/utils-lib/index.js', got %s", mainPath)
	}

	// Should have wildcard pattern mappings
	if utilsPath, ok := importMap.Imports["utils-lib/utils/"]; !ok {
		t.Error("Expected 'utils-lib/utils/' wildcard mapping in import map")
	} else if utilsPath != "/node_modules/utils-lib/src/utils/" {
		t.Errorf("Expected '/node_modules/utils-lib/src/utils/', got %s", utilsPath)
	}

	if helpersPath, ok := importMap.Imports["utils-lib/helpers/"]; !ok {
		t.Error("Expected 'utils-lib/helpers/' wildcard mapping in import map")
	} else if helpersPath != "/node_modules/utils-lib/dist/helpers/" {
		t.Errorf("Expected '/node_modules/utils-lib/dist/helpers/', got %s", helpersPath)
	}
}

// TestImportMap_MultipleWorkspacesWithExports tests multiple workspace packages with exports
func TestImportMap_MultipleWorkspacesWithExports(t *testing.T) {
	tmpDir := t.TempDir()

	rootPackageJSON := `{
		"name": "monorepo-root",
		"workspaces": ["packages/*"]
	}`

	if err := os.WriteFile(filepath.Join(tmpDir, "package.json"), []byte(rootPackageJSON), 0644); err != nil {
		t.Fatalf("Failed to write package.json: %v", err)
	}

	// First workspace package
	pkg1Dir := filepath.Join(tmpDir, "packages", "ui")
	if err := os.MkdirAll(pkg1Dir, 0755); err != nil {
		t.Fatalf("Failed to create workspace directory: %v", err)
	}

	pkg1JSON := `{
		"name": "@myorg/ui",
		"version": "1.0.0",
		"exports": {
			".": "./index.js",
			"./button": "./button.js",
			"./input": "./input.js"
		}
	}`

	if err := os.WriteFile(filepath.Join(pkg1Dir, "package.json"), []byte(pkg1JSON), 0644); err != nil {
		t.Fatalf("Failed to write workspace package.json: %v", err)
	}

	// Second workspace package
	pkg2Dir := filepath.Join(tmpDir, "packages", "icons")
	if err := os.MkdirAll(pkg2Dir, 0755); err != nil {
		t.Fatalf("Failed to create workspace directory: %v", err)
	}

	pkg2JSON := `{
		"name": "@myorg/icons",
		"version": "1.0.0",
		"exports": {
			".": "./icons.js",
			"./check": "./svg/check.js",
			"./close": "./svg/close.js"
		}
	}`

	if err := os.WriteFile(filepath.Join(pkg2Dir, "package.json"), []byte(pkg2JSON), 0644); err != nil {
		t.Fatalf("Failed to write workspace package.json: %v", err)
	}

	importMap, err := Generate(tmpDir, nil)
	if err != nil {
		t.Fatalf("GenerateImportMap failed: %v", err)
	}

	if importMap == nil {
		t.Fatal("Expected import map, got nil")
	}

	// Check first workspace package
	if mainPath, ok := importMap.Imports["@myorg/ui"]; !ok {
		t.Error("Expected '@myorg/ui' in import map")
	} else if mainPath != "/packages/ui/index.js" {
		t.Errorf("Expected '/packages/ui/index.js', got %s", mainPath)
	}

	if buttonPath, ok := importMap.Imports["@myorg/ui/button"]; !ok {
		t.Error("Expected '@myorg/ui/button' in import map")
	} else if buttonPath != "/packages/ui/button.js" {
		t.Errorf("Expected '/packages/ui/button.js', got %s", buttonPath)
	}

	// Check second workspace package
	if mainPath, ok := importMap.Imports["@myorg/icons"]; !ok {
		t.Error("Expected '@myorg/icons' in import map")
	} else if mainPath != "/packages/icons/icons.js" {
		t.Errorf("Expected '/packages/icons/icons.js', got %s", mainPath)
	}

	if checkPath, ok := importMap.Imports["@myorg/icons/check"]; !ok {
		t.Error("Expected '@myorg/icons/check' in import map")
	} else if checkPath != "/packages/icons/svg/check.js" {
		t.Errorf("Expected '/packages/icons/svg/check.js', got %s", checkPath)
	}
}
