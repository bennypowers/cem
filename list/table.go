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
package list

import (
	"errors"
	"fmt"
	"net/url"
	"path"
	"regexp"
	"strings"

	"slices"

	lipgloss "charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/table"
	"github.com/agext/levenshtein"
	"github.com/charmbracelet/x/ansi"

	M "bennypowers.dev/cem/manifest"
)

var (
	isCodeStyleHeader = regexp.MustCompile(`Code|Type|Syntax|Static|Default|DOM Property`)
	isURLHeader       = regexp.MustCompile(`(?i)^(URL|Source|Href)$`)
)

// RenderOptions provides options for the Render function.
type RenderOptions struct {
	Columns         []string
	IncludeSections []string
	Markdown        bool
}

// Render recursively renders a Renderable, creating sectioned tables.
func Render(r M.Renderable, opts RenderOptions, pred M.PredicateFunc) (string, error) {
	if r == nil {
		return "", nil
	}

	var builder strings.Builder

	// Only print the main header if we are not filtering by section.
	if len(opts.IncludeSections) == 0 {
		builder.WriteString(sectionStyle.Render(r.Label()) + "\n\n")

		if d, ok := r.(M.Describable); ok {
			if summary := d.Summary(); summary != "" {
				builder.WriteString(summary + "\n")
			}
			if description := d.Description(); description != "" {
				builder.WriteString(description + "\n")
			}
		}
	}

	if sdp, ok := r.(M.SectionDataProvider); ok {
		for _, section := range sdp.Sections() {
			if len(section.Items) == 0 {
				continue
			}

			// If IncludeSections is specified, only render those sections.
			if len(opts.IncludeSections) > 0 && !slices.Contains(opts.IncludeSections, section.Title) {
				continue
			}

			// Filter items by predicate
			var filteredItems []M.Renderable
			for _, item := range section.Items {
				if pred(item) {
					filteredItems = append(filteredItems, item)
				}
			}

			if len(filteredItems) == 0 {
				continue // no items match predicate
			}

			// The title for a subsection is smaller
			builder.WriteString(sectionStyle.Render(section.Title) + "\n\n")
			headers := filteredItems[0].ColumnHeadings()
			rows := MapToTableRows(filteredItems)

			// Only filter empty columns if the user did not specify columns
			if len(opts.Columns) == 0 {
				headers, rows = RemoveEmptyColumns(headers, rows)
			}

			if len(headers) == 0 {
				continue // all columns empty, nothing to display
			}
			if str, err := formatTable(headers, rows, opts.Columns); err != nil {
				return "", err
			} else {
				builder.WriteString(str + "\n")
			}
		}
	} else {
		for _, child := range r.Children() {
			// Only render children that match the predicate or have matching descendants
			if pred(child) || hasMatchingDescendants(child, pred) {
				// Recursive calls pass the options and predicate down.
				if s, err := Render(child, opts, pred); err != nil {
					return "", err
				} else {
					if strings.TrimSpace(s) != "" {
						builder.WriteString(s + "\n")
					}
				}
			}
		}
	}

	return strings.TrimRight(builder.String(), "\n") + "\n\n", nil
}

// hasMatchingDescendants checks if a renderable has any descendants that match the predicate
func hasMatchingDescendants(r M.Renderable, pred M.PredicateFunc) bool {
	if sdp, ok := r.(M.SectionDataProvider); ok {
		for _, section := range sdp.Sections() {
			if slices.ContainsFunc(section.Items, pred) {
				return true
			}
		}
	}
	for _, child := range r.Children() {
		if pred(child) || hasMatchingDescendants(child, pred) {
			return true
		}
	}
	return false
}

// RenderTable renders a simple table with a title.
func RenderTable(title string, headers []string, rows [][]string, columns []string) (string, error) {
	var builder strings.Builder

	builder.WriteString(sectionStyle.Render(title) + "\n\n")
	str, err := formatTable(headers, rows, columns)
	if err != nil {
		return "", err
	}
	builder.WriteString(str)
	return builder.String(), nil
}

func RenderModulesTable(manifest *M.Package, opts RenderOptions) (string, error) {
	headers := []string{"Module path", "Custom Elements"}
	rows := make([][]string, 0)

	for _, mod := range manifest.Modules {
		customElements := make([]string, 0)
		for _, decl := range mod.Declarations {
			if ce, ok := decl.(*M.CustomElementDeclaration); ok {
				customElements = append(customElements, fmt.Sprintf("<%s>", ce.TagName))
			}
		}
		rows = append(rows, []string{mod.Path, strings.Join(customElements, ", ")})
	}

	if opts.Markdown {
		finalHeaders, finalRows, err := buildTableData(headers, rows, opts.Columns, true)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("### Modules\n\n%s", FormatMarkdownTable(finalHeaders, finalRows)), nil
	}

	str, err := formatTable(headers, rows, opts.Columns)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("## Modules\n\n%s", str), nil
}

func RenderTagsTable(manifest *M.Package, opts RenderOptions) (string, error) {
	headers := []string{"Tag Name", "Class", "Module", "Summary"}
	rows := make([][]string, 0)

	for _, mod := range manifest.Modules {
		for _, decl := range mod.Declarations {
			if ce, ok := decl.(*M.CustomElementDeclaration); ok {
				rows = append(rows, []string{
					"<" + ce.TagName + ">",
					ce.Name(),
					mod.Path,
					ce.Summary,
				})
			}
		}
	}

	if opts.Markdown {
		finalHeaders, finalRows, err := buildTableData(headers, rows, opts.Columns, true)
		if err != nil {
			return "", err
		}
		return FormatMarkdownTable(finalHeaders, finalRows), nil
	}

	return formatTable(headers, rows, opts.Columns)
}

// MapToTableRows maps a slice of Renderables to [][]string.
func MapToTableRows[T M.Renderable](items []T) [][]string {
	rows := make([][]string, 0, len(items))
	for _, item := range items {
		rows = append(rows, item.ToTableRow())
	}
	return rows
}

var (
	tablePadding = lipgloss.NewStyle().PaddingLeft(1).PaddingRight(1)
	codeStyle    = lipgloss.NewStyle().PaddingLeft(1).PaddingRight(1).Foreground(lipgloss.Cyan)
	headerStyle  = lipgloss.NewStyle().PaddingLeft(1).PaddingRight(1).Bold(true)
)

func isCodeColumn(headers []string, col int) bool {
	return col == 0 || (col < len(headers) && isCodeStyleHeader.MatchString(headers[col]))
}

// formatTable renders a styled terminal table.
func formatTable(headers []string, rows [][]string, columns []string) (string, error) {
	if len(columns) == 0 {
		headers, rows = RemoveEmptyColumns(headers, rows)
	}
	finalHeaders, finalRows, err := buildTableData(headers, rows, columns, false)
	if err != nil {
		return "", err
	}
	if len(finalRows) == 0 {
		return "", nil
	}
	finalRows = hyperlinkURLColumns(finalHeaders, finalRows)
	t := table.New().
		Headers(finalHeaders...).
		Rows(finalRows...).
		Border(lipgloss.NormalBorder()).
		BorderTop(false).
		BorderBottom(false).
		BorderLeft(false).
		BorderRight(false).
		BorderHeader(true).
		BorderColumn(true).
		BorderRow(false).
		StyleFunc(func(row, col int) lipgloss.Style {
			if row == table.HeaderRow {
				return headerStyle
			}
			if isCodeColumn(finalHeaders, col) {
				return codeStyle
			}
			return tablePadding
		})
	return t.Render(), nil
}

// buildTableData prepares the filtered headers and rows for the table, given column selection.
// When markdown is true, code columns are wrapped in backticks.
func buildTableData(headers []string, rows [][]string, columns []string, markdown bool) ([]string, [][]string, error) {
	if err := checkUnknownColumns(headers, columns); err != nil {
		return nil, nil, err
	}
	finalHeaders, finalRows := filterTableColumns(headers, rows, columns)
	if markdown {
		finalHeaders, finalRows = backtickCodeColumns(finalHeaders, finalRows)
	}
	return finalHeaders, finalRows, nil
}

func shortenURL(rawURL string) string {
	if rawURL == "" {
		return ""
	}
	u, err := url.Parse(rawURL)
	if err != nil {
		return rawURL
	}
	base := path.Base(u.Path)
	if base == "" || base == "." || base == "/" {
		return rawURL
	}
	return base
}

func formatURLColumns(headers []string, rows [][]string, linkFn func(short, full string) string) [][]string {
	out := make([][]string, len(rows))
	for i, row := range rows {
		newRow := make([]string, len(row))
		copy(newRow, row)
		for j, cell := range row {
			if j < len(headers) && isURLHeader.MatchString(headers[j]) && cell != "" {
				newRow[j] = linkFn(shortenURL(cell), cell)
			}
		}
		out[i] = newRow
	}
	return out
}

var linkStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Blue)

func hyperlinkURLColumns(headers []string, rows [][]string) [][]string {
	return formatURLColumns(headers, rows, func(short, full string) string {
		return linkStyle.Hyperlink(full).Render(short)
	})
}

func markdownLinkURLColumns(headers []string, rows [][]string) [][]string {
	return formatURLColumns(headers, rows, func(short, full string) string {
		return fmt.Sprintf("[%s](%s)", short, full)
	})
}

// Wraps the first cell in a row in backticks, for markdown output
func backtickCodeColumns(
	headers []string,
	rows [][]string,
) (outheaders []string, outrows [][]string) {
	outrows = make([][]string, len(rows))
	for i, row := range rows {
		if len(row) == 0 {
			outrows[i] = nil
			continue
		}
		// Copy the row (so we don't mutate the source)
		newRow := make([]string, len(row))
		copy(newRow, row)
		for j := range row {
			if j == 0 || isCodeStyleHeader.MatchString(headers[j]) {
				if newRow[j] != "" {
					newRow[j] = "`" + row[j] + "`"
				}
			}
		}
		outrows[i] = newRow
	}
	return headers, outrows
}


// checkUnknownColumns returns an error if any column name is not in headers, case-insensitive.
func checkUnknownColumns(headers []string, columns []string) error {
	headerSet := make(map[string]string, len(headers)) // lower-case -> original
	for _, h := range headers {
		headerSet[strings.ToLower(h)] = h
	}
	for _, col := range columns {
		colLower := strings.ToLower(col)
		if _, found := headerSet[colLower]; !found {
			suggestion := closestHeader(col, headers)
			msg := fmt.Sprintf("unknown column %q.", col)
			if suggestion != "" && !strings.EqualFold(suggestion, col) {
				msg += fmt.Sprintf(" Did you mean %q?", suggestion)
			}
			return errors.New(msg)
		}
	}
	return nil
}

// closestHeader returns the header with the smallest Levenshtein distance to col, case-insensitive.
func closestHeader(col string, headers []string) string {
	colLower := strings.ToLower(col)
	minDist := 1000
	closest := ""
	for _, h := range headers {
		dist := levenshtein.Distance(strings.ToLower(h), colLower, nil)
		if dist < minDist {
			minDist = dist
			closest = h
		}
	}
	if minDist <= 3 {
		return closest
	}
	return ""
}

// filterTableColumns filters headers and rows to only include the first column and any named columns in 'columns' (case-insensitive).
// If the 'Summary' column is present in the filtered set but all its values are empty, it is omitted from the results.
// Returns new headers and rows in the filtered order.
func filterTableColumns(headers []string, rows [][]string, columns []string) ([]string, [][]string) {
	var selectedIdx []int
	if len(columns) == 0 {
		// Include all columns by default
		selectedIdx = make([]int, len(headers))
		for i := range headers {
			selectedIdx[i] = i
		}
	} else {
		// Always keep the first column.
		selectedIdx = []int{0}
		colSet := map[int]struct{}{0: {}}
		headerLC := make([]string, len(headers))
		for i, h := range headers {
			headerLC[i] = strings.ToLower(h)
		}
		for _, colName := range columns {
			colNameLC := strings.ToLower(colName)
			for idx, hlc := range headerLC {
				if idx == 0 {
					continue
				}
				if hlc == colNameLC {
					if _, exists := colSet[idx]; !exists {
						selectedIdx = append(selectedIdx, idx)
						colSet[idx] = struct{}{}
					}
				}
			}
		}
	}

	// Always check for "Summary" column in the selected set, and remove if all values are empty
	summaryCol := -1
	for i, idx := range selectedIdx {
		if strings.ToLower(headers[idx]) == "summary" {
			summaryCol = i // position in selectedIdx
			break
		}
	}
	if summaryCol != -1 {
		allEmpty := true
		for _, row := range rows {
			colIdx := selectedIdx[summaryCol]
			if colIdx < len(row) && strings.TrimSpace(row[colIdx]) != "" {
				allEmpty = false
				break
			}
		}
		if allEmpty {
			selectedIdx = slices.Delete(selectedIdx, summaryCol, summaryCol+1)
		}
	}

	// Build filtered headers and rows as before
	filteredHeaders := make([]string, len(selectedIdx))
	for i, idx := range selectedIdx {
		filteredHeaders[i] = headers[idx]
	}
	filteredRows := make([][]string, len(rows))
	for i, row := range rows {
		filtered := make([]string, len(selectedIdx))
		for j, idx := range selectedIdx {
			if idx < len(row) {
				filtered[j] = row[idx]
			}
		}
		filteredRows[i] = filtered
	}
	return filteredHeaders, filteredRows
}

// RenderMarkdown renders a Renderable as valid markdown tables (no ANSI).
func RenderMarkdown(r M.Renderable, opts RenderOptions, pred M.PredicateFunc) (string, error) {
	if r == nil {
		return "", nil
	}

	var builder strings.Builder

	if len(opts.IncludeSections) == 0 {
		label := ansi.Strip(r.Label())
		switch r.(type) {
		case *M.RenderableCustomElementDeclaration:
			if strings.HasPrefix(label, "<") {
				label = "`" + label + "`"
			}
		}
		builder.WriteString("# " + label + "\n\n")

		if d, ok := r.(M.Describable); ok {
			if summary := d.Summary(); summary != "" {
				builder.WriteString(summary + "\n")
			}
			if description := d.Description(); description != "" {
				builder.WriteString(description + "\n")
			}
		}
	}

	if sdp, ok := r.(M.SectionDataProvider); ok {
		for _, section := range sdp.Sections() {
			if len(section.Items) == 0 {
				continue
			}

			if len(opts.IncludeSections) > 0 && !slices.Contains(opts.IncludeSections, section.Title) {
				continue
			}

			var filteredItems []M.Renderable
			for _, item := range section.Items {
				if pred(item) {
					filteredItems = append(filteredItems, item)
				}
			}

			if len(filteredItems) == 0 {
				continue
			}

			builder.WriteString("## " + ansi.Strip(section.Title) + "\n\n")
			headers := filteredItems[0].ColumnHeadings()
			rows := MapToTableRows(filteredItems)

			if len(opts.Columns) == 0 {
				headers, rows = RemoveEmptyColumns(headers, rows)
			}

			if len(headers) == 0 {
				continue
			}
			finalHeaders, finalRows, err := buildTableData(headers, rows, opts.Columns, true)
			if err != nil {
				return "", err
			}
			finalRows = markdownLinkURLColumns(finalHeaders, finalRows)
			if len(finalRows) > 0 {
				builder.WriteString(FormatMarkdownTable(finalHeaders, finalRows))
				builder.WriteString("\n")
			}
		}
	} else {
		for _, child := range r.Children() {
			if pred(child) || hasMatchingDescendants(child, pred) {
				s, err := RenderMarkdown(child, opts, pred)
				if err != nil {
					return "", err
				}
				if strings.TrimSpace(s) != "" {
					builder.WriteString(s + "\n")
				}
			}
		}
	}

	return strings.TrimRight(builder.String(), "\n") + "\n\n", nil
}

// FormatMarkdownTable renders headers and rows as a valid markdown table.
func FormatMarkdownTable(headers []string, rows [][]string) string {
	if len(rows) == 0 {
		return ""
	}

	widths := make([]int, len(headers))
	for i, h := range headers {
		widths[i] = len(h)
	}
	for _, row := range rows {
		for i, cell := range row {
			if i < len(widths) {
				widths[i] = max(widths[i], len(cell))
			}
		}
	}

	var builder strings.Builder

	// Header row
	for i, h := range headers {
		if i > 0 {
			builder.WriteString(" | ")
		}
		fmt.Fprintf(&builder, "%-*s", widths[i], h)
	}
	builder.WriteString("\n")

	// Separator row
	for i, w := range widths {
		if i > 0 {
			builder.WriteString(" | ")
		}
		builder.WriteString(strings.Repeat("-", max(3, w)))
	}
	builder.WriteString("\n")

	// Data rows
	for _, row := range rows {
		for i := range headers {
			if i > 0 {
				builder.WriteString(" | ")
			}
			cell := ""
			if i < len(row) {
				cell = row[i]
			}
			fmt.Fprintf(&builder, "%-*s", widths[i], cell)
		}
		builder.WriteString("\n")
	}

	return builder.String()
}
