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
package platform

import (
	"bytes"
	"io"
	"io/fs"
	"path/filepath"
	"strings"
	"sync/atomic"
	"testing/fstest"
)

// MapFS wraps testing/fstest.MapFS to implement our FileSystem interface
// This provides an in-memory filesystem for testing with predictable paths
type MapFS struct {
	fstest.MapFS
}

// NewMapFS creates a new in-memory filesystem from a map of file contents
func NewMapFS(files map[string]string) *MapFS {
	mapFS := make(fstest.MapFS)
	for path, content := range files {
		mapFS[path] = &fstest.MapFile{
			Data: []byte(content),
			Mode: 0644,
		}
	}
	return &MapFS{MapFS: mapFS}
}

func cleanMapFSPath(name string) string {
	return strings.TrimPrefix(filepath.ToSlash(filepath.Clean(name)), "/")
}

func (m *MapFS) WriteFile(name string, data []byte, perm fs.FileMode) error {
	name = cleanMapFSPath(name)
	m.MapFS[name] = &fstest.MapFile{
		Data: data,
		Mode: perm,
	}
	return nil
}

func (m *MapFS) ReadFile(name string) ([]byte, error) {
	return fs.ReadFile(m.MapFS, cleanMapFSPath(name))
}

func (m *MapFS) Create(name string) (io.WriteCloser, error) {
	name = cleanMapFSPath(name)
	return &mapFSWriter{fs: m, name: name}, nil
}

func (m *MapFS) CreateTemp(dir, pattern string) (TempFile, error) {
	name := tempName(dir, pattern, &mapFSTempCounter)
	return &mapFSWriter{fs: m, name: name}, nil
}

func (m *MapFS) Rename(oldpath, newpath string) error {
	oldpath = cleanMapFSPath(oldpath)
	newpath = cleanMapFSPath(newpath)
	entry, ok := m.MapFS[oldpath]
	if !ok {
		return &fs.PathError{Op: "rename", Path: oldpath, Err: fs.ErrNotExist}
	}
	m.MapFS[newpath] = entry
	delete(m.MapFS, oldpath)
	return nil
}

func (m *MapFS) Remove(name string) error {
	delete(m.MapFS, cleanMapFSPath(name))
	return nil
}

func (m *MapFS) MkdirAll(path string, perm fs.FileMode) error {
	// MapFS doesn't need explicit directories
	return nil
}

func (m *MapFS) MkdirTemp(dir, pattern string) (string, error) {
	name := tempName(dir, pattern, &mapFSTempCounter)
	m.MapFS[cleanMapFSPath(name)] = &fstest.MapFile{Mode: fs.ModeDir | 0755}
	return name, nil
}

func (m *MapFS) ReadDir(name string) ([]fs.DirEntry, error) {
	return fs.ReadDir(m.MapFS, cleanMapFSPath(name))
}

func (m *MapFS) TempDir() string {
	return "/tmp"
}

func (m *MapFS) Stat(name string) (fs.FileInfo, error) {
	return fs.Stat(m.MapFS, cleanMapFSPath(name))
}

func (m *MapFS) Exists(path string) bool {
	_, err := fs.Stat(m.MapFS, cleanMapFSPath(path))
	return err == nil
}

func (m *MapFS) Glob(pattern string) ([]string, error) {
	return fs.Glob(m.MapFS, cleanMapFSPath(pattern))
}

func (m *MapFS) Open(name string) (fs.File, error) {
	return m.MapFS.Open(cleanMapFSPath(name))
}

var mapFSTempCounter atomic.Int64

type mapFSWriter struct {
	fs   *MapFS
	name string
	buf  bytes.Buffer
}

func (w *mapFSWriter) Write(p []byte) (int, error) {
	return w.buf.Write(p)
}

func (w *mapFSWriter) Close() error {
	w.fs.MapFS[cleanMapFSPath(w.name)] = &fstest.MapFile{
		Data: w.buf.Bytes(),
		Mode: 0644,
	}
	return nil
}

func (w *mapFSWriter) Name() string {
	return w.name
}
