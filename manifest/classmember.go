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

// Declaration is a union of several types.
type ClassMember interface {
	Deprecatable
	isClassMember()
	GetStartByte() uint
}

func unmarshalClassMember(data json.RawMessage) (ClassMember, error) {
	var kindWrap struct {
		Kind string `json:"kind"`
	}
	if err := json.Unmarshal(data, &kindWrap); err != nil {
		return nil, err
	}
	switch kindWrap.Kind {
	case "field":
		// Probe for custom element field properties
		type probeField struct {
			Attribute *string `json:"attribute"`
			Reflects  *bool   `json:"reflects"`
		}
		var probe probeField
		if err := json.Unmarshal(data, &probe); err != nil {
			return nil, err
		}
		if probe.Attribute != nil || probe.Reflects != nil {
			var f CustomElementField
			if err := json.Unmarshal(data, &f); err == nil {
				// NB: it's not clear to me why these assignments are necessary,
				// but empirically, they are
				f.Attribute = *probe.Attribute
				f.Reflects = *probe.Reflects
				return &f, nil
			} else {
				// Only fallback if there's a clear reason (e.g., missing required fields for custom element field)
				return nil, fmt.Errorf("cannot unmarshal as CustomElementField: %v", err)
			}
		}
		var f ClassField
		if err := json.Unmarshal(data, &f); err == nil {
			return &f, nil
		}
		return nil, fmt.Errorf("cannot unmarshal as ClassField")
	case "method":
		var m ClassMethod
		if err := json.Unmarshal(data, &m); err == nil {
			return &m, nil
		}
		return nil, fmt.Errorf("cannot unmarshal as ClassMethod")
	default:
		return nil, fmt.Errorf("unknown class member kind: %s", kindWrap.Kind)
	}
}
