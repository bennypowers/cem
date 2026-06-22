package tui

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Inline: output depends on terminal state and chroma internals; golden files would be brittle.
func TestHighlight_ValidYAML(t *testing.T) {
	var buf bytes.Buffer
	err := Highlight(&buf, "key: value\n", "yaml")
	require.NoError(t, err)
	assert.Contains(t, buf.String(), "key")
	assert.Contains(t, buf.String(), "value")
}

func TestHighlight_ValidJSON(t *testing.T) {
	var buf bytes.Buffer
	err := Highlight(&buf, `{"key": "value"}`+"\n", "json")
	require.NoError(t, err)
	assert.Contains(t, buf.String(), "key")
	assert.Contains(t, buf.String(), "value")
}

func TestHighlight_UnknownLang_FallsBack(t *testing.T) {
	var buf bytes.Buffer
	err := Highlight(&buf, "plain text\n", "nonexistent-lang-xyz")
	require.NoError(t, err)
	assert.Contains(t, buf.String(), "plain text")
}

func TestHighlight_NonTTY_NoANSI(t *testing.T) {
	var buf bytes.Buffer
	err := Highlight(&buf, "key: value\n", "yaml")
	require.NoError(t, err)
	assert.False(t, strings.Contains(buf.String(), "\x1b["),
		"non-TTY writer should not contain ANSI escape codes, got: %q", buf.String())
}

