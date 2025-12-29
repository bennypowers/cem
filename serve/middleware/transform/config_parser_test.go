/*
Copyright © 2025 Benny Powers <web@bennypowers.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

package transform_test

import (
	"strings"
	"testing"

	"bennypowers.dev/cem/internal/platform/testutil"
	"bennypowers.dev/cem/serve/middleware/transform"
)

// TestParseConfigFileURLRewrites_Valid tests parsing of valid config file
func TestParseConfigFileURLRewrites_Valid(t *testing.T) {
	fs := testutil.NewFixtureFS(t, "path-mappings/in-place", "/test")

	validConfig := `serve:
  urlRewrites:
    - urlPattern: "/dist/:path*"
      urlTemplate: "/src/{{.path}}"
    - urlPattern: "/build/:name"
      urlTemplate: "/source/{{.name}}"
`
	fs.AddFile("/test/cem.yaml", validConfig, 0644)

	rewrites, visitedFiles, err := transform.ParseConfigFileURLRewrites("/test/cem.yaml", fs)
	if err != nil {
		t.Fatalf("ParseConfigFileURLRewrites failed: %v", err)
	}

	if len(rewrites) != 2 {
		t.Fatalf("Expected 2 URL rewrites, got %d", len(rewrites))
	}

	if rewrites[0].URLPattern != "/dist/:path*" {
		t.Errorf("Expected first URLPattern '/dist/:path*', got %q", rewrites[0].URLPattern)
	}

	if rewrites[0].URLTemplate != "/src/{{.path}}" {
		t.Errorf("Expected first URLTemplate '/src/{{.path}}', got %q", rewrites[0].URLTemplate)
	}

	if rewrites[1].URLPattern != "/build/:name" {
		t.Errorf("Expected second URLPattern '/build/:name', got %q", rewrites[1].URLPattern)
	}

	if len(visitedFiles) != 1 {
		t.Errorf("Expected 1 visited file, got %d", len(visitedFiles))
	}
}

// TestParseConfigFileURLRewrites_Empty tests empty config file
func TestParseConfigFileURLRewrites_Empty(t *testing.T) {
	fs := testutil.NewFixtureFS(t, "path-mappings/in-place", "/test")

	emptyConfig := `serve:
  urlRewrites: []
`
	fs.AddFile("/test/cem.yaml", emptyConfig, 0644)

	rewrites, visitedFiles, err := transform.ParseConfigFileURLRewrites("/test/cem.yaml", fs)
	if err != nil {
		t.Fatalf("ParseConfigFileURLRewrites failed: %v", err)
	}

	if len(rewrites) != 0 {
		t.Errorf("Expected 0 URL rewrites for empty config, got %d", len(rewrites))
	}

	if len(visitedFiles) != 1 {
		t.Errorf("Expected 1 visited file, got %d", len(visitedFiles))
	}
}

// TestParseConfigFileURLRewrites_NoServeSection tests config without serve section
func TestParseConfigFileURLRewrites_NoServeSection(t *testing.T) {
	fs := testutil.NewFixtureFS(t, "path-mappings/in-place", "/test")

	noServeConfig := `other:
  setting: value
`
	fs.AddFile("/test/cem.yaml", noServeConfig, 0644)

	rewrites, visitedFiles, err := transform.ParseConfigFileURLRewrites("/test/cem.yaml", fs)
	if err != nil {
		t.Fatalf("ParseConfigFileURLRewrites failed: %v", err)
	}

	if len(rewrites) != 0 {
		t.Errorf("Expected 0 URL rewrites when no serve section, got %d", len(rewrites))
	}

	if len(visitedFiles) != 1 {
		t.Errorf("Expected 1 visited file, got %d", len(visitedFiles))
	}
}

// TestParseConfigFileURLRewrites_MissingFile tests handling of non-existent file
func TestParseConfigFileURLRewrites_MissingFile(t *testing.T) {
	fs := testutil.NewFixtureFS(t, "path-mappings/in-place", "/test")

	rewrites, visitedFiles, err := transform.ParseConfigFileURLRewrites("/test/nonexistent.yaml", fs)
	if err != nil {
		t.Errorf("Expected no error for missing file, got: %v", err)
	}

	if len(rewrites) != 0 {
		t.Errorf("Expected 0 URL rewrites for missing file, got %d", len(rewrites))
	}

	if len(visitedFiles) != 0 {
		t.Errorf("Expected 0 visited files for missing file, got %d", len(visitedFiles))
	}
}

// TestParseConfigFileURLRewrites_InvalidYAML tests invalid YAML handling
func TestParseConfigFileURLRewrites_InvalidYAML(t *testing.T) {
	fs := testutil.NewFixtureFS(t, "path-mappings/in-place", "/test")

	invalidYAML := `serve:
  urlRewrites:
    - urlPattern: "/dist/:path*"
      urlTemplate: "/src/{{.path}}"
    - this is not valid yaml
`
	fs.AddFile("/test/cem.yaml", invalidYAML, 0644)

	_, _, err := transform.ParseConfigFileURLRewrites("/test/cem.yaml", fs)
	if err == nil {
		t.Error("Expected error for invalid YAML, got nil")
	}

	if !strings.Contains(err.Error(), "parsing config YAML") {
		t.Errorf("Expected 'parsing config YAML' in error message, got: %v", err)
	}
}

// TestParseConfigFileURLRewrites_AbsolutePath tests that returned path is absolute
func TestParseConfigFileURLRewrites_AbsolutePath(t *testing.T) {
	fs := testutil.NewFixtureFS(t, "path-mappings/in-place", "/test")

	validConfig := `serve:
  urlRewrites:
    - urlPattern: "/dist/:path*"
      urlTemplate: "/src/{{.path}}"
`
	fs.AddFile("/test/cem.yaml", validConfig, 0644)

	_, visitedFiles, err := transform.ParseConfigFileURLRewrites("/test/cem.yaml", fs)
	if err != nil {
		t.Fatalf("ParseConfigFileURLRewrites failed: %v", err)
	}

	if len(visitedFiles) != 1 {
		t.Fatalf("Expected 1 visited file, got %d", len(visitedFiles))
	}

	// Path should be absolute (starts with /)
	if !strings.HasPrefix(visitedFiles[0], "/") {
		t.Errorf("Expected absolute path, got: %s", visitedFiles[0])
	}
}
