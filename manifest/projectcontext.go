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
	"strings"

	"bennypowers.dev/cem/cmd/config"
	"github.com/bmatcuk/doublestar"
	"github.com/pterm/pterm"
	"gopkg.in/yaml.v3"
)

var ErrNoManifest = errors.New("no package.json found, could not derive custom-elements.json")
var ErrNoPackageCustomElements = errors.New("package does not specify a custom elements manifest")

// isGlobPattern checks if a string contains any common glob pattern metacharacters.
// This is a heuristic and may produce false positives for file paths that
// legitimately contain one of these characters, but it covers most common cases.
func isGlobPattern(pattern string) bool {
	// The set of characters that are special in glob patterns.
	// We include '*' for wildcards, '?' for single characters,
	// '[' and ']' for character classes, and '{' and '}' for brace expansion.
	globChars := "*?[]"
	return strings.ContainsAny(pattern, globChars)
}

// WorkspaceContext abstracts access to project resources, regardless of source (local or remote).
type WorkspaceContext interface {
	// Performs validation/discovery and caches results as needed.
	Init() error
	// ConfigFile Returns the path to the config file
	ConfigFile() string
	// Config returns the parsed and initialized config
	Config() (*config.CemConfig, error)
	// Returns the project's parsed PackageJSON.
	PackageJSON() (*PackageJSON, error)
	// Manifest returns the project's parsed custom elements manifest.
	Manifest() (*Package, error)
	// ReadFile returns an io.ReadCloser for a file within the project.
	ReadFile(path string) (io.ReadCloser, error)
	// Glob returns a list of file paths matching the given pattern (e.g., *.ts).
	Glob(pattern string) ([]string, error)
	// Writes outputs to paths
	OutputWriter(path string) (io.WriteCloser, error)
	// Root returns the canonical root path or name for the project.
	Root() string
	// Cleanup releases any resources (e.g., tempdirs) held by the context.
	Cleanup() error
}

var _ WorkspaceContext = (*FileSystemWorkspaceContext)(nil)

// FileSystemWorkspaceContext implements ProjectContext for a local filesystem project.
type FileSystemWorkspaceContext struct {
	root            string
	config          *config.CemConfig
	manifestPath    string
	packageJSONPath string
	// Cache parsed results if desired
	manifest    *Package
	packageJSON *PackageJSON
}

func (c *FileSystemWorkspaceContext) initConfig() (*config.CemConfig, error) {
	var config config.CemConfig
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
	// Make output path project-root-relative if needed
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

func (c *FileSystemWorkspaceContext) PackageJSON() (*PackageJSON, error) {
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
	p, err := decodeJSON[PackageJSON](rc)
	if err != nil {
		return nil, err
	}
	c.packageJSON = p
	return p, err
}

func (c *FileSystemWorkspaceContext) Config() (*config.CemConfig, error) {
	return c.config, nil
}

func (c *FileSystemWorkspaceContext) Manifest() (*Package, error) {
	if pkg, err := c.PackageJSON(); err != nil {
		return nil, err
	} else if pkg == nil {
		// TODO: an input flag?
		// TODO: fall back to `custom-elements.json`?
		return nil, ErrNoManifest
	} else if pkg.CustomElements == "" {
		return nil, ErrNoPackageCustomElements
	} else if rc, err := c.ReadFile(filepath.Join(c.root, pkg.CustomElements)); err != nil {
		return nil, err
	} else {
		return decodeJSON[Package](rc)
	}
}

// Init discovers package.json file, caches paths/parsed results.
func (c *FileSystemWorkspaceContext) Init() error {
	var err error
	c.config, err = c.initConfig()
	if err != nil {
		return err
	}
	// Discover package.json
	_, err = c.PackageJSON()
	if !errors.Is(err, ErrNoPackageCustomElements) {
		return err
	}
	_, err = c.Manifest()
	if err != nil {
		return err
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

func (c *FileSystemWorkspaceContext) Root() string {
	return c.root
}

func (c *FileSystemWorkspaceContext) Cleanup() error {
	// Nothing to clean up for local projects
	return nil
}

var _ WorkspaceContext = (*RemoteWorkspaceContext)(nil)

// RemoteWorkspaceContext implements ProjectContext for remote/package-based projects.
type RemoteWorkspaceContext struct {
	tempdir string
	// Add fields for cached files, manifest URL, etc.
}

var ErrRemoteUnsupported = fmt.Errorf("Remote project context is not yet supported: %w", errors.ErrUnsupported)

func NewRemoteProjectContext(tempdir string) *RemoteWorkspaceContext {
	return &RemoteWorkspaceContext{tempdir: tempdir}
}

func (c *RemoteWorkspaceContext) Init() error {
	return ErrRemoteUnsupported
}

func (c *RemoteWorkspaceContext) ConfigFile() string {
	// TODO: Download or extract config file to tempdir, open and return it
	return ""
}

func (c *RemoteWorkspaceContext) Config() (*config.CemConfig, error) {
	// TODO: Download or extract config file to tempdir, open and return it
	return nil, ErrRemoteUnsupported
}

func (c *RemoteWorkspaceContext) Manifest() (*Package, error) {
	// TODO: Download or extract manifest file to tempdir, open and return it
	return nil, ErrRemoteUnsupported
}

func (c *RemoteWorkspaceContext) PackageJSON() (*PackageJSON, error) {
	// TODO: Download or extract package.json file to tempdir, open and return it
	return nil, ErrRemoteUnsupported
}

func (c *RemoteWorkspaceContext) ReadFile(path string) (io.ReadCloser, error) {
	// TODO: Download or extract file to tempdir, open and return it
	return nil, ErrRemoteUnsupported
}

func (c *RemoteWorkspaceContext) Glob(pattern string) ([]string, error) {
	// TODO: List files in the tempdir matching pattern
	return nil, ErrRemoteUnsupported
}

func (c *RemoteWorkspaceContext) OutputWriter(path string) (io.WriteCloser, error) {
	return nil, ErrRemoteUnsupported
}

func (c *RemoteWorkspaceContext) Root() string {
	return c.tempdir
}

func (c *RemoteWorkspaceContext) Cleanup() error {
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
