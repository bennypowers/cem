/*
Copyright © 2026 Benny Powers <web@bennypowers.com>

SPDX-License-Identifier: GPL-3.0-or-later

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
package validate_test

import (
	"errors"
	"io"
	"io/fs"
	"path/filepath"
	"testing"
	"time"

	"bennypowers.dev/cem/validate"
)

// The public aliases must be nameable from an external package.
var _ validate.FileSystem = validate.NewOSFileSystem()
var _ validate.FileSystem = (*memoryFS)(nil)
var _ validate.TempFile = (*memoryTempFile)(nil)

var errNotImplemented = errors.New("not implemented")

type memoryTempFile struct {
	name string
	data []byte
}

func (f *memoryTempFile) Write(p []byte) (int, error) {
	f.data = append(f.data, p...)
	return len(p), nil
}

func (f *memoryTempFile) Close() error { return nil }
func (f *memoryTempFile) Name() string { return f.name }

type memoryFileInfo struct {
	name string
	size int64
}

func (i memoryFileInfo) Name() string       { return i.name }
func (i memoryFileInfo) Size() int64        { return i.size }
func (i memoryFileInfo) Mode() fs.FileMode  { return 0o644 }
func (i memoryFileInfo) ModTime() time.Time { return time.Time{} }
func (i memoryFileInfo) IsDir() bool        { return false }
func (i memoryFileInfo) Sys() any           { return nil }

// memoryFS implements validate.FileSystem from scratch. It proves external
// callers do not need to embed NewOSFileSystem or import an internal package.
type memoryFS struct {
	files map[string][]byte
}

func (m *memoryFS) WriteFile(name string, data []byte, perm fs.FileMode) error {
	m.files[name] = append([]byte(nil), data...)
	return nil
}

func (m *memoryFS) ReadFile(name string) ([]byte, error) {
	data, ok := m.files[name]
	if !ok {
		return nil, fs.ErrNotExist
	}
	return data, nil
}

func (m *memoryFS) Remove(name string) error {
	delete(m.files, name)
	return nil
}

func (m *memoryFS) Create(name string) (io.WriteCloser, error) {
	return nil, errNotImplemented
}

func (m *memoryFS) CreateTemp(dir, pattern string) (validate.TempFile, error) {
	return &memoryTempFile{name: filepath.Join(dir, pattern)}, nil
}

func (m *memoryFS) Rename(oldpath, newpath string) error {
	return errNotImplemented
}

func (m *memoryFS) MkdirAll(path string, perm fs.FileMode) error {
	return nil
}

func (m *memoryFS) MkdirTemp(dir, pattern string) (string, error) {
	return "", errNotImplemented
}

func (m *memoryFS) ReadDir(name string) ([]fs.DirEntry, error) {
	return nil, errNotImplemented
}

func (m *memoryFS) TempDir() string {
	return ""
}

func (m *memoryFS) Stat(name string) (fs.FileInfo, error) {
	data, ok := m.files[name]
	if !ok {
		return nil, fs.ErrNotExist
	}
	return memoryFileInfo{name: name, size: int64(len(data))}, nil
}

func (m *memoryFS) Exists(path string) bool {
	_, ok := m.files[path]
	return ok
}

func (m *memoryFS) Glob(pattern string) ([]string, error) {
	return nil, errNotImplemented
}

func (m *memoryFS) Open(name string) (fs.File, error) {
	return nil, errNotImplemented
}

// TestCustomFileSystemImplementation validates a manifest through a
// from-scratch FileSystem implementation and returns a custom TempFile.
func TestCustomFileSystemImplementation(t *testing.T) {
	fsys := &memoryFS{files: map[string][]byte{
		"custom-elements.json": []byte(`{"schemaVersion":"2.1.0","modules":[]}`),
	}}

	result, err := validate.Validate(fsys, "custom-elements.json", validate.ValidationOptions{
		IncludeWarnings: true,
	})
	if err != nil {
		t.Fatalf("Validate: %v", err)
	}
	if !result.IsValid {
		t.Errorf("Validate reported invalid manifest: %+v", result.Errors)
	}

	temp, err := fsys.CreateTemp("", "cem-test")
	if err != nil {
		t.Fatalf("CreateTemp: %v", err)
	}
	if _, err := temp.Write([]byte("hello")); err != nil {
		t.Fatalf("TempFile.Write: %v", err)
	}
	if temp.Name() == "" {
		t.Error("TempFile.Name() returned an empty name")
	}
	if err := temp.Close(); err != nil {
		t.Fatalf("TempFile.Close: %v", err)
	}
}

// TestPublicValidateAPI exercises the documented external usage of the
// package. It deliberately imports only the public validate package and the
// standard library, so it fails to compile if an internal type leaks into the
// public example.
func TestPublicValidateAPI(t *testing.T) {
	fsys := validate.NewOSFileSystem()
	manifestPath := filepath.Join(t.TempDir(), "custom-elements.json")
	manifest := []byte(`{"schemaVersion":"2.1.0","modules":[]}`)
	if err := fsys.WriteFile(manifestPath, manifest, 0o644); err != nil {
		t.Fatalf("writing manifest: %v", err)
	}

	result, err := validate.Validate(fsys, manifestPath, validate.ValidationOptions{
		IncludeWarnings: true,
	})
	if err != nil {
		t.Fatalf("Validate: %v", err)
	}
	if !result.IsValid {
		t.Errorf("Validate reported invalid manifest: %+v", result.Errors)
	}

	result, err = validate.ValidateFile(manifestPath, validate.ValidationOptions{})
	if err != nil {
		t.Fatalf("ValidateFile: %v", err)
	}
	if !result.IsValid {
		t.Errorf("ValidateFile reported invalid manifest: %+v", result.Errors)
	}
}
