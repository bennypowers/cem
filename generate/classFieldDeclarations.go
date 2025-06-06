package generate

import (
	"errors"
	"regexp"
	"strings"

	"bennypowers.dev/cem/manifest"
	"bennypowers.dev/cem/set"
	"dario.cat/mergo"
	ts "github.com/tree-sitter/go-tree-sitter"
)

var ignoredStaticFields = set.NewSet("formAssociated", "styles", "shadowRootOptions")

func createClassFieldFromAccessorMatch(
	fieldName string,
	accessor string,
	isStatic bool,
	captures CaptureMap,
) (err error, field manifest.CustomElementField) {
	field.Kind = "field"
	field.Static = isStatic
	field.Name = fieldName
	field.Readonly = accessor == "get"

	for _, x := range captures["field.type"] {
		typeText := x.Text
		if typeText != "" {
			field.Type = &manifest.Type{
				Text: typeText,
			}
		}
	}

	for _, x := range captures["field.jsdoc"] {
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
					field.Type = &manifest.Type{
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
				field.Type = &manifest.Type{
					Text: "boolean",
				}
			} else if regexp.MustCompile(`^[\d._]+$`).MatchString(field.Default) {
				field.Type = &manifest.Type{
					Text: "number",
				}
			} else if regexp.MustCompile(`^('|")`).MatchString(field.Default) {
				field.Type = &manifest.Type{
					Text: "string",
				}
			}
		}
	}

	for _, x := range captures["field.privacy"] {
		switch x.Text {
			case "private":
				field.Privacy = manifest.Private
			case "protected":
				field.Privacy = manifest.Protected
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

func createClassFieldFromFieldMatch(fieldName string, isStatic bool, captures CaptureMap) (err error, field manifest.CustomElementField) {
	_, readonly := captures["field.readonly"]

	field.Kind = "field"
	field.Static = isStatic
	field.Name = fieldName
	field.Readonly = readonly

	for _, x := range captures["field.type"] {
		typeText := x.Text
		if typeText != "" {
			field.Type = &manifest.Type{
				Text: typeText,
			}
		}
	}

	for _, x := range captures["field.jsdoc"] {
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
				field.Type = &manifest.Type{
					Text: info.Type,
				}
			}
		}
	}

	for _, x := range captures["field.initializer"] {
		field.Default = x.Text
		if field.Type == nil {
			if (field.Default == "true" || field.Default == "false") {
				field.Type = &manifest.Type{
					Text: "boolean",
				}
			} else if regexp.MustCompile(`^[\d._]+$`).MatchString(field.Default) {
				field.Type = &manifest.Type{
					Text: "number",
				}
			} else if regexp.MustCompile(`^('|")`).MatchString(field.Default) {
				field.Type = &manifest.Type{
					Text: "string",
				}
			}
		}
	}

	for _, x := range captures["field.privacy"] {
		switch x.Text {
			case "private":
				field.Privacy = manifest.Private
			case "protected":
				field.Privacy = manifest.Protected
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

func getClassFieldsFromClassDeclarationNode(code []byte, node *ts.Node) (err error, field []manifest.CustomElementField) {
	qm, closeQm := NewQueryMatcher("classFieldDeclaration", Languages.Typescript)
	defer closeQm()

	staticsSet := set.NewSet[string]()
	fieldsSet := set.NewSet[string]()
	gettersSet := set.NewSet[string]()
	settersSet := set.NewSet[string]()

	// to collate first instance of accessor pair / accessor
	accessorIndices := make(map[string]int)

	// final list of fields, in source order
	fields := make([]manifest.CustomElementField, 0)

	for match := range qm.AllQueryMatches(node, code) {
		captures := qm.GetCapturesFromMatch(match, code)
		fieldName := captures["field.name"][0].Text
		_, isAccessor := captures["accessor"]
		_, isField := captures["field"]
		_, isStatic := captures["field.static"]

		if isStatic && ignoredStaticFields.Has(fieldName) {
			continue
		}

		if (isField && isStatic && !staticsSet.Has(fieldName)) {
			staticsSet.Add(fieldName)
			error, field := createClassFieldFromFieldMatch(fieldName, isStatic, captures)
			if error != nil {
				err = errors.Join(err, error)
			} else {
				fields = append(fields, field)
			}
		} else if (isField && !isStatic && !fieldsSet.Has(fieldName)) {
			fieldsSet.Add(fieldName)
			error, field := createClassFieldFromFieldMatch(fieldName, isStatic, captures)
			if error != nil {
				err = errors.Join(err, error)
			} else {
				fields = append(fields, field)
			}
		} else if isAccessor {
			if !gettersSet.Has(fieldName) && !settersSet.Has(fieldName) {
				accessorIndices[fieldName] = len(fields)
			}

			if gettersSet.Has(fieldName) && settersSet.Has(fieldName) {
				continue
			}

			fieldsSet.Add(fieldName)
			accessor := captures["field.accessor"][0].Text
			var hasPair bool
			switch accessor {
				case "get":
					if gettersSet.Has(fieldName) {
						continue
					} else {
						gettersSet.Add(fieldName)
					}
					hasPair = settersSet.Has(fieldName)
				case "set":
					if settersSet.Has(fieldName) {
						continue
					} else {
						settersSet.Add(fieldName)
					}
					hasPair = gettersSet.Has(fieldName)
			}
			index := accessorIndices[fieldName]
			if hasPair {
				// field is a *cem.CustomElementField
				error, field := createClassFieldFromAccessorMatch(fieldName, accessor, isStatic, captures)
				if error != nil {
					err = errors.Join(err, error)
				} else {
					// this next line fails at runtime with "dst must be a pointer"
					err := mergo.Merge(&fields[index], field, mergo.WithOverride)
					if err != nil {
						err = errors.Join(err, error)
					} else {
						fields[index].Readonly = false
					}
				}
			} else {
				error, field := createClassFieldFromAccessorMatch(fieldName, accessor, isStatic, captures)
				if error != nil {
					err = errors.Join(err, error)
				} else {
					fields = append(fields, field)
				}
			}
		}
	}

	return err, fields
}

