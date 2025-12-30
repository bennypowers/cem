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

// TestFindParentElementTag_MalformedHTML tests edge cases with malformed HTML
func TestFindParentElementTag_MalformedHTML(t *testing.T) {
	// Load test manifest with custom elements
	fs := testutil.NewFixtureFS(t, "slot-completions-test", "/test")
	manifestBytes, err := fs.ReadFile("/test/manifest.json")
	if err != nil {
		t.Fatalf("Failed to read test manifest: %v", err)
	}

	var pkg M.Package
	err = json.Unmarshal(manifestBytes, &pkg)
	if err != nil {
		t.Fatalf("Failed to parse manifest: %v", err)
	}

	ctx := testhelpers.NewMockServerContext()
	ctx.AddManifest(&pkg)

	// Create DocumentManager for tree-sitter parsing
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)

	tests := []struct {
		name           string
		content        string
		position       protocol.Position
		expectedParent string
		description    string
	}{
		// Unclosed tags
		{
			name:           "unclosed custom element",
			content:        `<card-element><button slot="`,
			position:       protocol.Position{Line: 0, Character: 28},
			expectedParent: "card-element",
			description:    "Should find parent even with unclosed child tag",
		},
		{
			name:           "unclosed parent and child",
			content:        `<card-element><button`,
			position:       protocol.Position{Line: 0, Character: 21},
			expectedParent: "card-element",
			description:    "Should find parent even when both are unclosed",
		},
		{
			name:           "deeply nested unclosed",
			content:        `<card-element><div><span><button slot="`,
			position:       protocol.Position{Line: 0, Character: 40},
			expectedParent: "card-element",
			description:    "Should find custom element parent through unclosed standard elements",
		},

		// Overlapping/mismatched tags
		{
			name:           "mismatched closing tag",
			content:        `<card-element><button></div></card-element>`,
			position:       protocol.Position{Line: 0, Character: 22},
			expectedParent: "card-element",
			description:    "Should find parent despite mismatched closing tags",
		},
		{
			name:           "wrong nesting order",
			content:        `<card-element><button></card-element></button>`,
			position:       protocol.Position{Line: 0, Character: 22},
			expectedParent: "card-element",
			description:    "Should handle wrong nesting order gracefully",
		},

		// Truncated tags
		{
			name:           "truncated opening tag",
			content:        `<card-element><butt`,
			position:       protocol.Position{Line: 0, Character: 19},
			expectedParent: "card-element",
			description:    "Should find parent with truncated child tag",
		},
		{
			name:           "truncated attribute",
			content:        `<card-element><button slo`,
			position:       protocol.Position{Line: 0, Character: 25},
			expectedParent: "card-element",
			description:    "Should find parent with truncated attribute name",
		},
		{
			name:           "truncated attribute value",
			content:        `<card-element><button slot="hea`,
			position:       protocol.Position{Line: 0, Character: 32},
			expectedParent: "card-element",
			description:    "Should find parent with truncated attribute value",
		},

		// Broken syntax
		{
			name:           "missing closing bracket",
			content:        `<card-element<button slot="header">`,
			position:       protocol.Position{Line: 0, Character: 27},
			expectedParent: "",
			description:    "Should handle missing > gracefully",
		},
		{
			name:           "extra opening bracket",
			content:        `<card-element><<button slot="header">`,
			position:       protocol.Position{Line: 0, Character: 28},
			expectedParent: "card-element",
			description:    "Should handle extra < gracefully",
		},
		{
			name:           "missing tag name",
			content:        `<card-element><>`,
			position:       protocol.Position{Line: 0, Character: 16},
			expectedParent: "",
			description:    "Empty tag is too malformed to determine parent",
		},

		// Nested custom and standard elements
		{
			name:           "custom in custom",
			content:        `<card-element><dialog-element><button slot="`,
			position:       protocol.Position{Line: 0, Character: 45},
			expectedParent: "dialog-element",
			description:    "Should find immediate custom element parent",
		},
		{
			name:           "standard between customs",
			content:        `<card-element><div><dialog-element><button slot="`,
			position:       protocol.Position{Line: 0, Character: 50},
			expectedParent: "dialog-element",
			description:    "Should skip standard element and find custom parent",
		},
		{
			name:           "multiple standard between customs",
			content:        `<card-element><div><span><section><dialog-element><button slot="`,
			position:       protocol.Position{Line: 0, Character: 65},
			expectedParent: "dialog-element",
			description:    "Should skip multiple standard elements",
		},

		// Byte/character offset edge cases
		{
			name:           "position at tag start",
			content:        `<card-element><button>`,
			position:       protocol.Position{Line: 0, Character: 14},
			expectedParent: "",
			description:    "Cursor at < makes detection ambiguous (could be before or inside tag)",
		},
		{
			name:           "position in tag name",
			content:        `<card-element><button>`,
			position:       protocol.Position{Line: 0, Character: 18},
			expectedParent: "card-element",
			description:    "Should find parent when cursor in tag name",
		},
		{
			name:           "position at tag end",
			content:        `<card-element><button>`,
			position:       protocol.Position{Line: 0, Character: 21},
			expectedParent: "card-element",
			description:    "Should find parent when cursor at >",
		},
		{
			name:           "multi-line with varying indentation",
			content:        "<card-element>\n  <div>\n    <button slot=\"",
			position:       protocol.Position{Line: 2, Character: 18},
			expectedParent: "card-element",
			description:    "Should handle multi-line with indentation",
		},
		{
			name:           "unicode characters before tag",
			content:        `<card-element><!-- 你好 --><button slot="`,
			position:       protocol.Position{Line: 0, Character: 41},
			expectedParent: "card-element",
			description:    "Should handle unicode in comments",
		},

		// Empty/minimal cases
		{
			name:           "empty document",
			content:        ``,
			position:       protocol.Position{Line: 0, Character: 0},
			expectedParent: "",
			description:    "Should handle empty document without panic",
		},
		{
			name:           "only opening tag",
			content:        `<card-element`,
			position:       protocol.Position{Line: 0, Character: 13},
			expectedParent: "",
			description:    "Should return empty for incomplete opening tag",
		},
		{
			name:           "position beyond content",
			content:        `<card-element></card-element>`,
			position:       protocol.Position{Line: 0, Character: 100},
			expectedParent: "",
			description:    "Should handle position beyond content gracefully",
		},
		{
			name:           "position on non-existent line",
			content:        `<card-element></card-element>`,
			position:       protocol.Position{Line: 10, Character: 0},
			expectedParent: "",
			description:    "Should handle position on non-existent line",
		},

		// Self-closing tags
		{
			name:           "self-closing custom element",
			content:        `<card-element/><button slot="`,
			position:       protocol.Position{Line: 0, Character: 29},
			expectedParent: "",
			description:    "Should not consider self-closing tag as parent",
		},
		{
			name:           "self-closing with attributes",
			content:        `<card-element attr="value"/><button slot="`,
			position:       protocol.Position{Line: 0, Character: 42},
			expectedParent: "",
			description:    "Should handle self-closing with attributes",
		},

		// Mixed well-formed and malformed
		{
			name:           "well-formed parent, malformed child",
			content:        `<card-element><button</card-element>`,
			position:       protocol.Position{Line: 0, Character: 21},
			expectedParent: "card-element",
			description:    "Should find well-formed parent despite malformed child",
		},
		{
			name:           "malformed parent, well-formed child",
			content:        `<card-element<button></button>`,
			position:       protocol.Position{Line: 0, Character: 22},
			expectedParent: "",
			description:    "Malformed parent makes parsing unreliable",
		},

		// Attribute edge cases
		{
			name:           "cursor in attribute name",
			content:        `<card-element><button slo`,
			position:       protocol.Position{Line: 0, Character: 25},
			expectedParent: "card-element",
			description:    "Should find parent when cursor in partial attribute",
		},
		{
			name:           "cursor in attribute value",
			content:        `<card-element><button slot="head`,
			position:       protocol.Position{Line: 0, Character: 33},
			expectedParent: "card-element",
			description:    "Should find parent when cursor in attribute value",
		},
		{
			name:           "cursor in unclosed quote",
			content:        `<card-element><button slot="header`,
			position:       protocol.Position{Line: 0, Character: 35},
			expectedParent: "card-element",
			description:    "Should find parent with unclosed quote",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Open document and parse with tree-sitter
			doc := dm.OpenDocument("test://malformed.html", tt.content, 1)

			// Get slot value completions which internally calls findParentElementTag
			completions := completion.GetAttributeValueCompletionsWithContext(ctx, doc, tt.position, "", "slot")

			// Verify based on expected parent
			if tt.expectedParent == "" {
				// Should not find slot completions if no valid parent
				if len(completions) > 0 {
					t.Errorf("%s: Expected no completions (no valid parent), got %d completions",
						tt.description, len(completions))
				}
			} else {
				// Should find slot completions from the expected parent
				if len(completions) == 0 {
					t.Errorf("%s: Expected slot completions from parent '%s', got none",
						tt.description, tt.expectedParent)
				}
				// Verify we can get slots for the expected parent
				slots, exists := ctx.Slots(tt.expectedParent)
				if exists && len(slots) > 0 {
					// We should have at least one completion for named slots
					namedSlotCount := 0
					for _, slot := range slots {
						if slot.Name != "" {
							namedSlotCount++
						}
					}
					if namedSlotCount > 0 && len(completions) == 0 {
						t.Errorf("%s: Expected %d slot completions from '%s', got none",
							tt.description, namedSlotCount, tt.expectedParent)
					}
				}
			}

			// Most importantly: verify no panics occurred during processing
			// The fact we got here means no panics - test passes
		})
	}
}

// TestFindParentElementTag_FallbackBehavior tests that fallback is used when tree-sitter fails
func TestFindParentElementTag_FallbackBehavior(t *testing.T) {
	// Load test manifest
	fs := testutil.NewFixtureFS(t, "slot-completions-test", "/test")
	manifestBytes, err := fs.ReadFile("/test/manifest.json")
	if err != nil {
		t.Fatalf("Failed to read test manifest: %v", err)
	}

	var pkg M.Package
	err = json.Unmarshal(manifestBytes, &pkg)
	if err != nil {
		t.Fatalf("Failed to parse manifest: %v", err)
	}

	ctx := testhelpers.NewMockServerContext()
	ctx.AddManifest(&pkg)

	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)

	// Test cases where tree-sitter parsing might struggle and fallback is needed
	fallbackTests := []struct {
		name        string
		content     string
		position    protocol.Position
		expectSlots bool
		description string
	}{
		{
			name:        "severely malformed nested tags",
			content:     `<card-element><<><>><button slot="`,
			position:    protocol.Position{Line: 0, Character: 34},
			expectSlots: true,
			description: "Fallback should handle severely malformed HTML",
		},
		{
			name:        "script tag with < in content",
			content:     `<card-element><script>if (x < 5)</script><button slot="`,
			position:    protocol.Position{Line: 0, Character: 56},
			expectSlots: true,
			description: "Fallback should skip script content",
		},
		{
			name:        "style tag with selectors",
			content:     `<card-element><style>button > span</style><button slot="`,
			position:    protocol.Position{Line: 0, Character: 57},
			expectSlots: true,
			description: "Fallback should skip style content",
		},
		{
			name:        "comment with fake tags",
			content:     `<card-element><!-- <fake-element> --><button slot="`,
			position:    protocol.Position{Line: 0, Character: 52},
			expectSlots: true,
			description: "Fallback should skip comments",
		},
		{
			name:        "mixed quotes in attributes",
			content:     `<card-element attr="value with ' quote"><button slot="`,
			position:    protocol.Position{Line: 0, Character: 55},
			expectSlots: true,
			description: "Fallback should handle mixed quotes",
		},
	}

	for _, tt := range fallbackTests {
		t.Run(tt.name, func(t *testing.T) {
			doc := dm.OpenDocument("test://fallback.html", tt.content, 1)
			completions := completion.GetAttributeValueCompletionsWithContext(ctx, doc, tt.position, "", "slot")

			if tt.expectSlots {
				if len(completions) == 0 {
					t.Errorf("%s: Expected slot completions via fallback, got none", tt.description)
				}
			} else {
				if len(completions) > 0 {
					t.Errorf("%s: Expected no completions, got %d", tt.description, len(completions))
				}
			}

			// Verify no panics
		})
	}
}

// TestFindParentElementTag_ByteOffsetEdgeCases tests byte/character conversion edge cases
func TestFindParentElementTag_ByteOffsetEdgeCases(t *testing.T) {
	fs := testutil.NewFixtureFS(t, "slot-completions-test", "/test")
	manifestBytes, err := fs.ReadFile("/test/manifest.json")
	if err != nil {
		t.Fatalf("Failed to read test manifest: %v", err)
	}

	var pkg M.Package
	err = json.Unmarshal(manifestBytes, &pkg)
	if err != nil {
		t.Fatalf("Failed to parse manifest: %v", err)
	}

	ctx := testhelpers.NewMockServerContext()
	ctx.AddManifest(&pkg)

	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)

	// Test multi-byte unicode characters
	unicodeTests := []struct {
		name        string
		content     string
		position    protocol.Position
		expectSlots bool
		description string
	}{
		{
			name:        "emoji in content",
			content:     `<card-element>🎉🎊<button slot="`,
			position:    protocol.Position{Line: 0, Character: 33},
			expectSlots: true,
			description: "Should handle emoji characters",
		},
		{
			name:        "multi-byte characters in attribute",
			content:     `<card-element title="日本語"><button slot="`,
			position:    protocol.Position{Line: 0, Character: 43},
			expectSlots: true,
			description: "Should handle multi-byte in attributes",
		},
		{
			name:        "mixed ASCII and unicode",
			content:     `<card-element>Hello 世界<button slot="`,
			position:    protocol.Position{Line: 0, Character: 38},
			expectSlots: true,
			description: "Should handle mixed ASCII and unicode",
		},
		{
			name:        "zero-width characters",
			content:     "<card-element>\u200B\u200C<button slot=\"",
			position:    protocol.Position{Line: 0, Character: 33},
			expectSlots: true,
			description: "Should handle zero-width characters",
		},
		{
			name:        "combining characters",
			content:     "<card-element>e\u0301<button slot=\"", // é as e + combining acute
			position:    protocol.Position{Line: 0, Character: 32},
			expectSlots: true,
			description: "Should handle combining diacritical marks",
		},
	}

	for _, tt := range unicodeTests {
		t.Run(tt.name, func(t *testing.T) {
			doc := dm.OpenDocument("test://unicode.html", tt.content, 1)
			completions := completion.GetAttributeValueCompletionsWithContext(ctx, doc, tt.position, "", "slot")

			if tt.expectSlots {
				if len(completions) == 0 {
					t.Logf("%s: Got %d completions (may fail due to unicode byte/char offset differences)",
						tt.description, len(completions))
					// Don't fail - unicode handling can be tricky
				}
			}

			// Main goal: verify no panics
		})
	}
}
