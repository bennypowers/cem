package definition

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchesExportPattern(t *testing.T) {
	tests := []struct {
		name       string
		modulePath string
		exportKey  string
		expected   bool
	}{
		{
			name:       "exact match",
			modulePath: "components/button.js",
			exportKey:  "./components/button.js",
			expected:   true,
		},
		{
			name:       "wildcard match - ./*",
			modulePath: "anything/here.js",
			exportKey:  "./*",
			expected:   true,
		},
		{
			name:       "wildcard match - ./lib/*",
			modulePath: "lib/utils.js",
			exportKey:  "./lib/*",
			expected:   true,
		},
		{
			name:       "no match - different path",
			modulePath: "other/path.js",
			exportKey:  "./components/button.js",
			expected:   false,
		},
		{
			name:       "no match - wrong prefix for wildcard",
			modulePath: "src/utils.js",
			exportKey:  "./lib/*",
			expected:   false,
		},
		{
			name:       "exact match without extension",
			modulePath: "index",
			exportKey:  "./index",
			expected:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := matchesExportPattern(tt.modulePath, tt.exportKey)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestResolveExportPattern(t *testing.T) {
	tests := []struct {
		name          string
		modulePath    string
		exportKey     string
		exportValue   string
		workspaceRoot string
		expected      string
	}{
		{
			name:          "wildcard substitution - ./* to ./dist/*",
			modulePath:    "button.js",
			exportKey:     "./*",
			exportValue:   "./dist/*",
			workspaceRoot: "/workspace",
			expected:      "/workspace/dist/button.js",
		},
		{
			name:          "wildcard substitution - ./lib/* to ./src/*",
			modulePath:    "lib/utils.js",
			exportKey:     "./lib/*",
			exportValue:   "./src/*",
			workspaceRoot: "/workspace",
			expected:      "/workspace/src/utils.js",
		},
		{
			name:          "no wildcard - exact match",
			modulePath:    "index.js",
			exportKey:     "./index.js",
			exportValue:   "./dist/index.js",
			workspaceRoot: "/workspace",
			expected:      "/workspace/dist/index.js",
		},
		{
			name:          "no wildcard - no match",
			modulePath:    "other.js",
			exportKey:     "./index.js",
			exportValue:   "./dist/index.js",
			workspaceRoot: "/workspace",
			expected:      "",
		},
		{
			name:          "wildcard key without wildcard value",
			modulePath:    "foo.js",
			exportKey:     "./*",
			exportValue:   "./dist/bundle.js",
			workspaceRoot: "/workspace",
			expected:      "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := resolveExportPattern(tt.modulePath, tt.exportKey, tt.exportValue, tt.workspaceRoot)
			assert.Equal(t, tt.expected, result)
		})
	}
}
