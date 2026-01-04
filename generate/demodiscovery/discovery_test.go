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
package demodiscovery

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	C "bennypowers.dev/cem/cmd/config"
	"bennypowers.dev/cem/internal/platform/testutil"
	Q "bennypowers.dev/cem/queries"
	W "bennypowers.dev/cem/workspace"
)

// Helper function to run fixture-based tests
func runFixtureTest[T any](t *testing.T, fixtureDir string, test func(t *testing.T, inputPath string, expected *T)) {
	t.Helper()

	fixtures, err := os.ReadDir(filepath.Join("testdata", fixtureDir))
	if err != nil {
		t.Fatalf("Failed to read fixture directory %s: %v", fixtureDir, err)
	}

	for _, fixture := range fixtures {
		if !fixture.IsDir() {
			continue
		}

		t.Run(fixture.Name(), func(t *testing.T) {
			fixturePath := filepath.Join(fixtureDir, fixture.Name())
			inputPath := filepath.Join(fixturePath, "input.html")
			expectedPath := filepath.Join(fixturePath, "expected.json")

			// Read expected output
			var expected T
			expectedBytes := testutil.LoadFixtureFile(t, expectedPath)
			if err := json.Unmarshal(expectedBytes, &expected); err != nil {
				t.Fatalf("Failed to unmarshal expected JSON: %v", err)
			}

			// Run test
			test(t, inputPath, &expected)

			// If update flag is set, write actual output as new golden file
			if *testutil.Update {
				// This would be implemented per test type
				t.Log("Golden file update not yet implemented for this test")
			}
		})
	}
}

func TestExtractDemoMetadata(t *testing.T) {
	runFixtureTest[DemoMetadata](t, "extract-metadata", func(t *testing.T, inputPath string, expected *DemoMetadata) {
		// Create temporary copy for testing
		tmpDir := t.TempDir()
		tmpPath := filepath.Join(tmpDir, "input.html")

		input := testutil.LoadFixtureFile(t, inputPath)
		if err := os.WriteFile(tmpPath, input, 0644); err != nil {
			t.Fatalf("Failed to write test file: %v", err)
		}

		// Extract metadata
		ctx := W.NewFileSystemWorkspaceContext(tmpDir)
		result, err := extractDemoMetadata(ctx, tmpPath)
		if err != nil {
			t.Fatalf("extractDemoMetadata failed: %v", err)
		}

		// Compare results
		if result.URL != expected.URL {
			t.Errorf("URL mismatch: got %q, want %q", result.URL, expected.URL)
		}
		if result.Description != expected.Description {
			t.Errorf("Description mismatch: got %q, want %q", result.Description, expected.Description)
		}
		if len(result.DemoFor) != len(expected.DemoFor) {
			t.Errorf("DemoFor length mismatch: got %d, want %d", len(result.DemoFor), len(expected.DemoFor))
		} else {
			for i, elem := range result.DemoFor {
				if elem != expected.DemoFor[i] {
					t.Errorf("DemoFor[%d] mismatch: got %q, want %q", i, elem, expected.DemoFor[i])
				}
			}
		}
	})
}

func TestGenerateFallbackURL(t *testing.T) {
	ctx := W.NewFileSystemWorkspaceContext(t.TempDir())

	tests := []struct {
		name        string
		config      *C.CemConfig
		demoPath    string
		tagAliases  map[string]string
		expected    string
		expectError bool
	}{
		{
			name: "URLPattern syntax with explicit aliasing",
			config: &C.CemConfig{
				Generate: C.GenerateConfig{
					DemoDiscovery: C.DemoDiscoveryConfig{
						URLPattern:  "/components/:element/demo/:demo.html",
						URLTemplate: "https://site.com/components/{{.element | alias}}/demo/{{.demo}}/",
					},
				},
			},
			demoPath:   "/components/button/demo/primary.html",
			tagAliases: map[string]string{"button": "button-alias"},
			expected:   "https://site.com/components/button-alias/demo/primary/",
		},
		{
			name: "no URLPattern configured",
			config: &C.CemConfig{
				Generate: C.GenerateConfig{
					DemoDiscovery: C.DemoDiscoveryConfig{
						URLPattern:  "",
						URLTemplate: "",
					},
				},
			},
			demoPath:   "/elements/button/demo/primary.html",
			tagAliases: map[string]string{},
			expected:   "",
		},
		{
			name: "invalid URLPattern",
			config: &C.CemConfig{
				Generate: C.GenerateConfig{
					DemoDiscovery: C.DemoDiscoveryConfig{
						URLPattern:  ":::invalid[[[",
						URLTemplate: "https://site.com/",
					},
				},
			},
			demoPath:    "/components/button/demo/primary.html",
			tagAliases:  map[string]string{},
			expectError: true,
		},
		{
			name: "accordion edge case in URLPattern with explicit aliasing",
			config: &C.CemConfig{
				Generate: C.GenerateConfig{
					DemoDiscovery: C.DemoDiscoveryConfig{
						URLPattern:  "/components/:element/demo/:demo.html",
						URLTemplate: "https://site.com/components/{{.element | alias}}/demo/{{.demo}}/",
					},
				},
			},
			demoPath:   "/components/my-accordion-header/demo/primary.html",
			tagAliases: map[string]string{"my-accordion-header": "accordion-header"},
			expected:   "https://site.com/components/accordion-header/demo/primary/",
		},
		{
			name: "template without aliasing uses original values",
			config: &C.CemConfig{
				Generate: C.GenerateConfig{
					DemoDiscovery: C.DemoDiscoveryConfig{
						URLPattern:  "/components/:element/demo/:demo.html",
						URLTemplate: "https://site.com/components/{{.element}}/demo/{{.demo}}/",
					},
				},
			},
			demoPath:   "/components/my-element/demo/primary.html",
			tagAliases: map[string]string{"my-element": "element-alias"},
			expected:   "https://site.com/components/my-element/demo/primary/", // No aliasing applied
		},
		{
			name: "template with slug and upper functions",
			config: &C.CemConfig{
				Generate: C.GenerateConfig{
					DemoDiscovery: C.DemoDiscoveryConfig{
						URLPattern:  "/components/:element/demo/:demo.html",
						URLTemplate: "https://site.com/components/{{.element | slug}}/demo/{{.demo | upper}}/",
					},
				},
			},
			demoPath:   "/components/my-element/demo/primary-demo.html",
			tagAliases: map[string]string{},
			expected:   "https://site.com/components/my-element/demo/PRIMARY-DEMO/",
		},
		{
			name: "template with chained functions",
			config: &C.CemConfig{
				Generate: C.GenerateConfig{
					DemoDiscovery: C.DemoDiscoveryConfig{
						URLPattern:  "/components/:element/demo/:demo.html",
						URLTemplate: "https://site.com/components/{{.element | alias | slug}}/demo/{{.demo | lower}}/",
					},
				},
			},
			demoPath:   "/components/MyElement/demo/Primary-Demo.html",
			tagAliases: map[string]string{"MyElement": "My Custom Element"},
			expected:   "https://site.com/components/my-custom-element/demo/primary-demo/",
		},
		{
			name: "SSG-like routing: index.html should map to root path",
			config: &C.CemConfig{
				Generate: C.GenerateConfig{
					DemoDiscovery: C.DemoDiscoveryConfig{
						URLPattern:  "/elements/:tag/demo/:demo.html",
						URLTemplate: "https://ux.redhat.com/elements/{{.tag | alias | slug}}/demo/{{.demo}}/",
					},
				},
			},
			demoPath:   "/elements/accordion/demo/index.html",
			tagAliases: map[string]string{"accordion": "rh-accordion"},
			expected:   "https://ux.redhat.com/elements/rh-accordion/demo/",
		},
		{
			name: "SSG-like routing: index.html with path containing 'index' directory",
			config: &C.CemConfig{
				Generate: C.GenerateConfig{
					DemoDiscovery: C.DemoDiscoveryConfig{
						URLPattern:  "/elements/:tag/demo/:demo.html",
						URLTemplate: "https://site.com/elements/{{.tag}}/demo/{{.demo}}/",
					},
				},
			},
			demoPath:   "/elements/index/demo/index.html",
			tagAliases: map[string]string{},
			expected:   "https://site.com/elements/index/demo/",
		},
		{
			name: "SSG-like routing: regular demo file should not be affected",
			config: &C.CemConfig{
				Generate: C.GenerateConfig{
					DemoDiscovery: C.DemoDiscoveryConfig{
						URLPattern:  "/elements/:tag/demo/:demo.html",
						URLTemplate: "https://site.com/elements/{{.tag}}/demo/{{.demo}}/",
					},
				},
			},
			demoPath:   "/elements/button/demo/primary.html",
			tagAliases: map[string]string{},
			expected:   "https://site.com/elements/button/demo/primary/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := generateFallbackURL(ctx, tt.config, tt.demoPath, tt.tagAliases)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
				return
			}

			if err != nil {
				t.Fatalf("generateFallbackURL failed: %v", err)
			}

			if result != tt.expected {
				t.Errorf("URL mismatch: got %q, want %q", result, tt.expected)
			}
		})
	}
}

type extractTagsConfig struct {
	ElementAliases map[string]string `json:"elementAliases"`
	DemoPath       string            `json:"demoPath"`
}

func TestExtractDemoTags(t *testing.T) {
	runFixtureTest[[]string](t, "extract-tags", func(t *testing.T, inputPath string, expected *[]string) {
		// Read config
		configPath := filepath.Join(filepath.Dir(inputPath), "config.json")
		configBytes := testutil.LoadFixtureFile(t, configPath)

		var config extractTagsConfig
		if err := json.Unmarshal(configBytes, &config); err != nil {
			t.Fatalf("Failed to unmarshal config JSON: %v", err)
		}

		// Create temporary directory structure
		tmpDir := t.TempDir()
		tmpPath := filepath.Join(tmpDir, config.DemoPath)
		if err := os.MkdirAll(filepath.Dir(tmpPath), 0755); err != nil {
			t.Fatalf("Failed to create directories: %v", err)
		}

		input := testutil.LoadFixtureFile(t, inputPath)
		if err := os.WriteFile(tmpPath, input, 0644); err != nil {
			t.Fatalf("Failed to write test file: %v", err)
		}

		// Extract tags
		ctx := W.NewFileSystemWorkspaceContext(tmpDir)
		result, err := extractDemoTags(ctx, tmpPath, config.ElementAliases)
		if err != nil {
			t.Fatalf("extractDemoTags failed: %v", err)
		}

		// Compare results (order-independent for content-based fallback)
		if len(result) != len(*expected) {
			t.Errorf("Tag count mismatch: got %d (%v), want %d (%v)",
				len(result), result, len(*expected), *expected)
			return
		}

		// Use order-independent comparison for content-based tests
		resultSet := make(map[string]bool)
		for _, tag := range result {
			resultSet[tag] = true
		}
		for _, expectedTag := range *expected {
			if !resultSet[expectedTag] {
				t.Errorf("Expected tag %q not found in result %v", expectedTag, result)
			}
		}
	})
}

func TestExtractParameterValues(t *testing.T) {
	tests := []struct {
		name     string
		demoPath string
		pattern  string
		expected map[string]bool // Use map to avoid order dependency
	}{
		{
			name:     "single parameter",
			demoPath: "/components/button/demo/basic.html",
			pattern:  "/components/:element/demo/:demo.html",
			expected: map[string]bool{"button": true, "basic": true},
		},
		{
			name:     "shop example - exact match",
			demoPath: "/shop/shop/primary.html",
			pattern:  "/shop/:element/:demo.html",
			expected: map[string]bool{"shop": true, "primary": true},
		},
		{
			name:     "shop example - my-shop",
			demoPath: "/shop/my-shop/demo.html",
			pattern:  "/shop/:element/:demo.html",
			expected: map[string]bool{"my-shop": true, "demo": true},
		},
		{
			name:     "no match - path doesn't match pattern",
			demoPath: "/different/path/structure.html",
			pattern:  "/shop/:element/:demo.html",
			expected: map[string]bool{},
		},
		{
			name:     "accordion edge case",
			demoPath: "/components/my-accordion-header/demo/basic.html",
			pattern:  "/components/:element/demo/:demo.html",
			expected: map[string]bool{"my-accordion-header": true, "basic": true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := extractParameterValues(tt.demoPath, tt.pattern)
			if err != nil {
				t.Fatalf("extractParameterValues failed: %v", err)
			}

			if len(result) != len(tt.expected) {
				t.Errorf("Value count mismatch: got %d (%v), want %d (%v)",
					len(result), result, len(tt.expected), tt.expected)
				return
			}

			// Check that all returned values are expected
			for _, value := range result {
				if !tt.expected[value] {
					t.Errorf("Unexpected value: got %q", value)
				}
			}

			// Check that all expected values are present
			resultMap := make(map[string]bool)
			for _, value := range result {
				resultMap[value] = true
			}
			for expectedValue := range tt.expected {
				if !resultMap[expectedValue] {
					t.Errorf("Missing expected value: %q", expectedValue)
				}
			}
		})
	}
}

type newDemoMapConfig struct {
	ElementAliases map[string]string `json:"elementAliases"`
}

func TestNewDemoMap(t *testing.T) {
	// Read config
	configPath := filepath.Join("new-demo-map", "config.json")
	configBytes := testutil.LoadFixtureFile(t, configPath)

	var config newDemoMapConfig
	if err := json.Unmarshal(configBytes, &config); err != nil {
		t.Fatalf("Failed to unmarshal config JSON: %v", err)
	}

	// Create temporary directory and copy fixtures
	tmpDir := t.TempDir()

	demo1Input := testutil.LoadFixtureFile(t, filepath.Join("new-demo-map", "demo1", "input.html"))
	demo1Path := filepath.Join(tmpDir, "demo1.html")
	if err := os.WriteFile(demo1Path, demo1Input, 0644); err != nil {
		t.Fatalf("Failed to write demo1: %v", err)
	}

	demo2Input := testutil.LoadFixtureFile(t, filepath.Join("new-demo-map", "demo2", "input.html"))
	demo2Path := filepath.Join(tmpDir, "demo2.html")
	if err := os.WriteFile(demo2Path, demo2Input, 0644); err != nil {
		t.Fatalf("Failed to write demo2: %v", err)
	}

	// Create workspace context for test
	ctx := W.NewFileSystemWorkspaceContext(tmpDir)

	demoMap, err := NewDemoMap(ctx, []string{demo1Path, demo2Path}, config.ElementAliases)
	if err != nil {
		t.Fatalf("NewDemoMap failed: %v", err)
	}

	// Check rh-button mapping
	buttonDemos := demoMap["rh-button"]
	if len(buttonDemos) != 1 || !strings.Contains(buttonDemos[0], "demo1.html") {
		t.Errorf("Expected rh-button to map to demo1.html, got: %v", buttonDemos)
	}

	// Check rh-card mapping
	cardDemos := demoMap["rh-card"]
	if len(cardDemos) != 1 || !strings.Contains(cardDemos[0], "demo2.html") {
		t.Errorf("Expected rh-card to map to demo2.html, got: %v", cardDemos)
	}
}

func TestMicrodataExtraction(t *testing.T) {
	inputPath := filepath.Join("microdata-extraction", "input.html")
	expectedPath := filepath.Join("microdata-extraction", "expected.txt")

	code := testutil.LoadFixtureFile(t, inputPath)
	expectedBytes := testutil.LoadFixtureFile(t, expectedPath)
	expected := strings.TrimSpace(string(expectedBytes))

	parser := Q.GetHTMLParser()
	defer Q.PutHTMLParser(parser)
	tree := parser.Parse(code, nil)
	defer tree.Close()
	root := tree.RootNode()

	demoFor := extractMicrodata(root, code, "demo-for")

	if demoFor != expected {
		t.Errorf("Expected %q, got %q", expected, demoFor)
	}
}

func TestScriptMarkdownExtraction(t *testing.T) {
	inputPath := filepath.Join("script-markdown-extraction", "input.html")
	expectedPath := filepath.Join("script-markdown-extraction", "expected.txt")

	code := testutil.LoadFixtureFile(t, inputPath)
	expectedBytes := testutil.LoadFixtureFile(t, expectedPath)
	expected := string(expectedBytes)

	parser := Q.GetHTMLParser()
	defer Q.PutHTMLParser(parser)
	tree := parser.Parse(code, nil)
	defer tree.Close()
	root := tree.RootNode()

	description := extractMicrodata(root, code, "description")

	if description != expected {
		t.Errorf("Expected %q, got %q", expected, description)
	}
}

func TestValidateDemoDiscoveryConfig(t *testing.T) {
	tests := []struct {
		name          string
		config        *C.CemConfig
		tagAliases    map[string]string
		expectError   bool
		errorContains string
	}{
		{
			name: "valid URLPattern and template",
			config: &C.CemConfig{
				Generate: C.GenerateConfig{
					DemoDiscovery: C.DemoDiscoveryConfig{
						URLPattern:  "/elements/:tag/demo/:demo.html",
						URLTemplate: "https://site.com/{{.tag | alias}}/{{.demo}}/",
					},
				},
			},
			tagAliases:  map[string]string{"my-element": "element-alias"},
			expectError: false,
		},
		{
			name: "invalid URLPattern - regex syntax",
			config: &C.CemConfig{
				Generate: C.GenerateConfig{
					DemoDiscovery: C.DemoDiscoveryConfig{
						URLPattern:  "elements/(?P<tag>[\\w-]+)/demo/(?P<demo>[\\w-]+).html",
						URLTemplate: "https://site.com/{{.tag}}/{{.demo}}/",
					},
				},
			},
			tagAliases:    map[string]string{},
			expectError:   true,
			errorContains: "invalid demo discovery urlPattern",
		},
		{
			name: "invalid template syntax",
			config: &C.CemConfig{
				Generate: C.GenerateConfig{
					DemoDiscovery: C.DemoDiscoveryConfig{
						URLPattern:  "/elements/:tag/demo/:demo.html",
						URLTemplate: "https://site.com/{{.tag | badfunction}}/{{.demo}}/",
					},
				},
			},
			tagAliases:    map[string]string{},
			expectError:   true,
			errorContains: "invalid demo discovery URLTemplate",
		},
		{
			name: "no demo discovery config",
			config: &C.CemConfig{
				Generate: C.GenerateConfig{
					DemoDiscovery: C.DemoDiscoveryConfig{},
				},
			},
			tagAliases:  map[string]string{},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateDemoDiscoveryConfig(tt.config, tt.tagAliases)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
					return
				}
				if !strings.Contains(err.Error(), tt.errorContains) {
					t.Errorf("Expected error to contain %q, got: %v", tt.errorContains, err)
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error but got: %v", err)
				}
			}
		})
	}
}
