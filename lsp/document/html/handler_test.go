package html

import (
	"testing"

	"bennypowers.dev/cem/internal/platform/testutil"
	treesitter "bennypowers.dev/cem/internal/treesitter"
	"bennypowers.dev/cem/lsp/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Inline: pure function, table-driven
// extractImportPath, extractDynamicImportPath, and parseImportStatements are
// pure string parsers. ParseScriptTags, ParseImportMap, and FindCustomElements
// use fixtures but validate structured output with scalar assertions that
// would be awkward as goldens (nested struct fields, map lookups).

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

func TestParseScriptTags(t *testing.T) {
	qm, err := treesitter.NewQueryManager(treesitter.LSPQueries())
	require.NoError(t, err)
	defer qm.Close()

	handler, err := NewHandler(qm)
	require.NoError(t, err)
	defer handler.Close()

	mfs := testutil.LoadTestdataFS(t, "testdata", "/")

	tests := []struct {
		name     string
		fixture  string
		validate func(t *testing.T, tags []types.ScriptTag)
	}{
		{
			name:    "module script with src",
			fixture: "/script-tags/module-script.html",
			validate: func(t *testing.T, tags []types.ScriptTag) {
				require.Len(t, tags, 1)
				tag := tags[0]
				assert.Equal(t, "module", tag.Type)
				assert.Equal(t, "./app.js", tag.Src)
				assert.True(t, tag.IsModule)
				assert.Empty(t, tag.Imports, "external script should have no inline imports")
			},
		},
		{
			name:    "inline script with imports",
			fixture: "/script-tags/inline-script.html",
			validate: func(t *testing.T, tags []types.ScriptTag) {
				require.Len(t, tags, 1)
				tag := tags[0]
				assert.Equal(t, "module", tag.Type)
				assert.Empty(t, tag.Src, "inline script should have no src")
				assert.True(t, tag.IsModule)
				require.Len(t, tag.Imports, 2, "should parse both import statements")
				assert.Equal(t, "lit", tag.Imports[0].ImportPath)
				assert.Equal(t, "static", tag.Imports[0].Type)
				assert.Equal(t, "./my-element.js", tag.Imports[1].ImportPath)
				assert.Equal(t, "static", tag.Imports[1].Type)
			},
		},
		{
			name:    "no scripts",
			fixture: "/script-tags/no-scripts.html",
			validate: func(t *testing.T, tags []types.ScriptTag) {
				assert.Empty(t, tags)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			content := testutil.ReadFixture(t, mfs, tt.fixture)
			doc := handler.CreateDocument("file:///test.html", string(content), 1)
			defer doc.Close()
			tags, err := handler.ParseScriptTags(doc)
			require.NoError(t, err)
			tt.validate(t, tags)
		})
	}
}

func TestParseImportMap(t *testing.T) {
	qm, err := treesitter.NewQueryManager(treesitter.LSPQueries())
	require.NoError(t, err)
	defer qm.Close()

	handler, err := NewHandler(qm)
	require.NoError(t, err)
	defer handler.Close()

	mfs := testutil.LoadTestdataFS(t, "testdata", "/")

	tests := []struct {
		name     string
		fixture  string
		validate func(t *testing.T, importMap map[string]string, err error)
	}{
		{
			name:    "valid importmap",
			fixture: "/import-map/valid-importmap.html",
			validate: func(t *testing.T, importMap map[string]string, err error) {
				require.NoError(t, err)
				require.NotNil(t, importMap)
				assert.Len(t, importMap, 2)
				assert.Equal(t, "/vendor/lit.js", importMap["lit"])
				assert.Equal(t, "/vendor/scope-pkg.js", importMap["@scope/pkg"])
			},
		},
		{
			name:    "no importmap",
			fixture: "/import-map/no-importmap.html",
			validate: func(t *testing.T, importMap map[string]string, err error) {
				require.NoError(t, err)
				assert.Nil(t, importMap)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			content := testutil.ReadFixture(t, mfs, tt.fixture)
			doc := handler.CreateDocument("file:///test.html", string(content), 1)
			defer doc.Close()
			importMap, err := handler.ParseImportMap(doc)
			tt.validate(t, importMap, err)
		})
	}
}

func TestFindCustomElements(t *testing.T) {
	qm, err := treesitter.NewQueryManager(treesitter.LSPQueries())
	require.NoError(t, err)
	defer qm.Close()

	handler, err := NewHandler(qm)
	require.NoError(t, err)
	defer handler.Close()

	mfs := testutil.LoadTestdataFS(t, "testdata", "/")

	tests := []struct {
		name     string
		fixture  string
		validate func(t *testing.T, elements []types.CustomElementMatch)
	}{
		{
			name:    "with attributes",
			fixture: "/custom-elements/with-attributes.html",
			validate: func(t *testing.T, elements []types.CustomElementMatch) {
				require.Len(t, elements, 2)
				myEl := elements[0]
				assert.Equal(t, "my-element", myEl.TagName)
				require.Contains(t, myEl.Attributes, "disabled")
				require.Contains(t, myEl.Attributes, "variant")
				assert.Equal(t, "primary", myEl.Attributes["variant"].Value)
				otherEl := elements[1]
				assert.Equal(t, "other-element", otherEl.TagName)
				require.Contains(t, otherEl.Attributes, "name")
				assert.Equal(t, "test", otherEl.Attributes["name"].Value)
			},
		},
		{
			name:    "no custom elements",
			fixture: "/custom-elements/no-custom-elements.html",
			validate: func(t *testing.T, elements []types.CustomElementMatch) {
				assert.Empty(t, elements)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			content := testutil.ReadFixture(t, mfs, tt.fixture)
			doc := handler.CreateDocument("file:///test.html", string(content), 1)
			defer doc.Close()
			elements, err := handler.FindCustomElements(doc)
			require.NoError(t, err)
			tt.validate(t, elements)
		})
	}
}
