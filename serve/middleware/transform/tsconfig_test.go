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
	"strings"
	"testing"
)

// TestTransformTypeScript_WithTsconfigRaw tests that custom tsconfig settings are respected
func TestTransformTypeScript_WithTsconfigRaw(t *testing.T) {
	source := `
class MyElement {
	@property({ type: String })
	name = 'default';
}
`

	// Transform with experimentalDecorators enabled via TsconfigRaw
	result, err := TransformTypeScript([]byte(source), TransformOptions{
		Loader:    LoaderTS,
		Target:    ES2020,
		Sourcemap: SourceMapInline,
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
	source := `
async function test() {
	const obj = { a: 1, b: 2 };
	const { a, ...rest } = obj;
	return rest;
}
`

	// Transform without TsconfigRaw - should use default (importHelpers: false)
	result, err := TransformTypeScript([]byte(source), TransformOptions{
		Loader:    LoaderTS,
		Target:    ES2020,
		Sourcemap: SourceMapInline,
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
