/*
Copyright © 2025 Benny Powers

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
package jsdoc

import (
	"strings"

	M "bennypowers.dev/cem/manifest"
	Q "bennypowers.dev/cem/queries"
	ts "github.com/tree-sitter/go-tree-sitter"
)

// ExtractFromNode returns the immediately preceding JSDoc comment (/** ... */) for the given node,
// skipping over decorator nodes. If there is no such comment, returns "".
func ExtractFromNode(node *ts.Node, source []byte) string {
	if node == nil {
		return ""
	}
	for prev := node.PrevNamedSibling(); prev != nil; prev = prev.PrevNamedSibling() {
		kind := prev.GrammarName()
		switch kind {
		case "decorator":
			continue
		case "comment":
			text := prev.Utf8Text(source)
			if strings.HasPrefix(text, "/**") {
				return text
			} else {
				return ""
			}
		default:
			return ""
		}
	}
	return ""
}

// EnrichClassWithJSDoc parses JSDoc comment text and applies the extracted information
// to a ClassDeclaration.
func EnrichClassWithJSDoc(jsdocText string, decl *M.ClassDeclaration, queryManager *Q.QueryManager) error {
	info, err := parseForClass(jsdocText, queryManager)
	if err != nil {
		return err
	}
	applyToClassDeclaration(info, decl)
	return nil
}

// EnrichCustomElementWithJSDoc parses JSDoc comment text and applies the extracted information
// to a CustomElementDeclaration.
func EnrichCustomElementWithJSDoc(jsdocText string, decl *M.CustomElementDeclaration, queryManager *Q.QueryManager) error {
	info, err := parseForClass(jsdocText, queryManager)
	if err != nil {
		return err
	}
	applyToCustomElementDeclaration(info, decl)
	return nil
}

// EnrichMethodWithJSDoc parses JSDoc comment text and applies the extracted information
// to a ClassMethod.
func EnrichMethodWithJSDoc(jsdocText string, method *M.ClassMethod, queryManager *Q.QueryManager) error {
	info, err := parseForMethod(jsdocText, queryManager)
	if err != nil {
		return err
	}
	applyToMethod(info, method)
	return nil
}

// EnrichFunctionWithJSDoc parses JSDoc comment text and applies the extracted information
// to a FunctionDeclaration.
func EnrichFunctionWithJSDoc(jsdocText string, fn *M.FunctionDeclaration, queryManager *Q.QueryManager) error {
	info, err := parseForMethod(jsdocText, queryManager)
	if err != nil {
		return err
	}
	applyToFunctionDeclaration(info, fn)
	return nil
}

// EnrichPropertyWithJSDoc parses JSDoc comment text and applies the extracted information
// to a PropertyLike.
func EnrichPropertyWithJSDoc(jsdocText string, prop *M.PropertyLike, queryManager *Q.QueryManager) error {
	info, err := parseForProperty(jsdocText, queryManager)
	if err != nil {
		return err
	}
	applyToPropertyLike(info, prop)
	return nil
}

// EnrichCSSPropertyWithJSDoc parses JSDoc comment text and applies the extracted information
// to a CssCustomProperty.
func EnrichCSSPropertyWithJSDoc(jsdocText string, prop *M.CssCustomProperty, queryManager *Q.QueryManager) error {
	info, err := parseForCSSProperty(jsdocText, queryManager)
	if err != nil {
		return err
	}
	applyToCSSProperty(info, prop)
	return nil
}

// ExtractAliasFromJSDoc parses JSDoc comment text and returns the @alias value if present.
func ExtractAliasFromJSDoc(jsdocText string, queryManager *Q.QueryManager) (string, error) {
	info, err := parseForClass(jsdocText, queryManager)
	if err != nil {
		return "", err
	}
	return info.Alias, nil
}