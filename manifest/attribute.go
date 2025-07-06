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

var _ Deprecatable = (*Attribute)(nil)

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

func (x *Attribute) IsDeprecated() bool {
	if x == nil {
		return false
	}
	return x.Deprecated != nil
}

func (a *Attribute) UnmarshalJSON(data []byte) error {
	type Alias Attribute
	aux := &struct {
		Deprecated json.RawMessage `json:"deprecated"`
		*Alias
	}{
		Alias: (*Alias)(a),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if len(aux.Deprecated) > 0 && string(aux.Deprecated) != "null" {
		var dep Deprecated
		if !decodeDeprecatedField(&dep, aux.Deprecated) {
			return fmt.Errorf("invalid type for deprecated field")
		}
		a.Deprecated = dep
	}
	return nil
}

