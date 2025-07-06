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

// Module may expand in future; currently only JavaScriptModule.
type Module = JavaScriptModule

func NewModule(file string) *Module {
	return &Module{
		Kind: "javascript-module",
		Path: normalizePath(file),
	}
}

type JavaScriptModule struct {
	Kind         string        `json:"kind"` // 'javascript-module'
	Path         string        `json:"path"`
	Summary      string        `json:"summary,omitempty"`
	Description  string        `json:"description,omitempty"`
	Declarations []Declaration `json:"declarations,omitempty"`
	Exports      []Export      `json:"exports,omitempty"`
	Deprecated   Deprecated    `json:"deprecated,omitempty"` // bool or string
}

func (m *Module) UnmarshalJSON(data []byte) error {
	type Alias Module
	aux := &struct {
		Declarations []json.RawMessage `json:"declarations"`
		Exports      []json.RawMessage `json:"exports"`
		*Alias
	}{
		Alias: (*Alias)(m),
	}
	// Remove m.Declarations before unmarshaling so we control its population
	m.Declarations = nil

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	for _, d := range aux.Declarations {
		decl, err := unmarshalDeclaration(d)
		if err != nil {
			return fmt.Errorf("cannot unmarshal declaration: %w", err)
		}
		m.Declarations = append(m.Declarations, decl)
	}
	if m.Declarations == nil {
		m.Declarations = []Declaration{}
	}

	for _, e := range aux.Exports {
		export, err := unmarshalExport(e)
		if err != nil {
			return fmt.Errorf("cannot unmarshal export: %w", err)
		}
		m.Exports = append(m.Exports, export)
	}

	return nil
}
