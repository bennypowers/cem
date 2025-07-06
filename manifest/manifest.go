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

// Embedded type for Types with Name, Summary, and Description
type FullyQualified struct {
	Name        string     `json:"name"`
	Summary     string     `json:"summary,omitempty"`
	Description string     `json:"description,omitempty"`
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

// PropertyLike is the common interface of variables, class fields, and function parameters.
type PropertyLike struct {
	FullyQualified
	StartByte   uint       `json:"-"`
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

