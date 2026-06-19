//go:build e2e

package cmd_test

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/charmbracelet/x/ansi"
	"github.com/nsf/jsondiff"
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

	for _, pkg := range []string{"button", "card"} {
		manifestPath := filepath.Join(projectDir, "packages", pkg, "custom-elements.json")
		data, err := os.ReadFile(manifestPath)
		require.NoError(t, err, "manifest should exist for packages/%s", pkg)
		compareGolden(t, workspaceGolden("generate-"+pkg, true), string(data), true)
	}
}

func TestWorkspaceGenerate_OutputFlagErrors(t *testing.T) {
	projectDir := setupTest(t, "workspace-generate")
	_, stderr := runCemCommand(t, projectDir, "generate", "-o", "out.json")
	compareGolden(t, workspaceGolden("generate-output-flag-error", false), stderr, false)
}

func TestWorkspaceGenerate_SinglePackageOverride(t *testing.T) {
	projectDir := setupTest(t, "workspace-generate")
	runCemCommand(t, projectDir, "generate", "-p", "packages/button", "src/**/*.ts")

	buttonManifest := filepath.Join(projectDir, "packages", "button", "custom-elements.json")
	data, err := os.ReadFile(buttonManifest)
	require.NoError(t, err)
	compareGolden(t, workspaceGolden("generate-single-package-button", true), string(data), true)

	cardManifest := filepath.Join(projectDir, "packages", "card", "custom-elements.json")
	_, err = os.ReadFile(cardManifest)
	assert.True(t, os.IsNotExist(err), "card manifest should not exist when only button was targeted")
}

func TestWorkspaceValidate_ValidatesAllPackages(t *testing.T) {
	projectDir := generateWorkspaceFixture(t)
	stdout, _ := runCemCommand(t, projectDir, "validate")
	cleaned := ansi.Strip(stdout)
	compareGolden(t, workspaceGolden("validate", false), cleaned, false)
}

var (
	reSpaces = regexp.MustCompile(`  +`)
	reTmpDir = regexp.MustCompile(`/tmp/cem-test-\d+/workspace-generate`)
)

func normalizeTextOutput(s string) string {
	var lines []string
	for line := range strings.SplitSeq(s, "\n") {
		line = strings.TrimRight(line, " \t")
		line = reSpaces.ReplaceAllString(line, " ")
		lines = append(lines, line)
	}
	result := strings.Join(lines, "\n")
	result = reTmpDir.ReplaceAllString(result, "$TMPDIR")
	return result
}

func compareGolden(t *testing.T, goldenPath, actual string, isJSON bool) {
	t.Helper()

	if *update {
		if err := os.MkdirAll(filepath.Dir(goldenPath), 0755); err != nil {
			t.Fatalf("failed to create golden dir: %v", err)
		}
		content := actual
		if !isJSON {
			content = normalizeTextOutput(content)
		}
		if err := os.WriteFile(goldenPath, []byte(content), 0644); err != nil {
			t.Fatalf("failed to write golden file: %v", err)
		}
		return
	}

	expected, err := os.ReadFile(goldenPath)
	if err != nil {
		t.Fatalf("golden file missing: %s (run with -update to generate)", goldenPath)
	}

	if isJSON {
		opts := jsondiff.DefaultConsoleOptions()
		diff, report := jsondiff.Compare(expected, []byte(actual), &opts)
		if diff != jsondiff.FullMatch {
			t.Errorf("JSON diff:\n%s", report)
		}
	} else {
		got := normalizeTextOutput(actual)
		want := string(expected)
		if got != want {
			t.Errorf("output mismatch vs %s\n--- expected ---\n%s\n--- got ---\n%s", goldenPath, want, got)
		}
	}
}

func workspaceGolden(name string, isJSON bool) string {
	ext := ".golden"
	if isJSON {
		ext = ".golden.json"
	}
	return filepath.Join("testdata", "goldens", "workspace."+name+ext)
}

func TestWorkspaceHealth_ScoresAllPackages(t *testing.T) {
	projectDir := generateWorkspaceFixture(t)
	stdout, _ := runCemCommand(t, projectDir, "health")
	cleaned := ansi.Strip(stdout)
	compareGolden(t, workspaceGolden("health-all", false), cleaned, false)
}

func TestWorkspaceHealth_TagNameFilter(t *testing.T) {
	projectDir := generateWorkspaceFixture(t)
	stdout, _ := runCemCommand(t, projectDir, "health", "-t", "test-button")
	cleaned := ansi.Strip(stdout)
	compareGolden(t, workspaceGolden("health-tag-filter", false), cleaned, false)
}

func TestWorkspaceHealth_FormatJSON(t *testing.T) {
	projectDir := generateWorkspaceFixture(t)
	stdout, _ := runCemCommand(t, projectDir, "health", "--format", "json")
	compareGolden(t, workspaceGolden("health-json", true), stdout, true)
}

func TestWorkspaceHealth_FormatJSON_WithTagFilter(t *testing.T) {
	projectDir := generateWorkspaceFixture(t)
	stdout, _ := runCemCommand(t, projectDir, "health", "--format", "json", "-t", "test-button")
	compareGolden(t, workspaceGolden("health-json-tag-filter", true), stdout, true)
}

func TestWorkspaceHealth_FormatJSON_NoMatch(t *testing.T) {
	projectDir := generateWorkspaceFixture(t)
	stdout, _ := runCemCommand(t, projectDir, "health", "--format", "json", "-t", "nonexistent-element")
	compareGolden(t, workspaceGolden("health-json-no-match", true), stdout, true)
}

func TestWorkspaceHealth_DeprecatedComponentFlag(t *testing.T) {
	projectDir := generateWorkspaceFixture(t)
	stdout, stderr := runCemCommand(t, projectDir, "health", "--component", "test-button")
	cleaned := ansi.Strip(stdout)
	compareGolden(t, workspaceGolden("health-deprecated-component", false), cleaned, false)
	// Justification for inline assertion: deprecation warning is cobra-generated
	// stderr whose exact format is cobra's concern, not ours. Golden-matching
	// would couple tests to cobra internals.
	assert.Contains(t, stderr, "deprecated", "stderr should warn that --component is deprecated")
}

func TestWorkspaceListTags_ShowsAllPackages(t *testing.T) {
	projectDir := generateWorkspaceFixture(t)
	stdout, _ := runCemCommand(t, projectDir, "list", "tags")
	cleaned := ansi.Strip(stdout)
	compareGolden(t, workspaceGolden("list-tags", false), cleaned, false)
}

func TestWorkspaceListModules_ShowsAllPackages(t *testing.T) {
	projectDir := generateWorkspaceFixture(t)
	stdout, _ := runCemCommand(t, projectDir, "list", "modules")
	cleaned := ansi.Strip(stdout)
	compareGolden(t, workspaceGolden("list-modules", false), cleaned, false)
}

func TestWorkspaceList_ShowsAllPackages(t *testing.T) {
	projectDir := generateWorkspaceFixture(t)
	stdout, _ := runCemCommand(t, projectDir, "list")
	cleaned := ansi.Strip(stdout)
	compareGolden(t, workspaceGolden("list-all", false), cleaned, false)
}

func TestWorkspaceSearch_FindsAcrossPackages(t *testing.T) {
	projectDir := generateWorkspaceFixture(t)
	stdout, _ := runCemCommand(t, projectDir, "search", "button")
	cleaned := ansi.Strip(stdout)
	compareGolden(t, workspaceGolden("search-button", false), cleaned, false)
}

func TestWorkspaceSearch_NoResults(t *testing.T) {
	projectDir := generateWorkspaceFixture(t)
	stdout, _ := runCemCommand(t, projectDir, "search", "zzz-nonexistent-zzz")
	cleaned := ansi.Strip(stdout)
	compareGolden(t, workspaceGolden("search-no-results", false), cleaned, false)
}

func TestWorkspaceExport_SucceedsWithConfig(t *testing.T) {
	projectDir := generateWorkspaceFixture(t)
	stdout, _ := runCemCommand(t, projectDir, "export")
	cleaned := ansi.Strip(stdout)
	compareGolden(t, workspaceGolden("export", false), cleaned, false)
}

func TestWorkspaceExport_WithFormat(t *testing.T) {
	projectDir := generateWorkspaceFixture(t)
	stdout, _ := runCemCommand(t, projectDir, "export", "--format", "react")
	cleaned := ansi.Strip(stdout)
	compareGolden(t, workspaceGolden("export-react", false), cleaned, false)
}

func TestWorkspaceExport_OutputPerPackage(t *testing.T) {
	projectDir := generateWorkspaceFixture(t)
	runCemCommand(t, projectDir, "export", "--format", "react")

	// Justification for inline assertions: these verify filesystem side effects
	// (directory existence), not command output. Golden files compare text/JSON
	// output; filesystem structure checks have no output to golden-match.
	for _, pkg := range []string{"packages/button", "packages/card"} {
		reactDir := filepath.Join(projectDir, pkg, "react")
		_, err := os.Stat(reactDir)
		assert.NoError(t, err, "react export dir should exist for %s", pkg)
	}
	_, err := os.Stat(filepath.Join(projectDir, "react"))
	assert.True(t, os.IsNotExist(err), "react dir should not exist at workspace root")
}

func TestWorkspaceGenerate_DemoDiscovery(t *testing.T) {
	projectDir := generateWorkspaceFixture(t)

	for _, pkg := range []string{"button", "card"} {
		manifestPath := filepath.Join(projectDir, "packages", pkg, "custom-elements.json")
		data, err := os.ReadFile(manifestPath)
		require.NoError(t, err)
		compareGolden(t, workspaceGolden("generate-demo-"+pkg, true), string(data), true)
	}
}

func TestWorkspaceListAttributes_FindsInCorrectPackage(t *testing.T) {
	projectDir := generateWorkspaceFixture(t)
	stdout, _ := runCemCommand(t, projectDir, "list", "attributes", "-t", "test-button")
	cleaned := ansi.Strip(stdout)
	compareGolden(t, workspaceGolden("list-attributes-button", false), cleaned, false)
}
