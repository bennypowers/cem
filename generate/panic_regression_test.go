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

// TestGenerateConfigFileNotFoundNoPanic tests the exact scenario that originally caused a panic.
// This is a regression test for the nil pointer dereference that occurred when using --config
// with a file path that caused path resolution issues.
//
// Original panic:
// panic: runtime error: invalid memory address or nil pointer dereference
// [signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x9eeb8e]
// goroutine 1 [running]:
// bennypowers.dev/cem/cmd.init.func1(0x147b3c0, {0xc000325f50, 0x1, 0x3})
//
//	/home/bennyp/Developer/cem/cmd/generate.go:150 +0xc0e
func TestGenerateConfigFileNotFoundNoPanic(t *testing.T) {
	// Change to the test fixture directory to test the exact problematic scenario
	fixtureDir := filepath.Join("test", "fixtures", "project-source-hrefs")

	// Use go run from the project root but run the command in the fixture directory
	// This simulates the exact path resolution issue that caused the original panic
	cmd := exec.Command("go", "run", "../../../../main.go", "generate", "--config", ".config/cem.yaml", "src/source-hrefs.ts")
	cmd.Dir = fixtureDir

	cmdOutput, err := cmd.CombinedOutput()

	// Should exit with error (not panic)
	if err == nil {
		t.Error("Expected command to fail due to file path resolution, but it succeeded")
	}

	outputStr := string(cmdOutput)

	// The main check: Should NOT contain panic trace
	// This is what we're testing for - the original issue was a panic
	hasPanic := strings.Contains(outputStr, "panic:") ||
		strings.Contains(outputStr, "runtime error:") ||
		strings.Contains(outputStr, "nil pointer dereference")

	if hasPanic {
		t.Errorf("REGRESSION: Command panicked (this was the original bug that was fixed). Output:\n%s", outputStr)
	} else {
		// This is the expected behavior after the fix
		t.Logf("✓ Command correctly failed with clean error (no panic)")
	}

	// Should contain proper error message about file resolution
	if !strings.Contains(outputStr, "no such file or directory") {
		t.Errorf("Expected file resolution error, got: %s", outputStr)
	}

	// Should contain the incorrectly resolved path that causes the issue
	if !strings.Contains(outputStr, ".config/src/source-hrefs.ts") {
		t.Errorf("Expected to see the problematic path resolution, got: %s", outputStr)
	}
}

// TestGenerateSuccessCodePathNoPanic tests that the success code path doesn't panic
// when writing output files. This tests the specific line that was panicking in
// cmd/generate.go:150 (pterm.Success.Printf)
func TestGenerateSuccessCodePathNoPanic(t *testing.T) {
	// Test successful generation with output file (the code path that was panicking)
	fixtureDir := filepath.Join("test", "fixtures", "project-source-hrefs")
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
