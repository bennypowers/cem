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
	"fmt"
	"path/filepath"
	"strings"

	"bennypowers.dev/cem/modulegraph"
	"bennypowers.dev/cem/queries"
	"github.com/evanw/esbuild/pkg/api"
)

// Loader specifies the file type for transformation
type Loader string

const (
	LoaderTS  Loader = "ts"
	LoaderTSX Loader = "tsx"
	LoaderJS  Loader = "js"
	LoaderJSX Loader = "jsx"
)

// Target specifies the ECMAScript target version
type Target string

const (
	ES2015 Target = "es2015"
	ES2016 Target = "es2016"
	ES2017 Target = "es2017"
	ES2018 Target = "es2018"
	ES2019 Target = "es2019"
	ES2020 Target = "es2020"
	ES2021 Target = "es2021"
	ES2022 Target = "es2022"
	ES2023 Target = "es2023"
	ESNext Target = "esnext"
)

// ValidTargets returns all valid target values
func ValidTargets() []Target {
	return []Target{ES2015, ES2016, ES2017, ES2018, ES2019, ES2020, ES2021, ES2022, ES2023, ESNext}
}

// IsValidTarget checks if a target string is valid
func IsValidTarget(target string) bool {
	switch Target(target) {
	case ES2015, ES2016, ES2017, ES2018, ES2019, ES2020, ES2021, ES2022, ES2023, ESNext:
		return true
	default:
		return false
	}
}

// SourceMapMode specifies how source maps are generated
type SourceMapMode string

const (
	SourceMapInline   SourceMapMode = "inline"
	SourceMapExternal SourceMapMode = "external"
	SourceMapNone     SourceMapMode = "none"
)

// TransformOptions configures the transformation
type TransformOptions struct {
	Loader      Loader
	Target      Target
	Sourcemap   SourceMapMode
	TsconfigRaw string // Optional tsconfig.json content as JSON string
	Sourcefile  string // Original source file path for source maps
}

// TransformTypeScript transforms TypeScript source code to JavaScript using esbuild
func TransformTypeScript(source []byte, opts TransformOptions) (*TransformResult, error) {
	// Convert our options to esbuild options
	loader := api.LoaderTS
	switch opts.Loader {
	case LoaderTSX:
		loader = api.LoaderTSX
	case LoaderJS:
		loader = api.LoaderJS
	case LoaderJSX:
		loader = api.LoaderJSX
	}

	target := api.ES2020
	switch opts.Target {
	case ES2015:
		target = api.ES2015
	case ES2016:
		target = api.ES2016
	case ES2017:
		target = api.ES2017
	case ES2018:
		target = api.ES2018
	case ES2019:
		target = api.ES2019
	case ES2020:
		target = api.ES2020
	case ES2021:
		target = api.ES2021
	case ES2022:
		target = api.ES2022
	case ES2023:
		target = api.ES2023
	case ESNext:
		target = api.ESNext
	}

	sourcemap := api.SourceMapInline
	switch opts.Sourcemap {
	case SourceMapExternal:
		sourcemap = api.SourceMapExternal
	case SourceMapNone:
		sourcemap = api.SourceMapNone
	}

	// Use provided tsconfig or default configuration
	tsconfigRaw := opts.TsconfigRaw
	if tsconfigRaw == "" {
		// Default: inline helpers to avoid tslib dependency issues
		tsconfigRaw = `{
			"compilerOptions": {
				"importHelpers": false
			}
		}`
	}

	// Transform using esbuild
	result := api.Transform(string(source), api.TransformOptions{
		Loader:      loader,
		Target:      target,
		Format:      api.FormatESModule,
		Sourcemap:   sourcemap,
		Sourcefile:  opts.Sourcefile, // Set source file path for better source maps
		TsconfigRaw: tsconfigRaw,
	})

	// Check for errors
	if len(result.Errors) > 0 {
		// Format errors for user
		errMsg := "Transform failed:\n"
		for _, err := range result.Errors {
			errMsg += fmt.Sprintf("  %s\n", err.Text)
		}
		return nil, fmt.Errorf("%s", errMsg)
	}

	// Extract dependencies from source code using tree-sitter
	dependencies := extractDependencies(source, opts.Sourcefile)

	// Return result
	return &TransformResult{
		Code:         result.Code,
		Map:          result.Map,
		Dependencies: dependencies,
	}, nil
}

// extractDependencies extracts import paths from TypeScript/JavaScript source using tree-sitter
func extractDependencies(source []byte, sourcePath string) []string {
	// Create a dependency tracker to collect imports
	dependencyTracker := modulegraph.NewDependencyTracker()
	exportTracker := modulegraph.NewExportTracker()

	// Get global QueryManager for tree-sitter parsing
	queryManager, err := queries.GetGlobalQueryManager()
	if err != nil {
		// If QueryManager fails to initialize, dependency tracking won't work
		// but the transform can still succeed
		return nil
	}

	// Use the DefaultExportParser to parse imports/exports
	parser := &modulegraph.DefaultExportParser{}
	err = parser.ParseExportsFromContent(sourcePath, source, exportTracker, dependencyTracker, queryManager)
	if err != nil {
		// Log error but don't fail the transform - dependency tracking is best-effort
		// The transform can still succeed even if we can't track dependencies
		return nil
	}

	// Get the dependencies that were found
	deps := dependencyTracker.GetModuleDependencies(sourcePath)

	// Resolve relative imports to absolute paths
	resolvedDeps := make([]string, 0, len(deps))
	for _, dep := range deps {
		if isRelativeImport(dep) {
			resolved := resolveImport(sourcePath, dep)
			resolvedDeps = append(resolvedDeps, resolved)
		}
		// Skip bare specifiers (e.g., 'lit', '@rhds/elements') as they're not local files
	}

	return resolvedDeps
}

// isRelativeImport checks if an import path is relative (./ or ../)
// Also handles normalized relative imports where the ./ or ../ prefix was removed
func isRelativeImport(path string) bool {
	// Check for explicit relative path markers
	if strings.HasPrefix(path, "./") || strings.HasPrefix(path, "../") {
		return true
	}

	// Handle normalized relative imports (where ./ or ../ prefix was removed by NormalizeImportPath)
	// A path is NOT relative if it's a bare specifier (npm package):
	// 1. Scoped package starting with @ (e.g., @rhds/elements)
	// 2. Simple package name without / and without file extension (e.g., lit)
	// 3. Absolute path starting with /

	if strings.HasPrefix(path, "@") || strings.HasPrefix(path, "/") {
		return false // Scoped package or absolute path
	}

	if !strings.Contains(path, "/") {
		// If there's no /, it could be either:
		// - A bare specifier like 'lit' (no extension)
		// - A normalized relative import like 'component.css' (has extension)
		hasExtension := strings.Contains(path, ".")
		return hasExtension // True if it's a file with extension, false if it's a bare package
	}

	// If we get here, it's a path like "shared/utils.js"
	// which are normalized relative imports
	return true
}

// resolveImport resolves a relative import path against a source file
func resolveImport(sourcePath, importPath string) string {
	// Get directory of source file
	sourceDir := filepath.Dir(sourcePath)

	// Join and clean the path
	resolved := filepath.Join(sourceDir, importPath)

	// Normalize separators
	resolved = filepath.Clean(resolved)

	return resolved
}

// stringToTemplateLiteral escapes a string for safe inclusion in a JS template literal
// Based on Lit's stringToTemplateLiteral: /\\|`|\$(?={)|(?<=<)\//g
func stringToTemplateLiteral(str string) string {
	var result strings.Builder
	result.Grow(len(str) + 20) // Pre-allocate with some buffer for escapes

	prevChar := rune(0)
	for i, char := range str {
		switch char {
		case '\\', '`':
			// Escape backslashes and backticks
			result.WriteRune('\\')
			result.WriteRune(char)
		case '$':
			// Only escape $ if followed by {  (\$(?={))
			if i+1 < len(str) && str[i+1] == '{' {
				result.WriteString("\\$")
			} else {
				result.WriteRune(char)
			}
		case '/':
			// Escape / if preceded by <  ((?<=<)\/)
			if prevChar == '<' {
				result.WriteString("\\/")
			} else {
				result.WriteRune(char)
			}
		default:
			result.WriteRune(char)
		}
		prevChar = char
	}

	return result.String()
}

// TransformCSS transforms CSS to a JavaScript module exporting a CSSStyleSheet
func TransformCSS(source []byte, path string) string {
	css := stringToTemplateLiteral(string(source))

	// Use embedded template with source map
	// Template format matches Lit's CSS module pattern
	return fmt.Sprintf(`// [served] %s
const sheet = new CSSStyleSheet();
sheet.replaceSync(%s);
export default sheet;
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbXSwibmFtZXMiOltdLCJtYXBwaW5ncyI6IiJ9
`, path, "`"+css+"`")
}
