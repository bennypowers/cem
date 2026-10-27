/*
Copyright © 2026 Benny Powers <web@bennypowers.com>

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
package markup_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"bennypowers.dev/cem/internal/platform/testutil"
	"bennypowers.dev/cem/manifest"
	"bennypowers.dev/cem/validate/markup"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func loadPackage(t *testing.T) *manifest.Package {
	t.Helper()

	fs := testutil.LoadTestdataFS(t, "testdata", "/")
	data := testutil.ReadFixture(t, fs, "/custom-elements.json")

	var pkg manifest.Package
	require.NoError(t, json.Unmarshal(data, &pkg))
	return &pkg
}

func TestValidateGolden(t *testing.T) {
	fs := testutil.LoadTestdataFS(t, "testdata", "/")
	pkg := loadPackage(t)

	tests := []struct {
		name   string
		golden string
		file   string
		format string
		opts   markup.Options
	}{
		{
			name:   "clean html",
			golden: "clean_html.json",
			file:   "clean.html",
			format: "html",
			opts:   markup.Options{ValidateAttrs: true, ValidateSlots: true},
		},
		{
			name:   "unknown attribute and invalid enum value",
			golden: "attributes.json",
			file:   "attributes.html",
			format: "html",
			opts:   markup.Options{ValidateAttrs: true},
		},
		{
			name:   "unknown and deprecated slots",
			golden: "slots.json",
			file:   "slots.html",
			format: "html",
			opts:   markup.Options{ValidateSlots: true},
		},
		{
			name:   "deprecated element and attribute",
			golden: "deprecated.json",
			file:   "deprecated.html",
			format: "html",
			opts:   markup.Options{ValidateAttrs: true, ValidateSlots: true},
		},
		{
			name:   "prose guidelines conflicts and redundancy",
			golden: "prose.json",
			file:   "prose.html",
			format: "html",
			opts:   markup.Options{ValidateAttrs: true, ValidateSlots: true, Strict: true},
		},
		{
			name:   "prose findings omitted without strict",
			golden: "prose_non_strict.json",
			file:   "prose.html",
			format: "html",
			opts:   markup.Options{ValidateAttrs: true, ValidateSlots: true},
		},
		{
			name:   "slot content guidelines",
			golden: "guidelines.json",
			file:   "guidelines.html",
			format: "html",
			opts:   markup.Options{ValidateSlots: true, Strict: true},
		},
		{
			name:   "attribute conflicts",
			golden: "conflicts.json",
			file:   "conflicts.html",
			format: "html",
			opts:   markup.Options{ValidateAttrs: true, Strict: true},
		},
		{
			name:   "content attribute redundancy",
			golden: "redundancy.json",
			file:   "redundancy.html",
			format: "html",
			opts:   markup.Options{ValidateSlots: true, Strict: true},
		},
		{
			name:   "twig with interpolation and balanced control flow",
			golden: "twig_basic.json",
			file:   "twig-basic.twig",
			format: "twig",
			opts:   markup.Options{ValidateAttrs: true, ValidateSlots: true},
		},
		{
			name:   "twig with unbalanced control flow",
			golden: "twig_unbalanced.json",
			file:   "twig-unbalanced.twig",
			format: "twig",
			opts:   markup.Options{ValidateAttrs: true},
		},
		{
			name:   "twig markup-like text outside html ranges",
			golden: "twig_non_markup.json",
			file:   "twig-non-markup.twig",
			format: "twig",
			opts:   markup.Options{ValidateAttrs: true, ValidateSlots: true, Strict: true},
		},
		{
			name:   "twig mixed static and expression values are dynamic",
			golden: "twig_mixed_values.json",
			file:   "twig-mixed-values.twig",
			format: "twig",
			opts:   markup.Options{ValidateAttrs: true, ValidateSlots: true},
		},
		{
			name:   "twig unclosed custom elements are discovered",
			golden: "twig_unclosed.json",
			file:   "twig-unclosed.twig",
			format: "twig",
			opts:   markup.Options{ValidateAttrs: true},
		},
		{
			name:   "self closing custom elements are validated",
			golden: "self_closing.json",
			file:   "self-closing.html",
			format: "html",
			opts:   markup.Options{ValidateAttrs: true},
		},
		{
			name:   "unclosed custom elements are discovered",
			golden: "unclosed.json",
			file:   "unclosed.html",
			format: "html",
			opts:   markup.Options{ValidateAttrs: true},
		},
		{
			name:   "html template expression value keeps literal validation",
			golden: "html_template_value.json",
			file:   "html-template-value.html",
			format: "html",
			opts:   markup.Options{ValidateAttrs: true},
		},
		{
			name:   "hbs template expression value keeps template semantics",
			golden: "hbs_template_value.json",
			file:   "hbs-template-value.hbs",
			format: "hbs",
			opts:   markup.Options{ValidateAttrs: true},
		},
		{
			name:   "hbs block helpers and helper arguments are not markup",
			golden: "hbs_helpers.json",
			file:   "hbs-helpers.hbs",
			format: "hbs",
			opts:   markup.Options{ValidateAttrs: true},
		},
		{
			name:   "duplicate attributes keep the first value",
			golden: "duplicates.json",
			file:   "duplicates.html",
			format: "html",
			opts:   markup.Options{ValidateAttrs: true},
		},
		{
			name:   "uppercase names normalize to lowercase",
			golden: "uppercase.json",
			file:   "uppercase.html",
			format: "html",
			opts:   markup.Options{ValidateAttrs: true},
		},
		{
			name:   "html character references decode in static values",
			golden: "entities.json",
			file:   "entities.html",
			format: "html",
			opts:   markup.Options{ValidateAttrs: true},
		},
		{
			name:   "raw text and rcdata elements are not parsed as markup",
			golden: "rcdata.json",
			file:   "rcdata.html",
			format: "html",
			opts:   markup.Options{ValidateAttrs: true, ValidateSlots: true, Strict: true},
		},
		{
			name:   "non ascii columns are byte based",
			golden: "unicode.json",
			file:   "unicode.html",
			format: "html",
			opts:   markup.Options{ValidateAttrs: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			source := testutil.ReadFixture(t, fs, "/markup/"+tt.file)

			opts := tt.opts
			opts.Format = tt.format
			issues, err := markup.Validate(pkg, string(source), opts)
			require.NoError(t, err)
			if issues == nil {
				// Goldens represent the empty result as [] rather than null.
				issues = []markup.Issue{}
			}

			var buf bytes.Buffer
			encoder := json.NewEncoder(&buf)
			encoder.SetEscapeHTML(false)
			encoder.SetIndent("", "  ")
			require.NoError(t, encoder.Encode(issues))
			data := buf.Bytes()

			testutil.CheckGolden(t, tt.golden, data, testutil.GoldenOptions{
				Dir:       "/goldens",
				FS:        fs,
				SourceDir: "testdata",
			})
		})
	}
}

// These error paths are pure input validation, so inline assertions are clearer
// than a fixture directory per error.
func TestValidateErrors(t *testing.T) {
	pkg := loadPackage(t)

	tests := []struct {
		name    string
		pkg     *manifest.Package
		format  string
		wantErr error
	}{
		{name: "nil package", pkg: nil, format: "html", wantErr: markup.ErrNilPackage},
		{name: "unsupported format", pkg: pkg, format: "markdown", wantErr: markup.ErrUnsupportedFormat},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := markup.Validate(tt.pkg, "<my-button></my-button>", markup.Options{Format: tt.format})
			require.ErrorIs(t, err, tt.wantErr)
		})
	}
}

// Empty input is a pure edge case; inline assertions keep the intent obvious.
func TestValidateEmptyMarkup(t *testing.T) {
	pkg := loadPackage(t)

	issues, err := markup.Validate(pkg, "", markup.Options{ValidateAttrs: true, ValidateSlots: true})
	require.NoError(t, err)
	assert.Empty(t, issues)
}

// Defaults are part of the API contract and are easier to read inline than in goldens.
func TestValidateDefaultFormat(t *testing.T) {
	pkg := loadPackage(t)

	issues, err := markup.Validate(pkg, `<unknown-element bogus="1"></unknown-element>`, markup.Options{})
	require.NoError(t, err)
	require.Len(t, issues, 1)
	assert.Equal(t, markup.RuleUnknownElement, issues[0].Rule)
	assert.Equal(t, markup.ConfidenceHigh, issues[0].Confidence)
}

// Dynamic template values are unknown when validating, so conflict heuristics
// must skip them. This filters one rule because a golden for the same input
// would also pin unrelated deprecation and unknown-attribute findings.
func TestValidateAttributeCombinations(t *testing.T) {
	pkg := loadPackage(t)

	conflicts := func(t *testing.T, source string) []markup.Issue {
		t.Helper()

		issues, err := markup.Validate(pkg, source, markup.Options{
			Format:        "twig",
			ValidateAttrs: true,
			Strict:        true,
		})
		require.NoError(t, err)

		var found []markup.Issue
		for _, issue := range issues {
			if issue.Rule == markup.RuleAttributeConflict {
				found = append(found, issue)
			}
		}
		return found
	}

	tests := []struct {
		name   string
		source string
		want   int
	}{
		{
			name:   "dynamic value is unknown",
			source: `<my-card loading="eager" lazy="{{ dynamic }}">Card</my-card>`,
			want:   0,
		},
		{
			name:   "eager loading with a true lazy attribute conflicts",
			source: `<my-card loading="eager" lazy="true">Card</my-card>`,
			want:   1,
		},
		{
			name:   "eager loading with a bare lazy attribute conflicts",
			source: `<my-card loading="eager" lazy>Card</my-card>`,
			want:   1,
		},
		{
			name:   "lazy loading does not conflict",
			source: `<my-card loading="lazy" lazy>Card</my-card>`,
			want:   0,
		},
		{
			name:   "disabled and interactive conflict",
			source: `<my-card disabled interactive>Card</my-card>`,
			want:   1,
		},
		{
			name:   "disabled alone does not conflict",
			source: `<my-card disabled>Card</my-card>`,
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Len(t, conflicts(t, tt.source), tt.want)
		})
	}
}

// Options gate rule families independently. A golden per flag combination
// would duplicate the same markup, so these assertions name each flag's effect.
func TestValidateOptionGating(t *testing.T) {
	pkg := loadPackage(t)

	source := `<my-button bogus="1"><span slot="nope">x</span></my-button>`

	rules := func(t *testing.T, opts markup.Options) []string {
		t.Helper()

		issues, err := markup.Validate(pkg, source, opts)
		require.NoError(t, err)

		names := make([]string, len(issues))
		for i, issue := range issues {
			names[i] = issue.Rule
		}
		return names
	}

	assert.Equal(t, []string{markup.RuleUnknownAttribute},
		rules(t, markup.Options{ValidateAttrs: true}))
	assert.Equal(t, []string{markup.RuleUnknownSlot},
		rules(t, markup.Options{ValidateSlots: true}))
	assert.Equal(t, []string{markup.RuleUnknownAttribute, markup.RuleUnknownSlot},
		rules(t, markup.Options{ValidateAttrs: true, ValidateSlots: true}))
	assert.Empty(t, rules(t, markup.Options{Strict: true}),
		"Strict alone adds no findings")
}

// Liquid comments never render, so markup inside them must not be validated;
// markup outside them must still be. Inline assertions state both halves.
func TestValidateLiquidComments(t *testing.T) {
	pkg := loadPackage(t)

	commented, err := markup.Validate(pkg,
		`{% comment %}<my-button variant="bogus">{% endcomment %}`,
		markup.Options{Format: "liquid", ValidateAttrs: true})
	require.NoError(t, err)
	assert.Empty(t, commented)

	live, err := markup.Validate(pkg,
		`{% comment %}ignored{% endcomment %}<my-button variant="bogus">`,
		markup.Options{Format: "liquid", ValidateAttrs: true})
	require.NoError(t, err)
	require.Len(t, live, 1)
	assert.Equal(t, markup.RuleInvalidAttributeValue, live[0].Rule)

	// A tag-looking string inside an output expression is not a comment.
	tricky, err := markup.Validate(pkg,
		`{{ "{% comment %}" }}<my-button variant="bogus">{% comment %}real{% endcomment %}`,
		markup.Options{Format: "liquid", ValidateAttrs: true})
	require.NoError(t, err)
	require.Len(t, tricky, 1)
	assert.Equal(t, markup.RuleInvalidAttributeValue, tricky[0].Rule)

	// A literal comment opener inside a raw block is not a comment.
	raw, err := markup.Validate(pkg,
		`{% raw %}{% comment %}{% endraw %}<my-button variant="bogus">{% comment %} real {% endcomment %}`,
		markup.Options{Format: "liquid", ValidateAttrs: true})
	require.NoError(t, err)
	require.Len(t, raw, 1)
	assert.Equal(t, markup.RuleInvalidAttributeValue, raw[0].Rule)
}

// These parser regressions each exercise one masking quirk, so a focused table
// states the expected rule presence without unrelated output.
func TestValidateTemplateEdgeCases(t *testing.T) {
	pkg := loadPackage(t)

	hasRule := func(t *testing.T, source, format, rule string) bool {
		t.Helper()

		issues, err := markup.Validate(pkg, source, markup.Options{
			Format:        format,
			ValidateAttrs: true,
			ValidateSlots: true,
		})
		require.NoError(t, err)
		for _, issue := range issues {
			if issue.Rule == rule {
				return true
			}
		}
		return false
	}

	t.Run("unquoted twig value is dynamic", func(t *testing.T) {
		assert.False(t, hasRule(t, `<my-button variant=prim{{ suffix }}>`, "twig",
			markup.RuleInvalidAttributeValue))
	})

	t.Run("unquoted twig value keeps the element", func(t *testing.T) {
		assert.True(t, hasRule(t, `{% if x %}<my-unknown variant={{ v }}>{% endif %}`, "twig",
			markup.RuleUnknownElement))
	})

	t.Run("unquoted hbs value is dynamic", func(t *testing.T) {
		assert.False(t, hasRule(t, `<my-button variant={{ v }}>`, "hbs",
			markup.RuleInvalidAttributeValue))
	})

	t.Run("hbs comment apostrophe does not corrupt masking", func(t *testing.T) {
		assert.False(t, hasRule(t, `{{!-- don't --}}<my-button {{#if x}}variant="primary"{{/if}}>`, "hbs",
			markup.RuleUnknownAttribute))
	})

	t.Run("escaped mustache is literal text", func(t *testing.T) {
		elements, err := markup.FindCustomElements(
			`<p>Use \{{ braces</p><my-card {{#if x}}variant="primary"{{/if}}></my-card>`,
			markup.Options{Format: "hbs"})
		require.NoError(t, err)

		tags := make([]string, len(elements))
		for i, element := range elements {
			tags[i] = element.TagName
		}
		assert.Contains(t, tags, "my-card")
	})

	t.Run("camel case manifest attribute matches", func(t *testing.T) {
		assert.False(t, hasRule(t, `<my-card flipBoundary="x">`, "html",
			markup.RuleUnknownAttribute))
	})

	t.Run("dynamic tag name is skipped", func(t *testing.T) {
		assert.False(t, hasRule(t, `<my-{{ x }} variant="bogus">`, "twig",
			markup.RuleUnknownElement))
	})

	t.Run("dynamic attribute name is skipped", func(t *testing.T) {
		assert.False(t, hasRule(t, `<my-button {{ n }}ariant="bogus">`, "twig",
			markup.RuleUnknownAttribute))
	})

	t.Run("noscript children are markup", func(t *testing.T) {
		assert.True(t, hasRule(t, `<noscript><my-button variant="bogus"></my-button></noscript>`, "html",
			markup.RuleInvalidAttributeValue))
	})

	t.Run("conditional twig attribute is still validated", func(t *testing.T) {
		assert.True(t, hasRule(t, `<my-button {% if x %}variant="bogus"{% endif %}>`, "twig",
			markup.RuleInvalidAttributeValue))
	})

	t.Run("conditional hbs attribute is still validated", func(t *testing.T) {
		assert.True(t, hasRule(t, `<my-button {{#if x}}variant="bogus"{{/if}}>`, "hbs",
			markup.RuleInvalidAttributeValue))
	})

	t.Run("twig statement after a static attribute does not suppress it", func(t *testing.T) {
		assert.True(t, hasRule(t, `<my-button variant="bogus"{% if x %}class="a"{% endif %}>`, "twig",
			markup.RuleInvalidAttributeValue))
	})

	t.Run("twig comment before an attribute does not suppress it", func(t *testing.T) {
		assert.True(t, hasRule(t, `<my-button {# c #}variant="bogus">`, "twig",
			markup.RuleInvalidAttributeValue))
	})

	t.Run("unterminated literal mustache does not stop masking", func(t *testing.T) {
		assert.True(t, hasRule(t, `<p>Use "{{" here</p><my-button {{#if x}}variant="bogus"{{/if}}>`, "hbs",
			markup.RuleInvalidAttributeValue))
	})

	t.Run("empty hbs comment is masked", func(t *testing.T) {
		assert.True(t, hasRule(t, `<my-button {{!----}}variant="bogus">`, "hbs",
			markup.RuleInvalidAttributeValue))
	})

	t.Run("dynamic hbs attribute name does not drop later attributes", func(t *testing.T) {
		assert.True(t, hasRule(t, `<my-button {{x}}="v" bogus="1">`, "hbs",
			markup.RuleUnknownAttribute))
	})

	t.Run("dynamic twig attribute name does not drop later attributes", func(t *testing.T) {
		assert.True(t, hasRule(t, `<my-button {{x}}=v bogus="1">`, "twig",
			markup.RuleUnknownAttribute))
	})

	t.Run("dynamic tag name is not discovered", func(t *testing.T) {
		elements, err := markup.FindCustomElements(`<my-{{ x }}>`, markup.Options{Format: "twig"})
		require.NoError(t, err)
		assert.Empty(t, elements)
	})

	t.Run("dynamic attribute name is not redundant", func(t *testing.T) {
		assert.False(t, hasRule(t, `<my-card {{ x }}label="Save">`, "twig",
			markup.RuleContentAttributeRedundancy))
	})

	t.Run("html comment quote does not corrupt masking", func(t *testing.T) {
		source := `<!-- don't --><my-button variant=prim{{ x }}>`

		assert.False(t, hasRule(t, source, "twig", markup.RuleInvalidAttributeValue))

		elements, err := markup.FindCustomElements(source, markup.Options{Format: "twig"})
		require.NoError(t, err)
		assert.Len(t, elements, 1)
	})

	t.Run("mask state advances past a value span", func(t *testing.T) {
		assert.True(t, hasRule(t, `<my-card flipBoundary={{ x }} {# c #}loading="bogus">`, "twig",
			markup.RuleInvalidAttributeValue))
	})

	t.Run("dynamic name does not trigger a conflict", func(t *testing.T) {
		assert.False(t, hasRule(t, `<my-card {{ x }}lazy="true" loading="eager">`, "twig",
			markup.RuleAttributeConflict))
	})

	t.Run("dynamic slot name is not an unknown slot", func(t *testing.T) {
		assert.False(t, hasRule(t, `<my-card><span {{ x }}slot="nope">x</span></my-card>`, "twig",
			markup.RuleUnknownSlot))
	})

	t.Run("whitespace control block does not suppress an attribute", func(t *testing.T) {
		assert.True(t, hasRule(t, `<my-card {{~#if x~}}loading="bogus"{{~/if~}}>`, "hbs",
			markup.RuleInvalidAttributeValue))
	})

	t.Run("comment before equals does not suppress an attribute", func(t *testing.T) {
		assert.True(t, hasRule(t, `<my-button variant{# c #}="bogus">`, "twig",
			markup.RuleInvalidAttributeValue))
	})

	t.Run("statement before equals does not suppress an attribute", func(t *testing.T) {
		assert.True(t, hasRule(t, `<my-button variant{% if x %}{% endif %}="bogus">`, "twig",
			markup.RuleInvalidAttributeValue))
	})

	t.Run("dollar braces are literal in twig", func(t *testing.T) {
		assert.True(t, hasRule(t, `<my-button variant="${bogus}">`, "twig",
			markup.RuleInvalidAttributeValue))
	})

	t.Run("jinja statement is literal in hbs values", func(t *testing.T) {
		assert.True(t, hasRule(t, `<my-button variant="{% bogus %}">`, "hbs",
			markup.RuleInvalidAttributeValue))
	})

	t.Run("expression containing closing braces is measured correctly", func(t *testing.T) {
		assert.False(t, hasRule(t, `<my-card {{ "}}" }}variant="primary">`, "twig",
			markup.RuleUnknownAttribute))
	})

	t.Run("else-prefixed helper is an expression", func(t *testing.T) {
		assert.False(t, hasRule(t, `<my-button {{elsewhere}}variant="bogus">`, "hbs",
			markup.RuleInvalidAttributeValue))
	})

	t.Run("raw-only content is validated", func(t *testing.T) {
		assert.True(t, hasRule(t, `{% raw %}<my-bad-el bogus>{% endraw %}`, "twig",
			markup.RuleUnknownElement))
	})

	t.Run("plaintext consumes following markup", func(t *testing.T) {
		assert.False(t, hasRule(t, `<plaintext><my-card variant="bogus">`, "html",
			markup.RuleInvalidAttributeValue))
	})

	t.Run("self-closing textarea consumes following markup", func(t *testing.T) {
		assert.False(t, hasRule(t, `<textarea/><my-card variant="bogus">`, "html",
			markup.RuleInvalidAttributeValue))
	})

	t.Run("unclosed textarea consumes following markup", func(t *testing.T) {
		assert.False(t, hasRule(t, `<textarea><my-card variant="bogus">`, "html",
			markup.RuleInvalidAttributeValue))
	})
}

// Issue is a JSON-serializable result type; a consumer that marshals the
// slice must not silently lose the rule-specific details.
func TestIssueJSONIncludesDetails(t *testing.T) {
	pkg := loadPackage(t)

	issues, err := markup.Validate(pkg, `<my-button bogus="1">Save</my-button>`, markup.Options{ValidateAttrs: true})
	require.NoError(t, err)
	require.Len(t, issues, 1)

	data, err := json.Marshal(issues)
	require.NoError(t, err)

	var decoded []map[string]any
	require.NoError(t, json.Unmarshal(data, &decoded))
	require.Len(t, decoded, 1)

	details, ok := decoded[0]["details"].(map[string]any)
	require.True(t, ok, "serialized issue must include details")
	assert.Equal(t, "bogus", details[markup.DetailAttribute])
}

// FindCustomElements is a pure discovery API; a table-like inline assertion is
// clearer than a golden because the expected value is just tag order.
func TestFindCustomElements(t *testing.T) {
	source := `<MY-CHILD><my-child/><span><my-nested></my-nested></span></MY-CHILD><my-el>`

	elements, err := markup.FindCustomElements(source, markup.Options{Format: "html"})
	require.NoError(t, err)

	tagNames := make([]string, len(elements))
	for i, element := range elements {
		tagNames[i] = element.TagName
	}
	assert.Equal(t, []string{"my-child", "my-child", "my-nested", "my-el"}, tagNames)
	assert.Equal(t, 2, elements[0].Column, "outer element column")
	assert.Equal(t, 12, elements[1].Column, "self-closing element column")
}

// FuzzValidate checks that arbitrary markup, including malformed template
// syntax, never panics the parser or validator. Seeds cover both formats.
func FuzzValidate(f *testing.F) {
	f.Add(`<my-el variant="a"><span slot="label">x</span></my-el>`)
	f.Add(`{% if x %}<my-el bogus>{% endif %}`)
	f.Add(`{# <my-el bogus> #}`)
	f.Add(`{% raw %}<my-el bogus>{% endraw %}`)
	f.Add(``)

	pkg := fuzzPackage()

	f.Fuzz(func(t *testing.T, source string) {
		for _, format := range []string{"html", "twig", "hbs"} {
			issues, err := markup.Validate(pkg, source, markup.Options{
				Format:        format,
				ValidateAttrs: true,
				ValidateSlots: true,
				Strict:        true,
			})
			if err != nil {
				t.Fatalf("Validate(%q, %q) returned error: %v", format, source, err)
			}
			for _, issue := range issues {
				if issue.Confidence != markup.ConfidenceHigh && issue.Confidence != markup.ConfidenceLow {
					t.Fatalf("issue has unknown confidence %q: %+v", issue.Confidence, issue)
				}
			}
		}
	})
}

func fuzzPackage() *manifest.Package {
	const data = `{
		"schemaVersion": "2.1.0",
		"modules": [{
			"kind": "javascript-module",
			"path": "fuzz.js",
			"declarations": [{
				"kind": "class",
				"name": "MyEl",
				"tagName": "my-el",
				"customElement": true,
				"attributes": [{"name": "variant", "type": {"text": "\"a\" | \"b\""}}],
				"slots": [{"name": "label", "description": "Only interactive content must avoid plain text."}]
			}]
		}]
	}`

	var pkg manifest.Package
	if err := json.Unmarshal([]byte(data), &pkg); err != nil {
		panic(err)
	}
	return &pkg
}
