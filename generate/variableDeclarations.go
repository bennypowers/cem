package generate

import (
	"errors"

	M "bennypowers.dev/cem/manifest"
	Q "bennypowers.dev/cem/generate/queries"
)

func generateVarDeclaration(
	captures Q.CaptureMap,
	queryManager *Q.QueryManager,
) (err error, declaration *M.VariableDeclaration) {
	nameNodes, ok := captures["variable.name"]
	if (!ok || len(nameNodes) <= 0) {
		return errors.Join(err, &Q.NoCaptureError{ Capture: "variable.name", Query: "variable" }), nil
	}

	varName := nameNodes[0].Text

	declaration = &M.VariableDeclaration{
		Kind: "variable",
		PropertyLike: M.PropertyLike{
			Name: varName,
		},
	}

	typeNodes, ok := captures["variable.type"]
	if (ok && len(typeNodes) > 0) {
		declaration.Type = &M.Type{
			Text: typeNodes[0].Text,
		}
	}

	jsdoc, ok := captures["variable.jsdoc"]
	if (ok && len(jsdoc) > 0) {
		error, info := NewPropertyInfo(jsdoc[0].Text, queryManager)
		if error != nil {
			err = errors.Join(err, error)
		} else {
			info.MergeToPropertyLike(&declaration.PropertyLike)
		}
	}

	return nil, declaration
}

