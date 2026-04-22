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

package elements

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestTranspileElements(t *testing.T) {
	tests := []struct {
		name        string
		setupFiles  map[string]string // filename -> content
		wantError   bool
		checkOutput map[string]bool // filename -> should exist
	}{
		{
			name: "transpiles single TypeScript file",
			setupFiles: map[string]string{
				"test-element.ts": `export class TestElement extends HTMLElement {
					connectedCallback() {
						this.textContent = 'Hello';
					}
				}`,
			},
			wantError: false,
			checkOutput: map[string]bool{
				"test-element.js": true,
			},
		},
		{
			name: "transpiles multiple TypeScript files",
			setupFiles: map[string]string{
				"element-a.ts": `export class ElementA extends HTMLElement {}`,
				"element-b.ts": `export class ElementB extends HTMLElement {}`,
			},
			wantError: false,
			checkOutput: map[string]bool{
				"element-a.js": true,
				"element-b.js": true,
			},
		},
		{
			name: "skips test files",
			setupFiles: map[string]string{
				"element.ts":      `export class Element extends HTMLElement {}`,
				"element.test.ts": `import { Element } from './element';`,
			},
			wantError: false,
			checkOutput: map[string]bool{
				"element.js":      true,
				"element.test.js": false,
			},
		},
		{
			name: "handles directory with no TypeScript files",
			setupFiles: map[string]string{
				"readme.md": `# Readme`,
			},
			wantError: false,
			checkOutput: map[string]bool{
				"readme.md": true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create temporary directory
			tmpDir := t.TempDir()

			// Set up test files
			for filename, content := range tt.setupFiles {
				path := filepath.Join(tmpDir, filename)
				if err := os.WriteFile(path, []byte(content), 0644); err != nil {
					t.Fatalf("Failed to create test file %s: %v", filename, err)
				}
			}

			// Run transpilation
			err := TranspileElements(tmpDir)

			// Check error
			if (err != nil) != tt.wantError {
				t.Errorf("TranspileElements() error = %v, wantError %v", err, tt.wantError)
				return
			}

			// Check output files
			for filename, shouldExist := range tt.checkOutput {
				path := filepath.Join(tmpDir, filename)
				_, err := os.Stat(path)
				exists := err == nil

				if exists != shouldExist {
					t.Errorf("Output file %s: exists = %v, want %v", filename, exists, shouldExist)
				}

				// For .js files, verify they contain valid JavaScript
				if shouldExist && exists && strings.HasSuffix(filename, ".js") {
					content, err := os.ReadFile(path)
					if err != nil {
						t.Errorf("Failed to read output file %s: %v", filename, err)
						continue
					}
					if len(content) == 0 {
						t.Errorf("Output file %s is empty", filename)
					}
					// Basic sanity check - should contain 'export'
					if !strings.Contains(string(content), "export") {
						t.Errorf("Output file %s doesn't appear to be valid JavaScript: %s", filename, string(content))
					}
				}
			}
		})
	}
}

func TestTranspileElements_Subdirectories(t *testing.T) {
	// Create temporary directory structure
	tmpDir := t.TempDir()

	// Create subdirectories with TypeScript files
	subdirs := []string{"cem-drawer", "cem-panel"}
	for _, subdir := range subdirs {
		dir := filepath.Join(tmpDir, subdir)
		if err := os.MkdirAll(dir, 0755); err != nil {
			t.Fatalf("Failed to create subdirectory %s: %v", subdir, err)
		}

		tsFile := filepath.Join(dir, subdir+".ts")
		content := `export class MyElement extends HTMLElement {}`
		if err := os.WriteFile(tsFile, []byte(content), 0644); err != nil {
			t.Fatalf("Failed to create TypeScript file: %v", err)
		}
	}

	// Run transpilation
	if err := TranspileElements(tmpDir); err != nil {
		t.Fatalf("TranspileElements() error = %v", err)
	}

	// Verify .js files were created in subdirectories
	for _, subdir := range subdirs {
		jsFile := filepath.Join(tmpDir, subdir, subdir+".js")
		if _, err := os.Stat(jsFile); err != nil {
			t.Errorf("Expected .js file not found: %s", jsFile)
		}
	}
}
