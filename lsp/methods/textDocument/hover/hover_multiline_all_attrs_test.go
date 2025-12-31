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

	"bennypowers.dev/cem/lsp/document"
	"bennypowers.dev/cem/lsp/methods/textDocument/hover"
	"bennypowers.dev/cem/lsp/testhelpers"
	M "bennypowers.dev/cem/manifest"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// TestHover_MultilineAllAttributes tests hover on all attributes that failed in the benchmark.
// This is a regression test for issues #176 and #179, ensuring our attribute value matching fix
// works for ALL attributes on a multi-line element.
//
// Benchmark failures (on docs/lsp/benchmarks branch) showed these 5 attributes all failed:
// - variant, size, disabled, loading, icon
//
// Root cause: Benchmark test positions were outdated, but this test validates the fix works.
func TestHover_MultilineAllAttributes(t *testing.T) {
	fixtureDir := "testdata/multiline-all-attributes-regression"

	// Read input HTML
	inputHTML, err := os.ReadFile(filepath.Join(fixtureDir, "input.html"))
	if err != nil {
		t.Fatalf("Failed to read input.html: %v", err)
	}

	// Read and parse manifest
	manifestBytes, err := os.ReadFile(filepath.Join(fixtureDir, "manifest.json"))
	if err != nil {
		t.Fatalf("Failed to read manifest.json: %v", err)
	}

	var pkg M.Package
	if err := json.Unmarshal(manifestBytes, &pkg); err != nil {
		t.Fatalf("Failed to parse manifest: %v", err)
	}

	// Setup context
	ctx := testhelpers.NewMockServerContext()
	ctx.AddManifest(&pkg)

	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)

	uri := "file:///test.html"
	doc := dm.OpenDocument(uri, string(inputHTML), 1)
	ctx.AddDocument(uri, doc)

	// Test each attribute that failed in the benchmark
	testCases := []struct {
		name     string
		position protocol.Position
	}{
		{"variant", protocol.Position{Line: 6, Character: 13}},
		{"size", protocol.Position{Line: 7, Character: 13}},
		{"disabled", protocol.Position{Line: 8, Character: 13}},
		{"loading", protocol.Position{Line: 9, Character: 13}},
		{"icon", protocol.Position{Line: 10, Character: 13}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Load expected result
			expectedFile := filepath.Join(fixtureDir, "expected-"+tc.name+".json")
			expectedBytes, err := os.ReadFile(expectedFile)
			if err != nil {
				t.Fatalf("Failed to read %s: %v", expectedFile, err)
			}

			var expected protocol.Hover
			if err := json.Unmarshal(expectedBytes, &expected); err != nil {
				t.Fatalf("Failed to parse %s: %v", expectedFile, err)
			}

			// Call hover
			params := &protocol.HoverParams{
				TextDocumentPositionParams: protocol.TextDocumentPositionParams{
					TextDocument: protocol.TextDocumentIdentifier{URI: uri},
					Position:     tc.position,
				},
			}

			result, err := hover.Hover(ctx, nil, params)
			if err != nil {
				t.Fatalf("Hover failed: %v", err)
			}

			if result == nil {
				t.Fatal("Expected hover result, got nil")
			}

			// Structured comparison
			actualContents, ok := result.Contents.(protocol.MarkupContent)
			if !ok {
				t.Fatalf("Expected Contents to be MarkupContent, got %T", result.Contents)
			}

			expectedContents, ok := expected.Contents.(protocol.MarkupContent)
			if !ok {
				t.Fatalf("Expected Contents in expected.json to be MarkupContent, got %T", expected.Contents)
			}

			if actualContents.Kind != expectedContents.Kind {
				t.Errorf("Expected kind %s, got %s", expectedContents.Kind, actualContents.Kind)
			}

			if actualContents.Value != expectedContents.Value {
				t.Errorf("Expected value:\n%s\n\nGot:\n%s", expectedContents.Value, actualContents.Value)
			}

			// Verify range
			if expected.Range != nil && result.Range != nil {
				if result.Range.Start.Line != expected.Range.Start.Line ||
					result.Range.Start.Character != expected.Range.Start.Character ||
					result.Range.End.Line != expected.Range.End.Line ||
					result.Range.End.Character != expected.Range.End.Character {
					t.Errorf("Range mismatch.\nExpected: %+v\nGot: %+v", expected.Range, result.Range)
				}
			}
		})
	}
}
