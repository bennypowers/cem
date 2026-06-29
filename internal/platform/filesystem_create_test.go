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
	"testing/synctest"

	"bennypowers.dev/cem/internal/platform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMapFS_Create(t *testing.T) {
	fs := platform.NewMapFS(nil)

	w, err := fs.Create("created.txt")
	require.NoError(t, err)

	_, err = io.WriteString(w, "hello from Create")
	require.NoError(t, err)
	require.NoError(t, w.Close())

	data, err := fs.ReadFile("created.txt")
	require.NoError(t, err)
	assert.Equal(t, "hello from Create", string(data))
}

func TestMapFS_CreateTemp(t *testing.T) {
	fs := platform.NewMapFS(nil)

	tf, err := fs.CreateTemp("tmp", "test-*.txt")
	require.NoError(t, err)

	name := tf.Name()
	assert.Contains(t, name, "tmp")
	assert.Contains(t, name, "test-")

	_, err = io.WriteString(tf, "temp content")
	require.NoError(t, err)
	require.NoError(t, tf.Close())

	data, err := fs.ReadFile(name)
	require.NoError(t, err)
	assert.Equal(t, "temp content", string(data))
}

func TestMapFS_Rename(t *testing.T) {
	fs := platform.NewMapFS(map[string]string{"old.txt": "content"})

	err := fs.Rename("old.txt", "new.txt")
	require.NoError(t, err)

	assert.False(t, fs.Exists("old.txt"))
	data, err := fs.ReadFile("new.txt")
	require.NoError(t, err)
	assert.Equal(t, "content", string(data))
}

func TestMapFS_Rename_NotExist(t *testing.T) {
	fs := platform.NewMapFS(nil)
	err := fs.Rename("missing.txt", "dest.txt")
	assert.Error(t, err)
}

func TestMapFS_MkdirTemp(t *testing.T) {
	fs := platform.NewMapFS(nil)

	dir, err := fs.MkdirTemp("", "test-dir-*")
	require.NoError(t, err)
	assert.Contains(t, dir, "test-dir-")
}

func TestMapFS_Glob(t *testing.T) {
	fs := platform.NewMapFS(map[string]string{
		"src/a.ts":      "a",
		"src/b.ts":      "b",
		"src/c.js":      "c",
		"other/d.ts":    "d",
	})

	matches, err := fs.Glob("src/*.ts")
	require.NoError(t, err)
	assert.ElementsMatch(t, []string{"src/a.ts", "src/b.ts"}, matches)
}

func TestMapFS_Glob_NoMatch(t *testing.T) {
	fs := platform.NewMapFS(map[string]string{"a.txt": "content"})

	matches, err := fs.Glob("*.js")
	require.NoError(t, err)
	assert.Empty(t, matches)
}

func TestMapFileSystem_Create(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		mfs := platform.NewMapFileSystem(nil)

		w, err := mfs.Create("/created.txt")
		require.NoError(t, err)

		_, err = io.WriteString(w, "hello from Create")
		require.NoError(t, err)
		require.NoError(t, w.Close())

		data, err := mfs.ReadFile("/created.txt")
		require.NoError(t, err)
		assert.Equal(t, "hello from Create", string(data))
	})
}

func TestMapFileSystem_CreateTemp(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		mfs := platform.NewMapFileSystem(nil)

		tf, err := mfs.CreateTemp("/tmp", "test-*.txt")
		require.NoError(t, err)

		name := tf.Name()
		assert.Contains(t, name, "tmp")
		assert.Contains(t, name, "test-")

		_, err = io.WriteString(tf, "temp content")
		require.NoError(t, err)
		require.NoError(t, tf.Close())

		data, err := mfs.ReadFile(name)
		require.NoError(t, err)
		assert.Equal(t, "temp content", string(data))
	})
}

func TestMapFileSystem_Rename(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		mfs := platform.NewMapFileSystem(nil)
		mfs.AddFile("old.txt", "content", 0644)

		err := mfs.Rename("/old.txt", "/new.txt")
		require.NoError(t, err)

		assert.False(t, mfs.Exists("/old.txt"))
		data, err := mfs.ReadFile("/new.txt")
		require.NoError(t, err)
		assert.Equal(t, "content", string(data))
	})
}

func TestMapFileSystem_Rename_NotExist(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		mfs := platform.NewMapFileSystem(nil)
		err := mfs.Rename("/missing.txt", "/dest.txt")
		assert.Error(t, err)
	})
}

func TestMapFileSystem_MkdirTemp(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		mfs := platform.NewMapFileSystem(nil)

		dir, err := mfs.MkdirTemp("", "test-dir-*")
		require.NoError(t, err)
		assert.Contains(t, dir, "test-dir-")

		assert.True(t, mfs.Exists(dir))
	})
}

func TestMapFileSystem_Glob(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		mfs := platform.NewMapFileSystem(nil)
		mfs.AddFile("src/a.ts", "a", 0644)
		mfs.AddFile("src/b.ts", "b", 0644)
		mfs.AddFile("src/c.js", "c", 0644)
		mfs.AddFile("other/d.ts", "d", 0644)

		matches, err := mfs.Glob("src/*.ts")
		require.NoError(t, err)
		assert.ElementsMatch(t, []string{"src/a.ts", "src/b.ts"}, matches)
	})
}

func TestMapFileSystem_Glob_NoMatch(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		mfs := platform.NewMapFileSystem(nil)
		mfs.AddFile("a.txt", "content", 0644)

		matches, err := mfs.Glob("*.js")
		require.NoError(t, err)
		assert.Empty(t, matches)
	})
}

func TestTempDirFileSystem_Create(t *testing.T) {
	fs, err := platform.NewTempDirFileSystem()
	require.NoError(t, err)
	defer func() { _ = fs.Cleanup() }()

	w, err := fs.Create("created.txt")
	require.NoError(t, err)

	_, err = io.WriteString(w, "hello from Create")
	require.NoError(t, err)
	require.NoError(t, w.Close())

	data, err := fs.ReadFile("created.txt")
	require.NoError(t, err)
	assert.Equal(t, "hello from Create", string(data))
}

func TestTempDirFileSystem_CreateTemp(t *testing.T) {
	fs, err := platform.NewTempDirFileSystem()
	require.NoError(t, err)
	defer func() { _ = fs.Cleanup() }()

	require.NoError(t, fs.MkdirAll("sub", 0755))

	tf, err := fs.CreateTemp("sub", "test-*.txt")
	require.NoError(t, err)

	name := tf.Name()
	assert.Contains(t, name, "test-")

	_, err = io.WriteString(tf, "temp content")
	require.NoError(t, err)
	require.NoError(t, tf.Close())
}

func TestTempDirFileSystem_Rename(t *testing.T) {
	fs, err := platform.NewTempDirFileSystem()
	require.NoError(t, err)
	defer func() { _ = fs.Cleanup() }()

	require.NoError(t, fs.WriteFile("old.txt", []byte("content"), 0644))

	err = fs.Rename("old.txt", "new.txt")
	require.NoError(t, err)

	assert.False(t, fs.Exists("old.txt"))
	data, err := fs.ReadFile("new.txt")
	require.NoError(t, err)
	assert.Equal(t, "content", string(data))
}

func TestTempDirFileSystem_MkdirTemp(t *testing.T) {
	fs, err := platform.NewTempDirFileSystem()
	require.NoError(t, err)
	defer func() { _ = fs.Cleanup() }()

	require.NoError(t, fs.MkdirAll("base", 0755))

	dir, err := fs.MkdirTemp("base", "test-dir-*")
	require.NoError(t, err)
	assert.Contains(t, dir, "test-dir-")
}

func TestTempDirFileSystem_Glob(t *testing.T) {
	fs, err := platform.NewTempDirFileSystem()
	require.NoError(t, err)
	defer func() { _ = fs.Cleanup() }()

	require.NoError(t, fs.MkdirAll("src", 0755))
	require.NoError(t, fs.WriteFile("src/a.ts", []byte("a"), 0644))
	require.NoError(t, fs.WriteFile("src/b.ts", []byte("b"), 0644))
	require.NoError(t, fs.WriteFile("src/c.js", []byte("c"), 0644))

	matches, err := fs.Glob("src/*.ts")
	require.NoError(t, err)
	assert.Len(t, matches, 2)
}

func TestMapFS_CreateTemp_ReplacesAsterisk(t *testing.T) {
	fs := platform.NewMapFS(nil)

	tf, err := fs.CreateTemp("tmp", "cem-*.json")
	require.NoError(t, err)

	name := tf.Name()
	assert.NotContains(t, name, "*", "asterisk should be replaced in temp file name")
	assert.Contains(t, name, "cem-")
	assert.Contains(t, name, ".json")
}

func TestMapFileSystem_CreateTemp_ReplacesAsterisk(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		mfs := platform.NewMapFileSystem(nil)

		tf, err := mfs.CreateTemp("/tmp", "cem-*.json")
		require.NoError(t, err)

		name := tf.Name()
		assert.NotContains(t, name, "*", "asterisk should be replaced in temp file name")
		assert.Contains(t, name, "cem-")
		assert.Contains(t, name, ".json")
	})
}

func TestDirFS_Open(t *testing.T) {
	fs := platform.NewMapFS(map[string]string{
		"root/sub/file.txt": "hello",
	})

	dirFS := platform.DirFS(fs, "root")
	f, err := dirFS.Open("sub/file.txt")
	require.NoError(t, err)
	defer func() { _ = f.Close() }()

	buf := make([]byte, 5)
	n, err := f.Read(buf)
	require.NoError(t, err)
	assert.Equal(t, 5, n)
	assert.Equal(t, "hello", string(buf))
}

func TestDirFS_RejectsPathTraversal(t *testing.T) {
	fs := platform.NewMapFS(map[string]string{
		"secret.txt": "secret",
	})

	dirFS := platform.DirFS(fs, "root")
	_, err := dirFS.Open("../secret.txt")
	assert.Error(t, err, "path traversal should be rejected")
}

func TestDirFS_RejectsSiblingDirectory(t *testing.T) {
	fs := platform.NewMapFS(map[string]string{
		"foobar/secret.txt": "secret",
	})
	dirFS := platform.DirFS(fs, "foo")
	_, err := dirFS.Open("../foobar/secret.txt")
	assert.Error(t, err)
}

func TestMapFS_MkdirTemp_Exists(t *testing.T) {
	fs := platform.NewMapFS(nil)
	dir, err := fs.MkdirTemp("", "test-*")
	require.NoError(t, err)
	assert.True(t, fs.Exists(dir))
}

func TestMapFS_Create_AbsolutePath(t *testing.T) {
	fs := platform.NewMapFS(nil)

	w, err := fs.Create("/abs/path/file.txt")
	require.NoError(t, err)
	_, _ = io.WriteString(w, "content")
	require.NoError(t, w.Close())

	data, err := fs.ReadFile("abs/path/file.txt")
	require.NoError(t, err)
	assert.Equal(t, "content", string(data))
}
