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
	"io/fs"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"bennypowers.dev/cem/validate"
)

// writeFile writes through the optional write half of a ReadFS. It proves
// callers can write with the operating-system implementation without a wider
// public interface.
func writeFile(t *testing.T, fsys validate.ReadFS, name string, data []byte) {
	t.Helper()
	writer, ok := fsys.(interface {
		WriteFile(name string, data []byte, perm fs.FileMode) error
	})
	if !ok {
		t.Fatalf("filesystem %T does not implement WriteFile", fsys)
	}
	if err := writer.WriteFile(name, data, 0o644); err != nil {
		t.Fatalf("writing %s: %v", name, err)
	}
}

// The public interfaces and aliases must be nameable from an external package.
// The minimal type proves each entry point needs only the methods it declares.
var _ validate.ReadFS = (*minimalReadFS)(nil)
var _ validate.ReadFS = validate.NewOSFileSystem()
var _ validate.ReadFS = (*memoryFS)(nil)
var _ validate.TempFile = (*memoryTempFile)(nil)

// minimalReadFS implements only ReadFS. Schema and Validate must not require
// more.
type minimalReadFS map[string][]byte

func (m minimalReadFS) ReadFile(name string) ([]byte, error) {
	data, ok := m[name]
	if !ok {
		return nil, fs.ErrNotExist
	}
	return data, nil
}

func (m minimalReadFS) Stat(name string) (fs.FileInfo, error) {
	data, ok := m[name]
	if !ok {
		return nil, fs.ErrNotExist
	}
	return memoryFileInfo{name: name, size: int64(len(data))}, nil
}

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

// memoryFS implements ReadFS plus write methods so tests can cache schemas. It
// proves external callers do not need to embed NewOSFileSystem or import an
// internal package.
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

func (m *memoryFS) CreateTemp(dir, pattern string) (validate.TempFile, error) {
	return &memoryTempFile{name: filepath.Join(dir, pattern)}, nil
}

func (m *memoryFS) MkdirAll(path string, perm fs.FileMode) error {
	return nil
}

func (m *memoryFS) Stat(name string) (fs.FileInfo, error) {
	data, ok := m.files[name]
	if !ok {
		return nil, fs.ErrNotExist
	}
	return memoryFileInfo{name: name, size: int64(len(data))}, nil
}

// TestPublicSchemaMinimalReadFS proves external callers can use Schema with an
// implementation of only the two ReadFS methods.
//
// Inline assertions: this test checks the embedded result and the error
// contract, where a golden file would add indirection only.
func TestPublicSchemaMinimalReadFS(t *testing.T) {
	fsys := minimalReadFS{}

	embedded, err := validate.Schema(fsys, "2.1.0")
	if err != nil {
		t.Fatalf("Schema: %v", err)
	}
	if len(embedded) == 0 {
		t.Error("Schema returned an empty embedded schema")
	}

	_, err = validate.Schema(fsys, "9.9.9", validate.WithoutNetwork())
	if err == nil {
		t.Fatal("Schema returned no error for an unknown version with WithoutNetwork")
	}
	// Require the exact no-network contract so a no-op WithoutNetwork, which
	// would surface a fetch or filesystem error instead, cannot pass this test.
	if !strings.Contains(err.Error(), `schema for version "9.9.9" is not embedded or cached and network access is disabled`) {
		t.Errorf("error = %q, want the unknown-version network-disabled error", err)
	}
}

// TestPublicValidateMinimalReadFS proves external callers can validate with an
// implementation of only the two ReadFS methods when the schema is embedded.
//
// Inline assertions: this test checks one validity result, where a golden file
// would add indirection only.
func TestPublicValidateMinimalReadFS(t *testing.T) {
	fsys := minimalReadFS{
		"custom-elements.json": []byte(`{"schemaVersion":"2.1.0","modules":[]}`),
	}

	result, err := validate.Validate(fsys, "custom-elements.json", validate.ValidationOptions{
		NoNetwork: true,
	})
	if err != nil {
		t.Fatalf("Validate: %v", err)
	}
	if !result.IsValid {
		t.Errorf("Validate reported invalid manifest: %+v", result.Errors)
	}
}

// TestPublicGetSchemaCompatibility verifies the deprecated shim released
// callers can still use.
//
// Inline assertions: this test checks that the shim returns an embedded
// schema, where a golden file would add indirection only.
func TestPublicGetSchemaCompatibility(t *testing.T) {
	data, err := validate.GetSchema(minimalReadFS{}, "2.1.0")
	if err != nil {
		t.Fatalf("GetSchema: %v", err)
	}
	if len(data) == 0 {
		t.Error("GetSchema returned an empty schema")
	}
}

// TestCustomFileSystemImplementation validates a manifest through a
// from-scratch FileSystem implementation and returns a custom TempFile.
//
// Inline assertions: this test checks that a hand-written interface
// implementation works, where a golden file would not express the contract.
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
//
// Inline assertions: this test checks the public success contract on a
// temporary manifest, where a golden file would add indirection only.
func TestPublicValidateAPI(t *testing.T) {
	fsys := validate.NewOSFileSystem()
	manifestPath := filepath.Join(t.TempDir(), "custom-elements.json")
	manifest := []byte(`{"schemaVersion":"2.1.0","modules":[]}`)
	writeFile(t, fsys, manifestPath, manifest)

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

// TestPublicValidateAPIRejectsUnsafeSchemaVersion verifies the behavior
// external callers rely on: a path-unsafe schemaVersion can never produce a
// valid result through the exported API.
//
// Inline assertions: this test checks two fields of one result, where a golden
// file would add indirection only.
func TestPublicValidateAPIRejectsUnsafeSchemaVersion(t *testing.T) {
	fsys := validate.NewOSFileSystem()
	manifestPath := filepath.Join(t.TempDir(), "custom-elements.json")
	manifest := []byte(`{"schemaVersion":"../../tmp/x","modules":[]}`)
	writeFile(t, fsys, manifestPath, manifest)

	result, err := validate.ValidateFile(manifestPath, validate.ValidationOptions{})
	if err != nil {
		t.Fatalf("ValidateFile: %v", err)
	}
	if result.IsValid {
		t.Errorf("ValidateFile reported an unsafe schemaVersion as valid: %+v", result.Errors)
	}
	if len(result.Errors) == 0 || result.Errors[0].ID != "schema-version-invalid" {
		t.Errorf("Errors = %+v, want a schema-version-invalid error", result.Errors)
	}
}
