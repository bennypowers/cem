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
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

// TestImportMap_TransitiveDependenciesWithScopes tests that transitive dependencies
// are added to the scopes block so imports within packages can be resolved.
// Only packages in the dependency tree from root should get scopes (not all of node_modules)
func TestImportMap_TransitiveDependenciesWithScopes(t *testing.T) {
	tmpDir := t.TempDir()

	// Create root package.json with lit dependency
	rootPkg := map[string]interface{}{
		"name": "my-app",
		"dependencies": map[string]string{
			"lit": "^3.0.0",
		},
	}
	rootPkgJSON, _ := json.Marshal(rootPkg)
	if err := os.WriteFile(filepath.Join(tmpDir, "package.json"), rootPkgJSON, 0644); err != nil {
		t.Fatal(err)
	}

	// Create node_modules/lit
	litDir := filepath.Join(tmpDir, "node_modules", "lit")
	if err := os.MkdirAll(litDir, 0755); err != nil {
		t.Fatal(err)
	}

	// lit depends on @lit/reactive-element
	litPkg := map[string]interface{}{
		"name": "lit",
		"dependencies": map[string]string{
			"@lit/reactive-element": "^2.0.0",
		},
		"exports": map[string]interface{}{
			".": "./index.js",
			"./decorators.js": "./decorators.js",
		},
	}
	litPkgJSON, _ := json.Marshal(litPkg)
	if err := os.WriteFile(filepath.Join(litDir, "package.json"), litPkgJSON, 0644); err != nil {
		t.Fatal(err)
	}

	// Create node_modules/@lit/reactive-element
	reactiveElementDir := filepath.Join(tmpDir, "node_modules", "@lit", "reactive-element")
	if err := os.MkdirAll(reactiveElementDir, 0755); err != nil {
		t.Fatal(err)
	}

	reactiveElementPkg := map[string]interface{}{
		"name": "@lit/reactive-element",
		"exports": map[string]interface{}{
			".": "./reactive-element.js",
			"./decorators/*": "./decorators/*.js",
		},
	}
	reactiveElementPkgJSON, _ := json.Marshal(reactiveElementPkg)
	if err := os.WriteFile(filepath.Join(reactiveElementDir, "package.json"), reactiveElementPkgJSON, 0644); err != nil {
		t.Fatal(err)
	}

	// Generate import map
	importMap, err := Generate(tmpDir, nil)
	if err != nil {
		t.Fatalf("GenerateImportMap failed: %v", err)
	}

	// Check that lit is in imports
	if importMap.Imports["lit"] != "/node_modules/lit/index.js" {
		t.Errorf("Expected lit mapping, got: %v", importMap.Imports["lit"])
	}

	// @lit/reactive-element should NOT be in top-level imports since it's not
	// in root package.json dependencies - it's a transitive dependency
	// Instead, it should only be in lit's scope

	// Check that scopes exist for lit
	litScope, ok := importMap.Scopes["/node_modules/lit/"]
	if !ok {
		t.Fatal("Expected scope for /node_modules/lit/")
	}

	// Check that @lit/reactive-element is in lit's scope
	if litScope["@lit/reactive-element"] != "/node_modules/@lit/reactive-element/reactive-element.js" {
		t.Errorf("Expected @lit/reactive-element in lit scope, got: %v", litScope["@lit/reactive-element"])
	}

	// Check wildcard mapping
	if litScope["@lit/reactive-element/decorators/"] != "/node_modules/@lit/reactive-element/decorators/" {
		t.Errorf("Expected @lit/reactive-element/decorators/ in lit scope, got: %v", litScope["@lit/reactive-element/decorators/"])
	}
}

// TestImportMap_OnlyDependencyTreeGetsScopes verifies that only packages
// in the dependency tree get scopes, not all packages in node_modules
func TestImportMap_OnlyDependencyTreeGetsScopes(t *testing.T) {
	tmpDir := t.TempDir()

	// Root only depends on pkg-a
	rootPkg := map[string]interface{}{
		"name": "my-app",
		"dependencies": map[string]string{
			"pkg-a": "^1.0.0",
		},
	}
	rootPkgJSON, _ := json.Marshal(rootPkg)
	if err := os.WriteFile(filepath.Join(tmpDir, "package.json"), rootPkgJSON, 0644); err != nil {
		t.Fatal(err)
	}

	// Create pkg-a (depends on pkg-b)
	pkgADir := filepath.Join(tmpDir, "node_modules", "pkg-a")
	if err := os.MkdirAll(pkgADir, 0755); err != nil {
		t.Fatal(err)
	}
	pkgA := map[string]interface{}{
		"name": "pkg-a",
		"dependencies": map[string]string{
			"pkg-b": "^1.0.0",
		},
		"exports": "./index.js",
	}
	pkgAJSON, _ := json.Marshal(pkgA)
	if err := os.WriteFile(filepath.Join(pkgADir, "package.json"), pkgAJSON, 0644); err != nil {
		t.Fatal(err)
	}

	// Create pkg-b (no dependencies)
	pkgBDir := filepath.Join(tmpDir, "node_modules", "pkg-b")
	if err := os.MkdirAll(pkgBDir, 0755); err != nil {
		t.Fatal(err)
	}
	pkgB := map[string]interface{}{
		"name":    "pkg-b",
		"exports": "./index.js",
	}
	pkgBJSON, _ := json.Marshal(pkgB)
	if err := os.WriteFile(filepath.Join(pkgBDir, "package.json"), pkgBJSON, 0644); err != nil {
		t.Fatal(err)
	}

	// Create pkg-unused (exists in node_modules but not in dependency tree)
	pkgUnusedDir := filepath.Join(tmpDir, "node_modules", "pkg-unused")
	if err := os.MkdirAll(pkgUnusedDir, 0755); err != nil {
		t.Fatal(err)
	}
	pkgUnused := map[string]interface{}{
		"name": "pkg-unused",
		"dependencies": map[string]string{
			"pkg-unused-dep": "^1.0.0",
		},
		"exports": "./index.js",
	}
	pkgUnusedJSON, _ := json.Marshal(pkgUnused)
	if err := os.WriteFile(filepath.Join(pkgUnusedDir, "package.json"), pkgUnusedJSON, 0644); err != nil {
		t.Fatal(err)
	}

	// Create pkg-unused-dep
	pkgUnusedDepDir := filepath.Join(tmpDir, "node_modules", "pkg-unused-dep")
	if err := os.MkdirAll(pkgUnusedDepDir, 0755); err != nil {
		t.Fatal(err)
	}
	pkgUnusedDep := map[string]interface{}{
		"name":    "pkg-unused-dep",
		"exports": "./index.js",
	}
	pkgUnusedDepJSON, _ := json.Marshal(pkgUnusedDep)
	if err := os.WriteFile(filepath.Join(pkgUnusedDepDir, "package.json"), pkgUnusedDepJSON, 0644); err != nil {
		t.Fatal(err)
	}

	// Generate import map
	importMap, err := Generate(tmpDir, nil)
	if err != nil {
		t.Fatalf("GenerateImportMap failed: %v", err)
	}

	// pkg-a should have a scope with pkg-b
	if _, ok := importMap.Scopes["/node_modules/pkg-a/"]; !ok {
		t.Error("Expected scope for pkg-a")
	}
	if importMap.Scopes["/node_modules/pkg-a/"]["pkg-b"] != "/node_modules/pkg-b/index.js" {
		t.Errorf("Expected pkg-b in pkg-a scope, got: %v", importMap.Scopes["/node_modules/pkg-a/"])
	}

	// pkg-unused should NOT have a scope (not in dependency tree)
	if _, ok := importMap.Scopes["/node_modules/pkg-unused/"]; ok {
		t.Error("pkg-unused should not have a scope (not in dependency tree)")
	}

	// pkg-b should NOT have a scope (it has no dependencies)
	if _, ok := importMap.Scopes["/node_modules/pkg-b/"]; ok {
		t.Error("pkg-b should not have a scope (no dependencies)")
	}
}
