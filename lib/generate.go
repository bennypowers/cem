package generate

import (
	"embed"
	"fmt"
	"cmp"
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

	queryText, err := loadQueryFile("customElementDecorator")
	if err != nil {
		log.Fatal(nil)
	}
	query, qerr := ts.NewQuery(language, queryText)
	if qerr != nil {
		log.Fatal(qerr)
	}
	defer query.Close()

	captureNames := query.CaptureNames()

	cursor := ts.NewQueryCursor()
	defer cursor.Close()

	root := tree.RootNode()

	matches := cursor.Matches(query, root, code)

	module := cem.Module {
		Kind: "javascript-module",
		Path: file,
		Declarations: make([]cem.Declaration, 0),
		Exports: make([]cem.Export, 0),
	}

	for {
		match := matches.Next()
		if match == nil {
			break
		}
		var tagName, className string
		for _, capture := range match.Captures {
			name := captureNames[capture.Index]
			switch name {
			case "tag-name":
				tagName = capture.Node.Utf8Text(code)
			case "class-name":
				className = capture.Node.Utf8Text(code)
			}
		}
		if tagName != "" && className != "" {
			declaration := cem.Reference{ Name: className, Module: &file }
			module.Exports = append(module.Exports, &cem.CustomElementExport{
				Kind: "custom-element-definition",
				Name: tagName,
				Declaration: declaration,
			}, &cem.JavaScriptExport{
				Kind: "js",
				Name: className,
				Declaration: declaration,
			})
			module.Declarations = append(module.Declarations, &cem.CustomElementDeclaration{
				CustomElement: cem.CustomElement {
					CustomElement: true,
				},
				ClassDeclaration: cem.ClassDeclaration{
					Kind: "class",
					ClassLike: cem.ClassLike{
						Name: className,
						Members: []cem.ClassMember{},
					},
				},
			})
		}
	}
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

	manifest, err := cem.SerializeToString(pkg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(manifest)
}
