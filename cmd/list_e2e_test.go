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
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/pterm/pterm"
)

func TestListE2E(t *testing.T) {
	testCases := []struct {
		name          string
		command       []string
		expectedError string
	}{
		{
			name:    "All",
			command: []string{"list"},
		},
		{
			name:    "Tags",
			command: []string{"list", "tags"},
		},
		{
			name:          "JSON",
			command:       []string{"list", "tags", "--format", "json"},
			expectedError: "Error: unknown format: json",
		},
		{
			name:          "Markdown",
			command:       []string{"list", "tags", "--format", "markdown"},
			expectedError: "Error: unknown format: markdown",
		},
		{
			name:    "Attributes",
			command: []string{"list", "attributes", "--tag-name", "test-elem"},
		},
		{
			name:    "CSS Custom Properties",
			command: []string{"list", "css-custom-properties", "--tag-name", "test-elem"},
		},
		{
			name:    "CSS Custom States",
			command: []string{"list", "css-custom-states", "--tag-name", "test-elem"},
		},
		{
			name:    "CSS Parts",
			command: []string{"list", "css-parts", "--tag-name", "test-elem"},
		},
		{
			name:    "Events",
			command: []string{"list", "events", "--tag-name", "test-elem"},
		},
		{
			name:    "Methods",
			command: []string{"list", "methods", "--tag-name", "test-elem"},
		},
		{
			name:    "Slots",
			command: []string{"list", "slots", "--tag-name", "test-elem"},
		},
		{
			name:    "Modules",
			command: []string{"list", "modules"},
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

				golden := filepath.Join(
					"goldens",
					fmt.Sprintf(
						"list.%s.golden.md",
						strings.ReplaceAll(strings.ToLower(tc.name), " ", "-"),
					),
				)
				expected, err := os.ReadFile(golden)
				if err != nil {
					t.Fatalf("failed to read golden file: %v", err)
				}

				cleanedStdout := pterm.RemoveColorFromString(stdout)
				cleanedExpected := strings.TrimSpace(string(expected))

				if !strings.Contains(cleanedStdout, cleanedExpected) {
					t.Fatalf("expected output to contain table:\n%s\n\nGot:\n%s", cleanedExpected, cleanedStdout)
				}
			}
		})
	}
}
