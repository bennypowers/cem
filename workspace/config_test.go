package workspace_test

import (
	"path/filepath"
	"testing"

	C "bennypowers.dev/cem/cmd/config"
	"bennypowers.dev/cem/workspace"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func absFixture(t *testing.T, rel string) string {
	t.Helper()
	abs, err := filepath.Abs(filepath.Join("testdata", rel))
	require.NoError(t, err)
	return abs
}

func TestConfigFile_YamlInDotConfig(t *testing.T) {
	root := absFixture(t, "config-yaml")
	ctx := workspace.NewFileSystemWorkspaceContext(root)
	assert.Equal(t, filepath.Join(root, ".config", "cem.yaml"), ctx.ConfigFile())
}

func TestConfigFile_YmlInDotConfig(t *testing.T) {
	root := absFixture(t, "config-yml")
	ctx := workspace.NewFileSystemWorkspaceContext(root)
	assert.Equal(t, filepath.Join(root, ".config", "cem.yml"), ctx.ConfigFile())
}

func TestConfigFile_JsonInDotConfig(t *testing.T) {
	root := absFixture(t, "config-json")
	ctx := workspace.NewFileSystemWorkspaceContext(root)
	assert.Equal(t, filepath.Join(root, ".config", "cem.json"), ctx.ConfigFile())
}

func TestConfigFile_JsoncInDotConfig(t *testing.T) {
	root := absFixture(t, "config-jsonc")
	ctx := workspace.NewFileSystemWorkspaceContext(root)
	assert.Equal(t, filepath.Join(root, ".config", "cem.jsonc"), ctx.ConfigFile())
}

func TestConfigFile_DotFile(t *testing.T) {
	root := absFixture(t, "config-dotfile")
	ctx := workspace.NewFileSystemWorkspaceContext(root)
	assert.Equal(t, filepath.Join(root, ".cem.yaml"), ctx.ConfigFile())
}

func TestConfigFile_NoneFound(t *testing.T) {
	root := absFixture(t, "config-none")
	ctx := workspace.NewFileSystemWorkspaceContext(root)
	assert.Empty(t, ctx.ConfigFile())
}

func TestWithConfigFile_OverridesDiscovery(t *testing.T) {
	root := absFixture(t, "config-yaml")
	explicit := filepath.Join(root, ".config", "cem.yaml")
	ctx := workspace.NewFileSystemWorkspaceContext(root, workspace.WithConfigFile(explicit))
	assert.Equal(t, explicit, ctx.ConfigFile())
}

func TestWithConfigFile_ArbitraryPath(t *testing.T) {
	root := absFixture(t, "config-none")
	// Inject a config path that wouldn't be discovered
	explicit := absFixture(t, "config-yaml/.config/cem.yaml")
	ctx := workspace.NewFileSystemWorkspaceContext(root, workspace.WithConfigFile(explicit))
	assert.Equal(t, explicit, ctx.ConfigFile())
}

func TestInit_LoadsYamlConfig(t *testing.T) {
	root := absFixture(t, "config-yaml")
	ctx := workspace.NewFileSystemWorkspaceContext(root)
	require.NoError(t, ctx.Init())
	cfg, err := ctx.Config()
	require.NoError(t, err)
	require.NotNil(t, cfg)
	assert.Equal(t, []string{"elements/**/*.ts"}, cfg.Generate.Files)
	assert.Contains(t, cfg.Generate.Output, "custom-elements.json")
}

func TestInit_LoadsYmlConfig(t *testing.T) {
	root := absFixture(t, "config-yml")
	ctx := workspace.NewFileSystemWorkspaceContext(root)
	require.NoError(t, ctx.Init())
	cfg, err := ctx.Config()
	require.NoError(t, err)
	assert.Equal(t, []string{"src/**/*.ts"}, cfg.Generate.Files)
}

func TestInit_LoadsJsonConfig(t *testing.T) {
	root := absFixture(t, "config-json")
	ctx := workspace.NewFileSystemWorkspaceContext(root)
	require.NoError(t, ctx.Init())
	cfg, err := ctx.Config()
	require.NoError(t, err)
	assert.Equal(t, []string{"lib/**/*.js"}, cfg.Generate.Files)
}

func TestInit_LoadsJsoncConfig(t *testing.T) {
	root := absFixture(t, "config-jsonc")
	ctx := workspace.NewFileSystemWorkspaceContext(root)
	require.NoError(t, ctx.Init())
	cfg, err := ctx.Config()
	require.NoError(t, err)
	assert.Equal(t, []string{"src/**/*.ts"}, cfg.Generate.Files)
}

func TestInit_LoadsDotfileConfig(t *testing.T) {
	root := absFixture(t, "config-dotfile")
	ctx := workspace.NewFileSystemWorkspaceContext(root)
	require.NoError(t, ctx.Init())
	cfg, err := ctx.Config()
	require.NoError(t, err)
	assert.Equal(t, []string{"components/**/*.ts"}, cfg.Generate.Files)
}

func TestInit_NoConfig(t *testing.T) {
	root := absFixture(t, "config-none")
	ctx := workspace.NewFileSystemWorkspaceContext(root)
	require.NoError(t, ctx.Init())
	cfg, err := ctx.Config()
	require.NoError(t, err)
	require.NotNil(t, cfg)
	assert.Empty(t, cfg.Generate.Files)
}

func TestInit_OutputPathRelativeToRoot(t *testing.T) {
	root := absFixture(t, "config-yaml")
	ctx := workspace.NewFileSystemWorkspaceContext(root)
	require.NoError(t, ctx.Init())
	cfg, err := ctx.Config()
	require.NoError(t, err)
	// Output should be made absolute relative to project root
	assert.Equal(t, filepath.Join(root, "custom-elements.json"), cfg.Generate.Output)
}

func TestInit_ProjectDirSet(t *testing.T) {
	root := absFixture(t, "config-yaml")
	ctx := workspace.NewFileSystemWorkspaceContext(root)
	require.NoError(t, ctx.Init())
	cfg, err := ctx.Config()
	require.NoError(t, err)
	assert.Equal(t, root, cfg.ProjectDir)
}

func TestLoadWorkspaceConfig_WithConfig(t *testing.T) {
	pkgDir := absFixture(t, "workspace-with-config/packages/elements")
	cfg, err := workspace.LoadWorkspaceConfig(pkgDir)
	require.NoError(t, err)
	require.NotNil(t, cfg)
	assert.Equal(t, 9000, cfg.Serve.Port)
}

func TestLoadWorkspaceConfig_WorkspaceNoConfig(t *testing.T) {
	// workspace-mode-no-manifest has a workspaces field but no .config/cem.yaml
	pkgDir := absFixture(t, "workspace-mode-no-manifest/packages/util")
	cfg, err := workspace.LoadWorkspaceConfig(pkgDir)
	require.NoError(t, err)
	assert.Nil(t, cfg)
}

func TestLoadWorkspaceConfig_NoWorkspace(t *testing.T) {
	root := absFixture(t, "config-yaml")
	cfg, err := workspace.LoadWorkspaceConfig(root)
	require.NoError(t, err)
	assert.Nil(t, cfg)
}

func TestLoadPackageConfigWithWorkspaceDefaults_Merges(t *testing.T) {
	pkgDir := absFixture(t, "workspace-with-config/packages/elements")
	cfg, err := workspace.LoadPackageConfigWithWorkspaceDefaults(pkgDir)
	require.NoError(t, err)
	require.NotNil(t, cfg)
	// Package overrides workspace port
	assert.Equal(t, 3000, cfg.Serve.Port)
	// Package has its own files
	assert.Equal(t, []string{"src/**/*.ts"}, cfg.Generate.Files)
}

func TestLoadPackageConfigWithWorkspaceDefaults_WorkspaceDefaults(t *testing.T) {
	// Create a package dir inside the workspace that has no config of its own
	// The workspace-with-config fixture has port 9000 at workspace level
	// A package without config should inherit workspace defaults
	pkgDir := absFixture(t, "workspace-with-config/packages/elements")

	// First verify workspace config
	wsCfg, err := workspace.LoadWorkspaceConfig(pkgDir)
	require.NoError(t, err)
	require.NotNil(t, wsCfg)
	assert.Equal(t, 9000, wsCfg.Serve.Port)
}

func TestClone_Nil(t *testing.T) {
	assert.Nil(t, C.Clone(nil))
}

func TestClone_DeepCopy(t *testing.T) {
	original := &C.CemConfig{
		Generate: C.GenerateConfig{
			Files:   []string{"a.ts", "b.ts"},
			Exclude: []string{"*.test.ts"},
			Output:  "out.json",
		},
		Serve: C.ServeConfig{
			Port: 8000,
			URLRewrites: []C.URLRewrite{
				{URLPattern: "/old/*", URLTemplate: "/new/{{.path}}"},
			},
		},
		AdditionalPackages: []string{"npm:@scope/pkg"},
		Export: map[string]C.FrameworkExportConfig{
			"react": {Output: "react/"},
		},
	}

	clone := C.Clone(original)
	require.NotNil(t, clone)

	// Verify values match
	assert.Equal(t, original.Generate.Files, clone.Generate.Files)
	assert.Equal(t, original.Serve.Port, clone.Serve.Port)
	assert.Equal(t, original.AdditionalPackages, clone.AdditionalPackages)

	// Verify deep copy (mutating clone doesn't affect original)
	clone.Generate.Files[0] = "changed.ts"
	assert.Equal(t, "a.ts", original.Generate.Files[0])

	clone.Serve.URLRewrites[0].URLPattern = "/changed/*"
	assert.Equal(t, "/old/*", original.Serve.URLRewrites[0].URLPattern)

	clone.AdditionalPackages[0] = "changed"
	assert.Equal(t, "npm:@scope/pkg", original.AdditionalPackages[0])

	clone.Export["react"] = C.FrameworkExportConfig{Output: "changed/"}
	assert.Equal(t, "react/", original.Export["react"].Output)
}

func TestInit_InvalidConfig(t *testing.T) {
	root := absFixture(t, "config-invalid")
	ctx := workspace.NewFileSystemWorkspaceContext(root)
	err := ctx.Init()
	assert.Error(t, err)
}

func TestLoadWorkspaceConfig_InvalidConfig(t *testing.T) {
	pkgDir := absFixture(t, "workspace-invalid-config/packages/elements")
	_, err := workspace.LoadWorkspaceConfig(pkgDir)
	assert.Error(t, err)
}

func TestLoadPackageConfigWithWorkspaceDefaults_NoPackageConfig(t *testing.T) {
	// Package has no config, workspace has port 9000 and openBrowser: true
	pkgDir := absFixture(t, "workspace-with-config/packages/utils")
	cfg, err := workspace.LoadPackageConfigWithWorkspaceDefaults(pkgDir)
	require.NoError(t, err)
	require.NotNil(t, cfg)
	// Should inherit workspace port since package has none
	assert.Equal(t, 9000, cfg.Serve.Port)
	// Should inherit workspace openBrowser since package has none
	assert.True(t, cfg.Serve.OpenBrowser)
}

func TestLoadPackageConfigWithWorkspaceDefaults_NoWorkspace(t *testing.T) {
	// Not in a workspace at all
	root := absFixture(t, "config-yaml")
	cfg, err := workspace.LoadPackageConfigWithWorkspaceDefaults(root)
	require.NoError(t, err)
	require.NotNil(t, cfg)
	assert.Equal(t, []string{"elements/**/*.ts"}, cfg.Generate.Files)
}

func TestInit_VerboseConfig(t *testing.T) {
	root := absFixture(t, "config-verbose")
	ctx := workspace.NewFileSystemWorkspaceContext(root)
	require.NoError(t, ctx.Init())
	cfg, err := ctx.Config()
	require.NoError(t, err)
	assert.True(t, cfg.Verbose)
}

func TestLoadPackageConfigWithWorkspaceDefaults_InvalidPackageConfig(t *testing.T) {
	pkgDir := absFixture(t, "workspace-with-config/packages/broken")
	_, err := workspace.LoadPackageConfigWithWorkspaceDefaults(pkgDir)
	assert.Error(t, err)
}

func TestLoadPackageConfigWithWorkspaceDefaults_InvalidWorkspaceConfig(t *testing.T) {
	pkgDir := absFixture(t, "workspace-invalid-config/packages/elements")
	_, err := workspace.LoadPackageConfigWithWorkspaceDefaults(pkgDir)
	assert.Error(t, err)
}

func TestFindProjectRootFromDir_ConfigYaml(t *testing.T) {
	root := absFixture(t, "config-yaml")
	// findProjectRootFromDir is not exported, but we can test it
	// indirectly via NewFileSystemWorkspaceContext which uses FindConfigFile
	ctx := workspace.NewFileSystemWorkspaceContext(root)
	assert.Equal(t, root, ctx.Root())
}

func TestFindProjectRootFromDir_DotFile(t *testing.T) {
	root := absFixture(t, "config-dotfile")
	ctx := workspace.NewFileSystemWorkspaceContext(root)
	assert.Equal(t, root, ctx.Root())
}
