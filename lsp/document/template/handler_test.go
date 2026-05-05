package template

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
