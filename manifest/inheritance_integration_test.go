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
package manifest_test

import (
	"testing"

	"bennypowers.dev/cem/internal/platform/testutil"
	"bennypowers.dev/cem/manifest"
)

// TestInheritance_SimpleMixin tests basic mixin flattening
func TestInheritance_SimpleMixin(t *testing.T) {
	var pkg manifest.Package
	testutil.LoadJSONFixture(t, "inheritance_simple_mixin.json", &pkg)

	// Get the custom element that uses the mixin
	ced := pkg.Modules[1].Declarations[0].(*manifest.CustomElementDeclaration)

	// Test that flattened attributes include both class and mixin attributes
	attrs := ced.Attributes()
	if len(attrs) != 2 {
		t.Fatalf("Expected 2 flattened attributes, got %d", len(attrs))
	}

	// Check that mixin attribute has InheritedFrom set
	var disabledAttr *manifest.Attribute
	var labelAttr *manifest.Attribute
	for i := range attrs {
		switch attrs[i].Name {
		case "disabled":
			disabledAttr = &attrs[i]
		case "label":
			labelAttr = &attrs[i]
		}
	}

	if disabledAttr == nil {
		t.Fatal("disabled attribute not found in flattened attributes")
	}
	if disabledAttr.InheritedFrom == nil {
		t.Error("disabled attribute should have InheritedFrom set")
	} else if disabledAttr.InheritedFrom.Name != "DisabledMixin" {
		t.Errorf("disabled.InheritedFrom.Name = %s, want DisabledMixin", disabledAttr.InheritedFrom.Name)
	}

	if labelAttr == nil {
		t.Fatal("label attribute not found in flattened attributes")
	}
	if labelAttr.InheritedFrom != nil {
		t.Error("label attribute should not have InheritedFrom set (class's own attribute)")
	}
}

// TestInheritance_Superclass tests basic superclass inheritance
func TestInheritance_Superclass(t *testing.T) {
	var pkg manifest.Package
	testutil.LoadJSONFixture(t, "inheritance_superclass.json", &pkg)

	// Get the subclass
	ced := pkg.Modules[1].Declarations[0].(*manifest.CustomElementDeclaration)

	// Test that flattened attributes include both class and superclass attributes
	attrs := ced.Attributes()
	if len(attrs) != 2 {
		t.Fatalf("Expected 2 flattened attributes, got %d", len(attrs))
	}

	// Check that superclass attribute has InheritedFrom set
	var baseAttr *manifest.Attribute
	var childAttr *manifest.Attribute
	for i := range attrs {
		switch attrs[i].Name {
		case "base-prop":
			baseAttr = &attrs[i]
		case "child-prop":
			childAttr = &attrs[i]
		}
	}

	if baseAttr == nil {
		t.Fatal("base-prop attribute not found in flattened attributes")
	}
	if baseAttr.InheritedFrom == nil {
		t.Error("base-prop attribute should have InheritedFrom set")
	} else if baseAttr.InheritedFrom.Name != "BaseElement" {
		t.Errorf("base-prop.InheritedFrom.Name = %s, want BaseElement", baseAttr.InheritedFrom.Name)
	}

	if childAttr == nil {
		t.Fatal("child-prop attribute not found in flattened attributes")
	}
	if childAttr.InheritedFrom != nil {
		t.Error("child-prop attribute should not have InheritedFrom set (class's own attribute)")
	}
}

// TestInheritance_SuperclassAndMixins tests combined superclass + mixin inheritance
func TestInheritance_SuperclassAndMixins(t *testing.T) {
	var pkg manifest.Package
	testutil.LoadJSONFixture(t, "inheritance_superclass_and_mixins.json", &pkg)

	// Get the element that has both superclass and mixins
	ced := pkg.Modules[2].Declarations[0].(*manifest.CustomElementDeclaration)

	// Test that flattened attributes include superclass, mixin, and class attributes
	attrs := ced.Attributes()
	if len(attrs) != 3 {
		t.Fatalf("Expected 3 flattened attributes (superclass + mixin + class), got %d", len(attrs))
	}

	// Verify each attribute's InheritedFrom
	attrMap := make(map[string]*manifest.Attribute)
	for i := range attrs {
		attrMap[attrs[i].Name] = &attrs[i]
	}

	// Superclass attribute
	if attr, ok := attrMap["base-prop"]; !ok {
		t.Error("base-prop not found")
	} else if attr.InheritedFrom == nil {
		t.Error("base-prop should have InheritedFrom set")
	} else if attr.InheritedFrom.Name != "BaseElement" {
		t.Errorf("base-prop.InheritedFrom.Name = %s, want BaseElement", attr.InheritedFrom.Name)
	}

	// Mixin attribute
	if attr, ok := attrMap["disabled"]; !ok {
		t.Error("disabled not found")
	} else if attr.InheritedFrom == nil {
		t.Error("disabled should have InheritedFrom set")
	} else if attr.InheritedFrom.Name != "DisabledMixin" {
		t.Errorf("disabled.InheritedFrom.Name = %s, want DisabledMixin", attr.InheritedFrom.Name)
	}

	// Class's own attribute
	if attr, ok := attrMap["child-prop"]; !ok {
		t.Error("child-prop not found")
	} else if attr.InheritedFrom != nil {
		t.Error("child-prop should not have InheritedFrom set")
	}
}

// TestInheritance_ChainedMixins tests mixin A → B → C chain
func TestInheritance_ChainedMixins(t *testing.T) {
	var pkg manifest.Package
	testutil.LoadJSONFixture(t, "inheritance_chained_mixins.json", &pkg)

	// Get the element that uses a mixin which itself uses another mixin
	ced := pkg.Modules[2].Declarations[0].(*manifest.CustomElementDeclaration)

	// Test that all attributes from the mixin chain are included
	attrs := ced.Attributes()
	if len(attrs) < 2 {
		t.Fatalf("Expected at least 2 flattened attributes from mixin chain, got %d", len(attrs))
	}

	// Verify attributes from different levels of the chain
	attrMap := make(map[string]*manifest.Attribute)
	for i := range attrs {
		attrMap[attrs[i].Name] = &attrs[i]
	}

	// Verify inherited attributes have InheritedFrom set
	if attr, ok := attrMap["base-attr"]; ok {
		if attr.InheritedFrom == nil {
			t.Error("base-attr should have InheritedFrom set")
		}
	} else {
		t.Error("base-attr not found in chain")
	}

	if attr, ok := attrMap["middle-attr"]; ok {
		if attr.InheritedFrom == nil {
			t.Error("middle-attr should have InheritedFrom set")
		}
	} else {
		t.Error("middle-attr not found in chain")
	}

	// Verify class's own attribute does NOT have InheritedFrom
	if attr, ok := attrMap["element-attr"]; ok {
		if attr.InheritedFrom != nil {
			t.Error("element-attr should not have InheritedFrom set (class's own attribute)")
		}
	} else {
		t.Error("element-attr not found")
	}
}

// TestInheritance_DescriptionDeduplication tests smart description merging
func TestInheritance_DescriptionDeduplication(t *testing.T) {
	var pkg manifest.Package
	testutil.LoadJSONFixture(t, "inheritance_description_dedup.json", &pkg)

	ced := pkg.Modules[1].Declarations[0].(*manifest.CustomElementDeclaration)
	attrs := ced.Attributes()

	attrMap := make(map[string]*manifest.Attribute)
	for i := range attrs {
		attrMap[attrs[i].Name] = &attrs[i]
	}

	// Test case 1: Same description - should not duplicate
	if attr, ok := attrMap["same-desc"]; ok {
		// Description should appear only once (not "Same description\n\nSame description")
		if attr.Description != "Same description" {
			t.Errorf("same-desc description should not be duplicated, got: %q", attr.Description)
		}
	}

	// Test case 2: Different description - should concatenate
	if attr, ok := attrMap["different-desc"]; ok {
		// Should have both descriptions concatenated
		if attr.Description != "Child description\n\nParent description" {
			t.Errorf("different-desc should have concatenated descriptions, got: %q", attr.Description)
		}
	}

	// Test case 3: Empty child description - should use parent description
	if attr, ok := attrMap["empty-child-desc"]; ok {
		if attr.Description != "Parent description" {
			t.Errorf("empty-child-desc should use parent description, got: %q", attr.Description)
		}
	}
}

// TestInheritance_MixinOverride tests class overriding mixin attribute
func TestInheritance_MixinOverride(t *testing.T) {
	var pkg manifest.Package
	testutil.LoadJSONFixture(t, "inheritance_mixin_override.json", &pkg)

	ced := pkg.Modules[1].Declarations[0].(*manifest.CustomElementDeclaration)
	attrs := ced.Attributes()

	attrMap := make(map[string]*manifest.Attribute)
	for i := range attrs {
		attrMap[attrs[i].Name] = &attrs[i]
	}

	// When class overrides a mixin attribute, InheritedFrom should be nil
	if attr, ok := attrMap["disabled"]; ok {
		if attr.InheritedFrom != nil {
			t.Error("Overridden attribute should not have InheritedFrom set")
		}
		// Class's values should take precedence
		if attr.Type == nil || attr.Type.Text != "string" {
			t.Error("Overridden attribute should have class's type (string), not mixin's type (boolean)")
		}
	} else {
		t.Fatal("disabled attribute not found")
	}
}

// TestInheritance_CircularReference tests circular reference detection
func TestInheritance_CircularReference(t *testing.T) {
	var pkg manifest.Package
	testutil.LoadJSONFixture(t, "inheritance_circular.json", &pkg)

	ced := pkg.Modules[0].Declarations[0].(*manifest.CustomElementDeclaration)

	// Should not panic or infinite loop
	attrs := ced.Attributes()

	// Should still get attributes (circular ref is skipped)
	if len(attrs) == 0 {
		t.Error("Should still get attributes despite circular reference")
	}
}

// TestInheritance_SlotsAndEvents tests inheritance of slots and events
func TestInheritance_SlotsAndEvents(t *testing.T) {
	var pkg manifest.Package
	testutil.LoadJSONFixture(t, "inheritance_slots_events.json", &pkg)

	ced := pkg.Modules[1].Declarations[0].(*manifest.CustomElementDeclaration)

	// Test slots
	slots := ced.Slots()
	if len(slots) != 2 {
		t.Fatalf("Expected 2 flattened slots, got %d", len(slots))
	}

	var inheritedSlot *manifest.Slot
	for i := range slots {
		if slots[i].InheritedFrom != nil {
			inheritedSlot = &slots[i]
			break
		}
	}
	if inheritedSlot == nil {
		t.Error("Expected at least one inherited slot")
	}

	// Test events
	events := ced.Events()
	if len(events) != 2 {
		t.Fatalf("Expected 2 flattened events, got %d", len(events))
	}

	var inheritedEvent *manifest.Event
	for i := range events {
		if events[i].InheritedFrom != nil {
			inheritedEvent = &events[i]
			break
		}
	}
	if inheritedEvent == nil {
		t.Error("Expected at least one inherited event")
	}
}
