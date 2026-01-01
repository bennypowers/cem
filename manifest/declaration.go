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
type Declaration interface {
	Deprecatable
	isDeclaration()
	GetStartByte() uint
	Clone() Declaration
	Name() string
}

func unmarshalDeclaration(data json.RawMessage) (Declaration, error) {
	var kindWrap struct {
		Kind          string `json:"kind"`
		CustomElement bool   `json:"customElement"`
	}
	if err := json.Unmarshal(data, &kindWrap); err != nil {
		return nil, err
	}
	switch kindWrap.Kind {
	case "class":
		if kindWrap.CustomElement {
			var c CustomElementDeclaration
			if err := json.Unmarshal(data, &c); err != nil {
				return nil, fmt.Errorf("cannot unmarshal as CustomElementDeclaration: %w", err)
			}
			return &c, nil
		} else {
			var c ClassDeclaration
			if err := json.Unmarshal(data, &c); err != nil {
				return nil, fmt.Errorf("cannot unmarshal as ClassDeclaration: %w", err)
			}
			return &c, nil
		}
	case "custom-element-definition":
		var ce CustomElementDeclaration
		if err := json.Unmarshal(data, &ce); err != nil {
			return nil, fmt.Errorf("cannot unmarshal as CustomElementDeclaration: %w", err)
		}
		return &ce, nil
	case "mixin":
		if kindWrap.CustomElement {
			var m CustomElementMixinDeclaration
			if err := json.Unmarshal(data, &m); err != nil {
				return nil, fmt.Errorf("cannot unmarshal as CustomElementMixinDeclaration: %w", err)
			}
			return &m, nil
		} else {
			var m MixinDeclaration
			if err := json.Unmarshal(data, &m); err != nil {
				return nil, fmt.Errorf("cannot unmarshal as MixinDeclaration: %w", err)
			}
			return &m, nil
		}
	case "function":
		var f FunctionDeclaration
		if err := json.Unmarshal(data, &f); err != nil {
			return nil, fmt.Errorf("cannot unmarshal as FunctionDeclaration: %w", err)
		}
		return &f, nil
	case "variable":
		var v VariableDeclaration
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, fmt.Errorf("cannot unmarshal as VariableDeclaration: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unknown declaration kind: %s", kindWrap.Kind)
	}
}
