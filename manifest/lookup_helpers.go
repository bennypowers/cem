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

import "sync"

// Package-level caches for thread-safe lazy initialization.
// These are kept external to manifest structs to avoid affecting their semantic structure.
var (
	// Cache for attribute field maps: map[*CustomElementDeclaration]map[string]*CustomElementField
	attributeFieldCache = &sync.Map{}
	attributeFieldOnce  = &sync.Map{} // map[*CustomElementDeclaration]*sync.Once

	// Cache for export maps: map[*Module]exportMaps
	exportMapsCache = &sync.Map{}
	exportMapsOnce  = &sync.Map{} // map[*Module]*sync.Once
)

// exportMaps holds both CustomElementExport and JavaScriptExport maps
type exportMaps struct {
	customElementExports map[string]*CustomElementExport
	jsExports            map[string]*JavaScriptExport
}

// BuildAttributeFieldMap creates a map from attribute name to CustomElementField
// for fast O(1) lookups. Returns an empty map (not nil) if no fields found.
//
// This function is exported for testing purposes.
func BuildAttributeFieldMap(members []ClassMember) map[string]*CustomElementField {
	m := make(map[string]*CustomElementField)
	for _, member := range members {
		if cef, ok := member.(*CustomElementField); ok {
			if cef.Attribute != "" {
				m[cef.Attribute] = cef
			}
		}
	}
	return m
}

// getAttributeFieldMap returns the cached map or builds it if needed.
// This method uses lazy initialization with sync.Once for thread-safe concurrent access.
// Both the cache and sync.Once are stored externally to avoid adding non-semantic fields to the struct.
func (ced *CustomElementDeclaration) getAttributeFieldMap() map[string]*CustomElementField {
	// Check if already cached
	if cached, ok := attributeFieldCache.Load(ced); ok {
		return cached.(map[string]*CustomElementField)
	}

	// Get or create a sync.Once for this instance
	once, _ := attributeFieldOnce.LoadOrStore(ced, &sync.Once{})

	// Use the sync.Once to ensure single initialization
	once.(*sync.Once).Do(func() {
		m := BuildAttributeFieldMap(ced.Members)
		attributeFieldCache.Store(ced, m)
	})

	// Load the now-initialized cache
	cached, _ := attributeFieldCache.Load(ced)
	return cached.(map[string]*CustomElementField)
}

// LookupAttributeField performs an O(1) lookup to find the CustomElementField
// associated with the given attribute name. Returns nil if not found.
//
// This method uses an internal cache that is lazily initialized on first use.
// Subsequent calls reuse the cached map for optimal performance.
//
// Performance: O(1) lookup vs O(n) linear search through Members.
// Thread-safe: Uses sync.Once to ensure safe concurrent initialization.
func (ced *CustomElementDeclaration) LookupAttributeField(attrName string) *CustomElementField {
	if ced == nil {
		return nil
	}
	return ced.getAttributeFieldMap()[attrName]
}

// BuildExportMaps creates maps from declaration name to export for fast lookups.
// Returns two maps: one for CustomElementExports and one for JavaScriptExports.
// Filters exports by module path (includes if module is empty or matches modulePath).
//
// This function is exported for testing purposes.
func BuildExportMaps(exports []Export, modulePath string) (
	customElementExports map[string]*CustomElementExport,
	jsExports map[string]*JavaScriptExport,
) {
	customElementExports = make(map[string]*CustomElementExport)
	jsExports = make(map[string]*JavaScriptExport)

	for i, exp := range exports {
		if cee, ok := exp.(*CustomElementExport); ok {
			if cee != nil && cee.Declaration != nil {
				name := cee.Declaration.Name
				// Match module path logic from original implementation
				if cee.Declaration.Module == "" || cee.Declaration.Module == modulePath {
					customElementExports[name] = exports[i].(*CustomElementExport)
				}
			}
		}
		if je, ok := exp.(*JavaScriptExport); ok {
			if je != nil && je.Declaration != nil && je.Declaration.Name != "" {
				if je.Declaration.Module == "" || je.Declaration.Module == modulePath {
					jsExports[je.Declaration.Name] = exports[i].(*JavaScriptExport)
				}
			}
		}
	}

	return customElementExports, jsExports
}

// buildExportMaps builds and caches the export maps if not already built.
// This method uses lazy initialization with sync.Once for thread-safe concurrent access.
// Both the cache and sync.Once are stored externally to avoid adding non-semantic fields to the struct.
func (m *Module) buildExportMaps() *exportMaps {
	// Check if already cached
	if cached, ok := exportMapsCache.Load(m); ok {
		return cached.(*exportMaps)
	}

	// Get or create a sync.Once for this instance
	once, _ := exportMapsOnce.LoadOrStore(m, &sync.Once{})

	// Use the sync.Once to ensure single initialization
	once.(*sync.Once).Do(func() {
		ceMap, jsMap := BuildExportMaps(m.Exports, m.Path)
		maps := &exportMaps{
			customElementExports: ceMap,
			jsExports:            jsMap,
		}
		exportMapsCache.Store(m, maps)
	})

	// Load the now-initialized cache
	cached, _ := exportMapsCache.Load(m)
	return cached.(*exportMaps)
}

// LookupCustomElementExport finds the CustomElementExport for a declaration name.
// Returns nil if not found.
//
// This method uses an external cache that is lazily initialized on first use.
// Subsequent calls reuse the cached map for optimal performance.
//
// Performance: O(1) lookup vs O(n) linear search through Exports.
// Thread-safe: Uses sync.Once to ensure safe concurrent initialization.
func (m *Module) LookupCustomElementExport(declName string) *CustomElementExport {
	if m == nil {
		return nil
	}
	maps := m.buildExportMaps()
	return maps.customElementExports[declName]
}

// LookupJavaScriptExport finds the JavaScriptExport for a declaration name.
// Returns nil if not found.
//
// This method uses an external cache that is lazily initialized on first use.
// Subsequent calls reuse the cached map for optimal performance.
//
// Performance: O(1) lookup vs O(n) linear search through Exports.
// Thread-safe: Uses sync.Once to ensure safe concurrent initialization.
func (m *Module) LookupJavaScriptExport(declName string) *JavaScriptExport {
	if m == nil {
		return nil
	}
	maps := m.buildExportMaps()
	return maps.jsExports[declName]
}
