package generate

import (
	"errors"

	"bennypowers.dev/cem/manifest"
	A "github.com/IBM/fp-go/array"
	ts "github.com/tree-sitter/go-tree-sitter"
)

func generateCustomElementDeclaration(
	captures CaptureMap,
	root *ts.Node,
	code []byte,
) (err error, declaration *manifest.CustomElementDeclaration) {
	declaration = &manifest.CustomElementDeclaration{
		ClassDeclaration: manifest.ClassDeclaration{
			Kind: "class",
		},
	}

	tagNameNodes, ok := captures["tag-name"]
	if (!ok || len(tagNameNodes) <= 0) {
		return &NoCaptureError{ "tag-name", "customElementDeclaration"  }, nil
	}
	tagName := tagNameNodes[0].Text

	classNameNodes, ok := captures["class.name"]
	if (!ok || len(classNameNodes) <= 0) {
		return &NoCaptureError{ "class.name", "customElementDeclaration" }, nil
	}
	if (!ok || len(classNameNodes) <= 0) {
		return errors.Join(err, &NoCaptureError{ "class.name", "customElementDeclaration" }), nil
	}
	className := classNameNodes[0].Text

	declaration.ClassLike.Name = className

	error, fields := getClassFieldsFromClassDeclarationNode(code, root)
	if error != nil {
		return errors.Join(err, error), nil
	}
	error, methods := getClassMethodsFromClassDeclarationNode(code, root)
	if error != nil {
		return errors.Join(err, error), nil
	}

	for _, field := range fields {
		declaration.Members = append(declaration.Members, field)
	}

	for _, method := range methods {
		declaration.Members = append(declaration.Members, method)
	}

	if (tagName != "") {
		declaration.CustomElement = manifest.CustomElement{
			CustomElement: true,
			TagName:       tagName,
		}

		declaration.CustomElement.Attributes = A.Chain(func(field manifest.CustomElementField) []manifest.Attribute {
			if field.Attribute == "" {
				return []manifest.Attribute{}
			} else {
				return []manifest.Attribute{{
					Name:        field.Attribute,
					Summary:     field.Summary,
					Description: field.Description,
					Deprecated:  field.Deprecated,
					Default:     field.Default,
					Type:        field.Type,
					FieldName:   field.Name,
				}}
			}
		})(fields)
	}

	jsdoc, ok := captures["jsdoc"]
	if (ok && len(jsdoc) > 0) {
		error, classInfo := NewClassInfo(jsdoc[0].Text)
		if error != nil {
			return errors.Join(err, error), nil
		} else {
			classInfo.MergeToCustomElementDeclaration(declaration)
		}
	}

	return nil, declaration
}
