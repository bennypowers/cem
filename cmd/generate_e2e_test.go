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
	t.Run("WithOutputFlag", func(t *testing.T) {
		projectDir := setupTest(t, "my-project")

		// Define paths
		srcFilePath := filepath.Join(projectDir, "my-element.js")
		outputFilePath := filepath.Join(projectDir, "dist", "custom-elements.json")

		// Execute the generate command
		stdout, _ := runCemCommand(t, projectDir, "generate", srcFilePath, "--output", outputFilePath)

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
		if !strings.Contains(pterm.RemoveColorFromString(stdout), expectedLog) {
			t.Fatalf("log output does not contain expected string.\nExpected: %s\nGot: %s", expectedLog, stdout)
		}
	})
	t.Run("InPackageDir", func(t *testing.T) {
		projectDir := setupTest(t, "package-flag")
		outputFilePath := filepath.Join(projectDir, "dist", "custom-elements.json")

		// Execute the generate command
		stdout, _ := runCemCommand(t, projectDir, "generate")

		// Check if the output file was created in the correct location
		if _, err := os.Stat(outputFilePath); os.IsNotExist(err) {
			t.Fatalf("output file was not created: %s", outputFilePath)
		}

		// Check the log output for the correct relative path
		expectedLog := "Wrote manifest to dist/custom-elements.json"
		if !strings.Contains(pterm.RemoveColorFromString(stdout), expectedLog) {
			t.Fatalf("log output does not contain expected string.\nExpected: %s\nGot: %s", expectedLog, stdout)
		}
	})
	t.Run("WithPackageFlag", func(t *testing.T) {
		projectDir := setupTest(t, "package-flag")
		outputFilePath := filepath.Join(projectDir, "dist", "custom-elements.json")

		// Execute the generate command
		stdout, _ := runCemCommand(t, ".", "generate", "--package", projectDir)

		// Check if the output file was created in the correct location
		if _, err := os.Stat(outputFilePath); os.IsNotExist(err) {
			t.Fatalf("output file was not created: %s", outputFilePath)
		}

		// Check the log output for the correct relative path
		expectedLog := "Wrote manifest to dist/custom-elements.json"
		if !strings.Contains(pterm.RemoveColorFromString(stdout), expectedLog) {
			t.Fatalf("log output does not contain expected string.\nExpected: %s\nGot: %s", expectedLog, stdout)
		}
	})
}
