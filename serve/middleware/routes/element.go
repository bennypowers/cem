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
	"fmt"
	"html/template"
	"slices"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
)

var (
	// minifier is a shared minifier instance for CSS and HTML
	minifier = initMinifier()
)

// initMinifier creates and configures a minifier for CSS and HTML
func initMinifier() *minify.M {
	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	m.AddFunc("text/html", html.Minify)
	return m
}

// InlineStyleElements lists element names that should have inline CSS in their shadow roots.
// With CSS minification, we can inline all styles to get 0ms FOUC while keeping reasonable page size.
// Set to nil to inline all components, or specify individual component names to inline selectively.
var InlineStyleElements []string = nil // nil = inline all (minified)

// ElementWrapperData holds data for rendering the element wrapper template
type ElementWrapperData struct {
	ElementName string
	CSS         template.CSS // Only populated for elements in InlineStyleElements
	ShadowDOM   template.HTML
	LightDOM    template.HTML
}

// shouldInlineCSS checks if an element should have inline CSS
func shouldInlineCSS(elementName string) bool {
	// If InlineStyleElements is nil, inline everything
	if InlineStyleElements == nil {
		return true
	}

	// Otherwise check if element is in the list
	return slices.Contains(InlineStyleElements, elementName)
}

// RenderElementShadowRoot renders a custom element's shadow root with DSD
// Returns just the shadow root (template + styles + shadow DOM)
// Light DOM content should be provided by the caller in the template
// This function is recursive - it will render nested custom elements' shadow roots
func RenderElementShadowRoot(
	templates *TemplateRegistry,
	elementName string,
	data any,
) (template.HTML, error) {
	// Read the element's HTML template (shadow DOM structure)
	templatePath := fmt.Sprintf("templates/elements/%s/%s.html", elementName, elementName)
	templateContent, err := TemplatesFS.ReadFile(templatePath)
	if err != nil {
		return "", fmt.Errorf("failed to read element template %s: %w", templatePath, err)
	}

	// Parse the shadow DOM template with template functions available
	// Note: templates.getTemplateFuncs() includes renderElementShadowRoot, enabling recursion
	shadowTmpl, err := template.New(elementName).Funcs(templates.getTemplateFuncs()).Parse(string(templateContent))
	if err != nil {
		return "", fmt.Errorf("failed to parse element template %s: %w", elementName, err)
	}

	// Execute shadow DOM template with data
	// If the template contains {{renderElementShadowRoot "nested-element" dict}}, it will recurse
	var shadowBuf bytes.Buffer
	err = shadowTmpl.Execute(&shadowBuf, data)
	if err != nil {
		return "", fmt.Errorf("failed to execute element template %s: %w", elementName, err)
	}

	// Minify the shadow DOM HTML to reduce page size
	shadowHTML := shadowBuf.String()
	minifiedShadow, err := minifier.String("text/html", shadowHTML)
	if err != nil {
		// If minification fails, fall back to original HTML
		minifiedShadow = shadowHTML
	}

	// Load and execute the wrapper template
	wrapperData := ElementWrapperData{
		ElementName: elementName,
		ShadowDOM:   template.HTML(minifiedShadow),
		LightDOM:    "", // Light DOM is provided by caller
	}

	// Only inline CSS for critical layout elements to avoid FOUC
	// Other elements use <link> for better caching
	if shouldInlineCSS(elementName) {
		cssPath := fmt.Sprintf("templates/elements/%s/%s.css", elementName, elementName)
		cssContent, err := TemplatesFS.ReadFile(cssPath)
		if err != nil {
			return "", fmt.Errorf("failed to read element CSS %s: %w", cssPath, err)
		}

		// Minify CSS to reduce page size while keeping 0ms FOUC
		minifiedCSS, err := minifier.String("text/css", string(cssContent))
		if err != nil {
			// If minification fails, fall back to original CSS
			minifiedCSS = string(cssContent)
		}

		wrapperData.CSS = template.CSS(minifiedCSS)
	}

	var wrapperBuf bytes.Buffer
	err = templates.ElementWrapperTemplate.Execute(&wrapperBuf, wrapperData)
	if err != nil {
		return "", fmt.Errorf("failed to execute wrapper template: %w", err)
	}

	return template.HTML(wrapperBuf.String()), nil
}
