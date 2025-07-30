/*
Copyright © 2025 Benny Powers

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
	"errors"
	"strings"
	"testing"
)

func TestErrorStandardization(t *testing.T) {
	baseErr := errors.New("base error")

	tests := []struct {
		name     string
		errFunc  func() error
		expected string
	}{
		{
			name:     "WrapPreprocessError",
			errFunc:  func() error { return WrapPreprocessError(baseErr) },
			expected: "preprocess failed: base error",
		},
		{
			name:     "WrapProcessError",
			errFunc:  func() error { return WrapProcessError(baseErr) },
			expected: "process failed: base error",
		},
		{
			name:     "WrapPostprocessError",
			errFunc:  func() error { return WrapPostprocessError(baseErr) },
			expected: "postprocess failed: base error",
		},
		{
			name:     "WrapIncrementalError",
			errFunc:  func() error { return WrapIncrementalError("module processing", baseErr) },
			expected: "incremental module processing failed: base error",
		},
		{
			name:     "WrapFileError",
			errFunc:  func() error { return WrapFileError("read", "/path/to/file.js", baseErr) },
			expected: "read file \"/path/to/file.js\": base error",
		},
		{
			name:     "WrapModuleError",
			errFunc:  func() error { return WrapModuleError("test/module.js", baseErr) },
			expected: "module \"test/module.js\": base error",
		},
		{
			name:     "WrapComponentError",
			errFunc:  func() error { return WrapComponentError("slot", "header", baseErr) },
			expected: "slot \"header\": base error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.errFunc()
			if err == nil {
				t.Fatal("Expected non-nil error")
			}

			actual := err.Error()
			if actual != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, actual)
			}

			// Verify error wrapping works correctly
			if !errors.Is(err, baseErr) {
				t.Error("Error should unwrap to base error")
			}
		})
	}
}

func TestErrorUtilities(t *testing.T) {
	// Test nil handling
	if err := WrapPreprocessError(nil); err != nil {
		t.Error("Wrapping nil should return nil")
	}

	// Test NewError
	err := NewError("test message")
	if err.Error() != "test message" {
		t.Errorf("NewError: expected 'test message', got %q", err.Error())
	}

	// Test NewErrorf
	err = NewErrorf("test %s with %d", "message", 42)
	expected := "test message with 42"
	if err.Error() != expected {
		t.Errorf("NewErrorf: expected %q, got %q", expected, err.Error())
	}
}

func TestErrorConventions(t *testing.T) {
	// Test that error messages follow lowercase convention
	testCases := []func() error{
		func() error { return WrapPreprocessError(errors.New("base")) },
		func() error { return WrapProcessError(errors.New("base")) },
		func() error { return WrapPostprocessError(errors.New("base")) },
		func() error { return WrapIncrementalError("test", errors.New("base")) },
	}

	for _, errFunc := range testCases {
		err := errFunc()
		msg := err.Error()

		// Error messages should start with lowercase (Go convention)
		if len(msg) > 0 && strings.ToUpper(string(msg[0])) == string(msg[0]) {
			// Allow exceptions for proper nouns or acronyms
			words := strings.Fields(msg)
			if len(words) > 0 && !isProperNounOrAcronym(words[0]) {
				t.Errorf("Error message should start with lowercase: %q", msg)
			}
		}
	}
}

func isProperNounOrAcronym(word string) bool {
	// Common proper nouns and acronyms that should start with uppercase
	properNouns := []string{"QueryManager", "JSON", "HTML", "CSS", "LSP"}
	for _, proper := range properNouns {
		if strings.HasPrefix(word, proper) {
			return true
		}
	}
	return false
}
