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
			expectedError: "Error: requires a pattern argument (or use -i for interactive mode)",
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
		"requires a terminal",
	}

	for _, expected := range expectedStrings {
		if !strings.Contains(stdout, expected) {
			t.Errorf("Expected help output to contain %q, got:\n%s", expected, stdout)
		}
	}
}

func TestSearchInteractive_RequiresTerminal(t *testing.T) {
	// Subprocess has no TTY so term.IsTerminal returns false; the TTY guard
	// should fire before any manifest loading or TUI launch.
	projectDir := setupTest(t, "list-project")
	_, stderr := runCemCommand(t, projectDir, "search", "-i")
	if !strings.Contains(stderr, "interactive mode requires a terminal") {
		t.Errorf("expected TTY guard error, got: %s", stderr)
	}
}

func TestSearchInteractive_MutuallyExclusiveWithFormat(t *testing.T) {
	projectDir := setupTest(t, "list-project")
	_, stderr := runCemCommand(t, projectDir, "search", "-i", "--format", "tree", "button")
	if !strings.Contains(stderr, "if any flags in the group [interactive format] are set") {
		t.Errorf("expected mutual-exclusion error, got: %s", stderr)
	}
}
