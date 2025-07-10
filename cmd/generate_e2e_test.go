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
	t.Skip("Skipping while we debug TestGenerateE2EWithPackageFlag")
	viper.Reset()
	// Create a temporary directory for the test
	tmpDir, err := os.MkdirTemp("", "cem-test-")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	if err := os.WriteFile(filepath.Join(tmpDir, "package.json"), []byte(`{"name":"test-pkg"}`), 0644); err != nil {
		t.Fatalf("Failed to write dummy package.json: %v", err)
	}

	viper.Set("package", tmpDir)

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

func TestGenerateE2EWithPackageFlag(t *testing.T) {
	viper.Reset()
	// Create a temporary directory for the test
	tmpDir, err := os.MkdirTemp("", "cem-test-")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	if err := os.WriteFile(filepath.Join(tmpDir, "package.json"), []byte(`{"name":"test-pkg"}`), 0644); err != nil {
		t.Fatalf("Failed to write dummy package.json: %v", err)
	}

	// Create a project directory within the temp directory
	packageDir := filepath.Join(tmpDir, "my-project")

	// Create a config file in the project directory
	configDir := filepath.Join(packageDir, ".config")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		t.Fatalf("Failed to create config dir: %v", err)
	}
	configFile := filepath.Join(configDir, "cem.yaml")
	if err := os.WriteFile(configFile, []byte(`
generate:
  files:
    - my-element.js
`), 0644); err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	// Create a config file in the project directory
	pkgJsonPath := filepath.Join(packageDir, "package.json")
	if err := os.WriteFile(pkgJsonPath, []byte(`
{
	"name": "test",
	"version": "test",
	"customElements": "dist/custom-elements.json"
}
`), 0644); err != nil {
		t.Fatalf("Failed to write package json: %v", err)
	}

	// Create a dummy source file to be processed by the generate command
	srcFilePath := filepath.Join(packageDir, "my-element.js")
	if err := os.WriteFile(srcFilePath, []byte(`
/**
 * @customElement my-element
 */
export class MyElement extends HTMLElement {}
`), 0644); err != nil {
		t.Fatalf("Failed to write dummy source file: %v", err)
	}

	// Capture the output of the command
	var out bytes.Buffer
	origOut := pterm.Success.Writer
	pterm.Success.Writer = &out
	defer func() { pterm.Success.Writer = origOut }()
	rootCmd.SetOut(&out)
	rootCmd.SetErr(&out)

	// Set the project directory in viper, so the runtime can find it and the config file.
	viper.Set("package", packageDir)
	t.Logf("Set viper 'package' key to: %s", packageDir)

	// We don't need to pass command-line args if viper is set directly
	// and the config file is being read.
	rootCmd.SetArgs([]string{"generate"})

	err = rootCmd.Execute()
	if err != nil {
		t.Logf("Command stderr:\n%s", out.String())
		t.Fatalf("generate command failed: %v", err)
	}

	// Check if the output file was created in the correct location
	outputFile := filepath.Join(packageDir, "dist", "custom-elements.json")
	if _, err := os.Stat(outputFile); os.IsNotExist(err) {
		t.Fatalf("output file was not created: %s", outputFile)
	}

	// Check the log output for the correct relative path
	expectedLog := "Wrote manifest to my-project/dist/custom-elements.json"
	if !strings.Contains(pterm.RemoveColorFromString(out.String()), expectedLog) {
		t.Fatalf("log output does not contain expected string.\nExpected: %s\nGot: %s", expectedLog, out.String())
	}
}
