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

var _ Deprecatable = (*Parameter)(nil)

// Parameter is a function parameter.
type Parameter struct {
	PropertyLike
	Optional bool `json:"optional,omitempty"`
	Rest     bool `json:"rest,omitempty"`
}

func (p *Parameter) UnmarshalJSON(data []byte) error {
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
	type Alias Parameter
	if err := json.Unmarshal(rest, (*Alias)(p)); err != nil {
		return err
	}
	if len(depRaw) > 0 && string(depRaw) != "null" {
		var dep Deprecated
		if !decodeDeprecatedField(&dep, depRaw) {
			return fmt.Errorf("invalid type for deprecated field")
		}
		p.Deprecated = dep
	}
	return nil
}

