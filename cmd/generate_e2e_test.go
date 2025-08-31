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
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/nsf/jsondiff"
	"github.com/pterm/pterm"
)

var update = flag.Bool("update", false, "update golden files")

func TestGenerateE2E(t *testing.T) {
	testCases := []struct {
		name        string
		fixture     string
		command     []string
		outputFile  string
		expectedLog string
		workDir     string
	}{
		{
			name:        "WithOutputFlag",
			fixture:     "generate-project",
			command:     []string{"generate", "my-element.js", "--output", "dist/custom-elements.json"},
			outputFile:  "dist/custom-elements.json",
			expectedLog: "Wrote manifest to %s",
		},
		{
			name:        "InPackageDir",
			fixture:     "generate-project",
			command:     []string{"generate"},
			outputFile:  "dist/custom-elements.json",
			expectedLog: "Wrote manifest to dist/custom-elements.json",
		},
		{
			name:        "WithPackageFlag",
			fixture:     "generate-project",
			command:     []string{"generate", "--package"},
			outputFile:  "dist/custom-elements.json",
			expectedLog: "Wrote manifest to dist/custom-elements.json",
			workDir:     ".",
		},
		{
			name:        "WithExcludeFlag",
			fixture:     "generate-project",
			command:     []string{"generate", "--exclude", "my-element.js"},
			outputFile:  "dist/custom-elements.json",
			expectedLog: "Wrote manifest to dist/custom-elements.json",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			projectDir := setupTest(t, tc.fixture)
			outputFilePath := filepath.Join(projectDir, tc.outputFile)
			goldenPath := filepath.Join("goldens", tc.name+".json")

			command := tc.command
			if len(command) > 0 && command[len(command)-1] == "--package" {
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

			if *update {
				if err := os.WriteFile(goldenPath, content, 0644); err != nil {
					t.Fatalf("failed to write golden file: %v", err)
				}
			}

			expected, err := os.ReadFile(goldenPath)
			if err != nil {
				t.Fatalf("golden file missing: %s (have you run with -update?)\nerror: %v", goldenPath, err)
			}

			// Validate JSON
			var jsExpected, jsActual any
			if err := json.Unmarshal(expected, &jsExpected); err != nil {
				t.Fatalf("expected golden file is invalid JSON: %v", err)
			}
			if err := json.Unmarshal(content, &jsActual); err != nil {
				t.Fatalf("actual output is invalid JSON: %v\noutput:\n%s", err, content)
			}

			if !bytes.Equal(expected, content) {
				options := jsondiff.DefaultConsoleOptions()
				t.Error(jsondiff.Compare(expected, content, &options))
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
