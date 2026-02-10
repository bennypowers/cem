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
package health

import (
	"encoding/json"
	"flag"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var update = flag.Bool("update", false, "update golden files")

func TestAnalyze(t *testing.T) {
	tests := []struct {
		name    string
		fixture string
	}{
		{"well-documented", "well-documented"},
		{"undocumented", "undocumented"},
		{"partial-docs", "partial-docs"},
		{"no-declarations", "no-declarations"},
		{"rfc2119-keywords", "rfc2119-keywords"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fixturePath := filepath.Join("testdata", "fixtures", tt.fixture, "custom-elements.json")
			goldenPath := filepath.Join("testdata", "goldens", tt.fixture+".json")

			result, err := Analyze(fixturePath, HealthOptions{})
			if err != nil {
				t.Fatalf("Analyze() error = %v", err)
			}

			got, err := json.MarshalIndent(result, "", "  ")
			if err != nil {
				t.Fatalf("json.MarshalIndent() error = %v", err)
			}
			got = append(got, '\n')

			if *update {
				if err := os.MkdirAll(filepath.Dir(goldenPath), 0755); err != nil {
					t.Fatalf("os.MkdirAll() error = %v", err)
				}
				if err := os.WriteFile(goldenPath, got, 0644); err != nil {
					t.Fatalf("os.WriteFile() error = %v", err)
				}
				return
			}

			want, err := os.ReadFile(goldenPath)
			if err != nil {
				t.Fatalf("os.ReadFile(%s) error = %v\nRun tests with --update to generate golden files", goldenPath, err)
			}

			if diff := cmp.Diff(string(want), string(got)); diff != "" {
				t.Errorf("Analyze() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestAnalyze_ComponentFilter(t *testing.T) {
	fixturePath := filepath.Join("testdata", "fixtures", "well-documented", "custom-elements.json")

	t.Run("filter by tag name", func(t *testing.T) {
		result, err := Analyze(fixturePath, HealthOptions{Component: "multi-select"})
		if err != nil {
			t.Fatalf("Analyze() error = %v", err)
		}
		if len(result.Modules) != 1 {
			t.Fatalf("expected 1 module, got %d", len(result.Modules))
		}
		if len(result.Modules[0].Declarations) != 1 {
			t.Errorf("expected 1 declaration, got %d", len(result.Modules[0].Declarations))
		}
		if result.Modules[0].Declarations[0].TagName != "multi-select" {
			t.Errorf("expected tag name 'multi-select', got %q", result.Modules[0].Declarations[0].TagName)
		}
	})

	t.Run("filter by class name", func(t *testing.T) {
		result, err := Analyze(fixturePath, HealthOptions{Component: "MultiSelect"})
		if err != nil {
			t.Fatalf("Analyze() error = %v", err)
		}
		if len(result.Modules) != 1 {
			t.Fatalf("expected 1 module, got %d", len(result.Modules))
		}
		if len(result.Modules[0].Declarations) != 1 {
			t.Errorf("expected 1 declaration, got %d", len(result.Modules[0].Declarations))
		}
	})

	t.Run("filter no match", func(t *testing.T) {
		result, err := Analyze(fixturePath, HealthOptions{Component: "nonexistent"})
		if err != nil {
			t.Fatalf("Analyze() error = %v", err)
		}
		if len(result.Modules) != 0 {
			t.Errorf("expected 0 modules, got %d", len(result.Modules))
		}
	})
}

func TestAnalyze_ModuleFilter(t *testing.T) {
	fixturePath := filepath.Join("testdata", "fixtures", "well-documented", "custom-elements.json")

	t.Run("filter matching module", func(t *testing.T) {
		result, err := Analyze(fixturePath, HealthOptions{Modules: []string{"elements/multi-select/multi-select.js"}})
		if err != nil {
			t.Fatalf("Analyze() error = %v", err)
		}
		if len(result.Modules) != 1 {
			t.Fatalf("expected 1 module, got %d", len(result.Modules))
		}
		if result.Modules[0].Path != "elements/multi-select/multi-select.js" {
			t.Errorf("expected module path 'elements/multi-select/multi-select.js', got %q", result.Modules[0].Path)
		}
	})

	t.Run("filter multiple modules", func(t *testing.T) {
		result, err := Analyze(fixturePath, HealthOptions{
			Modules: []string{
				"elements/multi-select/multi-select.js",
				"nonexistent.js",
			},
		})
		if err != nil {
			t.Fatalf("Analyze() error = %v", err)
		}
		if len(result.Modules) != 1 {
			t.Fatalf("expected 1 module, got %d", len(result.Modules))
		}
		if result.Modules[0].Path != "elements/multi-select/multi-select.js" {
			t.Errorf("expected module path 'elements/multi-select/multi-select.js', got %q", result.Modules[0].Path)
		}
	})

	t.Run("filter no match", func(t *testing.T) {
		result, err := Analyze(fixturePath, HealthOptions{Modules: []string{"nonexistent.js"}})
		if err != nil {
			t.Fatalf("Analyze() error = %v", err)
		}
		if len(result.Modules) != 0 {
			t.Errorf("expected 0 modules, got %d", len(result.Modules))
		}
	})
}

func TestAnalyze_DisableCategories(t *testing.T) {
	fixturePath := filepath.Join("testdata", "fixtures", "well-documented", "custom-elements.json")

	result, err := Analyze(fixturePath, HealthOptions{Disable: []string{"demos"}})
	if err != nil {
		t.Fatalf("Analyze() error = %v", err)
	}

	for _, mod := range result.Modules {
		for _, comp := range mod.Declarations {
			for _, cat := range comp.Categories {
				if cat.ID == "demos" {
					t.Error("demos category should have been disabled")
				}
			}
		}
	}
}
