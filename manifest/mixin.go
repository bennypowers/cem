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

var _ Deprecatable = (*MixinDeclaration)(nil)
var _ Deprecatable = (*CustomElementMixinDeclaration)(nil)

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

// CustomElementMixinDeclaration extends MixinDeclaration and CustomElement.
type CustomElementMixinDeclaration struct {
	MixinDeclaration
	CustomElementDeclaration
	Name        string     `json:"name"`
	Summary     string     `json:"summary,omitempty"`
	Description string     `json:"description,omitempty"`
	Deprecated  Deprecated `json:"deprecated,omitempty"`
}

func (*MixinDeclaration) isDeclaration()       {}

func (x *MixinDeclaration) IsDeprecated() bool {
	if x == nil {
		return false
	}
	return x.Deprecated != nil
}

func (x *MixinDeclaration) GetStartByte() uint { return x.FunctionLike.StartByte }

func (m *MixinDeclaration) UnmarshalJSON(data []byte) error {
	type Alias MixinDeclaration
	aux := &struct {
		Deprecated json.RawMessage   `json:"deprecated"`
		Members    []json.RawMessage `json:"members"`
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
	// Handle members
	m.Members = nil
	for _, mem := range aux.Members {
		member, err := unmarshalClassMember(mem)
		if err == nil && member != nil {
			m.Members = append(m.Members, member)
		}
	}
	if m.Members == nil {
		m.Members = []ClassMember{}
	}
	return nil
}

func (*CustomElementMixinDeclaration) isDeclaration()       {}

func (x *CustomElementMixinDeclaration) IsDeprecated() bool {
	if x == nil {
		return false
	}
	return x.Deprecated != nil
}

func (x *CustomElementMixinDeclaration) GetStartByte() uint { return x.FunctionLike.StartByte }

func (m *CustomElementMixinDeclaration) UnmarshalJSON(data []byte) error {
	type Alias CustomElementMixinDeclaration
	aux := &struct {
		Members []json.RawMessage `json:"members"`
		*Alias
	}{Alias: (*Alias)(m)}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	m.Members = nil
	for _, mem := range aux.Members {
		member, err := unmarshalClassMember(mem)
		if err != nil {
			return fmt.Errorf("cannot unmarshal member: %w", err)
		}
		m.Members = append(m.Members, member)
	}
	return nil
}
