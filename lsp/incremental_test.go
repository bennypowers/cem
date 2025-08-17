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
package lsp

import (
	"testing"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TestIncrementalParser_AnalyzeChanges(t *testing.T) {
	parser := NewIncrementalParser(ParseStrategyAuto)

	tests := []struct {
		name                string
		changes             []protocol.TextDocumentContentChangeEvent
		oldContent          string
		expectedSmall       bool
		expectedIncremental bool
	}{
		{
			name: "Single small change should use incremental",
			changes: []protocol.TextDocumentContentChangeEvent{
				{
					Range: &protocol.Range{
						Start: protocol.Position{Line: 0, Character: 5},
						End:   protocol.Position{Line: 0, Character: 10},
					},
					Text: "hello",
				},
			},
			oldContent:          "some world text",
			expectedSmall:       true,
			expectedIncremental: true,
		},
		{
			name: "Large change should not use incremental",
			changes: []protocol.TextDocumentContentChangeEvent{
				{
					Range: &protocol.Range{
						Start: protocol.Position{Line: 0, Character: 0},
						End:   protocol.Position{Line: 50, Character: 0},
					},
					Text: string(make([]byte, 2000)), // Large change
				},
			},
			oldContent:          "small content",
			expectedSmall:       false,
			expectedIncremental: false,
		},
		{
			name: "Many changes should not use incremental",
			changes: func() []protocol.TextDocumentContentChangeEvent {
				changes := make([]protocol.TextDocumentContentChangeEvent, 10) // Too many changes
				for i := range changes {
					changes[i] = protocol.TextDocumentContentChangeEvent{
						Range: &protocol.Range{
							Start: protocol.Position{Line: 0, Character: uint32(i)},
							End:   protocol.Position{Line: 0, Character: uint32(i + 1)},
						},
						Text: "x",
					}
				}
				return changes
			}(),
			oldContent:          "0123456789abcdef",
			expectedSmall:       true,
			expectedIncremental: false,
		},
		{
			name: "Full document change should not use incremental",
			changes: []protocol.TextDocumentContentChangeEvent{
				{
					Range: nil, // Full document change
					Text:  "completely new content",
				},
			},
			oldContent:          "old content",
			expectedSmall:       true,
			expectedIncremental: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			analysis := parser.AnalyzeChanges(tt.changes, tt.oldContent)

			if analysis.IsSmallChange != tt.expectedSmall {
				t.Errorf("Expected IsSmallChange=%t, got %t", tt.expectedSmall, analysis.IsSmallChange)
			}

			if analysis.ShouldUseIncremental != tt.expectedIncremental {
				t.Errorf("Expected ShouldUseIncremental=%t, got %t", tt.expectedIncremental, analysis.ShouldUseIncremental)
			}

			if len(analysis.Edits) != len(tt.changes) {
				t.Errorf("Expected %d edits, got %d", len(tt.changes), len(analysis.Edits))
			}
		})
	}
}

func TestIncrementalParser_ParseStrategies(t *testing.T) {
	tests := []struct {
		name     string
		strategy ParseStrategy
	}{
		{"Full parsing strategy", ParseStrategyFull},
		{"Incremental parsing strategy", ParseStrategyIncremental},
		{"Auto parsing strategy", ParseStrategyAuto},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := NewIncrementalParser(tt.strategy)
			if parser == nil {
				t.Fatal("Failed to create incremental parser")
			}
			if parser.strategy != tt.strategy {
				t.Errorf("Expected strategy %v, got %v", tt.strategy, parser.strategy)
			}
		})
	}
}

func TestIncrementalParser_PositionToByteOffset(t *testing.T) {
	parser := NewIncrementalParser(ParseStrategyAuto)

	content := "line 1\nline 2\nline 3"

	tests := []struct {
		name     string
		position protocol.Position
		expected uint
	}{
		{
			name:     "Start of document",
			position: protocol.Position{Line: 0, Character: 0},
			expected: 0,
		},
		{
			name:     "Middle of first line",
			position: protocol.Position{Line: 0, Character: 3},
			expected: 3,
		},
		{
			name:     "Start of second line",
			position: protocol.Position{Line: 1, Character: 0},
			expected: 7, // "line 1\n" = 7 bytes
		},
		{
			name:     "Middle of second line",
			position: protocol.Position{Line: 1, Character: 2},
			expected: 9, // "line 1\nli" = 9 bytes
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parser.positionToByteOffset(tt.position, content)
			if result != tt.expected {
				t.Errorf("Expected byte offset %d, got %d", tt.expected, result)
			}
		})
	}
}

func TestSplitLines(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected []string
	}{
		{
			name:     "Empty content",
			content:  "",
			expected: []string{},
		},
		{
			name:     "Single line",
			content:  "hello world",
			expected: []string{"hello world"},
		},
		{
			name:     "Multiple lines",
			content:  "line 1\nline 2\nline 3",
			expected: []string{"line 1", "line 2", "line 3"},
		},
		{
			name:     "Trailing newline",
			content:  "line 1\nline 2\n",
			expected: []string{"line 1", "line 2"},
		},
		{
			name:     "Only newlines",
			content:  "\n\n\n",
			expected: []string{"", "", ""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := splitLines(tt.content)
			if len(result) != len(tt.expected) {
				t.Errorf("Expected %d lines, got %d", len(tt.expected), len(result))
				return
			}

			for i, line := range result {
				if line != tt.expected[i] {
					t.Errorf("Line %d: expected %q, got %q", i, tt.expected[i], line)
				}
			}
		})
	}
}
