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
package cmd

import (
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/pmezard/go-difflib/difflib"
	"github.com/spf13/cobra"
)

var (
	diffAddStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("2"))
	diffRemoveStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("1"))
	diffHunkStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("6"))
	diffHeaderStyle = lipgloss.NewStyle().Bold(true)
)

func printColorDiff(cmd *cobra.Command, diff string) {
	for line := range strings.SplitSeq(diff, "\n") {
		switch {
		case strings.HasPrefix(line, "+++"), strings.HasPrefix(line, "---"):
			cmd.Println(diffHeaderStyle.Render(line))
		case strings.HasPrefix(line, "@@"):
			cmd.Println(diffHunkStyle.Render(line))
		case strings.HasPrefix(line, "+"):
			cmd.Println(diffAddStyle.Render(line))
		case strings.HasPrefix(line, "-"):
			cmd.Println(diffRemoveStyle.Render(line))
		default:
			cmd.Println(line)
		}
	}
}

func unifiedDiff(existing, proposed, filename string) string {
	diff := difflib.UnifiedDiff{
		A:        difflib.SplitLines(existing),
		B:        difflib.SplitLines(proposed),
		FromFile: filename,
		ToFile:   filename,
		Context:  3,
	}
	result, err := difflib.GetUnifiedDiffString(diff)
	if err != nil {
		return ""
	}
	return result
}
