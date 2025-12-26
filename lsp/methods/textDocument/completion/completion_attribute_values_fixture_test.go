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
	"bennypowers.dev/cem/lsp/methods/textDocument/completion"
	"bennypowers.dev/cem/lsp/testhelpers"
	M "bennypowers.dev/cem/manifest"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// Completion test configuration for each fixture (in Go code, not JSON)
var completionTestConfigs = map[string]struct {
	TagName       string
	AttributeName string
}{
	// Original test-component fixtures
	"union-type":         {TagName: "test-component", AttributeName: "theme"},
	"size-attr":          {TagName: "test-component", AttributeName: "size"},
	"variant-attr":       {TagName: "test-component", AttributeName: "variant"},
	"string-attr":        {TagName: "test-component", AttributeName: "color"},
	"boolean-attr":       {TagName: "test-component", AttributeName: "disabled"},
	"number-attr":        {TagName: "test-component", AttributeName: "count"},
	"array-type":         {TagName: "test-component", AttributeName: "items"},
	"non-custom-element": {TagName: "div", AttributeName: "class"},
	"non-existent-attr":  {TagName: "test-component", AttributeName: "nonexistent"},

	// test-alert fixtures (converted from completion_unit_test.go)
	"test-alert-state":    {TagName: "test-alert", AttributeName: "state"},
	"test-alert-variant":  {TagName: "test-alert", AttributeName: "variant"},
	"test-alert-disabled": {TagName: "test-alert", AttributeName: "disabled"},
}

// TestAttributeValueCompletions_NewFixtures runs all attribute value completion tests
func TestAttributeValueCompletions_NewFixtures(t *testing.T) {
	testutil.RunLSPFixtures(t, "testdata/attribute-value-completions", func(t *testing.T, fixture *testutil.LSPFixture) {
		// Get test config from map
		config, ok := completionTestConfigs[fixture.Name]
		if !ok {
			t.Fatalf("No test config defined for fixture %s", fixture.Name)
		}

		// Create mock server context
		ctx := testhelpers.NewMockServerContext()

		// Load manifest if present
		if len(fixture.Manifest) > 0 {
			var manifestData CompletionManifestFixture
			if err := json.Unmarshal(fixture.Manifest, &manifestData); err != nil {
				t.Fatalf("Failed to parse manifest.json: %v", err)
			}

			// Add elements and attributes to context
			for tagName, elementData := range manifestData.Elements {
				// Add attributes
				attrMap := make(map[string]*M.Attribute)
				for attrName, attrData := range elementData.Attributes {
					attr := &M.Attribute{
						FullyQualified: M.FullyQualified{
							Name: attrData.Name,
						},
						Default: attrData.Default,
					}
					if attrData.Type != nil {
						attr.Type = &M.Type{
							Text: attrData.Type.Text,
						}
					}
					attrMap[attrName] = attr
				}
				ctx.AddAttributes(tagName, attrMap)
			}
		}

		// Load expected completions (actual LSP CompletionItem[] response)
		var expected []protocol.CompletionItem
		err := fixture.GetExpected("expected", &expected)
		if err != nil {
			t.Fatalf("Failed to load expected completions: %v", err)
		}

		// Get actual completions
		actual := completion.GetAttributeValueCompletions(ctx, config.TagName, config.AttributeName)

		// Verify completion count matches
		if len(actual) != len(expected) {
			t.Errorf("Expected %d completions, got %d", len(expected), len(actual))
			t.Logf("Expected labels: %v", getLabels(expected))
			t.Logf("Actual labels: %v", getLabels(actual))
		}

		// Verify each expected completion is present
		for i, exp := range expected {
			if i >= len(actual) {
				t.Errorf("Missing completion at index %d: %+v", i, exp)
				continue
			}

			act := actual[i]

			// Compare key fields
			if act.Label != exp.Label {
				t.Errorf("Completion %d: expected label %q, got %q", i, exp.Label, act.Label)
			}

			// Compare Kind (handle pointer)
			if exp.Kind != nil && act.Kind != nil {
				if *exp.Kind != *act.Kind {
					t.Errorf("Completion %d (%s): expected kind %d, got %d", i, exp.Label, *exp.Kind, *act.Kind)
				}
			} else if exp.Kind != nil || act.Kind != nil {
				t.Errorf("Completion %d (%s): kind mismatch (one is nil)", i, exp.Label)
			}
			// Compare InsertText (handle pointer)
			expInsert := ""
			if exp.InsertText != nil {
				expInsert = *exp.InsertText
			}
			actInsert := ""
			if act.InsertText != nil {
				actInsert = *act.InsertText
			}
			if actInsert != expInsert && expInsert != "" {
				t.Errorf("Completion %d (%s): expected insertText %q, got %q", i, exp.Label, expInsert, actInsert)
			}
		}
	})
}

// Helper to extract labels from completions for debugging
func getLabels(items []protocol.CompletionItem) []string {
	labels := make([]string, len(items))
	for i, item := range items {
		labels[i] = item.Label
	}
	return labels
}

// CompletionManifestFixture represents simplified manifest data for completion tests
type CompletionManifestFixture struct {
	Elements map[string]struct {
		TagName    string `json:"tagName"`
		Attributes map[string]struct {
			Name    string `json:"name"`
			Type    *struct {
				Text string `json:"text"`
			} `json:"type,omitempty"`
			Default string `json:"default,omitempty"`
		} `json:"attributes"`
	} `json:"elements"`
}

// ExpectedCompletions represents the OLD expected completions format (for migration)
type ExpectedCompletions struct {
	TagName        string   `json:"tagName"`
	AttributeName  string   `json:"attributeName"`
	ExpectedLabels []string `json:"expectedLabels"`
	Description    string   `json:"description"`
}
