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

// Nunjucks Custom Element Extraction
// ============================================================================

func TestNunjucks_FindCustomElements(t *testing.T) {
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create document manager: %v", err)
	}
	defer dm.Close()

	content := readTemplateFixture(t, "nunjucks-conditional-attr.njk")
	elements := openAndFindElements(t, dm, "test://page.njk", content)

	assertElementsFound(t, elements, []string{"rh-tile", "uxdot-example", "site-header", "nav-menu", "nav-item"})
}

func TestNunjucks_InTagTemplateSyntax(t *testing.T) {
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create document manager: %v", err)
	}
	defer dm.Close()

	content := readTemplateFixture(t, "nunjucks-conditional-attr.njk")
	elements := openAndFindElements(t, dm, "test://page.njk", content)

	el := findElement(elements, "rh-tile")
	if el == nil {
		t.Fatal("rh-tile not found")
	}

	// Template syntax should NOT appear as attributes
	for _, bad := range []string{"{%", "if", "comingSoon", "%}disabled{%", "endif", "%}"} {
		if _, ok := el.Attributes[bad]; ok {
			t.Errorf("Template syntax %q should not be an attribute", bad)
		}
	}

	// Real attributes should be present
	if _, ok := el.Attributes["compact"]; !ok {
		t.Error("Expected 'compact' attribute on rh-tile")
	}
	if _, ok := el.Attributes["bleed"]; !ok {
		t.Error("Expected 'bleed' attribute on rh-tile")
	}
	// `disabled` from {% if comingSoon %}disabled{% endif %} should survive stripping
	if _, ok := el.Attributes["disabled"]; !ok {
		t.Error("Expected 'disabled' attribute on rh-tile (from conditional template block)")
	}
}

func TestNunjucks_PositionPreservation(t *testing.T) {
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create document manager: %v", err)
	}
	defer dm.Close()

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
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create document manager: %v", err)
	}
	defer dm.Close()

	content := readTemplateFixture(t, "no-elements.njk")
	elements := openAndFindElements(t, dm, "test://plain.njk", content)

	if len(elements) != 0 {
		t.Errorf("Expected 0 custom elements, got %d: %v", len(elements), tagNames(elements))
	}
}

func TestNunjucks_AttributeExtraction(t *testing.T) {
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create document manager: %v", err)
	}
	defer dm.Close()

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
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create document manager: %v", err)
	}
	defer dm.Close()

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
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create document manager: %v", err)
	}
	defer dm.Close()

	content := readTemplateFixture(t, "handlebars-conditional-attr.hbs")
	elements := openAndFindElements(t, dm, "test://page.hbs", content)

	assertElementsFound(t, elements, []string{"site-header", "nav-menu", "nav-item", "rh-tile", "uxdot-example"})
}

func TestHandlebars_InTagTemplateSyntax(t *testing.T) {
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create document manager: %v", err)
	}
	defer dm.Close()

	content := readTemplateFixture(t, "handlebars-conditional-attr.hbs")
	elements := openAndFindElements(t, dm, "test://page.hbs", content)

	el := findElement(elements, "rh-tile")
	if el == nil {
		t.Fatal("rh-tile not found")
	}

	// Template syntax should NOT appear as attributes
	for _, bad := range []string{"{{#if", "this.comingSoon}}", "{{/if}}"} {
		if _, ok := el.Attributes[bad]; ok {
			t.Errorf("Template syntax %q should not be an attribute", bad)
		}
	}

	// Real attributes should be present
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
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create document manager: %v", err)
	}
	defer dm.Close()

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
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create document manager: %v", err)
	}
	defer dm.Close()

	content := readTemplateFixture(t, "no-elements.hbs")
	elements := openAndFindElements(t, dm, "test://plain.hbs", content)

	if len(elements) != 0 {
		t.Errorf("Expected 0 custom elements, got %d: %v", len(elements), tagNames(elements))
	}
}

func TestHandlebars_AttributeExtraction(t *testing.T) {
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create document manager: %v", err)
	}
	defer dm.Close()

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

// Twig In-Tag Template Syntax (regression test)
// ============================================================================

func TestTwig_InTagTemplateSyntax(t *testing.T) {
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create document manager: %v", err)
	}
	defer dm.Close()

	content := readTemplateFixture(t, "twig-intag-conditional.html.twig")
	elements := openAndFindElements(t, dm, "test://page.html.twig", content)

	assertElementsFound(t, elements, []string{"site-header", "rh-tile"})

	el := findElement(elements, "rh-tile")
	if el == nil {
		t.Fatal("rh-tile not found")
	}

	// Template syntax should NOT appear as attributes
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
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create document manager: %v", err)
	}
	defer dm.Close()

	// .twig extension (without .html prefix) should also use template handler
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

// Extension Routing
// ============================================================================

func TestTemplateExtensionRouting(t *testing.T) {
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create document manager: %v", err)
	}
	defer dm.Close()

	content := readTemplateFixture(t, "nunjucks-conditional-attr.njk")

	for _, ext := range []string{".njk", ".j2", ".jinja", ".jinja2", ".liquid"} {
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

func TestTemplateCompoundExtensionRouting(t *testing.T) {
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create document manager: %v", err)
	}
	defer dm.Close()

	content := readTemplateFixture(t, "nunjucks-conditional-attr.njk")

	for _, ext := range []string{".html.twig", ".html.j2", ".html.jinja2"} {
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
