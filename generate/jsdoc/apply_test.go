package jsdoc

import (
	"testing"

	M "bennypowers.dev/cem/manifest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAppendWithSeparator(t *testing.T) {
	tests := []struct {
		name      string
		existing  string
		new       string
		separator string
		want      string
	}{
		{name: "both empty", existing: "", new: "", separator: "\n", want: ""},
		{name: "new empty", existing: "existing", new: "", separator: "\n", want: "existing"},
		{name: "existing empty", existing: "", new: "new", separator: "\n", want: "new"},
		{name: "both non-empty", existing: "A", new: "B", separator: "\n\n", want: "A\n\nB"},
		{name: "dash separator", existing: "X", new: "Y", separator: " - ", want: "X - Y"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, appendWithSeparator(tt.existing, tt.new, tt.separator))
		})
	}
}

func TestApplyToClassDeclaration(t *testing.T) {
	tests := []struct {
		name  string
		info  classInfo
		check func(t *testing.T, decl *M.ClassDeclaration)
	}{
		{
			name: "sets description and summary",
			info: classInfo{
				Description: "A component",
				Summary:     "Brief",
			},
			check: func(t *testing.T, decl *M.ClassDeclaration) {
				assert.Equal(t, "A component", decl.Description)
				assert.Equal(t, "Brief", decl.Summary)
			},
		},
		{
			name: "sets deprecated flag",
			info: classInfo{
				Deprecated: M.NewDeprecated(true),
			},
			check: func(t *testing.T, decl *M.ClassDeclaration) {
				require.NotNil(t, decl.Deprecated)
				assert.Equal(t, true, decl.Deprecated.Value())
			},
		},
		{
			name: "sets deprecated with reason",
			info: classInfo{
				Deprecated: M.NewDeprecated("Use v2"),
			},
			check: func(t *testing.T, decl *M.ClassDeclaration) {
				require.NotNil(t, decl.Deprecated)
				assert.Equal(t, "Use v2", decl.Deprecated.Value())
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decl := &M.ClassDeclaration{}
			applyToClassDeclaration(&tt.info, decl)
			tt.check(t, decl)
		})
	}
}

func TestApplyToPropertyLike(t *testing.T) {
	tests := []struct {
		name  string
		info  propertyInfo
		init  M.PropertyLike
		check func(t *testing.T, prop *M.PropertyLike)
	}{
		{
			name: "sets description",
			info: propertyInfo{Description: "A property"},
			check: func(t *testing.T, prop *M.PropertyLike) {
				assert.Equal(t, "A property", prop.Description)
			},
		},
		{
			name: "appends description with separator",
			info: propertyInfo{Description: "extra info"},
			init: M.PropertyLike{FullyQualified: M.FullyQualified{Description: "original"}},
			check: func(t *testing.T, prop *M.PropertyLike) {
				assert.Equal(t, "original\n\nextra info", prop.Description)
			},
		},
		{
			name: "sets type",
			info: propertyInfo{Type: "string"},
			check: func(t *testing.T, prop *M.PropertyLike) {
				require.NotNil(t, prop.Type)
				assert.Equal(t, "string", prop.Type.Text)
			},
		},
		{
			name: "sets default",
			info: propertyInfo{Default: "42"},
			check: func(t *testing.T, prop *M.PropertyLike) {
				assert.Equal(t, "42", prop.Default)
			},
		},
		{
			name: "does not overwrite default with empty",
			info: propertyInfo{Default: ""},
			init: M.PropertyLike{Default: "existing"},
			check: func(t *testing.T, prop *M.PropertyLike) {
				assert.Equal(t, "existing", prop.Default)
			},
		},
		{
			name: "sets deprecated",
			info: propertyInfo{Deprecated: M.NewDeprecated(true)},
			check: func(t *testing.T, prop *M.PropertyLike) {
				require.NotNil(t, prop.Deprecated)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prop := tt.init
			applyToPropertyLike(&tt.info, &prop)
			tt.check(t, &prop)
		})
	}
}

func TestApplyToCSSProperty(t *testing.T) {
	tests := []struct {
		name  string
		info  cssPropertyInfo
		init  M.CssCustomProperty
		check func(t *testing.T, prop *M.CssCustomProperty)
	}{
		{
			name: "sets description",
			info: cssPropertyInfo{Description: "Color property"},
			check: func(t *testing.T, prop *M.CssCustomProperty) {
				assert.Equal(t, "Color property", prop.Description)
			},
		},
		{
			name: "appends description",
			info: cssPropertyInfo{Description: "more details"},
			init: M.CssCustomProperty{FullyQualified: M.FullyQualified{Description: "original"}},
			check: func(t *testing.T, prop *M.CssCustomProperty) {
				assert.Equal(t, "original\n\nmore details", prop.Description)
			},
		},
		{
			name: "sets syntax",
			info: cssPropertyInfo{Syntax: "<color>"},
			check: func(t *testing.T, prop *M.CssCustomProperty) {
				assert.Equal(t, "<color>", prop.Syntax)
			},
		},
		{
			name: "does not overwrite syntax with empty",
			info: cssPropertyInfo{Syntax: ""},
			init: M.CssCustomProperty{Syntax: "<length>"},
			check: func(t *testing.T, prop *M.CssCustomProperty) {
				assert.Equal(t, "<length>", prop.Syntax)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prop := tt.init
			applyToCSSProperty(&tt.info, &prop)
			tt.check(t, &prop)
		})
	}
}

func TestApplyToMethod(t *testing.T) {
	tests := []struct {
		name  string
		info  methodInfo
		check func(t *testing.T, m *M.ClassMethod)
	}{
		{
			name: "sets description and summary",
			info: methodInfo{
				Description: "Does something",
				Summary:     "Brief",
			},
			check: func(t *testing.T, m *M.ClassMethod) {
				assert.Equal(t, "Does something", m.Description)
				assert.Equal(t, "Brief", m.Summary)
			},
		},
		{
			name: "sets deprecated",
			info: methodInfo{
				Deprecated: M.NewDeprecated("old"),
			},
			check: func(t *testing.T, m *M.ClassMethod) {
				require.NotNil(t, m.Deprecated)
				assert.Equal(t, "old", m.Deprecated.Value())
			},
		},
		{
			name: "applies return info",
			info: methodInfo{
				Return: &returnInfo{Type: "string", Description: "the result"},
			},
			check: func(t *testing.T, m *M.ClassMethod) {
				require.NotNil(t, m.Return)
				assert.Equal(t, "the result", m.Return.Description)
				require.NotNil(t, m.Return.Type)
				assert.Equal(t, "string", m.Return.Type.Text)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &M.ClassMethod{}
			applyToMethod(&tt.info, m)
			tt.check(t, m)
		})
	}
}

func TestApplyToFunctionDeclaration(t *testing.T) {
	info := &methodInfo{
		Description: "A function",
		Summary:     "Brief",
		Deprecated:  M.NewDeprecated(true),
		Return:      &returnInfo{Type: "void"},
	}
	decl := &M.FunctionDeclaration{}
	applyToFunctionDeclaration(info, decl)
	assert.Equal(t, "A function", decl.Description)
	assert.Equal(t, "Brief", decl.Summary)
	require.NotNil(t, decl.Deprecated)
	require.NotNil(t, decl.Return)
	assert.Equal(t, "void", decl.Return.Type.Text)
}

func TestApplyToFunctionLikeParameters(t *testing.T) {
	info := &methodInfo{
		Parameters: []parameterInfo{
			{Name: "name", Type: "string", Description: "The name", Optional: true, Default: "world"},
		},
	}
	decl := &M.FunctionLike{
		Parameters: []M.Parameter{
			{PropertyLike: M.PropertyLike{FullyQualified: M.FullyQualified{Name: "name"}, Type: &M.Type{Text: "any"}}},
		},
	}
	applyToFunctionLike(info, decl)
	require.Len(t, decl.Parameters, 1)
	assert.Equal(t, "The name", decl.Parameters[0].Description)
	assert.Equal(t, "string", decl.Parameters[0].Type.Text)
	assert.True(t, decl.Parameters[0].Optional)
	assert.Equal(t, "world", decl.Parameters[0].Default)
}

func TestApplyToCustomElementDeclaration(t *testing.T) {
	t.Run("merges existing attributes with jsdoc", func(t *testing.T) {
		info := &classInfo{
			Description: "My element",
			TagName:     "my-el",
			Attrs: []M.Attribute{
				{FullyQualified: M.FullyQualified{Name: "disabled", Description: "jsdoc desc"}, Type: &M.Type{Text: "boolean"}},
				{FullyQualified: M.FullyQualified{Name: "new-attr", Description: "brand new"}},
			},
		}
		decl := &M.CustomElementDeclaration{}
		decl.CustomElement.Attributes = []M.Attribute{
			{FullyQualified: M.FullyQualified{Name: "disabled"}},
		}
		applyToCustomElementDeclaration(info, decl)
		assert.Equal(t, "my-el", decl.TagName)
		require.Len(t, decl.CustomElement.Attributes, 2)
		// Existing attribute gets enriched
		assert.Equal(t, "disabled", decl.CustomElement.Attributes[0].Name)
		assert.Equal(t, "jsdoc desc", decl.CustomElement.Attributes[0].Description)
		require.NotNil(t, decl.CustomElement.Attributes[0].Type)
		assert.Equal(t, "boolean", decl.CustomElement.Attributes[0].Type.Text)
		// New attribute appended
		assert.Equal(t, "new-attr", decl.CustomElement.Attributes[1].Name)
	})

	t.Run("merges existing slots with jsdoc", func(t *testing.T) {
		info := &classInfo{
			Slots: []M.Slot{
				{FullyQualified: M.FullyQualified{Name: "header", Description: "jsdoc header"}},
				{FullyQualified: M.FullyQualified{Name: "footer", Description: "jsdoc footer"}},
			},
		}
		decl := &M.CustomElementDeclaration{}
		decl.CustomElement.Slots = []M.Slot{
			{FullyQualified: M.FullyQualified{Name: "header"}},
		}
		applyToCustomElementDeclaration(info, decl)
		require.Len(t, decl.CustomElement.Slots, 2)
		assert.Equal(t, "header", decl.CustomElement.Slots[0].Name)
		assert.Equal(t, "jsdoc header", decl.CustomElement.Slots[0].Description)
		assert.Equal(t, "footer", decl.CustomElement.Slots[1].Name)
	})

	t.Run("merges existing css parts with jsdoc", func(t *testing.T) {
		info := &classInfo{
			CssParts: []M.CssPart{
				{FullyQualified: M.FullyQualified{Name: "button", Description: "jsdoc button"}},
			},
		}
		decl := &M.CustomElementDeclaration{}
		decl.CustomElement.CssParts = []M.CssPart{
			{FullyQualified: M.FullyQualified{Name: "button"}},
		}
		applyToCustomElementDeclaration(info, decl)
		require.Len(t, decl.CustomElement.CssParts, 1)
		assert.Equal(t, "jsdoc button", decl.CustomElement.CssParts[0].Description)
	})

	t.Run("does not overwrite existing attr description", func(t *testing.T) {
		info := &classInfo{
			Attrs: []M.Attribute{
				{FullyQualified: M.FullyQualified{Name: "name", Description: "jsdoc desc"}},
			},
		}
		decl := &M.CustomElementDeclaration{}
		decl.CustomElement.Attributes = []M.Attribute{
			{FullyQualified: M.FullyQualified{Name: "name", Description: "original desc"}},
		}
		applyToCustomElementDeclaration(info, decl)
		assert.Equal(t, "original desc", decl.CustomElement.Attributes[0].Description)
	})
}
