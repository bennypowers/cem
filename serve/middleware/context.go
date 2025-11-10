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

package middleware

import (
	"bennypowers.dev/cem/serve/logger"
)

// Logger is a type alias for the logger.Logger interface
type Logger = logger.Logger

// ImportMap is a marker interface for import map types
type ImportMap interface {
	IsImportMap()
}

// WorkspacePackage represents a package in workspace mode
type WorkspacePackage struct {
	Name     string // Package name from package.json
	Path     string // Absolute path to package directory
	Manifest []byte // Generated custom elements manifest
}

// PackageJSON represents parsed package.json data
type PackageJSON struct {
	Name    string
	Version string
}

// DevServerContext provides middlewares access to dev server state and functionality
type DevServerContext interface {
	// WatchDir returns the directory being watched for changes
	WatchDir() string

	// IsWorkspace returns true if running in workspace/monorepo mode
	IsWorkspace() bool

	// WorkspacePackages returns all discovered workspace packages (workspace mode only)
	WorkspacePackages() []WorkspacePackage

	// Manifest returns the current custom elements manifest (single-package mode)
	Manifest() ([]byte, error)

	// ImportMap returns the pre-computed import map (may be nil)
	ImportMap() ImportMap

	// DemoRoutes returns the pre-computed demo routing table (both workspace and single-package mode)
	// The return type is map[string]*routes.DemoRouteEntry but we use any to avoid circular imports
	DemoRoutes() any

	// Logger returns the server's logger
	Logger() Logger

	// PackageJSON returns parsed package.json (single-package mode only)
	PackageJSON() (*PackageJSON, error)
}
