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
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestValidateSingleManifestJSONIsArray(t *testing.T) {
	projectDir := setupTest(t, "valid-manifest")
	stdout, _ := runCemCommand(t, projectDir, "validate", "--format=json")

	var results []json.RawMessage
	if err := json.Unmarshal([]byte(stdout), &results); err != nil {
		t.Fatalf("--format=json must produce a JSON array, got parse error: %v\nstdout: %s", err, stdout)
	}
	if len(results) != 1 {
		t.Errorf("expected 1 result, got %d", len(results))
	}
}

func TestValidateWorkspaceJSON(t *testing.T) {
	projectDir := setupTest(t, "workspace-validate")
	stdout, _ := runCemCommand(t, projectDir, "validate", "--format=json")

	var results []json.RawMessage
	if err := json.Unmarshal([]byte(stdout), &results); err != nil {
		t.Fatalf("workspace --format=json must produce a single JSON array, got parse error: %v\nstdout: %s", err, stdout)
	}
	if len(results) != 2 {
		t.Errorf("expected 2 package results, got %d", len(results))
	}
}

func TestValidateE2E(t *testing.T) {
	testCases := []struct {
		name              string
		fixture           string
		args              []string
		expectedStdout    string
		notExpectedStdout string
		expectedStderr    string
		notExpectedStderr string
		expectSuccess     bool
	}{
		{
			name:           "valid_complex_manifest",
			fixture:        "valid-manifest",
			args:           []string{"validate"},
			expectedStderr: "Manifest is valid",
			expectSuccess:  true,
		},
		{
			name:           "valid_manifest_with_warnings",
			fixture:        "warning-lifecycle-methods",
			args:           []string{"validate"},
			expectedStderr: "Manifest valid with",
			expectSuccess:  true,
		},
		{
			name:           "invalid_manifest_multiple_errors",
			fixture:        "multiple-modules-errors",
			args:           []string{"validate"},
			expectedStderr: "Validation failed with",
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
			expectedStderr: "Manifest valid with",
			expectSuccess:  true,
		},
		{
			name:              "quiet_suppresses_success",
			fixture:           "valid-manifest",
			args:              []string{"validate", "-q"},
			notExpectedStderr: "Manifest is valid",
			expectSuccess:     true,
		},
		{
			name:           "default_shows_success",
			fixture:        "valid-manifest",
			args:           []string{"validate"},
			expectedStderr: "Manifest is valid",
			expectSuccess:  true,
		},
		{
			name:           "debug_shows_schema_version",
			fixture:        "valid-manifest",
			args:           []string{"validate", "-vv"},
			expectedStderr: "Schema version: 2.1.1",
			expectSuccess:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Check if fixture exists, skip if not (some may not exist yet)
			fixturePath := filepath.Join("testdata", "fixtures", tc.fixture, "custom-elements.json")
			if _, err := os.Stat(fixturePath); os.IsNotExist(err) {
				t.Skipf("Fixture %s does not exist", tc.fixture)
			}

			projectDir := setupTest(t, tc.fixture)
			stdout, stderr := runCemCommand(t, projectDir, tc.args...)

			// Check expected output
			if tc.expectedStdout != "" && !strings.Contains(stdout, tc.expectedStdout) {
				t.Errorf("stdout did not contain expected string.\nExpected: %q\nGot: %q", tc.expectedStdout, stdout)
			}

			if tc.notExpectedStdout != "" && strings.Contains(stdout, tc.notExpectedStdout) {
				t.Errorf("stdout should not contain %q but did.\nGot: %q", tc.notExpectedStdout, stdout)
			}

			if tc.expectedStderr != "" && !strings.Contains(stderr, tc.expectedStderr) {
				t.Errorf("stderr did not contain expected string.\nExpected: %q\nGot: %q", tc.expectedStderr, stderr)
			}

			if tc.notExpectedStderr != "" && strings.Contains(stderr, tc.notExpectedStderr) {
				t.Errorf("stderr should not contain %q but did.\nGot: %q", tc.notExpectedStderr, stderr)
			}
		})
	}
}
