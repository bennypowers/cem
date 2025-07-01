// Package cem represents the schema for Custom Elements Manifest.
// Generated from schema.d.ts of https://github.com/webcomponents/custom-elements-manifest
package manifest

import (
	"encoding/json"
	"errors"
	"regexp"
)

func normalizePath(path string) string {
	return regexp.MustCompile(`\.ts$`).ReplaceAllString(path, ".js")
}

// Package is the top-level interface of a custom elements manifest file.
type Package struct {
	SchemaVersion string    `json:"schemaVersion"`
	Readme        *string   `json:"readme,omitempty"`
	Modules       []Module  `json:"modules"`
	Deprecated    any       `json:"deprecated,omitempty"` // bool or string
}

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
func (x *Package) GetAllTagNamesWithContext() (tags []CustomElementWithContext) {
	for _, m := range x.Modules {
		mrs := make(map[string]CustomElementWithContext)
		for _, d := range m.Declarations {
			if ced, ok := d.(*CustomElementDeclaration); ok {
				mrs[ced.TagName] = CustomElementWithContext{CustomElementDeclaration: *ced, JavaScriptModule: m, TagName: ced.TagName}
			}
		}
		for _, e := range m.Exports {
			if cee, ok := e.(*CustomElementExport); ok {
				r := mrs[cee.Name]
				r.CustomElementExport = *cee
			}
		}
		for _, r := range mrs {
			tags = append(tags, r)
		}
	}
	return tags
}
func (x *Package) GetTagAttrsWithContext(tagName string) (attrs []AttributeWithContext, err error) {
	attrMap := make(map[string]AttributeWithContext)
	var ceName string
  var ceExport CustomElementExport
	for _, m := range x.Modules {
		for _, d := range m.Declarations {
			if ced, ok := d.(*CustomElementDeclaration); ok {
				if ced.TagName == tagName {
					ceName = ced.TagName
					fieldMap := make(map[string]CustomElementField)
					for _, member := range ced.Members {
						if cef, ok := member.(*CustomElementField); ok {
							fieldMap[cef.Attribute] = *cef
						}
					}
					for _, attr := range ced.Attributes {
						attrMap[attr.Name] = AttributeWithContext{
							Name: attr.Name,
							Attribute: attr,
							JavaScriptModule: m,
							CustomElementDeclaration: *ced,
							CustomElementField: fieldMap[attr.Name],
						}
					}
				}
				break
			}
		}
		for _, e := range m.Exports {
			if cee, ok := e.(*CustomElementExport); ok {
				if cee.Name == ceName {
					ceExport = *cee
					break
				}
			}
		}
	}
	length := len(attrMap)
	if length == 0 {
		return nil, errors.New("Tag not found: " + tagName)
	}
	attrs = make([]AttributeWithContext, 0)
	for name, awc := range attrMap{
		if name != "" {
			awc.CustomElementExport = ceExport
			attrs = append(attrs, awc)
		}
	}
	return attrs, nil
}

func NewPackage(modules []Module) Package {
	return Package{
		SchemaVersion: "1.0.0",
		Modules: modules,
	}
}

// Module may expand in future; currently only JavaScriptModule.
type Module = JavaScriptModule

func NewModule(file string) (*Module) {
	return &Module{
		Kind: "javascript-module",
		Path: normalizePath(file),
	}
}

type JavaScriptModule struct {
	Kind         string        `json:"kind"` // 'javascript-module'
	Path         string        `json:"path"`
	Summary      string        `json:"summary,omitempty"`
	Description  string        `json:"description,omitempty"`
	Declarations []Declaration `json:"declarations,omitempty"`
	Exports      []Export      `json:"exports,omitempty"`
	Deprecated   any           `json:"deprecated,omitempty"` // bool or string
}

// Export is a union type: JavaScriptExport or CustomElementExport.
type Export interface {
	isExport()
	GetStartByte() uint
}

// JavaScriptExport represents a JS export.
type JavaScriptExport struct {
	StartByte   uint			 `json:"-"`
	Kind        string     `json:"kind"` // 'js'
	Name        string     `json:"name"`
	Declaration *Reference `json:"declaration"`
	Deprecated  any        `json:"deprecated,omitempty"` // bool or string
}
func (*JavaScriptExport) isExport() {}
func (e *JavaScriptExport) GetStartByte() uint {
	return e.StartByte
}

// CustomElementExport represents a custom element definition.
type CustomElementExport struct {
	StartByte   uint			 `json:"-"`
	Kind        string     `json:"kind"` // 'custom-element-definition'
	Name        string     `json:"name"`
	Declaration *Reference `json:"declaration"`
	Deprecated  any        `json:"deprecated,omitempty"` // bool or string
}
func (*CustomElementExport) isExport() {}
func (e *CustomElementExport) GetStartByte() uint {
	return e.StartByte
}
func NewCustomElementExport(
	tagName string,
	declaration *Reference,
	startByte uint,
	deprecated *Deprecated,
) *CustomElementExport {
	ce := &CustomElementExport{
		Kind: "custom-element-definition",
		StartByte: startByte,
		Name: tagName,
		Declaration: declaration,
	}
	if deprecated != nil {
		ce.Deprecated = deprecated
	}
	return ce
}

// Declaration is a union of several types.
type Declaration interface {
	isDeclaration()
	GetStartByte() uint
}

// Reference to an export of a module.
type Reference struct {
	Name    string `json:"name"`
	Package string `json:"package,omitempty"`
	Module  string `json:"module,omitempty"`
}

func NewReference(name string, pkg string, module string) *Reference {
	return &Reference{
		Name: name,
		Package: pkg,
		Module: normalizePath(module),
	}
}

// SourceReference is a reference to the source of a declaration or member.
type SourceReference struct {
	Href string `json:"href"`
}

// CustomElementDeclaration extends ClassDeclaration and CustomElement.
type CustomElementDeclaration struct {
	ClassDeclaration
	CustomElement
}
func (*CustomElementDeclaration) isDeclaration() {}
func (x *CustomElementDeclaration) GetStartByte() uint { return x.StartByte }

type CustomElementWithContext struct {
	TagName string
	CustomElementDeclaration
	JavaScriptModule
	CustomElementExport
}

// CustomElement adds fields to classes/mixins for custom elements.
type CustomElement struct {
	TagName       string              `json:"tagName,omitempty"`
	Attributes    []Attribute         `json:"attributes,omitempty"`
	Events        []Event             `json:"events,omitempty"`
	Slots         []Slot              `json:"slots,omitempty"`
	CssParts      []CssPart           `json:"cssParts,omitempty"`
	CssProperties []CssCustomProperty `json:"cssProperties,omitempty"`
	CssStates     []CssCustomState    `json:"cssStates,omitempty"`
	Demos         []Demo              `json:"demos,omitempty"`
	CustomElement bool                `json:"customElement"`
}

// Attribute for custom elements.
type Attribute struct {
	Name          string     `json:"name"`
	Summary       string     `json:"summary,omitempty"`
	Description   string     `json:"description,omitempty"`
	InheritedFrom *Reference `json:"inheritedFrom,omitempty"`
	Type          *Type      `json:"type,omitempty"`
	Default       string     `json:"default,omitempty"`
	FieldName     string     `json:"fieldName,omitempty"`
	Deprecated    Deprecated `json:"deprecated,omitempty"` // bool or string
	StartByte     uint       `json:"-"`
}

type AttributeWithContext struct {
	Name string
	Attribute
	CustomElementDeclaration
	JavaScriptModule
	CustomElementExport
  CustomElementField
}

// Event emitted by a custom element.
type Event struct {
	Name          string     `json:"name"`
	Summary       string     `json:"summary,omitempty"`
	Description   string     `json:"description,omitempty"`
	Type          *Type      `json:"type,omitempty"`
	InheritedFrom *Reference `json:"inheritedFrom,omitempty"`
	Deprecated    Deprecated `json:"deprecated,omitempty"` // bool or string
}

// Slot in a custom element.
type Slot struct {
	Name        string     `json:"name"`
	Summary     string     `json:"summary,omitempty"`
	Description string     `json:"description,omitempty"`
	Deprecated  Deprecated `json:"deprecated,omitempty"` // bool or string
}

// CssPart describes a CSS part.
type CssPart struct {
	Name        string     `json:"name"`
	Summary     string     `json:"summary,omitempty"`
	Description string     `json:"description,omitempty"`
	Deprecated  Deprecated `json:"deprecated,omitempty"` // bool or string
}

// CssCustomState describes a CSS custom state.
type CssCustomState struct {
	Name        string     `json:"name"`
	Summary     string     `json:"summary,omitempty"`
	Description string     `json:"description,omitempty"`
	Deprecated  Deprecated `json:"deprecated,omitempty"` // bool or string
}

// CssCustomProperty describes a CSS custom property.
type CssCustomProperty struct {
	StartByte   uint			 `json:"-"`
	Name        string     `json:"name"`
	Syntax      string     `json:"syntax,omitempty"`
	Default     string     `json:"default,omitempty"`
	Summary     string     `json:"summary,omitempty"`
	Description string     `json:"description,omitempty"`
	Deprecated  Deprecated `json:"deprecated,omitempty"` // bool or string
}

// Type string representation.
type Type struct {
	Text       string           `json:"text,omitempty"`
	References []TypeReference  `json:"references,omitempty"`
	Source     *SourceReference `json:"source,omitempty"`
}

// TypeReference associates with a type string and optionally a range.
type TypeReference struct {
	Reference
	Start int `json:"start,omitempty"`
	End   int `json:"end,omitempty"`
}

// ClassLike is the common interface of classes and mixins.
type ClassLike struct {
	StartByte		uint			`json:"-"`
	Name        string           `json:"name"`
	Summary     string           `json:"summary,omitempty"`
	Description string           `json:"description,omitempty"`
	Superclass  *Reference       `json:"superclass,omitempty"`
	Mixins      []Reference      `json:"mixins,omitempty"`
	Members     []ClassMember    `json:"members,omitempty"`
	Source      *SourceReference `json:"source,omitempty"`
	Deprecated  Deprecated        `json:"deprecated,omitempty"` // bool or string
}

// ClassDeclaration is a class.
type ClassDeclaration struct {
	ClassLike
	Kind string `json:"kind"` // 'class'
}
func (*ClassDeclaration) isDeclaration() {}
func (x *ClassDeclaration) GetStartByte() uint { return x.StartByte }

// Declaration is a union of several types.
type ClassMember interface {
	isClassMember()
	GetStartByte() uint
}

// ClassField is a class field.
type ClassField struct {
	PropertyLike
	Kind          string       			`json:"kind"` // 'field'
	Static        bool        			`json:"static,omitempty"`
	Privacy       Privacy      			`json:"privacy,omitempty"` // 'public', 'private', 'protected'
	InheritedFrom *Reference 				`json:"inheritedFrom,omitempty"`
	Source        *SourceReference	`json:"source,omitempty"`
}
func (ClassField) isClassMember() {}
func (f ClassField) GetStartByte() uint { return f.StartByte }

// ClassMethod is a method.
type ClassMethod struct {
	FunctionLike
	Kind          string           `json:"kind"` // 'method'
	Static        bool             `json:"static,omitempty"`
	Privacy       Privacy          `json:"privacy,omitempty"` // 'public', 'private', 'protected'
	InheritedFrom *Reference       `json:"inheritedFrom,omitempty"`
	Source        *SourceReference `json:"source,omitempty"`
}
func (ClassMethod) isClassMember() {}
func (x ClassMethod) GetStartByte() uint { return x.StartByte }

// PropertyLike is the common interface of variables, class fields, and function parameters.
type PropertyLike struct {
	StartByte   uint       `json:"-"`
	Name        string     `json:"name"`
	Summary     string     `json:"summary,omitempty"`
	Description string     `json:"description,omitempty"`
	Type        *Type      `json:"type,omitempty"`
	Default     string     `json:"default,omitempty"`
	Deprecated  Deprecated `json:"deprecated,omitempty"`
	Readonly    bool       `json:"readonly,omitempty"`
}

// CustomElementField extends ClassField with attribute/reflects.
type CustomElementField struct {
	ClassField
	Attribute string `json:"attribute,omitempty"`
	Reflects  bool   `json:"reflects,omitempty"`
}
func (x CustomElementField) isClassMember() {}
func (x CustomElementField) GetStartByte() uint { return x.ClassField.StartByte }

// MixinDeclaration describes a class mixin.
type MixinDeclaration struct {
	ClassLike
	FunctionLike
	Name        string     `json:"name"`
	Summary     string     `json:"summary,omitempty"`
	Description string     `json:"description,omitempty"`
	Deprecated  Deprecated `json:"deprecated,omitempty"`
	Kind string            `json:"kind"` // 'mixin'
}
func (*MixinDeclaration) isDeclaration() {}
func (x *MixinDeclaration) GetStartByte() uint { return x.FunctionLike.StartByte }

// CustomElementMixinDeclaration extends MixinDeclaration and CustomElement.
type CustomElementMixinDeclaration struct {
	MixinDeclaration
	CustomElementDeclaration
	Name        string     `json:"name"`
	Summary     string     `json:"summary,omitempty"`
	Description string     `json:"description,omitempty"`
	Deprecated  Deprecated `json:"deprecated,omitempty"`
}
func (*CustomElementMixinDeclaration) isDeclaration() {}
func (x *CustomElementMixinDeclaration) GetStartByte() uint { return x.FunctionLike.StartByte }

// VariableDeclaration is a variable.
type VariableDeclaration struct {
	PropertyLike
	Kind   string           `json:"kind"` // 'variable'
	Source *SourceReference `json:"source,omitempty"`
}
func (*VariableDeclaration) isDeclaration() {}
func (x *VariableDeclaration) GetStartByte() uint { return x.StartByte }

// FunctionDeclaration is a function.
type FunctionDeclaration struct {
	FunctionLike
	Kind   string           `json:"kind"` // 'function'
	Source *SourceReference `json:"source,omitempty"`
}
func (*FunctionDeclaration) isDeclaration() {}
func (x *FunctionDeclaration) GetStartByte() uint { return x.StartByte }

// Parameter is a function parameter.
type Parameter struct {
	PropertyLike
	Optional bool `json:"optional,omitempty"`
	Rest     bool `json:"rest,omitempty"`
}

type Deprecated interface {
	isDeprecated()
	Value() any
}

type DeprecatedFlag bool
func (DeprecatedFlag) isDeprecated() {}
func (d DeprecatedFlag) Value() any { return bool(d) }
func (d DeprecatedFlag) MarshalJSON() ([]byte, error) {
	return json.Marshal(bool(d))
}

type DeprecatedReason string
func (DeprecatedReason) isDeprecated() {}
func (d DeprecatedReason) Value() any { return string(d) }
func (d DeprecatedReason) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(d))
}

type Privacy string

const (
	Public    Privacy = "public"
	Protected Privacy = "protected"
	Private   Privacy = "private"
)

// FunctionLike is the common interface of functions and mixins.
type FunctionLike struct {
	StartByte   uint        `json:"-"`
	Name        string      `json:"name"`
	Summary     string      `json:"summary,omitempty"`
	Description string      `json:"description,omitempty"`
	Deprecated  Deprecated  `json:"deprecated,omitempty"` // bool or string
	Parameters  []Parameter `json:"parameters,omitempty"`
	Return      *Return     `json:"return,omitempty"`
}

// Return value for functions.
type Return struct {
	Type        *Type   `json:"type,omitempty"`
	Summary     string  `json:"summary,omitempty"`
	Description string  `json:"description,omitempty"`
}

// Demo for custom elements.
type Demo struct {
	Description string         `json:"description,omitempty"`
	URL         string          `json:"url"`
	Source      *SourceReference `json:"source,omitempty"`
}

func SerializeToBytes(pkg *Package) ([]byte, error) {
	return json.MarshalIndent(pkg, "", "  ")
}

func SerializeToString(pkg *Package) (string, error) {
	b, err := SerializeToBytes(pkg)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
