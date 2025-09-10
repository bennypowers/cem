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
package tools

import (
	"bytes"
	"embed"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"

	"bennypowers.dev/cem/mcp/types"
)

//go:embed templates/*.md
var templateFiles embed.FS

// TemplateData represents the data passed to templates
type TemplateData struct {
	TagName       string
	Description   string
	CssProperties []types.CssProperty
	CssParts      []types.CssPart
	CssStates     []types.CssState
	Theme         string
	Context       string
	Options       map[string]string
}

// NewTemplateData creates template data from element info and args
func NewTemplateData(element types.ElementInfo, args SuggestCssIntegrationArgs) TemplateData {
	return TemplateData{
		TagName:       element.TagName(),
		Description:   element.Description(),
		CssProperties: element.CssProperties(),
		CssParts:      element.CssParts(),
		CssStates:     element.CssStates(),
		Theme:         args.Theme,
		Context:       args.Context,
		Options:       args.Options,
	}
}

// renderTemplate loads and executes a template with the given data
func renderTemplate(templateName string, data TemplateData) (string, error) {
	// Create template with helper functions
	tmpl := template.New(templateName).Funcs(template.FuncMap{
		"title": strings.Title,
		"len": func(slice interface{}) int {
			switch s := slice.(type) {
			case []types.CssProperty:
				return len(s)
			case []types.CssPart:
				return len(s)
			case []types.CssState:
				return len(s)
			default:
				return 0
			}
		},
		"gt": func(a, b int) bool {
			return a > b
		},
	})

	// Load template content
	templatePath := filepath.Join("templates", templateName+".md")
	content, err := templateFiles.ReadFile(templatePath)
	if err != nil {
		return "", fmt.Errorf("failed to read template %s: %w", templateName, err)
	}

	// Parse template
	tmpl, err = tmpl.Parse(string(content))
	if err != nil {
		return "", fmt.Errorf("failed to parse template %s: %w", templateName, err)
	}

	// Execute template
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to execute template %s: %w", templateName, err)
	}

	return buf.String(), nil
}

// getTemplatesDir returns the directory containing template files
func getTemplatesDir() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return ""
	}
	return filepath.Join(filepath.Dir(filename), "templates")
}
