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
package export

import (
	"strings"
	"sync"
	"text/template"
)

// reactFuncMap returns template functions for the React exporter.
// Templates also use the built-in `or` function (Go 1.23+).
func reactFuncMap() template.FuncMap {
	return template.FuncMap{
		// jsKey wraps attribute names containing special characters in quotes
		// for use as JS object keys. Assumes keys don't contain single quotes,
		// which is safe for custom-element attribute names.
		"jsKey": func(s string) string {
			if strings.ContainsAny(s, "-. ") {
				return "'" + s + "'"
			}
			return s
		},
	}
}

var (
	getReactTemplate = sync.OnceValues(func() (*template.Template, error) {
		return template.New("react.ts.tmpl").Funcs(reactFuncMap()).ParseFS(templateFS, "templates/react.ts.tmpl")
	})
	getVueTemplate = sync.OnceValues(func() (*template.Template, error) {
		return template.New("vue.vue.tmpl").ParseFS(templateFS, "templates/vue.vue.tmpl")
	})
	getAngularTemplate = sync.OnceValues(func() (*template.Template, error) {
		return template.New("angular.ts.tmpl").ParseFS(templateFS, "templates/angular.ts.tmpl")
	})
	getAngularModuleTemplate = sync.OnceValues(func() (*template.Template, error) {
		return template.New("angular-module.ts.tmpl").ParseFS(templateFS, "templates/angular-module.ts.tmpl")
	})
)
