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
package workspace_test

import (
	"os"
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/workspace"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestFindWorkspaceRoot tests the workspace root detection logic
func TestFindWorkspaceRoot(t *testing.T) {
	// Create a temporary directory structure for testing
	tmpDir := t.TempDir()

	// Create a workspace structure:
	// tmpDir/
	//   .git/
	//   package.json (with workspaces)
	//   elements/
	//     package.json
	//     src/
	//       my-button.ts

	// Create .git directory
	gitDir := filepath.Join(tmpDir, ".git")
	err := os.Mkdir(gitDir, 0755)
	require.NoError(t, err)

	// Create root package.json with workspaces
	rootPackageJSON := filepath.Join(tmpDir, "package.json")
	err = os.WriteFile(rootPackageJSON, []byte(`{
		"name": "@test/monorepo",
		"workspaces": ["./elements"]
	}`), 0644)
	require.NoError(t, err)

	// Create elements directory
	elementsDir := filepath.Join(tmpDir, "elements")
	err = os.Mkdir(elementsDir, 0755)
	require.NoError(t, err)

	// Create elements package.json
	elementsPackageJSON := filepath.Join(elementsDir, "package.json")
	err = os.WriteFile(elementsPackageJSON, []byte(`{
		"name": "@test/elements"
	}`), 0644)
	require.NoError(t, err)

	// Create src directory
	srcDir := filepath.Join(elementsDir, "src")
	err = os.Mkdir(srcDir, 0755)
	require.NoError(t, err)

	t.Run("Find workspace root from subdirectory", func(t *testing.T) {
		// Start from elements/src/ and should find tmpDir
		root, err := workspace.FindWorkspaceRoot(srcDir)
		require.NoError(t, err)
		assert.Equal(t, tmpDir, root, "Expected to find workspace root at tmpDir")
	})

	t.Run("Find workspace root from workspace package", func(t *testing.T) {
		// Start from elements/ and should find tmpDir (parent with .git)
		root, err := workspace.FindWorkspaceRoot(elementsDir)
		require.NoError(t, err)
		assert.Equal(t, tmpDir, root, "Expected to find workspace root at tmpDir, not elements/")
	})

	t.Run("Find workspace root from workspace root itself", func(t *testing.T) {
		// Start from tmpDir and should find tmpDir
		root, err := workspace.FindWorkspaceRoot(tmpDir)
		require.NoError(t, err)
		assert.Equal(t, tmpDir, root, "Expected to find workspace root at tmpDir")
	})

	t.Run("Find workspace root from file path", func(t *testing.T) {
		// Create a test file
		testFile := filepath.Join(srcDir, "my-button.ts")
		err := os.WriteFile(testFile, []byte("// test"), 0644)
		require.NoError(t, err)

		// Start from file path and should find tmpDir
		root, err := workspace.FindWorkspaceRoot(testFile)
		require.NoError(t, err)
		assert.Equal(t, tmpDir, root, "Expected to find workspace root from file path")
	})
}

// TestFindWorkspaceRoot_PnpmWorkspace tests detection of pnpm workspaces
func TestFindWorkspaceRoot_PnpmWorkspace(t *testing.T) {
	tmpDir := t.TempDir()

	// Create pnpm-workspace.yaml
	pnpmWorkspace := filepath.Join(tmpDir, "pnpm-workspace.yaml")
	err := os.WriteFile(pnpmWorkspace, []byte(`packages:
  - 'packages/*'
`), 0644)
	require.NoError(t, err)

	// Create packages directory
	packagesDir := filepath.Join(tmpDir, "packages")
	err = os.Mkdir(packagesDir, 0755)
	require.NoError(t, err)

	t.Run("Find pnpm workspace root", func(t *testing.T) {
		root, err := workspace.FindWorkspaceRoot(packagesDir)
		require.NoError(t, err)
		assert.Equal(t, tmpDir, root, "Expected to find pnpm workspace root")
	})
}

// TestFindWorkspaceRoot_NoWorkspace tests behavior when no workspace is found
func TestFindWorkspaceRoot_NoWorkspace(t *testing.T) {
	tmpDir := t.TempDir()

	// Create a simple directory with no workspace indicators
	subDir := filepath.Join(tmpDir, "some", "nested", "dir")
	err := os.MkdirAll(subDir, 0755)
	require.NoError(t, err)

	t.Run("Return original path when no workspace found", func(t *testing.T) {
		root, err := workspace.FindWorkspaceRoot(subDir)
		require.NoError(t, err)
		// Should return the original path as fallback
		assert.Equal(t, subDir, root, "Expected to return original path when no workspace found")
	})
}

// TestFindWorkspaceRoot_GitSubmoduleBoundary tests VCS boundary detection
// This ensures that Git submodules are treated as hard boundaries and we don't
// climb to parent repositories when the current directory only has VCS markers
func TestFindWorkspaceRoot_GitSubmoduleBoundary(t *testing.T) {
	tmpDir := t.TempDir()

	// Create parent repo structure:
	// tmpDir/
	//   .git/
	//   package.json (with workspaces)
	//   submodules/
	//     my-submodule/
	//       .git/
	//       package.json (NO workspaces - just a simple package)
	//       src/
	//         index.ts

	// Create parent .git directory
	parentGitDir := filepath.Join(tmpDir, ".git")
	err := os.Mkdir(parentGitDir, 0755)
	require.NoError(t, err)

	// Create parent package.json with workspaces
	parentPackageJSON := filepath.Join(tmpDir, "package.json")
	err = os.WriteFile(parentPackageJSON, []byte(`{
		"name": "@test/parent-repo",
		"workspaces": ["./submodules/*"]
	}`), 0644)
	require.NoError(t, err)

	// Create submodules directory
	submodulesDir := filepath.Join(tmpDir, "submodules")
	err = os.Mkdir(submodulesDir, 0755)
	require.NoError(t, err)

	// Create submodule directory
	submoduleDir := filepath.Join(submodulesDir, "my-submodule")
	err = os.Mkdir(submoduleDir, 0755)
	require.NoError(t, err)

	// Create submodule .git directory (this is a VCS boundary)
	submoduleGitDir := filepath.Join(submoduleDir, ".git")
	err = os.Mkdir(submoduleGitDir, 0755)
	require.NoError(t, err)

	// Create submodule package.json WITHOUT workspaces (just a simple package)
	submodulePackageJSON := filepath.Join(submoduleDir, "package.json")
	err = os.WriteFile(submodulePackageJSON, []byte(`{
		"name": "@test/my-submodule"
	}`), 0644)
	require.NoError(t, err)

	// Create src directory inside submodule
	srcDir := filepath.Join(submoduleDir, "src")
	err = os.Mkdir(srcDir, 0755)
	require.NoError(t, err)

	t.Run("Git submodule is a hard boundary - should NOT climb to parent", func(t *testing.T) {
		// Start from submodule/src/ and should find submoduleDir, NOT tmpDir
		// This is because submoduleDir has .git (VCS marker) even though it doesn't
		// have workspace metadata (pnpm-workspace.yaml or package.json with workspaces)
		root, err := workspace.FindWorkspaceRoot(srcDir)
		require.NoError(t, err)
		assert.Equal(t, submoduleDir, root,
			"Expected to stop at Git submodule boundary, not climb to parent repo")
	})

	t.Run("Starting from submodule root should return submodule root", func(t *testing.T) {
		// Start from submoduleDir itself
		root, err := workspace.FindWorkspaceRoot(submoduleDir)
		require.NoError(t, err)
		assert.Equal(t, submoduleDir, root,
			"Expected to return submodule root, not climb to parent")
	})
}

// TestFindWorkspaceRoot_WorkspaceMetadataClimb tests that we DO climb when
// workspace metadata (not just VCS markers) is present
func TestFindWorkspaceRoot_WorkspaceMetadataClimb(t *testing.T) {
	tmpDir := t.TempDir()

	// Create structure:
	// tmpDir/
	//   pnpm-workspace.yaml (workspace metadata - should climb)
	//   packages/
	//     my-package/
	//       package.json (with workspaces - workspace metadata)
	//       elements/
	//         package.json

	// Create pnpm-workspace.yaml at root
	pnpmWorkspace := filepath.Join(tmpDir, "pnpm-workspace.yaml")
	err := os.WriteFile(pnpmWorkspace, []byte(`packages:
  - 'packages/*'
`), 0644)
	require.NoError(t, err)

	// Create packages directory
	packagesDir := filepath.Join(tmpDir, "packages")
	err = os.Mkdir(packagesDir, 0755)
	require.NoError(t, err)

	// Create my-package directory
	myPackageDir := filepath.Join(packagesDir, "my-package")
	err = os.Mkdir(myPackageDir, 0755)
	require.NoError(t, err)

	// Create my-package package.json WITH workspaces (workspace metadata)
	myPackageJSON := filepath.Join(myPackageDir, "package.json")
	err = os.WriteFile(myPackageJSON, []byte(`{
		"name": "@test/my-package",
		"workspaces": ["./elements"]
	}`), 0644)
	require.NoError(t, err)

	// Create elements directory
	elementsDir := filepath.Join(myPackageDir, "elements")
	err = os.Mkdir(elementsDir, 0755)
	require.NoError(t, err)

	// Create elements package.json
	elementsPackageJSON := filepath.Join(elementsDir, "package.json")
	err = os.WriteFile(elementsPackageJSON, []byte(`{
		"name": "@test/elements"
	}`), 0644)
	require.NoError(t, err)

	t.Run("Should climb when workspace metadata is present", func(t *testing.T) {
		// Start from elements/ and should climb to tmpDir
		// Because both myPackageDir and tmpDir have workspace metadata,
		// we should prefer the topmost workspace
		root, err := workspace.FindWorkspaceRoot(elementsDir)
		require.NoError(t, err)
		assert.Equal(t, tmpDir, root,
			"Expected to climb to topmost workspace when workspace metadata is present")
	})
}
