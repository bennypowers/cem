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
	path, err := filepath.Abs("../test/fixtures/")
	if err != nil {
		b.Fatalf("BenchmarkGenerate failed to resolve project dir: %v", err)
	}

	ctx := manifest.NewFileSystemWorkspaceContext(path)
	if err := ctx.Init(); err != nil {
		b.Fatalf("BenchmarkGenerate failed to init context: %v", err)
	}

	// Gather all .ts files in the test-fixtures directory as input.
	matches, err := DS.Glob(filepath.Join(path, "**/*.ts"))
	if err != nil {
		b.Fatal(err)
	}

	if len(matches) == 0 {
		b.Skip("No test fixtures found")
	}

	// Run the Generate function, measuring its speed.
	cfg, err := ctx.Config()
	if err != nil {
		b.Fatalf("BenchmarkGenerate failed to load config: %v", err)
	}

	cfg.Generate.Files = matches

	var lastOut string

	for b.Loop() {
		out, err := generate.Generate(ctx)
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
