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

package transform_test

import (
	"strings"
	"testing"

	"bennypowers.dev/cem/serve/middleware/transform"
)

// TestRewriteImportAttributes_SingleCSSImport tests rewriting a single CSS import with type attribute
func TestRewriteImportAttributes_SingleCSSImport(t *testing.T) {
	input := []byte(`import styles from './foo.css' with { type: 'css' };`)

	result, err := transform.RewriteImportAttributes(input)
	if err != nil {
		t.Fatalf("RewriteImportAttributes failed: %v", err)
	}

	output := string(result)

	// Should rewrite to include query parameter
	if !strings.Contains(output, "__cem-import-attrs[type]=css") {
		t.Errorf("Expected query parameter in output, got: %s", output)
	}

	// Should preserve the rest of the import statement
	if !strings.Contains(output, "import styles from") {
		t.Errorf("Import statement structure changed: %s", output)
	}
}

// TestRewriteImportAttributes_MultipleAttributes tests import with multiple attributes
func TestRewriteImportAttributes_MultipleAttributes(t *testing.T) {
	input := []byte(`import data from './data.json' with { type: 'json', integrity: 'sha256-abc' };`)

	result, err := transform.RewriteImportAttributes(input)
	if err != nil {
		t.Fatalf("RewriteImportAttributes failed: %v", err)
	}

	output := string(result)

	// Should include both attributes as query parameters
	if !strings.Contains(output, "__cem-import-attrs[type]=json") {
		t.Errorf("Expected type attribute in output, got: %s", output)
	}
	if !strings.Contains(output, "__cem-import-attrs[integrity]=") {
		t.Errorf("Expected integrity attribute in output, got: %s", output)
	}
}

// TestRewriteImportAttributes_NoAttributes tests import without attributes remains unchanged
func TestRewriteImportAttributes_NoAttributes(t *testing.T) {
	input := []byte(`import styles from './foo.css';`)

	result, err := transform.RewriteImportAttributes(input)
	if err != nil {
		t.Fatalf("RewriteImportAttributes failed: %v", err)
	}

	output := string(result)

	// Should not add query parameters
	if strings.Contains(output, "__cem-import-attrs") {
		t.Errorf("Unexpected query parameter in output: %s", output)
	}

	// Output should match input
	if output != string(input) {
		t.Errorf("Expected unchanged output, got: %s", output)
	}
}

// TestRewriteImportAttributes_MultipleImports tests multiple imports in the same file
func TestRewriteImportAttributes_MultipleImports(t *testing.T) {
	input := []byte(`import styles from './foo.css' with { type: 'css' };
import data from './data.json' with { type: 'json' };
import regular from './regular.js';`)

	result, err := transform.RewriteImportAttributes(input)
	if err != nil {
		t.Fatalf("RewriteImportAttributes failed: %v", err)
	}

	output := string(result)

	// Should rewrite both imports with attributes
	if !strings.Contains(output, "./foo.css?__cem-import-attrs[type]=css") {
		t.Errorf("Expected rewritten CSS import, got: %s", output)
	}
	if !strings.Contains(output, "./data.json?__cem-import-attrs[type]=json") {
		t.Errorf("Expected rewritten JSON import, got: %s", output)
	}

	// Should not rewrite regular import
	if strings.Contains(output, "./regular.js?") {
		t.Errorf("Regular import should not be rewritten: %s", output)
	}
}

// TestRewriteImportAttributes_DoubleQuotes tests import with double quotes
func TestRewriteImportAttributes_DoubleQuotes(t *testing.T) {
	input := []byte(`import styles from "./foo.css" with { type: "css" };`)

	result, err := transform.RewriteImportAttributes(input)
	if err != nil {
		t.Fatalf("RewriteImportAttributes failed: %v", err)
	}

	output := string(result)

	// Should handle double quotes correctly
	if !strings.Contains(output, "__cem-import-attrs[type]=css") {
		t.Errorf("Expected query parameter in output, got: %s", output)
	}
}

// TestRewriteImportAttributes_CompleteFile tests rewriting in a complete TypeScript file
func TestRewriteImportAttributes_CompleteFile(t *testing.T) {
	input := []byte("import { LitElement, html } from 'lit';\n" +
		"import styles from './component.css' with { type: 'css' };\n\n" +
		"export class MyElement extends LitElement {\n" +
		"  static styles = styles;\n\n" +
		"  render() {\n" +
		"    return html`<div>Hello</div>`;\n" +
		"  }\n" +
		"}")

	result, err := transform.RewriteImportAttributes(input)
	if err != nil {
		t.Fatalf("RewriteImportAttributes failed: %v", err)
	}

	output := string(result)

	// Should rewrite the CSS import
	if !strings.Contains(output, "./component.css?__cem-import-attrs[type]=css") {
		t.Errorf("Expected rewritten CSS import, got: %s", output)
	}

	// Should preserve other imports
	if !strings.Contains(output, "from 'lit'") {
		t.Errorf("Regular import modified: %s", output)
	}

	// Should preserve class definition
	if !strings.Contains(output, "export class MyElement") {
		t.Errorf("Class definition modified: %s", output)
	}
}

// TestRewriteImportAttributes_EmptySource tests handling of empty source
func TestRewriteImportAttributes_EmptySource(t *testing.T) {
	input := []byte(``)

	result, err := transform.RewriteImportAttributes(input)
	if err != nil {
		t.Fatalf("RewriteImportAttributes failed: %v", err)
	}

	if len(result) != 0 {
		t.Errorf("Expected empty output, got: %s", string(result))
	}
}
