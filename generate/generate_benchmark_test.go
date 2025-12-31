/*
Copyright © 2025 Benny Powers

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package generate_test

import (
	"os"
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/generate"
	W "bennypowers.dev/cem/workspace"
	DS "github.com/bmatcuk/doublestar"
)

// BenchmarkGenerate runs the Generate function on all test fixtures to measure performance.
func BenchmarkGenerate(b *testing.B) {
	path, err := filepath.Abs("testdata/fixtures/")
	if err != nil {
		b.Fatalf("BenchmarkGenerate failed to resolve project dir: %v", err)
	}

	ctx := W.NewFileSystemWorkspaceContext(path)
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
