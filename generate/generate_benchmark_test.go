package generate

import (
	"os"
	"path/filepath"
	"testing"

	C "bennypowers.dev/cem/cmd/config"
)

// BenchmarkGenerate runs the Generate function on all test fixtures to measure performance.
func BenchmarkGenerate(b *testing.B) {
	// Gather all .ts files in the test-fixtures directory as input.
	matches, err := filepath.Glob("../test/fixtures/**/*.ts")
	if err != nil {
		b.Fatal(err)
	}
	if len(matches) == 0 {
		b.Skip("No test fixtures found")
	}

	var lastOut string

	for b.Loop() {
		// Run the Generate function, measuring its speed.
		specs, err := filepath.Glob("../test/fixtures/*/design-tokens.json")
		if err != nil {
			b.Errorf("Failed to load design tokens: %v", err)
		}
		cfg := C.CemConfig{
			SourceControlRootUrl: "",
			Generate: C.GenerateConfig{
				Files: matches,
				DesignTokens: C.DesignTokensConfig{
					Spec: specs[0], // todo: should accept a slice
					Prefix: "token",
				},
				DemoDiscovery: C.DemoDiscoveryConfig{
					FileGlob: "demos/.",
					URLPattern: "",
					URLTemplate: "",
				},
			},
		}
		out, err := Generate(&cfg)
		if err != nil {
			b.Errorf("Generate returned error: %v", err)
		}
		lastOut = *out
	}

	err = os.WriteFile("../docs/site/benchmark-last-output.json", []byte(lastOut), 0644)

	if err != nil {
		b.Errorf("Could not write output: %v", err)
	}
}
