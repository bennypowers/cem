package generate

import (
	"errors"
	"slices"

	"bennypowers.dev/cem/manifest"

	A "github.com/IBM/fp-go/array"
	ts "github.com/tree-sitter/go-tree-sitter"
	tsts "github.com/tree-sitter/tree-sitter-typescript/bindings/go"
)

func generateModule(file string, code []byte) (err error, module manifest.Module) {
	language := ts.NewLanguage(tsts.LanguageTypescript())
	parser := ts.NewParser()
	defer parser.Close()
	parser.SetLanguage(language)
	tree := parser.Parse(code, nil)
	defer tree.Close()
	root := tree.RootNode()

	module = manifest.NewModule(file)

	qm, closeQm := NewQueryMatcher("customElementDeclaration", Languages.Typescript)
	defer closeQm()
	captureNames := qm.query.CaptureNames()

	for match := range qm.AllQueryMatches(root, code) {
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


			error, fields := getClassFieldsFromClassDeclarationNode(code, &declNode)
			if error != nil {
				err = errors.Join(err, error)
			}
			error, methods := getClassMethodsFromClassDeclarationNode(code, &declNode)
			if err != nil {
				err = errors.Join(err, error)
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

			error, classInfo := NewClassInfo(jsdoc)
			if error != nil {
				err = errors.Join(err, error)
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

			for _, field := range fields {
				declaration.Members = append(declaration.Members, field)
			}

			for _, method := range methods {
				declaration.Members = append(declaration.Members, method)
			}

			module.Declarations = append(module.Declarations, declaration)
		}
	}
	return err, module
}

