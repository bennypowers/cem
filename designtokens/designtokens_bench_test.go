/*
Copyright 2026 Benny Powers. All rights reserved.
Use of this source code is governed by the GPLv3
license that can be found in the LICENSE file.
*/

package designtokens

import (
	"encoding/json"
	"fmt"
	"testing"
)

// generateTokenFile creates a DTCG-format token file with the specified number of tokens
func generateTokenFile(numTokens int) []byte {
	tokens := make(map[string]any)
	tokens["$schema"] = "https://www.designtokens.org/schemas/draft.json"

	// Create color group
	colors := make(map[string]any)
	colors["$type"] = "color"
	for i := 0; i < numTokens/3; i++ {
		colors[fmt.Sprintf("color%d", i)] = map[string]any{
			"$value":       fmt.Sprintf("#%06x", i*1000),
			"$description": fmt.Sprintf("Color token %d", i),
		}
	}
	tokens["color"] = colors

	// Create spacing group
	spacing := make(map[string]any)
	spacing["$type"] = "dimension"
	for i := 0; i < numTokens/3; i++ {
		spacing[fmt.Sprintf("space%d", i)] = map[string]any{
			"$value":       fmt.Sprintf("%dpx", i*4),
			"$description": fmt.Sprintf("Spacing token %d", i),
		}
	}
	tokens["spacing"] = spacing

	// Create typography group with nesting
	typography := make(map[string]any)
	heading := make(map[string]any)
	heading["$type"] = "dimension"
	for i := 0; i < numTokens/3; i++ {
		heading[fmt.Sprintf("h%d", i)] = map[string]any{
			"$value":       fmt.Sprintf("%drem", i+1),
			"$description": fmt.Sprintf("Heading %d size", i),
		}
	}
	typography["heading"] = heading
	tokens["typography"] = typography

	data, err := json.Marshal(tokens)
	if err != nil {
		panic(fmt.Sprintf("generateTokenFile: %v", err))
	}
	return data
}

func BenchmarkParseTokensWithAsimonim(b *testing.B) {
	sizes := []int{100, 500, 1000, 5000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("tokens=%d", size), func(b *testing.B) {
			data := generateTokenFile(size)

			b.ResetTimer()
			b.ReportAllocs()

			for b.Loop() {
				_, err := ParseTokensWithAsimonim(data, ParseOptions{})
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}
