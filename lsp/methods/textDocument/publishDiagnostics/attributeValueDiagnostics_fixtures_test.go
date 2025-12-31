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

// TestAttributeValueDiagnostics_MultilineBooleanAttrs tests issue #179
func TestAttributeValueDiagnostics_MultilineBooleanAttrs(t *testing.T) {
	testutil.RunLSPFixtures(t, "testdata/multiline-boolean-attrs-issue-179", func(t *testing.T, fixture *testutil.LSPFixture) {
		runBooleanParsingTest(t, fixture, 179)
	})
}

// TestAttributeValueDiagnostics_BooleanBeforeValued tests issue #176
func TestAttributeValueDiagnostics_BooleanBeforeValued(t *testing.T) {
	testutil.RunLSPFixtures(t, "testdata/boolean-before-valued-issue-176", func(t *testing.T, fixture *testutil.LSPFixture) {
		runBooleanParsingTest(t, fixture, 176)
	})
}

func runBooleanParsingTest(t *testing.T, fixture *testutil.LSPFixture, issueNumber int) {
	// Create server context
	ctx := testhelpers.NewMockServerContext()

	// Parse and add manifest if present
	if fixture.Manifest != nil {
		var pkg M.Package
		err := json.Unmarshal(fixture.Manifest, &pkg)
		if err != nil {
			t.Fatalf("Failed to parse manifest: %v", err)
		}
		ctx.AddManifest(&pkg)
	}

	// Create DocumentManager and document
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)

	uri := "file:///test.html"
	doc := dm.OpenDocument(uri, fixture.InputContent, 1)
	ctx.AddDocument(uri, doc)

	// Analyze diagnostics
	diagnostics := publishDiagnostics.AnalyzeAttributeValueDiagnosticsForTest(ctx, doc)

	// Load expected diagnostics
	var expectedDiagnostics []protocol.Diagnostic
	err = fixture.GetExpected("expected", &expectedDiagnostics)
	if err != nil {
		t.Fatalf("Failed to get expected diagnostics: %v", err)
	}

	// Check for specific regression errors based on issue number
	if issueNumber == 179 {
		for _, diag := range diagnostics {
			if diag.Message == "Boolean attribute 'disabled' should not have value 'star'. Use <my-button disabled> instead." {
				t.Errorf("REGRESSION #179: disabled attribute incorrectly detected as having value 'star'")
			}
			if diag.Message == "Boolean attribute 'loading' should not have value 'star'. Use <my-button loading> instead." {
				t.Errorf("REGRESSION #179: loading attribute incorrectly detected as having value 'star'")
			}
		}
	}

	if issueNumber == 176 {
		for _, diag := range diagnostics {
			if diag.Message == "Boolean attribute 'elevated' should not have value '16'. Use <mdw-card elevated> instead." {
				t.Errorf("REGRESSION #176: elevated attribute incorrectly detected as having value '16'")
			}
		}
	}

	// Compare diagnostic counts
	if len(diagnostics) != len(expectedDiagnostics) {
		t.Errorf("Expected %d diagnostics, got %d:", len(expectedDiagnostics), len(diagnostics))
		for i, diag := range diagnostics {
			t.Errorf("  Diagnostic %d: %s", i, diag.Message)
		}
	}
}
