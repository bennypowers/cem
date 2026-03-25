package config_test

import (
	"strings"
	"testing"

	"bennypowers.dev/cem/internal/config"
)

func TestStripComments_LineComments(t *testing.T) {
	input := []byte(`{
  // this is a comment
  "key": "value"
}`)
	got := config.StripComments(input)
	if len(got) == 0 {
		t.Fatal("StripComments returned empty")
	}
	if strings.Contains(string(got), "// this is") {
		t.Error("StripComments did not remove line comment")
	}
}

func TestStripComments_BlockComments(t *testing.T) {
	input := []byte(`{
  /* block comment */
  "key": "value"
}`)
	got := config.StripComments(input)
	if strings.Contains(string(got), "block comment") {
		t.Error("StripComments did not remove block comment")
	}
}

func TestStripComments_TrailingCommas(t *testing.T) {
	input := []byte(`{
  "a": 1,
  "b": 2,
}`)
	got := config.StripComments(input)
	if len(got) == 0 {
		t.Fatal("StripComments returned empty")
	}
	s := string(got)
	if strings.Contains(s, ",\n}") || strings.Contains(s, ",}") {
		t.Error("StripComments did not remove trailing comma")
	}
}

func TestStripComments_NoComments(t *testing.T) {
	input := []byte(`{"key": "value"}`)
	got := config.StripComments(input)
	if string(got) != string(input) {
		t.Errorf("StripComments modified comment-free JSON: got %q", string(got))
	}
}

func TestStripComments_EmptyInput(t *testing.T) {
	got := config.StripComments([]byte{})
	if len(got) != 0 {
		t.Errorf("StripComments of empty input returned %d bytes", len(got))
	}
}
