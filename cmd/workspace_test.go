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

// generateWorkspaceFixture runs cem generate on the workspace-generate fixture
// and returns the project dir. Most workspace tests need manifests to exist first.
func generateWorkspaceFixture(t *testing.T) string {
	t.Helper()
	projectDir := setupTest(t, "workspace-generate")
	runCemCommand(t, projectDir, "generate")
	return projectDir
}

func TestWorkspaceGenerate_PartitionsFilesByPackage(t *testing.T) {
	projectDir := generateWorkspaceFixture(t)

	for _, pkg := range []struct {
		dir       string
		wantClass string
		wantPath  string
		excludes  string
	}{
		{"packages/button", "test-button", "src/button.js", "test-card"},
		{"packages/card", "test-card", "src/card.js", "test-button"},
	} {
		manifestPath := filepath.Join(projectDir, pkg.dir, "custom-elements.json")
		data, err := os.ReadFile(manifestPath)
		require.NoError(t, err, "manifest should exist for %s", pkg.dir)

		content := string(data)

		assert.Contains(t, content, `"path": "`+pkg.wantPath+`"`,
			"module path in %s should be package-relative", pkg.dir)

		assert.NotContains(t, content, `"path": "packages/`,
			"module path in %s must not have workspace-root prefix", pkg.dir)

		assert.True(t, strings.Contains(content, pkg.wantClass),
			"%s manifest should contain %s", pkg.dir, pkg.wantClass)

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
	runCemCommand(t, projectDir, "generate", "-p", "packages/button", "src/**/*.ts")

	buttonManifest := filepath.Join(projectDir, "packages", "button", "custom-elements.json")
	data, err := os.ReadFile(buttonManifest)
	require.NoError(t, err)
	assert.Contains(t, string(data), "TestButton")

	cardManifest := filepath.Join(projectDir, "packages", "card", "custom-elements.json")
	_, err = os.ReadFile(cardManifest)
	assert.True(t, os.IsNotExist(err), "card manifest should not exist when only button was targeted")
}

func TestWorkspaceValidate_ValidatesAllPackages(t *testing.T) {
	projectDir := generateWorkspaceFixture(t)
	stdout, stderr := runCemCommand(t, projectDir, "validate")
	combined := stdout + stderr

	assert.Contains(t, combined, "@test/button")
	assert.Contains(t, combined, "@test/card")
	assert.Contains(t, combined, "Validated manifests")
}

func TestWorkspaceHealth_ScoresAllPackages(t *testing.T) {
	projectDir := generateWorkspaceFixture(t)
	stdout, stderr := runCemCommand(t, projectDir, "health")
	combined := stdout + stderr

	assert.Contains(t, combined, "@test/button")
	assert.Contains(t, combined, "@test/card")
	assert.Contains(t, combined, "Checked health")
}

func TestWorkspaceListTags_ShowsAllPackages(t *testing.T) {
	projectDir := generateWorkspaceFixture(t)
	stdout, _ := runCemCommand(t, projectDir, "list", "tags")

	assert.Contains(t, stdout, "@test/button")
	assert.Contains(t, stdout, "@test/card")
	assert.Contains(t, stdout, "TestButton")
	assert.Contains(t, stdout, "TestCard")
}

func TestWorkspaceListModules_ShowsAllPackages(t *testing.T) {
	projectDir := generateWorkspaceFixture(t)
	stdout, _ := runCemCommand(t, projectDir, "list", "modules")

	assert.Contains(t, stdout, "@test/button")
	assert.Contains(t, stdout, "@test/card")
	assert.Contains(t, stdout, "src/button.js")
	assert.Contains(t, stdout, "src/card.js")
}

func TestWorkspaceList_ShowsAllPackages(t *testing.T) {
	projectDir := generateWorkspaceFixture(t)
	stdout, _ := runCemCommand(t, projectDir, "list")

	assert.Contains(t, stdout, "@test/button")
	assert.Contains(t, stdout, "@test/card")
}

func TestWorkspaceSearch_FindsAcrossPackages(t *testing.T) {
	projectDir := generateWorkspaceFixture(t)
	stdout, _ := runCemCommand(t, projectDir, "search", "button")

	assert.Contains(t, stdout, "@test/button")
	assert.Contains(t, stdout, "src/button.js")
}

func TestWorkspaceSearch_NoResults(t *testing.T) {
	projectDir := generateWorkspaceFixture(t)
	stdout, _ := runCemCommand(t, projectDir, "search", "zzz-nonexistent-zzz")

	assert.Contains(t, stdout, "Searched manifests")
	assert.NotContains(t, stdout, "TestButton")
}

func TestWorkspaceExport_SucceedsWithConfig(t *testing.T) {
	// Root config has export.react.output, which cascades to packages
	projectDir := generateWorkspaceFixture(t)
	stdout, stderr := runCemCommand(t, projectDir, "export")
	combined := stdout + stderr

	assert.Contains(t, combined, "Exported wrappers")
}

func TestWorkspaceExport_WithFormat(t *testing.T) {
	projectDir := generateWorkspaceFixture(t)
	stdout, stderr := runCemCommand(t, projectDir, "export", "--format", "react")
	combined := stdout + stderr

	// Export runs per-package and reports results
	assert.Contains(t, combined, "Exported wrappers")
}

func TestWorkspaceExport_OutputPerPackage(t *testing.T) {
	// Root config has export.react.output: "react".
	// Each package should get its own react/ dir, not a shared one at the root.
	projectDir := generateWorkspaceFixture(t)
	runCemCommand(t, projectDir, "export", "--format", "react")

	// Verify each package has its own react output (not at workspace root)
	for _, pkg := range []string{"packages/button", "packages/card"} {
		reactDir := filepath.Join(projectDir, pkg, "react")
		_, err := os.Stat(reactDir)
		assert.NoError(t, err, "react export dir should exist for %s", pkg)
	}
	// Workspace root should NOT have a react/ dir
	_, err := os.Stat(filepath.Join(projectDir, "react"))
	assert.True(t, os.IsNotExist(err), "react dir should not exist at workspace root")
}

func TestWorkspaceGenerate_DemoDiscovery(t *testing.T) {
	// Root config has demoDiscovery with fileGlob, urlPattern, and urlTemplate.
	// In workspace mode, demo files are package-relative but the urlPattern
	// is root-level. The workspace fallback should prepend the package dir name
	// so the urlPattern matches and URLs are generated.
	projectDir := generateWorkspaceFixture(t)

	for _, pkg := range []struct {
		dir     string
		wantURL string
	}{
		{"packages/button", "https://test.example.com/components/button/demo/basic/"},
		{"packages/card", "https://test.example.com/components/card/demo/basic/"},
	} {
		manifestPath := filepath.Join(projectDir, pkg.dir, "custom-elements.json")
		data, err := os.ReadFile(manifestPath)
		require.NoError(t, err)
		content := string(data)

		assert.Contains(t, content, `"kind": "javascript-module"`,
			"%s manifest should have modules", pkg.dir)

		assert.Contains(t, content, `"demos"`,
			"%s manifest should have demos attached to declarations", pkg.dir)

		assert.Contains(t, content, pkg.wantURL,
			"%s manifest should contain generated demo URL", pkg.dir)
	}
}

func TestWorkspaceListAttributes_FindsInCorrectPackage(t *testing.T) {
	projectDir := generateWorkspaceFixture(t)
	stdout, _ := runCemCommand(t, projectDir, "list", "attributes", "-t", "test-button")

	assert.Contains(t, stdout, "Listed from manifests")
}
