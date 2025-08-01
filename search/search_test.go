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
package search_test

import (
	"testing"

	"bennypowers.dev/cem/manifest"
	"bennypowers.dev/cem/search"
	"github.com/pterm/pterm"
	"github.com/stretchr/testify/assert"
)

// mockRenderable implements manifest.Renderable for testing
type mockRenderable struct {
	name        string
	label       string
	summary     string
	description string
}

func (m *mockRenderable) Name() string                     { return m.name }
func (m *mockRenderable) Label() string                    { return m.label }
func (m *mockRenderable) Summary() string                  { return m.summary }
func (m *mockRenderable) Description() string              { return m.description }
func (m *mockRenderable) IsDeprecated() bool               { return false }
func (m *mockRenderable) Deprecation() manifest.Deprecated { return nil }
func (m *mockRenderable) ColumnHeadings() []string         { return []string{"Name"} }
func (m *mockRenderable) ToTableRow() []string             { return []string{m.name} }
func (m *mockRenderable) ToTreeNode(pred manifest.PredicateFunc) pterm.TreeNode {
	return pterm.TreeNode{Text: m.name}
}
func (m *mockRenderable) Children() []manifest.Renderable { return nil }

func TestCreateSearchPredicate(t *testing.T) {
	tests := []struct {
		name        string
		pattern     string
		renderable  *mockRenderable
		shouldMatch bool
	}{
		{
			name:    "Empty pattern matches everything",
			pattern: "",
			renderable: &mockRenderable{
				name: "test-element",
			},
			shouldMatch: true,
		},
		{
			name:    "Literal string in name",
			pattern: "button",
			renderable: &mockRenderable{
				name: "my-button-element",
			},
			shouldMatch: true,
		},
		{
			name:    "Literal string in summary",
			pattern: "click",
			renderable: &mockRenderable{
				name:    "test-element",
				summary: "A clickable button component",
			},
			shouldMatch: true,
		},
		{
			name:    "Literal string in description",
			pattern: "navigation",
			renderable: &mockRenderable{
				name:        "nav-item",
				description: "Used for navigation between pages",
			},
			shouldMatch: true,
		},
		{
			name:    "Literal string in label",
			pattern: "header",
			renderable: &mockRenderable{
				name:  "title-element",
				label: "Page Header Component",
			},
			shouldMatch: true,
		},
		{
			name:    "Case insensitive matching",
			pattern: "BUTTON",
			renderable: &mockRenderable{
				name: "my-button-element",
			},
			shouldMatch: true,
		},
		{
			name:    "No match",
			pattern: "nonexistent",
			renderable: &mockRenderable{
				name:    "test-element",
				summary: "A simple component",
			},
			shouldMatch: false,
		},
		{
			name:    "Regex pattern - starts with",
			pattern: "^my-",
			renderable: &mockRenderable{
				name: "my-button-element",
			},
			shouldMatch: true,
		},
		{
			name:    "Regex pattern - ends with",
			pattern: "-button$",
			renderable: &mockRenderable{
				name: "rh-button",
			},
			shouldMatch: true,
		},
		{
			name:    "Regex pattern - OR",
			pattern: "button|dialog",
			renderable: &mockRenderable{
				name: "my-dialog-element",
			},
			shouldMatch: true,
		},
		{
			name:    "Invalid regex falls back to literal",
			pattern: "[invalid",
			renderable: &mockRenderable{
				name: "[invalid-element",
			},
			shouldMatch: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			predicate := search.CreateSearchPredicate(tt.pattern)
			result := predicate(tt.renderable)
			assert.Equal(t, tt.shouldMatch, result)
		})
	}
}

func TestRenderSearchResults(t *testing.T) {
	// Create a minimal manifest for testing
	manifest := &manifest.Package{
		SchemaVersion: "1.0.0",
		Modules:       []manifest.JavaScriptModule{},
	}

	tests := []struct {
		name     string
		pattern  string
		format   string
		wantErr  bool
		contains string
	}{
		{
			name:     "Tree format with pattern that doesn't exist",
			pattern:  "button",
			format:   "tree",
			wantErr:  false,
			contains: "", // Empty manifest might return empty string
		},
		{
			name:     "Table format with pattern",
			pattern:  "test",
			format:   "table",
			wantErr:  false,
			contains: "<root>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := search.RenderSearchResults(manifest, tt.pattern, tt.format)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			if tt.contains != "" {
				assert.Contains(t, result, tt.contains)
			}
		})
	}
}
