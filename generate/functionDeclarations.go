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
package generate

import (
	"errors"
	"fmt"
	"slices"

	"bennypowers.dev/cem/generate/jsdoc"
	Q "bennypowers.dev/cem/internal/treesitter"
	M "bennypowers.dev/cem/manifest"
	ts "github.com/tree-sitter/go-tree-sitter"
)

func (mp *ModuleProcessor) generateFunctionDeclaration(
	captures Q.CaptureMap,
) (declaration *M.FunctionDeclaration, err error) {
	nameNodes, ok := captures["function.name"]
	if !ok || len(nameNodes) <= 0 {
		return nil, errors.Join(err, &Q.NoCaptureError{
			Capture: "function.name",
			Query:   "functionDeclaration",
		})
	}

	funcName := nameNodes[0].Text

	declaration = &M.FunctionDeclaration{
		Kind: "function",
		FullyQualified: M.FullyQualified{
			Name: funcName,
		},
	}

	// Add source reference if available
	nameNode := Q.GetDescendantById(mp.root, nameNodes[0].NodeId)
	sourceRef, sourceErr := mp.generateSourceReference(nameNode)
	if sourceErr != nil {
		err = errors.Join(err, fmt.Errorf("failed to generate source reference for function %s: %w", funcName, sourceErr))
	} else {
		declaration.Source = sourceRef
	}

	// Generic type parameter names from `<Value extends string = string>`;
	// parameter and return types that reference them get a Type.source
	// pointer to the parameter's declaration
	var bindings []typeParameterBinding
	if typeParams, ok := captures["function.typeParameters"]; ok && len(typeParams) > 0 {
		if node := Q.GetDescendantById(mp.root, typeParams[0].NodeId); node != nil {
			bindings = typeParameterBindings(node, mp.code)
		}
	}

	// One `function.params` capture per parameter node; each query match
	// binds a single parameter, so ParentCaptures aggregates them in match
	// order. Sort by source position to restore declaration order.
	if params, ok := captures["function.params"]; ok && len(params) > 0 {
		slices.SortStableFunc(params, func(a, b Q.CaptureInfo) int {
			return int(a.StartByte) - int(b.StartByte)
		})
		for _, paramCapture := range params {
			node := Q.GetDescendantById(mp.root, paramCapture.NodeId)
			if node == nil {
				continue
			}
			parameter, ok := parameterFromNode(node, mp.code)
			if !ok {
				continue
			}
			declaration.Parameters = append(declaration.Parameters, parameter)
		}
	}

	// Return type from the explicit `: T` annotation, set before JSDoc
	// enrichment so @returns descriptions merge onto it and an @returns
	// type annotation overrides it (mirrors method handling and parameter
	// JSDoc-type behavior).
	if returns, ok := captures["function.return.type"]; ok && len(returns) > 0 {
		declaration.Return = &M.Return{
			Type: &M.Type{
				Text: returns[0].Text,
			},
		}
	}

	jsdocNodes, ok := captures["function.jsdoc"]
	if ok && len(jsdocNodes) > 0 {
		jsdocErr := jsdoc.EnrichFunctionWithJSDoc(jsdocNodes[0].Text, declaration, mp.queryManager)
		if jsdocErr != nil {
			return nil, errors.Join(err, jsdocErr)
		}
	}

	// After JSDoc enrichment so @type-provided types are annotated too
	types := make([]*M.Type, 0, len(declaration.Parameters)+1)
	for _, param := range declaration.Parameters {
		types = append(types, param.Type)
	}
	if declaration.Return != nil {
		types = append(types, declaration.Return.Type)
	}
	mp.annotateTypeParameterSources(bindings, types...)

	return declaration, nil
}

// parameterFromNode builds a manifest Parameter from a required_parameter or
// optional_parameter tree-sitter node. The `?` marker maps to Optional, and
// the `=` initializer's source text to Default; a defaulted parameter is also
// optional, since callers may omit it. Parameter shapes we do not model
// (destructuring patterns, `this` parameters) report false and are omitted.
func parameterFromNode(node *ts.Node, code []byte) (M.Parameter, bool) {
	var parameter M.Parameter
	if node == nil {
		return parameter, false
	}
	switch node.GrammarName() {
	case "required_parameter", "optional_parameter":
	default:
		return parameter, false
	}
	parameter.Optional = node.GrammarName() == "optional_parameter"

	pattern := node.ChildByFieldName("pattern")
	if pattern == nil {
		return parameter, false
	}
	switch pattern.GrammarName() {
	case "identifier":
		parameter.Name = pattern.Utf8Text(code)
	case "rest_pattern":
		parameter.Rest = true
		// the identifier inside a rest_pattern carries no field name
		for i := range pattern.ChildCount() {
			if child := pattern.Child(i); child.GrammarName() == "identifier" {
				parameter.Name = child.Utf8Text(code)
				break
			}
		}
		// e.g. `...[a, b]`: destructuring rest has no name to expose
		if parameter.Name == "" {
			return parameter, false
		}
	default:
		return parameter, false
	}

	if typeAnnotation := node.ChildByFieldName("type"); typeAnnotation != nil {
		if typeExpr := typeAnnotation.NamedChild(0); typeExpr != nil {
			parameter.Type = &M.Type{
				Text: typeExpr.Utf8Text(code),
			}
		}
	}
	if value := node.ChildByFieldName("value"); value != nil {
		parameter.Default = value.Utf8Text(code)
		parameter.Optional = true
	}
	return parameter, true
}
