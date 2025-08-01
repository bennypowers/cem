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
package list_test

import (
	"encoding/json"
	"flag"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"bennypowers.dev/cem/list"
	"bennypowers.dev/cem/manifest"
	M "bennypowers.dev/cem/manifest"
	"github.com/stretchr/testify/assert"
)

var update = flag.Bool("update", false, "update golden files")
var ansiRegexp = regexp.MustCompile(`\x1b\[[0-9;]*m`)

func stripANSI(s string) string {
	return ansiRegexp.ReplaceAllString(s, "")
}

func loadTestFixture(t *testing.T) *manifest.Package {
	t.Helper()
	bytes, err := os.ReadFile(filepath.Join("fixtures", filepath.Base(t.Name())+".json"))
	assert.NoError(t, err)
	var pkg manifest.Package
	err = json.Unmarshal(bytes, &pkg)
	assert.NoError(t, err)
	return &pkg
}

func checkGolden(t *testing.T, actual []byte) {
	t.Helper()
	goldenPath := filepath.Join("goldens", filepath.Base(t.Name())+".md")
	if *update {
		err := os.WriteFile(goldenPath, actual, 0644)
		assert.NoError(t, err, "failed to update golden file")
		return
	}

	expected, err := os.ReadFile(goldenPath)
	assert.NoError(t, err, "failed to read golden file")

	// Normalize line endings for comparison
	expectedNormalized := strings.ReplaceAll(string(expected), "\r\n", "\n")
	actualNormalized := strings.ReplaceAll(string(actual), "\r\n", "\n")

	assert.Equal(t, expectedNormalized, actualNormalized)
}

func TestRender(t *testing.T) {
	t.Run("custom-element-table-coverage", func(t *testing.T) {
		pkg := loadTestFixture(t)
		renderable := manifest.NewRenderablePackage(pkg)

		opts := list.RenderOptions{}
		output, err := list.Render(renderable, opts, M.True)
		assert.NoError(t, err)

		checkGolden(t, []byte(stripANSI(output)))
	})
}
