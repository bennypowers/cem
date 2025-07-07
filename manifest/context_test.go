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
								FullyQualified: FullyQualified{
									Name: "FooElement",
								},
								Members: []ClassMember{
									&CustomElementField{
										ClassField: ClassField{
											PropertyLike: PropertyLike{
												FullyQualified: FullyQualified{
													Name:    "bar",
													Summary: "bar summary",
												},
											},
										},
										Attribute: "bar",
										Reflects:  true,
									},
									&ClassMethod{
										FullyQualified: FullyQualified{
											Name:    "doAThing",
											Summary: "does something",
										},
										FunctionLike: FunctionLike{
											Return: &Return{Type: &Type{Text: "string"}},
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
								{
									StartByte: 5,
									FullyQualified: FullyQualified{
										Name:    "bar",
										Summary: "bar summary",
									},
								},
								{
									StartByte: 7,
									FullyQualified: FullyQualified{
										Name:    "baz",
										Summary: "baz summary",
									},
								},
							},
							Events: []Event{
								{
									FullyQualified: FullyQualified{
										Name:    "foo-event",
										Summary: "foo event summary",
									},
									Type: &Type{Text: "CustomEvent"},
								},
							},
							Slots: []Slot{
								{FullyQualified: FullyQualified{Name: "", Summary: "default slot"}},
								{FullyQualified: FullyQualified{Name: "named", Summary: "named slot"}},
							},
							CssParts: []CssPart{
								{FullyQualified: FullyQualified{Name: "part1", Summary: "part1 summary"}},
							},
							CssProperties: []CssCustomProperty{
								{FullyQualified: FullyQualified{Name: "--foo-bar", Summary: "foo-bar summary"}, Syntax: "string", Default: "qux", StartByte: 9},
							},
							CssStates: []CssCustomState{
								{FullyQualified: FullyQualified{Name: "--active", Summary: "active state"}},
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

func TestRenderableCustomElementDeclarations(t *testing.T) {
	pkg := makeTestPackage()
	tags := pkg.RenderableCustomElementDeclarations()
	if len(tags) != 1 {
		t.Fatalf("expected 1 tag, got %d", len(tags))
	}
	tag := tags[0]
	if tag.Name() != "foo-el" {
		t.Errorf("expected TagName 'foo-el', got %q", tag.Name())
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

func TestTagRenderableAttributes(t *testing.T) {
	pkg := makeTestPackage()
	attrs, err := pkg.TagRenderableAttributes("foo-el")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(attrs) != 2 {
		t.Fatalf("expected 2 attributes, got %d", len(attrs))
	}
	attrNames := []string{attrs[0].Name(), attrs[1].Name()}
	if !reflect.DeepEqual(attrNames, []string{"bar", "baz"}) && !reflect.DeepEqual(attrNames, []string{"baz", "bar"}) {
		t.Errorf("expected attr names [bar baz], got %v", attrNames)
	}
	for _, attr := range attrs {
		if attr.Name() == "bar" && (attr.CustomElementField == nil || !attr.CustomElementField.Reflects) {
			t.Errorf("expected bar to have CustomElementField with Reflects=true")
		}
	}
}

func TestTagRenderableSlots(t *testing.T) {
	pkg := makeTestPackage()
	slots, err := pkg.TagRenderableSlots("foo-el")
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

func TestTagRenderableCssProperties(t *testing.T) {
	pkg := makeTestPackage()
	props, err := pkg.TagRenderableCssProperties("foo-el")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(props) != 1 {
		t.Fatalf("expected 1 css property, got %d", len(props))
	}
	if props[0].Name() != "--foo-bar" {
		t.Errorf("expected property '--foo-bar', got %q", props[0].Name())
	}
}

func TestTagRenderableCssStates(t *testing.T) {
	pkg := makeTestPackage()
	states, err := pkg.TagRenderableCssStates("foo-el")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(states) != 1 || states[0].Name() != "--active" {
		t.Errorf("expected 1 state '--active', got %+v", states)
	}
}

func TestTagRenderableCssParts(t *testing.T) {
	pkg := makeTestPackage()
	parts, err := pkg.TagRenderableCssParts("foo-el")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(parts) != 1 || parts[0].Name() != "part1" {
		t.Errorf("expected 1 part 'part1', got %+v", parts)
	}
}

func TestTagRenderableEvents(t *testing.T) {
	pkg := makeTestPackage()
	events, err := pkg.TagRenderableEvents("foo-el")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(events) != 1 || events[0].Name() != "foo-event" {
		t.Errorf("expected 1 event 'foo-event', got %+v", events)
	}
}

func TestTagRenderableMethods(t *testing.T) {
	pkg := makeTestPackage()
	methods, err := pkg.TagRenderableMethods("foo-el")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(methods) != 1 || methods[0].Name() != "doAThing" {
		t.Errorf("expected 1 method 'doAThing', got %+v", methods)
	}
}

func TestRenderableAttributeToTableRow(t *testing.T) {
	attr := NewRenderableAttribute(
		&Attribute{
			FullyQualified: FullyQualified{
				Name:    "foo",
				Summary: "attr summary",
			},
		},
		&CustomElementDeclaration{
			ClassDeclaration: ClassDeclaration{
				ClassLike: ClassLike{
					Members: []ClassMember{
						&CustomElementField{
							Attribute: "foo",
							Reflects:   true,
							ClassField: ClassField{
								PropertyLike: PropertyLike{
									FullyQualified: FullyQualified{
										Name: "foo",
									},
								},
							},
						},
					},
				},
			},
		},
		nil,
		nil,
	)
	row := attr.ToTableRow()
	if !reflect.DeepEqual(row, []string{"foo", "foo", "✅", "attr summary", "", ""}) {
		t.Errorf("unexpected RenderableAttribute ToTableRow: %#v", row)
	}
}

func TestRenderableSlotToTableRow(t *testing.T) {
	slot := RenderableSlot{
		Slot: &Slot{
			FullyQualified: FullyQualified{
				Summary: "default",
			},
		},
	}
	row := slot.ToTableRow()
	if !reflect.DeepEqual(row, []string{"<default>", "default"}) {
		t.Errorf("unexpected RenderableSlot ToTableRow: %#v", row)
	}
}

func TestRenderableCssCustomPropertyToTableRow(t *testing.T) {
	cssProp := RenderableCssCustomProperty{
		CssCustomProperty: &CssCustomProperty{
			Syntax: "string",
			Default: "bar",
			FullyQualified: FullyQualified{
				Name: "--foo",
				Summary: "baz",
			},
		},
	}
	row := cssProp.ToTableRow()
	if !reflect.DeepEqual(row, []string{"--foo", "string", "bar", "baz"}) {
		t.Errorf("unexpected RenderableCssCustomProperty ToTableRow: %#v", row)
	}
}

func TestRenderableCssCustomStateToTableRow(t *testing.T) {
	state := RenderableCssCustomState{
		CssCustomState: &CssCustomState{
			FullyQualified: FullyQualified{
				Name: "--active",
				Summary: "active state",
			},
		},
	}
	row := state.ToTableRow()
	if !reflect.DeepEqual(row, []string{"--active", "active state"}) {
		t.Errorf("unexpected RenderableCssCustomState ToTableRow: %#v", row)
	}
}

func TestRenderableCssPartToTableRow(t *testing.T) {
	part := RenderableCssPart{
		CssPart: &CssPart{
			FullyQualified: FullyQualified{
				Name: "part1",
				Summary: "part summary",
			},
		},
	}
	row := part.ToTableRow()
	if !reflect.DeepEqual(row, []string{"part1", "part summary"}) {
		t.Errorf("unexpected RenderableCssPart ToTableRow: %#v", row)
	}
}

func TestRenderableEventToTableRow(t *testing.T) {
	event := RenderableEvent{
		Event: &Event{
			Type: &Type{
				Text: "CustomEvent",
			},
			FullyQualified: FullyQualified{
				Name: "evt",
				Summary: "summ",
			},
		},
	}
	row := event.ToTableRow()
	if !reflect.DeepEqual(row, []string{"evt", "CustomEvent", "summ"}) {
		t.Errorf("unexpected RenderableEvent ToTableRow: %#v", row)
	}
}

func TestRenderableMethodToTableRow(t *testing.T) {
	method := RenderableClassMethod{
		Method: &ClassMethod{
			FunctionLike: FunctionLike{
				Return: &Return{Type: &Type{Text: "string"}},
			},
			FullyQualified: FullyQualified{
				Name: "doIt",
				Summary: "summ",
			},
			Privacy: "private",
			Static:  false,
		},
	}
	row := method.ToTableRow()
	if !reflect.DeepEqual(row, []string{"doIt", "string", "private", "false", "summ"}) {
		t.Errorf("unexpected RenderableMethod ToTableRow: %#v", row)
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
	_, err := pkg.TagRenderableAttributes("does-not-exist")
	if err == nil {
		t.Error("expected error for missing tag, got nil")
	}
}
