/*
Copyright © 2025 Benny Powers <web@bennypowers.com>

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
package relationships

import (
	"testing"

	"bennypowers.dev/cem/manifest"
)

func TestDetector_Superclass(t *testing.T) {
	d := NewDetector()

	// Base element
	d.AddElement(ElementData{
		TagName:   "pf-base",
		ClassName: "PfBase",
	})

	// Derived element
	d.AddElement(ElementData{
		TagName:   "pf-button",
		ClassName: "PfButton",
		Superclass: &manifest.Reference{
			Name: "PfBase",
		},
	})

	// Check relationships from derived element
	rels := d.DetectRelationships("pf-button")
	found := false
	for _, r := range rels {
		if r.TargetTagName == "pf-base" && r.Type == Superclass {
			found = true
			if r.Via != "PfBase" {
				t.Errorf("expected Via to be 'PfBase', got %q", r.Via)
			}
		}
	}
	if !found {
		t.Error("expected superclass relationship to pf-base")
	}

	// Check relationships from base element (should show subclass)
	rels = d.DetectRelationships("pf-base")
	found = false
	for _, r := range rels {
		if r.TargetTagName == "pf-button" && r.Type == Subclass {
			found = true
		}
	}
	if !found {
		t.Error("expected subclass relationship to pf-button")
	}
}

func TestDetector_SharedMixins(t *testing.T) {
	d := NewDetector()

	// Two elements sharing a mixin
	d.AddElement(ElementData{
		TagName:   "pf-button",
		ClassName: "PfButton",
		Mixins: []manifest.Reference{
			{Name: "InteractiveMixin"},
		},
	})

	d.AddElement(ElementData{
		TagName:   "pf-link",
		ClassName: "PfLink",
		Mixins: []manifest.Reference{
			{Name: "InteractiveMixin"},
		},
	})

	// Check that pf-button sees pf-link as sharing InteractiveMixin
	rels := d.DetectRelationships("pf-button")
	found := false
	for _, r := range rels {
		if r.TargetTagName == "pf-link" && r.Type == Mixin && r.Via == "InteractiveMixin" {
			found = true
		}
	}
	if !found {
		t.Error("expected mixin relationship between pf-button and pf-link")
	}
}

func TestDetector_ModuleCoLocation(t *testing.T) {
	d := NewDetector()

	// Two elements in the same module
	d.AddElement(ElementData{
		TagName:    "pf-button",
		ClassName:  "PfButton",
		ModulePath: "elements/button.js",
	})

	d.AddElement(ElementData{
		TagName:    "pf-button-group",
		ClassName:  "PfButtonGroup",
		ModulePath: "elements/button.js",
	})

	rels := d.DetectRelationships("pf-button")
	found := false
	for _, r := range rels {
		if r.TargetTagName == "pf-button-group" && r.Type == Module {
			found = true
		}
	}
	if !found {
		t.Error("expected module relationship between pf-button and pf-button-group")
	}
}

func TestDetector_PackageSiblings(t *testing.T) {
	d := NewDetector()

	// Two elements in same package but different modules
	d.AddElement(ElementData{
		TagName:     "pf-button",
		ClassName:   "PfButton",
		ModulePath:  "elements/button.js",
		PackageName: "@patternfly/elements",
	})

	d.AddElement(ElementData{
		TagName:     "pf-icon",
		ClassName:   "PfIcon",
		ModulePath:  "elements/icon.js",
		PackageName: "@patternfly/elements",
	})

	rels := d.DetectRelationships("pf-button")
	found := false
	for _, r := range rels {
		if r.TargetTagName == "pf-icon" && r.Type == Package {
			found = true
		}
	}
	if !found {
		t.Error("expected package relationship between pf-button and pf-icon")
	}
}

func TestDetector_RelationshipPriority(t *testing.T) {
	d := NewDetector()

	// Elements with multiple relationship types
	d.AddElement(ElementData{
		TagName:     "pf-base",
		ClassName:   "PfBase",
		ModulePath:  "elements/base.js",
		PackageName: "@patternfly/elements",
	})

	d.AddElement(ElementData{
		TagName:     "pf-button",
		ClassName:   "PfButton",
		ModulePath:  "elements/base.js", // Same module as base
		PackageName: "@patternfly/elements",
		Superclass:  &manifest.Reference{Name: "PfBase"},
	})

	rels := d.DetectRelationships("pf-button")

	// Should have superclass relationship
	hasSuperclass := false
	hasModule := false
	for _, r := range rels {
		if r.TargetTagName == "pf-base" {
			if r.Type == Superclass {
				hasSuperclass = true
			}
			if r.Type == Module {
				hasModule = true
			}
		}
	}

	if !hasSuperclass {
		t.Error("expected superclass relationship")
	}
	// Module relationship should be skipped since superclass is a stronger relationship
	if hasModule {
		t.Error("module relationship should be skipped when superclass exists")
	}
}

func TestDetector_NoSelfRelationship(t *testing.T) {
	d := NewDetector()

	d.AddElement(ElementData{
		TagName:     "pf-button",
		ClassName:   "PfButton",
		ModulePath:  "elements/button.js",
		PackageName: "@patternfly/elements",
	})

	rels := d.DetectRelationships("pf-button")

	for _, r := range rels {
		if r.TargetTagName == "pf-button" {
			t.Error("element should not have relationship to itself")
		}
	}
}

func TestDetector_UnknownElement(t *testing.T) {
	d := NewDetector()

	rels := d.DetectRelationships("unknown-element")

	if len(rels) != 0 {
		t.Errorf("expected no relationships for unknown element, got %v", rels)
	}
}

func TestRelationship_Label(t *testing.T) {
	tests := []struct {
		rel      Relationship
		expected string
	}{
		{Relationship{Type: Superclass, Via: "BaseClass"}, "extends BaseClass"},
		{Relationship{Type: Subclass, TargetTagName: "child-element"}, "extended by child-element"},
		{Relationship{Type: Mixin, Via: "InteractiveMixin"}, "shares InteractiveMixin"},
		{Relationship{Type: Module}, "same module"},
		{Relationship{Type: Package}, "same package"},
	}

	for _, tt := range tests {
		t.Run(string(tt.rel.Type), func(t *testing.T) {
			if got := tt.rel.Label(); got != tt.expected {
				t.Errorf("got %q, want %q", got, tt.expected)
			}
		})
	}
}
