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

// renderElementShadowRoot renders a custom element's shadow root with DSD
// Returns just the shadow root (template + styles + shadow DOM)
// Light DOM content should be provided by the caller in the template
func renderElementShadowRoot(elementName string, attrs map[string]string) (template.HTML, error) {
	// Read the element's HTML template (shadow DOM structure)
	templatePath := fmt.Sprintf("templates/elements/%s/%s.html", elementName, elementName)
	templateContent, err := templatesFS.ReadFile(templatePath)
	if err != nil {
		return "", fmt.Errorf("failed to read element template %s: %w", templatePath, err)
	}

	// Read the element's CSS
	cssPath := fmt.Sprintf("templates/elements/%s/%s.css", elementName, elementName)
	cssContent, err := templatesFS.ReadFile(cssPath)
	if err != nil {
		return "", fmt.Errorf("failed to read element CSS %s: %w", cssPath, err)
	}

	// Parse the shadow DOM template
	shadowTmpl, err := template.New(elementName).Parse(string(templateContent))
	if err != nil {
		return "", fmt.Errorf("failed to parse element template %s: %w", elementName, err)
	}

	// Execute shadow DOM template with attrs
	var shadowBuf bytes.Buffer
	err = shadowTmpl.Execute(&shadowBuf, attrs)
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
	err = ElementWrapperTemplate.Execute(&wrapperBuf, wrapperData)
	if err != nil {
		return "", fmt.Errorf("failed to execute wrapper template: %w", err)
	}

	return template.HTML(wrapperBuf.String()), nil
}
