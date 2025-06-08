package generate

import (
	"errors"
	"slices"

	M "bennypowers.dev/cem/manifest"
	S "bennypowers.dev/cem/set"
	ts "github.com/tree-sitter/go-tree-sitter"
	tsts "github.com/tree-sitter/tree-sitter-typescript/bindings/go"
)

func generateModule(file string, code []byte) (errs error, module *M.Module) {
	language := ts.NewLanguage(tsts.LanguageTypescript())
	parser := ts.NewParser()
	defer parser.Close()
	parser.SetLanguage(language)
	tree := parser.Parse(code, nil)
	defer tree.Close()
	root := tree.RootNode()

	module = M.NewModule(file)

	classNamesAdded := S.NewSet[string]()

	queryName := "variableDeclaration"
	qm, closeQm := NewQueryMatcher(queryName, Languages.Typescript)
	defer closeQm()
	for match := range qm.AllQueryMatches(root, code) {
		captures := qm.GetCapturesFromMatch(match, code)
		err, declaration := generateVarDeclaration(captures)
		if err != nil {
			errs = errors.Join(errs, err)
		} else {
			module.Declarations = append(module.Declarations, declaration)
		}
	}

	queryName = "classDeclaration"
	qm, closeQm = NewQueryMatcher(queryName, Languages.Typescript)
	defer closeQm()
	for match := range qm.AllQueryMatches(root, code) {
		captures := qm.GetCapturesFromMatch(match, code)
		err, d := generateClassDeclaration(captures, root, code)
		if err != nil {
			errs = errors.Join(errs, err)
		} else {
			declaration, ok := d.(*M.ClassDeclaration)
			if ok && !classNamesAdded.Has(declaration.Name) {
				classNamesAdded.Add(declaration.Name)
				module.Declarations = append(module.Declarations, declaration)
			} else {
				declaration, ok := d.(*M.CustomElementDeclaration)
				if ok && !classNamesAdded.Has(declaration.Name) {
					classNamesAdded.Add(declaration.Name)
					module.Declarations = append(module.Declarations, declaration)
				}
			}
		}
	}

	queryName = "functionDeclaration"
	qm, closeQm = NewQueryMatcher(queryName, Languages.Typescript)
	defer closeQm()
	for match := range qm.AllQueryMatches(root, code) {
		captures := qm.GetCapturesFromMatch(match, code)
		err, declaration := generateFunctionDeclaration(captures, root, code)
		if err != nil {
			errs = errors.Join(errs, err)
		} else {
			module.Declarations = append(module.Declarations, declaration)
		}
	}

	queryName = "exportStatement"
	qm, closeQm = NewQueryMatcher(queryName, Languages.Typescript)
	defer closeQm()
	for match := range qm.AllQueryMatches(root, code) {
		captures := qm.GetCapturesFromMatch(match, code)
		if exportNodes, ok := captures["export"]; (!ok || len(exportNodes) <= 0) {
			errs = errors.Join(errs, &NoCaptureError{ "export", queryName })
		} else if declarationNameNodes, ok := captures["declaration.name"]; (!ok || len(declarationNameNodes) <= 0) {
			errs = errors.Join(errs, &NoCaptureError{ "declaration.name", queryName })
		} else {
			declaration := declarationNameNodes[0]
			export := &M.JavaScriptExport{
				Kind: "js",
				Name: declaration.Text,
				Declaration: M.NewReference(declaration.Text, "", file),
			}
			module.Exports = append(module.Exports, export)
		}
	}

	for name := range classNamesAdded {
		reference := M.NewReference(name, "", file)
		index := slices.IndexFunc(module.Declarations, func(d M.Declaration) bool {
			if ce, ok := d.(*M.CustomElementDeclaration); ok {
				return ce.Name == name
			} else {
				return false
			}
		})
		if index >= 0 {
			d := module.Declarations[index]
			if declaration, ok := d.(*M.CustomElementDeclaration); ok {
				module.Exports = append(module.Exports, &M.CustomElementExport{
					Kind: "custom-element-definition",
					Name: declaration.TagName,
					Declaration: reference,
				})
			}
		}
	}

	return nil, module
}

