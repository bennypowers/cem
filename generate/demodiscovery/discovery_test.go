package demodiscovery

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	C "bennypowers.dev/cem/cmd/config"
	Q "bennypowers.dev/cem/generate/queries"
)

func TestExtractDemoMetadata(t *testing.T) {
	tests := []struct {
		name     string
		html     string
		expected DemoMetadata
	}{
		{
			name: "microdata URL and description",
			html: `<!DOCTYPE html>
<html>
<head>
<meta itemprop="demo-url" content="/elements/call-to-action/demo/">
<meta itemprop="description" content="Primary variant demonstration">
</head>
<body>
<pf-call-to-action variant="primary">Click me</pf-call-to-action>
</body>
</html>`,
			expected: DemoMetadata{
				URL:         "/elements/call-to-action/demo/",
				Description: "Primary variant demonstration",
				DemoFor:     nil,
			},
		},
		{
			name: "microdata with demo-for association",
			html: `<!DOCTYPE html>
<html>
<head>
<meta itemprop="demo-url" content="/elements/button/demo/">
<meta itemprop="demo-for" content="rh-button pf-button">
</head>
<body>
<rh-button>Click me</rh-button>
</body>
</html>`,
			expected: DemoMetadata{
				URL:         "/elements/button/demo/",
				Description: "",
				DemoFor:     []string{"rh-button", "pf-button"},
			},
		},
		{
			name: "script tag with markdown description",
			html: `<!DOCTYPE html>
<html>
<head>
<meta itemprop="demo-url" content="/elements/card/demo/">
<script type="text/markdown" itemprop="description">
# Card Demo
Showcases different card variants with accessibility features.
</script>
</head>
<body>
<rh-card>Content</rh-card>
</body>
</html>`,
			expected: DemoMetadata{
				URL:         "/elements/card/demo/",
				Description: "# Card Demo\nShowcases different card variants with accessibility features.",
				DemoFor:     nil,
			},
		},
		{
			name: "no metadata",
			html: `<!DOCTYPE html>
<html>
<body>
<rh-button>Click me</rh-button>
</body>
</html>`,
			expected: DemoMetadata{
				URL:         "",
				Description: "",
				DemoFor:     nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create temporary file
			tmpDir := t.TempDir()
			filePath := filepath.Join(tmpDir, "demo.html")
			err := os.WriteFile(filePath, []byte(tt.html), 0644)
			if err != nil {
				t.Fatalf("Failed to write test file: %v", err)
			}

			// Extract metadata
			result, err := extractDemoMetadata(filePath)
			if err != nil {
				t.Fatalf("extractDemoMetadata failed: %v", err)
			}

			// Compare results
			if result.URL != tt.expected.URL {
				t.Errorf("URL mismatch: got %q, want %q", result.URL, tt.expected.URL)
			}
			if result.Description != tt.expected.Description {
				t.Errorf("Description mismatch: got %q, want %q", result.Description, tt.expected.Description)
			}
			if len(result.DemoFor) != len(tt.expected.DemoFor) {
				t.Errorf("DemoFor length mismatch: got %d, want %d", len(result.DemoFor), len(tt.expected.DemoFor))
			} else {
				for i, elem := range result.DemoFor {
					if elem != tt.expected.DemoFor[i] {
						t.Errorf("DemoFor[%d] mismatch: got %q, want %q", i, elem, tt.expected.DemoFor[i])
					}
				}
			}
		})
	}
}

func TestGenerateFallbackURL(t *testing.T) {
	tests := []struct {
		name        string
		config      *C.CemConfig
		demoPath    string
		tagAliases  map[string]string
		expected    string
		expectError bool
	}{
		{
			name: "URLPattern syntax",
			config: &C.CemConfig{
				Generate: C.GenerateConfig{
					DemoDiscovery: C.DemoDiscoveryConfig{
						URLPattern:  "/components/:element/demo/:demo.html",
						URLTemplate: "https://site.com/components/{{.element}}/demo/{{.demo}}/",
					},
				},
			},
			demoPath:   "/components/button/demo/primary.html",
			tagAliases: map[string]string{"rh-button": "button"},
			expected:   "https://site.com/components/button/demo/primary/",
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
			name: "accordion edge case in URLPattern",
			config: &C.CemConfig{
				Generate: C.GenerateConfig{
					DemoDiscovery: C.DemoDiscoveryConfig{
						URLPattern:  "/components/:element/demo/:demo.html",
						URLTemplate: "https://site.com/components/{{.element}}/demo/{{.demo}}/",
					},
				},
			},
			demoPath:   "/components/my-accordion-header/demo/primary.html",
			tagAliases: map[string]string{"my-accordion-header": "accordion-header"},
			expected:   "https://site.com/components/accordion-header/demo/primary/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := generateFallbackURL(tt.config, tt.demoPath, tt.tagAliases)

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

func TestExtractDemoTags(t *testing.T) {
	tests := []struct {
		name           string
		html           string
		elementAliases map[string]string
		demoPath       string
		expected       []string
	}{
		{
			name: "explicit microdata association",
			html: `<!DOCTYPE html>
<html>
<head>
<meta itemprop="demo-for" content="rh-button pf-button">
</head>
<body>
<rh-card>Should be ignored</rh-card>
</body>
</html>`,
			elementAliases: map[string]string{"rh-button": "button"},
			demoPath:       "/components/button/demo/primary.html",
			expected:       []string{"rh-button", "pf-button"},
		},
		{
			name: "path-based association",
			html: `<!DOCTYPE html>
<html>
<body>
<div>No custom elements</div>
</body>
</html>`,
			elementAliases: map[string]string{"rh-button": "button", "rh-card": "card"},
			demoPath:       "/components/button/demo/primary.html",
			expected:       []string{"rh-button"},
		},
		{
			name: "content-based fallback",
			html: `<!DOCTYPE html>
<html>
<body>
<rh-button>Click me</rh-button>
<rh-card>Content</rh-card>
</body>
</html>`,
			elementAliases: map[string]string{"rh-button": "button"},
			demoPath:       "/demos/mixed/example.html",      // Path that doesn't match aliases
			expected:       []string{"rh-button", "rh-card"}, // Should find both from content
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create temporary file
			tmpDir := t.TempDir()
			filePath := filepath.Join(tmpDir, tt.demoPath)
			err := os.MkdirAll(filepath.Dir(filePath), 0755)
			if err != nil {
				t.Fatalf("Failed to create directories: %v", err)
			}
			err = os.WriteFile(filePath, []byte(tt.html), 0644)
			if err != nil {
				t.Fatalf("Failed to write test file: %v", err)
			}

			// Extract tags
			result, err := extractDemoTags(filePath, tt.elementAliases)
			if err != nil {
				t.Fatalf("extractDemoTags failed: %v", err)
			}

			// t.Logf("Test %s: got tags %v, expected %v", tt.name, result, tt.expected)

			// Compare results
			if len(result) != len(tt.expected) {
				t.Errorf("Tag count mismatch: got %d (%v), want %d (%v)",
					len(result), result, len(tt.expected), tt.expected)
				return
			}

			for i, tag := range result {
				if tag != tt.expected[i] {
					t.Errorf("Tag[%d] mismatch: got %q, want %q", i, tag, tt.expected[i])
				}
			}
		})
	}
}

func TestExtractParameterPositions(t *testing.T) {
	tests := []struct {
		name     string
		pattern  string
		expected []int
	}{
		{
			name:     "single parameter",
			pattern:  "/components/:element/demo/",
			expected: []int{1},
		},
		{
			name:     "multiple parameters",
			pattern:  "/shop/:element/:demo.html",
			expected: []int{1, 2},
		},
		{
			name:     "no parameters",
			pattern:  "/static/demo/page.html",
			expected: []int{},
		},
		{
			name:     "parameter with extension",
			pattern:  "/docs/:page.html",
			expected: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := extractParameterPositions(tt.pattern)
			if err != nil {
				t.Fatalf("extractParameterPositions failed: %v", err)
			}

			if len(result) != len(tt.expected) {
				t.Errorf("Position count mismatch: got %v, want %v", result, tt.expected)
				return
			}

			for i, pos := range result {
				if pos != tt.expected[i] {
					t.Errorf("Position[%d] mismatch: got %d, want %d", i, pos, tt.expected[i])
				}
			}
		})
	}
}

func TestExtractPathBasedTagsWithPattern(t *testing.T) {
	tests := []struct {
		name           string
		demoPath       string
		urlPattern     string
		elementAliases map[string]string
		expected       []string
	}{
		{
			name:       "accordion exact match in parameter position",
			demoPath:   "/components/my-accordion/demo/basic.html",
			urlPattern: "/components/:element/demo/:demo.html",
			elementAliases: map[string]string{
				"my-accordion":        "accordion",
				"my-accordion-header": "accordion-header",
			},
			expected: []string{}, // "accordion" doesn't match "my-accordion" exactly
		},
		{
			name:       "alias matches parameter position",
			demoPath:   "/components/accordion/demo/basic.html",
			urlPattern: "/components/:element/demo/:demo.html",
			elementAliases: map[string]string{
				"my-accordion":        "accordion",
				"my-accordion-header": "accordion-header",
			},
			expected: []string{"my-accordion"},
		},
		{
			name:       "shop example - my-shop should not match shop",
			demoPath:   "/shop/my-shop/demo.html",
			urlPattern: "/shop/:element/:demo.html",
			elementAliases: map[string]string{
				"my-shop": "shop",
			},
			expected: []string{}, // "shop" doesn't match "my-shop"
		},
		{
			name:       "shop example - shop matches exactly",
			demoPath:   "/shop/shop/demo.html",
			urlPattern: "/shop/:element/:demo.html",
			elementAliases: map[string]string{
				"my-shop": "shop",
			},
			expected: []string{"my-shop"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := extractPathBasedTagsWithPattern(tt.demoPath, tt.urlPattern, tt.elementAliases)
			if err != nil {
				t.Fatalf("extractPathBasedTagsWithPattern failed: %v", err)
			}

			if len(result) != len(tt.expected) {
				t.Errorf("Tag count mismatch: got %d (%v), want %d (%v)",
					len(result), result, len(tt.expected), tt.expected)
				return
			}

			for i, tag := range result {
				if tag != tt.expected[i] {
					t.Errorf("Tag[%d] mismatch: got %q, want %q", i, tag, tt.expected[i])
				}
			}
		})
	}
}

func TestExtractPathBasedTags(t *testing.T) {
	tests := []struct {
		name           string
		demoPath       string
		elementAliases map[string]string
		expected       []string
	}{
		{
			name:     "exact alias match",
			demoPath: "/components/button/demo/primary.html",
			elementAliases: map[string]string{
				"rh-button": "button",
				"rh-card":   "card",
			},
			expected: []string{"rh-button"},
		},
		{
			name:     "multiple matches with scoring",
			demoPath: "/components/button-card/demo/with-button.html",
			elementAliases: map[string]string{
				"rh-button": "button",
				"rh-card":   "card",
			},
			expected: []string{"rh-button", "rh-card"}, // button scores higher due to exact directory match
		},
		{
			name:     "no matches",
			demoPath: "/components/table/demo/primary.html",
			elementAliases: map[string]string{
				"rh-button": "button",
				"rh-card":   "card",
			},
			expected: []string{},
		},
		{
			name:     "accordion edge case - exact match",
			demoPath: "/components/my-accordion/demo/basic.html",
			elementAliases: map[string]string{
				"my-accordion":        "accordion",
				"my-accordion-header": "accordion-header",
			},
			expected: []string{"my-accordion"},
		},
		{
			name:     "accordion edge case - header should not match accordion",
			demoPath: "/components/my-accordion-header/demo/basic.html",
			elementAliases: map[string]string{
				"my-accordion":        "accordion",
				"my-accordion-header": "accordion-header",
			},
			expected: []string{"my-accordion-header"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractPathBasedTags(tt.demoPath, tt.elementAliases)

			if len(result) != len(tt.expected) {
				t.Errorf("Tag count mismatch: got %d (%v), want %d (%v)",
					len(result), result, len(tt.expected), tt.expected)
				return
			}

			for i, tag := range result {
				if tag != tt.expected[i] {
					t.Errorf("Tag[%d] mismatch: got %q, want %q", i, tag, tt.expected[i])
				}
			}
		})
	}
}

func TestNewDemoMap(t *testing.T) {
	// Create temporary demo files
	tmpDir := t.TempDir()

	// Demo 1: Explicit microdata
	demo1Path := filepath.Join(tmpDir, "demo1.html")
	demo1Content := `<!DOCTYPE html>
<html>
<head>
<meta itemprop="demo-for" content="rh-button">
</head>
<body><rh-button>Click</rh-button></body>
</html>`
	err := os.WriteFile(demo1Path, []byte(demo1Content), 0644)
	if err != nil {
		t.Fatalf("Failed to write demo1: %v", err)
	}

	// Demo 2: Content-based discovery
	demo2Path := filepath.Join(tmpDir, "demo2.html")
	demo2Content := `<!DOCTYPE html>
<html>
<body><rh-card>Content</rh-card></body>
</html>`
	err = os.WriteFile(demo2Path, []byte(demo2Content), 0644)
	if err != nil {
		t.Fatalf("Failed to write demo2: %v", err)
	}

	elementAliases := map[string]string{
		"rh-button": "button",
		"rh-card":   "card",
	}

	demoMap, err := NewDemoMap([]string{demo1Path, demo2Path}, elementAliases)
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
	html := `<!DOCTYPE html>
<html>
<head>
<meta itemprop="demo-for" content="rh-button pf-button">
</head>
<body>
<rh-card>Should be ignored</rh-card>
</body>
</html>`

	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "test.html")
	err := os.WriteFile(filePath, []byte(html), 0644)
	if err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	code, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read test file: %v", err)
	}

	parser := Q.GetHTMLParser()
	defer Q.PutHTMLParser(parser)
	tree := parser.Parse(code, nil)
	defer tree.Close()
	root := tree.RootNode()

	demoFor := extractMicrodata(root, code, "demo-for")

	if demoFor != "rh-button pf-button" {
		t.Errorf("Expected 'rh-button pf-button', got %q", demoFor)
	}
}

func TestScriptMarkdownExtraction(t *testing.T) {
	html := `<!DOCTYPE html>
<html>
<head>
<script type="text/markdown" itemprop="description">
# Card Demo
Showcases different card variants with accessibility features.
</script>
</head>
<body>
<rh-card>Content</rh-card>
</body>
</html>`

	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "test.html")
	err := os.WriteFile(filePath, []byte(html), 0644)
	if err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	code, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read test file: %v", err)
	}

	parser := Q.GetHTMLParser()
	defer Q.PutHTMLParser(parser)
	tree := parser.Parse(code, nil)
	defer tree.Close()
	root := tree.RootNode()

	description := extractMicrodata(root, code, "description")
	expected := "# Card Demo\nShowcases different card variants with accessibility features."

	if description != expected {
		t.Errorf("Expected %q, got %q", expected, description)
	}
}
