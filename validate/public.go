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
// name the type when they need to. Callers do not import the internal package:
// use NewOSFileSystem for the host filesystem or pass a custom implementation
// that satisfies the method set documented here.
type FileSystem = platform.FileSystem

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
