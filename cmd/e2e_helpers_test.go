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
	"fmt"

	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

// These tests are end-to-end tests that build the `cem` binary and run it as a
// subprocess, simulating real-world usage from a shell. This provides higher
// confidence than calling the cobra command's Execute() method directly.
//
// To capture code coverage from these tests, we build the binary with the
// `-cover` flag. When the tests run the binary, they set the `GOCOVERDIR`
// environment variable. The instrumented binary writes coverage data to this
// directory upon exit. Finally, the `TestMain` function merges the data from
// all test runs into a single `coverage.e2e.out` file.

var cemBinary string
var coverDir string

func TestMain(m *testing.M) {
	// Create a temporary directory for the test binary
	binaryDir, err := os.MkdirTemp("", "cem-test-binary-")
	if err != nil {
		panic("Failed to create temp dir for binary: " + err.Error())
	}
	defer os.RemoveAll(binaryDir)

	// Create a temporary directory for the coverage data
	coverDir, err = os.MkdirTemp("", "cem-test-cover-")
	if err != nil {
		panic("Failed to create temp dir for coverage: " + err.Error())
	}
	defer os.RemoveAll(coverDir)

	// Build the binary with coverage instrumentation from the project root
	buildCmd := exec.Command(
		"go",
		"build",
		"-o",
		binaryDir,
		"-cover",
		"-coverpkg",
		"./...",
		"./...",
	)
	buildCmd.Dir = ".."
	buildCmd.Stderr = os.Stderr
	if err := buildCmd.Run(); err != nil {
		panic("Failed to build cem binary: " + err.Error())
	}

	cemBinary = filepath.Join(binaryDir, "cem")

	// Run the tests
	code := m.Run()

	// Merge coverage data from the E2E tests into binary format.
	// This directory is placed in the `cmd` directory where `go test` is run.
	// The Makefile will later merge this with unit test coverage.
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	e2eCoverDir := filepath.Join(cwd, "coverage.e2e")

	// Remove old coverage data if it exists, then create fresh directory
	os.RemoveAll(e2eCoverDir)
	if err := os.MkdirAll(e2eCoverDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create E2E coverage directory: %v\n", err)
	}

	// Merge coverage from all E2E test runs into binary format
	mergeCmd := exec.Command("go", "tool", "covdata", "merge", "-i", coverDir, "-o", e2eCoverDir)
	mergeCmd.Stdout = os.Stdout
	mergeCmd.Stderr = os.Stderr
	if err := mergeCmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to merge E2E coverage data: %v\n", err)
	}

	// Note: We keep the binary format directory (coverage.e2e/) for the Makefile to merge
	// with unit test coverage. The Makefile's 'coverage' target will handle final conversion.

	os.Exit(code)
}

func setupTest(t *testing.T, fixtureName string) (projectDir string) {
	t.Helper()
	// Create a temporary directory for the test
	tmpDir, err := os.MkdirTemp("", "cem-test-")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	t.Cleanup(func() { os.RemoveAll(tmpDir) })

	// Create a project directory within the temp directory
	projectDir = filepath.Join(tmpDir, fixtureName)
	err = os.CopyFS(projectDir, os.DirFS(filepath.Join("testdata", "fixtures", fixtureName)))
	if err != nil {
		t.Fatalf("%s", err)
	}

	return projectDir
}

func runCemCommand(t *testing.T, workDir string, args ...string) (stdout, stderr string) {
	t.Helper()
	cmd := exec.Command(cemBinary, args...)
	cmd.Dir = workDir
	cmd.Env = append(os.Environ(), "GOCOVERDIR="+coverDir)
	var out, errOut bytes.Buffer
	t.Logf("coverDir=%q\n", coverDir)
	cmd.Stdout = &out
	cmd.Stderr = &errOut

	err := cmd.Run()
	if err != nil {
		t.Logf("cem command failed: %v\n", err)
	}

	return out.String(), errOut.String()
}
