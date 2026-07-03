/*
Copyright © 2026 Benny Powers <web@bennypowers.com>

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
package breaking

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsSemverTag(t *testing.T) {
	tests := []struct {
		tag  string
		want bool
	}{
		{"1.0.0", true},
		{"v1.0.0", true},
		{"v0.11.0", true},
		{"v2.1.1", true},
		{"1.2.3-alpha.1", true},
		{"v1.0.0-rc.1", true},
		{"v1.0.0-beta", true},
		{"1.2", false},
		{"v1", false},
		{"", false},
		{"latest", false},
		{"v1.x.3", false},
		{"v1.2.x", false},
		{"release-1.0", false},
		{"v.1.2.3", false},
	}
	for _, tt := range tests {
		t.Run(tt.tag, func(t *testing.T) {
			assert.Equal(t, tt.want, isSemverTag(tt.tag))
		})
	}
}

func TestResolveGitPath(t *testing.T) {
	// inline assertions: testing path resolution logic with real git repo
	root, err := gitTopLevel(".")
	require.NoError(t, err)

	t.Run("relative path", func(t *testing.T) {
		got, err := resolveGitPath(".", "custom-elements.json")
		require.NoError(t, err)
		assert.NotEmpty(t, got)
		assert.NotContains(t, got, root)
	})

	t.Run("absolute path within repo", func(t *testing.T) {
		absPath := filepath.Join(root, "custom-elements.json")
		got, err := resolveGitPath(".", absPath)
		require.NoError(t, err)
		assert.Equal(t, "custom-elements.json", got)
	})
}
