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
package resources

import (
	"encoding/json"
	"testing"

	"bennypowers.dev/cem/mcp/types"
)

// TestDataTraversal_BasicPath tests basic path resolution without cross-references
func TestDataTraversal_BasicPath(t *testing.T) {
	tests := []struct {
		name     string
		source   string
		path     string
		data     map[string]any
		args     map[string]any
		expected any
		wantErr  bool
	}{
		{
			name:   "simple property access",
			source: "data",
			path:   "name",
			data: map[string]any{
				"name": "test-element",
				"type": "element",
			},
			args:     map[string]any{},
			expected: "test-element",
			wantErr:  false,
		},
		{
			name:   "nested property access",
			source: "data",
			path:   "properties.version",
			data: map[string]any{
				"properties": map[string]any{
					"version": "1.0.0",
					"author":  "test",
				},
			},
			args:     map[string]any{},
			expected: "1.0.0",
			wantErr:  false,
		},
		{
			name:   "array element access",
			source: "data",
			path:   "attributes[0].name",
			data: map[string]any{
				"attributes": []any{
					map[string]any{"name": "color", "type": "string"},
					map[string]any{"name": "size", "type": "number"},
				},
			},
			args:     map[string]any{},
			expected: "color",
			wantErr:  false,
		},
		{
			name:   "nonexistent path",
			source: "data",
			path:   "missing.property",
			data: map[string]any{
				"name": "test-element",
			},
			args:    map[string]any{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create data sources
			sources := map[string]any{
				"data": tt.data,
				"args": tt.args,
			}

			engine := NewPathTraversalEngine()
			result, err := engine.ResolvePath(tt.source, tt.path, sources)

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

// TestDataTraversal_CrossReference tests path resolution with variable substitution
func TestDataTraversal_CrossReference(t *testing.T) {
	tests := []struct {
		name     string
		source   string
		path     string
		registry map[string]any
		args     map[string]any
		expected any
		wantErr  bool
	}{
		{
			name:   "element lookup by tagName",
			source: "registry",
			path:   "elements[$.tagName]",
			registry: map[string]any{
				"elements": map[string]any{
					"button-element": map[string]any{
						"tagName": "button-element",
						"name":    "Button Element",
					},
					"input-element": map[string]any{
						"tagName": "input-element",
						"name":    "Input Element",
					},
				},
			},
			args: map[string]any{
				"tagName": "button-element",
			},
			expected: map[string]any{
				"tagName": "button-element",
				"name":    "Button Element",
			},
			wantErr: false,
		},
		{
			name:   "nested property access with variable",
			source: "registry",
			path:   "elements[$.tagName].attributes",
			registry: map[string]any{
				"elements": map[string]any{
					"button-element": map[string]any{
						"tagName": "button-element",
						"attributes": []any{
							map[string]any{"name": "color", "type": "string"},
							map[string]any{"name": "disabled", "type": "boolean"},
						},
					},
				},
			},
			args: map[string]any{
				"tagName": "button-element",
			},
			expected: []any{
				map[string]any{"name": "color", "type": "string"},
				map[string]any{"name": "disabled", "type": "boolean"},
			},
			wantErr: false,
		},
		{
			name:   "array filter with variable",
			source: "registry",
			path:   "elements[$.tagName].attributes.#(name==$.attrName)",
			registry: map[string]any{
				"elements": map[string]any{
					"button-element": map[string]any{
						"tagName": "button-element",
						"attributes": []any{
							map[string]any{"name": "color", "type": "string"},
							map[string]any{"name": "disabled", "type": "boolean"},
						},
					},
				},
			},
			args: map[string]any{
				"tagName":  "button-element",
				"attrName": "color",
			},
			expected: map[string]any{"name": "color", "type": "string"},
			wantErr: false,
		},
		{
			name:   "missing variable in args",
			source: "registry",
			path:   "elements[$.tagName]",
			registry: map[string]any{
				"elements": map[string]any{},
			},
			args:    map[string]any{}, // missing tagName
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create data sources
			sources := map[string]any{
				"registry": tt.registry,
				"args":     tt.args,
			}

			engine := NewPathTraversalEngine()
			result, err := engine.ResolvePath(tt.source, tt.path, sources)

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			// Convert to JSON and back for deep comparison
			expectedJSON, _ := json.Marshal(tt.expected)
			resultJSON, _ := json.Marshal(result)

			if string(expectedJSON) != string(resultJSON) {
				t.Errorf("expected %s, got %s", string(expectedJSON), string(resultJSON))
			}
		})
	}
}

// TestDataTraversal_Filtering tests result filtering and transformation
func TestDataTraversal_Filtering(t *testing.T) {
	tests := []struct {
		name     string
		source   string
		path     string
		filter   string
		data     map[string]any
		expected any
		wantErr  bool
	}{
		{
			name:   "first filter on array",
			source: "data",
			path:   "attributes",
			filter: "first",
			data: map[string]any{
				"attributes": []any{
					map[string]any{"name": "color", "type": "string"},
					map[string]any{"name": "disabled", "type": "boolean"},
				},
			},
			expected: map[string]any{"name": "color", "type": "string"},
			wantErr:  false,
		},
		{
			name:   "count filter on array",
			source: "data",
			path:   "attributes",
			filter: "count",
			data: map[string]any{
				"attributes": []any{
					map[string]any{"name": "color"},
					map[string]any{"name": "disabled"},
				},
			},
			expected: 2,
			wantErr:  false,
		},
		{
			name:   "exists filter - true",
			source: "data",
			path:   "attributes",
			filter: "exists",
			data: map[string]any{
				"attributes": []any{
					map[string]any{"name": "color"},
				},
			},
			expected: true,
			wantErr:  false,
		},
		{
			name:   "exists filter - false",
			source: "data",
			path:   "missing",
			filter: "exists",
			data: map[string]any{
				"attributes": []any{},
			},
			expected: false,
			wantErr:  false,
		},
		{
			name:   "first filter on empty array",
			source: "data",
			path:   "attributes",
			filter: "first",
			data: map[string]any{
				"attributes": []any{},
			},
			expected: nil,
			wantErr:  false,
		},
		{
			name:   "invalid filter",
			source: "data",
			path:   "attributes",
			filter: "invalid",
			data: map[string]any{
				"attributes": []any{},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sources := map[string]any{
				"data": tt.data,
			}

			engine := NewPathTraversalEngine()
			result, err := engine.ResolvePathWithFilter(tt.source, tt.path, tt.filter, sources)

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			// Convert to JSON for comparison
			expectedJSON, _ := json.Marshal(tt.expected)
			resultJSON, _ := json.Marshal(result)

			if string(expectedJSON) != string(resultJSON) {
				t.Errorf("expected %s, got %s", string(expectedJSON), string(resultJSON))
			}
		})
	}
}

// TestDataFetcher_NewFormat tests the new DataFetcher format with path resolution
func TestDataFetcher_NewFormat(t *testing.T) {
	tests := []struct {
		name     string
		fetcher  types.DataFetcher
		sources  map[string]any
		expected any
		wantErr  bool
	}{
		{
			name: "simple path fetcher",
			fetcher: types.DataFetcher{
				Name:   "element",
				Source: "registry",
				Path:   "elements[$.tagName]",
			},
			sources: map[string]any{
				"registry": map[string]any{
					"elements": map[string]any{
						"button-element": map[string]any{
							"tagName": "button-element",
							"name":    "Button",
						},
					},
				},
				"args": map[string]any{
					"tagName": "button-element",
				},
			},
			expected: map[string]any{
				"tagName": "button-element",
				"name":    "Button",
			},
			wantErr: false,
		},
		{
			name: "filtered path fetcher",
			fetcher: types.DataFetcher{
				Name:   "firstAttribute",
				Source: "registry",
				Path:   "elements[$.tagName].attributes",
				Filter: "first",
			},
			sources: map[string]any{
				"registry": map[string]any{
					"elements": map[string]any{
						"button-element": map[string]any{
							"attributes": []any{
								map[string]any{"name": "color", "type": "string"},
								map[string]any{"name": "disabled", "type": "boolean"},
							},
						},
					},
				},
				"args": map[string]any{
					"tagName": "button-element",
				},
			},
			expected: map[string]any{"name": "color", "type": "string"},
			wantErr:  false,
		},
		{
			name: "cross-reference between fetchers",
			fetcher: types.DataFetcher{
				Name:   "attributes",
				Source: "element",
				Path:   "attributes",
			},
			sources: map[string]any{
				"element": map[string]any{
					"tagName": "button-element",
					"attributes": []any{
						map[string]any{"name": "color", "type": "string"},
					},
				},
			},
			expected: []any{
				map[string]any{"name": "color", "type": "string"},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			engine := NewPathTraversalEngine()
			result, err := engine.ExecuteDataFetcher(tt.fetcher, tt.sources)

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			// Convert to JSON for comparison
			expectedJSON, _ := json.Marshal(tt.expected)
			resultJSON, _ := json.Marshal(result)

			if string(expectedJSON) != string(resultJSON) {
				t.Errorf("expected %s, got %s", string(expectedJSON), string(resultJSON))
			}
		})
	}
}

// TestDataTraversal_SubResource tests sub-resource filtering with the new system
func TestDataTraversal_SubResource(t *testing.T) {
	tests := []struct {
		name        string
		fetchers    []types.DataFetcher
		args        map[string]any
		registry    map[string]any
		expected    map[string]any
		wantErr     bool
	}{
		{
			name: "sub-resource attribute filtering",
			fetchers: []types.DataFetcher{
				{
					Name:   "element",
					Source: "registry",
					Path:   "elements[$.tagName]",
				},
				{
					Name:   "attributes",
					Source: "element",
					Path:   "attributes.#(name==$.subResource)",
				},
			},
			args: map[string]any{
				"tagName":     "button-element",
				"subResource": "color",
			},
			registry: map[string]any{
				"elements": map[string]any{
					"button-element": map[string]any{
						"tagName": "button-element",
						"attributes": []any{
							map[string]any{"name": "color", "type": "string", "description": "Button color"},
							map[string]any{"name": "disabled", "type": "boolean", "description": "Disabled state"},
						},
					},
				},
			},
			expected: map[string]any{
				"element": map[string]any{
					"tagName": "button-element",
					"attributes": []any{
						map[string]any{"name": "color", "type": "string", "description": "Button color"},
						map[string]any{"name": "disabled", "type": "boolean", "description": "Disabled state"},
					},
				},
				"attributes": map[string]any{"name": "color", "type": "string", "description": "Button color"},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sources := map[string]any{
				"registry": tt.registry,
				"args":     tt.args,
			}

			engine := NewPathTraversalEngine()
			result, err := engine.ExecuteDataFetchers(tt.fetchers, sources)

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			// Convert to JSON for comparison
			expectedJSON, _ := json.Marshal(tt.expected)
			resultJSON, _ := json.Marshal(result)

			if string(expectedJSON) != string(resultJSON) {
				t.Errorf("expected %s, got %s", string(expectedJSON), string(resultJSON))
			}
		})
	}
}