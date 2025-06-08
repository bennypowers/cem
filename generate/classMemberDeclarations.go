package generate

import (
	"errors"
	"regexp"
	"strings"

	M "bennypowers.dev/cem/manifest"
	"bennypowers.dev/cem/set"
	"dario.cat/mergo"
	ts "github.com/tree-sitter/go-tree-sitter"
)

var ignoredStaticFields = set.NewSet("formAssociated", "styles", "shadowRootOptions")

func createClassFieldFromAccessorMatch(
	fieldName string,
	accessor string,
	isStatic bool,
	isReadonly bool,
	captures CaptureMap,
) (err error, field M.CustomElementField) {
	field.Kind = "field"
	field.Static = isStatic
	field.Name = fieldName
	field.Readonly = isReadonly

	for _, x := range captures["field.type"] {
		typeText := x.Text
		if typeText != "" {
			field.Type = &M.Type{
				Text: typeText,
			}
		}
	}

	for _, x := range captures["member.jsdoc"] {
		source := x.Text
		if strings.HasPrefix(source, `/**`) {
			error, info := NewFieldInfo(source)
			if error != nil {
				err = errors.Join(err, error)
			} else {
				field.Description += info.Description
				field.Summary += info.Summary
				field.Deprecated = info.Deprecated
				if info.Type != "" {
					field.Type = &M.Type{
						Text: info.Type,
					}
				}
			}
		}
	}

	for _, x := range captures["field.initializer"] {
		field.Default = x.Text
		if field.Type == nil {
			if (field.Default == "true" || field.Default == "false") {
				field.Type = &M.Type{
					Text: "boolean",
				}
			} else if regexp.MustCompile(`^[\d._]+$`).MatchString(field.Default) {
				field.Type = &M.Type{
					Text: "number",
				}
			} else if regexp.MustCompile(`^('|")`).MatchString(field.Default) {
				field.Type = &M.Type{
					Text: "string",
				}
			}
		}
	}

	for _, x := range captures["field.privacy"] {
		switch x.Text {
			case "private":
				field.Privacy = M.Private
			case "protected":
				field.Privacy = M.Protected
		}
	}

	field.Reflects = !isStatic && len(captures["field.attr.reflects"]) > 0 && captures["field.attr.reflects"][0].Text == "true"

	_, hasDecorators := captures["decorator.name"]
	var isProperty bool
	if hasDecorators {
		for _, deco := range captures["decorator.name"] {
			if deco.Text == "property" {
				isProperty = true
				break
			}
		}
	}

	if (isProperty) {
		attributeBoolNodes := captures["field.attr.bool"]
		attributeNameNodes := captures["field.attr.name"]

		if (!isStatic &&
				len(attributeBoolNodes) == 0 ||
				(len(attributeBoolNodes) > 0 && attributeBoolNodes[0].Node.GrammarName() != "false")) {
			if len(attributeNameNodes) > 0 {
				field.Attribute = attributeNameNodes[0].Text
			} else {
				field.Attribute = strings.ToLower(field.Name)
			}
		}
	}

	return err, field
}

func createClassFieldFromFieldMatch(fieldName string, isStatic bool, captures CaptureMap) (err error, field M.CustomElementField) {
	_, readonly := captures["field.readonly"]

	field.Kind = "field"
	field.Static = isStatic
	field.Name = fieldName
	field.Readonly = readonly

	for _, x := range captures["field.type"] {
		typeText := x.Text
		if typeText != "" {
			field.Type = &M.Type{
				Text: typeText,
			}
		}
	}

	for _, x := range captures["member.jsdoc"] {
		source := x.Text
		if strings.HasPrefix(source, `/**`) {
			error, info := NewFieldInfo(source)
			if error != nil {
				err = errors.Join(err, error)
			}
			field.Description += info.Description
			field.Summary += info.Summary
			field.Deprecated = info.Deprecated
			if info.Type != "" {
				field.Type = &M.Type{
					Text: info.Type,
				}
			}
		}
	}

	for _, x := range captures["field.initializer"] {
		field.Default = x.Text
		if field.Type == nil {
			if (field.Default == "true" || field.Default == "false") {
				field.Type = &M.Type{
					Text: "boolean",
				}
			} else if regexp.MustCompile(`^[\d._]+$`).MatchString(field.Default) {
				field.Type = &M.Type{
					Text: "number",
				}
			} else if regexp.MustCompile(`^('|")`).MatchString(field.Default) {
				field.Type = &M.Type{
					Text: "string",
				}
			}
		}
	}

	for _, x := range captures["field.privacy"] {
		switch x.Text {
			case "private":
				field.Privacy = M.Private
			case "protected":
				field.Privacy = M.Protected
		}
	}

	field.Reflects = !field.Static && len(captures["field.attr.reflects"]) > 0 && captures["field.attr.reflects"][0].Text == "true"

	attributeBoolNodes := captures["field.attr.bool"]
	attributeNameNodes := captures["field.attr.name"]

	if (!isStatic &&
			len(attributeBoolNodes) == 0 ||
			(len(attributeBoolNodes) > 0 && attributeBoolNodes[0].Node.GrammarName() != "false")) {
		if len(attributeNameNodes) > 0 {
			field.Attribute = attributeNameNodes[0].Text
		} else {
			field.Attribute = strings.ToLower(field.Name)
		}
	}

	return err, field
}

func getClassMembersFromClassDeclarationNode(code []byte, root *ts.Node) (errs error, members []M.ClassMember) {
	qm, closeQm := NewQueryMatcher("classMemberDeclaration", Languages.Typescript)
	defer closeQm()

	staticsSet := set.NewSet[string]()
	fieldsSet := set.NewSet[string]()
	gettersSet := set.NewSet[string]()
	settersSet := set.NewSet[string]()
	methodsSet := set.NewSet[string]()

	// to collate first instance of accessor pair / accessor
	accessorIndices := make(map[string]int)

	for match := range qm.AllQueryMatches(root, code) {
		captures := qm.GetCapturesFromMatch(match, code)
		_, isAccessor := captures["accessor"]
		_, isStatic := captures["field.static"]
		_, isField := captures["field"]
		_, isMethod := captures["method"]

		switch {
		case isField:
			fieldName := captures["field.name"][0].Text
			if isStatic && ignoredStaticFields.Has(fieldName) {
				continue
			}

			if (isField && isStatic && !staticsSet.Has(fieldName)) {
				staticsSet.Add(fieldName)
				error, field := createClassFieldFromFieldMatch(fieldName, isStatic, captures)
				if error != nil {
					errs = errors.Join(errs, error)
				} else {
					members = append(members, field)
				}
			} else if (isField && !isStatic && !fieldsSet.Has(fieldName)) {
				fieldsSet.Add(fieldName)
				error, field := createClassFieldFromFieldMatch(fieldName, isStatic, captures)
				if error != nil {
					errs = errors.Join(errs, error)
				} else {
					members = append(members, field)
				}
			}
		case isAccessor:
			fieldName := captures["field.name"][0].Text
			if !gettersSet.Has(fieldName) && !settersSet.Has(fieldName) {
				accessorIndices[fieldName] = len(members)
			}

			if gettersSet.Has(fieldName) && settersSet.Has(fieldName) {
				continue
			}

			fieldsSet.Add(fieldName)
			accessorKind := captures["field.accessor"][0].Text
			var hasPair bool
			switch accessorKind {
			case "get":
				hasPair = settersSet.Has(fieldName)
				if gettersSet.Has(fieldName) {
					continue
				} else {
					gettersSet.Add(fieldName)
				}
			case "set":
				hasPair = gettersSet.Has(fieldName)
				if settersSet.Has(fieldName) {
					continue
				} else {
					settersSet.Add(fieldName)
				}
			}
			index := accessorIndices[fieldName]
			if hasPair {
				// field is a *cem.CustomElementField
				error, field := createClassFieldFromAccessorMatch(fieldName, accessorKind, isStatic, false, captures)
				if error != nil {
					errs = errors.Join(errs, error)
				} else {
					item := members[index].(M.CustomElementField)
					item.Readonly = false
					err := mergo.Merge(&item, field, mergo.WithOverride)
					if err != nil {
						err = errors.Join(err, error)
					} else {
						members[index] = item
					}
				}
			} else {
				error, field := createClassFieldFromAccessorMatch(fieldName, accessorKind, isStatic, accessorKind == "get", captures)
				if error != nil {
					errs = errors.Join(errs, error)
				} else {
					members = append(members, field)
				}
			}
		case isMethod:
			method := M.ClassMethod{Kind: "method"}

			captures := qm.GetCapturesFromMatch(match, code)

			methodNames, ok := captures["method.name"]
			methodName := methodNames[0]
			if ok {
				method.Name = methodName.Text
			}

			if methodsSet.Has(method.Name) {
				continue
			}

			methodsSet.Add(method.Name)

			privacyNodes, ok := captures["method.privacy"]
			if ok && len(privacyNodes) > 0 {
				privacy := privacyNodes[0]
				method.Privacy = M.Privacy(privacy.Text)
			}

			staticNodes, ok := captures["method.static"]
			if ok && len(staticNodes) > 0 {
				method.Static = true
			}

			params, ok := captures["params"]
			if ok && len(params) > 0 {
				node := params[0].Node
				for _, param := range node.NamedChildren(node.Walk()) {
					parameter := M.Parameter{ }
					nameNode := param.ChildByFieldName("pattern")
					typeParentNode := param.ChildByFieldName("type")
					if typeParentNode != nil {
						typeNode := typeParentNode.NamedChild(0)
						if typeNode != nil {
							parameter.Type = &M.Type{
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
						method.Parameters = append(method.Parameters, parameter)
					}
				}
			}

			returnNodes, ok := captures["method.returns"]
			if ok && len(returnNodes) > 0 {
				returns := returnNodes[0]
				method.Return = &M.Return{
					Type: &M.Type{
						Text: returns.Text,
					},
				}
			}

			jsdocs, ok := captures["member.jsdoc"]
			if ok {
				for _, jsdoc := range jsdocs {
					merr, info := NewMethodInfo(jsdoc.Text)
					if merr != nil {
						errs = errors.Join(merr)
					} else {
						method.Description = info.Description
						method.Deprecated = info.Deprecated
						method.Summary = info.Summary
						if (info.Return != nil) {
							if method.Return == nil {
								method.Return = &M.Return{}
							}
							method.Return.Description = info.Return.Description
							if info.Return.Type != "" {
								method.Return.Type = &M.Type{
									Text: info.Return.Type,
								}
							}
						}
						if info.Privacy != "" {
							method.Privacy = info.Privacy
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
			members = append(members, method)

		}
	}

	return errs, members
}

