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

	"bennypowers.dev/cem/lsp"
	"bennypowers.dev/cem/lsp/methods/textDocument/publishDiagnostics"
	"bennypowers.dev/cem/lsp/testhelpers"
	M "bennypowers.dev/cem/manifest"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TestAttributeValueDiagnostics_BooleanAttributes(t *testing.T) {
	ctx := testhelpers.NewMockServerContext()
	content := `<my-element disabled="false" hidden="true" readonly>Content</my-element>`

	// Create DocumentManager and document
	dm, err := lsp.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)
	doc := dm.OpenDocument("test.html", content, 1)
	ctx.AddDocument("test.html", doc)

	// Set up boolean attributes for my-element
	ctx.AddAttributes("my-element", map[string]*M.Attribute{
		"disabled": {
			FullyQualified: M.FullyQualified{Name: "disabled"},
			Type:           &M.Type{Text: "boolean"},
		},
		"hidden": {
			FullyQualified: M.FullyQualified{Name: "hidden"},
			Type:           &M.Type{Text: "boolean"},
		},
		"readonly": {
			FullyQualified: M.FullyQualified{Name: "readonly"},
			Type:           &M.Type{Text: "boolean"},
		},
	})

	diagnostics := publishDiagnostics.AnalyzeAttributeValueDiagnosticsForTest(ctx, doc)

	// Should have 2 diagnostics: disabled="false" (warning) and hidden="true" (info)
	if len(diagnostics) != 2 {
		t.Errorf("Expected 2 diagnostics for boolean attributes, got %d", len(diagnostics))
		for i, diag := range diagnostics {
			t.Errorf("Diagnostic %d: %s", i, diag.Message)
		}
		return
	}

	// Check disabled="false" diagnostic (should be warning)
	found := false
	for _, diag := range diagnostics {
		if diag.Message == "Boolean attribute 'disabled' with value 'false' is still true. Remove the attribute entirely to make it false." {
			if *diag.Severity != protocol.DiagnosticSeverityWarning {
				t.Errorf("Expected warning severity for disabled='false', got %v", *diag.Severity)
			}
			found = true
			break
		}
	}
	if !found {
		t.Error("Expected diagnostic for disabled='false' not found")
	}

	// Check hidden="true" diagnostic (should be info)
	found = false
	for _, diag := range diagnostics {
		if diag.Message == "Boolean attribute 'hidden' with value 'true' is redundant. Use <my-element hidden> instead." {
			if *diag.Severity != protocol.DiagnosticSeverityInformation {
				t.Errorf("Expected info severity for hidden='true', got %v", *diag.Severity)
			}
			found = true
			break
		}
	}
	if !found {
		t.Error("Expected diagnostic for hidden='true' not found")
	}
}

func TestAttributeValueDiagnostics_UnionTypes(t *testing.T) {
	ctx := testhelpers.NewMockServerContext()
	content := `<my-element size="smll" theme="Dark" variant="invalid">Content</my-element>`

	// Create DocumentManager and document
	dm, err := lsp.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)
	doc := dm.OpenDocument("test.html", content, 1)
	ctx.AddDocument("test.html", doc)

	// Set up union type attributes for my-element
	ctx.AddAttributes("my-element", map[string]*M.Attribute{
		"size": {
			FullyQualified: M.FullyQualified{Name: "size"},
			Type:           &M.Type{Text: `"small" | "medium" | "large"`},
		},
		"theme": {
			FullyQualified: M.FullyQualified{Name: "theme"},
			Type:           &M.Type{Text: `"light" | "dark" | "auto"`},
		},
		"variant": {
			FullyQualified: M.FullyQualified{Name: "variant"},
			Type:           &M.Type{Text: `"primary" | "secondary"`},
		},
	})

	diagnostics := publishDiagnostics.AnalyzeAttributeValueDiagnosticsForTest(ctx, doc)

	// Should have 3 diagnostics for invalid union values
	if len(diagnostics) != 3 {
		t.Errorf("Expected 3 diagnostics for union types, got %d", len(diagnostics))
		for i, diag := range diagnostics {
			t.Errorf("Diagnostic %d: %s", i, diag.Message)
		}
		return
	}

	// Check that we get suggestions for close matches
	expectedMessages := map[string]bool{
		"Expected one of: 'small', 'medium' or 'large' for attribute 'size', got 'smll'. Did you mean 'small'?": false,
		"Expected one of: 'light', 'dark' or 'auto' for attribute 'theme', got 'Dark'. Did you mean 'dark'?":      false,
		"Expected one of: 'primary' or 'secondary' for attribute 'variant', got 'invalid'":                       false,
	}

	for _, diag := range diagnostics {
		if _, exists := expectedMessages[diag.Message]; exists {
			expectedMessages[diag.Message] = true
		} else {
			t.Errorf("Unexpected diagnostic message: %s", diag.Message)
		}
	}

	for msg, found := range expectedMessages {
		if !found {
			t.Errorf("Expected diagnostic message not found: %s", msg)
		}
	}
}

func TestAttributeValueDiagnostics_LiteralTypes(t *testing.T) {
	ctx := testhelpers.NewMockServerContext()
	content := `<my-element role="Primary" status="Inactive">Content</my-element>`

	// Create DocumentManager and document
	dm, err := lsp.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)
	doc := dm.OpenDocument("test.html", content, 1)
	ctx.AddDocument("test.html", doc)

	// Set up literal type attributes for my-element
	ctx.AddAttributes("my-element", map[string]*M.Attribute{
		"role": {
			FullyQualified: M.FullyQualified{Name: "role"},
			Type:           &M.Type{Text: `"primary"`},
		},
		"status": {
			FullyQualified: M.FullyQualified{Name: "status"},
			Type:           &M.Type{Text: `'active'`},
		},
	})

	diagnostics := publishDiagnostics.AnalyzeAttributeValueDiagnosticsForTest(ctx, doc)

	// Should have 2 diagnostics for case mismatches
	if len(diagnostics) != 2 {
		t.Errorf("Expected 2 diagnostics for literal types, got %d", len(diagnostics))
		for i, diag := range diagnostics {
			t.Errorf("Diagnostic %d: %s", i, diag.Message)
		}
		return
	}

	// Check for case mismatch diagnostics
	expectedMessages := []string{
		"Expected literal value 'primary', got 'Primary' (case mismatch)",
		"Expected literal value 'active' for attribute 'status', got 'Inactive'",
	}

	foundMessages := make(map[string]bool)
	for _, diag := range diagnostics {
		foundMessages[diag.Message] = true
	}

	for _, expected := range expectedMessages {
		if !foundMessages[expected] {
			t.Errorf("Expected diagnostic message not found: %s", expected)
		}
	}
}

func TestAttributeValueDiagnostics_NumberTypes(t *testing.T) {
	ctx := testhelpers.NewMockServerContext()
	content := `<my-element count="abc" max="42" min="">Content</my-element>`

	// Create DocumentManager and document
	dm, err := lsp.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)
	doc := dm.OpenDocument("test.html", content, 1)
	ctx.AddDocument("test.html", doc)

	// Set up number type attributes for my-element
	ctx.AddAttributes("my-element", map[string]*M.Attribute{
		"count": {
			FullyQualified: M.FullyQualified{Name: "count"},
			Type:           &M.Type{Text: "number"},
		},
		"max": {
			FullyQualified: M.FullyQualified{Name: "max"},
			Type:           &M.Type{Text: "number"},
		},
		"min": {
			FullyQualified: M.FullyQualified{Name: "min"},
			Type:           &M.Type{Text: "number"},
		},
	})

	diagnostics := publishDiagnostics.AnalyzeAttributeValueDiagnosticsForTest(ctx, doc)

	// Should have 2 diagnostics: count="abc" (invalid number) and min="" (empty value)
	if len(diagnostics) != 2 {
		t.Errorf("Expected 2 diagnostics for number types, got %d", len(diagnostics))
		for i, diag := range diagnostics {
			t.Errorf("Diagnostic %d: %s", i, diag.Message)
		}
		return
	}

	expectedMessages := []string{
		"Expected number for attribute 'count', got 'abc'",
		"Number attribute 'min' requires a numeric value",
	}

	foundMessages := make(map[string]bool)
	for _, diag := range diagnostics {
		foundMessages[diag.Message] = true
	}

	for _, expected := range expectedMessages {
		if !foundMessages[expected] {
			t.Errorf("Expected diagnostic message not found: %s", expected)
		}
	}
}

func TestAttributeValueDiagnostics_ArrayTypes(t *testing.T) {
	ctx := testhelpers.NewMockServerContext()
	content := `<my-element items='["a","b"]' tags="tag1,tag2">Content</my-element>`

	// Create DocumentManager and document
	dm, err := lsp.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)
	doc := dm.OpenDocument("test.html", content, 1)
	ctx.AddDocument("test.html", doc)

	// Set up array type attributes for my-element
	ctx.AddAttributes("my-element", map[string]*M.Attribute{
		"items": {
			FullyQualified: M.FullyQualified{Name: "items"},
			Type:           &M.Type{Text: "string[]"},
		},
		"tags": {
			FullyQualified: M.FullyQualified{Name: "tags"},
			Type:           &M.Type{Text: "Array<string>"},
		},
	})

	diagnostics := publishDiagnostics.AnalyzeAttributeValueDiagnosticsForTest(ctx, doc)

	// Should have 2 informational diagnostics for array types
	if len(diagnostics) != 2 {
		t.Errorf("Expected 2 diagnostics for array types, got %d", len(diagnostics))
		for i, diag := range diagnostics {
			t.Errorf("Diagnostic %d: %s", i, diag.Message)
		}
		return
	}

	expectedMessage := "Array attributes support multiple formats (JSON, comma-separated, space-separated). Refer to component documentation."
	for _, diag := range diagnostics {
		if diag.Message != expectedMessage {
			t.Errorf("Expected array info message, got: %s", diag.Message)
		}
		if *diag.Severity != protocol.DiagnosticSeverityInformation {
			t.Errorf("Expected info severity for array types, got %v", *diag.Severity)
		}
	}
}

func TestAttributeValueDiagnostics_StringTypes(t *testing.T) {
	ctx := testhelpers.NewMockServerContext()
	content := `<my-element label="Hello World" description="">Content</my-element>`

	// Create DocumentManager and document
	dm, err := lsp.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)
	doc := dm.OpenDocument("test.html", content, 1)
	ctx.AddDocument("test.html", doc)

	// Set up string type attributes for my-element
	ctx.AddAttributes("my-element", map[string]*M.Attribute{
		"label": {
			FullyQualified: M.FullyQualified{Name: "label"},
			Type:           &M.Type{Text: "string"},
		},
		"description": {
			FullyQualified: M.FullyQualified{Name: "description"},
			Type:           &M.Type{Text: "string"},
		},
	})

	diagnostics := publishDiagnostics.AnalyzeAttributeValueDiagnosticsForTest(ctx, doc)

	// Should have no diagnostics for valid string attributes
	if len(diagnostics) != 0 {
		t.Errorf("Expected 0 diagnostics for string types, got %d", len(diagnostics))
		for i, diag := range diagnostics {
			t.Errorf("Diagnostic %d: %s", i, diag.Message)
		}
	}
}

func TestAttributeValueDiagnostics_NoTypeInformation(t *testing.T) {
	ctx := testhelpers.NewMockServerContext()
	content := `<my-element custom="value">Content</my-element>`

	// Create DocumentManager and document
	dm, err := lsp.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)
	doc := dm.OpenDocument("test.html", content, 1)
	ctx.AddDocument("test.html", doc)

	// Set up attribute without type information
	ctx.AddAttributes("my-element", map[string]*M.Attribute{
		"custom": {
			FullyQualified: M.FullyQualified{Name: "custom"},
			Type:           nil, // No type information
		},
	})

	diagnostics := publishDiagnostics.AnalyzeAttributeValueDiagnosticsForTest(ctx, doc)

	// Should have no diagnostics when no type information is available
	if len(diagnostics) != 0 {
		t.Errorf("Expected 0 diagnostics for attributes without type information, got %d", len(diagnostics))
		for i, diag := range diagnostics {
			t.Errorf("Diagnostic %d: %s", i, diag.Message)
		}
	}
}

func TestAttributeValueDiagnostics_StandardHTMLElements(t *testing.T) {
	ctx := testhelpers.NewMockServerContext()
	content := `<div invalid-attr="value">Content</div>`

	// Create DocumentManager and document
	dm, err := lsp.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)
	doc := dm.OpenDocument("test.html", content, 1)
	ctx.AddDocument("test.html", doc)

	diagnostics := publishDiagnostics.AnalyzeAttributeValueDiagnosticsForTest(ctx, doc)

	// Should have no diagnostics for standard HTML elements (outside CEM scope)
	if len(diagnostics) != 0 {
		t.Errorf("Expected 0 diagnostics for standard HTML elements, got %d", len(diagnostics))
		for i, diag := range diagnostics {
			t.Errorf("Diagnostic %d: %s", i, diag.Message)
		}
	}
}

func TestAttributeValueDiagnostics_EdgeCases(t *testing.T) {
	ctx := testhelpers.NewMockServerContext()
	content := `<my-element unquoted=value quoted="quoted value" empty="">Content</my-element>`

	// Create DocumentManager and document
	dm, err := lsp.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)
	doc := dm.OpenDocument("test.html", content, 1)
	ctx.AddDocument("test.html", doc)

	// Set up various attribute types
	ctx.AddAttributes("my-element", map[string]*M.Attribute{
		"unquoted": {
			FullyQualified: M.FullyQualified{Name: "unquoted"},
			Type:           &M.Type{Text: "string"},
		},
		"quoted": {
			FullyQualified: M.FullyQualified{Name: "quoted"},
			Type:           &M.Type{Text: "string"},
		},
		"empty": {
			FullyQualified: M.FullyQualified{Name: "empty"},
			Type:           &M.Type{Text: "string"},
		},
	})

	diagnostics := publishDiagnostics.AnalyzeAttributeValueDiagnosticsForTest(ctx, doc)

	// Should have no diagnostics for edge case string values
	if len(diagnostics) != 0 {
		t.Errorf("Expected 0 diagnostics for edge case values, got %d", len(diagnostics))
		for i, diag := range diagnostics {
			t.Errorf("Diagnostic %d: %s", i, diag.Message)
		}
	}
}