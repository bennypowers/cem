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
	"embed"

	"bennypowers.dev/cem/mcp/templates"
)

//go:embed templates/*.md
var resourcesTemplateFiles embed.FS

func init() {
	// Register resources templates with the global template pool
	templates.RegisterTemplateSource("resources", &resourcesTemplateFiles)
}

// ManifestContext represents context combining schema definitions with user manifest data
type ManifestContext struct {
	// Schema context - what fields mean according to custom elements spec
	SchemaVersion     string                     // Active schema version
	SchemaDefinitions map[string]interface{}     // Schema definitions for semantic understanding

	// User manifest data - your specific values and patterns
	Overview            string                // Overview of manifest statistics
	ElementCount        int                   // Number of elements in manifests
	CommonPrefixes      []string              // Common element prefixes found
	CSSProperties       []string              // All CSS custom properties
	SchemaVersions      []string              // Manifest schema versions in use
	ElementPatterns     []ElementPattern      // Patterns found across elements
	AttributePatterns   []AttributePattern    // Common attribute usage patterns
	SlotPatterns        []SlotPattern         // Common slot usage patterns
	ExtractedGuidelines []ExtractedGuideline  // Guidelines extracted from descriptions
}

// ElementPattern represents a pattern found across multiple elements
type ElementPattern struct {
	Type        string   // Type of pattern (e.g., "naming", "structure")
	Description string   // Description of the pattern
	Examples    []string // Example elements that follow this pattern
}

// AttributePattern represents common attribute usage
type AttributePattern struct {
	Name        string // Attribute name
	UsageCount  int    // Number of elements using this attribute
	Description string // Description of usage pattern
}

// SlotPattern represents common slot usage
type SlotPattern struct {
	Name        string // Slot name
	UsageCount  int    // Number of elements using this slot
	Description string // Description of usage pattern
}

// ExtractedGuideline represents a guideline extracted from manifest descriptions
type ExtractedGuideline struct {
	Source    string // Source element/attribute (e.g., "my-button" or "my-button.variant")
	Type      string // Type of source ("element" or "attribute")
	Guideline string // The extracted guideline text
}