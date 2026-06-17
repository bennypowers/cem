package cmd

import (
	"testing"

	IC "bennypowers.dev/cem/internal/config"
)

// Inline assertions justified: testing unexported cmd wiring functions,
// pure logic with no I/O dependencies.

func TestDeduplicateErrors_Empty(t *testing.T) {
	result := deduplicateErrors(nil)
	if len(result) != 0 {
		t.Errorf("expected empty, got %v", result)
	}
}

func TestDeduplicateErrors_NoDuplicates(t *testing.T) {
	errs := []IC.ValidationError{
		{Field: "a", Message: "error a"},
		{Field: "b", Message: "error b"},
	}
	result := deduplicateErrors(errs)
	if len(result) != 2 {
		t.Fatalf("expected 2, got %d", len(result))
	}
	if result[0].Field != "a" || result[1].Field != "b" {
		t.Errorf("unexpected order: %v", result)
	}
}

func TestDeduplicateErrors_LastWins(t *testing.T) {
	errs := []IC.ValidationError{
		{Field: "serve.port", Message: "schema error"},
		{Field: "serve.port", Message: "semantic error", Value: "99999"},
	}
	result := deduplicateErrors(errs)
	if len(result) != 1 {
		t.Fatalf("expected 1 after dedup, got %d", len(result))
	}
	if result[0].Message != "semantic error" {
		t.Errorf("expected semantic error to win, got %q", result[0].Message)
	}
	if result[0].Value != "99999" {
		t.Errorf("expected value preserved, got %q", result[0].Value)
	}
}

func TestDeduplicateErrors_PreservesOrder(t *testing.T) {
	errs := []IC.ValidationError{
		{Field: "a", Message: "first"},
		{Field: "b", Message: "second"},
		{Field: "a", Message: "replaced"},
	}
	result := deduplicateErrors(errs)
	if len(result) != 2 {
		t.Fatalf("expected 2, got %d", len(result))
	}
	if result[0].Field != "a" || result[0].Message != "replaced" {
		t.Errorf("expected 'a' replaced at index 0, got %v", result[0])
	}
	if result[1].Field != "b" {
		t.Errorf("expected 'b' at index 1, got %v", result[1])
	}
}

func TestEnrichPositions_YAMLFormat(t *testing.T) {
	yaml := []byte("serve:\n  port: 3000\n")
	errs := []IC.ValidationError{
		{Field: "serve.port"},
	}
	enrichPositions(errs, yaml, "yaml")
	if errs[0].Line != 2 || errs[0].Column != 9 {
		t.Errorf("expected line 2 col 9, got line %d col %d", errs[0].Line, errs[0].Column)
	}
}

func TestEnrichPositions_NoMatch(t *testing.T) {
	yaml := []byte("serve:\n  port: 3000\n")
	errs := []IC.ValidationError{
		{Field: "nonexistent.field"},
	}
	enrichPositions(errs, yaml, "yaml")
	if errs[0].Line != 0 || errs[0].Column != 0 {
		t.Errorf("expected 0/0 for no match, got %d/%d", errs[0].Line, errs[0].Column)
	}
}

func TestEnrichPositions_ParentFallback(t *testing.T) {
	yaml := []byte("generate:\n  demoDiscovery:\n    fileGlob: \"*.html\"\n")
	errs := []IC.ValidationError{
		{Field: "generate.demoDiscovery"},
	}
	enrichPositions(errs, yaml, "yaml")
	if errs[0].Line == 0 {
		t.Error("expected parent fallback to find position")
	}
}
