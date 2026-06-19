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
	"testing"

	"bennypowers.dev/cem/internal/platform/testutil"
	"bennypowers.dev/cem/list"
	M "bennypowers.dev/cem/manifest"
)

func TestRender(t *testing.T) {
	t.Run("custom-element-table-coverage", func(t *testing.T) {
		var pkg M.Package
		testutil.LoadJSONFixture(t, "custom-element-table-coverage/input.json", &pkg)

		renderable := M.NewRenderablePackage(&pkg)

		opts := list.RenderOptions{}
		output, err := list.Render(renderable, opts, M.True)
		if err != nil {
			t.Fatalf("Render failed: %v", err)
		}

		testutil.CheckGolden(t, "custom-element-table-coverage/expected.md", []byte(output), testutil.GoldenOptions{
			Dir:          "testdata",
			StripANSI:    true,
			NormalizeEOL: true,
		})
	})
}

func TestRenderMarkdown(t *testing.T) {
	t.Run("custom-element-table-coverage", func(t *testing.T) {
		var pkg M.Package
		testutil.LoadJSONFixture(t, "custom-element-table-coverage/input.json", &pkg)

		renderable := M.NewRenderablePackage(&pkg)

		opts := list.RenderOptions{}
		output, err := list.RenderMarkdown(renderable, opts, M.True)
		if err != nil {
			t.Fatalf("RenderMarkdown failed: %v", err)
		}

		testutil.CheckGolden(t, "custom-element-table-coverage/expected-markdown.md", []byte(output), testutil.GoldenOptions{
			Dir:          "testdata",
			NormalizeEOL: true,
		})
	})
}

func TestFormatMarkdownTable(t *testing.T) {
	t.Run("basic table", func(t *testing.T) {
		headers := []string{"Name", "Type", "Summary"}
		rows := [][]string{
			{"foo", "string", "A foo"},
			{"bar", "number", ""},
		}
		got := list.FormatMarkdownTable(headers, rows)
		testutil.CheckGolden(t, "markdown-table-basic.md", []byte(got), testutil.GoldenOptions{
			Dir:          "testdata",
			NormalizeEOL: true,
		})
	})

	t.Run("pipe in cell", func(t *testing.T) {
		headers := []string{"Name", "Type"}
		rows := [][]string{
			{"variant", `"primary" | "secondary"`},
		}
		got := list.FormatMarkdownTable(headers, rows)
		testutil.CheckGolden(t, "markdown-table-pipe.md", []byte(got), testutil.GoldenOptions{
			Dir:          "testdata",
			NormalizeEOL: true,
		})
	})

	// inline: empty input is a zero-value identity check, no golden needed
	t.Run("empty rows", func(t *testing.T) {
		headers := []string{"Name", "Type"}
		var rows [][]string
		got := list.FormatMarkdownTable(headers, rows)
		if got != "" {
			t.Errorf("expected empty string for empty rows, got %q", got)
		}
	})
}
