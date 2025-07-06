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
	"errors"
	"slices"
	"strconv"
	"strings"
)

func (x *Package) GetAllTagNames() (tags []string) {
	// Write index.html
	for _, m := range x.Modules {
		for _, decl := range m.Declarations {
			ced, ok := decl.(*CustomElementDeclaration)
			if !ok || ced.TagName == "" {
				continue
			}
			tags = append(tags, ced.TagName)
		}
	}
	return tags
}

type RenderableMemberWithContext interface {
	ToTableRow() []string
	IsDeprecated() bool
	GetMemberType() string
}

type ModuleWithContext struct {
	Path                     string
	Module                   *Module
	CustomElementExports     []CustomElementExport
}
func (x ModuleWithContext) ToTableRow() []string {
	tags := make([]string, 0)
	for _, cee := range x.CustomElementExports {
		tags = append(tags, cee.Name)
	}
	return []string{
		x.Path,
		strings.Join(tags, ", "),
	}
}

func (x ModuleWithContext) IsDeprecated() bool {
	return x.Module != nil && x.Module.Deprecated != nil
}

func (x ModuleWithContext) GetMemberType() string {
	return "module"
}

type CustomElementWithContext struct {
	TagName                  string
	Module                   *Module
	CustomElementDeclaration *CustomElementDeclaration
	CustomElementExport      *CustomElementExport
}

// Renders a CustomElement as a table row.
// Columns:
//   Tag (tag name), Class (class name), Module (module path), Summary
func (c CustomElementWithContext) ToTableRow() []string {
	modulePath := ""
	if c.Module != nil {
		modulePath = c.Module.Path
	}
	className := ""
	summary := ""
	if c.CustomElementDeclaration != nil {
		className = c.CustomElementDeclaration.Name
		summary = c.CustomElementDeclaration.Summary
	}
	return []string{
		c.TagName,
		className,
		modulePath,
		summary,
	}
}

func (c CustomElementWithContext) IsDeprecated() bool {
	if c.CustomElementExport != nil && c.CustomElementExport.Deprecated != nil {
		return true
	}
	return c.CustomElementDeclaration != nil && c.CustomElementDeclaration.Deprecated != nil
}

func (c CustomElementWithContext) GetMemberType() string {
	return "element"
}

type AttributeWithContext struct {
	Name                     string
	Attribute                *Attribute
	JavaScriptModule         *JavaScriptModule
	CustomElementField       *CustomElementField
	CustomElementDeclaration *CustomElementDeclaration
	CustomElementExport      *CustomElementExport
}

// Renders a CustomElement as a table row.
// Columns:
//   Name, DOM Property, Reflects, Summary
func (a AttributeWithContext) ToTableRow() []string {
	domProp := ""
	reflects := "❌"
	if a.CustomElementField != nil {
		domProp = a.CustomElementField.Name
		if a.CustomElementField.Reflects {
			reflects = "✅"
		}
	}
	return []string{
		a.Name,
		domProp,
		reflects,
		a.Attribute.Summary,
	}
}

func (a AttributeWithContext) IsDeprecated() bool {
	return a.Attribute != nil && a.Attribute.Deprecated != nil
}

func (a AttributeWithContext) GetMemberType() string {
	return "attribute"
}

type SlotWithContext struct {
	Name                     string
	Slot                     *Slot
	JavaScriptModule         *JavaScriptModule
	CustomElementDeclaration *CustomElementDeclaration
	CustomElementExport      *CustomElementExport
}

// Renders a Slot as a table row.
// Columns:
//   Name, Summary
func (s SlotWithContext) ToTableRow() []string {
	slotName := s.Name
	if slotName == "" {
		slotName = "<default>"
	}
	return []string{
		slotName,
		s.Slot.Summary,
	}
}

func (s SlotWithContext) IsDeprecated() bool {
	return s.Slot != nil && s.Slot.Deprecated != nil
}

func (s SlotWithContext) GetMemberType() string {
	return "slot"
}

type CssCustomPropertyWithContext struct {
	Name                     string
	CssCustomProperty        *CssCustomProperty
	CustomElementDeclaration *CustomElementDeclaration
	CustomElementExport      *CustomElementExport
	JavaScriptModule         *JavaScriptModule
}

// Renders a CSS CssCustomProperty as a table row.
// Columns:
//   Name, Syntax, Default, Summary
func (c CssCustomPropertyWithContext) ToTableRow() []string {
	return []string{
		c.Name,
		c.CssCustomProperty.Syntax,
		c.CssCustomProperty.Default,
		c.CssCustomProperty.Summary,
	}
}

func (c CssCustomPropertyWithContext) IsDeprecated() bool {
	return c.CssCustomProperty != nil && c.CssCustomProperty.Deprecated != nil
}

func (c CssCustomPropertyWithContext) GetMemberType() string {
	return "css-property"
}

type CssCustomStateWithContext struct {
	Name                     string
	CssCustomState           *CssCustomState
	CustomElementDeclaration *CustomElementDeclaration
	CustomElementExport      *CustomElementExport
	JavaScriptModule         *JavaScriptModule
}

// Renders a CssCustomState as a table row.
// Columns:
//   Name, Summary
func (c CssCustomStateWithContext) ToTableRow() []string {
	return []string{
		c.Name,
		c.CssCustomState.Summary,
	}
}

func (c CssCustomStateWithContext) IsDeprecated() bool {
	return c.CssCustomState != nil && c.CssCustomState.Deprecated != nil
}

func (c CssCustomStateWithContext) GetMemberType() string {
	return "css-state"
}

type CssPartWithContext struct {
	Name                     string
	CssPart                  *CssPart
	CustomElementDeclaration *CustomElementDeclaration
	CustomElementExport      *CustomElementExport
	JavaScriptModule         *JavaScriptModule
}

// Renders a CssPart as a table row.
// Columns:
//   Name, Summary
func (c CssPartWithContext) ToTableRow() []string {
	return []string{
		c.Name,
		c.CssPart.Summary,
	}
}

func (c CssPartWithContext) IsDeprecated() bool {
	return c.CssPart != nil && c.CssPart.Deprecated != nil
}

func (c CssPartWithContext) GetMemberType() string {
	return "css-part"
}

type EventWithContext struct {
	Name                     string
	Event                    *Event
	CustomElementDeclaration *CustomElementDeclaration
	CustomElementExport      *CustomElementExport
	JavaScriptModule         *JavaScriptModule
}

// Renders an Event as a table row.
// Columns:
//   Name, Type, Summary
func (e EventWithContext) ToTableRow() []string {
	eventType := ""
	if e.Event.Type != nil {
		eventType = e.Event.Type.Text
	}
	return []string{
		e.Name,
		eventType,
		e.Event.Summary,
	}
}

func (e EventWithContext) IsDeprecated() bool {
	return e.Event != nil && e.Event.Deprecated != nil
}

func (e EventWithContext) GetMemberType() string {
	return "event"
}

type MethodWithContext struct {
	Name                     string
	Method                   *ClassMethod
	CustomElementDeclaration *CustomElementDeclaration
	CustomElementExport      *CustomElementExport
	JavaScriptModule         *JavaScriptModule
}

// Renders an Event as a table row.
// Columns:
//   Name, Return Type, Privacy, Static, Summary
func (m MethodWithContext) ToTableRow() []string {
	returnType := "void"
	privacy := string(m.Method.Privacy)
	if privacy == "" {
		privacy = "public"
	}
	if m.Method.Return != nil && m.Method.Return.Type != nil {
		returnType = m.Method.Return.Type.Text
	}
	return []string{
		m.Name,
		returnType,
		privacy,
		strconv.FormatBool(m.Method.Static),
		m.Method.Summary,
	}
}

func (m MethodWithContext) IsDeprecated() bool {
	return m.Method != nil && m.Method.Deprecated != nil
}

func (m MethodWithContext) GetMemberType() string {
	return "method"
}

type FieldWithContext struct {
	Name                     string
	Field                    ClassMember // Could be ClassField or CustomElementField
	CustomElementDeclaration *CustomElementDeclaration
	CustomElementExport      *CustomElementExport
	JavaScriptModule         *JavaScriptModule
}

// Renders a Field as a table row.
// Columns:
//   Name, Type, Privacy, Static, Summary
func (f FieldWithContext) ToTableRow() []string {
	var fieldType, privacy, summary string
	var static bool

	switch field := f.Field.(type) {
	case *ClassField:
		if field.Type != nil {
			fieldType = field.Type.Text
		}
		privacy = string(field.Privacy)
		static = field.Static
		summary = field.Summary
	case *CustomElementField:
		if field.Type != nil {
			fieldType = field.Type.Text
		}
		privacy = string(field.Privacy)
		static = field.Static
		summary = field.Summary
	}

	if privacy == "" {
		privacy = "public"
	}

	return []string{
		f.Name,
		fieldType,
		privacy,
		strconv.FormatBool(static),
		summary,
	}
}

func (f FieldWithContext) IsDeprecated() bool {
	switch field := f.Field.(type) {
	case *ClassField:
		return field.Deprecated != nil
	case *CustomElementField:
		return field.Deprecated != nil
	}
	return false
}

func (f FieldWithContext) GetMemberType() string {
	return "field"
}

// GetAllModulesWithContext returns a slice of ModuleWithContext for all modules.
func (x *Package) GetAllModulesWithContext() (modules []ModuleWithContext) {
	ms := make(map[string]ModuleWithContext)
	for i := range x.Modules {
		module := &x.Modules[i]
		ms[module.Path] = ModuleWithContext{Path: module.Path, Module: module}
		for _, e := range module.Exports {
			if cee, ok := e.(*CustomElementExport); ok {
				c := ms[module.Path]
				c.CustomElementExports = append(ms[module.Path].CustomElementExports, *cee)
				ms[module.Path] = c
			}
		}
	}
	for _, m := range ms {
		modules = append(modules, m)
	}
	slices.SortStableFunc(modules, func(a ModuleWithContext, b ModuleWithContext) int {
		return strings.Compare(a.Path, b.Path)
	})
	return modules
}

// GetAllTagNamesWithContext returns a slice of CustomElementWithContext for all custom elements in all modules.
func (x *Package) GetAllTagNamesWithContext() (tags []CustomElementWithContext) {
	for _, m := range x.Modules {
		mrs := make(map[string]CustomElementWithContext)
		for _, d := range m.Declarations {
			if ced, ok := d.(*CustomElementDeclaration); ok {
				mrs[ced.TagName] = CustomElementWithContext{CustomElementDeclaration: ced, Module: &m, TagName: ced.TagName}
			}
		}
		for _, e := range m.Exports {
			if cee, ok := e.(*CustomElementExport); ok {
				r := mrs[cee.Name]
				r.CustomElementExport = cee
				mrs[cee.Name] = r
			}
		}
		for _, r := range mrs {
			tags = append(tags, r)
		}
	}
	slices.SortStableFunc(tags, func(a CustomElementWithContext, b CustomElementWithContext) int {
		return int(a.CustomElementDeclaration.StartByte - b.CustomElementDeclaration.StartByte)
	})
	return tags
}

// findCustomElementContext locates the first CustomElementDeclaration, CustomElementExport, and JavaScriptModule for a given tagName.
func (x *Package) findCustomElementContext(tagName string) (*CustomElementDeclaration, *CustomElementExport, *JavaScriptModule, error) {
	for _, m := range x.Modules {
		for _, d := range m.Declarations {
			if ced, ok := d.(*CustomElementDeclaration); ok && ced.TagName == tagName {
				var ceExport *CustomElementExport
				for _, e := range m.Exports {
					if cee, ok := e.(*CustomElementExport); ok && cee.Name == tagName {
						ceExport = cee
						break
					}
				}
				return ced, ceExport, &m, nil
			}
		}
	}
	return nil, nil, nil, errors.New("Tag not found: " + tagName)
}

// GetTagAttrsWithContext returns attributes for a given tag name with context.
func (x *Package) GetTagAttrsWithContext(tagName string) ([]AttributeWithContext, error) {
	ced, ceExport, m, err := x.findCustomElementContext(tagName)
	if err != nil {
		return nil, err
	}
	fieldMap := make(map[string]CustomElementField)
	for _, member := range ced.Members {
		if cef, ok := member.(*CustomElementField); ok {
			fieldMap[cef.Attribute] = *cef
		}
	}
	attrMap := make(map[string]AttributeWithContext)
	for _, attr := range ced.Attributes {
		attrCopy := attr // Create a copy of the loop variable
		var field *CustomElementField
		if f, ok := fieldMap[attrCopy.Name]; ok {
			field = &f
		}
		attrMap[attrCopy.Name] = AttributeWithContext{
			Name:                     attrCopy.Name,
			Attribute:                &attrCopy,
			CustomElementDeclaration: ced,
			CustomElementField:       field,
			CustomElementExport:      ceExport,
			JavaScriptModule:         m,
		}
	}
	attrs := make([]AttributeWithContext, 0, len(attrMap))
	for _, awc := range attrMap {
		attrs = append(attrs, awc)
	}
	slices.SortStableFunc(attrs, func(a AttributeWithContext, b AttributeWithContext) int {
		if a.Attribute != nil && b.Attribute != nil {
			return int(a.Attribute.StartByte - b.Attribute.StartByte)
		}
		return 0
	})
	return attrs, nil
}

// GetTagSlotsWithContext returns slots for a given tag name with context.
func (x *Package) GetTagSlotsWithContext(tagName string) ([]SlotWithContext, error) {
	ced, ceExport, m, err := x.findCustomElementContext(tagName)
	if err != nil {
		return nil, err
	}
	var slots []SlotWithContext
	for i := range ced.Slots {
		slot := &ced.Slots[i]
		slots = append(slots, SlotWithContext{
			Name:                     slot.Name,
			Slot:                     slot,
			CustomElementDeclaration: ced,
			CustomElementExport:      ceExport,
			JavaScriptModule:         m,
		})
	}
	return slots, nil
}

// GetTagCssPropertiesWithContext returns CSS custom properties for a given tag name with context.
func (x *Package) GetTagCssPropertiesWithContext(tagName string) ([]CssCustomPropertyWithContext, error) {
	ced, ceExport, m, err := x.findCustomElementContext(tagName)
	if err != nil {
		return nil, err
	}
	var props []CssCustomPropertyWithContext
	for i := range ced.CssProperties {
		prop := &ced.CssProperties[i]
		props = append(props, CssCustomPropertyWithContext{
			Name:                     prop.Name,
			CssCustomProperty:        prop,
			CustomElementDeclaration: ced,
			CustomElementExport:      ceExport,
			JavaScriptModule:         m,
		})
	}
	return props, nil
}

// GetTagCssStatesWithContext returns CSS custom states for a given tag name with context.
func (x *Package) GetTagCssStatesWithContext(tagName string) ([]CssCustomStateWithContext, error) {
	ced, ceExport, m, err := x.findCustomElementContext(tagName)
	if err != nil {
		return nil, err
	}
	var states []CssCustomStateWithContext
	for i := range ced.CssStates {
		state := &ced.CssStates[i]
		states = append(states, CssCustomStateWithContext{
			Name:                     state.Name,
			CssCustomState:           state,
			CustomElementDeclaration: ced,
			CustomElementExport:      ceExport,
			JavaScriptModule:         m,
		})
	}
	return states, nil
}

// GetTagCssPartsWithContext returns CSS shadow parts for a given tag name with context.
func (x *Package) GetTagCssPartsWithContext(tagName string) ([]CssPartWithContext, error) {
	ced, ceExport, m, err := x.findCustomElementContext(tagName)
	if err != nil {
		return nil, err
	}
	var parts []CssPartWithContext
	for i := range ced.CssParts {
		part := &ced.CssParts[i]
		parts = append(parts, CssPartWithContext{
			Name:                     part.Name,
			CssPart:                  part,
			CustomElementDeclaration: ced,
			CustomElementExport:      ceExport,
			JavaScriptModule:         m,
		})
	}
	return parts, nil
}

// GetTagEventsWithContext returns events for a given tag name with context.
func (x *Package) GetTagEventsWithContext(tagName string) ([]EventWithContext, error) {
	ced, ceExport, m, err := x.findCustomElementContext(tagName)
	if err != nil {
		return nil, err
	}
	var events []EventWithContext
	for i := range ced.Events {
		event := &ced.Events[i]
		events = append(events, EventWithContext{
			Name:                     event.Name,
			Event:                    event,
			CustomElementDeclaration: ced,
			CustomElementExport:      ceExport,
			JavaScriptModule:         m,
		})
	}
	return events, nil
}

// GetTagMethodsWithContext returns methods for a given tag name with context.
func (x *Package) GetTagMethodsWithContext(tagName string) ([]MethodWithContext, error) {
	ced, ceExport, m, err := x.findCustomElementContext(tagName)
	if err != nil {
		return nil, err
	}
	var methods []MethodWithContext
	for _, member := range ced.Members {
		if method, ok := member.(*ClassMethod); ok {
			methods = append(methods, MethodWithContext{
				Name:                     method.Name,
				Method:                   method,
				CustomElementDeclaration: ced,
				CustomElementExport:      ceExport,
				JavaScriptModule:         m,
			})
		}
	}
	return methods, nil
}

// GetTagFieldsWithContext returns fields for a given tag name with context.
func (x *Package) GetTagFieldsWithContext(tagName string) ([]FieldWithContext, error) {
	ced, ceExport, m, err := x.findCustomElementContext(tagName)
	if err != nil {
		return nil, err
	}
	var fields []FieldWithContext
	for _, member := range ced.Members {
		switch field := member.(type) {
		case *ClassField:
			fields = append(fields, FieldWithContext{
				Name:                     field.Name,
				Field:                    field,
				CustomElementDeclaration: ced,
				CustomElementExport:      ceExport,
				JavaScriptModule:         m,
			})
		case *CustomElementField:
			fields = append(fields, FieldWithContext{
				Name:                     field.Name,
				Field:                    field,
				CustomElementDeclaration: ced,
				CustomElementExport:      ceExport,
				JavaScriptModule:         m,
			})
		}
	}
	return fields, nil
}

// GetTagAllMembersWithContext returns all class members grouped by type for a given tag name.
func (x *Package) GetTagAllMembersWithContext(tagName string) (map[string][]RenderableMemberWithContext, error) {
	result := make(map[string][]RenderableMemberWithContext)

	// Get attributes
	if attrs, err := x.GetTagAttrsWithContext(tagName); err == nil {
		var members []RenderableMemberWithContext
		for _, attr := range attrs {
			members = append(members, attr)
		}
		if len(members) > 0 {
			result["attributes"] = members
		}
	}

	// Get fields
	if fields, err := x.GetTagFieldsWithContext(tagName); err == nil {
		var members []RenderableMemberWithContext
		for _, field := range fields {
			members = append(members, field)
		}
		if len(members) > 0 {
			result["fields"] = members
		}
	}

	// Get methods
	if methods, err := x.GetTagMethodsWithContext(tagName); err == nil {
		var members []RenderableMemberWithContext
		for _, method := range methods {
			members = append(members, method)
		}
		if len(members) > 0 {
			result["methods"] = members
		}
	}

	// Get events
	if events, err := x.GetTagEventsWithContext(tagName); err == nil {
		var members []RenderableMemberWithContext
		for _, event := range events {
			members = append(members, event)
		}
		if len(members) > 0 {
			result["events"] = members
		}
	}

	// Get slots
	if slots, err := x.GetTagSlotsWithContext(tagName); err == nil {
		var members []RenderableMemberWithContext
		for _, slot := range slots {
			members = append(members, slot)
		}
		if len(members) > 0 {
			result["slots"] = members
		}
	}

	return result, nil
}
