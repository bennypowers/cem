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

package elements

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/evanw/esbuild/pkg/api"
)

// TranspileElements transpiles TypeScript element files to JavaScript
// It looks for .ts files in sourceDir and outputs .js files to the same location
func TranspileElements(sourceDir string) error {
	// Find all TypeScript files in the elements directory
	var tsFiles []string
	err := filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".ts") && !strings.HasSuffix(path, ".test.ts") {
			tsFiles = append(tsFiles, path)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("walking source directory: %w", err)
	}

	if len(tsFiles) == 0 {
		return nil // No TypeScript files to transpile
	}

	// Build entry points for esbuild
	entryPoints := make([]string, len(tsFiles))
	copy(entryPoints, tsFiles)

	// Run esbuild to transpile all files
	result := api.Build(api.BuildOptions{
		EntryPoints: entryPoints,
		Outdir:      sourceDir,
		Bundle:      false, // Don't bundle, just transpile
		Write:       true,  // Write to disk
		Format:      api.FormatESModule,
		Target:      api.ES2020,
		Loader: map[string]api.Loader{
			".ts": api.LoaderTS,
		},
		LogLevel:       api.LogLevelWarning,
		Sourcemap:      api.SourceMapInline,
		OutExtension:   map[string]string{".js": ".js"},
		AllowOverwrite: true, // Allow overwriting existing .js files
	})

	if len(result.Errors) > 0 {
		var errMsgs []string
		for _, err := range result.Errors {
			errMsgs = append(errMsgs, err.Text)
		}
		return fmt.Errorf("esbuild errors: %s", strings.Join(errMsgs, "; "))
	}

	return nil
}
