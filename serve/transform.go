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

package serve

import (
	"fmt"
	"strings"

	"github.com/evanw/esbuild/pkg/api"
)

// Loader specifies the file type for transformation
type Loader string

const (
	LoaderTS  Loader = "ts"
	LoaderTSX Loader = "tsx"
	LoaderJS  Loader = "js"
	LoaderJSX Loader = "jsx"
)

// Target specifies the ECMAScript target version
type Target string

const (
	ES2015   Target = "es2015"
	ES2016   Target = "es2016"
	ES2017   Target = "es2017"
	ES2018   Target = "es2018"
	ES2019   Target = "es2019"
	ES2020   Target = "es2020"
	ES2021   Target = "es2021"
	ES2022   Target = "es2022"
	ES2023   Target = "es2023"
	ESNext   Target = "esnext"
)

// SourceMapMode specifies how source maps are generated
type SourceMapMode string

const (
	SourceMapInline   SourceMapMode = "inline"
	SourceMapExternal SourceMapMode = "external"
	SourceMapNone     SourceMapMode = "none"
)

// TransformOptions configures the transformation
type TransformOptions struct {
	Loader    Loader
	Target    Target
	Sourcemap SourceMapMode
}

// TransformResult contains the transformed code and optional source map
type TransformResult struct {
	Code []byte
	Map  []byte
}

// TransformTypeScript transforms TypeScript source code to JavaScript using esbuild
func TransformTypeScript(source []byte, opts TransformOptions) (*TransformResult, error) {
	// Convert our options to esbuild options
	loader := api.LoaderTS
	switch opts.Loader {
	case LoaderTSX:
		loader = api.LoaderTSX
	case LoaderJS:
		loader = api.LoaderJS
	case LoaderJSX:
		loader = api.LoaderJSX
	}

	target := api.ES2020
	switch opts.Target {
	case ES2015:
		target = api.ES2015
	case ES2016:
		target = api.ES2016
	case ES2017:
		target = api.ES2017
	case ES2018:
		target = api.ES2018
	case ES2019:
		target = api.ES2019
	case ES2020:
		target = api.ES2020
	case ES2021:
		target = api.ES2021
	case ES2022:
		target = api.ES2022
	case ES2023:
		target = api.ES2023
	case ESNext:
		target = api.ESNext
	}

	sourcemap := api.SourceMapInline
	switch opts.Sourcemap {
	case SourceMapExternal:
		sourcemap = api.SourceMapExternal
	case SourceMapNone:
		sourcemap = api.SourceMapNone
	}

	// Transform using esbuild
	result := api.Transform(string(source), api.TransformOptions{
		Loader:    loader,
		Target:    target,
		Format:    api.FormatESModule,
		Sourcemap: sourcemap,
	})

	// Check for errors
	if len(result.Errors) > 0 {
		// Format errors for user
		errMsg := "Transform failed:\n"
		for _, err := range result.Errors {
			errMsg += fmt.Sprintf("  %s\n", err.Text)
		}
		return nil, fmt.Errorf("%s", errMsg)
	}

	// Return result
	return &TransformResult{
		Code: result.Code,
		Map:  result.Map,
	}, nil
}

// TransformCSS transforms CSS to a JavaScript module exporting a CSSStyleSheet
func TransformCSS(source []byte) string {
	// Escape backticks and backslashes in CSS
	css := string(source)
	css = strings.ReplaceAll(css, "\\", "\\\\")
	css = strings.ReplaceAll(css, "`", "\\`")
	css = strings.ReplaceAll(css, "${", "\\${")

	// Wrap in constructable stylesheet template
	return fmt.Sprintf(`// CSS transformed to constructable stylesheet
const sheet = new CSSStyleSheet();
sheet.replaceSync(%s);
export default sheet;
`, "`"+css+"`")
}
