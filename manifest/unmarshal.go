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

// --- Helpers ---

func (f *ClassField) UnmarshalJSON(data []byte) error {
	type Alias ClassField
	aux := &struct {
		Deprecated json.RawMessage `json:"deprecated"`
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if len(aux.Deprecated) != 0 && string(aux.Deprecated) != "null" {
		// Try bool
		var b bool
		if err := json.Unmarshal(aux.Deprecated, &b); err == nil {
			f.Deprecated = DeprecatedFlag(b)
			return nil
		}
		// Try string
		var s string
		if err := json.Unmarshal(aux.Deprecated, &s); err == nil {
			f.Deprecated = DeprecatedReason(s)
			return nil
		}
		// Ignore other types
	}
	return nil
}

func decodeDeprecatedField(dst *Deprecated, data json.RawMessage) {
	if dst == nil {
		return
	}
	if len(data) == 0 || string(data) == "null" {
		*dst = nil
		return
	}
	var b bool
	if err := json.Unmarshal(data, &b); err == nil {
		*dst = DeprecatedFlag(b)
		return
	}
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		*dst = DeprecatedReason(s)
	}
}

// --- Core Package Unmarshalling ---

func UnmarshalPackage(data []byte) (*Package, error) {
	var raw struct {
		SchemaVersion string            `json:"schemaVersion"`
		Readme        *string           `json:"readme,omitempty"`
		Modules       []json.RawMessage `json:"modules"`
		Deprecated    json.RawMessage   `json:"deprecated,omitempty"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, fmt.Errorf("unmarshal package prelude: %w", err)
	}
	pkg := &Package{
		SchemaVersion: raw.SchemaVersion,
		Readme:        raw.Readme,
	}
	decodeDeprecatedField(&pkg.Deprecated, raw.Deprecated)
	for _, modData := range raw.Modules {
		mod, err := unmarshalModule(modData)
		if err != nil {
			log.Printf("Failed to unmarshal module: %v", err)
			continue
		}
		pkg.Modules = append(pkg.Modules, *mod)
	}
	if pkg.Modules == nil {
		pkg.Modules = []Module{}
	}
	return pkg, nil
}

func unmarshalModule(data []byte) (*Module, error) {
	var raw struct {
		Kind         string            `json:"kind"`
		Path         string            `json:"path"`
		Summary      string            `json:"summary,omitempty"`
		Description  string            `json:"description,omitempty"`
		Declarations []json.RawMessage `json:"declarations,omitempty"`
		Exports      []json.RawMessage `json:"exports,omitempty"`
		Deprecated   json.RawMessage   `json:"deprecated,omitempty"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}
	switch raw.Kind {
	case "javascript-module":
		mod := &JavaScriptModule{
			Kind:        raw.Kind,
			Path:        raw.Path,
			Summary:     raw.Summary,
			Description: raw.Description,
		}
		decodeDeprecatedField(&mod.Deprecated, raw.Deprecated)
		for _, d := range raw.Declarations {
			decl, err := unmarshalDeclaration(d)
			if err == nil && decl != nil {
				mod.Declarations = append(mod.Declarations, decl)
			}
		}
		for _, e := range raw.Exports {
			exp, err := unmarshalExport(e)
			if err == nil && exp != nil {
				mod.Exports = append(mod.Exports, exp)
			}
		}
		if mod.Declarations == nil {
			mod.Declarations = []Declaration{}
		}
		if mod.Exports == nil {
			mod.Exports = []Export{}
		}
		return mod, nil
	default:
		return nil, fmt.Errorf("unknown module kind: %s", raw.Kind)
	}
}

// --- Declarations/Exports ---

func unmarshalDeclaration(data []byte) (Declaration, error) {
	var kindWrap struct{ Kind string `json:"kind"` }
	if err := json.Unmarshal(data, &kindWrap); err != nil {
		return nil, err
	}
	switch kindWrap.Kind {
	case "class":
		return unmarshalClassDeclaration(data)
	case "mixin":
		return unmarshalMixinDeclaration(data)
	case "function":
		return unmarshalFunctionDeclaration(data)
	case "variable":
		var v VariableDeclaration
		if err := json.Unmarshal(data, &v); err == nil {
			return &v, nil
		} else {
			return nil, err
		}
	case "custom-element":
		return unmarshalCustomElementDeclaration(data)
	case "custom-element-mixin":
		return unmarshalCustomElementMixinDeclaration(data)
	default:
		return nil, fmt.Errorf("unknown declaration kind: %s", kindWrap.Kind)
	}
}

func unmarshalExport(data []byte) (Export, error) {
	var kindWrap struct{ Kind string `json:"kind"` }
	if err := json.Unmarshal(data, &kindWrap); err != nil {
		return nil, err
	}
	switch kindWrap.Kind {
	case "js":
		var e JavaScriptExport
		if err := json.Unmarshal(data, &e); err == nil {
			return &e, nil
		} else {
			return nil, err
		}
	case "custom-element-definition":
		var e CustomElementExport
		if err := json.Unmarshal(data, &e); err == nil {
			return &e, nil
		} else {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unknown export kind: %s", kindWrap.Kind)
	}
}

// --- Class/Mixin Declarations ---

func unmarshalClassDeclaration(data []byte) (*ClassDeclaration, error) {
	var raw struct {
		Name        string            `json:"name"`
		Kind        string            `json:"kind"`
		Superclass  *Reference        `json:"superclass"`
		Mixins      []Reference       `json:"mixins"`
		Members     []json.RawMessage `json:"members"`
		Deprecated  json.RawMessage   `json:"deprecated"`
		// add other fields as needed
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}
	c := &ClassDeclaration{
		Kind:       raw.Kind,
		ClassLike: ClassLike{
			Name:       raw.Name,
			Superclass: raw.Superclass,
			Mixins:     raw.Mixins,
		},
	}
	decodeDeprecatedField(&c.Deprecated, raw.Deprecated)

	for _, m := range raw.Members {
		member, err := unmarshalClassMember(m)
		if err == nil && member != nil {
			c.Members = append(c.Members, member)
		}
	}
	if c.Members == nil {
		c.Members = []ClassMember{}
	}
	if c.Mixins == nil {
		c.Mixins = []Reference{}
	}
	return c, nil
}

func unmarshalMixinDeclaration(data []byte) (*MixinDeclaration, error) {
	var raw struct {
		Name       string            `json:"name"`
		Kind       string            `json:"kind"`
		Members    []json.RawMessage `json:"members"`
		Deprecated json.RawMessage   `json:"deprecated"`
		// Add fields as needed
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}
	m := &MixinDeclaration{
		Name: raw.Name,
		Kind: raw.Kind,
	}
	decodeDeprecatedField(&m.Deprecated, raw.Deprecated)
	for _, mm := range raw.Members {
		member, err := unmarshalClassMember(mm)
		if err == nil && member != nil {
			m.Members = append(m.Members, member)
		}
	}
	if m.Members == nil {
		m.Members = []ClassMember{}
	}
	return m, nil
}

// --- Custom Element Declarations ---

func unmarshalCustomElementDeclaration(data []byte) (*CustomElementDeclaration, error) {
	var raw struct {
		// flatten ClassDeclaration for embedded struct fields
		Name        string            `json:"name"`
		Kind        string            `json:"kind"`
		Superclass  *Reference        `json:"superclass"`
		Mixins      []Reference       `json:"mixins"`
		Members     []json.RawMessage `json:"members"`
		Attributes  []json.RawMessage `json:"attributes"`
		Events      []json.RawMessage `json:"events"`
		Slots       []json.RawMessage `json:"slots"`
		CssParts    []json.RawMessage `json:"cssParts"`
		CssProps    []json.RawMessage `json:"cssProperties"`
		CssStates   []json.RawMessage `json:"cssStates"`
		Demos       []Demo            `json:"demos"`
		Deprecated  json.RawMessage   `json:"deprecated"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}
	ce := &CustomElementDeclaration{
		CustomElement: CustomElement{
			Demos: raw.Demos,
		},
		ClassDeclaration: ClassDeclaration{
			Kind:       raw.Kind,
			ClassLike: ClassLike{
				Name:       raw.Name,
				Superclass: raw.Superclass,
				Mixins:     raw.Mixins,
			},
		},
	}
	decodeDeprecatedField(&ce.Deprecated, raw.Deprecated)
	for _, mm := range raw.Members {
		member, err := unmarshalClassMember(mm)
		if err == nil && member != nil {
			ce.Members = append(ce.Members, member)
		}
	}
	for _, a := range raw.Attributes {
		attr, err := unmarshalAttribute(a)
		if err == nil {
			ce.Attributes = append(ce.Attributes, attr)
		}
	}
	for _, e := range raw.Events {
		ev, err := unmarshalEvent(e)
		if err == nil {
			ce.Events = append(ce.Events, ev)
		}
	}
	for _, s := range raw.Slots {
		slot, err := unmarshalSlot(s)
		if err == nil {
			ce.Slots = append(ce.Slots, slot)
		}
	}
	for _, cp := range raw.CssParts {
		part, err := unmarshalCssPart(cp)
		if err == nil {
			ce.CssParts = append(ce.CssParts, part)
		}
	}
	for _, cpp := range raw.CssProps {
		prop, err := unmarshalCssCustomProperty(cpp)
		if err == nil {
			ce.CssProperties = append(ce.CssProperties, prop)
		}
	}
	for _, cs := range raw.CssStates {
		state, err := unmarshalCssCustomState(cs)
		if err == nil {
			ce.CssStates = append(ce.CssStates, state)
		}
	}
	// Ensure non-nil slices
	if ce.Members == nil {
		ce.Members = []ClassMember{}
	}
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
	if ce.Mixins == nil {
		ce.Mixins = []Reference{}
	}
	return ce, nil
}

func unmarshalCustomElementMixinDeclaration(data []byte) (*CustomElementMixinDeclaration, error) {
	var raw struct {
		Name       string            `json:"name"`
		Kind       string            `json:"kind"`
		Members    []json.RawMessage `json:"members"`
		Deprecated json.RawMessage   `json:"deprecated"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}
	cem := &CustomElementMixinDeclaration{
		Name: raw.Name,
		MixinDeclaration: MixinDeclaration{
			Kind: raw.Kind,
		},
	}
	decodeDeprecatedField(&cem.Deprecated, raw.Deprecated)
	for _, mm := range raw.Members {
		member, err := unmarshalClassMember(mm)
		if err == nil && member != nil {
			cem.Members = append(cem.Members, member)
		}
	}
	if cem.Members == nil {
		cem.Members = []ClassMember{}
	}
	return cem, nil
}

// --- Function Declaration ---

func unmarshalFunctionDeclaration(data []byte) (*FunctionDeclaration, error) {
	var f FunctionDeclaration
	if err := json.Unmarshal(data, &f); err != nil {
		return nil, err
	}
	if f.Parameters == nil {
		f.Parameters = []Parameter{}
	}
	return &f, nil
}

// --- Class Members ---

func unmarshalClassMember(data json.RawMessage) (ClassMember, error) {
	var kindWrap struct{ Kind string `json:"kind"` }
	if err := json.Unmarshal(data, &kindWrap); err != nil {
		return nil, err
	}
	switch kindWrap.Kind {
	case "field":
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
			return nil, fmt.Errorf("field has 'attribute' or 'reflects' but cannot unmarshal as CustomElementField")
		}
		var f ClassField
		if err := json.Unmarshal(data, &f); err == nil {
			return &f, nil
		}
		return nil, fmt.Errorf("cannot unmarshal as ClassField")
	case "method":
		var m ClassMethod
		if err := json.Unmarshal(data, &m); err == nil {
			return &m, nil
		}
		return nil, fmt.Errorf("cannot unmarshal as ClassMethod")
	default:
		return nil, fmt.Errorf("unknown class member kind: %s", kindWrap.Kind)
	}
}

// --- Leaf Types ---

func unmarshalAttribute(data json.RawMessage) (Attribute, error) {
	var a Attribute
	if err := json.Unmarshal(data, &a); err != nil {
		return a, err
	}
	return a, nil
}
func unmarshalEvent(data json.RawMessage) (Event, error) {
	var e Event
	if err := json.Unmarshal(data, &e); err != nil {
		return e, err
	}
	return e, nil
}
func unmarshalSlot(data json.RawMessage) (Slot, error) {
	var s Slot
	if err := json.Unmarshal(data, &s); err != nil {
		return s, err
	}
	return s, nil
}
func unmarshalCssPart(data json.RawMessage) (CssPart, error) {
	var cp CssPart
	if err := json.Unmarshal(data, &cp); err != nil {
		return cp, err
	}
	return cp, nil
}
func unmarshalCssCustomProperty(data json.RawMessage) (CssCustomProperty, error) {
	var cpp CssCustomProperty
	if err := json.Unmarshal(data, &cpp); err != nil {
		return cpp, err
	}
	return cpp, nil
}
func unmarshalCssCustomState(data json.RawMessage) (CssCustomState, error) {
	var cs CssCustomState
	if err := json.Unmarshal(data, &cs); err != nil {
		return cs, err
	}
	return cs, nil
}
