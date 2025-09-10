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
	Attributes    []types.Attribute
	Theme         string
	Context       string
	Options       map[string]string
}

// AttributeTemplateData represents the data passed to attribute-specific templates
type AttributeTemplateData struct {
	types.Attribute
	TagName string
	Context string
}

// ElementWithCapabilities wraps an element with its capability strings
type ElementWithCapabilities struct {
	types.ElementInfo
	Capabilities []string
}

// RegistryQueryData represents the data for registry query templates
type RegistryQueryData struct {
	Elements   map[string]types.ElementInfo
	Categories map[string][]ElementWithCapabilities
	Filter     string
	Search     string
}

// NewTemplateData creates template data from element info and args
func NewTemplateData(element types.ElementInfo, args SuggestCssIntegrationArgs) TemplateData {
	return TemplateData{
		TagName:       element.TagName(),
		Description:   element.Description(),
		CssProperties: element.CssProperties(),
		CssParts:      element.CssParts(),
		CssStates:     element.CssStates(),
		Attributes:    element.Attributes(),
		Theme:         args.Theme,
		Context:       args.Context,
		Options:       args.Options,
	}
}

// NewAttributeTemplateData creates template data for attribute-specific templates
func NewAttributeTemplateData(element types.ElementInfo, context string) TemplateData {
	return TemplateData{
		TagName:       element.TagName(),
		Description:   element.Description(),
		CssProperties: element.CssProperties(),
		CssParts:      element.CssParts(),
		CssStates:     element.CssStates(),
		Attributes:    element.Attributes(),
		Context:       context,
	}
}

// NewSingleAttributeData creates data for rendering a single attribute
func NewSingleAttributeData(attr types.Attribute, tagName, context string) AttributeTemplateData {
	return AttributeTemplateData{
		Attribute: attr,
		TagName:   tagName,
		Context:   context,
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
			case []types.Attribute:
				return len(s)
			case []string:
				return len(s)
			default:
				return 0
			}
		},
		"index": func(slice interface{}, i int) interface{} {
			switch s := slice.(type) {
			case []string:
				if i >= 0 && i < len(s) {
					return s[i]
				}
			}
			return ""
		},
		"gt": func(a, b int) bool {
			return a > b
		},
		"join": func(slice []string, sep string) string {
			return strings.Join(slice, sep)
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

// renderAttributeTemplate loads and executes a template with AttributeTemplateData
func renderAttributeTemplate(templateName string, data AttributeTemplateData) (string, error) {
	// Create template with helper functions
	tmpl := template.New(templateName).Funcs(template.FuncMap{
		"title": strings.Title,
		"len": func(slice interface{}) int {
			switch s := slice.(type) {
			case []string:
				return len(s)
			default:
				return 0
			}
		},
		"gt": func(a, b int) bool {
			return a > b
		},
		"join": func(slice []string, sep string) string {
			return strings.Join(slice, sep)
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

// renderRegistryTemplate loads and executes a template with RegistryQueryData
func renderRegistryTemplate(templateName string, data RegistryQueryData) (string, error) {
	// Create template with helper functions
	tmpl := template.New(templateName).Funcs(template.FuncMap{
		"title": strings.Title,
		"len": func(slice interface{}) int {
			switch s := slice.(type) {
			case map[string]types.ElementInfo:
				return len(s)
			case map[string][]ElementWithCapabilities:
				return len(s)
			case []ElementWithCapabilities:
				return len(s)
			case []string:
				return len(s)
			default:
				return 0
			}
		},
		"gt": func(a, b int) bool {
			return a > b
		},
		"join": func(slice []string, sep string) string {
			return strings.Join(slice, sep)
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
