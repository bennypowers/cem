package document

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLanguageFromURI(t *testing.T) {
	tests := []struct {
		name     string
		uri      string
		expected string
	}{
		{
			name:     "HTML file",
			uri:      "file:///project/index.html",
			expected: "html",
		},
		{
			name:     "HTM file",
			uri:      "file:///project/index.htm",
			expected: "html",
		},
		{
			name:     "PHP file",
			uri:      "file:///project/index.php",
			expected: "php",
		},
		{
			name:     "Blade PHP file",
			uri:      "file:///project/view.blade.php",
			expected: "blade",
		},
		{
			name:     "TypeScript file",
			uri:      "file:///project/app.ts",
			expected: "typescript",
		},
		{
			name:     "JavaScript file",
			uri:      "file:///project/app.js",
			expected: "typescript",
		},
		{
			name:     "TSX file",
			uri:      "file:///project/App.tsx",
			expected: "tsx",
		},
		{
			name:     "JSX file",
			uri:      "file:///project/App.jsx",
			expected: "tsx",
		},
		{
			name:     "Nunjucks file",
			uri:      "file:///project/template.njk",
			expected: "template",
		},
		{
			name:     "Jinja2 file",
			uri:      "file:///project/template.j2",
			expected: "template",
		},
		{
			name:     "Jinja file",
			uri:      "file:///project/template.jinja",
			expected: "template",
		},
		{
			name:     "Jinja2 long extension",
			uri:      "file:///project/template.jinja2",
			expected: "template",
		},
		{
			name:     "Liquid file",
			uri:      "file:///project/template.liquid",
			expected: "template",
		},
		{
			name:     "Handlebars file",
			uri:      "file:///project/template.hbs",
			expected: "template",
		},
		{
			name:     "Twig file",
			uri:      "file:///project/template.twig",
			expected: "template",
		},
		{
			name:     "ERB file",
			uri:      "file:///project/template.erb",
			expected: "template",
		},
		{
			name:     "EJS file",
			uri:      "file:///project/template.ejs",
			expected: "template",
		},
		{
			name:     "unknown extension defaults to html",
			uri:      "file:///project/readme.txt",
			expected: "html",
		},
		{
			name:     "no extension defaults to html",
			uri:      "file:///project/Makefile",
			expected: "html",
		},
		{
			name:     "uppercase blade.php is detected",
			uri:      "file:///project/view.BLADE.PHP",
			expected: "blade",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getLanguageFromURI(tt.uri)
			assert.Equal(t, tt.expected, result)
		})
	}
}
