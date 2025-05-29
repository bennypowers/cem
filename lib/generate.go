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

	exportCaptureIndex, ok := query.CaptureIndexForName("export")
	if !ok {
		log.Fatal("could not find export capture in query")
	}

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

	for {
		match, idx := captures.Next()
		if match == nil {
			break
		}
		if idx != exportCaptureIndex {
			continue
		}
		for _, exportNode := range match.NodesForCaptureIndex(exportCaptureIndex) {
			nodeCursor := exportNode.Walk()
			decoratorNodes := exportNode.ChildrenByFieldName("decorator", nodeCursor)
			decoratorNodeIndex := slices.IndexFunc(decoratorNodes, func(n ts.Node) bool {
				callExpressionNode := n.NamedChild(0)
				if callExpressionNode.Kind() != "call_expression" { return false }
				functionName := callExpressionNode.ChildByFieldName("function").Utf8Text(code)
				return functionName == "customElement"
			})
			if decoratorNodeIndex < 0 {
				continue
			}
			decoratorNode := decoratorNodes[decoratorNodeIndex]
			decoratorCallExpressionNode := decoratorNode.NamedChild(0)
			decoratorArgumentsStringStringFragmentNode :=
				decoratorCallExpressionNode.NamedChild(1).NamedChild(0).NamedChild(0)
			classDeclarationNode := exportNode.ChildByFieldName("declaration")
			tagName := decoratorArgumentsStringStringFragmentNode.Utf8Text(code)
			className := classDeclarationNode.ChildByFieldName("name").Utf8Text(code)
			declaration := cem.Reference{ Name: className, Module: &file, }
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

	modules := make(chan cem.Module, len(files))

	for _, file := range files {
		if !excludeSet.Has(file) {
			wg.Add(1)
			go generateModule(file, modules, &wg)
		}
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

	slices.SortFunc(pkg.Modules, func(a, b cem.Module) int {
		return cmp.Compare(a.Path, b.Path)
	})

	manifest, err := cem.SerializeToString(&pkg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(manifest)
}
