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

package routes

import (
	"bytes"

	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/microcosm-cc/bluemonday"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/extension"
	goldmarkhtml "github.com/yuin/goldmark/renderer/html"
)

var (
	// Markdown renderer with GFM, syntax highlighting, and HTML escaping
	md = goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			highlighting.NewHighlighting(
				highlighting.WithStyle("github"),
				highlighting.WithFormatOptions(
					chromahtml.WithClasses(true),
				),
			),
		),
		goldmark.WithRendererOptions(
			goldmarkhtml.WithXHTML(),
		),
	)

	// HTML sanitization policy for user-generated content
	// Allows safe HTML elements while stripping dangerous tags and attributes
	htmlSanitizer = bluemonday.UGCPolicy()
)

// markdownToHTML converts markdown text to sanitized HTML
// The output is sanitized using bluemonday to prevent XSS attacks
// Returns an error if markdown conversion fails, allowing callers to handle it appropriately
func markdownToHTML(text string) (string, error) {
	var buf bytes.Buffer
	err := md.Convert([]byte(text), &buf)
	if err != nil {
		return "", err
	}
	// Sanitize the HTML output to prevent XSS attacks
	sanitized := htmlSanitizer.Sanitize(buf.String())
	return sanitized, nil
}
