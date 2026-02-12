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

import (
	"strings"
	"unicode"
)

// ToPascalCase converts a kebab-case string to PascalCase.
// e.g. "my-button" -> "MyButton"
func ToPascalCase(s string) string {
	parts := strings.Split(s, "-")
	var b strings.Builder
	for _, part := range parts {
		if part == "" {
			continue
		}
		b.WriteString(upperFirst(part))
	}
	return b.String()
}

// ToCamelCase converts a kebab-case string to camelCase.
// e.g. "my-event" -> "myEvent"
func ToCamelCase(s string) string {
	pascal := ToPascalCase(s)
	if pascal == "" {
		return ""
	}
	return lowerFirst(pascal)
}

// ToKebabCase converts a PascalCase or camelCase string to kebab-case.
// e.g. "DemoButton" -> "demo-button"
func ToKebabCase(s string) string {
	if s == "" {
		return ""
	}
	var b strings.Builder
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				b.WriteByte('-')
			}
			b.WriteRune(unicode.ToLower(r))
		} else {
			b.WriteRune(r)
		}
	}
	return b.String()
}

// TagNameToComponentName converts a custom element tag name to a component name,
// optionally stripping a prefix.
// e.g. TagNameToComponentName("demo-button", "demo-") -> "Button"
// e.g. TagNameToComponentName("my-element", "") -> "MyElement"
func TagNameToComponentName(tagName, stripPrefix string) string {
	name := tagName
	if stripPrefix != "" {
		name = strings.TrimPrefix(name, stripPrefix)
	}
	return ToPascalCase(name)
}

// ToReactEventName converts a kebab-case event name to a React event handler prop name.
// e.g. "my-event" -> "onMyEvent"
func ToReactEventName(eventName string) string {
	return "on" + ToPascalCase(eventName)
}

// ToAngularEventName converts a kebab-case event name to an Angular camelCase event name.
// e.g. "my-event" -> "myEvent"
func ToAngularEventName(eventName string) string {
	return ToCamelCase(eventName)
}

// ToVueEventName returns the event name as-is for Vue (kebab-case is idiomatic).
// e.g. "my-event" -> "my-event"
func ToVueEventName(eventName string) string {
	return eventName
}

func upperFirst(s string) string {
	if s == "" {
		return ""
	}
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func lowerFirst(s string) string {
	if s == "" {
		return ""
	}
	r := []rune(s)
	r[0] = unicode.ToLower(r[0])
	return string(r)
}
