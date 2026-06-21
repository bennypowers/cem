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
package health

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"text/template"

	"charm.land/bubbles/v2/progress"
	lipgloss "charm.land/lipgloss/v2"

	"bennypowers.dev/cem/internal/logging"
)

var (
	headerStyle    = lipgloss.NewStyle().Bold(true)
	sectionStyle   = lipgloss.NewStyle().Bold(true).Underline(true)
	greenStyle     = lipgloss.NewStyle().Foreground(lipgloss.Green)
	yellowStyle    = lipgloss.NewStyle().Foreground(lipgloss.Yellow)
	redStyle       = lipgloss.NewStyle().Foreground(lipgloss.Red)
	lightBlueStyle = lipgloss.NewStyle().Foreground(lipgloss.BrightBlue)
)

//go:embed templates/report.md.tmpl
var reportTemplate string

// DisplayOptions configures health output formatting.
type DisplayOptions struct {
	Format string
}

// PrintHealthResult prints the health result with the specified format.
func PrintHealthResult(w io.Writer, result *HealthResult, options DisplayOptions) error {
	switch options.Format {
	case "json":
		return printHealthResultJSON(w, result)
	case "markdown":
		return writeMarkdownReport(w, result)
	case "text", "":
		return printHealthResultText(w, result)
	default:
		return fmt.Errorf("invalid format: %s. Use 'text', 'json', or 'markdown'", options.Format)
	}
}

// PrintWorkspaceHealthResults prints collected workspace health results as a single output.
func PrintWorkspaceHealthResults(w io.Writer, results []PackageHealthResult, options DisplayOptions) error {
	switch options.Format {
	case "json":
		encoder := json.NewEncoder(w)
		encoder.SetIndent("", "  ")
		return encoder.Encode(results)
	case "markdown":
		for _, r := range results {
			if _, err := fmt.Fprintf(w, "# %s\n\n", r.Package); err != nil {
				return err
			}
			if err := writeMarkdownReport(w, r.Result); err != nil {
				return err
			}
			if _, err := fmt.Fprintln(w); err != nil {
				return err
			}
		}
		return nil
	default:
		return fmt.Errorf("invalid format: %s", options.Format)
	}
}

func printHealthResultJSON(w io.Writer, result *HealthResult) error {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	return encoder.Encode(result)
}

func printHealthResultText(w io.Writer, result *HealthResult) error {
	if len(result.Modules) == 0 {
		logging.Warning("No declarations found in manifest")
		return nil
	}

	if _, err := lipgloss.Fprintln(w, headerStyle.Render("Component Health Report")); err != nil {
		return err
	}
	if _, err := lipgloss.Fprintln(w); err != nil {
		return err
	}

	for _, mod := range result.Modules {
		if err := printModuleReport(w, mod); err != nil {
			return err
		}
	}

	pct := percentage(result.OverallScore, result.OverallMax)
	style := scoreStyle(pct)
	if _, err := lipgloss.Fprintln(w); err != nil {
		return err
	}
	if _, err := lipgloss.Fprintf(w, "Overall: %s %d/%d\n",
		style.Render(buildBar(pct, 40)),
		result.OverallScore,
		result.OverallMax); err != nil {
		return err
	}
	if _, err := lipgloss.Fprintln(w); err != nil {
		return err
	}

	if len(result.Recommendations) > 0 {
		if _, err := lipgloss.Fprintln(w, sectionStyle.Render("Recommendations")); err != nil {
			return err
		}
		for _, rec := range result.Recommendations {
			if _, err := lipgloss.Fprintf(w, "  • %s\n", rec); err != nil {
				return err
			}
		}
	}
	return nil
}

func printModuleReport(w io.Writer, mod ModuleReport) error {
	modPct := percentage(mod.Score, mod.MaxScore)
	modStyle := scoreStyle(modPct)

	if _, err := lipgloss.Fprintf(w, "\n%s %s %d/%d\n",
		lightBlueStyle.Render(mod.Path),
		modStyle.Render(buildBar(modPct, 20)),
		mod.Score,
		mod.MaxScore); err != nil {
		return err
	}

	for _, decl := range mod.Declarations {
		if err := printComponentReport(w, decl); err != nil {
			return err
		}
	}
	return nil
}

func printComponentReport(w io.Writer, comp ComponentReport) error {
	name := comp.TagName
	if name == "" {
		name = comp.Name
	}

	pct := percentage(comp.Score, comp.MaxScore)
	style := scoreStyle(pct)

	if _, err := lipgloss.Fprintf(w, "\n  %s %s %d/%d\n",
		name,
		style.Render(buildBar(pct, 20)),
		comp.Score,
		comp.MaxScore); err != nil {
		return err
	}

	for _, cat := range comp.Categories {
		icon := greenStyle.Render("✓")
		switch cat.Status {
		case "warn":
			icon = yellowStyle.Render("⚠")
		case "fail":
			icon = redStyle.Render("✗")
		}

		msg := cat.Category
		if cat.Status != "pass" {
			for _, f := range cat.Findings {
				if f.Message != "" && f.Points < f.Max {
					msg = fmt.Sprintf("%s -- %s", cat.Category, f.Message)
					break
				}
			}
		}

		if _, err := lipgloss.Fprintf(w, "    %s %s (%d/%d)\n", icon, msg, cat.Points, cat.MaxPoints); err != nil {
			return err
		}
	}
	return nil
}

func buildBar(pct, width int) string {
	bar := progress.New(
		progress.WithWidth(width),
		progress.WithoutPercentage(),
		progress.WithDefaultBlend(),
	)
	return bar.ViewAs(float64(pct) / 100)
}

func percentage(score, max int) int {
	if max == 0 {
		return 0
	}
	return score * 100 / max
}

func scoreStyle(pct int) lipgloss.Style {
	if pct >= 80 {
		return greenStyle
	}
	if pct >= 40 {
		return yellowStyle
	}
	return redStyle
}

var markdownTmpl = template.Must(template.New("report").Funcs(template.FuncMap{
	"percentage": percentage,
	"statusEmoji": func(score, max int) string {
		pct := percentage(score, max)
		if pct >= 80 {
			return ":white_check_mark:"
		}
		if pct >= 40 {
			return ":warning:"
		}
		return ":x:"
	},
	"summaryLabel": func(tagName, name string) string {
		if tagName != "" && tagName != name {
			return fmt.Sprintf("<code>%s</code> (%s)", tagName, name)
		}
		if tagName != "" {
			return fmt.Sprintf("<code>%s</code>", tagName)
		}
		return fmt.Sprintf("<code>%s</code>", name)
	},
	"inc": func(i int) int { return i + 1 },
}).Parse(reportTemplate))

func writeMarkdownReport(w io.Writer, result *HealthResult) error {
	return markdownTmpl.Execute(w, result)
}
