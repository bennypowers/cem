package base

import (
	"testing"

	"github.com/stretchr/testify/assert"
	ts "github.com/tree-sitter/go-tree-sitter"
)

func TestPoolLifecycle_CloseReturnsParserViaCallback(t *testing.T) {
	var returned *ts.Parser
	callback := func(p *ts.Parser) { returned = p }

	doc := NewBaseDocument("test.html", "<div></div>", 1, "html", callback, nil)
	parser := ts.NewParser()
	doc.SetParser(parser)

	doc.Close()

	assert.Same(t, parser, returned, "Close should return parser via callback")
	assert.Nil(t, doc.Parser(), "parser should be nil after Close")
	if returned != nil {
		returned.Close()
	}
}

func TestPoolLifecycle_SetParserReturnsOldParserViaCallback(t *testing.T) {
	var returned *ts.Parser
	callback := func(p *ts.Parser) { returned = p }

	doc := NewBaseDocument("test.ts", "", 1, "typescript", callback, nil)
	old := ts.NewParser()
	t.Cleanup(old.Close)
	new := ts.NewParser()
	t.Cleanup(new.Close)

	doc.SetParser(old)
	doc.SetParser(new)

	assert.Same(t, old, returned, "SetParser should return old parser via callback")
}

func TestPoolLifecycle_CrossPoolAssignment(t *testing.T) {
	var htmlReturned *ts.Parser
	htmlCallback := func(p *ts.Parser) { htmlReturned = p }

	var tsReturned *ts.Parser
	tsCallback := func(p *ts.Parser) { tsReturned = p }

	htmlDoc := NewBaseDocument("test.html", "", 1, "html", htmlCallback, nil)
	tsDoc := NewBaseDocument("test.ts", "", 1, "typescript", tsCallback, nil)

	parser := ts.NewParser()

	htmlDoc.SetParser(parser)
	tsDoc.SetParser(parser)

	tsDoc.Close()

	assert.Nil(t, htmlReturned, "HTML callback should not be called when TS doc closes")
	assert.Same(t, parser, tsReturned, "TS callback should receive the parser")
	if tsReturned != nil {
		tsReturned.Close()
	}
}

func TestPoolLifecycle_DoubleCloseDoesNotPanic(t *testing.T) {
	var returned *ts.Parser
	callback := func(p *ts.Parser) { returned = p }
	doc := NewBaseDocument("test.html", "", 1, "html", callback, nil)
	parser := ts.NewParser()
	doc.SetParser(parser)

	assert.NotPanics(t, func() {
		doc.Close()
		doc.Close()
	}, "double Close should not panic")
	if returned != nil {
		returned.Close()
	}
}
