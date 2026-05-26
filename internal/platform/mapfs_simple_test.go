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
package platform_test

import (
	"io"
	"testing"

	"bennypowers.dev/cem/internal/platform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Inline: pure function, scalar assertions
func TestNewMapFS(t *testing.T) {
	fs := platform.NewMapFS(map[string]string{
		"a.js":     "console.log('a')",
		"dir/b.js": "console.log('b')",
	})

	data, err := fs.ReadFile("a.js")
	require.NoError(t, err)
	assert.Equal(t, "console.log('a')", string(data))

	data, err = fs.ReadFile("dir/b.js")
	require.NoError(t, err)
	assert.Equal(t, "console.log('b')", string(data))

	_, err = fs.ReadFile("nonexistent.js")
	assert.Error(t, err)
}

func TestMapFS_WriteAndRead(t *testing.T) {
	fs := platform.NewMapFS(nil)
	err := fs.WriteFile("new.js", []byte("hello"), 0o644)
	require.NoError(t, err)

	data, err := fs.ReadFile("new.js")
	require.NoError(t, err)
	assert.Equal(t, "hello", string(data))
}

func TestMapFS_Remove(t *testing.T) {
	fs := platform.NewMapFS(map[string]string{"a.js": "content"})
	assert.True(t, fs.Exists("a.js"))

	err := fs.Remove("a.js")
	require.NoError(t, err)
	assert.False(t, fs.Exists("a.js"))
}

func TestMapFS_Exists(t *testing.T) {
	fs := platform.NewMapFS(map[string]string{"a.js": "content"})
	assert.True(t, fs.Exists("a.js"))
	assert.False(t, fs.Exists("b.js"))
}

func TestMapFS_Stat(t *testing.T) {
	fs := platform.NewMapFS(map[string]string{"a.js": "content"})
	info, err := fs.Stat("a.js")
	require.NoError(t, err)
	assert.Equal(t, int64(7), info.Size())

	_, err = fs.Stat("nonexistent.js")
	assert.Error(t, err)
}

func TestMapFS_ReadDir(t *testing.T) {
	fs := platform.NewMapFS(map[string]string{
		"dir/a.js": "a",
		"dir/b.js": "b",
	})
	entries, err := fs.ReadDir("dir")
	require.NoError(t, err)
	assert.Len(t, entries, 2)
}

func TestMapFS_Open(t *testing.T) {
	fs := platform.NewMapFS(map[string]string{"a.js": "hello"})
	f, err := fs.Open("a.js")
	require.NoError(t, err)
	defer func() { _ = f.Close() }()

	buf := make([]byte, 5)
	n, err := f.Read(buf)
	require.NoError(t, err)
	assert.Equal(t, 5, n)
	assert.Equal(t, "hello", string(buf))

	n, err = f.Read(buf)
	assert.Equal(t, 0, n)
	assert.ErrorIs(t, err, io.EOF)
}

func TestMapFS_MkdirAll(t *testing.T) {
	fs := platform.NewMapFS(nil)
	err := fs.MkdirAll("a/b/c", 0o755)
	assert.NoError(t, err)
}

func TestMapFS_TempDir(t *testing.T) {
	fs := platform.NewMapFS(nil)
	assert.Equal(t, "/tmp", fs.TempDir())
}
