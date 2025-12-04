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
)

// ElementWrapperData holds data for rendering the element wrapper template
type ElementWrapperData struct {
	CSS       template.CSS
	ShadowDOM template.HTML
	LightDOM  template.HTML
}

// RenderElementShadowRoot renders a custom element's shadow root with DSD
// Returns just the shadow root (template + styles + shadow DOM)
// Light DOM content should be provided by the caller in the template
// This function is recursive - it will render nested custom elements' shadow roots
func RenderElementShadowRoot(templates *TemplateRegistry, elementName string, data interface{}) (template.HTML, error) {
	// Read the element's HTML template (shadow DOM structure)
	templatePath := fmt.Sprintf("templates/elements/%s/%s.html", elementName, elementName)
	templateContent, err := TemplatesFS.ReadFile(templatePath)
	if err != nil {
		return "", fmt.Errorf("failed to read element template %s: %w", templatePath, err)
	}

	// Read the element's CSS
	cssPath := fmt.Sprintf("templates/elements/%s/%s.css", elementName, elementName)
	cssContent, err := TemplatesFS.ReadFile(cssPath)
	if err != nil {
		return "", fmt.Errorf("failed to read element CSS %s: %w", cssPath, err)
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

	// Load and execute the wrapper template
	wrapperData := ElementWrapperData{
		CSS:       template.CSS(cssContent),
		ShadowDOM: template.HTML(shadowBuf.String()),
		LightDOM:  "", // Light DOM is provided by caller
	}

	var wrapperBuf bytes.Buffer
	err = templates.ElementWrapperTemplate.Execute(&wrapperBuf, wrapperData)
	if err != nil {
		return "", fmt.Errorf("failed to execute wrapper template: %w", err)
	}

	return template.HTML(wrapperBuf.String()), nil
}
