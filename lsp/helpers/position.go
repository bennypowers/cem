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

import protocol "github.com/bennypowers/glsp/protocol_3_17"

// IsPositionInRange checks if a position falls within a range (inclusive).
func IsPositionInRange(pos protocol.Position, r protocol.Range) bool {
	if pos.Line < r.Start.Line || pos.Line > r.End.Line {
		return false
	}

	if pos.Line == r.Start.Line && pos.Character < r.Start.Character {
		return false
	}

	if pos.Line == r.End.Line && pos.Character > r.End.Character {
		return false
	}

	return true
}

// PositionBefore checks if pos1 is strictly before pos2.
func PositionBefore(pos1, pos2 protocol.Position) bool {
	if pos1.Line < pos2.Line {
		return true
	}
	if pos1.Line == pos2.Line && pos1.Character < pos2.Character {
		return true
	}
	return false
}

// PositionAfter checks if pos1 is strictly after pos2.
func PositionAfter(pos1, pos2 protocol.Position) bool {
	if pos1.Line > pos2.Line {
		return true
	}
	if pos1.Line == pos2.Line && pos1.Character > pos2.Character {
		return true
	}
	return false
}

// RangesOverlap checks if two ranges overlap.
// Range 1 starts before range 2 ends, AND range 1 ends after range 2 starts.
func RangesOverlap(start1, end1, start2, end2 protocol.Position) bool {
	return !PositionAfter(start1, end2) && !PositionBefore(end1, start2)
}

// InRange checks if the inner range overlaps with the outer range.
func InRange(inner, outer protocol.Range) bool {
	if inner.End.Line < outer.Start.Line {
		return false
	}
	if inner.Start.Line > outer.End.Line {
		return false
	}
	if inner.End.Line == outer.Start.Line && inner.End.Character < outer.Start.Character {
		return false
	}
	if inner.Start.Line == outer.End.Line && inner.Start.Character >= outer.End.Character {
		return false
	}
	return true
}
