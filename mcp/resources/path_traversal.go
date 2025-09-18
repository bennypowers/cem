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
	"fmt"
	"reflect"
	"strings"

	"bennypowers.dev/cem/manifest/traversal"
	"bennypowers.dev/cem/mcp/types"
)

// PathTraversalEngine handles MCP-specific data path resolution
// extending the base manifest traversal engine with MCP-specific features
type PathTraversalEngine struct {
	engine *traversal.Engine
}

// NewPathTraversalEngine creates a new MCP path traversal engine
func NewPathTraversalEngine() *PathTraversalEngine {
	return &PathTraversalEngine{
		engine: traversal.NewEngine(),
	}
}

// ResolvePath delegates to the base engine with MCP variable resolution
func (e *PathTraversalEngine) ResolvePath(source, path string, sources map[string]any) (any, error) {
	// Handle MCP-specific variable substitution in path ($.varName)
	resolvedPath := path
	if strings.Contains(path, "$.") {
		var err error
		resolvedPath, err = e.resolvePathVariables(path, sources["args"])
		if err != nil {
			return nil, fmt.Errorf("failed to resolve path variables: %w", err)
		}
	}

	return e.engine.ResolvePath(source, resolvedPath, sources)
}

// ResolvePathWithFilter resolves a path and applies a filter to the result (backward compatibility)
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

	return e.applyMCPFilter(result, filter)
}

// ExecuteDataFetcher executes a single DataFetcher using MCP-specific extensions
func (e *PathTraversalEngine) ExecuteDataFetcher(fetcher types.DataFetcher, sources map[string]any) (any, error) {
	// Pre-process path variables (MCP-specific feature)
	resolvedPath := fetcher.Path
	if strings.Contains(resolvedPath, "$.") {
		var err error
		resolvedPath, err = e.resolvePathVariables(resolvedPath, sources["args"])
		if err != nil {
			return nil, fmt.Errorf("failed to resolve path variables: %w", err)
		}
	}

	// Convert to manifest traversal.DataFetcher
	manifestFetcher := traversal.DataFetcher{
		Name:     fetcher.Name,
		Source:   fetcher.Source,
		Path:     resolvedPath,
		Filter:   "", // We'll handle filters separately for MCP-specific ones
		Required: fetcher.Required,
	}

	// Execute using base engine
	result, err := e.engine.ExecuteDataFetcher(manifestFetcher, sources)
	if err != nil {
		return nil, err
	}

	// Apply MCP-specific filters if needed
	if fetcher.Filter != "" {
		result, err = e.applyMCPFilter(result, fetcher.Filter)
		if err != nil {
			return nil, err
		}
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
			if (c < 'a' || c > 'z') && (c < 'A' || c > 'Z') && (c < '0' || c > '9') && c != '_' {
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

// applyMCPFilter applies MCP-specific filters to transform results
func (e *PathTraversalEngine) applyMCPFilter(result any, filter string) (any, error) {
	switch filter {
	case "attributes":
		return e.getElementAttributes(result), nil
	case "packages_with_metadata":
		return e.getPackagesWithMetadata(result), nil
	case "first":
		return e.getFirst(result), nil
	case "count":
		return e.getCount(result), nil
	case "exists":
		return e.getExists(result), nil
	default:
		return nil, fmt.Errorf("unknown MCP filter: %s", filter)
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