package sourcepos_test

import (
	"encoding/json"
	"path/filepath"
	"strings"
	"testing"

	"bennypowers.dev/cem/internal/platform/testutil"
	"bennypowers.dev/cem/internal/sourcepos"
)

func TestBuildPositionMap(t *testing.T) {
	mfs := testutil.LoadTestdataFS(t, filepath.Join("testdata"), "/")

	formats := []struct {
		file   string
		format string
	}{
		{"simple.yaml", "yaml"},
		{"simple.json", "json"},
		{"simple.jsonc", "jsonc"},
	}

	for _, tt := range formats {
		t.Run(tt.format, func(t *testing.T) {
			data, err := mfs.ReadFile(filepath.Join("/fixtures", tt.file))
			if err != nil {
				t.Fatalf("failed to read fixture: %v", err)
			}

			posMap := sourcepos.BuildPositionMap(data, tt.format)

			jsonBytes, err := json.MarshalIndent(posMap, "", "  ")
			if err != nil {
				t.Fatalf("failed to marshal position map: %v", err)
			}

			testutil.CheckGolden(t, tt.format, jsonBytes, testutil.GoldenOptions{
				Dir:         filepath.Join("testdata", "goldens"),
				Extension:   ".json",
				UseJSONDiff: true,
			})
		})
	}
}

// Inline: trivial boundary check, golden overhead unjustified.
func TestBuildPositionMap_Empty(t *testing.T) {
	posMap := sourcepos.BuildPositionMap([]byte{}, "yaml")
	if len(posMap) != 0 {
		t.Errorf("expected empty map for empty input, got %v", posMap)
	}
}

// Inline: trivial boundary check, golden overhead unjustified.
func TestBuildPositionMap_UnknownFormat(t *testing.T) {
	posMap := sourcepos.BuildPositionMap([]byte("foo: bar"), "toml")
	if len(posMap) != 0 {
		t.Errorf("expected empty map for unknown format, got %v", posMap)
	}
}

// Inline: pure function with synthetic map input, golden adds no value.
func TestResolve_ExactMatch(t *testing.T) {
	posMap := map[string]sourcepos.Position{
		"/serve/port": {Line: 3, Column: 9},
	}
	pos, ok := sourcepos.Resolve(posMap, "/serve/port")
	if !ok {
		t.Fatal("expected match")
	}
	if pos.Line != 3 || pos.Column != 9 {
		t.Errorf("got %+v, want {3 9}", pos)
	}
}

// Inline: pure function with synthetic map input, golden adds no value.
func TestResolve_ParentFallback(t *testing.T) {
	posMap := map[string]sourcepos.Position{
		"/serve":       {Line: 1, Column: 1},
		"/serve/demos": {Line: 3, Column: 3},
	}
	pos, ok := sourcepos.Resolve(posMap, "/serve/demos/rendering")
	if !ok {
		t.Fatal("expected parent fallback match")
	}
	if pos.Line != 3 || pos.Column != 3 {
		t.Errorf("got %+v, want {3 3}", pos)
	}
}

// Inline: pure function with synthetic map input, golden adds no value.
func TestResolve_NoMatch(t *testing.T) {
	posMap := map[string]sourcepos.Position{}
	_, ok := sourcepos.Resolve(posMap, "/nonexistent")
	if ok {
		t.Error("expected no match")
	}
}

// Inline: table-driven pure function test, golden overhead unjustified.
func TestFieldToJSONPointer(t *testing.T) {
	tests := []struct {
		field string
		want  string
	}{
		{"serve.port", "/serve/port"},
		{"serve.demos.rendering", "/serve/demos/rendering"},
		{"(root)", ""},
		{"generate.files", "/generate/files"},
	}
	for _, tt := range tests {
		t.Run(tt.field, func(t *testing.T) {
			got := sourcepos.FieldToJSONPointer(tt.field)
			if got != tt.want {
				t.Errorf("FieldToJSONPointer(%q) = %q, want %q", tt.field, got, tt.want)
			}
		})
	}
}

// Inline: testing RFC 6901 escaping edge case, no fixture needed.
func TestFieldToJSONPointer_RFC6901Escaping(t *testing.T) {
	got := sourcepos.FieldToJSONPointer("a~b.c/d")
	want := "/a~0b/c~1d"
	if got != want {
		t.Errorf("FieldToJSONPointer(\"a~b.c/d\") = %q, want %q", got, want)
	}
}

// Inline: roundtrip check with synthetic map, golden adds no value.
func TestFieldToJSONPointer_Roundtrip(t *testing.T) {
	posMap := map[string]sourcepos.Position{
		"/serve/port": {Line: 2, Column: 9},
	}
	pointer := sourcepos.FieldToJSONPointer("serve.port")
	pos, ok := sourcepos.Resolve(posMap, pointer)
	if !ok {
		t.Fatal("roundtrip: FieldToJSONPointer -> Resolve failed")
	}
	if pos.Line != 2 {
		t.Errorf("roundtrip got line %d, want 2", pos.Line)
	}
}

// Inline: pure function with synthetic map input, golden adds no value.
func TestResolve_DeepParentFallback(t *testing.T) {
	posMap := map[string]sourcepos.Position{
		"/generate": {Line: 5, Column: 1},
	}
	pos, ok := sourcepos.Resolve(posMap, "/generate/demoDiscovery/urlPattern")
	if !ok {
		t.Fatal("expected deep parent fallback")
	}
	if pos.Line != 5 {
		t.Errorf("got line %d, want 5", pos.Line)
	}
}

// Inline: verifying specific line numbers from synthetic YAML, golden adds no value.
func TestBuildPositionMap_ArrayElements(t *testing.T) {
	yaml := strings.Join([]string{
		"generate:",
		"  files:",
		"    - \"src/**/*.ts\"",
		"    - \"src/**/*.js\"",
	}, "\n")
	posMap := sourcepos.BuildPositionMap([]byte(yaml), "yaml")

	if pos, ok := posMap["/generate/files/0"]; !ok {
		t.Error("expected /generate/files/0 in map")
	} else if pos.Line != 3 {
		t.Errorf("/generate/files/0 line = %d, want 3", pos.Line)
	}

	if pos, ok := posMap["/generate/files/1"]; !ok {
		t.Error("expected /generate/files/1 in map")
	} else if pos.Line != 4 {
		t.Errorf("/generate/files/1 line = %d, want 4", pos.Line)
	}
}
