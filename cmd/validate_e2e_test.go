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
			expectedStdout: "Manifest is valid.",
		},
		{
			fixture:        "invalid-class",
			expectedStderr: "Manifest is invalid:\n- /modules/0/declarations/0: missing properties: 'name'\n",
		},
		{
			fixture:        "invalid-mixin",
			expectedStderr: "Manifest is invalid:\n- /modules/0/declarations/1: missing properties: 'name'\n",
		},
		{
			fixture:        "invalid-variable",
			expectedStderr: "Manifest is invalid:\n- /modules/0/declarations/2: missing properties: 'name'\n",
		},
		{
			fixture:        "invalid-function",
			expectedStderr: "Manifest is invalid:\n- /modules/0/declarations/3: missing properties: 'name'\n",
		},
		{
			fixture:        "invalid-field",
			expectedStderr: "Manifest is invalid:\n- /modules/0/declarations/0/members/0: missing properties: 'name'\n",
		},
		{
			fixture:        "invalid-method",
			expectedStderr: "Manifest is invalid:\n- /modules/0/declarations/0/members/1: missing properties: 'name'\n",
		},
		{
			fixture:        "invalid-event",
			expectedStderr: "Manifest is invalid:\n- /modules/0/declarations/0/events/0: missing properties: 'name'\n",
		},
		{
			fixture:        "invalid-attribute",
			expectedStderr: "Manifest is invalid:\n- /modules/0/declarations/0/attributes/0: missing properties: 'name'\n",
		},
		{
			fixture:        "invalid-slot",
			expectedStderr: "Manifest is invalid:\n- /modules/0/declarations/0/slots/0: missing properties: 'name'\n",
		},
		{
			fixture:        "invalid-css-part",
			expectedStderr: "Manifest is invalid:\n- /modules/0/declarations/0/cssParts/0: missing properties: 'name'\n",
		},
		{
			fixture:        "invalid-css-custom-property",
			expectedStderr: "Manifest is invalid:\n- /modules/0/declarations/0/cssProperties/0: missing properties: 'name'\n",
		},
		{
			fixture:        "unsupported-schema",
			expectedStderr: "Schema version unsupported",
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
