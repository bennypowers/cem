package generate

import (
	"errors"

	"bennypowers.dev/cem/manifest"
	"bennypowers.dev/cem/set"
	ts "github.com/tree-sitter/go-tree-sitter"
)

func getClassMethodsFromClassDeclarationNode(code []byte, node *ts.Node) (err error, methods []manifest.ClassMethod) {
	qm, closeQm := NewQueryMatcher("classMethod", Languages.Typescript)
	defer closeQm()

	set := set.NewSet[string]()

	for match := range qm.AllQueryMatches(node, code) {
		method := manifest.ClassMethod{Kind: "method"}

		captures := qm.GetCapturesFromMatch(match, code)

		methodNames, ok := captures["method.name"]
		methodName := methodNames[0]
		if ok {
			method.Name = methodName.Text
		}

		if set.Has(method.Name) {
			continue
		}

		set.Add(method.Name)

		privacyNodes, ok := captures["method.privacy"]
		if ok && len(privacyNodes) > 0 {
			privacy := privacyNodes[0]
			method.Privacy = manifest.Privacy(privacy.Text)
		}

		staticNodes, ok := captures["method.static"]
		if ok && len(staticNodes) > 0 {
			method.Static = true
		}
		params, ok := captures["params"]
		if ok && len(params) > 0 {
			node := params[0].Node
			for _, param := range node.NamedChildren(node.Walk()) {
				nameNode := param.ChildByFieldName("pattern")
				typeNode := param.ChildByFieldName("type").NamedChild(0)
				parameter := manifest.Parameter{ }
				if nameNode != nil {
					if nameNode.GrammarName() == "rest_pattern" {
						parameter.Rest = true
						nameNode = nameNode.NamedChild(0)
					}
					parameter.Name = nameNode.Utf8Text(code)
				}
				if typeNode != nil {
					parameter.Type = &manifest.Type{
						Text: typeNode.Utf8Text(code),
					}
				}
				if parameter.Name != "" {
					method.Parameters = append(method.Parameters, parameter)
				}
			}
		}

		returnNodes, ok := captures["method.returns"]
		if ok && len(returnNodes) > 0 {
			returns := returnNodes[0]
			method.Return = &manifest.Return{
				Type: &manifest.Type{
					Text: returns.Text,
				},
			}
		}

		jsdocs, ok := captures["method.jsdoc"]
		if ok {
			for _, jsdoc := range jsdocs {
				merr, info := NewMethodInfo(jsdoc.Text)
				if merr != nil {
					err = errors.Join(merr)
				} else {
					method.Description = info.Description
					method.Deprecated = info.Deprecated
					method.Summary = info.Summary
					if (info.Return != nil) {
						if method.Return == nil {
							method.Return = &manifest.Return{}
						}
					}
					method.Return.Description = info.Return.Description
					if info.Privacy != "" {
						method.Privacy = info.Privacy
					}
					if info.Return.Type != "" {
						method.Return.Type = &manifest.Type{
							Text: info.Return.Type,
						}
					}
					for _, iparam := range info.Parameters {
						for i, _ := range method.Parameters {
							if method.Parameters[i].Name == iparam.Name {
								method.Parameters[i].Description = iparam.Description
								method.Parameters[i].Deprecated = iparam.Deprecated
								if iparam.Optional {
									method.Parameters[i].Optional = true
								}
								if iparam.Type != "" {
									method.Parameters[i].Type.Text = iparam.Type
								}
								if iparam.Default != "" {
									method.Parameters[i].Default = iparam.Default
								}
							}
						}
					}
				}
			}
		}

		methods = append(methods, method)
	}

	return err, methods
}
