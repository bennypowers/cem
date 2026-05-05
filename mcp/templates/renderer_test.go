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
package templates

import (
	"embed"
	"strings"
	"sync"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// testdata embeds the testdata directory. The FS paths are prefixed with
// "testdata/", so we cannot use RegisterTemplateSource directly (it expects
// the FS to contain "templates/<name>.md" at the root). Instead we use it
// to test createTemplate's error paths and exercise RegisterTemplateSource
// itself, then use newTestPool for full-pipeline Render tests.
//
//go:embed testdata
var testdataFS embed.FS

// newTestPool returns a TemplatePool pre-loaded with a hand-parsed template
// named "simple" that exercises the complete Render execution path.
func newTestPool() *TemplatePool {
	pool := NewTemplatePool()
	simpleTmpl := template.Must(
		template.New("simple").Funcs(pool.funcMap).Parse("Hello, {{.Name}}!\n"),
	)
	funcsTmpl := template.Must(
		template.New("with-funcs").Funcs(pool.funcMap).Parse("# {{title .Title}}\n\n{{len .Items}} items: {{join .Items \", \"}}\n"),
	)
	errTmpl := template.Must(
		template.New("exec-error").Funcs(pool.funcMap).Parse("{{.Name.Invalid.Method}}\n"),
	)

	pool.poolsMu.Lock()
	pool.pools["simple"] = &sync.Pool{New: func() any { return simpleTmpl }}
	pool.pools["with-funcs"] = &sync.Pool{New: func() any { return funcsTmpl }}
	pool.pools["exec-error"] = &sync.Pool{New: func() any { return errTmpl }}
	pool.poolsMu.Unlock()

	return pool
}

// --- createSecureFuncMap: exercise each template function ---

func TestFuncMap_Title(t *testing.T) {
	fm := createSecureFuncMap()
	fn := fm["title"].(func(string) string)

	tests := []struct {
		name string
		in   string
		want string
	}{
		{"simple", "hello world", "Hello World"},
		{"already titled", "Hello", "Hello"},
		{"empty", "", ""},
		{"single word", "foo", "Foo"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, fn(tc.in))
		})
	}
}

func TestFuncMap_Len(t *testing.T) {
	fm := createSecureFuncMap()
	fn := fm["len"].(func(any) int)

	tests := []struct {
		name string
		in   any
		want int
	}{
		{"string slice", []string{"a", "b"}, 2},
		{"empty string slice", []string{}, 0},
		{"map string any", map[string]any{"k": 1}, 1},
		{"map string string", map[string]string{"a": "b", "c": "d"}, 2},
		{"map string string slice", map[string][]string{"x": {"1", "2"}}, 1},
		{"nil", nil, 0},
		{"int slice via reflection", []int{1, 2, 3}, 3},
		{"int array via reflection", [2]int{1, 2}, 2},
		{"unsupported type returns 0", 42, 0},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, fn(tc.in))
		})
	}
}

// lenner implements the Len() int interface for testing
type lenner struct{ n int }

func (l lenner) Len() int { return l.n }

func TestFuncMap_Len_Interface(t *testing.T) {
	fm := createSecureFuncMap()
	fn := fm["len"].(func(any) int)
	assert.Equal(t, 7, fn(lenner{7}))
}

func TestFuncMap_Index(t *testing.T) {
	fm := createSecureFuncMap()
	fn := fm["index"].(func(any, int) any)

	slice := []string{"a", "b", "c"}

	tests := []struct {
		name  string
		slice any
		idx   int
		want  any
	}{
		{"first", slice, 0, "a"},
		{"last", slice, 2, "c"},
		{"negative oob", slice, -1, ""},
		{"positive oob", slice, 5, ""},
		{"non-string slice returns empty", []int{1, 2}, 0, ""},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, fn(tc.slice, tc.idx))
		})
	}
}

func TestFuncMap_Gt(t *testing.T) {
	fm := createSecureFuncMap()
	fn := fm["gt"].(func(int, int) bool)

	assert.True(t, fn(2, 1))
	assert.False(t, fn(1, 2))
	assert.False(t, fn(1, 1))
}

func TestFuncMap_Eq(t *testing.T) {
	fm := createSecureFuncMap()
	fn := fm["eq"].(func(any, any) bool)

	assert.True(t, fn(1, 1))
	assert.True(t, fn("x", "x"))
	assert.False(t, fn(1, 2))
	assert.False(t, fn("a", "b"))
	assert.True(t, fn(nil, nil))
}

func TestFuncMap_Add(t *testing.T) {
	fm := createSecureFuncMap()
	fn := fm["add"].(func(int, int) int)

	assert.Equal(t, 3, fn(1, 2))
	assert.Equal(t, 0, fn(-1, 1))
	assert.Equal(t, 0, fn(0, 0))
}

func TestFuncMap_Join(t *testing.T) {
	fm := createSecureFuncMap()
	fn := fm["join"].(func([]string, string) string)

	assert.Equal(t, "a, b, c", fn([]string{"a", "b", "c"}, ", "))
	assert.Equal(t, "x", fn([]string{"x"}, ","))
	assert.Equal(t, "", fn([]string{}, ","))
}

func TestFuncMap_Sanitize(t *testing.T) {
	fm := createSecureFuncMap()
	fn := fm["sanitize"].(func(string) string)

	// sanitize delegates to security.SanitizeDescription; verify plumbing
	assert.Equal(t, "", fn(""))
	assert.Equal(t, "hello", fn("hello"))
	// collapse whitespace
	assert.Equal(t, "a b", fn("a   b"))
}

func TestFuncMap_SchemaDesc(t *testing.T) {
	fm := createSecureFuncMap()
	fn := fm["schemaDesc"].(func(any, string) string)

	schema := map[string]any{
		"definitions": map[string]any{
			"Foo": map[string]any{
				"description": "Foo desc",
			},
		},
	}
	assert.Equal(t, "Foo desc", fn(schema, "Foo"))
	assert.Equal(t, "", fn(schema, "Missing"))
	assert.Equal(t, "", fn("not a map", "Foo"))
}

func TestFuncMap_SchemaFieldDesc(t *testing.T) {
	fm := createSecureFuncMap()
	fn := fm["schemaFieldDesc"].(func(any, string, string) string)

	schema := map[string]any{
		"definitions": map[string]any{
			"Bar": map[string]any{
				"properties": map[string]any{
					"name": map[string]any{
						"description": "The name",
					},
				},
			},
		},
	}
	assert.Equal(t, "The name", fn(schema, "Bar", "name"))
	assert.Equal(t, "", fn(schema, "Bar", "missing"))
	assert.Equal(t, "", fn(schema, "NoType", "name"))
	assert.Equal(t, "", fn(42, "Bar", "name"))
}

// --- getSchemaDescription ---

func TestGetSchemaDescription(t *testing.T) {
	tests := []struct {
		name     string
		schema   any
		typeName string
		want     string
	}{
		{
			name:     "from definitions",
			schema:   map[string]any{"definitions": map[string]any{"A": map[string]any{"description": "desc-A"}}},
			typeName: "A",
			want:     "desc-A",
		},
		{
			name:     "from properties when definitions missing",
			schema:   map[string]any{"properties": map[string]any{"B": map[string]any{"description": "desc-B"}}},
			typeName: "B",
			want:     "desc-B",
		},
		{
			name:     "definitions take precedence",
			schema:   map[string]any{"definitions": map[string]any{"C": map[string]any{"description": "def"}}, "properties": map[string]any{"C": map[string]any{"description": "prop"}}},
			typeName: "C",
			want:     "def",
		},
		{
			name:     "not a map",
			schema:   "nope",
			typeName: "X",
			want:     "",
		},
		{
			name:     "type not found",
			schema:   map[string]any{"definitions": map[string]any{}},
			typeName: "Z",
			want:     "",
		},
		{
			name:     "no description field",
			schema:   map[string]any{"definitions": map[string]any{"D": map[string]any{"type": "object"}}},
			typeName: "D",
			want:     "",
		},
		{
			name:     "nil schema",
			schema:   nil,
			typeName: "X",
			want:     "",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, getSchemaDescription(tc.schema, tc.typeName))
		})
	}
}

// --- getSchemaFieldDescription ---

func TestGetSchemaFieldDescription(t *testing.T) {
	tests := []struct {
		name      string
		schema    any
		typeName  string
		fieldName string
		want      string
	}{
		{
			name: "field found in definitions",
			schema: map[string]any{
				"definitions": map[string]any{
					"Elem": map[string]any{
						"properties": map[string]any{
							"slot": map[string]any{"description": "slot desc"},
						},
					},
				},
			},
			typeName:  "Elem",
			fieldName: "slot",
			want:      "slot desc",
		},
		{
			name: "field found in properties section",
			schema: map[string]any{
				"properties": map[string]any{
					"Elem": map[string]any{
						"properties": map[string]any{
							"attr": map[string]any{"description": "attr desc"},
						},
					},
				},
			},
			typeName:  "Elem",
			fieldName: "attr",
			want:      "attr desc",
		},
		{
			name:      "not a map",
			schema:    123,
			typeName:  "X",
			fieldName: "y",
			want:      "",
		},
		{
			name: "type exists but field missing",
			schema: map[string]any{
				"definitions": map[string]any{
					"T": map[string]any{
						"properties": map[string]any{},
					},
				},
			},
			typeName:  "T",
			fieldName: "nope",
			want:      "",
		},
		{
			name: "type exists but no properties key",
			schema: map[string]any{
				"definitions": map[string]any{
					"T": map[string]any{"type": "string"},
				},
			},
			typeName:  "T",
			fieldName: "x",
			want:      "",
		},
		{
			name:      "type not found anywhere",
			schema:    map[string]any{},
			typeName:  "Missing",
			fieldName: "f",
			want:      "",
		},
		{
			name:      "nil schema",
			schema:    nil,
			typeName:  "X",
			fieldName: "f",
			want:      "",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, getSchemaFieldDescription(tc.schema, tc.typeName, tc.fieldName))
		})
	}
}

// --- Render pipeline tests using pre-parsed templates ---

func TestRender_SimpleTemplate(t *testing.T) {
	pool := newTestPool()
	result, err := pool.Render("simple", map[string]string{"Name": "world"})
	require.NoError(t, err)
	assert.Equal(t, "Hello, world!\n", result)
}

func TestRender_TemplateFuncs(t *testing.T) {
	pool := newTestPool()
	result, err := pool.Render("with-funcs", map[string]any{
		"Title": "my title",
		"Items": []string{"a", "b", "c"},
	})
	require.NoError(t, err)
	assert.Contains(t, result, "My Title")
	assert.Contains(t, result, "3 items")
	assert.Contains(t, result, "a, b, c")
}

func TestRender_ExecutionError(t *testing.T) {
	pool := NewTemplatePool()

	// A template that calls a method on data, but the data type lacks that method.
	// template.Must panics on parse error, so we parse a valid template that
	// will fail at execution time when given incompatible data.
	badTmpl := template.Must(
		template.New("exec-error").Funcs(pool.funcMap).Parse("{{.Boom}}"),
	)
	pool.poolsMu.Lock()
	pool.pools["exec-error"] = &sync.Pool{New: func() any { return badTmpl }}
	pool.poolsMu.Unlock()

	// Pass a type that has no Boom field and uses missingkey=error
	badTmpl.Option("missingkey=error")
	_, err := pool.Render("exec-error", struct{}{})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to execute template")
}

func TestRender_TemplateNotFoundFromFS(t *testing.T) {
	pool := NewTemplatePool()
	pool.RegisterTemplateSource("test", &testdataFS)

	// testdataFS has paths like "testdata/templates/simple.md",
	// but createTemplate looks for "templates/<name>.md", so it won't find it.
	// This exercises the err != nil path in createTemplate.
	_, err := pool.Render("nonexistent", nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "template not found")
}

func TestRender_NoSourcesRegistered(t *testing.T) {
	pool := NewTemplatePool()

	// With no template sources, createTemplate's loop body never runs.
	// err stays nil, content stays empty, so an empty template is parsed.
	// Render succeeds with empty output -- this is expected behavior.
	result, err := pool.Render("anything", nil)
	require.NoError(t, err)
	assert.Equal(t, "", result)
}

// --- Render path traversal ---

func TestRender_PathTraversal(t *testing.T) {
	pool := NewTemplatePool()

	tests := []struct {
		name         string
		templateName string
	}{
		{"dot-dot", "../secret"},
		{"forward slash", "sub/dir"},
		{"backslash", `sub\dir`},
		{"dot-dot with slash", "../../etc/passwd"},
		{"embedded dots", "a..b"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := pool.Render(tc.templateName, nil)
			require.Error(t, err)
			assert.Contains(t, err.Error(), "invalid template name")
		})
	}
}

// --- RegisterTemplateSource ---

func TestRegisterTemplateSource(t *testing.T) {
	t.Run("stores FS without panic", func(t *testing.T) {
		pool := NewTemplatePool()
		pool.RegisterTemplateSource("pkg-a", &testdataFS)

		pool.templateSourcesMu.RLock()
		_, exists := pool.templateSources["pkg-a"]
		pool.templateSourcesMu.RUnlock()
		assert.True(t, exists)
	})

	t.Run("overwrite replaces previous", func(t *testing.T) {
		pool := NewTemplatePool()
		var other embed.FS
		pool.RegisterTemplateSource("pkg", &testdataFS)
		pool.RegisterTemplateSource("pkg", &other)

		pool.templateSourcesMu.RLock()
		got := pool.templateSources["pkg"]
		pool.templateSourcesMu.RUnlock()
		assert.Equal(t, &other, got)
	})
}

// --- createTemplate ---

func TestCreateTemplate_NotFound(t *testing.T) {
	pool := NewTemplatePool()
	// Register an FS that has no "templates/" directory at root level
	pool.RegisterTemplateSource("test", &testdataFS)

	tmpl := pool.createTemplate("nonexistent")
	assert.Nil(t, tmpl, "createTemplate should return nil for missing templates")
}

// --- getTemplatePool ---

func TestGetTemplatePool_CachesPool(t *testing.T) {
	pool := NewTemplatePool()

	p1 := pool.getTemplatePool("test-cache")
	p2 := pool.getTemplatePool("test-cache")

	// Same pointer means the pool was cached
	assert.Equal(t, p1, p2)
}

func TestGetTemplatePool_DifferentNames(t *testing.T) {
	pool := NewTemplatePool()

	p1 := pool.getTemplatePool("alpha")
	p2 := pool.getTemplatePool("beta")

	assert.NotEqual(t, p1, p2)
}

// --- Global wrappers ---

func TestGlobalRegisterTemplateSource(t *testing.T) {
	// Should not panic
	RegisterTemplateSource("renderer-test-pkg", &testdataFS)
}

func TestGlobalRenderTemplate_InvalidName(t *testing.T) {
	_, err := RenderTemplate("../bad", nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid template name")
}

func TestGlobalRenderTemplate_MissingFromFS(t *testing.T) {
	// The global pool may have template sources registered by other tests.
	// With testdataFS registered (wrong path layout), a truly missing name
	// should return "template not found" if sources are registered,
	// or empty string if no sources hit the ReadFile error path.
	// Either way the call should not panic.
	result, err := RenderTemplate("completely-missing-template-xyz", nil)
	if err != nil {
		assert.Contains(t, err.Error(), "template not found")
	} else {
		assert.Equal(t, "", result)
	}
}

// --- Render: nil template from pool ---

func TestRender_NilTemplateInPool(t *testing.T) {
	pool := NewTemplatePool()

	// Manually insert a pool that returns nil (simulating createTemplate failure)
	pool.poolsMu.Lock()
	pool.pools["bad"] = &sync.Pool{New: func() any { return nil }}
	pool.poolsMu.Unlock()

	_, err := pool.Render("bad", nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "template not found")
}

// --- Render: wrong type from pool ---

func TestRender_WrongTypeInPool(t *testing.T) {
	pool := NewTemplatePool()

	// Pool returns a non-template value
	pool.poolsMu.Lock()
	pool.pools["wrong"] = &sync.Pool{New: func() any { return "not a template" }}
	pool.poolsMu.Unlock()

	_, err := pool.Render("wrong", nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "template not found")
}

// --- Concurrent safety ---

func TestRender_ConcurrentSafety(t *testing.T) {
	pool := newTestPool()

	const goroutines = 20
	errs := make(chan error, goroutines)

	for range goroutines {
		go func() {
			_, err := pool.Render("simple", map[string]string{"Name": "concurrent"})
			errs <- err
		}()
	}

	for range goroutines {
		assert.NoError(t, <-errs)
	}
}

// --- FuncMap completeness ---

func TestFuncMap_AllExpectedFunctionsExist(t *testing.T) {
	fm := createSecureFuncMap()

	expected := []string{
		"title", "schemaDesc", "schemaFieldDesc",
		"len", "index", "gt", "eq", "join", "add", "sanitize",
	}
	for _, name := range expected {
		t.Run(name, func(t *testing.T) {
			_, exists := fm[name]
			assert.True(t, exists, "function map should contain %q", name)
		})
	}
}

// --- Render validates template name before pool lookup ---

func TestRender_ValidNameButMissing(t *testing.T) {
	pool := NewTemplatePool()
	pool.RegisterTemplateSource("test", &testdataFS)

	// testdataFS paths don't match the expected "templates/<name>.md" layout,
	// so ReadFile fails for all sources. This exercises the "not found" path.
	_, err := pool.Render("valid-but-missing", nil)
	require.Error(t, err)
	// Should get "template not found", not "invalid template name"
	assert.Contains(t, err.Error(), "template not found")
	assert.NotContains(t, err.Error(), "invalid template name")
}

// --- Edge: render with empty data ---

func TestRender_EmptyData(t *testing.T) {
	pool := NewTemplatePool()

	tmpl := template.Must(
		template.New("static").Funcs(pool.funcMap).Parse("static content"),
	)
	pool.poolsMu.Lock()
	pool.pools["static"] = &sync.Pool{New: func() any { return tmpl }}
	pool.poolsMu.Unlock()

	result, err := pool.Render("static", nil)
	require.NoError(t, err)
	assert.Equal(t, "static content", result)
}

// --- Render returns clean output without trailing artifacts ---

func TestRender_OutputIntegrity(t *testing.T) {
	pool := newTestPool()

	result, err := pool.Render("simple", map[string]string{"Name": "test"})
	require.NoError(t, err)
	assert.Equal(t, "Hello, test!\n", result)
	// No double newlines or other artifacts
	assert.False(t, strings.HasSuffix(result, "\n\n"))
}
