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
	Name        string `json:"name"`
	Summary     string `json:"summary,omitempty"`
	Description string `json:"description,omitempty"`
}

// Clone creates a deep copy of the FullyQualified structure.
func (f FullyQualified) Clone() FullyQualified {
	return FullyQualified{
		Name:        f.Name,
		Summary:     f.Summary,
		Description: f.Description,
	}
}

type Describable interface {
	Summary() string
	Description() string
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

// Clone creates a deep copy of the Reference.
func (r Reference) Clone() Reference {
	return Reference{
		Name:    r.Name,
		Package: r.Package,
		Module:  r.Module,
	}
}

// SourceReference is a reference to the source of a declaration or member.
type SourceReference struct {
	Href string `json:"href"`
}

// Clone creates a deep copy of the SourceReference.
func (s SourceReference) Clone() SourceReference {
	return SourceReference{
		Href: s.Href,
	}
}

// PropertyLike is the common interface of variables, class fields, and function parameters.
type PropertyLike struct {
	FullyQualified
	StartByte  uint       `json:"-"`
	Type       *Type      `json:"type,omitempty"`
	Default    string     `json:"default,omitempty"`
	Deprecated Deprecated `json:"deprecated,omitempty"`
	Readonly   bool       `json:"readonly,omitempty"`
}

func (x *PropertyLike) IsDeprecated() bool {
	if x == nil {
		return false
	}
	return x.Deprecated != nil
}

// Clone creates a deep copy of the PropertyLike structure.
func (p PropertyLike) Clone() PropertyLike {
	cloned := PropertyLike{
		FullyQualified: p.FullyQualified.Clone(),
		StartByte:      p.StartByte,
		Default:        p.Default,
		Readonly:       p.Readonly,
	}

	if p.Deprecated != nil {
		cloned.Deprecated = p.Deprecated.Clone()
	}

	if p.Type != nil {
		cloned.Type = p.Type.Clone()
	}

	return cloned
}

var _ Deprecatable = (*PropertyLike)(nil)

type Deprecated interface {
	isDeprecated()
	Value() any
	Clone() Deprecated
}

type DeprecatedFlag bool

func (DeprecatedFlag) isDeprecated()       {}
func (d DeprecatedFlag) Value() any        { return bool(d) }
func (d DeprecatedFlag) Clone() Deprecated { return d }
func (d DeprecatedFlag) MarshalJSON() ([]byte, error) {
	return json.Marshal(bool(d))
}

type DeprecatedReason string

func (DeprecatedReason) isDeprecated()       {}
func (d DeprecatedReason) Value() any        { return string(d) }
func (d DeprecatedReason) Clone() Deprecated { return d }
func (d DeprecatedReason) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(d))
}

func NewDeprecated(x any) Deprecated {
	switch v := x.(type) {
	case bool:
		return DeprecatedFlag(v)
	case string:
		return DeprecatedReason(v)
	default:
		return nil
	}
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

// Clone helpers for efficient deep copying without JSON serialization overhead.
// These utility functions are used by Clone methods throughout the manifest package.

// cloneStringPtr creates a deep copy of a string pointer.
// Returns nil if the input is nil, otherwise returns a new pointer to a copy of the string.
func cloneStringPtr(s *string) *string {
	if s == nil {
		return nil
	}
	cloned := *s
	return &cloned
}

// cloneStringSlice creates a deep copy of a string slice.
// Returns nil if the input is nil, otherwise returns a new slice with copied strings.
func cloneStringSlice(s []string) []string {
	if s == nil {
		return nil
	}
	cloned := make([]string, len(s))
	copy(cloned, s)
	return cloned
}
