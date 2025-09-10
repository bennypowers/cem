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
package mcp_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/types"
	"bennypowers.dev/cem/workspace"
	"github.com/stretchr/testify/require"
)

// createTestWorkspace creates a temporary workspace for testing
func createTestWorkspace(t *testing.T, fixtureName string) (types.WorkspaceContext, func()) {
	t.Helper()

	// Create temporary directory
	tempDir, err := os.MkdirTemp("", "mcp-test-"+fixtureName+"-*")
	require.NoError(t, err, "Failed to create temp directory")

	// Copy fixture files to temp directory
	fixtureDir := filepath.Join("test-fixtures", fixtureName)
	copyFixtureFiles(t, fixtureDir, tempDir)

	// Create workspace context and initialize it
	workspace := workspace.NewFileSystemWorkspaceContext(tempDir)
	err = workspace.Init()
	require.NoError(t, err, "Failed to initialize workspace")

	cleanup := func() {
		_ = os.RemoveAll(tempDir)
	}

	return workspace, cleanup
}

// copyFixtureFiles copies files from the fixture directory to the target directory
func copyFixtureFiles(t *testing.T, fixtureDir, targetDir string) {
	t.Helper()

	if _, err := os.Stat(fixtureDir); os.IsNotExist(err) {
		// If fixture directory doesn't exist, create a basic manifest
		createBasicFixture(t, targetDir)
		return
	}

	err := filepath.Walk(fixtureDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Get relative path from fixture directory
		relPath, err := filepath.Rel(fixtureDir, path)
		if err != nil {
			return err
		}

		targetPath := filepath.Join(targetDir, relPath)

		if info.IsDir() {
			return os.MkdirAll(targetPath, info.Mode())
		}

		// Copy file
		srcContent, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		err = os.MkdirAll(filepath.Dir(targetPath), 0755)
		if err != nil {
			return err
		}

		return os.WriteFile(targetPath, srcContent, info.Mode())
	})

	require.NoError(t, err, "Failed to copy fixture files")
}

// createBasicFixture creates a basic manifest for testing
func createBasicFixture(t *testing.T, dir string) {
	t.Helper()

	manifest := map[string]interface{}{
		"schemaVersion": "2.1.1",
		"modules": []map[string]interface{}{
			{
				"kind": "javascript-module",
				"path": "test-element.js",
				"declarations": []map[string]interface{}{
					{
						"kind":          "class",
						"name":          "TestElement",
						"tagName":       "test-element",
						"customElement": true,
						"description":   "A test element for MCP tests",
						"attributes": []map[string]interface{}{
							{
								"name": "variant",
								"type": map[string]interface{}{
									"text": "\"primary\" | \"secondary\"",
								},
								"description": "Element variant",
								"default":     "\"primary\"",
							},
							{
								"name": "disabled",
								"type": map[string]interface{}{
									"text": "boolean",
								},
								"description": "Whether element is disabled",
							},
						},
						"slots": []map[string]interface{}{
							{
								"name":        "",
								"description": "Default slot content",
							},
							{
								"name":        "header",
								"description": "Header slot content",
							},
						},
						"events": []map[string]interface{}{
							{
								"name": "test-event",
								"type": map[string]interface{}{
									"text": "CustomEvent",
								},
								"description": "Test event",
							},
						},
						"cssProperties": []map[string]interface{}{
							{
								"name":        "--test-color",
								"syntax":      "<color>",
								"default":     "blue",
								"description": "Test color property",
							},
						},
						"cssParts": []map[string]interface{}{
							{
								"name":        "container",
								"description": "Main container part",
							},
						},
						"cssStates": []map[string]interface{}{
							{
								"name":        "active",
								"description": "Active state",
							},
						},
					},
				},
				"exports": []map[string]interface{}{
					{
						"kind": "custom-element-definition",
						"name": "test-element",
						"declaration": map[string]interface{}{
							"name":   "TestElement",
							"module": "test-element.js",
						},
					},
				},
			},
		},
	}

	manifestBytes, err := json.MarshalIndent(manifest, "", "  ")
	require.NoError(t, err, "Failed to marshal test manifest")

	err = os.WriteFile(filepath.Join(dir, "custom-elements.json"), manifestBytes, 0644)
	require.NoError(t, err, "Failed to write test manifest")

	// Create package.json that references the manifest
	packageJSON := map[string]interface{}{
		"name":           "test-package",
		"version":        "1.0.0",
		"customElements": "custom-elements.json",
	}

	packageBytes, err := json.MarshalIndent(packageJSON, "", "  ")
	require.NoError(t, err, "Failed to marshal package.json")

	err = os.WriteFile(filepath.Join(dir, "package.json"), packageBytes, 0644)
	require.NoError(t, err, "Failed to write package.json")
}
