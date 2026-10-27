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

// Package manifest provides Go types for the Custom Elements Manifest (CEM)
// schema, together with helpers for decoding, querying, cloning, and encoding
// manifests.
//
// The types mirror schema 2.1.x from
// https://github.com/webcomponents/custom-elements-manifest. Properties that
// are not part of the schema are ignored while decoding and are not preserved,
// so a manifest that adds unrecognized properties still decodes. Unknown
// declaration, export, and class member kinds are not supported and cause a
// decoding error.
//
// # External use
//
// This package is intended for external consumption and is part of cem's
// public Go API. It follows the 0.x compatibility policy described in
// CONTRIBUTING.md: treat releases as prerelease, pin an exact version, and
// upgrade deliberately.
//
// # Decoding a manifest
//
//	var pkg manifest.Package
//	if err := json.Unmarshal(data, &pkg); err != nil {
//		return err
//	}
//
//	element := pkg.FindCustomElementDeclaration("my-element")
//	if element != nil {
//		fmt.Println(element.TagName, element.Summary)
//	}
//
// # Building and encoding a manifest
//
//	pkg := manifest.NewPackage([]manifest.Module{
//		*manifest.NewModule("src/my-element.ts"),
//	})
//	data, err := manifest.SerializeToBytes(&pkg)
//
// # Copying manifests
//
// Manifests can contain large graphs of declarations with backreferences.
// Use Clone for deep copies instead of JSON round-trips when performance
// matters:
//
//	cloned := pkg.Clone()
package manifest
