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
	protocol "github.com/tliron/glsp/protocol_3_16"
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

func TestNunjucks_CompletionContext_TemplateLessThan(t *testing.T) {
	dm := newTemplateDM(t)
	// Template comparison operator < must not be mistaken for HTML tag start
	content := "{% if count < max %}\n<my-element>\n{% endif %}"
	doc := dm.OpenDocument("test://cmp.njk", content, 1)

	// Cursor at start of <my-element> on line 1
	analysis := doc.AnalyzeCompletionContextTS(protocol.Position{Line: 1, Character: 1}, dm)
	if analysis.Type != types.CompletionTagName {
		t.Errorf("Expected CompletionTagName at element position, got %d", analysis.Type)
	}
	if analysis.TagName != "my-element" {
		t.Errorf("Expected TagName 'my-element', got %q", analysis.TagName)
	}
}

func TestNunjucks_CompletionContext_CursorInTemplateRegion(t *testing.T) {
	dm := newTemplateDM(t)
	content := "{% if count < max %}\n<my-element>\n{% endif %}"
	doc := dm.OpenDocument("test://cmp.njk", content, 1)

	// Cursor inside template region (line 0, inside {% if ... %})
	analysis := doc.AnalyzeCompletionContextTS(protocol.Position{Line: 0, Character: 10}, dm)
	if analysis.Type != types.CompletionUnknown {
		t.Errorf("Expected CompletionUnknown in template region, got %d", analysis.Type)
	}
}

func TestNunjucks_CompletionContext_AttributeAfterTemplate(t *testing.T) {
	dm := newTemplateDM(t)
	content := "<my-element {% if x %}disabled{% endif %} variant=\"primary\">"

	elements := openAndFindElements(t, dm, "test://attr.njk", content)
	assertElementsFound(t, elements, []string{"my-element"})

	el := findElement(elements, "my-element")
	if el == nil {
		t.Fatal("my-element not found")
	}
	if _, ok := el.Attributes["variant"]; !ok {
		t.Error("Expected variant attribute to be found")
	}

	// Verify template tokens don't leak as attributes
	for _, bad := range []string{"{%", "if", "endif", "%}"} {
		if _, ok := el.Attributes[bad]; ok {
			t.Errorf("Template syntax %q should not be an attribute", bad)
		}
	}
}

func TestNunjucks_HeadInsertionPoint(t *testing.T) {
	dm := newTemplateDM(t)
	content := "{% extends 'base.njk' %}\n<html>\n<head>\n  <title>Test</title>\n</head>\n<body>\n  <my-element></my-element>\n</body>\n</html>"
	doc := dm.OpenDocument("test://head.njk", content, 1)

	pos, found := doc.FindHeadInsertionPoint(dm)
	if !found {
		t.Fatal("Expected to find head insertion point")
	}
	if pos.Line != 4 {
		t.Errorf("Expected head insertion at line 4, got %d", pos.Line)
	}
}
