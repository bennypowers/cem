package generate

import (
	"errors"

	Q "bennypowers.dev/cem/generate/queries"
	M "bennypowers.dev/cem/manifest"
)

func generateVarDeclaration(
	captures Q.CaptureMap,
	queryManager *Q.QueryManager,
) (declaration *M.VariableDeclaration, errs error) {
	nameNodes, ok := captures["variable.name"]
	if !ok || len(nameNodes) <= 0 {
		return nil, errors.Join(errs, &Q.NoCaptureError{Capture: "variable.name", Query: "variable"})
	}

	varName := nameNodes[0].Text

	declaration = &M.VariableDeclaration{
		Kind: "variable",
		PropertyLike: M.PropertyLike{
			Name: varName,
		},
	}

	typeNodes, ok := captures["variable.type"]
	if ok && len(typeNodes) > 0 {
		declaration.Type = &M.Type{
			Text: typeNodes[0].Text,
		}
	}

	jsdoc, ok := captures["variable.jsdoc"]
	if ok && len(jsdoc) > 0 {
		info, err := NewPropertyInfo(jsdoc[0].Text, queryManager)
		if err != nil {
			errs = errors.Join(errs, err)
		} else {
			info.MergeToPropertyLike(&declaration.PropertyLike)
		}
	}

	return declaration, errs
}
