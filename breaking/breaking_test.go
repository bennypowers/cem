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
package breaking_test

import (
	"encoding/json"
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/breaking"
	"bennypowers.dev/cem/internal/platform"
	"bennypowers.dev/cem/internal/platform/testutil"
	M "bennypowers.dev/cem/manifest"
	"github.com/stretchr/testify/require"
)

var fixtureNames = []string{
	"element-removed",
	"attribute-changes",
	"mixed-changes",
	"no-changes",
}

func loadManifest(t *testing.T, mfs *platform.MapFileSystem, path string) *M.Package {
	t.Helper()
	data, err := mfs.ReadFile(path)
	require.NoError(t, err)
	var pkg M.Package
	require.NoError(t, json.Unmarshal(data, &pkg))
	return &pkg
}

func TestCompare(t *testing.T) {
	mfs := testutil.LoadTestdataFS(t, "testdata/fixtures", "/fixtures")

	for _, name := range fixtureNames {
		t.Run(name, func(t *testing.T) {
			base := loadManifest(t, mfs, filepath.Join("/fixtures", name, "base.json"))
			head := loadManifest(t, mfs, filepath.Join("/fixtures", name, "head.json"))
			result := breaking.Compare(base, head, breaking.Options{})
			actual, err := json.MarshalIndent(result, "", "  ")
			require.NoError(t, err)
			actual = append(actual, '\n')
			testutil.CheckGolden(t, name, actual, testutil.GoldenOptions{
				Dir:         "testdata/goldens",
				Extension:   ".json",
				UseJSONDiff: true,
			})
		})
	}
}

func TestCompareWithDisable(t *testing.T) {
	mfs := testutil.LoadTestdataFS(t, "testdata/fixtures", "/fixtures")
	base := loadManifest(t, mfs, "/fixtures/mixed-changes/base.json")
	head := loadManifest(t, mfs, "/fixtures/mixed-changes/head.json")
	result := breaking.Compare(base, head, breaking.Options{
		Disable: []string{"element-added", "css-custom-property-added", "css-part-added", "event-added", "method-added"},
	})

	for _, c := range result.Changes {
		switch c.Rule {
		case "element-added", "css-custom-property-added", "css-part-added", "event-added", "method-added":
			t.Errorf("disabled rule %q should not appear in results", c.Rule)
		}
	}
}
