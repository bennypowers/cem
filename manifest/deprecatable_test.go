/*
Copyright © 2026 Benny Powers

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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsDeprecated_AllTypes(t *testing.T) {
	type deprecatableCase struct {
		name         string
		nilCheck     func() bool
		notDepr      func() bool
		withFlag     func() bool
		withReason   func() bool
	}

	cases := []deprecatableCase{
		{
			"Attribute",
			func() bool { return (*Attribute)(nil).IsDeprecated() },
			func() bool { return (&Attribute{}).IsDeprecated() },
			func() bool { return (&Attribute{Deprecated: DeprecatedFlag(true)}).IsDeprecated() },
			func() bool { return (&Attribute{Deprecated: DeprecatedReason("old")}).IsDeprecated() },
		},
		{
			"Slot",
			func() bool { return (*Slot)(nil).IsDeprecated() },
			func() bool { return (&Slot{}).IsDeprecated() },
			func() bool { return (&Slot{Deprecated: DeprecatedFlag(true)}).IsDeprecated() },
			func() bool { return (&Slot{Deprecated: DeprecatedReason("old")}).IsDeprecated() },
		},
		{
			"Event",
			func() bool { return (*Event)(nil).IsDeprecated() },
			func() bool { return (&Event{}).IsDeprecated() },
			func() bool { return (&Event{Deprecated: DeprecatedFlag(true)}).IsDeprecated() },
			func() bool { return (&Event{Deprecated: DeprecatedReason("old")}).IsDeprecated() },
		},
		{
			"CssCustomProperty",
			func() bool { return (*CssCustomProperty)(nil).IsDeprecated() },
			func() bool { return (&CssCustomProperty{}).IsDeprecated() },
			func() bool { return (&CssCustomProperty{Deprecated: DeprecatedFlag(true)}).IsDeprecated() },
			func() bool { return (&CssCustomProperty{Deprecated: DeprecatedReason("old")}).IsDeprecated() },
		},
		{
			"CssCustomState",
			func() bool { return (*CssCustomState)(nil).IsDeprecated() },
			func() bool { return (&CssCustomState{}).IsDeprecated() },
			func() bool { return (&CssCustomState{Deprecated: DeprecatedFlag(true)}).IsDeprecated() },
			func() bool { return (&CssCustomState{Deprecated: DeprecatedReason("old")}).IsDeprecated() },
		},
		{
			"CssPart",
			func() bool { return (*CssPart)(nil).IsDeprecated() },
			func() bool { return (&CssPart{}).IsDeprecated() },
			func() bool { return (&CssPart{Deprecated: DeprecatedFlag(true)}).IsDeprecated() },
			func() bool { return (&CssPart{Deprecated: DeprecatedReason("old")}).IsDeprecated() },
		},
		{
			"PropertyLike",
			func() bool { return (*PropertyLike)(nil).IsDeprecated() },
			func() bool { return (&PropertyLike{}).IsDeprecated() },
			func() bool { return (&PropertyLike{Deprecated: DeprecatedFlag(true)}).IsDeprecated() },
			func() bool { return (&PropertyLike{Deprecated: DeprecatedReason("old")}).IsDeprecated() },
		},
		{
			"ClassField",
			func() bool { return (*ClassField)(nil).IsDeprecated() },
			func() bool { return (&ClassField{}).IsDeprecated() },
			func() bool {
				return (&ClassField{PropertyLike: PropertyLike{Deprecated: DeprecatedFlag(true)}}).IsDeprecated()
			},
			func() bool {
				return (&ClassField{PropertyLike: PropertyLike{Deprecated: DeprecatedReason("old")}}).IsDeprecated()
			},
		},
		{
			"CustomElementField",
			func() bool { return (*CustomElementField)(nil).IsDeprecated() },
			func() bool { return (&CustomElementField{}).IsDeprecated() },
			func() bool {
				return (&CustomElementField{ClassField: ClassField{PropertyLike: PropertyLike{Deprecated: DeprecatedFlag(true)}}}).IsDeprecated()
			},
			func() bool {
				return (&CustomElementField{ClassField: ClassField{PropertyLike: PropertyLike{Deprecated: DeprecatedReason("old")}}}).IsDeprecated()
			},
		},
		{
			"ClassMethod",
			func() bool { return (*ClassMethod)(nil).IsDeprecated() },
			func() bool { return (&ClassMethod{}).IsDeprecated() },
			func() bool { return (&ClassMethod{Deprecated: DeprecatedFlag(true)}).IsDeprecated() },
			func() bool { return (&ClassMethod{Deprecated: DeprecatedReason("old")}).IsDeprecated() },
		},
		{
			"JavaScriptExport",
			func() bool { return (*JavaScriptExport)(nil).IsDeprecated() },
			func() bool { return (&JavaScriptExport{}).IsDeprecated() },
			func() bool { return (&JavaScriptExport{Deprecated: DeprecatedFlag(true)}).IsDeprecated() },
			func() bool { return (&JavaScriptExport{Deprecated: DeprecatedReason("old")}).IsDeprecated() },
		},
		{
			"CustomElementExport",
			func() bool { return (*CustomElementExport)(nil).IsDeprecated() },
			func() bool { return (&CustomElementExport{}).IsDeprecated() },
			func() bool { return (&CustomElementExport{Deprecated: DeprecatedFlag(true)}).IsDeprecated() },
			func() bool { return (&CustomElementExport{Deprecated: DeprecatedReason("old")}).IsDeprecated() },
		},
		{
			"FunctionDeclaration",
			func() bool { return (*FunctionDeclaration)(nil).IsDeprecated() },
			func() bool { return (&FunctionDeclaration{}).IsDeprecated() },
			func() bool { return (&FunctionDeclaration{Deprecated: DeprecatedFlag(true)}).IsDeprecated() },
			func() bool { return (&FunctionDeclaration{Deprecated: DeprecatedReason("old")}).IsDeprecated() },
		},
		{
			"VariableDeclaration",
			func() bool { return (*VariableDeclaration)(nil).IsDeprecated() },
			func() bool { return (&VariableDeclaration{}).IsDeprecated() },
			func() bool {
				return (&VariableDeclaration{PropertyLike: PropertyLike{Deprecated: DeprecatedFlag(true)}}).IsDeprecated()
			},
			func() bool {
				return (&VariableDeclaration{PropertyLike: PropertyLike{Deprecated: DeprecatedReason("old")}}).IsDeprecated()
			},
		},
		{
			"Module",
			func() bool { return (*Module)(nil).IsDeprecated() },
			func() bool { return (&Module{}).IsDeprecated() },
			func() bool { return (&Module{Deprecated: DeprecatedFlag(true)}).IsDeprecated() },
			func() bool { return (&Module{Deprecated: DeprecatedReason("old")}).IsDeprecated() },
		},
		{
			"Package",
			func() bool { return (*Package)(nil).IsDeprecated() },
			func() bool { return (&Package{}).IsDeprecated() },
			func() bool { return (&Package{Deprecated: DeprecatedFlag(true)}).IsDeprecated() },
			func() bool { return (&Package{Deprecated: DeprecatedReason("old")}).IsDeprecated() },
		},
		{
			"MixinDeclaration",
			func() bool { return (*MixinDeclaration)(nil).IsDeprecated() },
			func() bool { return (&MixinDeclaration{}).IsDeprecated() },
			func() bool { return (&MixinDeclaration{Deprecated: DeprecatedFlag(true)}).IsDeprecated() },
			func() bool { return (&MixinDeclaration{Deprecated: DeprecatedReason("old")}).IsDeprecated() },
		},
		{
			"CustomElementMixinDeclaration",
			func() bool { return (*CustomElementMixinDeclaration)(nil).IsDeprecated() },
			func() bool { return (&CustomElementMixinDeclaration{}).IsDeprecated() },
			func() bool {
				return (&CustomElementMixinDeclaration{
					MixinDeclaration: MixinDeclaration{Deprecated: DeprecatedFlag(true)},
				}).IsDeprecated()
			},
			func() bool {
				return (&CustomElementMixinDeclaration{
					MixinDeclaration: MixinDeclaration{Deprecated: DeprecatedReason("old")},
				}).IsDeprecated()
			},
		},
		{
			"CustomElementDeclaration",
			func() bool { return (*CustomElementDeclaration)(nil).IsDeprecated() },
			func() bool { return (&CustomElementDeclaration{}).IsDeprecated() },
			func() bool {
				return (&CustomElementDeclaration{
					ClassDeclaration: ClassDeclaration{Deprecated: DeprecatedFlag(true)},
				}).IsDeprecated()
			},
			func() bool {
				return (&CustomElementDeclaration{
					ClassDeclaration: ClassDeclaration{Deprecated: DeprecatedReason("old")},
				}).IsDeprecated()
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Run("nil receiver", func(t *testing.T) {
				assert.False(t, tc.nilCheck())
			})
			t.Run("not deprecated", func(t *testing.T) {
				assert.False(t, tc.notDepr())
			})
			t.Run("deprecated flag", func(t *testing.T) {
				assert.True(t, tc.withFlag())
			})
			t.Run("deprecated reason", func(t *testing.T) {
				assert.True(t, tc.withReason())
			})
		})
	}
}
