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
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/internal/platform/testutil"
)

func TestWriteMarkdownReport(t *testing.T) {
	mfs := testutil.LoadTestdataFS(t, "testdata", "/")

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

			result, err := Analyze(fixturePath, Options{})
			if err != nil {
				t.Fatalf("Analyze() error = %v", err)
			}

			var buf bytes.Buffer
			if err := writeMarkdownReport(&buf, result); err != nil {
				t.Fatalf("writeMarkdownReport() error = %v", err)
			}
			got := buf.Bytes()

			testutil.CheckGolden(t, "markdown-"+tt.fixture+".md", got, testutil.GoldenOptions{
				Dir: "goldens",
				FS:  mfs,
			})
		})
	}
}
