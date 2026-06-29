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
package workspace

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestGlob_CrossPlatform(t *testing.T) {
	// Create temporary directory structure
	tmpDir := t.TempDir()

	// Create test files
	testFiles := []string{
		"src/component.ts",
		"src/utils/helper.ts",
		"dist/component.js",
		"tests/component.test.ts",
	}

	for _, file := range testFiles {
		fullPath := filepath.Join(tmpDir, file)
		if err := os.MkdirAll(filepath.Dir(fullPath), 0755); err != nil {
			t.Fatalf("Failed to create directory: %v", err)
		}
		if err := os.WriteFile(fullPath, []byte("test content"), 0644); err != nil {
			t.Fatalf("Failed to create test file: %v", err)
		}
	}

	ctx := NewFileSystemWorkspaceContext(tmpDir)

	tests := []struct {
		name     string
		pattern  string
		expected []string
	}{
		{
			name:     "glob_pattern_ts_files",
			pattern:  "src/**/*.ts",
			expected: []string{"src/component.ts", "src/utils/helper.ts"},
		},
		{
			name:     "relative_path",
			pattern:  "src/component.ts",
			expected: []string{"src/component.ts"},
		},
		{
			name:     "path_with_separators",
			pattern:  filepath.Join("src", "utils", "helper.ts"),
			expected: []string{filepath.Join("src", "utils", "helper.ts")},
		},
		{
			name:     "absolute_path_within_project",
			pattern:  filepath.Join(tmpDir, "dist", "component.js"),
			expected: []string{filepath.Join("dist", "component.js")},
		},
		{
			name:     "absolute_path_outside_project",
			pattern:  "/tmp/external/file.ts",
			expected: []string{}, // Should return empty for external files
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ctx.Glob(tt.pattern)
			if err != nil {
				t.Fatalf("Glob failed: %v", err)
			}

			// Normalize expected paths for cross-platform comparison
			normalizedExpected := make([]string, len(tt.expected))
			for i, path := range tt.expected {
				normalizedExpected[i] = filepath.Clean(path)
			}

			if len(result) != len(normalizedExpected) {
				t.Fatalf("Expected %d results, got %d: %v", len(normalizedExpected), len(result), result)
			}

			// Check each result (order-independent)
			for _, expected := range normalizedExpected {
				found := false
				for _, actual := range result {
					if filepath.Clean(actual) == expected {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("Expected result %q not found in %v", expected, result)
				}
			}
		})
	}
}

func TestMakeRelativeToRoot_CrossPlatform(t *testing.T) {
	tmpDir := t.TempDir()
	ctx := NewFileSystemWorkspaceContext(tmpDir)

	tests := []struct {
		name        string
		absPath     string
		expectError bool
		expected    string
	}{
		{
			name:        "path_within_project",
			absPath:     filepath.Join(tmpDir, "src", "component.ts"),
			expectError: false,
			expected:    filepath.Join("src", "component.ts"),
		},
		{
			name:        "path_at_project_root",
			absPath:     filepath.Join(tmpDir, "package.json"),
			expectError: false,
			expected:    "package.json",
		},
		{
			name:        "path_outside_project",
			absPath:     "/tmp/external/file.ts",
			expectError: true,
		},
	}

	if runtime.GOOS == "windows" {
		// Add Windows-specific test cases
		tests = append(tests, struct {
			name        string
			absPath     string
			expectError bool
			expected    string
		}{
			name:        "windows_path_within_project",
			absPath:     filepath.Join(tmpDir, "src\\component.ts"),
			expectError: false,
			expected:    filepath.Join("src", "component.ts"),
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ctx.makeRelativeToRoot(tt.absPath)

			if tt.expectError {
				if err == nil {
					t.Error("Expected error but got none")
				}
				return
			}

			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			expected := filepath.Clean(tt.expected)
			actual := filepath.Clean(result)

			if actual != expected {
				t.Errorf("Expected %q, got %q", expected, actual)
			}
		})
	}
}

// Inline assertions: testing error sentinel types via errors.Is; goldens cannot capture this.
// Real filesystem: doublestar.Glob requires actual disk paths, not MapFS.

func TestGlob_AllOutsideRoot(t *testing.T) {
	tmpDir := t.TempDir()

	siblingFiles := []string{
		"packages/components/button/src/button.ts",
		"packages/components/card/src/card.ts",
	}
	for _, file := range siblingFiles {
		fullPath := filepath.Join(tmpDir, file)
		if err := os.MkdirAll(filepath.Dir(fullPath), 0755); err != nil {
			t.Fatalf("Failed to create directory: %v", err)
		}
		if err := os.WriteFile(fullPath, []byte("export class Foo {}"), 0644); err != nil {
			t.Fatalf("Failed to create test file: %v", err)
		}
	}

	projectRoot := filepath.Join(tmpDir, "website")
	if err := os.MkdirAll(projectRoot, 0755); err != nil {
		t.Fatalf("Failed to create project dir: %v", err)
	}

	ctx := NewFileSystemWorkspaceContext(projectRoot)

	result, err := ctx.Glob("../packages/components/**/*.ts")

	if len(result) != 0 {
		t.Errorf("Expected empty result, got %v", result)
	}
	if !errors.Is(err, ErrGlobAllOutsideRoot) {
		t.Errorf("Expected ErrGlobAllOutsideRoot, got %v", err)
	}
}

func TestGlob_SomeOutsideRoot(t *testing.T) {
	tmpDir := t.TempDir()

	// Files both inside and outside project root
	allFiles := []string{
		"packages/components/button/src/button.ts", // outside website/
		"website/src/local.ts",                     // inside website/
	}
	for _, file := range allFiles {
		fullPath := filepath.Join(tmpDir, file)
		if err := os.MkdirAll(filepath.Dir(fullPath), 0755); err != nil {
			t.Fatalf("Failed to create directory: %v", err)
		}
		if err := os.WriteFile(fullPath, []byte("export class Foo {}"), 0644); err != nil {
			t.Fatalf("Failed to create test file: %v", err)
		}
	}

	projectRoot := filepath.Join(tmpDir, "website")
	ctx := NewFileSystemWorkspaceContext(projectRoot)

	// Patterns starting with ".." are rejected upfront as outside-root
	// with DirFS-scoped globbing -- they cannot partially match inside root.
	result, err := ctx.Glob("../**/*.ts")

	if len(result) != 0 {
		t.Errorf("Expected empty result, got %v", result)
	}
	if !errors.Is(err, ErrGlobAllOutsideRoot) {
		t.Errorf("Expected ErrGlobAllOutsideRoot, got %v", err)
	}
}

func TestGlob_NoneMatched(t *testing.T) {
	tmpDir := t.TempDir()
	ctx := NewFileSystemWorkspaceContext(tmpDir)

	result, err := ctx.Glob("nonexistent/**/*.ts")

	if len(result) != 0 {
		t.Errorf("Expected empty result, got %v", result)
	}
	if !errors.Is(err, ErrGlobNoneMatched) {
		t.Errorf("Expected ErrGlobNoneMatched, got %v", err)
	}
}

func TestHandleNonGlobPattern_CrossPlatform(t *testing.T) {
	tmpDir := t.TempDir()
	ctx := NewFileSystemWorkspaceContext(tmpDir)

	// Test path cleaning with different separators
	testPattern := "src/./component/../component.ts"
	result, err := ctx.handleNonGlobPattern(testPattern)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(result) != 1 {
		t.Fatalf("Expected 1 result, got %d", len(result))
	}

	expected := filepath.Join("src", "component.ts")
	actual := filepath.Clean(result[0])

	if actual != expected {
		t.Errorf("Expected cleaned path %q, got %q", expected, actual)
	}
}
