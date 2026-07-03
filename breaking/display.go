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
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"text/template"

	lipgloss "charm.land/lipgloss/v2"

	"bennypowers.dev/cem/internal/tui"
)

//go:embed templates/report.md.tmpl
var reportTemplate string

type DisplayOptions struct {
	Format string
}

func PrintResult(w io.Writer, result *Result, opts DisplayOptions) error {
	switch opts.Format {
	case "json":
		return printResultJSON(w, result)
	case "markdown":
		return printResultMarkdown(w, result)
	case "text", "":
		return printResultText(w, result)
	default:
		return fmt.Errorf("invalid format: %s. Use 'text', 'json', or 'markdown'", opts.Format)
	}
}

func printResultJSON(w io.Writer, result *Result) error {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	return encoder.Encode(result)
}

func printResultText(w io.Writer, result *Result) error {
	if len(result.Changes) == 0 {
		_, err := lipgloss.Fprintln(w, tui.SuccessStyle.Render("No breaking changes detected"))
		return err
	}

	sections := []struct {
		severity Severity
		title    string
		style    lipgloss.Style
		icon     string
	}{
		{Breaking, "Breaking Changes", tui.ErrorStyle, "✗"},
		{Dangerous, "Dangerous Changes", tui.WarnStyle, "⚠"},
		{Safe, "Safe Changes", tui.SuccessStyle, "✓"},
	}

	for _, sec := range sections {
		var items []Change
		for _, c := range result.Changes {
			if c.Severity == sec.severity {
				items = append(items, c)
			}
		}
		if len(items) == 0 {
			continue
		}

		header := fmt.Sprintf("%s (%d)", sec.title, len(items))
		if _, err := lipgloss.Fprintln(w, tui.SectionStyle.Render(header)); err != nil {
			return err
		}
		for _, c := range items {
			icon := sec.style.Render(sec.icon)
			if _, err := lipgloss.Fprintf(w, "  %s %s\n", icon, c.Message); err != nil {
				return err
			}
		}
		if _, err := lipgloss.Fprintln(w); err != nil {
			return err
		}
	}

	return nil
}

var markdownTmpl = template.Must(template.New("report").Funcs(template.FuncMap{
	"escapeCell": func(s string) string {
		return strings.ReplaceAll(s, "|", "\\|")
	},
	"severities": func() []Severity {
		return []Severity{Breaking, Dangerous, Safe}
	},
	"severityTitle": func(s Severity) string {
		switch s {
		case Breaking:
			return "Breaking"
		case Dangerous:
			return "Dangerous"
		case Safe:
			return "Safe"
		default:
			return "Unknown"
		}
	},
	"filterBySeverity": func(changes []Change, s Severity) []Change {
		var filtered []Change
		for _, c := range changes {
			if c.Severity == s {
				filtered = append(filtered, c)
			}
		}
		return filtered
	},
}).Parse(reportTemplate))

func printResultMarkdown(w io.Writer, result *Result) error {
	return markdownTmpl.Execute(w, result)
}
