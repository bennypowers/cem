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
	"io/fs"
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

func (m *MapFS) WriteFile(name string, data []byte, perm fs.FileMode) error {
	m.MapFS[name] = &fstest.MapFile{
		Data: data,
		Mode: perm,
	}
	return nil
}

func (m *MapFS) ReadFile(name string) ([]byte, error) {
	return fs.ReadFile(m.MapFS, name)
}

func (m *MapFS) Remove(name string) error {
	delete(m.MapFS, name)
	return nil
}

func (m *MapFS) MkdirAll(path string, perm fs.FileMode) error {
	// MapFS doesn't need explicit directories
	return nil
}

func (m *MapFS) ReadDir(name string) ([]fs.DirEntry, error) {
	return fs.ReadDir(m.MapFS, name)
}

func (m *MapFS) TempDir() string {
	return "/tmp"
}

func (m *MapFS) Stat(name string) (fs.FileInfo, error) {
	return fs.Stat(m.MapFS, name)
}

func (m *MapFS) Exists(path string) bool {
	_, err := fs.Stat(m.MapFS, path)
	return err == nil
}

func (m *MapFS) Open(name string) (fs.File, error) {
	return m.MapFS.Open(name)
}
