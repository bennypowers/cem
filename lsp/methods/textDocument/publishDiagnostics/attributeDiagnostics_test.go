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
	"testing"

	"bennypowers.dev/cem/lsp/document"
	"bennypowers.dev/cem/lsp/methods/textDocument/publishDiagnostics"
	"bennypowers.dev/cem/lsp/testhelpers"
	M "bennypowers.dev/cem/manifest"
)

func TestAttributeDiagnostics_GlobalAttributes(t *testing.T) {
	ctx := testhelpers.NewMockServerContext()
	content := `<div class="test" id="main" data-value="42">Hello</div>`
	// Create DocumentManager and document
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)
	doc := dm.OpenDocument("test.html", content, 1)
	ctx.AddDocument("test.html", doc)

	diagnostics, err := publishDiagnostics.AnalyzeAttributeDiagnosticsForTest(ctx, doc)
	if err != nil {
		t.Fatalf("Failed to analyze attributes: %v", err)
	}

	// Should have no diagnostics for global attributes
	if len(diagnostics) != 0 {
		t.Errorf("Expected 0 diagnostics for global attributes, got %d", len(diagnostics))
	}
}

func TestAttributeDiagnostics_CustomElementValidAttribute(t *testing.T) {
	ctx := testhelpers.NewMockServerContext()
	content := `<my-element size="large" color="red">Content</my-element>`
	// Create DocumentManager and document
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)
	doc := dm.OpenDocument("test.html", content, 1)
	ctx.AddDocument("test.html", doc)

	// Set up attributes for my-element
	ctx.AddAttributes("my-element", map[string]*M.Attribute{
		"size":  {FullyQualified: M.FullyQualified{Name: "size"}},
		"color": {FullyQualified: M.FullyQualified{Name: "color"}},
	})

	diagnostics, err := publishDiagnostics.AnalyzeAttributeDiagnosticsForTest(ctx, doc)
	if err != nil {
		t.Fatalf("Failed to analyze attributes: %v", err)
	}

	// Should have no diagnostics for valid custom element attributes
	if len(diagnostics) != 0 {
		t.Errorf("Expected 0 diagnostics for valid custom element attributes, got %d", len(diagnostics))
	}
}

func TestAttributeDiagnostics_CustomElementInvalidAttribute(t *testing.T) {
	ctx := testhelpers.NewMockServerContext()
	content := `<my-element siz="large" colour="red">Content</my-element>`
	// Create DocumentManager and document
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)
	doc := dm.OpenDocument("test.html", content, 1)
	ctx.AddDocument("test.html", doc)

	// Set up attributes for my-element
	ctx.AddAttributes("my-element", map[string]*M.Attribute{
		"size":  {FullyQualified: M.FullyQualified{Name: "size"}},
		"color": {FullyQualified: M.FullyQualified{Name: "color"}},
	})

	diagnostics, err := publishDiagnostics.AnalyzeAttributeDiagnosticsForTest(ctx, doc)
	if err != nil {
		t.Fatalf("Failed to analyze attributes: %v", err)
	}

	// Should have 2 diagnostics for invalid attributes
	if len(diagnostics) != 2 {
		t.Errorf("Expected 2 diagnostics for invalid attributes, got %d", len(diagnostics))
		return
	}

	// Check first diagnostic (siz -> size)
	if diagnostics[0].Message != "Unknown attribute 'siz'. Did you mean 'size'?" {
		t.Errorf("Unexpected diagnostic message: %s", diagnostics[0].Message)
	}

	// Check second diagnostic (colour -> color)
	if diagnostics[1].Message != "Unknown attribute 'colour'. Did you mean 'color'?" {
		t.Errorf("Unexpected diagnostic message: %s", diagnostics[1].Message)
	}
}

func TestAttributeDiagnostics_ScriptTypeModule(t *testing.T) {
	ctx := testhelpers.NewMockServerContext()
	content := `<script type="module">import './my-element.js';</script>`
	// Create DocumentManager and document
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)
	doc := dm.OpenDocument("test.html", content, 1)
	ctx.AddDocument("test.html", doc)

	diagnostics, err := publishDiagnostics.AnalyzeAttributeDiagnosticsForTest(ctx, doc)
	if err != nil {
		t.Fatalf("Failed to analyze attributes: %v", err)
	}

	// Should have 0 diagnostics - script[type] is outside CEM scope
	if len(diagnostics) != 0 {
		t.Errorf("Expected 0 diagnostics for script[type] (outside CEM scope), got %d", len(diagnostics))
		for _, diag := range diagnostics {
			t.Errorf("Unexpected diagnostic: %s", diag.Message)
		}
	}
}

func TestAttributeDiagnostics_GlobalAttributeTypoOnCustomElement(t *testing.T) {
	ctx := testhelpers.NewMockServerContext()
	content := `<my-custom-element clas="test" titl="tooltip">Hello</my-custom-element>`
	// Create DocumentManager and document
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)
	doc := dm.OpenDocument("test.html", content, 1)
	ctx.AddDocument("test.html", doc)

	// Custom element with no defined attributes
	ctx.AddAttributes("my-custom-element", map[string]*M.Attribute{})

	diagnostics, err := publishDiagnostics.AnalyzeAttributeDiagnosticsForTest(ctx, doc)
	if err != nil {
		t.Fatalf("Failed to analyze attributes: %v", err)
	}

	// Should have 2 diagnostics for global attribute typos on custom elements
	if len(diagnostics) != 2 {
		t.Errorf("Expected 2 diagnostics for global attribute typos on custom element, got %d", len(diagnostics))
		for i, diag := range diagnostics {
			t.Errorf("Diagnostic %d: %s", i, diag.Message)
		}
		return
	}

	// Check diagnostic messages
	found := make(map[string]bool)
	for _, diag := range diagnostics {
		found[diag.Message] = true
	}

	expected := []string{
		"Unknown attribute 'clas'. Did you mean 'class'?",
		"Unknown attribute 'titl'. Did you mean 'title'?",
	}

	for _, exp := range expected {
		if !found[exp] {
			t.Errorf("Expected diagnostic message not found: %s", exp)
		}
	}
}

func TestAttributeDiagnostics_StandardElementIgnored(t *testing.T) {
	ctx := testhelpers.NewMockServerContext()
	content := `<div clas="test" titl="tooltip">Hello</div>`
	// Create DocumentManager and document
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)
	doc := dm.OpenDocument("test.html", content, 1)
	ctx.AddDocument("test.html", doc)

	diagnostics, err := publishDiagnostics.AnalyzeAttributeDiagnosticsForTest(ctx, doc)
	if err != nil {
		t.Fatalf("Failed to analyze attributes: %v", err)
	}

	// Should have 0 diagnostics - standard HTML elements are outside CEM scope (except slot)
	if len(diagnostics) != 0 {
		t.Errorf("Expected 0 diagnostics for standard HTML elements (outside CEM scope), got %d", len(diagnostics))
		for _, diag := range diagnostics {
			t.Errorf("Unexpected diagnostic: %s", diag.Message)
		}
	}
}
