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
package completion_test

import (
	"encoding/json"
	"testing"

	"bennypowers.dev/cem/internal/platform/testutil"
	"bennypowers.dev/cem/lsp/document"
	"bennypowers.dev/cem/lsp/methods/textDocument/completion"
	"bennypowers.dev/cem/lsp/testhelpers"
	M "bennypowers.dev/cem/manifest"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// TestSlotAttributeNameSuggestionRegression ensures slot attribute suggestions work correctly
// and prevents future regressions of this critical completion feature
func TestSlotAttributeNameSuggestionRegression(t *testing.T) {
	// Load test manifest with multiple elements that have different slot configurations
	fs := testutil.NewFixtureFS(t, "legacy/slot-attribute-regression", "/test")
	manifestBytes, err := fs.ReadFile("/test/manifest.json")
	if err != nil {
		t.Fatalf("Failed to read regression test manifest: %v", err)
	}

	var pkg M.Package
	err = json.Unmarshal(manifestBytes, &pkg)
	if err != nil {
		t.Fatalf("Failed to parse regression test manifest: %v", err)
	}

	// Create context and add the test manifest
	ctx := testhelpers.NewMockServerContext()
	ctx.AddManifest(&pkg)

	// Create and set a real DocumentManager
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)

	// Regression test cases covering all scenarios where slot attribute should/shouldn't be suggested
	tests := []struct {
		name                string
		html                string
		position            protocol.Position
		tagName             string
		shouldSuggestSlot   bool
		expectedCompletions []string // Additional completions that should be present
		description         string
	}{
		// POSITIVE CASES: Should suggest slot attribute
		{
			name:              "Standard HTML element inside custom element with named slots",
			html:              `<card-layout><div `,
			position:          protocol.Position{Line: 0, Character: 17},
			tagName:           "div",
			shouldSuggestSlot: true,
			description:       "div inside card-layout should suggest slot attribute",
		},
		{
			name:              "Button inside navigation component",
			html:              `<nav-menu><button `,
			position:          protocol.Position{Line: 0, Character: 18},
			tagName:           "button",
			shouldSuggestSlot: true,
			description:       "button inside nav-menu should suggest slot attribute",
		},
		{
			name:              "Span inside modal with header/footer slots",
			html:              `<modal-dialog><span `,
			position:          protocol.Position{Line: 0, Character: 20},
			tagName:           "span",
			shouldSuggestSlot: true,
			description:       "span inside modal-dialog should suggest slot attribute",
		},
		{
			name:                "Custom element inside another custom element with slots",
			html:                `<card-layout><icon-button `,
			position:            protocol.Position{Line: 0, Character: 26},
			tagName:             "icon-button",
			shouldSuggestSlot:   true,
			expectedCompletions: []string{"variant", "disabled"}, // icon-button's own attributes
			description:         "icon-button inside card-layout should suggest both slot and its own attributes",
		},
		{
			name:              "Nested structure - grandchild of slotted element",
			html:              `<card-layout><header><h1 `,
			position:          protocol.Position{Line: 0, Character: 24},
			tagName:           "h1",
			shouldSuggestSlot: true,
			description:       "h1 nested inside header inside card-layout should suggest slot (header is direct child)",
		},
		{
			name:              "Complex nesting with multiple custom elements",
			html:              `<page-layout><card-layout><button `,
			position:          protocol.Position{Line: 0, Character: 34},
			tagName:           "button",
			shouldSuggestSlot: true,
			description:       "button inside card-layout inside page-layout should suggest slot (card-layout is direct parent)",
		},

		// NEGATIVE CASES: Should NOT suggest slot attribute
		{
			name:              "Element inside standard HTML element",
			html:              `<div><button `,
			position:          protocol.Position{Line: 0, Character: 13},
			tagName:           "button",
			shouldSuggestSlot: false,
			description:       "button inside div should not suggest slot attribute",
		},
		{
			name:              "Top-level element with no parent",
			html:              `<button `,
			position:          protocol.Position{Line: 0, Character: 8},
			tagName:           "button",
			shouldSuggestSlot: false,
			description:       "top-level button should not suggest slot attribute",
		},
		{
			name:              "Element inside custom element with only anonymous slot",
			html:              `<simple-wrapper><p `,
			position:          protocol.Position{Line: 0, Character: 19},
			tagName:           "p",
			shouldSuggestSlot: false,
			description:       "p inside simple-wrapper (only default slot) should not suggest slot attribute",
		},
		{
			name:              "Element inside custom element with no slots",
			html:              `<data-provider><span `,
			position:          protocol.Position{Line: 0, Character: 21},
			tagName:           "span",
			shouldSuggestSlot: false,
			description:       "span inside data-provider (no slots) should not suggest slot attribute",
		},
		{
			name:              "Deeply nested in standard HTML",
			html:              `<div><section><article><p `,
			position:          protocol.Position{Line: 0, Character: 26},
			tagName:           "p",
			shouldSuggestSlot: false,
			description:       "p deeply nested in standard HTML should not suggest slot attribute",
		},

		// EDGE CASES: Boundary conditions and complex scenarios
		{
			name:              "Self-closing tag inside slotted element",
			html:              `<card-layout><img `,
			position:          protocol.Position{Line: 0, Character: 18},
			tagName:           "img",
			shouldSuggestSlot: true,
			description:       "self-closing img inside card-layout should suggest slot attribute",
		},
		{
			name:              "Element with existing attributes inside slotted element",
			html:              `<nav-menu><button class="primary" `,
			position:          protocol.Position{Line: 0, Character: 34},
			tagName:           "button",
			shouldSuggestSlot: true,
			description:       "button with existing attributes inside nav-menu should still suggest slot",
		},
		{
			name:              "Custom element with mixed slot types",
			html:              `<complex-layout><section `,
			position:          protocol.Position{Line: 0, Character: 25},
			tagName:           "section",
			shouldSuggestSlot: true,
			description:       "section inside complex-layout (has both named and default slots) should suggest slot",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a real document using the DocumentManager
			doc := dm.OpenDocument("test://test.html", tt.html, 1)

			// Call the attribute completion function with context
			completions := completion.GetAttributeCompletionsWithContext(ctx, doc, tt.position, tt.tagName)

			// Check if slot attribute is suggested
			foundSlot := false
			slotCompletion := protocol.CompletionItem{}
			allLabels := make([]string, len(completions))

			for i, completion := range completions {
				allLabels[i] = completion.Label
				if completion.Label == "slot" {
					foundSlot = true
					slotCompletion = completion
				}
			}

			// Verify slot suggestion expectation
			if tt.shouldSuggestSlot && !foundSlot {
				t.Errorf("REGRESSION: Expected slot attribute to be suggested but it wasn't.\nHTML: %s\nPosition: line=%d, char=%d\nTag: %s\nGot completions: %v",
					tt.html, tt.position.Line, tt.position.Character, tt.tagName, allLabels)
			}

			if !tt.shouldSuggestSlot && foundSlot {
				t.Errorf("REGRESSION: Expected slot attribute NOT to be suggested but it was.\nHTML: %s\nPosition: line=%d, char=%d\nTag: %s\nGot completions: %v",
					tt.html, tt.position.Line, tt.position.Character, tt.tagName, allLabels)
			}

			// If slot should be suggested, verify its structure
			if tt.shouldSuggestSlot && foundSlot {
				// Verify slot completion has proper structure
				if slotCompletion.InsertText == nil || *slotCompletion.InsertText != `slot="$0"` {
					t.Errorf("REGRESSION: slot completion should have snippet insert text 'slot=\"$0\"', got: %v", slotCompletion.InsertText)
				}

				if slotCompletion.Kind == nil || *slotCompletion.Kind != protocol.CompletionItemKindProperty {
					t.Errorf("REGRESSION: slot completion should have Property kind, got: %v", slotCompletion.Kind)
				}

				if slotCompletion.Detail == nil || *slotCompletion.Detail != "HTML slot attribute" {
					t.Errorf("REGRESSION: slot completion should have proper detail, got: %v", slotCompletion.Detail)
				}
			}

			// Verify expected additional completions are present
			for _, expectedCompletion := range tt.expectedCompletions {
				found := false
				for _, label := range allLabels {
					if label == expectedCompletion {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("REGRESSION: Expected completion '%s' not found in results: %v", expectedCompletion, allLabels)
				}
			}

			// Log test results for debugging
			t.Logf("✓ %s", tt.description)
			t.Logf("  Found %d completions: %v", len(completions), allLabels)
			if foundSlot {
				t.Logf("  ✓ slot attribute suggested correctly")
			}
		})
	}
}

// TestSlotAttributeParentDetectionRegression tests the parent element detection logic
// that was fixed to properly find parent elements instead of current elements
func TestSlotAttributeParentDetectionRegression(t *testing.T) {
	// This is a focused regression test for the specific bug that was fixed:
	// findParentElementTag was returning the current element instead of parent element

	fs := testutil.NewFixtureFS(t, "legacy/slot-attribute-regression", "/test")
	manifestBytes, err := fs.ReadFile("/test/manifest.json")
	if err != nil {
		t.Fatalf("Failed to read regression test manifest: %v", err)
	}

	var pkg M.Package
	err = json.Unmarshal(manifestBytes, &pkg)
	if err != nil {
		t.Fatalf("Failed to parse regression test manifest: %v", err)
	}

	ctx := testhelpers.NewMockServerContext()
	ctx.AddManifest(&pkg)

	// Create and set a real DocumentManager
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)

	// Test cases specifically for the parent detection bug
	parentDetectionTests := []struct {
		name        string
		html        string
		position    protocol.Position
		tagName     string
		shouldFind  bool
		description string
	}{
		{
			name:        "Custom element inside custom element - should skip self and find parent",
			html:        `<card-layout><icon-button `,
			position:    protocol.Position{Line: 0, Character: 26},
			tagName:     "icon-button",
			shouldFind:  true,
			description: "Should find card-layout as parent, not icon-button as self",
		},
		{
			name:        "Standard element inside custom element - should find parent",
			html:        `<nav-menu><button `,
			position:    protocol.Position{Line: 0, Character: 18},
			tagName:     "button",
			shouldFind:  true,
			description: "Should find nav-menu as parent when current element is standard HTML",
		},
		{
			name:        "Element inside standard HTML - should not find slotted parent",
			html:        `<div><button `,
			position:    protocol.Position{Line: 0, Character: 13},
			tagName:     "button",
			shouldFind:  false,
			description: "Should not find slotted parent when inside standard HTML div",
		},
		{
			name:        "Deeply nested custom elements - should find immediate parent",
			html:        `<page-layout><card-layout><icon-button `,
			position:    protocol.Position{Line: 0, Character: 39},
			tagName:     "icon-button",
			shouldFind:  true,
			description: "Should find card-layout as immediate parent, not page-layout",
		},
	}

	for _, tt := range parentDetectionTests {
		t.Run(tt.name, func(t *testing.T) {
			doc := dm.OpenDocument("test://test.html", tt.html, 1)
			completions := completion.GetAttributeCompletionsWithContext(ctx, doc, tt.position, tt.tagName)

			foundSlot := false
			for _, completion := range completions {
				if completion.Label == "slot" {
					foundSlot = true
					break
				}
			}

			if tt.shouldFind && !foundSlot {
				t.Errorf("PARENT DETECTION REGRESSION: Expected to find slot attribute (parent detection failed).\nHTML: %s\nPosition: line=%d, char=%d\nTag: %s",
					tt.html, tt.position.Line, tt.position.Character, tt.tagName)
			}

			if !tt.shouldFind && foundSlot {
				t.Errorf("PARENT DETECTION REGRESSION: Expected NOT to find slot attribute (false positive parent detection).\nHTML: %s\nPosition: line=%d, char=%d\nTag: %s",
					tt.html, tt.position.Line, tt.position.Character, tt.tagName)
			}

			t.Logf("✓ %s", tt.description)
		})
	}
}

// TestSlotAttributeCompletionStructureRegression ensures slot completions have proper LSP structure
func TestSlotAttributeCompletionStructureRegression(t *testing.T) {
	fs := testutil.NewFixtureFS(t, "legacy/slot-attribute-regression", "/test")
	manifestBytes, err := fs.ReadFile("/test/manifest.json")
	if err != nil {
		t.Fatalf("Failed to read regression test manifest: %v", err)
	}

	var pkg M.Package
	err = json.Unmarshal(manifestBytes, &pkg)
	if err != nil {
		t.Fatalf("Failed to parse regression test manifest: %v", err)
	}

	ctx := testhelpers.NewMockServerContext()
	ctx.AddManifest(&pkg)

	// Create and set a real DocumentManager
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)

	// Test that slot completion has proper LSP protocol structure
	doc := dm.OpenDocument("test://test.html", `<card-layout><button `, 1)
	position := protocol.Position{Line: 0, Character: 20}
	completions := completion.GetAttributeCompletionsWithContext(ctx, doc, position, "button")

	var slotCompletion *protocol.CompletionItem
	for _, completion := range completions {
		if completion.Label == "slot" {
			slotCompletion = &completion
			break
		}
	}

	if slotCompletion == nil {
		t.Fatal("STRUCTURE REGRESSION: slot completion not found")
	}

	// Test all required fields are present and correct
	tests := []struct {
		name     string
		test     func() bool
		expected string
		actual   string
	}{
		{
			name:     "Label should be 'slot'",
			test:     func() bool { return slotCompletion.Label == "slot" },
			expected: "slot",
			actual:   slotCompletion.Label,
		},
		{
			name: "Kind should be Property",
			test: func() bool {
				return slotCompletion.Kind != nil && *slotCompletion.Kind == protocol.CompletionItemKindProperty
			},
			expected: "Property",
			actual: func() string {
				if slotCompletion.Kind == nil {
					return "nil"
				}
				return string(*slotCompletion.Kind)
			}(),
		},
		{
			name: "Detail should be 'HTML slot attribute'",
			test: func() bool {
				return slotCompletion.Detail != nil && *slotCompletion.Detail == "HTML slot attribute"
			},
			expected: "HTML slot attribute",
			actual: func() string {
				if slotCompletion.Detail == nil {
					return "nil"
				}
				return *slotCompletion.Detail
			}(),
		},
		{
			name: "InsertText should be snippet 'slot=\"$0\"'",
			test: func() bool {
				return slotCompletion.InsertText != nil && *slotCompletion.InsertText == `slot="$0"`
			},
			expected: `slot="$0"`,
			actual: func() string {
				if slotCompletion.InsertText == nil {
					return "nil"
				}
				return *slotCompletion.InsertText
			}(),
		},
		{
			name: "InsertTextFormat should be Snippet",
			test: func() bool {
				return slotCompletion.InsertTextFormat != nil && *slotCompletion.InsertTextFormat == protocol.InsertTextFormatSnippet
			},
			expected: "Snippet",
			actual: func() string {
				if slotCompletion.InsertTextFormat == nil {
					return "nil"
				}
				return string(*slotCompletion.InsertTextFormat)
			}(),
		},
		{
			name: "Documentation should be present",
			test: func() bool {
				return slotCompletion.Documentation != nil
			},
			expected: "present",
			actual: func() string {
				if slotCompletion.Documentation == nil {
					return "nil"
				}
				return "present"
			}(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if !test.test() {
				t.Errorf("STRUCTURE REGRESSION: %s\nExpected: %s\nActual: %s", test.name, test.expected, test.actual)
			} else {
				t.Logf("✓ %s", test.name)
			}
		})
	}
}
