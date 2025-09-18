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
package resources

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"bennypowers.dev/cem/mcp/types"
	"github.com/tidwall/gjson"
)

// PathTraversalEngine handles declarative data path resolution
type PathTraversalEngine struct {
	// Future: could add caching, optimizations, etc.
}

// NewPathTraversalEngine creates a new path traversal engine
func NewPathTraversalEngine() *PathTraversalEngine {
	return &PathTraversalEngine{}
}

// ResolvePath resolves a JSONPath expression against a data source
func (e *PathTraversalEngine) ResolvePath(source, path string, sources map[string]any) (any, error) {
	// Get the source data
	sourceData, exists := sources[source]
	if !exists {
		return nil, fmt.Errorf("data source '%s' not found", source)
	}

	// Handle variable substitution in path ($.varName)
	resolvedPath, err := e.resolvePathVariables(path, sources["args"])
	if err != nil {
		return nil, fmt.Errorf("failed to resolve path variables: %w", err)
	}

	// Convert source data to JSON for gjson processing
	sourceJSON, err := json.Marshal(sourceData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal source data: %w", err)
	}

	// Use gjson to resolve the path
	result := gjson.GetBytes(sourceJSON, resolvedPath)
	finalPath := resolvedPath
	if !result.Exists() {
		// Handle special cases where gjson path might need adjustment
		if strings.Contains(resolvedPath, "[") {
			// This might be a map access like elements[button-element]
			// Convert to gjson map syntax: elements.button-element
			adjustedPath := convertMapAccessToGjsonPath(resolvedPath)
			finalPath = adjustedPath
			result = gjson.GetBytes(sourceJSON, adjustedPath)
		}

		if !result.Exists() {
			return nil, fmt.Errorf("path '%s' not found in source '%s'", finalPath, source)
		}
	}

	// Parse the result back to Go types
	var parsedResult any
	if err := json.Unmarshal([]byte(result.Raw), &parsedResult); err != nil {
		return nil, fmt.Errorf("failed to parse result: %w", err)
	}

	return parsedResult, nil
}

// ResolvePathWithFilter resolves a path and applies a filter to the result
func (e *PathTraversalEngine) ResolvePathWithFilter(source, path, filter string, sources map[string]any) (any, error) {
	result, err := e.ResolvePath(source, path, sources)
	if err != nil {
		// For exists filter, return false instead of error when path doesn't exist
		if filter == "exists" {
			return false, nil
		}
		return nil, err
	}

	if filter == "" {
		return result, nil
	}

	return e.applyFilter(result, filter)
}

// ExecuteDataFetcher executes a single DataFetcher using the new path-based system
func (e *PathTraversalEngine) ExecuteDataFetcher(fetcher types.DataFetcher, sources map[string]any) (any, error) {
	// Validate required fields
	if fetcher.Source == "" {
		return nil, fmt.Errorf("fetcher '%s' missing required 'source' field", fetcher.Name)
	}

	// If path is empty, return the source data directly
	var result any
	var err error
	if fetcher.Path == "" {
		sourceData, exists := sources[fetcher.Source]
		if !exists {
			return nil, fmt.Errorf("data source '%s' not found", fetcher.Source)
		}
		result = sourceData
		// Apply filter if specified
		if fetcher.Filter != "" {
			result, err = e.applyFilter(result, fetcher.Filter)
		}
	} else {
		// Resolve the path
		result, err = e.ResolvePathWithFilter(fetcher.Source, fetcher.Path, fetcher.Filter, sources)
	}
	if err != nil {
		if fetcher.Required {
			return nil, fmt.Errorf("required fetcher '%s' failed: %w", fetcher.Name, err)
		}
		return nil, nil // Non-required fetcher failure returns nil
	}

	return result, nil
}

// ExecuteDataFetchers executes multiple data fetchers and returns combined results
func (e *PathTraversalEngine) ExecuteDataFetchers(fetchers []types.DataFetcher, sources map[string]any) (map[string]any, error) {
	results := make(map[string]any)

	for _, fetcher := range fetchers {
		result, err := e.ExecuteDataFetcher(fetcher, sources)
		if err != nil {
			return nil, err
		}

		// Store result and make it available for subsequent fetchers
		if result != nil {
			results[fetcher.Name] = result
			sources[fetcher.Name] = result
		}
	}

	return results, nil
}

// resolvePathVariables replaces $.varName with actual values from args
func (e *PathTraversalEngine) resolvePathVariables(path string, args any) (string, error) {
	if !strings.Contains(path, "$.") {
		return path, nil
	}

	argsMap, ok := args.(map[string]any)
	if !ok {
		return "", fmt.Errorf("args must be a map[string]any for variable resolution")
	}

	resolvedPath := path

	// Find all $. variables in the path
	for strings.Contains(resolvedPath, "$.") {
		start := strings.Index(resolvedPath, "$.")
		if start == -1 {
			break
		}

		// Find the end of the variable name
		end := start + 2
		for end < len(resolvedPath) {
			c := resolvedPath[end]
			if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') || c == '_') {
				break
			}
			end++
		}

		if end == start+2 {
			return "", fmt.Errorf("invalid variable syntax at position %d", start)
		}

		// Extract variable name and get value from args
		varName := resolvedPath[start+2 : end]
		value, exists := argsMap[varName]
		if !exists {
			return "", fmt.Errorf("variable '%s' not found in args", varName)
		}

		// Convert value to string, handling gjson filter context
		var valueStr string

		// Check if we're inside a gjson filter expression #(...)
		if isInGjsonFilterExpression(resolvedPath, start) {
			// For gjson filter expressions, string values don't need quotes
			valueStr = fmt.Sprintf("%v", value)
		} else {
			// For regular path access, use the value as-is
			valueStr = fmt.Sprintf("%v", value)
		}

		// Replace the variable with its value
		resolvedPath = resolvedPath[:start] + valueStr + resolvedPath[end:]
	}

	return resolvedPath, nil
}

// applyFilter applies various filters to transform results
func (e *PathTraversalEngine) applyFilter(result any, filter string) (any, error) {
	switch filter {
	case "first":
		return e.getFirst(result), nil
	case "count":
		return e.getCount(result), nil
	case "exists":
		return e.getExists(result), nil
	case "attributes":
		return e.getElementAttributes(result), nil
	case "packages_with_metadata":
		return e.getPackagesWithMetadata(result), nil
	default:
		return nil, fmt.Errorf("unknown filter: %s", filter)
	}
}

// getFirst returns the first element of an array, or nil if empty/not array
func (e *PathTraversalEngine) getFirst(result any) any {
	v := reflect.ValueOf(result)
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		return result // If not an array, return as-is
	}

	if v.Len() == 0 {
		return nil
	}

	return v.Index(0).Interface()
}

// getCount returns the length of an array or 1 for non-arrays
func (e *PathTraversalEngine) getCount(result any) int {
	v := reflect.ValueOf(result)
	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		return v.Len()
	}
	return 1 // Non-arrays have a count of 1
}

// getExists returns true if the result exists and is not nil
func (e *PathTraversalEngine) getExists(result any) bool {
	if result == nil {
		return false
	}

	v := reflect.ValueOf(result)
	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		return v.Len() > 0
	}

	return true
}

// getElementAttributes extracts attributes from an ElementInfo interface
func (e *PathTraversalEngine) getElementAttributes(result any) any {
	if elementInfo, ok := result.(types.ElementInfo); ok {
		attrs := elementInfo.Attributes()
		result := make([]map[string]any, len(attrs))
		for i, attr := range attrs {
			result[i] = map[string]any{
				"name":        attr.Name,
				"type":        attr.Type,
				"description": attr.Description,
				"default":     attr.Default,
				"fieldName":   attr.FieldName,
				"deprecated":  attr.IsDeprecated(),
			}
		}
		return result
	}
	return result
}

// getPackagesWithMetadata extracts packages data with metadata in the expected format
func (e *PathTraversalEngine) getPackagesWithMetadata(result any) any {
	if registryData, ok := result.(map[string]any); ok {
		packages, hasPackages := registryData["packages"]

		if hasPackages {
			if packagesArray, ok := packages.([]map[string]any); ok {
				return map[string]any{
					"packages": packages,
					"metadata": map[string]any{
						"totalPackages": len(packagesArray),
					},
				}
			}
		}
	}
	return result
}

// filterArrayByName filters an array for items that have a "name" field matching the target
func (e *PathTraversalEngine) filterArrayByName(dataArray []any, targetName string) (any, error) {
	for _, item := range dataArray {
		if itemMap, ok := item.(map[string]any); ok {
			if name, ok := itemMap["name"].(string); ok && name == targetName {
				return item, nil
			}
		}
	}
	return nil, nil // Not found, but not an error
}

// convertMapAccessToGjsonPath converts path syntax like "elements[button-element]" to gjson-compatible "elements.button-element"
func convertMapAccessToGjsonPath(path string) string {
	// Handle simple map access: elements[key] -> elements.key
	result := path
	searchStart := 0

	// Replace [key] with .key for simple map access, but preserve filter expressions
	for {
		start := strings.Index(result[searchStart:], "[")
		if start == -1 {
			break
		}
		start += searchStart

		end := strings.Index(result[start:], "]")
		if end == -1 {
			break
		}
		end += start

		// Extract the key
		key := result[start+1 : end]

		// Skip if this looks like a filter expression (containing ? or @)
		if strings.Contains(key, "?") || strings.Contains(key, "@") {
			// This is a filter expression, leave it as is but continue past it
			searchStart = end + 1
			continue
		}

		// Convert [key] to .key for simple map access
		result = result[:start] + "." + key + result[end+1:]
		// Continue searching from after the replacement
		searchStart = start + 1 + len(key)
	}

	return result
}

// isInGjsonFilterExpression checks if a position in the path is inside a gjson filter expression #(...)
func isInGjsonFilterExpression(path string, position int) bool {
	// Look backwards from position to find the most recent #( or )
	for i := position - 1; i >= 0; i-- {
		if i < len(path)-1 && path[i] == '#' && path[i+1] == '(' {
			// Found opening #(, now look forward for closing )
			for j := position; j < len(path); j++ {
				if path[j] == ')' {
					return true // We're between #( and )
				}
			}
			return false // Found #( but no closing )
		}
		if path[i] == ')' {
			return false // Found ) before #(, we're not in a filter
		}
	}
	return false // No #( found
}