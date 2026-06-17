package cmd

import (
	"testing"

	IC "bennypowers.dev/cem/internal/config"
)

// Inline assertions justified: testing unexported cmd wiring functions,
// pure logic with no I/O dependencies.

func TestDeduplicateErrors_BothEmpty(t *testing.T) {
	result := deduplicateErrors(nil, nil)
	if len(result) != 0 {
		t.Errorf("expected empty, got %v", result)
	}
}

func TestDeduplicateErrors_SchemaOnly(t *testing.T) {
	schema := []IC.ValidationError{
		{Field: "a", Message: "type error"},
		{Field: "b", Message: "enum error"},
	}
	result := deduplicateErrors(schema, nil)
	if len(result) != 2 {
		t.Fatalf("expected 2, got %d", len(result))
	}
}

func TestDeduplicateErrors_SemanticReplacesSchema(t *testing.T) {
	schema := []IC.ValidationError{
		{Field: "serve.port", Message: "maximum: got 99999, want 65535", Severity: IC.SeverityError},
	}
	semantic := []IC.ValidationError{
		{Field: "serve.port", Message: "must be between 0 and 65535", Value: "99999", Severity: IC.SeverityError},
	}
	result := deduplicateErrors(schema, semantic)
	if len(result) != 1 {
		t.Fatalf("expected 1 after dedup, got %d", len(result))
	}
	if result[0].Value != "99999" {
		t.Errorf("expected semantic error to win, got %q", result[0].Message)
	}
}

func TestDeduplicateErrors_PreservesMultipleSchemaOnSameField(t *testing.T) {
	schema := []IC.ValidationError{
		{Field: "serve.port", Message: "type error", Severity: IC.SeverityError},
		{Field: "serve.port", Message: "range warning", Severity: IC.SeverityWarning},
	}
	result := deduplicateErrors(schema, nil)
	if len(result) != 2 {
		t.Fatalf("expected 2 (different severities preserved), got %d", len(result))
	}
}

func TestDeduplicateErrors_PreservesDifferentSeverities(t *testing.T) {
	schema := []IC.ValidationError{
		{Field: "serve.port", Message: "schema error", Severity: IC.SeverityError},
	}
	semantic := []IC.ValidationError{
		{Field: "serve.port", Message: "semantic warning", Severity: IC.SeverityWarning},
	}
	result := deduplicateErrors(schema, semantic)
	if len(result) != 2 {
		t.Fatalf("expected 2 (error + warning preserved), got %d", len(result))
	}
}

func TestDeduplicateErrors_PreservesRootErrors(t *testing.T) {
	schema := []IC.ValidationError{
		{Field: "(root)", Message: "parse error"},
	}
	semantic := []IC.ValidationError{
		{Field: "(root)", Message: "another root error"},
	}
	result := deduplicateErrors(schema, semantic)
	if len(result) != 1 {
		t.Fatalf("expected 1 (same field+severity deduped), got %d", len(result))
	}
}

func TestDeduplicateErrors_NormalizesEmptySeverity(t *testing.T) {
	schema := []IC.ValidationError{
		{Field: "serve.port", Message: "maximum: got 99999", Severity: IC.SeverityError},
	}
	semantic := []IC.ValidationError{
		{Field: "serve.port", Message: "must be between 0 and 65535", Value: "99999"},
	}
	result := deduplicateErrors(schema, semantic)
	if len(result) != 1 {
		t.Fatalf("expected 1 after dedup with empty severity, got %d", len(result))
	}
	if result[0].Value != "99999" {
		t.Errorf("expected semantic error to win, got %q", result[0].Message)
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
