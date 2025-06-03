package generate

import (
	"cmp"
	"embed"
	"fmt"
	"iter"
	"log"
	"os"
	"slices"
	"strings"
	"sync"

	"github.com/bennypowers/cemgen/cem"
	"github.com/bennypowers/cemgen/set"

	A "github.com/IBM/fp-go/array"
	ts "github.com/tree-sitter/go-tree-sitter"
	tsts "github.com/tree-sitter/tree-sitter-typescript/bindings/go"
)

//go:embed queries/*.scm
var queries embed.FS

func loadQueryFile(queryName string) (string, error) {
	path := fmt.Sprintf("queries/%s.scm", queryName)
	data, err := queries.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func allMatches(qc *ts.QueryCursor, q *ts.Query, node *ts.Node, text []byte) iter.Seq2[*ts.QueryMatch, *ts.Query] {
	qm := qc.Matches(q, node, text)
	return func(yield func (qm *ts.QueryMatch, q *ts.Query) bool) {
		for {
			m := qm.Next()
			if m == nil {
				break
			}
			if !yield(m, q) {
				return
			}
		}
	}
}

func getClassMethodsFromClassDeclarationNode(
	language *ts.Language,
	code []byte,
	node *ts.Node,
) []cem.ClassMethod {
	methods := make([]cem.ClassMethod, 0)
	queryText, err := loadQueryFile("customElementClassMethod")
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
	for match := range allMatches(cursor, query, node, code) {
		var methodName string
		for _, capture := range match.Captures {
			name := query.CaptureNames()[capture.Index]
			switch name {
			case "method.name":
				methodName = capture.Node.Utf8Text(code)
			}
		}
		if methodName != "" {
			methods = append(methods, cem.ClassMethod{
				Kind: "method",
				FunctionLike: cem.FunctionLike{
					Name: methodName,
				},
			})
		}
	}
	return methods
}

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
		attributeCaptureIndex, _ := query.CaptureIndexForName("field.attr")
		attributeNameCaptureIndex, _ := query.CaptureIndexForName("field.attr.name")
		reflectsCaptureIndex, _ := query.CaptureIndexForName("field.attr.reflects")
		fieldNameCaptureIndex, _ := query.CaptureIndexForName("field.name")
		fieldTypeCaptureIndex, _ := query.CaptureIndexForName("field.type")
		fieldInitializerCaptureIndex, _ := query.CaptureIndexForName("field.initializer")

		jsdocNodes := match.NodesForCaptureIndex(jsdocCaptureIndex)
		fieldNodes := match.NodesForCaptureIndex(fieldCaptureIndex)
		attributeNodes := match.NodesForCaptureIndex(attributeCaptureIndex)
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
					reflects := len(reflectsNodes) > 0 && reflectsNodes[0].Utf8Text(code) != "true"
				  attrFalse := len(attributeNodes) > 0 && attributeNodes[0].GrammarName() == "false"
					if !attrFalse && len(attributeNameNodes) > 0 {
						attribute = attributeNameNodes[0].Utf8Text(code)
					}
				  if reflects && attribute == "" {
						attribute = strings.ToLower(classField.Name)
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

func generateModule(file string, channel chan<- cem.Module, wg *sync.WaitGroup) {
	defer wg.Done()

	code, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	language := ts.NewLanguage(tsts.LanguageTypescript())
	parser := ts.NewParser()
	defer parser.Close()
	parser.SetLanguage(language)
	tree := parser.Parse(code, nil)
	defer tree.Close()
	root := tree.RootNode()

	module := cem.NewModule(file)

	queryText, err := loadQueryFile("customElementDeclaration")
	if err != nil {
		log.Fatal(err)
	}
	query, qerr := ts.NewQuery(language, queryText)
	defer query.Close()
	captureNames := query.CaptureNames()
	if qerr != nil {
		log.Fatal(qerr)
	}
	cursor := ts.NewQueryCursor()

	for match := range allMatches(cursor, query, root, code) {
		var declNode ts.Node
		var tagName, className, jsdoc string

		for _, capture := range match.Captures {
			name := captureNames[capture.Index]
			switch name {
				case "class.declaration":
					declNode = capture.Node
				case "jsdoc":
					jsdoc = capture.Node.Utf8Text(code)
				case "tag-name":
					tagName = capture.Node.Utf8Text(code)
				case "class.name":
					className = capture.Node.Utf8Text(code)
			}
		}

		if tagName != "" && className != "" {
			reference := cem.Reference{ Name: className, Module: file }
			module.Exports = append(module.Exports, &cem.CustomElementExport{
				Kind: "custom-element-definition",
				Name: tagName,
				Declaration: &reference,
			}, &cem.JavaScriptExport{
				Kind: "js",
				Name: className,
				Declaration: &reference,
			})


			fields := getClassFieldsFromClassDeclarationNode(language, code, &declNode)
			methods := getClassMethodsFromClassDeclarationNode(language, code, &declNode)

			fieldAttrs := A.Chain(func(field cem.CustomElementField) ([]cem.Attribute) {
				if field.Attribute == "" {
					return []cem.Attribute{}
				} else {
					return []cem.Attribute{{
						Name: field.Attribute,
						Summary: field.Summary,
						Description: field.Description,
						Deprecated: field.Deprecated,
						Default: field.Default,
						Type: field.Type,
						FieldName: field.Name,
					}}
				}
			})(fields)

			classInfo := NewClassInfo(jsdoc)

			declaration := &cem.CustomElementDeclaration{
				CustomElement: cem.CustomElement {
					Attributes: slices.Concat(classInfo.Attrs, fieldAttrs),
					CustomElement: true,
					TagName: tagName,
					Slots: classInfo.Slots,
					Events: classInfo.Events,
					CssProperties: classInfo.CssProperties,
					CssParts: classInfo.CssParts,
				},
				ClassDeclaration: cem.ClassDeclaration{
					Kind: "class",
					ClassLike: cem.ClassLike{
						Name: className,
						Deprecated: classInfo.Deprecated,
						Description: classInfo.Description,
						Summary: classInfo.Summary,
					},
				},
			}

			for _, field := range fields { declaration.Members = append(declaration.Members, field) }
			for _, method := range methods { declaration.Members = append(declaration.Members, method) }

			module.Declarations = append(module.Declarations, declaration)
		}
	}

	cursor.Close()
	channel <- module
}

// Generates a custom-elements manifest from a list of typescript files
func Generate(files []string, exclude []string) (string) {
	excludeSet := set.NewSet(exclude...)

	var wg sync.WaitGroup

	modulesChan := make(chan cem.Module, len(files))

	for _, file := range files {
		if !excludeSet.Has(file) {
			wg.Add(1)
			go generateModule(file, modulesChan, &wg)
		}
	}

	wg.Wait()
	close(modulesChan)

	modules := make([]cem.Module, 0)
	for module := range modulesChan { modules = append(modules, module) }

	pkg := cem.NewPackage(modules)

	slices.SortFunc(pkg.Modules, func(a, b cem.Module) int {
		return cmp.Compare(a.Path, b.Path)
	})

	manifest, err := cem.SerializeToString(&pkg)
	if err != nil {
		log.Fatal(err)
	}
	return manifest
}
