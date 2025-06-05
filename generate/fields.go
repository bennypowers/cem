package generate

import (
	"log"
	"regexp"
	"strings"

	"bennypowers.dev/cem/cem"
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
) (*cem.CustomElementField) {
	classField := cem.ClassField{
		Kind: "field",
		Static: isStatic,
		PropertyLike: cem.PropertyLike{
			Name: fieldName,
			Readonly: accessor == "get",
		},
	}

	for _, x := range captures["field.type"] {
		typeText := x.Text
		if typeText != "" {
			classField.Type = &cem.Type{
				Text: typeText,
			}
		}
	}

	for _, x := range captures["field.jsdoc"] {
		source := x.Text
		if strings.HasPrefix(source, `/**`) {
			info := NewFieldInfo(source)
			classField.Description += info.Description
			classField.Summary += info.Summary
			classField.Deprecated = info.Deprecated
			if info.Type != "" {
				classField.Type = &cem.Type{
					Text: info.Type,
				}
			}
		}
	}

	for _, x := range captures["field.initializer"] {
		classField.Default = x.Text
		if classField.Type == nil {
			if (classField.Default == "true" || classField.Default == "false") {
				classField.Type = &cem.Type{
					Text: "boolean",
				}
			} else if regexp.MustCompile(`^[\d._]+$`).MatchString(classField.Default) {
				classField.Type = &cem.Type{
					Text: "number",
				}
			} else if regexp.MustCompile(`^('|")`).MatchString(classField.Default) {
				classField.Type = &cem.Type{
					Text: "string",
				}
			}
		}
	}

	for _, x := range captures["field.privacy"] {
		switch x.Text {
			case "private":
				classField.Privacy = cem.Private
			case "protected":
				classField.Privacy = cem.Protected
		}
	}

	ceField := cem.CustomElementField{
		ClassField: classField,
		Reflects: !classField.Static && len(captures["field.attr.reflects"]) > 0 && captures["field.attr.reflects"][0].Text == "true",
	}

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

		if (!classField.Static &&
				len(attributeBoolNodes) == 0 ||
				(len(attributeBoolNodes) > 0 && attributeBoolNodes[0].Node.GrammarName() != "false")) {
			if len(attributeNameNodes) > 0 {
				ceField.Attribute = attributeNameNodes[0].Text
			} else {
				ceField.Attribute = strings.ToLower(classField.Name)
			}
		}
	}

	for k := range captures { delete(captures, k) }
	return &ceField
}

func createClassFieldFromFieldMatch(fieldName string, isStatic bool, captures CaptureMap) (*cem.CustomElementField) {
	_, readonly := captures["field.readonly"]

	classField := cem.ClassField{
		Kind: "field",
		Static: isStatic,
		PropertyLike: cem.PropertyLike{
			Name: fieldName,
			Readonly: readonly,
		},
	}

	for _, x := range captures["field.type"] {
		typeText := x.Text
		if typeText != "" {
			classField.Type = &cem.Type{
				Text: typeText,
			}
		}
	}

	for _, x := range captures["field.jsdoc"] {
		source := x.Text
		if strings.HasPrefix(source, `/**`) {
			info := NewFieldInfo(source)
			classField.Description += info.Description
			classField.Summary += info.Summary
			classField.Deprecated = info.Deprecated
			if info.Type != "" {
				classField.Type = &cem.Type{
					Text: info.Type,
				}
			}
		}
	}

	for _, x := range captures["field.initializer"] {
		classField.Default = x.Text
		if classField.Type == nil {
			if (classField.Default == "true" || classField.Default == "false") {
				classField.Type = &cem.Type{
					Text: "boolean",
				}
			} else if regexp.MustCompile(`^[\d._]+$`).MatchString(classField.Default) {
				classField.Type = &cem.Type{
					Text: "number",
				}
			} else if regexp.MustCompile(`^('|")`).MatchString(classField.Default) {
				classField.Type = &cem.Type{
					Text: "string",
				}
			}
		}
	}

	for _, x := range captures["field.privacy"] {
		switch x.Text {
			case "private":
				classField.Privacy = cem.Private
			case "protected":
				classField.Privacy = cem.Protected
		}
	}

	ceField := cem.CustomElementField{
		ClassField: classField,
		Reflects: !classField.Static && len(captures["field.attr.reflects"]) > 0 && captures["field.attr.reflects"][0].Text == "true",
	}

	attributeBoolNodes := captures["field.attr.bool"]
	attributeNameNodes := captures["field.attr.name"]

	if (!classField.Static &&
			len(attributeBoolNodes) == 0 ||
			(len(attributeBoolNodes) > 0 && attributeBoolNodes[0].Node.GrammarName() != "false")) {
		if len(attributeNameNodes) > 0 {
			ceField.Attribute = attributeNameNodes[0].Text
		} else {
			ceField.Attribute = strings.ToLower(classField.Name)
		}
	}

	for k := range captures { delete(captures, k) }
	return &ceField
}

func getClassFieldsFromClassDeclarationNode(code []byte, node *ts.Node) []cem.CustomElementField {
	queryText, err := LoadQueryFile("classField")
	if err != nil {
		log.Fatal(err)
	}
	query, qerr := ts.NewQuery(Typescript, queryText)
	defer query.Close()
	if qerr != nil {
		log.Fatal(qerr)
	}
	cursor := ts.NewQueryCursor()
	defer cursor.Close()

	staticsSet := set.NewSet[string]()
	fieldsSet := set.NewSet[string]()
	gettersSet := set.NewSet[string]()
	settersSet := set.NewSet[string]()

	// to collate first instance of accessor pair / accessor
	accessorIndices := make(map[string]int)

	// final list of fields, in source order
	fields := make([]cem.CustomElementField, 0)

	for match := range AllQueryMatches(cursor, query, node, code) {
		captures := GetCapturesFromMatch(match, query, code)
		fieldName := captures["field.name"][0].Text
		_, isAccessor := captures["accessor"]
		_, isField := captures["field"]
		_, isStatic := captures["field.static"]

		if isStatic && ignoredStaticFields.Has(fieldName) {
			continue
		}

		if (isField && isStatic && !staticsSet.Has(fieldName)) {
			staticsSet.Add(fieldName)
			field := createClassFieldFromFieldMatch(fieldName, isStatic, captures)
			fields = append(fields, *field)
		} else if (isField && !isStatic && !fieldsSet.Has(fieldName)) {
			fieldsSet.Add(fieldName)
			field := createClassFieldFromFieldMatch(fieldName, isStatic, captures)
			fields = append(fields, *field)
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
				field := createClassFieldFromAccessorMatch(fieldName, accessor, isStatic, captures)
				// this next line fails at runtime with "dst must be a pointer"
				err := mergo.Merge(&fields[index], field, mergo.WithOverride)
				if err != nil {
					log.Fatal(err)
				}
				fields[index].Readonly = false
			} else {
				field := createClassFieldFromAccessorMatch(fieldName, accessor, isStatic, captures)
				fields = append(fields, *field)
			}
		}
	}

	return fields
}

