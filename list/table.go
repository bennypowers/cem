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
	"regexp"
	"strings"

	"slices"

	"github.com/agext/levenshtein"
	"github.com/pterm/pterm"

	M "bennypowers.dev/cem/manifest"
)

var isCodeStyleHeader = regexp.MustCompile(`Code|Type|Syntax|Static|Default|DOM Property`)

// RenderOptions provides options for the Render function.
type RenderOptions struct {
	Columns         []string
	IncludeSections []string
}

// Render recursively renders a Renderable, creating sectioned tables.
func Render(r M.Renderable, opts RenderOptions, pred M.PredicateFunc) (string, error) {
	if r == nil {
		return "", nil
	}

	var builder strings.Builder

	// Only print the main header if we are not filtering by section.
	if len(opts.IncludeSections) == 0 {
		label := r.Label()
		switch r.(type) {
		case *M.RenderableCustomElementDeclaration:
			if strings.HasPrefix(label, "<") {
				label = "`" + label + "`"
			}
		}
		formatted := pterm.DefaultSection.WithLevel(1).Sprint(label)
		builder.WriteString(strings.TrimSpace(formatted) + "\n\n")

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
			formatted := pterm.DefaultSection.WithLevel(2).Sprint(section.Title)
			builder.WriteString(strings.TrimSpace(formatted) + "\n\n")
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

	formatted := pterm.DefaultSection.Sprint(title)
	builder.WriteString(strings.TrimSpace(formatted) + "\n\n")
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
				customElements = append(customElements, fmt.Sprintf("`<%s>`", ce.TagName))
			}
		}
		rows = append(rows, []string{mod.Path, strings.Join(customElements, ", ")})
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

// formatTable renders a basic table, used by the main Render function.
func formatTable(headers []string, rows [][]string, columns []string) (string, error) {
	// If the user did not specify columns, filter out all-empty columns (except the first column).
	if len(columns) == 0 {
		headers, rows = RemoveEmptyColumns(headers, rows)
	}
	finalHeaders, finalRows, err := buildTableData(headers, rows, columns)
	if err != nil {
		return "", err
	}
	if len(finalRows) == 0 {
		return "", nil
	}
	table := pterm.DefaultTable.
		WithHasHeader(true).
		WithBoxed(false)
	data := pterm.TableData{finalHeaders}
	data = append(data, finalRows...)
	return table.WithData(data).Srender()
}

// buildTableData prepares the filtered headers and rows for the table, given column selection.
// Returns error if unknown columns are requested.
func buildTableData(headers []string, rows [][]string, columns []string) ([]string, [][]string, error) {
	if err := checkUnknownColumns(headers, columns); err != nil {
		return nil, nil, err
	}
	finalHeaders, finalRows := insertMarkdownHeaderRow(backtickCodeColumns(filterTableColumns(headers, rows, columns)))
	return finalHeaders, finalRows, nil
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

func insCopy(dest [][]string, row []string, i int) {
	// Copy the row (so we don't mutate the source)
	newRow := make([]string, len(row))
	copy(newRow, row)
	dest[i] = newRow
}

func insertMarkdownHeaderRow(headers []string, rows [][]string) ([]string, [][]string) {
	if len(rows) == 0 {
		return headers, rows
	}

	out := make([][]string, len(rows)+2)

	columnWidths := make([]int, len(headers))

	for i, header := range headers {
		columnWidths[i] = len(header)
	}
	for _, row := range rows {
		for i, cell := range row {
			columnWidths[i] = max(columnWidths[i], len(cell))
		}
	}

	// don't insert headers, that happened in formatTable
	sep := make([]string, len(headers))
	for j, width := range columnWidths {
		sep[j] = strings.Repeat("-", max(3, width))
	}
	out[1] = sep

	for i, row := range rows {
		insCopy(out, row, i+2)
	}
	return headers, out
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
