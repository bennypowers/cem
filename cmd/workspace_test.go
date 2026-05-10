//go:build e2e

package cmd_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWorkspaceGenerate_PartitionsFilesByPackage(t *testing.T) {
	projectDir := setupTest(t, "workspace-generate")
	runCemCommand(t, projectDir, "generate")

	for _, pkg := range []struct {
		dir       string
		wantClass string
		wantPath  string
		excludes  string
	}{
		{"packages/button", "TestButton", "src/button.js", "TestCard"},
		{"packages/card", "TestCard", "src/card.js", "TestButton"},
	} {
		manifestPath := filepath.Join(projectDir, pkg.dir, "custom-elements.json")
		data, err := os.ReadFile(manifestPath)
		require.NoError(t, err, "manifest should exist for %s", pkg.dir)

		content := string(data)

		// Module paths must be package-relative
		assert.Contains(t, content, `"path": "`+pkg.wantPath+`"`,
			"module path in %s should be package-relative", pkg.dir)

		// Must not contain workspace-root-relative paths
		assert.NotContains(t, content, `"path": "packages/`,
			"module path in %s must not have workspace-root prefix", pkg.dir)

		// Each package should contain its own class
		assert.True(t, strings.Contains(content, pkg.wantClass),
			"%s manifest should contain %s", pkg.dir, pkg.wantClass)

		// Each package must NOT contain other packages' classes
		assert.False(t, strings.Contains(content, pkg.excludes),
			"%s manifest must not contain %s", pkg.dir, pkg.excludes)
	}
}

func TestWorkspaceGenerate_OutputFlagErrors(t *testing.T) {
	projectDir := setupTest(t, "workspace-generate")
	stdout, stderr := runCemCommand(t, projectDir, "generate", "-o", "out.json")
	combined := stdout + stderr
	assert.Contains(t, combined, "cannot use --output in workspace mode")
}

func TestWorkspaceGenerate_SinglePackageOverride(t *testing.T) {
	projectDir := setupTest(t, "workspace-generate")
	// With -p, only the targeted package is processed (single-package mode).
	// Pass the file patterns explicitly since per-package config doesn't exist.
	runCemCommand(t, projectDir, "generate", "-p", "packages/button", "src/**/*.ts")

	buttonManifest := filepath.Join(projectDir, "packages", "button", "custom-elements.json")
	data, err := os.ReadFile(buttonManifest)
	require.NoError(t, err)
	assert.Contains(t, string(data), "TestButton")

	cardManifest := filepath.Join(projectDir, "packages", "card", "custom-elements.json")
	_, err = os.ReadFile(cardManifest)
	assert.True(t, os.IsNotExist(err), "card manifest should not exist when only button was targeted")
}
