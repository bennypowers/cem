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

	"bennypowers.dev/cem/internal/platform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestReadPnpmWorkspace tests parsing of pnpm-workspace.yaml
func TestReadPnpmWorkspace(t *testing.T) {
	t.Run("exists", func(t *testing.T) {
		rootDir := filepath.Join("testdata", "pnpm-workspace")
		ws, err := ReadPnpmWorkspace(rootDir, platform.NewOSFileSystem())
		require.NoError(t, err)
		require.NotNil(t, ws)
		assert.Equal(t, []string{"packages/*", "packages/sub/**", "!packages/excluded"}, ws.Packages)
	})

	t.Run("does not exist", func(t *testing.T) {
		rootDir := filepath.Join("testdata", "non-workspace")
		ws, err := ReadPnpmWorkspace(rootDir, platform.NewOSFileSystem())
		require.NoError(t, err)
		assert.Nil(t, ws)
	})
}

// TestHasPnpmWorkspace tests detection of pnpm-workspace.yaml
func TestHasPnpmWorkspace(t *testing.T) {
	t.Run("has pnpm workspace", func(t *testing.T) {
		rootDir := filepath.Join("testdata", "pnpm-workspace")
		assert.True(t, HasPnpmWorkspace(rootDir, platform.NewOSFileSystem()))
	})

	t.Run("no pnpm workspace", func(t *testing.T) {
		rootDir := filepath.Join("testdata", "non-workspace")
		assert.False(t, HasPnpmWorkspace(rootDir, platform.NewOSFileSystem()))
	})

	t.Run("pnpm-only workspace", func(t *testing.T) {
		rootDir := filepath.Join("testdata", "pnpm-only")
		assert.True(t, HasPnpmWorkspace(rootDir, platform.NewOSFileSystem()))
	})
}

// TestDiscoverPnpmPackages tests package discovery from pnpm-workspace.yaml
func TestDiscoverPnpmPackages(t *testing.T) {
	t.Run("finds packages from patterns", func(t *testing.T) {
		rootDir := filepath.Join("testdata", "pnpm-workspace")
		ws, err := ReadPnpmWorkspace(rootDir, platform.NewOSFileSystem())
		require.NoError(t, err)
		require.NotNil(t, ws)

		packages, err := DiscoverPnpmPackages(rootDir, ws, platform.NewOSFileSystem())
		require.NoError(t, err)

		// Should find alpha, beta, nested (sub/**), but NOT excluded
		assert.Len(t, packages, 3)
		assert.Contains(t, packages, "@pnpm/alpha")
		assert.Contains(t, packages, "@pnpm/beta")
		assert.Contains(t, packages, "@pnpm/nested")
		assert.NotContains(t, packages, "@pnpm/excluded")
	})
}

// TestIsWorkspaceMode_PnpmWorkspace tests that pnpm workspaces are detected
func TestIsWorkspaceMode_PnpmWorkspace(t *testing.T) {
	t.Run("pnpm-workspace.yaml only", func(t *testing.T) {
		rootDir := filepath.Join("testdata", "pnpm-workspace")
		assert.True(t, IsWorkspaceMode(rootDir, platform.NewOSFileSystem()))
	})

	t.Run("pnpm-only (no npm workspaces field)", func(t *testing.T) {
		rootDir := filepath.Join("testdata", "pnpm-only")
		assert.True(t, IsWorkspaceMode(rootDir, platform.NewOSFileSystem()))
	})

	t.Run("MapFS with pnpm workspace", func(t *testing.T) {
		fsys := platform.NewMapFS(map[string]string{
			"workspace/package.json": `{"name": "test-root", "version": "1.0.0"}`,
			"workspace/pnpm-workspace.yaml": `packages:
  - 'packages/*'
`,
			"workspace/packages/ele/package.json": `{"name": "@test/ele", "version": "1.0.0"}`,
		})
		assert.True(t, IsWorkspaceMode("workspace", fsys))
	})
}

// TestFindPackagesWithManifests_Pnpm tests FindPackagesWithManifests with pnpm workspaces
func TestFindPackagesWithManifests_Pnpm(t *testing.T) {
	t.Run("pnpm workspace with negated patterns", func(t *testing.T) {
		rootDir := filepath.Join("testdata", "pnpm-workspace")
		packages, err := FindPackagesWithManifests(rootDir, platform.NewOSFileSystem())
		require.NoError(t, err)

		// Should find 3 packages (alpha, beta, nested), excluded should be filtered out
		assert.Len(t, packages, 3)

		names := make(map[string]bool)
		for _, pkg := range packages {
			names[pkg.Name] = true
		}
		assert.True(t, names["@pnpm/alpha"])
		assert.True(t, names["@pnpm/beta"])
		assert.True(t, names["@pnpm/nested"])
		assert.False(t, names["@pnpm/excluded"])
	})

	t.Run("pnpm-only workspace", func(t *testing.T) {
		rootDir := filepath.Join("testdata", "pnpm-only")
		packages, err := FindPackagesWithManifests(rootDir, platform.NewOSFileSystem())
		require.NoError(t, err)

		assert.Len(t, packages, 2)

		names := make(map[string]bool)
		for _, pkg := range packages {
			names[pkg.Name] = true
		}
		assert.True(t, names["@pnpm/pkg-a"])
		assert.True(t, names["@pnpm/pkg-b"])
	})
}

// TestFindPackagesWithManifests_Pnpm_MapFS tests pnpm discovery with MapFS
func TestFindPackagesWithManifests_Pnpm_MapFS(t *testing.T) {
	fsys := platform.NewMapFS(map[string]string{
		"workspace/package.json": `{"name": "pnpm-root", "version": "1.0.0"}`,
		"workspace/pnpm-workspace.yaml": `packages:
  - 'packages/*'
  - '!packages/test-fixtures'
`,
		"workspace/packages/component-a/package.json": `{"name": "@pnpm/comp-a", "version": "1.0.0", "customElements": "custom-elements.json"}`,
		"workspace/packages/component-b/package.json": `{"name": "@pnpm/comp-b", "version": "1.0.0", "customElements": "custom-elements.json"}`,
		"workspace/packages/test-fixtures/package.json": `{"name": "@pnpm/test-fix", "version": "1.0.0", "customElements": "custom-elements.json"}`,
	})

	packages, err := FindPackagesWithManifests("workspace", fsys)
	require.NoError(t, err)

	assert.Len(t, packages, 2)

	names := make(map[string]bool)
	for _, pkg := range packages {
		names[pkg.Name] = true
	}
	assert.True(t, names["@pnpm/comp-a"])
	assert.True(t, names["@pnpm/comp-b"])
	assert.False(t, names["@pnpm/test-fix"])
}

// TestLoadWorkspaceConfig_Pnpm tests config loading from pnpm workspace root
func TestLoadWorkspaceConfig_Pnpm(t *testing.T) {
	t.Run("pnpm workspace with config", func(t *testing.T) {
		// Use workspace-with-config which has a .cem.yaml
		rootDir := filepath.Join("testdata", "workspace-with-config")
		cfg, err := LoadWorkspaceConfig(rootDir, platform.NewOSFileSystem())
		require.NoError(t, err)
		assert.NotNil(t, cfg)
	})

	t.Run("pnpm-only workspace no config returns nil", func(t *testing.T) {
		rootDir := filepath.Join("testdata", "pnpm-only")
		cfg, err := LoadWorkspaceConfig(filepath.Join(rootDir, "packages", "pkg-a"), platform.NewOSFileSystem())
		require.NoError(t, err)
		// No config file in pnpm-only fixture, so nil is expected
		assert.Nil(t, cfg)
	})

	t.Run("pnpm-only workspace root no config returns nil", func(t *testing.T) {
		rootDir := filepath.Join("testdata", "pnpm-only")
		cfg, err := LoadWorkspaceConfig(rootDir, platform.NewOSFileSystem())
		require.NoError(t, err)
		// No config file in pnpm-only fixture, so nil is expected
		assert.Nil(t, cfg)
	})
}

// TestDiscoverPnpmPackages_NegatedPattern tests negation patterns in pnpm-workspace.yaml
func TestDiscoverPnpmPackages_NegatedPattern(t *testing.T) {
	t.Run("MapFS with negated patterns", func(t *testing.T) {
		fsys := platform.NewMapFS(map[string]string{
			"workspace/pnpm-workspace.yaml": `packages:
  - 'packages/*'
  - '!packages/private'
`,
			"workspace/packages/public/package.json": `{"name": "@test/pub", "version": "1.0.0"}`,
			"workspace/packages/private/package.json": `{"name": "@test/private", "version": "1.0.0"}`,
		})

		ws, err := ReadPnpmWorkspace("workspace", fsys)
		require.NoError(t, err)
		require.NotNil(t, ws)

		packages, err := DiscoverPnpmPackages("workspace", ws, fsys)
		require.NoError(t, err)

		assert.Len(t, packages, 1)
		assert.Contains(t, packages, "@test/pub")
		assert.NotContains(t, packages, "@test/private")
	})
}

// TestReadPnpmWorkspace_Empty tests parsing of empty/minimal pnpm-workspace.yaml
func TestReadPnpmWorkspace_Empty(t *testing.T) {
	t.Run("empty file", func(t *testing.T) {
		fsys := platform.NewMapFS(map[string]string{
			"workspace/pnpm-workspace.yaml": "",
		})
		ws, err := ReadPnpmWorkspace("workspace", fsys)
		require.NoError(t, err)
		assert.NotNil(t, ws) // parsed ok, but packages is empty
		assert.Empty(t, ws.Packages)
	})

	t.Run("only whitespace", func(t *testing.T) {
		fsys := platform.NewMapFS(map[string]string{
			"workspace/pnpm-workspace.yaml": "\n  \n\n",
		})
		ws, err := ReadPnpmWorkspace("workspace", fsys)
		require.NoError(t, err)
		assert.NotNil(t, ws)
		assert.Empty(t, ws.Packages)
	})
}

// TestHasPnpmWorkspace_EmptyTests tests that empty pnpm-workspace.yaml is not detected as workspace
func TestHasPnpmWorkspace_EmptyTests(t *testing.T) {
	fsys := platform.NewMapFS(map[string]string{
		"workspace/package.json": `{"name": "root", "version": "1.0.0"}`,
		"workspace/pnpm-workspace.yaml": "",
	})
	assert.False(t, HasPnpmWorkspace("workspace", fsys))
}

// TestIsWorkspaceMode_PnpmAndNpm tests that both pnpm and npm workspaces are detected
func TestIsWorkspaceMode_PnpmAndNpm(t *testing.T) {
	t.Run("npm workspaces only", func(t *testing.T) {
		fsys := platform.NewMapFS(map[string]string{
			"workspace/package.json": `{"name": "root", "workspaces": ["packages/*"]}`,
		})
		assert.True(t, IsWorkspaceMode("workspace", fsys))
	})

	t.Run("pnpm only", func(t *testing.T) {
		fsys := platform.NewMapFS(map[string]string{
			"workspace/package.json": `{"name": "root", "version": "1.0.0"}`,
			"workspace/pnpm-workspace.yaml": `packages:
  - 'packages/*'
`,
		})
		assert.True(t, IsWorkspaceMode("workspace", fsys))
	})

	t.Run("neither", func(t *testing.T) {
		fsys := platform.NewMapFS(map[string]string{
			"workspace/package.json": `{"name": "root", "version": "1.0.0"}`,
		})
		assert.False(t, IsWorkspaceMode("workspace", fsys))
	})
}
