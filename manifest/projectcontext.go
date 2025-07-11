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

	"github.com/bmatcuk/doublestar"
)

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
	// Returns the path to the config file
	ConfigFile() string
	// Returns the project's parsed PackageJSON.
	PackageJSON() (*PackageJSON, error)
	// Manifest returns the project's parsed custom elements manifest.
	Manifest() (*Package, error)
	// SourceFile returns an io.ReadCloser for a file within the project.
	ReadFile(path string) (io.ReadCloser, error)
	// SourceFile returns an io.ReadCloser for a TypeScript source file within the project.
	SourceFile(path string) (io.ReadCloser, error)
	// Glob returns a list of file paths matching the given pattern (e.g., *.ts).
	Glob(pattern string) ([]string, error)
	// Writes outputs to paths
	OutputWriter(path string) (io.WriteCloser, error)
	// Root returns the canonical root path or name for the project.
	Root() string
	// Cleanup releases any resources (e.g., tempdirs) held by the context.
	Cleanup() error
}

var _ WorkspaceContext = (*LocalFSProjectContext)(nil)

// LocalFSProjectContext implements ProjectContext for a local filesystem project.
type LocalFSProjectContext struct {
	root            string
	manifestPath    string
	packageJSONPath string
	// Cache parsed results if desired
	manifest    *Package
	packageJSON *PackageJSON
}

func NewLocalFSProjectContext(root string) *LocalFSProjectContext {
	return &LocalFSProjectContext{root: root}
}

// ConfigFile Returns the path to the config file, or an empty string if it does not exist.
func (c *LocalFSProjectContext) ConfigFile() string {
	return filepath.Join(c.root, ".config", "cem.yaml")
}

func (c *LocalFSProjectContext) PackageJSON() (*PackageJSON, error) {
	rc, err := c.ReadFile("package.json")
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	p, err := decodeJSON[PackageJSON](rc)
	if err == nil && p != nil {
		c.packageJSON = p
	}
	return p, err
}

func (c *LocalFSProjectContext) Manifest() (*Package, error) {
	if pkg, err := c.PackageJSON(); err != nil {
		return nil, err
	} else if pkg.CustomElements == "" {
		return nil, errors.New("package does not specify a custom elements manifest")
	} else if rc, err := c.ReadFile(pkg.CustomElements); err != nil {
		return nil, err
	} else {
		return decodeJSON[Package](rc)
	}
}

// Init discovers package.json file, caches paths/parsed results.
func (c *LocalFSProjectContext) Init() error {
	// Discover package.json
	packageJSONPath := filepath.Join(c.root, "package.json")
	if _, err := os.Stat(packageJSONPath); err == nil {
		c.packageJSONPath = packageJSONPath
		if pkg, err := c.PackageJSON(); err == nil {
			c.packageJSON = pkg
		}
	}
	return nil
}

func (c *LocalFSProjectContext) ReadFile(path string) (io.ReadCloser, error) {
	abs := ""
	if filepath.IsAbs(path) {
		abs = path
	} else {
		abs = filepath.Join(c.root, path)
	}
	rc, err := os.Open(abs)
	if err != nil {
		return nil, fmt.Errorf("LocalFSProjectContext could not open: %w", err)
	}
	return rc, nil
}

func (c *LocalFSProjectContext) SourceFile(path string) (io.ReadCloser, error) {
	return c.ReadFile(path)
}

func (c *LocalFSProjectContext) Glob(pattern string) ([]string, error) {
	if isGlobPattern(pattern) {
		return doublestar.Glob(filepath.Join(c.root, pattern))
	} else {
		return []string{pattern}, nil
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

var _ WorkspaceContext = (*RemoteProjectContext)(nil)

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

func (c *RemoteProjectContext) ConfigFile() string {
	// TODO: Download or extract config file to tempdir, open and return it
	return ""
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

func (c *RemoteProjectContext) Glob(pattern string) ([]string, error) {
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
