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
	"testing"

	M "bennypowers.dev/cem/manifest"
	"github.com/google/go-cmp/cmp"
)

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
		TagName: "my-button",
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

	html, err := RenderKnobsHTML(knobs)
	if err != nil {
		t.Fatalf("RenderKnobsHTML failed: %v", err)
	}

	htmlStr := string(html)

	// Verify structure is present
	if htmlStr == "" {
		t.Error("Expected non-empty HTML output")
	}

	// Verify sections are present
	expectedStrings := []string{
		"knobs-container",
		"Attributes",
		"Properties",
		"CSS Custom Properties",
		"disabled",
		"variant",
		"label",
		"--button-color",
	}

	for _, expected := range expectedStrings {
		if !contains(htmlStr, expected) {
			t.Errorf("Expected HTML to contain '%s'", expected)
		}
	}
}

// Helper function for string contains
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 || (len(s) > 0 && len(substr) > 0 && hasSubstring(s, substr)))
}

func hasSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
