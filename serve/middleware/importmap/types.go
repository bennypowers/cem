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

// DependencyGraph tracks package dependencies for incremental updates.
// This is an opaque wrapper around mappa's DependencyGraph.
type DependencyGraph struct {
	graph any // *resolve.DependencyGraph from mappa
}

// IncrementalResult contains both the import map and dependency graph.
// Use this for initial resolution when you plan to do incremental updates.
type IncrementalResult struct {
	// ImportMap is the generated import map.
	ImportMap *ImportMap

	// Graph tracks package dependencies for future incremental updates.
	Graph *DependencyGraph
}

// IncrementalUpdate describes which packages changed for incremental resolution.
type IncrementalUpdate struct {
	// ChangedPackages lists the package names that were modified.
	// Can be package names from node_modules (e.g., "lit") or workspace packages.
	ChangedPackages []string

	// PreviousMap is the import map from the last resolution.
	PreviousMap *ImportMap

	// PreviousGraph is the dependency graph from the last resolution.
	PreviousGraph *DependencyGraph
}
