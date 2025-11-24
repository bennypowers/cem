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

	"bennypowers.dev/cem/serve/middleware"
)

//go:embed templates/default-index.html
var defaultIndexTemplate string

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

//go:embed templates/**
var TemplatesFS embed.FS

//go:embed templates/js/*.js templates/css/*.css templates/images/* templates/elements/*.css templates/elements/**/*
var InternalModules embed.FS

// errorBroadcaster holds the error broadcaster for template errors
// Set by SetErrorBroadcaster when routes middleware is initialized
var errorBroadcaster middleware.DevServerContext

// SetErrorBroadcaster sets the error broadcaster for template error reporting
func SetErrorBroadcaster(ctx middleware.DevServerContext) {
	errorBroadcaster = ctx
}

// Template functions available to templates
// getTemplateFuncs returns the template function map
// This is a function to avoid initialization cycles with renderElementShadowRoot
func getTemplateFuncs() template.FuncMap {
	return template.FuncMap{
		"contains": strings.Contains,
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
				if errorBroadcaster != nil {
					_ = errorBroadcaster.BroadcastError(
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
			html, err := RenderElementShadowRoot(elementName, data)
			if err != nil {
				errMsg := "<!-- Error rendering element " + elementName + ": " + err.Error() + " -->"

				// Broadcast error to connected clients via WebSocket overlay
				if errorBroadcaster != nil {
					_ = errorBroadcaster.BroadcastError(
						"Element Render Error",
						"Failed to render element: "+err.Error(),
						elementName,
					)
				}

				return template.HTML(errMsg)
			}
			return html
		},
	}
}

// DefaultIndexTemplate is the parsed template for the default index page
var DefaultIndexTemplate = template.Must(template.New("default-index").Parse(defaultIndexTemplate))

// WorkspaceListingTemplate is the parsed template for workspace package listing
var WorkspaceListingTemplate = template.Must(template.New("workspace-listing").Funcs(getTemplateFuncs()).Parse(workspaceListingTemplate))

// NavigationTemplate is the parsed template for navigation drawer
var NavigationTemplate = template.Must(template.New("navigation").Funcs(getTemplateFuncs()).Parse(navigationTemplate))

// NotFoundTemplate is the parsed template for 404 page content
var NotFoundTemplate = template.Must(template.New("404").Funcs(getTemplateFuncs()).Parse(notFoundTemplate))

// ElementWrapperTemplate is the parsed template for wrapping custom elements with DSD
var ElementWrapperTemplate = template.Must(template.New("element-wrapper").Parse(elementWrapperTemplate))

// KnobsTemplate is the parsed template for knobs controls
var KnobsTemplate = template.Must(template.New("knobs").Funcs(getTemplateFuncs()).Parse(knobsTemplate))

// DemoChromeTemplate is the parsed template for demo chrome wrapper
var DemoChromeTemplate = template.Must(template.New("demo-chrome").Funcs(getTemplateFuncs()).Parse(demoChromeTemplate))
