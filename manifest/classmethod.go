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

import (
	"encoding/json"
	"fmt"
)

var _ Deprecatable = (*ClassMethod)(nil)

// ClassMethod is a method.
type ClassMethod struct {
	FunctionLike
	FullyQualified
	Kind          string           `json:"kind"` // 'method'
	Static        bool             `json:"static,omitempty"`
	Privacy       Privacy          `json:"privacy,omitempty"` // 'public', 'private', 'protected'
	InheritedFrom *Reference       `json:"inheritedFrom,omitempty"`
	Source        *SourceReference `json:"source,omitempty"`
	Deprecated    Deprecated       `json:"deprecated,omitempty"` // bool or string
}

func (*ClassMethod) isClassMember() {}

func (x *ClassMethod) GetStartByte() uint { return x.StartByte }

func (x *ClassMethod) IsDeprecated() bool { return x.Deprecated != nil }

func (m *ClassMethod) UnmarshalJSON(data []byte) error {
	type Alias ClassMethod
	aux := &struct {
		Deprecated json.RawMessage   `json:"deprecated"`
		Parameters []json.RawMessage `json:"parameters"`
		*Alias
	}{
		Alias: (*Alias)(m),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	// Handle deprecated
	if len(aux.Deprecated) > 0 && string(aux.Deprecated) != "null" {
		var dep Deprecated
		if !decodeDeprecatedField(&dep, aux.Deprecated) {
			return fmt.Errorf("invalid type for deprecated field")
		}
		m.Deprecated = dep
	}
	// Handle parameters
	m.Parameters = nil
	for _, p := range aux.Parameters {
		var param Parameter
		if err := json.Unmarshal(p, &param); err != nil {
			return fmt.Errorf("unmarshal parameter: %w", err)
		}
		m.Parameters = append(m.Parameters, param)
	}
	return nil
}
