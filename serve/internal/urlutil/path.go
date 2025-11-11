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

// Package urlutil provides utility functions for URL path matching and manipulation.
package urlutil

import (
	"net/url"
)

// ContainsPath checks if a full URL contains a path as an exact match or valid prefix.
// It returns true if:
// - fullURL exactly equals path, or
// - fullURL starts with path followed by a valid delimiter ('/', '?', or '#')
//
// Both fullURL and path are URL-decoded before comparison to handle
// cases like "/demo/my%20file.html" matching "/demo/my file.html".
//
// This prevents false positive matches like "/elements/button" matching "/elements/buttonly".
func ContainsPath(fullURL, path string) bool {
	if len(path) == 0 {
		return false
	}

	// Decode both URLs to normalize encoding differences
	decodedFull, err := url.PathUnescape(fullURL)
	if err != nil {
		// If decoding fails, use original string
		decodedFull = fullURL
	}
	decodedPath, err := url.PathUnescape(path)
	if err != nil {
		// If decoding fails, use original string
		decodedPath = path
	}

	// Exact match
	if decodedFull == decodedPath {
		return true
	}
	// Prefix match only if followed by valid delimiter ('/', '?', or '#')
	if len(decodedFull) > len(decodedPath) && decodedFull[:len(decodedPath)] == decodedPath {
		// Check the character immediately after the path
		nextChar := decodedFull[len(decodedPath)]
		if nextChar == '/' || nextChar == '?' || nextChar == '#' {
			return true
		}
	}
	return false
}
