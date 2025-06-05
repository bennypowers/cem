package generate

import (
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
	"sync"

	"bennypowers.dev/cem/manifest"
	"bennypowers.dev/cem/set"

	A "github.com/IBM/fp-go/array"
	ts "github.com/tree-sitter/go-tree-sitter"
	tsts "github.com/tree-sitter/tree-sitter-typescript/bindings/go"
)

func generateModule(file string, channel chan<-manifest.Module, wg *sync.WaitGroup) {
	defer wg.Done()

	code, err := os.ReadFile(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		return
	}

	language := ts.NewLanguage(tsts.LanguageTypescript())
	parser := ts.NewParser()
	defer parser.Close()
	parser.SetLanguage(language)
	tree := parser.Parse(code, nil)
	defer tree.Close()
	root := tree.RootNode()

	module := manifest.NewModule(file)

	queryText, err := LoadQueryFile("customElementDeclaration")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
	}
	query, qerr := ts.NewQuery(language, queryText)
	defer query.Close()
	if qerr != nil {
		log.Fatal(qerr)
	}
	captureNames := query.CaptureNames()
	cursor := ts.NewQueryCursor()

	for match := range AllQueryMatches(cursor, query, root, code) {
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
			reference := manifest.Reference{ Name: className, Module: file }
			module.Exports = append(module.Exports, &manifest.CustomElementExport{
				Kind: "custom-element-definition",
				Name: tagName,
				Declaration: &reference,
			}, &manifest.JavaScriptExport{
				Kind: "js",
				Name: className,
				Declaration: &reference,
			})


			err, fields := getClassFieldsFromClassDeclarationNode(code, &declNode)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s", err)
			}
			err, methods := getClassMethodsFromClassDeclarationNode(code, &declNode)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s", err)
			}

			fieldAttrs := A.Chain(func(field manifest.CustomElementField) ([]manifest.Attribute) {
				if field.Attribute == "" {
					return []manifest.Attribute{}
				} else {
					return []manifest.Attribute{{
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

			err, classInfo := NewClassInfo(jsdoc)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s", err)
			}

			declaration := &manifest.CustomElementDeclaration{
				CustomElement: manifest.CustomElement {
					CustomElement: true,
					TagName: tagName,
					Attributes: slices.Concat(classInfo.Attrs, fieldAttrs),
					Slots: classInfo.Slots,
					Events: classInfo.Events,
					CssProperties: classInfo.CssProperties,
					CssParts: classInfo.CssParts,
					CssStates: classInfo.CssStates,
				},
				ClassDeclaration: manifest.ClassDeclaration{
					Kind: "class",
					ClassLike: manifest.ClassLike{
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

	modulesChan := make(chan manifest.Module, len(files))

	for _, file := range files {
		if !excludeSet.Has(file) {
			wg.Add(1)
			go generateModule(file, modulesChan, &wg)
		}
	}

	wg.Wait()
	close(modulesChan)

	modules := make([]manifest.Module, 0)
	for module := range modulesChan {
		modules = append(modules, module)
	}

	pkg := manifest.NewPackage(modules)

	slices.SortFunc(pkg.Modules, func(a, b manifest.Module) int {
		return cmp.Compare(a.Path, b.Path)
	})

	manifest, err := manifest.SerializeToString(&pkg)

	if err != nil {
		log.Fatal(err)
	}

	return manifest
}
