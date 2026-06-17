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
	"strings"
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
		input := `<html><head><script type="importmap">{"imports":{}}</script></head><body></body></html>`
		result := injectImportMap(input, `{"imports":{"lit":"/lit.js"}}`)
		assert.Contains(t, result, `{"imports":{"lit":"/lit.js"}}`)
		assert.NotContains(t, result, `{"imports":{}}`)
	})

	t.Run("injects before existing scripts", func(t *testing.T) {
		input := `<html><head><title>Test</title><script type="module" src="/app.js"></script></head><body></body></html>`
		result := injectImportMap(input, `{"imports":{}}`)
		importMapIdx := strings.Index(result, `<script type="importmap">`)
		scriptIdx := strings.Index(result, `<script type="module"`)
		assert.Greater(t, importMapIdx, -1, "import map should be present")
		assert.Greater(t, scriptIdx, -1, "module script should be present")
		assert.Less(t, importMapIdx, scriptIdx, "import map must appear before module script")
	})

	t.Run("appends to head when no scripts", func(t *testing.T) {
		input := `<html><head><title>Test</title><link rel="stylesheet" href="/style.css"></head><body></body></html>`
		result := injectImportMap(input, `{"imports":{}}`)
		assert.Contains(t, result, `<script type="importmap">`)
		importMapIdx := strings.Index(result, `<script type="importmap">`)
		headCloseIdx := strings.Index(result, `</head>`)
		assert.Less(t, importMapIdx, headCloseIdx, "import map should be inside head")
	})

	t.Run("replaces existing and keeps before scripts", func(t *testing.T) {
		input := `<html><head><script type="module" src="/a.js"></script><script type="importmap">{"old":true}</script></head><body></body></html>`
		result := injectImportMap(input, `{"imports":{"new":"/new.js"}}`)
		importMapIdx := strings.Index(result, `<script type="importmap">`)
		scriptIdx := strings.Index(result, `<script type="module"`)
		assert.Less(t, importMapIdx, scriptIdx, "import map must appear before module script even after replacement")
		assert.NotContains(t, result, `"old"`)
	})
}

func TestInjectImportMapFallback(t *testing.T) {
	t.Run("injects before first script in head", func(t *testing.T) {
		input := `<html><head><script src="/app.js"></script></head><body></body></html>`
		result := injectImportMapFallback(input, `{}`)
		importMapIdx := strings.Index(result, `<script type="importmap">`)
		scriptIdx := strings.Index(result, `<script src="/app.js">`)
		assert.Greater(t, importMapIdx, -1)
		assert.Less(t, importMapIdx, scriptIdx)
	})

	t.Run("injects before head close when no scripts", func(t *testing.T) {
		input := `<html><head><title>T</title></head><body></body></html>`
		result := injectImportMapFallback(input, `{}`)
		assert.Contains(t, result, `<script type="importmap">`)
	})

	t.Run("falls back to body", func(t *testing.T) {
		input := `<html><body>content</body></html>`
		result := injectImportMapFallback(input, `{}`)
		assert.Contains(t, result, `<script type="importmap">`)
	})

	t.Run("no suitable location returns unchanged", func(t *testing.T) {
		input := `just text`
		result := injectImportMapFallback(input, `{}`)
		assert.Equal(t, input, result)
	})
}
