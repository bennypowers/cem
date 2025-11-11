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

// TestTransformTypeScript_WithTsconfigRaw tests that custom tsconfig settings are respected
func TestTransformTypeScript_WithTsconfigRaw(t *testing.T) {
	// Read fixture file with decorator syntax
	source := testutil.LoadFixtureFile(t, "transforms/tsconfig-decorators/input.ts")

	// Transform with experimentalDecorators enabled via TsconfigRaw
	result, err := transform.TransformTypeScript(source, transform.TransformOptions{
		Loader:    transform.LoaderTS,
		Target:    transform.ES2020,
		Sourcemap: transform.SourceMapInline,
		TsconfigRaw: `{
			"compilerOptions": {
				"experimentalDecorators": true,
				"useDefineForClassFields": false
			}
		}`,
	})

	if err != nil {
		t.Fatalf("Transform failed: %v", err)
	}

	code := string(result.Code)

	// Should NOT contain the @property decorator syntax (it should be transformed)
	if strings.Contains(code, "@property") {
		t.Error("Decorators not transformed - still contains @property syntax")
	}

	// Should contain __decorate or __decorateClass helper (decorator runtime code)
	if !strings.Contains(code, "__decorate") && !strings.Contains(code, "__decorateClass") {
		t.Logf("Output:\n%s", code)
		t.Error("Missing decorator helper function calls in output")
	}
}

// TestTransformTypeScript_DefaultTsconfigDisablesImportHelpers tests the default behavior
func TestTransformTypeScript_DefaultTsconfigDisablesImportHelpers(t *testing.T) {
	// Read fixture file with features that could use tslib
	source := testutil.LoadFixtureFile(t, "transforms/tsconfig-default/input.ts")

	// Transform without TsconfigRaw - should use default (importHelpers: false)
	result, err := transform.TransformTypeScript(source, transform.TransformOptions{
		Loader:    transform.LoaderTS,
		Target:    transform.ES2020,
		Sourcemap: transform.SourceMapInline,
	})

	if err != nil {
		t.Fatalf("Transform failed: %v", err)
	}

	code := string(result.Code)

	// Should NOT import from tslib
	if strings.Contains(code, `from "tslib"`) || strings.Contains(code, `from 'tslib'`) {
		t.Error("Code imports from tslib despite importHelpers: false default")
	}
}
