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

package importmap

import (
	"encoding/json"

	"bennypowers.dev/cem/internal/platform"
	"bennypowers.dev/cem/serve/middleware"
	"bennypowers.dev/cem/serve/middleware/types"
)

// ImportMap represents an ES module import map
type ImportMap struct {
	Imports map[string]string            `json:"imports"`
	Scopes  map[string]map[string]string `json:"scopes,omitempty"`
}

// IsImportMap is a marker method to implement the middleware.ImportMap interface
func (*ImportMap) IsImportMap() {}

// ToJSON converts the import map to indented JSON string (returns empty string if nil or empty)
func (im *ImportMap) ToJSON() string {
	if im == nil || len(im.Imports) == 0 {
		return ""
	}

	bytes, err := json.MarshalIndent(im, "  ", "  ")
	if err != nil {
		return ""
	}

	return string(bytes)
}

// Config configures import map generation
type Config struct {
	InputMapPath      string                        // Path to user override file
	ConfigOverride    *ImportMap                    // Config file overrides (imports and scopes) - uses ImportMap directly
	Logger            types.Logger                  // Logger for warnings
	WorkspacePackages []middleware.WorkspacePackage // If set, generate workspace-mode import map (flattened scopes)
	FS                platform.FileSystem           // Filesystem abstraction (defaults to OS filesystem if nil)
}
