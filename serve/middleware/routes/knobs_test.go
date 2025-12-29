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

package routes

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	M "bennypowers.dev/cem/manifest"
	"github.com/google/go-cmp/cmp"
)

// testTemplatesForKnobs creates a template registry for testing (with nil context)
func testTemplatesForKnobs() *TemplateRegistry {
	return NewTemplateRegistry(nil)
}

func TestGenerateKnobs_SimpleButton(t *testing.T) {
	// Load manifest fixture
	manifestPath := filepath.Join("testdata", "knobs", "simple-button-manifest.json")
	manifestBytes, err := os.ReadFile(manifestPath)
	if err != nil {
		t.Fatalf("Failed to read manifest fixture: %v", err)
	}

	var pkg M.Package
	err = json.Unmarshal(manifestBytes, &pkg)
	if err != nil {
		t.Fatalf("Failed to parse manifest: %v", err)
	}

	// Get the first custom element declaration
	var declaration *M.CustomElementDeclaration
	for _, mod := range pkg.Modules {
		for _, decl := range mod.Declarations {
			if customEl, ok := decl.(*M.CustomElementDeclaration); ok {
				declaration = customEl
				break
			}
		}
		if declaration != nil {
			break
		}
	}

	if declaration == nil {
		t.Fatal("No custom element declaration found in manifest")
	}

	// Load demo HTML
	demoPath := filepath.Join("testdata", "knobs", "simple-button-demo.html")
	demoHTML, err := os.ReadFile(demoPath)
	if err != nil {
		t.Fatalf("Failed to read demo HTML: %v", err)
	}

	// Generate knobs with all categories enabled
	knobs, err := GenerateKnobs(declaration, demoHTML, "attributes properties css-properties")
	if err != nil {
		t.Fatalf("GenerateKnobs failed: %v", err)
	}

	// Verify tag name
	if knobs.TagName != "my-button" {
		t.Errorf("Expected TagName = 'my-button', got '%s'", knobs.TagName)
	}

	// Verify attribute knobs were generated (variant and size should be attributes only)
	if len(knobs.AttributeKnobs) == 0 {
		t.Error("Expected attribute knobs to be generated, got 0")
	}

	// Verify we have variant attribute (enum)
	var variantKnob *KnobData
	for i, knob := range knobs.AttributeKnobs {
		if knob.Name == "variant" {
			variantKnob = &knobs.AttributeKnobs[i]
			break
		}
	}
	if variantKnob == nil {
		t.Error("Expected 'variant' attribute knob, not found")
	} else {
		if variantKnob.Type != KnobTypeEnum {
			t.Errorf("Expected variant knob type = enum, got %s", variantKnob.Type)
		}
		expectedEnumValues := []string{"primary", "secondary", "danger"}
		if diff := cmp.Diff(expectedEnumValues, variantKnob.EnumValues); diff != "" {
			t.Errorf("Variant enum values mismatch (-want +got):\n%s", diff)
		}
		// Current value should be "primary" from demo HTML
		if variantKnob.CurrentValue != "primary" {
			t.Errorf("Expected variant current value = 'primary', got '%s'", variantKnob.CurrentValue)
		}
	}

	// Verify property knobs
	if len(knobs.PropertyKnobs) == 0 {
		t.Error("Expected property knobs to be generated, got 0")
	}

	// Verify we have disabled property (not attribute, due to deduplication)
	var disabledKnob *KnobData
	for i, knob := range knobs.PropertyKnobs {
		if knob.Name == "disabled" {
			disabledKnob = &knobs.PropertyKnobs[i]
			break
		}
	}
	if disabledKnob == nil {
		t.Error("Expected 'disabled' property knob, not found")
	} else {
		if disabledKnob.Type != KnobTypeBoolean {
			t.Errorf("Expected disabled knob type = boolean, got %s", disabledKnob.Type)
		}
	}

	// Verify CSS property knobs
	if len(knobs.CSSPropertyKnobs) == 0 {
		t.Error("Expected CSS property knobs to be generated, got 0")
	}

	// Check for --button-color (should be color type)
	var colorKnob *KnobData
	for i, knob := range knobs.CSSPropertyKnobs {
		if knob.Name == "--button-color" {
			colorKnob = &knobs.CSSPropertyKnobs[i]
			break
		}
	}
	if colorKnob == nil {
		t.Error("Expected '--button-color' CSS property knob, not found")
	} else {
		if colorKnob.Type != KnobTypeColor {
			t.Errorf("Expected --button-color knob type = color, got %s", colorKnob.Type)
		}
	}
}

func TestParseEnabledKnobs(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		wantAttributes bool
		wantProperties bool
		wantCSS        bool
	}{
		{
			name:           "all enabled",
			input:          "attributes properties css-properties",
			wantAttributes: true,
			wantProperties: true,
			wantCSS:        true,
		},
		{
			name:           "only attributes",
			input:          "attributes",
			wantAttributes: true,
			wantProperties: false,
			wantCSS:        false,
		},
		{
			name:           "attributes and properties",
			input:          "attributes properties",
			wantAttributes: true,
			wantProperties: true,
			wantCSS:        false,
		},
		{
			name:           "empty",
			input:          "",
			wantAttributes: false,
			wantProperties: false,
			wantCSS:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseEnabledKnobs(tt.input)

			if got[KnobCategoryAttribute] != tt.wantAttributes {
				t.Errorf("attributes: got %v, want %v", got[KnobCategoryAttribute], tt.wantAttributes)
			}
			if got[KnobCategoryProperty] != tt.wantProperties {
				t.Errorf("properties: got %v, want %v", got[KnobCategoryProperty], tt.wantProperties)
			}
			if got[KnobCategoryCSSProperty] != tt.wantCSS {
				t.Errorf("css-properties: got %v, want %v", got[KnobCategoryCSSProperty], tt.wantCSS)
			}
		})
	}
}

func TestParseType(t *testing.T) {
	tests := []struct {
		name           string
		typeText       string
		wantType       KnobType
		wantEnumValues []string
	}{
		{
			name:           "boolean",
			typeText:       "boolean",
			wantType:       KnobTypeBoolean,
			wantEnumValues: nil,
		},
		{
			name:           "number",
			typeText:       "number",
			wantType:       KnobTypeNumber,
			wantEnumValues: nil,
		},
		{
			name:           "string",
			typeText:       "string",
			wantType:       KnobTypeString,
			wantEnumValues: nil,
		},
		{
			name:           "enum",
			typeText:       "'primary' | 'secondary' | 'danger'",
			wantType:       KnobTypeEnum,
			wantEnumValues: []string{"primary", "secondary", "danger"},
		},
		{
			name:           "enum with spaces",
			typeText:       "'a' | 'b'",
			wantType:       KnobTypeEnum,
			wantEnumValues: []string{"a", "b"},
		},
		{
			name:           "single literal single quotes",
			typeText:       "'bordered'",
			wantType:       KnobTypeEnum,
			wantEnumValues: []string{"bordered"},
		},
		{
			name:           "single literal double quotes",
			typeText:       "\"solid\"",
			wantType:       KnobTypeEnum,
			wantEnumValues: []string{"solid"},
		},
		{
			name:           "single literal with spaces",
			typeText:       " 'compact' ",
			wantType:       KnobTypeEnum,
			wantEnumValues: []string{"compact"},
		},
		{
			name:           "empty",
			typeText:       "",
			wantType:       KnobTypeString,
			wantEnumValues: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotType, gotEnumValues := parseType(tt.typeText)

			if gotType != tt.wantType {
				t.Errorf("type: got %v, want %v", gotType, tt.wantType)
			}

			if diff := cmp.Diff(tt.wantEnumValues, gotEnumValues); diff != "" {
				t.Errorf("enum values mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestRenderKnobsHTML(t *testing.T) {
	// Create sample knobs data
	knobs := &KnobsData{
		TagName:   "my-button",
		ElementID: "my-button-0",
		AttributeKnobs: []KnobData{
			{
				Name:         "disabled",
				Category:     KnobCategoryAttribute,
				Type:         KnobTypeBoolean,
				CurrentValue: "",
				Description:  "Whether the button is disabled",
			},
			{
				Name:         "variant",
				Category:     KnobCategoryAttribute,
				Type:         KnobTypeEnum,
				CurrentValue: "primary",
				EnumValues:   []string{"primary", "secondary", "danger"},
				Description:  "Button variant",
				Default:      "primary",
			},
		},
		PropertyKnobs: []KnobData{
			{
				Name:         "label",
				Category:     KnobCategoryProperty,
				Type:         KnobTypeString,
				CurrentValue: "",
				Description:  "Button label text",
			},
		},
		CSSPropertyKnobs: []KnobData{
			{
				Name:         "--button-color",
				Category:     KnobCategoryCSSProperty,
				Type:         KnobTypeColor,
				CurrentValue: "",
				Description:  "Button background color",
				Default:      "#0066cc",
			},
		},
	}

	// Wrap in ElementKnobGroup slice
	knobGroups := []ElementKnobGroup{
		{
			TagName:       "my-button",
			ElementID:     "my-button-0",
			Label:         "#my-button-0",
			InstanceIndex: 0,
			IsPrimary:     true,
			Knobs:         knobs,
		},
	}

	html, err := RenderKnobsHTML(testTemplatesForKnobs(), knobGroups)
	if err != nil {
		t.Fatalf("RenderKnobsHTML failed: %v", err)
	}

	htmlStr := string(html)

	// Verify structure is present
	if htmlStr == "" {
		t.Fatal("Expected non-empty HTML output - template may have failed to render")
	}

	// Verify sections are present
	expectedStrings := []string{
		"cem-serve-knobs",
		"Attributes",
		"Properties",
		"CSS Custom Properties",
		"disabled",
		"variant",
		"label",
		"--button-color",
	}

	for _, expected := range expectedStrings {
		if !strings.Contains(htmlStr, expected) {
			t.Errorf("Expected HTML to contain '%s'", expected)
		}
	}
}

// TestDiscoverElementInstances tests multi-element discovery
func TestDiscoverElementInstances(t *testing.T) {
	tests := []struct {
		name          string
		demoFile      string
		expectedCount int
		expectedIDs   []string // Expected element IDs in source order
	}{
		{
			name:          "three instances with IDs",
			demoFile:      "multi-element-with-ids-demo.html",
			expectedCount: 3,
			expectedIDs:   []string{"card-1", "card-2", "card-3"},
		},
		{
			name:          "three instances without IDs",
			demoFile:      "multi-element-fallback-demo.html",
			expectedCount: 3,
			expectedIDs:   []string{"", "", ""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Load demo HTML
			demoPath := filepath.Join("testdata", "knobs", tt.demoFile)
			demoHTML, err := os.ReadFile(demoPath)
			if err != nil {
				t.Fatalf("Failed to read demo HTML: %v", err)
			}

			// Discover instances
			instances, err := discoverElementInstances("my-card", demoHTML)
			if err != nil {
				t.Fatalf("discoverElementInstances failed: %v", err)
			}

			// Verify count
			if len(instances) != tt.expectedCount {
				t.Errorf("Expected %d instances, got %d", tt.expectedCount, len(instances))
			}

			// Verify IDs match (in source order)
			for i, expectedID := range tt.expectedIDs {
				if i >= len(instances) {
					break
				}
				if instances[i].ID != expectedID {
					t.Errorf("Instance %d: expected ID '%s', got '%s'", i, expectedID, instances[i].ID)
				}
			}
		})
	}
}

// TestGenerateElementLabel tests the label generation priority system
func TestGenerateElementLabel(t *testing.T) {
	tests := []struct {
		name          string
		demoFile      string
		instanceIndex int
		expectedLabel string
	}{
		{
			name:          "label from ID",
			demoFile:      "multi-element-with-ids-demo.html",
			instanceIndex: 0,
			expectedLabel: "#card-1",
		},
		{
			name:          "label from text content",
			demoFile:      "multi-element-with-text-demo.html",
			instanceIndex: 0,
			expectedLabel: "User Settings",
		},
		{
			name:          "label from long text (truncated)",
			demoFile:      "multi-element-with-text-demo.html",
			instanceIndex: 1,
			expectedLabel: "This is a very long…",
		},
		{
			name:          "label from aria-label",
			demoFile:      "multi-element-with-aria-demo.html",
			instanceIndex: 0,
			expectedLabel: "Main dashboard card",
		},
		{
			name:          "fallback label",
			demoFile:      "multi-element-fallback-demo.html",
			instanceIndex: 1,
			expectedLabel: "my-card No. 2",
		},
		{
			name:          "label from nested text (h2)",
			demoFile:      "multi-element-nested-demo.html",
			instanceIndex: 0,
			expectedLabel: "User Profile",
		},
		{
			name:          "label from deeply nested text",
			demoFile:      "multi-element-nested-demo.html",
			instanceIndex: 1,
			expectedLabel: "Settings Panel",
		},
		{
			name:          "label from span text",
			demoFile:      "multi-element-nested-demo.html",
			instanceIndex: 2,
			expectedLabel: "Dashboard",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Load demo HTML
			demoPath := filepath.Join("testdata", "knobs", tt.demoFile)
			demoHTML, err := os.ReadFile(demoPath)
			if err != nil {
				t.Fatalf("Failed to read demo HTML: %v", err)
			}

			// Discover instances
			instances, err := discoverElementInstances("my-card", demoHTML)
			if err != nil {
				t.Fatalf("discoverElementInstances failed: %v", err)
			}

			if tt.instanceIndex >= len(instances) {
				t.Fatalf("Instance index %d out of bounds (have %d instances)", tt.instanceIndex, len(instances))
			}

			// Generate label
			label := generateElementLabel(instances[tt.instanceIndex], "my-card", tt.instanceIndex)

			if label != tt.expectedLabel {
				t.Errorf("Expected label '%s', got '%s'", tt.expectedLabel, label)
			}
		})
	}
}

// TestGenerateMultiInstanceKnobs tests knob generation for multiple element instances
func TestGenerateMultiInstanceKnobs(t *testing.T) {
	// Load manifest fixture
	manifestPath := filepath.Join("testdata", "knobs", "multi-element-manifest.json")
	manifestBytes, err := os.ReadFile(manifestPath)
	if err != nil {
		t.Fatalf("Failed to read manifest fixture: %v", err)
	}

	var pkg M.Package
	err = json.Unmarshal(manifestBytes, &pkg)
	if err != nil {
		t.Fatalf("Failed to parse manifest: %v", err)
	}

	// Get the custom element declaration
	var declaration *M.CustomElementDeclaration
	for _, mod := range pkg.Modules {
		for _, decl := range mod.Declarations {
			if customEl, ok := decl.(*M.CustomElementDeclaration); ok {
				declaration = customEl
				break
			}
		}
		if declaration != nil {
			break
		}
	}

	if declaration == nil {
		t.Fatal("No custom element declaration found in manifest")
	}

	// Load demo HTML with multiple instances
	demoPath := filepath.Join("testdata", "knobs", "multi-element-with-ids-demo.html")
	demoHTML, err := os.ReadFile(demoPath)
	if err != nil {
		t.Fatalf("Failed to read demo HTML: %v", err)
	}

	// Generate knobs for all instances
	knobGroups, err := GenerateMultiInstanceKnobs(declaration, demoHTML, "attributes properties css-properties")
	if err != nil {
		t.Fatalf("GenerateMultiInstanceKnobs failed: %v", err)
	}

	// Verify we got 3 knob groups (3 instances in demo)
	if len(knobGroups) != 3 {
		t.Errorf("Expected 3 knob groups, got %d", len(knobGroups))
	}

	// Verify first group is marked as primary
	if len(knobGroups) > 0 && !knobGroups[0].IsPrimary {
		t.Error("Expected first knob group to be marked as primary")
	}

	// Verify other groups are not primary
	for i := 1; i < len(knobGroups); i++ {
		if knobGroups[i].IsPrimary {
			t.Errorf("Expected knob group %d to not be primary", i)
		}
	}

	// Verify labels are generated
	expectedLabels := []string{"#card-1", "#card-2", "#card-3"}
	for i, group := range knobGroups {
		if i >= len(expectedLabels) {
			break
		}
		if group.Label != expectedLabels[i] {
			t.Errorf("Group %d: expected label '%s', got '%s'", i, expectedLabels[i], group.Label)
		}
	}

	// Verify each group has the correct tag name
	for i, group := range knobGroups {
		if group.TagName != "my-card" {
			t.Errorf("Group %d: expected TagName 'my-card', got '%s'", i, group.TagName)
		}
	}

	// Verify each group has knobs data
	for i, group := range knobGroups {
		if group.Knobs == nil {
			t.Errorf("Group %d: expected non-nil Knobs", i)
			continue
		}

		// Verify knobs were generated
		if len(group.Knobs.AttributeKnobs) == 0 {
			t.Errorf("Group %d: expected attribute knobs", i)
		}
	}

	// Verify current values are extracted per-instance
	// card-1 has variant="primary", card-2 has variant="warning", card-3 has no variant
	if len(knobGroups) >= 3 {
		// Find variant knob in each group
		for i, expectedVariant := range []string{"primary", "warning", ""} {
			var variantKnob *KnobData
			for j := range knobGroups[i].Knobs.AttributeKnobs {
				if knobGroups[i].Knobs.AttributeKnobs[j].Name == "variant" {
					variantKnob = &knobGroups[i].Knobs.AttributeKnobs[j]
					break
				}
			}

			if variantKnob == nil {
				t.Errorf("Group %d: variant knob not found", i)
				continue
			}

			if variantKnob.CurrentValue != expectedVariant {
				t.Errorf("Group %d: expected variant='%s', got '%s'", i, expectedVariant, variantKnob.CurrentValue)
			}
		}
	}
}

// TestRenderMultiInstanceKnobsHTML tests rendering of knob groups with <details> structure
func TestRenderMultiInstanceKnobsHTML(t *testing.T) {
	// Create sample multi-instance knobs data
	knobGroups := []ElementKnobGroup{
		{
			TagName:   "my-card",
			Label:     "#card-1",
			IsPrimary: true,
			Knobs: &KnobsData{
				TagName: "my-card",
				AttributeKnobs: []KnobData{
					{
						Name:         "variant",
						Category:     KnobCategoryAttribute,
						Type:         KnobTypeEnum,
						CurrentValue: "primary",
						EnumValues:   []string{"default", "primary", "warning"},
					},
				},
			},
		},
		{
			TagName:   "my-card",
			Label:     "#card-2",
			IsPrimary: false,
			Knobs: &KnobsData{
				TagName: "my-card",
				AttributeKnobs: []KnobData{
					{
						Name:         "variant",
						Category:     KnobCategoryAttribute,
						Type:         KnobTypeEnum,
						CurrentValue: "warning",
						EnumValues:   []string{"default", "primary", "warning"},
					},
				},
			},
		},
		{
			TagName:   "my-card",
			Label:     "#card-3",
			IsPrimary: false,
			Knobs: &KnobsData{
				TagName: "my-card",
				AttributeKnobs: []KnobData{
					{
						Name:         "variant",
						Category:     KnobCategoryAttribute,
						Type:         KnobTypeEnum,
						CurrentValue: "",
						EnumValues:   []string{"default", "primary", "warning"},
					},
				},
			},
		},
	}

	html, err := RenderKnobsHTML(testTemplatesForKnobs(), knobGroups)
	if err != nil {
		t.Fatalf("RenderKnobsHTML failed: %v", err)
	}

	htmlStr := string(html)

	// Verify structure is present
	if htmlStr == "" {
		t.Error("Expected non-empty HTML output")
	}

	// Verify cem-serve-knobs wrapper exists
	if !strings.Contains(htmlStr, "cem-serve-knobs") {
		t.Error("Expected HTML to contain cem-serve-knobs element")
	}

	// Count knobs-panel divs (one per instance)
	if !strings.Contains(htmlStr, `class="knobs-panel"`) {
		t.Error("Expected HTML to contain knobs-panel divs")
	}

	// Verify all three instances are present
	if !strings.Contains(htmlStr, "instance-0") || !strings.Contains(htmlStr, "instance-1") || !strings.Contains(htmlStr, "instance-2") {
		t.Error("Expected HTML to contain all three instance panels (instance-0, instance-1, instance-2)")
	}

	// Verify labels appear in nav links
	expectedLabels := []string{"#card-1", "#card-2", "#card-3"}
	for _, label := range expectedLabels {
		if !strings.Contains(htmlStr, label) {
			t.Errorf("Expected HTML to contain label '%s'", label)
		}
	}

	// Verify tag name appears
	if !strings.Contains(htmlStr, "my-card") {
		t.Error("Expected HTML to contain tag name 'my-card'")
	}
}
