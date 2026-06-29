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
package generate_test

import (
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/internal/platform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Inline: VFS abstraction test, verifying path conversion logic

// TestModuleProcessor_VirtualFS_PathHandling proves that the path construction
// used by NewModuleProcessor (filepath.Join(root, file) then fsys.ReadFile)
// works correctly with MapFileSystem, which strips leading slashes internally.
// This verifies that code reading modules through the FS abstraction does not
// depend on the host filesystem.
func TestModuleProcessor_VirtualFS_PathHandling(t *testing.T) {
	// MapFileSystem (mocks.go) cleans paths in all methods, stripping leading "/".
	// MapFS (mapfs.go) does NOT clean paths in ReadFile/Open/Stat/Exists.
	// NewModuleProcessor builds: path = filepath.Join(ctx.Root(), file)
	// When root is "/project", this produces "/project/src/my-element.ts".
	// MapFileSystem handles this; MapFS does not.

	const root = "/project"
	const relFile = "src/my-element.ts"
	absPath := filepath.Join(root, relFile)

	moduleContent := `import {LitElement, html} from 'lit';
import {customElement, property} from 'lit/decorators.js';

@customElement('my-element')
export class MyElement extends LitElement {
  @property() name: string = 'World';

  render() {
    return html` + "`" + `<p>Hello, ${this.name}!</p>` + "`" + `;
  }
}
`

	t.Run("MapFileSystem handles absolute paths from filepath.Join", func(t *testing.T) {
		mapFS := platform.NewMapFileSystem(nil)
		mapFS.AddFile(absPath, moduleContent, 0644)

		// Read using the absolute path, exactly as NewModuleProcessor does
		data, err := mapFS.ReadFile(absPath)
		require.NoError(t, err, "MapFileSystem.ReadFile should handle absolute path %q", absPath)
		assert.Equal(t, moduleContent, string(data))
	})

	t.Run("MapFileSystem Stat works with absolute paths", func(t *testing.T) {
		mapFS := platform.NewMapFileSystem(nil)
		mapFS.AddFile(absPath, moduleContent, 0644)

		info, err := mapFS.Stat(absPath)
		require.NoError(t, err, "MapFileSystem.Stat should handle absolute path %q", absPath)
		assert.False(t, info.IsDir())
	})

	t.Run("MapFileSystem Exists works with absolute paths", func(t *testing.T) {
		mapFS := platform.NewMapFileSystem(nil)
		mapFS.AddFile(absPath, moduleContent, 0644)

		assert.True(t, mapFS.Exists(absPath),
			"MapFileSystem.Exists should find file at absolute path %q", absPath)
	})

	t.Run("MapFileSystem Open works with absolute paths", func(t *testing.T) {
		mapFS := platform.NewMapFileSystem(nil)
		mapFS.AddFile(absPath, moduleContent, 0644)

		f, err := mapFS.Open(absPath)
		require.NoError(t, err, "MapFileSystem.Open should handle absolute path %q", absPath)
		require.NoError(t, f.Close())
	})

	t.Run("simple MapFS fails with absolute paths", func(t *testing.T) {
		// This documents the limitation: the simple MapFS (mapfs.go) does not
		// clean paths in read methods, so absolute paths fail. Code that calls
		// NewModuleProcessor must use MapFileSystem, not MapFS.
		simpleFS := platform.NewMapFS(map[string]string{
			"project/src/my-element.ts": moduleContent,
		})

		_, err := simpleFS.ReadFile(absPath)
		assert.Error(t, err,
			"simple MapFS should fail on absolute path %q because it does not strip leading /", absPath)
	})
}
