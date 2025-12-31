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
package helpers

import (
	"testing"
)

func TestPathsMatch(t *testing.T) {
	tests := []struct {
		name          string
		importPath    string
		elementSource string
		expected      bool
		description   string
	}{
		{
			name:          "ExactMatch",
			importPath:    "my-package/button.js",
			elementSource: "my-package/button.js",
			expected:      true,
			description:   "Exact path match should return true",
		},
		{
			name:          "PackageLevelImport_IndexFile",
			importPath:    "large-component-library",
			elementSource: "large-component-library/index.js",
			expected:      true,
			description:   "Package-level import should match elements from that package's index",
		},
		{
			name:          "PackageLevelImport_Submodule",
			importPath:    "large-component-library",
			elementSource: "large-component-library/button/button.js",
			expected:      true,
			description:   "Package-level import should match all elements from that package",
		},
		{
			name:          "ScopedPackageLevelImport",
			importPath:    "@rhds/elements",
			elementSource: "@rhds/elements/rh-card/rh-card.js",
			expected:      true,
			description:   "Scoped package import should match elements from that package",
		},
		{
			name:          "SubpathImport",
			importPath:    "my-package/utils/helpers.js",
			elementSource: "my-package/utils/helpers.js",
			expected:      true,
			description:   "Subpath imports should match exactly",
		},
		{
			name:          "RelativeImport",
			importPath:    "./button.js",
			elementSource: "button.js",
			expected:      true,
			description:   "Relative imports should match when normalized",
		},
		{
			name:          "DifferentPackage",
			importPath:    "other-package",
			elementSource: "my-package/index.js",
			expected:      false,
			description:   "Different packages should not match",
		},
		{
			name:          "SubpathNotMatching",
			importPath:    "my-package/button.js",
			elementSource: "my-package/card.js",
			expected:      false,
			description:   "Different subpaths should not match",
		},
		{
			name:          "PackagePrefixNotMatching",
			importPath:    "my-component",
			elementSource: "my-component-library/index.js",
			expected:      false,
			description:   "Package name prefix should not match if it's a different package",
		},
		{
			name:          "SelfImport_LocalPackage",
			importPath:    "cem-lsp-demo",
			elementSource: "cem-lsp-demo/components/button-element.js",
			expected:      true,
			description:   "Local package importing itself by name should match its own elements",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PathsMatch(tt.importPath, tt.elementSource)
			if result != tt.expected {
				t.Errorf("PathsMatch(%q, %q) = %v, expected %v\n  Description: %s",
					tt.importPath, tt.elementSource, result, tt.expected, tt.description)
			}
		})
	}
}

func TestNormalizePathForMatching(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "BarePackageName",
			input:    "large-component-library",
			expected: "large-component-library",
		},
		{
			name:     "PackageWithSubpath",
			input:    "large-component-library/index.js",
			expected: "large-component-library/index.js",
		},
		{
			name:     "ScopedPackage",
			input:    "@rhds/elements",
			expected: "@rhds/elements",
		},
		{
			name:     "ScopedPackageWithSubpath",
			input:    "@rhds/elements/rh-card/rh-card.js",
			expected: "rh-card/rh-card.js",
		},
		{
			name:     "RelativePath",
			input:    "./button.js",
			expected: "button.js",
		},
		{
			name:     "ParentRelativePath",
			input:    "../shared/utils.js",
			expected: "shared/utils.js",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NormalizePathForMatching(tt.input)
			if result != tt.expected {
				t.Errorf("NormalizePathForMatching(%q) = %q, expected %q",
					tt.input, result, tt.expected)
			}
		})
	}
}
