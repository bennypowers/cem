package generate

import (
	"path/filepath"
	"testing"
	C "bennypowers.dev/cem/cmd/config"
)

// BenchmarkGenerate runs the Generate function on all test fixtures to measure performance.
func BenchmarkGenerate(b *testing.B) {
	// Gather all .ts files in the test-fixtures directory as input.
	matches, err := filepath.Glob("test-fixtures/*.ts")
	if err != nil {
		b.Fatal(err)
	}
	if len(matches) == 0 {
		b.Skip("No test fixtures found")
	}

	for b.Loop() {
		// Run the Generate function, measuring its speed.
		cfg := C.CemConfig{
			TagPrefix: "demo",
			Generate: C.GenerateConfig{
				Files: matches,
				DesignTokensSpec: "./" + filepath.Join("test-fixtures", "design-tokens.json"),
				DesignTokensPrefix: "token",
				DemoDiscovery: C.DemoDiscoveryConfig{
					FilePattern: "demos/.",
					// Canonical public source control URL corresponding to project root on primary branch.
					// e.g. https://github.com/bennypowers/cem/tree/main/
					SourceControlUrl: "",
					URLPattern: "",
					URLTemplate: "",
				},
			},
		}
		_, err := Generate(&cfg)
		if err != nil {
			b.Errorf("Generate returned error: %v", err)
		}
	}
}
