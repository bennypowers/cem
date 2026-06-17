package mcp_test

import (
	"os"
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/internal/workspace"
	"bennypowers.dev/cem/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Inline assertions justified: testing cache invalidation behavior,
// values are single scalars compared before/after mutation.

func TestWorkspace_InvalidateConfig(t *testing.T) {
	dir := t.TempDir()
	cfgPath := filepath.Join(dir, ".config", "cem.yaml")
	require.NoError(t, os.MkdirAll(filepath.Dir(cfgPath), 0o755))
	require.NoError(t, os.WriteFile(cfgPath, []byte("serve:\n  port: 5000\n"), 0o644))
	require.NoError(t, os.WriteFile(filepath.Join(dir, "package.json"), []byte(`{"name":"test"}`), 0o644))

	ws := workspace.NewFileSystemWorkspaceContext(dir)
	require.NoError(t, ws.Init())

	cfg, err := ws.Config()
	require.NoError(t, err)
	assert.Equal(t, 5000, cfg.Serve.Port)

	require.NoError(t, os.WriteFile(cfgPath, []byte("serve:\n  port: 6000\n"), 0o644))

	ws.InvalidateConfig()

	cfg2, err := ws.Config()
	require.NoError(t, err)
	assert.Equal(t, 6000, cfg2.Serve.Port)
}

func TestWorkspace_ConcurrentConfigAccess(t *testing.T) {
	dir := t.TempDir()
	cfgPath := filepath.Join(dir, ".config", "cem.yaml")
	require.NoError(t, os.MkdirAll(filepath.Dir(cfgPath), 0o755))
	require.NoError(t, os.WriteFile(cfgPath, []byte("serve:\n  port: 9000\n"), 0o644))
	require.NoError(t, os.WriteFile(filepath.Join(dir, "package.json"), []byte(`{"name":"test"}`), 0o644))

	ws := workspace.NewFileSystemWorkspaceContext(dir)
	require.NoError(t, ws.Init())

	done := make(chan struct{})
	go func() {
		defer close(done)
		for range 100 {
			ws.InvalidateConfig()
		}
	}()

	for range 100 {
		_, _ = ws.Config()
	}
	<-done
}

func TestMCPContext_InvalidateConfig(t *testing.T) {
	dir := t.TempDir()
	cfgPath := filepath.Join(dir, ".config", "cem.yaml")
	require.NoError(t, os.MkdirAll(filepath.Dir(cfgPath), 0o755))
	require.NoError(t, os.WriteFile(cfgPath, []byte("serve:\n  port: 7000\n"), 0o644))
	require.NoError(t, os.WriteFile(filepath.Join(dir, "package.json"), []byte(`{"name":"test"}`), 0o644))

	ws := workspace.NewFileSystemWorkspaceContext(dir)
	require.NoError(t, ws.Init())

	ctx, err := mcp.NewMCPContext(ws)
	require.NoError(t, err)
	require.NoError(t, ctx.LoadManifests())

	cfg, err := ctx.Config()
	require.NoError(t, err)
	assert.Equal(t, 7000, cfg.Serve.Port)

	require.NoError(t, os.WriteFile(cfgPath, []byte("serve:\n  port: 8000\n"), 0o644))

	ctx.InvalidateConfig()

	cfg2, err := ctx.Config()
	require.NoError(t, err)
	assert.Equal(t, 8000, cfg2.Serve.Port)
}
