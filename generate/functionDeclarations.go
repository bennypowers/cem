package generate

import (
	"errors"

	"bennypowers.dev/cem/manifest"
)

func generateFunctionDeclaration(captures CaptureMap, _ any, code []byte) (err error, declaration *manifest.FunctionDeclaration) {
	nameNodes, ok := captures["function.name"]
	if (!ok || len(nameNodes) <= 0) {
		return errors.Join(err, &NoCaptureError{ "function.name", "functionDeclaration" }), nil
	}

	funcName := nameNodes[0].Text

	declaration = &manifest.FunctionDeclaration{
		Kind: "function",
		FunctionLike: manifest.FunctionLike{
			Name: funcName,
		},
	}

	paramNodes, ok := captures["function.params"]
	if ok && len(paramNodes) > 0 {
		node := paramNodes[0].Node
		for _, param := range node.NamedChildren(node.Walk()) {
			parameter := manifest.Parameter{ }
			nameNode := param.ChildByFieldName("pattern")
			typeParentNode := param.ChildByFieldName("type")
			if typeParentNode != nil {
				typeNode := typeParentNode.NamedChild(0)
				if typeNode != nil {
					parameter.Type = &manifest.Type{
						Text: typeNode.Utf8Text(code),
					}
				}
			}
			if nameNode != nil {
				if nameNode.GrammarName() == "rest_pattern" {
					parameter.Rest = true
					nameNode = nameNode.NamedChild(0)
				}
				parameter.Name = nameNode.Utf8Text(code)
			}
			if parameter.Name != "" {
				declaration.Parameters = append(declaration.Parameters, parameter)
			}
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
