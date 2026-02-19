/*
Copyright 2026 Benny Powers. All rights reserved.
Use of this source code is governed by the GPLv3
license that can be found in the LICENSE file.
*/

package designtokens

import (
	"encoding/json"
	"fmt"
	"io"
	"path/filepath"
	"strings"
	"testing"

	C "bennypowers.dev/cem/cmd/config"
	"bennypowers.dev/cem/internal/platform/testutil"
	M "bennypowers.dev/cem/manifest"
	"bennypowers.dev/cem/types"
)

// expectedToken represents the expected token structure in golden files.
type expectedToken struct {
	Value       string `json:"value"`
	Description string `json:"description"`
	Syntax      string `json:"syntax"`
}

// stubWorkspaceContext implements types.WorkspaceContext for testing LoadDesignTokens.
type stubWorkspaceContext struct {
	root   string
	config *C.CemConfig
}

func (s *stubWorkspaceContext) Root() string                         { return s.root }
func (s *stubWorkspaceContext) Config() (*C.CemConfig, error)        { return s.config, nil }
func (s *stubWorkspaceContext) Init() error                          { return nil }
func (s *stubWorkspaceContext) ConfigFile() string                   { return "" }
func (s *stubWorkspaceContext) PackageJSON() (*M.PackageJSON, error) { return nil, nil }
func (s *stubWorkspaceContext) Manifest() (*M.Package, error)        { return nil, nil }
func (s *stubWorkspaceContext) CustomElementsManifestPath() string   { return "" }
func (s *stubWorkspaceContext) ReadFile(path string) (io.ReadCloser, error) {
	return nil, fmt.Errorf("stubWorkspaceContext.ReadFile(%q): not implemented", path)
}
func (s *stubWorkspaceContext) Glob(string) ([]string, error)               { return nil, nil }
func (s *stubWorkspaceContext) OutputWriter(string) (io.WriteCloser, error) { return nil, nil }
func (s *stubWorkspaceContext) Cleanup() error                              { return nil }
func (s *stubWorkspaceContext) ModulePathToFS(string) string                { return "" }
func (s *stubWorkspaceContext) FSPathToModule(string) (string, error)       { return "", nil }
func (s *stubWorkspaceContext) ResolveModuleDependency(string, string) (string, error) {
	return "", nil
}
func (s *stubWorkspaceContext) DesignTokensCache() types.DesignTokensCache { return nil }

func TestValidatePrefix(t *testing.T) {
	t.Run("rejects dash prefix", func(t *testing.T) {
		err := validatePrefix("--rh")
		if err == nil {
			t.Fatal("expected error for --rh prefix, got nil")
		}
		want := `design token prefix "--rh" should not start with dashes (use "rh" instead)`
		if err.Error() != want {
			t.Errorf("unexpected error message:\n got: %s\nwant: %s", err.Error(), want)
		}
	})

	t.Run("accepts valid prefix", func(t *testing.T) {
		if err := validatePrefix("rh"); err != nil {
			t.Errorf("unexpected error for valid prefix: %v", err)
		}
	})

	t.Run("accepts empty prefix", func(t *testing.T) {
		if err := validatePrefix(""); err != nil {
			t.Errorf("unexpected error for empty prefix: %v", err)
		}
	})
}

func TestLoadDesignTokens(t *testing.T) {
	fixtureDirs := []string{
		"draft-basic",
		"draft-aliases",
		"v2025-structured-colors",
	}

	for _, dir := range fixtureDirs {
		t.Run(dir, func(t *testing.T) {
			fixtureRoot, err := filepath.Abs(filepath.Join("testdata", dir))
			if err != nil {
				t.Fatalf("failed to resolve fixture path: %v", err)
			}

			ctx := &stubWorkspaceContext{
				root: fixtureRoot,
				config: &C.CemConfig{
					Generate: C.GenerateConfig{
						DesignTokens: C.DesignTokensConfig{
							Spec:   "./tokens.json",
							Prefix: "test",
						},
					},
				},
			}

			dt, err := LoadDesignTokens(ctx)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			// Build actual output in golden format from loaded tokens
			// and verify Get() returns each token
			actual := make(map[string]expectedToken)
			for _, tok := range dt.tokens.All() {
				name := tok.CSSVariableName()
				if _, ok := dt.Get(name); !ok {
					t.Errorf("Get(%q) returned false", name)
				}
				actual[name] = expectedToken{
					Value:       tok.DisplayValue(),
					Description: tok.Description,
					Syntax:      tok.CSSSyntax(),
				}
			}

			actualJSON, err := json.MarshalIndent(actual, "", "  ")
			if err != nil {
				t.Fatalf("failed to marshal actual tokens: %v", err)
			}

			testutil.CheckGolden(t, "expected", append(actualJSON, '\n'), testutil.GoldenOptions{
				Dir:         filepath.Join("testdata", dir),
				Extension:   ".json",
				UseJSONDiff: true,
			})
		})
	}
}

func TestLoadDesignTokensFetcherWiring(t *testing.T) {
	fixtureRoot, err := filepath.Abs(filepath.Join("testdata", "draft-basic"))
	if err != nil {
		t.Fatalf("failed to resolve fixture path: %v", err)
	}

	tests := []struct {
		name string
		spec string
	}{
		{"npm", "npm:@nonexistent-cem-test-pkg/tokens/tokens.json"},
		{"jsr", "jsr:@nonexistent-cem-test-pkg/tokens/tokens.json"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := &stubWorkspaceContext{
				root: fixtureRoot,
				config: &C.CemConfig{
					Generate: C.GenerateConfig{
						DesignTokens: C.DesignTokensConfig{
							Spec:   tt.spec,
							Prefix: "test",
						},
					},
				},
			}

			_, err := LoadDesignTokens(ctx)
			if err == nil {
				t.Fatalf("expected error for nonexistent %s package, got nil", tt.name)
			}

			// The error should mention esm.sh, indicating the fetcher was wired up
			// and attempted a CDN fallback after local resolution failed.
			errMsg := err.Error()
			if !strings.Contains(errMsg, "esm.sh") {
				t.Errorf("expected error to mention esm.sh (CDN fallback), got: %s", errMsg)
			}
		})
	}
}

func TestLoadDesignTokensRejectsDashPrefix(t *testing.T) {
	fixtureRoot, err := filepath.Abs(filepath.Join("testdata", "invalid-dash-prefix"))
	if err != nil {
		t.Fatalf("failed to resolve fixture path: %v", err)
	}

	ctx := &stubWorkspaceContext{
		root: fixtureRoot,
		config: &C.CemConfig{
			Generate: C.GenerateConfig{
				DesignTokens: C.DesignTokensConfig{
					Spec:   "./tokens.json",
					Prefix: "--rh",
				},
			},
		},
	}

	_, err = LoadDesignTokens(ctx)
	if err == nil {
		t.Fatal("expected error for --rh prefix, got nil")
	}
}
