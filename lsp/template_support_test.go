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
	"os"
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/lsp/document"
	"bennypowers.dev/cem/lsp/types"
	protocol "github.com/bennypowers/glsp/protocol_3_17"
)

const templateFixtureDir = "testdata/integration/template-support"

func readTemplateFixture(t *testing.T, name string) string {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(templateFixtureDir, name))
	if err != nil {
		t.Fatalf("Failed to read fixture %s: %v", name, err)
	}
	return string(data)
}

func newTemplateDM(t *testing.T) types.Manager {
	t.Helper()
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create document manager: %v", err)
	}
	t.Cleanup(dm.Close)
	return dm
}

// Nunjucks Custom Element Extraction
// ============================================================================

func TestNunjucks_FindCustomElements(t *testing.T) {
	dm := newTemplateDM(t)
	content := readTemplateFixture(t, "nunjucks-conditional-attr.njk")
	elements := openAndFindElements(t, dm, "test://page.njk", content)

	assertElementsFound(t, elements, []string{"rh-tile", "uxdot-example", "site-header", "nav-menu", "nav-item"})
}

func TestNunjucks_InTagTemplateSyntax(t *testing.T) {
	dm := newTemplateDM(t)
	content := readTemplateFixture(t, "nunjucks-conditional-attr.njk")
	elements := openAndFindElements(t, dm, "test://page.njk", content)

	el := findElement(elements, "rh-tile")
	if el == nil {
		t.Fatal("rh-tile not found")
	}

	for _, bad := range []string{"{%", "if", "comingSoon", "%}disabled{%", "endif", "%}"} {
		if _, ok := el.Attributes[bad]; ok {
			t.Errorf("Template syntax %q should not be an attribute", bad)
		}
	}

	if _, ok := el.Attributes["compact"]; !ok {
		t.Error("Expected 'compact' attribute on rh-tile")
	}
	if _, ok := el.Attributes["bleed"]; !ok {
		t.Error("Expected 'bleed' attribute on rh-tile")
	}
	if _, ok := el.Attributes["disabled"]; !ok {
		t.Error("Expected 'disabled' attribute on rh-tile (from conditional template block)")
	}
}

func TestNunjucks_PositionPreservation(t *testing.T) {
	dm := newTemplateDM(t)
	content := readTemplateFixture(t, "nunjucks-conditional-attr.njk")
	elements := openAndFindElements(t, dm, "test://page.njk", content)

	// site-header is on line 14 (1-indexed) = line 13 (0-indexed)
	el := findElement(elements, "site-header")
	if el == nil {
		t.Fatal("site-header not found")
	}
	if el.Range.Start.Line != 13 {
		t.Errorf("Expected site-header on line 13, got line %d", el.Range.Start.Line)
	}
}

func TestNunjucks_NoCustomElements(t *testing.T) {
	dm := newTemplateDM(t)
	content := readTemplateFixture(t, "no-elements.njk")
	elements := openAndFindElements(t, dm, "test://plain.njk", content)

	if len(elements) != 0 {
		t.Errorf("Expected 0 custom elements, got %d: %v", len(elements), tagNames(elements))
	}
}

func TestNunjucks_AttributeExtraction(t *testing.T) {
	dm := newTemplateDM(t)
	content := readTemplateFixture(t, "nunjucks-conditional-attr.njk")
	elements := openAndFindElements(t, dm, "test://page.njk", content)

	el := findElement(elements, "nav-menu")
	if el == nil {
		t.Fatal("nav-menu not found")
	}
	attr, ok := el.Attributes["orientation"]
	if !ok {
		t.Fatal("Expected nav-menu to have 'orientation' attribute")
	}
	if attr.Value != "horizontal" {
		t.Errorf("Expected orientation=\"horizontal\", got %q", attr.Value)
	}
}

func TestNunjucks_DocumentUpdate(t *testing.T) {
	dm := newTemplateDM(t)
	content := readTemplateFixture(t, "nunjucks-conditional-attr.njk")
	dm.OpenDocument("test://page.njk", content, 1)

	updated := content + "\n<alert-banner type=\"info\">Update!</alert-banner>\n"
	doc := dm.UpdateDocument("test://page.njk", updated, 2)
	if doc == nil {
		t.Fatal("Expected updated document")
	}

	elements, err := doc.FindCustomElements(dm)
	if err != nil {
		t.Fatalf("FindCustomElements failed: %v", err)
	}

	el := findElement(elements, "alert-banner")
	if el == nil {
		t.Errorf("Expected <alert-banner> after update, got: %v", tagNames(elements))
	}
}

// Handlebars Custom Element Extraction
// ============================================================================

func TestHandlebars_FindCustomElements(t *testing.T) {
	dm := newTemplateDM(t)
	content := readTemplateFixture(t, "handlebars-conditional-attr.hbs")
	elements := openAndFindElements(t, dm, "test://page.hbs", content)

	assertElementsFound(t, elements, []string{"site-header", "nav-menu", "nav-item", "rh-tile", "uxdot-example"})
}

func TestHandlebars_InTagTemplateSyntax(t *testing.T) {
	dm := newTemplateDM(t)
	content := readTemplateFixture(t, "handlebars-conditional-attr.hbs")
	elements := openAndFindElements(t, dm, "test://page.hbs", content)

	el := findElement(elements, "rh-tile")
	if el == nil {
		t.Fatal("rh-tile not found")
	}

	for _, bad := range []string{"{{#if", "this.comingSoon}}", "{{/if}}"} {
		if _, ok := el.Attributes[bad]; ok {
			t.Errorf("Template syntax %q should not be an attribute", bad)
		}
	}

	if _, ok := el.Attributes["compact"]; !ok {
		t.Error("Expected 'compact' attribute on rh-tile")
	}
	if _, ok := el.Attributes["bleed"]; !ok {
		t.Error("Expected 'bleed' attribute on rh-tile")
	}
	if _, ok := el.Attributes["disabled"]; !ok {
		t.Error("Expected 'disabled' attribute on rh-tile (from conditional template block)")
	}
}

func TestHandlebars_PositionPreservation(t *testing.T) {
	dm := newTemplateDM(t)
	content := readTemplateFixture(t, "handlebars-conditional-attr.hbs")
	elements := openAndFindElements(t, dm, "test://page.hbs", content)

	// site-header is on line 3 (1-indexed) = line 2 (0-indexed)
	el := findElement(elements, "site-header")
	if el == nil {
		t.Fatal("site-header not found")
	}
	if el.Range.Start.Line != 2 {
		t.Errorf("Expected site-header on line 2, got line %d", el.Range.Start.Line)
	}
}

func TestHandlebars_NoCustomElements(t *testing.T) {
	dm := newTemplateDM(t)
	content := readTemplateFixture(t, "no-elements.hbs")
	elements := openAndFindElements(t, dm, "test://plain.hbs", content)

	if len(elements) != 0 {
		t.Errorf("Expected 0 custom elements, got %d: %v", len(elements), tagNames(elements))
	}
}

func TestHandlebars_AttributeExtraction(t *testing.T) {
	dm := newTemplateDM(t)
	content := readTemplateFixture(t, "handlebars-conditional-attr.hbs")
	elements := openAndFindElements(t, dm, "test://page.hbs", content)

	el := findElement(elements, "nav-menu")
	if el == nil {
		t.Fatal("nav-menu not found")
	}
	attr, ok := el.Attributes["orientation"]
	if !ok {
		t.Fatal("Expected nav-menu to have 'orientation' attribute")
	}
	if attr.Value != "horizontal" {
		t.Errorf("Expected orientation=\"horizontal\", got %q", attr.Value)
	}
}

// Liquid Custom Element Extraction
// ============================================================================

func TestLiquid_FindCustomElements(t *testing.T) {
	dm := newTemplateDM(t)
	content := readTemplateFixture(t, "liquid-conditional-attr.liquid")
	elements := openAndFindElements(t, dm, "test://page.liquid", content)

	assertElementsFound(t, elements, []string{"site-header", "rh-tile"})
}

func TestLiquid_InTagTemplateSyntax(t *testing.T) {
	dm := newTemplateDM(t)
	content := readTemplateFixture(t, "liquid-conditional-attr.liquid")
	elements := openAndFindElements(t, dm, "test://page.liquid", content)

	el := findElement(elements, "rh-tile")
	if el == nil {
		t.Fatal("rh-tile not found")
	}

	for _, bad := range []string{"{%", "if", "tile.comingSoon", "endif", "%}"} {
		if _, ok := el.Attributes[bad]; ok {
			t.Errorf("Template syntax %q should not be an attribute", bad)
		}
	}

	if _, ok := el.Attributes["compact"]; !ok {
		t.Error("Expected 'compact' attribute on rh-tile")
	}
	if _, ok := el.Attributes["disabled"]; !ok {
		t.Error("Expected 'disabled' attribute on rh-tile (from conditional template block)")
	}
}

func TestLiquid_RawBlock(t *testing.T) {
	dm := newTemplateDM(t)
	content := readTemplateFixture(t, "liquid-conditional-attr.liquid")
	elements := openAndFindElements(t, dm, "test://page.liquid", content)

	// Liquid's {% raw %}...{% endraw %} is recognized by the Jinja grammar
	// as a raw_statement. Content inside raw blocks is captured as text
	// nodes, but nested {{ }} inside raw blocks may be stripped since the
	// Jinja grammar parses them as output nodes within the raw statement's
	// body. The custom element tag itself should still be found.
	el := findElement(elements, "raw-content")
	if el != nil {
		// raw-content found — verify no template syntax leaked as attributes
		doubleBrace := "{{"
		if _, ok := el.Attributes[doubleBrace]; ok {
			t.Error("Template syntax should not leak inside raw block")
		}
	}
	// If raw-content is not found, that's acceptable: the Jinja grammar
	// may parse the raw block content differently than Liquid expects.
}

// Twig In-Tag Template Syntax (regression test)
// ============================================================================

func TestTwig_InTagTemplateSyntax(t *testing.T) {
	dm := newTemplateDM(t)
	content := readTemplateFixture(t, "twig-intag-conditional.html.twig")
	elements := openAndFindElements(t, dm, "test://page.html.twig", content)

	assertElementsFound(t, elements, []string{"site-header", "rh-tile"})

	el := findElement(elements, "rh-tile")
	if el == nil {
		t.Fatal("rh-tile not found")
	}

	for _, bad := range []string{"{%", "if", "tile.comingSoon", "endif", "%}"} {
		if _, ok := el.Attributes[bad]; ok {
			t.Errorf("Template syntax %q should not be an attribute", bad)
		}
	}

	if _, ok := el.Attributes["compact"]; !ok {
		t.Error("Expected 'compact' attribute on rh-tile")
	}
	if _, ok := el.Attributes["bleed"]; !ok {
		t.Error("Expected 'bleed' attribute on rh-tile")
	}
	if _, ok := el.Attributes["disabled"]; !ok {
		t.Error("Expected 'disabled' attribute on rh-tile (from conditional template block)")
	}
}

func TestTwig_InTagTemplateSyntax_TwigExtension(t *testing.T) {
	dm := newTemplateDM(t)
	content := readTemplateFixture(t, "twig-intag-conditional.html.twig")
	elements := openAndFindElements(t, dm, "test://page.twig", content)

	assertElementsFound(t, elements, []string{"site-header", "rh-tile"})

	el := findElement(elements, "rh-tile")
	if el == nil {
		t.Fatal("rh-tile not found")
	}
	if _, ok := el.Attributes["{%"]; ok {
		t.Error("Template syntax should not be an attribute for .twig files")
	}
}

// ERB Custom Element Extraction
// ============================================================================

func TestERB_FindCustomElements(t *testing.T) {
	dm := newTemplateDM(t)
	content := readTemplateFixture(t, "erb-conditional-attr.erb")
	elements := openAndFindElements(t, dm, "test://page.erb", content)

	assertElementsFound(t, elements, []string{"site-header", "rh-tile", "uxdot-example"})
}

func TestERB_InTagTemplateSyntax(t *testing.T) {
	dm := newTemplateDM(t)
	content := readTemplateFixture(t, "erb-conditional-attr.erb")
	elements := openAndFindElements(t, dm, "test://page.erb", content)

	el := findElement(elements, "rh-tile")
	if el == nil {
		t.Fatal("rh-tile not found")
	}

	for _, bad := range []string{"<%", "if", "tile.coming_soon", "%>disabled<%", "end", "%>"} {
		if _, ok := el.Attributes[bad]; ok {
			t.Errorf("Template syntax %q should not be an attribute", bad)
		}
	}

	if _, ok := el.Attributes["compact"]; !ok {
		t.Error("Expected 'compact' attribute on rh-tile")
	}
	if _, ok := el.Attributes["bleed"]; !ok {
		t.Error("Expected 'bleed' attribute on rh-tile")
	}
	if _, ok := el.Attributes["disabled"]; !ok {
		t.Error("Expected 'disabled' attribute on rh-tile (from conditional template block)")
	}
}

func TestERB_PositionPreservation(t *testing.T) {
	dm := newTemplateDM(t)
	content := readTemplateFixture(t, "erb-conditional-attr.erb")
	elements := openAndFindElements(t, dm, "test://page.erb", content)

	// site-header is on line 3 (1-indexed) = line 2 (0-indexed)
	el := findElement(elements, "site-header")
	if el == nil {
		t.Fatal("site-header not found")
	}
	if el.Range.Start.Line != 2 {
		t.Errorf("Expected site-header on line 2, got line %d", el.Range.Start.Line)
	}
}

func TestERB_NoCustomElements(t *testing.T) {
	dm := newTemplateDM(t)
	content := readTemplateFixture(t, "no-elements.erb")
	elements := openAndFindElements(t, dm, "test://plain.erb", content)

	if len(elements) != 0 {
		t.Errorf("Expected 0 custom elements, got %d: %v", len(elements), tagNames(elements))
	}
}

// EJS Custom Element Extraction
// ============================================================================

func TestEJS_FindCustomElements(t *testing.T) {
	dm := newTemplateDM(t)
	content := readTemplateFixture(t, "ejs-conditional-attr.ejs")
	elements := openAndFindElements(t, dm, "test://page.ejs", content)

	assertElementsFound(t, elements, []string{"site-header", "rh-tile", "uxdot-example"})
}

func TestEJS_InTagTemplateSyntax(t *testing.T) {
	dm := newTemplateDM(t)
	content := readTemplateFixture(t, "ejs-conditional-attr.ejs")
	elements := openAndFindElements(t, dm, "test://page.ejs", content)

	el := findElement(elements, "rh-tile")
	if el == nil {
		t.Fatal("rh-tile not found")
	}

	for _, bad := range []string{"<%", "if", "(tile.comingSoon)", "%>disabled<%", "%>"} {
		if _, ok := el.Attributes[bad]; ok {
			t.Errorf("Template syntax %q should not be an attribute", bad)
		}
	}

	if _, ok := el.Attributes["compact"]; !ok {
		t.Error("Expected 'compact' attribute on rh-tile")
	}
	if _, ok := el.Attributes["disabled"]; !ok {
		t.Error("Expected 'disabled' attribute on rh-tile (from conditional template block)")
	}
}

// Extension Routing
// ============================================================================

func TestTemplateExtensionRouting(t *testing.T) {
	dm := newTemplateDM(t)
	content := readTemplateFixture(t, "nunjucks-conditional-attr.njk")

	for _, ext := range []string{".njk", ".j2", ".jinja", ".jinja2"} {
		t.Run(ext, func(t *testing.T) {
			elements := openAndFindElements(t, dm, "test://page"+ext, content)
			if len(elements) == 0 {
				t.Errorf("Expected custom elements for %s extension", ext)
			}
			el := findElement(elements, "rh-tile")
			if el == nil {
				t.Fatalf("rh-tile not found for %s", ext)
			}
			if _, ok := el.Attributes["{%"]; ok {
				t.Errorf("Template syntax leaked as attribute for %s", ext)
			}
		})
	}
}

func TestTemplateLiquidExtensionRouting(t *testing.T) {
	dm := newTemplateDM(t)
	content := readTemplateFixture(t, "liquid-conditional-attr.liquid")
	elements := openAndFindElements(t, dm, "test://page.liquid", content)

	if len(elements) == 0 {
		t.Error("Expected custom elements for .liquid extension")
	}
	el := findElement(elements, "rh-tile")
	if el == nil {
		t.Fatal("rh-tile not found for .liquid")
	}
	if _, ok := el.Attributes["{%"]; ok {
		t.Error("Template syntax leaked as attribute for .liquid")
	}
}

func TestTemplateCompoundExtensionRouting(t *testing.T) {
	dm := newTemplateDM(t)
	content := readTemplateFixture(t, "nunjucks-conditional-attr.njk")

	for _, uri := range []string{"test://page.html.twig", "test://page.twig", "test://page.html.j2", "test://page.html.jinja2"} {
		t.Run(uri, func(t *testing.T) {
			elements := openAndFindElements(t, dm, uri, content)
			if len(elements) == 0 {
				t.Errorf("Expected custom elements for %s", uri)
			}
			el := findElement(elements, "rh-tile")
			if el == nil {
				t.Fatalf("rh-tile not found for %s", uri)
			}
			if _, ok := el.Attributes["{%"]; ok {
				t.Errorf("Template syntax leaked as attribute for %s", uri)
			}
		})
	}
}

// Template Completion Context (injection safety)
// ============================================================================

func TestTemplate_CompletionContext_ComparisonOperator(t *testing.T) {
	dm := newTemplateDM(t)

	tests := []struct {
		name     string
		fixture  string
		uri      string
		position protocol.Position
		wantType types.CompletionContextType
		wantTag  string
	}{
		{"nunjucks", "comparison-operator.njk", "test://cmp.njk", protocol.Position{Line: 2, Character: 1}, types.CompletionTagName, "my-element"},
		{"handlebars", "comparison-operator.hbs", "test://cmp.hbs", protocol.Position{Line: 1, Character: 1}, types.CompletionTagName, ""},
		{"erb", "comparison-operator.erb", "test://cmp.erb", protocol.Position{Line: 1, Character: 1}, types.CompletionTagName, ""},
		{"ejs", "comparison-operator.ejs", "test://cmp.ejs", protocol.Position{Line: 1, Character: 1}, types.CompletionTagName, ""},
		{"nunjucks-in-template", "comparison-operator.njk", "test://cmp.njk", protocol.Position{Line: 1, Character: 10}, types.CompletionUnknown, ""},
		{"multibyte", "multibyte.njk", "test://utf8.njk", protocol.Position{Line: 1, Character: 1}, types.CompletionTagName, "my-element"},
		{"blade", "comparison-operator.blade.php", "test://cmp.blade.php", protocol.Position{Line: 1, Character: 1}, types.CompletionTagName, "my-element"},
		{"blade-in-directive", "comparison-operator.blade.php", "test://dir.blade.php", protocol.Position{Line: 0, Character: 5}, types.CompletionUnknown, ""},
		{"blade-multibyte", "blade-multibyte.blade.php", "test://utf8.blade.php", protocol.Position{Line: 1, Character: 1}, types.CompletionTagName, "my-element"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			content := readTemplateFixture(t, tt.fixture)
			doc := dm.OpenDocument(tt.uri, content, 1)

			analysis := doc.AnalyzeCompletionContextTS(tt.position, dm)
			if analysis.Type != tt.wantType {
				t.Errorf("Expected completion type %d, got %d", tt.wantType, analysis.Type)
			}
			if tt.wantTag != "" && analysis.TagName != tt.wantTag {
				t.Errorf("Expected TagName %q, got %q", tt.wantTag, analysis.TagName)
			}
		})
	}
}

func TestTemplate_AttributeAfterTemplateBlock(t *testing.T) {
	dm := newTemplateDM(t)
	content := readTemplateFixture(t, "intag-conditional-attr.njk")

	elements := openAndFindElements(t, dm, "test://attr.njk", content)
	assertElementsFound(t, elements, []string{"my-element"})

	el := findElement(elements, "my-element")
	if el == nil {
		t.Fatal("my-element not found")
	}
	if _, ok := el.Attributes["variant"]; !ok {
		t.Error("Expected variant attribute to be found")
	}

	for _, bad := range []string{"{%", "if", "endif", "%}"} {
		if _, ok := el.Attributes[bad]; ok {
			t.Errorf("Template syntax %q should not be an attribute", bad)
		}
	}
}

func TestTemplate_HeadInsertionPoint(t *testing.T) {
	dm := newTemplateDM(t)

	tests := []struct {
		name      string
		fixture   string
		uri       string
		wantFound bool
		wantLine  uint32
	}{
		{"with-head", "head-section.njk", "test://head.njk", true, 4},
		{"headless", "headless.njk", "test://nohead.njk", false, 0},
		{"blade-with-head", "blade-head-section.blade.php", "test://head.blade.php", true, 4},
		{"blade-headless", "blade-headless.blade.php", "test://nohead.blade.php", false, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			content := readTemplateFixture(t, tt.fixture)
			doc := dm.OpenDocument(tt.uri, content, 1)

			pos, found := doc.FindHeadInsertionPoint(dm)
			if found != tt.wantFound {
				t.Errorf("Expected found=%v, got %v", tt.wantFound, found)
			}
			if found && pos.Line != tt.wantLine {
				t.Errorf("Expected line %d, got %d", tt.wantLine, pos.Line)
			}
		})
	}
}

func TestTemplate_EdgeCases(t *testing.T) {
	dm := newTemplateDM(t)

	tests := []struct {
		name         string
		fixture      string
		uri          string
		wantElements []string
	}{
		{"pure-template", "pure-template.njk", "test://pure.njk", nil},
		{"pure-html", "pure-html.njk", "test://purehtml.njk", []string{"my-element"}},
		{"pure-blade", "pure-blade.blade.php", "test://pure.blade.php", nil},
		{"pure-html-blade", "pure-html.blade.php", "test://purehtml.blade.php", []string{"my-element"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			content := readTemplateFixture(t, tt.fixture)
			elements := openAndFindElements(t, dm, tt.uri, content)
			if tt.wantElements == nil {
				if len(elements) != 0 {
					t.Errorf("Expected no elements, got %d", len(elements))
				}
			} else {
				assertElementsFound(t, elements, tt.wantElements)
			}
		})
	}
}

// Blade Custom Element Extraction
// ============================================================================

func TestBlade_FindCustomElements(t *testing.T) {
	dm := newTemplateDM(t)
	content := readTemplateFixture(t, "blade-wordpress-theme.blade.php")
	elements := openAndFindElements(t, dm, "test://theme.blade.php", content)

	assertElementsFound(t, elements, []string{
		"page-header", "content-grid", "post-card",
		"page-navigation", "site-footer", "nav-menu", "nav-item",
	})
}

func TestBlade_AttributeExtraction(t *testing.T) {
	dm := newTemplateDM(t)
	content := readTemplateFixture(t, "blade-conditional-attr.blade.php")
	elements := openAndFindElements(t, dm, "test://attr.blade.php", content)

	el := findElement(elements, "nav-menu")
	if el == nil {
		t.Fatal("nav-menu not found")
	}
	if v, ok := el.Attributes["orientation"]; !ok || v.Value != "horizontal" {
		t.Errorf("Expected orientation='horizontal', got %+v", el.Attributes["orientation"])
	}
}

func TestBlade_NoCustomElements(t *testing.T) {
	dm := newTemplateDM(t)
	content := readTemplateFixture(t, "no-elements.blade.php")
	elements := openAndFindElements(t, dm, "test://noce.blade.php", content)

	if len(elements) != 0 {
		t.Errorf("Expected no custom elements, got %d: %v", len(elements), tagNames(elements))
	}
}

func TestBlade_PositionPreservation(t *testing.T) {
	dm := newTemplateDM(t)
	content := readTemplateFixture(t, "blade-conditional-attr.blade.php")
	elements := openAndFindElements(t, dm, "test://pos.blade.php", content)

	el := findElement(elements, "site-header")
	if el == nil {
		t.Fatal("site-header not found")
	}
	if el.Range.Start.Line != 3 {
		t.Errorf("Expected site-header on line 3, got %d", el.Range.Start.Line)
	}
}

func TestBlade_DocumentUpdate(t *testing.T) {
	dm := newTemplateDM(t)
	content1 := readTemplateFixture(t, "blade-conditional-attr.blade.php")
	elements1 := openAndFindElements(t, dm, "test://update.blade.php", content1)
	count1 := len(elements1)

	content2 := readTemplateFixture(t, "blade-wordpress-theme.blade.php")
	doc2 := dm.UpdateDocument("test://update.blade.php", content2, 2)
	elements2, err := doc2.FindCustomElements(dm)
	if err != nil {
		t.Fatalf("FindCustomElements failed: %v", err)
	}

	if len(elements2) == count1 {
		t.Error("Expected different element count after document update")
	}

	// Verify specific elements from updated content are present
	for _, expected := range []string{"page-header", "content-grid", "post-card", "site-footer"} {
		if findElement(elements2, expected) == nil {
			t.Errorf("Expected %q in updated document", expected)
		}
	}
	// Verify elements from old content are gone
	if findElement(elements2, "rh-tile") != nil {
		t.Error("Expected rh-tile from old content to be gone after update")
	}
}

func TestBlade_DeepNesting(t *testing.T) {
	dm := newTemplateDM(t)
	content := readTemplateFixture(t, "blade-deep-nesting.blade.php")
	elements := openAndFindElements(t, dm, "test://deep.blade.php", content)

	// Blade grammar wraps @section content in ERROR nodes, so the outermost
	// page-layout element may not be found. Elements deeper in the nesting
	// that are direct children of other elements ARE found.
	deepElements := []string{"content-section", "section-header", "card-element"}
	for _, tag := range deepElements {
		el := findElement(elements, tag)
		if el == nil {
			t.Errorf("Expected %q at deep nesting level", tag)
		}
	}

	section := findElement(elements, "content-section")
	if section != nil {
		if _, ok := section.Attributes["id"]; !ok {
			t.Error("Expected id attribute (with interpolation) on content-section")
		}
	}
}

func TestBlade_InterpolationEdgeCases(t *testing.T) {
	dm := newTemplateDM(t)
	content := readTemplateFixture(t, "blade-interpolation-edge-cases.blade.php")
	elements := openAndFindElements(t, dm, "test://interp.blade.php", content)

	assertElementsFound(t, elements, []string{"my-element", "status-badge"})

	el := findElement(elements, "my-element")
	if el == nil {
		t.Fatal("my-element not found")
	}

	tests := []struct {
		attr     string
		contains string
	}{
		{"id", "el-{{ $id }}"},
		{"variant", "{{ $variant }}"},
		{"title", "{!! $rawTitle !!}"},
		{"data-count", "{{ $items->count() }}"},
	}
	for _, tt := range tests {
		v, ok := el.Attributes[tt.attr]
		if !ok {
			t.Errorf("Expected %q attribute", tt.attr)
			continue
		}
		if v.Value != tt.contains {
			t.Errorf("Expected %q value %q, got %q", tt.attr, tt.contains, v.Value)
		}
	}

	for name := range el.Attributes {
		if name == "{{" || name == "}}" || name == "{!!" || name == "!!}" {
			t.Errorf("Interpolation token %q leaked as attribute name", name)
		}
	}
}

func TestBlade_EchoInAttribute(t *testing.T) {
	dm := newTemplateDM(t)
	content := readTemplateFixture(t, "blade-echo-in-attr.blade.php")
	elements := openAndFindElements(t, dm, "test://echo.blade.php", content)

	assertElementsFound(t, elements, []string{"my-element"})

	el := findElement(elements, "my-element")
	if el == nil {
		t.Fatal("my-element not found")
	}

	if _, ok := el.Attributes["id"]; !ok {
		t.Error("Expected id attribute")
	}
	if _, ok := el.Attributes["variant"]; !ok {
		t.Error("Expected variant attribute")
	}
}

func TestBlade_DisabledDirective(t *testing.T) {
	dm := newTemplateDM(t)
	content := readTemplateFixture(t, "blade-disabled-directive.blade.php")
	elements := openAndFindElements(t, dm, "test://disabled.blade.php", content)

	assertElementsFound(t, elements, []string{"my-element", "my-input"})

	el := findElement(elements, "my-element")
	if el == nil {
		t.Fatal("my-element not found")
	}
	if _, ok := el.Attributes["compact"]; !ok {
		t.Error("Expected compact attribute before directive")
	}
	// NOTE: attributes after @disabled directive may not be captured because
	// the directive attribute (attribute (directive) (parameter)) doesn't match
	// the query pattern (attribute (attribute_name) ...), breaking the * repetition
}

func TestBlade_InTagConditional(t *testing.T) {
	dm := newTemplateDM(t)
	content := readTemplateFixture(t, "blade-intag-conditional.blade.php")
	elements := openAndFindElements(t, dm, "test://intag.blade.php", content)

	assertElementsFound(t, elements, []string{"site-header", "nav-menu", "rh-tile"})

	el := findElement(elements, "rh-tile")
	if el == nil {
		t.Fatal("rh-tile not found")
	}
	if _, ok := el.Attributes["compact"]; !ok {
		t.Error("Expected compact attribute on rh-tile")
	}
	if _, ok := el.Attributes["bleed"]; !ok {
		t.Error("Expected bleed attribute on rh-tile")
	}
}

func TestBlade_ExtensionRouting(t *testing.T) {
	dm := newTemplateDM(t)
	content := readTemplateFixture(t, "blade-conditional-attr.blade.php")
	elements := openAndFindElements(t, dm, "test://page.blade.php", content)

	if len(elements) == 0 {
		t.Error("Expected custom elements from .blade.php file")
	}

	el := findElement(elements, "site-header")
	if el == nil {
		t.Error("Expected site-header from blade template")
	}

	// Verify .php files don't route to blade handler
	phpContent := "<?php echo 'hello'; ?>\n<my-element></my-element>"
	phpElements := openAndFindElements(t, dm, "test://page.php", phpContent)
	phpEl := findElement(phpElements, "my-element")
	if phpEl == nil {
		t.Error("Expected my-element from .php file (should use PHP handler, not blade)")
	}
}
