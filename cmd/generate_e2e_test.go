package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func resetFlags(cmd *cobra.Command) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		f.Value.Set(f.DefValue)
	})
	for _, c := range cmd.Commands() {
		resetFlags(c)
	}
}

func resetCobraAndViper(t *testing.T) {
	t.Helper()
	rootCmd.Flags().VisitAll(func(f *pflag.Flag) {
		f.Value.Set(f.DefValue) // Reset flag to default (if possible)
	})

	// Optionally reset persistent flags if you use them
	resetFlags(rootCmd)

	// If you set up PersistentPreRunE, consider resetting side effects (like global context)
	// e.g., ctx = nil
	initialCWD = ""
	generateFiles = nil
	start = time.Time{}

	viper.Reset() // Reset config state
}

func TestGenerateE2E(t *testing.T) {
	resetCobraAndViper(t)
	// Create a temporary directory for the test
	tmpDir, err := os.MkdirTemp("", "cem-test-")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create a project directory within the temp directory
	err = os.CopyFS(tmpDir, os.DirFS(filepath.Clean("./fixture/package-flag")))
	if err != nil {
		t.Fatalf("%s", err)
	}

	viper.Set("package", tmpDir)

	// Create a dummy source file to be processed by the generate command
	srcFilePath := filepath.Join(tmpDir, "my-element.js")
	outputDir := filepath.Join(tmpDir, "dist")
	outputFilePath := filepath.Join(outputDir, "custom-elements.json")

	// Capture the output of the command
	var out bytes.Buffer
	rootCmd.SetOut(&out)
	rootCmd.SetErr(&out)
	origOut := pterm.Success.Writer
	pterm.Success.Writer = &out
	defer func() { pterm.Success.Writer = origOut }()

	// Execute the generate command
	args := []string{"generate", srcFilePath, "-o", outputFilePath}
	rootCmd.SetArgs(args)
	err = rootCmd.Execute()
	if err != nil {
		t.Fatalf("generate command failed: %v", err)
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
	expectedLog := "Wrote manifest to dist/custom-elements.json"
	actualLog := out.String()
	t.Log(actualLog)
	if !strings.Contains(pterm.RemoveColorFromString(actualLog), expectedLog) {
		t.Fatalf("log output does not contain expected string.\nExpected: %s\nGot: %s", expectedLog, actualLog)
	}
}

func TestGenerateE2EWithPackageFlag(t *testing.T) {
	resetCobraAndViper(t)
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

	packageDir := filepath.Join(tmpDir, "package-flag")
	outputFilePath := filepath.Join(packageDir, "dist", "custom-elements.json")
	viper.Set("package", packageDir)

	// Set the project directory in viper, so the runtime can find it and the config file.
	t.Logf("Set viper 'package' key to: %s", viper.GetString("package"))

	// Capture the output of the command
	var out bytes.Buffer
	origOut := pterm.Success.Writer
	pterm.Success.Writer = &out
	defer func() { pterm.Success.Writer = origOut }()
	rootCmd.SetOut(&out)
	rootCmd.SetErr(&out)

	// We don't need to pass command-line args if viper is set directly
	// and the config file is being read.
	rootCmd.SetArgs([]string{"generate"})

	err = rootCmd.Execute()
	if err != nil {
		t.Fatalf("generate command failed: %v", err)
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
