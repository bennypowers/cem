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

func makeTestManifest() *M.Package {
	return &M.Package{
		SchemaVersion: "2.1.0",
		Modules: []M.Module{{
			Kind: "javascript-module",
			Path: "elements/my-button/my-button.js",
			Declarations: []M.Declaration{
				&M.CustomElementDeclaration{
					ClassDeclaration: M.ClassDeclaration{
						ClassLike: M.ClassLike{
							FullyQualified: M.FullyQualified{Name: "MyButton"},
						},
					},
					CustomElement: M.CustomElement{
						TagName:       "my-button",
						CustomElement: true,
						Attributes: []M.Attribute{
							{
								FullyQualified: M.FullyQualified{
									Name:        "variant",
									Description: "The button's theme variant. If you set `variant` to 'icon', you MUST also set `accessible-label`.",
								},
								Type: &M.Type{Text: "'primary' | 'secondary' | 'icon'"},
							},
							{
								FullyQualified: M.FullyQualified{
									Name:        "accessible-label",
									Description: "Accessible label for the button.",
								},
								Type: &M.Type{Text: "string"},
							},
							{
								FullyQualified: M.FullyQualified{
									Name:        "icon",
									Description: "The icon name. You MUST also set `accessible-label`.",
								},
								Type: &M.Type{Text: "string"},
							},
							{
								FullyQualified: M.FullyQualified{
									Name:        "loading",
									Description: "Shows a loading spinner. When `loading` is set, you MUST provide `accessible-label`.",
								},
								Type: &M.Type{Text: "boolean"},
							},
							{
								FullyQualified: M.FullyQualified{
									Name:        "compact",
									Description: "Compact mode.",
								},
								Type: &M.Type{Text: "boolean"},
							},
						},
					},
				},
			},
		}},
	}
}

func TestRequirementDiagnostics_ViolationWhenAttrValueMissing(t *testing.T) {
	ctx := testhelpers.NewMockServerContext()

	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)

	ctx.AddManifest(makeTestManifest())

	// variant="icon" without accessible-label -> should produce diagnostic
	html := `<my-button variant="icon" icon="doggy"></my-button>`
	doc := dm.OpenDocument("test.html", html, 1)
	ctx.AddDocument("test.html", doc)

	diagnostics := publishDiagnostics.AnalyzeRequirementDiagnosticsForTest(ctx, doc)

	// We expect violations:
	// 1. variant="icon" requires accessible-label (from variant's description)
	// 2. icon is set, requires accessible-label (from icon's description, unconditional)
	if len(diagnostics) < 1 {
		t.Fatalf("expected at least 1 diagnostic, got %d", len(diagnostics))
	}

	foundVariantViolation := false
	for _, d := range diagnostics {
		t.Logf("diagnostic: %s", d.Message)
		if d.Severity != nil && *d.Severity == 1 { // Error
			foundVariantViolation = true
		}
	}

	if !foundVariantViolation {
		t.Error("expected an error-severity diagnostic for missing accessible-label")
	}
}

func TestRequirementDiagnostics_NoViolationWhenSatisfied(t *testing.T) {
	ctx := testhelpers.NewMockServerContext()

	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)

	ctx.AddManifest(makeTestManifest())

	// variant="icon" with accessible-label -> no diagnostic
	html := `<my-button variant="icon" icon="doggy" accessible-label="Dog icon"></my-button>`
	doc := dm.OpenDocument("test.html", html, 1)
	ctx.AddDocument("test.html", doc)

	diagnostics := publishDiagnostics.AnalyzeRequirementDiagnosticsForTest(ctx, doc)

	if len(diagnostics) != 0 {
		for _, d := range diagnostics {
			t.Errorf("unexpected diagnostic: %s", d.Message)
		}
	}
}

func TestRequirementDiagnostics_NoViolationWhenConditionNotMet(t *testing.T) {
	ctx := testhelpers.NewMockServerContext()

	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)

	ctx.AddManifest(makeTestManifest())

	// variant="primary" (not "icon") -> no requirement for accessible-label
	html := `<my-button variant="primary"></my-button>`
	doc := dm.OpenDocument("test.html", html, 1)
	ctx.AddDocument("test.html", doc)

	diagnostics := publishDiagnostics.AnalyzeRequirementDiagnosticsForTest(ctx, doc)

	if len(diagnostics) != 0 {
		for _, d := range diagnostics {
			t.Errorf("unexpected diagnostic: %s", d.Message)
		}
	}
}

func TestRequirementDiagnostics_UnconditionalRequirement(t *testing.T) {
	ctx := testhelpers.NewMockServerContext()

	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)

	ctx.AddManifest(makeTestManifest())

	// icon is set without accessible-label -> violation (unconditional MUST from icon's description)
	html := `<my-button icon="star"></my-button>`
	doc := dm.OpenDocument("test.html", html, 1)
	ctx.AddDocument("test.html", doc)

	diagnostics := publishDiagnostics.AnalyzeRequirementDiagnosticsForTest(ctx, doc)

	if len(diagnostics) == 0 {
		t.Fatal("expected at least 1 diagnostic for unconditional requirement violation")
	}
}

func TestRequirementDiagnostics_LoadingRequiresLabel(t *testing.T) {
	ctx := testhelpers.NewMockServerContext()

	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)

	ctx.AddManifest(makeTestManifest())

	// loading is set without accessible-label -> violation
	html := `<my-button loading></my-button>`
	doc := dm.OpenDocument("test.html", html, 1)
	ctx.AddDocument("test.html", doc)

	diagnostics := publishDiagnostics.AnalyzeRequirementDiagnosticsForTest(ctx, doc)

	if len(diagnostics) == 0 {
		t.Fatal("expected diagnostic for loading without accessible-label")
	}
}
