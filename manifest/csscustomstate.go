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

var _ Deprecatable = (*CssCustomState)(nil)

// CssCustomState describes a CSS custom state.
type CssCustomState struct {
	FullyQualified
	Deprecated Deprecated `json:"deprecated,omitempty"` // bool or string
}

func (x *CssCustomState) IsDeprecated() bool {
	if x == nil {
		return false
	}
	return x.Deprecated != nil
}

func (c *CssCustomState) UnmarshalJSON(data []byte) error {
	type Alias CssCustomState
	aux := &struct {
		Deprecated json.RawMessage `json:"deprecated"`
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if len(aux.Deprecated) > 0 && string(aux.Deprecated) != "null" {
		var dep Deprecated
		if !decodeDeprecatedField(&dep, aux.Deprecated) {
			return fmt.Errorf("invalid type for deprecated field")
		}
		c.Deprecated = dep
	}
	return nil
}
