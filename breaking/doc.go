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

// Package breaking compares the custom element declarations in two Custom
// Elements Manifest documents and classifies changes to their public API as
// safe, dangerous, or breaking.
//
// Rules cover element additions and removals, plus the attributes, slots,
// events, CSS custom properties, parts, states, and members of each custom
// element. Compare indexes tagged custom element declarations, so it does not
// report changes to other declaration kinds, such as exported functions. Each
// detected change carries a rule ID, a severity, and a human-readable message.
// Use Options.Disable to silence rule IDs that do not apply to a project.
//
// # External use
//
// This package is intended for external consumption and is part of cem's
// public Go API. It follows the 0.x compatibility policy described in
// CONTRIBUTING.md: treat releases as prerelease, pin an exact version, and
// upgrade deliberately.
//
// # Comparing manifests
//
//	var base, head manifest.Package
//	if err := json.Unmarshal(baseData, &base); err != nil {
//		return err
//	}
//	if err := json.Unmarshal(headData, &head); err != nil {
//		return err
//	}
//
//	result := breaking.Compare(&base, &head, breaking.Options{})
//	if result.Breaking > 0 {
//		return breaking.PrintResult(os.Stdout, result, breaking.DisplayOptions{Format: "text"})
//	}
//
// PrintResult supports "text", "markdown", and "json" output. The Result
// value is also directly inspectable through its Changes field.
package breaking
