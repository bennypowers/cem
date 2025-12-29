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
package publishDiagnostics_test

import (
	"encoding/json"
	"testing"

	"bennypowers.dev/cem/internal/platform/testutil"
	"bennypowers.dev/cem/lsp/document"
	"bennypowers.dev/cem/lsp/methods/textDocument/publishDiagnostics"
	"bennypowers.dev/cem/lsp/testhelpers"
	M "bennypowers.dev/cem/manifest"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// ManifestFixture represents simplified manifest data for tests
type ManifestFixture struct {
	Attributes map[string]map[string]M.Attribute `json:"attributes"`
}

// TestAttributeDiagnostics_Fixtures runs all attribute diagnostics tests from fixtures
func TestAttributeDiagnostics_Fixtures(t *testing.T) {
	testutil.RunLSPFixtures(t, "testdata/attribute-diagnostics", func(t *testing.T, fixture *testutil.LSPFixture) {
		// Create mock server context
		ctx := testhelpers.NewMockServerContext()

		// Create DocumentManager and document
		dm, err := document.NewDocumentManager()
		if err != nil {
			t.Fatalf("Failed to create DocumentManager: %v", err)
		}
		defer dm.Close()
		ctx.SetDocumentManager(dm)

		// Open document with fixture content
		doc := dm.OpenDocument("test.html", fixture.InputHTML, 1)
		ctx.AddDocument("test.html", doc)

		// Load manifest if present
		if len(fixture.Manifest) > 0 {
			var manifestData ManifestFixture
			if err := json.Unmarshal(fixture.Manifest, &manifestData); err != nil {
				t.Fatalf("Failed to parse manifest.json: %v", err)
			}

			// Add attributes to context
			for tagName, attrs := range manifestData.Attributes {
				attrMap := make(map[string]*M.Attribute)
				for attrName, attr := range attrs {
					attr.Name = attrName
					attrCopy := attr
					attrMap[attrName] = &attrCopy
				}
				ctx.AddAttributes(tagName, attrMap)
			}
		}

		// Run diagnostics analysis
		diagnostics, err := publishDiagnostics.AnalyzeAttributeDiagnosticsForTest(ctx, doc)
		if err != nil {
			t.Fatalf("Failed to analyze attributes: %v", err)
		}

		// Load expected Diagnostic[] from expected.json
		var expected []protocol.Diagnostic
		if err := fixture.GetExpected("expected", &expected); err != nil {
			t.Fatalf("Failed to load expected diagnostics: %v", err)
		}

		// Verify diagnostic count
		if len(diagnostics) != len(expected) {
			t.Errorf("Expected %d diagnostics, got %d", len(expected), len(diagnostics))
			for i, diag := range diagnostics {
				t.Errorf("  Diagnostic %d: %s", i, diag.Message)
			}
			return
		}

		// Verify each diagnostic matches expected
		for i, exp := range expected {
			if i >= len(diagnostics) {
				break
			}
			act := diagnostics[i]

			// Verify message
			if exp.Message != act.Message {
				t.Errorf("Diagnostic %d message mismatch:\n  expected: %s\n  got: %s", i, exp.Message, act.Message)
			}

			// Verify range
			if exp.Range.Start.Line != act.Range.Start.Line || exp.Range.Start.Character != act.Range.Start.Character {
				t.Errorf("Diagnostic %d range start mismatch:\n  expected: line %d, char %d\n  got: line %d, char %d",
					i, exp.Range.Start.Line, exp.Range.Start.Character, act.Range.Start.Line, act.Range.Start.Character)
			}

			// Verify severity (compare values, not pointers)
			if (exp.Severity == nil) != (act.Severity == nil) {
				t.Errorf("Diagnostic %d severity nil mismatch:\n  expected nil: %v\n  got nil: %v",
					i, exp.Severity == nil, act.Severity == nil)
			} else if exp.Severity != nil && act.Severity != nil && *exp.Severity != *act.Severity {
				t.Errorf("Diagnostic %d severity mismatch:\n  expected: %d\n  got: %d", i, *exp.Severity, *act.Severity)
			}
		}
	})
}
