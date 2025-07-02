package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/agext/levenshtein"
	"github.com/pterm/pterm"

	M "bennypowers.dev/cem/manifest"
	"slices"
)

// MapToTableRows maps a slice of RenderableMemberWithContext to [][]string.
func MapToTableRows[T M.RenderableMemberWithContext](items []T) [][]string {
	rows := make([][]string, 0, len(items))
	for _, item := range items {
		rows = append(rows, item.ToTableRow())
	}
	return rows
}

// RenderTable renders a table given headers, rows, and columns to display.
// If columns is empty, renders all columns. Otherwise, always renders the first column and any columns listed in 'columns' (by header name), without duplicates.
// Now returns error for error handling.
func RenderTable(title string, headers []string, rows [][]string, columns []string) error {
	if err := checkUnknownColumns(headers, columns); err != nil {
		return err
	}
	table := pterm.DefaultTable.
		WithHasHeader(true).
		WithBoxed(false)
	finalHeaders, finalRows := filterTableColumns(headers, rows, columns)
	if len(finalRows) == 0 {
		return nil
	}
	data := pterm.TableData{finalHeaders}
	data = append(data, finalRows...)
	out, err := table.WithData(data).Srender()
	if err != nil {
		return err
	}
	pterm.DefaultSection.Println(title)
	pterm.Println(out)
	return nil
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
