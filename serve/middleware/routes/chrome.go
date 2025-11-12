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
	"html/template"

	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	goldmarkhtml "github.com/yuin/goldmark/renderer/html"
)

// ChromeData represents template data for demo chrome
type ChromeData struct {
	TagName        string
	DemoTitle      string
	DemoHTML       template.HTML
	EnabledKnobs   string
	KnobsHTML      template.HTML // Rendered knobs controls HTML
	ImportMap      template.HTML // Use HTML instead of JS for importmap script content
	Description    template.HTML
	ShadowMode     bool
	DemoSwitcher   template.HTML
	SourceURL      string        // Source file URL (e.g., GitHub blob URL)
	CanonicalURL   string        // Canonical demo URL
	PackageName    string        // Package name for title (listing pages)
	NavigationHTML template.HTML // Navigation drawer HTML (listing pages)
}

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
			goldmarkhtml.WithHardWraps(),
			goldmarkhtml.WithXHTML(),
		),
	)
)

// renderDemoChrome renders the demo chrome template with given data
func renderDemoChrome(data ChromeData) (string, error) {
	// Render markdown description if present
	if data.Description != "" {
		var buf bytes.Buffer
		err := md.Convert([]byte(data.Description), &buf)
		if err != nil {
			return "", err
		}
		data.Description = template.HTML(buf.String())
	}

	var buf bytes.Buffer
	err := DemoChromeTemplate.Execute(&buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
