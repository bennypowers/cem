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
	"os"
	"path/filepath"
	"testing"
)

// LSPFixture represents a single LSP test scenario loaded from fixtures
type LSPFixture struct {
	Name         string          // Test scenario name (directory name)
	InputHTML    string          // HTML content from input.html (deprecated: use InputContent)
	InputContent string          // Content from input.html or input.ts
	InputType    string          // File type: "html" or "ts"
	Manifest     json.RawMessage // Optional manifest data from manifest.json
	ExpectedMap  map[string]any  // Expected results from expected-*.json or expected.json
}

// RunLSPFixtures discovers and runs LSP tests from a testdata directory.
// Each subdirectory in testdata/ represents one test scenario with:
//   - input.html or input.ts (required): Content to test
//   - manifest.json (optional): Manifest data for custom elements
//   - expected.json or expected-*.json (optional): Expected results for assertions
//
// Example usage:
//
//	RunLSPFixtures(t, "testdata/attribute-diagnostics", func(t *testing.T, fixture *LSPFixture) {
//	    // Your test logic here using fixture.InputContent, fixture.Manifest, etc.
//	})
func RunLSPFixtures(t *testing.T, testdataDir string, testFunc func(*testing.T, *LSPFixture)) {
	t.Helper()

	entries, err := os.ReadDir(testdataDir)
	if err != nil {
		t.Fatalf("Failed to read testdata directory %s: %v", testdataDir, err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		scenarioName := entry.Name()
		scenarioDir := filepath.Join(testdataDir, scenarioName)

		t.Run(scenarioName, func(t *testing.T) {
			fixture := loadLSPFixture(t, scenarioDir, scenarioName)
			testFunc(t, fixture)
		})
	}
}

// loadLSPFixture loads a single LSP test fixture from a scenario directory
func loadLSPFixture(t *testing.T, scenarioDir, scenarioName string) *LSPFixture {
	t.Helper()

	fixture := &LSPFixture{
		Name:        scenarioName,
		ExpectedMap: make(map[string]any),
	}

	// Try input.html first, then input.ts
	inputPath := filepath.Join(scenarioDir, "input.html")
	inputBytes, err := os.ReadFile(inputPath)
	if err != nil {
		// Try TypeScript input
		inputPath = filepath.Join(scenarioDir, "input.ts")
		inputBytes, err = os.ReadFile(inputPath)
		if err != nil {
			t.Fatalf("Failed to read input.html or input.ts for scenario %s: %v", scenarioName, err)
		}
		fixture.InputType = "ts"
	} else {
		fixture.InputType = "html"
	}

	fixture.InputContent = string(inputBytes)
	fixture.InputHTML = string(inputBytes) // Backward compatibility

	// Load manifest.json (optional)
	manifestPath := filepath.Join(scenarioDir, "manifest.json")
	if manifestBytes, err := os.ReadFile(manifestPath); err == nil {
		fixture.Manifest = json.RawMessage(manifestBytes)
	}

	// Load all expected-*.json files (optional)
	entries, err := os.ReadDir(scenarioDir)
	if err == nil {
		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}
			name := entry.Name()
			if filepath.Ext(name) == ".json" && filepath.Base(name) != "manifest.json" {
				fullPath := filepath.Join(scenarioDir, name)
				data, err := os.ReadFile(fullPath)
				if err != nil {
					continue
				}
				var expectedData any
				if err := json.Unmarshal(data, &expectedData); err != nil {
					t.Logf("Warning: failed to parse %s: %v", name, err)
					continue
				}
				// Handle both "expected.json" and "expected-*.json" patterns
				var key string
				if name == "expected.json" {
					key = "expected"
				} else if len(name) > len("expected-") && name[:len("expected-")] == "expected-" {
					key = name[len("expected-") : len(name)-len(".json")]
				} else {
					// Other JSON files, use filename without extension
					key = name[:len(name)-len(".json")]
				}
				fixture.ExpectedMap[key] = expectedData
			}
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
