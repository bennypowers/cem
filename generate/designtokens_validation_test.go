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
package generate

import (
	"errors"
	"io"
	"strings"
	"testing"

	C "bennypowers.dev/cem/cmd/config"
	M "bennypowers.dev/cem/manifest"
	"bennypowers.dev/cem/types"
)

// Mock design tokens cache for testing validation scenarios
type mockDesignTokensCache struct {
	returnValue types.DesignTokens
	returnError error
}

func (m *mockDesignTokensCache) LoadOrReuse(ctx types.WorkspaceContext) (types.DesignTokens, error) {
	return m.returnValue, m.returnError
}

func (m *mockDesignTokensCache) Clear() {
	// No-op for testing
}

// Mock design tokens for testing
type mockDesignTokens struct {
	name string
}

func (m *mockDesignTokens) Get(name string) (types.TokenResult, bool) {
	return nil, false
}

// Mock workspace context for testing
type mockWorkspaceContext struct {
	cache types.DesignTokensCache
}

func (m *mockWorkspaceContext) DesignTokensCache() types.DesignTokensCache {
	return m.cache
}

// Implement other required methods as no-ops for testing
func (m *mockWorkspaceContext) Root() string                                { return "/test" }
func (m *mockWorkspaceContext) ConfigFile() string                          { return "" }
func (m *mockWorkspaceContext) Config() (*C.CemConfig, error)               { return nil, nil }
func (m *mockWorkspaceContext) Init() error                                 { return nil }
func (m *mockWorkspaceContext) Glob(string) ([]string, error)               { return nil, nil }
func (m *mockWorkspaceContext) Cleanup() error                              { return nil }
func (m *mockWorkspaceContext) PackageJSON() (*M.PackageJSON, error)        { return nil, nil }
func (m *mockWorkspaceContext) Manifest() (*M.Package, error)               { return nil, nil }
func (m *mockWorkspaceContext) CustomElementsManifestPath() string          { return "" }
func (m *mockWorkspaceContext) ReadFile(string) (io.ReadCloser, error)      { return nil, nil }
func (m *mockWorkspaceContext) OutputWriter(string) (io.WriteCloser, error) { return nil, nil }
func (m *mockWorkspaceContext) ModulePathToFS(string) string                { return "" }
func (m *mockWorkspaceContext) FSPathToModule(string) (string, error)       { return "", nil }
func (m *mockWorkspaceContext) ResolveModuleDependency(string, string) (string, error) {
	return "", nil
}

func TestValidateAndLoadDesignTokens(t *testing.T) {
	tests := []struct {
		name          string
		cacheReturn   types.DesignTokens
		cacheError    error
		expectError   bool
		errorContains string
	}{
		{
			name:        "successful_validation",
			cacheReturn: &mockDesignTokens{name: "test"},
			cacheError:  nil,
			expectError: false,
		},
		{
			name:          "cache_load_error",
			cacheReturn:   nil,
			cacheError:    errors.New("network error"),
			expectError:   true,
			errorContains: "failed to load design tokens",
		},
		{
			name:          "nil_return_from_cache",
			cacheReturn:   nil,
			cacheError:    nil,
			expectError:   true,
			errorContains: "design tokens cache returned nil",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCache := &mockDesignTokensCache{
				returnValue: tt.cacheReturn,
				returnError: tt.cacheError,
			}

			mockCtx := &mockWorkspaceContext{
				cache: mockCache,
			}

			result, err := validateAndLoadDesignTokens(mockCtx)

			if tt.expectError {
				if err == nil {
					t.Error("Expected error but got none")
					return
				}

				if tt.errorContains != "" && !strings.Contains(err.Error(), tt.errorContains) {
					t.Errorf("Expected error to contain %q, but got: %v", tt.errorContains, err)
				}

				if result != nil {
					t.Error("Expected nil result when error occurs")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
					return
				}

				if result == nil {
					t.Error("Expected non-nil result when no error occurs")
				}
			}
		})
	}
}

func TestValidateAndLoadDesignTokens_Success(t *testing.T) {
	// Test that valid design tokens are returned successfully
	mockTokens := &mockDesignTokens{name: "test"}
	mockCache := &mockDesignTokensCache{
		returnValue: mockTokens,
		returnError: nil,
	}

	mockCtx := &mockWorkspaceContext{
		cache: mockCache,
	}

	result, err := validateAndLoadDesignTokens(mockCtx)

	if err != nil {
		t.Fatalf("Expected no error but got: %v", err)
	}

	if result == nil {
		t.Error("Expected design tokens object but got nil")
	}

	if result != mockTokens {
		t.Error("Expected same design tokens object that was cached")
	}
}

func TestValidateAndLoadDesignTokens_NilReturn(t *testing.T) {
	// Test when cache returns nil (valid case with no design tokens configured)
	mockCache := &mockDesignTokensCache{
		returnValue: nil,
		returnError: nil,
	}

	mockCtx := &mockWorkspaceContext{
		cache: mockCache,
	}

	result, err := validateAndLoadDesignTokens(mockCtx)

	if err == nil {
		t.Fatal("Expected error but got none")
	}

	if result != nil {
		t.Error("Expected nil result when cache returns nil")
	}

	// Verify the error message is helpful for debugging
	errorMsg := err.Error()
	if !strings.Contains(errorMsg, "design tokens cache returned nil") {
		t.Errorf("Expected error about nil cache return, got: %s", errorMsg)
	}
}

func TestValidateAndLoadDesignTokens_CacheErrorHandling(t *testing.T) {
	// Test that cache errors are properly wrapped and propagated
	originalError := errors.New("network timeout while fetching design tokens")
	mockCache := &mockDesignTokensCache{
		returnValue: nil,
		returnError: originalError,
	}

	mockCtx := &mockWorkspaceContext{
		cache: mockCache,
	}

	result, err := validateAndLoadDesignTokens(mockCtx)

	if err == nil {
		t.Fatal("Expected error but got none")
	}

	if result != nil {
		t.Error("Expected nil result when cache fails")
	}

	// Verify the original error is wrapped and preserved
	if !strings.Contains(err.Error(), "failed to load design tokens") {
		t.Error("Error should be wrapped with descriptive message")
	}

	if !strings.Contains(err.Error(), originalError.Error()) {
		t.Error("Original error should be preserved in wrapped error")
	}
}
