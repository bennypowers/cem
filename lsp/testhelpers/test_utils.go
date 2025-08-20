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
package testhelpers

import (
	"os"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

// GetCompletionLabels extracts labels from completion items for easy testing
func GetCompletionLabels(completions []protocol.CompletionItem) []string {
	labels := make([]string, len(completions))
	for i, completion := range completions {
		labels[i] = completion.Label
	}
	return labels
}

// CopyFile copies a file from src to dst
func CopyFile(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	return os.WriteFile(dst, data, 0644)
}
