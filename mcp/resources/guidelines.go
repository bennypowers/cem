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
	"strings"

	"bennypowers.dev/cem/mcp/templates"
	"bennypowers.dev/cem/mcp/types"
	V "bennypowers.dev/cem/validate"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// handleGuidelinesResource provides manifest-derived context and patterns
func handleGuidelinesResource(ctx context.Context, req *mcp.ReadResourceRequest, registry types.MCPContext) (*mcp.ReadResourceResult, error) {
	// Collect context from actual manifest data
	contextData := collectManifestContext(registry)

	// Render using simplified template
	response, err := templates.RenderTemplate("manifest_context", contextData)
	if err != nil {
		return nil, fmt.Errorf("failed to render manifest context: %w", err)
	}

	return &mcp.ReadResourceResult{
		Contents: []*mcp.ResourceContents{{
			URI:      req.Params.URI,
			MIMEType: "text/markdown",
			Text:     response,
		}},
	}, nil
}

// collectManifestContext gathers context combining schema definitions with user manifest data
func collectManifestContext(registry types.MCPContext) ManifestContext {
	elementMap := registry.AllElements()

	// Convert to slice for processing
	elements := make([]types.ElementInfo, 0, len(elementMap))
	for _, element := range elementMap {
		elements = append(elements, element)
	}

	// Get schema context for semantic understanding
	schemaVersions := registry.GetManifestSchemaVersions()
	var schemaVersion string
	var schemaDefinitions map[string]interface{}

	if len(schemaVersions) > 0 {
		schemaVersion = schemaVersions[0] // Use first/primary schema version
		if schema, err := getSchemaDefinitions(registry); err == nil {
			schemaDefinitions = schema
		}
	}

	return ManifestContext{
		// Schema context
		SchemaVersion:     schemaVersion,
		SchemaDefinitions: schemaDefinitions,

		// User manifest data
		Overview:            generateOverview(elements, registry),
		Elements:            elements,
		ElementCount:        len(elements),
		CommonPrefixes:      registry.CommonPrefixes(),
		CSSProperties:       registry.AllCSSProperties(),
		SchemaVersions:      schemaVersions,
		ElementPatterns:     extractElementPatterns(elements),
		AttributePatterns:   extractAttributePatterns(elements),
		SlotPatterns:        extractSlotPatterns(elements),
		CssPartPatterns:     extractCssPartPatterns(elements),
		CssStatePatterns:    extractCssStatePatterns(elements),
		ExtractedGuidelines: extractAllGuidelines(elements),
	}
}

// generateOverview creates overview from actual manifest data
func generateOverview(elements []types.ElementInfo, registry types.MCPContext) string {
	cssProps := registry.AllCSSProperties()
	prefixes := registry.CommonPrefixes()
	versions := registry.GetManifestSchemaVersions()

	return fmt.Sprintf(`# Component Context from Your Manifests

Your workspace contains **%d custom elements** with real constraints and patterns.

## Your Component Library Statistics
- **%d elements** across **%d common prefixes** (%s)
- **%d CSS custom properties** defined
- **Schema versions**: %s

This context is derived from your actual manifest data, not generic guidelines.`,
		len(elements),
		len(elements),
		len(prefixes),
		strings.Join(prefixes, ", "),
		len(cssProps),
		strings.Join(versions, ", "))
}

// extractElementPatterns finds common patterns across elements
func extractElementPatterns(elements []types.ElementInfo) []ElementPattern {
	var patterns []ElementPattern

	// Group by common prefixes
	prefixGroups := make(map[string][]string)
	for _, element := range elements {
		tagName := element.TagName()
		if parts := strings.Split(tagName, "-"); len(parts) > 1 {
			prefix := parts[0]
			prefixGroups[prefix] = append(prefixGroups[prefix], tagName)
		}
	}

	// Generate patterns from groupings
	for prefix, tagNames := range prefixGroups {
		if len(tagNames) > 1 {
			patterns = append(patterns, ElementPattern{
				Type:        "naming",
				Description: fmt.Sprintf("Elements with '%s-' prefix", prefix),
				Examples:    tagNames,
			})
		}
	}

	return patterns
}

// extractAttributePatterns finds common attribute patterns
func extractAttributePatterns(elements []types.ElementInfo) []AttributePattern {
	attributeFreq := make(map[string]int)

	// Count attribute usage across elements
	for _, element := range elements {
		for _, attr := range element.Attributes() {
			attributeFreq[attr.Name()]++
		}
	}

	var patterns []AttributePattern
	for attrName, count := range attributeFreq {
		if count > 1 {
			patterns = append(patterns, AttributePattern{
				Name:        attrName,
				UsageCount:  count,
				Description: fmt.Sprintf("Used in %d elements", count),
			})
		}
	}

	return patterns
}

// extractSlotPatterns finds common slot patterns
func extractSlotPatterns(elements []types.ElementInfo) []SlotPattern {
	slotFreq := make(map[string]int)

	// Count slot usage across elements
	for _, element := range elements {
		for _, slot := range element.Slots() {
			slotName := slot.Name()
			if slotName == "" {
				slotName = "default"
			}
			slotFreq[slotName]++
		}
	}

	var patterns []SlotPattern
	for slotName, count := range slotFreq {
		if count > 1 {
			patterns = append(patterns, SlotPattern{
				Name:        slotName,
				UsageCount:  count,
				Description: fmt.Sprintf("Used in %d elements", count),
			})
		}
	}

	return patterns
}

// extractCssPartPatterns finds common CSS part patterns
func extractCssPartPatterns(elements []types.ElementInfo) []CssPartPattern {
	partFreq := make(map[string]int)

	// Count CSS part usage across elements
	for _, element := range elements {
		for _, part := range element.CssParts() {
			partFreq[part.Name()]++
		}
	}

	var patterns []CssPartPattern
	for partName, count := range partFreq {
		if count > 1 {
			patterns = append(patterns, CssPartPattern{
				Name:        partName,
				UsageCount:  count,
				Description: fmt.Sprintf("Used in %d elements", count),
			})
		}
	}

	return patterns
}

// extractCssStatePatterns finds common CSS state patterns
func extractCssStatePatterns(elements []types.ElementInfo) []CssStatePattern {
	stateFreq := make(map[string]int)

	// Count CSS state usage across elements
	for _, element := range elements {
		for _, state := range element.CssStates() {
			stateFreq[state.Name()]++
		}
	}

	var patterns []CssStatePattern
	for stateName, count := range stateFreq {
		if count > 1 {
			patterns = append(patterns, CssStatePattern{
				Name:        stateName,
				UsageCount:  count,
				Description: fmt.Sprintf("Used in %d elements", count),
			})
		}
	}

	return patterns
}

// extractAllGuidelines extracts RFC 2119 guidelines from all descriptions
func extractAllGuidelines(elements []types.ElementInfo) []ExtractedGuideline {
	var guidelines []ExtractedGuideline

	for _, element := range elements {
		// Extract from element description
		elementGuidelines := extractGuidelinesFromText(element.Description())
		for _, guideline := range elementGuidelines {
			guidelines = append(guidelines, ExtractedGuideline{
				Source:    element.TagName(),
				Type:      "element",
				Guideline: guideline,
			})
		}

		// Extract from attribute descriptions
		for _, attr := range element.Attributes() {
			attrGuidelines := extractGuidelinesFromText(attr.Description())
			for _, guideline := range attrGuidelines {
				guidelines = append(guidelines, ExtractedGuideline{
					Source:    element.TagName() + "." + attr.Name(),
					Type:      "attribute",
					Guideline: guideline,
				})
			}
		}
	}

	return guidelines
}

// extractGuidelinesFromText extracts RFC 2119 keywords from text
func extractGuidelinesFromText(text string) []string {
	var guidelines []string
	if text == "" {
		return guidelines
	}

	sentences := strings.Split(text, ".")
	for _, sentence := range sentences {
		sentence = strings.TrimSpace(sentence)
		lower := strings.ToLower(sentence)
		if strings.Contains(lower, "should") ||
			strings.Contains(lower, "must") ||
			strings.Contains(lower, "use") ||
			strings.Contains(lower, "avoid") {
			if sentence != "" {
				guidelines = append(guidelines, sentence+".")
			}
		}
	}

	return guidelines
}

// getSchemaDefinitions loads the complete schema for template functions
func getSchemaDefinitions(registry types.MCPContext) (map[string]interface{}, error) {
	// Get schema versions from manifests
	versions := registry.GetManifestSchemaVersions()
	if len(versions) == 0 {
		return make(map[string]interface{}), nil
	}

	// Use the first version to load schema
	schemaVersion := versions[0]
	schemaData, err := V.GetSchema(schemaVersion)
	if err != nil {
		return nil, fmt.Errorf("failed to load schema %s: %w", schemaVersion, err)
	}

	// Parse schema JSON - return the complete schema for template functions
	var schema map[string]interface{}
	if err := json.Unmarshal(schemaData, &schema); err != nil {
		return nil, fmt.Errorf("failed to parse schema JSON: %w", err)
	}

	// Return complete schema so template functions can navigate properly
	return schema, nil
}