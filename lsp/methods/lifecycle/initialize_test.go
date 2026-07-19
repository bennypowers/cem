package lifecycle

import (
	"testing"

	"bennypowers.dev/cem/lsp/testhelpers"
	"go.lsp.dev/protocol"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Inline: integration test, scalar assertions
// TestInitializeDiagnosticIdentifier validates LSP capability negotiation.
// TestParseStringSlice is a pure function, table-driven.

func TestInitializeDiagnosticIdentifier(t *testing.T) {
	t.Run("pull diagnostics identifier is set", func(t *testing.T) {
		ctx := testhelpers.NewMockServerContext()
		params := &protocol.InitializeParams{
			Capabilities: protocol.ClientCapabilities{
				TextDocument: &protocol.TextDocumentClientCapabilities{
					Diagnostic: &protocol.DiagnosticClientCapabilities{},
				},
			},
		}

		result, err := Initialize(ctx, params)
		require.NoError(t, err)

		require.NotNil(t, result.Capabilities.DiagnosticProvider)
		diagOpts, ok := result.Capabilities.DiagnosticProvider.(*protocol.DiagnosticOptions)
		require.True(t, ok)
		require.NotNil(t, diagOpts.Identifier)
		assert.Equal(t, "cem", *diagOpts.Identifier)
	})

	t.Run("no diagnostic provider without client support", func(t *testing.T) {
		ctx := testhelpers.NewMockServerContext()
		params := &protocol.InitializeParams{
			Capabilities: protocol.ClientCapabilities{},
		}

		result, err := Initialize(ctx, params)
		require.NoError(t, err)

		assert.Nil(t, result.Capabilities.DiagnosticProvider)
	})
}

func TestParseStringSlice(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected []string
	}{
		{
			name:     "string slice",
			input:    []string{"foo", "bar", "baz"},
			expected: []string{"foo", "bar", "baz"},
		},
		{
			name:     "[]any with strings",
			input:    []any{"alpha", "beta"},
			expected: []string{"alpha", "beta"},
		},
		{
			name:     "[]any with mixed types - only strings kept",
			input:    []any{"hello", 42, true, "world"},
			expected: []string{"hello", "world"},
		},
		{
			name:     "non-slice returns nil",
			input:    "not a slice",
			expected: nil,
		},
		{
			name:     "nil returns nil",
			input:    nil,
			expected: nil,
		},
		{
			name:     "integer returns nil",
			input:    42,
			expected: nil,
		},
		{
			name:     "empty []any",
			input:    []any{},
			expected: []string{},
		},
		{
			name:     "empty []string",
			input:    []string{},
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseStringSlice(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
