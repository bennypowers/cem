/*
Copyright © 2026 Benny Powers <web@bennypowers.com>

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
package breaking_test

import (
	"testing"

	"bennypowers.dev/cem/breaking"
	M "bennypowers.dev/cem/manifest"
	"github.com/stretchr/testify/assert"
)

func makeElements(elements ...*M.CustomElementDeclaration) map[string]*M.CustomElementDeclaration {
	m := make(map[string]*M.CustomElementDeclaration)
	for _, el := range elements {
		m[el.TagName] = el
	}
	return m
}

func ced(tagName string, opts ...func(*M.CustomElementDeclaration)) *M.CustomElementDeclaration {
	el := &M.CustomElementDeclaration{}
	el.TagName = tagName
	el.CustomElement.CustomElement = true // named field, can't simplify
	for _, opt := range opts {
		opt(el)
	}
	return el
}

func withAttributes(attrs ...M.Attribute) func(*M.CustomElementDeclaration) {
	return func(el *M.CustomElementDeclaration) {
		el.CustomElement.Attributes = attrs
	}
}

func withSlots(slots ...M.Slot) func(*M.CustomElementDeclaration) {
	return func(el *M.CustomElementDeclaration) {
		el.CustomElement.Slots = slots
	}
}

func withEvents(events ...M.Event) func(*M.CustomElementDeclaration) {
	return func(el *M.CustomElementDeclaration) {
		el.CustomElement.Events = events
	}
}

func withCssProperties(props ...M.CssCustomProperty) func(*M.CustomElementDeclaration) {
	return func(el *M.CustomElementDeclaration) {
		el.CustomElement.CssProperties = props
	}
}

func withCssParts(parts ...M.CssPart) func(*M.CustomElementDeclaration) {
	return func(el *M.CustomElementDeclaration) {
		el.CustomElement.CssParts = parts
	}
}

func withCssStates(states ...M.CssCustomState) func(*M.CustomElementDeclaration) {
	return func(el *M.CustomElementDeclaration) {
		el.CustomElement.CssStates = states
	}
}

func withMembers(members ...M.ClassMember) func(*M.CustomElementDeclaration) {
	return func(el *M.CustomElementDeclaration) {
		el.Members = members
	}
}

func attr(name string, opts ...func(*M.Attribute)) M.Attribute {
	a := M.Attribute{}
	a.Name = name
	for _, opt := range opts {
		opt(&a)
	}
	return a
}

func withType(text string) func(*M.Attribute) {
	return func(a *M.Attribute) {
		a.Type = &M.Type{Text: text}
	}
}

func withDefault(val string) func(*M.Attribute) {
	return func(a *M.Attribute) {
		a.Default = val
	}
}

func slot(name string) M.Slot {
	s := M.Slot{}
	s.Name = name
	return s
}

func event(name string, typeText ...string) M.Event {
	e := M.Event{}
	e.Name = name
	if len(typeText) > 0 {
		e.Type = &M.Type{Text: typeText[0]}
	}
	return e
}

func cssProp(name string, defaultVal ...string) M.CssCustomProperty {
	p := M.CssCustomProperty{}
	p.Name = name
	if len(defaultVal) > 0 {
		p.Default = defaultVal[0]
	}
	return p
}

func cssPart(name string) M.CssPart {
	p := M.CssPart{}
	p.Name = name
	return p
}

func cssState(name string) M.CssCustomState {
	s := M.CssCustomState{}
	s.Name = name
	return s
}

func method(name string, privacy M.Privacy, params []M.Parameter, ret *M.Return) *M.ClassMethod {
	m := &M.ClassMethod{
		Kind:    "method",
		Privacy: privacy,
	}
	m.Name = name
	m.Parameters = params
	m.Return = ret
	return m
}

func field(name string, privacy M.Privacy, typeText string) *M.ClassField {
	f := &M.ClassField{
		Kind:    "field",
		Privacy: privacy,
	}
	f.Name = name
	if typeText != "" {
		f.Type = &M.Type{Text: typeText}
	}
	return f
}

func param(name string, typeText string) M.Parameter {
	p := M.Parameter{}
	p.Name = name
	if typeText != "" {
		p.Type = &M.Type{Text: typeText}
	}
	return p
}

// Tier 1: inline assertions on pure rule logic
func TestElementRules(t *testing.T) {
	t.Run("element-removed", func(t *testing.T) {
		base := makeElements(ced("my-element"))
		head := makeElements()
		rule := findRule("element-removed")
		changes := rule.Check(base, head)
		assert.Len(t, changes, 1)
		assert.Equal(t, breaking.Breaking, changes[0].Severity)
		assert.Equal(t, "my-element", changes[0].Element)
	})

	t.Run("element-added", func(t *testing.T) {
		base := makeElements()
		head := makeElements(ced("my-element"))
		rule := findRule("element-added")
		changes := rule.Check(base, head)
		assert.Len(t, changes, 1)
		assert.Equal(t, breaking.Safe, changes[0].Severity)
		assert.Equal(t, "my-element", changes[0].Element)
	})

	t.Run("no changes", func(t *testing.T) {
		base := makeElements(ced("my-element"))
		head := makeElements(ced("my-element"))
		removed := findRule("element-removed")
		added := findRule("element-added")
		assert.Empty(t, removed.Check(base, head))
		assert.Empty(t, added.Check(base, head))
	})
}

// Tier 1: inline assertions on pure rule logic
func TestAttributeRules(t *testing.T) {
	t.Run("attribute-removed", func(t *testing.T) {
		base := makeElements(ced("my-el", withAttributes(attr("variant"))))
		head := makeElements(ced("my-el"))
		changes := findRule("attribute-removed").Check(base, head)
		assert.Len(t, changes, 1)
		assert.Equal(t, breaking.Breaking, changes[0].Severity)
		assert.Equal(t, "variant", changes[0].Subject)
	})

	t.Run("attribute-added", func(t *testing.T) {
		base := makeElements(ced("my-el"))
		head := makeElements(ced("my-el", withAttributes(attr("color"))))
		changes := findRule("attribute-added").Check(base, head)
		assert.Len(t, changes, 1)
		assert.Equal(t, breaking.Safe, changes[0].Severity)
	})

	t.Run("attribute-type-changed", func(t *testing.T) {
		base := makeElements(ced("my-el", withAttributes(attr("size", withType("string")))))
		head := makeElements(ced("my-el", withAttributes(attr("size", withType("number")))))
		changes := findRule("attribute-type-changed").Check(base, head)
		assert.Len(t, changes, 1)
		assert.Equal(t, breaking.Breaking, changes[0].Severity)
	})

	t.Run("attribute-type-unchanged", func(t *testing.T) {
		base := makeElements(ced("my-el", withAttributes(attr("size", withType("string")))))
		head := makeElements(ced("my-el", withAttributes(attr("size", withType("string")))))
		changes := findRule("attribute-type-changed").Check(base, head)
		assert.Empty(t, changes)
	})

	t.Run("attribute-default-changed", func(t *testing.T) {
		base := makeElements(ced("my-el", withAttributes(attr("size", withDefault(`"medium"`)))))
		head := makeElements(ced("my-el", withAttributes(attr("size", withDefault(`"large"`)))))
		changes := findRule("attribute-default-changed").Check(base, head)
		assert.Len(t, changes, 1)
		assert.Equal(t, breaking.Dangerous, changes[0].Severity)
	})

	t.Run("skips removed element", func(t *testing.T) {
		base := makeElements(ced("my-el", withAttributes(attr("variant"))))
		head := makeElements()
		assert.Empty(t, findRule("attribute-removed").Check(base, head))
	})
}

// Tier 1: inline assertions on pure rule logic
func TestSlotRules(t *testing.T) {
	t.Run("slot-removed", func(t *testing.T) {
		base := makeElements(ced("my-el", withSlots(slot("header"))))
		head := makeElements(ced("my-el"))
		changes := findRule("slot-removed").Check(base, head)
		assert.Len(t, changes, 1)
		assert.Equal(t, breaking.Breaking, changes[0].Severity)
	})

	t.Run("slot-added", func(t *testing.T) {
		base := makeElements(ced("my-el"))
		head := makeElements(ced("my-el", withSlots(slot("footer"))))
		changes := findRule("slot-added").Check(base, head)
		assert.Len(t, changes, 1)
		assert.Equal(t, breaking.Safe, changes[0].Severity)
	})

	t.Run("default slot removed", func(t *testing.T) {
		base := makeElements(ced("my-el", withSlots(slot(""))))
		head := makeElements(ced("my-el"))
		changes := findRule("slot-removed").Check(base, head)
		assert.Len(t, changes, 1)
		assert.Contains(t, changes[0].Message, "(default)")
	})
}

// Tier 1: inline assertions on pure rule logic
func TestEventRules(t *testing.T) {
	t.Run("event-removed", func(t *testing.T) {
		base := makeElements(ced("my-el", withEvents(event("change"))))
		head := makeElements(ced("my-el"))
		changes := findRule("event-removed").Check(base, head)
		assert.Len(t, changes, 1)
		assert.Equal(t, breaking.Breaking, changes[0].Severity)
	})

	t.Run("event-added", func(t *testing.T) {
		base := makeElements(ced("my-el"))
		head := makeElements(ced("my-el", withEvents(event("input"))))
		changes := findRule("event-added").Check(base, head)
		assert.Len(t, changes, 1)
		assert.Equal(t, breaking.Safe, changes[0].Severity)
	})

	t.Run("event-type-changed", func(t *testing.T) {
		base := makeElements(ced("my-el", withEvents(event("change", "CustomEvent"))))
		head := makeElements(ced("my-el", withEvents(event("change", "Event"))))
		changes := findRule("event-type-changed").Check(base, head)
		assert.Len(t, changes, 1)
		assert.Equal(t, breaking.Breaking, changes[0].Severity)
	})
}

// Tier 1: inline assertions on pure rule logic
func TestCssRules(t *testing.T) {
	t.Run("css-custom-property-removed", func(t *testing.T) {
		base := makeElements(ced("my-el", withCssProperties(cssProp("--bg"))))
		head := makeElements(ced("my-el"))
		changes := findRule("css-custom-property-removed").Check(base, head)
		assert.Len(t, changes, 1)
		assert.Equal(t, breaking.Breaking, changes[0].Severity)
	})

	t.Run("css-custom-property-added", func(t *testing.T) {
		base := makeElements(ced("my-el"))
		head := makeElements(ced("my-el", withCssProperties(cssProp("--color"))))
		changes := findRule("css-custom-property-added").Check(base, head)
		assert.Len(t, changes, 1)
		assert.Equal(t, breaking.Safe, changes[0].Severity)
	})

	t.Run("css-custom-property-default-changed", func(t *testing.T) {
		base := makeElements(ced("my-el", withCssProperties(cssProp("--bg", "red"))))
		head := makeElements(ced("my-el", withCssProperties(cssProp("--bg", "blue"))))
		changes := findRule("css-custom-property-default-changed").Check(base, head)
		assert.Len(t, changes, 1)
		assert.Equal(t, breaking.Dangerous, changes[0].Severity)
	})

	t.Run("css-part-removed", func(t *testing.T) {
		base := makeElements(ced("my-el", withCssParts(cssPart("button"))))
		head := makeElements(ced("my-el"))
		changes := findRule("css-part-removed").Check(base, head)
		assert.Len(t, changes, 1)
		assert.Equal(t, breaking.Breaking, changes[0].Severity)
	})

	t.Run("css-part-added", func(t *testing.T) {
		base := makeElements(ced("my-el"))
		head := makeElements(ced("my-el", withCssParts(cssPart("icon"))))
		changes := findRule("css-part-added").Check(base, head)
		assert.Len(t, changes, 1)
		assert.Equal(t, breaking.Safe, changes[0].Severity)
	})

	t.Run("css-state-removed", func(t *testing.T) {
		base := makeElements(ced("my-el", withCssStates(cssState("active"))))
		head := makeElements(ced("my-el"))
		changes := findRule("css-state-removed").Check(base, head)
		assert.Len(t, changes, 1)
		assert.Equal(t, breaking.Breaking, changes[0].Severity)
	})

	t.Run("css-state-added", func(t *testing.T) {
		base := makeElements(ced("my-el"))
		head := makeElements(ced("my-el", withCssStates(cssState("focused"))))
		changes := findRule("css-state-added").Check(base, head)
		assert.Len(t, changes, 1)
		assert.Equal(t, breaking.Safe, changes[0].Severity)
	})
}

// Tier 1: inline assertions on pure rule logic
func TestMemberRules(t *testing.T) {
	t.Run("method-removed", func(t *testing.T) {
		base := makeElements(ced("my-el", withMembers(method("doStuff", M.Public, nil, nil))))
		head := makeElements(ced("my-el"))
		changes := findRule("method-removed").Check(base, head)
		assert.Len(t, changes, 1)
		assert.Equal(t, breaking.Breaking, changes[0].Severity)
	})

	t.Run("private method removed is not breaking", func(t *testing.T) {
		base := makeElements(ced("my-el", withMembers(method("_internal", M.Private, nil, nil))))
		head := makeElements(ced("my-el"))
		changes := findRule("method-removed").Check(base, head)
		assert.Empty(t, changes)
	})

	t.Run("method-added", func(t *testing.T) {
		base := makeElements(ced("my-el"))
		head := makeElements(ced("my-el", withMembers(method("doStuff", M.Public, nil, nil))))
		changes := findRule("method-added").Check(base, head)
		assert.Len(t, changes, 1)
		assert.Equal(t, breaking.Safe, changes[0].Severity)
	})

	t.Run("method-return-type-changed", func(t *testing.T) {
		base := makeElements(ced("my-el", withMembers(
			method("getValue", M.Public, nil, &M.Return{Type: &M.Type{Text: "string"}}),
		)))
		head := makeElements(ced("my-el", withMembers(
			method("getValue", M.Public, nil, &M.Return{Type: &M.Type{Text: "number"}}),
		)))
		changes := findRule("method-return-type-changed").Check(base, head)
		assert.Len(t, changes, 1)
		assert.Equal(t, breaking.Breaking, changes[0].Severity)
	})

	t.Run("method-parameter-changed", func(t *testing.T) {
		base := makeElements(ced("my-el", withMembers(
			method("doStuff", M.Public, []M.Parameter{param("name", "string")}, nil),
		)))
		head := makeElements(ced("my-el", withMembers(
			method("doStuff", M.Public, []M.Parameter{param("name", "string"), param("value", "number")}, nil),
		)))
		changes := findRule("method-parameter-changed").Check(base, head)
		assert.Len(t, changes, 1)
		assert.Equal(t, breaking.Breaking, changes[0].Severity)
	})

	t.Run("return-type-only change does not trigger parameter-changed", func(t *testing.T) {
		base := makeElements(ced("my-el", withMembers(
			method("doStuff", M.Public, []M.Parameter{param("name", "string")}, &M.Return{Type: &M.Type{Text: "void"}}),
		)))
		head := makeElements(ced("my-el", withMembers(
			method("doStuff", M.Public, []M.Parameter{param("name", "string")}, &M.Return{Type: &M.Type{Text: "Promise<void>"}}),
		)))
		changes := findRule("method-parameter-changed").Check(base, head)
		assert.Empty(t, changes)
	})

	// Overload sets: manifests carry one member per overload signature, so
	// same-name methods compare per overload, matched by parameter
	// signature. A name counts as removed only when absent entirely.
	t.Run("overload set: added overload is not breaking", func(t *testing.T) {
		base := makeElements(ced("my-el", withMembers(
			method("format", M.Public, []M.Parameter{param("value", "string")}, nil),
		)))
		head := makeElements(ced("my-el", withMembers(
			method("format", M.Public, []M.Parameter{param("value", "string")}, nil),
			method("format", M.Public, []M.Parameter{param("value", "number")}, nil),
		)))
		assert.Empty(t, findRule("method-removed").Check(base, head))
		assert.Empty(t, findRule("method-parameter-changed").Check(base, head))
		assert.Empty(t, findRule("method-return-type-changed").Check(base, head))
	})

	t.Run("overload set: removed overload is breaking naming the overload", func(t *testing.T) {
		base := makeElements(ced("my-el", withMembers(
			method("format", M.Public, []M.Parameter{param("value", "string")}, nil),
			method("format", M.Public, []M.Parameter{param("value", "number")}, nil),
		)))
		head := makeElements(ced("my-el", withMembers(
			method("format", M.Public, []M.Parameter{param("value", "string")}, nil),
		)))
		assert.Empty(t, findRule("method-removed").Check(base, head))
		changes := findRule("method-parameter-changed").Check(base, head)
		assert.Len(t, changes, 1)
		assert.Equal(t, breaking.Breaking, changes[0].Severity)
		assert.Contains(t, changes[0].Message, "(number)")
	})

	t.Run("overload set: changed params on one overload is breaking naming that overload", func(t *testing.T) {
		base := makeElements(ced("my-el", withMembers(
			method("format", M.Public, []M.Parameter{param("value", "string")}, nil),
			method("format", M.Public, []M.Parameter{param("value", "number")}, nil),
		)))
		head := makeElements(ced("my-el", withMembers(
			method("format", M.Public, []M.Parameter{param("value", "string")}, nil),
			method("format", M.Public, []M.Parameter{param("value", "number"), param("radix", "number")}, nil),
		)))
		changes := findRule("method-parameter-changed").Check(base, head)
		assert.Len(t, changes, 1)
		assert.Equal(t, breaking.Breaking, changes[0].Severity)
		assert.Contains(t, changes[0].Message, "(number)")
	})

	t.Run("overload set: return type compared on the param-matched overload", func(t *testing.T) {
		base := makeElements(ced("my-el", withMembers(
			method("format", M.Public, []M.Parameter{param("value", "string")}, &M.Return{Type: &M.Type{Text: "string"}}),
			method("format", M.Public, []M.Parameter{param("value", "number")}, &M.Return{Type: &M.Type{Text: "number"}}),
		)))
		head := makeElements(ced("my-el", withMembers(
			method("format", M.Public, []M.Parameter{param("value", "string")}, &M.Return{Type: &M.Type{Text: "string"}}),
			method("format", M.Public, []M.Parameter{param("value", "number")}, &M.Return{Type: &M.Type{Text: "string"}}),
		)))
		changes := findRule("method-return-type-changed").Check(base, head)
		assert.Len(t, changes, 1)
		assert.Equal(t, breaking.Breaking, changes[0].Severity)
		assert.Contains(t, changes[0].Message, "number")
	})

	t.Run("overload set: method name kept when any overload remains", func(t *testing.T) {
		base := makeElements(ced("my-el", withMembers(
			method("format", M.Public, []M.Parameter{param("value", "string")}, nil),
		)))
		head := makeElements(ced("my-el", withMembers(
			method("format", M.Public, []M.Parameter{param("value", "number")}, nil),
		)))
		assert.Empty(t, findRule("method-removed").Check(base, head))
		changes := findRule("method-parameter-changed").Check(base, head)
		assert.Len(t, changes, 1)
		assert.Equal(t, breaking.Breaking, changes[0].Severity)
	})

	t.Run("overload set: optional to required is breaking naming the overload", func(t *testing.T) {
		baseOptional := param("value", "string")
		baseOptional.Optional = true
		base := makeElements(ced("my-el", withMembers(
			method("format", M.Public, []M.Parameter{baseOptional}, nil),
			method("format", M.Public, []M.Parameter{param("value", "number")}, nil),
		)))
		head := makeElements(ced("my-el", withMembers(
			method("format", M.Public, []M.Parameter{param("value", "string")}, nil),
			method("format", M.Public, []M.Parameter{param("value", "number")}, nil),
		)))
		assert.Empty(t, findRule("method-removed").Check(base, head))
		changes := findRule("method-parameter-changed").Check(base, head)
		assert.Len(t, changes, 1)
		assert.Equal(t, breaking.Breaking, changes[0].Severity)
		assert.Contains(t, changes[0].Message, "(string?)")
	})

	t.Run("rest to non-rest is breaking", func(t *testing.T) {
		baseRest := param("values", "string[]")
		baseRest.Rest = true
		base := makeElements(ced("my-el", withMembers(
			method("emit", M.Public, []M.Parameter{baseRest}, nil),
		)))
		head := makeElements(ced("my-el", withMembers(
			method("emit", M.Public, []M.Parameter{param("values", "string[]")}, nil),
		)))
		changes := findRule("method-parameter-changed").Check(base, head)
		assert.Len(t, changes, 1)
		assert.Equal(t, breaking.Breaking, changes[0].Severity)
		assert.Contains(t, changes[0].Message, "(...string[])")
		assert.Contains(t, changes[0].Message, "(string[])")
	})

	t.Run("static and instance methods of one name are distinct", func(t *testing.T) {
		baseStatic := method("create", M.Public, []M.Parameter{param("input", "string")}, nil)
		baseStatic.Static = true
		base := makeElements(ced("my-el", withMembers(
			baseStatic,
			method("create", M.Public, []M.Parameter{param("input", "string")}, nil),
		)))
		head := makeElements(ced("my-el", withMembers(
			method("create", M.Public, []M.Parameter{param("input", "string")}, nil),
		)))
		changes := findRule("method-removed").Check(base, head)
		assert.Len(t, changes, 1)
		assert.Equal(t, breaking.Breaking, changes[0].Severity)
		assert.Contains(t, changes[0].Message, "create")
	})

	t.Run("instance method removed while same-name static method added is reported", func(t *testing.T) {
		headStatic := method("foo", M.Public, []M.Parameter{param("x", "string")}, nil)
		headStatic.Static = true
		base := makeElements(ced("my-el", withMembers(
			method("foo", M.Public, []M.Parameter{param("x", "string")}, nil),
		)))
		head := makeElements(ced("my-el", withMembers(headStatic)))
		changes := findRule("method-removed").Check(base, head)
		assert.Len(t, changes, 1)
		assert.Contains(t, changes[0].Message, "foo")
	})

	t.Run("overload set: duplicate signatures pair one-to-one", func(t *testing.T) {
		base := makeElements(ced("my-el", withMembers(
			method("format", M.Public, []M.Parameter{param("value", "string")}, &M.Return{Type: &M.Type{Text: "string"}}),
			method("format", M.Public, []M.Parameter{param("value", "string")}, &M.Return{Type: &M.Type{Text: "number"}}),
		)))
		head := makeElements(ced("my-el", withMembers(
			method("format", M.Public, []M.Parameter{param("value", "string")}, &M.Return{Type: &M.Type{Text: "string"}}),
		)))
		changes := findRule("method-parameter-changed").Check(base, head)
		assert.Len(t, changes, 1)
		assert.Contains(t, changes[0].Message, "(string)")
	})

	t.Run("field-added", func(t *testing.T) {
		base := makeElements(ced("my-el"))
		head := makeElements(ced("my-el", withMembers(field("value", M.Public, "string"))))
		changes := findRule("field-added").Check(base, head)
		assert.Len(t, changes, 1)
		assert.Equal(t, breaking.Safe, changes[0].Severity)
	})

	t.Run("field-removed", func(t *testing.T) {
		base := makeElements(ced("my-el", withMembers(field("value", M.Public, "string"))))
		head := makeElements(ced("my-el"))
		changes := findRule("field-removed").Check(base, head)
		assert.Len(t, changes, 1)
		assert.Equal(t, breaking.Breaking, changes[0].Severity)
	})

	t.Run("field-type-changed", func(t *testing.T) {
		base := makeElements(ced("my-el", withMembers(field("value", M.Public, "string"))))
		head := makeElements(ced("my-el", withMembers(field("value", M.Public, "number"))))
		changes := findRule("field-type-changed").Check(base, head)
		assert.Len(t, changes, 1)
		assert.Equal(t, breaking.Dangerous, changes[0].Severity)
	})

	t.Run("private field removed is not breaking", func(t *testing.T) {
		base := makeElements(ced("my-el", withMembers(field("_data", M.Private, "string"))))
		head := makeElements(ced("my-el"))
		changes := findRule("field-removed").Check(base, head)
		assert.Empty(t, changes)
	})
}

func findRule(id string) breaking.BreakingRule {
	for _, r := range breaking.AllRules() {
		if r.ID() == id {
			return r
		}
	}
	panic("rule not found: " + id)
}
