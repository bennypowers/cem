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

package manifest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	C "bennypowers.dev/cem/cmd/config"
	"github.com/bmatcuk/doublestar"
	"github.com/pterm/pterm"
	"github.com/spf13/viper"
)

// ProjectContext abstracts access to project resources, regardless of source (local or remote).
type ProjectContext interface {
	// Performs validation/discovery and caches results as needed.
	// Initializes the project config, based on the config file, if present
	Init() error
	// Returns the project config
	Config() (*C.CemConfig, error)
	// Returns the project's parsed PackageJSON.
	PackageJSON() (*PackageJSON, error)
	// Manifest returns the project's parsed custom elements manifest.
	Manifest() (*Package, error)
	// SourceFile returns an io.ReadCloser for a file within the project.
	ReadFile(path string) (io.ReadCloser, error)
	// SourceFile returns an io.ReadCloser for a TypeScript source file within the project.
	SourceFile(path string) (io.ReadCloser, error)
	// ListFiles returns a list of file paths matching the given pattern (e.g., *.ts).
	ListFiles(pattern string) ([]string, error)
	// Writes outputs to paths
	OutputWriter(path string) (io.WriteCloser, error)
	// Root returns the canonical root path or name for the project.
	Root() string
	// Cleanup releases any resources (e.g., tempdirs) held by the context.
	Cleanup() error
}

var _ ProjectContext = (*LocalFSProjectContext)(nil)

// LocalFSProjectContext implements ProjectContext for a local filesystem project.
type LocalFSProjectContext struct {
	root            string
	manifestPath    string
	packageJSONPath string
	config          *C.CemConfig
	// Cache parsed results if desired
	manifest    *Package
	packageJSON *PackageJSON
}

func NewLocalFSProjectContext(root string) *LocalFSProjectContext {
	return &LocalFSProjectContext{root: root}
}

func (c *LocalFSProjectContext) Config() (*C.CemConfig, error) {
	if c.config != nil {
		return c.config, nil
	} else {
		return nil, errors.New("config uninitialized")
	}
}

func (c *LocalFSProjectContext) initConfig() (*C.CemConfig, error) {
	cfgFile := filepath.Join(c.root, ".config", "cem.yaml")
	if _, err := os.Stat(cfgFile); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			cfgFile = ""
		} else {
			return nil, err
		}
	}

	cfg := &C.CemConfig{
		ProjectDir: c.root,
	}

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
		if err := viper.ReadInConfig(); err != nil {
			return nil, err
		} else {
			cfg.ConfigFile = cfgFile
			cfg.Generate.NoDefaultExcludes = viper.GetBool("generate.noDefaultExcludes")
			cfg.Generate.Output = viper.GetString("generate.output")
			cfg.Generate.Exclude = viper.GetStringSlice("generate.exclude")
			cfg.Generate.DesignTokens.Spec = viper.GetString("generate.designTokens.spec")
			cfg.Generate.DesignTokens.Prefix = viper.GetString("generate.designTokens.prefix")
			cfg.Generate.DemoDiscovery.FileGlob = viper.GetString("generate.demoDiscovery.fileGlob")
			cfg.Generate.DemoDiscovery.URLPattern = viper.GetString("generate.demoDiscovery.urlPattern")
			cfg.Generate.DemoDiscovery.URLTemplate = viper.GetString("generate.demoDiscovery.urlTemplate")
		}
	} else {
		// If no config file was found, return a default config
		return cfg, nil
	}

	rc, err := c.ReadFile(cfgFile)
	if err != nil {
		// If config file not found, return default config
		return &C.CemConfig{
			ProjectDir: c.Root(),
			ConfigFile: cfgFile,
		}, nil
	}
	defer rc.Close()

	// Make output path project-root-relative if needed
	if cfg.Generate.Output != "" && !filepath.IsAbs(cfg.Generate.Output) {
		cfg.Generate.Output = filepath.Join(cfg.ProjectDir, cfg.Generate.Output)
	}

	if cfg.Generate.Output == "" {
		cfg.Generate.Output = filepath.Join(c.root, c.packageJSON.CustomElements)
	}

	// Set debug verbosity
	if cfg.Verbose {
		pterm.EnableDebugMessages()
	} else {
		pterm.DisableDebugMessages()
	}

	if cfg.ConfigFile != "" {
		pterm.Debug.Println("Using config file: ", cfgFile)
	}

	return cfg, nil
}

func (c *LocalFSProjectContext) initPackageJSON() (*PackageJSON, error) {
	// Discover package.json
	packageJSONPath := filepath.Join(c.root, "package.json")
	if _, err := os.Stat(packageJSONPath); err == nil {
		c.packageJSONPath = packageJSONPath
	} else {
		return nil, errors.New("package.json not found at project root")
	}
	rc, err := c.ReadFile("package.json")
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	return decodeJSON[PackageJSON](rc)
}

func (c *LocalFSProjectContext) PackageJSON() (*PackageJSON, error) {
	if c.packageJSON != nil {
		return c.packageJSON, nil
	} else {
		return nil, errors.New("package.json uninitialized")
	}
}

func (c *LocalFSProjectContext) Manifest() (*Package, error) {
	if c.manifest != nil {
		return c.manifest, nil
	}
	if pkg, err := c.PackageJSON(); err != nil {
		return nil, err
	} else if pkg.CustomElements == "" {
		return nil, errors.New("package does not specify a custom elements manifest")
	} else if rc, err := c.ReadFile(pkg.CustomElements); err != nil {
		return nil, err
	} else {
		manifest, err := decodeJSON[Package](rc)
		if err != nil {
			return nil, err
		}
		c.manifest = manifest
		return c.manifest, nil
	}
}

// Init discovers package.json file, caches paths/parsed results.
func (c *LocalFSProjectContext) Init() error {
	pkgJson, err := c.initPackageJSON()
	if err != nil {
		return err
	}
	c.packageJSON = pkgJson
	cfg, err := c.initConfig()
	if err != nil {
		return err
	}
	c.config = cfg
	return nil
}

func (c *LocalFSProjectContext) ReadFile(path string) (io.ReadCloser, error) {
	return os.Open(filepath.Join(c.root, path))
}

func (c *LocalFSProjectContext) SourceFile(path string) (io.ReadCloser, error) {
	return c.ReadFile(path)
}

func (c *LocalFSProjectContext) ListFiles(pattern string) ([]string, error) {
	fmt.Fprintf(os.Stderr, "LocalFSProjectContext.ListFiles(%s)", pattern)
	if filepath.IsAbs(pattern) {
		return []string{pattern}, nil
	} else {
		return doublestar.Glob(filepath.Join(c.root, pattern))
	}
}

func (c *LocalFSProjectContext) OutputWriter(path string) (io.WriteCloser, error) {
	absPath := path
	if !filepath.IsAbs(path) {
		absPath = filepath.Join(c.root, path)
	}
	dir := filepath.Dir(absPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}
	return os.Create(absPath)
}

func (c *LocalFSProjectContext) Root() string {
	return c.root
}

func (c *LocalFSProjectContext) Cleanup() error {
	// Nothing to clean up for local projects
	return nil
}

var _ ProjectContext = (*RemoteProjectContext)(nil)

// RemoteProjectContext implements ProjectContext for remote/package-based projects.
type RemoteProjectContext struct {
	tempdir string
	// Add fields for cached files, manifest URL, etc.
}

var ErrRemoteUnsupported = fmt.Errorf("Remote project context is not yet supported: %w", errors.ErrUnsupported)

func NewRemoteProjectContext(tempdir string) *RemoteProjectContext {
	return &RemoteProjectContext{tempdir: tempdir}
}

func (c *RemoteProjectContext) Init() error {
	return ErrRemoteUnsupported
}

func (c *RemoteProjectContext) Config() (*C.CemConfig, error) {
	// TODO: Download or extract config file to tempdir, open and return it
	return nil, ErrRemoteUnsupported
}

func (c *RemoteProjectContext) ConfigFile() (string, error) {
	// TODO: Download or extract config file to tempdir, open and return it
	return "", ErrRemoteUnsupported
}

func (c *RemoteProjectContext) Manifest() (*Package, error) {
	// TODO: Download or extract manifest file to tempdir, open and return it
	return nil, ErrRemoteUnsupported
}

func (c *RemoteProjectContext) PackageJSON() (*PackageJSON, error) {
	// TODO: Download or extract package.json file to tempdir, open and return it
	return nil, ErrRemoteUnsupported
}

func (c *RemoteProjectContext) ReadFile(path string) (io.ReadCloser, error) {
	// TODO: Download or extract file to tempdir, open and return it
	return nil, ErrRemoteUnsupported
}

func (c *RemoteProjectContext) SourceFile(path string) (io.ReadCloser, error) {
	return c.ReadFile(path)
}

func (c *RemoteProjectContext) ListFiles(pattern string) ([]string, error) {
	// TODO: List files in the tempdir matching pattern
	return nil, ErrRemoteUnsupported
}

func (c *RemoteProjectContext) OutputWriter(path string) (io.WriteCloser, error) {
	return nil, ErrRemoteUnsupported
}

func (c *RemoteProjectContext) Root() string {
	return c.tempdir
}

func (c *RemoteProjectContext) Cleanup() error {
	// TODO: Remove tempdir and any downloaded files
	return nil
}

func decodeJSON[T any](rc io.ReadCloser) (*T, error) {
	defer rc.Close()
	var out T
	if err := json.NewDecoder(rc).Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}
