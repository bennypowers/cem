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
package ephemeral

import (
	"sync"

	lspTypes "bennypowers.dev/cem/lsp/types"
	M "bennypowers.dev/cem/manifest"
)

// ephemeralElementDefinition implements types.ElementDefinition for ephemeral elements
type ephemeralElementDefinition struct {
	element    *M.CustomElement
	modulePath string
}

func (d *ephemeralElementDefinition) ModulePath() string        { return d.modulePath }
func (d *ephemeralElementDefinition) PackageName() string       { return "" }
func (d *ephemeralElementDefinition) SourceHref() string        { return "" }
func (d *ephemeralElementDefinition) Element() *M.CustomElement { return d.element }

// ephemeralEntry tracks a declaration and the URI it came from
type ephemeralEntry struct {
	uri  string
	decl *M.CustomElementDeclaration
}

// Registry stores synthesized CustomElementDeclaration objects from open
// documents. It acts as a fallback for the main manifest registry, providing
// hover, completion, and other LSP features for elements defined in the
// current file that aren't part of any CEM manifest.
type Registry struct {
	mu    sync.RWMutex
	byURI map[string]*M.Package      // document URI → synthesized package
	byTag map[string]*ephemeralEntry // tag name → entry (flat index)
}

// NewRegistry creates a new ephemeral registry.
func NewRegistry() *Registry {
	return &Registry{
		byURI: make(map[string]*M.Package),
		byTag: make(map[string]*ephemeralEntry),
	}
}

// Update atomically replaces all ephemeral data for a document URI.
// It rebuilds the flat tag-name index after the update.
func (r *Registry) Update(uri string, pkg *M.Package) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.byURI[uri] = pkg
	r.rebuildIndex()
}

// Remove removes all ephemeral data for a document URI.
func (r *Registry) Remove(uri string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.byURI, uri)
	r.rebuildIndex()
}

// rebuildIndex rebuilds the flat tag-name → entry index from all stored packages.
// If multiple URIs define the same tag name, the last one visited wins (Go map
// iteration is non-deterministic). In practice this is unlikely — locally-defined
// elements are typically unique per file — and the main registry takes precedence
// over ephemeral data regardless.
// Must be called with r.mu held.
func (r *Registry) rebuildIndex() {
	r.byTag = make(map[string]*ephemeralEntry)
	for uri, pkg := range r.byURI {
		for i := range pkg.Modules {
			for j := range pkg.Modules[i].Declarations {
				if ce, ok := pkg.Modules[i].Declarations[j].(*M.CustomElementDeclaration); ok {
					if ce.TagName != "" {
						r.byTag[ce.TagName] = &ephemeralEntry{
							uri:  uri,
							decl: ce,
						}
					}
				}
			}
		}
	}
}

// FindCustomElementDeclaration returns the full declaration for a tag name,
// or nil if not found in the ephemeral registry.
func (r *Registry) FindCustomElementDeclaration(tagName string) *M.CustomElementDeclaration {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if entry, ok := r.byTag[tagName]; ok {
		return entry.decl
	}
	return nil
}

// Element returns the CustomElement for a tag name.
func (r *Registry) Element(tagName string) (*M.CustomElement, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if entry, ok := r.byTag[tagName]; ok {
		return &entry.decl.CustomElement, true
	}
	return nil, false
}

// Attributes returns the attribute map for a tag name.
func (r *Registry) Attributes(tagName string) (map[string]*M.Attribute, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	entry, ok := r.byTag[tagName]
	if !ok {
		return nil, false
	}

	attrs := entry.decl.Attributes()
	if len(attrs) == 0 {
		return nil, false
	}

	attrMap := make(map[string]*M.Attribute, len(attrs))
	for i := range attrs {
		attrMap[attrs[i].Name] = &attrs[i]
	}
	return attrMap, true
}

// Slots returns the slots for a tag name.
func (r *Registry) Slots(tagName string) ([]M.Slot, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if entry, ok := r.byTag[tagName]; ok {
		slots := entry.decl.CustomElement.Slots
		return slots, len(slots) > 0
	}
	return nil, false
}

// AllTagNames returns all tag names known to the ephemeral registry.
func (r *Registry) AllTagNames() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	tags := make([]string, 0, len(r.byTag))
	for tag := range r.byTag {
		tags = append(tags, tag)
	}
	return tags
}

// ElementDefinition returns an ElementDefinition for a tag name.
// The module path is the document URI where the element is defined.
func (r *Registry) ElementDefinition(tagName string) (lspTypes.ElementDefinition, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	entry, ok := r.byTag[tagName]
	if !ok {
		return nil, false
	}

	return &ephemeralElementDefinition{
		element:    &entry.decl.CustomElement,
		modulePath: entry.uri,
	}, true
}

// ElementDescription returns the description or summary for a tag name.
func (r *Registry) ElementDescription(tagName string) (string, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	entry, ok := r.byTag[tagName]
	if !ok {
		return "", false
	}

	if entry.decl.Description != "" {
		return entry.decl.Description, true
	}
	if entry.decl.Summary != "" {
		return entry.decl.Summary, true
	}
	return "", false
}

// ElementSource returns the source specifier for a tag name.
// For ephemeral elements, this is the document URI.
func (r *Registry) ElementSource(tagName string) (string, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	entry, ok := r.byTag[tagName]
	if !ok {
		return "", false
	}

	return entry.uri, true
}

// Has returns true if the tag name exists in the ephemeral registry.
func (r *Registry) Has(tagName string) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	_, ok := r.byTag[tagName]
	return ok
}
