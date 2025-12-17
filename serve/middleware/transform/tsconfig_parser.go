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

package transform

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"

	"bennypowers.dev/cem/internal/platform"
)

// TsConfig represents the structure of a tsconfig.json file.
// We only parse the fields we need for path resolution.
type TsConfig struct {
	CompilerOptions *TsConfigCompilerOptions `json:"compilerOptions"`
	Extends         string                   `json:"extends"`
}

// TsConfigCompilerOptions represents compiler options we care about
type TsConfigCompilerOptions struct {
	RootDir string `json:"rootDir"`
	OutDir  string `json:"outDir"`
}

// ParseTsConfig reads a tsconfig.json file and extracts path mappings.
// Returns a map of URL prefix -> source directory (e.g., {"/dist/": "/src/"}).
//
// Handles tsconfig inheritance via the "extends" field (max depth: 5).
// Missing rootDir or outDir default to ".".
// Relative paths are normalized to the tsconfig directory.
func ParseTsConfig(path string, fs platform.FileSystem) (map[string]string, error) {
	return parseTsConfigRecursive(path, fs, 0, make(map[string]bool))
}

// parseTsConfigRecursive handles tsconfig parsing with inheritance support
func parseTsConfigRecursive(
	path string,
	fs platform.FileSystem,
	depth int,
	visited map[string]bool,
) (map[string]string, error) {
	// Prevent infinite recursion
	const maxDepth = 5
	if depth > maxDepth {
		return nil, fmt.Errorf("tsconfig extends depth exceeded (max: %d)", maxDepth)
	}

	// Prevent circular extends
	absPath, err := filepath.Abs(path)
	if err != nil {
		absPath = path
	}
	if visited[absPath] {
		return nil, fmt.Errorf("circular tsconfig extends detected: %s", absPath)
	}
	visited[absPath] = true

	// Read tsconfig file
	data, err := fs.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading tsconfig: %w", err)
	}

	// Parse JSON
	var config TsConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("parsing tsconfig JSON: %w", err)
	}

	// Merge compiler options from base config if extends is present
	var rootDir, outDir string

	if config.Extends != "" {
		// Resolve extends path relative to current tsconfig directory
		tsconfigDir := filepath.Dir(path)
		extendsPath := filepath.Join(tsconfigDir, config.Extends)

		// If extends doesn't have .json extension, add it
		if !strings.HasSuffix(extendsPath, ".json") {
			extendsPath += ".json"
		}

		// Recursively parse base config (handles chained inheritance and depth tracking)
		baseMappings, err := parseTsConfigRecursive(extendsPath, fs, depth+1, visited)
		if err != nil {
			return nil, err
		}

		// Parse base config to get compiler options
		baseData, err := fs.ReadFile(extendsPath)
		if err != nil {
			return nil, fmt.Errorf("reading extended tsconfig %s: %w", extendsPath, err)
		}

		var baseConfig TsConfig
		if err := json.Unmarshal(baseData, &baseConfig); err != nil {
			return nil, fmt.Errorf("parsing extended tsconfig JSON %s: %w", extendsPath, err)
		}

		// Start with base values
		if baseConfig.CompilerOptions != nil {
			rootDir = baseConfig.CompilerOptions.RootDir
			outDir = baseConfig.CompilerOptions.OutDir
		}

		// If base had mappings from further extends, we would merge them here
		// For now, we just handle direct rootDir/outDir inheritance
		_ = baseMappings
	}

	// Override with current config values
	if config.CompilerOptions != nil {
		if config.CompilerOptions.RootDir != "" {
			rootDir = config.CompilerOptions.RootDir
		}
		if config.CompilerOptions.OutDir != "" {
			outDir = config.CompilerOptions.OutDir
		}
	}

	// Default to "." if still not specified
	if rootDir == "" {
		rootDir = "."
	}
	if outDir == "" {
		outDir = "."
	}

	// If rootDir and outDir are the same, this is in-place compilation
	// No path mapping needed
	if rootDir == outDir {
		return make(map[string]string), nil
	}

	// Normalize paths relative to tsconfig directory
	tsconfigDir := filepath.Dir(path)

	// Resolve rootDir and outDir to absolute paths, then make them relative to cwd
	// This ensures they work correctly when watchDir might be different
	absRootDir := filepath.Join(tsconfigDir, rootDir)
	absOutDir := filepath.Join(tsconfigDir, outDir)

	// Make them relative to the tsconfig directory for the mapping
	// We want the mapping to be relative to watchDir, which is typically the tsconfig directory
	relRootDir, err := filepath.Rel(tsconfigDir, absRootDir)
	if err != nil {
		relRootDir = rootDir
	}
	relOutDir, err := filepath.Rel(tsconfigDir, absOutDir)
	if err != nil {
		relOutDir = outDir
	}

	// Normalize to forward slashes for URL paths
	relRootDir = filepath.ToSlash(relRootDir)
	relOutDir = filepath.ToSlash(relOutDir)

	// Ensure leading/trailing slashes for URL prefixes
	if !strings.HasPrefix(relOutDir, "/") {
		relOutDir = "/" + relOutDir
	}
	if !strings.HasSuffix(relOutDir, "/") {
		relOutDir = relOutDir + "/"
	}

	if !strings.HasPrefix(relRootDir, "/") {
		relRootDir = "/" + relRootDir
	}
	if !strings.HasSuffix(relRootDir, "/") {
		relRootDir = relRootDir + "/"
	}

	// Create mapping from the merged compiler options
	result := make(map[string]string)
	result[relOutDir] = relRootDir

	return result, nil
}
