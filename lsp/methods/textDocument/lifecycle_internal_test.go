package textDocument

import (
	"testing"

	protocol "github.com/bennypowers/glsp/protocol_3_17"
	"github.com/stretchr/testify/assert"
)

func ptrRange(startLine, startChar, endLine, endChar uint32) *protocol.Range {
	return &protocol.Range{
		Start: protocol.Position{Line: startLine, Character: startChar},
		End:   protocol.Position{Line: endLine, Character: endChar},
	}
}

func TestApplyIncrementalChange(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		change   protocol.TextDocumentContentChangeEvent
		expected string
	}{
		{
			name:    "insert at start",
			content: "hello world",
			change: protocol.TextDocumentContentChangeEvent{
				Range: ptrRange(0, 0, 0, 0),
				Text:  "prefix ",
			},
			expected: "prefix hello world",
		},
		{
			name:    "insert at end",
			content: "hello world",
			change: protocol.TextDocumentContentChangeEvent{
				Range: ptrRange(0, 11, 0, 11),
				Text:  " suffix",
			},
			expected: "hello world suffix",
		},
		{
			name:    "insert in middle",
			content: "hello world",
			change: protocol.TextDocumentContentChangeEvent{
				Range: ptrRange(0, 5, 0, 5),
				Text:  " beautiful",
			},
			expected: "hello beautiful world",
		},
		{
			name:    "delete range",
			content: "hello world",
			change: protocol.TextDocumentContentChangeEvent{
				Range: ptrRange(0, 5, 0, 11),
				Text:  "",
			},
			expected: "hello",
		},
		{
			name:    "replace range",
			content: "hello world",
			change: protocol.TextDocumentContentChangeEvent{
				Range: ptrRange(0, 6, 0, 11),
				Text:  "there",
			},
			expected: "hello there",
		},
		{
			name:    "multi-line insert",
			content: "line1\nline2",
			change: protocol.TextDocumentContentChangeEvent{
				Range: ptrRange(0, 5, 0, 5),
				Text:  "\nnewline\n",
			},
			expected: "line1\nnewline\n\nline2",
		},
		{
			name:    "empty change text (no-op insert)",
			content: "hello world",
			change: protocol.TextDocumentContentChangeEvent{
				Range: ptrRange(0, 5, 0, 5),
				Text:  "",
			},
			expected: "hello world",
		},
		{
			name:    "nil range returns full text replacement",
			content: "hello world",
			change: protocol.TextDocumentContentChangeEvent{
				Range: nil,
				Text:  "completely new",
			},
			expected: "completely new",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := applyIncrementalChange(tt.content, tt.change)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestExtractChangeParameters(t *testing.T) {
	tests := []struct {
		name     string
		change   protocol.TextDocumentContentChangeEvent
		lines    []string
		expected changeParameters
	}{
		{
			name: "valid params",
			change: protocol.TextDocumentContentChangeEvent{
				Range: ptrRange(1, 3, 2, 5),
				Text:  "replacement",
			},
			lines: []string{"line0", "line1", "line2"},
			expected: changeParameters{
				startLine: 1,
				startChar: 3,
				endLine:   2,
				endChar:   5,
			},
		},
		{
			name: "out-of-bounds lines",
			change: protocol.TextDocumentContentChangeEvent{
				Range: ptrRange(10, 0, 20, 0),
				Text:  "text",
			},
			lines: []string{"line0"},
			expected: changeParameters{
				startLine: 10,
				startChar: 0,
				endLine:   20,
				endChar:   0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractChangeParameters(tt.change, tt.lines)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestApplySingleLineChange(t *testing.T) {
	tests := []struct {
		name     string
		lines    []string
		params   changeParameters
		text     string
		expected string
	}{
		{
			name:  "replace within line",
			lines: []string{"hello world"},
			params: changeParameters{
				startLine: 0, startChar: 6, endLine: 0, endChar: 11,
			},
			text:     "there",
			expected: "hello there",
		},
		{
			name:  "insert at start of line",
			lines: []string{"world"},
			params: changeParameters{
				startLine: 0, startChar: 0, endLine: 0, endChar: 0,
			},
			text:     "hello ",
			expected: "hello world",
		},
		{
			name:  "insert at end of line",
			lines: []string{"hello"},
			params: changeParameters{
				startLine: 0, startChar: 5, endLine: 0, endChar: 5,
			},
			text:     " world",
			expected: "hello world",
		},
		{
			name:  "delete partial",
			lines: []string{"hello world"},
			params: changeParameters{
				startLine: 0, startChar: 5, endLine: 0, endChar: 11,
			},
			text:     "",
			expected: "hello",
		},
		{
			name:  "multi-line content with change on second line",
			lines: []string{"first", "second", "third"},
			params: changeParameters{
				startLine: 1, startChar: 0, endLine: 1, endChar: 6,
			},
			text:     "SECOND",
			expected: "first\nSECOND\nthird",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := applySingleLineChange(tt.lines, tt.params, tt.text)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestApplyMultiLineChange(t *testing.T) {
	tests := []struct {
		name     string
		lines    []string
		params   changeParameters
		text     string
		expected string
	}{
		{
			name:  "replace spanning two lines",
			lines: []string{"hello", "world", "end"},
			params: changeParameters{
				startLine: 0, startChar: 3, endLine: 1, endChar: 3,
			},
			text:     "LO WO",
			expected: "helLO WOld\nend",
		},
		{
			name:  "delete spanning two lines",
			lines: []string{"hello", "world", "end"},
			params: changeParameters{
				startLine: 0, startChar: 5, endLine: 1, endChar: 5,
			},
			text:     "",
			expected: "hello\nend",
		},
		{
			name:  "insert expanding across lines",
			lines: []string{"hello", "world"},
			params: changeParameters{
				startLine: 0, startChar: 5, endLine: 1, endChar: 0,
			},
			text:     "\nnew line\n",
			expected: "hello\nnew line\nworld",
		},
		{
			name:  "replace spanning all lines",
			lines: []string{"aaa", "bbb", "ccc"},
			params: changeParameters{
				startLine: 0, startChar: 0, endLine: 2, endChar: 3,
			},
			text:     "replaced",
			expected: "replaced",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := applyMultiLineChange(tt.lines, tt.params, tt.text)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestAdjustEndLineIfNeeded(t *testing.T) {
	tests := []struct {
		name   string
		params changeParameters
		lines  []string
		want   changeParameters
	}{
		{
			"within bounds",
			changeParameters{startLine: 0, endLine: 1, endChar: 3},
			[]string{"hello", "world"},
			changeParameters{startLine: 0, endLine: 1, endChar: 3},
		},
		{
			"beyond end",
			changeParameters{startLine: 0, endLine: 10, endChar: 5},
			[]string{"hello", "world"},
			changeParameters{startLine: 0, endLine: 1, endChar: 5},
		},
		{
			"exactly at boundary",
			changeParameters{startLine: 0, endLine: 2, endChar: 0},
			[]string{"hello", "world"},
			changeParameters{startLine: 0, endLine: 1, endChar: 5},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := adjustEndLineIfNeeded(tc.params, tc.lines)
			assert.Equal(t, tc.want.endLine, result.endLine)
			assert.Equal(t, tc.want.endChar, result.endChar)
		})
	}
}

func TestHandleOutOfBoundsChange(t *testing.T) {
	t.Run("appends to end", func(t *testing.T) {
		lines := []string{"hello", "world"}
		result := handleOutOfBoundsChange(lines, 5, "new text")
		assert.Contains(t, result, "new text")
	})

	t.Run("pads with empty lines", func(t *testing.T) {
		lines := []string{"hello"}
		result := handleOutOfBoundsChange(lines, 3, "far away")
		assert.Contains(t, result, "far away")
	})
}

func TestSplitLinesInternal(t *testing.T) {
	assert.Len(t, splitLines("a\nb\nc"), 3)
	assert.Len(t, splitLines("a\r\nb\r\nc"), 3)
	assert.Len(t, splitLines("a\rb\rc"), 3)
	assert.Len(t, splitLines("hello"), 1)
	assert.Len(t, splitLines(""), 1)
}

func TestApplySingleLineChange_ClampBounds(t *testing.T) {
	lines := []string{"hi"}
	params := changeParameters{startLine: 0, startChar: 100, endLine: 0, endChar: 200}
	result := applySingleLineChange(lines, params, "!")
	assert.Equal(t, "hi!", result)
}
