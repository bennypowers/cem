package generate

import (
	"errors"
	"regexp"
	"slices"
	"strings"

	Q "bennypowers.dev/cem/generate/queries"
	M "bennypowers.dev/cem/manifest"
	S "bennypowers.dev/cem/set"

	ts "github.com/tree-sitter/go-tree-sitter"
)

var ignoredStaticFieldsHTML = S.NewSet(
	"formAssociated",
	"observedAttributes",
)

var ignoredStaticFieldsLit = S.NewSet(
	"shadowRootOptions",
	"styles",
)

var ignoredInstanceMethodsHTML = S.NewSet(
  "adoptedCallback",
  "connectedCallback",
  "connectedMoveCallback",
  "disconnectedCallback",
  "formDisabledCallback",
  "formStateRestoreCallback",
)

var ignoredInstanceMethodsLit = S.NewSet(
  "render",
  "update",
  "getUpdateComplete",
  "updated",
  "willUpdate",
)

var (
	numberLiteralRegexp = regexp.MustCompile(`^[\d._]+$`)
	stringLiteralRegexp = regexp.MustCompile(`^('|")`)
)

func isIgnoredMember(memberName string, superclass string, isStatic bool) bool {
	switch superclass {
	case "HTMLElement":
		if ((isStatic && ignoredStaticFieldsHTML.Has(memberName)) ||
				(!isStatic && ignoredInstanceMethodsHTML.Has(memberName))) {
			return true
		}
	case "LitElement":
		if ((isStatic && ignoredStaticFieldsHTML.Has(memberName)) ||
				(!isStatic && ignoredInstanceMethodsHTML.Has(memberName))) {
			return true
		}
		if (isStatic && ignoredStaticFieldsLit.Has(memberName)) ||
				(!isStatic && ignoredInstanceMethodsLit.Has(memberName)) {
			return true
		}
	}
	return false
}

func isPropertyField(captures Q.CaptureMap) bool {
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

func amendFieldTypeWithCaptures(captures Q.CaptureMap, field *M.ClassField) {
	for _, capture := range captures["field.type"] {
		typeText := capture.Text
		if typeText != "" {
			field.Type = &M.Type{
				Text: typeText,
			}
		}
	}
}

func amendFieldPrivacyWithCaptures(captures Q.CaptureMap, field *M.ClassField) {
	for _, capture := range captures["field.privacy"] {
		switch capture.Text {
			case "private":
				field.Privacy = M.Private
			case "protected":
				field.Privacy = M.Protected
		}
	}
}

func amendFieldWithJsdocCaptures(queryManager *Q.QueryManager, captures Q.CaptureMap, field *M.ClassField) (error, bool) {
	for _, x := range captures["member.jsdoc"] {
		source := x.Text
		error, info := NewPropertyInfo(source, queryManager)
		if error != nil {
			return error, false
		} else {
			info.MergeToPropertyLike(&field.PropertyLike)
		}
	}
	return nil, true
}

func amendFieldWithPropertyConfigCaptures(captures Q.CaptureMap, field *M.CustomElementField) {
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
	classType string,
	captures Q.CaptureMap,
	queryManager *Q.QueryManager,
) (err error, field M.CustomElementField) {
	field.Kind = "field"
	field.Static = isStatic
	field.Name = fieldName
	field.Readonly = isReadonly

	amendFieldTypeWithCaptures(captures, &field.ClassField)
	amendFieldPrivacyWithCaptures(captures, &field.ClassField)

	isCustomElement := classType == "HTMLElement" || classType == "LitElement"
	isProperty := isCustomElement && !isStatic && isPropertyField(captures)

	if isProperty {
		amendFieldWithPropertyConfigCaptures(captures, &field)
	}

	amendFieldWithJsdocCaptures(queryManager, captures, &field.ClassField)

	return err, field
}

func createClassFieldFromFieldMatch(
	fieldName string,
	isStatic bool,
	superclass string,
	captures Q.CaptureMap,
	queryManager *Q.QueryManager,
) (err error, field M.CustomElementField) {
	_, readonly := captures["field.readonly"]

	field = M.CustomElementField{
		ClassField: M.ClassField{
			Kind:    "field",
			Static:  isStatic,
			PropertyLike: M.PropertyLike{
				Name:    fieldName,
				Readonly: readonly,
			},
		},
	}

	amendFieldTypeWithCaptures(captures, &field.ClassField)
	amendFieldPrivacyWithCaptures(captures, &field.ClassField)

	for _, x := range captures["field.initializer"] {
		field.Default = strings.ReplaceAll(x.Text, "\n", "\\n")
		if field.Type == nil {
			switch {
			case field.Default == "true" || field.Default == "false":
				field.Type = &M.Type{Text: "boolean"}
			case numberLiteralRegexp.MatchString(field.Default):
				field.Type = &M.Type{Text: "number"}
			case stringLiteralRegexp.MatchString(field.Default):
				field.Type = &M.Type{Text: "string"}
			}
		}
	}

	isCustomElement := superclass == "HTMLElement" || superclass == "LitElement"
	isProperty := isCustomElement && !isStatic && isPropertyField(captures)

	if isProperty {
		amendFieldWithPropertyConfigCaptures(captures, &field)
	}

	amendFieldWithJsdocCaptures(queryManager, captures, &field.ClassField)
	return nil, field
}

func (mp *ModuleProcessor) getClassMembersFromClassDeclarationNode(
	className string,
	classDeclarationNode *ts.Node,
	superclass string,
) (
	members []M.ClassMember,
	errs error,
) {
	matcher, err := Q.NewQueryMatcher(mp.queryManager, "typescript", "classMemberDeclaration")
	if err != nil {
		return members, errors.Join(errs, err)
	}
	defer matcher.Close()
	matcher.SetByteRange(classDeclarationNode.ByteRange())

	// Key: className + "." + fieldName + "." + kind + ".static"
	type memberKey struct {
		name     string
		kind     string // field, accessor, method
		static   bool
	}
	memberMap := make(map[memberKey]M.ClassMember)
	accessorPairIndices := make(map[string]memberKey) // for merging get/set

	for captures := range matcher.ParentCaptures(mp.root, mp.code, "member") {
		var (
			memberName string
			isStatic   bool
			kind       string
		)

		// Identify kind and key
		switch {
		case captures["method"] != nil:
			kind = "method"
			memberName = captures["method.name"][0].Text
			isStatic = len(captures["method.static"]) > 0
		case captures["accessor"] != nil:
			kind = "accessor"
			memberName = captures["field.name"][0].Text
			isStatic = len(captures["field.static"]) > 0
		case captures["field"] != nil:
			kind = "field"
			memberName = captures["field.name"][0].Text
			isStatic = len(captures["field.static"]) > 0
		default:
			continue
		}
		if isIgnoredMember(memberName, superclass, isStatic) {
			continue
		}
		key := memberKey{name: memberName, kind: kind, static: isStatic}

		switch kind {
		case "field":
			error, field := createClassFieldFromFieldMatch(memberName, isStatic, superclass, captures, mp.queryManager)
			field.ClassField.StartByte = captures["field"][0].StartByte
			if error != nil {
				errs = errors.Join(errs, error)
			} else {
				memberMap[key] = field
			}
		case "accessor":
			accessorKind := captures["field.accessor"][0].Text // "get" or "set"
			// Merge get/set into one field for same name+static
			pairKeyStr := className + "." + memberName + "." + isStaticToTypeFlag(isStatic)
			var isReadonly bool = accessorKind == "get"

			error, field := createClassFieldFromAccessorMatch(memberName, isStatic, isReadonly, superclass, captures, mp.queryManager)
			field.ClassField.StartByte = captures["accessor"][0].StartByte
			if error != nil {
				errs = errors.Join(errs, error)
				continue
			}

			// If we've seen the other half of the pair, merge readonly
			if prevKey, ok := accessorPairIndices[pairKeyStr]; ok {
				existing := memberMap[prevKey].(M.CustomElementField)
				existing.Readonly = false // If both get/set, not readonly
				// Merge doc/types/join
				if field.Type != nil && (existing.Type == nil || existing.Type.Text == "") {
					existing.Type = field.Type
				}
				if field.Privacy != "" && existing.Privacy == "" {
					existing.Privacy = field.Privacy
				}
				// Merge jsdoc, attribute, etc. as needed
				if field.PropertyLike.Description != "" && existing.PropertyLike.Description == "" {
					existing.PropertyLike.Description = field.PropertyLike.Description
				}
				memberMap[prevKey] = existing
			} else {
				memberMap[key] = field
				accessorPairIndices[pairKeyStr] = key
			}
		case "method":
			method := M.ClassMethod{Kind: "method"}
			method.Name = memberName
			method.Static = isStatic
			method.StartByte = captures["method"][0].StartByte

			// Privacy
			if nodes, ok := captures["method.privacy"]; ok && len(nodes) > 0 {
				method.Privacy = M.Privacy(nodes[0].Text)
			}

			// Parameters
			if params, hasParams := captures["params"]; hasParams && len(params) > 0 {
				for i := range params {
					var isRest bool
					if _, ok := captures["param.rest"]; ok {
						isRest = true
					}
					if _, hasName := captures["param.name"]; hasName {
						parameter := M.Parameter{
							Rest: isRest,
							PropertyLike: M.PropertyLike{
								Name: captures["param.name"][i].Text,
							},
						}
						if _, hasType := captures["param.type"]; hasType {
							parameter.Type = &M.Type{
								Text: captures["param.type"][i].Text,
							}
						}
						method.Parameters = append(method.Parameters, parameter)
					}
				}
			}

			// Return type
			if returns, ok := captures["method.returns"]; ok && len(returns) > 0 {
				method.Return = &M.Return{
					Type: &M.Type{
						Text: returns[0].Text,
					},
				}
			}
			// JSDoc
			if jsdocs, ok := captures["member.jsdoc"]; ok {
				for _, jsdoc := range jsdocs {
					merr, info := NewMethodInfo(jsdoc.Text, mp.queryManager)
					if merr != nil {
						errs = errors.Join(merr)
					} else {
						info.MergeToFunctionLike(&method.FunctionLike)
					}
				}
			}
			memberMap[key] = method
		}
	}

	// Collect in stable order (optional: sort if you want)
	for _, member := range memberMap {
		members = append(members, member)
	}

	slices.SortStableFunc(members, func(a, b M.ClassMember) int {
    return int(a.GetStartByte() - b.GetStartByte())
})

	return members, errs
}

func isStaticToTypeFlag(b bool) string {
	if b {
		return "static"
	}
	return "instance"
}
