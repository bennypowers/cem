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
package hover_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/internal/platform/testutil"
	"bennypowers.dev/cem/lsp/document"
	"bennypowers.dev/cem/lsp/methods/textDocument/hover"
	"bennypowers.dev/cem/lsp/testhelpers"
	M "bennypowers.dev/cem/manifest"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// TestHover_MultilineAllAttributes tests hover on all attributes that failed in the benchmark
// This test verifies that our fix for attribute value matching (issue #176, #179) works for ALL
// attributes on a multi-line element, not just a single one.
//
// Benchmark failures showed these attributes all failed with "No result":
// - variant (line 6, char 13)
// - size (line 7, char 13)
// - disabled (line 8, char 13)
// - loading (line 9, char 13)
// - icon (line 10, char 13)
func TestHover_MultilineAllAttributes(t *testing.T) {
	// Load fixture data manually
	fixtureDir := "testdata/multiline-all-attributes"

	// Read input HTML
	inputHTML, err := os.ReadFile(filepath.Join(fixtureDir, "input.html"))
	if err != nil {
		t.Fatalf("Failed to read input.html: %v", err)
	}

	// Read manifest
	manifestBytes, err := os.ReadFile(filepath.Join(fixtureDir, "manifest.json"))
	if err != nil {
		t.Fatalf("Failed to read manifest.json: %v", err)
	}

	// Create fixture-like structure for expected data loading
	fixture := &testutil.LSPFixture{
		Name:         "multiline-all-attributes",
		InputContent: string(inputHTML),
		InputType:    "html",
		Manifest:     json.RawMessage(manifestBytes),
		ExpectedMap:  make(map[string]any),
	}

	// Load all expected-*.json files
	entries, err := os.ReadDir(fixtureDir)
	if err == nil {
		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}
			name := entry.Name()
			if filepath.Ext(name) == ".json" && filepath.Base(name) != "manifest.json" {
				fullPath := filepath.Join(fixtureDir, name)
				data, err := os.ReadFile(fullPath)
				if err != nil {
					continue
				}
				var expectedData any
				if err := json.Unmarshal(data, &expectedData); err != nil {
					continue
				}
				// Handle "expected-*.json" pattern
				if len(name) > len("expected-") && name[:len("expected-")] == "expected-" {
					key := name[len("expected-") : len(name)-len(".json")]
					fixture.ExpectedMap[key] = expectedData
				}
			}
		}
	}

	// Parse manifest
	ctx := testhelpers.NewMockServerContext()
	if fixture.Manifest != nil {
		var pkg M.Package
		err := json.Unmarshal(fixture.Manifest, &pkg)
		if err != nil {
			t.Fatalf("Failed to parse manifest: %v", err)
		}
		ctx.AddManifest(&pkg)
	}

	// Create document manager
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)

	uri := "file:///test.html"
	doc := dm.OpenDocument(uri, fixture.InputContent, 1)
	ctx.AddDocument(uri, doc)

	// Test each attribute at its cursor position
	testCases := []struct {
		name            string
		position        protocol.Position
		expectedFileKey string
	}{
		{"variant", protocol.Position{Line: 6, Character: 13}, "variant"},
		{"size", protocol.Position{Line: 7, Character: 13}, "size"},
		{"disabled", protocol.Position{Line: 8, Character: 13}, "disabled"},
		{"loading", protocol.Position{Line: 9, Character: 13}, "loading"},
		{"icon", protocol.Position{Line: 10, Character: 13}, "icon"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Load expected hover response for this attribute
			var expected protocol.Hover
			err := fixture.GetExpected(tc.expectedFileKey, &expected)
			if err != nil {
				t.Fatalf("Failed to get expected hover for %s: %v", tc.name, err)
			}

			// Create hover params
			params := &protocol.HoverParams{
				TextDocumentPositionParams: protocol.TextDocumentPositionParams{
					TextDocument: protocol.TextDocumentIdentifier{URI: uri},
					Position:     tc.position,
				},
			}

			// Call hover
			result, err := hover.Hover(ctx, nil, params)
			if err != nil {
				t.Fatalf("Hover failed for %s: %v", tc.name, err)
			}

			if result == nil {
				t.Fatalf("Expected hover result for %s, got nil", tc.name)
			}

			// Verify result matches expected
			resultJSON, _ := json.MarshalIndent(result, "", "  ")
			expectedJSON, _ := json.MarshalIndent(expected, "", "  ")

			if string(resultJSON) != string(expectedJSON) {
				t.Errorf("Hover result for %s does not match expected.\n\nGot:\n%s\n\nExpected:\n%s",
					tc.name, string(resultJSON), string(expectedJSON))
			}
		})
	}
}
