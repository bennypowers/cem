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

import "fmt"

// RemoveEmptyColumns filters out columns where all values in all rows are empty ("").
// Always retains the first column (assumed identifier).
func RemoveEmptyColumns(headers []string, rows [][]string) ([]string, [][]string) {
	if len(headers) == 0 || len(rows) == 0 {
		return headers, rows
	}
	keep := make([]bool, len(headers))
	for col := range headers {
		for _, row := range rows {
			if col < len(row) && row[col] != "" {
				keep[col] = true
				break
			}
		}
	}
	// Always keep the first column
	keep[0] = true

	var outHeaders []string
	var colIdx []int
	for i, k := range keep {
		if k {
			outHeaders = append(outHeaders, headers[i])
			colIdx = append(colIdx, i)
		}
	}
	outRows := make([][]string, len(rows))
	for i, row := range rows {
		r := make([]string, 0, len(colIdx))
		for _, idx := range colIdx {
			if idx < len(row) {
				r = append(r, row[idx])
			} else {
				r = append(r, "")
			}
		}
		outRows[i] = r
	}
	return outHeaders, outRows
}

// filterToUserColumns filters headers and rows to requested columns, always retaining the first (identifier) column.
func filterToUserColumns(headers []string, rows [][]string, columns []string) ([]string, [][]string, error) {
    headerMap := make(map[string]int)
    for i, h := range headers {
        headerMap[h] = i
    }
    colIdx := []int{0} // always include first column
    used := map[int]bool{0: true}
    for _, c := range columns {
        if idx, ok := headerMap[c]; ok && !used[idx] {
            colIdx = append(colIdx, idx)
            used[idx] = true
        } else if !ok {
            return nil, nil, fmt.Errorf("unknown column: %s", c)
        }
    }
    filteredHeaders := make([]string, len(colIdx))
    for i, idx := range colIdx {
        filteredHeaders[i] = headers[idx]
    }
    filteredRows := make([][]string, len(rows))
    for i, row := range rows {
        filteredRows[i] = make([]string, len(colIdx))
        for j, idx := range colIdx {
            if idx < len(row) {
                filteredRows[i][j] = row[idx]
            } else {
                filteredRows[i][j] = ""
            }
        }
    }
    return filteredHeaders, filteredRows, nil
}
