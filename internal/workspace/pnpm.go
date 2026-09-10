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

package workspace

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"bennypowers.dev/cem/internal/platform"
	"gopkg.in/yaml.v3"
)

// pnpmWorkspace represents the parsed structure of pnpm-workspace.yaml
type pnpmWorkspace struct {
	Packages []string `yaml:"packages"`
}

// ReadPnpmWorkspace reads and parses a pnpm-workspace.yaml file.
// Returns nil (no error) if the file does not exist.
func ReadPnpmWorkspace(rootDir string, fsys platform.FileSystem) (*pnpmWorkspace, error) {
	path := filepath.Join(rootDir, "pnpm-workspace.yaml")
	data, err := fsys.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, nil
		}
		return nil, fmt.Errorf("reading pnpm-workspace.yaml at %s: %w", path, err)
	}

	var ws pnpmWorkspace
	if err := yaml.Unmarshal(data, &ws); err != nil {
		return nil, fmt.Errorf("parsing pnpm-workspace.yaml at %s: %w", path, err)
	}

	return &ws, nil
}

// DiscoverPnpmPackages discovers workspace packages from a pnpm-workspace.yaml file.
// Returns map of package name -> absolute path to package directory.
func DiscoverPnpmPackages(rootDir string, ws *pnpmWorkspace, fsys platform.FileSystem) (map[string]string, error) {
	if ws == nil {
		return nil, nil
	}
	return DiscoverWorkspacePackagesFromPatterns(rootDir, ws.Packages, fsys)
}

// HasPnpmWorkspace checks if a pnpm-workspace.yaml file exists in the given directory.
func HasPnpmWorkspace(dir string, fsys platform.FileSystem) bool {
	ws, err := ReadPnpmWorkspace(dir, fsys)
	return err == nil && ws != nil && len(ws.Packages) > 0
}
