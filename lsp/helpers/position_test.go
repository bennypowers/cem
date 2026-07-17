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
package helpers

import (
	"testing"

	"go.lsp.dev/protocol"
)

// Inline: pure function, table-driven

func pos(line, char uint32) protocol.Position {
	return protocol.Position{Line: line, Character: char}
}

func rng(startLine, startChar, endLine, endChar uint32) protocol.Range {
	return protocol.Range{
		Start: pos(startLine, startChar),
		End:   pos(endLine, endChar),
	}
}

func TestIsPositionInRange(t *testing.T) {
	tests := []struct {
		name     string
		pos      protocol.Position
		r        protocol.Range
		expected bool
	}{
		{
			name:     "InsideRange",
			pos:      pos(5, 10),
			r:        rng(3, 0, 8, 20),
			expected: true,
		},
		{
			name:     "BeforeRange",
			pos:      pos(1, 5),
			r:        rng(3, 0, 8, 20),
			expected: false,
		},
		{
			name:     "AfterRange",
			pos:      pos(10, 0),
			r:        rng(3, 0, 8, 20),
			expected: false,
		},
		{
			name:     "AtExactStart",
			pos:      pos(3, 0),
			r:        rng(3, 0, 8, 20),
			expected: true,
		},
		{
			name:     "AtExactEnd",
			pos:      pos(8, 20),
			r:        rng(3, 0, 8, 20),
			expected: true,
		},
		{
			name:     "SingleLineRange_Inside",
			pos:      pos(5, 10),
			r:        rng(5, 5, 5, 15),
			expected: true,
		},
		{
			name:     "SingleLineRange_Before",
			pos:      pos(5, 3),
			r:        rng(5, 5, 5, 15),
			expected: false,
		},
		{
			name:     "SingleLineRange_After",
			pos:      pos(5, 16),
			r:        rng(5, 5, 5, 15),
			expected: false,
		},
		{
			name:     "MultiLineRange_StartLineBeforeChar",
			pos:      pos(3, 2),
			r:        rng(3, 5, 8, 20),
			expected: false,
		},
		{
			name:     "MultiLineRange_EndLineAfterChar",
			pos:      pos(8, 21),
			r:        rng(3, 5, 8, 20),
			expected: false,
		},
		{
			name:     "MultiLineRange_MiddleLine",
			pos:      pos(5, 0),
			r:        rng(3, 5, 8, 20),
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsPositionInRange(tt.pos, tt.r)
			if result != tt.expected {
				t.Errorf("IsPositionInRange(%v, %v) = %v, expected %v",
					tt.pos, tt.r, result, tt.expected)
			}
		})
	}
}

func TestPositionBefore(t *testing.T) {
	tests := []struct {
		name     string
		pos1     protocol.Position
		pos2     protocol.Position
		expected bool
	}{
		{
			name:     "DifferentLine_Before",
			pos1:     pos(3, 10),
			pos2:     pos(5, 0),
			expected: true,
		},
		{
			name:     "DifferentLine_After",
			pos1:     pos(5, 0),
			pos2:     pos(3, 10),
			expected: false,
		},
		{
			name:     "SameLine_Before",
			pos1:     pos(5, 3),
			pos2:     pos(5, 10),
			expected: true,
		},
		{
			name:     "SameLine_After",
			pos1:     pos(5, 10),
			pos2:     pos(5, 3),
			expected: false,
		},
		{
			name:     "EqualPositions",
			pos1:     pos(5, 10),
			pos2:     pos(5, 10),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PositionBefore(tt.pos1, tt.pos2)
			if result != tt.expected {
				t.Errorf("PositionBefore(%v, %v) = %v, expected %v",
					tt.pos1, tt.pos2, result, tt.expected)
			}
		})
	}
}

func TestPositionAfter(t *testing.T) {
	tests := []struct {
		name     string
		pos1     protocol.Position
		pos2     protocol.Position
		expected bool
	}{
		{
			name:     "DifferentLine_After",
			pos1:     pos(5, 0),
			pos2:     pos(3, 10),
			expected: true,
		},
		{
			name:     "DifferentLine_Before",
			pos1:     pos(3, 10),
			pos2:     pos(5, 0),
			expected: false,
		},
		{
			name:     "SameLine_After",
			pos1:     pos(5, 10),
			pos2:     pos(5, 3),
			expected: true,
		},
		{
			name:     "SameLine_Before",
			pos1:     pos(5, 3),
			pos2:     pos(5, 10),
			expected: false,
		},
		{
			name:     "EqualPositions",
			pos1:     pos(5, 10),
			pos2:     pos(5, 10),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PositionAfter(tt.pos1, tt.pos2)
			if result != tt.expected {
				t.Errorf("PositionAfter(%v, %v) = %v, expected %v",
					tt.pos1, tt.pos2, result, tt.expected)
			}
		})
	}
}

func TestRangesOverlap(t *testing.T) {
	tests := []struct {
		name     string
		start1   protocol.Position
		end1     protocol.Position
		start2   protocol.Position
		end2     protocol.Position
		expected bool
	}{
		{
			name:     "Overlapping",
			start1:   pos(1, 0),
			end1:     pos(5, 10),
			start2:   pos(3, 0),
			end2:     pos(8, 0),
			expected: true,
		},
		{
			name:     "AdjacentNonOverlapping",
			start1:   pos(1, 0),
			end1:     pos(3, 0),
			start2:   pos(5, 0),
			end2:     pos(8, 0),
			expected: false,
		},
		{
			name:     "Identical",
			start1:   pos(3, 0),
			end1:     pos(5, 10),
			start2:   pos(3, 0),
			end2:     pos(5, 10),
			expected: true,
		},
		{
			name:     "Nested",
			start1:   pos(1, 0),
			end1:     pos(10, 0),
			start2:   pos(3, 0),
			end2:     pos(5, 0),
			expected: true,
		},
		{
			name:     "AdjacentSameLineNotOverlapping",
			start1:   pos(1, 0),
			end1:     pos(3, 5),
			start2:   pos(3, 6),
			end2:     pos(5, 0),
			expected: false,
		},
		{
			name:     "TouchingAtBoundary",
			start1:   pos(1, 0),
			end1:     pos(3, 5),
			start2:   pos(3, 5),
			end2:     pos(5, 0),
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RangesOverlap(tt.start1, tt.end1, tt.start2, tt.end2)
			if result != tt.expected {
				t.Errorf("RangesOverlap(%v, %v, %v, %v) = %v, expected %v",
					tt.start1, tt.end1, tt.start2, tt.end2, result, tt.expected)
			}
		})
	}
}

func TestInRange(t *testing.T) {
	tests := []struct {
		name     string
		inner    protocol.Range
		outer    protocol.Range
		expected bool
	}{
		{
			name:     "InnerInsideOuter",
			inner:    rng(3, 0, 5, 10),
			outer:    rng(1, 0, 8, 20),
			expected: true,
		},
		{
			name:     "InnerBeforeOuter",
			inner:    rng(1, 0, 2, 10),
			outer:    rng(5, 0, 8, 20),
			expected: false,
		},
		{
			name:     "InnerAfterOuter",
			inner:    rng(10, 0, 12, 10),
			outer:    rng(5, 0, 8, 20),
			expected: false,
		},
		{
			name:     "PartialOverlap",
			inner:    rng(3, 0, 6, 10),
			outer:    rng(5, 0, 8, 20),
			expected: true,
		},
		{
			name:     "Identical",
			inner:    rng(5, 0, 8, 20),
			outer:    rng(5, 0, 8, 20),
			expected: true,
		},
		{
			name:     "SameEndLine_InnerCharBeforeOuterStart",
			inner:    rng(5, 0, 5, 3),
			outer:    rng(5, 5, 8, 20),
			expected: false,
		},
		{
			name:     "SameStartLine_InnerStartAtOuterEnd_TouchingInclusive",
			inner:    rng(8, 20, 10, 0),
			outer:    rng(5, 0, 8, 20),
			expected: true,
		},
		{
			name:     "SameStartLine_InnerStartBeforeOuterEnd",
			inner:    rng(8, 19, 10, 0),
			outer:    rng(5, 0, 8, 20),
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := InRange(tt.inner, tt.outer)
			if result != tt.expected {
				t.Errorf("InRange(%v, %v) = %v, expected %v",
					tt.inner, tt.outer, result, tt.expected)
			}
		})
	}
}
