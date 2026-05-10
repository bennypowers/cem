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
package css

import (
	"fmt"

	csslang "bennypowers.dev/cem/internal/languages/css"
	"bennypowers.dev/cem/lsp/document/base"
)

// CSSDocument represents a CSS document with tree-sitter parsing.
type CSSDocument struct {
	*base.BaseDocument
}

// NewCSSDocument creates a new CSS document
func NewCSSDocument(uri, content string, version int32) *CSSDocument {
	doc := &CSSDocument{}
	doc.BaseDocument = base.NewBaseDocument(uri, content, version, "css", csslang.ReturnParser, doc)
	return doc
}

// Parse parses the CSS content using tree-sitter
func (d *CSSDocument) Parse(content string) error {
	d.UpdateContent(content, d.Version())

	parser := csslang.BorrowParser()
	if parser == nil {
		return fmt.Errorf("failed to get CSS parser")
	}
	d.SetParser(parser)

	tree := parser.Parse([]byte(content), nil)
	d.SetTree(tree)

	return nil
}
