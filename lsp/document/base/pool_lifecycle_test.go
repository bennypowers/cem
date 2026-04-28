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

	assert.Equal(t, parser, returned, "Close should return parser via callback")
	assert.Nil(t, doc.Parser(), "parser should be nil after Close")
}

func TestPoolLifecycle_SetParserReturnsOldParserViaCallback(t *testing.T) {
	var returned *ts.Parser
	callback := func(p *ts.Parser) { returned = p }

	doc := NewBaseDocument("test.ts", "", 1, "typescript", callback, nil)
	old := ts.NewParser()
	defer old.Close()
	new := ts.NewParser()
	defer new.Close()

	doc.SetParser(old)
	doc.SetParser(new)

	assert.Equal(t, old, returned, "SetParser should return old parser via callback")
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
	// Move parser from HTML doc to TS doc
	tsDoc.SetParser(parser)

	// Close TS doc -- should return via TS callback, not HTML callback
	tsDoc.Close()

	assert.Nil(t, htmlReturned, "HTML callback should not be called when TS doc closes")
	assert.Equal(t, parser, tsReturned, "TS callback should receive the parser")
}

func TestPoolLifecycle_DoubleCloseDoesNotPanic(t *testing.T) {
	callback := func(p *ts.Parser) {}
	doc := NewBaseDocument("test.html", "", 1, "html", callback, nil)
	parser := ts.NewParser()
	doc.SetParser(parser)

	assert.NotPanics(t, func() {
		doc.Close()
		doc.Close()
	}, "double Close should not panic")
}
