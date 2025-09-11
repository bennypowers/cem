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
package resources_test

import (
	"context"
	"encoding/json"
	"testing"

	"bennypowers.dev/cem/mcp"
	"bennypowers.dev/cem/mcp/resources"
	W "bennypowers.dev/cem/workspace"
	mcpSDK "github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// getTestRegistry creates a registry using the test fixtures following the existing pattern
func getTestRegistry(t *testing.T) *mcp.Registry {
	workspace := W.NewFileSystemWorkspaceContext("../test-fixtures/multiple-elements")
	err := workspace.Init()
	require.NoError(t, err)

	registry, err := mcp.NewRegistry(workspace)
	require.NoError(t, err)

	err = registry.LoadManifests()
	require.NoError(t, err)

	return registry
}

func TestResources_LoadingAndParsing(t *testing.T) {
	registry := getTestRegistry(t)
	registryAdapter := mcp.NewRegistryAdapter(registry)

	resourceDefs, err := resources.Resources(registryAdapter)
	require.NoError(t, err)
	require.NotEmpty(t, resourceDefs, "Should load resource definitions")

	// Verify all expected resources are loaded
	expectedResources := []string{"schema", "packages", "elements", "element", "guidelines", "accessibility"}
	resourceNames := make(map[string]bool)
	for _, def := range resourceDefs {
		resourceNames[def.Name] = true

		// Basic validation of each resource definition
		assert.NotEmpty(t, def.URI, "Resource %s should have URI", def.Name)
		assert.NotEmpty(t, def.MimeType, "Resource %s should have MIME type", def.Name)
		assert.NotNil(t, def.Handler, "Resource %s should have handler", def.Name)
	}

	for _, expected := range expectedResources {
		assert.True(t, resourceNames[expected], "Should have resource: %s", expected)
	}
}

func TestSchemaResource_Integration(t *testing.T) {
	registry := getTestRegistry(t)
	registryAdapter := mcp.NewRegistryAdapter(registry)

	resourceDefs, err := resources.Resources(registryAdapter)
	require.NoError(t, err)

	// Find schema resource
	var schemaResource *resources.ResourceDefinition
	for _, def := range resourceDefs {
		if def.Name == "schema" {
			schemaResource = &def
			break
		}
	}
	require.NotNil(t, schemaResource, "Should have schema resource")

	// Test schema resource handler
	req := &mcpSDK.ReadResourceRequest{
		Params: &mcpSDK.ReadResourceParams{
			URI: schemaResource.URI,
		},
	}

	result, err := schemaResource.Handler(context.Background(), req)
	require.NoError(t, err)
	require.Len(t, result.Contents, 1)

	content := result.Contents[0]
	assert.Equal(t, schemaResource.URI, content.URI)
	assert.Equal(t, "application/json", content.MIMEType)
	assert.NotEmpty(t, content.Text)

	// Validate that it's valid JSON
	var jsonData interface{}
	err = json.Unmarshal([]byte(content.Text), &jsonData)
	assert.NoError(t, err, "Schema should be valid JSON")
}

func TestElementsResource_Integration(t *testing.T) {
	registry := getTestRegistry(t)
	registryAdapter := mcp.NewRegistryAdapter(registry)

	resourceDefs, err := resources.Resources(registryAdapter)
	require.NoError(t, err)

	// Find elements resource
	var elementsResource *resources.ResourceDefinition
	for _, def := range resourceDefs {
		if def.Name == "elements" {
			elementsResource = &def
			break
		}
	}
	require.NotNil(t, elementsResource, "Should have elements resource")

	// Test elements resource handler
	req := &mcpSDK.ReadResourceRequest{
		Params: &mcpSDK.ReadResourceParams{
			URI: elementsResource.URI,
		},
	}

	result, err := elementsResource.Handler(context.Background(), req)
	require.NoError(t, err)
	require.Len(t, result.Contents, 1)

	content := result.Contents[0]
	assert.Equal(t, elementsResource.URI, content.URI)
	assert.Equal(t, "application/json", content.MIMEType)
	assert.NotEmpty(t, content.Text)

	// Parse and validate elements summary
	var elements []resources.ElementSummary
	err = json.Unmarshal([]byte(content.Text), &elements)
	require.NoError(t, err, "Elements should be valid JSON")
	assert.NotEmpty(t, elements, "Should have element summaries")

	// Verify element summaries have expected structure
	for _, element := range elements {
		assert.NotEmpty(t, element.TagName, "Element should have tag name")
		assert.NotEmpty(t, element.Capabilities, "Element should have capabilities")
		// Check for fixture elements
		if element.TagName == "button-element" {
			assert.Contains(t, element.Capabilities, "attributes", "Button should have attributes capability")
			assert.Greater(t, element.AttributeCount, 0, "Button should have attributes")
		}
	}
}

func TestElementResource_Integration(t *testing.T) {
	registry := getTestRegistry(t)
	registryAdapter := mcp.NewRegistryAdapter(registry)

	resourceDefs, err := resources.Resources(registryAdapter)
	require.NoError(t, err)

	// Find element resource
	var elementResource *resources.ResourceDefinition
	for _, def := range resourceDefs {
		if def.Name == "element" {
			elementResource = &def
			break
		}
	}
	require.NotNil(t, elementResource, "Should have element resource")

	tests := []struct {
		name        string
		uri         string
		expectError bool
		checkFunc   func(t *testing.T, content *mcpSDK.ResourceContents)
	}{
		{
			name:        "button-element details",
			uri:         "cem://element/button-element",
			expectError: false,
			checkFunc: func(t *testing.T, content *mcpSDK.ResourceContents) {
				assert.Contains(t, content.Text, "button-element", "Should contain element name")
				assert.Contains(t, content.Text, "variant", "Should contain attributes info")
				assert.Contains(t, content.Text, "A custom button element", "Should contain description")
			},
		},
		{
			name:        "card-element details",
			uri:         "cem://element/card-element",
			expectError: false,
			checkFunc: func(t *testing.T, content *mcpSDK.ResourceContents) {
				assert.Contains(t, content.Text, "card-element", "Should contain element name")
				assert.Contains(t, content.Text, "elevation", "Should contain attributes info")
			},
		},
		{
			name:        "non-existent element",
			uri:         "cem://element/non-existent",
			expectError: false, // Should handle gracefully
			checkFunc: func(t *testing.T, content *mcpSDK.ResourceContents) {
				assert.Contains(t, content.Text, "not found", "Should indicate element not found")
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := &mcpSDK.ReadResourceRequest{
				Params: &mcpSDK.ReadResourceParams{
					URI: test.uri,
				},
			}

			result, err := elementResource.Handler(context.Background(), req)

			if test.expectError {
				assert.Error(t, err, "Expected error for test: %s", test.name)
			} else {
				require.NoError(t, err, "Should not error for test: %s", test.name)
				require.Len(t, result.Contents, 1)

				content := result.Contents[0]
				assert.Equal(t, test.uri, content.URI)
				assert.NotEmpty(t, content.Text)

				if test.checkFunc != nil {
					test.checkFunc(t, content)
				}
			}
		})
	}
}

func TestPackagesResource_Integration(t *testing.T) {
	registry := getTestRegistry(t)
	registryAdapter := mcp.NewRegistryAdapter(registry)

	resourceDefs, err := resources.Resources(registryAdapter)
	require.NoError(t, err)

	// Find packages resource
	var packagesResource *resources.ResourceDefinition
	for _, def := range resourceDefs {
		if def.Name == "packages" {
			packagesResource = &def
			break
		}
	}
	require.NotNil(t, packagesResource, "Should have packages resource")

	// Test packages resource handler
	req := &mcpSDK.ReadResourceRequest{
		Params: &mcpSDK.ReadResourceParams{
			URI: packagesResource.URI,
		},
	}

	result, err := packagesResource.Handler(context.Background(), req)
	require.NoError(t, err)
	require.Len(t, result.Contents, 1)

	content := result.Contents[0]
	assert.Equal(t, packagesResource.URI, content.URI)
	assert.Equal(t, "application/json", content.MIMEType)
	assert.NotEmpty(t, content.Text)

	// Parse and validate packages data
	var packages interface{}
	err = json.Unmarshal([]byte(content.Text), &packages)
	assert.NoError(t, err, "Packages should be valid JSON")
}

func TestGuidelinesResource_Integration(t *testing.T) {
	registry := getTestRegistry(t)
	registryAdapter := mcp.NewRegistryAdapter(registry)

	resourceDefs, err := resources.Resources(registryAdapter)
	require.NoError(t, err)

	// Find guidelines resource
	var guidelinesResource *resources.ResourceDefinition
	for _, def := range resourceDefs {
		if def.Name == "guidelines" {
			guidelinesResource = &def
			break
		}
	}
	require.NotNil(t, guidelinesResource, "Should have guidelines resource")

	// Test guidelines resource handler
	req := &mcpSDK.ReadResourceRequest{
		Params: &mcpSDK.ReadResourceParams{
			URI: guidelinesResource.URI,
		},
	}

	result, err := guidelinesResource.Handler(context.Background(), req)
	require.NoError(t, err)
	require.Len(t, result.Contents, 1)

	content := result.Contents[0]
	assert.Equal(t, guidelinesResource.URI, content.URI)
	assert.Contains(t, content.MIMEType, "text/", "Should be text content")
	assert.NotEmpty(t, content.Text)
}

func TestAccessibilityResource_Integration(t *testing.T) {
	registry := getTestRegistry(t)
	registryAdapter := mcp.NewRegistryAdapter(registry)

	resourceDefs, err := resources.Resources(registryAdapter)
	require.NoError(t, err)

	// Find accessibility resource
	var accessibilityResource *resources.ResourceDefinition
	for _, def := range resourceDefs {
		if def.Name == "accessibility" {
			accessibilityResource = &def
			break
		}
	}
	require.NotNil(t, accessibilityResource, "Should have accessibility resource")

	// Test accessibility resource handler
	req := &mcpSDK.ReadResourceRequest{
		Params: &mcpSDK.ReadResourceParams{
			URI: accessibilityResource.URI,
		},
	}

	result, err := accessibilityResource.Handler(context.Background(), req)
	require.NoError(t, err)
	require.Len(t, result.Contents, 1)

	content := result.Contents[0]
	assert.Equal(t, accessibilityResource.URI, content.URI)
	assert.Contains(t, content.MIMEType, "text/", "Should be text content")
	assert.NotEmpty(t, content.Text)
}

func TestResourceFrontmatterParsing(t *testing.T) {
	// Test that resource definitions are parsed correctly from markdown files
	registry := getTestRegistry(t)
	registryAdapter := mcp.NewRegistryAdapter(registry)

	resourceDefs, err := resources.Resources(registryAdapter)
	require.NoError(t, err)

	tests := []struct {
		name         string
		resourceName string
		expectedURI  string
		expectedMime string
	}{
		{
			name:         "schema resource",
			resourceName: "schema",
			expectedURI:  "cem://schema",
			expectedMime: "application/json",
		},
		{
			name:         "elements resource",
			resourceName: "elements",
			expectedURI:  "cem://elements",
			expectedMime: "application/json",
		},
		{
			name:         "element resource",
			resourceName: "element",
			expectedURI:  "cem://element/{tagName}",
			expectedMime: "text/markdown",
		},
	}

	resourceMap := make(map[string]resources.ResourceDefinition)
	for _, def := range resourceDefs {
		resourceMap[def.Name] = def
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resource, exists := resourceMap[test.resourceName]
			require.True(t, exists, "Resource %s should exist", test.resourceName)

			assert.Equal(t, test.expectedURI, resource.URI, "URI should match for %s", test.resourceName)
			assert.Equal(t, test.expectedMime, resource.MimeType, "MIME type should match for %s", test.resourceName)
			assert.NotEmpty(t, resource.Description, "Should have description for %s", test.resourceName)
		})
	}
}

func TestResourceErrorHandling(t *testing.T) {
	registry := getTestRegistry(t)
	registryAdapter := mcp.NewRegistryAdapter(registry)

	resourceDefs, err := resources.Resources(registryAdapter)
	require.NoError(t, err)

	// Find element resource for testing error conditions
	var elementResource *resources.ResourceDefinition
	for _, def := range resourceDefs {
		if def.Name == "element" {
			elementResource = &def
			break
		}
	}
	require.NotNil(t, elementResource, "Should have element resource")

	tests := []struct {
		name        string
		uri         string
		expectError bool
	}{
		{
			name:        "malformed URI",
			uri:         "invalid-uri",
			expectError: false, // Should handle gracefully
		},
		{
			name:        "empty tag name",
			uri:         "cem://element/",
			expectError: false, // Should handle gracefully
		},
		{
			name:        "special characters in tag name",
			uri:         "cem://element/<script>alert('test')</script>",
			expectError: false, // Should handle gracefully
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := &mcpSDK.ReadResourceRequest{
				Params: &mcpSDK.ReadResourceParams{
					URI: test.uri,
				},
			}

			result, err := elementResource.Handler(context.Background(), req)

			if test.expectError {
				assert.Error(t, err, "Expected error for test: %s", test.name)
			} else {
				// Resource handlers should handle errors gracefully and return informative content
				assert.NoError(t, err, "Should handle errors gracefully for test: %s", test.name)
				if result != nil && len(result.Contents) > 0 {
					assert.NotEmpty(t, result.Contents[0].Text, "Should return meaningful content even for invalid requests")
				}
			}
		})
	}
}

func TestAllResourcesWithFixtures(t *testing.T) {
	// Test that all resources work with real fixture data
	registry := getTestRegistry(t)
	registryAdapter := mcp.NewRegistryAdapter(registry)

	resourceDefs, err := resources.Resources(registryAdapter)
	require.NoError(t, err)

	// Test each resource to ensure it works with fixture data
	for _, resource := range resourceDefs {
		t.Run(resource.Name, func(t *testing.T) {
			// Use appropriate URI for each resource type
			testURI := resource.URI
			if resource.Name == "element" {
				testURI = "cem://element/button-element" // Use known fixture element
			}

			req := &mcpSDK.ReadResourceRequest{
				Params: &mcpSDK.ReadResourceParams{
					URI: testURI,
				},
			}

			result, err := resource.Handler(context.Background(), req)

			assert.NoError(t, err, "Resource %s should handle requests without error", resource.Name)
			assert.NotNil(t, result, "Resource %s should return result", resource.Name)

			if result != nil {
				assert.NotEmpty(t, result.Contents, "Resource %s should return content", resource.Name)
				if len(result.Contents) > 0 {
					assert.NotEmpty(t, result.Contents[0].Text, "Resource %s should return non-empty text", resource.Name)
					assert.Equal(t, testURI, result.Contents[0].URI, "Resource %s should preserve URI", resource.Name)
				}
			}
		})
	}
}
