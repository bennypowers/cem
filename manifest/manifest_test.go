package manifest

import "testing"

func TestNormalizeSourcePath(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"src/button.ts", "src/button.js"},
		{"src/button.tsx", "src/button.js"},
		{"src/button.jsx", "src/button.js"},
		{"src/utils.mts", "src/utils.mjs"},
		{"src/utils.cts", "src/utils.cjs"},
		{"src/button.js", "src/button.js"},
		{"src/style.css", "src/style.css"},
		{"README.md", "README.md"},
		{"elements/rh-select/rh-select.ts", "elements/rh-select/rh-select.js"},
		{"deep/path/component.tsx", "deep/path/component.js"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := NormalizeSourcePath(tt.input)
			if got != tt.want {
				t.Errorf("NormalizeSourcePath(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
