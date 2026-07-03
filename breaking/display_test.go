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
	"bytes"
	"encoding/json"
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/breaking"
	"bennypowers.dev/cem/internal/platform/testutil"
	"github.com/stretchr/testify/require"
)

func TestPrintResultText(t *testing.T) {
	mfs := testutil.LoadTestdataFS(t, "testdata/fixtures", "/fixtures")

	for _, name := range fixtureNames {
		t.Run(name, func(t *testing.T) {
			base := loadManifest(t, mfs, filepath.Join("/fixtures", name, "base.json"))
			head := loadManifest(t, mfs, filepath.Join("/fixtures", name, "head.json"))
			result := breaking.Compare(base, head, breaking.Options{})

			var buf bytes.Buffer
			err := breaking.PrintResult(&buf, result, breaking.DisplayOptions{Format: "text"})
			require.NoError(t, err)

			testutil.CheckGolden(t, "text-"+name, buf.Bytes(), testutil.GoldenOptions{
				Dir:       "testdata/goldens",
				Extension: ".txt",
				StripANSI: true,
			})
		})
	}
}

func TestPrintResultMarkdown(t *testing.T) {
	mfs := testutil.LoadTestdataFS(t, "testdata/fixtures", "/fixtures")

	for _, name := range fixtureNames {
		t.Run(name, func(t *testing.T) {
			base := loadManifest(t, mfs, filepath.Join("/fixtures", name, "base.json"))
			head := loadManifest(t, mfs, filepath.Join("/fixtures", name, "head.json"))
			result := breaking.Compare(base, head, breaking.Options{})

			var buf bytes.Buffer
			err := breaking.PrintResult(&buf, result, breaking.DisplayOptions{Format: "markdown"})
			require.NoError(t, err)

			testutil.CheckGolden(t, "markdown-"+name, buf.Bytes(), testutil.GoldenOptions{
				Dir:       "testdata/goldens",
				Extension: ".md",
			})
		})
	}
}

func TestPrintResultJSON(t *testing.T) {
	mfs := testutil.LoadTestdataFS(t, "testdata/fixtures", "/fixtures")
	base := loadManifest(t, mfs, "/fixtures/mixed-changes/base.json")
	head := loadManifest(t, mfs, "/fixtures/mixed-changes/head.json")
	result := breaking.Compare(base, head, breaking.Options{})

	var buf bytes.Buffer
	err := breaking.PrintResult(&buf, result, breaking.DisplayOptions{Format: "json"})
	require.NoError(t, err)

	// inline assertions: verifying JSON round-trip serialization correctness
	var parsed breaking.Result
	require.NoError(t, json.Unmarshal(buf.Bytes(), &parsed))
	require.Equal(t, result.Breaking, parsed.Breaking)
	require.Equal(t, result.Dangerous, parsed.Dangerous)
	require.Equal(t, result.Safe, parsed.Safe)
	require.Len(t, parsed.Changes, len(result.Changes))
}

func TestSeverityUnmarshal(t *testing.T) {
	// inline assertions: verifying JSON round-trip for enum type
	data := []byte(`"breaking"`)
	var s breaking.Severity
	require.NoError(t, json.Unmarshal(data, &s))
	require.Equal(t, breaking.Breaking, s)
}
