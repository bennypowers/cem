//go:build e2e

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

package cmd_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestValidateE2E(t *testing.T) {
	testCases := []struct {
		name           string
		fixture        string
		args           []string
		expectedStdout string
		expectedStderr string
		expectSuccess  bool
	}{
		{
			name:           "valid_complex_manifest",
			fixture:        "valid-manifest",
			args:           []string{"validate"},
			expectedStdout: "✓ Manifest is valid",
			expectSuccess:  true,
		},
		{
			name:           "valid_manifest_with_warnings",
			fixture:        "warning-lifecycle-methods",
			args:           []string{"validate"},
			expectedStdout: "⚠ Manifest valid with",
			expectSuccess:  true,
		},
		{
			name:           "invalid_manifest_multiple_errors",
			fixture:        "multiple-modules-errors",
			args:           []string{"validate"},
			expectedStdout: "✗ Validation failed with",
			expectSuccess:  false,
		},
		{
			name:           "valid_manifest_json_output",
			fixture:        "valid-manifest",
			args:           []string{"validate", "--format=json"},
			expectedStdout: `"valid": true`,
			expectSuccess:  true,
		},
		{
			name:           "invalid_manifest_json_output",
			fixture:        "invalid-class",
			args:           []string{"validate", "--format=json"},
			expectedStdout: `"valid": false`,
			expectSuccess:  false,
		},
		{
			name:           "manifest_with_disabled_warnings",
			fixture:        "warning-lifecycle-methods",
			args:           []string{"validate", "--disable=lifecycle"},
			expectedStdout: "⚠ Manifest valid with", // Still has other warnings
			expectSuccess:  true,
		},
		{
			name:           "verbose_output",
			fixture:        "valid-manifest",
			args:           []string{"validate", "--verbose"},
			expectedStdout: "Schema version: 2.1.1",
			expectSuccess:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Check if fixture exists, skip if not (some may not exist yet)
			fixturePath := filepath.Join("fixture", tc.fixture, "custom-elements.json")
			if _, err := os.Stat(fixturePath); os.IsNotExist(err) {
				t.Skipf("Fixture %s does not exist", tc.fixture)
			}

			projectDir := setupTest(t, tc.fixture)
			stdout, stderr := runCemCommand(t, projectDir, tc.args...)

			// Check expected output
			if tc.expectedStdout != "" && !strings.Contains(stdout, tc.expectedStdout) {
				t.Errorf("stdout did not contain expected string.\nExpected: %q\nGot: %q", tc.expectedStdout, stdout)
			}

			if tc.expectedStderr != "" && !strings.Contains(stderr, tc.expectedStderr) {
				t.Errorf("stderr did not contain expected string.\nExpected: %q\nGot: %q", tc.expectedStderr, stderr)
			}
		})
	}
}
