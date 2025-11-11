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

package serve

import "testing"

// TestContainsPath_ExactMatch tests that exact matches return true
func TestContainsPath_ExactMatch(t *testing.T) {
	tests := []struct {
		fullURL string
		path    string
		want    bool
	}{
		{"/elements/button", "/elements/button", true},
		{"/demo/basic.html", "/demo/basic.html", true},
		{"/", "/", true},
	}

	for _, tt := range tests {
		t.Run(tt.fullURL, func(t *testing.T) {
			if got := containsPath(tt.fullURL, tt.path); got != tt.want {
				t.Errorf("containsPath(%q, %q) = %v, want %v", tt.fullURL, tt.path, got, tt.want)
			}
		})
	}
}

// TestContainsPath_ValidDelimiters tests that paths followed by valid delimiters return true
func TestContainsPath_ValidDelimiters(t *testing.T) {
	tests := []struct {
		name    string
		fullURL string
		path    string
		want    bool
	}{
		{"slash delimiter", "/elements/button/demo/", "/elements/button", true},
		{"query delimiter", "/elements/button?variant=primary", "/elements/button", true},
		{"fragment delimiter", "/elements/button#overview", "/elements/button", true},
		{"path ending with slash", "/elements/accordion/demo/", "/elements/accordion", true},
		{"nested path", "/elements/accordion/demo/accents", "/elements/accordion/demo", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsPath(tt.fullURL, tt.path); got != tt.want {
				t.Errorf("containsPath(%q, %q) = %v, want %v", tt.fullURL, tt.path, got, tt.want)
			}
		})
	}
}

// TestContainsPath_InvalidPrefixMatch tests that invalid prefix matches return false
// This is the key test that demonstrates the bug: "/elements/button" should NOT match "/elements/buttonly"
func TestContainsPath_InvalidPrefixMatch(t *testing.T) {
	tests := []struct {
		name    string
		fullURL string
		path    string
		want    bool
	}{
		{"partial word match", "/elements/buttonly", "/elements/button", false},
		{"partial with hyphen", "/elements/button-group", "/elements/button", false},
		{"partial with underscore", "/elements/button_component", "/elements/button", false},
		{"similar prefix", "/api/v1/users-list", "/api/v1/users", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsPath(tt.fullURL, tt.path); got != tt.want {
				t.Errorf("containsPath(%q, %q) = %v, want %v", tt.fullURL, tt.path, got, tt.want)
			}
		})
	}
}

// TestContainsPath_EdgeCases tests edge cases and invalid inputs
func TestContainsPath_EdgeCases(t *testing.T) {
	tests := []struct {
		name    string
		fullURL string
		path    string
		want    bool
	}{
		{"empty path", "/elements/button", "", false},
		{"path longer than fullURL", "/short", "/very/long/path", false},
		{"empty fullURL", "", "/path", false},
		{"both empty", "", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsPath(tt.fullURL, tt.path); got != tt.want {
				t.Errorf("containsPath(%q, %q) = %v, want %v", tt.fullURL, tt.path, got, tt.want)
			}
		})
	}
}
