package generate

import (
	"errors"

	M "bennypowers.dev/cem/manifest"
)

func generateFunctionDeclaration(captures CaptureMap, _ any, code []byte) (err error, declaration *M.FunctionDeclaration) {
	nameNodes, ok := captures["function.name"]
	if (!ok || len(nameNodes) <= 0) {
		return errors.Join(err, &NoCaptureError{ "function.name", "functionDeclaration" }), nil
	}

	funcName := nameNodes[0].Text

	declaration = &M.FunctionDeclaration{
		Kind: "function",
		FunctionLike: M.FunctionLike{
			Name: funcName,
		},
	}

	params, ok := captures["function.params"]
	if ok && len(params) > 0 {
		_, hasName := captures["function.param.name"]
		if hasName {
			_, isRest := captures["function.param.rest"]
			parameter := M.Parameter{
				Rest: isRest,
				PropertyLike: M.PropertyLike{
					Name: captures["function.param.name"][0].Text,
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
	if (ok && len(jsdoc) > 0) {
		error, info := NewMethodInfo(jsdoc[0].Text)
		if error != nil {
			return errors.Join(err, error), nil
		} else {
			info.MergeToFunctionLike(&declaration.FunctionLike)
		}
	}

	return nil, declaration
}
