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

package types

import (
	"io"

	C "bennypowers.dev/cem/cmd/config"
	M "bennypowers.dev/cem/manifest"
)

// WorkspaceContext abstracts access to package resources, regardless of source (local or remote).
type WorkspaceContext interface {
	// Performs validation/discovery and caches results as needed.
	Init() error
	// ConfigFile Returns the path to the config file
	ConfigFile() string
	// Config returns the parsed and initialized config
	Config() (*C.CemConfig, error)
	// Returns the package's parsed PackageJSON.
	PackageJSON() (*M.PackageJSON, error)
	// Manifest returns the package's parsed custom elements manifest.
	Manifest() (*M.Package, error)
	// CustomElementsManifestPath returns the path to the manifest file.
	CustomElementsManifestPath() string
	// ReadFile returns an io.ReadCloser for a file within the package.
	ReadFile(path string) (io.ReadCloser, error)
	// Glob returns a list of file paths matching the given pattern (e.g., *.ts).
	Glob(pattern string) ([]string, error)
	// Writes outputs to paths
	OutputWriter(path string) (io.WriteCloser, error)
	// Root returns the canonical root path or name for the package.
	Root() string
	// Cleanup releases any resources (e.g., tempdirs) held by the context.
	Cleanup() error

	// Path resolution utilities for consistent module/filesystem path mapping
	// ModulePathToFS converts a module path to filesystem path for watching
	ModulePathToFS(modulePath string) string
	// FSPathToModule converts a filesystem path to module path for manifest lookup
	FSPathToModule(fsPath string) (string, error)
	// ResolveModuleDependency resolves a dependency path relative to a module
	ResolveModuleDependency(modulePath, dependencyPath string) (string, error)

	// DesignTokensCache returns the design tokens cache for this workspace
	DesignTokensCache() DesignTokensCache
}
