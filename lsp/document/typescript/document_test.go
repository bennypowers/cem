package typescript

import (
	"os"
	"path/filepath"
	"sort"
	"testing"

	Q "bennypowers.dev/cem/internal/treesitter"
	"bennypowers.dev/cem/lsp/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetCaptureMapKeys(t *testing.T) {
	tests := []struct {
		name     string
		input    Q.CaptureMap
		expected []string
	}{
		{
			name: "populated map",
			input: Q.CaptureMap{
				"tag.name":   {{Text: "div"}},
				"attr.name":  {{Text: "class"}},
				"attr.value": {{Text: "foo"}},
			},
			expected: []string{"attr.name", "attr.value", "tag.name"},
		},
		{
			name:     "empty map",
			input:    Q.CaptureMap{},
			expected: []string{},
		},
		{
			name: "single key",
			input: Q.CaptureMap{
				"context": {{Text: "something"}},
			},
			expected: []string{"context"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getCaptureMapKeys(tt.input)
			sort.Strings(result)
			sort.Strings(tt.expected)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func newTSHandler(t *testing.T) *Handler {
	t.Helper()
	qm, err := Q.NewQueryManager(Q.LSPQueries())
	require.NoError(t, err)
	t.Cleanup(func() { qm.Close() })

	handler, err := NewHandler(qm)
	require.NoError(t, err)
	t.Cleanup(func() { handler.Close() })

	return handler
}

func TestFindCustomElements_TypeScript(t *testing.T) {
	handler := newTSHandler(t)

	tests := []struct {
		name     string
		fixture  string
		validate func(t *testing.T, elements []types.CustomElementMatch)
	}{
		{
			name:    "lit element with html template",
			fixture: filepath.Join("testdata", "lit-element.ts"),
			validate: func(t *testing.T, elements []types.CustomElementMatch) {
				require.GreaterOrEqual(t, len(elements), 2)
				tagNames := make([]string, len(elements))
				for i, e := range elements {
					tagNames[i] = e.TagName
				}
				assert.Contains(t, tagNames, "my-header")
				assert.Contains(t, tagNames, "my-card")
			},
		},
		{
			name:    "typescript with no html templates",
			fixture: filepath.Join("testdata", "no-templates.ts"),
			validate: func(t *testing.T, elements []types.CustomElementMatch) {
				assert.Empty(t, elements)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			content, err := os.ReadFile(tt.fixture)
			require.NoError(t, err)
			doc := handler.CreateDocument("file:///test.ts", string(content), 1)
			defer doc.Close()
			elements, err := handler.FindCustomElements(doc)
			require.NoError(t, err)
			tt.validate(t, elements)
		})
	}
}
