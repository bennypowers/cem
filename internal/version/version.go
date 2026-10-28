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

// Package version is an internal compatibility shim over the public
// bennypowers.dev/cem/version package. New code should import the public
// package directly.
package version

import publicversion "bennypowers.dev/cem/version"

// GetVersion returns the version string for the application.
func GetVersion() string {
	return publicversion.Get()
}

// GetFullVersion returns detailed version information.
func GetFullVersion() string {
	return publicversion.Full()
}

// GetBuildInfo returns detailed build information.
func GetBuildInfo() map[string]string {
	return publicversion.Info()
}
