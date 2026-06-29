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
	"path/filepath"
	"strings"
	"testing"

	"bennypowers.dev/cem/internal/platform"
)

// TestDiscoverWorkspacePackages_NegatedPattern tests that negated patterns (starting with !)
// are properly excluded from the workspace discovery
func TestDiscoverWorkspacePackages_NegatedPattern(t *testing.T) {
	workspacesField := []any{
		"packages/*",
		"!packages/private",
	}

	assertNegatedPattern := func(t *testing.T, packages map[string]string) {
		t.Helper()
		// Should find @test/public and @test/internal, but NOT @test/private
		expectedPackages := map[string]bool{
			"@test/public":   false,
			"@test/internal": false,
		}

		for name := range packages {
			if name == "@test/private" {
				t.Errorf("Found excluded package @test/private - negated pattern not respected")
			}
			if _, expected := expectedPackages[name]; expected {
				expectedPackages[name] = true
			} else if name != "@test/private" {
				t.Errorf("Found unexpected package %s", name)
			}
		}

		// Verify we found all expected packages
		for name, found := range expectedPackages {
			if !found {
				t.Errorf("Expected to find package %s but it was not discovered", name)
			}
		}
	}

	t.Run("os", func(t *testing.T) {
		rootDir := filepath.Join("testdata", "negation-test")
		packages, err := DiscoverWorkspacePackages(rootDir, workspacesField, platform.NewOSFileSystem())
		if err != nil {
			t.Fatalf("DiscoverWorkspacePackages failed: %v", err)
		}
		assertNegatedPattern(t, packages)
	})

	t.Run("mapfs", func(t *testing.T) {
		fsys := platform.NewMapFS(map[string]string{
			"workspace/packages/public/package.json":   `{"name": "@test/public", "version": "1.0.0"}`,
			"workspace/packages/private/package.json":  `{"name": "@test/private", "version": "1.0.0"}`,
			"workspace/packages/internal/package.json": `{"name": "@test/internal", "version": "1.0.0"}`,
		})
		packages, err := DiscoverWorkspacePackages("workspace", workspacesField, fsys)
		if err != nil {
			t.Fatalf("DiscoverWorkspacePackages failed: %v", err)
		}
		assertNegatedPattern(t, packages)
	})
}

// TestDiscoverWorkspacePackages_NegatedPatternObject tests negated patterns with object format
func TestDiscoverWorkspacePackages_NegatedPatternObject(t *testing.T) {
	// Test with object format: { "packages": ["apps/*", "!apps/internal-*"] }
	workspacesField := map[string]any{
		"packages": []any{
			"apps/*",
			"!apps/internal-*",
		},
	}

	assertNegatedObject := func(t *testing.T, packages map[string]string) {
		t.Helper()
		if _, found := packages["@test/internal-app"]; found {
			t.Errorf("Found excluded package @test/internal-app - negated pattern not respected")
		}
		if _, found := packages["@test/app1"]; !found {
			t.Errorf("Expected to find package @test/app1")
		}
		if len(packages) != 1 {
			t.Errorf("Expected 1 package, found %d: %v", len(packages), packages)
		}
	}

	t.Run("os", func(t *testing.T) {
		rootDir := filepath.Join("testdata", "negation-object-test")
		packages, err := DiscoverWorkspacePackages(rootDir, workspacesField, platform.NewOSFileSystem())
		if err != nil {
			t.Fatalf("DiscoverWorkspacePackages failed: %v", err)
		}
		assertNegatedObject(t, packages)
	})

	t.Run("mapfs", func(t *testing.T) {
		fsys := platform.NewMapFS(map[string]string{
			"workspace/apps/app1/package.json":         `{"name": "@test/app1", "version": "1.0.0"}`,
			"workspace/apps/internal-app/package.json": `{"name": "@test/internal-app", "version": "1.0.0"}`,
		})
		packages, err := DiscoverWorkspacePackages("workspace", workspacesField, fsys)
		if err != nil {
			t.Fatalf("DiscoverWorkspacePackages failed: %v", err)
		}
		assertNegatedObject(t, packages)
	})
}

// TestDiscoverWorkspacePackages_MultipleNegations tests multiple negation patterns
func TestDiscoverWorkspacePackages_MultipleNegations(t *testing.T) {
	workspacesField := []any{
		"packages/*",
		"!packages/private",
		"!packages/internal",
	}

	assertMultipleNegations := func(t *testing.T, packages map[string]string) {
		t.Helper()
		if len(packages) != 1 {
			t.Errorf("Expected 1 package, found %d: %v", len(packages), packages)
		}
		if _, found := packages["@test/public"]; !found {
			t.Errorf("Expected to find package @test/public")
		}
		if _, found := packages["@test/private"]; found {
			t.Errorf("Found excluded package @test/private")
		}
		if _, found := packages["@test/internal"]; found {
			t.Errorf("Found excluded package @test/internal")
		}
	}

	t.Run("os", func(t *testing.T) {
		rootDir := filepath.Join("testdata", "negation-test")
		packages, err := DiscoverWorkspacePackages(rootDir, workspacesField, platform.NewOSFileSystem())
		if err != nil {
			t.Fatalf("DiscoverWorkspacePackages failed: %v", err)
		}
		assertMultipleNegations(t, packages)
	})

	t.Run("mapfs", func(t *testing.T) {
		fsys := platform.NewMapFS(map[string]string{
			"workspace/packages/public/package.json":   `{"name": "@test/public", "version": "1.0.0"}`,
			"workspace/packages/private/package.json":  `{"name": "@test/private", "version": "1.0.0"}`,
			"workspace/packages/internal/package.json": `{"name": "@test/internal", "version": "1.0.0"}`,
		})
		packages, err := DiscoverWorkspacePackages("workspace", workspacesField, fsys)
		if err != nil {
			t.Fatalf("DiscoverWorkspacePackages failed: %v", err)
		}
		assertMultipleNegations(t, packages)
	})
}

// TestDiscoverWorkspacePackages_NoNegation tests that normal patterns still work
func TestDiscoverWorkspacePackages_NoNegation(t *testing.T) {
	workspacesField := []any{
		"packages/*",
	}

	t.Run("os", func(t *testing.T) {
		rootDir := filepath.Join("testdata", "negation-test")
		packages, err := DiscoverWorkspacePackages(rootDir, workspacesField, platform.NewOSFileSystem())
		if err != nil {
			t.Fatalf("DiscoverWorkspacePackages failed: %v", err)
		}
		if len(packages) != 3 {
			t.Errorf("Expected 3 packages, found %d: %v", len(packages), packages)
		}
	})

	t.Run("mapfs", func(t *testing.T) {
		fsys := platform.NewMapFS(map[string]string{
			"workspace/packages/public/package.json":   `{"name": "@test/public", "version": "1.0.0"}`,
			"workspace/packages/private/package.json":  `{"name": "@test/private", "version": "1.0.0"}`,
			"workspace/packages/internal/package.json": `{"name": "@test/internal", "version": "1.0.0"}`,
		})
		packages, err := DiscoverWorkspacePackages("workspace", workspacesField, fsys)
		if err != nil {
			t.Fatalf("DiscoverWorkspacePackages failed: %v", err)
		}
		if len(packages) != 3 {
			t.Errorf("Expected 3 packages, found %d: %v", len(packages), packages)
		}
	})
}

// TestIsWorkspaceMode_SinglePackageWorkspace tests that workspace mode is detected
// even when only one package has a customElements manifest
func TestIsWorkspaceMode_SinglePackageWorkspace(t *testing.T) {
	t.Run("os", func(t *testing.T) {
		rootDir := filepath.Join("testdata", "workspace-mode-single-package")
		if !IsWorkspaceMode(rootDir, platform.NewOSFileSystem()) {
			t.Error("Expected IsWorkspaceMode to return true for workspace with single package")
		}
	})

	t.Run("mapfs", func(t *testing.T) {
		fsys := platform.NewMapFS(map[string]string{
			"workspace/package.json": `{"name": "single-package-workspace", "workspaces": ["packages/*"]}`,
			"workspace/packages/elements/package.json":        `{"name": "@test/elements", "version": "1.0.0", "customElements": "custom-elements.json"}`,
			"workspace/packages/elements/custom-elements.json": `{"schemaVersion": "2.0.0", "readme": "", "modules": []}`,
		})
		if !IsWorkspaceMode("workspace", fsys) {
			t.Error("Expected IsWorkspaceMode to return true for workspace with single package")
		}
	})
}

// TestIsWorkspaceMode_NoManifests tests that workspace mode is detected
// even when no packages have customElements manifests
func TestIsWorkspaceMode_NoManifests(t *testing.T) {
	t.Run("os", func(t *testing.T) {
		rootDir := filepath.Join("testdata", "workspace-mode-no-manifest")
		if !IsWorkspaceMode(rootDir, platform.NewOSFileSystem()) {
			t.Error("Expected IsWorkspaceMode to return true for workspace without manifests")
		}
	})

	t.Run("mapfs", func(t *testing.T) {
		fsys := platform.NewMapFS(map[string]string{
			"workspace/package.json":                  `{"name": "workspace-no-manifests", "workspaces": ["packages/*"]}`,
			"workspace/packages/util/package.json":    `{"name": "@test/util", "version": "1.0.0"}`,
		})
		if !IsWorkspaceMode("workspace", fsys) {
			t.Error("Expected IsWorkspaceMode to return true for workspace without manifests")
		}
	})
}

// TestIsWorkspaceMode_NonWorkspace tests that non-workspace directories are detected
func TestIsWorkspaceMode_NonWorkspace(t *testing.T) {
	t.Run("os", func(t *testing.T) {
		rootDir := filepath.Join("testdata", "non-workspace")
		if IsWorkspaceMode(rootDir, platform.NewOSFileSystem()) {
			t.Error("Expected IsWorkspaceMode to return false for non-workspace directory")
		}
	})

	t.Run("mapfs", func(t *testing.T) {
		fsys := platform.NewMapFS(map[string]string{
			"workspace/package.json": `{"name": "regular-package", "version": "1.0.0"}`,
		})
		if IsWorkspaceMode("workspace", fsys) {
			t.Error("Expected IsWorkspaceMode to return false for non-workspace directory")
		}
	})
}

// TestIsWorkspaceMode_MultiplePackages tests that workspace mode works with multiple packages
func TestIsWorkspaceMode_MultiplePackages(t *testing.T) {
	t.Run("os", func(t *testing.T) {
		rootDir := filepath.Join("testdata", "negation-test")
		if !IsWorkspaceMode(rootDir, platform.NewOSFileSystem()) {
			t.Error("Expected IsWorkspaceMode to return true for workspace with multiple packages")
		}
	})

	t.Run("mapfs", func(t *testing.T) {
		fsys := platform.NewMapFS(map[string]string{
			"workspace/package.json":                           `{"name": "test-workspace", "workspaces": ["packages/*", "!packages/private"]}`,
			"workspace/packages/public/package.json":           `{"name": "@test/public", "version": "1.0.0"}`,
			"workspace/packages/private/package.json":          `{"name": "@test/private", "version": "1.0.0"}`,
			"workspace/packages/internal/package.json":         `{"name": "@test/internal", "version": "1.0.0"}`,
		})
		if !IsWorkspaceMode("workspace", fsys) {
			t.Error("Expected IsWorkspaceMode to return true for workspace with multiple packages")
		}
	})
}

// TestFindPackagesForFiles tests that files are correctly mapped to packages.
// Both rootDir and filePaths must use consistent path forms -- either both
// absolute or both relative -- since package paths come from fsys.Glob.
func TestFindPackagesForFiles(t *testing.T) {
	rootDir, err := filepath.Abs(filepath.Join("testdata", "workspace-mode-single-package"))
	if err != nil {
		t.Fatalf("Failed to get absolute path: %v", err)
	}

	tests := []struct {
		name          string
		filePaths     []string
		expectedCount int
		expectedName  string // Optional: check first package name
	}{
		{
			name: "single file in package",
			filePaths: []string{
				filepath.Join(rootDir, "packages", "elements", "src", "my-element.ts"),
			},
			expectedCount: 1,
			expectedName:  "@test/elements",
		},
		{
			name: "multiple files in same package",
			filePaths: []string{
				filepath.Join(rootDir, "packages", "elements", "src", "my-element.ts"),
				filepath.Join(rootDir, "packages", "elements", "src", "another-element.ts"),
			},
			expectedCount: 1, // Both files in same package, should deduplicate
		},
		{
			name: "file outside all packages",
			filePaths: []string{
				filepath.Join(rootDir, "some-other-file.ts"),
			},
			expectedCount: 0,
		},
		{
			name: "file in root node_modules",
			filePaths: []string{
				filepath.Join(rootDir, "node_modules", "some-dep", "index.js"),
			},
			expectedCount: 0,
		},
		{
			name: "mixed files in and outside packages",
			filePaths: []string{
				filepath.Join(rootDir, "packages", "elements", "src", "my-element.ts"),
				filepath.Join(rootDir, "some-file.ts"),
			},
			expectedCount: 1, // Only one file is in a package
		},
		{
			name:          "empty file list",
			filePaths:     []string{},
			expectedCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			packages, err := FindPackagesForFiles(rootDir, tt.filePaths, platform.NewOSFileSystem())
			if err != nil {
				t.Fatalf("FindPackagesForFiles failed: %v", err)
			}

			if len(packages) != tt.expectedCount {
				t.Errorf("Expected %d packages, got %d: %v", tt.expectedCount, len(packages), packages)
			}

			if tt.expectedName != "" && len(packages) > 0 {
				if packages[0].Name != tt.expectedName {
					t.Errorf("Expected package name %s, got %s", tt.expectedName, packages[0].Name)
				}
			}
		})
	}
}

// TestFindPackagesForFiles_MultiplePackages tests that files in different packages
// correctly returns all affected packages
func TestFindPackagesForFiles_MultiplePackages(t *testing.T) {
	rootDir, err := filepath.Abs(filepath.Join("testdata", "multi-package-workspace"))
	if err != nil {
		t.Fatalf("Failed to get absolute path: %v", err)
	}

	// Files in two different packages
	filePaths := []string{
		filepath.Join(rootDir, "packages", "alpha", "src", "component.ts"),
		filepath.Join(rootDir, "packages", "beta", "src", "component.ts"),
	}

	packages, err := FindPackagesForFiles(rootDir, filePaths, platform.NewOSFileSystem())
	if err != nil {
		t.Fatalf("FindPackagesForFiles failed: %v", err)
	}

	if len(packages) != 2 {
		t.Fatalf("Expected 2 packages, got %d: %v", len(packages), packages)
	}

	// Verify both packages are returned
	foundAlpha := false
	foundBeta := false
	for _, pkg := range packages {
		if pkg.Name == "@test/alpha" {
			foundAlpha = true
		}
		if pkg.Name == "@test/beta" {
			foundBeta = true
		}
	}

	if !foundAlpha {
		t.Error("Expected to find @test/alpha package")
	}
	if !foundBeta {
		t.Error("Expected to find @test/beta package")
	}
}

// TestDiscoverWorkspacePackages_MapFS verifies workspace discovery works with an
// in-memory filesystem, ensuring no host filesystem dependency (filepath.Abs).
func TestDiscoverWorkspacePackages_MapFS(t *testing.T) {
	fsys := platform.NewMapFS(map[string]string{
		"workspace/package.json": `{
			"name": "root",
			"workspaces": ["packages/*"]
		}`,
		"workspace/packages/alpha/package.json": `{
			"name": "@test/alpha",
			"customElements": "custom-elements.json"
		}`,
		"workspace/packages/beta/package.json": `{
			"name": "@test/beta",
			"customElements": "custom-elements.json"
		}`,
		"workspace/packages/gamma/package.json": `{
			"name": "@test/gamma"
		}`,
	})

	workspacesField := []any{"packages/*"}
	packages, err := DiscoverWorkspacePackages("workspace", workspacesField, fsys)
	if err != nil {
		t.Fatalf("DiscoverWorkspacePackages failed: %v", err)
	}

	if len(packages) != 3 {
		t.Fatalf("expected 3 packages, got %d: %v", len(packages), packages)
	}

	// Paths should match what fsys.Glob returns -- no filepath.Abs transformation
	for name, path := range packages {
		wantPrefix := "workspace/packages/"
		if !strings.HasPrefix(path, wantPrefix) {
			t.Errorf("package %s: path %q does not start with %q", name, path, wantPrefix)
		}
	}

	expectedPaths := map[string]string{
		"@test/alpha": "workspace/packages/alpha",
		"@test/beta":  "workspace/packages/beta",
		"@test/gamma": "workspace/packages/gamma",
	}
	for name, wantPath := range expectedPaths {
		gotPath, ok := packages[name]
		if !ok {
			t.Errorf("expected package %s not found", name)
			continue
		}
		if gotPath != wantPath {
			t.Errorf("package %s: got path %q, want %q", name, gotPath, wantPath)
		}
	}
}

// TestDiscoverWorkspacePackages_MapFS_NegatedPattern verifies negated patterns
// work with an in-memory filesystem.
func TestDiscoverWorkspacePackages_MapFS_NegatedPattern(t *testing.T) {
	fsys := platform.NewMapFS(map[string]string{
		"workspace/packages/public/package.json":   `{"name": "@test/public"}`,
		"workspace/packages/private/package.json":  `{"name": "@test/private"}`,
		"workspace/packages/internal/package.json": `{"name": "@test/internal"}`,
	})

	workspacesField := []any{"packages/*", "!packages/private"}
	packages, err := DiscoverWorkspacePackages("workspace", workspacesField, fsys)
	if err != nil {
		t.Fatalf("DiscoverWorkspacePackages failed: %v", err)
	}

	if _, found := packages["@test/private"]; found {
		t.Error("negated package @test/private should be excluded")
	}
	if _, found := packages["@test/public"]; !found {
		t.Error("expected @test/public to be included")
	}
	if _, found := packages["@test/internal"]; !found {
		t.Error("expected @test/internal to be included")
	}
	if len(packages) != 2 {
		t.Errorf("expected 2 packages, got %d: %v", len(packages), packages)
	}
}

// TestFindPackagesForFiles_MapFS verifies that FindPackagesForFiles works with
// an in-memory filesystem, ensuring no host filesystem dependency (filepath.Abs).
func TestFindPackagesForFiles_MapFS(t *testing.T) {
	fsys := platform.NewMapFS(map[string]string{
		"workspace/package.json": `{
			"name": "root",
			"workspaces": ["packages/*"]
		}`,
		"workspace/packages/alpha/package.json": `{
			"name": "@test/alpha",
			"customElements": "custom-elements.json"
		}`,
		"workspace/packages/alpha/custom-elements.json": `{"schemaVersion": "2.1.0"}`,
		"workspace/packages/alpha/src/component.ts":     `export class Alpha {}`,
		"workspace/packages/beta/package.json": `{
			"name": "@test/beta",
			"customElements": "custom-elements.json"
		}`,
		"workspace/packages/beta/custom-elements.json": `{"schemaVersion": "2.1.0"}`,
		"workspace/packages/beta/src/component.ts":     `export class Beta {}`,
	})

	tests := []struct {
		name          string
		filePaths     []string
		expectedCount int
		expectedNames []string
	}{
		{
			name:          "file in alpha package",
			filePaths:     []string{"workspace/packages/alpha/src/component.ts"},
			expectedCount: 1,
			expectedNames: []string{"@test/alpha"},
		},
		{
			name: "files in both packages",
			filePaths: []string{
				"workspace/packages/alpha/src/component.ts",
				"workspace/packages/beta/src/component.ts",
			},
			expectedCount: 2,
			expectedNames: []string{"@test/alpha", "@test/beta"},
		},
		{
			name:          "file outside any package",
			filePaths:     []string{"workspace/README.md"},
			expectedCount: 0,
		},
		{
			name:          "empty file list",
			filePaths:     []string{},
			expectedCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			packages, err := FindPackagesForFiles("workspace", tt.filePaths, fsys)
			if err != nil {
				t.Fatalf("FindPackagesForFiles failed: %v", err)
			}
			if len(packages) != tt.expectedCount {
				t.Errorf("expected %d packages, got %d: %v", tt.expectedCount, len(packages), packages)
			}
			for _, wantName := range tt.expectedNames {
				found := false
				for _, pkg := range packages {
					if pkg.Name == wantName {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("expected to find package %s", wantName)
				}
			}
		})
	}
}
