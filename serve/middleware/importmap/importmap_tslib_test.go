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

// TestImportMap_TslibExports tests that tslib resolves to the ES6 module version
// with proper exports, not the CommonJS version
func TestImportMap_TslibExports(t *testing.T) {
	tmpDir := t.TempDir()

	// Create root package.json with tslib dependency
	rootPkg := map[string]interface{}{
		"name": "my-app",
		"dependencies": map[string]string{
			"tslib": "^2.8.0",
		},
	}
	rootPkgJSON, _ := json.Marshal(rootPkg)
	if err := os.WriteFile(filepath.Join(tmpDir, "package.json"), rootPkgJSON, 0644); err != nil {
		t.Fatal(err)
	}

	// Create node_modules/tslib with realistic package.json
	tslibDir := filepath.Join(tmpDir, "node_modules", "tslib")
	if err := os.MkdirAll(tslibDir, 0755); err != nil {
		t.Fatal(err)
	}

	// tslib has complex conditional exports
	tslibPkg := map[string]interface{}{
		"name": "tslib",
		"main": "tslib.js",
		"exports": map[string]interface{}{
			".": map[string]interface{}{
				"module": map[string]interface{}{
					"types":   "./modules/index.d.ts",
					"default": "./tslib.es6.mjs",
				},
				"import": map[string]interface{}{
					"node": "./modules/index.js",
					"default": map[string]interface{}{
						"types":   "./modules/index.d.ts",
						"default": "./tslib.es6.mjs",
					},
				},
				"default": "./tslib.js",
			},
		},
	}
	tslibPkgJSON, _ := json.Marshal(tslibPkg)
	if err := os.WriteFile(filepath.Join(tslibDir, "package.json"), tslibPkgJSON, 0644); err != nil {
		t.Fatal(err)
	}

	// Generate import map
	importMap, err := Generate(tmpDir, nil)
	if err != nil {
		t.Fatalf("GenerateImportMap failed: %v", err)
	}

	// Check that tslib is mapped to the ES6 module version (tslib.es6.mjs)
	// NOT to tslib.js (CommonJS version)
	tslibMapping := importMap.Imports["tslib"]
	if tslibMapping != "/node_modules/tslib/tslib.es6.mjs" {
		t.Errorf("Expected tslib to map to ES6 module version, got: %v", tslibMapping)
	}
}
