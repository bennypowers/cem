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
package manifest

import (
	"slices"
	"sync"
)

// External caches for flattened members (pattern from lookup_helpers.go)
var (
	flattenedMembersCache = &sync.Map{}
	flattenedMembersOnce  = &sync.Map{}
)

// flattenedMembers holds all flattened member types for a custom element.
// Members from mixins will have their InheritedFrom field populated.
type flattenedMembers struct {
	attributes    []Attribute
	slots         []Slot
	events        []Event
	cssProperties []CssCustomProperty
	cssParts      []CssPart
	cssStates     []CssCustomState
	fields        []ClassMember
	methods       []ClassMember
}

// resolveSuperclass recursively resolves the superclass chain for a ClassLike.
// Returns superclasses in application order (base → derived).
// Tracks visited superclasses to prevent infinite loops from circular references.
func resolveSuperclass(pkg *Package, classLike *ClassLike, visited map[string]bool) []*ClassLike {
	if pkg == nil || classLike == nil || classLike.Superclass == nil {
		return nil
	}

	superclassRef := *classLike.Superclass

	// Skip if already visited (circular reference detection)
	if visited[superclassRef.Name] {
		return nil
	}

	// Find the superclass declaration
	decl := pkg.FindDeclaration(superclassRef)
	if decl == nil {
		// Superclass not found, skip gracefully
		return nil
	}

	// Extract ClassLike from the declaration
	var superclassLike *ClassLike
	switch d := decl.(type) {
	case *ClassDeclaration:
		superclassLike = &d.ClassLike
	case *CustomElementDeclaration:
		superclassLike = &d.ClassLike
	case *MixinDeclaration:
		superclassLike = &d.ClassLike
	case *CustomElementMixinDeclaration:
		superclassLike = &d.ClassLike
	default:
		// Not a valid superclass, skip
		return nil
	}

	// Mark as visited
	visited[superclassRef.Name] = true

	// Recursively resolve the superclass's superclass (depth-first)
	var result []*ClassLike
	nestedSuperclasses := resolveSuperclass(pkg, superclassLike, visited)
	result = append(result, nestedSuperclasses...)

	// Append this superclass
	result = append(result, superclassLike)

	return result
}

// resolveMixins recursively resolves all mixins for a ClassLike.
// Returns mixins in application order (bottom-up: base → derived).
// Tracks visited mixins to prevent infinite loops from circular references.
func resolveMixins(pkg *Package, classLike *ClassLike, visited map[string]bool) []Declaration {
	if pkg == nil || classLike == nil {
		return nil
	}

	var result []Declaration

	for _, mixinRef := range classLike.Mixins {
		// Skip if already visited (circular reference detection)
		if visited[mixinRef.Name] {
			continue
		}

		// Find the mixin declaration
		decl := pkg.FindDeclaration(mixinRef)
		if decl == nil {
			// Mixin not found, skip gracefully
			continue
		}

		// Extract ClassLike from the mixin declaration for recursive resolution
		var mixinClassLike *ClassLike
		switch m := decl.(type) {
		case *MixinDeclaration:
			mixinClassLike = &m.ClassLike
		case *CustomElementMixinDeclaration:
			mixinClassLike = &m.ClassLike
		default:
			// Not a mixin, skip
			continue
		}

		// Mark as visited
		visited[mixinRef.Name] = true

		// Recursively resolve this mixin's mixins (depth-first)
		nestedMixins := resolveMixins(pkg, mixinClassLike, visited)
		result = append(result, nestedMixins...)

		// Append this mixin declaration (not just ClassLike)
		result = append(result, decl)
	}

	return result
}

// mergeAttributes combines attributes from class and inherited sources (superclass/mixins).
// Inherited attributes have InheritedFrom set to the source reference.
// Class attributes take precedence and have InheritedFrom = nil.
// Description handling (avoid duplication):
//   - If class description is empty → use inherited description
//   - If class description == inherited description → use once (no duplication)
//   - If both have different descriptions → concatenate: class + "\n\n" + inherited
func mergeAttributes(classAttrs []Attribute, inheritedAttrs []Attribute, inheritedRef Reference) []Attribute {
	// Create map for O(1) lookup
	attrMap := make(map[string]Attribute)

	// Add all inherited attributes with InheritedFrom set
	for _, attr := range inheritedAttrs {
		// Only set InheritedFrom if not already set (preserve existing inheritance info)
		if attr.InheritedFrom == nil && inheritedRef.Name != "" {
			attr.InheritedFrom = &inheritedRef
		}
		attrMap[attr.Name] = attr
	}

	// Add class attributes (override inherited attributes)
	for _, attr := range classAttrs {
		if existing, exists := attrMap[attr.Name]; exists {
			// Merge: smart description handling, class precedence for values

			// Summary: use class if present, otherwise inherited
			if attr.Summary == "" {
				attr.Summary = existing.Summary
			}

			// Description: avoid duplication
			if attr.Description == "" {
				// Class has no description → use inherited
				attr.Description = existing.Description
			} else if attr.Description != existing.Description && existing.Description != "" {
				// Both have different descriptions → concatenate
				attr.Description = attr.Description + "\n\n" + existing.Description
			}
			// If descriptions match → use class description (no duplication)

			// Class attr overrides, so clear InheritedFrom
			attr.InheritedFrom = nil
		}
		attrMap[attr.Name] = attr
	}

	// Convert map to slice
	result := make([]Attribute, 0, len(attrMap))
	for _, attr := range attrMap {
		result = append(result, attr)
	}

	return result
}

// mergeSlots combines slots from class and inherited sources (superclass/mixins).
// Uses smart description deduplication like mergeAttributes.
func mergeSlots(classSlots []Slot, inheritedSlots []Slot, inheritedRef Reference) []Slot {
	slotMap := make(map[string]Slot)

	for _, slot := range inheritedSlots {
		// Only set InheritedFrom if not already set (preserve existing inheritance info)
		if slot.InheritedFrom == nil && inheritedRef.Name != "" {
			slot.InheritedFrom = &inheritedRef
		}
		slotMap[slot.Name] = slot
	}

	for _, slot := range classSlots {
		if existing, exists := slotMap[slot.Name]; exists {
			if slot.Summary == "" {
				slot.Summary = existing.Summary
			}
			// Description: avoid duplication
			if slot.Description == "" {
				slot.Description = existing.Description
			} else if slot.Description != existing.Description && existing.Description != "" {
				slot.Description = slot.Description + "\n\n" + existing.Description
			}
			slot.InheritedFrom = nil
		}
		slotMap[slot.Name] = slot
	}

	result := make([]Slot, 0, len(slotMap))
	for _, slot := range slotMap {
		result = append(result, slot)
	}

	return result
}

// mergeEvents combines events from class and mixins.
// Uses smart description deduplication like mergeAttributes.
func mergeEvents(classEvents []Event, mixinEvents []Event, mixinRef Reference) []Event {
	eventMap := make(map[string]Event)

	for _, event := range mixinEvents {
		// Only set InheritedFrom if not already set (preserve existing inheritance info)
		if event.InheritedFrom == nil && mixinRef.Name != "" {
			event.InheritedFrom = &mixinRef
		}
		eventMap[event.Name] = event
	}

	for _, event := range classEvents {
		if existing, exists := eventMap[event.Name]; exists {
			if event.Summary == "" {
				event.Summary = existing.Summary
			}
			// Description: avoid duplication
			if event.Description == "" {
				event.Description = existing.Description
			} else if event.Description != existing.Description && existing.Description != "" {
				event.Description = event.Description + "\n\n" + existing.Description
			}
			// If descriptions match → use class description (no duplication)
			event.InheritedFrom = nil
		}
		eventMap[event.Name] = event
	}

	result := make([]Event, 0, len(eventMap))
	for _, event := range eventMap {
		result = append(result, event)
	}

	return result
}

// mergeCssProperties combines CSS properties from class and mixins.
// Uses smart description deduplication like mergeAttributes.
func mergeCssProperties(classProps []CssCustomProperty, mixinProps []CssCustomProperty, mixinRef Reference) []CssCustomProperty {
	propMap := make(map[string]CssCustomProperty)

	for _, prop := range mixinProps {
		// Only set InheritedFrom if not already set (preserve existing inheritance info)
		if prop.InheritedFrom == nil && mixinRef.Name != "" {
			prop.InheritedFrom = &mixinRef
		}
		propMap[prop.Name] = prop
	}

	for _, prop := range classProps {
		if existing, exists := propMap[prop.Name]; exists {
			if prop.Summary == "" {
				prop.Summary = existing.Summary
			}
			// Description: avoid duplication
			if prop.Description == "" {
				prop.Description = existing.Description
			} else if prop.Description != existing.Description && existing.Description != "" {
				prop.Description = prop.Description + "\n\n" + existing.Description
			}
			// If descriptions match → use class description (no duplication)
			prop.InheritedFrom = nil
		}
		propMap[prop.Name] = prop
	}

	result := make([]CssCustomProperty, 0, len(propMap))
	for _, prop := range propMap {
		result = append(result, prop)
	}

	return result
}

// mergeCssParts combines CSS parts from class and mixins.
// Uses smart description deduplication like mergeAttributes.
func mergeCssParts(classParts []CssPart, mixinParts []CssPart, mixinRef Reference) []CssPart {
	partMap := make(map[string]CssPart)

	for _, part := range mixinParts {
		// Only set InheritedFrom if not already set (preserve existing inheritance info)
		if part.InheritedFrom == nil && mixinRef.Name != "" {
			part.InheritedFrom = &mixinRef
		}
		partMap[part.Name] = part
	}

	for _, part := range classParts {
		if existing, exists := partMap[part.Name]; exists {
			if part.Summary == "" {
				part.Summary = existing.Summary
			}
			// Description: avoid duplication
			if part.Description == "" {
				part.Description = existing.Description
			} else if part.Description != existing.Description && existing.Description != "" {
				part.Description = part.Description + "\n\n" + existing.Description
			}
			// If descriptions match → use class description (no duplication)
			part.InheritedFrom = nil
		}
		partMap[part.Name] = part
	}

	result := make([]CssPart, 0, len(partMap))
	for _, part := range partMap {
		result = append(result, part)
	}

	return result
}

// mergeCssStates combines CSS states from class and mixins.
// Uses smart description deduplication like mergeAttributes.
func mergeCssStates(classStates []CssCustomState, mixinStates []CssCustomState, mixinRef Reference) []CssCustomState {
	stateMap := make(map[string]CssCustomState)

	for _, state := range mixinStates {
		// Only set InheritedFrom if not already set (preserve existing inheritance info)
		if state.InheritedFrom == nil && mixinRef.Name != "" {
			state.InheritedFrom = &mixinRef
		}
		stateMap[state.Name] = state
	}

	for _, state := range classStates {
		if existing, exists := stateMap[state.Name]; exists {
			if state.Summary == "" {
				state.Summary = existing.Summary
			}
			// Description: avoid duplication
			if state.Description == "" {
				state.Description = existing.Description
			} else if state.Description != existing.Description && existing.Description != "" {
				state.Description = state.Description + "\n\n" + existing.Description
			}
			// If descriptions match → use class description (no duplication)
			state.InheritedFrom = nil
		}
		stateMap[state.Name] = state
	}

	result := make([]CssCustomState, 0, len(stateMap))
	for _, state := range stateMap {
		result = append(result, state)
	}

	return result
}

// mergeClassMembers combines class members (fields/methods) from class and mixins.
// Members can be fields or methods, we merge by name.
func mergeClassMembers(classMembers []ClassMember, mixinMembers []ClassMember, mixinRef Reference) []ClassMember {
	memberMap := make(map[string]ClassMember)

	for _, member := range mixinMembers {
		// Set InheritedFrom based on member type (only if not already set)
		switch m := member.(type) {
		case *ClassField:
			if m.InheritedFrom == nil && mixinRef.Name != "" {
				m.InheritedFrom = &mixinRef
			}
		case *ClassMethod:
			if m.InheritedFrom == nil && mixinRef.Name != "" {
				m.InheritedFrom = &mixinRef
			}
		}
		// Use a key combining name and kind for unique identification
		var key string
		switch m := member.(type) {
		case *ClassField:
			key = "field:" + m.Name
		case *ClassMethod:
			key = "method:" + m.Name
		default:
			continue
		}
		memberMap[key] = member
	}

	for _, member := range classMembers {
		var key string
		switch m := member.(type) {
		case *ClassField:
			key = "field:" + m.Name
			// Clear InheritedFrom for class members
			m.InheritedFrom = nil
		case *ClassMethod:
			key = "method:" + m.Name
			m.InheritedFrom = nil
		default:
			continue
		}

		if existing, exists := memberMap[key]; exists {
			// Merge summaries/descriptions with smart deduplication
			switch m := member.(type) {
			case *ClassField:
				if existingField, ok := existing.(*ClassField); ok {
					if m.Summary == "" {
						m.Summary = existingField.Summary
					}
					// Description: avoid duplication
					if m.Description == "" {
						m.Description = existingField.Description
					} else if m.Description != existingField.Description && existingField.Description != "" {
						m.Description = m.Description + "\n\n" + existingField.Description
					}
					// If descriptions match → use class description (no duplication)
				}
			case *ClassMethod:
				if existingMethod, ok := existing.(*ClassMethod); ok {
					if m.Summary == "" {
						m.Summary = existingMethod.Summary
					}
					// Description: avoid duplication
					if m.Description == "" {
						m.Description = existingMethod.Description
					} else if m.Description != existingMethod.Description && existingMethod.Description != "" {
						m.Description = m.Description + "\n\n" + existingMethod.Description
					}
					// If descriptions match → use class description (no duplication)
				}
			}
		}
		memberMap[key] = member
	}

	result := make([]ClassMember, 0, len(memberMap))
	for _, member := range memberMap {
		result = append(result, member)
	}

	return result
}

// flattenMembers resolves and flattens all members from class and mixins.
// Uses lazy caching with sync.Once for thread-safe performance.
// Mixin members will have InheritedFrom populated.
func (ced *CustomElementDeclaration) flattenMembers(pkg *Package) *flattenedMembers {
	if ced == nil || pkg == nil {
		return &flattenedMembers{}
	}

	// Create cache key
	cacheKey := ced

	// Try to get existing result
	if cached, ok := flattenedMembersCache.Load(cacheKey); ok {
		return cached.(*flattenedMembers)
	}

	// Ensure only one goroutine initializes
	onceValue, _ := flattenedMembersOnce.LoadOrStore(cacheKey, &sync.Once{})
	once := onceValue.(*sync.Once)

	var result *flattenedMembers
	once.Do(func() {
		result = &flattenedMembers{}

		// Start with empty slices
		var (
			attrs     []Attribute
			slots     []Slot
			events    []Event
			cssProps  []CssCustomProperty
			cssParts  []CssPart
			cssStates []CssCustomState
			fields    []ClassMember
		)

		// Resolution order: Superclass → Mixins → Class (matches JavaScript semantics)

		// 1. Resolve and process superclass chain first (simpler, linear inheritance)
		visitedSuperclass := make(map[string]bool)
		superclasses := resolveSuperclass(pkg, &ced.ClassLike, visitedSuperclass)

		for i, superclass := range superclasses {
			// Find the superclass reference for InheritedFrom
			// For the first superclass, use ced.Superclass
			// For chain members, find their reference
			var superclassRef Reference
			if i == 0 && ced.Superclass != nil {
				superclassRef = *ced.Superclass
			} else if superclass.Superclass != nil {
				superclassRef = *superclass.Superclass
			}

			// Check if this is a CustomElementDeclaration (has attributes, slots, etc.)
			// We need to find the full declaration to access CustomElement members
			if superclassRef.Name != "" {
				superclassDecl := pkg.FindDeclaration(superclassRef)
				if cedSuper, ok := superclassDecl.(*CustomElementDeclaration); ok {
					// CustomElement superclass - merge all custom element members
					attrs = mergeAttributes(attrs, cedSuper.OwnAttributes(), superclassRef)
					slots = mergeSlots(slots, cedSuper.OwnSlots(), superclassRef)
					events = mergeEvents(events, cedSuper.OwnEvents(), superclassRef)
					cssProps = mergeCssProperties(cssProps, cedSuper.OwnCssProperties(), superclassRef)
					cssParts = mergeCssParts(cssParts, cedSuper.OwnCssParts(), superclassRef)
					cssStates = mergeCssStates(cssStates, cedSuper.OwnCssStates(), superclassRef)
				}
			}

			// Merge superclass members (all superclasses contribute via ClassLike.Members)
			fields = mergeClassMembers(fields, superclass.Members, superclassRef)
		}

		// 2. Resolve and process mixins (can be multiple, more complex)
		visitedMixins := make(map[string]bool)
		mixins := resolveMixins(pkg, &ced.ClassLike, visitedMixins)

		// Process each mixin in order (base → derived)
		for _, mixinDecl := range mixins {
			// Extract name for the InheritedFrom reference
			var mixinName string
			var mixinClassLike *ClassLike
			switch m := mixinDecl.(type) {
			case *MixinDeclaration:
				mixinName = m.Name
				mixinClassLike = &m.ClassLike
			case *CustomElementMixinDeclaration:
				mixinName = m.Name
				mixinClassLike = &m.ClassLike
			}

			// Create reference for InheritedFrom
			mixinRef := Reference{Name: mixinName}

			// Find module path for this mixin
			for i := range pkg.Modules {
				if slices.Contains(pkg.Modules[i].Declarations, mixinDecl) {
					mixinRef.Module = pkg.Modules[i].Path
				}
				if mixinRef.Module != "" {
					break
				}
			}

			// Merge mixin members
			// For CustomElementMixin, merge custom element members (attributes, slots, etc.)
			if cemMixin, ok := mixinDecl.(*CustomElementMixinDeclaration); ok {
				// CustomElementMixin has attributes, slots, events, CSS
				// Use Own* methods to get direct members (not flattened)
				attrs = mergeAttributes(attrs, cemMixin.OwnAttributes(), mixinRef)
				slots = mergeSlots(slots, cemMixin.OwnSlots(), mixinRef)
				events = mergeEvents(events, cemMixin.OwnEvents(), mixinRef)
				cssProps = mergeCssProperties(cssProps, cemMixin.OwnCssProperties(), mixinRef)
				cssParts = mergeCssParts(cssParts, cemMixin.OwnCssParts(), mixinRef)
				cssStates = mergeCssStates(cssStates, cemMixin.OwnCssStates(), mixinRef)
			}

			// All mixins have class members (from ClassLike)
			if mixinClassLike != nil {
				fields = mergeClassMembers(fields, mixinClassLike.Members, mixinRef)
			}
		}

		// Finally merge with class's own members
		// Class members override mixin members
		result.attributes = mergeAttributes(ced.OwnAttributes(), attrs, Reference{})
		result.slots = mergeSlots(ced.OwnSlots(), slots, Reference{})
		result.events = mergeEvents(ced.OwnEvents(), events, Reference{})
		result.cssProperties = mergeCssProperties(ced.OwnCssProperties(), cssProps, Reference{})
		result.cssParts = mergeCssParts(ced.OwnCssParts(), cssParts, Reference{})
		result.cssStates = mergeCssStates(ced.OwnCssStates(), cssStates, Reference{})
		result.fields = mergeClassMembers(ced.Members, fields, Reference{})
		result.methods = result.fields // Both fields and methods are in Members

		// Store in cache
		flattenedMembersCache.Store(cacheKey, result)
	})

	// Return cached result (either just created or previously cached)
	if cached, ok := flattenedMembersCache.Load(cacheKey); ok {
		return cached.(*flattenedMembers)
	}

	return result
}
