//go:build e2e

package cmd_test

import (
	"strings"
	"testing"
)

func TestSearchE2E(t *testing.T) {
	testCases := []struct {
		name           string
		command        []string
		expectedError  string
		expectContains string
	}{
		{
			name:           "Basic search - test",
			command:        []string{"search", "test"},
			expectContains: "test-elem",
		},
		{
			name:           "Search with table format",
			command:        []string{"search", "test", "--format", "table"},
			expectContains: "test-elem",
		},
		{
			name:           "Search with tree format",
			command:        []string{"search", "test", "--format", "tree"},
			expectContains: "Search Results for: test",
		},
		{
			name:          "Empty search pattern",
			command:       []string{"search", ""},
			expectedError: "Error: search pattern cannot be empty",
		},
		{
			name:          "Missing search pattern",
			command:       []string{"search"},
			expectedError: "Error: accepts 1 arg(s), received 0",
		},
		{
			name:          "Invalid format",
			command:       []string{"search", "test", "--format", "json"},
			expectedError: "Error: unknown format: json",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			projectDir := setupTest(t, "list-project")
			stdout, stderr := runCemCommand(t, projectDir, tc.command...)

			if tc.expectedError != "" {
				if !strings.Contains(stderr, tc.expectedError) {
					t.Fatalf("expected stderr to contain '%s', got:\n%s", tc.expectedError, stderr)
				}
			} else {
				if stderr != "" {
					t.Fatal(stderr)
				}

				if tc.expectContains != "" {
					if !strings.Contains(stdout, tc.expectContains) {
						t.Errorf("Expected output to contain %q, got:\n%s", tc.expectContains, stdout)
					}
				}
			}
		})
	}
}

func TestSearchHelp(t *testing.T) {
	projectDir := setupTest(t, "list-project")
	stdout, stderr := runCemCommand(t, projectDir, "search", "--help")

	if stderr != "" {
		t.Fatal(stderr)
	}

	expectedStrings := []string{
		"Search through the custom elements manifest for any element matching the given pattern",
		"The search pattern is treated as a regular expression by default",
		"Examples:",
		"cem search button",
		"cem search \"^my-.*button$\"",
	}

	for _, expected := range expectedStrings {
		if !strings.Contains(stdout, expected) {
			t.Errorf("Expected help output to contain %q, got:\n%s", expected, stdout)
		}
	}
}
