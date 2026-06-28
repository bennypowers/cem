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
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
)

// TempFile is returned by FileSystem.CreateTemp. It provides write access
// and the generated name, without leaking *os.File into the abstraction.
type TempFile interface {
	io.WriteCloser
	Name() string
}

// FileSystem provides an abstraction over filesystem operations.
// This interface enables:
// - Testing with mock filesystems
// - WASM compatibility (where os package may be limited)
// - Cloud function environments
// - Embedded systems with custom storage
type FileSystem interface {
	// File operations
	WriteFile(name string, data []byte, perm fs.FileMode) error
	ReadFile(name string) ([]byte, error)
	Remove(name string) error
	Create(name string) (io.WriteCloser, error)
	CreateTemp(dir, pattern string) (TempFile, error)
	Rename(oldpath, newpath string) error

	// Directory operations
	MkdirAll(path string, perm fs.FileMode) error
	MkdirTemp(dir, pattern string) (string, error)
	ReadDir(name string) ([]fs.DirEntry, error)
	TempDir() string

	// File system queries
	Stat(name string) (fs.FileInfo, error)
	Exists(path string) bool
	Glob(pattern string) ([]string, error)

	// fs.FS compatibility - allows use with platform.WalkDir
	Open(name string) (fs.File, error)
}

// OSFileSystem implements FileSystem using the standard os package.
// This is the production implementation.
type OSFileSystem struct{}

// NewOSFileSystem creates a new filesystem that uses the standard os package.
func NewOSFileSystem() *OSFileSystem {
	return &OSFileSystem{}
}

func (fs *OSFileSystem) WriteFile(name string, data []byte, perm fs.FileMode) error {
	return os.WriteFile(name, data, perm)
}

func (fs *OSFileSystem) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}

func (fs *OSFileSystem) Remove(name string) error {
	return os.Remove(name)
}

func (fs *OSFileSystem) Create(name string) (io.WriteCloser, error) {
	return os.Create(name)
}

func (fs *OSFileSystem) CreateTemp(dir, pattern string) (TempFile, error) {
	return os.CreateTemp(dir, pattern)
}

func (fs *OSFileSystem) Rename(oldpath, newpath string) error {
	return os.Rename(oldpath, newpath)
}

func (fs *OSFileSystem) MkdirAll(path string, perm fs.FileMode) error {
	return os.MkdirAll(path, perm)
}

func (fs *OSFileSystem) MkdirTemp(dir, pattern string) (string, error) {
	return os.MkdirTemp(dir, pattern)
}

func (fs *OSFileSystem) TempDir() string {
	return os.TempDir()
}

func (fs *OSFileSystem) Stat(name string) (fs.FileInfo, error) {
	return os.Stat(name)
}

func (fs *OSFileSystem) Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func (fs *OSFileSystem) ReadDir(name string) ([]fs.DirEntry, error) {
	return os.ReadDir(name)
}

func (fs *OSFileSystem) Glob(pattern string) ([]string, error) {
	return filepath.Glob(pattern)
}

func (fs *OSFileSystem) Open(name string) (fs.File, error) {
	return os.Open(name)
}

// dirFS wraps a FileSystem into an fs.FS rooted at a directory.
// Open calls are resolved relative to dir via filepath.Join.
type dirFS struct {
	fsys FileSystem
	dir  string
}

// DirFS returns an fs.FS rooted at dir, backed by the given FileSystem.
// This is the platform.FileSystem equivalent of os.DirFS.
func DirFS(fsys FileSystem, dir string) fs.FS {
	return &dirFS{fsys: fsys, dir: dir}
}

func (d *dirFS) Open(name string) (fs.File, error) {
	full := filepath.Join(d.dir, name)
	if !strings.HasPrefix(filepath.Clean(full), filepath.Clean(d.dir)) {
		return nil, &fs.PathError{Op: "open", Path: name, Err: fs.ErrPermission}
	}
	return d.fsys.Open(full)
}

func tempName(dir, pattern string, counter *atomic.Int64) string {
	suffix := fmt.Sprintf("%d", counter.Add(1))
	if i := strings.LastIndex(pattern, "*"); i >= 0 {
		pattern = pattern[:i] + suffix + pattern[i+1:]
	} else {
		pattern += suffix
	}
	return filepath.ToSlash(filepath.Join(dir, pattern))
}
