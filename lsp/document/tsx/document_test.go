package tsx

import (
	"testing"

	"bennypowers.dev/cem/internal/platform/testutil"
	Q "bennypowers.dev/cem/internal/treesitter"
	"bennypowers.dev/cem/lsp/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Inline: pure function, table-driven
// extractCustomElementFromText, extractAttributeNameFromText, and scoreMatch
// are pure functions. FindCustomElements tests use fixtures but validate
// structured output with scalar assertions (tag names, counts).

func TestExtractCustomElementFromText(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		expected string
	}{
		{
			name:     "simple custom element",
			text:     "<my-element",
			expected: "my-element",
		},
		{
			name:     "custom element with attributes",
			text:     "<my-element disabled",
			expected: "my-element",
		},
		{
			name:     "standard HTML element",
			text:     "<div",
			expected: "",
		},
		{
			name:     "namespaced custom element",
			text:     "<my-ns-element",
			expected: "my-ns-element",
		},
		{
			name:     "empty string",
			text:     "",
			expected: "",
		},
		{
			name:     "whitespace only",
			text:     "   ",
			expected: "",
		},
		{
			name:     "no angle bracket",
			text:     "my-element",
			expected: "",
		},
		{
			name:     "custom element with closing bracket",
			text:     "<my-element>",
			expected: "my-element",
		},
		{
			name:     "multiple angle brackets picks last",
			text:     "<div><my-element",
			expected: "my-element",
		},
		{
			name:     "standard element after custom element",
			text:     "<my-element><div",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractCustomElementFromText(tt.text)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestExtractAttributeNameFromText(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		expected string
	}{
		{
			name:     "simple attribute",
			text:     "disabled",
			expected: "disabled",
		},
		{
			name:     "hyphenated attribute",
			text:     "aria-label",
			expected: "aria-label",
		},
		{
			name:     "lit event binding",
			text:     "@click",
			expected: "@click",
		},
		{
			name:     "lit property binding",
			text:     ".prop",
			expected: ".prop",
		},
		{
			name:     "empty string",
			text:     "",
			expected: "",
		},
		{
			name:     "attribute with equals sign",
			text:     "disabled=",
			expected: "disabled",
		},
		{
			name:     "text with spaces returns last word",
			text:     "<my-element disabled",
			expected: "disabled",
		},
		{
			name:     "text with angle brackets only",
			text:     "<div>",
			expected: "",
		},
		{
			name:     "attribute with value",
			text:     "class= foo",
			expected: "foo",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractAttributeNameFromText(tt.text)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestScoreMatch(t *testing.T) {
	tests := []struct {
		name        string
		captureType string
		captureMap  Q.CaptureMap
		capture     Q.CaptureInfo
		byteOffset  uint
		expected    int
	}{
		{
			name:        "tag name base score",
			captureType: "tag.name.completion",
			captureMap:  Q.CaptureMap{"a": nil, "b": nil},
			capture:     Q.CaptureInfo{StartByte: 100, EndByte: 105},
			byteOffset:  200,
			expected:    SCORE_TAG_NAME_BASE,
		},
		{
			name:        "attr name base score",
			captureType: "attr.name.completion",
			captureMap:  Q.CaptureMap{"a": nil, "b": nil},
			capture:     Q.CaptureInfo{StartByte: 100, EndByte: 105},
			byteOffset:  200,
			expected:    SCORE_ATTR_NAME_BASE,
		},
		{
			name:        "attr value base score",
			captureType: "attr.value.completion",
			captureMap:  Q.CaptureMap{"a": nil, "b": nil},
			capture:     Q.CaptureInfo{StartByte: 100, EndByte: 105},
			byteOffset:  200,
			expected:    SCORE_ATTR_VALUE_BASE,
		},
		{
			name:        "cursor inside capture gets bonus",
			captureType: "tag.name.completion",
			captureMap:  Q.CaptureMap{"a": nil, "b": nil},
			capture:     Q.CaptureInfo{StartByte: 10, EndByte: 20},
			byteOffset:  15,
			expected:    SCORE_TAG_NAME_BASE - SCORE_CURSOR_BONUS,
		},
		{
			name:        "cursor at start of capture gets bonus",
			captureType: "tag.name.completion",
			captureMap:  Q.CaptureMap{"a": nil, "b": nil},
			capture:     Q.CaptureInfo{StartByte: 10, EndByte: 20},
			byteOffset:  10,
			expected:    SCORE_TAG_NAME_BASE - SCORE_CURSOR_BONUS,
		},
		{
			name:        "cursor at end of capture gets bonus",
			captureType: "tag.name.completion",
			captureMap:  Q.CaptureMap{"a": nil, "b": nil},
			capture:     Q.CaptureInfo{StartByte: 10, EndByte: 20},
			byteOffset:  20,
			expected:    SCORE_TAG_NAME_BASE - SCORE_CURSOR_BONUS,
		},
		{
			name:        "broad capture map penalty",
			captureType: "tag.name.completion",
			captureMap:  Q.CaptureMap{"a": nil, "b": nil, "c": nil},
			capture:     Q.CaptureInfo{StartByte: 100, EndByte: 105},
			byteOffset:  200,
			expected:    SCORE_TAG_NAME_BASE + 3*SCORE_BROAD_CAPTURE_PENALTY,
		},
		{
			name:        "long capture length penalty",
			captureType: "tag.name.completion",
			captureMap:  Q.CaptureMap{"a": nil, "b": nil},
			capture:     Q.CaptureInfo{StartByte: 0, EndByte: 50},
			byteOffset:  100,
			expected:    SCORE_TAG_NAME_BASE + 50,
		},
		{
			name:        "unknown capture type gets zero base",
			captureType: "unknown",
			captureMap:  Q.CaptureMap{"a": nil, "b": nil},
			capture:     Q.CaptureInfo{StartByte: 100, EndByte: 105},
			byteOffset:  200,
			expected:    0,
		},
		{
			name:        "all bonuses and penalties combined",
			captureType: "attr.name.completion",
			captureMap:  Q.CaptureMap{"a": nil, "b": nil, "c": nil, "d": nil},
			capture:     Q.CaptureInfo{StartByte: 0, EndByte: 20},
			byteOffset:  10,
			// base(20) - cursor_bonus(50) + broad_penalty(4*5) + length_penalty(20)
			expected: SCORE_ATTR_NAME_BASE - SCORE_CURSOR_BONUS + 4*SCORE_BROAD_CAPTURE_PENALTY + 20,
		},
		{
			name:        "capture length at threshold is not penalized",
			captureType: "tag.name.completion",
			captureMap:  Q.CaptureMap{"a": nil, "b": nil},
			capture:     Q.CaptureInfo{StartByte: 0, EndByte: MAX_REASONABLE_TAG_LENGTH},
			byteOffset:  100,
			expected:    SCORE_TAG_NAME_BASE,
		},
		{
			name:        "capture length just over threshold is penalized",
			captureType: "tag.name.completion",
			captureMap:  Q.CaptureMap{"a": nil, "b": nil},
			capture:     Q.CaptureInfo{StartByte: 0, EndByte: MAX_REASONABLE_TAG_LENGTH + 1},
			byteOffset:  100,
			expected:    SCORE_TAG_NAME_BASE + int(MAX_REASONABLE_TAG_LENGTH+1),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := scoreMatch(tt.captureType, tt.captureMap, tt.capture, tt.byteOffset)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func newTSXHandler(t *testing.T) *Handler {
	t.Helper()
	qm, err := Q.NewQueryManager(Q.LSPQueries())
	require.NoError(t, err)
	t.Cleanup(func() { qm.Close() })

	handler, err := NewHandler(qm)
	require.NoError(t, err)
	t.Cleanup(func() { handler.Close() })

	return handler
}

func TestFindCustomElements_TSX(t *testing.T) {
	handler := newTSXHandler(t)
	mfs := testutil.LoadTestdataFS(t, "testdata", "/")

	tests := []struct {
		name     string
		fixture  string
		validate func(t *testing.T, elements []types.CustomElementMatch)
	}{
		{
			name:    "tsx with custom elements",
			fixture: "/with-custom-elements.tsx",
			validate: func(t *testing.T, elements []types.CustomElementMatch) {
				require.GreaterOrEqual(t, len(elements), 2)
				tagNames := make([]string, len(elements))
				for i, e := range elements {
					tagNames[i] = e.TagName
				}
				assert.Contains(t, tagNames, "my-element")
				assert.Contains(t, tagNames, "other-element")
			},
		},
		{
			name:    "tsx no custom elements",
			fixture: "/no-custom-elements.tsx",
			validate: func(t *testing.T, elements []types.CustomElementMatch) {
				assert.Empty(t, elements)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			content := testutil.ReadFixture(t, mfs, tt.fixture)
			doc := handler.CreateDocument("file:///test.tsx", string(content), 1)
			defer doc.Close()
			elements, err := handler.FindCustomElements(doc)
			require.NoError(t, err)
			tt.validate(t, elements)
		})
	}
}
