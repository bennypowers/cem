package manifest

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

// DRY Test Helpers

func loadFixture(t *testing.T, name string) []byte {
	t.Helper()
	data, err := os.ReadFile(filepath.Join("fixtures", name))
	if err != nil {
		t.Fatalf("failed to load fixture %s: %v", name, err)
	}
	return data
}

func mustUnmarshalPackage(t *testing.T, manifestJSON []byte) *Package {
	t.Helper()
	pkg, err := UnmarshalPackage(manifestJSON)
	if err != nil {
		t.Fatalf("UnmarshalPackage failed: %v", err)
	}
	return pkg
}

func mustFirstModule(t *testing.T, pkg *Package) *Module {
	t.Helper()
	if len(pkg.Modules) == 0 {
		t.Fatalf("No modules found in the package")
	}
	return &pkg.Modules[0]
}

func mustModuleDecls(t *testing.T, mod *Module, want int) []Declaration {
	t.Helper()
	if len(mod.Declarations) != want {
		t.Fatalf("len(Declarations) = %d, want %d", len(mod.Declarations), want)
	}
	return mod.Declarations
}

func mustCustomElementDecl(t *testing.T, decl Declaration) *CustomElementDeclaration {
	t.Helper()
	ce, ok := decl.(*CustomElementDeclaration)
	if !ok {
		t.Fatalf("Declaration is not a CustomElementDeclaration: %T", decl)
	}
	return ce
}

func mustClassDecl(t *testing.T, decl Declaration) *ClassDeclaration {
	t.Helper()
	cl, ok := decl.(*ClassDeclaration)
	if !ok {
		t.Fatalf("Declaration is not a ClassDeclaration: %T", decl)
	}
	return cl
}

func mustMixinDecl(t *testing.T, decl Declaration) *MixinDeclaration {
	t.Helper()
	m, ok := decl.(*MixinDeclaration)
	if !ok {
		t.Fatalf("Declaration is not a MixinDeclaration: %T", decl)
	}
	return m
}

func mustFunctionDecl(t *testing.T, decl Declaration) *FunctionDeclaration {
	t.Helper()
	fn, ok := decl.(*FunctionDeclaration)
	if !ok {
		t.Fatalf("Declaration is not a FunctionDeclaration: %T", decl)
	}
	return fn
}

func mustVariableDecl(t *testing.T, decl Declaration) *VariableDeclaration {
	t.Helper()
	v, ok := decl.(*VariableDeclaration)
	if !ok {
		t.Fatalf("Declaration is not a VariableDeclaration: %T", decl)
	}
	return v
}

func mustCustomElementMixinDecl(t *testing.T, decl Declaration) *CustomElementMixinDeclaration {
	t.Helper()
	cem, ok := decl.(*CustomElementMixinDeclaration)
	if !ok {
		t.Fatalf("Declaration is not a CustomElementMixinDeclaration: %T", decl)
	}
	return cem
}

func mustClassField(t *testing.T, member any) *ClassField {
	t.Helper()
	f, ok := member.(*ClassField)
	if !ok {
		t.Fatalf("Member = %T, want *ClassField", member)
	}
	return f
}

func mustCustomElementField(t *testing.T, member any) *CustomElementField {
	t.Helper()
	cem, ok := member.(*CustomElementField)
	if !ok {
		t.Fatalf("Member = %T, want *CustomElementField", member)
	}
	return cem
}

// --- Tests ---

func TestUnmarshalPackage_ClassWithCustomElementTrue_YieldsCustomElementDeclaration(t *testing.T) {
	manifestJSON := loadFixture(t, "class_with_custom_element_true.json")
	pkg := mustUnmarshalPackage(t, manifestJSON)
	mod := mustFirstModule(t, pkg)
	decls := mustModuleDecls(t, mod, 2)

	ce := mustCustomElementDecl(t, decls[0])
	if ce.Name != "MyCard" || ce.TagName != "my-card" {
		t.Errorf("CustomElementDeclaration: got Name=%q, TagName=%q; want 'MyCard', 'my-card'", ce.Name, ce.TagName)
	}
	if len(ce.Members) != 1 {
		t.Errorf("CustomElementDeclaration members = %+v, want 1 member", ce.Members)
	}
	field := mustClassField(t, ce.Members[0])
	if field.Name != "foo" || field.Type == nil || field.Type.Text != "string" {
		t.Errorf("CustomElementDeclaration first member = %+v, want field named foo of type string", ce.Members[0])
	}

	cl := mustClassDecl(t, decls[1])
	if cl.Name != "CardBase" {
		t.Errorf("ClassDeclaration: got Name=%q, want 'CardBase'", cl.Name)
	}
	field2 := mustClassField(t, cl.Members[0])
	if field2.Name != "bar" || field2.Type == nil || field2.Type.Text != "number" {
		t.Errorf("ClassDeclaration first member = %+v, want field named bar of type number", cl.Members[0])
	}
}

func TestUnmarshalPackage_CustomElement_BasicFields(t *testing.T) {
	manifestJSON := loadFixture(t, "custom_element_basic_fields.json")
	pkg := mustUnmarshalPackage(t, manifestJSON)
	mod := mustFirstModule(t, pkg)
	ce := mustCustomElementDecl(t, mustModuleDecls(t, mod, 1)[0])
	if ce.TagName != "my-card" {
		t.Errorf("TagName = %q, want 'my-card'", ce.TagName)
	}
	if ce.Name != "MyCard" {
		t.Errorf("Name = %q, want 'MyCard'", ce.Name)
	}
}

func TestUnmarshalPackage_CustomElement_MemberWithAttributeYieldsCustomElementField(t *testing.T) {
	manifestJSON := loadFixture(t, "custom_element_member_with_attribute.json")
	pkg := mustUnmarshalPackage(t, manifestJSON)
	ce := mustCustomElementDecl(t, mustFirstModule(t, pkg).Declarations[0])
	if len(ce.Members) != 2 {
		t.Fatalf("Members = %+v, want 2 items", ce.Members)
	}
	cem := mustCustomElementField(t, ce.Members[0])
	if cem.Name != "open" {
		t.Errorf("CustomElementField: want name=open got '%s'", cem.Name)
	}
	if cem.Attribute != "open" {
		t.Errorf("CustomElementField: want attribute=open got '%s'", cem.Attribute)
	}
	if cem.Type == nil {
		t.Errorf("CustomElementField: want non-nil type got '%+v'", cem.Type)
	}
	if  cem.Type.Text != "boolean" {
		t.Errorf("CustomElementField: want type.text=boolean got '%s'", cem.Type.Text)
	}
	if cem.Default != "false" {
		t.Errorf("CustomElementField: want default=false got '%s'", cem.Default)
	}
	if !cem.Reflects {
		t.Errorf("CustomElementField: want reflects=true got '%+v'", cem.Reflects)
	}
	plain := mustClassField(t, ce.Members[1])
	if plain.Name != "plainField" || plain.Type == nil || plain.Type.Text != "string" {
		t.Errorf("ClassField = %+v, want name=plainField type=string", plain)
	}
}

func TestUnmarshalPackage_CustomElement_Attributes(t *testing.T) {
	manifestJSON := loadFixture(t, "custom_element_attributes.json")
	pkg := mustUnmarshalPackage(t, manifestJSON)
	ce := mustCustomElementDecl(t, mustFirstModule(t, pkg).Declarations[0])
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
	manifestJSON := loadFixture(t, "custom_element_events.json")
	pkg := mustUnmarshalPackage(t, manifestJSON)
	ce := mustCustomElementDecl(t, mustFirstModule(t, pkg).Declarations[0])
	if len(ce.Events) != 1 || ce.Events[0].Name != "my-card-selected" {
		t.Errorf("Events = %+v, want 1 my-card-selected", ce.Events)
	}
}

func TestUnmarshalPackage_CustomElement_Slots(t *testing.T) {
	manifestJSON := loadFixture(t, "custom_element_slots.json")
	pkg := mustUnmarshalPackage(t, manifestJSON)
	ce := mustCustomElementDecl(t, mustFirstModule(t, pkg).Declarations[0])
	if len(ce.Slots) != 1 || ce.Slots[0].Description != "Default slot" {
		t.Errorf("Slots = %+v, want 1 Default slot", ce.Slots)
	}
}

func TestUnmarshalPackage_CustomElement_CssPartsPropertiesStates(t *testing.T) {
	manifestJSON := loadFixture(t, "custom_element_css_parts_properties_states.json")
	pkg := mustUnmarshalPackage(t, manifestJSON)
	ce := mustCustomElementDecl(t, mustFirstModule(t, pkg).Declarations[0])
	if len(ce.CssParts) != 1 || ce.CssParts[0].Name != "header" {
		t.Errorf("CssParts = %+v, want 1 header", ce.CssParts)
	}
	if len(ce.CssProperties) != 1 || ce.CssProperties[0].Name != "--my-card-color" {
		t.Errorf("CssProperties = %+v, want 1 --my-card-color", ce.CssProperties)
	}
	if len(ce.CssStates) != 1 || ce.CssStates[0].Name != "--active" {
		t.Errorf("CssStates = %+v, want 1 --active", ce.CssStates)
	}
}

func TestUnmarshalPackage_CustomElement_Demos(t *testing.T) {
	manifestJSON := loadFixture(t, "custom_element_demos.json")
	pkg := mustUnmarshalPackage(t, manifestJSON)
	ce := mustCustomElementDecl(t, mustFirstModule(t, pkg).Declarations[0])
	if len(ce.Demos) != 1 || ce.Demos[0].Description != "Basic demo" || ce.Demos[0].URL != "demo.html" {
		t.Errorf("Demos = %+v, want 1 Basic demo with url demo.html", ce.Demos)
	}
}

func TestUnmarshalPackage_CustomElement_Mixins(t *testing.T) {
	manifestJSON := loadFixture(t, "custom_element_mixins.json")
	pkg := mustUnmarshalPackage(t, manifestJSON)
	ce := mustCustomElementDecl(t, mustFirstModule(t, pkg).Declarations[0])
	if len(ce.Mixins) != 1 || ce.Mixins[0].Name != "MyMixin" {
		t.Errorf("Mixins = %+v, want 1 MyMixin", ce.Mixins)
	}
}

func TestUnmarshalPackage_CustomElement_Members(t *testing.T) {
	manifestJSON := loadFixture(t, "custom_element_members.json")
	pkg := mustUnmarshalPackage(t, manifestJSON)
	ce := mustCustomElementDecl(t, mustFirstModule(t, pkg).Declarations[0])
	if len(ce.Members) != 1 {
		t.Errorf("Members = %+v, want 1 item", ce.Members)
	}
	member := mustClassField(t, ce.Members[0])
	if member.Name != "prop1" || member.Type == nil || member.Type.Text != "string" {
		t.Errorf("Member = %+v, want prop1: string", member)
	}
}

func TestUnmarshalPackage_CustomElement_SliceNils(t *testing.T) {
	manifestJSON := loadFixture(t, "custom_element_slice_nils.json")
	pkg := mustUnmarshalPackage(t, manifestJSON)
	ce := mustCustomElementDecl(t, mustFirstModule(t, pkg).Declarations[0])
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
	manifestJSON := loadFixture(t, "class.json")
	pkg := mustUnmarshalPackage(t, manifestJSON)
	mod := mustFirstModule(t, pkg)
	c := mustClassDecl(t, mustModuleDecls(t, mod, 1)[0])
	if c.Name != "MyCardBase" {
		t.Errorf("Name = %q, want 'MyCardBase'", c.Name)
	}
	if len(c.Members) != 1 {
		t.Errorf("Members = %+v, want 1 member", c.Members)
	}
	member := mustClassField(t, c.Members[0])
	if member.Name != "baseProp" || member.Type.Text != "number" {
		t.Errorf("Member = %+v, want baseProp: number", member)
	}
	if c.Members == nil || c.Mixins == nil {
		t.Errorf("ClassDeclaration slice fields must not be nil")
	}
}

func TestUnmarshalPackage_Mixin(t *testing.T) {
	manifestJSON := loadFixture(t, "mixin.json")
	pkg := mustUnmarshalPackage(t, manifestJSON)
	mod := mustFirstModule(t, pkg)
	m := mustMixinDecl(t, mustModuleDecls(t, mod, 1)[0])
	if m.Name != "MyMixin" {
		t.Errorf("Name = %+v, want 'MyMixin'", m)
	}
	if len(m.Members) != 1 {
		t.Errorf("Members = %+v, want 1 member", m.Members)
	}
	member := mustClassField(t, m.Members[0])
	if member.Name != "mixinProp" || member.Type == nil || member.Type.Text != "boolean" {
		t.Errorf("Members = %+v, want 1 mixinProp boolean", m.Members)
	}
}

func TestUnmarshalPackage_Function(t *testing.T) {
	manifestJSON := loadFixture(t, "function.json")
	pkg := mustUnmarshalPackage(t, manifestJSON)
	mod := mustFirstModule(t, pkg)
	fn := mustFunctionDecl(t, mustModuleDecls(t, mod, 1)[0])
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
	manifestJSON := loadFixture(t, "variable.json")
	pkg := mustUnmarshalPackage(t, manifestJSON)
	mod := mustFirstModule(t, pkg)
	v := mustVariableDecl(t, mustModuleDecls(t, mod, 1)[0])
	if v.Name != "CONST_VAL" {
		t.Errorf("Name = %q, want 'CONST_VAL'", v.Name)
	}
	if v.Type == nil || v.Type.Text != "string" {
		t.Errorf("Type = %+v, want string", v.Type)
	}
}

func TestUnmarshalPackage_CustomElementMixin(t *testing.T) {
	manifestJSON := loadFixture(t, "custom_element_mixin.json")
	pkg := mustUnmarshalPackage(t, manifestJSON)
	mod := mustFirstModule(t, pkg)
	cem := mustCustomElementMixinDecl(t, mustModuleDecls(t, mod, 1)[0])
	if cem.Name != "MyCustomElementMixin" {
		t.Errorf("Name = %q, want 'MyCustomElementMixin'", cem.FunctionLike.Name)
	}
	if len(cem.Members) != 1 {
		t.Errorf("Members = %+v, want 1 item", cem.Members)
	}
	member := mustClassField(t, cem.Members[0])
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
	manifestJSON := loadFixture(t, "empty_slices.json")
	pkg := mustUnmarshalPackage(t, manifestJSON)
	ce := mustCustomElementDecl(t, mustFirstModule(t, pkg).Declarations[0])
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
	manifestJSON := loadFixture(t, "unknown_kind.json")
	pkg := mustUnmarshalPackage(t, manifestJSON)
	mod := mustFirstModule(t, pkg)
	if len(mod.Declarations) != 0 {
		t.Errorf("len(Declarations) = %d, want 0 for unknown kind", len(mod.Declarations))
	}
}
