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
	"encoding/json"
	"fmt"
	"log"
)

// UnmarshalPackage unmarshals a CEM manifest JSON, handling Declarations interface types and ClassMember interface types.
// It also ensures that all slice fields are non-nil (empty slices if not present in JSON).
func UnmarshalPackage(data []byte) (*Package, error) {
	var raw struct {
		SchemaVersion string            `json:"schemaVersion"`
		Readme        *string           `json:"readme,omitempty"`
		Modules       []json.RawMessage `json:"modules"`
		Deprecated    any               `json:"deprecated,omitempty"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, fmt.Errorf("unmarshal package prelude: %w", err)
	}
	pkg := &Package{
		SchemaVersion: raw.SchemaVersion,
		Readme:        raw.Readme,
		Deprecated:    raw.Deprecated,
	}
	for _, modData := range raw.Modules {
		var modRaw struct {
			Kind         string            `json:"kind"`
			Path         string            `json:"path"`
			Summary      string            `json:"summary,omitempty"`
			Description  string            `json:"description,omitempty"`
			Declarations []json.RawMessage `json:"declarations,omitempty"`
			Exports      []json.RawMessage `json:"exports,omitempty"`
			Deprecated   any               `json:"deprecated,omitempty"`
		}
		if err := json.Unmarshal(modData, &modRaw); err != nil {
			return nil, fmt.Errorf("unmarshal module: %w", err)
		}
		var decls []Declaration
		for _, declData := range modRaw.Declarations {
			// Peek at kind
			var kindWrap struct {
				Kind string `json:"kind"`
			}
			if err := json.Unmarshal(declData, &kindWrap); err != nil {
				continue
			}
			switch kindWrap.Kind {
			case "class":
				// Peek for customElement: true
				var probe struct {
					CustomElement bool `json:"customElement"`
				}
				if err := json.Unmarshal(declData, &probe); err == nil && probe.CustomElement {
					var ce CustomElementDeclaration
					if err := unmarshalCustomElementDeclaration(declData, &ce); err == nil {
						decls = append(decls, &ce)
					} else {
						log.Printf("Failed to unmarshal custom-element (class+customElement): %v", err)
					}
				} else {
					var c ClassDeclaration
					if err := unmarshalClassDeclaration(declData, &c); err == nil {
						decls = append(decls, &c)
					} else {
						log.Printf("Failed to unmarshal class: %v", err)
					}
				}
			case "mixin":
				var m MixinDeclaration
				if err := unmarshalMixinDeclaration(declData, &m); err == nil {
					decls = append(decls, &m)
				} else {
					log.Printf("Failed to unmarshal mixin: %v", err)
				}
			case "function":
				var f FunctionDeclaration
				if err := json.Unmarshal(declData, &f); err == nil {
					ensureNonNilSlicesFunction(&f.FunctionLike)
					decls = append(decls, &f)
				}
			case "variable":
				var v VariableDeclaration
				if err := json.Unmarshal(declData, &v); err == nil {
					decls = append(decls, &v)
				}
			case "custom-element":
				var ce CustomElementDeclaration
				if err := unmarshalCustomElementDeclaration(declData, &ce); err == nil {
					decls = append(decls, &ce)
				} else {
					log.Printf("Failed to unmarshal custom-element: %v", err)
				}
			case "custom-element-mixin":
				var cem CustomElementMixinDeclaration
				if err := unmarshalCustomElementMixinDeclaration(declData, &cem); err == nil {
					decls = append(decls, &cem)
				} else {
					log.Printf("Failed to unmarshal custom-element-mixin: %v", err)
				}
			default:
				// fallback: skip unknown kind
			}
		}
		if decls == nil {
			decls = []Declaration{}
		}

		// --- Exports unmarshalling ---
		var exports []Export
		for _, exportData := range modRaw.Exports {
			var kindWrap struct {
				Kind string `json:"kind"`
			}
			if err := json.Unmarshal(exportData, &kindWrap); err != nil {
				continue
			}
			switch kindWrap.Kind {
			case "js":
				var e JavaScriptExport
				if err := json.Unmarshal(exportData, &e); err == nil {
					exports = append(exports, &e)
				} else {
					log.Printf("Failed to unmarshal js export: %v", err)
				}
			case "custom-element-definition":
				var e CustomElementExport
				if err := json.Unmarshal(exportData, &e); err == nil {
					exports = append(exports, &e)
				} else {
					log.Printf("Failed to unmarshal custom element export: %v", err)
				}
			default:
				// fallback: skip unknown kind
			}
		}
		if exports == nil {
			exports = []Export{}
		}
		// --- end Exports unmarshalling ---

		mod := JavaScriptModule{
			Kind:         modRaw.Kind,
			Path:         modRaw.Path,
			Summary:      modRaw.Summary,
			Description:  modRaw.Description,
			Declarations: decls,
			Exports:      exports,
			Deprecated:   modRaw.Deprecated,
		}
		pkg.Modules = append(pkg.Modules, mod)
	}
	if pkg.Modules == nil {
		pkg.Modules = []Module{}
	}
	return pkg, nil
}

// --- Custom unmarshalling for ClassDeclaration and subtypes ---

// unmarshalClassDeclaration handles members as []interface{} of ClassMember
func unmarshalClassDeclaration(data []byte, c *ClassDeclaration) error {
	type Alias ClassDeclaration
	aux := &struct {
		Members []json.RawMessage `json:"members"`
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	c.Members = nil
	for _, m := range aux.Members {
		member, err := unmarshalClassMember(m)
		if err == nil && member != nil {
			c.Members = append(c.Members, member)
		}
	}
	ensureNonNilSlicesClass(&c.ClassLike)
	return nil
}

func unmarshalMixinDeclaration(data []byte, m *MixinDeclaration) error {
	type Alias MixinDeclaration
	aux := &struct {
		Members []json.RawMessage `json:"members"`
		*Alias
	}{
		Alias: (*Alias)(m),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	m.Members = nil
	for _, mm := range aux.Members {
		member, err := unmarshalClassMember(mm)
		if err == nil && member != nil {
			m.Members = append(m.Members, member)
		}
	}
	ensureNonNilSlicesMixin(m)
	return nil
}

func unmarshalCustomElementDeclaration(data []byte, ce *CustomElementDeclaration) error {
	type Alias CustomElementDeclaration
	aux := &struct {
		Members []json.RawMessage `json:"members"`
		*Alias
	}{
		Alias: (*Alias)(ce),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	ce.Members = nil
	for _, mm := range aux.Members {
		member, err := unmarshalClassMember(mm)
		if err == nil && member != nil {
			ce.Members = append(ce.Members, member)
		}
	}
	ensureNonNilSlicesCustomElement(ce)
	return nil
}

func unmarshalCustomElementMixinDeclaration(data []byte, cem *CustomElementMixinDeclaration) error {
	type Alias CustomElementMixinDeclaration
	aux := &struct {
		Members []json.RawMessage `json:"members"`
		*Alias
	}{
		Alias: (*Alias)(cem),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	cem.Members = nil
	for _, mm := range aux.Members {
		member, err := unmarshalClassMember(mm)
		if err == nil && member != nil {
			cem.Members = append(cem.Members, member)
		}
	}
	ensureNonNilSlicesCustomElementMixin(cem)
	return nil
}

// unmarshalClassMember unmarshals a ClassMember from raw JSON.
func unmarshalClassMember(data json.RawMessage) (ClassMember, error) {
	var kindWrap struct {
		Kind string `json:"kind"`
	}
	if err := json.Unmarshal(data, &kindWrap); err != nil {
		return nil, err
	}
	switch kindWrap.Kind {
	case "field":
		// Probe for either "attribute" or "reflects"
		// PERFORMANCE: could be improved perhaps with a custom unmarshaller
		var probe struct {
			Attribute *string `json:"attribute"`
			Reflects  *bool   `json:"reflects"`
		}
		_ = json.Unmarshal(data, &probe)
		if probe.Attribute != nil || probe.Reflects != nil {
			var f CustomElementField
			if err := json.Unmarshal(data, &f); err == nil {
				return &f, nil
			}
			// If attribute or reflects is present, but not a valid CustomElementField, surface the error.
			return nil, fmt.Errorf("field has 'attribute' or 'reflects' but cannot unmarshal as CustomElementField")
		}
		// No attribute or reflects: try as ClassField
		var f ClassField
		if err := json.Unmarshal(data, &f); err == nil {
			return &f, nil
		}
		return nil, fmt.Errorf("field does not have 'attribute' or 'reflects' and cannot unmarshal as ClassField")
	case "method":
		var m ClassMethod
		if err := json.Unmarshal(data, &m); err == nil {
			return &m, nil
		}
		return nil, fmt.Errorf("cannot unmarshal as ClassMethod")
		// Add other kinds as needed
	}
	return nil, fmt.Errorf("unknown class member kind: %s", kindWrap.Kind)
}

// Helpers to ensure non-nil slices for all relevant structs:

func ensureNonNilSlicesCustomElement(ce *CustomElementDeclaration) {
	if ce.Attributes == nil {
		ce.Attributes = []Attribute{}
	}
	if ce.Events == nil {
		ce.Events = []Event{}
	}
	if ce.Slots == nil {
		ce.Slots = []Slot{}
	}
	if ce.CssParts == nil {
		ce.CssParts = []CssPart{}
	}
	if ce.CssProperties == nil {
		ce.CssProperties = []CssCustomProperty{}
	}
	if ce.CssStates == nil {
		ce.CssStates = []CssCustomState{}
	}
	if ce.Demos == nil {
		ce.Demos = []Demo{}
	}
	ensureNonNilSlicesClass(&ce.ClassDeclaration.ClassLike)
}

func ensureNonNilSlicesCustomElementMixin(cem *CustomElementMixinDeclaration) {
	ensureNonNilSlicesMixin(&cem.MixinDeclaration)
	ensureNonNilSlicesCustomElement(&cem.CustomElementDeclaration)
}

func ensureNonNilSlicesClass(c *ClassLike) {
	if c.Members == nil {
		c.Members = []ClassMember{}
	}
	if c.Mixins == nil {
		c.Mixins = []Reference{}
	}
}

func ensureNonNilSlicesFunction(f *FunctionLike) {
	if f.Parameters == nil {
		f.Parameters = []Parameter{}
	}
}

func ensureNonNilSlicesMixin(m *MixinDeclaration) {
	ensureNonNilSlicesFunction(&m.FunctionLike)
	ensureNonNilSlicesClass(&m.ClassLike)
}
