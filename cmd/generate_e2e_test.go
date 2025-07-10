package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/pterm/pterm"
	"github.com/spf13/viper"
)

func TestGenerateE2E(t *testing.T) {
	viper.Reset()
	// Create a temporary directory for the test
	tmpDir, err := os.MkdirTemp("", "cem-test-")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create a dummy source file to be processed by the generate command
	srcFilePath := filepath.Join(tmpDir, "my-element.js")
	srcFileContent := `
/**
 * @customElement my-element
 */
export class MyElement extends HTMLElement {}
`
	if err := os.WriteFile(srcFilePath, []byte(srcFileContent), 0644); err != nil {
		t.Fatalf("Failed to write dummy source file: %v", err)
	}

	// Define the output path in a subdirectory that doesn't exist yet
	outputDir := filepath.Join(tmpDir, "dist")
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		t.Fatalf("Failed to create output dir: %v", err)
	}
	outputFile := filepath.Join(outputDir, "custom-elements.json")

	// Capture the output of the command
	var out bytes.Buffer
	rootCmd.SetOut(&out)
	rootCmd.SetErr(&out)
	origOut := pterm.Success.Writer
	pterm.Success.Writer = &out
	defer func() { pterm.Success.Writer = origOut }()

	// Execute the generate command
	args := []string{"generate", srcFilePath, "-o", outputFile}
	rootCmd.SetArgs(args)
	// viper.BindPFlag("generate.output", generateCmd.Flags().Lookup("output"))
	err = rootCmd.Execute()
	if err != nil {
		t.Fatalf("generate command failed: %v", err)
	}

	// Check if the output file was created
	if _, err := os.Stat(outputFile); os.IsNotExist(err) {
		t.Fatalf("output file was not created: %s", outputFile)
	}

	// Check if the output file contains the expected content
	content, err := os.ReadFile(outputFile)
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}

	expected := `"tagName": "my-element"`
	if !bytes.Contains(content, []byte(expected)) {
		t.Fatalf("output file does not contain expected content.\nExpected: %s\nGot: %s", expected, content)
	}

	// Check the log output
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("could not get cwd: %v", err)
	}
	rel, err := filepath.Rel(cwd, outputFile)
	if err != nil {
		t.Fatalf("could not get relative path: %v", err)
	}
	expectedLog := "Wrote manifest to " + rel
	if !strings.Contains(pterm.RemoveColorFromString(out.String()), expectedLog) {
		t.Fatalf("log output does not contain expected string.\nExpected: %s\nGot: %s", expectedLog, out.String())
	}
}

func TestGenerateE2EWithProjectDir(t *testing.T) {
	viper.Reset()
	// Create a temporary directory for the test
	tmpDir, err := os.MkdirTemp("", "cem-test-")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create a project directory within the temp directory
	projectDir := filepath.Join(tmpDir, "my-project")
	configDir := filepath.Join(projectDir, ".config")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		t.Fatalf("Failed to create config dir: %v", err)
	}

	// Create a dummy source file to be processed by the generate command
	srcFilePath := filepath.Join(projectDir, "my-element.js")
	if err := os.WriteFile(srcFilePath, []byte(`
/**
 * @customElement my-element
 */
export class MyElement extends HTMLElement {}
`), 0644); err != nil {
		t.Fatalf("Failed to write dummy source file: %v", err)
	}

	// Create a config file in the project directory
	configFile := filepath.Join(configDir, "cem.yaml")
	if err := os.WriteFile(configFile, []byte(`
generate:
  files:
    - my-element.js
  output: dist/custom-elements.json
`), 0644); err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	// Capture the output of the command
	var out bytes.Buffer
	origOut := pterm.Success.Writer
	pterm.Success.Writer = &out
	defer func() { pterm.Success.Writer = origOut }()
	rootCmd.SetOut(&out)
	rootCmd.SetErr(&out)

	// Execute the generate command with the --package flag
	rootCmd.SetArgs([]string{"generate", "--package", filepath.Join("./my-project")})
	err = rootCmd.Execute()
	if err != nil {
		t.Fatalf("generate command failed: %v", err)
	}

	// Check if the output file was created in the correct location
	outputFile := filepath.Join(projectDir, "dist", "custom-elements.json")
	if _, err := os.Stat(outputFile); os.IsNotExist(err) {
		t.Fatalf("output file was not created: %s", outputFile)
	}

	// Check the log output for the correct relative path
	expectedLog := "Wrote manifest to my-project/dist/custom-elements.json"
	if !strings.Contains(pterm.RemoveColorFromString(out.String()), expectedLog) {
		t.Fatalf("log output does not contain expected string.\nExpected: %s\nGot: %s", expectedLog, out.String())
	}
}
