package jsdoc

import (
	"slices"

	M "bennypowers.dev/cem/manifest"
)

// appendWithSeparator appends new content to existing content with a separator if both are non-empty
func appendWithSeparator(existing, new, separator string) string {
	if new == "" {
		return existing
	}
	if existing == "" {
		return new
	}
	return existing + separator + new
}

func applyToClassDeclaration(info *classInfo, declaration *M.ClassDeclaration) {
	declaration.Deprecated = info.Deprecated
	declaration.Description = info.Description
	declaration.Summary = info.Summary
}

func applyToCustomElementDeclaration(info *classInfo, declaration *M.CustomElementDeclaration) {
	applyToClassDeclaration(info, &declaration.ClassDeclaration)
	declaration.Attributes = slices.Concat(
		info.Attrs,
		declaration.Attributes,
	)

	jsdocSlots := make(map[string]M.Slot)
	for _, jsdocSlot := range info.Slots {
		jsdocSlots[jsdocSlot.Name] = jsdocSlot
	}

	for i := range declaration.Slots {
		if jsdocSlot, ok := jsdocSlots[declaration.Slots[i].Name]; ok {
			if declaration.Slots[i].Description == "" {
				declaration.Slots[i].Description = jsdocSlot.Description
			}
			if declaration.Slots[i].Summary == "" {
				declaration.Slots[i].Summary = jsdocSlot.Summary
			}
			if declaration.Slots[i].Deprecated == nil {
				declaration.Slots[i].Deprecated = jsdocSlot.Deprecated
			}
			delete(jsdocSlots, jsdocSlot.Name)
		}
	}

	for _, jsdocSlot := range jsdocSlots {
		declaration.Slots = append(declaration.Slots, jsdocSlot)
	}

	declaration.Events = info.Events
	declaration.CssProperties = info.CssProperties

	jsdocParts := make(map[string]M.CssPart)
	for _, jsdocPart := range info.CssParts {
		jsdocParts[jsdocPart.Name] = jsdocPart
	}

	for i := range declaration.CssParts {
		if jsdocPart, ok := jsdocParts[declaration.CssParts[i].Name]; ok {
			if declaration.CssParts[i].Description == "" {
				declaration.CssParts[i].Description = jsdocPart.Description
			}
			if declaration.CssParts[i].Summary == "" {
				declaration.CssParts[i].Summary = jsdocPart.Summary
			}
			if declaration.CssParts[i].Deprecated == nil {
				declaration.CssParts[i].Deprecated = jsdocPart.Deprecated
			}
			delete(jsdocParts, jsdocPart.Name)
		}
	}

	for _, jsdocPart := range jsdocParts {
		declaration.CssParts = append(declaration.CssParts, jsdocPart)
	}

	declaration.CssStates = info.CssStates
	declaration.Demos = info.Demos
	if info.TagName != "" {
		declaration.TagName = info.TagName
	}
}

func applyToCSSProperty(info *cssPropertyInfo, declaration *M.CssCustomProperty) {
	declaration.Description = appendWithSeparator(declaration.Description, info.Description, "\n\n")
	declaration.Summary = appendWithSeparator(declaration.Summary, info.Summary, "\n\n")
	declaration.Deprecated = info.Deprecated
	if info.Syntax != "" {
		declaration.Syntax = info.Syntax
	}
}

func applyToPropertyLike(info *propertyInfo, declaration *M.PropertyLike) {
	declaration.Description = appendWithSeparator(declaration.Description, info.Description, "\n\n")
	declaration.Summary = appendWithSeparator(declaration.Summary, info.Summary, "\n\n")
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
