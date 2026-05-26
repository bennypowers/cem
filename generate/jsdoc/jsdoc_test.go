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

// Inline: pure function, scalar assertions

func TestEnrichClassWithJSDoc(t *testing.T) {
	qm := newTestQueryManager(t)

	jsdoc := `/**
 * A base class.
 * @summary Brief class summary
 * @deprecated Use NewClass
 */`
	decl := &M.ClassDeclaration{}
	err := EnrichClassWithJSDoc(jsdoc, decl, qm)
	require.NoError(t, err)
	assert.Equal(t, "A base class.", decl.Description)
	assert.Equal(t, "Brief class summary", decl.Summary)
	require.NotNil(t, decl.Deprecated)
	assert.Equal(t, "Use NewClass", decl.Deprecated.Value())
}

func TestEnrichCustomElementWithJSDoc(t *testing.T) {
	qm := newTestQueryManager(t)

	jsdoc := `/**
 * My element description.
 * @element my-element
 * @slot - Default slot
 * @slot header - Header slot
 * @fires {CustomEvent} change - Emitted on change
 * @csspart button - The button part
 * @cssprop --text-color - Text color
 * @attr {string} name - The name attribute
 */`
	decl := &M.CustomElementDeclaration{}
	err := EnrichCustomElementWithJSDoc(jsdoc, decl, qm)
	require.NoError(t, err)

	assert.Equal(t, "My element description.", decl.Description)
	assert.Equal(t, "my-element", decl.TagName)
	require.Len(t, decl.CustomElement.Slots, 2)
	assert.Equal(t, "", decl.CustomElement.Slots[0].Name)
	assert.Equal(t, "header", decl.CustomElement.Slots[1].Name)
	require.Len(t, decl.CustomElement.Events, 1)
	assert.Equal(t, "change", decl.CustomElement.Events[0].Name)
	require.Len(t, decl.CustomElement.CssParts, 1)
	assert.Equal(t, "button", decl.CustomElement.CssParts[0].Name)
	require.Len(t, decl.CustomElement.CssProperties, 1)
	assert.Equal(t, "--text-color", decl.CustomElement.CssProperties[0].Name)
	require.Len(t, decl.CustomElement.Attributes, 1)
	assert.Equal(t, "name", decl.CustomElement.Attributes[0].Name)
}

func TestEnrichMethodWithJSDoc(t *testing.T) {
	qm := newTestQueryManager(t)

	jsdoc := `/**
 * Greets someone.
 * @summary Greeting method
 * @param {string} name - The person's name
 * @returns {string}
 */`
	method := &M.ClassMethod{}
	method.Parameters = []M.Parameter{
		{PropertyLike: M.PropertyLike{FullyQualified: M.FullyQualified{Name: "name"}, Type: &M.Type{Text: "any"}}},
	}
	err := EnrichMethodWithJSDoc(jsdoc, method, qm)
	require.NoError(t, err)
	assert.Equal(t, "Greets someone.", method.Description)
	assert.Equal(t, "Greeting method", method.Summary)
	require.NotNil(t, method.Return)
	assert.Equal(t, "string", method.Return.Type.Text)
	assert.Equal(t, "string", method.Parameters[0].Type.Text)
}

func TestEnrichFunctionWithJSDoc(t *testing.T) {
	qm := newTestQueryManager(t)

	jsdoc := `/**
 * Adds two numbers.
 * @param {number} a - First number
 * @param {number} b - Second number
 * @returns {number}
 * @deprecated Use sum() instead
 */`
	fn := &M.FunctionDeclaration{}
	fn.Parameters = []M.Parameter{
		{PropertyLike: M.PropertyLike{FullyQualified: M.FullyQualified{Name: "a"}, Type: &M.Type{Text: "any"}}},
		{PropertyLike: M.PropertyLike{FullyQualified: M.FullyQualified{Name: "b"}, Type: &M.Type{Text: "any"}}},
	}
	err := EnrichFunctionWithJSDoc(jsdoc, fn, qm)
	require.NoError(t, err)
	assert.Equal(t, "Adds two numbers.", fn.Description)
	require.NotNil(t, fn.Deprecated)
	require.NotNil(t, fn.Return)
	assert.Equal(t, "number", fn.Return.Type.Text)
	assert.Equal(t, "number", fn.Parameters[0].Type.Text)
	assert.Equal(t, "number", fn.Parameters[1].Type.Text)
}

func TestEnrichPropertyWithJSDoc(t *testing.T) {
	qm := newTestQueryManager(t)

	jsdoc := `/**
 * The color property.
 * @type {string}
 * @default red
 */`
	prop := &M.PropertyLike{}
	err := EnrichPropertyWithJSDoc(jsdoc, prop, qm)
	require.NoError(t, err)
	assert.Equal(t, "The color property.", prop.Description)
	require.NotNil(t, prop.Type)
	assert.Equal(t, "string", prop.Type.Text)
	assert.Equal(t, "red", prop.Default)
}

func TestEnrichCSSPropertyWithJSDoc(t *testing.T) {
	qm := newTestQueryManager(t)

	jsdoc := `/**
 * Background color of the component.
 * @summary Background color
 */`
	prop := &M.CssCustomProperty{}
	err := EnrichCSSPropertyWithJSDoc(jsdoc, prop, qm)
	require.NoError(t, err)
	assert.Equal(t, "Background color of the component.", prop.Description)
	assert.Equal(t, "Background color", prop.Summary)
}

func TestExtractAliasFromJSDoc(t *testing.T) {
	qm := newTestQueryManager(t)

	t.Run("with alias", func(t *testing.T) {
		jsdoc := `/**
 * @alias MyAlias
 */`
		alias, err := ExtractAliasFromJSDoc(jsdoc, qm)
		require.NoError(t, err)
		assert.Equal(t, "MyAlias", alias)
	})

	t.Run("without alias", func(t *testing.T) {
		jsdoc := `/** Just a comment */`
		alias, err := ExtractAliasFromJSDoc(jsdoc, qm)
		require.NoError(t, err)
		assert.Equal(t, "", alias)
	})
}

func TestHasIgnoreTag(t *testing.T) {
	qm := newTestQueryManager(t)

	tests := []struct {
		name   string
		input  string
		want   bool
	}{
		{name: "empty string", input: "", want: false},
		{name: "whitespace only", input: "   ", want: false},
		{name: "no ignore tag", input: "/**\n * A thing\n */", want: false},
		{name: "`@ignore` tag", input: "/**\n * @ignore\n */", want: true},
		{name: "`@internal` tag", input: "/**\n * @internal\n */", want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HasIgnoreTag(tt.input, qm)
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestHasElementTag(t *testing.T) {
	qm := newTestQueryManager(t)

	tests := []struct {
		name   string
		input  string
		want   bool
	}{
		{name: "empty string", input: "", want: false},
		{name: "whitespace only", input: "   ", want: false},
		{name: "no element tag", input: "/**\n * A class\n */", want: false},
		{name: "`@element` tag", input: "/**\n * @element my-el\n */", want: true},
		{name: "`@customElement` tag", input: "/**\n * @customElement my-el\n */", want: true},
		{name: "`@tagName` tag", input: "/**\n * @tagName my-el\n */", want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HasElementTag(tt.input, qm)
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

// TestQueryManagerCreation verifies we can create QueryManagers with the right queries.
func TestQueryManagerCreation(t *testing.T) {
	t.Run("valid jsdoc query", func(t *testing.T) {
		qm, err := Q.NewQueryManager(Q.QuerySelector{"jsdoc": {"jsdoc"}})
		require.NoError(t, err)
		defer qm.Close()
	})

	t.Run("invalid query name", func(t *testing.T) {
		_, err := Q.NewQueryManager(Q.QuerySelector{"jsdoc": {"nonexistent"}})
		assert.Error(t, err)
	})
}
