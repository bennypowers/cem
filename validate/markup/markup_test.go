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

// goldenIssue augments the public Issue with the optional structured details
// so the golden files pin both the public API and the data MCP renders.
type goldenIssue struct {
	markup.Issue
	Details map[string]string `json:"details,omitempty"`
}

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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			source := testutil.ReadFixture(t, fs, "/markup/"+tt.file)

			issues, err := markup.Validate(pkg, string(source), tt.opts)
			require.NoError(t, err)

			actual := make([]goldenIssue, len(issues))
			for i, issue := range issues {
				actual[i] = goldenIssue{
					Issue:   issue,
					Details: issue.Details(),
				}
			}

			var buf bytes.Buffer
			encoder := json.NewEncoder(&buf)
			encoder.SetEscapeHTML(false)
			encoder.SetIndent("", "  ")
			require.NoError(t, encoder.Encode(actual))
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
		name   string
		pkg    *manifest.Package
		format string
	}{
		{name: "nil package", pkg: nil, format: "html"},
		{name: "unsupported format", pkg: pkg, format: "markdown"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := markup.Validate(tt.pkg, "<my-button></my-button>", markup.Options{Format: tt.format})
			assert.Error(t, err)
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
