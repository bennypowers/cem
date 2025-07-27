//go:build e2e

package cmd_test

import (
	"strings"
	"testing"
)

func TestValidateE2E(t *testing.T) {
	testCases := []struct {
		fixture        string
		expectedStdout string
		expectedStderr string
	}{
		{
			fixture:        "valid-manifest",
			expectedStdout: "✓ Manifest is valid",
		},
		{
			fixture:        "invalid-class",
			expectedStdout: "✗ Validation failed with 1 issue",
		},
		{
			fixture:        "invalid-mixin",
			expectedStdout: "✗ Validation failed with 1 issue",
		},
		{
			fixture:        "invalid-variable",
			expectedStdout: "✗ Validation failed with 1 issue",
		},
		{
			fixture:        "invalid-function",
			expectedStdout: "✗ Validation failed with 1 issue",
		},
		{
			fixture:        "invalid-field",
			expectedStdout: "✗ Validation failed with 1 issue",
		},
		{
			fixture:        "invalid-method",
			expectedStdout: "✗ Validation failed with 1 issue",
		},
		{
			fixture:        "invalid-event",
			expectedStdout: "✗ Validation failed with 1 issue",
		},
		{
			fixture:        "invalid-attribute",
			expectedStdout: "✗ Validation failed with 1 issue",
		},
		{
			fixture:        "invalid-slot",
			expectedStdout: "✗ Validation failed with 1 issue",
		},
		{
			fixture:        "invalid-css-part",
			expectedStdout: "✗ Validation failed with 1 issue",
		},
		{
			fixture:        "invalid-css-custom-property",
			expectedStdout: "✗ Validation failed with 1 issue",
		},
		{
			fixture:        "unsupported-schema",
			expectedStdout: "⚠ Manifest valid with 1 warning",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.fixture, func(t *testing.T) {
			projectDir := setupTest(t, tc.fixture)
			// The command should be run from the project root, and it should find the manifest via package.json
			stdout, stderr := runCemCommand(t, projectDir, "validate")

			if !strings.Contains(stdout, tc.expectedStdout) {
				t.Errorf("stdout did not contain expected string.\nExpected: %q\nGot: %q", tc.expectedStdout, stdout)
			}

			if !strings.Contains(stderr, tc.expectedStderr) {
				t.Errorf("stderr did not contain expected string.\nExpected: %q\nGot: %q", tc.expectedStderr, stderr)
			}
		})
	}
}
