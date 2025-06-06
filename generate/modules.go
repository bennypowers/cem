package generate

import (
	"errors"

	"bennypowers.dev/cem/manifest"
	"bennypowers.dev/cem/set"

	ts "github.com/tree-sitter/go-tree-sitter"
	tsts "github.com/tree-sitter/tree-sitter-typescript/bindings/go"
)

func generateModule(file string, code []byte) (err error, module *manifest.Module) {
	language := ts.NewLanguage(tsts.LanguageTypescript())
	parser := ts.NewParser()
	defer parser.Close()
	parser.SetLanguage(language)
	tree := parser.Parse(code, nil)
	defer tree.Close()
	root := tree.RootNode()

	module = manifest.NewModule(file)

	classNamesAdded := set.NewSet[string]()

	queryName := "customElementDeclaration"
	qm, closeQm := NewQueryMatcher(queryName, Languages.Typescript)
	defer closeQm()
	for match := range qm.AllQueryMatches(root, code) {
		captures := qm.GetCapturesFromMatch(match, code)
		err, declaration := generateCustomElementDeclaration(captures, root, code)
		if err != nil {
			return err, nil
		} else {
			classNamesAdded.Add(declaration.Name)
			module.Declarations = append(module.Declarations, declaration)
			reference := manifest.Reference{ Name: declaration.Name, Module: file }
			module.Exports = append(module.Exports, &manifest.CustomElementExport{
				Kind: "custom-element-definition",
				Name: declaration.TagName,
				Declaration: &reference,
			})
		}
	}

	queryName = "variableDeclaration"
	qm, closeQm = NewQueryMatcher(queryName, Languages.Typescript)
	defer closeQm()
	for match := range qm.AllQueryMatches(root, code) {
		captures := qm.GetCapturesFromMatch(match, code)
		err, declaration := generateVarDeclaration(captures)
		if err != nil {
			return err, nil
		} else {
			module.Declarations = append(module.Declarations, declaration)
		}
	}

	queryName = "classDeclaration"
	qm, closeQm = NewQueryMatcher(queryName, Languages.Typescript)
	defer closeQm()
	for match := range qm.AllQueryMatches(root, code) {
		captures := qm.GetCapturesFromMatch(match, code)
		err, declaration := generateClassDeclaration(captures, root, code)
		if err != nil {
			return err, nil
		} else if !classNamesAdded.Has(declaration.Name) {
			classNamesAdded.Add(declaration.Name)
			module.Declarations = append(module.Declarations, declaration)
		}
	}

	queryName = "functionDeclaration"
	qm, closeQm = NewQueryMatcher(queryName, Languages.Typescript)
	defer closeQm()
	for match := range qm.AllQueryMatches(root, code) {
		captures := qm.GetCapturesFromMatch(match, code)
		err, declaration := generateFunctionDeclaration(captures, root, code)
		if err != nil {
			return err, nil
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
			return errors.Join(err, &NoCaptureError{ "export", queryName }), nil
		} else if declarationNameNodes, ok := captures["declaration.name"]; (!ok || len(declarationNameNodes) <= 0) {
			return errors.Join(err, &NoCaptureError{ "declaration.name", queryName }), nil
		} else {
			declaration := declarationNameNodes[0]
			module.Exports = append(module.Exports, &manifest.JavaScriptExport{
				Kind: "js",
				Name: declaration.Text,
				Declaration: &manifest.Reference{
					Name: declaration.Text,
					Module: file,
				},
			})
		}
	}


	return nil, module
}

