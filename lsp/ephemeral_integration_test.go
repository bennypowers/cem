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
package lsp_test

import (
	"slices"
	"testing"

	"bennypowers.dev/cem/internal/platform/testutil"
	"bennypowers.dev/cem/lsp"
	"bennypowers.dev/cem/lsp/methods/textDocument/publishDiagnostics"
	M "bennypowers.dev/cem/manifest"
	W "bennypowers.dev/cem/workspace"
)

// newEphemeralTestServer creates a server and opens the fixture's input.ts,
// triggering ephemeral synthesis. Returns the server and the document URI.
func newEphemeralTestServer(t *testing.T, fixture *testutil.LSPFixture) (*lsp.Server, string) {
	t.Helper()

	workspace := W.NewFileSystemWorkspaceContext("/test")
	if err := workspace.Init(); err != nil {
		t.Fatalf("Failed to init workspace: %v", err)
	}

	server, err := lsp.NewServer(workspace, lsp.TransportStdio)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	dm, err := server.DocumentManager()
	if err != nil {
		t.Fatalf("Failed to get document manager: %v", err)
	}

	uri := "file:///test.ts"
	dm.OpenDocument(uri, fixture.InputContent, 1)

	return server, uri
}

func TestEphemeralIntegration(t *testing.T) {
	testutil.RunLSPFixtures(t, "testdata/ephemeral-integration", func(t *testing.T, fixture *testutil.LSPFixture) {
		switch fixture.Name {
		case "synthesis":
			testSynthesis(t, fixture)
		case "precedence":
			testPrecedence(t, fixture)
		case "skip-when-in-main":
			testSkipWhenInMain(t, fixture)
		case "cleanup-on-close":
			testCleanupOnClose(t, fixture)
		case "local-definition-decorator":
			testLocalDefinitionNoDiagnostic(t, fixture, "test-el")
		case "local-definition-define":
			testLocalDefinitionNoDiagnostic(t, fixture, "test-el")
		case "local-definition-unknown-sibling":
			testLocalDefinitionUnknownSibling(t, fixture)
		default:
			t.Fatalf("Unknown fixture: %s", fixture.Name)
		}
	})
}

// testSynthesis verifies that opening a TypeScript file defining a custom
// element via @customElement causes the element to appear in the server's
// registry fallback.
func testSynthesis(t *testing.T, fixture *testutil.LSPFixture) {
	server, uri := newEphemeralTestServer(t, fixture)
	defer func() { _ = server.Close() }()

	// Before synthesis, element should not exist
	_, exists := server.Element("test-greeting")
	if exists {
		t.Error("Element should not exist before synthesis")
	}

	server.SynthesizeEphemeralElements(uri)

	el, exists := server.Element("test-greeting")
	if !exists {
		t.Fatal("Expected test-greeting element to exist after synthesis")
	}
	if el.TagName != "test-greeting" {
		t.Errorf("Expected tag name 'test-greeting', got %q", el.TagName)
	}

	attrs, hasAttrs := server.Attributes("test-greeting")
	if !hasAttrs {
		t.Fatal("Expected attributes for test-greeting")
	}
	if attrs["name"] == nil {
		t.Error("Expected 'name' attribute")
	}

	tags := server.AllTagNames()
	if !slices.Contains(tags, "test-greeting") {
		t.Error("Expected AllTagNames to include 'test-greeting'")
	}

	decl := server.FindCustomElementDeclaration("test-greeting")
	if decl == nil {
		t.Fatal("Expected FindCustomElementDeclaration to return non-nil")
	}
	if decl.Summary == "" && decl.Description == "" {
		t.Error("Expected declaration to have summary or description from JSDoc")
	}

	def, hasDef := server.ElementDefinition("test-greeting")
	if !hasDef {
		t.Fatal("Expected ElementDefinition to return true")
	}
	if def.ModulePath() != uri {
		t.Errorf("Expected module path %q, got %q", uri, def.ModulePath())
	}
}

// testPrecedence verifies that ephemeral synthesis works when the element
// is not yet in the main registry.
func testPrecedence(t *testing.T, fixture *testutil.LSPFixture) {
	server, uri := newEphemeralTestServer(t, fixture)
	defer func() { _ = server.Close() }()

	server.SynthesizeEphemeralElements(uri)

	_, exists := server.Element("test-button")
	if !exists {
		t.Fatal("Expected test-button to exist via ephemeral fallback")
	}
}

// testSkipWhenInMain verifies that the ephemeral synthesis pipeline does not
// duplicate work for elements already in the main manifest registry.
func testSkipWhenInMain(t *testing.T, fixture *testutil.LSPFixture) {
	server, uri := newEphemeralTestServer(t, fixture)
	defer func() { _ = server.Close() }()

	// Add test-button to the main registry first
	server.AddManifest(&M.Package{
		SchemaVersion: "2.1.0",
		Modules: []M.Module{{
			Kind: "javascript-module",
			Path: "test-button.js",
			Declarations: []M.Declaration{
				&M.CustomElementDeclaration{
					ClassDeclaration: M.ClassDeclaration{
						ClassLike: M.ClassLike{
							FullyQualified: M.FullyQualified{
								Name:    "TestButton",
								Summary: "From main registry",
							},
						},
						Kind: "class",
					},
					CustomElement: M.CustomElement{
						TagName:       "test-button",
						CustomElement: true,
					},
				},
			},
		}},
	})

	server.SynthesizeEphemeralElements(uri)

	decl := server.FindCustomElementDeclaration("test-button")
	if decl == nil {
		t.Fatal("Expected to find declaration")
	}
	if decl.Summary != "From main registry" {
		t.Errorf("Expected main registry data ('From main registry'), got %q — ephemeral data overwrote main registry", decl.Summary)
	}
}

// testCleanupOnClose verifies that ephemeral data is removed when a
// document is closed.
func testCleanupOnClose(t *testing.T, fixture *testutil.LSPFixture) {
	server, uri := newEphemeralTestServer(t, fixture)
	defer func() { _ = server.Close() }()

	server.SynthesizeEphemeralElements(uri)

	_, exists := server.Element("temp-element")
	if !exists {
		t.Fatal("Expected temp-element to exist after synthesis")
	}

	dm, err := server.DocumentManager()
	if err != nil {
		t.Fatalf("Failed to get document manager: %v", err)
	}
	dm.CloseDocument(uri)

	// SynthesizeEphemeralElements on a closed doc cleans up immediately
	server.SynthesizeEphemeralElements(uri)

	_, exists = server.Element("temp-element")
	if exists {
		t.Error("Expected temp-element to be removed after document close")
	}
}

// testLocalDefinitionNoDiagnostic verifies that elements defined in the
// current file via @customElement or customElements.define produce zero
// "unknown element" diagnostics when used in that same file's templates.
func testLocalDefinitionNoDiagnostic(t *testing.T, fixture *testutil.LSPFixture, expectedTag string) {
	server, uri := newEphemeralTestServer(t, fixture)
	defer func() { _ = server.Close() }()

	// Synthesis must run before diagnostics (mirrors lifecycle order)
	server.SynthesizeEphemeralElements(uri)

	doc := server.Document(uri)
	if doc == nil {
		t.Fatal("Expected document to be open")
	}

	diagnostics := publishDiagnostics.AnalyzeTagNameDiagnosticsForTest(server, doc)
	if len(diagnostics) != 0 {
		t.Errorf("Expected 0 diagnostics for locally-defined element %q, got %d:", expectedTag, len(diagnostics))
		for i, d := range diagnostics {
			t.Errorf("  %d: %s", i, d.Message)
		}
	}
}

// testLocalDefinitionUnknownSibling verifies that an unknown element used
// alongside a locally-defined one still produces a diagnostic.
func testLocalDefinitionUnknownSibling(t *testing.T, fixture *testutil.LSPFixture) {
	server, uri := newEphemeralTestServer(t, fixture)
	defer func() { _ = server.Close() }()

	server.SynthesizeEphemeralElements(uri)

	doc := server.Document(uri)
	if doc == nil {
		t.Fatal("Expected document to be open")
	}

	diagnostics := publishDiagnostics.AnalyzeTagNameDiagnosticsForTest(server, doc)
	if len(diagnostics) != 1 {
		t.Errorf("Expected 1 diagnostic for unknown-el, got %d:", len(diagnostics))
		for i, d := range diagnostics {
			t.Errorf("  %d: %s", i, d.Message)
		}
		return
	}

	if diagnostics[0].Message == "" {
		t.Error("Expected diagnostic message for unknown-el")
	}
}
