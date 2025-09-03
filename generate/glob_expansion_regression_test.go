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
package generate_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"testing/synctest"

	G "bennypowers.dev/cem/generate"
	M "bennypowers.dev/cem/manifest"
	W "bennypowers.dev/cem/workspace"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestGlobPatternExpansion tests the regression fix for glob pattern expansion
// where glob patterns like "elements/*/rh-*.ts" were treated as literal file paths
// instead of being expanded to actual matching files.
func TestGlobPatternExpansion(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
	// Create a temporary workspace
	tempDir := t.TempDir()

	// Create directory structure matching the pattern elements/*/rh-*.ts
	elementsDir := filepath.Join(tempDir, "elements", "rh-alert")
	err := os.MkdirAll(elementsDir, 0755)
	require.NoError(t, err)

	// Create a TypeScript file that matches the glob pattern
	rhAlertFile := filepath.Join(elementsDir, "rh-alert.ts")
	rhAlertContent := `import {LitElement} from 'lit';
import {customElement, property} from 'lit/decorators.js';

@customElement('rh-alert')
export class RhAlert extends LitElement {
  @property() state: 'info' | 'warning' | 'error' = 'info';
}`
	err = os.WriteFile(rhAlertFile, []byte(rhAlertContent), 0644)
	require.NoError(t, err)

	// Create a package.json with customElements reference
	packageJSON := `{
  "name": "test-package",
  "customElements": "dist/custom-elements.json"
}`
	err = os.WriteFile(filepath.Join(tempDir, "package.json"), []byte(packageJSON), 0644)
	require.NoError(t, err)

	// Create .config directory and cem.yaml with glob pattern
	configDir := filepath.Join(tempDir, ".config")
	err = os.MkdirAll(configDir, 0755)
	require.NoError(t, err)

	cemConfig := `generate:
  files:
    - elements/*/rh-*.ts  # This is the glob pattern that was failing
`
	err = os.WriteFile(filepath.Join(configDir, "cem.yaml"), []byte(cemConfig), 0644)
	require.NoError(t, err)

	// Create workspace context
	workspace := W.NewFileSystemWorkspaceContext(tempDir)
	err = workspace.Init()
	require.NoError(t, err)

	// Create generate session and test that glob patterns are expanded correctly
	session, err := G.NewGenerateSession(workspace)
	require.NoError(t, err)
	defer session.Close()

	// Generate the manifest - this should not fail with "no such file or directory"
	// for the glob pattern "elements/*/rh-*.ts"
	pkg, err := session.GenerateFullManifest(context.Background())
	require.NoError(t, err, "Glob pattern expansion should work without 'no such file or directory' errors")
	require.NotNil(t, pkg, "Generated package should not be nil")

	// Verify that the manifest contains the expected element
	assert.Len(t, pkg.Modules, 1, "Should have exactly one module")

	module := pkg.Modules[0]
	assert.Contains(t, module.Path, "elements/rh-alert/rh-alert.js", "Module path should contain the expected pattern")
	assert.Len(t, module.Declarations, 1, "Should have exactly one declaration")

	// Verify the custom element was detected
	customElementDecl, ok := module.Declarations[0].(*M.CustomElementDeclaration)
	require.True(t, ok, "Declaration should be a CustomElementDeclaration")
	assert.Equal(t, "rh-alert", customElementDecl.CustomElement.TagName, "Tag name should be correct")

	// Verify the attribute was detected
	assert.Len(t, customElementDecl.CustomElement.Attributes, 1, "Should have one attribute")
	attr := customElementDecl.CustomElement.Attributes[0]
	assert.Equal(t, "state", attr.Name, "Attribute name should be 'state'")
	assert.Equal(t, "'info' | 'warning' | 'error'", attr.Type.Text, "Attribute type should be the union type")
	}) // End synctest.Test
}

// TestGlobPatternWithWatchSession tests that glob patterns work correctly
// with watch sessions, ensuring the fix works for both generate and watch modes.
func TestGlobPatternWithWatchSession(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
	// Create a temporary workspace
	tempDir := t.TempDir()

	// Create directory structure
	elementsDir := filepath.Join(tempDir, "elements", "my-element")
	err := os.MkdirAll(elementsDir, 0755)
	require.NoError(t, err)

	// Create a TypeScript file
	tsFile := filepath.Join(elementsDir, "my-element.ts")
	tsContent := `import {LitElement} from 'lit';
import {customElement} from 'lit/decorators.js';

@customElement('my-element')
export class MyElement extends LitElement {}`
	err = os.WriteFile(tsFile, []byte(tsContent), 0644)
	require.NoError(t, err)

	// Create config with glob pattern
	configDir := filepath.Join(tempDir, ".config")
	err = os.MkdirAll(configDir, 0755)
	require.NoError(t, err)

	cemConfig := `generate:
  files:
    - elements/*/my-*.ts
`
	err = os.WriteFile(filepath.Join(configDir, "cem.yaml"), []byte(cemConfig), 0644)
	require.NoError(t, err)

	// Create workspace context
	workspace := W.NewFileSystemWorkspaceContext(tempDir)
	err = workspace.Init()
	require.NoError(t, err)

	// Test that watch session can be created without "no such file or directory" errors
	session, err := G.NewWatchSession(workspace, []string{"elements/*/my-*.ts"})
	require.NoError(t, err, "Watch session creation should work with glob patterns")
	require.NotNil(t, session, "Watch session should not be nil")

	// Clean up
	session.Close()
	}) // End synctest.Test
}
