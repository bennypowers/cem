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
package workspace

import (
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/internal/platform"
	"github.com/stretchr/testify/assert"
)

// Inline: pure function under test, scalar assertions

// TestFindProjectRootFromDir_VirtualFS proves that findProjectRootFromDir works
// with an in-memory filesystem, without depending on os.Getwd() or the host disk.
// findProjectRootFromDir walks upward from startDir checking for project markers
// (package.json, .git, tsconfig.json, CEM config paths) via fsys.Exists.
func TestFindProjectRootFromDir_VirtualFS(t *testing.T) {
	t.Run("finds root with package.json", func(t *testing.T) {
		mapFS := platform.NewMapFileSystem(nil)
		mapFS.AddFile("/project/package.json", `{"name": "test"}`, 0644)
		mapFS.AddFile("/project/src/components/button.ts", `export class Button {}`, 0644)

		root, found := findProjectRootFromDir("/project/src/components", mapFS)
		assert.True(t, found, "should find project root")
		assert.Equal(t, filepath.FromSlash("/project"), root)
	})

	t.Run("finds root with .git directory", func(t *testing.T) {
		mapFS := platform.NewMapFileSystem(nil)
		// AddDir is not needed; placing a file under .git/ is enough for Exists
		mapFS.AddFile("/repo/.git/HEAD", "ref: refs/heads/main", 0644)
		mapFS.AddFile("/repo/src/index.ts", `export {}`, 0644)

		root, found := findProjectRootFromDir("/repo/src", mapFS)
		assert.True(t, found, "should find project root via .git")
		assert.Equal(t, filepath.FromSlash("/repo"), root)
	})

	t.Run("finds root with tsconfig.json", func(t *testing.T) {
		mapFS := platform.NewMapFileSystem(nil)
		mapFS.AddFile("/app/tsconfig.json", `{"compilerOptions":{}}`, 0644)
		mapFS.AddFile("/app/lib/utils.ts", `export function noop() {}`, 0644)

		root, found := findProjectRootFromDir("/app/lib", mapFS)
		assert.True(t, found, "should find project root via tsconfig.json")
		assert.Equal(t, filepath.FromSlash("/app"), root)
	})

	t.Run("finds root with CEM config in .config", func(t *testing.T) {
		mapFS := platform.NewMapFileSystem(nil)
		mapFS.AddFile("/workspace/.config/cem.yaml", "generate:\n  files: []\n", 0644)
		mapFS.AddFile("/workspace/elements/my-el/my-el.ts", `export class MyEl {}`, 0644)

		root, found := findProjectRootFromDir("/workspace/elements/my-el", mapFS)
		assert.True(t, found, "should find project root via .config/cem.yaml")
		assert.Equal(t, filepath.FromSlash("/workspace"), root)
	})

	t.Run("prefers closest ancestor with marker", func(t *testing.T) {
		mapFS := platform.NewMapFileSystem(nil)
		// Two package.json at different levels
		mapFS.AddFile("/mono/package.json", `{"name": "monorepo"}`, 0644)
		mapFS.AddFile("/mono/packages/ui/package.json", `{"name": "@mono/ui"}`, 0644)
		mapFS.AddFile("/mono/packages/ui/src/button.ts", `export class Button {}`, 0644)

		root, found := findProjectRootFromDir("/mono/packages/ui/src", mapFS)
		assert.True(t, found, "should find project root")
		// findProjectRootFromDir walks upward and returns the FIRST match,
		// which is the closest ancestor containing a marker file.
		assert.Equal(t, filepath.FromSlash("/mono/packages/ui"), root)
	})

	t.Run("returns false when no marker found", func(t *testing.T) {
		mapFS := platform.NewMapFileSystem(nil)
		// Only data files, no project markers
		mapFS.AddFile("/data/readme.txt", "hello", 0644)

		_, found := findProjectRootFromDir("/data", mapFS)
		assert.False(t, found, "should not find project root without any markers")
	})
}
