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
func (ced *CustomElementDeclaration) getAttributeFieldMap() map[string]*CustomElementField {
	ced.attributeFieldOnce.Do(func() {
		ced.attributeFieldMap = BuildAttributeFieldMap(ced.Members)
	})
	return ced.attributeFieldMap
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
func (m *Module) buildExportMaps() {
	m.exportMapsOnce.Do(func() {
		m.customElementExportMap, m.jsExportMap = BuildExportMaps(m.Exports, m.Path)
	})
}

// LookupCustomElementExport finds the CustomElementExport for a declaration name.
// Returns nil if not found.
//
// This method uses an internal cache that is lazily initialized on first use.
// Subsequent calls reuse the cached map for optimal performance.
//
// Performance: O(1) lookup vs O(n) linear search through Exports.
// Thread-safe: Uses sync.Once to ensure safe concurrent initialization.
func (m *Module) LookupCustomElementExport(declName string) *CustomElementExport {
	if m == nil {
		return nil
	}
	m.buildExportMaps()
	return m.customElementExportMap[declName]
}

// LookupJavaScriptExport finds the JavaScriptExport for a declaration name.
// Returns nil if not found.
//
// This method uses an internal cache that is lazily initialized on first use.
// Subsequent calls reuse the cached map for optimal performance.
//
// Performance: O(1) lookup vs O(n) linear search through Exports.
// Thread-safe: Uses sync.Once to ensure safe concurrent initialization.
func (m *Module) LookupJavaScriptExport(declName string) *JavaScriptExport {
	if m == nil {
		return nil
	}
	m.buildExportMaps()
	return m.jsExportMap[declName]
}
