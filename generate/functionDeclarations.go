package generate

import (
	"errors"

	M "bennypowers.dev/cem/manifest"
	Q "bennypowers.dev/cem/generate/queries"
)

func generateFunctionDeclaration(
	captures Q.CaptureMap,
	_ any,
	_ []byte,
	queryManager *Q.QueryManager,
) (declaration *M.FunctionDeclaration, err error) {
	nameNodes, ok := captures["function.name"]
	if !ok || len(nameNodes) <= 0 {
		return nil, errors.Join(err, &Q.NoCaptureError{
			Capture: "function.name",
			Query: "functionDeclaration",
		})
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
			rest, isRest := captures["function.param.rest"]
			parameter := M.Parameter{
				Rest: isRest && len(rest) > 0,
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
		error, info := NewMethodInfo(jsdoc[0].Text, queryManager)
		if error != nil {
			return nil, errors.Join(err, error)
		} else {
			info.MergeToFunctionLike(&declaration.FunctionLike)
		}
	}

	return declaration, nil
}
