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
package diagnostic_test

import (
	"testing"

	"bennypowers.dev/cem/lsp/document"
	"bennypowers.dev/cem/lsp/methods/textDocument/diagnostic"
	"bennypowers.dev/cem/lsp/testhelpers"
	M "bennypowers.dev/cem/manifest"
	protocol "github.com/bennypowers/glsp/protocol_3_17"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDocumentDiagnostic_NilDocument(t *testing.T) {
	ctx := testhelpers.NewMockServerContext()

	result, err := diagnostic.DocumentDiagnostic(ctx, nil, &protocol.DocumentDiagnosticParams{
		TextDocument: protocol.TextDocumentIdentifier{URI: "nonexistent.html"},
	})

	require.NoError(t, err)
	report, ok := result.(protocol.RelatedFullDocumentDiagnosticReport)
	require.True(t, ok)
	assert.Equal(t, string(protocol.DocumentDiagnosticReportKindFull), report.Kind)
	assert.Empty(t, report.Items)
}

func TestDocumentDiagnostic_ValidDocument(t *testing.T) {
	ctx := testhelpers.NewMockServerContext()
	dm, err := document.NewDocumentManager()
	require.NoError(t, err)
	defer dm.Close()
	ctx.SetDocumentManager(dm)

	html := `<my-button></my-button>`
	doc := dm.OpenDocument("test.html", html, 1)
	ctx.AddDocument("test.html", doc)

	result, err := diagnostic.DocumentDiagnostic(ctx, nil, &protocol.DocumentDiagnosticParams{
		TextDocument: protocol.TextDocumentIdentifier{URI: "test.html"},
	})

	require.NoError(t, err)
	report, ok := result.(protocol.RelatedFullDocumentDiagnosticReport)
	require.True(t, ok)
	assert.Equal(t, string(protocol.DocumentDiagnosticReportKindFull), report.Kind)
	assert.NotNil(t, report.Items)
}

func TestDocumentDiagnostic_WithDiagnostics(t *testing.T) {
	ctx := testhelpers.NewMockServerContext()
	dm, err := document.NewDocumentManager()
	require.NoError(t, err)
	defer dm.Close()
	ctx.SetDocumentManager(dm)

	html := `<my-button wrongattr="val"></my-button>`
	doc := dm.OpenDocument("test.html", html, 1)
	ctx.AddDocument("test.html", doc)

	ctx.AddAttributes("my-button", map[string]*M.Attribute{
		"disabled": {FullyQualified: M.FullyQualified{Name: "disabled"}, Type: &M.Type{Text: "Boolean"}},
	})

	result, err := diagnostic.DocumentDiagnostic(ctx, nil, &protocol.DocumentDiagnosticParams{
		TextDocument: protocol.TextDocumentIdentifier{URI: "test.html"},
	})

	require.NoError(t, err)
	report, ok := result.(protocol.RelatedFullDocumentDiagnosticReport)
	require.True(t, ok)
	assert.Greater(t, len(report.Items), 0)
}
