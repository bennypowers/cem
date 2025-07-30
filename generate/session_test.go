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
	"time"

	W "bennypowers.dev/cem/workspace"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setupTestContext creates a test workspace context
func setupTestContext(t *testing.T, fixture string) W.WorkspaceContext {
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
				// Test concurrent access to in-memory manifest
				done := make(chan bool, 2)

				// Goroutine 1: Generate manifest
				go func() {
					_, err := session.GenerateFullManifest(context.Background())
					assert.NoError(t, err)
					done <- true
				}()

				// Goroutine 2: Read in-memory manifest
				go func() {
					// Wait a bit, then try to read
					time.Sleep(10 * time.Millisecond)
					manifest := session.InMemoryManifest()
					// Should not panic or cause race condition
					_ = manifest
					done <- true
				}()

				// Wait for both to complete
				<-done
				<-done
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

func BenchmarkGenerateSession_SingleGeneration(b *testing.B) {
	ctx := W.NewFileSystemWorkspaceContext("test/fixtures/project-classes")
	require.NoError(b, ctx.Init())

	cfg, err := ctx.Config()
	require.NoError(b, err)
	cfg.Generate.Files = []string{"src/class-fields.ts"}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		session, err := NewGenerateSession(ctx)
		require.NoError(b, err)

		_, err = session.GenerateFullManifest(context.Background())
		require.NoError(b, err)

		session.Close()
	}
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

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err = session.GenerateFullManifest(context.Background())
		require.NoError(b, err)
	}
}
