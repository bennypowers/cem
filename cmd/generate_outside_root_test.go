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
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGenerateOutsideProjectRoot(t *testing.T) {
	// Setup: Create a project in a temporary directory
	tmpDir, err := os.MkdirTemp("", "cem-test-")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Copy the design token fixture to the project directory
	projectDir := filepath.Join(tmpDir, "test-project")
	fixtureDir := "fixtures/project-design-tokens"

	// Create the project directory first
	err = os.MkdirAll(projectDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create project dir: %v", err)
	}

	err = os.CopyFS(projectDir, os.DirFS(fixtureDir))
	if err != nil {
		t.Fatalf("Failed to copy fixture: %v", err)
	}

	// Create a separate working directory outside the project
	workDir := filepath.Join(tmpDir, "outside-work-dir")
	err = os.MkdirAll(workDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create work dir: %v", err)
	}

	// Run cem generate from outside the project root using absolute paths
	designTokensPath := filepath.Join(projectDir, "design-tokens.json")
	sourceFilePath := filepath.Join(projectDir, "src/component.ts")
	outputPath := filepath.Join(workDir, "custom-elements.json")

	stdout, stderr := runCemCommand(t, workDir,
		"generate",
		"--verbose",
		"--output", outputPath,
		"--design-tokens-prefix", "cem",
		"--design-tokens", designTokensPath,
		sourceFilePath,
	)

	t.Logf("stdout: %s", stdout)
	t.Logf("stderr: %s", stderr)

	// Verify output file was created
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		t.Fatalf("output file was not created: %s", outputPath)
	}

	// Read and parse the output
	content, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}

	var manifest map[string]interface{}
	if err := json.Unmarshal(content, &manifest); err != nil {
		t.Fatalf("Failed to parse output JSON: %v", err)
	}

	// Check that module paths are relative, not absolute
	modules, ok := manifest["modules"].([]interface{})
	if !ok || len(modules) == 0 {
		t.Fatalf("No modules found in manifest")
	}

	module, ok := modules[0].(map[string]interface{})
	if !ok {
		t.Fatalf("Invalid module structure")
	}

	modulePath, ok := module["path"].(string)
	if !ok {
		t.Fatalf("Module path not found")
	}

	// The path should be relative to the project root, not absolute
	if filepath.IsAbs(modulePath) {
		t.Errorf("Expected relative module path, got absolute: %s", modulePath)
	}

	// Check that design token descriptions are merged
	manifestStr := string(content)
	if !strings.Contains(manifestStr, "DESIGN TOKEN") {
		t.Errorf("Expected design token descriptions to be merged into manifest, but 'DESIGN TOKEN' not found in output")
	}
}

func TestGenerateOutsideProjectRootRelativePaths(t *testing.T) {
	// Setup: Create a project in a temporary directory
	tmpDir, err := os.MkdirTemp("", "cem-test-")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Copy the design token fixture to the project directory
	projectDir := filepath.Join(tmpDir, "test-project")
	fixtureDir := "fixtures/project-design-tokens"

	// Create the project directory first
	err = os.MkdirAll(projectDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create project dir: %v", err)
	}

	err = os.CopyFS(projectDir, os.DirFS(fixtureDir))
	if err != nil {
		t.Fatalf("Failed to copy fixture: %v", err)
	}

	// Create a separate working directory outside the project
	workDir := filepath.Join(tmpDir, "outside-work-dir")
	err = os.MkdirAll(workDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create work dir: %v", err)
	}

	// Run cem generate from outside the project root using relative paths
	// This tests the scenario where user provides relative paths that need to be resolved
	outputPath := filepath.Join(workDir, "custom-elements.json")

	stdout, stderr := runCemCommand(t, workDir,
		"generate",
		"--verbose",
		"--output", outputPath,
		"--design-tokens-prefix", "cem",
		"--design-tokens", "../test-project/design-tokens.json",
		"../test-project/src/component.ts",
	)

	t.Logf("stdout: %s", stdout)
	t.Logf("stderr: %s", stderr)

	// Verify output file was created
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		t.Fatalf("output file was not created: %s", outputPath)
	}

	// Read and parse the output
	content, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}

	var manifest map[string]interface{}
	if err := json.Unmarshal(content, &manifest); err != nil {
		t.Fatalf("Failed to parse output JSON: %v", err)
	}

	// Check that module paths are relative, not absolute
	modules, ok := manifest["modules"].([]interface{})
	if !ok || len(modules) == 0 {
		t.Fatalf("No modules found in manifest")
	}

	module, ok := modules[0].(map[string]interface{})
	if !ok {
		t.Fatalf("Invalid module structure")
	}

	modulePath, ok := module["path"].(string)
	if !ok {
		t.Fatalf("Module path not found")
	}

	// The path should be relative to the project root, not absolute
	if filepath.IsAbs(modulePath) {
		t.Errorf("Expected relative module path, got absolute: %s", modulePath)
	}

	// For relative paths, the module path should still be relative to the detected project root
	// It should NOT include the "../test-project/" part since that was just how we got there
	if strings.Contains(modulePath, "../test-project/") {
		t.Errorf("Expected module path relative to project root, but got path relative to working directory: %s", modulePath)
	}

	// Check that design token descriptions are merged
	manifestStr := string(content)
	if !strings.Contains(manifestStr, "DESIGN TOKEN") {
		t.Errorf("Expected design token descriptions to be merged into manifest, but 'DESIGN TOKEN' not found in output")
	}
}
