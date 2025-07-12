//go:build e2e

package cmd_test

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/pterm/pterm"
)

func TestGenerateE2E(t *testing.T) {
	testCases := []struct {
		name            string
		fixture         string
		command         []string
		outputFile      string
		expectedLog     string
		expectedContent string
		workDir         string
	}{
		{
			name:            "WithOutputFlag",
			fixture:         "generate-project",
			command:         []string{"generate", "my-element.js", "--output", "dist/custom-elements.json"},
			outputFile:      "dist/custom-elements.json",
			expectedLog:     "Wrote manifest to %s",
			expectedContent: `"tagName": "my-element"`,
		},
		{
			name:            "InPackageDir",
			fixture:         "generate-project",
			command:         []string{"generate"},
			outputFile:      "dist/custom-elements.json",
			expectedLog:     "Wrote manifest to dist/custom-elements.json",
			expectedContent: `"tagName": "my-element"`,
		},
		{
			name:            "WithPackageFlag",
			fixture:         "generate-project",
			command:         []string{"generate", "--package"},
			outputFile:      "dist/custom-elements.json",
			expectedLog:     "Wrote manifest to dist/custom-elements.json",
			expectedContent: `"tagName": "my-element"`,
			workDir:         ".",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			projectDir := setupTest(t, tc.fixture)
			outputFilePath := filepath.Join(projectDir, tc.outputFile)

			command := tc.command
			if tc.name == "WithPackageFlag" {
				command = append(command, projectDir)
			}

			workDir := projectDir
			if tc.workDir != "" {
				workDir = tc.workDir
			}

			stdout, _ := runCemCommand(t, workDir, command...)

			if _, err := os.Stat(outputFilePath); os.IsNotExist(err) {
				t.Fatalf("output file was not created: %s", outputFilePath)
			}

			content, err := os.ReadFile(outputFilePath)
			if err != nil {
				t.Fatalf("Failed to read output file: %v", err)
			}

			if !bytes.Contains(content, []byte(tc.expectedContent)) {
				t.Fatalf("output file does not contain expected content.\nExpected: %s\nGot: %s", tc.expectedContent, content)
			}

			expectedLog := tc.expectedLog
			if tc.name == "WithOutputFlag" {
				expectedLog = fmt.Sprintf(expectedLog, tc.outputFile)
			}

			if !strings.Contains(pterm.RemoveColorFromString(stdout), expectedLog) {
				t.Fatalf("log output does not contain expected string.\nExpected: %s\nGot: %s", expectedLog, stdout)
			}
		})
	}
}

