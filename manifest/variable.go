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

// VariableDeclaration is a variable.
type VariableDeclaration struct {
	PropertyLike
	Kind   string           `json:"kind"` // 'variable'
	Source *SourceReference `json:"source,omitempty"`
}

func (*VariableDeclaration) isDeclaration() {}

func (x *VariableDeclaration) IsDeprecated() bool {
	if x == nil {
		return false
	}
	return x.Deprecated != nil
}

func (x *VariableDeclaration) GetStartByte() uint { return x.StartByte }

func (v *VariableDeclaration) UnmarshalJSON(data []byte) error {
	var proxy map[string]json.RawMessage
	if err := json.Unmarshal(data, &proxy); err != nil {
		return err
	}
	var depRaw json.RawMessage
	if raw, ok := proxy["deprecated"]; ok {
		depRaw = raw
		delete(proxy, "deprecated")
	}
	rest, err := json.Marshal(proxy)
	if err != nil {
		return err
	}
	type Alias VariableDeclaration
	if err := json.Unmarshal(rest, (*Alias)(v)); err != nil {
		return err
	}
	if len(depRaw) > 0 && string(depRaw) != "null" {
		var dep Deprecated
		if !decodeDeprecatedField(&dep, depRaw) {
			return fmt.Errorf("invalid type for deprecated field")
		}
		v.Deprecated = dep
	}
	return nil
}
