/*
Copyright © 2026 Benny Powers <web@bennypowers.com>

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
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"testing"

	C "bennypowers.dev/cem/cmd/config"
	IC "bennypowers.dev/cem/internal/config"
	"bennypowers.dev/cem/internal/platform"
	"bennypowers.dev/cem/internal/platform/testutil"
	M "bennypowers.dev/cem/manifest"
	"bennypowers.dev/cem/types"
	"github.com/bmatcuk/doublestar/v4"
)

var _ types.WorkspaceContext = (*MapWorkspaceContext)(nil)
var _ types.CacheInvalidator = (*MapWorkspaceContext)(nil)

// MapWorkspaceContext implements types.WorkspaceContext backed by a platform.MapFileSystem.
// All reads come from in-memory MapFS, no disk I/O after construction.
type MapWorkspaceContext struct {
	mfs                        *platform.MapFileSystem
	root                       string
	config                     *C.CemConfig
	configErr                  error
	configMu                   sync.RWMutex
	customElementsManifestPath string
	manifest                   *M.Package
	manifestMu                 sync.RWMutex
	packageJSON                *M.PackageJSON
	designTokensCache          types.DesignTokensCache
}

// NewMapWorkspaceContext loads dir into MapFS and returns a WorkspaceContext backed by it.
// dir is relative to the test's working directory.
func NewMapWorkspaceContext(t testing.TB, dir string) *MapWorkspaceContext {
	t.Helper()
	mfs := testutil.LoadTestdataFS(t, dir, "/")
	return &MapWorkspaceContext{
		mfs:               mfs,
		root:              "/",
		designTokensCache: &noopDesignTokensCache{},
	}
}

// NewMapWorkspaceContextFromFS creates a MapWorkspaceContext from an existing MapFileSystem.
// Use when building the filesystem directly with AddFile instead of loading from disk.
func NewMapWorkspaceContextFromFS(mfs *platform.MapFileSystem, root string) *MapWorkspaceContext {
	return &MapWorkspaceContext{
		mfs:               mfs,
		root:              root,
		designTokensCache: &noopDesignTokensCache{},
	}
}

// NewMapWorkspaceContextWithRoot loads dir into MapFS rooted at rootPath
// and returns a WorkspaceContext backed by it. Use this when the test depends
// on the workspace root directory name (e.g., workspace fallback logic that
// uses filepath.Base(root) as a package name).
func NewMapWorkspaceContextWithRoot(t testing.TB, dir string, rootPath string) *MapWorkspaceContext {
	t.Helper()
	mfs := testutil.LoadTestdataFS(t, dir, rootPath)
	return &MapWorkspaceContext{
		mfs:               mfs,
		root:              rootPath,
		designTokensCache: &noopDesignTokensCache{},
	}
}

func (c *MapWorkspaceContext) Init() error {
	data, err := c.mfs.ReadFile(filepath.Join(c.root, "package.json"))
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		return fmt.Errorf("MapWorkspaceContext: reading package.json: %w", err)
	}

	if err == nil {
		var pkg M.PackageJSON
		if err := json.Unmarshal(data, &pkg); err != nil {
			return fmt.Errorf("MapWorkspaceContext: parsing package.json: %w", err)
		}
		c.packageJSON = &pkg

		if pkg.CustomElements != "" {
			c.customElementsManifestPath = filepath.Join(c.root, pkg.CustomElements)
		}
	}

	if c.customElementsManifestPath == "" {
		c.customElementsManifestPath = filepath.Join(c.root, "custom-elements.json")
	}

	// Discover and load config from MapFS, mirroring FileSystemWorkspaceContext.initConfig
	if err := c.loadConfig(); err != nil {
		return err
	}

	return nil
}

func (c *MapWorkspaceContext) loadConfig() error {
	if configPath := IC.FindConfigFile(c.root, c.mfs); configPath != "" {
		cfg, err := IC.LoadConfig(configPath, c.mfs)
		if err != nil {
			return fmt.Errorf("MapWorkspaceContext: loading config: %w", err)
		}
		cfg.ProjectDir = c.root
		if cfg.Generate.Files == nil {
			cfg.Generate.Files = make([]string, 0)
		}
		if cfg.Generate.Exclude == nil {
			cfg.Generate.Exclude = make([]string, 0)
		}
		c.config = cfg
		return nil
	}

	c.config = &IC.CemConfig{
		ProjectDir: c.root,
		Generate: IC.GenerateConfig{
			Files:   make([]string, 0),
			Exclude: make([]string, 0),
		},
	}
	return nil
}

func (c *MapWorkspaceContext) ConfigFile() string {
	return ""
}

func (c *MapWorkspaceContext) Config() (*C.CemConfig, error) {
	c.configMu.RLock()
	defer c.configMu.RUnlock()
	return c.config, c.configErr
}

// InvalidateConfig re-reads the configuration from the in-memory filesystem.
// On failure the previous config is cleared, matching FileSystemWorkspaceContext.
func (c *MapWorkspaceContext) InvalidateConfig() {
	c.configMu.Lock()
	defer c.configMu.Unlock()
	if err := c.loadConfig(); err != nil {
		c.config = nil
		c.configErr = err
		return
	}
	c.configErr = nil
}

func (c *MapWorkspaceContext) PackageJSON() (*M.PackageJSON, error) {
	return c.packageJSON, nil
}

// Manifest returns the parsed custom elements manifest, caching it after the
// first successful read. InvalidateManifest clears the cache so the next call
// re-reads the manifest.
func (c *MapWorkspaceContext) Manifest() (*M.Package, error) {
	c.manifestMu.RLock()
	if c.manifest != nil {
		manifest := c.manifest
		c.manifestMu.RUnlock()
		return manifest, nil
	}
	c.manifestMu.RUnlock()

	c.manifestMu.Lock()
	defer c.manifestMu.Unlock()

	if c.manifest != nil {
		return c.manifest, nil
	}

	if c.customElementsManifestPath == "" {
		return nil, errors.New("no manifest path configured")
	}

	rc, err := c.ReadFile(c.customElementsManifestPath)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rc.Close() }()

	var pkg M.Package
	if err := json.NewDecoder(rc).Decode(&pkg); err != nil {
		return nil, fmt.Errorf("MapWorkspaceContext: parsing manifest: %w", err)
	}

	c.manifest = &pkg
	return c.manifest, nil
}

// InvalidateManifest clears the cached manifest so the next Manifest call
// re-reads it. It is safe to call concurrently with Manifest and other
// readers.
func (c *MapWorkspaceContext) InvalidateManifest() {
	c.manifestMu.Lock()
	c.manifest = nil
	c.manifestMu.Unlock()
}

func (c *MapWorkspaceContext) CustomElementsManifestPath() string {
	return c.customElementsManifestPath
}

func (c *MapWorkspaceContext) ReadFile(path string) (io.ReadCloser, error) {
	clean := c.cleanPath(path)
	data, err := c.mfs.ReadFile(clean)
	if err != nil {
		return nil, fmt.Errorf("MapWorkspaceContext: %w", err)
	}
	return io.NopCloser(bytes.NewReader(data)), nil
}

func (c *MapWorkspaceContext) Stat(path string) (fs.FileInfo, error) {
	clean := c.cleanPath(path)
	return c.mfs.Stat(clean)
}

func (c *MapWorkspaceContext) ReadDir(path string) ([]fs.DirEntry, error) {
	clean := c.cleanPath(path)
	return c.mfs.ReadDir(clean)
}

func (c *MapWorkspaceContext) Glob(pattern string) ([]string, error) {
	mapFS := c.mfs.GetMapFS()
	seen := make(map[string]bool)
	var matches []string

	for filePath := range mapFS {
		// Match the file itself
		if matched, _ := doublestar.Match(pattern, filePath); matched {
			if !seen[filePath] {
				seen[filePath] = true
				matches = append(matches, filePath)
			}
		}
		// Match intermediate directories derived from file paths
		dir := filePath
		for {
			dir = filepath.Dir(dir)
			if dir == "." || dir == "/" || dir == "" {
				break
			}
			if seen[dir] {
				break
			}
			if matched, _ := doublestar.Match(pattern, dir); matched {
				seen[dir] = true
				matches = append(matches, dir)
			}
		}
	}
	sort.Strings(matches)
	return matches, nil
}

func (c *MapWorkspaceContext) OutputWriter(path string) (io.WriteCloser, error) {
	return &mapFSWriter{mfs: c.mfs, path: path}, nil
}

func (c *MapWorkspaceContext) Root() string {
	return c.root
}

func (c *MapWorkspaceContext) Cleanup() error {
	return nil
}

func (c *MapWorkspaceContext) ModulePathToFS(modulePath string) string {
	if filepath.IsAbs(modulePath) {
		return modulePath
	}
	return filepath.Join(c.root, modulePath)
}

func (c *MapWorkspaceContext) FSPathToModule(fsPath string) (string, error) {
	rel, err := filepath.Rel(c.root, fsPath)
	if err != nil {
		return "", fmt.Errorf("failed to compute relative path from %s to %s: %w", c.root, fsPath, err)
	}
	return rel, nil
}

func (c *MapWorkspaceContext) ResolveModuleDependency(modulePath, dependencyPath string) (string, error) {
	if filepath.IsAbs(dependencyPath) {
		rel, err := filepath.Rel(c.root, dependencyPath)
		if err != nil {
			return "", err
		}
		return rel, nil
	}
	moduleDir := filepath.Dir(modulePath)
	resolved := filepath.Join(moduleDir, dependencyPath)
	return filepath.Clean(resolved), nil
}

func (c *MapWorkspaceContext) DesignTokensCache() types.DesignTokensCache {
	return c.designTokensCache
}

// FileSystem returns the underlying MapFileSystem for use as a platform.FileSystem
// argument in functions that require both a WorkspaceContext and a FileSystem.
func (c *MapWorkspaceContext) FileSystem() *platform.MapFileSystem {
	return c.mfs
}

func (c *MapWorkspaceContext) cleanPath(path string) string {
	return strings.TrimPrefix(filepath.Clean(path), "/")
}

type mapFSWriter struct {
	mfs  *platform.MapFileSystem
	path string
	buf  bytes.Buffer
}

func (w *mapFSWriter) Write(p []byte) (int, error) {
	return w.buf.Write(p)
}

func (w *mapFSWriter) Close() error {
	return w.mfs.WriteFile(w.path, w.buf.Bytes(), 0644)
}

type noopDesignTokensCache struct{}

func (n *noopDesignTokensCache) LoadOrReuse(_ types.WorkspaceContext) (types.DesignTokens, error) {
	return nil, nil
}

func (n *noopDesignTokensCache) Clear() {}
