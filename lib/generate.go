package generate

import (
	"cmp"
	"embed"
	"fmt"
	"iter"
	"log"
	"os"
	"slices"
	"sync"

	"github.com/bennypowers/cemgen/cem"
	"github.com/bennypowers/cemgen/set"

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
) []cem.ClassField {
	fields := make([]cem.ClassField, 0)
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
		jsdocCaptureIndex, _ := query.CaptureIndexForName("jsdoc")
		fieldNameCaptureIndex, _ := query.CaptureIndexForName("field.name")
		fieldTypeCaptureIndex, _ := query.CaptureIndexForName("field.type")
		fieldInitializerCaptureIndex, _ := query.CaptureIndexForName("field.initializer")
		nameNodes := match.NodesForCaptureIndex(fieldNameCaptureIndex)
		typeNodes := match.NodesForCaptureIndex(fieldTypeCaptureIndex)
		initializerNodes := match.NodesForCaptureIndex(fieldInitializerCaptureIndex)
		jsdocNodes := match.NodesForCaptureIndex(jsdocCaptureIndex)
		for _, node := range nameNodes {
			name := node.Utf8Text(code)
			if name != "" && !fieldsSet[name] {
				fieldsSet[name] = true
				var defaultValue, typeText string
				var jsdocInfo FieldInfo
				var deprecated cem.Deprecated
				for _, node := range typeNodes { typeText = node.Utf8Text(code) }
				for _, node := range initializerNodes { defaultValue = node.Utf8Text(code) }
				for _, jsdocNode := range jsdocNodes {
					jsdocInfo = getFieldInfoFromJsdoc(jsdocNode.Utf8Text(code))
				}
				fields = append(fields, cem.ClassField{
					Kind: "field",
					PropertyLike: cem.PropertyLike{
						Name: name,
						Description: jsdocInfo.Description,
						Default: defaultValue,
						Summary: jsdocInfo.Summary,
						Deprecated: deprecated,
						Type: &cem.Type{
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
						},
					},
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

			// todo: combine jsdoc and field attrs

			classInfo := NewClassInfo(jsdoc)

			declaration := &cem.CustomElementDeclaration{
				CustomElement: cem.CustomElement {
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
func Generate(files []string, exclude []string) {
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
	fmt.Println(manifest)
}
