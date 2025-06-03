package generate

import (
	"log"
	"strings"

	"github.com/bennypowers/cemgen/cem"
	ts "github.com/tree-sitter/go-tree-sitter"
)

func getClassFieldsFromClassDeclarationNode(
	language *ts.Language,
	code []byte,
	node *ts.Node,
) []cem.CustomElementField {
	fields := make([]cem.CustomElementField, 0)
	queryText, err := loadQueryFile("customElementClassField")
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
	fieldsSet := make(map[string]bool)
	for match := range allMatches(cursor, query, node, code) {
		fieldCaptureIndex, _ := query.CaptureIndexForName("field")
		jsdocCaptureIndex, _ := query.CaptureIndexForName("field.jsdoc")
		attributeBoolCaptureIndex, _ := query.CaptureIndexForName("field.attr.bool")
		attributeNameCaptureIndex, _ := query.CaptureIndexForName("field.attr.name")
		reflectsCaptureIndex, _ := query.CaptureIndexForName("field.attr.reflects")
		fieldNameCaptureIndex, _ := query.CaptureIndexForName("field.name")
		fieldTypeCaptureIndex, _ := query.CaptureIndexForName("field.type")
		fieldInitializerCaptureIndex, _ := query.CaptureIndexForName("field.initializer")

		jsdocNodes := match.NodesForCaptureIndex(jsdocCaptureIndex)
		fieldNodes := match.NodesForCaptureIndex(fieldCaptureIndex)
		attributeBoolNodes := match.NodesForCaptureIndex(attributeBoolCaptureIndex)
		attributeNameNodes := match.NodesForCaptureIndex(attributeNameCaptureIndex)
		reflectsNodes := match.NodesForCaptureIndex(reflectsCaptureIndex)
		nameNodes := match.NodesForCaptureIndex(fieldNameCaptureIndex)
		typeNodes := match.NodesForCaptureIndex(fieldTypeCaptureIndex)
		initializerNodes := match.NodesForCaptureIndex(fieldInitializerCaptureIndex)

		for _, node := range nameNodes {
			name := node.Utf8Text(code)
			if name != "" && !fieldsSet[name] {
				fieldsSet[name] = true
				var static bool
				var readonly bool
				var privacy cem.Privacy
				var defaultValue, typeText string
				var jsdocInfo FieldInfo
				var deprecated cem.Deprecated
				for _, node := range typeNodes { typeText = node.Utf8Text(code) }
				for _, node := range initializerNodes { defaultValue = node.Utf8Text(code) }
				for _, node := range jsdocNodes { jsdocInfo = NewFieldInfo(node.Utf8Text(code)) }
				for _, node := range fieldNodes {
					for _, node := range node.Children(node.Walk()) {
						text := node.Utf8Text(code)
						switch text {
							case "static":
								static = true
							case "readonly":
								readonly = true
							case "private":
								privacy = cem.Private
							case "protected":
								privacy = cem.Protected
						}
					}
				}

				fieldType := &cem.Type{
					Text: (func() (string) {
						if jsdocInfo.Type != "" {
							return jsdocInfo.Type
						} else if typeText != "" {
							return typeText
						} else if defaultValue == "true" || defaultValue == "false" {
							return "boolean"
						} else {
							return ""
						}
					})(),
				}

				classField := cem.ClassField{
					Kind: "field",
					Static: static,
					Privacy: privacy,
					PropertyLike: cem.PropertyLike{
						Name: name,
						Description: jsdocInfo.Description,
						Default: defaultValue,
						Summary: jsdocInfo.Summary,
						Deprecated: deprecated,
						Type: fieldType,
						Readonly: readonly,
					},
				}

					var attribute string
					reflects := len(reflectsNodes) > 0 && reflectsNodes[0].Utf8Text(code) == "true"
				  attrFalse := len(attributeBoolNodes) > 0 && attributeBoolNodes[0].GrammarName() == "false"
					if !attrFalse {
						if len(attributeNameNodes) > 0 {
							attribute = attributeNameNodes[0].Utf8Text(code)
						} else {
							attribute = strings.ToLower(classField.Name)
						}
					}

					fields = append(fields, cem.CustomElementField{
						ClassField: classField,
						Attribute: attribute,
						Reflects: reflects,
					})
			}
		}
	}
	return fields
}

