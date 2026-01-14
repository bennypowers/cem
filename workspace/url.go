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
	"fmt"
	"io"
	"net/url"
	"path"
	"strings"

	C "bennypowers.dev/cem/cmd/config"
	M "bennypowers.dev/cem/manifest"
	"bennypowers.dev/cem/types"
)

var _ types.WorkspaceContext = (*URLWorkspaceContext)(nil)

// URLWorkspaceContext implements WorkspaceContext for packages loaded from URLs.
// It fetches package.json and the manifest from the given URL, using HTTP caching.
type URLWorkspaceContext struct {
	baseURL           string
	cacheDir          string
	cache             *HTTPCache
	packageJSON       *M.PackageJSON
	manifest          *M.Package
	manifestPath      string
	designTokensCache types.DesignTokensCache
}

// NewURLWorkspaceContext creates a new context for loading packages from a URL.
// The baseURL should point to the package root (where package.json is located).
// The cacheDir is used for HTTP response caching.
func NewURLWorkspaceContext(baseURL, cacheDir string) *URLWorkspaceContext {
	// Ensure URL ends with /
	if !strings.HasSuffix(baseURL, "/") {
		baseURL += "/"
	}
	return &URLWorkspaceContext{
		baseURL:           baseURL,
		cacheDir:          cacheDir,
		designTokensCache: NewDesignTokensCache(nil),
	}
}

func (c *URLWorkspaceContext) Init() error {
	// Create HTTP cache
	c.cache = NewHTTPCache(c.cacheDir)

	// Fetch package.json
	pkgJSONURL := c.baseURL + "package.json"
	content, err := c.cache.Fetch(pkgJSONURL)
	if err != nil {
		return fmt.Errorf("failed to fetch package.json from %s: %w", pkgJSONURL, err)
	}

	var pkgJSON M.PackageJSON
	if err := json.Unmarshal(content, &pkgJSON); err != nil {
		return fmt.Errorf("failed to parse package.json: %w", err)
	}
	c.packageJSON = &pkgJSON

	// Verify customElements field is set
	if c.packageJSON.CustomElements == "" {
		return ErrNoPackageCustomElements
	}

	// Fetch the manifest
	manifestPath := c.packageJSON.CustomElements
	// Handle relative paths (./custom-elements.json -> custom-elements.json)
	manifestPath = strings.TrimPrefix(manifestPath, "./")
	c.manifestPath = manifestPath

	manifestURL := c.baseURL + manifestPath
	manifestContent, err := c.cache.Fetch(manifestURL)
	if err != nil {
		return fmt.Errorf("failed to fetch manifest from %s: %w", manifestURL, err)
	}

	var manifest M.Package
	if err := json.Unmarshal(manifestContent, &manifest); err != nil {
		return fmt.Errorf("failed to parse manifest: %w", err)
	}
	c.manifest = &manifest

	return nil
}

func (c *URLWorkspaceContext) ConfigFile() string {
	return ""
}

func (c *URLWorkspaceContext) Config() (*C.CemConfig, error) {
	// Remote URLs don't have a config file
	return &C.CemConfig{}, nil
}

func (c *URLWorkspaceContext) PackageJSON() (*M.PackageJSON, error) {
	if c.packageJSON == nil {
		return nil, ErrNoPackageJSON
	}
	return c.packageJSON, nil
}

func (c *URLWorkspaceContext) Manifest() (*M.Package, error) {
	if c.manifest == nil {
		return nil, ErrManifestNotFound
	}
	return c.manifest, nil
}

func (c *URLWorkspaceContext) CustomElementsManifestPath() string {
	return c.manifestPath
}

func (c *URLWorkspaceContext) ReadFile(filePath string) (io.ReadCloser, error) {
	// Fetch the file from the URL
	fileURL := c.resolveURL(filePath)
	content, err := c.cache.Fetch(fileURL)
	if err != nil {
		return nil, err
	}
	return io.NopCloser(strings.NewReader(string(content))), nil
}

func (c *URLWorkspaceContext) Glob(pattern string) ([]string, error) {
	// Glob is not supported for URL contexts
	return nil, nil
}

func (c *URLWorkspaceContext) OutputWriter(filePath string) (io.WriteCloser, error) {
	// Writing is not supported for URL contexts
	return nil, ErrRemoteUnsupported
}

func (c *URLWorkspaceContext) Root() string {
	return c.baseURL
}

func (c *URLWorkspaceContext) Cleanup() error {
	// Nothing to clean up for URL contexts
	return nil
}

func (c *URLWorkspaceContext) ModulePathToFS(modulePath string) string {
	return modulePath
}

func (c *URLWorkspaceContext) FSPathToModule(fsPath string) (string, error) {
	return fsPath, nil
}

func (c *URLWorkspaceContext) ResolveModuleDependency(modulePath, dependencyPath string) (string, error) {
	// Simple path resolution for dependencies
	if strings.HasPrefix(dependencyPath, ".") {
		dir := path.Dir(modulePath)
		return path.Join(dir, dependencyPath), nil
	}
	return dependencyPath, nil
}

func (c *URLWorkspaceContext) DesignTokensCache() types.DesignTokensCache {
	return c.designTokensCache
}

func (c *URLWorkspaceContext) resolveURL(filePath string) string {
	// Handle relative paths
	filePath = strings.TrimPrefix(filePath, "./")

	// Parse base URL and join with file path
	base, err := url.Parse(c.baseURL)
	if err != nil {
		return c.baseURL + filePath
	}

	ref, err := url.Parse(filePath)
	if err != nil {
		return c.baseURL + filePath
	}

	return base.ResolveReference(ref).String()
}
