package workspace_test

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/internal/workspace"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestForEachPackage_RunsAllPackages(t *testing.T) {
	root := absFixture(t, "multi-package-workspace")
	var count int
	results := workspace.ForEachPackage(root, func(pkg workspace.PackageInfo) error {
		count++
		return nil
	})
	assert.Equal(t, 2, count)
	assert.Len(t, results, 2)
	for _, r := range results {
		assert.NoError(t, r.Err)
	}
}

func TestForEachPackage_CollectsErrors(t *testing.T) {
	root := absFixture(t, "multi-package-workspace")
	results := workspace.ForEachPackage(root, func(pkg workspace.PackageInfo) error {
		if pkg.Name == "@test/alpha" {
			return errors.New("alpha broke")
		}
		return nil
	})
	assert.Len(t, results, 2)
	var errCount int
	for _, r := range results {
		if r.Err != nil {
			errCount++
			assert.Equal(t, "@test/alpha", r.Package.Name)
			assert.Contains(t, r.Err.Error(), "alpha broke")
		}
	}
	assert.Equal(t, 1, errCount)
}

func TestForEachPackage_NoPackages(t *testing.T) {
	root := absFixture(t, "workspace-mode-no-manifest")
	results := workspace.ForEachPackage(root, func(pkg workspace.PackageInfo) error {
		t.Fatal("should not be called")
		return nil
	})
	assert.Empty(t, results)
}

func TestForEachPackage_InvalidRoot(t *testing.T) {
	results := workspace.ForEachPackage("/nonexistent/path", func(pkg workspace.PackageInfo) error {
		t.Fatal("should not be called")
		return nil
	})
	require.Len(t, results, 1)
	assert.Error(t, results[0].Err)
}

func TestResolveWorkspaceFiles_PartitionsByPackage(t *testing.T) {
	root := absFixture(t, "multi-package-workspace")
	alphaDir := filepath.Join(root, "packages", "alpha")
	betaDir := filepath.Join(root, "packages", "beta")

	// multi-package-workspace has packages/alpha/ and packages/beta/
	// Create test source files
	require.NoError(t, os.MkdirAll(filepath.Join(alphaDir, "src"), 0755))
	require.NoError(t, os.MkdirAll(filepath.Join(betaDir, "src"), 0755))
	require.NoError(t, os.WriteFile(filepath.Join(alphaDir, "src", "alpha.ts"), []byte("// alpha"), 0644))
	require.NoError(t, os.WriteFile(filepath.Join(betaDir, "src", "beta.ts"), []byte("// beta"), 0644))
	defer func() {
		_ = os.RemoveAll(filepath.Join(alphaDir, "src"))
		_ = os.RemoveAll(filepath.Join(betaDir, "src"))
	}()

	// Root-relative pattern should partition files by package
	alphaFiles, err := workspace.ResolveWorkspaceFiles(root, []string{"packages/*/src/**/*.ts"}, alphaDir)
	require.NoError(t, err)
	assert.Equal(t, []string{"src/alpha.ts"}, alphaFiles)

	betaFiles, err := workspace.ResolveWorkspaceFiles(root, []string{"packages/*/src/**/*.ts"}, betaDir)
	require.NoError(t, err)
	assert.Equal(t, []string{"src/beta.ts"}, betaFiles)
}

func TestResolveWorkspaceFiles_NoMatchReturnsEmpty(t *testing.T) {
	root := absFixture(t, "multi-package-workspace")
	alphaDir := filepath.Join(root, "packages", "alpha")

	files, err := workspace.ResolveWorkspaceFiles(root, []string{"nonexistent/**/*.ts"}, alphaDir)
	require.NoError(t, err)
	assert.Empty(t, files)
}

func TestReportResults_NoPackages(t *testing.T) {
	err := workspace.ReportResults("Generated manifests", nil)
	assert.Error(t, err)
}

func TestReportResults_AllSucceeded(t *testing.T) {
	results := []workspace.PackageResult{
		{Package: workspace.PackageInfo{Name: "@test/a", CustomElementsRef: "custom-elements.json"}, Err: nil},
		{Package: workspace.PackageInfo{Name: "@test/b", CustomElementsRef: "custom-elements.json"}, Err: nil},
	}
	err := workspace.ReportResults("Generated manifests", results)
	assert.NoError(t, err)
}

func TestReportResults_PartialFailure(t *testing.T) {
	results := []workspace.PackageResult{
		{Package: workspace.PackageInfo{Name: "@test/a", CustomElementsRef: "custom-elements.json"}, Err: nil},
		{Package: workspace.PackageInfo{Name: "@test/b"}, Err: errors.New("parse error")},
	}
	err := workspace.ReportResults("Generated manifests", results)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "1 of 2")
}
