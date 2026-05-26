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
package generate

import (
	"testing"

	Q "bennypowers.dev/cem/internal/treesitter"
	"github.com/stretchr/testify/assert"
)

// Inline: pure function, table-driven

func TestIsIgnoredMember(t *testing.T) {
	tests := []struct {
		name       string
		memberName string
		superclass string
		isStatic   bool
		want       bool
	}{
		// HTMLElement ignored members
		{"HTMLElement static formAssociated", "formAssociated", "HTMLElement", true, true},
		{"HTMLElement static observedAttributes", "observedAttributes", "HTMLElement", true, true},
		{"HTMLElement instance connectedCallback", "connectedCallback", "HTMLElement", false, true},
		{"HTMLElement instance disconnectedCallback", "disconnectedCallback", "HTMLElement", false, true},
		{"HTMLElement instance attributeChangedCallback", "attributeChangedCallback", "HTMLElement", false, true},
		{"HTMLElement instance adoptedCallback", "adoptedCallback", "HTMLElement", false, true},
		{"HTMLElement instance formDisabledCallback", "formDisabledCallback", "HTMLElement", false, true},
		{"HTMLElement instance formStateRestoreCallback", "formStateRestoreCallback", "HTMLElement", false, true},
		{"HTMLElement instance connectedMoveCallback", "connectedMoveCallback", "HTMLElement", false, true},
		{"HTMLElement regular method kept", "onClick", "HTMLElement", false, false},
		{"HTMLElement regular field kept", "myProp", "HTMLElement", true, false},

		// LitElement includes HTMLElement + Lit-specific
		{"LitElement static styles", "styles", "LitElement", true, true},
		{"LitElement static shadowRootOptions", "shadowRootOptions", "LitElement", true, true},
		{"LitElement static formAssociated", "formAssociated", "LitElement", true, true},
		{"LitElement instance render", "render", "LitElement", false, true},
		{"LitElement instance updated", "updated", "LitElement", false, true},
		{"LitElement instance firstUpdated", "firstUpdated", "LitElement", false, true},
		{"LitElement instance update", "update", "LitElement", false, true},
		{"LitElement instance willUpdate", "willUpdate", "LitElement", false, true},
		{"LitElement instance getUpdateComplete", "getUpdateComplete", "LitElement", false, true},
		{"LitElement instance connectedCallback", "connectedCallback", "LitElement", false, true},
		{"LitElement regular method kept", "handleClick", "LitElement", false, false},

		// Unknown superclass
		{"unknown superclass", "render", "SomeOtherClass", false, false},
		{"empty superclass", "connectedCallback", "", false, false},

		// Static vs instance mismatch
		{"HTMLElement instance field not in static set", "formAssociated", "HTMLElement", false, false},
		{"HTMLElement static method not in instance set", "connectedCallback", "HTMLElement", true, false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, isIgnoredMember(tc.memberName, tc.superclass, tc.isStatic))
		})
	}
}

func TestIsPropertyField(t *testing.T) {
	tests := []struct {
		name     string
		captures Q.CaptureMap
		want     bool
	}{
		{
			"has property decorator",
			Q.CaptureMap{"decorator.name": {{Text: "property"}}},
			true,
		},
		{
			"has other decorator",
			Q.CaptureMap{"decorator.name": {{Text: "state"}}},
			false,
		},
		{
			"multiple decorators with property",
			Q.CaptureMap{"decorator.name": {{Text: "state"}, {Text: "property"}}},
			true,
		},
		{
			"no decorators key",
			Q.CaptureMap{"field": {{Text: "myField"}}},
			false,
		},
		{
			"empty capture map",
			Q.CaptureMap{},
			false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, isPropertyField(tc.captures))
		})
	}
}

func TestGetMemberKindFromCaptures(t *testing.T) {
	tests := []struct {
		name     string
		captures Q.CaptureMap
		want     string
	}{
		{"method", Q.CaptureMap{"method": {{Text: "foo"}}}, "method"},
		{"accessor", Q.CaptureMap{"accessor": {{Text: "bar"}}}, "accessor"},
		{"field", Q.CaptureMap{"field": {{Text: "baz"}}}, "field"},
		{"constructor parameter", Q.CaptureMap{"constructor.parameter": {{Text: "x"}}}, "constructor-parameter"},
		{"empty", Q.CaptureMap{}, ""},
		{"unknown key", Q.CaptureMap{"other": {{Text: "x"}}}, ""},
		{"method takes precedence", Q.CaptureMap{"method": {{Text: "a"}}, "field": {{Text: "b"}}}, "method"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, getMemberKindFromCaptures(tc.captures))
		})
	}
}

func TestIsStaticToTypeFlag(t *testing.T) {
	assert.Equal(t, "static", isStaticToTypeFlag(true))
	assert.Equal(t, "instance", isStaticToTypeFlag(false))
}
