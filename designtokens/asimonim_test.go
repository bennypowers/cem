/*
Copyright 2026 Benny Powers. All rights reserved.
Use of this source code is governed by the GPLv3
license that can be found in the LICENSE file.
*/

package designtokens

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

// expectedToken represents the expected token structure in golden files.
type expectedToken struct {
	Value       string `json:"value"`
	Description string `json:"description"`
	Syntax      string `json:"syntax"`
}

func TestParseTokensWithAsimonim(t *testing.T) {
	testCases := []struct {
		name       string
		fixtureDir string
	}{
		{"draft-basic", "draft-basic"},
		{"draft-aliases", "draft-aliases"},
		{"v2025-structured-colors", "v2025-structured-colors"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tokensPath := filepath.Join("testdata", tc.fixtureDir, "tokens.json")
			expectedPath := filepath.Join("testdata", tc.fixtureDir, "expected.json")

			// Read input tokens
			tokensData, err := os.ReadFile(tokensPath)
			if err != nil {
				t.Fatalf("failed to read tokens.json: %v", err)
			}

			// Read expected output
			expectedData, err := os.ReadFile(expectedPath)
			if err != nil {
				t.Fatalf("failed to read expected.json: %v", err)
			}

			var expected map[string]expectedToken
			if err := json.Unmarshal(expectedData, &expected); err != nil {
				t.Fatalf("failed to parse expected.json: %v", err)
			}

			// Parse tokens
			tokenMap, err := ParseTokensWithAsimonim(tokensData, ParseOptions{Prefix: "test"})
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			// Verify token count
			if tokenMap.Len() != len(expected) {
				t.Errorf("expected %d tokens, got %d", len(expected), tokenMap.Len())
			}

			// Verify each expected token
			for name, exp := range expected {
				tok, ok := tokenMap.Get(name)
				if !ok {
					t.Errorf("missing token: %s", name)
					continue
				}
				if tok.DisplayValue() != exp.Value {
					t.Errorf("%s: expected value %q, got %q", name, exp.Value, tok.DisplayValue())
				}
				if tok.Description != exp.Description {
					t.Errorf("%s: expected description %q, got %q", name, exp.Description, tok.Description)
				}
				if tok.CSSSyntax() != exp.Syntax {
					t.Errorf("%s: expected syntax %q, got %q", name, exp.Syntax, tok.CSSSyntax())
				}
			}
		})
	}
}
