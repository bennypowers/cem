/*
Copyright © 2025 Benny Powers <web@bennypowers.com>

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
