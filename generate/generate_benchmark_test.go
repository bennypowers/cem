package generate_test

import (
	"os"
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/generate"
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

	path, err := filepath.Abs("../test/fixtures/")
	if err != nil {
		b.Fatalf("BenchmarkGenerate failed to resolve project dir: %v", err)
	}
	ctx := manifest.NewFileSystemWorkspaceContext(path)
	if err := ctx.Init(); err != nil {
		b.Fatalf("BenchmarkGenerate failed to init context: %v", err)
	}

	// Run the Generate function, measuring its speed.
	cfg, err := ctx.Config()
	if err != nil {
		b.Fatalf("BenchmarkGenerate failed to load config: %v", err)
	}

	cfg.Generate.Files = matches

	for b.Loop() {
		out, err := generate.Generate(ctx, cfg)
		if err != nil {
			b.Errorf("BenchmarkGenerate generate returned error: %v", err)
		}
		lastOut = *out
	}

	if err := os.MkdirAll("../docs/data/", 0755); err != nil {
		b.Fatalf("BenchmarkGenerate failed to create ../docs/data dir: %v", err)
	}
	if err := os.WriteFile("../docs/data/lastBenchmark.json", []byte(lastOut), 0644); err != nil {
		b.Fatalf("BenchmarkGenerate could not write output: %v", err)
	}
}
