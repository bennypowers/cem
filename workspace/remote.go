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
	"archive/tar"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	C "bennypowers.dev/cem/cmd/config"
	M "bennypowers.dev/cem/manifest"
	"bennypowers.dev/cem/types"
	"github.com/adrg/xdg"
	"github.com/bmatcuk/doublestar"
	"github.com/pterm/pterm"
	"gopkg.in/yaml.v3"
)

var _ types.WorkspaceContext = (*RemoteWorkspaceContext)(nil)

// RemoteWorkspaceContext implements WorkspaceContext for remote packages.
type RemoteWorkspaceContext struct {
	spec               string
	name               string
	version            string
	cacheDir           string
	packageJSONPath    string
	packageJSON        *M.PackageJSON
	customElementsPath string
	spinner            *pterm.SpinnerPrinter
	designTokensCache  types.DesignTokensCache
}

func NewRemoteWorkspaceContext(spec string) *RemoteWorkspaceContext {
	return &RemoteWorkspaceContext{
		spec:              spec,
		designTokensCache: NewDesignTokensCache(nil), // Remote contexts don't need design tokens loading
	}
}

func (c *RemoteWorkspaceContext) Init() error {
	name, version, err := parseNpmSpecifier(c.spec)
	if err != nil {
		return err
	}
	c.name = name
	c.version = version
	c.cacheDir = filepath.Join(xdg.CacheHome, "cem", "packages", pkgCacheDirName(name, version))
	c.packageJSONPath = filepath.Join(c.cacheDir, "package.json")

	// Check cache first
	if fileExists(c.packageJSONPath) {
		if err := c.readInPackageJSON(); err == nil {
			c.customElementsPath = filepath.Join(c.cacheDir, c.packageJSON.CustomElements)
			if fileExists(c.customElementsPath) {
				pterm.Debug.Printf("Using cached remote package %s\n", c.spec)
				return nil // Fully cached
			}
		}
	}

	if err := os.MkdirAll(c.cacheDir, 0755); err != nil {
		return fmt.Errorf("failed to create cache directory: %w", err)
	}

	c.spinner, _ = pterm.SpinnerPrinter.Start(*pterm.DefaultSpinner.WithRemoveWhenDone())
	defer func() {
		if err := c.spinner.Stop(); err != nil {
			// Spinner stop errors are not critical, just log them
			fmt.Printf("Warning: error stopping spinner: %v\n", err)
		}
	}()

	// Try fetching from CDNs first (jsdelivr preferred, then esm.sh, then unpkg as fallback)
	cdnFetchers := []func(string, string) error{
		c.fetchFromJsdelivr,
		c.fetchFromEsmsh,
		c.fetchFromUnpkg,
	}

	var lastCdnError error
	for _, fetcher := range cdnFetchers {
		err := fetcher(name, version)
		if err == nil {
			return nil // Success
		}

		if errors.Is(err, ErrPackageNotFound) {
			pterm.Debug.Println("Package not found on CDN, stopping.")
			return err // The package doesn't exist, no point in trying others.
		}

		lastCdnError = err // Store the error to return later if all fallbacks fail

		// If we successfully got package.json but failed to get the manifest,
		// then the manifest is confirmed to be missing. No need to try other sources.
		if c.packageJSON != nil && errors.Is(err, ErrManifestNotFound) {
			pterm.Debug.Println("Found package.json but manifest is missing. Stopping.")
			return ErrManifestNotFound
		}

		pterm.Debug.Printf("CDN fetcher failed, trying next source: %v\n", err)
	}

	// If we're here, it means we couldn't even get package.json from any CDN.
	// Fallback to the npm tarball.
	pterm.Debug.Println("All CDN fetchers failed. Falling back to npm tarball.")
	if err := c.fetchFromNpm(name, version); err != nil {
		// Join the npm error with the last CDN error for a more complete picture.
		return errors.Join(lastCdnError, err)
	}

	return nil // Success from npm tarball
}

func (c *RemoteWorkspaceContext) readInPackageJSON() error {
	pkgjson, err := c.ReadFile(c.packageJSONPath)
	if err != nil {
		return err
	}
	c.packageJSON, err = decodeJSON[M.PackageJSON](pkgjson)
	if err != nil {
		return err
	}
	if c.packageJSON.CustomElements == "" {
		c.packageJSON.CustomElements = "custom-elements.json"
	}
	return nil
}

func (c *RemoteWorkspaceContext) fetchFromCDN(name, version, cdnName, baseUrlPattern string) error {
	c.spinner.UpdateText(fmt.Sprintf("Fetching package from %s", cdnName))
	base := fmt.Sprintf(baseUrlPattern, name, version)
	if err := c.fetch(base+"package.json", c.packageJSONPath); err != nil {
		c.spinner.Warning(fmt.Sprintf("Failed to load package.json from %s", cdnName))
		return err
	}
	if err := c.readInPackageJSON(); err != nil {
		return err
	}
	c.customElementsPath = filepath.Join(c.cacheDir, c.packageJSON.CustomElements)
	if err := c.fetch(base+c.packageJSON.CustomElements, c.customElementsPath); err != nil {
		c.spinner.Warning(fmt.Sprintf("Failed to load %s from %s", c.packageJSON.CustomElements, cdnName))
		return err
	}
	return nil
}

func (c *RemoteWorkspaceContext) fetchFromJsdelivr(name, version string) error {
	return c.fetchFromCDN(name, version, "jsdelivr", "https://cdn.jsdelivr.net/npm/%s@%s/")
}

func (c *RemoteWorkspaceContext) fetchFromEsmsh(name, version string) error {
	return c.fetchFromCDN(name, version, "esm.sh", "https://esm.sh/%s@%s/")
}

func (c *RemoteWorkspaceContext) fetchFromUnpkg(name, version string) error {
	return c.fetchFromCDN(name, version, "unpkg", "https://unpkg.com/%s@%s/")
}

func (c *RemoteWorkspaceContext) fetchFromNpm(name, version string) error {
	c.spinner.UpdateText("Falling back to npm tarball")
	metaUrl := fmt.Sprintf("https://registry.npmjs.org/%s", strings.ReplaceAll(name, "/", "%2F"))
	resp, err := http.Get(metaUrl)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()
	type npmMeta struct {
		Versions map[string]struct {
			Dist struct{ Tarball string }
		}
	}
	var meta npmMeta
	if err := json.NewDecoder(resp.Body).Decode(&meta); err != nil {
		return err
	}
	tarballUrl := meta.Versions[version].Dist.Tarball
	resp, err = http.Get(tarballUrl)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()
	err = extractFilesFromTarGz(resp.Body, c.cacheDir, []string{"package.json", "custom-elements.json"})
	if err != nil {
		return err
	}
	c.cacheDir, err = setupTempdirFromCache(c.cacheDir)
	if err != nil {
		return err
	}
	return nil
}

func (c *RemoteWorkspaceContext) CustomElementsManifestPath() string {
	return c.customElementsPath
}

func (c *RemoteWorkspaceContext) ConfigFile() string {
	configPath := filepath.Join(c.cacheDir, ".config", "cem.yaml")
	if _, err := os.Stat(configPath); err == nil {
		return configPath
	}
	return ""
}

func (c *RemoteWorkspaceContext) Config() (*C.CemConfig, error) {
	configFile := c.ConfigFile()
	if configFile == "" {
		return nil, errors.New("no config file found in remote workspace")
	}
	f, err := os.Open(configFile)
	if err != nil {
		return nil, err
	}
	defer func() { _ = f.Close() }()
	var cfg C.CemConfig
	if err := yaml.NewDecoder(f).Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func (c *RemoteWorkspaceContext) Manifest() (*M.Package, error) {
	pkg, err := c.PackageJSON()
	if err != nil {
		return nil, err
	}
	if pkg != nil && pkg.CustomElements != "" {
		manifestPath := filepath.Join(c.cacheDir, pkg.CustomElements)
		rc, err := os.Open(manifestPath)
		if err == nil {
			defer func() { _ = rc.Close() }()
			return decodeJSON[M.Package](rc)
		}
	}
	return nil, errors.New("no custom-elements.json found in remote workspace")
}

func (c *RemoteWorkspaceContext) PackageJSON() (*M.PackageJSON, error) {
	rc, err := os.Open(c.packageJSONPath)
	if err != nil {
		return nil, fmt.Errorf("could not open package.json in remote workspace: %w", err)
	}
	defer func() { _ = rc.Close() }()
	return decodeJSON[M.PackageJSON](rc)
}

func (c *RemoteWorkspaceContext) ReadFile(path string) (io.ReadCloser, error) {
	absPath := path
	if !filepath.IsAbs(path) {
		absPath = filepath.Join(c.cacheDir, path)
	}
	return os.Open(absPath)
}

func (c *RemoteWorkspaceContext) Glob(pattern string) ([]string, error) {
	rooted := filepath.Join(c.cacheDir, pattern)
	return doublestar.Glob(rooted)
}

func (c *RemoteWorkspaceContext) OutputWriter(path string) (io.WriteCloser, error) {
	absPath := filepath.Join(c.cacheDir, path)
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}
	return os.Create(absPath)
}

func (c *RemoteWorkspaceContext) Cleanup() error {
	return nil
}

func (c *RemoteWorkspaceContext) Root() string {
	return c.cacheDir
}

// ModulePathToFS converts a module path to filesystem path for watching
func (c *RemoteWorkspaceContext) ModulePathToFS(modulePath string) string {
	if filepath.IsAbs(modulePath) {
		return modulePath
	}
	return filepath.Join(c.cacheDir, modulePath)
}

// FSPathToModule converts a filesystem path to module path for manifest lookup
func (c *RemoteWorkspaceContext) FSPathToModule(fsPath string) (string, error) {
	rel, err := filepath.Rel(c.cacheDir, fsPath)
	if err != nil {
		return "", fmt.Errorf("failed to compute relative path from %s to %s: %w", c.cacheDir, fsPath, err)
	}
	return rel, nil
}

// ResolveModuleDependency resolves a dependency path relative to a module
func (c *RemoteWorkspaceContext) ResolveModuleDependency(modulePath, dependencyPath string) (string, error) {
	if filepath.IsAbs(dependencyPath) {
		// Convert to module-relative path
		rel, err := filepath.Rel(c.cacheDir, dependencyPath)
		if err != nil {
			return "", fmt.Errorf("failed to resolve absolute dependency path %s relative to %s: %w", dependencyPath, c.cacheDir, err)
		}
		return rel, nil
	}

	// Resolve relative to module's directory
	moduleDir := filepath.Dir(modulePath)
	resolved := filepath.Join(moduleDir, dependencyPath)
	return filepath.Clean(resolved), nil
}

// DesignTokensCache returns the design tokens cache for this workspace
func (c *RemoteWorkspaceContext) DesignTokensCache() types.DesignTokensCache {
	return c.designTokensCache
}

func setupTempdirFromCache(cacheDir string) (string, error) {
	tempdir, err := os.MkdirTemp("", "cem-remote-*")
	if err != nil {
		return "", err
	}
	files, err := os.ReadDir(cacheDir)
	if err != nil {
		return "", err
	}
	for _, f := range files {
		src := filepath.Join(cacheDir, f.Name())
		dst := filepath.Join(tempdir, f.Name())
		_, _ = copyFile(src, dst)
	}
	return tempdir, nil
}

func (c *RemoteWorkspaceContext) fetch(url, dest string) error {
	c.spinner.UpdateText("Fetching " + url)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode == http.StatusNotFound {
		if strings.HasSuffix(url, "package.json") {
			return fmt.Errorf("%w: %s", ErrPackageNotFound, c.name)
		}
		return fmt.Errorf("%w: %s", ErrManifestNotFound, url)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch %s: status %s", url, resp.Status)
	}

	if err := os.MkdirAll(filepath.Dir(dest), 0755); err != nil {
		return err
	}

	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer func() { _ = out.Close() }()
	_, err = io.Copy(out, resp.Body)
	return err
}

func extractFilesFromTarGz(r io.Reader, dest string, wanted []string) error {
	gzr, err := gzip.NewReader(r)
	if err != nil {
		return err
	}
	defer func() { _ = gzr.Close() }()
	tr := tar.NewReader(gzr)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		for _, w := range wanted {
			if path.Base(hdr.Name) == w {
				out, err := os.Create(filepath.Join(dest, w))
				if err != nil {
					return err
				}
				if _, err := io.Copy(out, tr); err != nil {
					_ = out.Close()
					return err
				}
				_ = out.Close()
			}
		}
	}
	return nil
}

func pkgCacheDirName(name, version string) string {
	return fmt.Sprintf("%s@%s", strings.ReplaceAll(name, "/", "+"), version)
}
