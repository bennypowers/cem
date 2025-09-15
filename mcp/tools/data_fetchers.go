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
package tools

import (
	"encoding/json"
	"fmt"
	"strings"

	"bennypowers.dev/cem/mcp/types"
	V "bennypowers.dev/cem/validate"
	"github.com/tidwall/gjson"
)

// FetchedData represents the result of a data fetcher
type FetchedData map[string]interface{}

// ExecuteDataFetchers runs all data fetchers and returns the collected data
func ExecuteDataFetchers(fetchers []types.DataFetcher, registry types.MCPContext, args map[string]interface{}) (FetchedData, error) {
	data := make(FetchedData)

	for _, fetcher := range fetchers {
		result, err := executeSingleFetcher(fetcher, registry, args)
		if err != nil {
			if fetcher.Required {
				return nil, fmt.Errorf("required data fetcher '%s' failed: %w", fetcher.Name, err)
			}
			// Optional fetcher failed, continue
			continue
		}
		data[fetcher.Name] = result
	}

	return data, nil
}

// executeSingleFetcher executes one data fetcher based on its type
func executeSingleFetcher(fetcher types.DataFetcher, registry types.MCPContext, args map[string]interface{}) (interface{}, error) {
	// Interpolate template variables in the path
	path := interpolatePath(fetcher.Path, args)

	switch fetcher.Type {
	case "manifest_element":
		return fetchManifestElement(registry, path, args)
	case "schema_definitions":
		return fetchSchemaDefinitions(registry, path)
	case "element_collection":
		return fetchElementCollection(registry, path)
	case "attribute_collection":
		return fetchAttributeCollection(registry, path, args)
	default:
		return nil, fmt.Errorf("unknown data fetcher type: %s", fetcher.Type)
	}
}

// interpolatePath replaces template variables in paths with actual values
func interpolatePath(path string, args map[string]interface{}) string {
	result := path
	for key, value := range args {
		placeholder := "{" + key + "}"
		result = strings.ReplaceAll(result, placeholder, fmt.Sprintf("%v", value))
	}
	return result
}

// fetchManifestElement fetches a single element using gjson path
func fetchManifestElement(registry types.MCPContext, path string, args map[string]interface{}) (interface{}, error) {
	// Get the tag name from args for element lookup
	tagName, ok := args["tagName"].(string)
	if !ok {
		return nil, fmt.Errorf("tagName not found in arguments")
	}

	// Use the existing element lookup which already handles type conversion
	element, err := registry.ElementInfo(tagName)
	if err != nil {
		return nil, fmt.Errorf("element '%s' not found", tagName)
	}

	return element, nil
}

// fetchSchemaDefinitions fetches schema definitions for template functions
func fetchSchemaDefinitions(registry types.MCPContext, path string) (interface{}, error) {
	// Get schema versions from manifests
	versions := registry.GetManifestSchemaVersions()
	if len(versions) == 0 {
		return make(map[string]interface{}), nil
	}

	// Use the first version to load schema (same as existing implementation)
	schemaVersion := versions[0]
	schemaData, err := getSchemaData(schemaVersion)
	if err != nil {
		return nil, fmt.Errorf("failed to load schema %s: %w", schemaVersion, err)
	}

	// Parse schema JSON
	var schema map[string]interface{}
	if err := json.Unmarshal(schemaData, &schema); err != nil {
		return nil, fmt.Errorf("failed to parse schema JSON: %w", err)
	}

	// Use gjson to extract the requested path if specified
	if path != "" && path != "." {
		schemaJSON := string(schemaData)
		result := gjson.Get(schemaJSON, path)
		if !result.Exists() {
			return make(map[string]interface{}), nil
		}
		return result.Value(), nil
	}

	return schema, nil
}

// fetchElementCollection fetches multiple elements using gjson
func fetchElementCollection(registry types.MCPContext, path string) (interface{}, error) {
	// Get all elements from registry
	allElements := registry.AllElements()

	// Convert to JSON for gjson processing
	elementsJSON, err := json.Marshal(allElements)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal elements: %w", err)
	}

	// Use gjson to extract elements based on path
	result := gjson.Get(string(elementsJSON), path)
	if !result.Exists() {
		return []interface{}{}, nil
	}

	return result.Value(), nil
}

// fetchAttributeCollection fetches attributes for a specific element
func fetchAttributeCollection(registry types.MCPContext, path string, args map[string]interface{}) (interface{}, error) {
	tagName, ok := args["tagName"].(string)
	if !ok {
		return nil, fmt.Errorf("tagName not found in arguments")
	}

	element, err := registry.ElementInfo(tagName)
	if err != nil {
		return nil, fmt.Errorf("element '%s' not found", tagName)
	}

	// Return the attributes directly
	return element.Attributes(), nil
}

// getSchemaData retrieves schema data using the existing schema loading mechanism
func getSchemaData(version string) ([]byte, error) {
	// Use the existing V.GetSchema function from the validate package
	return V.GetSchema(version)
}