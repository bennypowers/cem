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

import (
	"io/fs"

	"bennypowers.dev/cem/internal/platform"
)

// ReadFS is the filesystem abstraction accepted by Schema, Validate,
// ValidateFile, and NewValidationPipeline. It is the smallest interface an
// embedder must implement to read a manifest and load embedded or cached
// schemas:
//
//	ReadFile(name string) ([]byte, error)
//	Stat(name string) (fs.FileInfo, error)
//
// Validation never requires write methods from ReadFS. When the supplied value
// also implements WriteFile and MkdirAll, Schema caches schemas fetched from
// unpkg; otherwise a fetched schema is returned without being cached.
type ReadFS interface {
	ReadFile(name string) ([]byte, error)
	Stat(name string) (fs.FileInfo, error)
}

// TempFile provides write access and the generated name without exposing
// *os.File. It embeds io.WriteCloser and adds Name:
//
//	type TempFile interface {
//		io.WriteCloser
//		Name() string
//	}
type TempFile = platform.TempFile

// NewOSFileSystem returns a ReadFS backed by the operating system. The
// concrete value also implements WriteFile and MkdirAll, so fetched schemas are
// cached automatically. It lets external callers use Validate,
// NewValidationPipeline, and Schema without importing an internal package or
// naming the interface type:
//
//	fsys := validate.NewOSFileSystem()
//	result, err := validate.Validate(fsys, "custom-elements.json", validate.ValidationOptions{})
func NewOSFileSystem() ReadFS {
	return platform.NewOSFileSystem()
}

// ValidateFile validates a manifest on the operating system filesystem.
//
// It is a convenience wrapper around Validate and NewOSFileSystem for callers
// that do not need to supply their own filesystem implementation. Set
// ValidationOptions.NoNetwork to prevent schema fetches.
func ValidateFile(manifestPath string, options ValidationOptions) (*ValidationResult, error) {
	return Validate(NewOSFileSystem(), manifestPath, options)
}
