package typescript

import (
	"sort"
	"testing"

	"bennypowers.dev/cem/internal/platform/testutil"
	Q "bennypowers.dev/cem/internal/treesitter"
	"bennypowers.dev/cem/lsp/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Inline: pure function, table-driven
// getCaptureMapKeys is a pure function. FindCustomElements tests use fixtures
// but validate structured output with scalar assertions (tag names).

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
	mfs := testutil.LoadTestdataFS(t, "testdata", "/")

	tests := []struct {
		name     string
		fixture  string
		validate func(t *testing.T, elements []types.CustomElementMatch)
	}{
		{
			name:    "lit element with html template",
			fixture: "/lit-element.ts",
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
			fixture: "/no-templates.ts",
			validate: func(t *testing.T, elements []types.CustomElementMatch) {
				assert.Empty(t, elements)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			content := testutil.ReadFixture(t, mfs, tt.fixture)
			doc := handler.CreateDocument("file:///test.ts", string(content), 1)
			defer doc.Close()
			elements, err := handler.FindCustomElements(doc)
			require.NoError(t, err)
			tt.validate(t, elements)
		})
	}
}
