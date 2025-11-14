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

package routes

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestTransformTemplateForClient_SimpleInterpolation tests removal of simple {{.Foo}} interpolations
func TestTransformTemplateForClient_SimpleInterpolation(t *testing.T) {
	input := `<button type="{{.Type}}">
  <slot name="icon-start"></slot>
  <slot></slot>
  <slot name="icon-end"></slot>
</button>`

	expected := `<button type="">
  <slot name="icon-start"></slot>
  <slot></slot>
  <slot name="icon-end"></slot>
</button>`

	result := transformTemplateForClient(input)
	if !equalHTML(result, expected) {
		t.Errorf("Simple interpolation transform failed.\nExpected:\n%s\nGot:\n%s", expected, result)
	}
}

// TestTransformTemplateForClient_ControlFlow tests conversion of control flow to comments
func TestTransformTemplateForClient_ControlFlow(t *testing.T) {
	input := `<div>
  {{if .Disabled}}
  <p class="disabled">This is disabled</p>
  {{end}}
  {{range .Items}}
  <li>{{.Name}}</li>
  {{end}}
</div>`

	result := transformTemplateForClient(input)

	// Debug: print result
	t.Logf("Result:\n%s", result)

	// Check for comment markers (preserves original casing from Go template)
	if !strings.Contains(result, "<!-- if:Disabled -->") {
		t.Errorf("Expected '<!-- if:Disabled -->' in result, got: %s", result)
	}
	if !strings.Contains(result, "<!-- range:Items -->") {
		t.Errorf("Expected '<!-- range:Items -->' in result, got: %s", result)
	}
	if !strings.Contains(result, "<!-- end -->") {
		t.Errorf("Expected '<!-- end -->' markers in result, got: %s", result)
	}

	// Ensure original {{...}} syntax is removed
	if strings.Contains(result, "{{if") || strings.Contains(result, "{{range") || strings.Contains(result, "{{end}}") {
		t.Error("Original Go template syntax should be removed")
	}
}

// TestTransformTemplateForClient_PreservesHTML tests that valid HTML structure is preserved
func TestTransformTemplateForClient_PreservesHTML(t *testing.T) {
	input := `<input id="text-input" type="text" placeholder="Enter text">`

	result := transformTemplateForClient(input)

	// Should preserve the input element and attributes
	if !strings.Contains(result, `id="text-input"`) {
		t.Error("Expected id attribute to be preserved")
	}
	if !strings.Contains(result, `type="text"`) {
		t.Error("Expected type attribute to be preserved")
	}
	if !strings.Contains(result, `placeholder="Enter text"`) {
		t.Error("Expected placeholder attribute to be preserved")
	}
}

// TestTransformTemplateForClient_ComplexTemplate uses fixture/golden pattern
func TestTransformTemplateForClient_ComplexTemplate(t *testing.T) {
	fixturePath := filepath.Join("testdata", "transform", "complex-template.html")
	goldenPath := filepath.Join("testdata", "transform", "expected-complex-template.html")

	// Read fixture
	input, err := os.ReadFile(fixturePath)
	if err != nil {
		// If fixture doesn't exist yet, create it for the test
		t.Skipf("Fixture not found: %s (run with --update to create)", fixturePath)
	}

	result := transformTemplateForClient(string(input))

	// Read or update golden file
	if *update {
		if err := os.MkdirAll(filepath.Dir(goldenPath), 0755); err != nil {
			t.Fatalf("Failed to create testdata directory: %v", err)
		}
		if err := os.WriteFile(goldenPath, []byte(result), 0644); err != nil {
			t.Fatalf("Failed to update golden file: %v", err)
		}
	}

	expected, err := os.ReadFile(goldenPath)
	if err != nil {
		t.Fatalf("Failed to read golden file: %v", err)
	}

	if !equalHTML(string(expected), result) {
		t.Errorf("Complex template transform failed.\nExpected:\n%s\nGot:\n%s", string(expected), result)
	}
}

// TestTransformTemplateForClient_EmptyTemplate tests empty input
func TestTransformTemplateForClient_EmptyTemplate(t *testing.T) {
	input := ""
	result := transformTemplateForClient(input)

	if result != input {
		t.Errorf("Empty template should remain empty, got: %s", result)
	}
}

// TestTransformTemplateForClient_MalformedHTML tests handling of malformed HTML
func TestTransformTemplateForClient_MalformedHTML(t *testing.T) {
	input := `<div><p>Unclosed paragraph`

	// Should not crash, should return something (original or fixed HTML)
	result := transformTemplateForClient(input)

	if result == "" {
		t.Error("Malformed HTML should not result in empty string")
	}
}

// equalHTML compares two HTML strings ignoring whitespace differences
func equalHTML(a, b string) bool {
	return strings.TrimSpace(a) == strings.TrimSpace(b)
}
