//go:build e2e

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
