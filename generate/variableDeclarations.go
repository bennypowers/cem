package generate

import (
	"errors"

	"bennypowers.dev/cem/manifest"
)

func generateVarDeclaration(captures CaptureMap, queryManager *QueryManager) (err error, declaration *manifest.VariableDeclaration) {
	nameNodes, ok := captures["variable.name"]
	if (!ok || len(nameNodes) <= 0) {
		return errors.Join(err, &NoCaptureError{ "variable.name", "variable" }), nil
	}

	varName := nameNodes[0].Text

	declaration = &manifest.VariableDeclaration{
		Kind: "variable",
		PropertyLike: manifest.PropertyLike{
			Name: varName,
		},
	}

	typeNodes, ok := captures["variable.type"]
	if (ok && len(typeNodes) > 0) {
		declaration.Type = &manifest.Type{
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

