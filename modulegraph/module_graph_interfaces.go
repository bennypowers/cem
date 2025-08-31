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
package modulegraph

import (
	"path/filepath"
	"time"

	"bennypowers.dev/cem/queries"
)

// DefaultMaxTransitiveDepth is the default maximum depth for transitive closure computation
// to prevent performance issues with deeply nested dependency chains.
//
// This limit of 5 levels was chosen based on practical analysis of real-world projects:
// - Most legitimate dependency chains are 2-3 levels deep
// - 5 levels covers complex library architectures while preventing pathological cases
// - Protects against circular dependencies and runaway calculations
// - Balances comprehensive element discovery with acceptable performance
const DefaultMaxTransitiveDepth = 5

// FileParser interface abstracts file reading and workspace traversal operations
// for improved testability and dependency injection
type FileParser interface {
	// ReadFile reads the content of a file at the given path
	ReadFile(path string) ([]byte, error)
	// WalkWorkspace walks through a workspace directory calling the walkFn for each file
	WalkWorkspace(workspaceRoot string, walkFn filepath.WalkFunc) error
}

// ExportParser interface abstracts the parsing of export statements from content
// This provides a higher-level abstraction over tree-sitter operations
type ExportParser interface {
	// ParseExportsFromContent parses export statements from file content
	// Returns any parsing errors, but implementations should be resilient
	ParseExportsFromContent(modulePath string, content []byte, exportTracker *ExportTracker, dependencyTracker *DependencyTracker, queryManager *queries.QueryManager) error
}

// ManifestResolver interface abstracts manifest-based path resolution
// This allows the module graph to integrate with existing manifest path resolution
type ManifestResolver interface {
	// FindManifestModulesForImportPath finds manifest modules that match an import path
	// Returns module paths from the manifest (e.g., "rh-tabs/rh-tab.js")
	FindManifestModulesForImportPath(importPath string) []string

	// GetManifestModulePath converts a file path to its corresponding manifest module path
	// Returns empty string if no manifest module exists for this file
	GetManifestModulePath(filePath string) string

	// GetElementsFromManifestModule returns all custom element tag names available from a manifest module
	// This is the key method that maps manifest modules to their available custom elements
	GetElementsFromManifestModule(manifestModulePath string) []string
}

// MetricsCollector interface abstracts metrics collection for observability
type MetricsCollector interface {
	// IncrementCounter increments a named counter metric
	IncrementCounter(name string)
	// RecordDuration records a duration metric in milliseconds
	RecordDuration(name string, duration time.Duration)
	// SetGauge sets a gauge metric value
	SetGauge(name string, value int64)
	// AddHistogramValue adds a value to a histogram metric
	AddHistogramValue(name string, value float64)
}
