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
package traversal

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/tidwall/gjson"
)

// Engine handles generic data path resolution using JSONPath expressions
type Engine struct {
	// Future: could add caching, optimizations, etc.
}

// NewEngine creates a new path traversal engine
func NewEngine() *Engine {
	return &Engine{}
}

// ResolvePath resolves a JSONPath expression against a data source
func (e *Engine) ResolvePath(source, path string, sources map[string]any) (any, error) {
	// Get the source data
	sourceData, exists := sources[source]
	if !exists {
		return nil, fmt.Errorf("source '%s' not found", source)
	}

	// If no path specified, return the entire source
	if path == "" {
		return sourceData, nil
	}

	// Convert to JSON for gjson processing
	jsonBytes, err := json.Marshal(sourceData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal source data: %w", err)
	}

	// Use gjson to resolve the path
	result := gjson.GetBytes(jsonBytes, path)
	finalPath := path
	if !result.Exists() {
		// Handle special cases where gjson path might need adjustment
		if strings.Contains(path, "[") && !strings.Contains(path, "?") && !strings.Contains(path, "@") {
			// This might be a map access like elements[button-element]
			// Convert to gjson map syntax: elements.button-element
			adjustedPath := convertMapAccessToGjsonPath(path)
			finalPath = adjustedPath
			result = gjson.GetBytes(jsonBytes, adjustedPath)
		}

		if !result.Exists() {
			return nil, fmt.Errorf("path '%s' not found in source '%s'", finalPath, source)
		}
	}

	// Convert result back to Go types
	var value any
	if err := json.Unmarshal([]byte(result.Raw), &value); err != nil {
		// If unmarshaling fails, return the raw string value
		return result.String(), nil
	}

	return value, nil
}

// ExecuteDataFetcher executes a single data fetcher against sources
func (e *Engine) ExecuteDataFetcher(fetcher DataFetcher, sources map[string]any) (any, error) {
	result, err := e.ResolvePath(fetcher.Source, fetcher.Path, sources)
	if err != nil {
		return nil, err
	}

	// Apply filters if specified
	if fetcher.Filter != "" {
		result = e.applyFilter(result, fetcher.Filter)
	}

	return result, nil
}

// applyFilter applies filtering logic to results
func (e *Engine) applyFilter(result any, filter string) any {
	switch filter {
	case "first":
		if reflect.TypeOf(result).Kind() == reflect.Slice {
			v := reflect.ValueOf(result)
			if v.Len() > 0 {
				return v.Index(0).Interface()
			}
		}
		return result
	case "count":
		if reflect.TypeOf(result).Kind() == reflect.Slice {
			return reflect.ValueOf(result).Len()
		}
		return 1
	case "exists":
		return result != nil
	default:
		// Unknown filter, return as-is
		return result
	}
}

// DataFetcher represents a generic data fetcher configuration
type DataFetcher struct {
	Name     string `yaml:"name"`
	Source   string `yaml:"source"`   // Data source name
	Path     string `yaml:"path"`     // JSONPath/gjson expression
	Filter   string `yaml:"filter"`   // Optional: first, count, exists
	Required bool   `yaml:"required"` // Whether this fetcher is required
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
