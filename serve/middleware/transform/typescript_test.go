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

	"bennypowers.dev/cem/internal/platform/testutil"
	"bennypowers.dev/cem/serve/middleware/transform"
)

// TestTypeScriptTransform_Basic tests basic TypeScript to JavaScript transformation
func TestTypeScriptTransform_Basic(t *testing.T) {
	// Read input TypeScript file from fixture
	input := testutil.LoadFixtureFile(t, "transforms/basic/input.ts")

	// Transform TypeScript to JavaScript
	result, err := transform.TransformTypeScript(input, transform.TransformOptions{
		Loader:    transform.LoaderTS,
		Target:    transform.ES2020,
		Sourcemap: transform.SourceMapInline,
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
