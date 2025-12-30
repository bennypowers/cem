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
package manifest_test

import (
	"encoding/json"
	"testing"

	"bennypowers.dev/cem/internal/platform/testutil"
	"bennypowers.dev/cem/manifest"
)

// TestPackage_FindCustomElementDeclaration tests finding custom element declarations by tag name
func TestPackage_FindCustomElementDeclaration(t *testing.T) {
	fs := testutil.NewFixtureFS(t, "element-description-lookup", "/test")
	manifestBytes, err := fs.ReadFile("/test/manifest.json")
	if err != nil {
		t.Fatalf("Failed to read test manifest: %v", err)
	}

	var pkg manifest.Package
	if err := json.Unmarshal(manifestBytes, &pkg); err != nil {
		t.Fatalf("Failed to parse manifest: %v", err)
	}

	tests := []struct {
		name                string
		tagName             string
		expectedName        string
		expectedSummary     string
		expectedDescription string
		shouldFind          bool
	}{
		{
			name:                "find my-button with summary and description",
			tagName:             "my-button",
			expectedName:        "MyButton",
			expectedSummary:     "Primary button component",
			expectedDescription: "A versatile button component for user interactions. Supports multiple variants, sizes, and states to match your design system.",
			shouldFind:          true,
		},
		{
			name:                "find my-card with summary and description",
			tagName:             "my-card",
			expectedName:        "MyCard",
			expectedSummary:     "Card container component",
			expectedDescription: "A flexible card layout component with slots for header, content, and footer sections.",
			shouldFind:          true,
		},
		{
			name:                "find my-input without summary or description",
			tagName:             "my-input",
			expectedName:        "MyInput",
			expectedSummary:     "",
			expectedDescription: "",
			shouldFind:          true,
		},
		{
			name:       "return nil for non-existent element",
			tagName:    "non-existent-element",
			shouldFind: false,
		},
		{
			name:       "return nil for empty tag name",
			tagName:    "",
			shouldFind: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decl := (&pkg).FindCustomElementDeclaration(tt.tagName)

			if !tt.shouldFind {
				if decl != nil {
					t.Errorf("Expected nil for tag name %q, got %v", tt.tagName, decl)
				}
				return
			}

			if decl == nil {
				t.Fatalf("Expected to find declaration for tag name %q, got nil", tt.tagName)
			}

			if decl.Name != tt.expectedName {
				t.Errorf("Expected name %q, got %q", tt.expectedName, decl.Name)
			}

			if decl.Summary != tt.expectedSummary {
				t.Errorf("Expected summary %q, got %q", tt.expectedSummary, decl.Summary)
			}

			if decl.Description != tt.expectedDescription {
				t.Errorf("Expected description %q, got %q", tt.expectedDescription, decl.Description)
			}

			if decl.TagName != tt.tagName {
				t.Errorf("Expected tagName %q, got %q", tt.tagName, decl.TagName)
			}
		})
	}
}

// TestPackage_FindCustomElementDeclaration_NilPackage tests that nil package returns nil
func TestPackage_FindCustomElementDeclaration_NilPackage(t *testing.T) {
	var pkg *manifest.Package
	decl := pkg.FindCustomElementDeclaration("my-button")
	if decl != nil {
		t.Errorf("Expected nil for nil package, got %v", decl)
	}
}
