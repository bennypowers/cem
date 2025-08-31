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
package manifest

// Interface implementation checks
var _ Deprecatable = (*PropertyLike)(nil)

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
