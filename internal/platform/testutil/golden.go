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

package testutil

import (
	"encoding/json"
	"flag"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/nsf/jsondiff"
)

// Update is the global --update flag for regenerating golden files
var Update = flag.Bool("update", false, "update golden files")

var ansiRegex = regexp.MustCompile(`\x1b\[[0-9;]*m`)

// GoldenOptions configures golden file comparison behavior
type GoldenOptions struct {
	// Dir specifies the directory for golden files (default: "goldens")
	Dir string
	// Extension specifies the file extension (default: inferred from name or ".txt")
	Extension string
	// StripANSI removes ANSI color codes before comparison
	StripANSI bool
	// NormalizeEOL normalizes line endings to \n before comparison
	NormalizeEOL bool
	// UseJSONDiff uses jsondiff for JSON comparison instead of string equality
	UseJSONDiff bool
	// AutoName uses t.Name() for the golden filename (ignores name parameter)
	AutoName bool
}

// CheckGolden compares actual output against a golden file.
// If --update flag is set, it updates the golden file instead of comparing.
//
// Example:
//
//	CheckGolden(t, "output", actual, GoldenOptions{Dir: "testdata", Extension: ".html"})
//	CheckGolden(t, "", actual, GoldenOptions{AutoName: true, StripANSI: true})
func CheckGolden(t *testing.T, name string, actual []byte, opts ...GoldenOptions) {
	t.Helper()

	// Resolve options with defaults
	opt := GoldenOptions{
		Dir:       "goldens",
		Extension: "",
	}
	if len(opts) > 0 {
		opt = opts[0]
	}

	// Use test name if AutoName is enabled
	if opt.AutoName {
		name = filepath.Base(t.Name())
	}

	// Infer extension from name if not explicitly set
	if opt.Extension == "" {
		ext := filepath.Ext(name)
		if ext != "" {
			opt.Extension = ext
			name = strings.TrimSuffix(name, ext)
		} else {
			opt.Extension = ".txt"
		}
	}

	goldenPath := filepath.Join(opt.Dir, name+opt.Extension)

	// Update mode: write golden file and return
	if *Update {
		if err := os.MkdirAll(filepath.Dir(goldenPath), 0755); err != nil {
			t.Fatalf("failed to create golden directory: %v", err)
		}
		if err := os.WriteFile(goldenPath, actual, 0644); err != nil {
			t.Fatalf("failed to update golden file: %v", err)
		}
		t.Logf("Updated golden file: %s", goldenPath)
		return
	}

	// Read golden file
	expected, err := os.ReadFile(goldenPath)
	if err != nil {
		t.Fatalf("golden file missing: %s (run with -update)\nerror: %v", goldenPath, err)
	}

	// Apply transformations
	actualProcessed := actual
	expectedProcessed := expected

	if opt.StripANSI {
		actualProcessed = []byte(StripANSI(string(actualProcessed)))
		expectedProcessed = []byte(StripANSI(string(expectedProcessed)))
	}

	if opt.NormalizeEOL {
		actualProcessed = []byte(strings.ReplaceAll(string(actualProcessed), "\r\n", "\n"))
		expectedProcessed = []byte(strings.ReplaceAll(string(expectedProcessed), "\r\n", "\n"))
	}

	// Compare
	if opt.UseJSONDiff {
		// Validate both are valid JSON first
		var jsExpected, jsActual any
		if err := json.Unmarshal(expectedProcessed, &jsExpected); err != nil {
			t.Fatalf("expected golden file is invalid JSON: %v", err)
		}
		if err := json.Unmarshal(actualProcessed, &jsActual); err != nil {
			t.Fatalf("actual output is invalid JSON: %v\noutput:\n%s", err, actualProcessed)
		}

		// Compare JSON semantically
		if string(expectedProcessed) != string(actualProcessed) {
			options := jsondiff.DefaultConsoleOptions()
			diff, str := jsondiff.Compare(expectedProcessed, actualProcessed, &options)
			if diff == jsondiff.FullMatch {
				t.Logf("Semantic match, string mismatch: %s", str)
			} else {
				t.Errorf("%s\n%s", diff, str)
				t.Log("Run 'make update' to update golden files")
			}
		}
	} else {
		// String comparison
		if string(expectedProcessed) != string(actualProcessed) {
			t.Errorf("Output differs from golden file.\nExpected:\n%s\n\nGot:\n%s",
				string(expectedProcessed), string(actualProcessed))
			t.Log("Run 'make update' to update golden files")
		}
	}
}

// LoadFixture loads a single fixture file, trying multiple common directories.
// Returns the file contents or fails the test if not found.
//
// Example:
//
//	data := LoadFixture(t, "chrome-rendering/basic-demo.html")
func LoadFixture(t *testing.T, path string) []byte {
	t.Helper()

	// Try multiple common fixture directories
	possibleDirs := []string{
		"fixtures",
		"testdata",
		"test/fixtures",
		filepath.Join("..", "fixtures"),
		filepath.Join("..", "testdata"),
	}

	for _, dir := range possibleDirs {
		fullPath := filepath.Join(dir, path)
		if data, err := os.ReadFile(fullPath); err == nil {
			return data
		}
	}

	t.Fatalf("Fixture not found: %s (tried: %v)", path, possibleDirs)
	return nil
}

// LoadJSONFixture loads and unmarshals a JSON fixture into the provided value.
//
// Example:
//
//	var pkg manifest.Package
//	LoadJSONFixture(t, "custom-elements.json", &pkg)
func LoadJSONFixture(t *testing.T, path string, v any) {
	t.Helper()

	data := LoadFixture(t, path)
	if err := json.Unmarshal(data, v); err != nil {
		t.Fatalf("Failed to unmarshal JSON fixture %s: %v", path, err)
	}
}

// StripANSI removes ANSI color codes from text
func StripANSI(s string) string {
	return ansiRegex.ReplaceAllString(s, "")
}
