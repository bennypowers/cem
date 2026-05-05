/*
Copyright © 2026 Benny Powers

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
package workspace

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseNpmSpecifier(t *testing.T) {
	tests := []struct {
		spec        string
		wantName    string
		wantVersion string
		wantErr     bool
	}{
		{"pkg@1.2.3", "pkg", "1.2.3", false},
		{"@scope/pkg@1.2.3", "@scope/pkg", "1.2.3", false},
		{"pkg@latest", "pkg", "latest", false},
		{"pkg", "pkg", "latest", false},
		{"@scope/pkg", "@scope/pkg", "latest", false},
		{"npm:pkg@1.0.0", "pkg", "1.0.0", false},
		{"npm:@scope/pkg@2.0.0", "@scope/pkg", "2.0.0", false},
		{"@scope/pkg@", "@scope/pkg", "latest", false},
		{"", "", "", true},
	}
	for _, tc := range tests {
		t.Run(tc.spec, func(t *testing.T) {
			name, version, err := parseNpmSpecifier(tc.spec)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tc.wantName, name)
			assert.Equal(t, tc.wantVersion, version)
		})
	}
}

func TestIsGlobPattern(t *testing.T) {
	tests := []struct {
		pattern string
		want    bool
	}{
		{"packages/*", true},
		{"packages/*/src", true},
		{"packages/{a,b}", true},
		{"[abc]*", true},
		{"file?.txt", true},
		{"packages/normal", false},
		{"", false},
		{"simple-name", false},
		{"path/to/file.js", false},
	}
	for _, tc := range tests {
		t.Run(tc.pattern, func(t *testing.T) {
			assert.Equal(t, tc.want, isGlobPattern(tc.pattern))
		})
	}
}
