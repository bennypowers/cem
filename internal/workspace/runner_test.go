package workspace_test

import (
	"errors"
	"sync/atomic"
	"testing"

	"bennypowers.dev/cem/internal/workspace"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestForEachPackage_RunsAllPackages(t *testing.T) {
	root := absFixture(t, "multi-package-workspace")
	var count atomic.Int32
	results := workspace.ForEachPackage(root, func(pkg workspace.PackageInfo) error {
		count.Add(1)
		return nil
	})
	assert.Equal(t, int32(2), count.Load())
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
