package generate

import (
	"errors"
	"fmt"
	"maps"
	"os"
	"path"
	"slices"
	"strings"

	M "bennypowers.dev/cem/manifest"
	S "bennypowers.dev/cem/set"
	ts "github.com/tree-sitter/go-tree-sitter"
)

func ammendStylesMapFromSource(
	props map[string]M.CssCustomProperty,
	queryManager *QueryManager,
	queryMatcher *QueryMatcher,
	parser *ts.Parser,
	code []byte,
) (errs error) {
	tree := parser.Parse(code, nil)
	defer tree.Close()
	root := tree.RootNode()
	for captures := range queryMatcher.ParentCaptures(root, code, "cssPropertyCallSite") {
		properties := captures["property"]
		name := properties[len(properties)-1].Text
		startByte := properties[len(properties)-1].StartByte
		p := props[name]
		p.Name = name
		p.StartByte = startByte
		_, has := props[name]
		if !has {
			props[name] = M.CssCustomProperty{
				Name: name,
				StartByte: startByte,
			}
		}
		defaultVals, ok := captures["default"]
		if ok && len(defaultVals) > 0{
			startNode := GetDescendantById(root, defaultVals[0].NodeId)
			endNode := GetDescendantById(root, defaultVals[len(defaultVals) - 1].NodeId)
			p.Default = string(code[startNode.StartByte():endNode.EndByte()])
		}
		comment, ok := captures["comment"]
		if ok {
			for _, comment := range comment {
				info, err := NewCssCustomPropertyInfo(comment.Text, queryManager)
				if err != nil {
					errs = errors.Join(errs, err)
				} else {
					info.MergeToCssCustomProperty(&p)
				}
			}
		}
		props[name] = p
	}
	return errs
}

func generateModule(file string, code []byte, queryManager *QueryManager) (errs error, module *M.Module) {
	parser := ts.NewParser()
	defer parser.Close()
	parser.SetLanguage(Languages.typescript)
	tree := parser.Parse(code, nil)
	defer tree.Close()
	root := tree.RootNode()

	module = M.NewModule(file)

	classNamesAdded := S.NewSet[string]()

	queryName := "variableDeclaration"
	qm, err := NewQueryMatcher(queryManager, "typescript", queryName)
	if err != nil {
		return errors.Join(errs, err), nil
	}
	defer qm.Close()
	for captures := range qm.ParentCaptures(root, code, "variable") {
		err, declaration := generateVarDeclaration(captures, queryManager)
		if err != nil {
			errs = errors.Join(errs, err)
		} else {
			module.Declarations = append(module.Declarations, declaration)
		}
	}

	queryName = "classDeclaration"
	qm, err = NewQueryMatcher(queryManager, "typescript", queryName)
	if err != nil {
		return errors.Join(errs, err), nil
	}
	defer qm.Close()
	styleImportsBindingToSpecMap := make(map[string]string)
	for sImport := range qm.ParentCaptures(root, code, "styleImport.spec") {
		styleImportsBindingToSpecMap[sImport["styleImport.binding"][0].Text] = sImport["styleImport.spec"][0].Text
	}
	for captures := range qm.ParentCaptures(root, code, "class") {
		d, err := generateClassDeclaration(queryManager, captures, root, code)
		if err != nil {
			errs = errors.Join(errs, err)
		} else {
			declaration, ok := d.(*M.ClassDeclaration)
			if ok && !classNamesAdded.Has(declaration.Name) {
				classNamesAdded.Add(declaration.Name)
				module.Declarations = append(module.Declarations, declaration)
			} else {
				declaration, ok := d.(*M.CustomElementDeclaration)
				props := make(map[string]M.CssCustomProperty)
				if ok && !classNamesAdded.Has(declaration.Name) {
					classNamesAdded.Add(declaration.Name)
					// get all style array or value bindings
					bindings, hasBindings := captures["style.binding"]
					styleStrings, hasStrings := captures["style.string"]
					if hasBindings || hasStrings {
						qm, err := NewQueryMatcher(queryManager, "css", "cssCustomProperties")
						parser := ts.NewParser()
						defer parser.Close()
						parser.SetLanguage(Languages.css)
						if err != nil {
							return errors.Join(errs, err), nil
						}
						if hasBindings {
							for _, binding := range bindings {
								spec, ok := styleImportsBindingToSpecMap[binding.Text]
								if ok && strings.HasPrefix(spec, "."){
									// i'm proposing to do file io here, read the referenced spec file,
									// parse it, and add any variables i find to the cssProperties slice.
									// perhaps a more performant or maintainable alternative is to leave a marker here,
									// and then do all that work on the back end.
									content, err := os.ReadFile(path.Join(path.Dir(module.Path), spec))
									if err != nil {
										errs = errors.Join(errs, errors.New(fmt.Sprintf("Could not read tokens spec %s", spec)), err)
									} else {

								err := ammendStylesMapFromSource(props, queryManager, qm, parser, content)
								if err != nil {
									errs = errors.Join(errs, err)
								}
									}
								}
								// later on, i could add an else path here which checks to see if
								// the css declaration is made elsewhere in the same module, then parse it's string content
								// if the value for that variable declaration is a css template tag
							}
						}
						if hasStrings {
							for _, styleString := range styleStrings {
								err := ammendStylesMapFromSource(props, queryManager, qm, parser, []byte(styleString.Text))
								if err != nil {
									errs = errors.Join(errs, err)
								}
							}
						}
					}
					cssProps := slices.Collect(maps.Values(props))
					declaration.CssProperties = append(declaration.CssProperties, cssProps...)
					slices.SortStableFunc(declaration.CssProperties, func(a M.CssCustomProperty, b M.CssCustomProperty) int {
						if a.StartByte == b.StartByte {
							return strings.Compare(a.Name, b.Name)
						}
						return int(a.StartByte) - int(b.StartByte)
					})
					// TODO: add css parts by parsing template
					// TODO: add slots by parsing template. How do do description, etc?
					module.Declarations = append(module.Declarations, declaration)
				}
			}
		}
	}

	queryName = "functionDeclaration"
	qm, err = NewQueryMatcher(queryManager, "typescript", queryName)
	if err != nil {
		return errors.Join(errs, err), nil
	}
	defer qm.Close()
	for captures := range qm.ParentCaptures(root, code, "function") {
		err, declaration := generateFunctionDeclaration(captures, root, code, queryManager)
		if err != nil {
			errs = errors.Join(errs, err)
		} else {
			module.Declarations = append(module.Declarations, declaration)
		}
	}

	queryName = "exportStatement"
	qm, err = NewQueryMatcher(queryManager, "typescript", queryName)
	if err != nil {
		return errors.Join(errs, err), nil
	}
	defer qm.Close()
	for captures := range qm.ParentCaptures(root, code, "export") {
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
				StartByte: declaration.StartByte,
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
					StartByte: declaration.StartByte,
				})
			}
		}
	}

	slices.SortStableFunc(module.Exports, func(a M.Export, b M.Export) int {
		var c uint
		var d uint
		if A, ok := a.(*M.CustomElementExport); ok {
			c = A.StartByte
		}
		if A, ok := a.(*M.JavaScriptExport); ok {
			c = A.StartByte
		}
		if B, ok := b.(*M.CustomElementExport); ok {
			d = B.StartByte
		}
		if B, ok := b.(*M.JavaScriptExport); ok {
			d = B.StartByte
		}
		return int(c) - int(d)
	})
	return nil, module
}

