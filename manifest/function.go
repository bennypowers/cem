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

var _ Deprecatable = (*FunctionDeclaration)(nil)

// Return value for functions.
type Return struct {
	Type        *Type  `json:"type,omitempty"`
	Summary     string `json:"summary,omitempty"`
	Description string `json:"description,omitempty"`
}

// FunctionLike is the common interface of functions and mixins.
type FunctionLike struct {
	StartByte   uint        `json:"-"`
	Parameters  []Parameter `json:"parameters,omitempty"`
	Return      *Return     `json:"return,omitempty"`
}

// CssPart describes a CSS part.
// FunctionDeclaration is a function.
type FunctionDeclaration struct {
	FunctionLike
	FullyQualified
	Deprecated  Deprecated  `json:"deprecated,omitempty"` // bool or string
	Kind   string           `json:"kind"` // 'function'
	Source *SourceReference `json:"source,omitempty"`
}

func (*FunctionDeclaration) isDeclaration()       {}

func (x *FunctionDeclaration) IsDeprecated() bool {
	if x == nil {
		return false
	}
	return x.Deprecated != nil
}

func (x *FunctionDeclaration) GetStartByte() uint { return x.StartByte }

func (f *FunctionDeclaration) UnmarshalJSON(data []byte) error {
	type Alias FunctionDeclaration
	aux := &struct {
		Parameters []json.RawMessage `json:"parameters"`
		Deprecated json.RawMessage   `json:"deprecated"`
		*Alias
	}{
		Alias: (*Alias)(f),
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
		f.Deprecated = dep
	}
	// Unmarshal parameters using custom logic (since Parameter has its own UnmarshalJSON)
	f.Parameters = nil
	for _, p := range aux.Parameters {
		var param Parameter
		if err := json.Unmarshal(p, &param); err != nil {
			return fmt.Errorf("unmarshal parameter: %w", err)
		}
		f.Parameters = append(f.Parameters, param)
	}
	return nil
}
