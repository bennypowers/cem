/*
Copyright © 2026 Benny Powers

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

import (
	"testing"

	M "bennypowers.dev/cem/manifest"
	"github.com/stretchr/testify/assert"
)

// Inline: pure function, table-driven

func TestExportKeySet(t *testing.T) {
	t.Run("empty exports", func(t *testing.T) {
		result := exportKeySet(nil)
		assert.Empty(t, result)
	})

	t.Run("javascript exports", func(t *testing.T) {
		exports := []M.Export{
			&M.JavaScriptExport{Kind: "js", Name: "MyClass"},
		}
		result := exportKeySet(exports)
		assert.True(t, result[exportKey{"js", "MyClass"}])
		assert.Len(t, result, 1)
	})

	t.Run("custom element exports", func(t *testing.T) {
		ref := M.NewReference("MyElement", "", "my-element.js")
		exports := []M.Export{
			M.NewCustomElementExport("my-element", ref, 0, nil),
		}
		result := exportKeySet(exports)
		assert.True(t, result[exportKey{"custom-element-definition", "my-element"}])
		assert.Len(t, result, 1)
	})

	t.Run("mixed export types", func(t *testing.T) {
		ref := M.NewReference("MyElement", "", "my-element.js")
		exports := []M.Export{
			&M.JavaScriptExport{Kind: "js", Name: "MyClass"},
			M.NewCustomElementExport("my-element", ref, 0, nil),
		}
		result := exportKeySet(exports)
		assert.Len(t, result, 2)
		assert.True(t, result[exportKey{"js", "MyClass"}])
		assert.True(t, result[exportKey{"custom-element-definition", "my-element"}])
	})

	t.Run("duplicate names different kinds", func(t *testing.T) {
		ref := M.NewReference("MyElement", "", "my-element.js")
		exports := []M.Export{
			&M.JavaScriptExport{Kind: "js", Name: "MyElement"},
			M.NewCustomElementExport("MyElement", ref, 0, nil),
		}
		result := exportKeySet(exports)
		assert.Len(t, result, 2)
		assert.True(t, result[exportKey{"js", "MyElement"}])
		assert.True(t, result[exportKey{"custom-element-definition", "MyElement"}])
	})
}

