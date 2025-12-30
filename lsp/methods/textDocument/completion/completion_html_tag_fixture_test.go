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
	"strings"
	"testing"

	"bennypowers.dev/cem/internal/platform/testutil"
	"bennypowers.dev/cem/lsp/document"
	"bennypowers.dev/cem/lsp/methods/textDocument/completion"
	"bennypowers.dev/cem/lsp/testhelpers"
	M "bennypowers.dev/cem/manifest"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// TestHTMLTagCompletions_Fixtures tests HTML tag name completions using fixture/golden pattern
func TestHTMLTagCompletions_Fixtures(t *testing.T) {
	testutil.RunLSPFixtures(t, "testdata/html-tag-completions", func(t *testing.T, fixture *testutil.LSPFixture) {
		// Create mock server context
		ctx := testhelpers.NewMockServerContext()

		// Load and parse manifest
		if len(fixture.Manifest) > 0 {
			var pkg M.Package
			if err := json.Unmarshal(fixture.Manifest, &pkg); err != nil {
				t.Fatalf("Failed to parse manifest: %v", err)
			}
			ctx.AddManifest(&pkg)
		}

		// Create and set DocumentManager
		dm, err := document.NewDocumentManager()
		if err != nil {
			t.Fatalf("Failed to create DocumentManager: %v", err)
		}
		defer dm.Close()
		ctx.SetDocumentManager(dm)

		// Open HTML document from fixture
		htmlURI := "file:///test/" + fixture.Name + ".html"
		doc := dm.OpenDocument(htmlURI, fixture.InputContent, 1)
		if doc == nil {
			t.Fatal("Failed to open HTML document")
		}
		ctx.AddDocument(htmlURI, doc)

		// Calculate cursor position - find the last < and position cursor after it
		lines := strings.Split(fixture.InputContent, "\n")
		var position protocol.Position
		found := false

		// Search backwards through lines to find the last standalone <
		for i := len(lines) - 1; i >= 0 && !found; i-- {
			line := lines[i]
			// Find last < in this line that's not part of </ (closing tag)
			for j := len(line) - 1; j >= 0; j-- {
				if line[j] == '<' {
					// Check if it's a closing tag
					if j+1 < len(line) && line[j+1] == '/' {
						continue // Skip closing tags
					}
					// Found opening <, position cursor after it
					position = protocol.Position{
						Line:      uint32(i),
						Character: uint32(j + 1),
					}
					found = true
					break
				}
			}
		}

		if !found {
			t.Fatal("No opening < found in input content")
		}

		// Create completion params
		params := &protocol.CompletionParams{
			TextDocumentPositionParams: protocol.TextDocumentPositionParams{
				TextDocument: protocol.TextDocumentIdentifier{URI: htmlURI},
				Position:     position,
			},
			Context: &protocol.CompletionContext{
				TriggerCharacter: stringPtr("<"),
			},
		}

		// Request completions
		result, err := completion.Completion(ctx, &glsp.Context{}, params)
		if err != nil {
			t.Fatalf("Completion failed: %v", err)
		}

		actual, ok := result.([]protocol.CompletionItem)
		if !ok {
			t.Fatalf("Expected []CompletionItem, got %T", result)
		}

		// Load expected completions from golden file
		var expected []CompletionItemMatcher
		err = fixture.GetExpected("expected", &expected)
		if err != nil {
			t.Fatalf("Failed to load expected completions: %v", err)
		}

		// Verify completion count
		if len(actual) != len(expected) {
			t.Errorf("Expected %d completions, got %d", len(expected), len(actual))
			t.Logf("Expected labels: %v", getMatcherLabels(expected))
			t.Logf("Actual labels: %v", testhelpers.GetCompletionLabels(actual))
		}

		// Verify each expected completion is present (by label only, order-independent)
		expectedLabels := make(map[string]bool)
		for _, exp := range expected {
			expectedLabels[exp.Label] = true
		}

		actualLabels := make(map[string]bool)
		for _, act := range actual {
			actualLabels[act.Label] = true
			if !expectedLabels[act.Label] {
				t.Errorf("Unexpected completion: %s", act.Label)
			}
		}

		for _, exp := range expected {
			if !actualLabels[exp.Label] {
				t.Errorf("Missing expected completion: %s", exp.Label)
			}
		}
	})
}

// CompletionItemMatcher represents expected completion fields for assertions
type CompletionItemMatcher struct {
	Label string `json:"label"`
}

// Helper to extract labels from matchers
func getMatcherLabels(items []CompletionItemMatcher) []string {
	labels := make([]string, len(items))
	for i, item := range items {
		labels[i] = item.Label
	}
	return labels
}
