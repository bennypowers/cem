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
package export

import (
	"encoding/json"
	"flag"
	"os"
	"path/filepath"
	"testing"

	M "bennypowers.dev/cem/manifest"
)

var update = flag.Bool("update", false, "update golden files")

func loadTestManifest(t *testing.T) *M.Package {
	t.Helper()
	data, err := os.ReadFile("testdata/input/custom-elements.json")
	if err != nil {
		t.Fatalf("reading test manifest: %v", err)
	}
	var pkg M.Package
	if err := json.Unmarshal(data, &pkg); err != nil {
		t.Fatalf("parsing test manifest: %v", err)
	}
	return &pkg
}

func TestBuildExportElements(t *testing.T) {
	pkg := loadTestManifest(t)
	elements := buildExportElements(pkg, "my-package")

	if len(elements) != 1 {
		t.Fatalf("expected 1 element, got %d", len(elements))
	}

	elem := elements[0]

	if elem.TagName != "my-button" {
		t.Errorf("TagName = %q, want %q", elem.TagName, "my-button")
	}
	if elem.ClassName != "MyButton" {
		t.Errorf("ClassName = %q, want %q", elem.ClassName, "MyButton")
	}
	if elem.ImportPath != "my-package/elements/my-button.js" {
		t.Errorf("ImportPath = %q, want %q", elem.ImportPath, "my-package/elements/my-button.js")
	}
	if elem.Summary != "A button component" {
		t.Errorf("Summary = %q, want %q", elem.Summary, "A button component")
	}

	// Attributes
	if len(elem.Attributes) != 2 {
		t.Fatalf("expected 2 attributes, got %d", len(elem.Attributes))
	}
	if elem.Attributes[0].Name != "variant" {
		t.Errorf("Attributes[0].Name = %q, want %q", elem.Attributes[0].Name, "variant")
	}
	if elem.Attributes[0].FieldName != "variant" {
		t.Errorf("Attributes[0].FieldName = %q, want %q", elem.Attributes[0].FieldName, "variant")
	}
	if elem.Attributes[1].IsBoolean != true {
		t.Errorf("Attributes[1].IsBoolean = %v, want true", elem.Attributes[1].IsBoolean)
	}

	// Properties (only the non-attribute property "data")
	if len(elem.Properties) != 1 {
		t.Fatalf("expected 1 property, got %d", len(elem.Properties))
	}
	if elem.Properties[0].Name != "data" {
		t.Errorf("Properties[0].Name = %q, want %q", elem.Properties[0].Name, "data")
	}

	// Events
	if len(elem.Events) != 1 {
		t.Fatalf("expected 1 event, got %d", len(elem.Events))
	}
	if elem.Events[0].Name != "my-click" {
		t.Errorf("Events[0].Name = %q, want %q", elem.Events[0].Name, "my-click")
	}
	if elem.Events[0].ReactName != "onMyClick" {
		t.Errorf("Events[0].ReactName = %q, want %q", elem.Events[0].ReactName, "onMyClick")
	}
	if elem.Events[0].AngularName != "myClick" {
		t.Errorf("Events[0].AngularName = %q, want %q", elem.Events[0].AngularName, "myClick")
	}

	// Slots
	if len(elem.Slots) != 2 {
		t.Fatalf("expected 2 slots, got %d", len(elem.Slots))
	}
	if !elem.HasDefaultSlot {
		t.Error("expected HasDefaultSlot = true")
	}
	if !elem.HasNamedSlots {
		t.Error("expected HasNamedSlots = true")
	}

	// CSS Parts
	if len(elem.CssParts) != 1 {
		t.Fatalf("expected 1 CSS part, got %d", len(elem.CssParts))
	}

	// CSS Properties
	if len(elem.CssProperties) != 1 {
		t.Fatalf("expected 1 CSS property, got %d", len(elem.CssProperties))
	}
}

func TestReactExporter(t *testing.T) {
	testExporter(t, &ReactExporter{}, "react")
}

func TestVueExporter(t *testing.T) {
	testExporter(t, &VueExporter{}, "vue")
}

func TestAngularExporter(t *testing.T) {
	testExporter(t, &AngularExporter{}, "angular")
}

func testExporter(t *testing.T, exporter FrameworkExporter, framework string) {
	t.Helper()
	pkg := loadTestManifest(t)
	elements := buildExportElements(pkg, "my-package")
	cfg := FrameworkExportConfig{
		Output:      filepath.Join(t.TempDir(), framework),
		StripPrefix: "my-",
	}

	goldenDir := filepath.Join("testdata", "golden", framework)

	for _, elem := range elements {
		files, err := exporter.ExportElement(elem, cfg)
		if err != nil {
			t.Fatalf("ExportElement(%s): %v", elem.TagName, err)
		}

		for filename, content := range files {
			goldenPath := filepath.Join(goldenDir, filename)
			if *update {
				if err := os.MkdirAll(goldenDir, 0o755); err != nil {
					t.Fatal(err)
				}
				if err := os.WriteFile(goldenPath, []byte(content), 0o644); err != nil {
					t.Fatal(err)
				}
				continue
			}

			expected, err := os.ReadFile(goldenPath)
			if err != nil {
				t.Fatalf("reading golden file %s: %v (run with --update to create)", goldenPath, err)
			}
			if string(expected) != content {
				t.Errorf("%s mismatch.\nExpected:\n%s\nGot:\n%s", filename, string(expected), content)
			}
		}
	}

	// Test index files
	indexFiles, err := exporter.ExportIndex(elements, cfg)
	if err != nil {
		t.Fatalf("ExportIndex: %v", err)
	}

	for filename, content := range indexFiles {
		goldenPath := filepath.Join(goldenDir, filename)
		if *update {
			if err := os.WriteFile(goldenPath, []byte(content), 0o644); err != nil {
				t.Fatal(err)
			}
			continue
		}

		expected, err := os.ReadFile(goldenPath)
		if err != nil {
			t.Fatalf("reading golden file %s: %v (run with --update to create)", goldenPath, err)
		}
		if string(expected) != content {
			t.Errorf("%s mismatch.\nExpected:\n%s\nGot:\n%s", filename, string(expected), content)
		}
	}
}
