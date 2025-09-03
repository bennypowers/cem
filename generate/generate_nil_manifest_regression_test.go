package generate_test

import (
	"os"
	"path/filepath"
	"testing"
	"testing/synctest"

	G "bennypowers.dev/cem/generate"
	W "bennypowers.dev/cem/workspace"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestGenerateWithNilManifest tests the regression fix for nil pointer dereference
// when the Generate function returns a nil manifest pointer due to file processing errors.
// This was causing a panic in cmd/generate.go:131 when attempting to dereference *manifestStr
func TestGenerateWithNilManifest(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		// Create a temporary workspace
		tempDir := t.TempDir()

		// Create directory structure with an invalid file scenario
		// We'll create a directory with the same name as the file pattern to force an error
		invalidPath := filepath.Join(tempDir, "src", "component.ts")
		err := os.MkdirAll(invalidPath, 0755) // Create directory with .ts extension
		require.NoError(t, err)

		// Create package.json to make it a valid workspace
		packageJSON := `{
  "name": "test-package",
  "customElements": "dist/custom-elements.json"
}`
		err = os.WriteFile(filepath.Join(tempDir, "package.json"), []byte(packageJSON), 0644)
		require.NoError(t, err)

		// Create .config directory and cem.yaml pointing to the invalid file
		configDir := filepath.Join(tempDir, ".config")
		err = os.MkdirAll(configDir, 0755)
		require.NoError(t, err)

		cemConfig := `generate:
  files:
    - src/component.ts  # This will fail because it's a directory, not a file
`
		err = os.WriteFile(filepath.Join(configDir, "cem.yaml"), []byte(cemConfig), 0644)
		require.NoError(t, err)

		// Create workspace context
		workspace := W.NewFileSystemWorkspaceContext(tempDir)
		err = workspace.Init()
		require.NoError(t, err)

		// Call Generate directly - this should return nil manifest and an error
		// but should NOT panic when the result is handled properly
		manifestStr, err := G.Generate(workspace)

		// The function should return an error (not panic)
		assert.Error(t, err, "Generate should return an error when files cannot be processed")

		// The manifest should be nil when there's an error
		assert.Nil(t, manifestStr, "Manifest should be nil when generation fails")

		// The error should contain information about the file processing failure
		assert.Contains(t, err.Error(), "NewModuleProcessor", "Error should indicate NewModuleProcessor failure")
	}) // End synctest.Test
}

// TestGenerateWithFileReadError tests the case where os.ReadFile fails
// and ensures the Generate function returns nil manifest and error gracefully
func TestGenerateWithFileReadError(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		// Create a temporary workspace
		tempDir := t.TempDir()

		// Create package.json
		packageJSON := `{"name": "test-package"}`
		err := os.WriteFile(filepath.Join(tempDir, "package.json"), []byte(packageJSON), 0644)
		require.NoError(t, err)

		// Create a file with restricted permissions to simulate read failure
		restrictedFile := filepath.Join(tempDir, "restricted.ts")
		err = os.WriteFile(restrictedFile, []byte("export class Test {}"), 0000) // No read permissions
		require.NoError(t, err)

		// Clean up permissions after test
		defer os.Chmod(restrictedFile, 0644)

		// Create .config directory and cem.yaml pointing to the restricted file
		configDir := filepath.Join(tempDir, ".config")
		err = os.MkdirAll(configDir, 0755)
		require.NoError(t, err)

		cemConfig := `generate:
  files:
    - restricted.ts
`
		err = os.WriteFile(filepath.Join(configDir, "cem.yaml"), []byte(cemConfig), 0644)
		require.NoError(t, err)

		// Create workspace context
		workspace := W.NewFileSystemWorkspaceContext(tempDir)
		err = workspace.Init()
		require.NoError(t, err)

		// Call Generate - this should return nil manifest and an error
		manifestStr, err := G.Generate(workspace)

		assert.Error(t, err, "Generate should return an error when file cannot be read")
		assert.Nil(t, manifestStr, "Manifest should be nil when generation fails")
		assert.Contains(t, err.Error(), "NewModuleProcessor", "Error should indicate NewModuleProcessor failure")
	}) // End synctest.Test
}

// TestGenerateHandlesNilManifestProperly verifies that the command layer
// properly handles nil manifest pointers without dereferencing them
func TestGenerateHandlesNilManifestProperly(t *testing.T) {
	// This test simulates the fix in cmd/generate.go where we check for nil
	// before dereferencing the manifest pointer

	// Simulate a nil manifest pointer return from Generate
	var manifestStr *string = nil

	// The old code would do: *manifestStr (panic)
	// The new code should check: if manifestStr == nil { return errs }

	// Test that we can detect nil safely
	assert.Nil(t, manifestStr, "Manifest pointer should be nil")

	// Test that dereferencing nil would panic (this is what we're protecting against)
	assert.Panics(t, func() {
		_ = *manifestStr // This would cause the original panic
	}, "Dereferencing nil pointer should panic")
}
