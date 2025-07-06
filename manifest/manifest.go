// Package cem represents the schema for Custom Elements Manifest.
// Generated from schema.d.ts of https://github.com/webcomponents/custom-elements-manifest
package manifest

import (
	"encoding/json"
	"regexp"
)

type Deprecatable interface {
	IsDeprecated() bool
}

func normalizePath(path string) string {
	return regexp.MustCompile(`\.ts$`).ReplaceAllString(path, ".js")
}

// Module may expand in future; currently only JavaScriptModule.
type Module = JavaScriptModule

func NewModule(file string) *Module {
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
	Deprecated   Deprecated    `json:"deprecated,omitempty"` // bool or string
}

// Export is a union type: JavaScriptExport or CustomElementExport.
type Export interface {
	isExport()
	GetStartByte() uint
}

// JavaScriptExport represents a JS export.
type JavaScriptExport struct {
	StartByte   uint       `json:"-"`
	Kind        string     `json:"kind"` // 'js'
	Name        string     `json:"name"`
	Declaration *Reference `json:"declaration"`
	Deprecated  Deprecated `json:"deprecated,omitempty"` // bool or string
}

func (*JavaScriptExport) isExport() {}
func (x *JavaScriptExport) IsDeprecated() bool {
	if x == nil {
		return false
	}
	return x.Deprecated != nil
}
func (e *JavaScriptExport) GetStartByte() uint {
	return e.StartByte
}
var _ Deprecatable = (*JavaScriptExport)(nil)

// CustomElementExport represents a custom element definition.
type CustomElementExport struct {
	StartByte   uint       `json:"-"`
	Kind        string     `json:"kind"` // 'custom-element-definition'
	Name        string     `json:"name"`
	Declaration *Reference `json:"declaration"`
	Deprecated  any        `json:"deprecated,omitempty"` // bool or string
}

func (*CustomElementExport) isExport() {}
func (x *CustomElementExport) IsDeprecated() bool {
	if x == nil {
		return false
	}
	return x.Deprecated != nil
}
func (e *CustomElementExport) GetStartByte() uint {
	return e.StartByte
}
var _ Deprecatable = (*CustomElementExport)(nil)

func NewCustomElementExport(
	tagName string,
	declaration *Reference,
	startByte uint,
	deprecated *Deprecated,
) *CustomElementExport {
	ce := &CustomElementExport{
		Kind:        "custom-element-definition",
		StartByte:   startByte,
		Name:        tagName,
		Declaration: declaration,
	}
	if deprecated != nil {
		ce.Deprecated = deprecated
	}
	return ce
}

// Declaration is a union of several types.
type Declaration interface {
	Deprecatable
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
		Name:    name,
		Package: pkg,
		Module:  normalizePath(module),
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

func (*CustomElementDeclaration) isDeclaration()       {}
func (x *CustomElementDeclaration) IsDeprecated() bool {
	if x == nil {
		return false
	}
	return x.Deprecated != nil
}
func (x *CustomElementDeclaration) GetStartByte() uint { return x.StartByte }

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
	StartByte   uint             `json:"-"`
	Name        string           `json:"name"`
	Summary     string           `json:"summary,omitempty"`
	Description string           `json:"description,omitempty"`
	Superclass  *Reference       `json:"superclass,omitempty"`
	Mixins      []Reference      `json:"mixins,omitempty"`
	Members     []ClassMember    `json:"members,omitempty"`
	Source      *SourceReference `json:"source,omitempty"`
	Deprecated  Deprecated       `json:"deprecated,omitempty"` // bool or string
}
func (x *ClassLike) IsDeprecated() bool {
	if x == nil {
		return false
	}
	return x.Deprecated != nil
}
var _ Deprecatable = (*ClassLike)(nil)

// ClassDeclaration is a class.
type ClassDeclaration struct {
	ClassLike
	Kind string `json:"kind"` // 'class'
}

func (*ClassDeclaration) isDeclaration()       {}
func (x *ClassDeclaration) IsDeprecated() bool {
	if x == nil {
		return false
	}
	return x.Deprecated != nil
}
func (x *ClassDeclaration) GetStartByte() uint { return x.StartByte }

// Declaration is a union of several types.
type ClassMember interface {
	Deprecatable
	isClassMember()
	GetStartByte() uint
}

// ClassMethod is a method.
type ClassMethod struct {
	FunctionLike
	Kind          string           `json:"kind"` // 'method'
	Static        bool             `json:"static,omitempty"`
	Privacy       Privacy          `json:"privacy,omitempty"` // 'public', 'private', 'protected'
	InheritedFrom *Reference       `json:"inheritedFrom,omitempty"`
	Source        *SourceReference `json:"source,omitempty"`
}

func (*ClassMethod) isClassMember()       {}
func (x *ClassMethod) GetStartByte() uint { return x.StartByte }
func (x *ClassMethod) IsDeprecated() bool { return x.Deprecated != nil }
var _ Deprecatable = (*ClassMethod)(nil)

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
func (x *PropertyLike) IsDeprecated() bool {
	if x == nil {
		return false
	}
	return x.Deprecated != nil
}
var _ Deprecatable = (*PropertyLike)(nil)

// MixinDeclaration describes a class mixin.
type MixinDeclaration struct {
	ClassLike
	FunctionLike
	Name        string     `json:"name"`
	Summary     string     `json:"summary,omitempty"`
	Description string     `json:"description,omitempty"`
	Deprecated  Deprecated `json:"deprecated,omitempty"`
	Kind        string     `json:"kind"` // 'mixin'
}

func (*MixinDeclaration) isDeclaration()       {}
func (x *MixinDeclaration) IsDeprecated() bool {
	if x == nil {
		return false
	}
	return x.Deprecated != nil
}
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

func (*CustomElementMixinDeclaration) isDeclaration()       {}
func (x *CustomElementMixinDeclaration) IsDeprecated() bool {
	if x == nil {
		return false
	}
	return x.Deprecated != nil
}
func (x *CustomElementMixinDeclaration) GetStartByte() uint { return x.FunctionLike.StartByte }

// FunctionDeclaration is a function.
type FunctionDeclaration struct {
	FunctionLike
	Kind   string           `json:"kind"` // 'function'
	Source *SourceReference `json:"source,omitempty"`
}

func (*FunctionDeclaration) isDeclaration()       {}
func (x *FunctionDeclaration) IsDeprecated() bool {
	if x == nil {
		return false
	}
	return x.Deprecated != nil
}
func (x *FunctionDeclaration) GetStartByte() uint { return x.StartByte }

type Deprecated interface {
	isDeprecated()
	Value() any
}

type DeprecatedFlag bool

func (DeprecatedFlag) isDeprecated() {}
func (d DeprecatedFlag) Value() any  { return bool(d) }
func (d DeprecatedFlag) MarshalJSON() ([]byte, error) {
	return json.Marshal(bool(d))
}

type DeprecatedReason string

func (DeprecatedReason) isDeprecated() {}
func (d DeprecatedReason) Value() any  { return string(d) }
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
	Type        *Type  `json:"type,omitempty"`
	Summary     string `json:"summary,omitempty"`
	Description string `json:"description,omitempty"`
}

// Demo for custom elements.
type Demo struct {
	Description string           `json:"description,omitempty"`
	URL         string           `json:"url"`
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

