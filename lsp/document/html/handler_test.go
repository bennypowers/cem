package html

import (
	"testing"

	"bennypowers.dev/cem/lsp/types"
	"github.com/stretchr/testify/assert"
)

func TestExtractImportPath(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		expected string
	}{
		{
			name:     "single-quoted bare import",
			line:     "import 'foo'",
			expected: "foo",
		},
		{
			name:     "double-quoted bare import",
			line:     `import "foo"`,
			expected: "foo",
		},
		{
			name:     "from keyword with single quotes",
			line:     "import { x } from 'foo'",
			expected: "foo",
		},
		{
			name:     "from keyword with double quotes",
			line:     `import { x } from "foo"`,
			expected: "foo",
		},
		{
			name:     "default import with from keyword",
			line:     "import MyComponent from '@scope/my-component'",
			expected: "@scope/my-component",
		},
		{
			name:     "star import",
			line:     "import * as utils from './utils.js'",
			expected: "./utils.js",
		},
		{
			name:     "no quotes returns empty",
			line:     "import foo",
			expected: "",
		},
		{
			name:     "empty string returns empty",
			line:     "",
			expected: "",
		},
		{
			name:     "path with special characters",
			line:     "import '@foo/bar-baz_qux.js'",
			expected: "@foo/bar-baz_qux.js",
		},
		{
			name:     "path with relative prefix",
			line:     "import './components/my-element.js'",
			expected: "./components/my-element.js",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractImportPath(tt.line)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestExtractDynamicImportPath(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		expected string
	}{
		{
			name:     "single-quoted dynamic import",
			line:     "import('./module.js')",
			expected: "./module.js",
		},
		{
			name:     "double-quoted dynamic import",
			line:     `import("./module.js")`,
			expected: "./module.js",
		},
		{
			name:     "nested in expression",
			line:     "const m = await import('./lazy.js');",
			expected: "./lazy.js",
		},
		{
			name:     "nested in then",
			line:     "import('./lazy.js').then(m => m.default)",
			expected: "./lazy.js",
		},
		{
			name:     "no closing single quote returns empty",
			line:     "import('./no-end",
			expected: "",
		},
		{
			name:     "no closing double quote returns empty",
			line:     `import("./no-end`,
			expected: "",
		},
		{
			name:     "empty string returns empty",
			line:     "",
			expected: "",
		},
		{
			name:     "no import( present returns empty",
			line:     "const x = 42;",
			expected: "",
		},
		{
			name:     "scoped package path",
			line:     "import('@scope/pkg')",
			expected: "@scope/pkg",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractDynamicImportPath(tt.line)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestParseImportStatements(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected []types.ImportStatement
	}{
		{
			name:    "single static import",
			content: "import './my-element.js';",
			expected: []types.ImportStatement{
				{ImportPath: "./my-element.js", Type: "static"},
			},
		},
		{
			name:    "multiple static imports",
			content: "import './a.js';\nimport './b.js';",
			expected: []types.ImportStatement{
				{ImportPath: "./a.js", Type: "static"},
				{ImportPath: "./b.js", Type: "static"},
			},
		},
		{
			name:    "single dynamic import",
			content: "import('./lazy.js');",
			expected: []types.ImportStatement{
				{ImportPath: "./lazy.js", Type: "dynamic"},
			},
		},
		{
			name:    "mixed static and dynamic imports",
			content: "import './a.js';\nimport('./b.js');",
			expected: []types.ImportStatement{
				{ImportPath: "./a.js", Type: "static"},
				{ImportPath: "./b.js", Type: "dynamic"},
			},
		},
		{
			name:     "no imports",
			content:  "const x = 42;\nconsole.log(x);",
			expected: nil,
		},
		{
			name:     "empty content",
			content:  "",
			expected: nil,
		},
		{
			name:     "malformed import with no quotes",
			content:  "import foo\nimport bar",
			expected: nil,
		},
		{
			name:    "import with from keyword",
			content: "import { html } from 'lit';",
			expected: []types.ImportStatement{
				{ImportPath: "lit", Type: "static"},
			},
		},
		{
			name:    "dynamic import nested in await",
			content: "const m = await import('./lazy.js');",
			expected: []types.ImportStatement{
				{ImportPath: "./lazy.js", Type: "dynamic"},
			},
		},
		{
			name:    "static import line also containing dynamic import",
			content: "import { load } from './loader.js';\nload(() => import('./chunk.js'));",
			expected: []types.ImportStatement{
				{ImportPath: "./loader.js", Type: "static"},
				{ImportPath: "./chunk.js", Type: "dynamic"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseImportStatements(tt.content)
			assert.Equal(t, tt.expected, result)
		})
	}
}
