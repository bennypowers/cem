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
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	C "bennypowers.dev/cem/cmd/config"
	DT "bennypowers.dev/cem/designtokens"
	M "bennypowers.dev/cem/manifest"
	"bennypowers.dev/cem/types"
	"github.com/bmatcuk/doublestar"
	"github.com/pterm/pterm"
	"gopkg.in/yaml.v3"
)

var _ types.WorkspaceContext = (*FileSystemWorkspaceContext)(nil)

// FileSystemWorkspaceContext implements WorkspaceContext for a local filesystem
// package.
type FileSystemWorkspaceContext struct {
	root                       string
	config                     *C.CemConfig
	customElementsManifestPath string
	// Cache parsed results if desired
	manifest          *M.Package
	packageJSON       *M.PackageJSON
	designTokensCache types.DesignTokensCache
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
		defer func() { _ = rc.Close() }()
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
	return &FileSystemWorkspaceContext{
		root:              root,
		designTokensCache: NewDesignTokensCache(DT.NewLoader()),
	}
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
	defer func() { _ = rc.Close() }()
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
	defer func() { _ = rc.Close() }()

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
	// Clean and normalize the pattern first
	pattern = filepath.Clean(pattern)

	if isGlobPattern(pattern) {
		// Determine the base directory for resolving relative patterns.
		// In monorepos, patterns in a package config should resolve relative
		// to the package directory (where the config is), not the monorepo root.
		baseDir := c.root
		if c.config != nil && c.config.ProjectDir != "" {
			// Use ProjectDir from config if available (handles package subdirectories)
			baseDir = c.config.ProjectDir
		}

		// Ensure baseDir is absolute
		if !filepath.IsAbs(baseDir) {
			cwd, err := os.Getwd()
			if err != nil {
				return nil, fmt.Errorf("failed to get current directory: %w", err)
			}
			baseDir = filepath.Join(cwd, baseDir)
		}

		// For glob patterns, join with base directory and execute glob
		globPath := filepath.Join(baseDir, pattern)
		result, err := doublestar.Glob(globPath)
		if err != nil {
			return nil, fmt.Errorf("glob pattern %q failed: %w", pattern, err)
		}

		// Convert absolute paths back to relative paths within base directory
		relativeResult := make([]string, 0, len(result))
		for _, absPath := range result {
			rel, err := filepath.Rel(baseDir, absPath)
			if err != nil {
				// Skip files that can't be made relative
				continue
			}
			// Skip files outside base directory
			if strings.HasPrefix(rel, ".."+string(filepath.Separator)) || rel == ".." {
				continue
			}
			relativeResult = append(relativeResult, rel)
		}
		return relativeResult, nil
	}

	// For non-glob patterns, handle as individual file paths
	return c.handleNonGlobPattern(pattern)
}

// makeRelativeToRoot converts an absolute path to relative path within project root.
// Returns error if path is outside project root.
func (c *FileSystemWorkspaceContext) makeRelativeToRoot(absPath string) (string, error) {
	// Clean both paths to ensure consistent separators
	cleanRoot := filepath.Clean(c.root)
	cleanPath := filepath.Clean(absPath)

	rel, err := filepath.Rel(cleanRoot, cleanPath)
	if err != nil {
		return "", fmt.Errorf("failed to compute relative path: %w", err)
	}

	// Check if path escapes project root (cross-platform check)
	if strings.HasPrefix(rel, ".."+string(filepath.Separator)) || rel == ".." {
		return "", fmt.Errorf("path %q is outside project root %q", absPath, c.root)
	}

	return rel, nil
}

// handleNonGlobPattern processes individual file patterns (not glob patterns)
func (c *FileSystemWorkspaceContext) handleNonGlobPattern(pattern string) ([]string, error) {
	if filepath.IsAbs(pattern) {
		// Convert absolute path to relative if within project
		if rel, err := c.makeRelativeToRoot(pattern); err == nil {
			return []string{rel}, nil
		}
		// If outside project root, return empty (don't process external files)
		return []string{}, nil
	}

	// For relative patterns, ensure they resolve within project
	absPath := filepath.Join(c.root, pattern)
	if rel, err := c.makeRelativeToRoot(absPath); err == nil {
		return []string{rel}, nil
	}

	// Fallback: return the original pattern if path resolution fails
	// This maintains backward compatibility for edge cases
	return []string{filepath.Clean(pattern)}, nil
}

func (c *FileSystemWorkspaceContext) OutputWriter(path string) (io.WriteCloser, error) {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}
	pterm.Debug.Printfln("Output: %q", path)
	return os.Create(path)
}

func (c *FileSystemWorkspaceContext) Root() string {
	return c.root
}

// DesignTokensCache returns the design tokens cache for this workspace
func (c *FileSystemWorkspaceContext) DesignTokensCache() types.DesignTokensCache {
	return c.designTokensCache
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
func (c *FileSystemWorkspaceContext) ResolveModuleDependency(
	modulePath,
	dependencyPath string,
) (string, error) {
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

// FindWorkspaceRoot searches upward from the given path to find the workspace root.
// It looks for workspace indicators like .git, pnpm-workspace.yaml, or package.json with workspaces field.
// If the path is already inside a workspace package subdirectory, it will find the parent workspace root.
//
// VCS Boundary Handling:
// - If a directory has ONLY VCS markers (.git) without workspace metadata, treat it as a hard boundary
// - This prevents crossing Git submodule boundaries
// - If a directory has workspace metadata (pnpm-workspace.yaml, package.json with workspaces),
//   continue climbing to find the topmost workspace root
func FindWorkspaceRoot(startPath string) (string, error) {
	// Resolve to absolute path
	absPath, err := filepath.Abs(startPath)
	if err != nil {
		return "", fmt.Errorf("failed to resolve path: %w", err)
	}

	// Ensure we're starting from a directory
	info, err := os.Stat(absPath)
	if err != nil {
		return "", fmt.Errorf("failed to stat path: %w", err)
	}
	if !info.IsDir() {
		absPath = filepath.Dir(absPath)
	}

	current := absPath
	checked := make(map[string]bool)

	// Track the last directory we found with any workspace indicator
	var lastRootCandidate string

	for !checked[current] {
		checked[current] = true

		// Check if current directory has VCS marker
		hasVCS := isVCSRoot(current)
		// Check if current directory has workspace metadata
		hasWorkspaceMeta := hasWorkspaceMetadata(current)

		// If we have either VCS marker or workspace metadata, this is a candidate root
		if hasVCS || hasWorkspaceMeta {
			// If we ONLY have VCS marker (no workspace metadata), treat as hard boundary
			// This prevents crossing Git submodule boundaries
			if hasVCS && !hasWorkspaceMeta {
				return current, nil
			}

			// We have workspace metadata - record this as a candidate
			// but continue climbing to see if there's a parent workspace
			lastRootCandidate = current
		}

		// Move to parent directory
		parent := filepath.Dir(current)
		if parent == current {
			// Reached filesystem root
			if lastRootCandidate != "" {
				return lastRootCandidate, nil
			}
			return absPath, nil // Return original path as fallback
		}
		current = parent
	}

	// If we exhausted all directories and found a candidate, return it
	if lastRootCandidate != "" {
		return lastRootCandidate, nil
	}

	return absPath, nil // Return original path as fallback
}

// isVCSRoot checks if a directory contains VCS markers (version control root)
// Currently only checks for .git directory
func isVCSRoot(dir string) bool {
	// Check for .git directory
	if _, err := os.Stat(filepath.Join(dir, ".git")); err == nil {
		return true
	}
	return false
}

// hasWorkspaceMetadata checks if a directory contains workspace metadata files
// These indicate actual workspace structure (not just VCS)
func hasWorkspaceMetadata(dir string) bool {
	// Check for pnpm-workspace.yaml
	if _, err := os.Stat(filepath.Join(dir, "pnpm-workspace.yaml")); err == nil {
		return true
	}

	// Check for package.json with workspaces field
	packageJSONPath := filepath.Join(dir, "package.json")
	if data, err := os.ReadFile(packageJSONPath); err == nil {
		var pkg struct {
			Workspaces interface{} `json:"workspaces"`
		}
		if err := json.Unmarshal(data, &pkg); err == nil && pkg.Workspaces != nil {
			return true
		}
	}

	return false
}
