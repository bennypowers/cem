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

	"bennypowers.dev/cem/types"
)

type contextKey string

const WorkspaceContextKey = contextKey("workspaceContext")

var ErrNoManifest = errors.New("no package.json found, could not derive custom-elements.json")
var ErrNoPackageJSON = errors.New("package.json not loaded")
var ErrRemoteUnsupported = fmt.Errorf("remote workspace context is not yet supported: %w", errors.ErrUnsupported)
var ErrNoPackageCustomElements = errors.New("package does not specify a custom elements manifest")
var ErrManifestNotFound = errors.New("manifest not found")
var ErrPackageNotFound = errors.New("package not found")

// designTokensCacheImpl implements DesignTokensCache
type designTokensCacheImpl struct {
	loader types.DesignTokensLoader
	spec   string
	tokens types.DesignTokens
	err    error
}

// NewDesignTokensCache creates a new design tokens cache with the given loader
func NewDesignTokensCache(loader types.DesignTokensLoader) types.DesignTokensCache {
	return &designTokensCacheImpl{
		loader: loader,
	}
}

// LoadOrReuse loads design tokens from the cache if available, or loads and caches them
func (cache *designTokensCacheImpl) LoadOrReuse(ctx types.WorkspaceContext) (types.DesignTokens, error) {
	// If no loader is provided, return nil (no design tokens)
	if cache.loader == nil {
		return nil, nil
	}

	cfg, err := ctx.Config()
	if err != nil {
		return nil, err
	}

	spec := cfg.Generate.DesignTokens.Spec

	// If cache hit (same spec), return cached result
	if cache.spec == spec {
		return cache.tokens, cache.err
	}

	// Cache miss or different spec, load fresh
	cache.spec = spec
	cache.tokens, cache.err = cache.loader.Load(ctx)
	return cache.tokens, cache.err
}

// Clear resets the cache
func (cache *designTokensCacheImpl) Clear() {
	cache.spec = ""
	cache.tokens = nil
	cache.err = nil
}

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
	defer func() { _ = rc.Close() }()
	var out T
	if err := json.NewDecoder(rc).Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

// IsPackageSpecifier checks if a string is an npm or jsr package specifier.
func IsPackageSpecifier(spec string) bool {
	return strings.HasPrefix(spec, "npm:") || strings.HasPrefix(spec, "jsr:")
}

// IsURLSpecifier checks if a string is an HTTP(S) URL.
func IsURLSpecifier(spec string) bool {
	return strings.HasPrefix(spec, "https://") || strings.HasPrefix(spec, "http://")
}
