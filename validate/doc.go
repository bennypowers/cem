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

// Package validate validates Custom Elements Manifest documents against the
// official JSON schemas and reports warnings for content that is valid but
// should not be part of a public API.
//
// Schema validation covers every manifest version that cem supports. The
// schemas are embedded in the binary, so validation does not need the network
// for the common case. Warning rules flag patterns such as lifecycle methods,
// private or internal methods, implementation details, and builtin
// superclasses leaking into the manifest.
//
// # External use
//
// This package is public and is part of cem's Go API. It follows the 0.x
// compatibility policy described in CONTRIBUTING.md: treat releases as
// prerelease, pin an exact version, and upgrade deliberately.
//
// # Validating a manifest
//
//	fsys := validate.NewOSFileSystem()
//	result, err := validate.Validate(fsys, "custom-elements.json", validate.ValidationOptions{
//		IncludeWarnings: true,
//	})
//	if err != nil {
//		return err
//	}
//	if !result.IsValid || len(result.Warnings) > 0 {
//		return validate.PrintValidationResult(os.Stderr, "custom-elements.json", result, validate.DisplayOptions{})
//	}
//
// ValidateFile is a convenience wrapper for callers that do not need to supply
// their own filesystem:
//
//	result, err := validate.ValidateFile("custom-elements.json", validate.ValidationOptions{
//		IncludeWarnings: true,
//	})
//
// The filesystem-aware entry points accept a FileSystem. Use NewOSFileSystem
// for the host operating system, or implement the interface for in-memory or
// sandboxed input. FileSystem and TempFile are exported aliases, and the
// FileSystem type documents the full method set, so callers do not need to
// import an internal package. For finer control, build a ValidationPipeline
// with NewValidationPipeline and call its Validate method directly.
//
// Warning rules can be disabled by ID with ValidationOptions.DisabledRules.
// Use PrintValidationResult for human-readable or JSON output, or inspect
// ValidationResult.Errors and ValidationResult.Warnings directly.
package validate
