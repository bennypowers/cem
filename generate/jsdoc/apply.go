package jsdoc

import (
	"slices"

	M "bennypowers.dev/cem/manifest"
)

func applyToClassDeclaration(info *classInfo, declaration *M.ClassDeclaration) {
	declaration.Deprecated = info.Deprecated
	declaration.Description = info.Description
	declaration.Summary = info.Summary
}

func applyToCustomElementDeclaration(info *classInfo, declaration *M.CustomElementDeclaration) {
	applyToClassDeclaration(info, &declaration.ClassDeclaration)
	declaration.CustomElement.Attributes = slices.Concat(
		info.Attrs,
		declaration.CustomElement.Attributes,
	)

	jsdocSlots := make(map[string]M.Slot)
	for _, jsdocSlot := range info.Slots {
		jsdocSlots[jsdocSlot.Name] = jsdocSlot
	}

	for i := range declaration.CustomElement.Slots {
		if jsdocSlot, ok := jsdocSlots[declaration.CustomElement.Slots[i].Name]; ok {
			if declaration.CustomElement.Slots[i].Description == "" {
				declaration.CustomElement.Slots[i].Description = jsdocSlot.Description
			}
			if declaration.CustomElement.Slots[i].Summary == "" {
				declaration.CustomElement.Slots[i].Summary = jsdocSlot.Summary
			}
			if declaration.CustomElement.Slots[i].Deprecated == nil {
				declaration.CustomElement.Slots[i].Deprecated = jsdocSlot.Deprecated
			}
			delete(jsdocSlots, jsdocSlot.Name)
		}
	}

	for _, jsdocSlot := range jsdocSlots {
		declaration.CustomElement.Slots = append(declaration.CustomElement.Slots, jsdocSlot)
	}

	declaration.CustomElement.Events = info.Events
	declaration.CustomElement.CssProperties = info.CssProperties

	jsdocParts := make(map[string]M.CssPart)
	for _, jsdocPart := range info.CssParts {
		jsdocParts[jsdocPart.Name] = jsdocPart
	}

	for i := range declaration.CustomElement.CssParts {
		if jsdocPart, ok := jsdocParts[declaration.CustomElement.CssParts[i].Name]; ok {
			if declaration.CustomElement.CssParts[i].Description == "" {
				declaration.CustomElement.CssParts[i].Description = jsdocPart.Description
			}
			if declaration.CustomElement.CssParts[i].Summary == "" {
				declaration.CustomElement.CssParts[i].Summary = jsdocPart.Summary
			}
			if declaration.CustomElement.CssParts[i].Deprecated == nil {
				declaration.CustomElement.CssParts[i].Deprecated = jsdocPart.Deprecated
			}
			delete(jsdocParts, jsdocPart.Name)
		}
	}

	for _, jsdocPart := range jsdocParts {
		declaration.CustomElement.CssParts = append(declaration.CustomElement.CssParts, jsdocPart)
	}

	declaration.CustomElement.CssStates = info.CssStates
	declaration.CustomElement.Demos = info.Demos
	if info.TagName != "" {
		declaration.CustomElement.TagName = info.TagName
	}
}

func applyToCSSProperty(info *cssPropertyInfo, declaration *M.CssCustomProperty) {
	declaration.Description += info.Description
	declaration.Summary += info.Summary
	declaration.Deprecated = info.Deprecated
	if info.Syntax != "" {
		declaration.Syntax = info.Syntax
	}
}

func applyToPropertyLike(info *propertyInfo, declaration *M.PropertyLike) {
	declaration.Description += info.Description
	declaration.Summary += info.Summary
	declaration.Deprecated = info.Deprecated
	if info.Type != "" {
		declaration.Type = &M.Type{
			Text: info.Type,
		}
	}
}

func applyToFunctionLike(info *methodInfo, declaration *M.FunctionLike) {
	if info.Return != nil {
		if declaration.Return == nil {
			declaration.Return = &M.Return{}
		}
		declaration.Return.Description = info.Return.Description
		if info.Return.Type != "" {
			declaration.Return.Type = &M.Type{
				Text: info.Return.Type,
			}
		}
	}
	for _, iparam := range info.Parameters {
		for i := range declaration.Parameters {
			if declaration.Parameters[i].Name == iparam.Name {
				declaration.Parameters[i].Description = iparam.Description
				declaration.Parameters[i].Deprecated = iparam.Deprecated
				if iparam.Optional {
					declaration.Parameters[i].Optional = true
				}
				if iparam.Type != "" {
					declaration.Parameters[i].Type.Text = iparam.Type
				}
				if iparam.Default != "" {
					declaration.Parameters[i].Default = iparam.Default
				}
			}
		}
	}
}

func applyToMethod(info *methodInfo, declaration *M.ClassMethod) {
	declaration.Description = normalizeJsdocLines(info.Description)
	declaration.Deprecated = info.Deprecated
	declaration.Summary = info.Summary
	applyToFunctionLike(info, &declaration.FunctionLike)
}

func applyToFunctionDeclaration(info *methodInfo, declaration *M.FunctionDeclaration) {
	declaration.Description = normalizeJsdocLines(info.Description)
	declaration.Deprecated = info.Deprecated
	declaration.Summary = info.Summary
	applyToFunctionLike(info, &declaration.FunctionLike)
}