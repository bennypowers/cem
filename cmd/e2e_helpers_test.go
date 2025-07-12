//go:build e2e

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

	cemBinary = filepath.Join(binaryDir, "cem")

	// Build the binary with coverage instrumentation from the project root
	buildCmd := exec.Command("go", "build", "-o", cemBinary, "-cover", "..")
	buildCmd.Stderr = os.Stderr
	if err := buildCmd.Run(); err != nil {
		panic("Failed to build cem binary: " + err.Error())
	}

	// Run the tests
	code := m.Run()

	// Merge coverage data from the E2E tests into a single profile.
	// This file is placed in the `cmd` directory where `go test` is run.
	mergeCmd := exec.Command("go", "tool", "covdata", "merge", "-i", coverDir, "-o", "coverage.e2e.out")
	mergeCmd.Stderr = os.Stderr
	if err := mergeCmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to merge coverage data: %v\n", err)
	}

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
	err = os.CopyFS(projectDir, os.DirFS(filepath.Join(".", "fixture", fixtureName)))
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
	cmd.Stdout = &out
	cmd.Stderr = &errOut

	err := cmd.Run()
	if err != nil {
		t.Logf("cem command failed: %v", err)
	}

	return out.String(), errOut.String()
}
