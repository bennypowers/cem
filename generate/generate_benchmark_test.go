package generate

import (
	"os"
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/cmd/config"
	"bennypowers.dev/cem/manifest"
	DS "github.com/bmatcuk/doublestar"
)

// BenchmarkGenerate runs the Generate function on all test fixtures to measure performance.
func BenchmarkGenerate(b *testing.B) {
	// Gather all .ts files in the test-fixtures directory as input.
	matches, err := DS.Glob("../test/fixtures/**/*.ts")
	if err != nil {
		b.Fatal(err)
	}
	if len(matches) == 0 {
		b.Skip("No test fixtures found")
	}

	var lastOut string

	ctx := manifest.NewLocalFSProjectContext("../test/fixtures")
	if err := ctx.Init(); err != nil {
		b.Fatalf("Failed to init context: %v", err)
	}

	for b.Loop() {
		// Run the Generate function, measuring its speed.
		specs, err := filepath.Glob("../test/fixtures/*/design-tokens.json")
		if err != nil {
			b.Errorf("Failed to load design tokens: %v", err)
		}
		cfg, err := config.LoadConfig(ctx)
		if err != nil {
			b.Fatalf("failed to load config: %v", err)
		}
		if err != nil {
			b.Fatalf("Failed to get config: %v", err)
		}

		cfg.Generate.Files = matches
		cfg.Generate.DesignTokens.Spec = specs[0]
		cfg.Generate.DesignTokens.Prefix = "token"
		cfg.Generate.DemoDiscovery.FileGlob = "demos/."

		out, err := Generate(ctx, cfg)
		if err != nil {
			b.Errorf("Generate returned error: %v", err)
		}
		lastOut = *out
	}

	if err := os.MkdirAll("../docs/data/", 0755); err != nil {
		b.Fatalf("failed to create ../docs/data dir: %v", err)
	}
	if err := os.WriteFile("../docs/data/lastBenchmark.json", []byte(lastOut), 0644); err != nil {
		b.Fatalf("Could not write output: %v", err)
	}
}
