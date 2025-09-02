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

	M "bennypowers.dev/cem/manifest"
	Q "bennypowers.dev/cem/queries"
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

	params, ok := captures["function.params"]
	if ok && len(params) > 0 {
		_, hasName := captures["function.param.name"]
		if hasName {
			rest, isRest := captures["function.param.rest"]
			parameter := M.Parameter{
				Rest: isRest && len(rest) > 0,
				PropertyLike: M.PropertyLike{
					FullyQualified: M.FullyQualified{
						Name: captures["function.param.name"][0].Text,
					},
				},
			}
			_, hasType := captures["function.param.type"]
			if hasType {
				parameter.Type = &M.Type{
					Text: captures["function.param.type"][0].Text,
				}
			}
			declaration.Parameters = append(declaration.Parameters, parameter)
		}
	}

	jsdoc, ok := captures["function.jsdoc"]
	if ok && len(jsdoc) > 0 {
		error, info := NewMethodInfo(jsdoc[0].Text, mp.queryManager)
		if error != nil {
			return nil, errors.Join(err, error)
		} else {
			info.MergeToFunctionDeclaration(declaration)
		}
	}

	return declaration, nil
}
