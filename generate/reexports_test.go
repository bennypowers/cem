/*
Copyright © 2025 Benny Powers

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
package generate

import "testing"

func TestIsTypeOnlyImport(t *testing.T) {
	tests := []struct {
		name string
		stmt string
		want bool
	}{
		// Type-only (should be filtered)
		{"import type named", "import type { Foo } from './foo.js';", true},
		{"import type named alias", "import type { Foo as Bar } from './foo.js';", true},
		{"import type namespace", "import type * as FooNS from './foo.js';", true},
		{"import type specifier", "import { type Foo } from './foo.js';", true},
		{"import type specifier alias", "import { type Foo as Bar } from './foo.js';", true},
		{"import type specifier multi", "import { type Foo, type Bar } from './foo.js';", true},

		// Value (should NOT be filtered)
		{"side-effect import", "import './foo.js';", false},
		{"named import", "import { Foo } from './foo.js';", false},
		{"named import alias", "import { Foo as Bar } from './foo.js';", false},
		{"namespace import", "import * as FooNS from './foo.js';", false},
		{"mixed import", "import { Foo, type Bar } from './foo.js';", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isTypeOnlyImport(tt.stmt); got != tt.want {
				t.Errorf("isTypeOnlyImport(%q) = %v, want %v", tt.stmt, got, tt.want)
			}
		})
	}
}

func TestIsTypeOnlyExport(t *testing.T) {
	tests := []struct {
		name string
		stmt string
		want bool
	}{
		// Type-only (should be filtered)
		{"export type named", "export type { Foo } from './foo.js';", true},
		{"export type named alias", "export type { Foo as Bar } from './foo.js';", true},
		{"export type specifier", "export { type Foo } from './foo.js';", true},
		{"export type specifier alias", "export { type Foo as Bar } from './foo.js';", true},
		{"export type specifier multi", "export { type Foo, type Bar } from './foo.js';", true},

		// Value (should NOT be filtered)
		{"named export", "export { Foo } from './foo.js';", false},
		{"named export alias", "export { Foo as Bar } from './foo.js';", false},
		{"mixed export", "export { Foo, type Bar } from './foo.js';", false},
		{"namespace export", "export * from './foo.js';", false},
		{"namespace alias export", "export * as FooNS from './foo.js';", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isTypeOnlyExport(tt.stmt); got != tt.want {
				t.Errorf("isTypeOnlyExport(%q) = %v, want %v", tt.stmt, got, tt.want)
			}
		})
	}
}
