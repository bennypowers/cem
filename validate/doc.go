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
// Manifests with schemaVersion 2.1.0 (and 2.1.1) are validated against an
// embedded 2.1.1-speculative schema rather than the official 2.1.0 schema.
// This works around upstream schema issues, but it means validation for those
// versions is not strict validation against the official 2.1.0 schema. The
// speculative schema rejects additional top-level properties, so an
// unrecognized key at the manifest's top level is reported as an error.
//
// Shortened schema versions are ranges, not exact versions. A bare major such
// as "2" selects the newest embedded schema with that major, and a
// major.minor spelling such as "2.1" selects the newest embedded schema with
// those components. For example, schemaVersion "2" is validated against the
// embedded 2.1.1-speculative schema, like 2.1.0 and 2.1.1, rather than against
// 2.0.0. Exact versions with no embedded schema, and path-safe non-semver
// values such as dist-tags, fall back to the schema cache and then to npm
// using the spelling the manifest supplied. The "schema-version-old" warning
// follows the schema actually used, so a range that resolves to 2.1.x or newer
// is not reported as older than 2.1.0.
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
//		if err := validate.PrintValidationResult(os.Stderr, "custom-elements.json", result, validate.DisplayOptions{}); err != nil {
//			return err
//		}
//	}
//	if !result.IsValid {
//		return fmt.Errorf("manifest has %d validation error(s)", len(result.Errors))
//	}
//
// Printing a result does not fail the caller: PrintValidationResult returns a
// nil error when it successfully writes the result, even when the manifest is
// invalid. Warnings are reported but do not fail validation unless the caller
// treats them as errors.
//
// ValidateFile is a convenience wrapper for callers that do not need to supply
// their own filesystem:
//
//	result, err := validate.ValidateFile("custom-elements.json", validate.ValidationOptions{
//		IncludeWarnings: true,
//	})
//
// The filesystem-aware entry points use narrow interfaces, so embedders
// implement only what the entry point needs. Schema accepts a ReadFS, which is
// just ReadFile and Stat. Validate and NewValidationPipeline accept a
// ValidateFS, which embeds ReadFS and adds WriteFile, MkdirAll, CreateTemp,
// and Remove. TempFile remains exported for CreateTemp implementations. Use
// NewOSFileSystem for the host operating system, or implement the interface
// for in-memory or sandboxed input. For finer control, build a
// ValidationPipeline with NewValidationPipeline and call its Validate method
// directly.
//
// Schema fetches a schema from unpkg when the version is neither embedded nor
// cached. Callers that must not use the network can pass WithoutNetwork to
// Schema, or set ValidationOptions.NoNetwork when using Validate. Both still
// resolve embedded and cached schemas.
//
// Warning rules for manifest content can be disabled by ID with
// ValidationOptions.DisabledRules. Schema-version warnings are always
// reported, independent of IncludeWarnings and DisabledRules, because they
// describe the validation itself. Use PrintValidationResult for human-readable
// or JSON output, or inspect ValidationResult.Errors and
// ValidationResult.Warnings directly.
package validate
