package generate

import (
	"fmt"
	"log"
	"os"
	"embed"
	"sync"

	"github.com/bennypowers/cemgen/cem"

	ts "github.com/tree-sitter/go-tree-sitter"
	tsts "github.com/tree-sitter/tree-sitter-typescript/bindings/go"
)

//go:embed queries/*.scm
var queries embed.FS

func LoadQueryFile(queryName string) (string, error) {
	path := fmt.Sprintf("queries/%s.scm", queryName)
	data, err := queries.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}


func generate(file string, channel chan<- cem.Module, wg *sync.WaitGroup) {
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

	queryText, err := LoadQueryFile("customElementDecorator")
	if err != nil {
		log.Fatal(nil)
	}
	query, qerr := ts.NewQuery(language, queryText)
	if qerr != nil {
		log.Fatal(qerr)
	}
	defer query.Close()

	cursor := ts.NewQueryCursor()
	defer cursor.Close()

	root := tree.RootNode()

	captures := cursor.Captures(query, root, code)

	module := cem.Module {
		Kind: "javascript-module",
		Path: file,
		Declarations: make([]cem.Declaration, 0),
		Exports: make([]cem.Export, 0),
	}

	for match, index := captures.Next(); match != nil; match, index = captures.Next() {
		cap := match.Captures[index]
		text := cap.Node.Utf8Text(code)
		switch index {
			case 0:
				cex := &cem.CustomElementExport{
					Kind: "custom-element-definition",
				}
				module.Exports = append(module.Exports, cex)
			case 1:
				module.Exports[len(module.Exports)-1].(*cem.CustomElementExport).Name = text
			case 2:
				module.Exports[len(module.Exports)-1].(*cem.CustomElementExport).Declaration = cem.Reference{
				Name: text,
			}
		}
	}

	channel <- module
}

func Generate(files []string) {
	modules := make(chan cem.Module, len(files))
	var wg sync.WaitGroup

	for _, file := range files {
		wg.Add(1)
		go generate(file, modules, &wg)
	}

	wg.Wait()
	close(modules)

	pkg := cem.Package{
		SchemaVersion: "1.0.0",
		Modules: make([]cem.Module, 0),
	}

	for module := range modules {
		pkg.Modules = append(pkg.Modules, module)
	}

	manifest, err := cem.SerializeToString(&pkg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(manifest)
}
