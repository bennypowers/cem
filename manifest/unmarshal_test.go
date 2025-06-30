package manifest

import (
	"reflect"
	"testing"
)

func TestUnmarshalPackage_ClassWithCustomElementTrue_YieldsCustomElementDeclaration(t *testing.T) {
	manifestJSON := []byte(`
{
  "schemaVersion": "1.0.0",
  "modules": [
    {
      "kind": "javascript-module",
      "path": "src/card.js",
      "declarations": [
        {
          "kind": "class",
          "name": "MyCard",
          "customElement": true,
          "tagName": "my-card",
          "members": [
            {
              "kind": "field",
              "name": "foo",
              "type": { "text": "string" }
            }
          ]
        },
        {
          "kind": "class",
          "name": "CardBase",
          "members": [
            {
              "kind": "field",
              "name": "bar",
              "type": { "text": "number" }
            }
          ]
        }
      ]
    }
  ]
}
`)
	pkg, err := UnmarshalPackage(manifestJSON)
	if err != nil {
		t.Fatalf("UnmarshalPackage failed: %v", err)
	}
	mod := pkg.Modules[0]
	if len(mod.Declarations) != 2 {
		t.Fatalf("len(Declarations) = %d, want 2", len(mod.Declarations))
	}

	// Check custom element (should be CustomElementDeclaration)
	ce, ok := mod.Declarations[0].(*CustomElementDeclaration)
	if !ok {
		t.Fatalf("First Declaration is not a CustomElementDeclaration: %T", mod.Declarations[0])
	}
	if ce.Name != "MyCard" || ce.TagName != "my-card" {
		t.Errorf("CustomElementDeclaration: got Name=%q, TagName=%q; want 'MyCard', 'my-card'", ce.Name, ce.TagName)
	}
	if len(ce.Members) != 1 {
		t.Errorf("CustomElementDeclaration members = %+v, want 1 member", ce.Members)
	}
	field, ok := ce.Members[0].(*ClassField)
	if !ok || field.Name != "foo" || field.Type == nil || field.Type.Text != "string" {
		t.Errorf("CustomElementDeclaration first member = %+v, want field named foo of type string", ce.Members[0])
	}

	// Check base class (should be ClassDeclaration)
	cl, ok := mod.Declarations[1].(*ClassDeclaration)
	if !ok {
		t.Fatalf("Second Declaration is not a ClassDeclaration: %T", mod.Declarations[1])
	}
	if cl.Name != "CardBase" {
		t.Errorf("ClassDeclaration: got Name=%q, want 'CardBase'", cl.Name)
	}
	if len(cl.Members) != 1 {
		t.Errorf("ClassDeclaration members = %+v, want 1 member", cl.Members)
	}
	field2, ok := cl.Members[0].(*ClassField)
	if !ok || field2.Name != "bar" || field2.Type == nil || field2.Type.Text != "number" {
		t.Errorf("ClassDeclaration first member = %+v, want field named bar of type number", cl.Members[0])
	}
}

func TestUnmarshalPackage_CustomElement_BasicFields(t *testing.T) {
	manifestJSON := []byte(`
{
  "schemaVersion": "1.0.0",
  "modules": [
    {
      "kind": "javascript-module",
      "path": "src/rh-card.js",
      "declarations": [
        {
          "kind": "custom-element",
          "name": "RhCard",
          "tagName": "rh-card"
        }
      ]
    }
  ]
}
`)
	pkg, err := UnmarshalPackage(manifestJSON)
	if err != nil {
		t.Fatalf("UnmarshalPackage failed: %v", err)
	}
	if len(pkg.Modules) != 1 {
		t.Fatalf("len(Modules) = %d, want 1", len(pkg.Modules))
	}
	mod := pkg.Modules[0]
	if len(mod.Declarations) != 1 {
		t.Fatalf("len(Declarations) = %d, want 1", len(mod.Declarations))
	}
	ce, ok := mod.Declarations[0].(*CustomElementDeclaration)
	if !ok {
		t.Fatalf("Declaration is not a CustomElementDeclaration")
	}
	if ce.TagName != "rh-card" {
		t.Errorf("TagName = %q, want 'rh-card'", ce.TagName)
	}
	if ce.Name != "RhCard" {
		t.Errorf("Name = %q, want 'RhCard'", ce.Name)
	}
}

func TestUnmarshalPackage_CustomElement_Attributes(t *testing.T) {
	manifestJSON := []byte(`
{
  "schemaVersion": "1.0.0",
  "modules": [
    {
      "kind": "javascript-module",
      "path": "src/rh-card.js",
      "declarations": [
        {
          "kind": "custom-element",
          "name": "RhCard",
          "tagName": "rh-card",
          "attributes": [
            {
              "name": "variant",
              "type": { "text": "\"primary\" | \"secondary\"" }
            },
            {
              "name": "elevated",
              "type": { "text": "boolean" }
            }
          ]
        }
      ]
    }
  ]
}
`)
	pkg, err := UnmarshalPackage(manifestJSON)
	if err != nil {
		t.Fatalf("UnmarshalPackage failed: %v", err)
	}
	ce := pkg.Modules[0].Declarations[0].(*CustomElementDeclaration)
	if len(ce.Attributes) != 2 {
		t.Errorf("len(Attributes) = %d, want 2", len(ce.Attributes))
	}
	if ce.Attributes[0].Name != "variant" || ce.Attributes[1].Name != "elevated" {
		t.Errorf("Attribute names = %q,%q, want %q,%q", ce.Attributes[0].Name, ce.Attributes[1].Name, "variant", "elevated")
	}
	if ce.Attributes[0].Type == nil || ce.Attributes[0].Type.Text != "\"primary\" | \"secondary\"" {
		t.Errorf("First attribute type = %+v, want union", ce.Attributes[0].Type)
	}
	if ce.Attributes[1].Type == nil || ce.Attributes[1].Type.Text != "boolean" {
		t.Errorf("Second attribute type = %+v, want boolean", ce.Attributes[1].Type)
	}
}

func TestUnmarshalPackage_CustomElement_Events(t *testing.T) {
	manifestJSON := []byte(`
{
  "schemaVersion": "1.0.0",
  "modules": [
    {
      "kind": "javascript-module",
      "path": "src/rh-card.js",
      "declarations": [
        {
          "kind": "custom-element",
          "name": "RhCard",
          "tagName": "rh-card",
          "events": [
            {
              "name": "rh-card-selected",
              "type": { "text": "CustomEvent" }
            }
          ]
        }
      ]
    }
  ]
}
`)
	pkg, err := UnmarshalPackage(manifestJSON)
	if err != nil {
		t.Fatalf("UnmarshalPackage failed: %v", err)
	}
	ce := pkg.Modules[0].Declarations[0].(*CustomElementDeclaration)
	if len(ce.Events) != 1 || ce.Events[0].Name != "rh-card-selected" {
		t.Errorf("Events = %+v, want 1 rh-card-selected", ce.Events)
	}
}

func TestUnmarshalPackage_CustomElement_Slots(t *testing.T) {
	manifestJSON := []byte(`
{
  "schemaVersion": "1.0.0",
  "modules": [
    {
      "kind": "javascript-module",
      "path": "src/rh-card.js",
      "declarations": [
        {
          "kind": "custom-element",
          "name": "RhCard",
          "tagName": "rh-card",
          "slots": [
            {
              "name": "",
              "description": "Default slot"
            }
          ]
        }
      ]
    }
  ]
}
`)
	pkg, err := UnmarshalPackage(manifestJSON)
	if err != nil {
		t.Fatalf("UnmarshalPackage failed: %v", err)
	}
	ce := pkg.Modules[0].Declarations[0].(*CustomElementDeclaration)
	if len(ce.Slots) != 1 || ce.Slots[0].Description != "Default slot" {
		t.Errorf("Slots = %+v, want 1 Default slot", ce.Slots)
	}
}

func TestUnmarshalPackage_CustomElement_CssPartsPropertiesStates(t *testing.T) {
	manifestJSON := []byte(`
{
  "schemaVersion": "1.0.0",
  "modules": [
    {
      "kind": "javascript-module",
      "path": "src/rh-card.js",
      "declarations": [
        {
          "kind": "custom-element",
          "name": "RhCard",
          "tagName": "rh-card",
          "cssParts": [
            { "name": "header", "description": "Header part" }
          ],
          "cssProperties": [
            { "name": "--rh-card-color", "description": "Card color" }
          ],
          "cssStates": [
            { "name": "--active", "description": "Active state" }
          ]
        }
      ]
    }
  ]
}
`)
	pkg, err := UnmarshalPackage(manifestJSON)
	if err != nil {
		t.Fatalf("UnmarshalPackage failed: %v", err)
	}
	ce := pkg.Modules[0].Declarations[0].(*CustomElementDeclaration)
	if len(ce.CssParts) != 1 || ce.CssParts[0].Name != "header" {
		t.Errorf("CssParts = %+v, want 1 header", ce.CssParts)
	}
	if len(ce.CssProperties) != 1 || ce.CssProperties[0].Name != "--rh-card-color" {
		t.Errorf("CssProperties = %+v, want 1 --rh-card-color", ce.CssProperties)
	}
	if len(ce.CssStates) != 1 || ce.CssStates[0].Name != "--active" {
		t.Errorf("CssStates = %+v, want 1 --active", ce.CssStates)
	}
}

func TestUnmarshalPackage_CustomElement_Demos(t *testing.T) {
	manifestJSON := []byte(`
{
  "schemaVersion": "1.0.0",
  "modules": [
    {
      "kind": "javascript-module",
      "path": "src/rh-card.js",
      "declarations": [
        {
          "kind": "custom-element",
          "name": "RhCard",
          "tagName": "rh-card",
          "demos": [
            { "description": "Basic demo", "url": "demo.html" }
          ]
        }
      ]
    }
  ]
}
`)
	pkg, err := UnmarshalPackage(manifestJSON)
	if err != nil {
		t.Fatalf("UnmarshalPackage failed: %v", err)
	}
	ce := pkg.Modules[0].Declarations[0].(*CustomElementDeclaration)
	if len(ce.Demos) != 1 || ce.Demos[0].Description != "Basic demo" || ce.Demos[0].URL != "demo.html" {
		t.Errorf("Demos = %+v, want 1 Basic demo with url demo.html", ce.Demos)
	}
}

func TestUnmarshalPackage_CustomElement_Mixins(t *testing.T) {
	manifestJSON := []byte(`
{
  "schemaVersion": "1.0.0",
  "modules": [
    {
      "kind": "javascript-module",
      "path": "src/rh-card.js",
      "declarations": [
        {
          "kind": "custom-element",
          "name": "RhCard",
          "tagName": "rh-card",
          "mixins": [
            { "name": "RhMixin", "package": "rh-mixins" }
          ]
        }
      ]
    }
  ]
}
`)
	pkg, err := UnmarshalPackage(manifestJSON)
	if err != nil {
		t.Fatalf("UnmarshalPackage failed: %v", err)
	}
	ce := pkg.Modules[0].Declarations[0].(*CustomElementDeclaration)
	if len(ce.Mixins) != 1 || ce.Mixins[0].Name != "RhMixin" {
		t.Errorf("Mixins = %+v, want 1 RhMixin", ce.Mixins)
	}
}

func TestUnmarshalPackage_CustomElement_Members(t *testing.T) {
	manifestJSON := []byte(`
{
  "schemaVersion": "1.0.0",
  "modules": [
    {
      "kind": "javascript-module",
      "path": "src/rh-card.js",
      "declarations": [
        {
          "kind": "custom-element",
          "name": "RhCard",
          "tagName": "rh-card",
          "members": [
            {
              "kind": "field",
              "name": "prop1",
              "type": { "text": "string" }
            }
          ]
        }
      ]
    }
  ]
}
`)
	pkg, err := UnmarshalPackage(manifestJSON)
	if err != nil {
		t.Fatalf("UnmarshalPackage failed: %v", err)
	}
	ce := pkg.Modules[0].Declarations[0].(*CustomElementDeclaration)
	if len(ce.Members) != 1 {
		t.Errorf("Members = %+v, want 1 item", ce.Members)
	}
	member, ok := ce.Members[0].(*ClassField)
	if !ok {
		t.Errorf("Member = %+v, want class field", member)
	}
	if member.Name != "prop1" || member.Type == nil || member.Type.Text != "string" {
		t.Errorf("Member = %+v, want prop1: string", member)
	}
}

func TestUnmarshalPackage_CustomElement_SliceNils(t *testing.T) {
	manifestJSON := []byte(`
{
  "schemaVersion": "1.0.0",
  "modules": [
    {
      "kind": "javascript-module",
      "path": "src/rh-card.js",
      "declarations": [
        {
          "kind": "custom-element",
          "name": "RhCard",
          "tagName": "rh-card"
        }
      ]
    }
  ]
}
`)
	pkg, err := UnmarshalPackage(manifestJSON)
	if err != nil {
		t.Fatalf("UnmarshalPackage failed: %v", err)
	}
	ce := pkg.Modules[0].Declarations[0].(*CustomElementDeclaration)
	val := reflect.ValueOf(*ce)
	for i := range val.NumField() {
		field := val.Field(i)
		ft := val.Type().Field(i)
		if field.Kind() == reflect.Slice {
			if field.IsNil() {
				t.Errorf("Field %s is nil, want empty slice", ft.Name)
			}
		}
	}
}

func TestUnmarshalPackage_Class(t *testing.T) {
	manifestJSON := []byte(`
{
  "schemaVersion": "1.0.0",
  "modules": [
    {
      "kind": "javascript-module",
      "path": "src/rh-card-base.js",
      "declarations": [
        {
          "kind": "class",
          "name": "RhCardBase",
          "members": [
            {
              "kind": "field",
              "name": "baseProp",
              "type": { "text": "number" }
            }
          ]
        }
      ]
    }
  ]
}
`)
	pkg, err := UnmarshalPackage(manifestJSON)
	if err != nil {
		t.Fatalf("UnmarshalPackage failed: %v", err)
	}
	mod := pkg.Modules[0]
	c, ok := mod.Declarations[0].(*ClassDeclaration)
	if !ok {
		t.Fatalf("Declaration is not a ClassDeclaration")
	}
	if c.Name != "RhCardBase" {
		t.Errorf("Name = %q, want 'RhCardBase'", c.Name)
	}
	if len(c.Members) != 1 {
		t.Errorf("Members = %+v, want 1 member", c.Members)
	}
	member, ok := c.Members[0].(*ClassField)
	if !ok {
		t.Errorf("Member = %+v, want class field", member)
	}
	if member.Name != "baseProp" || member.Type.Text != "number" {
		t.Errorf("Member = %+v, want baseProp: number", member)
	}
	if c.Members == nil || c.Mixins == nil {
		t.Errorf("ClassDeclaration slice fields must not be nil")
	}
}

func TestUnmarshalPackage_Mixin(t *testing.T) {
	manifestJSON := []byte(`
{
  "schemaVersion": "1.0.0",
  "modules": [
    {
      "kind": "javascript-module",
      "path": "src/rh-mixin.js",
      "declarations": [
        {
          "kind": "mixin",
          "name": "RhMixin",
          "members": [
            {
              "kind": "field",
              "name": "mixinProp",
              "type": { "text": "boolean" }
            }
          ]
        }
      ]
    }
  ]
}
`)
	pkg, err := UnmarshalPackage(manifestJSON)
	if err != nil {
		t.Fatalf("UnmarshalPackage failed: %v", err)
	}
	mod := pkg.Modules[0]
	m, ok := mod.Declarations[0].(*MixinDeclaration)
	if !ok {
		t.Fatalf("Declaration is not a MixinDeclaration")
	}
	if m.Name != "RhMixin" {
		t.Errorf("Name = %+v, want 'RhMixin'", m)
	}
	if len(m.Members) != 1 {
		t.Errorf("Members = %+v, want 1 member", m.Members)
	}
	member, ok := m.Members[0].(*ClassField)
	if !ok {
		t.Errorf("Member = %+v, want class field", member)
	}
	if member.Name != "mixinProp" || member.Type == nil || member.Type.Text != "boolean" {
		t.Errorf("Members = %+v, want 1 mixinProp boolean", m.Members)
	}
	if m.Members == nil || m.Mixins == nil {
		t.Errorf("MixinDeclaration slice fields must not be nil")
	}
}

func TestUnmarshalPackage_Function(t *testing.T) {
	manifestJSON := []byte(`
{
  "schemaVersion": "1.0.0",
  "modules": [
    {
      "kind": "javascript-module",
      "path": "src/functions.js",
      "declarations": [
        {
          "kind": "function",
          "name": "helperFn",
          "parameters": [
            { "name": "x", "type": { "text": "number" } }
          ]
        }
      ]
    }
  ]
}
`)
	pkg, err := UnmarshalPackage(manifestJSON)
	if err != nil {
		t.Fatalf("UnmarshalPackage failed: %v", err)
	}
	mod := pkg.Modules[0]
	fn, ok := mod.Declarations[0].(*FunctionDeclaration)
	if !ok {
		t.Fatalf("Declaration is not a FunctionDeclaration")
	}
	if fn.Name != "helperFn" {
		t.Errorf("Name = %q, want 'helperFn'", fn.Name)
	}
	if len(fn.Parameters) != 1 || fn.Parameters[0].Name != "x" || fn.Parameters[0].Type == nil || fn.Parameters[0].Type.Text != "number" {
		t.Errorf("Parameters = %+v, want 1 x number", fn.Parameters)
	}
	if fn.Parameters == nil {
		t.Errorf("FunctionDeclaration.Parameters must not be nil")
	}
}

func TestUnmarshalPackage_Variable(t *testing.T) {
	manifestJSON := []byte(`
{
  "schemaVersion": "1.0.0",
  "modules": [
    {
      "kind": "javascript-module",
      "path": "src/vars.js",
      "declarations": [
        {
          "kind": "variable",
          "name": "CONST_VAL",
          "type": { "text": "string" }
        }
      ]
    }
  ]
}
`)
	pkg, err := UnmarshalPackage(manifestJSON)
	if err != nil {
		t.Fatalf("UnmarshalPackage failed: %v", err)
	}
	mod := pkg.Modules[0]
	v, ok := mod.Declarations[0].(*VariableDeclaration)
	if !ok {
		t.Fatalf("Declaration is not a VariableDeclaration")
	}
	if v.Name != "CONST_VAL" {
		t.Errorf("Name = %q, want 'CONST_VAL'", v.Name)
	}
	if v.Type == nil || v.Type.Text != "string" {
		t.Errorf("Type = %+v, want string", v.Type)
	}
}

func TestUnmarshalPackage_CustomElementMixin(t *testing.T) {
	manifestJSON := []byte(`
{
  "schemaVersion": "1.0.0",
  "modules": [
    {
      "kind": "javascript-module",
      "path": "src/cem.js",
      "declarations": [
        {
          "kind": "custom-element-mixin",
          "name": "RhCustomElementMixin",
          "members": [
            {
              "kind": "field",
              "name": "cemProp",
              "type": { "text": "boolean" }
            }
          ],
          "attributes": [
            { "name": "cem-attr", "type": { "text": "string" } }
          ]
        }
      ]
    }
  ]
}
`)
	pkg, err := UnmarshalPackage(manifestJSON)
	if err != nil {
		t.Fatalf("UnmarshalPackage failed: %v", err)
	}
	mod := pkg.Modules[0]
	cem, ok := mod.Declarations[0].(*CustomElementMixinDeclaration)
	if !ok {
		t.Fatalf("Declaration is not a CustomElementMixinDeclaration")
	}
	if cem.Name != "RhCustomElementMixin" {
		t.Errorf("Name = %q, want 'RhCustomElementMixin'", cem.FunctionLike.Name)
	}
	if len(cem.Members) != 1 {
		t.Errorf("Members = %+v, want 1 item", cem.Members)
	}
	member, ok := cem.Members[0].(*ClassField)
	if !ok {
		t.Errorf("Member = %+v, want class field", member)
	}
	if member.Name != "cemProp" {
		t.Errorf("Members = %+v, want 1 cemProp", cem.Members)
	}
	if cem.Attributes == nil || cem.Members == nil || cem.Mixins == nil {
		t.Errorf("CustomElementMixinDeclaration slice fields must not be nil")
	}
	if len(cem.Attributes) != 1 || cem.Attributes[0].Name != "cem-attr" {
		t.Errorf("Attributes = %+v, want 1 cem-attr", cem.Attributes)
	}
}

func TestUnmarshalPackage_EmptySlices(t *testing.T) {
	manifestJSON := []byte(`
{
  "schemaVersion": "1.0.0",
  "modules": [
    {
      "kind": "javascript-module",
      "path": "src/empty.js",
      "declarations": [
        {
          "kind": "custom-element",
          "name": "EmptyElem",
          "tagName": "empty-elem"
        }
      ]
    }
  ]
}
`)
	pkg, err := UnmarshalPackage(manifestJSON)
	if err != nil {
		t.Fatalf("UnmarshalPackage failed: %v", err)
	}
	mod := pkg.Modules[0]
	ce, ok := mod.Declarations[0].(*CustomElementDeclaration)
	if !ok {
		t.Fatalf("Declaration is not a CustomElementDeclaration")
	}
	// All slices should be non-nil and empty
	val := reflect.ValueOf(*ce)
	for i := range val.NumField() {
		field := val.Field(i)
		ft := val.Type().Field(i)
		if field.Kind() == reflect.Slice {
			if field.IsNil() {
				t.Errorf("Field %s is nil, want empty slice", ft.Name)
			}
		}
	}
}

func TestUnmarshalPackage_InvalidJSON(t *testing.T) {
	_, err := UnmarshalPackage([]byte(`{invalid}`))
	if err == nil {
		t.Fatal("expected error for invalid json, got nil")
	}
}

func TestUnmarshalPackage_UnknownKind(t *testing.T) {
	manifestJSON := []byte(`
{
  "schemaVersion": "1.0.0",
  "modules": [
    {
      "kind": "javascript-module",
      "path": "src/unknown.js",
      "declarations": [
        { "kind": "not-a-real-kind", "foo": "bar" }
      ]
    }
  ]
}
`)
	pkg, err := UnmarshalPackage(manifestJSON)
	if err != nil {
		t.Fatalf("UnmarshalPackage failed: %v", err)
	}
	mod := pkg.Modules[0]
	if len(mod.Declarations) != 0 {
		t.Errorf("len(Declarations) = %d, want 0 for unknown kind", len(mod.Declarations))
	}
}
