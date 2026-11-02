package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"bennypowers.dev/cem/internal/platform/testutil"
)

func TestValidateSchema_ExampleConfigs(t *testing.T) {
	matches, err := filepath.Glob(filepath.Join("..", "..", "examples", "*", ".config", "cem.yaml"))
	if err != nil {
		t.Fatalf("failed to glob example configs: %v", err)
	}
	if len(matches) == 0 {
		t.Skip("no example configs found")
	}
	for _, path := range matches {
		name := filepath.Base(filepath.Dir(filepath.Dir(path)))
		t.Run(name, func(t *testing.T) {
			data, err := os.ReadFile(path)
			if err != nil {
				t.Fatalf("failed to read %s: %v", path, err)
			}
			errs := ValidateSchema(data, "yaml")
			for _, e := range errs {
				t.Errorf("%s", e.Error())
			}
		})
	}
}

func TestValidateSchema_TemplateFunctions(t *testing.T) {
	t.Run("accepts documented template functions", func(t *testing.T) {
		data := []byte(`
generate:
  demoDiscovery:
    urlPattern: "/elements/:tag/demo/:demo.html"
    urlTemplate: "https://example.com/{{.tag | alias | slug}}/demo/{{.demo | lower}}/"
serve:
  urlRewrites:
    - urlPattern: "/api/:path*"
      urlTemplate: "/backend/{{.path | upper}}"
`)
		for _, e := range ValidateSchema(data, "yaml") {
			t.Errorf("unexpected validation error: %s", e.Error())
		}
	})

	t.Run("rejects unknown template function", func(t *testing.T) {
		data := []byte(`
generate:
  demoDiscovery:
    urlTemplate: "{{.tag | nope}}"
`)
		errs := ValidateSchema(data, "yaml")
		if len(errs) == 0 {
			t.Fatal("expected validation error for unknown template function, got none")
		}
		if field := errs[0].Field; field != "generate.demoDiscovery.urlTemplate" {
			t.Errorf("expected error on generate.demoDiscovery.urlTemplate, got %q", field)
		}
	})
}

func TestValidateSchema(t *testing.T) {
	mfs := testutil.LoadTestdataFS(t, filepath.Join("testdata", "schema"), "/")

	fixtureEntries, err := mfs.ReadDir("/fixtures")
	if err != nil {
		t.Fatalf("failed to read fixture directory: %v", err)
	}

	for _, entry := range fixtureEntries {
		if entry.IsDir() {
			continue
		}
		name := strings.TrimSuffix(entry.Name(), filepath.Ext(entry.Name()))
		t.Run(name, func(t *testing.T) {
			data, err := mfs.ReadFile(filepath.Join("/fixtures", entry.Name()))
			if err != nil {
				t.Fatalf("failed to read fixture: %v", err)
			}

			format := FormatFromPath(entry.Name())
			errs := ValidateSchema(data, format)

			type result struct {
				Valid  bool              `json:"valid"`
				Errors []ValidationError `json:"errors,omitempty"`
			}
			r := result{Valid: len(errs) == 0, Errors: errs}

			jsonBytes, err := json.MarshalIndent(r, "", "  ")
			if err != nil {
				t.Fatalf("failed to marshal result: %v", err)
			}

			testutil.CheckGolden(t, name, jsonBytes, testutil.GoldenOptions{
				Dir:         filepath.Join("testdata", "schema", "goldens"),
				Extension:   ".json",
				UseJSONDiff: true,
			})
		})
	}
}
