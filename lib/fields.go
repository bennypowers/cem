package generate

import (
	"log"
	"regexp"
	"strings"

	"github.com/bennypowers/cemgen/cem"
	ts "github.com/tree-sitter/go-tree-sitter"
)

func handleMatch(
	match *ts.QueryMatch,
	query *ts.Query,
	code []byte,
	statics map[string]*cem.CustomElementField,
	fields map[string]*cem.CustomElementField,
) (*cem.CustomElementField, bool) {
	captures := *getCapturesFromMatch(match, query)

	fieldName := captures["field.name"][0].Utf8Text(code)

	_, static := captures["field.static"]

	if static {
		switch fieldName {
		case "formAssociated", "styles", "shadowRootOptions":
			return nil, false
		}
		if (statics[fieldName] != nil) {
			return nil, false
		}
	}

	if (fields[fieldName] != nil) {
		return nil, false
	}

	_, readonly := captures["field.readonly"]

	classField := cem.ClassField{
		Kind: "field",
		Static: static,
		PropertyLike: cem.PropertyLike{
			Name: fieldName,
			Readonly: readonly,
		},
	}

	for _, node := range captures["field.type"] {
		typeText := node.Utf8Text(code)
		if typeText != "" {
			classField.Type = &cem.Type{
				Text: typeText,
			}
		}
	}

	for _, node := range captures["field.jsdoc"] {
		source := node.Utf8Text(code)
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

	for _, node := range captures["field.initializer"] {
		classField.Default = node.Utf8Text(code)
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

	for _, node := range captures["field.privacy"] {
		switch node.Utf8Text(code) {
			case "private":
				classField.Privacy = cem.Private
			case "protected":
				classField.Privacy = cem.Protected
		}
	}

	// for _, node := captures["decorator"] {
	// 	reflects, attrTrue, attrName = parseDecorator(node.Utf8Text(code))
	// }

	ceField := cem.CustomElementField{
		ClassField: classField,
		Reflects: !classField.Static && len(captures["field.attr.reflects"]) > 0 && captures["field.attr.reflects"][0].Utf8Text(code) == "true",
	}

	attributeBoolNodes := captures["field.attr.bool"]
	attributeNameNodes := captures["field.attr.name"]

	if (!classField.Static &&
			len(attributeBoolNodes) == 0 ||
			(len(attributeBoolNodes) > 0 && attributeBoolNodes[0].GrammarName() != "false")) {
		if len(attributeNameNodes) > 0 {
			ceField.Attribute = attributeNameNodes[0].Utf8Text(code)
		} else {
			ceField.Attribute = strings.ToLower(classField.Name)
		}
	}

	for k := range captures { delete(captures, k) }
	return &ceField, true
}

func getClassFieldsFromClassDeclarationNode(
	language *ts.Language,
	code []byte,
	node *ts.Node,
) []cem.CustomElementField {
	queryText, err := loadQueryFile("classField")
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

	statics := make(map[string]*cem.CustomElementField)
	staticOrder := make([]string, 0)
	fields := make(map[string]*cem.CustomElementField)
	fieldOrder := make([]string, 0)

	for match := range allMatches(cursor, query, node, code) {
		field, ok := handleMatch(match, query, code, statics, fields)
		if ok {
			if field.Static {
				statics[field.Name] = field
				staticOrder = append(staticOrder, field.Name)
			} else {
				fields[field.Name] = field
				fieldOrder = append(fieldOrder, field.Name)
			}
		}
	}

	orderedFields := make([]cem.CustomElementField, 0)

	for _, name := range staticOrder {
		orderedFields = append(orderedFields, *statics[name])
	}
	for _, name := range fieldOrder {
		orderedFields = append(orderedFields, *fields[name])
	}
	return orderedFields
}

