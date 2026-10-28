/*
Copyright © 2025 Benny Powers <web@bennypowers.com>

SPDX-License-Identifier: GPL-3.0-or-later

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
package validate

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"bennypowers.dev/cem/internal/platform"
	"bennypowers.dev/cem/internal/platform/testutil"
	"github.com/adrg/xdg"
)

// TestMain installs a fail-fast fetch default so a test that unexpectedly
// reaches the network fails instead of requesting unpkg. Tests that exercise
// fetching install their own stub with stubHTTPGet.
func TestMain(m *testing.M) {
	httpGet = func(url string) (*http.Response, error) {
		return nil, fmt.Errorf("unexpected schema fetch in tests: %s", url)
	}
	os.Exit(m.Run())
}

// stubHTTPGet replaces the schema fetch seam for the duration of a test.
// Tests must not reach the real network.
func stubHTTPGet(t *testing.T, fn func(url string) (*http.Response, error)) {
	t.Helper()
	orig := httpGet
	t.Cleanup(func() { httpGet = orig })
	httpGet = fn
}

func schemaResponse(body []byte) *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Status:     "200 OK",
		Body:       io.NopCloser(bytes.NewReader(body)),
	}
}

// loadValidateFixtureMapFS loads a fixture manifest into a MapFileSystem
// and returns the MapFileSystem and virtual path.
func loadValidateFixtureMapFS(t *testing.T, fixture string) (*platform.MapFileSystem, string) {
	t.Helper()
	mfs := testutil.LoadTestdataFS(t, filepath.Join("..", "cmd", "testdata", "fixtures", fixture), "/")
	return mfs, "custom-elements.json"
}

func TestValidateGolden(t *testing.T) {
	testCases := []struct {
		fixture string
		name    string
	}{
		{"valid-manifest", "valid_manifest"},
		{"valid-generics", "valid_generics"},
		{"invalid-class", "invalid_class"},
		{"invalid-mixin", "invalid_mixin"},
		{"invalid-variable", "invalid_variable"},
		{"invalid-function", "invalid_function"},
		{"invalid-field", "invalid_field"},
		{"invalid-method", "invalid_method"},
		{"invalid-event", "invalid_event"},
		{"event-without-type", "event_without_type"},
		{"invalid-attribute", "invalid_attribute"},
		{"invalid-slot", "invalid_slot"},
		{"invalid-css-part", "invalid_css_part"},
		{"invalid-css-custom-property", "invalid_css_custom_property"},
		{"unsupported-schema", "unsupported_schema"},
		{"warning-lifecycle-methods", "warning_lifecycle_methods"},
		{"warning-implementation-details", "warning_implementation_details"},
		{"warning-superclass-modules", "warning_superclass_modules"},
		{"warning-css-verbose", "warning_css_verbose"},
		{"schema-upgrade-2-1-0", "schema_upgrade_2_1_0"},
	}

	t.Run("os", func(t *testing.T) {
		fsys := platform.NewOSFileSystem()
		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				fixturePath := filepath.Join("..", "cmd", "testdata", "fixtures", tc.fixture, "custom-elements.json")
				if _, err := os.Stat(fixturePath); os.IsNotExist(err) {
					t.Skipf("Fixture %s does not exist", fixturePath)
				}

				result, err := Validate(fsys, fixturePath, ValidationOptions{
					IncludeWarnings: true,
					DisabledRules:   []string{},
				})
				if err != nil {
					t.Fatalf("Validate returned error: %v", err)
				}

				result.Path = ""

				jsonBytes, err := json.MarshalIndent(result, "", "  ")
				if err != nil {
					t.Fatalf("Failed to marshal result to JSON: %v", err)
				}

				testutil.CheckGolden(t, tc.name, jsonBytes, testutil.GoldenOptions{
					Dir:         "testdata/goldens",
					Extension:   ".json",
					UseJSONDiff: true,
				})
			})
		}
	})

	t.Run("mapfs", func(t *testing.T) {
		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				mfs, virtualPath := loadValidateFixtureMapFS(t, tc.fixture)

				result, err := Validate(mfs, virtualPath, ValidationOptions{
					IncludeWarnings: true,
					DisabledRules:   []string{},
				})
				if err != nil {
					t.Fatalf("Validate returned error: %v", err)
				}

				result.Path = ""

				jsonBytes, err := json.MarshalIndent(result, "", "  ")
				if err != nil {
					t.Fatalf("Failed to marshal result to JSON: %v", err)
				}

				testutil.CheckGolden(t, tc.name, jsonBytes, testutil.GoldenOptions{
					Dir:         "testdata/goldens",
					Extension:   ".json",
					UseJSONDiff: true,
				})
			})
		}
	})
}

func testValidateWithDisabledWarnings(t *testing.T, fsys platform.FileSystem, fixturePath string) {
	t.Helper()

	// Test with warnings disabled by category
	result, err := Validate(fsys, fixturePath, ValidationOptions{
		IncludeWarnings: true,
		DisabledRules:   []string{"lifecycle"},
	})
	if err != nil {
		t.Fatalf("Validate returned error: %v", err)
	}

	for _, warning := range result.Warnings {
		if warning.Category == "lifecycle" {
			t.Errorf("Found lifecycle warning when it should be disabled: %s", warning.Message)
		}
	}

	// Test with warnings disabled by specific ID
	result2, err := Validate(fsys, fixturePath, ValidationOptions{
		IncludeWarnings: true,
		DisabledRules:   []string{"private-underscore-methods"},
	})
	if err != nil {
		t.Fatalf("Validate returned error: %v", err)
	}

	originalResult, err := Validate(fsys, fixturePath, ValidationOptions{
		IncludeWarnings: true,
		DisabledRules:   []string{},
	})
	if err != nil {
		t.Fatalf("Validate returned error: %v", err)
	}

	if len(result2.Warnings) >= len(originalResult.Warnings) {
		t.Errorf("Expected fewer warnings when disabling specific rule, got %d vs %d", len(result2.Warnings), len(originalResult.Warnings))
	}
}

func TestValidateWithDisabledWarnings(t *testing.T) {
	osFixturePath := filepath.Join("..", "cmd", "testdata", "fixtures", "warning-lifecycle-methods", "custom-elements.json")
	if _, err := os.Stat(osFixturePath); os.IsNotExist(err) {
		t.Skip("warning-lifecycle-methods fixture does not exist")
	}

	t.Run("os", func(t *testing.T) {
		testValidateWithDisabledWarnings(t, platform.NewOSFileSystem(), osFixturePath)
	})

	t.Run("mapfs", func(t *testing.T) {
		mfs, virtualPath := loadValidateFixtureMapFS(t, "warning-lifecycle-methods")
		testValidateWithDisabledWarnings(t, mfs, virtualPath)
	})
}

// TestTryFetchSchema_NoCacheDir verifies that tryFetchSchema works when
// xdg.CacheHome is empty, as it would be in wasm or sandboxed environments.
// The embedded schema path (getSchema -> getEmbeddedSchema) bypasses
// tryFetchSchema entirely; this test confirms tryFetchSchema itself is
// resilient when called for non-embedded versions without a cache directory.
func TestTryFetchSchema_NoCacheDir(t *testing.T) {
	// Save and restore xdg.CacheHome
	origCacheHome := xdg.CacheHome
	defer func() { xdg.CacheHome = origCacheHome }()
	xdg.CacheHome = ""

	fsys := platform.NewMapFS(nil)

	// getSchema for an embedded version should succeed without cache
	schemaData, err := getSchema(fsys, "2.1.0", schemaOptions{noNetwork: true})
	if err != nil {
		t.Fatalf("getSchema with empty CacheHome returned error: %v", err)
	}

	var schema map[string]any
	if err := json.Unmarshal(schemaData, &schema); err != nil {
		t.Fatalf("failed to parse schema JSON: %v", err)
	}

	if _, ok := schema["title"]; !ok {
		t.Error("expected schema to contain 'title' field")
	}
}

// TestGetSchemaUpgrade210 verifies that 2.1.0 and 2.1.1 manifests are
// automatically upgraded to use the 2.1.1-speculative schema to work around
// discriminated union issues in the upstream CEM schema.
// See: https://github.com/webcomponents/custom-elements-manifest/issues/138
func testGetSchemaUpgrade210(t *testing.T, fsys platform.FileSystem) {
	t.Helper()

	testCases := []struct {
		version    string
		expected   string
		checkTitle bool // Only 2.1.x schemas have title field
	}{
		{"2.1.0", "2.1.1-speculative", true},
		{"2.1.1", "2.1.1-speculative", true},
		{"2.1.1-speculative", "2.1.1-speculative", true},
		// A leading "v" is a valid semantic version spelling and must resolve
		// to the same embedded schema as the unprefixed form.
		{"v2.1.0", "2.1.1-speculative", true},
		{"v2.1.1-speculative", "2.1.1-speculative", true},
		// Shortened and build-metadata spellings are canonicalized before the
		// embedded lookup, so they resolve without a network request.
		{"2.1", "2.1.1-speculative", true},
		{"v2.1", "2.1.1-speculative", true},
		{"v2.1.0+build.5", "2.1.1-speculative", true},
		{"v2.0.0", "2.0.0", false},
		{"2.0.0", "2.0.0", false},
		{"1.0.0", "1.0.0", false},
	}

	for _, tc := range testCases {
		t.Run(tc.version, func(t *testing.T) {
			schemaData, err := getSchema(fsys, tc.version, schemaOptions{})
			if err != nil {
				t.Fatalf("getSchema(%q) returned error: %v", tc.version, err)
			}

			if !tc.checkTitle {
				// These schemas have no title, so compare the bytes against the
				// embedded schema to prove the intended file was returned rather
				// than a fetched one.
				want, err := getEmbeddedSchema(tc.expected)
				if err != nil {
					t.Fatalf("getEmbeddedSchema(%q): %v", tc.expected, err)
				}
				if !bytes.Equal(schemaData, want) {
					t.Errorf("getSchema(%q) returned a different schema than %q", tc.version, tc.expected)
				}
				return
			}

			var schema map[string]any
			if err := json.Unmarshal(schemaData, &schema); err != nil {
				t.Fatalf("Failed to parse schema JSON: %v", err)
			}

			title, ok := schema["title"].(string)
			if !ok {
				t.Fatalf("Schema missing 'title' field")
			}

			expectedTitle := "Custom Elements Manifest Schema " + tc.expected
			if title != expectedTitle {
				t.Errorf("getSchema(%q) loaded schema with title %q, expected %q", tc.version, title, expectedTitle)
			}
		})
	}
}

func TestGetSchemaUpgrade210(t *testing.T) {
	t.Run("os", func(t *testing.T) {
		testGetSchemaUpgrade210(t, platform.NewOSFileSystem())
	})

	t.Run("mapfs", func(t *testing.T) {
		// getSchema uses embedded schemas for known versions, so MapFS works
		// without needing to load any fixture files
		testGetSchemaUpgrade210(t, platform.NewMapFileSystem(nil))
	})
}

// TestGetSchemaShortVersionsResolveEmbedded verifies that shortened semantic
// versions are ranges, not exact versions: a bare major selects the newest
// embedded schema with that major. Treating "0" as the canonical 0.0.0 broke
// validation (0.0.0 is not embedded, so the fetch used to 404), and treating
// "2" as 2.0.0 silently loosened validation.
//
// Inline assertions: this test checks which embedded schema was selected,
// where a golden file would obscure the cache fixture used to keep the
// pre-fix behavior offline.
func TestGetSchemaShortVersionsResolveEmbedded(t *testing.T) {
	// Register the reload before Setenv so cleanups run in the right order:
	// the environment is restored first, then xdg re-reads it.
	t.Cleanup(xdg.Reload)
	cacheHome := t.TempDir()
	t.Setenv("XDG_CACHE_HOME", cacheHome)
	xdg.Reload()

	fsys := platform.NewMapFileSystem(nil)
	// Seed the cache under the canonical 0.0.0 spelling so the pre-fix code
	// resolves offline: it must not be mistaken for "0", which npm resolves to
	// 0.1.0.
	cacheDir := filepath.Join(cacheHome, "cem", "schemas")
	if err := fsys.WriteFile(filepath.Join(cacheDir, "0.0.0.json"), []byte(`{"title":"wrong schema"}`), 0o644); err != nil {
		t.Fatalf("writing cached schema: %v", err)
	}

	testCases := []struct {
		version  string
		embedded string
	}{
		{"0", "0.1.0"},
		{"2", speculativeSchemaVersion},
		// Build metadata on a shortened spelling names the same range as the
		// spelling without it, matching how canonical versions drop metadata.
		{"2.1+build", speculativeSchemaVersion},
		{"v2.1+build.5", speculativeSchemaVersion},
	}

	for _, tc := range testCases {
		t.Run(tc.version, func(t *testing.T) {
			// Network access is disabled so any resolution that misses the
			// embedded schemas fails loudly instead of reaching unpkg.
			got, err := getSchema(fsys, tc.version, schemaOptions{noNetwork: true})
			if err != nil {
				t.Fatalf("getSchema(%q): %v", tc.version, err)
			}
			want, err := getEmbeddedSchema(tc.embedded)
			if err != nil {
				t.Fatalf("getEmbeddedSchema(%q): %v", tc.embedded, err)
			}
			if !bytes.Equal(got, want) {
				t.Errorf("getSchema(%q) did not select the newest embedded %s schema", tc.version, tc.embedded)
			}
		})
	}
}

// TestValidateBareMajorUsesNewestEmbeddedSchema verifies the validation
// consequence of resolving a bare major as a range. The 2.0.0 schema ignores
// cssStates, while the 2.1.x and speculative schemas require a name on each
// entry. Resolving "2" to 2.0.0 therefore reported this manifest as valid.
//
// Inline assertions: this test checks the validity and error IDs of one
// result, where a golden file would add indirection only.
func TestValidateBareMajorUsesNewestEmbeddedSchema(t *testing.T) {
	manifest := []byte(`{
	  "schemaVersion": "2",
	  "modules": [
	    {
	      "kind": "javascript-module",
	      "path": "css-states.js",
	      "declarations": [
	        {
	          "kind": "class",
	          "name": "CssStatesElement",
	          "customElement": true,
	          "tagName": "css-states-element",
	          "cssStates": [{ "summary": "Missing name" }]
	        }
	      ]
	    }
	  ]
	}`)

	pipeline, err := NewValidationPipeline(platform.NewMapFileSystem(nil), manifest)
	if err != nil {
		t.Fatalf("NewValidationPipeline: %v", err)
	}
	result, err := pipeline.Validate(ValidationOptions{})
	if err != nil {
		t.Fatalf("Validate: %v", err)
	}
	if result.IsValid {
		t.Fatalf("Validate accepted a cssState without a name under schemaVersion \"2\": %+v", result.Errors)
	}

	found := false
	for _, issue := range result.Errors {
		if strings.Contains(issue.Message, "name") {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Errors = %+v, want a missing name error from the newest embedded 2.x schema", result.Errors)
	}
}

// TestValidateBareMajorDoesNotWarnAsOld verifies that the old-version warning
// follows the schema actually used. A bare major "2" selects the newest
// embedded 2.x schema, not the exact 2.0.0 spelling, so it must not be
// reported as older than 2.1.0.
//
// Inline assertions: this test checks for the absence of one warning ID,
// where a golden file would add indirection only.
func TestValidateBareMajorDoesNotWarnAsOld(t *testing.T) {
	manifest := []byte(`{"schemaVersion":"2","modules":[]}`)
	pipeline, err := NewValidationPipeline(platform.NewMapFileSystem(nil), manifest)
	if err != nil {
		t.Fatalf("NewValidationPipeline: %v", err)
	}
	result, err := pipeline.Validate(ValidationOptions{IncludeWarnings: true})
	if err != nil {
		t.Fatalf("Validate: %v", err)
	}
	for _, warning := range result.Warnings {
		if warning.ID == "schema-version-old" {
			t.Errorf("Warnings = %+v, want no schema-version-old warning for a range that resolves to the newest embedded 2.x schema", result.Warnings)
		}
	}
}

// TestGetSchemaCacheUsesInjectedFileSystem verifies that cache lookup and
// storage go through the caller's FileSystem. xdg.CacheFile creates
// directories on the host, which escaped in-memory and sandboxed filesystems.
//
// Inline assertions: this test checks one sentinel value and one host path,
// where a golden file would obscure the filesystem boundary under test.
func TestGetSchemaCacheUsesInjectedFileSystem(t *testing.T) {
	// Register the reload before Setenv so cleanups run in the right order:
	// the environment is restored first, then xdg re-reads it.
	t.Cleanup(xdg.Reload)
	cacheHome := filepath.Join(t.TempDir(), "cache")
	t.Setenv("XDG_CACHE_HOME", cacheHome)
	xdg.Reload()

	fsys := platform.NewMapFileSystem(nil)
	cacheDir := filepath.Join(cacheHome, "cem", "schemas")
	sentinel := []byte(`{"type":"object","required":["schemaVersion"]}`)
	if err := fsys.WriteFile(filepath.Join(cacheDir, "latest.json"), sentinel, 0o644); err != nil {
		t.Fatalf("writing cached schema: %v", err)
	}

	got, err := getSchema(fsys, "latest", schemaOptions{noNetwork: true})
	if err != nil {
		t.Fatalf("getSchema: %v", err)
	}
	if !bytes.Equal(got, sentinel) {
		t.Errorf("getSchema did not read the cached schema through the supplied FileSystem")
	}
	if _, err := os.Stat(cacheHome); !os.IsNotExist(err) {
		t.Errorf("getSchema created %s on the host filesystem; cache I/O must go through the supplied FileSystem", cacheHome)
	}
}

// TestSchemaWithoutNetworkDoesNotFetch verifies that WithoutNetwork turns a
// cache miss for an unknown version into an error instead of an unpkg request.
//
// Inline assertions: this test checks the error contract for one version,
// where a golden file would obscure the fetch seam under test.
func TestSchemaWithoutNetworkDoesNotFetch(t *testing.T) {
	origCacheHome := xdg.CacheHome
	t.Cleanup(func() { xdg.CacheHome = origCacheHome })
	xdg.CacheHome = "" // no cache directory
	stubHTTPGet(t, func(url string) (*http.Response, error) {
		t.Fatalf("Schema fetched %s even though WithoutNetwork was set", url)
		return nil, nil
	})

	// readOnlySchemaFS cannot write, so this also proves the no-network path
	// does not try to cache anything.
	_, err := Schema(readOnlySchemaFS{}, "9.9.9", WithoutNetwork())
	if err == nil {
		t.Fatal("Schema returned no error for an unknown version with WithoutNetwork")
	}
	if !strings.Contains(err.Error(), `"9.9.9"`) {
		t.Errorf("error = %q, want it to name the requested version", err)
	}
	if !strings.Contains(err.Error(), "network access is disabled") {
		t.Errorf("error = %q, want it to explain that network access is disabled", err)
	}
}

// TestSchemaReadsCacheBeforeNetwork verifies that a cached schema wins over a
// network fetch even when fetching is allowed.
//
// Inline assertions: this test checks one sentinel value, where a golden file
// would obscure the cache setup under test.
func TestSchemaReadsCacheBeforeNetwork(t *testing.T) {
	t.Cleanup(xdg.Reload)
	cacheHome := t.TempDir()
	t.Setenv("XDG_CACHE_HOME", cacheHome)
	xdg.Reload()
	stubHTTPGet(t, func(url string) (*http.Response, error) {
		t.Fatalf("Schema fetched %s for a cached version", url)
		return nil, nil
	})

	fsys := platform.NewMapFileSystem(nil)
	sentinel := []byte(`{"title":"cached schema"}`)
	cachePath := filepath.Join(cacheHome, "cem", "schemas", "9.9.9.json")
	if err := fsys.WriteFile(cachePath, sentinel, 0o644); err != nil {
		t.Fatalf("writing cached schema: %v", err)
	}

	got, err := Schema(fsys, "9.9.9")
	if err != nil {
		t.Fatalf("Schema: %v", err)
	}
	if !bytes.Equal(got, sentinel) {
		t.Errorf("Schema did not return the cached schema")
	}
}

// TestSchemaDefaultFetchesWhenNotEmbedded verifies that the default behavior
// still fetches an unknown version and caches the result through the supplied
// filesystem.
//
// Inline assertions: this test checks the fetched bytes, the URL, and the
// cached copy, where a golden file would obscure the fetch seam under test.
func TestSchemaDefaultFetchesWhenNotEmbedded(t *testing.T) {
	t.Cleanup(xdg.Reload)
	cacheHome := t.TempDir()
	t.Setenv("XDG_CACHE_HOME", cacheHome)
	xdg.Reload()

	fetched := []byte(`{"title":"fetched schema"}`)
	var gotURL string
	stubHTTPGet(t, func(url string) (*http.Response, error) {
		gotURL = url
		return schemaResponse(fetched), nil
	})

	fsys := platform.NewMapFileSystem(nil)
	got, err := Schema(fsys, "9.9.9")
	if err != nil {
		t.Fatalf("Schema: %v", err)
	}
	if !bytes.Equal(got, fetched) {
		t.Errorf("Schema returned %q, want the fetched schema", got)
	}
	if !strings.Contains(gotURL, "custom-elements-manifest@9.9.9") {
		t.Errorf("fetch URL = %q, want it to name version 9.9.9", gotURL)
	}

	cached, err := fsys.ReadFile(filepath.Join(cacheHome, "cem", "schemas", "9.9.9.json"))
	if err != nil {
		t.Fatalf("reading cached schema: %v", err)
	}
	if !bytes.Equal(cached, fetched) {
		t.Errorf("cached schema = %q, want the fetched schema", cached)
	}
}

// readOnlySchemaFS implements only ReadFS. It cannot cache fetched schemas.
type readOnlySchemaFS struct{}

func (readOnlySchemaFS) ReadFile(name string) ([]byte, error) { return nil, fs.ErrNotExist }

func (readOnlySchemaFS) Stat(name string) (fs.FileInfo, error) { return nil, fs.ErrNotExist }

// TestSchemaDefaultFetchesWithReadOnlyFS verifies that a ReadFS-only
// implementation can still fetch an unknown version: caching is optional.
//
// Inline assertions: this test checks the fetched bytes, where a golden file
// would obscure the fetch seam under test.
func TestSchemaDefaultFetchesWithReadOnlyFS(t *testing.T) {
	t.Cleanup(xdg.Reload)
	t.Setenv("XDG_CACHE_HOME", t.TempDir())
	xdg.Reload()

	fetched := []byte(`{"title":"fetched schema"}`)
	stubHTTPGet(t, func(url string) (*http.Response, error) {
		return schemaResponse(fetched), nil
	})

	got, err := Schema(readOnlySchemaFS{}, "9.9.9")
	if err != nil {
		t.Fatalf("Schema: %v", err)
	}
	if !bytes.Equal(got, fetched) {
		t.Errorf("Schema returned %q, want the fetched schema", got)
	}
}

// partialWriteFS writes half of the requested data and then fails, simulating
// a short write that leaves a truncated cache file. It implements Remove so
// the cache can clean up the partial entry.
type partialWriteFS struct {
	files map[string][]byte
}

func (p *partialWriteFS) ReadFile(name string) ([]byte, error) {
	data, ok := p.files[name]
	if !ok {
		return nil, fs.ErrNotExist
	}
	return data, nil
}

func (p *partialWriteFS) Stat(name string) (fs.FileInfo, error) {
	data, ok := p.files[name]
	if !ok {
		return nil, fs.ErrNotExist
	}
	return partialWriteFileInfo{name: name, size: int64(len(data))}, nil
}

func (p *partialWriteFS) WriteFile(name string, data []byte, perm fs.FileMode) error {
	p.files[name] = append([]byte(nil), data[:len(data)/2]...)
	return errors.New("simulated write failure")
}

func (p *partialWriteFS) MkdirAll(path string, perm fs.FileMode) error { return nil }

func (p *partialWriteFS) Remove(name string) error {
	delete(p.files, name)
	return nil
}

type partialWriteFileInfo struct {
	name string
	size int64
}

func (i partialWriteFileInfo) Name() string       { return i.name }
func (i partialWriteFileInfo) Size() int64        { return i.size }
func (i partialWriteFileInfo) Mode() fs.FileMode  { return 0o644 }
func (i partialWriteFileInfo) ModTime() time.Time { return time.Time{} }
func (i partialWriteFileInfo) IsDir() bool        { return false }
func (i partialWriteFileInfo) Sys() any           { return nil }

// TestSchemaCacheWriteErrorRemovesPartialFile verifies that a failed cache
// write does not leave truncated bytes that later resolve as a cached schema.
//
// Inline assertions: this test checks the returned bytes, the cache state, and
// the fetch count, where a golden file would obscure the short write under
// test.
func TestSchemaCacheWriteErrorRemovesPartialFile(t *testing.T) {
	t.Cleanup(xdg.Reload)
	cacheHome := t.TempDir()
	t.Setenv("XDG_CACHE_HOME", cacheHome)
	xdg.Reload()

	fetched := []byte(`{"title":"fetched schema"}`)
	fetches := 0
	stubHTTPGet(t, func(url string) (*http.Response, error) {
		fetches++
		return schemaResponse(fetched), nil
	})

	fsys := &partialWriteFS{files: map[string][]byte{}}
	first, err := Schema(fsys, "9.9.9")
	if err != nil {
		t.Fatalf("Schema: %v", err)
	}
	if !bytes.Equal(first, fetched) {
		t.Errorf("Schema returned %q, want the fetched schema", first)
	}

	cachePath := filepath.Join(cacheHome, "cem", "schemas", "9.9.9.json")
	if _, err := fsys.Stat(cachePath); !errors.Is(err, fs.ErrNotExist) {
		t.Errorf("partial cache file was left behind after a write error")
	}

	second, err := Schema(fsys, "9.9.9")
	if err != nil {
		t.Fatalf("second Schema: %v", err)
	}
	if fetches != 2 {
		t.Errorf("fetches = %d, want 2: a partial cache entry must not be reused", fetches)
	}
	if !bytes.Equal(second, fetched) {
		t.Errorf("second Schema returned %q, want the fetched schema", second)
	}
}

// TestValidateNoNetworkDoesNotFetch verifies that ValidationOptions.NoNetwork
// reaches the schema loader and prevents an unpkg request.
//
// Inline assertions: this test checks the error contract for one manifest,
// where a golden file would obscure the fetch seam under test.
func TestValidateNoNetworkDoesNotFetch(t *testing.T) {
	origCacheHome := xdg.CacheHome
	t.Cleanup(func() { xdg.CacheHome = origCacheHome })
	xdg.CacheHome = ""
	stubHTTPGet(t, func(url string) (*http.Response, error) {
		t.Fatalf("Validate fetched %s even though NoNetwork was set", url)
		return nil, nil
	})

	manifest := []byte(`{"schemaVersion":"9.9.9","modules":[]}`)
	pipeline, err := NewValidationPipeline(platform.NewMapFileSystem(nil), manifest)
	if err != nil {
		t.Fatalf("NewValidationPipeline: %v", err)
	}
	_, err = pipeline.Validate(ValidationOptions{NoNetwork: true})
	if err == nil {
		t.Fatal("Validate returned no error for an unknown version with NoNetwork")
	}
	if !strings.Contains(err.Error(), "9.9.9") {
		t.Errorf("error = %q, want it to name the requested version", err)
	}
}

// TestValidateNoNetworkAcceptsEmbeddedSchema verifies that NoNetwork does not
// affect versions served from embedded schemas.
//
// Inline assertions: this test checks one validity result, where a golden file
// would add indirection only.
func TestValidateNoNetworkAcceptsEmbeddedSchema(t *testing.T) {
	manifest := []byte(`{"schemaVersion":"2.1.0","modules":[]}`)
	pipeline, err := NewValidationPipeline(platform.NewMapFileSystem(nil), manifest)
	if err != nil {
		t.Fatalf("NewValidationPipeline: %v", err)
	}
	result, err := pipeline.Validate(ValidationOptions{NoNetwork: true})
	if err != nil {
		t.Fatalf("Validate: %v", err)
	}
	if !result.IsValid {
		t.Errorf("Validate rejected a manifest with an embedded schema under NoNetwork: %+v", result.Errors)
	}
}

// Hardening of pre-existing behavior: schemaVersion was previously passed
// unchecked into embedded-schema and cache path construction. Unsafe values
// must be rejected before they can become path components.
//
// Inline assertions: this test checks rejection of scalar version strings,
// where a golden file would only duplicate the input list.
func TestGetSchemaRejectsUnsafeVersions(t *testing.T) {
	fsys := platform.NewMapFileSystem(nil)
	versions := []string{
		"../../../../tmp/x",
		"/tmp/x",
		`..\..\..\..\tmp\x`,
		"2.1.0/../../x",
		// URL and query delimiters must not reach the fetch path.
		"2.1.0?x",
		"2.1.0 x",
	}

	for _, version := range versions {
		t.Run(version, func(t *testing.T) {
			_, err := getSchema(fsys, version, schemaOptions{})
			if err == nil {
				t.Fatalf("getSchema(%q) = nil error, want rejection", version)
			}
			if !strings.Contains(err.Error(), "invalid schema version") {
				t.Errorf("getSchema(%q) error = %q, want an invalid schema version error", version, err)
			}
		})
	}
}

// Inline assertions: this test checks classification of scalar version
// strings, where a golden file would only duplicate the input/output table.
func TestSchemaVersionClassification(t *testing.T) {
	canonical := []struct {
		version string
		want    string
	}{
		{"2.1.0", "2.1.0"},
		{"v2.1.0", "2.1.0"},
		{"2.1", "2.1.0"},
		{"v2.1", "2.1.0"},
		{"v2.1.0+build.5", "2.1.0"},
		{"2.1.1-speculative", "2.1.1-speculative"},
	}
	for _, tc := range canonical {
		got, ok := canonicalSchemaVersion(tc.version)
		if !ok || got != tc.want {
			t.Errorf("canonicalSchemaVersion(%q) = (%q, %v), want (%q, true)", tc.version, got, ok, tc.want)
		}
	}

	for _, version := range []string{"latest", "^2.0.0", "2.x", "next"} {
		if got, ok := canonicalSchemaVersion(version); ok {
			t.Errorf("canonicalSchemaVersion(%q) = (%q, true), want not a semantic version", version, got)
		}
	}

	safe := []string{"2.1.0", "v2.1.0", "latest", "^2.0.0", "2.x"}
	for _, version := range safe {
		if !isPathSafeSchemaVersion(version) {
			t.Errorf("isPathSafeSchemaVersion(%q) = false, want true", version)
		}
	}

	unsafe := []string{
		"", "../../x", "/tmp/x", `..\x`, "2.1.0/../../x", "2.1.0?x", "2.1.0 x",
		// Unicode whitespace and control characters must not reach the URL or
		// cache path either.
		" ", " ", "\u0085", "\u000b", "\u000c",
	}
	for _, version := range unsafe {
		if isPathSafeSchemaVersion(version) {
			t.Errorf("isPathSafeSchemaVersion(%q) = true, want false", version)
		}
	}
}

// TestValidateAcceptsPathSafeNonSemverVersionFromCache verifies that a
// path-safe non-semver schemaVersion, such as a dist-tag, still resolves
// through the schema cache instead of being rejected as unsafe.
func TestValidateAcceptsPathSafeNonSemverVersionFromCache(t *testing.T) {
	// Inline assertions: this test checks a boolean result and one warning,
	// where a golden file would obscure the cache setup under test.
	// Register the reload before Setenv so cleanups run in the right order:
	// the environment is restored first, then xdg re-reads it.
	t.Cleanup(xdg.Reload)
	t.Setenv("XDG_CACHE_HOME", t.TempDir())
	xdg.Reload()

	fsys := platform.NewMapFileSystem(nil)
	cacheDir := filepath.Join(xdg.CacheHome, "cem", "schemas")
	// Use a sentinel schema that accepts modules as an object. The official
	// schema requires an array, so a network fallback cannot make this test
	// pass by accident.
	sentinel := []byte(`{"type":"object","required":["schemaVersion"]}`)
	if err := fsys.WriteFile(filepath.Join(cacheDir, "latest.json"), sentinel, 0o644); err != nil {
		t.Fatalf("writing cached schema: %v", err)
	}

	pipeline, err := NewValidationPipeline(fsys, []byte(`{"schemaVersion":"latest","modules":{}}`))
	if err != nil {
		t.Fatalf("NewValidationPipeline: %v", err)
	}
	result, err := pipeline.Validate(ValidationOptions{NoNetwork: true})
	if err != nil {
		t.Fatalf("Validate: %v", err)
	}
	if !result.IsValid {
		t.Errorf("Validate rejected a dist-tag version that resolved from cache: %+v", result.Errors)
	}

	foundWarning := false
	for _, warning := range result.Warnings {
		if warning.ID == "schema-version-invalid" {
			foundWarning = true
		}
	}
	if !foundWarning {
		t.Errorf("missing schema-version-invalid warning: %+v", result.Warnings)
	}
}

// TestNewValidationPipelineClonesManifestData verifies that mutating the
// caller's buffer after construction cannot change the bytes that Validate
// checks.
func TestNewValidationPipelineClonesManifestData(t *testing.T) {
	// Inline assertions: this test checks a boolean result and one error,
	// where a golden file would obscure the aliasing scenario under test.
	manifest := []byte(`{"schemaVersion":"2.0.0","modules":[]}`)
	pipeline, err := NewValidationPipeline(platform.NewMapFileSystem(nil), manifest)
	if err != nil {
		t.Fatalf("NewValidationPipeline: %v", err)
	}

	// Same length, so the mutation happens in place.
	copy(manifest, `{"schemaVersion":"2.0.0","modules":{}}`)

	result, err := pipeline.Validate(ValidationOptions{})
	if err != nil {
		t.Fatalf("Validate: %v", err)
	}
	if !result.IsValid {
		t.Errorf("Validate checked the mutated caller buffer: %+v", result.Errors)
	}
}

// Hardening of pre-existing behavior: an unsafe schemaVersion still produces
// the invalid-version warning, but validation falls back to an embedded schema
// instead of using the raw value as a path component.
//
// Inline assertions: this test checks the error and warning IDs for scalar
// version strings, where a golden file would add indirection only.
func TestValidateFallsBackForUnsafeSchemaVersion(t *testing.T) {
	versions := []string{
		"../../../../tmp/x",
		"/tmp/x",
	}

	for _, version := range versions {
		t.Run(version, func(t *testing.T) {
			manifest := []byte(fmt.Sprintf(`{"schemaVersion":%q,"modules":[]}`, version))
			pipeline, err := NewValidationPipeline(platform.NewMapFileSystem(nil), manifest)
			if err != nil {
				t.Fatalf("NewValidationPipeline: %v", err)
			}

			result, err := pipeline.Validate(ValidationOptions{})
			if err != nil {
				t.Fatalf("Validate: %v", err)
			}

			if result.SchemaVersion != version {
				t.Errorf("SchemaVersion = %q, want the raw value %q", result.SchemaVersion, version)
			}
			// An unsafe schemaVersion must never be reported as valid, even when
			// the manifest passes the embedded schema used as a fallback.
			if result.IsValid {
				t.Errorf("Validate reported a manifest with unsafe schemaVersion %q as valid: %+v", version, result.Errors)
			}
			// The only error is the unsafe-version error: the manifest itself is
			// valid under the fallback schema, so no other errors are expected.
			if len(result.Errors) != 1 {
				t.Fatalf("Errors = %+v, want exactly one unsafe schema version error", result.Errors)
			}
			if result.Errors[0].ID != "schema-version-invalid" {
				t.Errorf("Errors[0].ID = %q, want %q", result.Errors[0].ID, "schema-version-invalid")
			}
			if !strings.Contains(result.Errors[0].Message, version) {
				t.Errorf("Errors[0].Message = %q, want it to carry the offending schemaVersion %q", result.Errors[0].Message, version)
			}

			found := false
			for _, warning := range result.Warnings {
				if warning.ID == "schema-version-invalid" {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("missing schema-version-invalid warning: %+v", result.Warnings)
			}
		})
	}

	t.Run("fallback validation still reports schema errors", func(t *testing.T) {
		// The speculative schema rejects additional top-level properties, so
		// these errors prove validation actually ran against the fallback
		// schema rather than being skipped.
		manifest := []byte(`{"schemaVersion":"../../../../tmp/x","modules":[],"unexpected":true}`)
		pipeline, err := NewValidationPipeline(platform.NewMapFileSystem(nil), manifest)
		if err != nil {
			t.Fatalf("NewValidationPipeline: %v", err)
		}

		result, err := pipeline.Validate(ValidationOptions{})
		if err != nil {
			t.Fatalf("Validate: %v", err)
		}

		if result.IsValid {
			t.Errorf("Validate reported a manifest with unsafe schemaVersion as valid: %+v", result.Errors)
		}
		if len(result.Errors) < 2 {
			t.Fatalf("Errors = %+v, want the unsafe schema version error plus fallback schema errors", result.Errors)
		}
		// The unsafe-version error comes first, then any errors produced by the
		// fallback validation pass.
		if result.Errors[0].ID != "schema-version-invalid" {
			t.Errorf("Errors[0].ID = %q, want %q", result.Errors[0].ID, "schema-version-invalid")
		}
		foundFallbackError := false
		for _, issue := range result.Errors[1:] {
			if strings.Contains(issue.Message, "unexpected") {
				foundFallbackError = true
				break
			}
		}
		if !foundFallbackError {
			t.Errorf("fallback validation did not report the unexpected top-level property: %+v", result.Errors)
		}
	})
}
