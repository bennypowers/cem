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

package validate

import "bennypowers.dev/cem/internal/platform"

// FileSystem is the filesystem abstraction accepted by Validate,
// NewValidationPipeline, and GetSchema.
//
// It is an alias for an internal interface, exported so external callers can
// name and implement it without importing the internal package. The full
// method set is:
//
//	WriteFile(name string, data []byte, perm fs.FileMode) error
//	ReadFile(name string) ([]byte, error)
//	Remove(name string) error
//	Create(name string) (io.WriteCloser, error)
//	CreateTemp(dir, pattern string) (TempFile, error)
//	Rename(oldpath, newpath string) error
//	MkdirAll(path string, perm fs.FileMode) error
//	MkdirTemp(dir, pattern string) (string, error)
//	ReadDir(name string) ([]fs.DirEntry, error)
//	TempDir() string
//	Stat(name string) (fs.FileInfo, error)
//	Exists(path string) bool
//	Glob(pattern string) ([]string, error)
//	Open(name string) (fs.File, error)
//
// CreateTemp returns a TempFile. Callers can use NewOSFileSystem, which
// implements the interface, or provide their own in-memory or sandboxed
// implementation.
type FileSystem = platform.FileSystem

// TempFile is returned by FileSystem.CreateTemp. It provides write access and
// the generated name without exposing *os.File. It embeds io.WriteCloser and
// adds Name:
//
//	type TempFile interface {
//		io.WriteCloser
//		Name() string
//	}
type TempFile = platform.TempFile

// NewOSFileSystem returns a FileSystem backed by the operating system. It lets
// external callers use Validate, NewValidationPipeline, and GetSchema without
// importing an internal package or naming the interface type:
//
//	fsys := validate.NewOSFileSystem()
//	result, err := validate.Validate(fsys, "custom-elements.json", validate.ValidationOptions{})
func NewOSFileSystem() FileSystem {
	return platform.NewOSFileSystem()
}

// ValidateFile validates a manifest on the operating system filesystem.
//
// It is a convenience wrapper around Validate and NewOSFileSystem for callers
// that do not need to supply their own filesystem implementation.
func ValidateFile(manifestPath string, options ValidationOptions) (*ValidationResult, error) {
	return Validate(NewOSFileSystem(), manifestPath, options)
}
