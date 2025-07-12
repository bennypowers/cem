package cmd

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/pterm/pterm"
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

func TestGenerateE2E(t *testing.T) {
	fixtureName := "package-flag"
	// Create a temporary directory for the test
	tmpDir, err := os.MkdirTemp("", "cem-test-")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create a project directory within the temp directory
	projectDir := filepath.Join(tmpDir, fixtureName)
	err = os.CopyFS(projectDir, os.DirFS(filepath.Join(".", "fixture", fixtureName)))
	if err != nil {
		t.Fatalf("%s", err)
	}

	// Define paths
	srcFilePath := filepath.Join(projectDir, "my-element.js")
	outputFilePath := filepath.Join(projectDir, "dist", "custom-elements.json")

	// Execute the generate command
	cmd := exec.Command(
		cemBinary,
		"generate",
		srcFilePath,
		"-o",
		outputFilePath,
	)
	cmd.Env = append(os.Environ(), "GOCOVERDIR="+coverDir)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err = cmd.Run()
	if err != nil {
		t.Fatalf("generate command failed: %v\nOutput: %s", err, out.String())
	}

	// Check if the output file was created
	if _, err := os.Stat(outputFilePath); os.IsNotExist(err) {
		t.Fatalf("output file was not created: %s", outputFilePath)
	}

	// Check if the output file contains the expected content
	content, err := os.ReadFile(outputFilePath)
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}

	expected := `"tagName": "my-element"`
	if !bytes.Contains(content, []byte(expected)) {
		t.Fatalf("output file does not contain expected content.\nExpected: %s\nGot: %s", expected, content)
	}

	// Check the log output for the correct relative path
	expectedLog := fmt.Sprintf("Wrote manifest to %s", outputFilePath)
	actualLog := out.String()
	t.Log(actualLog)
	if !strings.Contains(pterm.RemoveColorFromString(actualLog), expectedLog) {
		t.Fatalf("log output does not contain expected string.\nExpected: %s\nGot: %s", expectedLog, actualLog)
	}
}

func TestGenerateE2EWithPackageFlag(t *testing.T) {
	fixtureName := "package-flag"
	// Create a temporary directory for the test
	tmpDir, err := os.MkdirTemp("", "cem-test-")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create a project directory within the temp directory
	err = os.CopyFS(tmpDir, os.DirFS(filepath.Clean("./fixture")))
	if err != nil {
		t.Fatalf("%s", err)
	}

	packageDir := filepath.Join(tmpDir, fixtureName)
	outputFilePath := filepath.Join(packageDir, "dist", "custom-elements.json")

	// Execute the generate command
	cmd := exec.Command(cemBinary, "generate", "--package", packageDir)
	cmd.Env = append(os.Environ(), "GOCOVERDIR="+coverDir)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err = cmd.Run()
	if err != nil {
		t.Fatalf("generate command failed: %v\nOutput: %s", err, out.String())
	}

	// Check if the output file was created in the correct location
	if _, err := os.Stat(outputFilePath); os.IsNotExist(err) {
		t.Fatalf("output file was not created: %s", outputFilePath)
	}

	// Check the log output for the correct relative path
	expectedLog := "Wrote manifest to dist/custom-elements.json"
	actualLog := out.String()
	t.Log(actualLog)
	if !strings.Contains(pterm.RemoveColorFromString(actualLog), expectedLog) {
		t.Fatalf("log output does not contain expected string.\nExpected: %s\nGot: %s", expectedLog, actualLog)
	}
}
