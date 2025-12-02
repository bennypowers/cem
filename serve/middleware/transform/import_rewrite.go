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
	"net/url"
	"strings"

	"bennypowers.dev/cem/queries"
	ts "github.com/tree-sitter/go-tree-sitter"
)

// ImportAttribute represents a single import attribute key-value pair
type ImportAttribute struct {
	Key   string
	Value string
}

// ImportRewrite represents a detected import with attributes that needs rewriting
type ImportRewrite struct {
	OriginalSource string            // The original import path (e.g., './foo.css')
	Attributes     []ImportAttribute // The import attributes
	SourceNode     *ts.Node          // The source string node for rewriting
	StartByte      uint              // Start byte of the source string
	EndByte        uint              // End byte of the source string
}

// RewriteImportAttributes rewrites TypeScript/JavaScript source to convert import attributes
// into query parameters that can be preserved through esbuild transformation.
//
// Example transformation:
//
//	import styles from './foo.css' with { type: 'css' }
//	  becomes:
//	import styles from './foo.css?__cem-import-attrs[type]=css'
//
// This allows the dev server to detect and handle these imports appropriately.
func RewriteImportAttributes(source []byte) ([]byte, error) {
	// Get a TypeScript parser
	parser := queries.RetrieveTypeScriptParser()
	defer queries.PutTypeScriptParser(parser)

	// Parse the source
	tree := parser.Parse(source, nil)
	if tree == nil {
		// If parsing fails, return original source unchanged
		return source, nil
	}
	defer tree.Close()

	// Get the global query manager
	queryManager, err := queries.GetGlobalQueryManager()
	if err != nil {
		// If query manager unavailable, return original source unchanged
		return source, nil
	}

	// Create a query matcher for import attributes
	// Note: This query needs to be added to the imports query or created as a new query
	matcher, err := queries.NewQueryMatcher(queryManager, "typescript", "importAttributes")
	if err != nil {
		// If query not available, return original source unchanged
		return source, nil
	}
	defer matcher.Close()

	// Collect all imports that need rewriting
	rewrites := findImportsWithAttributes(tree.RootNode(), source, matcher)

	// If no rewrites needed, return original source
	if len(rewrites) == 0 {
		return source, nil
	}

	// Apply rewrites in reverse order (from end to start) to preserve byte offsets
	result := make([]byte, len(source))
	copy(result, source)

	for i := len(rewrites) - 1; i >= 0; i-- {
		rewrite := rewrites[i]
		newSource := buildRewrittenImportPath(rewrite.OriginalSource, rewrite.Attributes)

		// Replace the source string content (keeping the quotes)
		// The source node includes the quotes, so we need to preserve them
		before := result[:rewrite.StartByte]
		after := result[rewrite.EndByte:]

		// Extract the quote character (first byte of the source node)
		quoteChar := source[rewrite.StartByte]

		// Build new result with rewritten import path
		newResult := make([]byte, 0, len(before)+len(newSource)+len(after)+2) // +2 for quotes
		newResult = append(newResult, before...)
		newResult = append(newResult, quoteChar)
		newResult = append(newResult, []byte(newSource)...)
		newResult = append(newResult, quoteChar)
		newResult = append(newResult, after...)

		result = newResult
	}

	return result, nil
}

// findImportsWithAttributes finds all import statements with attributes in the AST
func findImportsWithAttributes(root *ts.Node, source []byte, matcher *queries.QueryMatcher) []ImportRewrite {
	var rewrites []ImportRewrite

	// Use ParentCaptures to group all captures by import statement
	for captureMap := range matcher.ParentCaptures(root, source, "import.with.attrs") {
		var rewrite ImportRewrite
		var sourceNodeID int

		// Extract the source string node
		if sourceNodes, ok := captureMap["import.source.node"]; ok && len(sourceNodes) > 0 {
			sourceNodeID = sourceNodes[0].NodeId
			sourceText := sourceNodes[0].Text
			// Remove quotes from source text
			if len(sourceText) >= 2 {
				sourceText = sourceText[1 : len(sourceText)-1]
			}
			rewrite.OriginalSource = sourceText
			rewrite.StartByte = sourceNodes[0].StartByte
			rewrite.EndByte = sourceNodes[0].EndByte
		}

		// Extract all attribute key-value pairs
		if attrKeys, ok := captureMap["import.attr.key"]; ok {
			if attrValues, ok := captureMap["import.attr.value"]; ok {
				// Match up keys and values
				// Note: The tree-sitter query captures string_fragment nodes,
				// so we get the content without quotes
				for i := range attrKeys {
					if i < len(attrValues) {
						key := attrKeys[i].Text
						value := attrValues[i].Text

						rewrite.Attributes = append(rewrite.Attributes, ImportAttribute{
							Key:   key,
							Value: value,
						})
					}
				}
			}
		}

		// Only add if we have both a source and attributes
		if rewrite.OriginalSource != "" && len(rewrite.Attributes) > 0 {
			// Get the actual node for potential future use
			if sourceNodeID != 0 {
				rewrite.SourceNode = queries.GetDescendantById(root, sourceNodeID)
			}
			rewrites = append(rewrites, rewrite)
		}
	}

	return rewrites
}

// buildRewrittenImportPath constructs a new import path with query parameters
// encoding the import attributes.
//
// Example: './foo.css' with attributes [{type: 'css'}]
//
//	becomes: './foo.css?__cem-import-attrs[type]=css'
func buildRewrittenImportPath(originalPath string, attributes []ImportAttribute) string {
	if len(attributes) == 0 {
		return originalPath
	}

	var queryParams []string
	for _, attr := range attributes {
		// URL-encode both key and value for safety
		encodedKey := url.QueryEscape(attr.Key)
		encodedValue := url.QueryEscape(attr.Value)
		param := "__cem-import-attrs[" + encodedKey + "]=" + encodedValue
		queryParams = append(queryParams, param)
	}

	separator := "?"
	if strings.Contains(originalPath, "?") {
		separator = "&"
	}

	return originalPath + separator + strings.Join(queryParams, "&")
}
