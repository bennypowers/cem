package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/agext/levenshtein"
	"github.com/pterm/pterm"

	M "bennypowers.dev/cem/manifest"
)

// MapToTableRows maps a slice of RenderableMemberWithContext to [][]string.
func MapToTableRows[T M.RenderableMemberWithContext](items []T) [][]string {
	rows := make([][]string, 0, len(items))
	for _, item := range items {
		rows = append(rows, item.ToTableRow())
	}
	return rows
}

// Render a table given headers, rows, and columns to display.
// If columns is empty, renders all columns. Otherwise, always renders the first column and any columns listed in 'columns' (by header name), without duplicates.
func RenderTable(headers []string, rows [][]string, columns []string) error {
	if err := checkUnknownColumns(headers, columns); err != nil {
		return err
	}
	table := pterm.DefaultTable.
		WithHasHeader(true).
		WithBoxed(false)
	finalHeaders, finalRows := filterTableColumns(headers, rows, columns)
	data := pterm.TableData{finalHeaders}
	data = append(data, finalRows...)
	return table.
		WithData(data).
		Render()
}

// checkUnknownColumns returns an error if any column name is not in headers.
// If a column is unknown, it suggests a close match.
func checkUnknownColumns(headers []string, columns []string) error {
	headerSet := make(map[string]struct{}, len(headers))
	for _, h := range headers {
		headerSet[h] = struct{}{}
	}
	for _, col := range columns {
		if _, found := headerSet[col]; !found {
			suggestion := closestHeader(col, headers)
			msg := fmt.Sprintf("unknown column %q.", col)
			if suggestion != "" && suggestion != col {
				msg += fmt.Sprintf(" Did you mean %q?", suggestion)
			}
			return errors.New(msg)
		}
	}
	return nil
}

// closestHeader returns the header with the smallest Levenshtein distance, or "" if none.
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
	if minDist <= 3 { // Only suggest if it's a reasonably close typo
		return closest
	}
	return ""
}

// filterTableColumns filters headers and rows to only include the first column and any named columns in 'columns'.
// Returns new headers and rows in the filtered order.
func filterTableColumns(headers []string, rows [][]string, columns []string) ([]string, [][]string) {
	if len(columns) == 0 {
		return headers, rows
	}
	// Always keep the first column.
	selectedIdx := []int{0}
	colSet := map[int]struct{}{0: {}}
	for _, colName := range columns {
		for idx, header := range headers {
			if idx == 0 {
				continue
			}
			if header == colName {
				if _, exists := colSet[idx]; !exists {
					selectedIdx = append(selectedIdx, idx)
					colSet[idx] = struct{}{}
				}
			}
		}
	}
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
