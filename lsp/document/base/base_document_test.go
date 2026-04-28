package base

import (
	"testing"

	"github.com/stretchr/testify/assert"
	ts "github.com/tree-sitter/go-tree-sitter"
)

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

	tree, content := doc.TreeAndContent()

	assert.Nil(t, tree, "tree should be nil when not set")
	assert.Equal(t, "<div>hello</div>", content)
}
