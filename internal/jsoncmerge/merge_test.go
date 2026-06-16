package jsoncmerge_test

import (
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/internal/jsoncmerge"
	"bennypowers.dev/cem/internal/platform/testutil"
	"github.com/tailscale/hujson"
)

var mergeValue = map[string]any{"command": "cem", "args": []string{"mcp"}}

func TestMerge(t *testing.T) {
	fs := testutil.LoadTestdataFS(t, "testdata", "/testdata")

	tests := []struct {
		name   string
		topKey string
		subKey string
	}{
		{"insert-new-topkey", "context_servers", "cem"},
		{"insert-into-existing", "context_servers", "cem"},
		{"replace-existing-subkey", "context_servers", "cem"},
		{"key-in-comment", "context_servers", "cem"},
		{"key-in-string-value", "context_servers", "cem"},
		{"trailing-comma", "context_servers", "cem"},
		{"empty-object", "context_servers", "cem"},
		{"block-comment", "context_servers", "cem"},
		{"real-zed-settings", "context_servers", "cem"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			inputPath := filepath.Join("testdata", tc.name+".input.jsonc")
			input, err := fs.ReadFile(inputPath)
			if err != nil {
				t.Fatalf("failed to read input: %v", err)
			}

			result, err := jsoncmerge.Merge(input, tc.topKey, tc.subKey, mergeValue)
			if err != nil {
				t.Fatalf("Merge failed: %v", err)
			}

			if _, err := hujson.Parse(result); err != nil {
				t.Errorf("result is not valid JSONC: %v\n%s", err, result)
			}

			testutil.CheckGolden(t, tc.name, result, testutil.GoldenOptions{
				Dir:       filepath.Join("testdata", "goldens"),
				Extension: ".golden.jsonc",
				FS:        fs,
			})
		})
	}
}

// Inline assertion: testing error return only, no output to golden.
func TestMerge_InvalidInput(t *testing.T) {
	_, err := jsoncmerge.Merge([]byte("not json"), "key", "sub", "val")
	if err == nil {
		t.Error("expected error for invalid input")
	}
}

// Inline assertion: testing error return only, no output to golden.
func TestMerge_NonObjectRoot(t *testing.T) {
	_, err := jsoncmerge.Merge([]byte(`["array"]`), "key", "sub", "val")
	if err == nil {
		t.Error("expected error for array root")
	}
}

// Inline assertion: comparing two merge outputs for equality, not a single golden shape.
func TestMerge_Idempotent(t *testing.T) {
	input := []byte(`{
  "context_servers": {
    "cem": {"command": "old"}
  }
}
`)
	first, err := jsoncmerge.Merge(input, "context_servers", "cem", mergeValue)
	if err != nil {
		t.Fatalf("first merge: %v", err)
	}
	second, err := jsoncmerge.Merge(first, "context_servers", "cem", mergeValue)
	if err != nil {
		t.Fatalf("second merge: %v", err)
	}
	if string(first) != string(second) {
		t.Errorf("not idempotent\n--- first ---\n%s\n--- second ---\n%s", first, second)
	}
}
