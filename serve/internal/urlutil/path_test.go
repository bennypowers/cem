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

package urlutil_test

import (
	"testing"

	"bennypowers.dev/cem/serve/internal/urlutil"
)

func TestContainsPath(t *testing.T) {
	tests := []struct {
		name    string
		fullURL string
		path    string
		want    bool
	}{
		{
			name:    "exact match",
			fullURL: "/demo/basic",
			path:    "/demo/basic",
			want:    true,
		},
		{
			name:    "prefix match with slash",
			fullURL: "/demo/basic/index.html",
			path:    "/demo/basic",
			want:    true,
		},
		{
			name:    "prefix match with query",
			fullURL: "/demo/basic?foo=bar",
			path:    "/demo/basic",
			want:    true,
		},
		{
			name:    "prefix match with hash",
			fullURL: "/demo/basic#section",
			path:    "/demo/basic",
			want:    true,
		},
		{
			name:    "no match - different path",
			fullURL: "/demo/advanced",
			path:    "/demo/basic",
			want:    false,
		},
		{
			name:    "no match - false prefix",
			fullURL: "/demo/basically",
			path:    "/demo/basic",
			want:    false,
		},
		{
			name:    "empty path",
			fullURL: "/demo/basic",
			path:    "",
			want:    false,
		},
		{
			name:    "URL encoded space in fullURL",
			fullURL: "/demo/my%20file.html",
			path:    "/demo/my file.html",
			want:    true,
		},
		{
			name:    "URL encoded space in path",
			fullURL: "/demo/my file.html",
			path:    "/demo/my%20file.html",
			want:    true,
		},
		{
			name:    "URL encoded spaces in both",
			fullURL: "/demo/my%20file.html",
			path:    "/demo/my%20file.html",
			want:    true,
		},
		{
			name:    "URL encoded special chars",
			fullURL: "/demo/%C3%A9l%C3%A9ment.html",
			path:    "/demo/élément.html",
			want:    true,
		},
		{
			name:    "prefix with encoded path",
			fullURL: "/demo/my%20file.html/index",
			path:    "/demo/my file.html",
			want:    true,
		},
		{
			name:    "root path",
			fullURL: "/",
			path:    "/",
			want:    true,
		},
		{
			name:    "trailing slash in fullURL",
			fullURL: "/demo/basic/",
			path:    "/demo/basic",
			want:    true,
		},
		{
			name:    "trailing slash in path",
			fullURL: "/demo/basic",
			path:    "/demo/basic/",
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := urlutil.ContainsPath(tt.fullURL, tt.path)
			if got != tt.want {
				t.Errorf("ContainsPath(%q, %q) = %v, want %v",
					tt.fullURL, tt.path, got, tt.want)
			}
		})
	}
}
