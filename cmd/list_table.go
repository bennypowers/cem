package cmd

import (
	"github.com/pterm/pterm"

	M "bennypowers.dev/cem/manifest"
)

// MapToTableRows maps a slice of RenderableMemberWithContext to [][]string,
// applying the columns filter. If columns is empty, renders all columns.
// Otherwise, always renders the first column and any columns listed in 'columns' (by header name), without duplicates.
func MapToTableRows[T M.RenderableMemberWithContext](items []T, headers []string, columns []string) [][]string {
	if len(columns) == 0 {
		// No filter, map all
		rows := make([][]string, 0, len(items))
		for _, item := range items {
			rows = append(rows, item.ToTableRow())
		}
		return rows
	}

	// Build selected column indices: always include first column, and any additional columns by name
	selected := []int{0}
	colSet := map[int]struct{}{0: {}}

	for _, colName := range columns {
		for idx, header := range headers {
			if idx == 0 {
				continue // already included
			}
			if header == colName {
				if _, exists := colSet[idx]; !exists {
					selected = append(selected, idx)
					colSet[idx] = struct{}{}
				}
			}
		}
	}

	// Map and filter each row
	rows := make([][]string, 0, len(items))
	for _, item := range items {
		row := item.ToTableRow()
		filtered := make([]string, len(selected))
		for i, idx := range selected {
			if idx < len(row) {
				filtered[i] = row[idx]
			}
		}
		rows = append(rows, filtered)
	}
	return rows
}

// Render a table given headers, rows, and columns to display.
// If columns is empty, renders all columns. Otherwise, always renders the first column and any columns listed in 'columns' (by header name), without duplicates.
func RenderTable(headers []string, rows [][]string, columns []string) {
	table := pterm.DefaultTable.
		WithHasHeader(true).
		WithBoxed(false)
	// If columns is empty, render as usual
	if len(columns) == 0 {
		data := pterm.TableData{headers}
		data = append(data, rows...)
		table.
			WithData(data).
			Render()
		return
	}

	// Build selected column indices: always include first column, and any additional columns by name
	selected := []int{0}
	colSet := map[int]struct{}{0: {}}

	for _, colName := range columns {
		for idx, header := range headers {
			if idx == 0 {
				continue // already included
			}
			if header == colName {
				if _, exists := colSet[idx]; !exists {
					selected = append(selected, idx)
					colSet[idx] = struct{}{}
				}
			}
		}
	}

	// Filter headers and rows
	filteredHeaders := make([]string, len(selected))
	for i, idx := range selected {
		filteredHeaders[i] = headers[idx]
	}
	filteredRows := make([][]string, len(rows))
	for i, row := range rows {
		filtered := make([]string, len(selected))
		for j, idx := range selected {
			if idx < len(row) {
				filtered[j] = row[idx]
			}
		}
		filteredRows[i] = filtered
	}

	data := pterm.TableData{filteredHeaders}
	data = append(data, filteredRows...)
	table.
		WithData(data).
		Render()
}
