/*
Copyright © 2025 Benny Powers

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
package search

import (
	"regexp"
	"strings"

	"bennypowers.dev/cem/list"
	M "bennypowers.dev/cem/manifest"
)

// CreateSearchPredicate creates a predicate function that searches for a pattern
// in names, descriptions, summaries, and labels of renderables.
// The pattern is treated as a regular expression, with fallback to literal string search
// if the regex is invalid.
func CreateSearchPredicate(pattern string) M.PredicateFunc {
	if pattern == "" {
		return M.True
	}

	// Try to compile as regex first
	var re *regexp.Regexp
	var err error

	// Make regex case-insensitive by default
	re, err = regexp.Compile("(?i)" + pattern)
	if err != nil {
		// If regex compilation fails, treat as literal string (case-insensitive)
		patternLower := strings.ToLower(pattern)
		return func(r M.Renderable) bool {
			return matchesLiteral(r, patternLower)
		}
	}

	return func(r M.Renderable) bool {
		return matchesRegex(r, re)
	}
}

// matchesRegex checks if a renderable matches the given regex pattern
func matchesRegex(r M.Renderable, re *regexp.Regexp) bool {
	// Search in name
	if re.MatchString(r.Name()) {
		return true
	}

	// Search in label
	if re.MatchString(r.Label()) {
		return true
	}

	// Search in summary and description if available
	if desc, ok := r.(M.Describable); ok {
		if re.MatchString(desc.Summary()) || re.MatchString(desc.Description()) {
			return true
		}
	}

	return false
}

// matchesLiteral checks if a renderable contains the given literal string (case-insensitive)
func matchesLiteral(r M.Renderable, pattern string) bool {
	// Search in name
	if strings.Contains(strings.ToLower(r.Name()), pattern) {
		return true
	}

	// Search in label
	if strings.Contains(strings.ToLower(r.Label()), pattern) {
		return true
	}

	// Search in summary and description if available
	if desc, ok := r.(M.Describable); ok {
		if strings.Contains(strings.ToLower(desc.Summary()), pattern) ||
			strings.Contains(strings.ToLower(desc.Description()), pattern) {
			return true
		}
	}

	return false
}

// RenderSearchResults renders search results using the existing list rendering functions
func RenderSearchResults(manifest *M.Package, pattern, format string) (string, error) {
	searchPred := CreateSearchPredicate(pattern)

	switch format {
	case "tree":
		title := "Search Results"
		if pattern != "" {
			title = "Search Results for: " + pattern
		}
		return list.RenderTree(title, M.NewRenderablePackage(manifest), searchPred)
	case "table":
		opts := list.RenderOptions{}
		return list.Render(M.NewRenderablePackage(manifest), opts, searchPred)
	default:
		title := "Search Results"
		if pattern != "" {
			title = "Search Results for: " + pattern
		}
		return list.RenderTree(title, M.NewRenderablePackage(manifest), searchPred)
	}
}
