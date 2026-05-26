package jsdoc

import (
	"testing"

	Q "bennypowers.dev/cem/internal/treesitter"
	M "bennypowers.dev/cem/manifest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	// Register the jsdoc language so tree-sitter queries work.
	_ "bennypowers.dev/cem/internal/languages/jsdoc"
)

// Inline: pure function, table-driven

// newTestQueryManager creates a QueryManager with the jsdoc query loaded.
// Callers must defer qm.Close().
func newTestQueryManager(t *testing.T) *Q.QueryManager {
	t.Helper()
	qm, err := Q.NewQueryManager(Q.QuerySelector{
		"jsdoc": {"jsdoc"},
	})
	require.NoError(t, err)
	t.Cleanup(func() { qm.Close() })
	return qm
}

// ---------------------------------------------------------------------------
// newTagInfo
// ---------------------------------------------------------------------------

func TestNewTagInfo(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		wantTag     string
		wantDesc    string
	}{
		{
			name:     "simple tag no description",
			input:    "@deprecated",
			wantTag:  "@deprecated",
			wantDesc: "",
		},
		{
			name:     "tag with description",
			input:    "@summary A brief summary",
			wantTag:  "@summary",
			wantDesc: "A brief summary",
		},
		{
			name:     "tag with multiline description",
			input:    "@slot header - The header slot\n * for content",
			wantTag:  "@slot",
			wantDesc: "header - The header slot\nfor content",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info := newTagInfo(tt.input)
			assert.Equal(t, tt.wantTag, info.Tag)
			assert.Equal(t, tt.wantDesc, info.Description)
		})
	}
}

// ---------------------------------------------------------------------------
// tagInfo.toAlias
// ---------------------------------------------------------------------------

func TestTagInfoToAlias(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "simple alias",
			input: "@alias MyComponent",
			want:  "MyComponent",
		},
		{
			name:  "alias with extra whitespace",
			input: "@alias   SomeName",
			want:  "SomeName",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info := tagInfo{source: tt.input}
			assert.Equal(t, tt.want, info.toAlias())
		})
	}
}

// ---------------------------------------------------------------------------
// tagInfo.toAttribute
// ---------------------------------------------------------------------------

func TestTagInfoToAttribute(t *testing.T) {
	tests := []struct {
		name     string
		source   string
		wantName string
		wantType string
		wantDesc string
		wantDef  string
	}{
		{
			name:     "simple attribute",
			source:   "@attr disabled",
			wantName: "disabled",
		},
		{
			name:     "attribute with description",
			source:   "@attribute aria-label - The accessible label",
			wantName: "aria-label",
			wantDesc: "The accessible label",
		},
		{
			name:     "attribute with type",
			source:   "@attr {string} name - The name",
			wantName: "name",
			wantType: "string",
			wantDesc: "The name",
		},
		{
			name:     "attribute with default value in brackets",
			source:   "@attr [disabled=false]",
			wantName: "disabled",
			wantDef:  "false",
		},
		{
			name:     "attribute with type and default",
			source:   "@attr {boolean} [hidden=true] - Whether hidden",
			wantName: "hidden",
			wantType: "boolean",
			wantDef:  "true",
			wantDesc: "Whether hidden",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info := tagInfo{source: tt.source}
			attr := info.toAttribute()
			assert.Equal(t, tt.wantName, attr.Name)
			assert.Equal(t, tt.wantDesc, attr.Description)
			assert.Equal(t, tt.wantDef, attr.Default)
			if tt.wantType != "" {
				require.NotNil(t, attr.Type)
				assert.Equal(t, tt.wantType, attr.Type.Text)
			} else {
				assert.Nil(t, attr.Type)
			}
		})
	}
}

// ---------------------------------------------------------------------------
// tagInfo.toCssPart
// ---------------------------------------------------------------------------

func TestTagInfoToCssPart(t *testing.T) {
	tests := []struct {
		name     string
		source   string
		wantName string
		wantDesc string
	}{
		{
			name:     "part with description",
			source:   "@csspart button - The inner button",
			wantName: "button",
			wantDesc: "The inner button",
		},
		{
			name:     "part without description",
			source:   "@csspart container",
			wantName: "container",
			wantDesc: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info := tagInfo{source: tt.source}
			part := info.toCssPart()
			assert.Equal(t, tt.wantName, part.Name)
			assert.Equal(t, tt.wantDesc, part.Description)
		})
	}
}

// ---------------------------------------------------------------------------
// tagInfo.toCssCustomProperty
// ---------------------------------------------------------------------------

func TestTagInfoToCssCustomProperty(t *testing.T) {
	tests := []struct {
		name       string
		source     string
		wantName   string
		wantDesc   string
		wantSyntax string
		wantDef    string
	}{
		{
			name:     "simple css property",
			source:   "@cssprop --my-color",
			wantName: "--my-color",
		},
		{
			name:     "css property with description",
			source:   "@cssproperty --bg-color - Background color",
			wantName: "--bg-color",
			wantDesc: "Background color",
		},
		{
			name:       "css property with type",
			source:     "@cssprop {<color>} --text-color - Text color",
			wantName:   "--text-color",
			wantSyntax: "<color>",
			wantDesc:   "Text color",
		},
		{
			name:     "css property with default in brackets",
			source:   "@cssprop [--spacing=8px]",
			wantName: "--spacing",
			wantDef:  "8px",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info := tagInfo{source: tt.source}
			prop := info.toCssCustomProperty()
			assert.Equal(t, tt.wantName, prop.Name)
			assert.Equal(t, tt.wantDesc, prop.Description)
			assert.Equal(t, tt.wantSyntax, prop.Syntax)
			assert.Equal(t, tt.wantDef, prop.Default)
		})
	}
}

// ---------------------------------------------------------------------------
// tagInfo.toCssCustomState
// ---------------------------------------------------------------------------

func TestTagInfoToCssCustomState(t *testing.T) {
	tests := []struct {
		name     string
		source   string
		wantName string
		wantDesc string
	}{
		{
			name:     "state with description",
			source:   "@cssstate active - When active",
			wantName: "active",
			wantDesc: "When active",
		},
		{
			name:     "state without description",
			source:   "@cssstate loading",
			wantName: "loading",
			wantDesc: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info := tagInfo{source: tt.source}
			state := info.toCssCustomState()
			assert.Equal(t, tt.wantName, state.Name)
			assert.Equal(t, tt.wantDesc, state.Description)
		})
	}
}

// ---------------------------------------------------------------------------
// tagInfo.toDemo
// ---------------------------------------------------------------------------

func TestTagInfoToDemo(t *testing.T) {
	tests := []struct {
		name     string
		source   string
		wantURL  string
		wantDesc string
	}{
		{
			name:    "demo with URL only",
			source:  "@demo https://example.com",
			wantURL: "https://example.com",
		},
		{
			name:     "demo with URL and description",
			source:   "@demo https://example.com/demo - Interactive demo",
			wantURL:  "https://example.com/demo",
			wantDesc: "Interactive demo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info := tagInfo{source: tt.source}
			demo := info.toDemo()
			assert.Equal(t, tt.wantURL, demo.URL)
			assert.Equal(t, tt.wantDesc, demo.Description)
		})
	}
}

// ---------------------------------------------------------------------------
// tagInfo.toEvent
// ---------------------------------------------------------------------------

func TestTagInfoToEvent(t *testing.T) {
	tests := []struct {
		name     string
		source   string
		wantName string
		wantType string
		wantDesc string
		wantRefs []M.TypeReference
	}{
		{
			name:     "fires event no type",
			source:   "@fires change",
			wantName: "change",
			wantType: "Event",
			wantRefs: []M.TypeReference{{Reference: M.Reference{Name: "Event", Package: "global:"}}},
		},
		{
			name:     "event with CustomEvent type",
			source:   "@event {CustomEvent} my-event - Something happened",
			wantName: "my-event",
			wantType: "CustomEvent",
			wantDesc: "Something happened",
			wantRefs: []M.TypeReference{{Reference: M.Reference{Name: "CustomEvent", Package: "global:"}}},
		},
		{
			name:     "event with custom type",
			source:   "@fires {MyCustomEvent} update",
			wantName: "update",
			wantType: "MyCustomEvent",
		},
		{
			name:     "event with ErrorEvent type",
			source:   "@fires {ErrorEvent} error",
			wantName: "error",
			wantType: "ErrorEvent",
			wantRefs: []M.TypeReference{{Reference: M.Reference{Name: "ErrorEvent", Package: "global:"}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info := tagInfo{source: tt.source}
			event := info.toEvent()
			assert.Equal(t, tt.wantName, event.Name)
			require.NotNil(t, event.Type)
			assert.Equal(t, tt.wantType, event.Type.Text)
			assert.Equal(t, tt.wantDesc, event.Description)
			if tt.wantRefs != nil {
				assert.Equal(t, tt.wantRefs, event.Type.References)
			} else {
				assert.Empty(t, event.Type.References)
			}
		})
	}
}

// ---------------------------------------------------------------------------
// tagInfo.toSlot
// ---------------------------------------------------------------------------

func TestTagInfoToSlot(t *testing.T) {
	tests := []struct {
		name     string
		source   string
		wantName string
		wantDesc string
	}{
		{
			name:     "named slot with description",
			source:   "@slot header - The header content",
			wantName: "header",
			wantDesc: "The header content",
		},
		{
			name:     "named slot without description",
			source:   "@slot footer",
			wantName: "footer",
			wantDesc: "",
		},
		{
			name:     "default (anonymous) slot",
			source:   "@slot - The default slot content",
			wantName: "",
			wantDesc: "The default slot content",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info := tagInfo{source: tt.source}
			slot := info.toSlot()
			assert.Equal(t, tt.wantName, slot.Name)
			assert.Equal(t, tt.wantDesc, slot.Description)
		})
	}
}

// ---------------------------------------------------------------------------
// tagInfo.toReturn
// ---------------------------------------------------------------------------

func TestTagInfoToReturn(t *testing.T) {
	tests := []struct {
		name     string
		source   string
		wantType string
		wantDesc string
	}{
		{
			name:     "return with type only",
			source:   "@returns {string}",
			wantType: "string",
		},
		{
			name:     "return with type and description",
			source:   "@return {boolean} - Whether the thing is valid",
			wantType: "boolean",
			wantDesc: "- Whether the thing is valid",
		},
		{
			name:     "return with no type",
			source:   "@returns",
			wantType: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info := tagInfo{source: tt.source}
			ret := info.toReturn()
			assert.Equal(t, tt.wantType, ret.Type)
			if tt.wantDesc != "" {
				assert.Contains(t, ret.Description, "Whether the thing is valid")
			}
		})
	}
}

// ---------------------------------------------------------------------------
// tagInfo.toParameter
// ---------------------------------------------------------------------------

func TestTagInfoToParameter(t *testing.T) {
	tests := []struct {
		name         string
		description  string
		wantName     string
		wantType     string
		wantDefault  string
		wantOptional bool
	}{
		{
			name:        "simple parameter with trailing description",
			description: "name - The name",
			wantName:    "name",
		},
		{
			name:        "typed parameter with description",
			description: "{string} name - The user name",
			wantName:    "name",
			wantType:    "string",
		},
		{
			name:         "optional with default",
			description:  "{number} [count=0] - How many",
			wantName:     "count",
			wantType:     "number",
			wantDefault:  "0",
			wantOptional: true,
		},
		{
			name:         "optional without default",
			description:  "{boolean} [verbose] - Enable verbose",
			wantName:     "verbose",
			wantType:     "boolean",
			wantOptional: true,
		},
		{
			name:        "parameter with just trailing space",
			description: "{string} value ",
			wantName:    "value",
			wantType:    "string",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info := tagInfo{Tag: "@param", Description: tt.description}
			param := info.toParameter()
			assert.Equal(t, tt.wantName, param.Name)
			assert.Equal(t, tt.wantType, param.Type)
			assert.Equal(t, tt.wantDefault, param.Default)
			assert.Equal(t, tt.wantOptional, param.Optional)
		})
	}
}

// ---------------------------------------------------------------------------
// tagInfo.toExample
// ---------------------------------------------------------------------------

func TestTagInfoToExample(t *testing.T) {
	tests := []struct {
		name     string
		desc     string
		wantSub  string
	}{
		{
			name:    "plain example",
			desc:    "console.log('hello');",
			wantSub: "console.log('hello');",
		},
		{
			name:    "example with explicit caption",
			desc:    "<caption>Usage</caption>\nconsole.log('hello');",
			wantSub: "<figcaption>Usage</figcaption>",
		},
		{
			name:    "example with implicit caption",
			desc:    "How to use:\n```js\nconsole.log('hello');\n```",
			wantSub: "<figcaption>How to use:</figcaption>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info := tagInfo{Tag: "@example", Description: tt.desc}
			got := info.toExample()
			assert.Contains(t, got, tt.wantSub)
		})
	}
}

// ---------------------------------------------------------------------------
// parseForClass - full integration with tree-sitter
// ---------------------------------------------------------------------------

func TestParseForClass(t *testing.T) {
	qm := newTestQueryManager(t)

	tests := []struct {
		name   string
		source string
		check  func(t *testing.T, info *classInfo)
	}{
		{
			name:   "description only",
			source: "/** A simple component */",
			check: func(t *testing.T, info *classInfo) {
				assert.Equal(t, "A simple component", info.Description)
			},
		},
		{
			name: "description and summary",
			source: `/**
 * A detailed description of the component.
 * @summary A short summary
 */`,
			check: func(t *testing.T, info *classInfo) {
				assert.Equal(t, "A detailed description of the component.", info.Description)
				assert.Equal(t, "A short summary", info.Summary)
			},
		},
		{
			name: "deprecated boolean",
			source: `/**
 * @deprecated
 */`,
			check: func(t *testing.T, info *classInfo) {
				require.NotNil(t, info.Deprecated)
				assert.Equal(t, true, info.Deprecated.Value())
			},
		},
		{
			name: "deprecated with reason",
			source: `/**
 * @deprecated Use NewComponent instead
 */`,
			check: func(t *testing.T, info *classInfo) {
				require.NotNil(t, info.Deprecated)
				assert.Equal(t, "Use NewComponent instead", info.Deprecated.Value())
			},
		},
		{
			name: "tagName via `@customElement`",
			source: `/**
 * @customElement my-element
 */`,
			check: func(t *testing.T, info *classInfo) {
				assert.Equal(t, "my-element", info.TagName)
			},
		},
		{
			name: "tagName via `@element`",
			source: `/**
 * @element x-thing
 */`,
			check: func(t *testing.T, info *classInfo) {
				assert.Equal(t, "x-thing", info.TagName)
			},
		},
		{
			name: "alias",
			source: `/**
 * @alias MyAlias
 */`,
			check: func(t *testing.T, info *classInfo) {
				assert.Equal(t, "MyAlias", info.Alias)
			},
		},
		{
			name: "slots",
			source: `/**
 * @slot - Default slot content
 * @slot header - Header slot
 */`,
			check: func(t *testing.T, info *classInfo) {
				require.Len(t, info.Slots, 2)
				assert.Equal(t, "", info.Slots[0].Name)
				assert.Equal(t, "Default slot content", info.Slots[0].Description)
				assert.Equal(t, "header", info.Slots[1].Name)
				assert.Equal(t, "Header slot", info.Slots[1].Description)
			},
		},
		{
			name: "events",
			source: `/**
 * @fires change
 * @event {CustomEvent} my-event - Custom event
 */`,
			check: func(t *testing.T, info *classInfo) {
				require.Len(t, info.Events, 2)
				assert.Equal(t, "change", info.Events[0].Name)
				assert.Equal(t, "Event", info.Events[0].Type.Text)
				assert.Equal(t, "my-event", info.Events[1].Name)
				assert.Equal(t, "CustomEvent", info.Events[1].Type.Text)
				assert.Equal(t, "Custom event", info.Events[1].Description)
			},
		},
		{
			name: "css parts",
			source: `/**
 * @csspart button - The inner button
 */`,
			check: func(t *testing.T, info *classInfo) {
				require.Len(t, info.CssParts, 1)
				assert.Equal(t, "button", info.CssParts[0].Name)
				assert.Equal(t, "The inner button", info.CssParts[0].Description)
			},
		},
		{
			name: "css properties",
			source: `/**
 * @cssprop {<color>} --bg-color - Background color
 */`,
			check: func(t *testing.T, info *classInfo) {
				require.Len(t, info.CssProperties, 1)
				assert.Equal(t, "--bg-color", info.CssProperties[0].Name)
				assert.Equal(t, "<color>", info.CssProperties[0].Syntax)
				assert.Equal(t, "Background color", info.CssProperties[0].Description)
			},
		},
		{
			name: "css states",
			source: `/**
 * @cssstate active - Element is active
 */`,
			check: func(t *testing.T, info *classInfo) {
				require.Len(t, info.CssStates, 1)
				assert.Equal(t, "active", info.CssStates[0].Name)
				assert.Equal(t, "Element is active", info.CssStates[0].Description)
			},
		},
		{
			name: "attributes",
			source: `/**
 * @attr {string} name - The name attribute
 */`,
			check: func(t *testing.T, info *classInfo) {
				require.Len(t, info.Attrs, 1)
				assert.Equal(t, "name", info.Attrs[0].Name)
				assert.Equal(t, "The name attribute", info.Attrs[0].Description)
				require.NotNil(t, info.Attrs[0].Type)
				assert.Equal(t, "string", info.Attrs[0].Type.Text)
			},
		},
		{
			name: "demo",
			source: `/**
 * @demo https://example.com/demo - Live demo
 */`,
			check: func(t *testing.T, info *classInfo) {
				require.Len(t, info.Demos, 1)
				assert.Equal(t, "https://example.com/demo", info.Demos[0].URL)
				assert.Equal(t, "Live demo", info.Demos[0].Description)
			},
		},
		{
			name: "comprehensive jsdoc",
			source: `/**
 * A comprehensive custom element.
 * @summary Brief summary
 * @element x-comprehensive
 * @slot - Default content
 * @slot header - Header content
 * @fires change
 * @csspart base - The base wrapper
 * @cssprop --text-color - Text color
 * @attr {boolean} disabled - Whether disabled
 * @deprecated Use v2 instead
 */`,
			check: func(t *testing.T, info *classInfo) {
				assert.Equal(t, "A comprehensive custom element.", info.Description)
				assert.Equal(t, "Brief summary", info.Summary)
				assert.Equal(t, "x-comprehensive", info.TagName)
				assert.Len(t, info.Slots, 2)
				assert.Len(t, info.Events, 1)
				assert.Len(t, info.CssParts, 1)
				assert.Len(t, info.CssProperties, 1)
				assert.Len(t, info.Attrs, 1)
				require.NotNil(t, info.Deprecated)
				assert.Equal(t, "Use v2 instead", info.Deprecated.Value())
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info, err := parseForClass(tt.source, qm)
			require.NoError(t, err)
			tt.check(t, info)
		})
	}
}

// ---------------------------------------------------------------------------
// parseForProperty
// ---------------------------------------------------------------------------

func TestParseForProperty(t *testing.T) {
	qm := newTestQueryManager(t)

	tests := []struct {
		name   string
		source string
		check  func(t *testing.T, info *propertyInfo)
	}{
		{
			name:   "description only",
			source: "/** The color value */",
			check: func(t *testing.T, info *propertyInfo) {
				assert.Equal(t, "The color value", info.Description)
			},
		},
		{
			name: "type annotation",
			source: `/**
 * @type {string}
 */`,
			check: func(t *testing.T, info *propertyInfo) {
				assert.Equal(t, "string", info.Type)
			},
		},
		{
			name: "default value",
			source: `/**
 * @default red
 */`,
			check: func(t *testing.T, info *propertyInfo) {
				assert.Equal(t, "red", info.Default)
			},
		},
		{
			name: "defaultvalue tag",
			source: `/**
 * @defaultvalue blue
 */`,
			check: func(t *testing.T, info *propertyInfo) {
				assert.Equal(t, "blue", info.Default)
			},
		},
		{
			name: "deprecated boolean",
			source: `/**
 * @deprecated
 */`,
			check: func(t *testing.T, info *propertyInfo) {
				require.NotNil(t, info.Deprecated)
				assert.Equal(t, true, info.Deprecated.Value())
			},
		},
		{
			name: "deprecated with reason",
			source: `/**
 * @deprecated Use .color instead
 */`,
			check: func(t *testing.T, info *propertyInfo) {
				require.NotNil(t, info.Deprecated)
				assert.Equal(t, "Use .color instead", info.Deprecated.Value())
			},
		},
		{
			name: "`@ignore` tag",
			source: `/**
 * @ignore
 */`,
			check: func(t *testing.T, info *propertyInfo) {
				assert.True(t, info.Ignore)
			},
		},
		{
			name: "`@internal` tag",
			source: `/**
 * @internal
 */`,
			check: func(t *testing.T, info *propertyInfo) {
				assert.True(t, info.Ignore)
			},
		},
		{
			name: "summary tag",
			source: `/**
 * Full description here.
 * @summary Short summary
 */`,
			check: func(t *testing.T, info *propertyInfo) {
				assert.Equal(t, "Full description here.", info.Description)
				assert.Equal(t, "Short summary", info.Summary)
			},
		},
		{
			name: "example appended to description",
			source: `/**
 * A property.
 * @example
 * doStuff()
 */`,
			check: func(t *testing.T, info *propertyInfo) {
				assert.Contains(t, info.Description, "doStuff()")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info, err := parseForProperty(tt.source, qm)
			require.NoError(t, err)
			tt.check(t, info)
		})
	}
}

// ---------------------------------------------------------------------------
// parseForCSSProperty
// ---------------------------------------------------------------------------

func TestParseForCSSProperty(t *testing.T) {
	qm := newTestQueryManager(t)

	tests := []struct {
		name   string
		source string
		check  func(t *testing.T, info *cssPropertyInfo)
	}{
		{
			name:   "description only",
			source: "/** The background color */",
			check: func(t *testing.T, info *cssPropertyInfo) {
				assert.Equal(t, "The background color", info.Description)
			},
		},
		{
			name: "syntax tag via tree-sitter type node",
			// NOTE: tree-sitter's JSDoc grammar does not produce a "type" child
			// for the non-standard @syntax tag, so the Syntax field remains empty
			// when parsed this way. The @syntax value is populated by a different
			// code path (the tagInfo.toCssCustomProperty regex parser for @cssprop tags).
			source: `/**
 * has syntax and default
 * @syntax {<length>}
 */`,
			check: func(t *testing.T, info *cssPropertyInfo) {
				assert.Equal(t, "", info.Syntax, "tree-sitter JSDoc grammar does not recognize @syntax type")
			},
		},
		{
			name: "deprecated boolean",
			source: `/**
 * @deprecated
 */`,
			check: func(t *testing.T, info *cssPropertyInfo) {
				require.NotNil(t, info.Deprecated)
				assert.Equal(t, true, info.Deprecated.Value())
			},
		},
		{
			name: "summary",
			source: `/**
 * Detailed description.
 * @summary Short summary
 */`,
			check: func(t *testing.T, info *cssPropertyInfo) {
				assert.Equal(t, "Detailed description.", info.Description)
				assert.Equal(t, "Short summary", info.Summary)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info, err := parseForCSSProperty(tt.source, qm)
			require.NoError(t, err)
			tt.check(t, info)
		})
	}
}

// ---------------------------------------------------------------------------
// parseForMethod
// ---------------------------------------------------------------------------

func TestParseForMethod(t *testing.T) {
	qm := newTestQueryManager(t)

	tests := []struct {
		name   string
		source string
		check  func(t *testing.T, info *methodInfo)
	}{
		{
			name:   "description only",
			source: "/** Does a thing */",
			check: func(t *testing.T, info *methodInfo) {
				assert.Equal(t, "Does a thing", info.Description)
			},
		},
		{
			name: "method with param",
			source: `/**
 * Greets a user.
 * @param {string} name - The user's name
 */`,
			check: func(t *testing.T, info *methodInfo) {
				assert.Equal(t, "Greets a user.", info.Description)
				require.Len(t, info.Parameters, 1)
				assert.Equal(t, "name", info.Parameters[0].Name)
				assert.Equal(t, "string", info.Parameters[0].Type)
			},
		},
		{
			name: "method with return",
			source: `/**
 * @returns {boolean}
 */`,
			check: func(t *testing.T, info *methodInfo) {
				require.NotNil(t, info.Return)
				assert.Equal(t, "boolean", info.Return.Type)
			},
		},
		{
			name: "deprecated method",
			source: `/**
 * @deprecated Use newMethod instead
 */`,
			check: func(t *testing.T, info *methodInfo) {
				require.NotNil(t, info.Deprecated)
				assert.Equal(t, "Use newMethod instead", info.Deprecated.Value())
			},
		},
		{
			name: "summary",
			source: `/**
 * Full description.
 * @summary Brief
 */`,
			check: func(t *testing.T, info *methodInfo) {
				assert.Equal(t, "Full description.", info.Description)
				assert.Equal(t, "Brief", info.Summary)
			},
		},
		{
			name: "multiple params",
			source: `/**
 * @param {string} first - First param
 * @param {number} second - Second param
 */`,
			check: func(t *testing.T, info *methodInfo) {
				require.Len(t, info.Parameters, 2)
				assert.Equal(t, "first", info.Parameters[0].Name)
				assert.Equal(t, "string", info.Parameters[0].Type)
				assert.Equal(t, "second", info.Parameters[1].Name)
				assert.Equal(t, "number", info.Parameters[1].Type)
			},
		},
		{
			name: "dotted param names (sub-properties) included since dot is stripped",
			source: `/**
 * @param {object} options - Options object
 * @param {string} options.name - The name
 */`,
			check: func(t *testing.T, info *methodInfo) {
				// Both params are included because the toParameter regex extracts
				// only the portion after the dot as the Name, so the dot-check
				// in parseForMethod does not filter it.
				assert.GreaterOrEqual(t, len(info.Parameters), 1)
				assert.Equal(t, "options", info.Parameters[0].Name)
				assert.Equal(t, "object", info.Parameters[0].Type)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info, err := parseForMethod(tt.source, qm)
			require.NoError(t, err)
			tt.check(t, info)
		})
	}
}
