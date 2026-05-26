package integration_test

import (
	"testing"

	"bennypowers.dev/cem/internal/platform/testutil"
	"github.com/stretchr/testify/assert"
)

// Inline: integration test, scalar assertions

// TestZedBinaryNamingConsistency verifies that the Zed extension references
// the correct binary names that match the go-release-workflows output.
func TestZedBinaryNamingConsistency(t *testing.T) {
	// Read Zed extension source
	mfs := testutil.LoadTestdataFS(t, "../../extensions/zed/src", "/")
	zedSource := testutil.ReadFixture(t, mfs, "/lib.rs")

	sourceStr := string(zedSource)

	// Assert on actual Rust mapping literals to catch logic regressions
	expectedArchMappings := []string{
		`zed::Architecture::Aarch64 => "arm64"`,
		`zed::Architecture::X8664 => "x64"`,
	}
	for _, mapping := range expectedArchMappings {
		assert.Contains(t, sourceStr, mapping,
			"Zed extension should map architecture via %q", mapping)
	}

	expectedOsMappings := []string{
		`zed::Os::Mac => ("darwin", "")`,
		`zed::Os::Linux => ("linux", "")`,
		`zed::Os::Windows => ("win32", ".exe")`,
	}
	for _, mapping := range expectedOsMappings {
		assert.Contains(t, sourceStr, mapping,
			"Zed extension should map OS via %q", mapping)
	}

	// Ensure old naming patterns are NOT present
	oldPatterns := []string{
		"amd64",
		"windows",
	}
	for _, pattern := range oldPatterns {
		assert.NotContains(t, sourceStr, pattern,
			"Zed extension should NOT reference old pattern %q", pattern)
	}
}
