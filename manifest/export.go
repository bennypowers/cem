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

var _ Deprecatable = (*JavaScriptExport)(nil)
var _ Deprecatable = (*CustomElementExport)(nil)

// Export is a union type: JavaScriptExport or CustomElementExport.
type Export interface {
	isExport()
	GetStartByte() uint
}

// JavaScriptExport represents a JS export.
type JavaScriptExport struct {
	StartByte   uint       `json:"-"`
	Kind        string     `json:"kind"` // 'js'
	Name        string     `json:"name"`
	Declaration *Reference `json:"declaration"`
	Deprecated  Deprecated `json:"deprecated,omitempty"` // bool or string
}

// CustomElementExport represents a custom element definition.
type CustomElementExport struct {
	StartByte   uint       `json:"-"`
	Kind        string     `json:"kind"` // 'custom-element-definition'
	Name        string     `json:"name"`
	Declaration *Reference `json:"declaration"`
	Deprecated  any        `json:"deprecated,omitempty"` // bool or string
}

func NewCustomElementExport(
	tagName string,
	declaration *Reference,
	startByte uint,
	deprecated *Deprecated,
) *CustomElementExport {
	ce := &CustomElementExport{
		Kind:        "custom-element-definition",
		StartByte:   startByte,
		Name:        tagName,
		Declaration: declaration,
	}
	if deprecated != nil {
		ce.Deprecated = deprecated
	}
	return ce
}

func (*JavaScriptExport) isExport() {}

func (x *JavaScriptExport) IsDeprecated() bool {
	if x == nil {
		return false
	}
	return x.Deprecated != nil
}

func (e *JavaScriptExport) GetStartByte() uint {
	return e.StartByte
}

// JavaScriptExport implements custom unmarshaling for standard JS exports.
func (e *JavaScriptExport) UnmarshalJSON(data []byte) error {
	type Rest JavaScriptExport
	aux := &struct {
		Deprecated json.RawMessage `json:"deprecated"`
		*Rest
	}{
		Rest: (*Rest)(e),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if len(aux.Deprecated) > 0 && string(aux.Deprecated) != "null" {
		var dep Deprecated
		if !decodeDeprecatedField(&dep, aux.Deprecated) {
			return fmt.Errorf("invalid type for deprecated field")
		}
		e.Deprecated = dep
	}
	return nil
}

func (*CustomElementExport) isExport() {}

func (x *CustomElementExport) IsDeprecated() bool {
	if x == nil {
		return false
	}
	return x.Deprecated != nil
}

func (e *CustomElementExport) GetStartByte() uint {
	return e.StartByte
}

// CustomElementExport implements custom unmarshaling for custom element exports.
func (e *CustomElementExport) UnmarshalJSON(data []byte) error {
	type Rest CustomElementExport
	aux := &struct {
		Deprecated json.RawMessage `json:"deprecated"`
		*Rest
	}{
		Rest: (*Rest)(e),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if len(aux.Deprecated) > 0 && string(aux.Deprecated) != "null" {
		var dep Deprecated
		if !decodeDeprecatedField(&dep, aux.Deprecated) {
			return fmt.Errorf("invalid type for deprecated field")
		}
		e.Deprecated = dep
	}
	return nil
}

func unmarshalExport(data json.RawMessage) (Export, error) {
	var kindWrap struct {
		Kind string `json:"kind"`
	}
	if err := json.Unmarshal(data, &kindWrap); err != nil {
		return nil, err
	}
	switch kindWrap.Kind {
	case "js":
		var e JavaScriptExport
		if err := json.Unmarshal(data, &e); err != nil {
			return nil, err
		}
		return &e, nil
	case "custom-element-definition":
		var e CustomElementExport
		if err := json.Unmarshal(data, &e); err != nil {
			return nil, err
		}
		return &e, nil
	default:
		return nil, fmt.Errorf("unknown export kind: %s", kindWrap.Kind)
	}
}
