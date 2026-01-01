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
//
// Cache Invalidation Policy:
// - Keys are *CustomElementDeclaration pointers, which are stable across the package lifecycle
// - Declarations are assumed to be immutable after initial package unmarshaling/population
// - No invalidation mechanism is provided, as declarations are not modified incrementally
// - If incremental updates are needed in the future, cache clearing should be added
//
// Memory Considerations:
// - Caches grow unbounded (no size limits or eviction policy)
// - For LSP servers or long-running processes that load many packages:
//   - Memory usage is proportional to the number of unique CustomElementDeclarations processed
//   - Each cached entry contains slices of attributes, slots, events, etc.
//   - Consider periodically clearing caches if memory becomes a concern
//
// - Typical usage (CLI tools processing a single package) has negligible memory impact
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
	// Use composite key to handle same-named classes across different modules
	visitKey := superclassRef.Module + ":" + superclassRef.Name
	if visited[visitKey] {
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
	visited[visitKey] = true

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
		// Use composite key to handle same-named mixins across different modules
		visitKey := mixinRef.Module + ":" + mixinRef.Name
		if visited[visitKey] {
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
		visited[visitKey] = true

		// Recursively resolve this mixin's mixins (depth-first)
		nestedMixins := resolveMixins(pkg, mixinClassLike, visited)
		result = append(result, nestedMixins...)

		// Append this mixin declaration (not just ClassLike)
		result = append(result, decl)
	}

	return result
}

// Merge Functions Design Note:
//
// The following merge functions (mergeAttributes, mergeSlots, mergeEvents, etc.) follow
// similar patterns but are implemented separately for each member type rather than using
// Go generics. This is intentional:
//
// 1. Type Safety: Each function works with specific types (Attribute, Slot, etc.) with
//    different field structures, making compile-time type checking straightforward.
//
// 2. Clarity: Explicit functions are easier to understand and debug than generic code
//    with interface constraints and type parameters.
//
// 3. Member-Specific Logic: While the patterns are similar, each type has unique fields
//    (e.g., Attribute.FieldName, Slot-specific handling) that would require complex
//    interface definitions or reflection to handle generically.
//
// 4. Performance: No runtime type assertions or reflection overhead.
//
// The code duplication (~200 lines) is a conscious trade-off for maintainability and
// type safety. If generics support improves or patterns diverge significantly, this
// decision can be revisited.

// mergeAttributes combines attributes from class and inherited sources (superclass/mixins).
// Inherited attributes have InheritedFrom set to the source reference.
// Class attributes take precedence and have InheritedFrom = nil.
// Preserves source order: inherited attributes first (not overridden), then class attributes.
// Description handling (avoid duplication):
//   - If class description is empty → use inherited description
//   - If class description == inherited description → use once (no duplication)
//   - If both have different descriptions → concatenate: class + "\n\n" + inherited
func mergeAttributes(classAttrs []Attribute, inheritedAttrs []Attribute, inheritedRef Reference) []Attribute {
	// Build map for O(1) override detection
	classMap := make(map[string]Attribute)
	for _, attr := range classAttrs {
		classMap[attr.Name] = attr
	}

	// Result slice to preserve order
	result := make([]Attribute, 0, len(inheritedAttrs)+len(classAttrs))

	// 1. Add inherited attributes that aren't overridden (preserves order from inheritedAttrs)
	for _, attr := range inheritedAttrs {
		// Set InheritedFrom if not already set
		if attr.InheritedFrom == nil && inheritedRef.Name != "" {
			// Create a copy of inheritedRef to avoid pointer aliasing
			ref := inheritedRef
			attr.InheritedFrom = &ref
		}

		// Only add if not overridden by class
		if _, overridden := classMap[attr.Name]; !overridden {
			result = append(result, attr)
		}
	}

	// 2. Add class attributes (preserves order from classAttrs, overrides inherited)
	for _, attr := range classAttrs {
		// Check if we're overriding an inherited attribute
		if inheritedAttr, exists := getAttributeFromSlice(inheritedAttrs, attr.Name); exists {
			// Merge summaries/descriptions with smart deduplication
			if attr.Summary == "" {
				attr.Summary = inheritedAttr.Summary
			}

			// Description: avoid duplication
			if attr.Description == "" {
				attr.Description = inheritedAttr.Description
			} else if attr.Description != inheritedAttr.Description && inheritedAttr.Description != "" {
				attr.Description = attr.Description + "\n\n" + inheritedAttr.Description
			}
		}

		// Clear InheritedFrom for class attributes
		attr.InheritedFrom = nil
		result = append(result, attr)
	}

	return result
}

// getAttributeFromSlice finds an attribute in a slice by name (for description merging)
func getAttributeFromSlice(attrs []Attribute, name string) (Attribute, bool) {
	for _, attr := range attrs {
		if attr.Name == name {
			return attr, true
		}
	}
	return Attribute{}, false
}

// mergeSlots combines slots from class and inherited sources (superclass/mixins).
// Preserves source order: inherited slots first (not overridden), then class slots.
// Uses smart description deduplication like mergeAttributes.
func mergeSlots(classSlots []Slot, inheritedSlots []Slot, inheritedRef Reference) []Slot {
	// Build map for O(1) override detection
	classMap := make(map[string]Slot)
	for _, slot := range classSlots {
		classMap[slot.Name] = slot
	}

	// Result slice to preserve order
	result := make([]Slot, 0, len(inheritedSlots)+len(classSlots))

	// 1. Add inherited slots that aren't overridden (preserves order from inheritedSlots)
	for _, slot := range inheritedSlots {
		// Set InheritedFrom if not already set
		if slot.InheritedFrom == nil && inheritedRef.Name != "" {
			// Create a copy of inheritedRef to avoid pointer aliasing
			ref := inheritedRef
			slot.InheritedFrom = &ref
		}

		// Only add if not overridden by class
		if _, overridden := classMap[slot.Name]; !overridden {
			result = append(result, slot)
		}
	}

	// 2. Add class slots (preserves order from classSlots, overrides inherited)
	for _, slot := range classSlots {
		// Check if we're overriding an inherited slot
		if inheritedSlot, exists := getSlotFromSlice(inheritedSlots, slot.Name); exists {
			// Merge summaries/descriptions with smart deduplication
			if slot.Summary == "" {
				slot.Summary = inheritedSlot.Summary
			}

			// Description: avoid duplication
			if slot.Description == "" {
				slot.Description = inheritedSlot.Description
			} else if slot.Description != inheritedSlot.Description && inheritedSlot.Description != "" {
				slot.Description = slot.Description + "\n\n" + inheritedSlot.Description
			}
		}

		// Clear InheritedFrom for class slots
		slot.InheritedFrom = nil
		result = append(result, slot)
	}

	return result
}

// getSlotFromSlice finds a slot in a slice by name (for description merging)
func getSlotFromSlice(slots []Slot, name string) (Slot, bool) {
	for _, slot := range slots {
		if slot.Name == name {
			return slot, true
		}
	}
	return Slot{}, false
}

// mergeEvents combines events from class and mixins.
// Preserves source order: inherited events first (not overridden), then class events.
// Uses smart description deduplication like mergeAttributes.
func mergeEvents(classEvents []Event, inheritedEvents []Event, inheritedRef Reference) []Event {
	// Build map for O(1) override detection
	classMap := make(map[string]Event)
	for _, event := range classEvents {
		classMap[event.Name] = event
	}

	// Result slice to preserve order
	result := make([]Event, 0, len(inheritedEvents)+len(classEvents))

	// 1. Add inherited events that aren't overridden (preserves order from inheritedEvents)
	for _, event := range inheritedEvents {
		// Set InheritedFrom if not already set
		if event.InheritedFrom == nil && inheritedRef.Name != "" {
			// Create a copy of inheritedRef to avoid pointer aliasing
			ref := inheritedRef
			event.InheritedFrom = &ref
		}

		// Only add if not overridden by class
		if _, overridden := classMap[event.Name]; !overridden {
			result = append(result, event)
		}
	}

	// 2. Add class events (preserves order from classEvents, overrides inherited)
	for _, event := range classEvents {
		// Check if we're overriding an inherited event
		if inheritedEvent, exists := getEventFromSlice(inheritedEvents, event.Name); exists {
			// Merge summaries/descriptions with smart deduplication
			if event.Summary == "" {
				event.Summary = inheritedEvent.Summary
			}

			// Description: avoid duplication
			if event.Description == "" {
				event.Description = inheritedEvent.Description
			} else if event.Description != inheritedEvent.Description && inheritedEvent.Description != "" {
				event.Description = event.Description + "\n\n" + inheritedEvent.Description
			}
		}

		// Clear InheritedFrom for class events
		event.InheritedFrom = nil
		result = append(result, event)
	}

	return result
}

// getEventFromSlice finds an event in a slice by name (for description merging)
func getEventFromSlice(events []Event, name string) (Event, bool) {
	for _, event := range events {
		if event.Name == name {
			return event, true
		}
	}
	return Event{}, false
}

// mergeCssProperties combines CSS properties from class and mixins.
// Preserves source order: inherited properties first (not overridden), then class properties.
// Uses smart description deduplication like mergeAttributes.
func mergeCssProperties(classProps []CssCustomProperty, inheritedProps []CssCustomProperty, inheritedRef Reference) []CssCustomProperty {
	// Build map for O(1) override detection
	classMap := make(map[string]CssCustomProperty)
	for _, prop := range classProps {
		classMap[prop.Name] = prop
	}

	// Result slice to preserve order
	result := make([]CssCustomProperty, 0, len(inheritedProps)+len(classProps))

	// 1. Add inherited properties that aren't overridden (preserves order from inheritedProps)
	for _, prop := range inheritedProps {
		// Set InheritedFrom if not already set
		if prop.InheritedFrom == nil && inheritedRef.Name != "" {
			// Create a copy of inheritedRef to avoid pointer aliasing
			ref := inheritedRef
			prop.InheritedFrom = &ref
		}

		// Only add if not overridden by class
		if _, overridden := classMap[prop.Name]; !overridden {
			result = append(result, prop)
		}
	}

	// 2. Add class properties (preserves order from classProps, overrides inherited)
	for _, prop := range classProps {
		// Check if we're overriding an inherited property
		if inheritedProp, exists := getCssPropertyFromSlice(inheritedProps, prop.Name); exists {
			// Merge summaries/descriptions with smart deduplication
			if prop.Summary == "" {
				prop.Summary = inheritedProp.Summary
			}

			// Description: avoid duplication
			if prop.Description == "" {
				prop.Description = inheritedProp.Description
			} else if prop.Description != inheritedProp.Description && inheritedProp.Description != "" {
				prop.Description = prop.Description + "\n\n" + inheritedProp.Description
			}
		}

		// Clear InheritedFrom for class properties
		prop.InheritedFrom = nil
		result = append(result, prop)
	}

	return result
}

// getCssPropertyFromSlice finds a CSS property in a slice by name (for description merging)
func getCssPropertyFromSlice(props []CssCustomProperty, name string) (CssCustomProperty, bool) {
	for _, prop := range props {
		if prop.Name == name {
			return prop, true
		}
	}
	return CssCustomProperty{}, false
}

// mergeCssParts combines CSS parts from class and mixins.
// Preserves source order: inherited parts first (not overridden), then class parts.
// Uses smart description deduplication like mergeAttributes.
func mergeCssParts(classParts []CssPart, inheritedParts []CssPart, inheritedRef Reference) []CssPart {
	// Build map for O(1) override detection
	classMap := make(map[string]CssPart)
	for _, part := range classParts {
		classMap[part.Name] = part
	}

	// Result slice to preserve order
	result := make([]CssPart, 0, len(inheritedParts)+len(classParts))

	// 1. Add inherited parts that aren't overridden (preserves order from inheritedParts)
	for _, part := range inheritedParts {
		// Set InheritedFrom if not already set
		if part.InheritedFrom == nil && inheritedRef.Name != "" {
			// Create a copy of inheritedRef to avoid pointer aliasing
			ref := inheritedRef
			part.InheritedFrom = &ref
		}

		// Only add if not overridden by class
		if _, overridden := classMap[part.Name]; !overridden {
			result = append(result, part)
		}
	}

	// 2. Add class parts (preserves order from classParts, overrides inherited)
	for _, part := range classParts {
		// Check if we're overriding an inherited part
		if inheritedPart, exists := getCssPartFromSlice(inheritedParts, part.Name); exists {
			// Merge summaries/descriptions with smart deduplication
			if part.Summary == "" {
				part.Summary = inheritedPart.Summary
			}

			// Description: avoid duplication
			if part.Description == "" {
				part.Description = inheritedPart.Description
			} else if part.Description != inheritedPart.Description && inheritedPart.Description != "" {
				part.Description = part.Description + "\n\n" + inheritedPart.Description
			}
		}

		// Clear InheritedFrom for class parts
		part.InheritedFrom = nil
		result = append(result, part)
	}

	return result
}

// getCssPartFromSlice finds a CSS part in a slice by name (for description merging)
func getCssPartFromSlice(parts []CssPart, name string) (CssPart, bool) {
	for _, part := range parts {
		if part.Name == name {
			return part, true
		}
	}
	return CssPart{}, false
}

// mergeCssStates combines CSS states from class and mixins.
// Preserves source order: inherited states first (not overridden), then class states.
// Uses smart description deduplication like mergeAttributes.
func mergeCssStates(classStates []CssCustomState, inheritedStates []CssCustomState, inheritedRef Reference) []CssCustomState {
	// Build map for O(1) override detection
	classMap := make(map[string]CssCustomState)
	for _, state := range classStates {
		classMap[state.Name] = state
	}

	// Result slice to preserve order
	result := make([]CssCustomState, 0, len(inheritedStates)+len(classStates))

	// 1. Add inherited states that aren't overridden (preserves order from inheritedStates)
	for _, state := range inheritedStates {
		// Set InheritedFrom if not already set
		if state.InheritedFrom == nil && inheritedRef.Name != "" {
			// Create a copy of inheritedRef to avoid pointer aliasing
			ref := inheritedRef
			state.InheritedFrom = &ref
		}

		// Only add if not overridden by class
		if _, overridden := classMap[state.Name]; !overridden {
			result = append(result, state)
		}
	}

	// 2. Add class states (preserves order from classStates, overrides inherited)
	for _, state := range classStates {
		// Check if we're overriding an inherited state
		if inheritedState, exists := getCssStateFromSlice(inheritedStates, state.Name); exists {
			// Merge summaries/descriptions with smart deduplication
			if state.Summary == "" {
				state.Summary = inheritedState.Summary
			}

			// Description: avoid duplication
			if state.Description == "" {
				state.Description = inheritedState.Description
			} else if state.Description != inheritedState.Description && inheritedState.Description != "" {
				state.Description = state.Description + "\n\n" + inheritedState.Description
			}
		}

		// Clear InheritedFrom for class states
		state.InheritedFrom = nil
		result = append(result, state)
	}

	return result
}

// getCssStateFromSlice finds a CSS state in a slice by name (for description merging)
func getCssStateFromSlice(states []CssCustomState, name string) (CssCustomState, bool) {
	for _, state := range states {
		if state.Name == name {
			return state, true
		}
	}
	return CssCustomState{}, false
}

// mergeClassMembers combines class members (fields/methods) from class and mixins.
// Members can be fields or methods, we merge by name.
// Preserves source order: inherited members first (not overridden), then class members.
func mergeClassMembers(classMembers []ClassMember, mixinMembers []ClassMember, mixinRef Reference) []ClassMember {
	// Build map for O(1) override detection
	classMap := make(map[string]ClassMember)
	for _, member := range classMembers {
		var key string
		switch m := member.(type) {
		case *ClassField:
			key = "field:" + m.Name
		case *ClassMethod:
			key = "method:" + m.Name
		default:
			continue
		}
		classMap[key] = member
	}

	// Result slice to preserve order
	result := make([]ClassMember, 0, len(mixinMembers)+len(classMembers))

	// 1. Add inherited members that aren't overridden (preserves order from mixinMembers)
	for _, member := range mixinMembers {
		var key string
		switch m := member.(type) {
		case *ClassField:
			key = "field:" + m.Name
			// Set InheritedFrom if not already set
			if m.InheritedFrom == nil && mixinRef.Name != "" {
				// Create a copy of mixinRef to avoid pointer aliasing
				ref := mixinRef
				m.InheritedFrom = &ref
			}
		case *ClassMethod:
			key = "method:" + m.Name
			if m.InheritedFrom == nil && mixinRef.Name != "" {
				// Create a copy of mixinRef to avoid pointer aliasing
				ref := mixinRef
				m.InheritedFrom = &ref
			}
		default:
			continue
		}

		// Only add if not overridden by class
		if _, overridden := classMap[key]; !overridden {
			result = append(result, member)
		}
	}

	// 2. Add class members (preserves order from classMembers, overrides inherited)
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

		// Check if we're overriding an inherited member
		if inheritedMember, exists := getMemberFromSlice(mixinMembers, key); exists {
			// Merge summaries/descriptions with smart deduplication
			switch m := member.(type) {
			case *ClassField:
				if existingField, ok := inheritedMember.(*ClassField); ok {
					if m.Summary == "" {
						m.Summary = existingField.Summary
					}
					// Description: avoid duplication
					if m.Description == "" {
						m.Description = existingField.Description
					} else if m.Description != existingField.Description && existingField.Description != "" {
						m.Description = m.Description + "\n\n" + existingField.Description
					}
				}
			case *ClassMethod:
				if existingMethod, ok := inheritedMember.(*ClassMethod); ok {
					if m.Summary == "" {
						m.Summary = existingMethod.Summary
					}
					// Description: avoid duplication
					if m.Description == "" {
						m.Description = existingMethod.Description
					} else if m.Description != existingMethod.Description && existingMethod.Description != "" {
						m.Description = m.Description + "\n\n" + existingMethod.Description
					}
				}
			}
		}

		result = append(result, member)
	}

	return result
}

// getMemberFromSlice finds a member in a slice by key (for description merging)
func getMemberFromSlice(members []ClassMember, key string) (ClassMember, bool) {
	for _, member := range members {
		var memberKey string
		switch m := member.(type) {
		case *ClassField:
			memberKey = "field:" + m.Name
		case *ClassMethod:
			memberKey = "method:" + m.Name
		default:
			continue
		}
		if memberKey == key {
			return member, true
		}
	}
	return nil, false
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

	once.Do(func() {
		// Defensive recovery to prevent panics from deadlocking other goroutines
		// If a panic occurs, we store an empty result rather than blocking forever
		defer func() {
			if r := recover(); r != nil {
				// Store empty result on panic to allow other goroutines to proceed
				flattenedMembersCache.Store(cacheKey, &flattenedMembers{})
			}
		}()

		fm := &flattenedMembers{}

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

		for _, superclass := range superclasses {
			// Construct a reference to the current superclass for InheritedFrom tracking
			// Note: resolveSuperclass returns [base → derived], so we construct the reference
			// from the current superclass in the chain, not from ced.Superclass
			superclassRef := Reference{
				Name: superclass.Name,
			}
			if superclass.Module != nil {
				superclassRef.Module = superclass.Module.Path
			}

			// Check if this is a CustomElementDeclaration (has attributes, slots, etc.)
			// We need to find the full declaration to access CustomElement members
			if superclassRef.Name != "" {
				superclassDecl := pkg.FindDeclaration(superclassRef)
				if cedSuper, ok := superclassDecl.(*CustomElementDeclaration); ok {
					// CustomElement superclass - append inherited members with InheritedFrom set
					// We don't use merge here because merging multiple inherited sources would
					// incorrectly clear InheritedFrom fields
					for _, attr := range cedSuper.OwnAttributes() {
						if attr.InheritedFrom == nil {
							ref := superclassRef
							attr.InheritedFrom = &ref
						}
						attrs = append(attrs, attr)
					}
					for _, slot := range cedSuper.OwnSlots() {
						if slot.InheritedFrom == nil {
							ref := superclassRef
							slot.InheritedFrom = &ref
						}
						slots = append(slots, slot)
					}
					for _, event := range cedSuper.OwnEvents() {
						if event.InheritedFrom == nil {
							ref := superclassRef
							event.InheritedFrom = &ref
						}
						events = append(events, event)
					}
					for _, prop := range cedSuper.OwnCssProperties() {
						if prop.InheritedFrom == nil {
							ref := superclassRef
							prop.InheritedFrom = &ref
						}
						cssProps = append(cssProps, prop)
					}
					for _, part := range cedSuper.OwnCssParts() {
						if part.InheritedFrom == nil {
							ref := superclassRef
							part.InheritedFrom = &ref
						}
						cssParts = append(cssParts, part)
					}
					for _, state := range cedSuper.OwnCssStates() {
						if state.InheritedFrom == nil {
							ref := superclassRef
							state.InheritedFrom = &ref
						}
						cssStates = append(cssStates, state)
					}
				}
			}

			// Append superclass members (all superclasses contribute via ClassLike.Members)
			for _, member := range superclass.Members {
				switch m := member.(type) {
				case *ClassField:
					if m.InheritedFrom == nil {
						ref := superclassRef
						m.InheritedFrom = &ref
					}
					fields = append(fields, m)
				case *ClassMethod:
					if m.InheritedFrom == nil {
						ref := superclassRef
						m.InheritedFrom = &ref
					}
					fields = append(fields, m)
				}
			}
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
				mixinName = m.Name()
				mixinClassLike = &m.ClassLike
			case *CustomElementMixinDeclaration:
				mixinName = m.Name()
				mixinClassLike = &m.ClassLike
			}

			// Create reference for InheritedFrom
			mixinRef := Reference{Name: mixinName}

			// Find module path for this mixin
			// Performance: O(modules × declarations) linear search
			// Acceptable for typical projects (< 100 modules, < 50 declarations/module)
			// Could be optimized with a reverse lookup map if profiling shows bottleneck
			for i := range pkg.Modules {
				if slices.Contains(pkg.Modules[i].Declarations, mixinDecl) {
					mixinRef.Module = pkg.Modules[i].Path
					break
				}
			}

			// Append mixin members
			// For CustomElementMixin, append custom element members (attributes, slots, etc.)
			if cemMixin, ok := mixinDecl.(*CustomElementMixinDeclaration); ok {
				// CustomElementMixin has attributes, slots, events, CSS
				// Use Own* methods to get direct members (not flattened)
				for _, attr := range cemMixin.OwnAttributes() {
					if attr.InheritedFrom == nil {
						ref := mixinRef
						attr.InheritedFrom = &ref
					}
					attrs = append(attrs, attr)
				}
				for _, slot := range cemMixin.OwnSlots() {
					if slot.InheritedFrom == nil {
						ref := mixinRef
						slot.InheritedFrom = &ref
					}
					slots = append(slots, slot)
				}
				for _, event := range cemMixin.OwnEvents() {
					if event.InheritedFrom == nil {
						ref := mixinRef
						event.InheritedFrom = &ref
					}
					events = append(events, event)
				}
				for _, prop := range cemMixin.OwnCssProperties() {
					if prop.InheritedFrom == nil {
						ref := mixinRef
						prop.InheritedFrom = &ref
					}
					cssProps = append(cssProps, prop)
				}
				for _, part := range cemMixin.OwnCssParts() {
					if part.InheritedFrom == nil {
						ref := mixinRef
						part.InheritedFrom = &ref
					}
					cssParts = append(cssParts, part)
				}
				for _, state := range cemMixin.OwnCssStates() {
					if state.InheritedFrom == nil {
						ref := mixinRef
						state.InheritedFrom = &ref
					}
					cssStates = append(cssStates, state)
				}
			}

			// All mixins have class members (from ClassLike)
			if mixinClassLike != nil {
				for _, member := range mixinClassLike.Members {
					switch m := member.(type) {
					case *ClassField:
						if m.InheritedFrom == nil {
							ref := mixinRef
							m.InheritedFrom = &ref
						}
						fields = append(fields, m)
					case *ClassMethod:
						if m.InheritedFrom == nil {
							ref := mixinRef
							m.InheritedFrom = &ref
						}
						fields = append(fields, m)
					}
				}
			}
		}

		// Finally merge with class's own members
		// Class members override mixin members
		fm.attributes = mergeAttributes(ced.OwnAttributes(), attrs, Reference{})
		fm.slots = mergeSlots(ced.OwnSlots(), slots, Reference{})
		fm.events = mergeEvents(ced.OwnEvents(), events, Reference{})
		fm.cssProperties = mergeCssProperties(ced.OwnCssProperties(), cssProps, Reference{})
		fm.cssParts = mergeCssParts(ced.OwnCssParts(), cssParts, Reference{})
		fm.cssStates = mergeCssStates(ced.OwnCssStates(), cssStates, Reference{})

		// Merge all members (fields + methods) in source order
		allMembers := mergeClassMembers(ced.Members, fields, Reference{})

		// Separate fields and methods while preserving source order
		fm.fields = make([]ClassMember, 0)
		fm.methods = make([]ClassMember, 0)
		for _, member := range allMembers {
			switch member.(type) {
			case *ClassField:
				fm.fields = append(fm.fields, member)
			case *ClassMethod:
				fm.methods = append(fm.methods, member)
			}
		}

		// All members preserve source order (inherited first, then class's own)
		// Order within each category follows: base → derived → class (inheritance application order)

		// Store in cache
		flattenedMembersCache.Store(cacheKey, fm)
	})

	// Return cached result
	cached, _ := flattenedMembersCache.Load(cacheKey)
	return cached.(*flattenedMembers)
}
