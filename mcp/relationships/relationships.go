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

// Package relationships provides element relationship detection for MCP context.
// It detects relationships based on JavaScript-derived type information from manifests.
package relationships

import (
	"bennypowers.dev/cem/manifest"
)

// Type indicates the relationship between elements.
type Type string

const (
	// Superclass indicates inheritance relationship (element inherits from target).
	Superclass Type = "superclass"
	// Subclass indicates inheritance relationship (target inherits from element).
	Subclass Type = "subclass"
	// Mixin indicates elements share a mixin.
	Mixin Type = "mixin"
	// Module indicates elements are defined in the same module.
	Module Type = "module"
	// Package indicates elements are from the same package.
	Package Type = "package"
)

// Relationship represents a relationship between two elements.
type Relationship struct {
	// TargetTagName is the tag name of the related element.
	TargetTagName string `json:"target"`
	// Type is the kind of relationship.
	Type Type `json:"type"`
	// Via describes how the relationship was determined (e.g., mixin name, module path).
	Via string `json:"via,omitempty"`
}

// ElementData contains the information needed to detect relationships for an element.
type ElementData struct {
	TagName     string
	ClassName   string
	Superclass  *manifest.Reference
	Mixins      []manifest.Reference
	ModulePath  string
	PackageName string
}

// Detector detects relationships between elements.
type Detector struct {
	// elements maps tag names to their ElementData
	elements map[string]ElementData
	// classToTag maps class names to tag names
	classToTag map[string]string
	// moduleToTags maps module paths to tag names
	moduleToTags map[string][]string
	// packageToTags maps package names to tag names
	packageToTags map[string][]string
	// mixinToTags maps mixin names to tag names that use them
	mixinToTags map[string][]string
}

// NewDetector creates a new relationship detector.
func NewDetector() *Detector {
	return &Detector{
		elements:      make(map[string]ElementData),
		classToTag:    make(map[string]string),
		moduleToTags:  make(map[string][]string),
		packageToTags: make(map[string][]string),
		mixinToTags:   make(map[string][]string),
	}
}

// AddElement adds an element to the detector for relationship analysis.
func (d *Detector) AddElement(data ElementData) {
	d.elements[data.TagName] = data

	// Index by class name
	if data.ClassName != "" {
		d.classToTag[data.ClassName] = data.TagName
	}

	// Index by module
	if data.ModulePath != "" {
		d.moduleToTags[data.ModulePath] = append(d.moduleToTags[data.ModulePath], data.TagName)
	}

	// Index by package
	if data.PackageName != "" {
		d.packageToTags[data.PackageName] = append(d.packageToTags[data.PackageName], data.TagName)
	}

	// Index by mixins
	for _, mixin := range data.Mixins {
		mixinName := mixin.Name
		d.mixinToTags[mixinName] = append(d.mixinToTags[mixinName], data.TagName)
	}
}

// DetectRelationships returns all relationships for the given element.
func (d *Detector) DetectRelationships(tagName string) []Relationship {
	data, ok := d.elements[tagName]
	if !ok {
		return nil
	}

	var rels []Relationship

	// Check superclass relationship
	if data.Superclass != nil && data.Superclass.Name != "" {
		if targetTag, ok := d.classToTag[data.Superclass.Name]; ok && targetTag != tagName {
			rels = append(rels, Relationship{
				TargetTagName: targetTag,
				Type:          Superclass,
				Via:           data.Superclass.Name,
			})
		}
	}

	// Check for elements that inherit from this element (subclass)
	for _, otherData := range d.elements {
		if otherData.TagName == tagName {
			continue
		}
		if otherData.Superclass != nil && otherData.Superclass.Name == data.ClassName {
			rels = append(rels, Relationship{
				TargetTagName: otherData.TagName,
				Type:          Subclass,
				Via:           data.ClassName,
			})
		}
	}

	// Check shared mixins
	for _, mixin := range data.Mixins {
		mixinName := mixin.Name
		for _, otherTag := range d.mixinToTags[mixinName] {
			if otherTag == tagName {
				continue
			}
			// Avoid duplicates
			if !containsRelationship(rels, otherTag, Mixin) {
				rels = append(rels, Relationship{
					TargetTagName: otherTag,
					Type:          Mixin,
					Via:           mixinName,
				})
			}
		}
	}

	// Check module co-location
	if data.ModulePath != "" {
		for _, otherTag := range d.moduleToTags[data.ModulePath] {
			if otherTag == tagName {
				continue
			}
			// Skip if already related through inheritance or mixin
			if containsRelationship(rels, otherTag, Superclass) ||
				containsRelationship(rels, otherTag, Subclass) ||
				containsRelationship(rels, otherTag, Mixin) {
				continue
			}
			rels = append(rels, Relationship{
				TargetTagName: otherTag,
				Type:          Module,
				Via:           data.ModulePath,
			})
		}
	}

	// Check package siblings (only if no closer relationship exists)
	if data.PackageName != "" {
		for _, otherTag := range d.packageToTags[data.PackageName] {
			if otherTag == tagName {
				continue
			}
			// Skip if already related through any other relationship type
			if containsTarget(rels, otherTag) {
				continue
			}
			rels = append(rels, Relationship{
				TargetTagName: otherTag,
				Type:          Package,
				Via:           data.PackageName,
			})
		}
	}

	return rels
}

// containsRelationship checks if relationships already contain a specific target and type.
func containsRelationship(rels []Relationship, target string, relType Type) bool {
	for _, r := range rels {
		if r.TargetTagName == target && r.Type == relType {
			return true
		}
	}
	return false
}

// containsTarget checks if relationships already contain a specific target.
func containsTarget(rels []Relationship, target string) bool {
	for _, r := range rels {
		if r.TargetTagName == target {
			return true
		}
	}
	return false
}

// RelationshipLabel returns a human-readable label for the relationship type.
func (r Relationship) Label() string {
	switch r.Type {
	case Superclass:
		return "extends " + r.Via
	case Subclass:
		return "extended by " + r.TargetTagName
	case Mixin:
		return "shares " + r.Via
	case Module:
		return "same module"
	case Package:
		return "same package"
	default:
		return string(r.Type)
	}
}
