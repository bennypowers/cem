package generate

import (
	"errors"

	M "bennypowers.dev/cem/manifest"
	A "github.com/IBM/fp-go/array"
	ts "github.com/tree-sitter/go-tree-sitter"
)

func generateClassDeclaration(
	queryManager *QueryManager,
	captures CaptureMap,
	root *ts.Node,
	code []byte,
) (declaration M.Declaration, err error) {
	_, isCustomElement := captures["customElement"]
	classDeclarationCaptures, hasClassDeclaration := captures["class.declaration"]
	if (hasClassDeclaration && len(classDeclarationCaptures) > 0) {
		classDeclarationNodeId := classDeclarationCaptures[0].NodeId
		classDeclarationNode := GetDescendantById(root, classDeclarationNodeId)
		if isCustomElement {
			return generateCustomElementClassDeclaration(queryManager, captures, root, classDeclarationNode, code)
		} else {
			return generateCommonClassDeclaration(queryManager, captures, root, classDeclarationNode, code, isCustomElement)
		}
	}
	return nil, errors.New("Could not find class declaration")
}

func generateCommonClassDeclaration(
	queryManager *QueryManager,
	captures CaptureMap,
	root *ts.Node,
	classDeclarationNode *ts.Node,
	code []byte,
	isCustomElement bool,
) (declaration *M.ClassDeclaration, errs error) {
	className, ok := captures["class.name"]
	if (!ok || len(className) <= 0) {
		return nil, errors.Join(errs, &NoCaptureError{ "class.name", "classDeclaration" })
	}

	declaration = &M.ClassDeclaration{
		Kind: "class",
		ClassLike: M.ClassLike{
			Name: className[0].Text,
		},
	}

	err, members := getClassMembersFromClassDeclarationNode(
		queryManager,
		code,
		declaration.ClassLike.Name,
		root,
		classDeclarationNode,
		isCustomElement,
	)
	if err != nil {
		errs = errors.Join(errs, err)
	}

	for _, method := range members {
		declaration.Members = append(declaration.Members, method)
	}

	superClassName, ok := captures["superclass.name"]
	if (ok && len(superClassName) > 0) {
		name := superClassName[0].Text
		pkg := ""
		module := ""
		switch name {
		case
			"Event",
			"CustomEvent",
			"ErrorEvent",
			"HTMLElement":
		  pkg = "global:"
		case "LitElement":
			pkg = "lit"
		case "ReactiveElement":
			pkg = "@lit/reactive-element"
		// TODO: compute package and module
		// default:
		}
		declaration.Superclass = M.NewReference(name, pkg, module)
	}

	jsdoc, ok := captures["class.jsdoc"]
	if (ok && len(jsdoc) > 0) {
		err, info := NewClassInfo(jsdoc[0].Text, queryManager)
		if err != nil {
			errs = errors.Join(errs, err)
		} else {
			info.MergeToClassDeclaration(declaration)
		}
	}

	return declaration, nil
}

func generateCustomElementClassDeclaration(
	queryManager *QueryManager,
	captures CaptureMap,
	root *ts.Node,
	classDeclarationNode *ts.Node,
	code []byte,
) (declaration *M.CustomElementDeclaration, errs error) {
	classDeclaration, err := generateCommonClassDeclaration(
		queryManager,
		captures,
		root,
		classDeclarationNode,
		code,
		true,
	)
	if err != nil {
		errs = errors.Join(errs, err)
	}

	declaration = &M.CustomElementDeclaration{
		ClassDeclaration: *classDeclaration,
	}

	tagNameNodes, ok := captures["tag-name"]
	if (!ok || len(tagNameNodes) < 1) {
		errs = errors.Join(errs, &NoCaptureError{ "tag-name", "customElementDeclaration"  })
	} else {

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
	}

	jsdoc, ok := captures["class.jsdoc"]
	if (ok && len(jsdoc) > 0) {
		error, classInfo := NewClassInfo(jsdoc[0].Text, queryManager)
		if error != nil {
			errs = errors.Join(errs, error)
		} else {
			classInfo.MergeToCustomElementDeclaration(declaration)
		}
	}

	return declaration, nil
}

