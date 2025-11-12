//go:build ignore
// +build ignore

// This file is intentionally excluded from the build because Phase 5 is on pause.
// GenerateKnobsHTML() is not yet implemented and the required testdata/ fixtures
// do not exist. Remove this build tag when implementing Phase 5b (Advanced Knobs).

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

// TestGenerateKnobs runs table-driven tests for knob generation
func TestGenerateKnobs(t *testing.T) {
	tests := []struct {
		name         string
		testdataDir  string
		skipManifest bool
	}{
		{
			name:        "BooleanAttribute",
			testdataDir: "boolean-attribute",
		},
		{
			name:        "StringAttribute",
			testdataDir: "string-attribute",
		},
		{
			name:        "NumberAttribute",
			testdataDir: "number-attribute",
		},
		{
			name:        "EnumAttribute",
			testdataDir: "enum-attribute",
		},
		{
			name:         "EmptyState",
			testdataDir:  "empty-state",
			skipManifest: true,
		},
		{
			name:        "FirstElementOnly",
			testdataDir: "first-element-only",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testKnobGeneration(t, tt.testdataDir, tt.skipManifest)
		})
	}
}

// testKnobGeneration is a shared test helper for all knob generation tests
func testKnobGeneration(t *testing.T, testdataDir string, skipManifest bool) {
	var manifestJSON []byte
	if !skipManifest {
		manifestPath := filepath.Join("testdata", testdataDir, "manifest.json")
		var err error
		manifestJSON, err = os.ReadFile(manifestPath)
		if err != nil {
			t.Fatalf("Failed to read manifest fixture: %v", err)
		}
	}

	demoPath := filepath.Join("testdata", testdataDir, "demo.html")
	demoHTML, err := os.ReadFile(demoPath)
	if err != nil {
		t.Fatalf("Failed to read demo fixture: %v", err)
	}

	knobsHTML, err := knobs.GenerateKnobsHTML(manifestJSON, demoHTML)
	if err != nil {
		t.Fatalf("Failed to generate knobs: %v", err)
	}

	goldenPath := filepath.Join("testdata", testdataDir, "expected.html")

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
