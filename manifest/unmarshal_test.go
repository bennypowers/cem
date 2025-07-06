package manifest

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"testing"
)

// DRY Test Helpers

// Helper: Skip test if FIXTURE_PATTERN is set and doesn't match the fixture name.
func mustRunFixture(t *testing.T) {
	pattern := os.Getenv("FIXTURE_PATTERN")
	if pattern == "" {
		return
	}
	matched, err := regexp.MatchString(pattern, t.Name())
	if err != nil {
		t.Fatalf("Invalid FIXTURE_PATTERN: %v", err)
	}
	if !matched {
		t.Skipf("Skipping because FIXTURE_PATTERN=%q does not match fixture %q", pattern, t.Name())
	}
}

func loadFixture(t *testing.T, name string) []byte {
	t.Helper()
	data, err := os.ReadFile(filepath.Join("fixtures", name))
	if err != nil {
		t.Fatalf("failed to load fixture %s: %v", name, err)
	}
	return data
}

func mustUnmarshalPackage(t *testing.T, data []byte) *Package {
	t.Helper()
	var pkg Package
	if err := json.Unmarshal(data, &pkg); err != nil {
		t.Fatalf("UnmarshalPackage failed: %v", err)
	}
	return &pkg
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

func getFirstDecl(t *testing.T, pkg *Package) any {
	t.Helper()
	if len(pkg.Modules) == 0 {
		t.Fatal("no modules in package")
	}
	if len(pkg.Modules[0].Declarations) == 0 {
		t.Fatal("no declarations in first module")
	}
	return pkg.Modules[0].Declarations[0]
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

func mustUnmarshalPackageEdge(t *testing.T, data []byte) (*Package, error) {
	t.Helper()
	var pkg Package
	if err := json.Unmarshal(data, &pkg); err != nil {
		if regexp.MustCompile(`unknown declaration kind`).MatchString(err.Error()) {
			return nil, err // let test expect/verify this
		}
		if regexp.MustCompile(`invalid type for deprecated field`).MatchString(err.Error()) {
			return nil, err // let test expect/verify this
		}
		t.Log(err)
		t.Fatalf("UnmarshalPackage failed: %v", err)
	}
	return &pkg, nil
}

// --- Tests ---

func TestUnmarshalPackage(t *testing.T) {
	t.Run("CustomElement", func(t *testing.T) {
		t.Run("ClassWithCustomElementTrueYieldsCustomElementDeclaration", func(t *testing.T) {
			mustRunFixture(t)
			data := loadFixture(t, "class-with-custom-element-true.json")
			t.Log(string(data))
			pkg := mustUnmarshalPackage(t, data)
			t.Log(pkg.Modules[0].Declarations[0])
			mod := mustFirstModule(t, pkg)
			decls := mustModuleDecls(t, mod, 2)

			ce := mustCustomElementDecl(t, decls[0])
			if ce.Name != "MyCard" || ce.TagName != "my-card" {
				t.Errorf("CustomElementDeclaration: got Name=%q, TagName=%q; want 'MyCard', 'my-card'", ce.Name, ce.TagName)
			}
			if len(ce.Members) != 1 {
				t.Fatalf("CustomElementDeclaration members = %+v, want 1 member", ce.Members)
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
		})

		t.Run("BasicFields", func(t *testing.T) {
			mustRunFixture(t)
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
		})

		t.Run("MemberWithAttributeYieldsCustomElementField", func(t *testing.T) {
			mustRunFixture(t)
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
		})

		t.Run("Attributes", func(t *testing.T) {
			mustRunFixture(t)
			manifestJSON := loadFixture(t, "custom_element_attributes.json")
			pkg := mustUnmarshalPackage(t, manifestJSON)
			ce := mustCustomElementDecl(t, mustFirstModule(t, pkg).Declarations[0])
			if len(ce.Attributes) != 2 {
				t.Fatalf("len(Attributes) = %d, want 2", len(ce.Attributes))
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
			t.Run("Attribute", func(t *testing.T) {
				t.Run("TypeText", func(t *testing.T) {
					data := loadFixture(t, "attribute-type-text.json")
					pkg := mustUnmarshalPackage(t, data)
					ce, ok := getFirstDecl(t, pkg).(*CustomElementDeclaration)
					if !ok {
						t.Fatalf("not a CustomElementDeclaration: %#v", getFirstDecl(t, pkg))
					}
					if len(ce.Attributes) != 1 || ce.Attributes[0].Type.Text != "number" {
						t.Errorf("unexpected attribute type: %#v", ce.Attributes)
					}
				})
				t.Run("Deprecation", func(t *testing.T) {
					t.Run("None", func(t *testing.T) {
						mustRunFixture(t)
						pkg := mustUnmarshalPackage(t, loadFixture(t, "custom-element-attr-deprecated-none.json"))
						mod := mustFirstModule(t, pkg)
						ce := mustCustomElementDecl(t, mustModuleDecls(t, mod, 1)[0])
						if len(ce.Attributes) != 1 {
							t.Fatalf("len(Attributes) = %d, want 1", len(ce.Attributes))
						}
						if ce.Attributes[0].Deprecated != nil {
							t.Errorf("Attribute.Deprecated = %v, want nil", ce.Attributes[0].Deprecated)
						}
					})

					t.Run("Bool", func(t *testing.T) {
						mustRunFixture(t)
						pkg := mustUnmarshalPackage(t, loadFixture(t, "custom-element-attr-deprecated-bool.json"))
						mod := mustFirstModule(t, pkg)
						ce := mustCustomElementDecl(t, mustModuleDecls(t, mod, 1)[0])
						if len(ce.Attributes) != 1 {
							t.Fatalf("len(Attributes) = %d, want 1", len(ce.Attributes))
						}
						dep := ce.Attributes[0].Deprecated
						if dep == nil {
							t.Errorf("Attribute.Deprecated = nil, want non-nil")
						} else if v, ok := dep.(DeprecatedFlag); !ok || !bool(v) {
							t.Errorf("Attribute.Deprecated = %#v, want DeprecatedFlag(true)", dep)
						}
					})

					t.Run("Reason", func(t *testing.T) {
						mustRunFixture(t)
						pkg := mustUnmarshalPackage(t, loadFixture(t, "custom-element-attr-deprecated-reason.json"))
						mod := mustFirstModule(t, pkg)
						ce := mustCustomElementDecl(t, mustModuleDecls(t, mod, 1)[0])
						if len(ce.Attributes) != 1 {
							t.Fatalf("len(Attributes) = %d, want 1", len(ce.Attributes))
						}
						dep := ce.Attributes[0].Deprecated
						if dep == nil {
							t.Errorf("Attribute.Deprecated = nil, want non-nil")
						} else if v, ok := dep.(DeprecatedReason); !ok || string(v) != "use something else" {
							t.Errorf("Attribute.Deprecated = %#v, want DeprecatedReason(\"use something else\")", dep)
						}
					})
				})
			})
			t.Run("Event", func(t *testing.T) {
				mustRunFixture(t)
				manifestJSON := loadFixture(t, "custom_element_events.json")
				pkg := mustUnmarshalPackage(t, manifestJSON)
				ce := mustCustomElementDecl(t, mustFirstModule(t, pkg).Declarations[0])
				if len(ce.Events) != 1 || ce.Events[0].Name != "my-card-selected" {
					t.Errorf("Events = %+v, want 1 my-card-selected", ce.Events)
				}
				t.Run("TypeText", func(t *testing.T) {
					data := loadFixture(t, "event-type-text.json")
					pkg := mustUnmarshalPackage(t, data)
					ce, ok := getFirstDecl(t, pkg).(*CustomElementDeclaration)
					if !ok {
						t.Fatalf("not a CustomElementDeclaration: %#v", getFirstDecl(t, pkg))
					}
					if len(ce.Events) != 1 || ce.Events[0].Type.Text != "CustomEvent" {
						t.Errorf("unexpected event type: %#v", ce.Events)
					}
				})
				t.Run("Deprecation", func(t *testing.T) {
					t.Run("Bool", func(t *testing.T) {
					mustRunFixture(t)
						manifestJSON := loadFixture(t, "event-deprecated-bool.json")
						pkg := mustUnmarshalPackage(t, manifestJSON)
						ce := mustCustomElementDecl(t, mustFirstModule(t, pkg).Declarations[0])
						evs := ce.Events
						if len(evs) != 1 {
							t.Fatalf("len(Events) = %d, want 1", len(evs))
						}
						if got, want := evs[0].Deprecated, DeprecatedFlag(true); got != want {
							t.Errorf("Deprecated = %#v, want %#v", got, want)
						}
					})

					t.Run("Reason", func(t *testing.T) {
						mustRunFixture(t)
						manifestJSON := loadFixture(t, "event-deprecated-reason.json")
						pkg := mustUnmarshalPackage(t, manifestJSON)
						ce := mustCustomElementDecl(t, mustFirstModule(t, pkg).Declarations[0])
						evs := ce.Events
						if len(evs) != 1 {
							t.Fatalf("len(Events) = %d, want 1", len(evs))
						}
						if got, want := evs[0].Deprecated, DeprecatedReason("use something else"); got != want {
							t.Errorf("Deprecated = %#v, want %#v", got, want)
						}
					})
				})
			})
		})

		t.Run("Slots", func(t *testing.T) {
			mustRunFixture(t)
			manifestJSON := loadFixture(t, "custom_element_slots.json")
			pkg := mustUnmarshalPackage(t, manifestJSON)
			ce := mustCustomElementDecl(t, mustFirstModule(t, pkg).Declarations[0])
			if len(ce.Slots) != 1 || ce.Slots[0].Description != "Default slot" {
				t.Errorf("Slots = %+v, want 1 Default slot", ce.Slots)
			}
		})

		t.Run("SlotDeprecation", func(t *testing.T) {
			t.Run("Bool", func(t *testing.T) {
				mustRunFixture(t)
				manifestJSON := loadFixture(t, "slot-deprecated-bool.json")
				pkg := mustUnmarshalPackage(t, manifestJSON)
				ce := mustCustomElementDecl(t, mustFirstModule(t, pkg).Declarations[0])
				slots := ce.Slots
				if len(slots) != 1 {
					t.Fatalf("len(Slots) = %d, want 1", len(slots))
				}
				if got, want := slots[0].Deprecated, DeprecatedFlag(true); got != want {
					t.Errorf("Deprecated = %#v, want %#v", got, want)
				}
			})

			t.Run("Reason", func(t *testing.T) {
				mustRunFixture(t)
				manifestJSON := loadFixture(t, "slot-deprecated-reason.json")
				pkg := mustUnmarshalPackage(t, manifestJSON)
				ce := mustCustomElementDecl(t, mustFirstModule(t, pkg).Declarations[0])
				slots := ce.Slots
				if len(slots) != 1 {
					t.Fatalf("len(Slots) = %d, want 1", len(slots))
				}
				if got, want := slots[0].Deprecated, DeprecatedReason("use default slot instead"); got != want {
					t.Errorf("Deprecated = %#v, want %#v", got, want)
				}
			})
		})

		t.Run("CssPartsPropertiesStates", func(t *testing.T) {
			mustRunFixture(t)
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
		})

		t.Run("CssPartDeprecation", func(t *testing.T) {
			t.Run("Bool", func(t *testing.T) {
				mustRunFixture(t)
				manifestJSON := loadFixture(t, "part-deprecated-bool.json")
				pkg := mustUnmarshalPackage(t, manifestJSON)
				ce := mustCustomElementDecl(t, mustFirstModule(t, pkg).Declarations[0])
				parts := ce.CssParts
				if len(parts) != 1 {
					t.Fatalf("len(CssParts) = %d, want 1", len(parts))
				}
				if got, want := parts[0].Deprecated, DeprecatedFlag(true); got != want {
					t.Errorf("Deprecated = %#v, want %#v", got, want)
				}
			})

			t.Run("Reason", func(t *testing.T) {
				mustRunFixture(t)
				manifestJSON := loadFixture(t, "part-deprecated-reason.json")
				pkg := mustUnmarshalPackage(t, manifestJSON)
				ce := mustCustomElementDecl(t, mustFirstModule(t, pkg).Declarations[0])
				parts := ce.CssParts
				if len(parts) != 1 {
					t.Fatalf("len(CssParts) = %d, want 1", len(parts))
				}
				if got, want := parts[0].Deprecated, DeprecatedReason("use another part"); got != want {
					t.Errorf("Deprecated = %#v, want %#v", got, want)
				}
			})
		})

		t.Run("CssPropertyDeprecation", func(t *testing.T) {
			t.Run("Bool", func(t *testing.T) {
				mustRunFixture(t)
				manifestJSON := loadFixture(t, "css-property-deprecated-bool.json")
				pkg := mustUnmarshalPackage(t, manifestJSON)
				ce := mustCustomElementDecl(t, mustFirstModule(t, pkg).Declarations[0])
				props := ce.CssProperties
				if len(props) != 1 {
					t.Fatalf("len(CssProperties) = %d, want 1", len(props))
				}
				if got, want := props[0].Deprecated, DeprecatedFlag(true); got != want {
					t.Errorf("Deprecated = %#v, want %#v", got, want)
				}
			})
			t.Run("Reason", func(t *testing.T) {
				mustRunFixture(t)
				manifestJSON := loadFixture(t, "css-property-deprecated-reason.json")
				pkg := mustUnmarshalPackage(t, manifestJSON)
				ce := mustCustomElementDecl(t, mustFirstModule(t, pkg).Declarations[0])
				props := ce.CssProperties
				if len(props) != 1 {
					t.Fatalf("len(CssProperties) = %d, want 1", len(props))
				}
				if got, want := props[0].Deprecated, DeprecatedReason("use --other instead"); got != want {
					t.Errorf("Deprecated = %#v, want %#v", got, want)
				}
			})
		})

		t.Run("CssStateDeprecation", func(t *testing.T) {
			t.Run("Bool", func(t *testing.T) {
				mustRunFixture(t)
				manifestJSON := loadFixture(t, "css-state-deprecated-bool.json")
				pkg := mustUnmarshalPackage(t, manifestJSON)
				ce := mustCustomElementDecl(t, mustFirstModule(t, pkg).Declarations[0])
				states := ce.CssStates
				if len(states) != 1 {
					t.Fatalf("len(CssStates) = %d, want 1", len(states))
				}
				if got, want := states[0].Deprecated, DeprecatedFlag(true); got != want {
					t.Errorf("Deprecated = %#v, want %#v", got, want)
				}
			})
			t.Run("Reason", func(t *testing.T) {
				mustRunFixture(t)
				manifestJSON := loadFixture(t, "css-state-deprecated-reason.json")
				pkg := mustUnmarshalPackage(t, manifestJSON)
				ce := mustCustomElementDecl(t, mustFirstModule(t, pkg).Declarations[0])
				states := ce.CssStates
				if len(states) != 1 {
					t.Fatalf("len(CssStates) = %d, want 1", len(states))
				}
				if got, want := states[0].Deprecated, DeprecatedReason("no longer needed"); got != want {
					t.Errorf("Deprecated = %#v, want %#v", got, want)
				}
			})
		})

		t.Run("Demos", func(t *testing.T) {
			mustRunFixture(t)
			manifestJSON := loadFixture(t, "custom_element_demos.json")
			pkg := mustUnmarshalPackage(t, manifestJSON)
			ce := mustCustomElementDecl(t, mustFirstModule(t, pkg).Declarations[0])
			if len(ce.Demos) != 1 || ce.Demos[0].Description != "Basic demo" || ce.Demos[0].URL != "demo.html" {
				t.Errorf("Demos = %+v, want 1 Basic demo with url demo.html", ce.Demos)
			}
		})

		t.Run("Mixins", func(t *testing.T) {
			mustRunFixture(t)
			manifestJSON := loadFixture(t, "custom_element_mixins.json")
			pkg := mustUnmarshalPackage(t, manifestJSON)
			ce := mustCustomElementDecl(t, mustFirstModule(t, pkg).Declarations[0])
			if len(ce.Mixins) != 1 || ce.Mixins[0].Name != "MyMixin" {
				t.Errorf("Mixins = %+v, want 1 MyMixin", ce.Mixins)
			}
		})

		t.Run("Members", func(t *testing.T) {
			mustRunFixture(t)
			manifestJSON := loadFixture(t, "custom_element_members.json")
			pkg := mustUnmarshalPackage(t, manifestJSON)
			ce := mustCustomElementDecl(t, mustFirstModule(t, pkg).Declarations[0])
			if len(ce.Members) != 1 {
				t.Fatalf("Members = %+v, want 1 item", ce.Members)
			}
			member := mustClassField(t, ce.Members[0])
			if member.Name != "prop1" || member.Type == nil || member.Type.Text != "string" {
				t.Errorf("Member = %+v, want prop1: string", member)
			}
		})

		t.Run("SliceNils", func(t *testing.T) {
			mustRunFixture(t)
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
		})
	})

	t.Run("Class", func(t *testing.T) {
		mustRunFixture(t)
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
		t.Run("DeprecatedBool", func(t *testing.T) {
			mustRunFixture(t)
			data := loadFixture(t, "class-deprecated-bool.json")
			pkg := mustUnmarshalPackage(t, data)
			class, ok := getFirstDecl(t, pkg).(*ClassDeclaration)
			if !ok {
				t.Fatalf("not a ClassDeclaration: %#v", getFirstDecl(t, pkg))
			}
			if class.Deprecated != DeprecatedFlag(true) {
				t.Errorf("expected Deprecated=true, got %#v", class.Deprecated)
			}
		})
		t.Run("DeprecatedReason", func(t *testing.T) {
			mustRunFixture(t)
			data := loadFixture(t, "class-deprecated-reason.json")
			pkg := mustUnmarshalPackage(t, data)
			class, ok := getFirstDecl(t, pkg).(*ClassDeclaration)
			if !ok {
				t.Fatalf("not a ClassDeclaration: %#v", getFirstDecl(t, pkg))
			}
			exp := DeprecatedReason("Use the new class instead")
			if class.Deprecated != exp {
				t.Errorf("expected Deprecated=%#v, got %#v", exp, class.Deprecated)
			}
		})
		t.Run("FieldTypeText", func(t *testing.T) {
			mustRunFixture(t)
			data := loadFixture(t, "class-field-type-text.json")
			pkg := mustUnmarshalPackage(t, data)
			class, ok := getFirstDecl(t, pkg).(*ClassDeclaration)
			if !ok {
				t.Fatalf("not a ClassDeclaration: %#v", getFirstDecl(t, pkg))
			}
			if len(class.Members) != 1 {
				t.Fatalf("expected 1 member, got %d", len(class.Members))
			}
			field, ok := class.Members[0].(*ClassField)
			if !ok {
				t.Fatalf("first member not a ClassField: %#v", class.Members[0])
			}
			if field.Type.Text != "number" {
				t.Errorf("unexpected field type: %#v", field.Type.Text)
			}
		})
		t.Run("MethodParameters", func(t *testing.T) {
			mustRunFixture(t)
			data := loadFixture(t, "class-method-parameters.json")
			pkg := mustUnmarshalPackage(t, data)
			class, ok := getFirstDecl(t, pkg).(*ClassDeclaration)
			if !ok {
				t.Fatalf("not a ClassDeclaration: %#v", getFirstDecl(t, pkg))
			}
			if len(class.Members) != 1 {
				t.Fatalf("expected 1 member, got %d", len(class.Members))
			}
			method, ok := class.Members[0].(*ClassMethod)
			if !ok {
				t.Fatalf("first member not a ClassMethod: %#v", class.Members[0])
			}
			if len(method.Parameters) != 1 || method.Parameters[0].Name != "bar" {
				t.Errorf("unexpected parameters: %#v", method.Parameters)
			}
			if method.Parameters[0].Type.Text != "string" {
				t.Errorf("unexpected parameter type: %#v", method.Parameters[0].Type.Text)
			}
		})
	})

	t.Run("Mixin", func(t *testing.T) {
			mustRunFixture(t)
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
		t.Run("DeprecatedBool", func(t *testing.T) {
			mustRunFixture(t)
			data := loadFixture(t, "mixin-deprecated-bool.json")
			pkg := mustUnmarshalPackage(t, data)
			mixin, ok := getFirstDecl(t, pkg).(*MixinDeclaration)
			if !ok {
				t.Fatalf("not a MixinDeclaration: %#v", getFirstDecl(t, pkg))
			}
			if mixin.Deprecated != DeprecatedFlag(true) {
				t.Errorf("expected Deprecated=true, got %#v", mixin.Deprecated)
			}
		})
		t.Run("DeprecatedReason", func(t *testing.T) {
			mustRunFixture(t)
			data := loadFixture(t, "mixin-deprecated-reason.json")
			pkg := mustUnmarshalPackage(t, data)
			mixin, ok := getFirstDecl(t, pkg).(*MixinDeclaration)
			if !ok {
				t.Fatalf("not a MixinDeclaration: %#v", getFirstDecl(t, pkg))
			}
			exp := DeprecatedReason("Mixins are not recommended")
			if mixin.Deprecated != exp {
				t.Errorf("expected Deprecated=%#v, got %#v", exp, mixin.Deprecated)
			}
		})
	})

	t.Run("Function", func(t *testing.T) {
		t.Run("Fields", func(t *testing.T) {
			mustRunFixture(t)
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
		})
		t.Run("Deprecated", func(t *testing.T) {
			t.Run("Bool", func(t *testing.T) {
				mustRunFixture(t)
				manifestJSON := loadFixture(t, "function-deprecated-bool.json")
				pkg := mustUnmarshalPackage(t, manifestJSON)
				mod := mustFirstModule(t, pkg)
				fn := mustFunctionDecl(t, mustModuleDecls(t, mod, 1)[0])
				if fn.Deprecated != DeprecatedFlag(true) {
					t.Errorf("expected Deprecated=true, got %#v", fn.Deprecated)
				}
			})
			t.Run("Reason", func(t *testing.T) {
				mustRunFixture(t)
				data := loadFixture(t, "function-deprecated-reason.json")
				pkg := mustUnmarshalPackage(t, data)
				fn, ok := getFirstDecl(t, pkg).(*FunctionDeclaration)
				if !ok {
					t.Fatalf("not a FunctionDeclaration: %#v", getFirstDecl(t, pkg))
				}
				exp := DeprecatedReason("Use anotherFunction instead")
				if fn.Deprecated != exp {
					t.Errorf("expected Deprecated=%#v, got %#v", exp, fn.Deprecated)
				}
			})
		})
		t.Run("Parameters", func(t *testing.T) {
			t.Run("Fields", func(t *testing.T) {
				mustRunFixture(t)
				data := loadFixture(t, "function-parameters.json")
				pkg := mustUnmarshalPackage(t, data)
				fn, ok := getFirstDecl(t, pkg).(*FunctionDeclaration)
				if !ok {
					t.Fatalf("not a FunctionDeclaration: %#v", getFirstDecl(t, pkg))
				}
				if len(fn.Parameters) != 2 {
					t.Fatalf("expected 2 parameters, got %d", len(fn.Parameters))
				}
				if fn.Parameters[0].Name != "x" || fn.Parameters[0].Type.Text != "number" {
					t.Errorf("unexpected first parameter: %#v", fn.Parameters[0])
				}
				if fn.Parameters[1].Name != "y" || fn.Parameters[1].Type.Text != "string" {
					t.Errorf("unexpected second parameter: %#v", fn.Parameters[1])
				}
				if fn.Return.Type.Text != "boolean" {
					t.Errorf("unexpected return type: %#v", fn.Return.Type.Text)
				}
			})
			t.Run("Deprecation", func(t *testing.T) {
				t.Run("Bool", func(t *testing.T) {
					mustRunFixture(t)
					data := loadFixture(t, "function-parameter-deprecated-bool.json")
					pkg := mustUnmarshalPackage(t, data)
					fn, ok := getFirstDecl(t, pkg).(*FunctionDeclaration)
					if !ok {
						t.Fatalf("not a FunctionDeclaration: %#v", getFirstDecl(t, pkg))
					}
					if len(fn.Parameters) != 1 {
						t.Fatalf("expected 1 parameter, got %d", len(fn.Parameters))
					}
					if fn.Parameters[0].Deprecated != DeprecatedFlag(true) {
						t.Errorf("expected Deprecated=true, got %#v", fn.Parameters[0].Deprecated)
					}
				})
				t.Run("Reason", func(t *testing.T) {
					mustRunFixture(t)
					data := loadFixture(t, "function-parameter-deprecated-reason.json")
					pkg := mustUnmarshalPackage(t, data)
					fn, ok := getFirstDecl(t, pkg).(*FunctionDeclaration)
					if !ok {
						t.Fatalf("not a FunctionDeclaration: %#v", getFirstDecl(t, pkg))
					}
					if len(fn.Parameters) != 1 {
						t.Fatalf("expected 1 parameter, got %d", len(fn.Parameters))
					}
					exp := DeprecatedReason("No longer used")
					if fn.Parameters[0].Deprecated != exp {
						t.Errorf("expected Deprecated=%#v, got %#v", exp, fn.Parameters[0].Deprecated)
					}
				})
			})
		})
		t.Run("ReturnTypeText", func(t *testing.T) {
			data := loadFixture(t, "function-return-type-text.json")
			pkg := mustUnmarshalPackage(t, data)
			fn, ok := getFirstDecl(t, pkg).(*FunctionDeclaration)
			if !ok {
				t.Fatalf("not a FunctionDeclaration: %#v", getFirstDecl(t, pkg))
			}
			if fn.Return.Type.Text != "Promise<string>" {
				t.Errorf("unexpected return type: %#v", fn.Return.Type.Text)
			}
		})
	})

	t.Run("Variable", func(t *testing.T) {
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
		t.Run("DeprecatedBool", func(t *testing.T) {
			data := loadFixture(t, "variable-deprecated-bool.json")
			pkg := mustUnmarshalPackage(t, data)
			variable, ok := getFirstDecl(t, pkg).(*VariableDeclaration)
			if !ok {
				t.Fatalf("not a VariableDeclaration: %#v", getFirstDecl(t, pkg))
			}
			if variable.Deprecated != DeprecatedFlag(true) {
				t.Errorf("expected Deprecated=true, got %#v", variable.Deprecated)
			}
		})
		t.Run("DeprecatedReason", func(t *testing.T) {
			data := loadFixture(t, "variable-deprecated-reason.json")
			pkg := mustUnmarshalPackage(t, data)
			variable, ok := getFirstDecl(t, pkg).(*VariableDeclaration)
			if !ok {
				t.Fatalf("not a VariableDeclaration: %#v", getFirstDecl(t, pkg))
			}
			exp := DeprecatedReason("No longer supported")
			if variable.Deprecated != exp {
				t.Errorf("expected Deprecated=%#v, got %#v", exp, variable.Deprecated)
			}
		})
		t.Run("TypeText", func(t *testing.T) {
			data := loadFixture(t, "variable-type-text.json")
			pkg := mustUnmarshalPackage(t, data)
			variable, ok := getFirstDecl(t, pkg).(*VariableDeclaration)
			if !ok {
				t.Fatalf("not a VariableDeclaration: %#v", getFirstDecl(t, pkg))
			}
			if variable.Type.Text != "Array<string>" {
				t.Errorf("unexpected type text: %#v", variable.Type.Text)
			}
		})
	})

	t.Run("CustomElementMixin", func(t *testing.T) {
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
		if len(cem.Attributes) != 1 || cem.Attributes[0].Name != "cem-attr" {
			t.Errorf("Attributes = %+v, want 1 cem-attr", cem.Attributes)
		}
	})

	t.Run("EmptySlices", func(t *testing.T) {
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
	})

	t.Run("Export", func(t *testing.T) {
		t.Run("JavaScriptExport", func(t *testing.T) {
			data := loadFixture(t, "export-js-basic.json")
			t.Log(string(data))
			pkg := mustUnmarshalPackage(t, data)
			t.Log(pkg)
			if len(pkg.Modules) == 0 || len(pkg.Modules[0].Exports) == 0 {
				t.Fatalf("missing exports in module")
			}
			exp, ok := pkg.Modules[0].Exports[0].(*JavaScriptExport)
			if !ok || exp.Name != "jsExport" {
				t.Errorf("unexpected export: %#v", pkg.Modules[0].Exports[0])
			}
		})

		t.Run("CustomElementDefinitionExport", func(t *testing.T) {
			data := loadFixture(t, "export-custom-element-definition-basic.json")
			pkg := mustUnmarshalPackage(t, data)
			if len(pkg.Modules) == 0 || len(pkg.Modules[0].Exports) == 0 {
				t.Fatalf("missing exports in module")
			}
			exp, ok := pkg.Modules[0].Exports[0].(*CustomElementExport)
			if !ok || exp.Name != "my-element" {
				t.Errorf("unexpected export: %#v", pkg.Modules[0].Exports[0])
			}
		})
	})

	t.Run("InvalidJSON", func(t *testing.T) {
		data := []byte(`{invalid}`)
		var pkg Package
		if err := json.Unmarshal(data, &pkg); err == nil {
			t.Fatal("expected error for invalid json, got nil")
		}
	})

	t.Run("EdgeCases", func(t *testing.T) {
		t.Run("InvalidMissingRequired", func(t *testing.T) {
			data := loadFixture(t, "invalid-missing-required.json")
			t.Log(string(data))
			var pkg Package
			if err := json.Unmarshal(data, &pkg); err != nil {
				t.Error("expected error for missing required fields, got nil")
			}
		})
		t.Run("InvalidDeprecatedWrongType", func(t *testing.T) {
			data := loadFixture(t, "invalid-deprecated-wrong-type.json")
			_, err := mustUnmarshalPackageEdge(t, data)
			if err == nil {
				t.Error("expected error for invalid deprecated type, got nil")
			}
		})
		t.Run("InvalidExtraFields", func(t *testing.T) {
			data := loadFixture(t, "invalid-extra-fields.json")
			var pkg Package
			if err := json.Unmarshal(data, &pkg); err != nil {
				t.Errorf("unexpected error for extra fields: %v", err)
			}
		})
	})

	t.Run("UnknownKind", func(t *testing.T) {
		manifestJSON := loadFixture(t, "unknown_kind.json")
		_, err := mustUnmarshalPackageEdge(t, manifestJSON)
		if err == nil {
			t.Errorf("expected error for unknown kind, got nil")
		}
	})
}
