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
	"io"
	"os"
	"path/filepath"

	C "bennypowers.dev/cem/cmd/config"
	M "bennypowers.dev/cem/manifest"
	"github.com/bmatcuk/doublestar"
	"github.com/pterm/pterm"
	"gopkg.in/yaml.v3"
)

var _ WorkspaceContext = (*FileSystemWorkspaceContext)(nil)

// FileSystemWorkspaceContext implements WorkspaceContext for a local filesystem
// package.
type FileSystemWorkspaceContext struct {
	root                       string
	config                     *C.CemConfig
	manifestPath               string
	packageJSONPath            string
	customElementsManifestPath string
	// Cache parsed results if desired
	manifest    *M.Package
	packageJSON *M.PackageJSON
}

func (c *FileSystemWorkspaceContext) initConfig() (*C.CemConfig, error) {
	var config C.CemConfig
	config.ProjectDir = c.Root()
	config.ConfigFile = c.ConfigFile()

	if _, err := os.Stat(config.ConfigFile); os.IsNotExist(err) {
		pterm.Debug.Println("no config file found")
	} else if err != nil {
		return nil, err
	} else {
		rc, err := c.ReadFile(config.ConfigFile)
		if err != nil {
			return nil, err
		}
		defer rc.Close()
		if err := yaml.NewDecoder(rc).Decode(&config); err != nil {
			return nil, err
		}
	}
	// Make output path package-root-relative if needed
	if config.Generate.Output != "" && !filepath.IsAbs(config.Generate.Output) {
		config.Generate.Output = filepath.Join(config.ProjectDir, config.Generate.Output)
	}
	if config.Generate.Files == nil {
		config.Generate.Files = make([]string, 0)
	}
	if config.Generate.Exclude == nil {
		config.Generate.Exclude = make([]string, 0)
	}
	// Set debug verbosity
	if config.Verbose {
		pterm.EnableDebugMessages()
	} else {
		pterm.DisableDebugMessages()
	}

	return &config, nil
}

func NewFileSystemWorkspaceContext(root string) *FileSystemWorkspaceContext {
	return &FileSystemWorkspaceContext{root: root}
}

// ConfigFile Returns the path to the config file, or an empty string if it does not exist.
func (c *FileSystemWorkspaceContext) ConfigFile() string {
	return filepath.Join(c.root, ".config", "cem.yaml")
}

func (c *FileSystemWorkspaceContext) PackageJSON() (*M.PackageJSON, error) {
	if c.packageJSON != nil {
		return c.packageJSON, nil
	}
	rc, err := c.ReadFile(filepath.Join(c.root, "package.json"))
	if errors.Is(err, os.ErrNotExist) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	defer rc.Close()
	p, err := decodeJSON[M.PackageJSON](rc)
	if err != nil {
		return nil, err
	}
	c.packageJSON = p
	return p, err
}

func (c *FileSystemWorkspaceContext) Config() (*C.CemConfig, error) {
	return c.config, nil
}

func (c *FileSystemWorkspaceContext) CustomElementsManifestPath() string {
	return c.customElementsManifestPath
}

func (c *FileSystemWorkspaceContext) Manifest() (*M.Package, error) {
	if c.manifest != nil {
		return c.manifest, nil
	}

	if c.customElementsManifestPath == "" {
		return nil, ErrNoManifest
	}

	rc, err := c.ReadFile(c.customElementsManifestPath)
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	m, err := decodeJSON[M.Package](rc)
	if err != nil {
		return nil, err
	}

	c.manifest = m
	return m, nil
}

// Init discovers package.json file, caches paths/parsed results.
func (c *FileSystemWorkspaceContext) Init() error {
	var err error
	c.config, err = c.initConfig()
	if err != nil {
		return err
	}

	// Discover package.json
	pkg, err := c.PackageJSON()
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}

	// Try to get the manifest path from package.json first
	if pkg != nil && pkg.CustomElements != "" {
		c.customElementsManifestPath = filepath.Join(c.root, pkg.CustomElements)
	} else if c.config != nil && c.config.Generate.Output != "" {
		// If that fails, try to get it from the config
		c.customElementsManifestPath = c.config.Generate.Output
	} else {
		// Fallback to default
		c.customElementsManifestPath = filepath.Join(c.root, "custom-elements.json")
	}

	return nil
}

func (c *FileSystemWorkspaceContext) ReadFile(path string) (io.ReadCloser, error) {
	rc, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("FileSystemWorkspaceContext could not open: %w", err)
	}
	return rc, nil
}

func (c *FileSystemWorkspaceContext) Glob(pattern string) ([]string, error) {
	if isGlobPattern(pattern) {
		return doublestar.Glob(filepath.Join(c.root, pattern))
	} else {
		return []string{pattern}, nil
	}
}

func (c *FileSystemWorkspaceContext) OutputWriter(path string) (io.WriteCloser, error) {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}
	pterm.Debug.Println("Actual output path", path)
	return os.Create(path)
}

func (c *FileSystemWorkspaceContext) Root() string {
	return c.root
}

func (c *FileSystemWorkspaceContext) Cleanup() error {
	// Nothing to clean up for local packages
	return nil
}

// ModulePathToFS converts a module path to filesystem path for watching
func (c *FileSystemWorkspaceContext) ModulePathToFS(modulePath string) string {
	if filepath.IsAbs(modulePath) {
		return modulePath
	}
	return filepath.Join(c.root, modulePath)
}

// FSPathToModule converts a filesystem path to module path for manifest lookup
func (c *FileSystemWorkspaceContext) FSPathToModule(fsPath string) (string, error) {
	rel, err := filepath.Rel(c.root, fsPath)
	if err != nil {
		return "", fmt.Errorf("failed to compute relative path from %s to %s: %w", c.root, fsPath, err)
	}
	return rel, nil
}

// ResolveModuleDependency resolves a dependency path relative to a module
func (c *FileSystemWorkspaceContext) ResolveModuleDependency(modulePath, dependencyPath string) (string, error) {
	if filepath.IsAbs(dependencyPath) {
		// Convert to module-relative path
		rel, err := filepath.Rel(c.root, dependencyPath)
		if err != nil {
			return "", fmt.Errorf("failed to resolve absolute dependency path %s relative to %s: %w", dependencyPath, c.root, err)
		}
		return rel, nil
	}

	// Resolve relative to module's directory
	moduleDir := filepath.Dir(modulePath)
	resolved := filepath.Join(moduleDir, dependencyPath)
	return filepath.Clean(resolved), nil
}
