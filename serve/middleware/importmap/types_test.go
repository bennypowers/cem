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
package importmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Inline: pure function, scalar assertions

func TestImportMap_IsImportMap(t *testing.T) {
	im := &ImportMap{}
	im.IsImportMap()
}

func TestImportMap_ToJSON(t *testing.T) {
	t.Run("nil returns empty", func(t *testing.T) {
		var im *ImportMap
		assert.Equal(t, "", im.ToJSON())
	})

	t.Run("empty imports returns empty", func(t *testing.T) {
		im := &ImportMap{}
		assert.Equal(t, "", im.ToJSON())
	})

	t.Run("with imports returns JSON", func(t *testing.T) {
		im := &ImportMap{
			Imports: map[string]string{"lit": "/node_modules/lit/index.js"},
		}
		result := im.ToJSON()
		assert.Contains(t, result, "lit")
		assert.Contains(t, result, "/node_modules/lit/index.js")
	})
}

func TestDependencyGraph_Dependents(t *testing.T) {
	t.Run("nil graph returns nil", func(t *testing.T) {
		var dg *DependencyGraph
		assert.Nil(t, dg.Dependents("pkg"))
	})

	t.Run("nil inner graph returns nil", func(t *testing.T) {
		dg := &DependencyGraph{}
		assert.Nil(t, dg.Dependents("pkg"))
	})
}

func TestInjectImportMap(t *testing.T) {
	t.Run("replaces existing import map", func(t *testing.T) {
		html := `<head><script type="importmap">{"imports":{}}</script></head>`
		result := injectImportMap(html, `{"imports":{"lit":"/lit.js"}}`)
		assert.Contains(t, result, `{"imports":{"lit":"/lit.js"}}`)
		assert.NotContains(t, result, `{"imports":{}}`)
	})

	t.Run("injects before head close", func(t *testing.T) {
		html := `<html><head><title>Test</title></head><body></body></html>`
		result := injectImportMap(html, `{"imports":{}}`)
		assert.Contains(t, result, `<script type="importmap">`)
		assert.Contains(t, result, `</head>`)
	})
}

func TestInjectImportMapFallback(t *testing.T) {
	t.Run("injects before head close", func(t *testing.T) {
		html := `<html><head></head><body></body></html>`
		result := injectImportMapFallback(html, `{}`)
		assert.Contains(t, result, `<script type="importmap">`)
	})

	t.Run("falls back to body", func(t *testing.T) {
		html := `<html><body>content</body></html>`
		result := injectImportMapFallback(html, `{}`)
		assert.Contains(t, result, `<script type="importmap">`)
	})

	t.Run("falls back to html tag", func(t *testing.T) {
		html := `<html>content</html>`
		result := injectImportMapFallback(html, `{}`)
		assert.Contains(t, result, `<script type="importmap">`)
	})

	t.Run("no suitable location returns unchanged", func(t *testing.T) {
		html := `just text`
		result := injectImportMapFallback(html, `{}`)
		assert.Equal(t, html, result)
	})
}
