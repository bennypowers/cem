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
	"os"
	"strings"
	"text/template"

	"github.com/pterm/pterm"
)

//go:embed templates/report.md.tmpl
var reportTemplate string

// DisplayOptions configures health output formatting.
type DisplayOptions struct {
	Format string
}

// PrintHealthResult prints the health result with the specified format.
func PrintHealthResult(result *HealthResult, options DisplayOptions) error {
	switch options.Format {
	case "json":
		return printHealthResultJSON(result)
	case "markdown":
		return writeMarkdownReport(os.Stdout, result)
	case "text", "":
		printHealthResultText(result)
		return nil
	default:
		return fmt.Errorf("invalid format: %s. Use 'text', 'json', or 'markdown'", options.Format)
	}
}

func printHealthResultJSON(result *HealthResult) error {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(result)
}

func printHealthResultText(result *HealthResult) {
	if len(result.Modules) == 0 {
		pterm.Info.Println("No declarations found in manifest")
		return
	}

	pterm.DefaultHeader.WithFullWidth().Println("Component Health Report")
	pterm.Println()

	for _, mod := range result.Modules {
		printModuleReport(mod)
	}

	// Overall score
	pct := percentage(result.OverallScore, result.OverallMax)
	style := scoreStyle(pct)
	pterm.Println()
	pterm.Printf("Overall: %s %d/%d\n",
		style.Sprint(buildBar(pct, 40)),
		result.OverallScore,
		result.OverallMax)
	pterm.Println()

	// Recommendations
	if len(result.Recommendations) > 0 {
		pterm.DefaultSection.Println("Recommendations")
		items := make([]pterm.BulletListItem, len(result.Recommendations))
		for i, rec := range result.Recommendations {
			items[i] = pterm.BulletListItem{
				Level: 0,
				Text:  rec,
			}
		}
		_ = pterm.DefaultBulletList.WithItems(items).Render()
	}
}

func printModuleReport(mod ModuleReport) {
	modPct := percentage(mod.Score, mod.MaxScore)
	modStyle := scoreStyle(modPct)

	pterm.DefaultSection.WithLevel(1).Println(
		fmt.Sprintf("%s %s %d/%d",
			pterm.FgLightBlue.Sprint(mod.Path),
			modStyle.Sprint(buildBar(modPct, 20)),
			mod.Score,
			mod.MaxScore))

	for _, decl := range mod.Declarations {
		printComponentReport(decl)
	}
}

func printComponentReport(comp ComponentReport) {
	name := comp.TagName
	if name == "" {
		name = comp.Name
	}

	pct := percentage(comp.Score, comp.MaxScore)
	style := scoreStyle(pct)

	pterm.DefaultSection.WithLevel(2).Println(
		fmt.Sprintf("%s %s %d/%d",
			name,
			style.Sprint(buildBar(pct, 20)),
			comp.Score,
			comp.MaxScore))

	items := make([]pterm.BulletListItem, 0, len(comp.Categories))
	for _, cat := range comp.Categories {
		icon := pterm.FgGreen.Sprint("✓")
		switch cat.Status {
		case "warn":
			icon = pterm.FgYellow.Sprint("⚠")
		case "fail":
			icon = pterm.FgRed.Sprint("✗")
		}

		msg := cat.Category
		if cat.Status != "pass" {
			for _, f := range cat.Findings {
				if f.Message != "" && f.Points < f.Max {
					msg = fmt.Sprintf("%s — %s", cat.Category, f.Message)
					break
				}
			}
		}

		items = append(items, pterm.BulletListItem{
			Level:      0,
			Text:       fmt.Sprintf("%s %s (%d/%d)", icon, msg, cat.Points, cat.MaxPoints),
			BulletStyle: pterm.NewStyle(),
			Bullet:     " ",
		})
	}
	_ = pterm.DefaultBulletList.WithItems(items).Render()
}

func buildBar(pct, width int) string {
	filled := min(pct*width/100, width)
	empty := width - filled
	return strings.Repeat("█", filled) + strings.Repeat("░", empty)
}

func percentage(score, max int) int {
	if max == 0 {
		return 0
	}
	return score * 100 / max
}

func scoreStyle(pct int) pterm.Style {
	if pct >= 80 {
		return *pterm.NewStyle(pterm.FgGreen)
	}
	if pct >= 40 {
		return *pterm.NewStyle(pterm.FgYellow)
	}
	return *pterm.NewStyle(pterm.FgRed)
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
