/*
Copyright © 2025 Benny Powers <web@bennypowers.com>

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
package export

import "testing"

func TestToPascalCase(t *testing.T) {
	tests := []struct {
		input, want string
	}{
		{"my-button", "MyButton"},
		{"demo-element", "DemoElement"},
		{"x", "X"},
		{"", ""},
		{"a-b-c", "ABC"},
		{"hello", "Hello"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := ToPascalCase(tt.input); got != tt.want {
				t.Errorf("ToPascalCase(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestToCamelCase(t *testing.T) {
	tests := []struct {
		input, want string
	}{
		{"my-event", "myEvent"},
		{"demo-button", "demoButton"},
		{"x", "x"},
		{"", ""},
		{"hello", "hello"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := ToCamelCase(tt.input); got != tt.want {
				t.Errorf("ToCamelCase(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestToKebabCase(t *testing.T) {
	tests := []struct {
		input, want string
	}{
		{"DemoButton", "demo-button"},
		{"MyElement", "my-element"},
		{"X", "x"},
		{"", ""},
		{"hello", "hello"},
		{"ABCDef", "a-b-c-def"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := ToKebabCase(tt.input); got != tt.want {
				t.Errorf("ToKebabCase(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestTagNameToComponentName(t *testing.T) {
	tests := []struct {
		tagName, prefix, want string
	}{
		{"demo-button", "demo-", "Button"},
		{"demo-input-field", "demo-", "InputField"},
		{"my-element", "", "MyElement"},
		{"demo-button", "other-", "DemoButton"},
	}
	for _, tt := range tests {
		t.Run(tt.tagName+"/"+tt.prefix, func(t *testing.T) {
			if got := TagNameToComponentName(tt.tagName, tt.prefix); got != tt.want {
				t.Errorf("TagNameToComponentName(%q, %q) = %q, want %q", tt.tagName, tt.prefix, got, tt.want)
			}
		})
	}
}

func TestToReactEventName(t *testing.T) {
	tests := []struct {
		input, want string
	}{
		{"my-event", "onMyEvent"},
		{"click", "onClick"},
		{"color-changed", "onColorChanged"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := ToReactEventName(tt.input); got != tt.want {
				t.Errorf("ToReactEventName(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestToAngularEventName(t *testing.T) {
	tests := []struct {
		input, want string
	}{
		{"my-event", "myEvent"},
		{"click", "click"},
		{"color-changed", "colorChanged"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := ToAngularEventName(tt.input); got != tt.want {
				t.Errorf("ToAngularEventName(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestToVueEventName(t *testing.T) {
	tests := []struct {
		input, want string
	}{
		{"my-event", "my-event"},
		{"click", "click"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := ToVueEventName(tt.input); got != tt.want {
				t.Errorf("ToVueEventName(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
