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
	"io/fs"
	"path/filepath"
	"strings"
	"testing"

	"bennypowers.dev/cem/internal/platform"
)

// LSPFixture represents a single LSP test scenario loaded from fixtures
type LSPFixture struct {
	Name            string            // Test scenario name (directory name)
	InputHTML       string            // HTML content from input.html (deprecated: use InputContent)
	InputContent    string            // Content from input.html or input.ts
	InputType       string            // File type: "html" or "ts"
	Manifest        json.RawMessage   // Optional manifest data from manifest.json
	ExpectedMap     map[string]any    // Expected results from expected-*.json or expected.json
	AdditionalFiles map[string]string // Additional source files in the fixture directory (filename → content)
}

// RunLSPFixtures discovers and runs LSP tests from a testdata directory.
// Each subdirectory in testdata/ represents one test scenario with:
//   - input.html or input.ts (required): Content to test
//   - manifest.json (optional): Manifest data for custom elements
//   - expected.json or expected-*.json (optional): Expected results for assertions
//
// All fixture files are preloaded into MapFS at the start — no disk I/O
// during individual test execution.
func RunLSPFixtures(t *testing.T, testdataDir string, testFunc func(*testing.T, *LSPFixture)) {
	t.Helper()

	mfs := LoadTestdataFS(t, testdataDir, "/")

	entries, err := fs.ReadDir(mfs.GetMapFS(), ".")
	if err != nil {
		t.Fatalf("Failed to read testdata directory %s: %v", testdataDir, err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		scenarioName := entry.Name()

		t.Run(scenarioName, func(t *testing.T) {
			fixture := loadLSPFixtureFromMapFS(t, mfs, scenarioName)
			testFunc(t, fixture)
		})
	}
}

// loadLSPFixtureFromMapFS loads a single LSP test fixture from MapFS
func loadLSPFixtureFromMapFS(t *testing.T, mfs *platform.MapFileSystem, scenarioName string) *LSPFixture {
	t.Helper()

	fixture := &LSPFixture{
		Name:            scenarioName,
		ExpectedMap:     make(map[string]any),
		AdditionalFiles: make(map[string]string),
	}

	// Try input.html first, then input.ts, then input.css
	found := false
	for _, candidate := range []struct {
		name     string
		fileType string
	}{
		{"input.html", "html"},
		{"input.ts", "ts"},
		{"input.css", "css"},
	} {
		path := filepath.Join(scenarioName, candidate.name)
		data, err := mfs.ReadFile(path)
		if err == nil {
			fixture.InputContent = string(data)
			fixture.InputHTML = string(data)
			fixture.InputType = candidate.fileType
			found = true
			break
		}
	}

	if !found {
		t.Fatalf("No input.html, input.ts, or input.css found for scenario %s", scenarioName)
	}

	// Load manifest.json (optional)
	manifestPath := filepath.Join(scenarioName, "manifest.json")
	if data, err := mfs.ReadFile(manifestPath); err == nil {
		fixture.Manifest = json.RawMessage(data)
	}

	// Load expected-*.json and additional source files
	scenarioEntries, err := fs.ReadDir(mfs.GetMapFS(), scenarioName)
	if err != nil {
		return fixture
	}

	primaryInput := "input." + fixture.InputType

	for _, entry := range scenarioEntries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		ext := filepath.Ext(name)
		filePath := filepath.Join(scenarioName, name)

		switch {
		case ext == ".json" && name != "manifest.json" && name != "package.json":
			data, err := mfs.ReadFile(filePath)
			if err != nil {
				continue
			}
			var expectedData any
			if err := json.Unmarshal(data, &expectedData); err != nil {
				t.Logf("Warning: failed to parse %s: %v", name, err)
				continue
			}
			var key string
			if name == "expected.json" {
				key = "expected"
			} else if strings.HasPrefix(name, "expected-") {
				key = name[len("expected-") : len(name)-len(".json")]
			} else {
				key = name[:len(name)-len(".json")]
			}
			fixture.ExpectedMap[key] = expectedData

		case (ext == ".ts" || ext == ".js" || ext == ".html" || ext == ".css") && name != primaryInput:
			data, err := mfs.ReadFile(filePath)
			if err != nil {
				t.Logf("Warning: failed to read additional file %s: %v", name, err)
				continue
			}
			fixture.AdditionalFiles[name] = string(data)
		}
	}

	return fixture
}

// GetExpected returns typed expected data from the fixture
func (f *LSPFixture) GetExpected(key string, v any) error {
	data, ok := f.ExpectedMap[key]
	if !ok {
		return nil // No expected data for this key
	}

	// Re-marshal and unmarshal to convert to desired type
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, v)
}
