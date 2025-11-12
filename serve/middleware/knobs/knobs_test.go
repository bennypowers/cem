//go:build ignore
// +build ignore

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

package knobs_test

import (
	"flag"
	"os"
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/serve/middleware/knobs"
)

var update = flag.Bool("update", false, "update golden files")

// TestGenerateKnobs_BooleanAttribute tests knob generation for boolean attributes
func TestGenerateKnobs_BooleanAttribute(t *testing.T) {
	manifestPath := filepath.Join("testdata", "boolean-attribute", "manifest.json")
	manifestJSON, err := os.ReadFile(manifestPath)
	if err != nil {
		t.Fatalf("Failed to read manifest fixture: %v", err)
	}

	demoPath := filepath.Join("testdata", "boolean-attribute", "demo.html")
	demoHTML, err := os.ReadFile(demoPath)
	if err != nil {
		t.Fatalf("Failed to read demo fixture: %v", err)
	}

	// Generate knobs HTML from manifest and demo
	knobsHTML, err := knobs.GenerateKnobsHTML(manifestJSON, demoHTML)
	if err != nil {
		t.Fatalf("Failed to generate knobs: %v", err)
	}

	// Compare with golden file
	goldenPath := filepath.Join("testdata", "boolean-attribute", "expected.html")

	if *update {
		err := os.WriteFile(goldenPath, []byte(knobsHTML), 0644)
		if err != nil {
			t.Fatalf("Failed to update golden file: %v", err)
		}
		t.Log("Updated golden file:", goldenPath)
		return
	}

	expected, err := os.ReadFile(goldenPath)
	if err != nil {
		t.Fatalf("Failed to read golden file: %v", err)
	}

	if knobsHTML != string(expected) {
		t.Errorf("Generated knobs do not match golden file.\nGot:\n%s\n\nExpected:\n%s", knobsHTML, string(expected))
		t.Log("Run 'go test ./serve/middleware/knobs/... --update' to update golden files")
	}
}

// TestGenerateKnobs_StringAttribute tests knob generation for string attributes
func TestGenerateKnobs_StringAttribute(t *testing.T) {
	manifestPath := filepath.Join("testdata", "string-attribute", "manifest.json")
	manifestJSON, err := os.ReadFile(manifestPath)
	if err != nil {
		t.Fatalf("Failed to read manifest fixture: %v", err)
	}

	demoPath := filepath.Join("testdata", "string-attribute", "demo.html")
	demoHTML, err := os.ReadFile(demoPath)
	if err != nil {
		t.Fatalf("Failed to read demo fixture: %v", err)
	}

	knobsHTML, err := knobs.GenerateKnobsHTML(manifestJSON, demoHTML)
	if err != nil {
		t.Fatalf("Failed to generate knobs: %v", err)
	}

	goldenPath := filepath.Join("testdata", "string-attribute", "expected.html")

	if *update {
		err := os.WriteFile(goldenPath, []byte(knobsHTML), 0644)
		if err != nil {
			t.Fatalf("Failed to update golden file: %v", err)
		}
		t.Log("Updated golden file:", goldenPath)
		return
	}

	expected, err := os.ReadFile(goldenPath)
	if err != nil {
		t.Fatalf("Failed to read golden file: %v", err)
	}

	if knobsHTML != string(expected) {
		t.Errorf("Generated knobs do not match golden file.\nGot:\n%s\n\nExpected:\n%s", knobsHTML, string(expected))
		t.Log("Run 'go test ./serve/middleware/knobs/... --update' to update golden files")
	}
}

// TestGenerateKnobs_NumberAttribute tests knob generation for number attributes
func TestGenerateKnobs_NumberAttribute(t *testing.T) {
	manifestPath := filepath.Join("testdata", "number-attribute", "manifest.json")
	manifestJSON, err := os.ReadFile(manifestPath)
	if err != nil {
		t.Fatalf("Failed to read manifest fixture: %v", err)
	}

	demoPath := filepath.Join("testdata", "number-attribute", "demo.html")
	demoHTML, err := os.ReadFile(demoPath)
	if err != nil {
		t.Fatalf("Failed to read demo fixture: %v", err)
	}

	knobsHTML, err := knobs.GenerateKnobsHTML(manifestJSON, demoHTML)
	if err != nil {
		t.Fatalf("Failed to generate knobs: %v", err)
	}

	goldenPath := filepath.Join("testdata", "number-attribute", "expected.html")

	if *update {
		err := os.WriteFile(goldenPath, []byte(knobsHTML), 0644)
		if err != nil {
			t.Fatalf("Failed to update golden file: %v", err)
		}
		t.Log("Updated golden file:", goldenPath)
		return
	}

	expected, err := os.ReadFile(goldenPath)
	if err != nil {
		t.Fatalf("Failed to read golden file: %v", err)
	}

	if knobsHTML != string(expected) {
		t.Errorf("Generated knobs do not match golden file.\nGot:\n%s\n\nExpected:\n%s", knobsHTML, string(expected))
		t.Log("Run 'go test ./serve/middleware/knobs/... --update' to update golden files")
	}
}

// TestGenerateKnobs_EnumAttribute tests knob generation for enum/union type attributes
func TestGenerateKnobs_EnumAttribute(t *testing.T) {
	manifestPath := filepath.Join("testdata", "enum-attribute", "manifest.json")
	manifestJSON, err := os.ReadFile(manifestPath)
	if err != nil {
		t.Fatalf("Failed to read manifest fixture: %v", err)
	}

	demoPath := filepath.Join("testdata", "enum-attribute", "demo.html")
	demoHTML, err := os.ReadFile(demoPath)
	if err != nil {
		t.Fatalf("Failed to read demo fixture: %v", err)
	}

	knobsHTML, err := knobs.GenerateKnobsHTML(manifestJSON, demoHTML)
	if err != nil {
		t.Fatalf("Failed to generate knobs: %v", err)
	}

	goldenPath := filepath.Join("testdata", "enum-attribute", "expected.html")

	if *update {
		err := os.WriteFile(goldenPath, []byte(knobsHTML), 0644)
		if err != nil {
			t.Fatalf("Failed to update golden file: %v", err)
		}
		t.Log("Updated golden file:", goldenPath)
		return
	}

	expected, err := os.ReadFile(goldenPath)
	if err != nil {
		t.Fatalf("Failed to read golden file: %v", err)
	}

	if knobsHTML != string(expected) {
		t.Errorf("Generated knobs do not match golden file.\nGot:\n%s\n\nExpected:\n%s", knobsHTML, string(expected))
		t.Log("Run 'go test ./serve/middleware/knobs/... --update' to update golden files")
	}
}

// TestGenerateKnobs_EmptyState tests empty state when no manifest is available
func TestGenerateKnobs_EmptyState(t *testing.T) {
	// No manifest file - just demo HTML
	demoPath := filepath.Join("testdata", "empty-state", "demo.html")
	demoHTML, err := os.ReadFile(demoPath)
	if err != nil {
		t.Fatalf("Failed to read demo fixture: %v", err)
	}

	knobsHTML, err := knobs.GenerateKnobsHTML(nil, demoHTML)
	if err != nil {
		t.Fatalf("Failed to generate knobs: %v", err)
	}

	goldenPath := filepath.Join("testdata", "empty-state", "expected.html")

	if *update {
		err := os.WriteFile(goldenPath, []byte(knobsHTML), 0644)
		if err != nil {
			t.Fatalf("Failed to update golden file: %v", err)
		}
		t.Log("Updated golden file:", goldenPath)
		return
	}

	expected, err := os.ReadFile(goldenPath)
	if err != nil {
		t.Fatalf("Failed to read golden file: %v", err)
	}

	if knobsHTML != string(expected) {
		t.Errorf("Generated knobs do not match golden file.\nGot:\n%s\n\nExpected:\n%s", knobsHTML, string(expected))
		t.Log("Run 'go test ./serve/middleware/knobs/... --update' to update golden files")
	}
}

// TestGenerateKnobs_FirstElementOnly tests that only the first custom element gets knobs
func TestGenerateKnobs_FirstElementOnly(t *testing.T) {
	manifestPath := filepath.Join("testdata", "first-element-only", "manifest.json")
	manifestJSON, err := os.ReadFile(manifestPath)
	if err != nil {
		t.Fatalf("Failed to read manifest fixture: %v", err)
	}

	// Demo with multiple custom elements
	demoPath := filepath.Join("testdata", "first-element-only", "demo.html")
	demoHTML, err := os.ReadFile(demoPath)
	if err != nil {
		t.Fatalf("Failed to read demo fixture: %v", err)
	}

	knobsHTML, err := knobs.GenerateKnobsHTML(manifestJSON, demoHTML)
	if err != nil {
		t.Fatalf("Failed to generate knobs: %v", err)
	}

	goldenPath := filepath.Join("testdata", "first-element-only", "expected.html")

	if *update {
		err := os.WriteFile(goldenPath, []byte(knobsHTML), 0644)
		if err != nil {
			t.Fatalf("Failed to update golden file: %v", err)
		}
		t.Log("Updated golden file:", goldenPath)
		return
	}

	expected, err := os.ReadFile(goldenPath)
	if err != nil {
		t.Fatalf("Failed to read golden file: %v", err)
	}

	if knobsHTML != string(expected) {
		t.Errorf("Generated knobs do not match golden file.\nGot:\n%s\n\nExpected:\n%s", knobsHTML, string(expected))
		t.Log("Run 'go test ./serve/middleware/knobs/... --update' to update golden files")
	}
}
