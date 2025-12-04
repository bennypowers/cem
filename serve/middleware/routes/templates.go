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
	"embed"
	"fmt"
	"html/template"
	"strings"

	M "bennypowers.dev/cem/manifest"
	"bennypowers.dev/cem/serve/middleware"
)

//go:embed templates/workspace-listing.html
var workspaceListingTemplate string

//go:embed templates/navigation.html
var navigationTemplate string

//go:embed templates/404.html
var notFoundTemplate string

//go:embed templates/element-wrapper.html
var elementWrapperTemplate string

//go:embed templates/knobs.html
var knobsTemplate string

//go:embed templates/demo-chrome.html
var demoChromeTemplate string

//go:embed templates/template-error.html
var templateErrorTemplate string

//go:embed templates/**
var TemplatesFS embed.FS

//go:embed templates/js/*.js templates/css/*.css templates/images/* templates/elements/*.css templates/elements/**/*
var InternalModules embed.FS

// TemplateRegistry holds parsed templates and the context for error broadcasting
type TemplateRegistry struct {
	WorkspaceListingTemplate *template.Template
	NavigationTemplate       *template.Template
	NotFoundTemplate         *template.Template
	ElementWrapperTemplate   *template.Template
	KnobsTemplate            *template.Template
	DemoChromeTemplate       *template.Template
	TemplateErrorTemplate    *template.Template
	context                  middleware.DevServerContext
}

// NewTemplateRegistry creates a new template registry with the given context
func NewTemplateRegistry(ctx middleware.DevServerContext) *TemplateRegistry {
	registry := &TemplateRegistry{
		context: ctx,
	}

	// Create template function map with the context
	funcs := registry.getTemplateFuncs()

	// Parse all templates with the function map
	registry.WorkspaceListingTemplate = template.Must(template.New("workspace-listing").Funcs(funcs).Parse(workspaceListingTemplate))
	registry.NavigationTemplate = template.Must(template.New("navigation").Funcs(funcs).Parse(navigationTemplate))
	registry.NotFoundTemplate = template.Must(template.New("404").Funcs(funcs).Parse(notFoundTemplate))
	registry.ElementWrapperTemplate = template.Must(template.New("element-wrapper").Parse(elementWrapperTemplate))
	registry.KnobsTemplate = template.Must(template.New("knobs").Funcs(funcs).Parse(knobsTemplate))
	registry.DemoChromeTemplate = template.Must(template.New("demo-chrome").Funcs(funcs).Parse(demoChromeTemplate))
	registry.TemplateErrorTemplate = template.Must(template.New("template-error").Funcs(funcs).Parse(templateErrorTemplate))

	return registry
}

// getTemplateFuncs returns the template function map with access to the registry's context
// This is a method to provide the context for error broadcasting
func (registry *TemplateRegistry) getTemplateFuncs() template.FuncMap {
	return template.FuncMap{
		"contains": func(haystack, needle interface{}) bool {
			// Handle string contains (substring search)
			if h, ok := haystack.(string); ok {
				if n, ok := needle.(string); ok {
					return strings.Contains(h, n)
				}
			}
			// Handle slice contains (element search)
			if slice, ok := haystack.([]string); ok {
				if str, ok := needle.(string); ok {
					for _, s := range slice {
						if s == str {
							return true
						}
					}
					return false
				}
			}
			return false
		},
		"join": func(elems []string, sep string) string {
			return strings.Join(elems, sep)
		},
		"dict": func(values ...interface{}) (map[string]string, error) {
			if len(values)%2 != 0 {
				return nil, fmt.Errorf("dict requires even number of arguments")
			}
			dict := make(map[string]string, len(values)/2)
			for i := 0; i < len(values); i += 2 {
				key, ok := values[i].(string)
				if !ok {
					return nil, fmt.Errorf("dict keys must be strings")
				}
				dict[key] = fmt.Sprint(values[i+1])
			}
			return dict, nil
		},
		"include": func(path string) template.CSS {
			content, err := TemplatesFS.ReadFile("templates/" + path)
			if err != nil {
				errMsg := "/* Error reading " + path + ": " + err.Error() + " */"

				// Broadcast error to connected clients via WebSocket overlay
				if registry.context != nil {
					_ = registry.context.BroadcastError(
						"Template Include Error",
						"Failed to include template file: "+err.Error(),
						path,
					)
				}

				return template.CSS(errMsg)
			}
			return template.CSS(content)
		},
		"renderElementShadowRoot": func(elementName string, data interface{}) template.HTML {
			html, err := RenderElementShadowRoot(registry, elementName, data)
			if err != nil {
				errMsg := "<!-- Error rendering element " + elementName + ": " + err.Error() + " -->"

				// Broadcast error to connected clients via WebSocket overlay
				if registry.context != nil {
					_ = registry.context.BroadcastError(
						"Element Render Error",
						"Failed to render element: "+err.Error(),
						elementName,
					)
				}

				return template.HTML(errMsg)
			}
			return html
		},
		"markdown": func(text string) template.HTML {
			html, err := markdownToHTML(text)
			if err != nil {
				// If markdown conversion fails, return escaped plain text
				return template.HTML(template.HTMLEscapeString(text))
			}
			return template.HTML(html)
		},
		"prettifyRoute":     prettifyRoute,
		"extractLocalRoute": extractLocalRoute,
		"asCustomElement":   asCustomElement,
		"asClass":           asClass,
		"asFunction":        asFunction,
		"asVariable":        asVariable,
		"asMixin":           asMixin,
		"hasAttr": func(data interface{}, attrName string) bool {
			// Check if an attribute exists in .Attributes map
			if dataMap, ok := data.(map[string]interface{}); ok {
				if attrs, ok := dataMap["Attributes"].(map[string]string); ok {
					_, exists := attrs[attrName]
					return exists
				}
			}
			return false
		},
		"hasMethodMembers": func(members []M.ClassMember) bool {
			// Check if any member is a method
			for _, member := range members {
				if _, ok := member.(*M.ClassMethod); ok {
					return true
				}
			}
			return false
		},
	}
}

// asCustomElement performs type assertion to extract CustomElementDeclaration from Declaration interface.
// Returns nil if the declaration is not a custom element.
// Template usage: {{$ce := asCustomElement .}} {{if $ce}}...{{end}}
func asCustomElement(decl M.Declaration) *M.CustomElementDeclaration {
	if ce, ok := decl.(*M.CustomElementDeclaration); ok {
		return ce
	}
	return nil
}

// asClass performs type assertion to extract ClassDeclaration from Declaration interface.
// Returns nil if the declaration is not a class.
func asClass(decl M.Declaration) *M.ClassDeclaration {
	if c, ok := decl.(*M.ClassDeclaration); ok {
		return c
	}
	return nil
}

// asFunction performs type assertion to extract FunctionDeclaration from Declaration interface.
// Returns nil if the declaration is not a function.
func asFunction(decl M.Declaration) *M.FunctionDeclaration {
	if f, ok := decl.(*M.FunctionDeclaration); ok {
		return f
	}
	return nil
}

// asVariable performs type assertion to extract VariableDeclaration from Declaration interface.
// Returns nil if the declaration is not a variable.
func asVariable(decl M.Declaration) *M.VariableDeclaration {
	if v, ok := decl.(*M.VariableDeclaration); ok {
		return v
	}
	return nil
}

// asMixin performs type assertion to extract MixinDeclaration from Declaration interface.
// Returns nil if the declaration is not a mixin.
func asMixin(decl M.Declaration) *M.MixinDeclaration {
	if m, ok := decl.(*M.MixinDeclaration); ok {
		return m
	}
	return nil
}
