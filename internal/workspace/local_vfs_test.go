/*
Copyright © 2026 Benny Powers

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
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/internal/platform"
	"bennypowers.dev/cem/internal/workspace"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Inline: integration test exercising VFS path, scalar assertions

// TestFindWorkspaceRoot_VirtualFS proves that FindWorkspaceRoot works with
// MapFileSystem when given absolute-style virtual paths. filepath.Abs passes
// through already-absolute paths unchanged, so the function works without
// touching the host filesystem.
//
// Limitation: relative start paths would be resolved against the host cwd by
// filepath.Abs, so callers must supply absolute paths when using a virtual FS.
func TestFindWorkspaceRoot_VirtualFS(t *testing.T) {
	t.Run("finds VCS root with .git", func(t *testing.T) {
		mapFS := platform.NewMapFileSystem(nil)
		mapFS.AddFile("/repo/.git/HEAD", "ref: refs/heads/main", 0644)
		mapFS.AddFile("/repo/src/index.ts", `export {}`, 0644)

		root, err := workspace.FindWorkspaceRoot("/repo/src", mapFS)
		require.NoError(t, err)
		assert.Equal(t, filepath.FromSlash("/repo"), root,
			"should find workspace root at VCS boundary")
	})

	t.Run("finds monorepo root when it has customElements", func(t *testing.T) {
		mapFS := platform.NewMapFileSystem(nil)
		mapFS.AddFile("/mono/.git/HEAD", "ref: refs/heads/main", 0644)
		mapFS.AddFile("/mono/package.json", `{
			"name": "monorepo",
			"workspaces": ["packages/*"],
			"customElements": "custom-elements.json"
		}`, 0644)
		mapFS.AddFile("/mono/packages/ui/package.json", `{
			"name": "@mono/ui",
			"customElements": "custom-elements.json"
		}`, 0644)
		mapFS.AddFile("/mono/packages/ui/src/button.ts", `export class Button {}`, 0644)

		root, err := workspace.FindWorkspaceRoot("/mono/packages/ui/src", mapFS)
		require.NoError(t, err)
		// FindWorkspaceRoot climbs to the .git boundary. /mono has workspace
		// metadata AND customElements, so the workspace root candidate wins.
		assert.Equal(t, filepath.FromSlash("/mono"), root,
			"should find monorepo root when it has customElements")
	})

	t.Run("prefers nested package when monorepo root lacks customElements", func(t *testing.T) {
		mapFS := platform.NewMapFileSystem(nil)
		mapFS.AddFile("/mono/.git/HEAD", "ref: refs/heads/main", 0644)
		mapFS.AddFile("/mono/package.json", `{
			"name": "monorepo",
			"workspaces": ["packages/*"]
		}`, 0644)
		mapFS.AddFile("/mono/packages/ui/package.json", `{
			"name": "@mono/ui",
			"customElements": "custom-elements.json"
		}`, 0644)
		mapFS.AddFile("/mono/packages/ui/src/button.ts", `export class Button {}`, 0644)

		root, err := workspace.FindWorkspaceRoot("/mono/packages/ui/src", mapFS)
		require.NoError(t, err)
		// When the monorepo root has workspaces but no customElements field,
		// FindWorkspaceRoot prefers the nearest package.json with customElements.
		assert.Equal(t, filepath.FromSlash("/mono/packages/ui"), root,
			"should prefer nested package with customElements over monorepo root without it")
	})

	t.Run("finds pnpm workspace root when it has customElements", func(t *testing.T) {
		mapFS := platform.NewMapFileSystem(nil)
		mapFS.AddFile("/pnpm-mono/.git/HEAD", "ref: refs/heads/main", 0644)
		mapFS.AddFile("/pnpm-mono/pnpm-workspace.yaml", "packages:\n  - 'elements/*'\n", 0644)
		mapFS.AddFile("/pnpm-mono/package.json", `{
			"name": "@ds/root",
			"customElements": "custom-elements.json"
		}`, 0644)
		mapFS.AddFile("/pnpm-mono/elements/alert/package.json", `{
			"name": "@ds/alert",
			"customElements": "custom-elements.json"
		}`, 0644)
		mapFS.AddFile("/pnpm-mono/elements/alert/src/rh-alert.ts", `export class RhAlert {}`, 0644)

		root, err := workspace.FindWorkspaceRoot("/pnpm-mono/elements/alert/src", mapFS)
		require.NoError(t, err)
		// pnpm-workspace.yaml makes /pnpm-mono a workspace root candidate,
		// and its package.json has customElements, so it wins over the nested package.
		assert.Equal(t, filepath.FromSlash("/pnpm-mono"), root,
			"should find pnpm workspace root")
	})

	t.Run("stops at VCS boundary without workspace metadata", func(t *testing.T) {
		// Simulates a Git submodule: nested .git without workspace fields
		mapFS := platform.NewMapFileSystem(nil)
		// Parent repo
		mapFS.AddFile("/parent/.git/HEAD", "ref: refs/heads/main", 0644)
		mapFS.AddFile("/parent/package.json", `{"name": "parent", "workspaces": ["sub"]}`, 0644)
		// Submodule with its own .git
		mapFS.AddFile("/parent/sub/.git/HEAD", "ref: refs/heads/main", 0644)
		mapFS.AddFile("/parent/sub/package.json", `{
			"name": "submod",
			"customElements": "custom-elements.json"
		}`, 0644)
		mapFS.AddFile("/parent/sub/src/el.ts", `export class El {}`, 0644)

		root, err := workspace.FindWorkspaceRoot("/parent/sub/src", mapFS)
		require.NoError(t, err)
		// The submodule's .git is a hard VCS boundary (no workspace metadata),
		// so traversal stops there.
		assert.Equal(t, filepath.FromSlash("/parent/sub"), root,
			"should stop at submodule boundary, not climb to parent")
	})

	t.Run("relative path requires host filesystem", func(t *testing.T) {
		// Known limitation: FindWorkspaceRoot calls filepath.Abs(startPath)
		// which resolves relative paths against the host cwd. The resolved
		// absolute path then fails Stat on the virtual FS because no matching
		// entry exists. Callers must supply absolute paths when using a virtual FS.
		mapFS := platform.NewMapFileSystem(nil)
		mapFS.AddFile("project/package.json", `{
			"name": "test",
			"customElements": "custom-elements.json"
		}`, 0644)

		_, err := workspace.FindWorkspaceRoot("project", mapFS)
		assert.Error(t, err,
			"relative paths fail with virtual FS because filepath.Abs resolves against host cwd")
		assert.Contains(t, err.Error(), "failed to stat path",
			"error should indicate the stat failure on the host-resolved path")
	})
}
