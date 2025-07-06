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

var _ Deprecatable = (*ClassLike)(nil)
var _ Deprecatable = (*ClassDeclaration)(nil)

// ClassLike is the common interface of classes and mixins.
type ClassLike struct {
	FullyQualified
	StartByte  uint             `json:"-"`
	Superclass *Reference       `json:"superclass,omitempty"`
	Mixins     []Reference      `json:"mixins,omitempty"`
	Members    []ClassMember    `json:"members,omitempty"`
	Source     *SourceReference `json:"source,omitempty"`
	Deprecated Deprecated       `json:"deprecated,omitempty"` // bool or string
}

func (x *ClassLike) IsDeprecated() bool {
	if x == nil {
		return false
	}
	return x.Deprecated != nil
}

// ClassDeclaration is a class.
type ClassDeclaration struct {
	ClassLike
	Kind string `json:"kind"` // 'class'
}

func (*ClassDeclaration) isDeclaration() {}

func (x *ClassDeclaration) IsDeprecated() bool {
	if x == nil {
		return false
	}
	return x.Deprecated != nil
}

func (x *ClassDeclaration) GetStartByte() uint { return x.StartByte }

func (c *ClassDeclaration) UnmarshalJSON(data []byte) error {
	type Alias ClassDeclaration
	aux := &struct {
		Deprecated json.RawMessage   `json:"deprecated"`
		Members    []json.RawMessage `json:"members"`
		*Alias
	}{
		Alias: (*Alias)(c),
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
		c.Deprecated = dep
	}
	// Handle members
	for _, m := range aux.Members {
		member, err := unmarshalClassMember(m)
		if err == nil && member != nil {
			c.Members = append(c.Members, member)
		}
	}
	return nil
}
