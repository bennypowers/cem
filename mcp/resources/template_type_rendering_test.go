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
package resources

import (
	"testing"

	"bennypowers.dev/cem/manifest"
	"bennypowers.dev/cem/mcp/relationships"
	"bennypowers.dev/cem/mcp/templates"
	MCPTypes "bennypowers.dev/cem/mcp/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// stubElementInfo adapts manifest data to ElementInfo for template tests.
type stubElementInfo struct {
	tagName    string
	attributes []manifest.Attribute
	events     []manifest.Event
}

func (s *stubElementInfo) Declaration() *manifest.CustomElementDeclaration { return nil }
func (s *stubElementInfo) TagName() string                                 { return s.tagName }
func (s *stubElementInfo) Name() string                                    { return s.tagName }
func (s *stubElementInfo) Summary() string                                 { return "" }
func (s *stubElementInfo) Description() string                             { return "" }
func (s *stubElementInfo) Module() string                                  { return "" }
func (s *stubElementInfo) Package() string                                 { return "" }
func (s *stubElementInfo) Attributes() []manifest.Attribute                { return s.attributes }
func (s *stubElementInfo) Events() []manifest.Event                        { return s.events }
func (s *stubElementInfo) Slots() []manifest.Slot                          { return nil }
func (s *stubElementInfo) CssParts() []manifest.CssPart                    { return nil }
func (s *stubElementInfo) CssProperties() []manifest.CssCustomProperty     { return nil }
func (s *stubElementInfo) CssStates() []manifest.CssCustomState            { return nil }
func (s *stubElementInfo) Methods() []*manifest.ClassMethod                { return nil }
func (s *stubElementInfo) Fields() []*manifest.ClassField                  { return nil }
func (s *stubElementInfo) Guidelines() []string                            { return nil }
func (s *stubElementInfo) Examples() []MCPTypes.Example                    { return nil }
func (s *stubElementInfo) Relationships() []relationships.Relationship     { return nil }

// Inline assertions: template rendering over a hand-built element, where the
// interesting assertion is substring presence at exact markup positions;
// table-driven goldens would obscure which template line regressed.

func TestElementAttributesTemplateRendersTypeSource(t *testing.T) {
	data := struct {
		Element           MCPTypes.ElementInfo
		SchemaDefinitions map[string]any
	}{
		Element: &stubElementInfo{
			tagName: "data-view",
			attributes: []manifest.Attribute{
				{
					FullyQualified: manifest.FullyQualified{Name: "data"},
					Type: &manifest.Type{
						Text:   "Data",
						Source: &manifest.SourceReference{Href: "https://github.com/example/repo/blob/main/src/data-view.ts#L8C23"},
					},
				},
				{
					FullyQualified: manifest.FullyQualified{Name: "plain"},
					Type:           &manifest.Type{Text: "string"},
				},
				{
					FullyQualified: manifest.FullyQualified{Name: "untyped"},
				},
			},
		},
	}

	result, err := templates.RenderTemplate("element-attributes", data)
	require.NoError(t, err, "element-attributes template should render")

	assert.Contains(t, result, "**Type:** `Data` ([declaration](https://github.com/example/repo/blob/main/src/data-view.ts#L8C23))")
	assert.Contains(t, result, "**Type:** `string`")
	assert.Contains(t, result, "**Type:** -")
	assert.NotContains(t, result, "0x", "type rendering must not leak struct pointers")
	assert.NotContains(t, result, "{Data ", "type rendering must not dump the Type struct")
}

func TestElementEventsTemplateRendersTypeSource(t *testing.T) {
	data := struct {
		Element           MCPTypes.ElementInfo
		SchemaDefinitions map[string]any
	}{
		Element: &stubElementInfo{
			tagName: "data-view",
			events: []manifest.Event{
				{
					FullyQualified: manifest.FullyQualified{Name: "data-change"},
					Type: &manifest.Type{
						Text:   "DataChangeEvent<Data>",
						Source: &manifest.SourceReference{Href: "https://github.com/example/repo/blob/main/src/data-view.ts#L12C30"},
					},
				},
			},
		},
	}

	result, err := templates.RenderTemplate("element-events", data)
	require.NoError(t, err, "element-events template should render")

	assert.Contains(t, result, "`DataChangeEvent<Data>` ([declaration](https://github.com/example/repo/blob/main/src/data-view.ts#L12C30))")
	assert.NotContains(t, result, "0x", "type rendering must not leak struct pointers")
}
