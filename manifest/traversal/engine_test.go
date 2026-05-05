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
package traversal

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConvertMapAccessToGjsonPath(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "simple map access",
			input:    "elements[button-element]",
			expected: "elements.button-element",
		},
		{
			name:     "no brackets",
			input:    "elements.name",
			expected: "elements.name",
		},
		{
			name:     "multiple map accesses",
			input:    "elements[button][variant]",
			expected: "elements.button.variant",
		},
		{
			name:     "preserves filter expression with question mark",
			input:    "elements[?(@.name=='foo')]",
			expected: "elements[?(@.name=='foo')]",
		},
		{
			name:     "preserves filter expression with at sign",
			input:    "elements[@this]",
			expected: "elements[@this]",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "mixed map access and filter",
			input:    "modules[my-module].declarations[?(@.kind=='class')]",
			expected: "modules.my-module.declarations[?(@.kind=='class')]",
		},
		{
			name:     "nested path with map access at end",
			input:    "a.b.c[key]",
			expected: "a.b.c.key",
		},
		{
			name:     "unclosed bracket",
			input:    "elements[broken",
			expected: "elements[broken",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := convertMapAccessToGjsonPath(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestApplyFilter(t *testing.T) {
	engine := NewEngine()

	t.Run("first filter", func(t *testing.T) {
		tests := []struct {
			name     string
			input    any
			expected any
		}{
			{
				name:     "returns first element of slice",
				input:    []any{"alpha", "beta", "gamma"},
				expected: "alpha",
			},
			{
				name:     "returns input for empty slice",
				input:    []any{},
				expected: []any{},
			},
			{
				name:     "returns non-slice as-is",
				input:    "scalar-value",
				expected: "scalar-value",
			},
			{
				name:     "returns first of int slice",
				input:    []any{42, 99},
				expected: 42,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := engine.applyFilter(tt.input, "first")
				assert.Equal(t, tt.expected, result)
			})
		}
	})

	t.Run("count filter", func(t *testing.T) {
		tests := []struct {
			name     string
			input    any
			expected any
		}{
			{
				name:     "counts slice elements",
				input:    []any{"a", "b", "c"},
				expected: 3,
			},
			{
				name:     "counts empty slice as zero",
				input:    []any{},
				expected: 0,
			},
			{
				name:     "returns 1 for non-slice",
				input:    "scalar",
				expected: 1,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := engine.applyFilter(tt.input, "count")
				assert.Equal(t, tt.expected, result)
			})
		}
	})

	t.Run("exists filter", func(t *testing.T) {
		tests := []struct {
			name     string
			input    any
			expected any
		}{
			{
				name:     "returns true for non-nil value",
				input:    "something",
				expected: true,
			},
			{
				name:     "returns false for nil",
				input:    nil,
				expected: false,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := engine.applyFilter(tt.input, "exists")
				assert.Equal(t, tt.expected, result)
			})
		}
	})

	t.Run("unknown filter returns input as-is", func(t *testing.T) {
		input := map[string]any{"key": "value"}
		result := engine.applyFilter(input, "nonexistent")
		assert.Equal(t, input, result)
	})
}

func TestResolvePath(t *testing.T) {
	engine := NewEngine()

	sources := map[string]any{
		"manifest": map[string]any{
			"schemaVersion": "2.1.0",
			"modules": []any{
				map[string]any{
					"kind": "javascript-module",
					"path": "src/button.js",
					"declarations": []any{
						map[string]any{
							"kind":    "class",
							"name":    "MyButton",
							"tagName": "my-button",
						},
					},
				},
			},
		},
		"elements": map[string]any{
			"button-element": map[string]any{
				"tagName":    "my-button",
				"attributes": []any{"disabled", "variant"},
			},
		},
	}

	t.Run("returns entire source when path is empty", func(t *testing.T) {
		result, err := engine.ResolvePath("manifest", "", sources)
		require.NoError(t, err)
		assert.NotNil(t, result)
		m, ok := result.(map[string]any)
		require.True(t, ok)
		assert.Equal(t, "2.1.0", m["schemaVersion"])
	})

	t.Run("resolves dot-path into nested object", func(t *testing.T) {
		result, err := engine.ResolvePath("manifest", "schemaVersion", sources)
		require.NoError(t, err)
		assert.Equal(t, "2.1.0", result)
	})

	t.Run("resolves array index", func(t *testing.T) {
		result, err := engine.ResolvePath("manifest", "modules.0.path", sources)
		require.NoError(t, err)
		assert.Equal(t, "src/button.js", result)
	})

	t.Run("resolves nested path through array", func(t *testing.T) {
		result, err := engine.ResolvePath("manifest", "modules.0.declarations.0.name", sources)
		require.NoError(t, err)
		assert.Equal(t, "MyButton", result)
	})

	t.Run("returns error for missing source", func(t *testing.T) {
		_, err := engine.ResolvePath("nonexistent", "foo", sources)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "source 'nonexistent' not found")
	})

	t.Run("returns error for missing path", func(t *testing.T) {
		_, err := engine.ResolvePath("manifest", "does.not.exist", sources)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "path 'does.not.exist' not found")
	})

	t.Run("resolves map access with bracket syntax", func(t *testing.T) {
		result, err := engine.ResolvePath("elements", "button-element.tagName", sources)
		require.NoError(t, err)
		assert.Equal(t, "my-button", result)
	})

	t.Run("resolves bracket syntax via fallback conversion", func(t *testing.T) {
		// This path uses bracket syntax that gjson won't resolve directly,
		// triggering the convertMapAccessToGjsonPath fallback.
		result, err := engine.ResolvePath("elements", "button-element[tagName]", sources)
		require.NoError(t, err)
		assert.Equal(t, "my-button", result)
	})

	t.Run("returns error for bracket path that still fails after conversion", func(t *testing.T) {
		_, err := engine.ResolvePath("elements", "nonexistent[key]", sources)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "not found in source")
	})
}

func TestExecuteDataFetcher(t *testing.T) {
	engine := NewEngine()

	sources := map[string]any{
		"manifest": map[string]any{
			"modules": []any{
				map[string]any{"path": "src/a.js"},
				map[string]any{"path": "src/b.js"},
				map[string]any{"path": "src/c.js"},
			},
		},
	}

	t.Run("fetches without filter", func(t *testing.T) {
		fetcher := DataFetcher{
			Name:   "module-path",
			Source: "manifest",
			Path:   "modules.0.path",
		}
		result, err := engine.ExecuteDataFetcher(fetcher, sources)
		require.NoError(t, err)
		assert.Equal(t, "src/a.js", result)
	})

	t.Run("fetches with first filter", func(t *testing.T) {
		fetcher := DataFetcher{
			Name:   "first-module",
			Source: "manifest",
			Path:   "modules",
			Filter: "first",
		}
		result, err := engine.ExecuteDataFetcher(fetcher, sources)
		require.NoError(t, err)
		assert.NotNil(t, result)
		m, ok := result.(map[string]any)
		require.True(t, ok)
		assert.Equal(t, "src/a.js", m["path"])
	})

	t.Run("fetches with count filter", func(t *testing.T) {
		fetcher := DataFetcher{
			Name:   "module-count",
			Source: "manifest",
			Path:   "modules",
			Filter: "count",
		}
		result, err := engine.ExecuteDataFetcher(fetcher, sources)
		require.NoError(t, err)
		assert.Equal(t, 3, result)
	})

	t.Run("fetches with exists filter", func(t *testing.T) {
		fetcher := DataFetcher{
			Name:   "modules-exist",
			Source: "manifest",
			Path:   "modules",
			Filter: "exists",
		}
		result, err := engine.ExecuteDataFetcher(fetcher, sources)
		require.NoError(t, err)
		assert.Equal(t, true, result)
	})

	t.Run("returns error for missing source", func(t *testing.T) {
		fetcher := DataFetcher{
			Name:   "bad-source",
			Source: "missing",
			Path:   "foo",
		}
		_, err := engine.ExecuteDataFetcher(fetcher, sources)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "source 'missing' not found")
	})

	t.Run("returns error for missing path", func(t *testing.T) {
		fetcher := DataFetcher{
			Name:   "bad-path",
			Source: "manifest",
			Path:   "nonexistent.deep.path",
		}
		_, err := engine.ExecuteDataFetcher(fetcher, sources)
		require.Error(t, err)
	})
}

func TestNewEngine(t *testing.T) {
	engine := NewEngine()
	assert.NotNil(t, engine)
}
