package manifest

import (
	"reflect"
	"testing"
)

func makeTestPackage() *Package {
	// Minimal manifest with a custom element and all supported details
	return &Package{
		Modules: []Module{
			{
				Kind: "javascript-module",
				Path: "src/foo.js",
				Declarations: []Declaration{
					&CustomElementDeclaration{
						ClassDeclaration: ClassDeclaration{
							ClassLike: ClassLike{
								Name:    "FooElement",
								Members: []ClassMember{
									&CustomElementField{
										ClassField: ClassField{
											PropertyLike: PropertyLike{
												Name:    "bar",
												Summary: "bar summary",
											},
										},
										Attribute: "bar",
										Reflects:  true,
									},
									&ClassMethod{
										FunctionLike: FunctionLike{
											Name:    "doAThing",
											Summary: "does something",
											Return:  &Return{Type: &Type{Text: "string"}},
										},
										Kind:    "method",
										Privacy: Public,
										Static:  true,
									},
								},
							},
							Kind: "class",
						},
						CustomElement: CustomElement{
							TagName: "foo-el",
							Attributes: []Attribute{
								{Name: "bar", Summary: "bar summary", StartByte: 5},
								{Name: "baz", Summary: "baz summary", StartByte: 7},
							},
							Events: []Event{
								{Name: "foo-event", Summary: "foo event summary", Type: &Type{Text: "CustomEvent"}},
							},
							Slots: []Slot{
								{Name: "", Summary: "default slot"},
								{Name: "named", Summary: "named slot"},
							},
							CssParts: []CssPart{
								{Name: "part1", Summary: "part1 summary"},
							},
							CssProperties: []CssCustomProperty{
								{Name: "--foo-bar", Syntax: "string", Default: "qux", Summary: "foo-bar summary", StartByte: 9},
							},
							CssStates: []CssCustomState{
								{Name: "--active", Summary: "active state"},
							},
							Demos:         nil,
							CustomElement: true,
						},
					},
				},
				Exports: []Export{
					&CustomElementExport{
						Name:      "foo-el",
						StartByte: 1,
					},
				},
			},
		},
	}
}

func TestGetAllTagNamesWithContext(t *testing.T) {
	pkg := makeTestPackage()
	tags := pkg.GetAllTagNamesWithContext()
	if len(tags) != 1 {
		t.Fatalf("expected 1 tag, got %d", len(tags))
	}
	tag := tags[0]
	if tag.TagName != "foo-el" {
		t.Errorf("expected TagName 'foo-el', got %q", tag.TagName)
	}
	if tag.Module == nil || tag.Module.Path != "src/foo.js" {
		t.Errorf("expected Module path src/foo.js, got %+v", tag.Module)
	}
	if tag.CustomElementDeclaration == nil {
		t.Error("expected CustomElementDeclaration to be set")
	}
	if tag.CustomElementExport == nil {
		t.Error("expected CustomElementExport to be set")
	}
}

func TestGetTagAttrsWithContext(t *testing.T) {
	pkg := makeTestPackage()
	attrs, err := pkg.GetTagAttrsWithContext("foo-el")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(attrs) != 2 {
		t.Fatalf("expected 2 attributes, got %d", len(attrs))
	}
	attrNames := []string{attrs[0].Name, attrs[1].Name}
	if !reflect.DeepEqual(attrNames, []string{"bar", "baz"}) && !reflect.DeepEqual(attrNames, []string{"baz", "bar"}) {
		t.Errorf("expected attr names [bar baz], got %v", attrNames)
	}
	for _, attr := range attrs {
		if attr.Name == "bar" && (attr.CustomElementField == nil || !attr.CustomElementField.Reflects) {
			t.Errorf("expected bar to have CustomElementField with Reflects=true")
		}
	}
}

func TestGetTagSlotsWithContext(t *testing.T) {
	pkg := makeTestPackage()
	slots, err := pkg.GetTagSlotsWithContext("foo-el")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(slots) != 2 {
		t.Fatalf("expected 2 slots, got %d", len(slots))
	}
	if slots[0].Slot == nil || slots[1].Slot == nil {
		t.Error("expected Slot to be set in all results")
	}
}

func TestGetTagCssPropertiesWithContext(t *testing.T) {
	pkg := makeTestPackage()
	props, err := pkg.GetTagCssPropertiesWithContext("foo-el")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(props) != 1 {
		t.Fatalf("expected 1 css property, got %d", len(props))
	}
	if props[0].Name != "--foo-bar" {
		t.Errorf("expected property '--foo-bar', got %q", props[0].Name)
	}
}

func TestGetTagCssStatesWithContext(t *testing.T) {
	pkg := makeTestPackage()
	states, err := pkg.GetTagCssStatesWithContext("foo-el")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(states) != 1 || states[0].Name != "--active" {
		t.Errorf("expected 1 state '--active', got %+v", states)
	}
}

func TestGetTagCssPartsWithContext(t *testing.T) {
	pkg := makeTestPackage()
	parts, err := pkg.GetTagCssPartsWithContext("foo-el")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(parts) != 1 || parts[0].Name != "part1" {
		t.Errorf("expected 1 part 'part1', got %+v", parts)
	}
}

func TestGetTagEventsWithContext(t *testing.T) {
	pkg := makeTestPackage()
	events, err := pkg.GetTagEventsWithContext("foo-el")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(events) != 1 || events[0].Name != "foo-event" {
		t.Errorf("expected 1 event 'foo-event', got %+v", events)
	}
}

func TestGetTagMethodsWithContext(t *testing.T) {
	pkg := makeTestPackage()
	methods, err := pkg.GetTagMethodsWithContext("foo-el")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(methods) != 1 || methods[0].Name != "doAThing" {
		t.Errorf("expected 1 method 'doAThing', got %+v", methods)
	}
}

func TestContextToTableRow(t *testing.T) {
	// AttributeWithContext
	attr := AttributeWithContext{
		Name: "foo",
		Attribute: &Attribute{
			Name:    "foo",
			Summary: "attr summary",
		},
		CustomElementField: &CustomElementField{
			ClassField: ClassField{PropertyLike: PropertyLike{Name: "foo"}},
			Reflects:   true,
		},
	}
	row := attr.ToTableRow()
	if !reflect.DeepEqual(row, []string{"foo", "foo", "✅", "attr summary"}) {
		t.Errorf("unexpected AttributeWithContext ToTableRow: %#v", row)
	}

	// SlotWithContext
	slot := SlotWithContext{
		Name: "",
		Slot: &Slot{Summary: "default"},
	}
	row = slot.ToTableRow()
	if !reflect.DeepEqual(row, []string{"<default>", "default"}) {
		t.Errorf("unexpected SlotWithContext ToTableRow: %#v", row)
	}

	// CssCustomPropertyWithContext
	cssProp := CssCustomPropertyWithContext{
		Name:              "--foo",
		CssCustomProperty: &CssCustomProperty{Syntax: "string", Default: "bar", Summary: "baz"},
	}
	row = cssProp.ToTableRow()
	if !reflect.DeepEqual(row, []string{"--foo", "string", "bar", "baz"}) {
		t.Errorf("unexpected CssCustomPropertyWithContext ToTableRow: %#v", row)
	}

	// CssCustomStateWithContext
	state := CssCustomStateWithContext{
		Name:           "--active",
		CssCustomState: &CssCustomState{Summary: "active"},
	}
	row = state.ToTableRow()
	if !reflect.DeepEqual(row, []string{"--active", "active"}) {
		t.Errorf("unexpected CssCustomStateWithContext ToTableRow: %#v", row)
	}

	// CssPartWithContext
	part := CssPartWithContext{
		Name:    "part1",
		CssPart: &CssPart{Summary: "part summary"},
	}
	row = part.ToTableRow()
	if !reflect.DeepEqual(row, []string{"part1", "part summary"}) {
		t.Errorf("unexpected CssPartWithContext ToTableRow: %#v", row)
	}

	// EventWithContext
	event := EventWithContext{
		Name:  "evt",
		Event: &Event{Type: &Type{Text: "CustomEvent"}, Summary: "summ"},
	}
	row = event.ToTableRow()
	if !reflect.DeepEqual(row, []string{"evt", "CustomEvent", "summ"}) {
		t.Errorf("unexpected EventWithContext ToTableRow: %#v", row)
	}

	// MethodWithContext
	method := MethodWithContext{
		Name: "doIt",
		Method: &ClassMethod{
			FunctionLike: FunctionLike{
				Return:  &Return{Type: &Type{Text: "string"}},
				Summary: "summ",
			},
			Privacy: "private",
			Static:  false,
		},
	}
	row = method.ToTableRow()
	if !reflect.DeepEqual(row, []string{"doIt", "string", "private", "false", "summ"}) {
		t.Errorf("unexpected MethodWithContext ToTableRow: %#v", row)
	}
}

func TestFindCustomElementContext_TagNotFound(t *testing.T) {
	pkg := makeTestPackage()
	_, _, _, err := pkg.findCustomElementContext("does-not-exist")
	if err == nil {
		t.Error("expected error for missing tag, got nil")
	}
}

func TestGetTagAttrsWithContext_TagNotFound(t *testing.T) {
	pkg := makeTestPackage()
	_, err := pkg.GetTagAttrsWithContext("does-not-exist")
	if err == nil {
		t.Error("expected error for missing tag, got nil")
	}
}
