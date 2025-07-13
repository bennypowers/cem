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
	"strings"

	C "bennypowers.dev/cem/cmd/config"
	M "bennypowers.dev/cem/manifest"
)

type contextKey string

const WorkspaceContextKey = contextKey("workspaceContext")

var ErrNoManifest = errors.New("no package.json found, could not derive custom-elements.json")
var ErrRemoteUnsupported = fmt.Errorf("Remote workspace context is not yet supported: %w", errors.ErrUnsupported)
var ErrNoPackageCustomElements = errors.New("package does not specify a custom elements manifest")
var ErrManifestNotFound = errors.New("manifest not found")
var ErrPackageNotFound = errors.New("package not found")

// isGlobPattern checks if a string contains any common glob pattern metacharacters.
// This is a heuristic and may produce false positives for file paths that
// legitimately contain one of these characters, but it covers most common cases.
func isGlobPattern(pattern string) bool {
	// The set of characters that are special in glob patterns.
	// We include '*' for wildcards, '?' for single characters,
	// '[' and ']' for character classes, and '{' and '}' for brace expansion.
	globChars := "*?[]{}"
	return strings.ContainsAny(pattern, globChars)
}

// parseNpmSpecifier parses a spec like "@scope/pkg@1.2.3" or "pkg@1.2.3"
func parseNpmSpecifier(spec string) (name, version string, err error) {
	spec = strings.TrimPrefix(spec, "npm:")
	atIndex := strings.LastIndex(spec, "@")

	if atIndex <= 0 { // <= 0 to handle scoped packages like @foo/bar
		name = spec
		version = "latest"
	} else {
		name = spec[:atIndex]
		version = spec[atIndex+1:]
	}

	if name == "" {
		return "", "", errors.New("invalid npm specifier: missing package name")
	}

	if version == "" {
		version = "latest"
	}

	return name, version, nil
}

// decodeJSON parses a JSON stream into a struct of type T.
func decodeJSON[T any](rc io.ReadCloser) (*T, error) {
	defer rc.Close()
	var out T
	if err := json.NewDecoder(rc).Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

// IsPackageSpecifier checks if a string is an npm package specifier.
func IsPackageSpecifier(spec string) bool {
	return strings.HasPrefix(spec, "npm:") || strings.HasPrefix(spec, "jsr:")
}

// WorkspaceContext abstracts access to package resources, regardless of source (local or remote).
type WorkspaceContext interface {
	// Performs validation/discovery and caches results as needed.
	Init() error
	// ConfigFile Returns the path to the config file
	ConfigFile() string
	// Config returns the parsed and initialized config
	Config() (*C.CemConfig, error)
	// Returns the package's parsed PackageJSON.
	PackageJSON() (*M.PackageJSON, error)
	// Manifest returns the package's parsed custom elements manifest.
	Manifest() (*M.Package, error)
	// ReadFile returns an io.ReadCloser for a file within the package.
	ReadFile(path string) (io.ReadCloser, error)
	// Glob returns a list of file paths matching the given pattern (e.g., *.ts).
	Glob(pattern string) ([]string, error)
	// Writes outputs to paths
	OutputWriter(path string) (io.WriteCloser, error)
	// Root returns the canonical root path or name for the package.
	Root() string
	// Cleanup releases any resources (e.g., tempdirs) held by the context.
	Cleanup() error
}
