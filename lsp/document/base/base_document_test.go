package base

import (
	"sync"
	"testing"
	"unicode/utf8"

	"github.com/stretchr/testify/assert"
	protocol "github.com/tliron/glsp/protocol_3_16"
	ts "github.com/tree-sitter/go-tree-sitter"
	tsHtml "github.com/tree-sitter/tree-sitter-html/bindings/go"
)

func parseHTML(content string) *ts.Tree {
	parser := ts.NewParser()
	defer parser.Close()
	parser.SetLanguage(ts.NewLanguage(tsHtml.Language()))
	return parser.Parse([]byte(content), nil)
}

func TestRefCountedTree_ReleaseFreesTree(t *testing.T) {
	tree := parseHTML("<div>hello</div>")
	rc := newRefCountedTree(tree)
	rc.release()
	// tree.Close() was called internally; no double-free means success
}

func TestRefCountedTree_AcquireKeepsTreeAlive(t *testing.T) {
	tree := parseHTML("<div>hello</div>")
	rc := newRefCountedTree(tree)

	rc.acquire()
	rc.release() // owner releases, but reader still holds
	// tree should still be usable
	assert.NotNil(t, tree.RootNode())

	rc.release() // reader releases, tree is now freed
}

func TestRefCountedTree_UnderflowPanics(t *testing.T) {
	tree := parseHTML("<div>hello</div>")
	rc := newRefCountedTree(tree)
	rc.release() // legitimate release, refs=0, tree closed

	assert.Panics(t, func() {
		rc.release() // underflow
	}, "double release should panic")
}

func TestRefCountedTree_ConcurrentAcquireRelease(t *testing.T) {
	tree := parseHTML("<p>concurrent</p>")
	rc := newRefCountedTree(tree)

	var wg sync.WaitGroup
	for range 50 {
		rc.acquire()
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = tree.RootNode().Kind()
			rc.release()
		}()
	}
	wg.Wait()
	rc.release() // owner release, all readers done, tree freed
}

func TestAcquireTree_NilTreeReturnsNoopRelease(t *testing.T) {
	doc := NewBaseDocument("test.html", "", 1, "html", nil, nil)
	tree, release := doc.AcquireTree()
	assert.Nil(t, tree)
	assert.NotPanics(t, release)
}

func TestAcquireTree_PreventsUseAfterSetTree(t *testing.T) {
	doc := NewBaseDocument("test.html", "<div>v1</div>", 1, "html", nil, nil)
	tree1 := parseHTML("<div>v1</div>")
	doc.SetTree(tree1)

	acquired, release := doc.AcquireTree()
	assert.NotNil(t, acquired)

	tree2 := parseHTML("<div>v2</div>")
	doc.SetTree(tree2) // replaces tree1, but reader still holds ref

	// tree1 should still be usable via the acquired reference
	assert.Equal(t, "document", acquired.RootNode().Kind())

	release() // now tree1 can be freed
	doc.Close()
}

func TestSetParser_CallbackFiresWithOldParser(t *testing.T) {
	var returned *ts.Parser
	callback := func(p *ts.Parser) { returned = p }

	doc := NewBaseDocument("test.html", "", 1, "html", callback, nil)
	oldParser := ts.NewParser()
	t.Cleanup(oldParser.Close)
	newParser := ts.NewParser()
	t.Cleanup(newParser.Close)

	doc.SetParser(oldParser)
	doc.SetParser(newParser)

	assert.Same(t, oldParser, returned, "callback should receive the old parser")
}

func TestSetParser_SameParserDoesNotFireCallback(t *testing.T) {
	called := false
	callback := func(p *ts.Parser) { called = true }

	doc := NewBaseDocument("test.html", "", 1, "html", callback, nil)
	parser := ts.NewParser()
	t.Cleanup(parser.Close)

	doc.SetParser(parser)
	doc.SetParser(parser)

	assert.False(t, called, "callback should not fire when setting the same parser")
}

func TestSetParser_NoPreviousParserDoesNotFireCallback(t *testing.T) {
	called := false
	callback := func(p *ts.Parser) { called = true }

	doc := NewBaseDocument("test.html", "", 1, "html", callback, nil)
	parser := ts.NewParser()
	t.Cleanup(parser.Close)

	doc.SetParser(parser)

	assert.False(t, called, "callback should not fire when no previous parser existed")
}

func TestClose_FiresCallbackWithParser(t *testing.T) {
	var returned *ts.Parser
	callback := func(p *ts.Parser) {
		returned = p
	}

	doc := NewBaseDocument("test.html", "", 1, "html", callback, nil)
	parser := ts.NewParser()
	t.Cleanup(func() {
		if returned != nil {
			returned.Close()
		}
	})

	doc.SetParser(parser)
	doc.Close()

	assert.Same(t, parser, returned, "Close should return parser via callback")
	assert.Nil(t, doc.Parser(), "parser should be nil after Close")
}

func TestClose_NoParserDoesNotFireCallback(t *testing.T) {
	called := false
	callback := func(p *ts.Parser) { called = true }

	doc := NewBaseDocument("test.html", "", 1, "html", callback, nil)
	doc.Close()

	assert.False(t, called, "callback should not fire when no parser exists")
}

func TestClose_NilCallbackDoesNotPanic(t *testing.T) {
	doc := NewBaseDocument("test.html", "", 1, "html", nil, nil)
	parser := ts.NewParser()
	t.Cleanup(parser.Close)

	doc.SetParser(parser)

	assert.NotPanics(t, func() {
		doc.Close()
	}, "Close with nil callback should not panic")
}

func TestTreeAndContent_ReturnsAtomicSnapshot(t *testing.T) {
	doc := NewBaseDocument("test.html", "<div>hello</div>", 1, "html", nil, nil)

	tree, content, release := doc.TreeAndContent()
	defer release()

	assert.Nil(t, tree, "tree should be nil when not set")
	assert.Equal(t, "<div>hello</div>", content)
}

func TestByteOffsetToPosition_UTF16(t *testing.T) {
	doc := NewBaseDocument("test.html", "", 1, "html", nil, nil)
	tests := []struct {
		name     string
		content  string
		offset   uint
		expected protocol.Position
	}{
		{
			name:     "ASCII",
			content:  "hello world",
			offset:   6,
			expected: protocol.Position{Line: 0, Character: 6},
		},
		{
			name:     "emoji counts as 2 UTF-16 code units",
			content:  "😀 hello",
			offset:   4, // byte offset after emoji
			expected: protocol.Position{Line: 0, Character: 2},
		},
		{
			name:     "after emoji and space",
			content:  "😀 hello",
			offset:   5, // byte offset after emoji + space
			expected: protocol.Position{Line: 0, Character: 3},
		},
		{
			name:     "Chinese characters (3 bytes, 1 UTF-16 unit each)",
			content:  "你好world",
			offset:   6, // after 2 Chinese chars
			expected: protocol.Position{Line: 0, Character: 2},
		},
		{
			name:     "multiline with emoji",
			content:  "line1\n😀 hello",
			offset:   10, // "line1\n" (6) + emoji (4)
			expected: protocol.Position{Line: 1, Character: 2},
		},
		{
			name:     "multiline ASCII",
			content:  "line1\nline2\nline3",
			offset:   14, // "line1\nline2\nli"
			expected: protocol.Position{Line: 2, Character: 2},
		},
		{
			name:    "offset at start",
			content: "hello",
			offset:  0,
			expected: protocol.Position{Line: 0, Character: 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pos := doc.ByteOffsetToPosition(tt.offset, tt.content)
			assert.Equal(t, tt.expected, pos)
		})
	}
}

func TestPositionToByteOffset_UTF16(t *testing.T) {
	doc := NewBaseDocument("test.html", "", 1, "html", nil, nil)
	tests := []struct {
		name     string
		content  string
		pos      protocol.Position
		expected uint
	}{
		{
			name:     "ASCII",
			content:  "hello world",
			pos:      protocol.Position{Line: 0, Character: 6},
			expected: 6,
		},
		{
			name:     "emoji (2 UTF-16 units = 4 bytes)",
			content:  "😀 hello",
			pos:      protocol.Position{Line: 0, Character: 2}, // after emoji
			expected: 4,
		},
		{
			name:     "after emoji and space",
			content:  "😀 hello",
			pos:      protocol.Position{Line: 0, Character: 3}, // after emoji + space
			expected: 5,
		},
		{
			name:     "Chinese characters",
			content:  "你好world",
			pos:      protocol.Position{Line: 0, Character: 2}, // after 2 Chinese chars
			expected: 6,
		},
		{
			name:     "multiline with emoji",
			content:  "line1\n😀 hello",
			pos:      protocol.Position{Line: 1, Character: 2}, // after emoji on line 2
			expected: 10,
		},
		{
			name:     "position at start",
			content:  "hello",
			pos:      protocol.Position{Line: 0, Character: 0},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			off := doc.PositionToByteOffset(tt.pos, tt.content)
			assert.Equal(t, tt.expected, off)
		})
	}
}

func TestPositionRoundTrip(t *testing.T) {
	doc := NewBaseDocument("test.html", "", 1, "html", nil, nil)
	contents := []string{
		"hello world",
		"😀 hello 🌍 world",
		"你好世界",
		"line1\n😀 hello\n你好",
	}

	for _, content := range contents {
		t.Run(content, func(t *testing.T) {
			// Only test rune-aligned byte offsets; mid-rune offsets snap forward.
			for offset := 0; offset < len(content); {
				pos := doc.ByteOffsetToPosition(uint(offset), content)
				back := doc.PositionToByteOffset(pos, content)
				if back != uint(offset) {
					t.Errorf("round-trip failed: byte %d -> pos{%d,%d} -> byte %d",
						offset, pos.Line, pos.Character, back)
				}
				_, size := utf8.DecodeRuneInString(content[offset:])
				offset += size
			}
		})
	}
}
