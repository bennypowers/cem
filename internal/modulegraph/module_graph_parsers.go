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
package modulegraph

import (
	"fmt"
	"io/fs"
	"strings"

	"bennypowers.dev/cem/internal/languages/typescript"
	"bennypowers.dev/cem/internal/platform"
	"bennypowers.dev/cem/internal/treesitter"
	"bennypowers.dev/cem/lsp/helpers"
	ts "github.com/tree-sitter/go-tree-sitter"
)

// OSFileParser implements FileParser using the platform.FileSystem abstraction.
type OSFileParser struct {
	FS platform.FileSystem
}

func NewOSFileParser(fsys platform.FileSystem) *OSFileParser {
	if fsys == nil {
		fsys = platform.NewOSFileSystem()
	}
	return &OSFileParser{FS: fsys}
}

func (p *OSFileParser) ReadFile(path string) ([]byte, error) {
	return p.FS.ReadFile(path)
}

func (p *OSFileParser) Exists(path string) bool {
	return p.FS.Exists(path)
}

func (p *OSFileParser) WorkspaceFS(workspaceRoot string) fs.FS {
	return platform.DirFS(p.FS, workspaceRoot)
}


// DefaultExportParser implements ExportParser using tree-sitter queries
type DefaultExportParser struct{}

// ParseExportsFromContent implements ExportParser using tree-sitter
func (p *DefaultExportParser) ParseExportsFromContent(modulePath string, content []byte, exportTracker *ExportTracker, dependencyTracker *DependencyTracker, queryManager *treesitter.QueryManager) error {
	// Get TypeScript parser from pool
	parser := typescript.BorrowParser()
	defer typescript.ReturnParser(parser)

	// Parse the content
	tree := parser.Parse(content, nil)
	if tree == nil {
		// Enhanced error reporting with content analysis
		contentSize := len(content)
		hasContent := contentSize > 0

		var errorDetails string
		if !hasContent {
			errorDetails = "empty file"
		} else if contentSize > 100000 {
			errorDetails = fmt.Sprintf("file too large (%d bytes)", contentSize)
		} else {
			// Check for common syntax issues by examining first few lines
			previewSize := min(contentSize, 500)
			lines := strings.Split(string(content[:previewSize]), "\n")

			// Check for common non-JS/TS file types
			checkSize := min(contentSize, 100)

			if len(lines) > 0 && strings.Contains(lines[0], "<?xml") {
				errorDetails = "XML file detected (not TypeScript/JavaScript)"
			} else if strings.Contains(string(content[:checkSize]), "\x00") {
				errorDetails = "binary file detected"
			} else {
				errorDetails = "syntax errors or invalid TypeScript/JavaScript syntax"
			}
		}

		helpers.SafeDebugLog("[MODULE_GRAPH] Failed to parse file '%s' - %s (size: %d bytes)", modulePath, errorDetails, contentSize)
		return fmt.Errorf("failed to parse file '%s': %s", modulePath, errorDetails)
	}
	defer tree.Close()

	// Try to use tree-sitter queries for accurate parsing
	return p.parseExportsWithQueries(tree, modulePath, content, dependencyTracker, queryManager)
}

// parseExportsWithQueries uses tree-sitter queries to parse export statements
func (p *DefaultExportParser) parseExportsWithQueries(tree *ts.Tree, modulePath string, content []byte, dependencyTracker *DependencyTracker, queryManager *treesitter.QueryManager) error {
	// PERFORMANCE OPTIMIZATION: Use injected QueryManager for dependency injection
	if queryManager == nil {
		return fmt.Errorf("query manager not available for parsing exports in module %s", modulePath)
	}

	// Get cached export matcher (thread-safe)
	exportMatcher, err := treesitter.GetCachedQueryMatcher(queryManager, "typescript", "exports")
	if err != nil {
		return fmt.Errorf("failed to get cached export matcher for module %s: %w", modulePath, err)
	}
	// Note: Don't defer Close() on cached singleton instances

	// Process export matches
	for match := range exportMatcher.AllQueryMatches(tree.RootNode(), content) {
		p.processExportMatch(match, exportMatcher, modulePath, content, dependencyTracker)
	}

	return nil
}

// processExportMatch processes a single export/import match from tree-sitter
func (p *DefaultExportParser) processExportMatch(match *ts.QueryMatch, matcher *treesitter.QueryMatcher, modulePath string, content []byte, dependencyTracker *DependencyTracker) {
	var exportName, sourceModule string
	var importSource string
	var isImport bool

	// Process all captures in this match
	for _, capture := range match.Captures {
		captureName := matcher.GetCaptureNameByIndex(capture.Index)
		captureText := strings.TrimSpace(capture.Node.Utf8Text(content))

		switch captureName {
		case "export.name":
			exportName = captureText
		case "export.source":
			sourceModule = captureText
		case "export.class.name", "export.function.name", "export.variable.name", "export.default.name":
			exportName = captureText
		case "import.source", "import.dynamic.source":
			importSource = captureText
			isImport = true
		}
	}

	// Handle import statements - track module dependencies
	if isImport && importSource != "" {
		// Normalize the import path to handle relative imports
		normalizedImportPath := helpers.NormalizeImportPath(importSource)
		dependencyTracker.AddModuleDependency(modulePath, normalizedImportPath)
		return
	}

	// Handle re-exports (export { X } from './module')
	if exportName != "" && sourceModule != "" {
		// Track re-export relationship at module level - this is sufficient because
		// resolveReExportChains() will later resolve all elements from source modules
		// to re-exporting modules using existing manifest data
		normalizedSourceModule := helpers.NormalizeImportPath(sourceModule)
		dependencyTracker.AddModuleDependency(modulePath, normalizedSourceModule)
		dependencyTracker.AddReExportChain(modulePath, normalizedSourceModule)
		return
	}

	// Handle direct exports (export class MyElement)
	if exportName != "" && sourceModule == "" {
		// Direct exports are handled by the existing manifest registry integration.
		// The registry already contains element definitions from manifests by module path,
		// so individual export tracking here would be redundant. The LSP system uses
		// the manifest registry as the authoritative source for element definitions.
		return
	}
}

// NoOpManifestResolver implements ManifestResolver but returns empty results
type NoOpManifestResolver struct{}

func (r *NoOpManifestResolver) FindManifestModulesForImportPath(importPath string) []string {
	return nil
}

func (r *NoOpManifestResolver) GetManifestModulePath(filePath string) string {
	return ""
}

func (r *NoOpManifestResolver) GetElementsFromManifestModule(manifestModulePath string) []string {
	return nil
}
