package config_test

import (
	"testing"

	"bennypowers.dev/cem/internal/config"
)

func TestStripComments_LineComments(t *testing.T) {
	input := []byte(`{
  // this is a comment
  "key": "value"
}`)
	got := config.StripComments(input)
	// Should produce valid JSON
	if len(got) == 0 {
		t.Fatal("StripComments returned empty")
	}
	// Should not contain the comment
	for _, b := range got {
		if b == '/' {
			// Check if it's part of a comment
			s := string(got)
			if s != "" && contains(s, "// this is") {
				t.Error("StripComments did not remove line comment")
			}
		}
	}
}

func TestStripComments_BlockComments(t *testing.T) {
	input := []byte(`{
  /* block comment */
  "key": "value"
}`)
	got := config.StripComments(input)
	if contains(string(got), "block comment") {
		t.Error("StripComments did not remove block comment")
	}
}

func TestStripComments_TrailingCommas(t *testing.T) {
	input := []byte(`{
  "a": 1,
  "b": 2,
}`)
	got := config.StripComments(input)
	// Should produce valid JSON (trailing comma removed)
	if len(got) == 0 {
		t.Fatal("StripComments returned empty")
	}
	// Verify it's valid JSON by checking it doesn't end with ,}
	s := string(got)
	if contains(s, ",\n}") || contains(s, ",}") {
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

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && stringContains(s, substr))
}

func stringContains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
