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

// baseFuncMap returns template functions shared across all framework exporters.
func baseFuncMap() template.FuncMap {
	return template.FuncMap{
		"or": func(a, b string) string {
			if a != "" {
				return a
			}
			return b
		},
	}
}

// reactFuncMap returns template functions for the React exporter.
func reactFuncMap() template.FuncMap {
	fm := baseFuncMap()
	fm["jsKey"] = func(s string) string {
		if strings.ContainsAny(s, "-. ") {
			return "'" + s + "'"
		}
		return s
	}
	return fm
}

var (
	reactTmpl   *template.Template
	vueTmpl     *template.Template
	angularTmpl *template.Template

	reactOnce   sync.Once
	vueOnce     sync.Once
	angularOnce sync.Once
)

func getReactTemplate() (*template.Template, error) {
	var err error
	reactOnce.Do(func() {
		reactTmpl, err = template.New("react.ts.tmpl").Funcs(reactFuncMap()).ParseFS(templateFS, "templates/react.ts.tmpl")
	})
	return reactTmpl, err
}

func getVueTemplate() (*template.Template, error) {
	var err error
	vueOnce.Do(func() {
		vueTmpl, err = template.New("vue.vue.tmpl").Funcs(baseFuncMap()).ParseFS(templateFS, "templates/vue.vue.tmpl")
	})
	return vueTmpl, err
}

func getAngularTemplate() (*template.Template, error) {
	var err error
	angularOnce.Do(func() {
		angularTmpl, err = template.New("angular.ts.tmpl").Funcs(baseFuncMap()).ParseFS(templateFS, "templates/angular.ts.tmpl")
	})
	return angularTmpl, err
}
