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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type contextKey string

const workspaceContextKey = contextKey("workspaceContext")

func GetWorkspaceContext(cmd *cobra.Command) (WorkspaceContext, error) {
	val := cmd.Context().Value(workspaceContextKey)
	if val != nil {
		return val.(WorkspaceContext), nil
	}

	configPath := viper.GetString("configFile")
	if p, _ := cmd.Flags().GetString("configFile"); p != "" {
		configPath = p
	}
	spec := viper.GetString("package")
	if p, _ := cmd.Flags().GetString("package"); p != "" {
		spec = p
	}

	// fall back to previous behavior: use the directory above .config
	if spec == "" {
		spec = filepath.Dir(configPath)
	}

	if ctx, err := getAppropriateContextForSpec(spec, cmd.Name()); err != nil {
		return nil, err
	} else if err := ctx.Init(); err != nil {
		return nil, err
	} else {
		// Store in cmd.Context for future calls
		newCtx := context.WithValue(cmd.Context(), workspaceContextKey, ctx)
		cmd.SetContext(newCtx)
		return ctx, nil
	}
}

func getAppropriateContextForSpec(spec, cmdName string) (ctx WorkspaceContext, err error) {
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
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
