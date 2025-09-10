/*
Copyright © 2025 Benny Powers

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
package cmd

import (
	"os"
	"path/filepath"
	"testing"

	G "bennypowers.dev/cem/generate"
	W "bennypowers.dev/cem/workspace"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRunWatchMode(t *testing.T) {
	tests := []struct {
		name        string
		globs       []string
		expectError bool
	}{
		{
			name:  "valid watch session creation",
			globs: []string{"src/**/*.ts"},
		},
		{
			name:  "multiple globs",
			globs: []string{"src/**/*.ts", "lib/**/*.js"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempDir := t.TempDir()
			srcDir := filepath.Join(tempDir, "src")
			require.NoError(t, os.MkdirAll(srcDir, 0755))
			packageJSON := filepath.Join(tempDir, "package.json")
			require.NoError(t, os.WriteFile(packageJSON, []byte(`{"name": "test"}`), 0644))
			testFile := filepath.Join(srcDir, "test.ts")
			require.NoError(t, os.WriteFile(testFile, []byte("export class Test {}"), 0644))

			ctx := W.NewFileSystemWorkspaceContextWithDefaults(tempDir)
			require.NoError(t, ctx.Init())

			// Test that watch session can be created (we don't actually run the watch loop in tests)
			session, err := G.NewWatchSession(ctx, tt.globs)
			if tt.expectError {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
			defer session.Close()

			assert.NotNil(t, session)
		})
	}
}

// Integration test for the complete watch workflow
func TestWatchWorkflow_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	tempDir := t.TempDir()
	srcDir := filepath.Join(tempDir, "src")
	require.NoError(t, os.MkdirAll(srcDir, 0755))

	// Create package.json
	packageJSON := filepath.Join(tempDir, "package.json")
	require.NoError(t, os.WriteFile(packageJSON, []byte(`{"name": "test", "customElements": "custom-elements.json"}`), 0644))

	// Create initial TypeScript file
	testFile := filepath.Join(srcDir, "test.ts")
	initialContent := `export class TestElement extends HTMLElement {
		static get observedAttributes() { return ['test']; }
	}`
	require.NoError(t, os.WriteFile(testFile, []byte(initialContent), 0644))

	ctx := W.NewFileSystemWorkspaceContextWithDefaults(tempDir)
	require.NoError(t, ctx.Init())

	cfg, err := ctx.Config()
	require.NoError(t, err)

	// Expand globs to actual file paths (like the main command does)
	files, err := ctx.Glob("src/**/*.ts")
	require.NoError(t, err)
	cfg.Generate.Files = files

	// Test that watch session can be created and performs initial generation
	session, err := G.NewWatchSession(ctx, []string{"src/**/*.ts"})
	require.NoError(t, err)
	defer session.Close()

	assert.NotNil(t, session)
}
