package generate

import (
	"errors"

	M "bennypowers.dev/cem/manifest"
	A "github.com/IBM/fp-go/array"
	ts "github.com/tree-sitter/go-tree-sitter"
)

func generateCustomElementDeclaration(
	captures CaptureMap,
	root *ts.Node,
	code []byte,
) (errs error, declaration *M.CustomElementDeclaration) {
	err, classDeclaration := generateClassDeclaration(captures, root, code)
	if err != nil {
		errs = errors.Join(errs, err)
	}

	declaration = &M.CustomElementDeclaration{
		ClassDeclaration: *classDeclaration	,
	}

	tagNameNodes, ok := captures["tag-name"]
	if (!ok || len(tagNameNodes) <= 0) {
		return &NoCaptureError{ "tag-name", "customElementDeclaration"  }, nil
	}
	tagName := tagNameNodes[0].Text

	if (tagName != "") {
		declaration.CustomElement = M.CustomElement{
			CustomElement: true,
			TagName:       tagName,
		}

		// for i, member := range declaration.Members {
		// 	if member.Kind == "field" {
		//
		// 	}
		// }

		declaration.CustomElement.Attributes = A.Chain(func(member M.ClassMember) []M.Attribute {
			field, ok := (member).(M.CustomElementField)
			if (ok && field.Attribute != "") {
				return []M.Attribute{{
					Name:        field.Attribute,
					Summary:     field.Summary,
					Description: field.Description,
					Deprecated:  field.Deprecated,
					Default:     field.Default,
					Type:        field.Type,
					FieldName:   field.Name,
				}}
			} else {
				return []M.Attribute{}
			}
		})(declaration.Members)
	}

	jsdoc, ok := captures["jsdoc"]
	if (ok && len(jsdoc) > 0) {
		error, classInfo := NewClassInfo(jsdoc[0].Text)
		if error != nil {
			return errors.Join(errs, error), nil
		} else {
			classInfo.MergeToCustomElementDeclaration(declaration)
		}
	}

	return nil, declaration
}
