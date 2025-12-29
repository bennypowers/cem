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

package workspace_test

import (
	"errors"
	"io"
	"strings"
	"testing"

	C "bennypowers.dev/cem/cmd/config"
	M "bennypowers.dev/cem/manifest"
	"bennypowers.dev/cem/types"
	W "bennypowers.dev/cem/workspace"
)

// mockDesignTokensLoader implements DesignTokensLoader interface for testing
type mockDesignTokensLoader struct {
	loadCount int
	result    types.DesignTokens
	err       error
}

func (m *mockDesignTokensLoader) Load(ctx types.WorkspaceContext) (types.DesignTokens, error) {
	m.loadCount++
	return m.result, m.err
}

// mockDesignTokens implements DesignTokens interface for testing
type mockDesignTokens struct {
	name string
}

func (m *mockDesignTokens) Get(name string) (types.TokenResult, bool) {
	return nil, false
}

// mockWorkspaceContext implements WorkspaceContext for testing
type mockWorkspaceContext struct {
	config    *C.CemConfig
	configErr error
}

func (m *mockWorkspaceContext) Init() error                                      { return nil }
func (m *mockWorkspaceContext) ConfigFile() string                               { return "" }
func (m *mockWorkspaceContext) Config() (*C.CemConfig, error)                    { return m.config, m.configErr }
func (m *mockWorkspaceContext) PackageJSON() (*M.PackageJSON, error)             { return nil, nil }
func (m *mockWorkspaceContext) Manifest() (*M.Package, error)                    { return nil, nil }
func (m *mockWorkspaceContext) CustomElementsManifestPath() string               { return "" }
func (m *mockWorkspaceContext) ReadFile(path string) (io.ReadCloser, error)      { return nil, nil }
func (m *mockWorkspaceContext) Glob(pattern string) ([]string, error)            { return nil, nil }
func (m *mockWorkspaceContext) OutputWriter(path string) (io.WriteCloser, error) { return nil, nil }
func (m *mockWorkspaceContext) Root() string                                     { return "" }
func (m *mockWorkspaceContext) Cleanup() error                                   { return nil }
func (m *mockWorkspaceContext) ModulePathToFS(modulePath string) string          { return "" }
func (m *mockWorkspaceContext) FSPathToModule(fsPath string) (string, error)     { return "", nil }
func (m *mockWorkspaceContext) ResolveModuleDependency(modulePath, dependencyPath string) (string, error) {
	return "", nil
}
func (m *mockWorkspaceContext) DesignTokensCache() types.DesignTokensCache { return nil }

func TestDesignTokensCache_WithLoader(t *testing.T) {
	mockTokens := &mockDesignTokens{name: "test-tokens"}
	loader := &mockDesignTokensLoader{result: mockTokens}
	cache := W.NewDesignTokensCache(loader)

	ctx := &mockWorkspaceContext{
		config: &C.CemConfig{
			Generate: C.GenerateConfig{
				DesignTokens: C.DesignTokensConfig{
					Spec: "test-spec",
				},
			},
		},
	}

	// First call should load tokens
	result1, err := cache.LoadOrReuse(ctx)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result1 == nil {
		t.Errorf("Expected design tokens object, got nil")
	}
	if loader.loadCount != 1 {
		t.Errorf("Expected load count 1, got %d", loader.loadCount)
	}

	// Second call with same spec should use cache
	result2, err := cache.LoadOrReuse(ctx)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result2 == nil {
		t.Errorf("Expected design tokens object, got nil")
	}
	if loader.loadCount != 1 {
		t.Errorf("Expected load count still 1 (cached), got %d", loader.loadCount)
	}

	// Change spec should trigger reload
	ctx.config.Generate.DesignTokens.Spec = "new-spec"
	result3, err := cache.LoadOrReuse(ctx)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result3 == nil {
		t.Errorf("Expected design tokens object, got nil")
	}
	if loader.loadCount != 2 {
		t.Errorf("Expected load count 2 (reloaded), got %d", loader.loadCount)
	}
}

func TestDesignTokensCache_WithNilLoader(t *testing.T) {
	cache := W.NewDesignTokensCache(nil)

	ctx := &mockWorkspaceContext{
		config: &C.CemConfig{
			Generate: C.GenerateConfig{
				DesignTokens: C.DesignTokensConfig{
					Spec: "test-spec",
				},
			},
		},
	}

	// Should return nil when no loader provided
	result, err := cache.LoadOrReuse(ctx)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result != nil {
		t.Errorf("Expected nil result, got %v", result)
	}
}

func TestDesignTokensCache_LoaderError(t *testing.T) {
	testError := errors.New("test error")
	loader := &mockDesignTokensLoader{err: testError}
	cache := W.NewDesignTokensCache(loader)

	ctx := &mockWorkspaceContext{
		config: &C.CemConfig{
			Generate: C.GenerateConfig{
				DesignTokens: C.DesignTokensConfig{
					Spec: "test-spec",
				},
			},
		},
	}

	// Should return error from loader
	result, err := cache.LoadOrReuse(ctx)
	if err != testError {
		t.Fatalf("Expected test error, got %v", err)
	}
	if result != nil {
		t.Errorf("Expected nil result on error, got %v", result)
	}

	// Second call should return cached error
	result2, err2 := cache.LoadOrReuse(ctx)
	if err2 != testError {
		t.Fatalf("Expected cached test error, got %v", err2)
	}
	if result2 != nil {
		t.Errorf("Expected nil result on cached error, got %v", result2)
	}
	if loader.loadCount != 1 {
		t.Errorf("Expected load count 1 (error cached), got %d", loader.loadCount)
	}
}

func TestDesignTokensCache_Clear(t *testing.T) {
	loader := &mockDesignTokensLoader{result: &mockDesignTokens{name: "test-tokens"}}
	cache := W.NewDesignTokensCache(loader)

	ctx := &mockWorkspaceContext{
		config: &C.CemConfig{
			Generate: C.GenerateConfig{
				DesignTokens: C.DesignTokensConfig{
					Spec: "test-spec",
				},
			},
		},
	}

	// Load tokens
	_, err := cache.LoadOrReuse(ctx)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if loader.loadCount != 1 {
		t.Errorf("Expected load count 1, got %d", loader.loadCount)
	}

	// Clear cache
	cache.Clear()

	// Next call should reload
	_, err = cache.LoadOrReuse(ctx)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if loader.loadCount != 2 {
		t.Errorf("Expected load count 2 (reloaded after clear), got %d", loader.loadCount)
	}
}

func TestDesignTokensCache_ConfigError(t *testing.T) {
	loader := &mockDesignTokensLoader{result: &mockDesignTokens{name: "test-tokens"}}
	cache := W.NewDesignTokensCache(loader)

	// mockWorkspaceContext that returns config error
	ctx := &mockWorkspaceContext{
		config:    nil,
		configErr: errors.New("config error"),
	}

	// Should return config error
	result, err := cache.LoadOrReuse(ctx)
	if err == nil || !strings.Contains(err.Error(), "config error") {
		t.Fatalf("Expected config error, got %v", err)
	}
	if result != nil {
		t.Errorf("Expected nil result on config error, got %v", result)
	}
	if loader.loadCount != 0 {
		t.Errorf("Expected load count 0 (config failed), got %d", loader.loadCount)
	}
}
