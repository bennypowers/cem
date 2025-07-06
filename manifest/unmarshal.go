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
	"errors"
	"fmt"
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

func (p *Parameter) UnmarshalJSON(data []byte) error {
	var proxy map[string]json.RawMessage
	if err := json.Unmarshal(data, &proxy); err != nil {
		return err
	}
	var depRaw json.RawMessage
	if raw, ok := proxy["deprecated"]; ok {
		depRaw = raw
		delete(proxy, "deprecated")
	}
	rest, err := json.Marshal(proxy)
	if err != nil {
		return err
	}
	type Alias Parameter
	if err := json.Unmarshal(rest, (*Alias)(p)); err != nil {
		return err
	}
	if len(depRaw) > 0 && string(depRaw) != "null" {
		var dep Deprecated
		if !decodeDeprecatedField(&dep, depRaw) {
			return fmt.Errorf("invalid type for deprecated field")
		}
		p.Deprecated = dep
	}
	return nil
}

func (v *VariableDeclaration) UnmarshalJSON(data []byte) error {
	var proxy map[string]json.RawMessage
	if err := json.Unmarshal(data, &proxy); err != nil {
		return err
	}
	var depRaw json.RawMessage
	if raw, ok := proxy["deprecated"]; ok {
		depRaw = raw
		delete(proxy, "deprecated")
	}
	rest, err := json.Marshal(proxy)
	if err != nil {
		return err
	}
	type Alias VariableDeclaration
	if err := json.Unmarshal(rest, (*Alias)(v)); err != nil {
		return err
	}
	if len(depRaw) > 0 && string(depRaw) != "null" {
		var dep Deprecated
		if !decodeDeprecatedField(&dep, depRaw) {
			return fmt.Errorf("invalid type for deprecated field")
		}
		v.Deprecated = dep
	}
	return nil
}

func decodeDeprecatedField(dst *Deprecated, data json.RawMessage) bool {
	if dst == nil {
		return true
	}
	if len(data) == 0 || string(data) == "null" {
		*dst = nil
		return true
	}
	var b bool
	if err := json.Unmarshal(data, &b); err == nil {
		*dst = DeprecatedFlag(b)
		return true
	}
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		*dst = DeprecatedReason(s)
		return true
	}
	return false // unknown type
}

// DRY helper for types with a Deprecated field. T is the struct type (not pointer).
func unmarshalWithDeprecated[T any](
	data json.RawMessage,
	setDeprecated func(*T, Deprecated),
) (T, error) {
	var proxy map[string]json.RawMessage
	var obj T
	var zero T
	if err := json.Unmarshal(data, &proxy); err != nil {
		return zero, err
	}
	var depRaw json.RawMessage
	if raw, ok := proxy["deprecated"]; ok {
		depRaw = raw
		delete(proxy, "deprecated")
	}
	rest, err := json.Marshal(proxy)
	if err != nil {
		return zero, err
	}
	if err := json.Unmarshal(rest, &obj); err != nil {
		return zero, err
	}
	if len(depRaw) > 0 && string(depRaw) != "null" {
		var dep Deprecated
		if !decodeDeprecatedField(&dep, depRaw) {
			return zero, fmt.Errorf("invalid type for deprecated field")
		}
		setDeprecated(&obj, dep)
	}
	return obj, nil
}

// --- Core Package Unmarshalling ---

func UnmarshalPackage(data []byte) (pkg *Package, errs error) {
	var raw struct {
		SchemaVersion string            `json:"schemaVersion"`
		Readme        *string           `json:"readme,omitempty"`
		Modules       []json.RawMessage `json:"modules"`
		Deprecated    json.RawMessage   `json:"deprecated,omitempty"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, fmt.Errorf("unmarshal package prelude: %w", err)
	}
	pkg = &Package{
		SchemaVersion: raw.SchemaVersion,
		Readme:        raw.Readme,
	}
	decodeDeprecatedField(&pkg.Deprecated, raw.Deprecated)
	for _, modData := range raw.Modules {
		mod, err := unmarshalModule(modData)
		if mod != nil {
			pkg.Modules = append(pkg.Modules, *mod)
		}
		if err != nil {
			errs = errors.Join(errs, err)
		}
	}
	return pkg, errs
}

func unmarshalModule(data []byte) (mod *Module, errs error) {
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
		mod = &JavaScriptModule{
			Kind:        raw.Kind,
			Path:        raw.Path,
			Summary:     raw.Summary,
			Description: raw.Description,
		}
		decodeDeprecatedField(&mod.Deprecated, raw.Deprecated)
		for _, d := range raw.Declarations {
			decl, err := unmarshalDeclaration(d)
			if err != nil {
				errs = errors.Join(errs, err)
			}
			if decl != nil {
				mod.Declarations = append(mod.Declarations, decl)
			}
		}
		for _, e := range raw.Exports {
			exp, err := unmarshalExport(e)
			if err == nil && exp != nil {
				mod.Exports = append(mod.Exports, exp)
			}
		}
		return mod, errs
	default:
		return nil, errors.Join(errs, fmt.Errorf("unknown module kind: %s", raw.Kind))
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
		// Check if customElement is true; if so, dispatch to unmarshalCustomElementDeclaration
		var probe struct {
			CustomElement bool `json:"customElement"`
		}
		if err := json.Unmarshal(data, &probe); err != nil {
			return nil, err
		}
		if probe.CustomElement {
			return unmarshalCustomElementDeclaration(data)
		}
		return unmarshalClassDeclaration(data)
	case "mixin":
		// Quick raw check for customElement
		var probe struct {
			Kind          string          `json:"kind"`
			CustomElement json.RawMessage `json:"customElement"`
		}
		if err := json.Unmarshal(data, &probe); err != nil {
			return nil, err
		}
		// If customElement is present and true, dispatch
		if probe.Kind == "mixin" && len(probe.CustomElement) > 0 && string(probe.CustomElement) != "null" {
			var isCE bool
			if err := json.Unmarshal(probe.CustomElement, &isCE); err == nil && isCE {
				return unmarshalCustomElementMixinDeclaration(data)
			}
		}
		return unmarshalMixinDeclaration(data)
	case "function":
		decl, err := unmarshalFunctionDeclaration(data)
		return &decl, err
	case "variable":
		var v VariableDeclaration
		if err := json.Unmarshal(data, &v); err == nil {
			return &v, nil
		} else {
			return nil, err
		}
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

// --- Classes ---

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

func unmarshalCustomElementDeclaration(data []byte) (ce *CustomElementDeclaration, errs error) {
	var raw struct {
		Name        string            `json:"name"`
		Kind        string            `json:"kind"`
		TagName     string            `json:"tagName"`
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
	ce = &CustomElementDeclaration{
		CustomElement: CustomElement{
			Demos:   raw.Demos,
			TagName: raw.TagName,
		},
		ClassDeclaration: ClassDeclaration{
			Kind: raw.Kind,
			ClassLike: ClassLike{
				Name:       raw.Name,
				Superclass: raw.Superclass,
				Mixins:     raw.Mixins,
			},
		},
	}
	decodeDeprecatedField(&ce.Deprecated, raw.Deprecated)

	// Unmarshal members with dispatch
	for _, mm := range raw.Members {
		member, err := unmarshalClassMember(mm)
		if err != nil || member == nil {
			continue
		}
		switch m := member.(type) {
		case *CustomElementField:
			ce.Members = append(ce.Members, m)
		case *ClassField:
			ce.Members = append(ce.Members, m)
		case *ClassMethod:
			ce.Members = append(ce.Members, m)
		default:
			// ignore unknown
		}
	}

	for _, a := range raw.Attributes {
		attr, err := unmarshalAttribute(a)
		if err == nil {
			ce.Attributes = append(ce.Attributes, attr)
		} else {
			errs = errors.Join(errs, err)
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
	return ce, errs
}

func unmarshalClassMember(data json.RawMessage) (ClassMember, error) {
	var kindWrap struct{ Kind string `json:"kind"` }
	if err := json.Unmarshal(data, &kindWrap); err != nil {
		return nil, err
	}
	switch kindWrap.Kind {
	case "field":
		// Probe for custom element field properties
		type probeField struct {
			Attribute *string `json:"attribute"`
			Reflects  *bool   `json:"reflects"`
		}
		var probe probeField
		if err := json.Unmarshal(data, &probe); err != nil {
			return nil, err
		}
		if probe.Attribute != nil || probe.Reflects != nil {
			var f CustomElementField
			if err := json.Unmarshal(data, &f); err == nil {
				f.Attribute = *probe.Attribute
				f.Reflects = *probe.Reflects
				return &f, nil
			}
			// If it fails, fallback to ClassField for robustness
			var fallback ClassField
			if err := json.Unmarshal(data, &fallback); err == nil {
				return &fallback, nil
			}
			return nil, fmt.Errorf("cannot unmarshal as CustomElementField or ClassField")
		}
		var f ClassField
		if err := json.Unmarshal(data, &f); err == nil {
			return &f, nil
		}
		return nil, fmt.Errorf("cannot unmarshal as ClassField")
	case "method":
		// Get parameters as raw and unmarshal the rest with deprecated handled
		var probe struct {
			Parameters []json.RawMessage `json:"parameters"`
		}
		if err := json.Unmarshal(data, &probe); err != nil {
			return nil, err
		}
		m, err := unmarshalWithDeprecated(data, func(x *ClassMethod, dep Deprecated) { x.Deprecated = dep })
		if err != nil {
			return nil, err
		}
		m.Parameters = nil
		for _, p := range probe.Parameters {
			param, err := unmarshalParameter(p)
			if err == nil {
				m.Parameters = append(m.Parameters, param)
			}
		}
		if m.Parameters == nil {
			m.Parameters = []Parameter{}
		}
		return &m, nil
	default:
		return nil, fmt.Errorf("unknown class member kind: %s", kindWrap.Kind)
	}
}

// --- Mixins ---

func unmarshalMixinDeclaration(data []byte) (*MixinDeclaration, error) {
	var raw struct {
		Name       string            `json:"name"`
		Kind       string            `json:"kind"`
		Members    []json.RawMessage `json:"members"`
		Deprecated json.RawMessage   `json:"deprecated"`
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

func unmarshalCustomElementMixinDeclaration(data []byte) (*CustomElementMixinDeclaration, error) {
	base, err := unmarshalMixinDeclaration(data)
	if err != nil {
		return nil, err
	}

	var raw struct {
		Attributes []json.RawMessage `json:"attributes"`
		Events     []json.RawMessage `json:"events"`
		Slots      []json.RawMessage `json:"slots"`
		CssParts   []json.RawMessage `json:"cssParts"`
		CssProps   []json.RawMessage `json:"cssProperties"`
		CssStates  []json.RawMessage `json:"cssStates"`
		Demos      []Demo            `json:"demos"`
		Members    []json.RawMessage `json:"members"` // for re-dispatch
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	cem := &CustomElementMixinDeclaration{
		MixinDeclaration: *base,
		Name:             base.Name,
	}

	cem.Members = nil
	for _, mm := range raw.Members {
		member, err := unmarshalClassMember(mm)
		if err != nil || member == nil {
			continue
		}
		cem.Members = append(cem.Members, member)
	}

	for _, a := range raw.Attributes {
		attr, err := unmarshalAttribute(a)
		if err == nil {
			cem.Attributes = append(cem.Attributes, attr)
		}
	}
	for _, e := range raw.Events {
		ev, err := unmarshalEvent(e)
		if err == nil {
			cem.Events = append(cem.Events, ev)
		}
	}
	for _, s := range raw.Slots {
		slot, err := unmarshalSlot(s)
		if err == nil {
			cem.Slots = append(cem.Slots, slot)
		}
	}
	for _, cp := range raw.CssParts {
		part, err := unmarshalCssPart(cp)
		if err == nil {
			cem.CssParts = append(cem.CssParts, part)
		}
	}
	for _, cpp := range raw.CssProps {
		prop, err := unmarshalCssCustomProperty(cpp)
		if err == nil {
			cem.CssProperties = append(cem.CssProperties, prop)
		}
	}
	for _, cs := range raw.CssStates {
		state, err := unmarshalCssCustomState(cs)
		if err == nil {
			cem.CssStates = append(cem.CssStates, state)
		}
	}
	cem.Demos = raw.Demos

	return cem, nil
}

// --- Functions Declaration ---

func unmarshalFunctionDeclaration(data []byte) (FunctionDeclaration, error) {
		// Read parameters as raw first
	var probe struct {
		Parameters []json.RawMessage `json:"parameters"`
	}
	if err := json.Unmarshal(data, &probe); err != nil {
		var zero FunctionDeclaration
		return zero, err
	}
	fn, err := unmarshalWithDeprecated(data, func(x *FunctionDeclaration, dep Deprecated) { x.Deprecated = dep })
	if err != nil {
		return fn, err
	}
	fn.Parameters = nil
	for _, p := range probe.Parameters {
		param, err := unmarshalParameter(p)
		if err == nil {
			fn.Parameters = append(fn.Parameters, param)
		}
	}
	return fn, nil
}

func unmarshalParameter(data json.RawMessage) (Parameter, error) {
	var proxy map[string]json.RawMessage
	if err := json.Unmarshal(data, &proxy); err != nil {
		var zero Parameter
		return zero, err
	}
	var depRaw json.RawMessage
	if raw, ok := proxy["deprecated"]; ok {
		depRaw = raw
		delete(proxy, "deprecated")
	}
	rest, err := json.Marshal(proxy)
	if err != nil {
		var zero Parameter
		return zero, err
	}
	var param Parameter
	if err := json.Unmarshal(rest, &param); err != nil {
		var zero Parameter
		return zero, err
	}
	if len(depRaw) > 0 && string(depRaw) != "null" {
		var dep Deprecated
		decodeDeprecatedField(&dep, depRaw)
		param.Deprecated = dep
	}
	return param, nil
}

// --- Leaf Types ---

func unmarshalAttribute(data json.RawMessage) (Attribute, error) {
	return unmarshalWithDeprecated(data, func(x *Attribute, dep Deprecated) { x.Deprecated = dep })
}

func unmarshalEvent(data json.RawMessage) (Event, error) {
	return unmarshalWithDeprecated(data, func(x *Event, dep Deprecated) { x.Deprecated = dep })
}

func unmarshalSlot(data json.RawMessage) (Slot, error) {
	return unmarshalWithDeprecated(data, func(x *Slot, dep Deprecated) { x.Deprecated = dep })
}

func unmarshalCssPart(data json.RawMessage) (CssPart, error) {
	return unmarshalWithDeprecated(data, func(x *CssPart, dep Deprecated) { x.Deprecated = dep })
}

func unmarshalCssCustomProperty(data json.RawMessage) (CssCustomProperty, error) {
	return unmarshalWithDeprecated(data, func(x *CssCustomProperty, dep Deprecated) { x.Deprecated = dep })
}

func unmarshalCssCustomState(data json.RawMessage) (CssCustomState, error) {
	return unmarshalWithDeprecated(data, func(x *CssCustomState, dep Deprecated) { x.Deprecated = dep })
}
