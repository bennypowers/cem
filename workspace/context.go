/*
Copyright © 2025 Benny Powers

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
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"bennypowers.dev/cem/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Files that indicate project root, in order of preference
var projectFiles = []string{
	".config/cem.yaml",
	"package.json",
	".git",
	"tsconfig.json",
}

func GetWorkspaceContext(cmd *cobra.Command) (types.WorkspaceContext, error) {
	val := cmd.Context().Value(WorkspaceContextKey)
	if val != nil {
		return val.(types.WorkspaceContext), nil
	}

	spec := viper.GetString("package")
	if p, _ := cmd.Flags().GetString("package"); p != "" {
		spec = p
	}

	// If no package specified, determine project root
	if spec == "" {
		configPath := viper.GetString("configFile")
		if p, _ := cmd.Flags().GetString("configFile"); p != "" {
			configPath = p
		}

		if configPath != "" {
			// Use explicit config file path
			spec = filepath.Dir(configPath)
		} else {
			// Auto-detect project root by looking for common project files
			if detectedRoot, found := findProjectRootFromCommand(cmd); found {
				spec = detectedRoot
			} else {
				spec = "." // fallback to current directory
			}
		}
	}

	if ctx, err := getAppropriateContextForSpec(spec, cmd.Name()); err != nil {
		return nil, err
	} else if err := ctx.Init(); err != nil {
		return nil, err
	} else {
		// Store in cmd.Context for future calls
		newCtx := context.WithValue(cmd.Context(), WorkspaceContextKey, ctx)
		cmd.SetContext(newCtx)
		return ctx, nil
	}
}

func getAppropriateContextForSpec(spec, cmdName string) (ctx types.WorkspaceContext, err error) {
	if !IsPackageSpecifier(spec) {
		// if it's a file path, use it
		return NewFileSystemWorkspaceContext(spec), nil
	} else {
		// if it's a specifier, try to get it from node_modules
		name, _, err := parseNpmSpecifier(spec)
		if err != nil {
			return nil, err
		}
		localPath := filepath.Join("node_modules", name)
		if stat, err := os.Stat(localPath); err == nil && stat.IsDir() {
			return NewFileSystemWorkspaceContext(localPath), nil
		}
		// finally, we fetch from the network
		if cmdName == "generate" {
			return nil, errors.New("generate command cannot be used with a remote package specifier")
		}
		ctx := NewRemoteWorkspaceContext(spec)
		return ctx, nil
	}
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func copyFile(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer func() { _ = source.Close() }()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer func() { _ = destination.Close() }()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

// findProjectRoot searches for project root by looking for common project files
// Returns the detected root path and whether it was found
func findProjectRoot() (string, bool) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", false
	}
	return findProjectRootFromDir(cwd)
}

// findProjectRootFromCommand attempts to find project root by looking at command arguments
// for file paths and searching upward from those locations
func findProjectRootFromCommand(cmd *cobra.Command) (string, bool) {
	// First try the standard findProjectRoot from current directory
	if root, found := findProjectRoot(); found {
		return root, true
	}

	// If that fails, look at command arguments for file paths
	args := cmd.Flags().Args()
	for _, arg := range args {
		var dir string
		if filepath.IsAbs(arg) {
			// For absolute paths, start looking from the directory containing the file
			dir = filepath.Dir(arg)
		} else {
			// For relative paths, resolve them relative to current working directory
			if absPath, err := filepath.Abs(arg); err == nil {
				dir = filepath.Dir(absPath)
			} else {
				continue // Skip if we can't resolve the path
			}
		}
		if root, found := findProjectRootFromDir(dir); found {
			return root, true
		}
	}

	// Also check design-tokens flag for additional hint
	if designTokensPath, _ := cmd.Flags().GetString("design-tokens"); designTokensPath != "" {
		var dir string
		if filepath.IsAbs(designTokensPath) {
			dir = filepath.Dir(designTokensPath)
		} else {
			// For relative paths, resolve them relative to current working directory
			if absPath, err := filepath.Abs(designTokensPath); err == nil {
				dir = filepath.Dir(absPath)
			}
		}
		if dir != "" {
			if root, found := findProjectRootFromDir(dir); found {
				return root, true
			}
		}
	}

	return "", false
}

// findProjectRootFromDir searches for project root starting from a specific directory
func findProjectRootFromDir(startDir string) (string, bool) {
	dir := startDir
	for {
		for _, file := range projectFiles {
			if fileExists(filepath.Join(dir, file)) {
				return dir, true
			}
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			// Reached filesystem root
			break
		}
		dir = parent
	}

	return "", false
}
