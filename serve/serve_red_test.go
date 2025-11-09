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

package serve_test

import (
	"strings"
	"testing"

	"bennypowers.dev/cem/serve"
)

// RED PHASE TESTS
// These tests verify that stub methods properly fail until implemented
// They should PASS in Phase 0 (verifying stubs return errors)
// They will be updated in their respective phases to test actual functionality

// TestImportMapStub_ReturnsNil is removed - Phase 2 tests are in importmap_test.go
// This stub test is no longer needed as we're implementing Phase 2

// TestTransformTypeScriptStub_ReturnsNotImplemented verifies TypeScript transform stub fails
func TestTransformTypeScriptStub_ReturnsNotImplemented(t *testing.T) {
	_, err := serve.TransformTypeScript([]byte("const x = 1;"))
	if err == nil {
		t.Fatal("Expected TransformTypeScript to return error (not implemented)")
	}
	if !strings.Contains(err.Error(), "not implemented") {
		t.Errorf("Expected 'not implemented' error, got: %v", err)
	}
	if !strings.Contains(err.Error(), "Phase 4") {
		t.Errorf("Expected error to mention 'Phase 4', got: %v", err)
	}
}

// TestTransformCSSStub_ReturnsNotImplemented verifies CSS transform stub fails
func TestTransformCSSStub_ReturnsNotImplemented(t *testing.T) {
	_, err := serve.TransformCSS([]byte("body { color: red; }"))
	if err == nil {
		t.Fatal("Expected TransformCSS to return error (not implemented)")
	}
	if !strings.Contains(err.Error(), "not implemented") {
		t.Errorf("Expected 'not implemented' error, got: %v", err)
	}
	if !strings.Contains(err.Error(), "Phase 4") {
		t.Errorf("Expected error to mention 'Phase 4', got: %v", err)
	}
}
