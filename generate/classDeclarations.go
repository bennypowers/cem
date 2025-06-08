package generate

import (
	"errors"
	"fmt"

	M "bennypowers.dev/cem/manifest"
	A "github.com/IBM/fp-go/array"
	ts "github.com/tree-sitter/go-tree-sitter"
)

func generateCommonClassDeclaration(
	captures CaptureMap,
	root *ts.Node,
	code []byte,
	isCustomElement bool,
) (errs error, declaration *M.ClassDeclaration) {
	declaration = &M.ClassDeclaration{ Kind: "class" }

	className, ok := captures["class.name"]
	if (!ok || len(className) <= 0) {
		str := string(code)
		fmt.Println(str)
		return errors.Join(errs, &NoCaptureError{ "class.name", "classDeclaration" }), nil
	}

	declaration.ClassLike.Name = className[0].Text

	err, members := getClassMembersFromClassDeclarationNode(code, root, isCustomElement)
	if err != nil {
		errs = errors.Join(errs, err)
	}

	for _, method := range members {
		declaration.Members = append(declaration.Members, method)
	}

	jsdoc, ok := captures["jsdoc"]
	if (ok && len(jsdoc) > 0) {
		err, info := NewClassInfo(jsdoc[0].Text)
		if err != nil {
			errs = errors.Join(errs, err)
		} else {
			info.MergeToClassDeclaration(declaration)
		}
	}

	return nil, declaration
}

func generateCustomElementClassDeclaration(
	captures CaptureMap,
	root *ts.Node,
	code []byte,
) (errs error, declaration *M.CustomElementDeclaration) {
	err, classDeclaration := generateCommonClassDeclaration(captures, root, code, true)
	if err != nil {
		errs = errors.Join(errs, err)
	}

	declaration = &M.CustomElementDeclaration{
		ClassDeclaration: *classDeclaration,
	}

	tagNameNodes, ok := captures["tag-name"]
	if (!ok || len(tagNameNodes) <= 0) {
		errs = errors.Join(errs, &NoCaptureError{ "tag-name", "customElementDeclaration"  })
	}

	tagName := tagNameNodes[0].Text

	if (tagName != "") {
		declaration.CustomElement = M.CustomElement{
			CustomElement: true,
			TagName:       tagName,
		}

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
			errs = errors.Join(errs, error)
		} else {
			classInfo.MergeToCustomElementDeclaration(declaration)
		}
	}

	return nil, declaration
}

func generateClassDeclaration(
	captures CaptureMap,
	root *ts.Node,
	code []byte,
) (err error, declaration M.Declaration) {
	_, isCustomElement := captures["customElement"]
	if isCustomElement {
		return generateCustomElementClassDeclaration(captures, root, code)
	} else {
		return generateCommonClassDeclaration(captures, root, code, isCustomElement)
	}
}

