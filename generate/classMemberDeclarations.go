/*
Copyright © 2025 Benny Powers

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package generate

import (
	"errors"
	"regexp"
	"slices"
	"strings"

	M "bennypowers.dev/cem/manifest"
	Q "bennypowers.dev/cem/queries"
	S "bennypowers.dev/cem/set"
	"bennypowers.dev/cem/generate/jsdoc"

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
		if (isStatic && ignoredStaticFieldsHTML.Has(memberName)) ||
			(!isStatic && ignoredInstanceMethodsHTML.Has(memberName)) {
			return true
		}
	case "LitElement":
		if (isStatic && ignoredStaticFieldsHTML.Has(memberName)) ||
			(!isStatic && ignoredInstanceMethodsHTML.Has(memberName)) {
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

func amendFieldTypeWithCaptures(captures Q.CaptureMap, field *M.ClassField, logger *LogCtx) {
	for _, capture := range captures["field.type"] {
		typeText := capture.Text
		if typeText != "" {
			// Debug: log type extraction for variant property
			if field.Name == "variant" {
				logger.Debug("Extracted type for variant property: '%s'", typeText)
			}
			field.Type = &M.Type{
				Text: typeText,
			}
		}
	}
}

func amendFieldPrivacyWithCaptures(captures Q.CaptureMap, field *M.ClassField) {
	for _, capture := range captures["member.privacy"] {
		switch capture.Text {
		case "private":
			field.Privacy = M.Private
		case "protected":
			field.Privacy = M.Protected
		}
	}
}

func (mp ModuleProcessor) amendFieldWithJsdoc(captures Q.CaptureMap, field *M.ClassField) error {
	for _, x := range captures["member"] {
		jsdocText := jsdoc.ExtractFromNode(Q.GetDescendantById(mp.root, x.NodeId), mp.code)
		if jsdocText != "" {
			err := jsdoc.EnrichPropertyWithJSDoc(jsdocText, &field.PropertyLike, mp.queryManager)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (mp ModuleProcessor) amendMethodWithJsdoc(captures Q.CaptureMap, method *M.ClassMethod) error {
	// JSDoc
	if members, ok := captures["member"]; ok {
		for _, member := range members {
			jsdocText := jsdoc.ExtractFromNode(Q.GetDescendantById(mp.root, member.NodeId), mp.code)
			if jsdocText != "" {
				err := jsdoc.EnrichMethodWithJSDoc(jsdocText, method, mp.queryManager)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
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

func (mp *ModuleProcessor) createClassFieldFromAccessorMatch(
	fieldName string,
	isStatic bool,
	isReadonly bool,
	classType string,
	captures Q.CaptureMap,
) (field M.CustomElementField, err error) {
	field.Kind = "field"
	field.Static = isStatic
	field.Name = fieldName
	field.Readonly = isReadonly

	amendFieldTypeWithCaptures(captures, &field.ClassField, mp.logger)
	amendFieldPrivacyWithCaptures(captures, &field.ClassField)
	isCustomElement := classType == "HTMLElement" || classType == "LitElement"
	isProperty := isCustomElement && !isStatic && isPropertyField(captures)

	if isProperty {
		amendFieldWithPropertyConfigCaptures(captures, &field)
	}

	field.StartByte = captures["accessor"][0].StartByte
	err = mp.amendFieldWithJsdoc(captures, &field.ClassField)
	return field, err
}

func (mp *ModuleProcessor) createClassFieldFromConstructorParameterMatch(
	fieldName string,
	superclass string,
	captures Q.CaptureMap,
) (field M.CustomElementField, err error) {
	// Constructor parameters are never static
	isStatic := false

	field = M.CustomElementField{
		ClassField: M.ClassField{
			Kind:   "field",
			Static: isStatic,
			PropertyLike: M.PropertyLike{
				FullyQualified: M.FullyQualified{
					Name: fieldName,
				},
			},
		},
	}

	amendFieldTypeWithCaptures(captures, &field.ClassField, mp.logger)
	amendFieldPrivacyWithCaptures(captures, &field.ClassField)

	// Constructor parameters on non-CE classes typically don't have @property decorators
	// but we'll keep the logic for completeness
	isCustomElement := superclass == "HTMLElement" || superclass == "LitElement"
	isProperty := isCustomElement && !isStatic && isPropertyField(captures)

	if isProperty {
		amendFieldWithPropertyConfigCaptures(captures, &field)
	}

	err = mp.amendFieldWithJsdoc(captures, &field.ClassField)
	return field, err
}

func (mp *ModuleProcessor) createClassFieldFromFieldMatch(
	fieldName string,
	isStatic bool,
	superclass string,
	captures Q.CaptureMap,
) (field M.CustomElementField, err error) {
	_, readonly := captures["field.readonly"]

	field = M.CustomElementField{
		ClassField: M.ClassField{
			Kind:   "field",
			Static: isStatic,
			PropertyLike: M.PropertyLike{
				Readonly: readonly,
				FullyQualified: M.FullyQualified{
					Name: fieldName,
				},
			},
		},
	}

	amendFieldTypeWithCaptures(captures, &field.ClassField, mp.logger)
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

	_ = mp.amendFieldWithJsdoc(captures, &field.ClassField)
	return field, nil
}

// Identify kind and key
func getMemberKindFromCaptures(captures Q.CaptureMap) string {
	switch {
	case captures["method"] != nil:
		return "method"
	case captures["accessor"] != nil:
		return "accessor"
	case captures["field"] != nil:
		return "field"
	case captures["constructor.parameter"] != nil:
		return "constructor-parameter"
	default:
		return ""
	}
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
		name   string
		kind   string // field, accessor, method
		static bool
	}
	memberMap := make(map[memberKey]M.ClassMember)
	seenAccessors := make(map[string]memberKey) // for merging get/set

	for captures := range matcher.ParentCaptures(mp.root, mp.code, "member") {
		memberName := captures["member.name"][0].Text
		isStatic := len(captures["member.static"]) > 0
		kind := getMemberKindFromCaptures(captures)
		key := memberKey{name: memberName, kind: kind, static: isStatic}
		if kind == "" || isIgnoredMember(memberName, superclass, isStatic) {
			continue
		}

		// Debug: log all variant property processing
		if memberName == "variant" {
			mp.logger.Debug("Processing variant member: kind=%s, static=%v", kind, isStatic)
		}

		switch kind {
		case "field":
			field, error := mp.createClassFieldFromFieldMatch(memberName, isStatic, superclass, captures)
			field.StartByte = captures["field"][0].StartByte
			if error != nil {
				errs = errors.Join(errs, error)
			} else {
				memberMap[key] = &field
			}
		case "constructor-parameter":
			field, err := mp.createClassFieldFromConstructorParameterMatch(memberName, superclass, captures)
			field.StartByte = captures["constructor.parameter"][0].StartByte
			if err != nil {
				errs = errors.Join(errs, err)
			} else {
				memberMap[key] = &field
			}
		case "accessor":
			accessorKind := captures["field.accessor"][0].Text // "get" or "set"
			// Merge get/set into one field for same name+static
			pairKeyStr := className + "." + memberName + "." + isStaticToTypeFlag(isStatic)
			isReadonly := accessorKind == "get"

			field, err := mp.createClassFieldFromAccessorMatch(memberName, isStatic, isReadonly, superclass, captures)
			if err != nil {
				errs = errors.Join(errs, err)
				continue
			}

			// If we've seen the other half of the pair, merge readonly
			if prevKey, ok := seenAccessors[pairKeyStr]; ok {
				needle := memberMap[prevKey]
				existing := needle.(*M.CustomElementField)
				existing.Readonly = false // If both get/set, not readonly
				// Merge doc/types/join
				if field.Type != nil && (existing.Type == nil || existing.Type.Text == "") {
					existing.Type = field.Type
				}
				if field.Privacy != "" && existing.Privacy == "" {
					existing.Privacy = field.Privacy
				}
				// Merge jsdoc, attribute, etc. as needed
				if field.Description != "" && existing.Description == "" {
					existing.Description = field.Description
				}
				memberMap[prevKey] = existing
			} else {
				memberMap[key] = &field
				seenAccessors[pairKeyStr] = key
			}
		case "method":
			method := M.ClassMethod{Kind: "method"}
			method.Name = memberName
			method.Static = isStatic
			method.StartByte = captures["method"][0].StartByte

			err := mp.amendMethodWithJsdoc(captures, &method)
			if err != nil {
				errs = errors.Join(errs, err)
			}

			// Privacy
			if nodes, ok := captures["member.privacy"]; ok && len(nodes) > 0 {
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
								FullyQualified: M.FullyQualified{
									Name: captures["param.name"][i].Text,
								},
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
			memberMap[key] = &method
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
