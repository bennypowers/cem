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
package queries

import (
	"iter"
	"slices"

	ts "github.com/tree-sitter/go-tree-sitter"
)

type QueryMatcher struct {
	query  *ts.Query
	cursor *ts.QueryCursor
}

func (qm QueryMatcher) Close() {
	// NOTE: we don't close queries here, only at the end of execution in QueryManager.Close
	// Close the cursor since we're not pooling
	qm.cursor.Close()
}

func (qm QueryMatcher) GetCaptureNameByIndex(index uint32) string {
	return qm.query.CaptureNames()[index]
}

func (qm QueryMatcher) CaptureCount() int {
	return len(qm.query.CaptureNames())
}

func (qm QueryMatcher) GetCaptureIndexForName(name string) (uint, bool) {
	return qm.query.CaptureIndexForName(name)
}

func (qm QueryMatcher) SetByteRange(start uint, end uint) {
	qm.cursor.SetByteRange(start, end)
}

func NewQueryMatcher(
	manager *QueryManager,
	language string,
	queryName string,
) (*QueryMatcher, error) {
	if manager == nil {
		return nil, ErrNoQueryManager
	}
	query, err := manager.getQuery(queryName, language)
	if err != nil {
		return nil, err
	}
	cursor := ts.NewQueryCursor()
	qm := QueryMatcher{query, cursor}
	return &qm, nil
}

func (q QueryMatcher) AllQueryMatches(node *ts.Node, text []byte) iter.Seq[*ts.QueryMatch] {
	qm := q.cursor.Matches(q.query, node, text)
	return func(yield func(qm *ts.QueryMatch) bool) {
		for {
			m := qm.Next()
			if m == nil {
				break
			}
			if !yield(m) {
				return
			}
		}
	}
}

// ParentCaptures returns an iterator over unique parent node captures as identified by the given parent capture name.
// For each unique parent node (e.g., a field or method), it aggregates all captures from all query matches sharing
// that parent node into a single CaptureMap. This allows you to collect all related captures (such as attributes,
// decorators, etc.) for each parent node in the source AST.
//
// Example usage:
//
//	for captures := range matcher.ParentCaptures(root, code, "field") {
//	  // captures represents the captured nodes for a single field
//	  addFieldToDeclaration(captures, declaration)
//	}
func (q *QueryMatcher) ParentCaptures(root *ts.Node, code []byte, parentCaptureName string) iter.Seq[CaptureMap] {
	names := q.query.CaptureNames()

	type pgroup struct {
		capMap    CaptureMap
		startByte uint
	}

	parentGroups := make(map[int]pgroup)

	for match := range q.AllQueryMatches(root, code) {
		var parentNode *ts.Node
		for _, cap := range match.Captures {
			name := names[cap.Index]
			if name == parentCaptureName {
				parentNode = &cap.Node
				break
			}
		}
		if parentNode == nil {
			continue
		}
		pid := int(parentNode.Id())
		startByte := parentNode.StartByte()
		_, ok := parentGroups[pid]
		if !ok {
			capmap := make(CaptureMap)
			parentGroups[pid] = pgroup{capmap, startByte}
		}
		for _, cap := range match.Captures {
			name := names[cap.Index]
			text := cap.Node.Utf8Text(code)
			ci := CaptureInfo{
				NodeId:    int(cap.Node.Id()),
				Text:      text,
				StartByte: cap.Node.StartByte(),
				EndByte:   cap.Node.EndByte(),
			}
			if _, hasMap := parentGroups[pid].capMap[name]; !hasMap {
				parentGroups[pid].capMap[name] = make([]CaptureInfo, 0)
			}
			if !slices.ContainsFunc(parentGroups[pid].capMap[name], func(m CaptureInfo) bool {
				return m.NodeId == ci.NodeId
			}) {
				parentGroups[pid].capMap[name] = append(parentGroups[pid].capMap[name], ci)
			}
		}
	}

	sorted := make([]pgroup, 0)
	for _, group := range parentGroups {
		sorted = append(sorted, group)
	}

	slices.SortStableFunc(sorted, func(a pgroup, b pgroup) int {
		return int(a.startByte) - int(b.startByte)
	})

	return func(yield func(CaptureMap) bool) {
		for _, group := range sorted {
			if !yield(group.capMap) {
				break
			}
		}
	}
}

// GetCachedQueryMatcher returns a query matcher with cached query but fresh cursor
// This prevents concurrent access to the same cursor which causes segmentation faults
func GetCachedQueryMatcher(manager *QueryManager, language, queryName string) (*QueryMatcher, error) {
	if manager == nil {
		return nil, ErrNoQueryManager
	}

	// Get the cached query (thread-safe to share)
	query, err := manager.getQuery(queryName, language)
	if err != nil {
		return nil, err
	}

	// Always create a fresh cursor (NOT thread-safe to share)
	cursor := ts.NewQueryCursor()
	matcher := QueryMatcher{query, cursor}
	return &matcher, nil
}
