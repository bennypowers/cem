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
package queries

import (
	"fmt"
	"sync"

	ts "github.com/tree-sitter/go-tree-sitter"

	tsHandlebars "bennypowers.dev/tree-sitter-handlebars/bindings/go"
	tsJinja "bennypowers.dev/tree-sitter-jinja-dialects/bindings/go"
	tsBlade "github.com/EmranMR/tree-sitter-blade/bindings/go"
	tsCss "github.com/tree-sitter/tree-sitter-css/bindings/go"
	tsET "github.com/tree-sitter/tree-sitter-embedded-template/bindings/go"
	tsHtml "github.com/tree-sitter/tree-sitter-html/bindings/go"
	tsJsdoc "github.com/tree-sitter/tree-sitter-jsdoc/bindings/go"
	tsPHP "github.com/tree-sitter/tree-sitter-php/bindings/go"
	tsTypescript "github.com/tree-sitter/tree-sitter-typescript/bindings/go"
)

// ---- Languages struct holds pre-initialized language grammars ----
var languages = struct {
	typescript       *ts.Language
	jsdoc            *ts.Language
	css              *ts.Language
	html             *ts.Language
	tsx              *ts.Language
	blade            *ts.Language
	jinja            *ts.Language
	handlebars       *ts.Language
	embeddedTemplate *ts.Language
	php              *ts.Language
}{
	ts.NewLanguage(tsTypescript.LanguageTypescript()),
	ts.NewLanguage(tsJsdoc.Language()),
	ts.NewLanguage(tsCss.Language()),
	ts.NewLanguage(tsHtml.Language()),
	ts.NewLanguage(tsTypescript.LanguageTSX()),
	ts.NewLanguage(tsBlade.Language()),
	ts.NewLanguage(tsJinja.Language()),
	ts.NewLanguage(tsHandlebars.Language()),
	ts.NewLanguage(tsET.Language()),
	ts.NewLanguage(tsPHP.LanguagePHP()),
}

// ---- Parser Pooling Section ----

func newParserPool(lang *ts.Language) *sync.Pool {
	return &sync.Pool{
		New: func() any {
			parser := ts.NewParser()
			if err := parser.SetLanguage(lang); err != nil {
				panic(fmt.Sprintf("failed to set language: %v", err))
			}
			return parser
		},
	}
}

func getParser(pool *sync.Pool) *ts.Parser {
	return pool.Get().(*ts.Parser)
}

func putParser(pool *sync.Pool, parser *ts.Parser) {
	parser.Reset()
	pool.Put(parser)
}

var (
	htmlParserPool       = newParserPool(languages.html)
	cssParserPool        = newParserPool(languages.css)
	jsdocParserPool      = newParserPool(languages.jsdoc)
	typescriptParserPool = newParserPool(languages.typescript)
	tsxParserPool        = newParserPool(languages.tsx)
	bladeParserPool      = newParserPool(languages.blade)
	jinjaParserPool      = newParserPool(languages.jinja)
	hbsParserPool        = newParserPool(languages.handlebars)
	etParserPool         = newParserPool(languages.embeddedTemplate)
	phpParserPool        = newParserPool(languages.php)
)

// Note: Removed cursor pooling due to stateful nature of QueryCursor objects.
// Cursors maintain internal state that can affect subsequent queries,
// leading to unpredictable behavior. Always create fresh cursors.

// Get*Parser returns a pooled parser for the given language.
// Always defer the matching Put*Parser call to return it to the pool.
func GetHTMLParser() *ts.Parser        { return getParser(htmlParserPool) }
func PutHTMLParser(parser *ts.Parser)  { putParser(htmlParserPool, parser) }
func GetCSSParser() *ts.Parser         { return getParser(cssParserPool) }
func PutCSSParser(parser *ts.Parser)   { putParser(cssParserPool, parser) }
func GetJSDocParser() *ts.Parser       { return getParser(jsdocParserPool) }
func PutJSDocParser(parser *ts.Parser) { putParser(jsdocParserPool, parser) }

func RetrieveTypeScriptParser() *ts.Parser        { return getParser(typescriptParserPool) }
func PutTypeScriptParser(parser *ts.Parser)       { putParser(typescriptParserPool, parser) }
func GetTSXParser() *ts.Parser                    { return getParser(tsxParserPool) }
func PutTSXParser(parser *ts.Parser)              { putParser(tsxParserPool, parser) }
func GetBladeParser() *ts.Parser                  { return getParser(bladeParserPool) }
func PutBladeParser(parser *ts.Parser)            { putParser(bladeParserPool, parser) }
func GetJinjaParser() *ts.Parser                  { return getParser(jinjaParserPool) }
func PutJinjaParser(parser *ts.Parser)            { putParser(jinjaParserPool, parser) }
func GetHandlebarsParser() *ts.Parser             { return getParser(hbsParserPool) }
func PutHandlebarsParser(parser *ts.Parser)       { putParser(hbsParserPool, parser) }
func GetEmbeddedTemplateParser() *ts.Parser       { return getParser(etParserPool) }
func PutEmbeddedTemplateParser(parser *ts.Parser) { putParser(etParserPool, parser) }
func GetPHPParser() *ts.Parser                    { return getParser(phpParserPool) }
func PutPHPParser(parser *ts.Parser)              { putParser(phpParserPool, parser) }

// ---- Language Accessors ----
// *Language returns the pre-initialized tree-sitter language grammar.
// Used by handlers that need to compile queries against a specific language.

// HTMLLanguage returns the tree-sitter HTML language.
func HTMLLanguage() *ts.Language             { return languages.html }
func CSSLanguage() *ts.Language              { return languages.css }
func JSDocLanguage() *ts.Language            { return languages.jsdoc }
func TypeScriptLanguage() *ts.Language       { return languages.typescript }
func TSXLanguage() *ts.Language              { return languages.tsx }
func BladeLanguage() *ts.Language            { return languages.blade }
func JinjaLanguage() *ts.Language            { return languages.jinja }
func HandlebarsLanguage() *ts.Language       { return languages.handlebars }
func EmbeddedTemplateLanguage() *ts.Language { return languages.embeddedTemplate }
func PHPLanguage() *ts.Language              { return languages.php }

// ---- End Parser Pooling Section ----
