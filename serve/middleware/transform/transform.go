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

package transform

import (
	"time"
)

// CacheKey uniquely identifies a cached transform based on file metadata
type CacheKey struct {
	Path    string
	ModTime time.Time
	Size    int64
}

// CacheEntry stores transformed code and its dependencies
type CacheEntry struct {
	Code         []byte
	Dependencies []string
	Size         int64
	AccessTime   time.Time
}

// TransformResult contains the transformed code and dependencies
type TransformResult struct {
	Code         []byte
	Map          []byte
	Dependencies []string
}
