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
	"context"
	"testing"
	"testing/synctest"
	"time"

	M "bennypowers.dev/cem/manifest"
	"bennypowers.dev/cem/types"
	W "bennypowers.dev/cem/workspace"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setupTestContext creates a test workspace context
func setupTestContext(t *testing.T, fixture string) types.WorkspaceContext {
	ctx := W.NewFileSystemWorkspaceContext(fixture)
	require.NoError(t, ctx.Init())
	return ctx
}

func TestNewGenerateSession(t *testing.T) {
	tests := []struct {
		name    string
		fixture string
		wantErr bool
	}{
		{
			name:    "valid project with classes",
			fixture: "test/fixtures/project-classes",
			wantErr: false,
		},
		{
			name:    "valid project with CSS",
			fixture: "test/fixtures/project-css",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := setupTestContext(t, tt.fixture)

			session, err := NewGenerateSession(ctx)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, session)
			} else {
				require.NoError(t, err)
				require.NotNil(t, session)

				// Verify setup context is initialized
				assert.NotNil(t, session.setupCtx)
				assert.NotNil(t, session.setupCtx.QueryManager())
				assert.Equal(t, ctx, session.setupCtx.WorkspaceContext)

				// Test cleanup
				session.Close()
			}
		})
	}
}

func TestGenerateSession_GenerateFullManifest(t *testing.T) {
	tests := []struct {
		name          string
		fixture       string
		files         []string
		expectModules bool
		expectErr     bool
		cancelContext bool
	}{
		{
			name:          "generate from class fields",
			fixture:       "test/fixtures/project-classes",
			files:         []string{"src/class-fields.ts"},
			expectModules: true,
			expectErr:     false,
		},
		{
			name:          "generate from multiple files",
			fixture:       "test/fixtures/project-classes",
			files:         []string{"src/class-fields.ts", "src/class-methods.ts"},
			expectModules: true,
			expectErr:     false,
		},
		{
			name:          "generate with cancelled context",
			fixture:       "test/fixtures/project-classes",
			files:         []string{"src/class-fields.ts"},
			expectModules: false,
			expectErr:     true,
			cancelContext: true,
		},
		{
			name:          "generate with no files",
			fixture:       "test/fixtures/project-classes",
			files:         []string{},
			expectModules: false,
			expectErr:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := setupTestContext(t, tt.fixture)

			// Set up files for generation
			cfg, err := ctx.Config()
			require.NoError(t, err)
			cfg.Generate.Files = tt.files

			session, err := NewGenerateSession(ctx)
			require.NoError(t, err)
			defer session.Close()

			// Prepare context
			genCtx := context.Background()
			if tt.cancelContext {
				var cancel context.CancelFunc
				genCtx, cancel = context.WithCancel(context.Background())
				cancel() // Cancel immediately
			}

			// Test generation
			pkg, err := session.GenerateFullManifest(genCtx)

			if tt.expectErr {
				assert.Error(t, err)
				if tt.cancelContext {
					assert.Equal(t, context.Canceled, err)
				}
			} else {
				require.NoError(t, err)
				require.NotNil(t, pkg)

				// Verify manifest content
				assert.Equal(t, "2.1.0", pkg.SchemaVersion)

				if tt.expectModules {
					assert.NotEmpty(t, pkg.Modules)
				}

				// Test in-memory manifest is updated
				inMemory := session.InMemoryManifest()
				require.NotNil(t, inMemory)
				assert.Equal(t, pkg.SchemaVersion, inMemory.SchemaVersion)
				assert.Equal(t, len(pkg.Modules), len(inMemory.Modules))
			}
		})
	}
}

func TestGenerateSession_GetInMemoryManifest(t *testing.T) {
	tests := []struct {
		name            string
		fixture         string
		files           []string
		generateFirst   bool
		expectNil       bool
		testConcurrency bool
	}{
		{
			name:          "nil before generation",
			fixture:       "test/fixtures/project-classes",
			generateFirst: false,
			expectNil:     true,
		},
		{
			name:          "populated after generation",
			fixture:       "test/fixtures/project-classes",
			files:         []string{"src/class-fields.ts"},
			generateFirst: true,
			expectNil:     false,
		},
		{
			name:            "concurrent access safety",
			fixture:         "test/fixtures/project-classes",
			files:           []string{"src/class-fields.ts"},
			generateFirst:   false,
			testConcurrency: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := setupTestContext(t, tt.fixture)

			if len(tt.files) > 0 {
				cfg, err := ctx.Config()
				require.NoError(t, err)
				cfg.Generate.Files = tt.files
			}

			session, err := NewGenerateSession(ctx)
			require.NoError(t, err)
			defer session.Close()

			if tt.generateFirst {
				_, err := session.GenerateFullManifest(context.Background())
				require.NoError(t, err)
			}

			if tt.testConcurrency {
				// Test concurrent access to in-memory manifest using virtual time
				synctest.Test(t, func(t *testing.T) {
					done := make(chan bool, 2)

					// Goroutine 1: Generate manifest
					go func() {
						_, err := session.GenerateFullManifest(context.Background())
						assert.NoError(t, err)
						done <- true
					}()

					// Goroutine 2: Read in-memory manifest
					go func() {
						// Virtual time coordination - no real delay needed
						manifest := session.InMemoryManifest()
						// Should not panic or cause race condition
						_ = manifest
						done <- true
					}()

					// Wait for both to complete
					<-done
					<-done
				})
			} else {
				manifest := session.InMemoryManifest()

				if tt.expectNil {
					assert.Nil(t, manifest)
				} else {
					assert.NotNil(t, manifest)
				}
			}
		})
	}
}

func TestGenerateSession_QueryManagerReuse(t *testing.T) {
	ctx := setupTestContext(t, "test/fixtures/project-classes")

	cfg, err := ctx.Config()
	require.NoError(t, err)
	cfg.Generate.Files = []string{"src/class-fields.ts"}

	session, err := NewGenerateSession(ctx)
	require.NoError(t, err)
	defer session.Close()

	// First generation
	start := time.Now()
	pkg1, err := session.GenerateFullManifest(context.Background())
	duration1 := time.Since(start)
	require.NoError(t, err)

	// Second generation (should be faster due to QueryManager reuse)
	start = time.Now()
	pkg2, err := session.GenerateFullManifest(context.Background())
	duration2 := time.Since(start)
	require.NoError(t, err)

	// Both should produce equivalent results
	assert.Equal(t, pkg1.SchemaVersion, pkg2.SchemaVersion)
	assert.Equal(t, len(pkg1.Modules), len(pkg2.Modules))

	// Second generation should be significantly faster (though this could be flaky)
	// We'll just verify it's not dramatically slower
	ratio := float64(duration2) / float64(duration1)
	assert.Less(t, ratio, 2.0, "Second generation should not be more than 2x slower than first")
}

func TestInMemoryManifest_Performance(t *testing.T) {
	gs := &GenerateSession{}

	// Test nil manifest
	if result := gs.InMemoryManifest(); result != nil {
		t.Error("InMemoryManifest should return nil when no manifest is set")
	}

	// Set a test manifest
	original := &M.Package{
		SchemaVersion: "2.1.0",
		Modules: []M.Module{
			{
				Kind:    "javascript-module",
				Path:    "test/module.js",
				Summary: "Test module",
			},
		},
	}

	gs.mu.Lock()
	gs.inMemoryManifest = original
	gs.mu.Unlock()

	// Get a copy (shallow for performance)
	copy := gs.InMemoryManifest()
	if copy == nil {
		t.Fatal("InMemoryManifest should return a copy of the manifest")
	}

	// Verify it's a different instance (thread-safe at the package level)
	if copy == original {
		t.Error("InMemoryManifest should return a different instance")
	}

	// Verify content is preserved
	if copy.SchemaVersion != original.SchemaVersion {
		t.Errorf("Content mismatch after copy: got %s, want %s", copy.SchemaVersion, original.SchemaVersion)
	}
}

func TestInMemoryManifestDeep_ThreadSafety(t *testing.T) {
	gs := &GenerateSession{}

	// Test nil manifest
	if result := gs.InMemoryManifestDeep(); result != nil {
		t.Error("InMemoryManifestDeep should return nil when no manifest is set")
	}

	// Set a test manifest
	original := &M.Package{
		SchemaVersion: "2.1.0",
		Modules: []M.Module{
			{
				Kind:    "javascript-module",
				Path:    "test/module.js",
				Summary: "Test module",
			},
		},
	}

	gs.mu.Lock()
	gs.inMemoryManifest = original
	gs.mu.Unlock()

	// Get a deep copy
	copy := gs.InMemoryManifestDeep()
	if copy == nil {
		t.Fatal("InMemoryManifestDeep should return a copy of the manifest")
	}

	// Verify it's a different instance (thread-safe)
	if copy == original {
		t.Error("InMemoryManifestDeep should return a different instance for thread safety")
	}

	// Verify content is preserved
	if copy.SchemaVersion != original.SchemaVersion {
		t.Errorf("Content mismatch after deep copy: got %s, want %s", copy.SchemaVersion, original.SchemaVersion)
	}

	// Test that modifying the deep copy doesn't affect the original
	if len(copy.Modules) > 0 {
		copy.Modules[0].Summary = "Modified summary"
		if original.Modules[0].Summary == "Modified summary" {
			t.Error("Modifying deep copy should not affect original (not properly deep copied)")
		}
	}
}

func TestModuleIndex_Performance(t *testing.T) {
	gs := &GenerateSession{}

	// Create test modules
	modules := []M.Module{
		{Kind: "javascript-module", Path: "test/module1.js"},
		{Kind: "javascript-module", Path: "test/module2.js"},
		{Kind: "javascript-module", Path: "test/module3.js"},
	}

	pkg := M.NewPackage(modules)

	// Set manifest and build index
	gs.mu.Lock()
	gs.inMemoryManifest = &pkg
	gs.moduleIndex = make(map[string]*M.Module)
	gs.rebuildModuleIndex()
	gs.mu.Unlock()

	// Test O(1) lookup
	module := gs.ModuleByPath("test/module2.js")
	if module == nil {
		t.Fatal("ModuleByPath should find existing module")
	}

	if module.Path != "test/module2.js" {
		t.Errorf("Wrong module returned: got %s, want test/module2.js", module.Path)
	}

	// Test non-existent module
	module = gs.ModuleByPath("test/nonexistent.js")
	if module != nil {
		t.Error("ModuleByPath should return nil for non-existent module")
	}
}

func TestMergeModulesIntoManifest_WithIndex(t *testing.T) {
	gs := &GenerateSession{
		moduleIndex: make(map[string]*M.Module),
	}

	// Test with empty manifest
	updatedModules := []M.Module{
		{Kind: "javascript-module", Path: "test/new.js", Summary: "New module"},
	}

	gs.MergeModulesIntoManifest(updatedModules)

	if gs.inMemoryManifest == nil {
		t.Fatal("Manifest should be created")
	}

	if len(gs.inMemoryManifest.Modules) != 1 {
		t.Errorf("Expected 1 module, got %d", len(gs.inMemoryManifest.Modules))
	}

	// Verify index was built
	module := gs.ModuleByPath("test/new.js")
	if module == nil {
		t.Error("Module should be findable in index")
	}

	// Test updating existing module
	existingModules := []M.Module{
		{Kind: "javascript-module", Path: "test/existing.js", Summary: "Original"},
	}

	pkg := M.NewPackage(existingModules)
	gs.mu.Lock()
	gs.inMemoryManifest = &pkg
	gs.rebuildModuleIndex()
	gs.mu.Unlock()

	// Update the module
	updatedModules = []M.Module{
		{Kind: "javascript-module", Path: "test/existing.js", Summary: "Updated"},
	}

	gs.MergeModulesIntoManifest(updatedModules)

	// Verify update worked
	module = gs.ModuleByPath("test/existing.js")
	if module == nil {
		t.Fatal("Updated module should be findable")
	}

	if module.Summary != "Updated" {
		t.Errorf("Module summary should be updated: got %s, want Updated", module.Summary)
	}

	// Should still only have one module
	if len(gs.inMemoryManifest.Modules) != 1 {
		t.Errorf("Should still have 1 module after update, got %d", len(gs.inMemoryManifest.Modules))
	}
}

func BenchmarkGenerateSession_SingleGeneration(b *testing.B) {
	ctx := W.NewFileSystemWorkspaceContext("test/fixtures/project-classes")
	require.NoError(b, ctx.Init())

	cfg, err := ctx.Config()
	require.NoError(b, err)
	cfg.Generate.Files = []string{"src/class-fields.ts"}

	for b.Loop() {
		session, err := NewGenerateSession(ctx)
		require.NoError(b, err)

		_, err = session.GenerateFullManifest(context.Background())
		require.NoError(b, err)

		session.Close()
	}
}

func TestSetMaxWorkers(t *testing.T) {
	ctx := setupTestContext(t, "test/fixtures/project-classes")

	session, err := NewGenerateSession(ctx)
	require.NoError(t, err)
	defer session.Close()

	// Test setting max workers
	session.SetMaxWorkers(4)

	// Verify worker count is set correctly
	workerCount := session.WorkerCount()
	assert.Equal(t, 4, workerCount, "WorkerCount should return the configured value")
}

func TestSetMaxWorkers_BeforeProcessing(t *testing.T) {
	ctx := setupTestContext(t, "test/fixtures/project-classes")

	cfg, err := ctx.Config()
	require.NoError(t, err)
	cfg.Generate.Files = []string{"src/class-fields.ts"}

	session, err := NewGenerateSession(ctx)
	require.NoError(t, err)
	defer session.Close()

	// Set max workers immediately after creation, before any processing
	session.SetMaxWorkers(2)

	// Generate manifest - this should use the configured worker count
	pkg, err := session.GenerateFullManifest(context.Background())
	require.NoError(t, err)
	assert.NotNil(t, pkg)

	// Verify the worker count is still configured correctly
	assert.Equal(t, 2, session.WorkerCount())
}

func TestSetMaxWorkers_ConcurrentAccess(t *testing.T) {
	ctx := setupTestContext(t, "test/fixtures/project-classes")

	session, err := NewGenerateSession(ctx)
	require.NoError(t, err)
	defer session.Close()

	// Test concurrent access using virtual time
	synctest.Test(t, func(t *testing.T) {
		done := make(chan bool, 2)

		// Goroutine 1: Set max workers
		go func() {
			session.SetMaxWorkers(4)
			done <- true
		}()

		// Goroutine 2: Read worker count
		go func() {
			_ = session.WorkerCount()
			done <- true
		}()

		// Wait for both to complete - should not deadlock or panic
		<-done
		<-done
	})

	// Verify final state is consistent
	workerCount := session.WorkerCount()
	assert.Greater(t, workerCount, 0, "WorkerCount should be positive")
}

func TestSetMaxWorkers_ZeroValue(t *testing.T) {
	ctx := setupTestContext(t, "test/fixtures/project-classes")

	session, err := NewGenerateSession(ctx)
	require.NoError(t, err)
	defer session.Close()

	// Setting to 0 should use default (NumCPU)
	session.SetMaxWorkers(0)

	// WorkerCount should fall back to default behavior
	workerCount := session.WorkerCount()
	assert.Greater(t, workerCount, 0, "WorkerCount should fall back to default when set to 0")
}

func BenchmarkGenerateSession_ReusedSession(b *testing.B) {
	ctx := W.NewFileSystemWorkspaceContext("test/fixtures/project-classes")
	require.NoError(b, ctx.Init())

	cfg, err := ctx.Config()
	require.NoError(b, err)
	cfg.Generate.Files = []string{"src/class-fields.ts"}

	session, err := NewGenerateSession(ctx)
	require.NoError(b, err)
	defer session.Close()

	for b.Loop() {
		_, err = session.GenerateFullManifest(context.Background())
		require.NoError(b, err)
	}
}
