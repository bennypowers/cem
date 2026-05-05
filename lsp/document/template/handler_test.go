package template

import (
	"os"
	"path/filepath"
	"testing"

	treesitter "bennypowers.dev/cem/internal/treesitter"
	htmldoc "bennypowers.dev/cem/lsp/document/html"
	"bennypowers.dev/cem/lsp/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTemplateFamily(t *testing.T) {
	tests := []struct {
		name     string
		uri      string
		expected string
	}{
		{
			name:     "nunjucks extension",
			uri:      "file:///project/template.njk",
			expected: "jinja",
		},
		{
			name:     "jinja2 extension",
			uri:      "file:///project/template.j2",
			expected: "jinja",
		},
		{
			name:     "jinja extension",
			uri:      "file:///project/template.jinja",
			expected: "jinja",
		},
		{
			name:     "jinja2 long extension",
			uri:      "file:///project/template.jinja2",
			expected: "jinja",
		},
		{
			name:     "handlebars extension",
			uri:      "file:///project/template.hbs",
			expected: "handlebars",
		},
		{
			name:     "erb extension",
			uri:      "file:///project/template.erb",
			expected: "embedded-template",
		},
		{
			name:     "ejs extension",
			uri:      "file:///project/template.ejs",
			expected: "embedded-template",
		},
		{
			name:     "liquid extension uses jinja",
			uri:      "file:///project/template.liquid",
			expected: "jinja",
		},
		{
			name:     "twig extension uses jinja",
			uri:      "file:///project/template.twig",
			expected: "jinja",
		},
		{
			name:     "unknown extension falls through to jinja",
			uri:      "file:///project/template.html",
			expected: "jinja",
		},
		{
			name:     "uppercase HBS is case insensitive",
			uri:      "file:///project/template.HBS",
			expected: "handlebars",
		},
		{
			name:     "uppercase ERB is case insensitive",
			uri:      "file:///project/template.ERB",
			expected: "embedded-template",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := templateFamily(tt.uri)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func newTemplateHandler(t *testing.T) *Handler {
	t.Helper()
	qm, err := treesitter.NewQueryManager(treesitter.LSPQueries())
	require.NoError(t, err)
	t.Cleanup(func() { qm.Close() })

	htmlHandler, err := htmldoc.NewHandler(qm)
	require.NoError(t, err)
	t.Cleanup(func() { htmlHandler.Close() })

	handler, err := NewHandler(htmlHandler)
	require.NoError(t, err)
	t.Cleanup(func() { handler.Close() })

	return handler
}

func TestFindCustomElements_Nunjucks(t *testing.T) {
	handler := newTemplateHandler(t)

	tests := []struct {
		name     string
		fixture  string
		validate func(t *testing.T, elements []types.CustomElementMatch)
	}{
		{
			name:    "nunjucks with custom elements",
			fixture: filepath.Join("testdata", "with-custom-elements.njk"),
			validate: func(t *testing.T, elements []types.CustomElementMatch) {
				require.Len(t, elements, 2)
				assert.Equal(t, "my-element", elements[0].TagName)
				require.Contains(t, elements[0].Attributes, "disabled")
				require.Contains(t, elements[0].Attributes, "variant")
				assert.Equal(t, "primary", elements[0].Attributes["variant"].Value)
				assert.Equal(t, "other-element", elements[1].TagName)
			},
		},
		{
			name:    "nunjucks no custom elements",
			fixture: filepath.Join("testdata", "no-custom-elements.njk"),
			validate: func(t *testing.T, elements []types.CustomElementMatch) {
				assert.Empty(t, elements)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			content, err := os.ReadFile(tt.fixture)
			require.NoError(t, err)
			doc := handler.CreateDocument("file:///test.njk", string(content), 1)
			defer doc.Close()
			elements, err := handler.FindCustomElements(doc)
			require.NoError(t, err)
			tt.validate(t, elements)
		})
	}
}

func TestFindCustomElements_Handlebars(t *testing.T) {
	handler := newTemplateHandler(t)

	content, err := os.ReadFile(filepath.Join("testdata", "with-custom-elements.hbs"))
	require.NoError(t, err)
	doc := handler.CreateDocument("file:///test.hbs", string(content), 1)
	defer doc.Close()
	elements, err := handler.FindCustomElements(doc)
	require.NoError(t, err)
	require.Len(t, elements, 1)
	assert.Equal(t, "my-card", elements[0].TagName)
}

func TestCreateDocument_SetsLanguage(t *testing.T) {
	handler := newTemplateHandler(t)
	doc := handler.CreateDocument("file:///test.njk", "<div>hello</div>", 1)
	defer doc.Close()
	assert.NotNil(t, doc)
}
