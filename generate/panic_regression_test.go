/*
Copyright © 2025 Benny Powers

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
package generate

import (
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// TestGenerateConfigRelativePathNoPanic tests that --config with a relative path
// correctly derives the project root and does not panic.
//
// Original bug: --config .config/cem.yaml used filepath.Dir() which gave ".config/"
// as the project root, causing files to resolve to ".config/src/source-hrefs.ts"
// and either panicking or failing with a confusing error.
//
// After the unified config fix, --config .config/cem.yaml correctly identifies
// the parent of .config/ as the project root, and the command succeeds.
func TestGenerateConfigRelativePathNoPanic(t *testing.T) {
	fixtureDir := filepath.Join("testdata", "fixtures", "project-source-hrefs")

	cmd := exec.Command("go", "run", "../../../../main.go", "generate", "--config", ".config/cem.yaml", "src/source-hrefs.ts")
	cmd.Dir = fixtureDir

	cmdOutput, err := cmd.CombinedOutput()
	outputStr := string(cmdOutput)

	// Should NOT contain panic trace
	hasPanic := strings.Contains(outputStr, "panic:") ||
		strings.Contains(outputStr, "runtime error:") ||
		strings.Contains(outputStr, "nil pointer dereference")

	if hasPanic {
		t.Errorf("REGRESSION: Command panicked. Output:\n%s", outputStr)
	}

	// With the unified config fix, --config .config/cem.yaml now correctly
	// derives the project root as the parent of .config/, so the command succeeds.
	if err != nil {
		t.Errorf("Command should succeed with --config .config/cem.yaml, but failed: %v\nOutput: %s", err, outputStr)
	}
}

// TestGenerateConfigMissingFileNoPanic tests that --config with a nonexistent
// path fails with a clean error, not a panic.
func TestGenerateConfigMissingFileNoPanic(t *testing.T) {
	fixtureDir := filepath.Join("testdata", "fixtures", "project-source-hrefs")

	cmd := exec.Command("go", "run", "../../../../main.go", "generate", "--config", "./nonexistent/cem.yaml", "src/source-hrefs.ts")
	cmd.Dir = fixtureDir

	cmdOutput, err := cmd.CombinedOutput()
	outputStr := string(cmdOutput)

	hasPanic := strings.Contains(outputStr, "panic:") ||
		strings.Contains(outputStr, "runtime error:")

	if hasPanic {
		t.Errorf("Command panicked with missing config. Output:\n%s", outputStr)
	}

	if err == nil {
		t.Error("Expected command to fail with missing config file, but it succeeded")
	}
}

// TestGenerateSuccessCodePathNoPanic tests that the success code path doesn't panic
// when writing output files. This tests the specific line that was panicking in
// cmd/generate.go:150 (pterm.Success.Printf)
func TestGenerateSuccessCodePathNoPanic(t *testing.T) {
	// Test successful generation with output file (the code path that was panicking)
	fixtureDir := filepath.Join("testdata", "fixtures", "project-source-hrefs")
	outputFile := filepath.Join(t.TempDir(), "test-manifest.json")

	// Use go run to test the success code path without hardcoded paths
	cmd := exec.Command("go", "run", "../../../../main.go", "generate",
		"src/source-hrefs.ts",
		"--source-control-root-url", "https://github.com/example/repo/tree/main/",
		"--output", outputFile)
	cmd.Dir = fixtureDir

	successOutput, err := cmd.CombinedOutput()

	// Should succeed without panic
	if err != nil {
		t.Errorf("Command should succeed, but failed with: %v\nOutput: %s", err, successOutput)
	}

	outputStr := string(successOutput)

	// Should NOT contain panic trace
	if strings.Contains(outputStr, "panic:") {
		t.Errorf("Command panicked during success path. Output:\n%s", outputStr)
	}

	// Should contain success message (the line that was panicking)
	if !strings.Contains(outputStr, "SUCCESS") || !strings.Contains(outputStr, "Wrote manifest") {
		t.Errorf("Expected success message from the previously panicking code path, got: %s", outputStr)
	}

	t.Logf("Success code path works correctly: %s", strings.TrimSpace(outputStr))
}
