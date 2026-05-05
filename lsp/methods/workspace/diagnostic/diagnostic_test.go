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
	workspaceDiag "bennypowers.dev/cem/lsp/methods/workspace/diagnostic"
	"bennypowers.dev/cem/lsp/testhelpers"
	protocol "github.com/bennypowers/glsp/protocol_3_17"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWorkspaceDiagnostic_Empty(t *testing.T) {
	ctx := testhelpers.NewMockServerContext()

	result, err := workspaceDiag.WorkspaceDiagnostic(ctx, nil, &protocol.WorkspaceDiagnosticParams{})

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Empty(t, result.Items)
}

func TestWorkspaceDiagnostic_MultipleDocuments(t *testing.T) {
	ctx := testhelpers.NewMockServerContext()
	dm, err := document.NewDocumentManager()
	require.NoError(t, err)
	defer dm.Close()
	ctx.SetDocumentManager(dm)

	doc1 := dm.OpenDocument("a.html", `<my-el></my-el>`, 1)
	doc2 := dm.OpenDocument("b.html", `<other-el></other-el>`, 1)
	ctx.AddDocument("a.html", doc1)
	ctx.AddDocument("b.html", doc2)

	result, err := workspaceDiag.WorkspaceDiagnostic(ctx, nil, &protocol.WorkspaceDiagnosticParams{})

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Len(t, result.Items, 2)
}
