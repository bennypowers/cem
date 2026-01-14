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
package workspace_test

import (
	"testing"

	"bennypowers.dev/cem/workspace"
	"github.com/stretchr/testify/assert"
)

func TestIsURLSpecifier(t *testing.T) {
	tests := []struct {
		name     string
		spec     string
		expected bool
	}{
		// HTTP URLs
		{name: "https URL", spec: "https://cdn.jsdelivr.net/npm/@vaadin/button/", expected: true},
		{name: "http URL", spec: "http://example.com/package/", expected: true},
		{name: "https with path and version", spec: "https://esm.sh/@patternfly/elements@4.0.0/", expected: true},
		{name: "https unpkg", spec: "https://unpkg.com/@shortfuse/materialdesignweb/", expected: true},

		// npm specifiers (should be false)
		{name: "npm specifier", spec: "npm:@vaadin/button@24.3.5", expected: false},
		{name: "npm specifier no version", spec: "npm:lodash", expected: false},

		// jsr specifiers (should be false)
		{name: "jsr specifier", spec: "jsr:@std/testing", expected: false},

		// Local paths (should be false)
		{name: "relative path", spec: "./node_modules/@vaadin/button", expected: false},
		{name: "absolute path", spec: "/home/user/project", expected: false},
		{name: "dot path", spec: ".", expected: false},

		// Edge cases
		{name: "empty string", spec: "", expected: false},
		{name: "https prefix but invalid", spec: "https", expected: false},
		{name: "http prefix but invalid", spec: "http", expected: false},
		{name: "contains https but not prefix", spec: "file://https://example.com", expected: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := workspace.IsURLSpecifier(tt.spec)
			assert.Equal(t, tt.expected, result, "IsURLSpecifier(%q)", tt.spec)
		})
	}
}

func TestIsPackageSpecifier(t *testing.T) {
	tests := []struct {
		name     string
		spec     string
		expected bool
	}{
		// npm specifiers
		{name: "npm scoped with version", spec: "npm:@vaadin/button@24.3.5", expected: true},
		{name: "npm scoped no version", spec: "npm:@patternfly/elements", expected: true},
		{name: "npm unscoped", spec: "npm:lodash", expected: true},

		// jsr specifiers
		{name: "jsr scoped", spec: "jsr:@std/testing", expected: true},
		{name: "jsr unscoped", spec: "jsr:testing", expected: true},

		// URLs (should be false)
		{name: "https URL", spec: "https://cdn.jsdelivr.net/npm/@vaadin/button/", expected: false},
		{name: "http URL", spec: "http://example.com/package/", expected: false},

		// Local paths (should be false)
		{name: "relative path", spec: "./node_modules/@vaadin/button", expected: false},
		{name: "absolute path", spec: "/home/user/project", expected: false},
		{name: "dot path", spec: ".", expected: false},

		// Edge cases
		{name: "empty string", spec: "", expected: false},
		{name: "npm prefix but invalid", spec: "npm", expected: false},
		{name: "jsr prefix but invalid", spec: "jsr", expected: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := workspace.IsPackageSpecifier(tt.spec)
			assert.Equal(t, tt.expected, result, "IsPackageSpecifier(%q)", tt.spec)
		})
	}
}
