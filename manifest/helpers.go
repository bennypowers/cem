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
	return []string{
		c.TagName,
		c.CustomElementDeclaration.Name,
		modulePath,
		c.CustomElementDeclaration.Summary,
	}
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

// GetTagAttrsWithContext returns attributes for a given tag name with context.
// NB: Only the first matching tag is considered per manifest, as tag names are assumed global.
// If duplicate attribute names are present, only one will be returned (see generate.go for possible warnings).
func (x *Package) GetTagAttrsWithContext(tagName string) (attrs []AttributeWithContext, err error) {
	attrMap := make(map[string]AttributeWithContext)
	tagFound := false
modules:
	for _, m := range x.Modules {
		for _, d := range m.Declarations {
			if ced, ok := d.(*CustomElementDeclaration); ok {
				if tagFound = ced.TagName == tagName; tagFound {
					fieldMap := make(map[string]CustomElementField)
					for _, member := range ced.Members {
						if cef, ok := member.(*CustomElementField); ok {
							fieldMap[cef.Attribute] = *cef
						}
					}
					var ceExport *CustomElementExport
					for _, e := range m.Exports {
						if cee, ok := e.(*CustomElementExport); ok {
							if cee.Name == tagName {
								ceExport = cee
								break
							}
						}
					}
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
							JavaScriptModule:         &m,
						}
					}
					break modules // NB: Tag names are assumed global; stop after first found
				}
			}
		}
	}
	if !tagFound {
		return nil, errors.New("Tag not found: " + tagName)
	}
	attrs = make([]AttributeWithContext, 0, len(attrMap))
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
func (x *Package) GetTagSlotsWithContext(tagName string) (slots []SlotWithContext, err error) {
	tagFound := false
modules:
	for _, m := range x.Modules {
		for _, d := range m.Declarations {
			if ced, ok := d.(*CustomElementDeclaration); ok {
				if tagFound = ced.TagName == tagName; tagFound {
					var ceExport *CustomElementExport
					for _, e := range m.Exports {
						if cee, ok := e.(*CustomElementExport); ok && cee.Name == tagName {
							ceExport = cee
							break
						}
					}
					for i := range ced.Slots {
						slot := &ced.Slots[i]
						slots = append(slots, SlotWithContext{
							Name:                     slot.Name,
							Slot:                     slot,
							CustomElementDeclaration: ced,
							CustomElementExport:      ceExport,
							JavaScriptModule:         &m,
						})
					}
					break modules // NB: Tag names are assumed global; stop after first found
				}
			}
		}
	}
	if !tagFound {
		return nil, errors.New("Tag not found: " + tagName)
	}
	return slots, nil
}

// GetTagCssPropertiesWithContext returns CSS custom properties for a given tag name with context.
func (x *Package) GetTagCssPropertiesWithContext(tagName string) (props []CssCustomPropertyWithContext, err error) {
	tagFound := false
modules:
	for _, m := range x.Modules {
		for _, d := range m.Declarations {
			if ced, ok := d.(*CustomElementDeclaration); ok {
				if tagFound = ced.TagName == tagName; tagFound {
					var ceExport *CustomElementExport
					for _, e := range m.Exports {
						if cee, ok := e.(*CustomElementExport); ok && cee.Name == tagName {
							ceExport = cee
							break
						}
					}
					for i := range ced.CssProperties {
						prop := &ced.CssProperties[i]
						props = append(props, CssCustomPropertyWithContext{
							Name:                     prop.Name,
							CssCustomProperty:        prop,
							CustomElementDeclaration: ced,
							CustomElementExport:      ceExport,
							JavaScriptModule:         &m,
						})
					}
					break modules // NB: Tag names are assumed global; stop after first found
				}
			}
		}
	}
	if !tagFound {
		return nil, errors.New("Tag not found: " + tagName)
	}
	return props, nil
}

// GetTagCssStatesWithContext returns CSS custom states for a given tag name with context.
func (x *Package) GetTagCssStatesWithContext(tagName string) (states []CssCustomStateWithContext, err error) {
	tagFound := false
modules:
	for _, m := range x.Modules {
		for _, d := range m.Declarations {
			if ced, ok := d.(*CustomElementDeclaration); ok {
				if tagFound = ced.TagName == tagName; tagFound {
					var ceExport *CustomElementExport
					for _, e := range m.Exports {
						if cee, ok := e.(*CustomElementExport); ok && cee.Name == tagName {
							ceExport = cee
							break
						}
					}
					for i := range ced.CssStates {
						state := &ced.CssStates[i]
						states = append(states, CssCustomStateWithContext{
							Name:                     state.Name,
							CssCustomState:           state,
							CustomElementDeclaration: ced,
							CustomElementExport:      ceExport,
							JavaScriptModule:         &m,
						})
					}
					break modules // NB: Tag names are assumed global; stop after first found
				}
			}
		}
	}
	if !tagFound {
		return nil, errors.New("Tag not found: " + tagName)
	}
	return states, nil
}

// GetTagCssPartsWithContext returns CSS shadow parts for a given tag name with context.
func (x *Package) GetTagCssPartsWithContext(tagName string) (parts []CssPartWithContext, err error) {
	tagFound := false
modules:
	for _, m := range x.Modules {
		for _, d := range m.Declarations {
			if ced, ok := d.(*CustomElementDeclaration); ok {
				if tagFound = ced.TagName == tagName; tagFound {
					var ceExport *CustomElementExport
					for _, e := range m.Exports {
						if cee, ok := e.(*CustomElementExport); ok && cee.Name == tagName {
							ceExport = cee
							break
						}
					}
					for i := range ced.CssParts {
						part := &ced.CssParts[i]
						parts = append(parts, CssPartWithContext{
							Name:                     part.Name,
							CssPart:                  part,
							CustomElementDeclaration: ced,
							CustomElementExport:      ceExport,
							JavaScriptModule:         &m,
						})
					}
					break modules // NB: Tag names are assumed global; stop after first found
				}
			}
		}
	}
	if tagFound {
		return nil, errors.New("Tag not found: " + tagName)
	}
	return parts, nil
}

// GetTagEventsWithContext returns events for a given tag name with context.
func (x *Package) GetTagEventsWithContext(tagName string) (events []EventWithContext, err error) {
	tagFound := false
modules:
	for _, m := range x.Modules {
		for _, d := range m.Declarations {
			if ced, ok := d.(*CustomElementDeclaration); ok {
				if tagFound = ced.TagName == tagName; tagFound {
					var ceExport *CustomElementExport
					for _, e := range m.Exports {
						if cee, ok := e.(*CustomElementExport); ok && cee.Name == tagName {
							ceExport = cee
							break
						}
					}
					for i := range ced.Events {
						event := &ced.Events[i]
						events = append(events, EventWithContext{
							Name:                     event.Name,
							Event:                    event,
							CustomElementDeclaration: ced,
							CustomElementExport:      ceExport,
							JavaScriptModule:         &m,
						})
					}
					break modules // NB: Tag names are assumed global; stop after first found
				}
			}
		}
	}
	if !tagFound {
		return nil, errors.New("Tag not found: " + tagName)
	}
	return events, nil
}

// GetTagMethodsWithContext returns methods for a given tag name with context.
func (x *Package) GetTagMethodsWithContext(tagName string) (methods []MethodWithContext, err error) {
	tagFound := false
modules:
	for _, m := range x.Modules {
		for _, d := range m.Declarations {
			if ced, ok := d.(*CustomElementDeclaration); ok {
				if tagFound = ced.TagName == tagName; tagFound {
					var ceExport *CustomElementExport
					for _, e := range m.Exports {
						if cee, ok := e.(*CustomElementExport); ok && cee.Name == tagName {
							ceExport = cee
							break
						}
					}
					for _, member := range ced.Members {
						if method, ok := member.(*ClassMethod); ok {
							methods = append(methods, MethodWithContext{
								Name:                     method.Name,
								Method:                   method,
								CustomElementDeclaration: ced,
								CustomElementExport:      ceExport,
								JavaScriptModule:         &m,
							})
						}
					}
					break modules // NB: Tag names are assumed global; stop after first found
				}
			}
		}
	}
	if !tagFound {
		return nil, errors.New("Tag not found: " + tagName)
	}
	return methods, nil
}
