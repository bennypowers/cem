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
	"testing"
)

// TestDiscoverWorkspacePackages_NegatedPattern tests that negated patterns (starting with !)
// are properly excluded from the workspace discovery
func TestDiscoverWorkspacePackages_NegatedPattern(t *testing.T) {
	// Test with array format: ["packages/*", "!packages/private"]
	rootDir := filepath.Join("testdata", "negation-test")

	workspacesField := []interface{}{
		"packages/*",
		"!packages/private",
	}

	packages, err := DiscoverWorkspacePackages(rootDir, workspacesField)
	if err != nil {
		t.Fatalf("DiscoverWorkspacePackages failed: %v", err)
	}

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

// TestDiscoverWorkspacePackages_NegatedPatternObject tests negated patterns with object format
func TestDiscoverWorkspacePackages_NegatedPatternObject(t *testing.T) {
	// Test with object format: { "packages": ["apps/*", "!apps/internal-*"] }
	rootDir := filepath.Join("testdata", "negation-object-test")

	workspacesField := map[string]interface{}{
		"packages": []interface{}{
			"apps/*",
			"!apps/internal-*",
		},
	}

	packages, err := DiscoverWorkspacePackages(rootDir, workspacesField)
	if err != nil {
		t.Fatalf("DiscoverWorkspacePackages failed: %v", err)
	}

	// Should find @test/app1, but NOT @test/internal-app
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

// TestDiscoverWorkspacePackages_MultipleNegations tests multiple negation patterns
func TestDiscoverWorkspacePackages_MultipleNegations(t *testing.T) {
	rootDir := filepath.Join("testdata", "negation-test")

	workspacesField := []interface{}{
		"packages/*",
		"!packages/private",
		"!packages/internal",
	}

	packages, err := DiscoverWorkspacePackages(rootDir, workspacesField)
	if err != nil {
		t.Fatalf("DiscoverWorkspacePackages failed: %v", err)
	}

	// Should only find @test/public
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

// TestDiscoverWorkspacePackages_NoNegation tests that normal patterns still work
func TestDiscoverWorkspacePackages_NoNegation(t *testing.T) {
	rootDir := filepath.Join("testdata", "negation-test")

	workspacesField := []interface{}{
		"packages/*",
	}

	packages, err := DiscoverWorkspacePackages(rootDir, workspacesField)
	if err != nil {
		t.Fatalf("DiscoverWorkspacePackages failed: %v", err)
	}

	// Should find all 3 packages when no negation patterns exist
	if len(packages) != 3 {
		t.Errorf("Expected 3 packages, found %d: %v", len(packages), packages)
	}
}

// TestIsWorkspaceMode_SinglePackageWorkspace tests that workspace mode is detected
// even when only one package has a customElements manifest
func TestIsWorkspaceMode_SinglePackageWorkspace(t *testing.T) {
	rootDir := filepath.Join("testdata", "workspace-mode-single-package")

	// Should return true because workspaces field exists in package.json
	if !IsWorkspaceMode(rootDir) {
		t.Error("Expected IsWorkspaceMode to return true for workspace with single package")
	}
}

// TestIsWorkspaceMode_NoManifests tests that workspace mode is detected
// even when no packages have customElements manifests
func TestIsWorkspaceMode_NoManifests(t *testing.T) {
	rootDir := filepath.Join("testdata", "workspace-mode-no-manifest")

	// Should return true because workspaces field exists, even with no manifests
	if !IsWorkspaceMode(rootDir) {
		t.Error("Expected IsWorkspaceMode to return true for workspace without manifests")
	}
}

// TestIsWorkspaceMode_NonWorkspace tests that non-workspace directories are detected
func TestIsWorkspaceMode_NonWorkspace(t *testing.T) {
	rootDir := filepath.Join("testdata", "non-workspace")

	// Should return false because no workspaces field exists
	if IsWorkspaceMode(rootDir) {
		t.Error("Expected IsWorkspaceMode to return false for non-workspace directory")
	}
}

// TestIsWorkspaceMode_MultiplePackages tests that workspace mode works with multiple packages
func TestIsWorkspaceMode_MultiplePackages(t *testing.T) {
	rootDir := filepath.Join("testdata", "negation-test")

	// Should return true because workspaces field exists
	if !IsWorkspaceMode(rootDir) {
		t.Error("Expected IsWorkspaceMode to return true for workspace with multiple packages")
	}
}

// TestFindPackagesForFiles tests that files are correctly mapped to packages
func TestFindPackagesForFiles(t *testing.T) {
	rootDir := filepath.Join("testdata", "workspace-mode-single-package")
	absRootDir, err := filepath.Abs(rootDir)
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
				filepath.Join(absRootDir, "packages", "elements", "src", "my-element.ts"),
			},
			expectedCount: 1,
			expectedName:  "@test/elements",
		},
		{
			name: "multiple files in same package",
			filePaths: []string{
				filepath.Join(absRootDir, "packages", "elements", "src", "my-element.ts"),
				filepath.Join(absRootDir, "packages", "elements", "src", "another-element.ts"),
			},
			expectedCount: 1, // Both files in same package, should deduplicate
		},
		{
			name: "file outside all packages",
			filePaths: []string{
				filepath.Join(absRootDir, "some-other-file.ts"),
			},
			expectedCount: 0,
		},
		{
			name: "file in root node_modules",
			filePaths: []string{
				filepath.Join(absRootDir, "node_modules", "some-dep", "index.js"),
			},
			expectedCount: 0,
		},
		{
			name: "mixed files in and outside packages",
			filePaths: []string{
				filepath.Join(absRootDir, "packages", "elements", "src", "my-element.ts"),
				filepath.Join(absRootDir, "some-file.ts"),
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
			packages, err := FindPackagesForFiles(rootDir, tt.filePaths)
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
