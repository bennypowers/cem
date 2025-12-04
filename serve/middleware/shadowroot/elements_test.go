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

package shadowroot

import (
	"embed"
	"testing"
)

//go:embed testdata/**
var testFS embed.FS

func TestBuildKnownElements(t *testing.T) {
	// Test with actual testFS using correct prefix
	knownElements := buildKnownElementsWithPrefix(testFS, "testdata/templates/elements")

	// Check what we found
	t.Logf("Found %d known elements", len(knownElements))

	// Verify expected elements are found
	expected := []string{"test-button", "outer-element", "inner-element", "test-icon"}
	for _, elem := range expected {
		if knownElements.Has(elem) {
			t.Logf("✓ %s found", elem)
		} else {
			t.Errorf("✗ %s NOT found", elem)
		}
	}
}
