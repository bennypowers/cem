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
package mcp

import (
	"strings"
	"testing"

	M "bennypowers.dev/cem/manifest"
	"bennypowers.dev/cem/mcp/constants"
	"bennypowers.dev/cem/mcp/relationships"
	MCPTypes "bennypowers.dev/cem/mcp/types"
	"github.com/stretchr/testify/assert"
)

func TestGenerateCacheKey(t *testing.T) {
	assert.Equal(t, "my-element", generateCacheKey("my-element", nil))
	assert.Equal(t, "x-foo", generateCacheKey("x-foo", &M.CustomElement{}))
}

func TestExtractGuidelinesFromElement(t *testing.T) {
	t.Run("no attributes", func(t *testing.T) {
		el := &M.CustomElement{}
		assert.Nil(t, extractGuidelinesFromElement(el))
	})

	t.Run("attributes with descriptions", func(t *testing.T) {
		el := &M.CustomElement{
			Attributes: []M.Attribute{
				{FullyQualified: M.FullyQualified{Name: "variant", Description: "Visual variant"}},
				{FullyQualified: M.FullyQualified{Name: "size"}},
				{FullyQualified: M.FullyQualified{Name: "disabled", Description: "Disables the element"}},
			},
		}
		result := extractGuidelinesFromElement(el)
		assert.Len(t, result, 2)
		joined := strings.Join(result, "\n")
		assert.Contains(t, joined, "variant: Visual variant")
		assert.Contains(t, joined, "disabled: Disables the element")
	})
}

func TestSelectBestSchemaVersion(t *testing.T) {
	ctx := &MCPContext{}

	t.Run("empty returns default", func(t *testing.T) {
		assert.Equal(t, constants.DefaultSchemaVersion, ctx.selectBestSchemaVersion(nil))
	})

	t.Run("single version", func(t *testing.T) {
		assert.Equal(t, "2.0.0", ctx.selectBestSchemaVersion([]string{"2.0.0"}))
	})

	t.Run("prefers speculative", func(t *testing.T) {
		result := ctx.selectBestSchemaVersion([]string{"2.0.0", "2.1.0-speculative"})
		assert.Equal(t, "2.1.0-speculative", result)
	})

	t.Run("highest non-speculative", func(t *testing.T) {
		result := ctx.selectBestSchemaVersion([]string{"1.0.0", "2.0.0", "1.5.0"})
		assert.Equal(t, "2.0.0", result)
	})

	t.Run("skips empty strings", func(t *testing.T) {
		result := ctx.selectBestSchemaVersion([]string{"", "", "1.0.0"})
		assert.Equal(t, "1.0.0", result)
	})

	t.Run("all empty returns default", func(t *testing.T) {
		result := ctx.selectBestSchemaVersion([]string{"", ""})
		assert.Equal(t, constants.DefaultSchemaVersion, result)
	})
}

type mockElementInfo struct {
	tagName       string
	cssProperties []M.CssCustomProperty
}

func (m *mockElementInfo) Declaration() *M.CustomElementDeclaration    { return nil }
func (m *mockElementInfo) TagName() string                             { return m.tagName }
func (m *mockElementInfo) Name() string                                { return m.tagName }
func (m *mockElementInfo) Summary() string                             { return "" }
func (m *mockElementInfo) Description() string                         { return "" }
func (m *mockElementInfo) Module() string                              { return "" }
func (m *mockElementInfo) Package() string                             { return "" }
func (m *mockElementInfo) Attributes() []M.Attribute                   { return nil }
func (m *mockElementInfo) Events() []M.Event                           { return nil }
func (m *mockElementInfo) Slots() []M.Slot                             { return nil }
func (m *mockElementInfo) CssParts() []M.CssPart                      { return nil }
func (m *mockElementInfo) CssProperties() []M.CssCustomProperty        { return m.cssProperties }
func (m *mockElementInfo) CssStates() []M.CssCustomState               { return nil }
func (m *mockElementInfo) Guidelines() []string                        { return nil }
func (m *mockElementInfo) Examples() []MCPTypes.Example                { return nil }
func (m *mockElementInfo) Relationships() []relationships.Relationship { return nil }

func TestComputeCommonPrefixes(t *testing.T) {
	ctx := &MCPContext{}

	t.Run("no elements", func(t *testing.T) {
		result := ctx.computeCommonPrefixes(nil)
		assert.Empty(t, result)
	})

	t.Run("common prefix", func(t *testing.T) {
		elements := []MCPTypes.ElementInfo{
			&mockElementInfo{tagName: "rh-button"},
			&mockElementInfo{tagName: "rh-card"},
			&mockElementInfo{tagName: "pf-icon"},
		}
		result := ctx.computeCommonPrefixes(elements)
		assert.Contains(t, result, "rh")
		assert.NotContains(t, result, "pf")
	})

	t.Run("single element no prefix", func(t *testing.T) {
		elements := []MCPTypes.ElementInfo{
			&mockElementInfo{tagName: "my-element"},
		}
		result := ctx.computeCommonPrefixes(elements)
		assert.Empty(t, result)
	})
}

func TestComputeAllCSSProperties(t *testing.T) {
	ctx := &MCPContext{}

	t.Run("collects unique properties", func(t *testing.T) {
		elements := []MCPTypes.ElementInfo{
			&mockElementInfo{
				tagName: "rh-button",
				cssProperties: []M.CssCustomProperty{
					{FullyQualified: M.FullyQualified{Name: "--color"}},
				},
			},
			&mockElementInfo{
				tagName: "rh-card",
				cssProperties: []M.CssCustomProperty{
					{FullyQualified: M.FullyQualified{Name: "--color"}},
					{FullyQualified: M.FullyQualified{Name: "--spacing"}},
				},
			},
		}
		result := ctx.computeAllCSSProperties(elements)
		assert.Len(t, result, 2)
		assert.Contains(t, result, "--color")
		assert.Contains(t, result, "--spacing")
	})
}
