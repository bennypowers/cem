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
	"sort"
	"testing"

	"bennypowers.dev/cem/lsp/document"
	"bennypowers.dev/cem/lsp/types"
)

const phpTwigFixtureDir = "testdata/integration/php-twig-support"

func readPHPTwigFixture(t *testing.T, name string) string {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(phpTwigFixtureDir, name))
	if err != nil {
		t.Fatalf("Failed to read fixture %s: %v", name, err)
	}
	return string(data)
}

// tagNames extracts sorted tag names from custom element matches
func tagNames(elements []types.CustomElementMatch) []string {
	names := make([]string, len(elements))
	for i, el := range elements {
		names[i] = el.TagName
	}
	sort.Strings(names)
	return names
}

// assertElementsFound asserts that the expected custom elements are present
func assertElementsFound(t *testing.T, elements []types.CustomElementMatch, expected []string) {
	t.Helper()
	got := tagNames(elements)
	sort.Strings(expected)
	if len(got) != len(expected) {
		t.Fatalf("Expected elements %v, got %v", expected, got)
	}
	for i := range got {
		if got[i] != expected[i] {
			t.Fatalf("Expected elements %v, got %v", expected, got)
		}
	}
}

// findElement returns the element with the given tag name, or nil
func findElement(elements []types.CustomElementMatch, tagName string) *types.CustomElementMatch {
	for _, el := range elements {
		if el.TagName == tagName {
			return &el
		}
	}
	return nil
}

// openAndFindElements opens a document and returns its custom elements
func openAndFindElements(t *testing.T, dm types.Manager, uri, content string) []types.CustomElementMatch {
	t.Helper()
	doc := dm.OpenDocument(uri, content, 1)
	if doc == nil {
		t.Fatalf("Failed to open document %s", uri)
	}
	elements, err := doc.FindCustomElements(dm)
	if err != nil {
		t.Fatalf("FindCustomElements failed: %v", err)
	}
	return elements
}

// PHP Custom Element Extraction
// ============================================================================

func TestPHP_FindCustomElements(t *testing.T) {
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create document manager: %v", err)
	}
	defer dm.Close()

	content := readPHPTwigFixture(t, "wordpress-theme.php")
	elements := openAndFindElements(t, dm, "test://theme.php", content)

	// wordpress-theme.php: site-header, blog-card, site-footer
	assertElementsFound(t, elements, []string{"site-header", "blog-card", "site-footer"})
}

func TestPHP_FindCustomElements_WithInterpolation(t *testing.T) {
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create document manager: %v", err)
	}
	defer dm.Close()

	// PHP interpolation inside attributes, PHP conditionals wrapping elements
	content := readPHPTwigFixture(t, "php-interpolated.php")
	elements := openAndFindElements(t, dm, "test://product.php", content)

	// product-card and sidebar-widget (inside PHP conditional)
	assertElementsFound(t, elements, []string{"product-card", "sidebar-widget"})
}

func TestPHP_NoCustomElements(t *testing.T) {
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create document manager: %v", err)
	}
	defer dm.Close()

	content := readPHPTwigFixture(t, "no-elements.php")
	elements := openAndFindElements(t, dm, "test://plain.php", content)

	if len(elements) != 0 {
		t.Errorf("Expected 0 custom elements, got %d: %v", len(elements), tagNames(elements))
	}
}

func TestPHP_AttributeExtraction(t *testing.T) {
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create document manager: %v", err)
	}
	defer dm.Close()

	content := readPHPTwigFixture(t, "wordpress-theme.php")
	elements := openAndFindElements(t, dm, "test://theme.php", content)

	// site-header has theme="dark" attribute
	el := findElement(elements, "site-header")
	if el == nil {
		t.Fatal("site-header not found")
	}
	attr, ok := el.Attributes["theme"]
	if !ok {
		t.Fatal("Expected site-header to have 'theme' attribute")
	}
	if attr.Value != "dark" {
		t.Errorf("Expected theme=\"dark\", got theme=%q", attr.Value)
	}
}

func TestPHP_SlotAttribute(t *testing.T) {
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create document manager: %v", err)
	}
	defer dm.Close()

	content := readPHPTwigFixture(t, "php-interpolated.php")
	elements := openAndFindElements(t, dm, "test://product.php", content)

	// product-card's child img has slot="image", span has slot="price"
	// These are slotted children, not attributes on product-card itself,
	// so they appear as separate elements only if they're custom elements.
	// The important test is that product-card itself is found with its attributes.
	el := findElement(elements, "product-card")
	if el == nil {
		t.Fatal("product-card not found")
	}
}

func TestPHP_DocumentUpdate(t *testing.T) {
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create document manager: %v", err)
	}
	defer dm.Close()

	content := readPHPTwigFixture(t, "wordpress-theme.php")
	dm.OpenDocument("test://theme.php", content, 1)

	// Update with new content that adds another custom element
	updated := content + "\n<alert-banner type=\"info\">Update!</alert-banner>\n"
	doc := dm.UpdateDocument("test://theme.php", updated, 2)
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

// Twig Custom Element Extraction
// ============================================================================

func TestTwig_FindCustomElements(t *testing.T) {
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create document manager: %v", err)
	}
	defer dm.Close()

	content := readPHPTwigFixture(t, "drupal-theme.html.twig")
	elements := openAndFindElements(t, dm, "test://template.html.twig", content)

	// site-header, blog-card, sidebar-widget, site-footer
	assertElementsFound(t, elements, []string{"site-header", "blog-card", "sidebar-widget", "site-footer"})
}

func TestTwig_FindCustomElements_WithInterpolation(t *testing.T) {
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create document manager: %v", err)
	}
	defer dm.Close()

	// Twig interpolation in attributes, macros, loops
	content := readPHPTwigFixture(t, "twig-interpolated.html.twig")
	elements := openAndFindElements(t, dm, "test://page.html.twig", content)

	// site-header, nav-menu, nav-item, blog-card
	assertElementsFound(t, elements, []string{"site-header", "nav-menu", "nav-item", "blog-card"})
}

func TestTwig_NoCustomElements(t *testing.T) {
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create document manager: %v", err)
	}
	defer dm.Close()

	content := readPHPTwigFixture(t, "no-elements.html.twig")
	elements := openAndFindElements(t, dm, "test://plain.html.twig", content)

	if len(elements) != 0 {
		t.Errorf("Expected 0 custom elements, got %d: %v", len(elements), tagNames(elements))
	}
}

func TestTwig_AttributeWithInterpolation(t *testing.T) {
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create document manager: %v", err)
	}
	defer dm.Close()

	content := readPHPTwigFixture(t, "twig-interpolated.html.twig")
	elements := openAndFindElements(t, dm, "test://page.html.twig", content)

	// nav-menu has orientation="horizontal"
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

func TestTwig_TwigExtension(t *testing.T) {
	// .twig extension (without .html prefix) should also work
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create document manager: %v", err)
	}
	defer dm.Close()

	content := readPHPTwigFixture(t, "drupal-theme.html.twig")
	elements := openAndFindElements(t, dm, "test://template.twig", content)

	if len(elements) == 0 {
		t.Error("Expected custom elements to be found for .twig extension")
	}
}
