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

	"bennypowers.dev/cem/generate/jsdoc"
	M "bennypowers.dev/cem/manifest"
	Q "bennypowers.dev/cem/queries"
)

func (mp *ModuleProcessor) generateVarDeclaration(
	captures Q.CaptureMap,
) (declaration *M.VariableDeclaration, errs error) {
	nameNodes, ok := captures["variable.name"]
	if !ok || len(nameNodes) <= 0 {
		return nil, errors.Join(errs, &Q.NoCaptureError{Capture: "variable.name", Query: "variable"})
	}

	varName := nameNodes[0].Text

	declaration = &M.VariableDeclaration{
		Kind: "variable",
		PropertyLike: M.PropertyLike{
			FullyQualified: M.FullyQualified{
				Name: varName,
			},
		},
	}

	// Add source reference if available
	nameNode := Q.GetDescendantById(mp.root, nameNodes[0].NodeId)
	sourceRef, sourceErr := mp.generateSourceReference(nameNode)
	if sourceErr != nil {
		errs = errors.Join(errs, fmt.Errorf("failed to generate source reference for variable %s: %w", varName, sourceErr))
	} else {
		declaration.Source = sourceRef
	}

	typeNodes, ok := captures["variable.type"]
	if ok && len(typeNodes) > 0 {
		declaration.Type = &M.Type{
			Text: typeNodes[0].Text,
		}
	}

	jsdocNodes, ok := captures["variable.jsdoc"]
	if ok && len(jsdocNodes) > 0 {
		err := jsdoc.EnrichPropertyWithJSDoc(jsdocNodes[0].Text, &declaration.PropertyLike, mp.queryManager)
		if err != nil {
			errs = errors.Join(errs, err)
		}
	}

	return declaration, errs
}
