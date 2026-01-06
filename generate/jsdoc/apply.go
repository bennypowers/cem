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

	jsdocAttrs := make(map[string]M.Attribute)
	for _, jsdocAttr := range info.Attrs {
		jsdocAttrs[jsdocAttr.Name] = jsdocAttr
	}

	for i := range declaration.CustomElement.Attributes {
		if jsdocAttr, ok := jsdocAttrs[declaration.CustomElement.Attributes[i].Name]; ok {
			if declaration.CustomElement.Attributes[i].Description == "" {
				declaration.CustomElement.Attributes[i].Description = jsdocAttr.Description
			}
			if declaration.CustomElement.Attributes[i].Summary == "" {
				declaration.CustomElement.Attributes[i].Summary = jsdocAttr.Summary
			}
			if declaration.CustomElement.Attributes[i].Type == nil {
				declaration.CustomElement.Attributes[i].Type = jsdocAttr.Type
			}
			if declaration.CustomElement.Attributes[i].Default == "" {
				declaration.CustomElement.Attributes[i].Default = jsdocAttr.Default
			}
			if declaration.CustomElement.Attributes[i].Deprecated == nil {
				declaration.CustomElement.Attributes[i].Deprecated = jsdocAttr.Deprecated
			}
			delete(jsdocAttrs, jsdocAttr.Name)
		}
	}

	// Collect remaining attribute names and sort for deterministic output
	remainingAttrNames := make([]string, 0, len(jsdocAttrs))
	for name := range jsdocAttrs {
		remainingAttrNames = append(remainingAttrNames, name)
	}
	slices.Sort(remainingAttrNames)

	// Append remaining JSDoc-only attributes in sorted order
	for _, name := range remainingAttrNames {
		declaration.CustomElement.Attributes = append(declaration.CustomElement.Attributes, jsdocAttrs[name])
	}

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

	// Collect remaining slot names and sort for deterministic output
	remainingSlotNames := make([]string, 0, len(jsdocSlots))
	for name := range jsdocSlots {
		remainingSlotNames = append(remainingSlotNames, name)
	}
	slices.Sort(remainingSlotNames)

	// Append remaining JSDoc-only slots in sorted order
	for _, name := range remainingSlotNames {
		declaration.CustomElement.Slots = append(declaration.CustomElement.Slots, jsdocSlots[name])
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

	// Collect remaining CSS part names and sort for deterministic output
	remainingPartNames := make([]string, 0, len(jsdocParts))
	for name := range jsdocParts {
		remainingPartNames = append(remainingPartNames, name)
	}
	slices.Sort(remainingPartNames)

	// Append remaining JSDoc-only CSS parts in sorted order
	for _, name := range remainingPartNames {
		declaration.CustomElement.CssParts = append(declaration.CustomElement.CssParts, jsdocParts[name])
	}

	declaration.CustomElement.CssStates = info.CssStates
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
