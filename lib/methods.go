package generate

import (
	"log"

	"github.com/bennypowers/cemgen/cem"
	ts "github.com/tree-sitter/go-tree-sitter"
)

func getClassMethodsFromClassDeclarationNode(
	language *ts.Language,
	code []byte,
	node *ts.Node,
) []cem.ClassMethod {
	methods := make([]cem.ClassMethod, 0)
	queryText, err := loadQueryFile("classMethod")
	if err != nil {
		log.Fatal(err)
	}
	query, qerr := ts.NewQuery(language, queryText)
	defer query.Close()
	if qerr != nil {
		log.Fatal(qerr)
	}
	cursor := ts.NewQueryCursor()
	defer cursor.Close()

	set := make(map[string]bool)

	for match := range allMatches(cursor, query, node, code) {
		method := cem.ClassMethod{Kind: "method"}

		captures := *getCapturesFromMatch(match, query)

		methodNames, ok := captures["method.name"]
		methodName := methodNames[0]
		if ok {
			method.Name = methodName.Utf8Text(code)
		}

		if set[method.Name] {
			continue
		}

		set[method.Name] = true

		privacyNodes, ok := captures["method.privacy"]
		if ok && len(privacyNodes) > 0 {
			privacy := privacyNodes[0]
			method.Privacy = cem.Privacy(privacy.Utf8Text(code))
		}

		staticNodes, ok := captures["method.static"]
		if ok && len(staticNodes) > 0 {
			method.Static = true
		}
		params, ok := captures["params"]
		if ok && len(params) > 0 {
			for _, param := range params[0].NamedChildren(params[0].Walk()) {
				nameNode := param.ChildByFieldName("pattern")
				typeNode := param.ChildByFieldName("type").NamedChild(0)
				parameter := cem.Parameter{ }
				if nameNode != nil {
					if nameNode.GrammarName() == "rest_pattern" {
						parameter.Rest = true
						nameNode = nameNode.NamedChild(0)
					}
					parameter.Name = nameNode.Utf8Text(code)
				}
				if typeNode != nil {
					parameter.Type = &cem.Type{
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
			method.Return = &cem.Return{
				Type: &cem.Type{
					Text: returns.Utf8Text(code),
				},
			}
		}

		jsdocs, ok := captures["method.jsdoc"]
		if ok {
			for _, jsdoc := range jsdocs {
				info := NewMethodInfo(jsdoc.Utf8Text(code))
				method.Description = info.Description
				method.Deprecated = info.Deprecated
				method.Summary = info.Summary
				if (info.Return != nil) {
					if method.Return == nil {
						method.Return = &cem.Return{}
					}
				}
				method.Return.Description = info.Return.Description
				if info.Privacy != "" {
					method.Privacy = info.Privacy
				}
				if info.Return.Type != "" {
					method.Return.Type = &cem.Type{
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

		methods = append(methods, method)
	}

	return methods
}
