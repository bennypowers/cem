package generate

import (
	"path/filepath"
	"testing"
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

	excludes := []string{}
	designTokens := "./" + filepath.Join("test-fixtures", "design-tokens.json")
	designTokensPrefix := "token"

	for b.Loop() {
		// Run the Generate function, measuring its speed.
		_, err := Generate(matches, excludes, designTokens, designTokensPrefix)
		if err != nil {
			b.Errorf("Generate returned error: %v", err)
		}
	}
}
