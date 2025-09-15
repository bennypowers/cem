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
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"bennypowers.dev/cem/mcp/templates"
	"bennypowers.dev/cem/mcp/types"
	V "bennypowers.dev/cem/validate"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/tidwall/gjson"
)

// DeclarativeResourceConfig holds the configuration for a declarative resource
type DeclarativeResourceConfig struct {
	URI          string
	Name         string
	MimeType     string
	URITemplate  bool
	DataFetchers []types.DataFetcher
	Template     string
	ResponseType string
}

// FetchedData represents the result of a data fetcher (ported from tools)
type FetchedData map[string]interface{}

// BaseTemplateData provides common template data fields (ported from tools)
type BaseTemplateData struct {
	Element           types.ElementInfo
	Context           string
	Options           map[string]string
	SchemaDefinitions interface{}
}

// NewBaseTemplateDataWithSchema creates base template data with schema context
func NewBaseTemplateDataWithSchema(element types.ElementInfo, context string, options map[string]string, schemaDefinitions interface{}) BaseTemplateData {
	return BaseTemplateData{
		Element:           element,
		Context:           context,
		Options:           options,
		SchemaDefinitions: schemaDefinitions,
	}
}

// handleDeclarativeResource is the generic handler for declarative resources
func handleDeclarativeResource(
	ctx context.Context,
	req *mcp.ReadResourceRequest,
	registry types.MCPContext,
	config DeclarativeResourceConfig,
) (*mcp.ReadResourceResult, error) {
	// Step 1: Parse URI parameters and extract arguments
	args, err := parseResourceURI(req.Params.URI, config.URI, config.URITemplate)
	if err != nil {
		return nil, fmt.Errorf("failed to parse resource URI: %w", err)
	}

	// Add resource context
	args["context"] = "resource-access"

	// Step 2: Execute data fetchers to collect required data
	fetchedData, err := executeDataFetchers(config.DataFetchers, registry, args)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %w", err)
	}

	// Step 3: Prepare template data by combining args and fetched data
	templateData, err := prepareTemplateData(args, fetchedData, registry)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare template data: %w", err)
	}

	// Step 4: Check if this should return JSON directly
	var response string
	if config.ResponseType == "json" {
		// Return JSON directly from the first data fetcher
		for _, data := range fetchedData {
			jsonBytes, err := json.MarshalIndent(data, "", "  ")
			if err != nil {
				return nil, fmt.Errorf("failed to marshal JSON response: %w", err)
			}
			response = string(jsonBytes)
			break // Use first fetcher result for JSON response
		}
	} else {
		// Step 5: Render the template with the prepared data
		templateName := config.Template
		if templateName == "" {
			templateName = config.Name
		}

		var err error
		response, err = templates.RenderTemplate(templateName, templateData)
		if err != nil {
			return nil, fmt.Errorf("failed to render template '%s': %w", templateName, err)
		}
	}

	// Step 6: Apply any sub-resource filtering if needed
	filteredResponse, err := applySubResourceFiltering(response, args, fetchedData)
	if err != nil {
		return nil, fmt.Errorf("failed to apply sub-resource filtering: %w", err)
	}

	// Step 7: Return the response as a resource
	return &mcp.ReadResourceResult{
		Contents: []*mcp.ResourceContents{{
			URI:      req.Params.URI,
			MIMEType: config.MimeType,
			Text:     filteredResponse,
		}},
	}, nil
}

// executeDataFetchers executes multiple data fetchers and returns combined results (ported from tools)
func executeDataFetchers(fetchers []types.DataFetcher, registry types.MCPContext, args map[string]interface{}) (FetchedData, error) {
	data := make(FetchedData)

	for _, fetcher := range fetchers {
		result, err := executeSingleFetcher(fetcher, registry, args)
		if err != nil {
			if fetcher.Required {
				return nil, fmt.Errorf("required data fetcher '%s' failed: %w", fetcher.Name, err)
			}
			// Skip non-required failed fetchers
			continue
		}
		data[fetcher.Name] = result
	}

	return data, nil
}

// executeSingleFetcher executes a single data fetcher (ported from tools)
func executeSingleFetcher(fetcher types.DataFetcher, registry types.MCPContext, args map[string]interface{}) (interface{}, error) {
	switch fetcher.Type {
	case "manifest_element":
		return fetchManifestElement(registry, args)
	case "schema_definitions":
		return fetchSchemaDefinitions(registry, fetcher.Path)
	case "attribute_collection":
		return fetchAttributeCollection(registry, args)
	case "element_collection":
		return fetchElementCollection(registry, fetcher.Path)
	case "all_elements_with_capabilities":
		return fetchAllElementsWithCapabilities(registry, fetcher.Path)
	case "schema_version_detection":
		return fetchSchemaVersionDetection(registry, fetcher.Path)
	case "package_collection":
		return fetchPackageCollection(registry, fetcher.Path)
	case "guidelines_extraction":
		return fetchGuidelinesExtraction(registry, fetcher.Path)
	case "accessibility_patterns":
		return fetchAccessibilityPatterns(registry, fetcher.Path)
	case "single_element_detailed":
		return fetchSingleElementDetailed(registry, args)
	case "manifest_schema":
		return fetchManifestSchema(registry, fetcher.Path)
	case "package_overview":
		return fetchPackageOverview(registry, fetcher.Path)
	case "workspace_manifest_overview":
		return fetchWorkspaceManifestOverview(registry, fetcher.Path)
	default:
		return nil, fmt.Errorf("unknown data fetcher type: %s", fetcher.Type)
	}
}

// fetchManifestElement fetches element information from the registry (ported from tools)
func fetchManifestElement(registry types.MCPContext, args map[string]interface{}) (interface{}, error) {
	tagName, ok := args["tagName"].(string)
	if !ok || tagName == "" {
		return nil, fmt.Errorf("tagName is required for manifest_element fetcher")
	}

	return registry.ElementInfo(tagName)
}

// fetchSchemaDefinitions fetches schema definitions with gjson path support (ported from tools)
func fetchSchemaDefinitions(registry types.MCPContext, path string) (interface{}, error) {
	// Get schema versions from manifests
	versions := registry.GetManifestSchemaVersions()
	if len(versions) == 0 {
		return make(map[string]interface{}), nil
	}

	// Use the first version to load schema
	schemaVersion := versions[0]
	schemaData, err := getSchemaData(schemaVersion)
	if err != nil {
		return nil, fmt.Errorf("failed to load schema %s: %w", schemaVersion, err)
	}

	// If no path specified, return the whole schema
	if path == "" {
		var schema map[string]interface{}
		if err := json.Unmarshal(schemaData, &schema); err != nil {
			return nil, fmt.Errorf("failed to parse schema JSON: %w", err)
		}
		return schema, nil
	}

	// Use gjson to extract the specific path
	result := gjson.GetBytes(schemaData, path)
	if !result.Exists() {
		return nil, fmt.Errorf("path '%s' not found in schema", path)
	}

	var extracted interface{}
	if err := json.Unmarshal([]byte(result.Raw), &extracted); err != nil {
		return nil, fmt.Errorf("failed to parse extracted schema data: %w", err)
	}

	return extracted, nil
}

// fetchAttributeCollection fetches attribute information (ported from tools)
func fetchAttributeCollection(registry types.MCPContext, args map[string]interface{}) (interface{}, error) {
	tagName, ok := args["tagName"].(string)
	if !ok || tagName == "" {
		return nil, fmt.Errorf("tagName is required for attribute_collection fetcher")
	}

	element, err := registry.ElementInfo(tagName)
	if err != nil {
		return nil, fmt.Errorf("failed to get element info: %w", err)
	}

	// Return the attributes directly
	return element.Attributes(), nil
}

// fetchElementCollection fetches multiple elements using gjson (ported from tools)
func fetchElementCollection(registry types.MCPContext, path string) (interface{}, error) {
	// Get all elements from registry
	allElements := registry.AllElements()

	// Convert to JSON for gjson processing
	elementsJSON, err := json.Marshal(allElements)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal elements: %w", err)
	}

	// If no path specified, return all elements
	if path == "" {
		return allElements, nil
	}

	// Use gjson to extract the specific path
	result := gjson.GetBytes(elementsJSON, path)
	if !result.Exists() {
		return nil, fmt.Errorf("path '%s' not found in elements", path)
	}

	var extracted interface{}
	if err := json.Unmarshal([]byte(result.Raw), &extracted); err != nil {
		return nil, fmt.Errorf("failed to parse extracted elements data: %w", err)
	}

	return extracted, nil
}

// fetchAllElementsWithCapabilities fetches all elements with capability analysis
func fetchAllElementsWithCapabilities(registry types.MCPContext, path string) (interface{}, error) {
	// Get all elements from registry
	elementMap := registry.AllElements()

	// Convert to element summaries with capabilities (logic from elements.go)
	elements := make([]map[string]interface{}, 0, len(elementMap))

	// Sort element keys for deterministic output
	elementKeys := make([]string, 0, len(elementMap))
	for tagName := range elementMap {
		elementKeys = append(elementKeys, tagName)
	}
	sort.Strings(elementKeys)

	for _, tagName := range elementKeys {
		element := elementMap[tagName]

		// Build capabilities list
		capabilities := buildElementCapabilities(element)

		summary := map[string]interface{}{
			"tagName":          element.TagName(),
			"name":             element.Name(),
			"package":          extractPackageFromElementInfo(element),
			"capabilities":     capabilities,
			"attributeCount":   len(element.Attributes()),
			"slotCount":        len(element.Slots()),
			"eventCount":       len(element.Events()),
			"cssPropertyCount": len(element.CssProperties()),
			"cssPartCount":     len(element.CssParts()),
			"cssStateCount":    len(element.CssStates()),
		}
		elements = append(elements, summary)
	}

	// Build response structure
	result := map[string]interface{}{
		"elements": elements,
		"metadata": map[string]interface{}{
			"totalElements": len(elements),
			"categories":    categorizeElementsByCapabilityList(elements),
		},
	}

	return result, nil
}

// fetchSchemaVersionDetection fetches schema version with detection logic
func fetchSchemaVersionDetection(registry types.MCPContext, path string) (interface{}, error) {
	// Schema detection logic from schema.go
	versions := registry.GetManifestSchemaVersions()

	defaultVersion := "2.1.1-speculative"
	selectedVersion := defaultVersion

	if len(versions) == 1 {
		selectedVersion = versions[0]
	} else if len(versions) > 1 {
		selectedVersion = selectBestVersion(versions)
	}

	// Get the actual schema data
	schemaData, err := V.GetSchema(selectedVersion)
	if err != nil {
		return nil, fmt.Errorf("failed to load schema: %w", err)
	}

	return map[string]interface{}{
		"version":           selectedVersion,
		"availableVersions": versions,
		"schemaData":        string(schemaData),
	}, nil
}

// fetchPackageCollection fetches package information
func fetchPackageCollection(registry types.MCPContext, path string) (interface{}, error) {
	// Package collection logic from packages.go - analyze elements to build package structure
	elementMap := registry.AllElements()

	// Build package map from elements
	packageMap := make(map[string]*packageInfo)

	for _, element := range elementMap {
		packageName := extractPackageFromElementInfo(element)
		if packageName == "" {
			packageName = "default"
		}

		if _, exists := packageMap[packageName]; !exists {
			packageMap[packageName] = &packageInfo{
				Name:     packageName,
				Elements: []string{},
				Modules:  []string{},
			}
		}

		pkg := packageMap[packageName]
		pkg.Elements = append(pkg.Elements, element.TagName())

		if module := element.Module(); module != "" {
			moduleExists := false
			for _, existingModule := range pkg.Modules {
				if existingModule == module {
					moduleExists = true
					break
				}
			}
			if !moduleExists {
				pkg.Modules = append(pkg.Modules, module)
			}
		}
	}

	// Convert to response format with stable sorting
	packageNames := make([]string, 0, len(packageMap))
	for name := range packageMap {
		packageNames = append(packageNames, name)
	}
	sort.Strings(packageNames)

	packages := make([]map[string]interface{}, 0, len(packageMap))
	for _, name := range packageNames {
		pkg := packageMap[name]

		// Sort elements and modules for stable output
		elements := make([]string, len(pkg.Elements))
		copy(elements, pkg.Elements)
		sort.Strings(elements)

		modules := make([]string, len(pkg.Modules))
		copy(modules, pkg.Modules)
		sort.Strings(modules)

		packages = append(packages, map[string]interface{}{
			"name":         name,
			"elements":     elements,
			"modules":      modules,
			"elementCount": len(elements),
			"moduleCount":  len(modules),
		})
	}

	result := map[string]interface{}{
		"packages": packages,
		"metadata": map[string]interface{}{
			"totalPackages": len(packages),
		},
	}

	return result, nil
}

// fetchGuidelinesExtraction fetches extracted guidelines
func fetchGuidelinesExtraction(registry types.MCPContext, path string) (interface{}, error) {
	// Guidelines extraction logic from guidelines.go
	guidelines := extractGuidelinesFromManifests(registry)

	result := map[string]interface{}{
		"guidelines": guidelines,
		"metadata": map[string]interface{}{
			"totalGuidelines": len(guidelines),
		},
	}

	return result, nil
}

// fetchAccessibilityPatterns fetches accessibility patterns
func fetchAccessibilityPatterns(registry types.MCPContext, path string) (interface{}, error) {
	// Accessibility patterns logic from accessibility.go
	patterns := extractAccessibilityPatterns(registry)

	result := map[string]interface{}{
		"patterns": patterns,
		"metadata": map[string]interface{}{
			"totalPatterns": len(patterns),
		},
	}

	return result, nil
}

// fetchSingleElementDetailed fetches detailed element information
func fetchSingleElementDetailed(registry types.MCPContext, args map[string]interface{}) (interface{}, error) {
	tagName, ok := args["tagName"].(string)
	if !ok || tagName == "" {
		return nil, fmt.Errorf("tagName is required for single_element_detailed fetcher")
	}

	element, err := registry.ElementInfo(tagName)
	if err != nil {
		return nil, fmt.Errorf("failed to get element info: %w", err)
	}

	// Detailed element with examples and usage patterns (logic from element.go)
	result := map[string]interface{}{
		"element":     element,
		"examples":    generateElementExamples(element),
		"usageNotes":  generateUsageNotes(element),
		"integration": generateIntegrationNotes(element),
	}

	return result, nil
}

// Helper types and functions for the new data fetchers

type packageInfo struct {
	Name     string
	Elements []string
	Modules  []string
}

func buildElementCapabilities(element types.ElementInfo) []string {
	var capabilities []string

	if len(element.Attributes()) > 0 {
		capabilities = append(capabilities, "configurable")
	}
	if len(element.Slots()) > 0 {
		capabilities = append(capabilities, "content-slots")
	}
	if len(element.Events()) > 0 {
		capabilities = append(capabilities, "interactive")
	}
	if len(element.CssProperties()) > 0 {
		capabilities = append(capabilities, "themeable")
	}
	if len(element.CssParts()) > 0 {
		capabilities = append(capabilities, "styleable")
	}
	if len(element.CssStates()) > 0 {
		capabilities = append(capabilities, "stateful")
	}

	// Add semantic capabilities based on tag name patterns
	tagName := element.TagName()
	if containsAnySubstring(tagName, []string{"button", "input", "field", "form"}) {
		capabilities = append(capabilities, "form-element")
	}
	if containsAnySubstring(tagName, []string{"grid", "layout", "container", "flex"}) {
		capabilities = append(capabilities, "layout-element")
	}
	if containsAnySubstring(tagName, []string{"nav", "menu", "tab", "breadcrumb"}) {
		capabilities = append(capabilities, "navigation-element")
	}

	return capabilities
}

func extractPackageFromElementInfo(element types.ElementInfo) string {
	module := element.Module()
	if module != "" {
		// Extract package name from module path
		parts := strings.Split(module, "/")
		if len(parts) > 0 {
			return parts[0]
		}
		return module
	}

	// Fallback: extract package from tag name (e.g., "button-element" -> "button")
	tagName := element.TagName()
	if strings.Contains(tagName, "-") {
		parts := strings.Split(tagName, "-")
		if len(parts) > 1 {
			return parts[0]
		}
	}

	return ""
}

func categorizeElementsByCapabilityList(elements []map[string]interface{}) map[string]int {
	categories := map[string]int{
		"configurable":        0,
		"content-slots":       0,
		"interactive":         0,
		"themeable":           0,
		"styleable":           0,
		"stateful":            0,
		"form-elements":       0,
		"layout-elements":     0,
		"navigation-elements": 0,
	}

	for _, element := range elements {
		capabilities, ok := element["capabilities"].([]string)
		if !ok {
			continue
		}

		for _, capability := range capabilities {
			if capability == "form-element" {
				categories["form-elements"]++
			} else if capability == "layout-element" {
				categories["layout-elements"]++
			} else if capability == "navigation-element" {
				categories["navigation-elements"]++
			} else if _, exists := categories[capability]; exists {
				categories[capability]++
			}
		}
	}

	return categories
}

func selectBestVersion(versions []string) string {
	if len(versions) == 0 {
		return "2.1.1-speculative"
	}

	var best string
	var hasSpeculative bool

	for _, version := range versions {
		if version == "" {
			continue
		}

		// Always prefer speculative versions as they are most complete
		if strings.Contains(version, "speculative") {
			if !hasSpeculative || version > best {
				best = version
				hasSpeculative = true
			}
		} else if !hasSpeculative {
			// Among non-speculative versions, take the highest
			if best == "" || version > best {
				best = version
			}
		}
	}

	if best == "" {
		return "2.1.1-speculative"
	}
	return best
}

func extractGuidelinesFromManifests(registry types.MCPContext) []map[string]interface{} {
	// Extract guidelines from element descriptions
	var guidelines []map[string]interface{}

	elementMap := registry.AllElements()
	for _, element := range elementMap {
		description := element.Description()
		if description != "" {
			guidelines = append(guidelines, map[string]interface{}{
				"source":    element.TagName(),
				"type":      "element",
				"guideline": description,
			})
		}

		// Extract from attributes
		for _, attr := range element.Attributes() {
			if attr.Description != "" {
				guidelines = append(guidelines, map[string]interface{}{
					"source":    element.TagName() + "." + attr.Name,
					"type":      "attribute",
					"guideline": attr.Description,
				})
			}
		}
	}

	return guidelines
}

func extractAccessibilityPatterns(registry types.MCPContext) []map[string]interface{} {
	// Extract accessibility patterns from manifests
	var patterns []map[string]interface{}

	elementMap := registry.AllElements()
	for _, element := range elementMap {
		// Look for accessibility-related information
		if containsAnySubstring(element.Description(), []string{"aria", "accessibility", "screen reader", "keyboard"}) {
			patterns = append(patterns, map[string]interface{}{
				"element":     element.TagName(),
				"pattern":     "accessibility-aware",
				"description": element.Description(),
			})
		}
	}

	return patterns
}

func generateElementExamples(element types.ElementInfo) []map[string]interface{} {
	// Generate usage examples for the element
	var examples []map[string]interface{}

	// Basic example
	examples = append(examples, map[string]interface{}{
		"title":       "Basic Usage",
		"description": "Standard implementation of " + element.TagName(),
		"html":        "<" + element.TagName() + "></" + element.TagName() + ">",
	})

	return examples
}

func generateUsageNotes(element types.ElementInfo) []string {
	var notes []string

	if len(element.Attributes()) > 0 {
		notes = append(notes, "This element supports "+fmt.Sprintf("%d", len(element.Attributes()))+" configurable attributes")
	}
	if len(element.Slots()) > 0 {
		notes = append(notes, "Content can be placed in "+fmt.Sprintf("%d", len(element.Slots()))+" available slots")
	}
	if len(element.Events()) > 0 {
		notes = append(notes, "Listen for "+fmt.Sprintf("%d", len(element.Events()))+" custom events for interactivity")
	}

	return notes
}

func generateIntegrationNotes(element types.ElementInfo) map[string]interface{} {
	return map[string]interface{}{
		"package": extractPackageFromElementInfo(element),
		"module":  element.Module(),
		"styling": len(element.CssProperties()) > 0 || len(element.CssParts()) > 0,
	}
}

func containsAnySubstring(s string, substrings []string) bool {
	for _, substr := range substrings {
		if strings.Contains(s, substr) {
			return true
		}
	}
	return false
}

// getSchemaData retrieves schema data using the existing schema loading mechanism (ported from tools)
func getSchemaData(version string) ([]byte, error) {
	// Use the existing V.GetSchema function from the validate package
	return V.GetSchema(version)
}

// prepareTemplateData combines arguments and fetched data for template rendering (ported from tools)
func prepareTemplateData(args map[string]interface{}, fetchedData FetchedData, registry types.MCPContext) (interface{}, error) {
	// Create the base template data structure
	templateData := BaseTemplateData{
		Context: getStringArg(args, "context", ""),
		Options: make(map[string]string),
	}

	// Add element data if available
	if element, ok := fetchedData["element"]; ok {
		if elementInfo, ok := element.(types.ElementInfo); ok {
			templateData.Element = elementInfo
		}
	}

	// Add schema definitions if available
	if schema, ok := fetchedData["schema"]; ok {
		templateData.SchemaDefinitions = schema
	}

	// Create a combined data structure that includes both base template data
	// and individual fetched data for template access
	combinedData := struct {
		BaseTemplateData
		FetchedData FetchedData
		Args        map[string]interface{}
	}{
		BaseTemplateData: templateData,
		FetchedData:      fetchedData,
		Args:             args,
	}

	// If we have element data, use the existing pattern for compatibility
	if templateData.Element != nil {
		return NewBaseTemplateDataWithSchema(
			templateData.Element,
			templateData.Context,
			templateData.Options,
			templateData.SchemaDefinitions,
		), nil
	}

	return combinedData, nil
}

// getStringArg safely extracts a string argument with a default value (ported from tools)
func getStringArg(args map[string]interface{}, key, defaultValue string) string {
	if value, ok := args[key]; ok {
		if str, ok := value.(string); ok {
			return str
		}
	}
	return defaultValue
}

// parseResourceURI extracts parameters from the URI based on the template
func parseResourceURI(actualURI, templateURI string, isTemplate bool) (map[string]interface{}, error) {
	args := make(map[string]interface{})

	if !isTemplate {
		// Simple URI, no parameters to extract
		return args, nil
	}

	// Extract parameters from URI template
	// Expected patterns:
	// cem://element/{tagName}/attributes -> extracts tagName
	// cem://element/{tagName}/attributes/{name} -> extracts tagName and name
	// cem://element/{tagName}/css/parts/{name} -> extracts tagName and name

	// Remove scheme and split into parts
	actualParts := strings.Split(strings.TrimPrefix(actualURI, "cem://"), "/")
	templateParts := strings.Split(strings.TrimPrefix(templateURI, "cem://"), "/")

	// Match template parts with actual parts
	for i, templatePart := range templateParts {
		if i >= len(actualParts) {
			continue
		}

		actualPart := actualParts[i]

		// Check if this is a parameter (surrounded by {})
		if strings.HasPrefix(templatePart, "{") && strings.HasSuffix(templatePart, "}") {
			paramName := strings.Trim(templatePart, "{}")
			args[paramName] = actualPart
		}
	}

	// Handle sub-resource access (optional parameters beyond template)
	if len(actualParts) > len(templateParts) {
		// Additional path segments indicate sub-resource access
		extraParts := actualParts[len(templateParts):]
		if len(extraParts) > 0 {
			// Store the sub-resource name for filtering
			args["subResource"] = strings.Join(extraParts, "/")
		}
	}

	return args, nil
}

// applySubResourceFiltering filters the response for sub-resource access
func applySubResourceFiltering(response string, args map[string]interface{}, fetchedData FetchedData) (string, error) {
	// Check if this is a sub-resource request
	subResource, hasSubResource := args["subResource"].(string)
	if !hasSubResource || subResource == "" {
		return response, nil
	}

	// Apply filtering based on the sub-resource type and name
	// This would be expanded based on specific filtering needs
	// For now, return the full response (filtering can be implemented per resource type)
	return response, nil
}

// makeDeclarativeResourceHandler creates a handler function for a declarative resource
func makeDeclarativeResourceHandler(registry types.MCPContext, config DeclarativeResourceConfig) mcp.ResourceHandler {
	return func(ctx context.Context, req *mcp.ReadResourceRequest) (*mcp.ReadResourceResult, error) {
		return handleDeclarativeResource(ctx, req, registry, config)
	}
}

// MakeDeclarativeResourceHandler is the exported version for testing
func MakeDeclarativeResourceHandler(registry types.MCPContext, config DeclarativeResourceConfig) mcp.ResourceHandler {
	return makeDeclarativeResourceHandler(registry, config)
}

// fetchManifestSchema fetches the manifest schema as JSON string
func fetchManifestSchema(registry types.MCPContext, path string) (interface{}, error) {
	schemaData, err := fetchSchemaDefinitions(registry, path)
	if err != nil {
		return nil, err
	}

	// Convert to pretty-printed JSON string for template output
	jsonBytes, err := json.MarshalIndent(schemaData, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to serialize schema to JSON: %w", err)
	}

	return string(jsonBytes), nil
}

// fetchPackageOverview fetches package overview information
func fetchPackageOverview(registry types.MCPContext, path string) (interface{}, error) {
	return fetchPackageCollection(registry, path)
}

// fetchWorkspaceManifestOverview fetches workspace manifest overview
func fetchWorkspaceManifestOverview(registry types.MCPContext, path string) (interface{}, error) {
	return fetchPackageCollection(registry, path)
}
