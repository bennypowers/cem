/*
Copyright © 2025 Benny Powers

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

// Clone creates a deep copy of the Type.
func (t *Type) Clone() *Type {
	if t == nil {
		return nil
	}

	cloned := &Type{
		Text: t.Text,
	}

	// Clone references
	if len(t.References) > 0 {
		cloned.References = make([]TypeReference, len(t.References))
		for i, ref := range t.References {
			cloned.References[i] = ref.Clone()
		}
	}

	// Clone source
	if t.Source != nil {
		source := t.Source.Clone()
		cloned.Source = &source
	}

	return cloned
}

// Clone creates a deep copy of the TypeReference.
func (tr TypeReference) Clone() TypeReference {
	return TypeReference{
		Reference: tr.Reference.Clone(),
		Start:     tr.Start,
		End:       tr.End,
	}
}
