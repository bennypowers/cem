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
	"sort"
	"strings"

	"bennypowers.dev/cem/manifest"
	"bennypowers.dev/cem/mcp/constants"
	"bennypowers.dev/cem/mcp/types"
	V "bennypowers.dev/cem/validate"
)

// DataSourceProvider creates standardized data sources for path traversal
type DataSourceProvider struct {
	registry types.MCPContext
}

// NewDataSourceProvider creates a new data source provider
func NewDataSourceProvider(registry types.MCPContext) *DataSourceProvider {
	return &DataSourceProvider{
		registry: registry,
	}
}

// CreateDataSources builds standardized data sources for declarative path resolution
func (p *DataSourceProvider) CreateDataSources(args map[string]any) (map[string]any, error) {
	sources := make(map[string]any)

	// Add args as a data source
	sources["args"] = args

	// Create registry data source
	registryData, err := p.createRegistryDataSource()
	if err != nil {
		return nil, fmt.Errorf("failed to create registry data source: %w", err)
	}
	sources["registry"] = registryData

	// Create schema data source
	schemaData, err := p.createSchemaDataSource()
	if err != nil {
		return nil, fmt.Errorf("failed to create schema data source: %w", err)
	}
	sources["schema"] = schemaData

	// Add direct element access if tagName is provided
	if tagName, ok := args["tagName"].(string); ok && tagName != "" {
		elementInfo, err := p.registry.ElementInfo(tagName)
		if err == nil {
			sources["elementInfo"] = elementInfo
		}
	}

	return sources, nil
}

// createRegistryDataSource converts the registry into a path-traversable structure
func (p *DataSourceProvider) createRegistryDataSource() (map[string]any, error) {
	// Get all elements from registry
	allElements := p.registry.AllElements()

	// Convert elements to a map structure for path traversal
	elementsMap := make(map[string]any)
	for tagName, element := range allElements {
		elementData := map[string]any{
			"tagName":     element.TagName(),
			"name":        element.Name(),
			"description": element.Description(),
			"module":      element.Module(),
			"attributes":  convertAttributes(element.Attributes()),
			"slots":       convertSlots(element.Slots()),
			"events":      convertEvents(element.Events()),
			"cssProperties": convertCssProperties(element.CssProperties()),
			"cssParts":    convertCssParts(element.CssParts()),
			"cssStates":   convertCssStates(element.CssStates()),
		}
		elementsMap[tagName] = elementData
	}

	// Build package structure
	packagesData, err := p.buildPackageStructure(allElements)
	if err != nil {
		return nil, fmt.Errorf("failed to build package structure: %w", err)
	}

	// Create summarized elements structure for compatibility
	elementsSummary, summaryMetadata := p.createElementsSummary(allElements)

	return map[string]any{
		"elements":        elementsMap,
		"elementsRaw":     allElements, // Keep original ElementInfo interfaces
		"elementsSummary": elementsSummary,
		"packages":        packagesData,
		"metadata": map[string]any{
			"totalElements":  len(elementsMap),
			"schemaVersions": p.registry.GetManifestSchemaVersions(),
			"summary":        summaryMetadata,
		},
	}, nil
}

// createElementsSummary creates a summarized elements structure with counts and capabilities
func (p *DataSourceProvider) createElementsSummary(allElements map[string]types.ElementInfo) (map[string]any, map[string]any) {
	var elements []map[string]any
	categoryCounts := make(map[string]int)

	for _, element := range allElements {
		// Calculate counts
		attributeCount := len(element.Attributes())
		slotCount := len(element.Slots())
		eventCount := len(element.Events())
		cssPropertyCount := len(element.CssProperties())
		cssPartCount := len(element.CssParts())
		cssStateCount := len(element.CssStates())

		// Determine capabilities based on counts and element properties
		var capabilities []string

		if attributeCount > 0 {
			capabilities = append(capabilities, "configurable")
			categoryCounts["configurable"]++
		}
		if slotCount > 0 {
			capabilities = append(capabilities, "content-slots")
			categoryCounts["content-slots"]++
		}
		if eventCount > 0 {
			capabilities = append(capabilities, "interactive")
			categoryCounts["interactive"]++
		}
		if cssPropertyCount > 0 {
			capabilities = append(capabilities, "themeable")
			categoryCounts["themeable"]++
		}
		if cssPartCount > 0 || cssPropertyCount > 0 {
			capabilities = append(capabilities, "styleable")
			categoryCounts["styleable"]++
		}

		// Check for form element indicators
		for _, attr := range element.Attributes() {
			if attr.Name == "disabled" || attr.Name == "required" || attr.Name == "value" {
				capabilities = append(capabilities, "form-element")
				categoryCounts["form-elements"]++
				break
			}
		}

		packageName := extractPackageFromElementInfo(element)
		if packageName == "" {
			packageName = "default"
		}

		elementSummary := map[string]any{
			"tagName":            element.TagName(),
			"name":               element.Name(),
			"package":            packageName,
			"attributeCount":     attributeCount,
			"slotCount":          slotCount,
			"eventCount":         eventCount,
			"cssPropertyCount":   cssPropertyCount,
			"cssPartCount":       cssPartCount,
			"cssStateCount":      cssStateCount,
			"capabilities":       capabilities,
		}

		elements = append(elements, elementSummary)
	}

	// Initialize all category counts to ensure zero values appear
	allCategories := []string{
		"configurable", "content-slots", "interactive", "themeable",
		"styleable", "form-elements", "layout-elements",
		"navigation-elements", "stateful",
	}
	for _, category := range allCategories {
		if _, exists := categoryCounts[category]; !exists {
			categoryCounts[category] = 0
		}
	}

	summary := map[string]any{
		"elements": elements,
		"metadata": map[string]any{
			"categories":     categoryCounts,
			"totalElements": len(elements),
		},
	}

	metadata := map[string]any{
		"categories":     categoryCounts,
		"totalElements": len(elements),
	}

	return summary, metadata
}

// createSchemaDataSource creates a schema data source with version detection
func (p *DataSourceProvider) createSchemaDataSource() (map[string]any, error) {
	versions := p.registry.GetManifestSchemaVersions()

	defaultVersion := constants.DefaultSchemaVersion
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

	// Parse schema JSON for path traversal
	var schema map[string]any
	if err := json.Unmarshal(schemaData, &schema); err != nil {
		return nil, fmt.Errorf("failed to parse schema JSON: %w", err)
	}

	// Create JSON string for template compatibility
	jsonBytes, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to serialize schema to JSON: %w", err)
	}

	return map[string]any{
		"version":           selectedVersion,
		"availableVersions": versions,
		"definitions":       schema["definitions"],
		"full":             schema,
		"json":             string(jsonBytes),
	}, nil
}

// buildPackageStructure analyzes elements to build package organization
func (p *DataSourceProvider) buildPackageStructure(elementMap map[string]types.ElementInfo) ([]map[string]any, error) {
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

	// Convert to sorted array for stable output
	packageNames := make([]string, 0, len(packageMap))
	for name := range packageMap {
		packageNames = append(packageNames, name)
	}
	sort.Strings(packageNames)

	packages := make([]map[string]any, 0, len(packageMap))
	for _, name := range packageNames {
		pkg := packageMap[name]

		// Sort elements and modules for stable output
		elements := make([]string, len(pkg.Elements))
		copy(elements, pkg.Elements)
		sort.Strings(elements)

		modules := make([]string, len(pkg.Modules))
		copy(modules, pkg.Modules)
		sort.Strings(modules)

		packages = append(packages, map[string]any{
			"name":         name,
			"elements":     elements,
			"modules":      modules,
			"elementCount": len(elements),
			"moduleCount":  len(modules),
		})
	}

	return packages, nil
}

// Helper functions to convert types to map[string]any for JSON path traversal

func convertAttributes(attrs []manifest.Attribute) []map[string]any {
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

func convertSlots(slots []manifest.Slot) []map[string]any {
	result := make([]map[string]any, len(slots))
	for i, slot := range slots {
		result[i] = map[string]any{
			"name":        slot.Name,
			"description": slot.Description,
		}
	}
	return result
}

func convertEvents(events []manifest.Event) []map[string]any {
	result := make([]map[string]any, len(events))
	for i, event := range events {
		result[i] = map[string]any{
			"name":        event.Name,
			"type":        event.Type,
			"description": event.Description,
		}
	}
	return result
}

func convertCssProperties(props []manifest.CssCustomProperty) []map[string]any {
	result := make([]map[string]any, len(props))
	for i, prop := range props {
		result[i] = map[string]any{
			"name":        prop.Name,
			"syntax":      prop.Syntax,
			"description": prop.Description,
			"default":     prop.Default,
		}
	}
	return result
}

func convertCssParts(parts []manifest.CssPart) []map[string]any {
	result := make([]map[string]any, len(parts))
	for i, part := range parts {
		result[i] = map[string]any{
			"name":        part.Name,
			"description": part.Description,
		}
	}
	return result
}

func convertCssStates(states []manifest.CssCustomState) []map[string]any {
	result := make([]map[string]any, len(states))
	for i, state := range states {
		result[i] = map[string]any{
			"name":        state.Name,
			"description": state.Description,
		}
	}
	return result
}

// packageInfo struct for building package organization (reused from existing code)
type packageInfo struct {
	Name     string
	Elements []string
	Modules  []string
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

func selectBestVersion(versions []string) string {
	if len(versions) == 0 {
		return constants.DefaultSchemaVersion
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
		return constants.DefaultSchemaVersion
	}
	return best
}