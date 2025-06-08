package generate

import (
	"errors"
	"regexp"
	"strings"

	M "bennypowers.dev/cem/manifest"
	"bennypowers.dev/cem/set"
	"dario.cat/mergo"
	ts "github.com/tree-sitter/go-tree-sitter"
)

var ignoredStaticFields = set.NewSet(
	"formAssociated",
	"observedAttributes",
	"shadowRootOptions",
	"styles",
)

var ignoredInstanceMethods = set.NewSet(
  "adoptedCallback",
  "connectedCallback",
  "connectedMoveCallback",
  "disconnectedCallback",
  "formDisabledCallback",
  "formStateRestoreCallback",
  "render",
  "update",
  "getUpdateComplete",
  "updated",
  "willUpdate",
)

func isPropertyField(captures CaptureMap) bool {
	_, hasDecorators := captures["decorator.name"]
	if hasDecorators {
		for _, name := range captures["decorator.name"] {
			if name.Text == "property" {
				return true
			}
		}
	}
	return false
}

func ammendFieldTypeWithCaptures(captures CaptureMap, field *M.ClassField) {
	for _, capture := range captures["field.type"] {
		typeText := capture.Text
		if typeText != "" {
			field.Type = &M.Type{
				Text: typeText,
			}
		}
	}
}

func ammendFieldPrivacyWithCaptures(captures CaptureMap, field *M.ClassField) {
	for _, capture := range captures["field.privacy"] {
		switch capture.Text {
			case "private":
				field.Privacy = M.Private
			case "protected":
				field.Privacy = M.Protected
		}
	}
}

func ammendFieldWithJsdocCaptures(captures CaptureMap, field *M.ClassField) (error, bool) {
	for _, x := range captures["member.jsdoc"] {
		source := x.Text
		error, info := NewPropertyInfo(source)
		if error != nil {
			return error, false
		} else {
			info.MergeToPropertyLike(&field.PropertyLike)
		}
	}
	return nil, true
}

func ammendFieldWithPropertyConfigCaptures(captures CaptureMap, field *M.CustomElementField) {
	field.Attribute = strings.ToLower(field.Name)
	bool, hasAttrBool := captures["field.attr.bool"]
	name, hasAttrName := captures["field.attr.name"]
	reflect, hasAttrReflect := captures["field.attr.reflect"]
	if hasAttrReflect {
		for _, capture := range reflect {
			field.Reflects = capture.Text == "true"
		}
	}
	if hasAttrBool {
		for _, capture := range bool {
			if capture.Text == "false" {
				field.Attribute = ""
			}
		}
	}
	if hasAttrName {
		for _, capture := range name {
			field.Attribute = capture.Text
		}
	}
}

func createClassFieldFromAccessorMatch(
	fieldName string,
	isStatic bool,
	isReadonly bool,
	isCustomElement bool,
	captures CaptureMap,
) (err error, field M.CustomElementField) {
	field.Kind = "field"
	field.Static = isStatic
	field.Name = fieldName
	field.Readonly = isReadonly

	ammendFieldTypeWithCaptures(captures, &field.ClassField)
	ammendFieldPrivacyWithCaptures(captures, &field.ClassField)

	isProperty := isCustomElement && !isStatic && isPropertyField(captures)

	if isProperty {
		ammendFieldWithPropertyConfigCaptures(captures, &field)
	}

	ammendFieldWithJsdocCaptures(captures, &field.ClassField)

	return err, field
}

func createClassFieldFromFieldMatch(fieldName string, isStatic bool, isCustomElement bool, captures CaptureMap) (err error, field M.CustomElementField) {
	_, readonly := captures["field.readonly"]

	field.Kind = "field"
	field.Static = isStatic
	field.Name = fieldName
	field.Readonly = readonly

	ammendFieldTypeWithCaptures(captures, &field.ClassField)
	ammendFieldPrivacyWithCaptures(captures, &field.ClassField)

	for _, x := range captures["field.initializer"] {
		field.Default = x.Text
		if field.Type == nil {
			if (field.Default == "true" || field.Default == "false") {
				field.Type = &M.Type{
					Text: "boolean",
				}
			} else if regexp.MustCompile(`^[\d._]+$`).MatchString(field.Default) {
				field.Type = &M.Type{
					Text: "number",
				}
			} else if regexp.MustCompile(`^('|")`).MatchString(field.Default) {
				field.Type = &M.Type{
					Text: "string",
				}
			}
		}
	}

	isProperty := isCustomElement && !isStatic && isPropertyField(captures)

	if isProperty {
		ammendFieldWithPropertyConfigCaptures(captures, &field)
	}

	ammendFieldWithJsdocCaptures(captures, &field.ClassField)

	return err, field
}

func getClassMembersFromClassDeclarationNode(
	code []byte,
	className string,
	root *ts.Node,
	classDeclarationNode *ts.Node,
	isCustomElement bool,
) (errs error, members []M.ClassMember) {
	qm, closeQm := NewQueryMatcher("classMemberDeclaration", Languages.Typescript)
	defer closeQm()
	qm.cursor.SetByteRange(classDeclarationNode.ByteRange())

	staticsSet := set.NewSet[string]()
	fieldsSet := set.NewSet[string]()
	gettersSet := set.NewSet[string]()
	settersSet := set.NewSet[string]()
	methodsSet := set.NewSet[string]()

	// to collate first instance of accessor pair / accessor
	accessorIndices := make(map[string]int)

	for captures := range qm.ParentCaptures(root, code, "accessor") {
		_, isStatic := captures["field.static"]
		fieldName := captures["field.name"][0].Text
		key := className + "." + fieldName
		if !gettersSet.Has(key) && !settersSet.Has(key) {
			accessorIndices[key] = len(members)
		}

		if gettersSet.Has(key) && settersSet.Has(key) {
			continue
		}

		fieldsSet.Add(key)
		accessorKind := captures["field.accessor"][0].Text
		var hasPair bool
		switch accessorKind {
		case "get":
			hasPair = settersSet.Has(key)
			if gettersSet.Has(key) {
				continue
			} else {
				gettersSet.Add(key)
			}
		case "set":
			hasPair = gettersSet.Has(key)
			if settersSet.Has(key) {
				continue
			} else {
				settersSet.Add(key)
			}
		}
		index := accessorIndices[key]
		if hasPair {
			// field is a *cem.CustomElementField
			isReadonly := false
			error, field := createClassFieldFromAccessorMatch(fieldName, isStatic, isReadonly, isCustomElement, captures)
			if error != nil {
				errs = errors.Join(errs, error)
			} else {
				item := members[index].(M.CustomElementField)
				item.Readonly = false
				err := mergo.Merge(&item, field, mergo.WithOverride)
				if err != nil {
					err = errors.Join(err, error)
				} else {
					members[index] = item
				}
			}
		} else {
			isReadonly := accessorKind == "get"
			error, field := createClassFieldFromAccessorMatch(fieldName, isStatic, isReadonly, isCustomElement, captures)
			if error != nil {
				errs = errors.Join(errs, error)
			} else {
				members = append(members, field)
			}
		}
	}

	for captures := range qm.ParentCaptures(root, code, "field") {
		_, isStatic := captures["field.static"]
		fieldName := captures["field.name"][0].Text
		key := className + "." + fieldName
		if (isStatic && ignoredStaticFields.Has(fieldName)) || (!isStatic && ignoredInstanceMethods.Has(fieldName)) {
			continue
		}

		if (isStatic && !staticsSet.Has(key)) {
			staticsSet.Add(key)
			error, field := createClassFieldFromFieldMatch(fieldName, isStatic, isCustomElement, captures)
			if error != nil {
				errs = errors.Join(errs, error)
			} else {
				members = append(members, field)
			}
		} else if (!isStatic && !fieldsSet.Has(key)) {
			fieldsSet.Add(key)
			error, field := createClassFieldFromFieldMatch(fieldName, isStatic, isCustomElement, captures)
			if error != nil {
				errs = errors.Join(errs, error)
			} else {
				members = append(members, field)
			}
		}
	}

	for captures := range qm.ParentCaptures(root, code, "method") {
		method := M.ClassMethod{Kind: "method"}
		methodNames, ok := captures["method.name"]
		methodName := methodNames[0]
		if ok {
			method.Name = methodName.Text
		}

		methodsSet.Add(method.Name)

		privacyNodes, ok := captures["method.privacy"]
		if ok && len(privacyNodes) > 0 {
			privacy := privacyNodes[0]
			method.Privacy = M.Privacy(privacy.Text)
		}

		_, isStatic := captures["method.static"]
		method.Static = isStatic

		params, hasParams := captures["params"]
		if hasParams && len(params) > 0 {
			for i, _ := range params {
				_, isRest := captures["param.rest"]
				_, hasName := captures["param.name"]
				if hasName {
					parameter := M.Parameter{
						Rest: isRest,
						PropertyLike: M.PropertyLike{
							Name: captures["param.name"][i].Text,
						},
					}
					_, hasType := captures["param.type"]
					if hasType {
						parameter.Type = &M.Type{
							Text: captures["param.type"][i].Text,
						}
					}
					method.Parameters = append(method.Parameters, parameter)
				}
			}
		}

		returnNodes, ok := captures["method.returns"]
		if ok && len(returnNodes) > 0 {
			returns := returnNodes[0]
			method.Return = &M.Return{
				Type: &M.Type{
					Text: returns.Text,
				},
			}
		}

		jsdocs, ok := captures["member.jsdoc"]
		if ok {
			for _, jsdoc := range jsdocs {
				merr, info := NewMethodInfo(jsdoc.Text)
				if merr != nil {
					errs = errors.Join(merr)
				} else {
					info.MergeToFunctionLike(&method.FunctionLike)
				}
			}
		}

		members = append(members, method)
	}

	return errs, members
}

