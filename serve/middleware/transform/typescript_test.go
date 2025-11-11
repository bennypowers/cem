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

package transform

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestTypeScriptTransform_Basic tests basic TypeScript to JavaScript transformation
func TestTypeScriptTransform_Basic(t *testing.T) {
	// Read input TypeScript file
	inputPath := filepath.Join("..", "..", "testdata", "transforms", "basic", "input.ts")
	input, err := os.ReadFile(inputPath)
	if err != nil {
		t.Fatalf("Failed to read input file: %v", err)
	}

	// Transform TypeScript to JavaScript
	result, err := TransformTypeScript(input, TransformOptions{
		Loader:    LoaderTS,
		Target:    ES2020,
		Sourcemap: SourceMapInline,
	})
	if err != nil {
		t.Fatalf("Transform failed: %v", err)
	}

	// Verify output is valid JavaScript
	if len(result.Code) == 0 {
		t.Fatal("Transform produced empty output")
	}

	// Verify TypeScript-specific syntax is removed
	output := string(result.Code)
	if strings.Contains(output, ": string") || strings.Contains(output, ": number") {
		t.Error("Output still contains TypeScript type annotations")
	}

	// Verify inline source map is present
	if !strings.Contains(output, "//# sourceMappingURL=") {
		t.Error("Output missing inline source map")
	}
}
