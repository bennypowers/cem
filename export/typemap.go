/*
Copyright © 2025 Benny Powers <web@bennypowers.com>

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
package export

import "strings"

// MapCEMType maps a CEM type string to a TypeScript type for use in generated wrappers.
// CEM types are already valid TypeScript, so this is mostly a passthrough.
// Empty types map to "unknown".
func MapCEMType(cemType string) string {
	t := strings.TrimSpace(cemType)
	if t == "" {
		return "unknown"
	}
	return t
}

// IsBooleanType returns true if the CEM type is a boolean type.
// Used to determine attribute handling (boolean attributes vs string attributes).
func IsBooleanType(cemType string) bool {
	t := strings.TrimSpace(cemType)
	return t == "boolean"
}
