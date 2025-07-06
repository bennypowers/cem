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
	"errors"
	"fmt"
)

// Demo for custom elements.
type Demo struct {
	Description string           `json:"description,omitempty"`
	URL         string           `json:"url"`
	Source      *SourceReference `json:"source,omitempty"`
}

// CustomElementDeclaration extends ClassDeclaration and CustomElement.
type CustomElementDeclaration struct {
	ClassDeclaration
	CustomElement
}

// CustomElement adds fields to classes/mixins for custom elements.
type CustomElement struct {
	TagName       string              `json:"tagName,omitempty"`
	Attributes    []Attribute         `json:"attributes,omitempty"`
	Events        []Event             `json:"events,omitempty"`
	Slots         []Slot              `json:"slots,omitempty"`
	CssParts      []CssPart           `json:"cssParts,omitempty"`
	CssProperties []CssCustomProperty `json:"cssProperties,omitempty"`
	CssStates     []CssCustomState    `json:"cssStates,omitempty"`
	Demos         []Demo              `json:"demos,omitempty"`
	CustomElement bool                `json:"customElement"`
}

func (*CustomElementDeclaration) isDeclaration()       {}

func (x *CustomElementDeclaration) IsDeprecated() bool {
	if x == nil {
		return false
	}
	return x.Deprecated != nil
}

func (x *CustomElementDeclaration) GetStartByte() uint {
	return x.StartByte
}

func (c *CustomElementDeclaration) UnmarshalJSON(data []byte) (errs error) {
	type Alias CustomElementDeclaration
	aux := &struct {
		Deprecated json.RawMessage   `json:"deprecated"`
		Members    []json.RawMessage `json:"members"`
		*Alias
	}{
		Alias: (*Alias)(c),
	}

	// First, unmarshal into the alias (fills ClassDeclaration and CustomElement)
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Second, unmarshal directly into the embedded CustomElement
	// This ensures CustomElement fields (TagName, Attributes, etc) are populated.
	if err := json.Unmarshal(data, &c.CustomElement); err != nil {
		return err
	}

	// Handle deprecated (custom logic)
	if len(aux.Deprecated) > 0 && string(aux.Deprecated) != "null" {
		var dep Deprecated
		if !decodeDeprecatedField(&dep, aux.Deprecated) {
			return fmt.Errorf("invalid type for deprecated field")
		}
		c.Deprecated = dep
	}

	// Handle members (custom logic)
	for _, m := range aux.Members {
		member, err := unmarshalClassMember(m)
		if member != nil {
			c.Members = append(c.Members, member)
		}
		if err != nil {
			errs = errors.Join(errs, err)
		}
	}

	return errs
}
