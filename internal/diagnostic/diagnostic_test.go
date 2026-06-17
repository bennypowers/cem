package diagnostic_test

import (
	"bytes"
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/internal/config"
	"bennypowers.dev/cem/internal/diagnostic"
	"bennypowers.dev/cem/internal/platform/testutil"
	"bennypowers.dev/cem/internal/sourcepos"
)

var sampleYAML = []byte(`serve:
  port: 99999
  demos:
    rendering: invalid-mode
generate:
  files:
    - "src/**/*.ts"
  designTokens:
    prefix: --bad
`)

func goldenOpts(t *testing.T) testutil.GoldenOptions {
	t.Helper()
	return testutil.GoldenOptions{
		Dir:       "goldens",
		Extension: ".txt",
		StripANSI: true,
		FS:        testutil.LoadTestdataFS(t, filepath.Join("testdata"), "/"),
	}
}

func TestRender_ErrorWithPosition(t *testing.T) {
	d := diagnostic.Diagnostic{
		File:     ".config/cem.yaml",
		Pos:      sourcepos.Position{Line: 4, Column: 16},
		Severity: config.SeverityError,
		Field:    "serve.demos.rendering",
		Message:  "must be one of: light, shadow, iframe, chromeless",
		Value:    "invalid-mode",
		Source:   sampleYAML,
	}

	var buf bytes.Buffer
	diagnostic.Render(&buf, d)

	testutil.CheckGolden(t, "error-with-position", buf.Bytes(), goldenOpts(t))
}

func TestRender_WarningWithPosition(t *testing.T) {
	d := diagnostic.Diagnostic{
		File:     ".config/cem.yaml",
		Pos:      sourcepos.Position{Line: 9, Column: 13},
		Severity: config.SeverityWarning,
		Field:    "generate.designTokens.prefix",
		Message:  `should not start with dashes (use "bad" instead)`,
		Value:    "--bad",
		Source:   sampleYAML,
	}

	var buf bytes.Buffer
	diagnostic.Render(&buf, d)

	testutil.CheckGolden(t, "warning-with-position", buf.Bytes(), goldenOpts(t))
}

func TestRender_NoPosition(t *testing.T) {
	d := diagnostic.Diagnostic{
		File:     ".config/cem.yaml",
		Severity: config.SeverityError,
		Field:    "serve.urlRewrites",
		Message:  "bad pattern",
	}

	var buf bytes.Buffer
	diagnostic.Render(&buf, d)

	testutil.CheckGolden(t, "no-position", buf.Bytes(), goldenOpts(t))
}

func TestRender_ErrorWithValue(t *testing.T) {
	d := diagnostic.Diagnostic{
		File:     ".config/cem.yaml",
		Pos:      sourcepos.Position{Line: 2, Column: 9},
		Severity: config.SeverityError,
		Field:    "serve.port",
		Message:  "must be between 0 and 65535",
		Value:    "99999",
		Source:   sampleYAML,
	}

	var buf bytes.Buffer
	diagnostic.Render(&buf, d)

	testutil.CheckGolden(t, "error-with-value", buf.Bytes(), goldenOpts(t))
}

func TestRender_FirstLine(t *testing.T) {
	source := []byte("port: 99999\nother: value\n")
	d := diagnostic.Diagnostic{
		File:     "cem.yaml",
		Pos:      sourcepos.Position{Line: 1, Column: 7},
		Severity: config.SeverityError,
		Field:    "port",
		Message:  "must be between 0 and 65535",
		Value:    "99999",
		Source:   source,
	}

	var buf bytes.Buffer
	diagnostic.Render(&buf, d)

	testutil.CheckGolden(t, "first-line", buf.Bytes(), goldenOpts(t))
}

func TestRender_LastLine(t *testing.T) {
	source := []byte("other: value\nport: 99999\n")
	d := diagnostic.Diagnostic{
		File:     "cem.yaml",
		Pos:      sourcepos.Position{Line: 2, Column: 7},
		Severity: config.SeverityError,
		Field:    "port",
		Message:  "must be between 0 and 65535",
		Value:    "99999",
		Source:   source,
	}

	var buf bytes.Buffer
	diagnostic.Render(&buf, d)

	testutil.CheckGolden(t, "last-line", buf.Bytes(), goldenOpts(t))
}

func TestRenderAll(t *testing.T) {
	diags := []diagnostic.Diagnostic{
		{
			File:     ".config/cem.yaml",
			Pos:      sourcepos.Position{Line: 2, Column: 9},
			Severity: config.SeverityError,
			Field:    "serve.port",
			Message:  "must be between 0 and 65535",
			Value:    "99999",
			Source:   sampleYAML,
		},
		{
			File:     ".config/cem.yaml",
			Pos:      sourcepos.Position{Line: 4, Column: 16},
			Severity: config.SeverityError,
			Field:    "serve.demos.rendering",
			Message:  "must be one of: light, shadow, iframe, chromeless",
			Value:    "invalid-mode",
			Source:   sampleYAML,
		},
	}

	var buf bytes.Buffer
	diagnostic.RenderAll(&buf, diags)

	testutil.CheckGolden(t, "render-all", buf.Bytes(), goldenOpts(t))
}
