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

//go:embed templates/demo-chromeless.html
var demoChromelessTemplate string

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
	DemoChromelessTemplate   *template.Template
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
	registry.DemoChromelessTemplate = template.Must(template.New("demo-chromeless").Parse(demoChromelessTemplate))
	registry.TemplateErrorTemplate = template.Must(template.New("template-error").Funcs(funcs).Parse(templateErrorTemplate))

	return registry
}

// getTemplateFuncs returns the template function map with access to the registry's context
// This is a method to provide the context for error broadcasting
func (registry *TemplateRegistry) getTemplateFuncs() template.FuncMap {
	return template.FuncMap{
		"asClass":           declarationAsKind[*M.ClassDeclaration],
		"asCustomElement":   declarationAsKind[*M.CustomElementDeclaration],
		"asFunction":        declarationAsKind[*M.FunctionDeclaration],
		"asMixin":           declarationAsKind[*M.MixinDeclaration],
		"asVariable":        declarationAsKind[*M.VariableDeclaration],
		"contains":          strings.Contains,
		"extractLocalRoute": extractLocalRoute,
		"hasMethodMembers":  hasMethodMembers,
		"markdown":          markdown,
		"prettifyRoute":     prettifyRoute,
	}
}

// markdown processes markdown text (e.g. a summary or description) into HTML
func markdown(text string) template.HTML {
	html, err := markdownToHTML(text)
	if err != nil {
		// If markdown conversion fails, return escaped plain text
		return template.HTML(template.HTMLEscapeString(text))
	}
	return template.HTML(html)
}

// hasMethodMembers checks if a class' `members` list contains at least one method
// recall that `members` consists of both fields (properties) and methods
func hasMethodMembers(members []M.ClassMember) bool {
	// Check if any member is a method
	for _, member := range members {
		if _, ok := member.(*M.ClassMethod); ok {
			return true
		}
	}
	return false
}

// declarationAsKind() performs type assertion to extract a subtype from Declaration interface.
// Returns nil if the declaration is not an instance of that subtype.
func declarationAsKind[K M.Declaration](decl M.Declaration) K {
	if m, ok := decl.(K); ok {
		return m
	}
	var zero K // For pointer types (all Declarations), this is nil
	return zero
}
