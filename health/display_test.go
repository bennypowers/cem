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
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestWriteMarkdownReport(t *testing.T) {
	tests := []struct {
		name    string
		fixture string
	}{
		{"well-documented", "well-documented"},
		{"undocumented", "undocumented"},
		{"partial-docs", "partial-docs"},
		{"no-declarations", "no-declarations"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fixturePath := filepath.Join("testdata", "fixtures", tt.fixture, "custom-elements.json")
			goldenPath := filepath.Join("testdata", "goldens", "markdown-"+tt.fixture+".md")

			result, err := Analyze(fixturePath, HealthOptions{})
			if err != nil {
				t.Fatalf("Analyze() error = %v", err)
			}

			var buf bytes.Buffer
			if err := writeMarkdownReport(&buf, result); err != nil {
				t.Fatalf("writeMarkdownReport() error = %v", err)
			}
			got := buf.Bytes()

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
				t.Errorf("writeMarkdownReport() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
