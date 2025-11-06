/*
Copyright © 2025 Benny Powers <web@bennypowers.com>
*/
package lsp_test

import (
	"testing"

	"bennypowers.dev/cem/lsp"
	"bennypowers.dev/cem/workspace"
	"github.com/stretchr/testify/require"
)

// TestDebugElementDefinitions checks element definitions and package names
func TestDebugElementDefinitions(t *testing.T) {
	pfePath := "/home/bennyp/Developer/patternfly/patternfly-elements"

	wsCtx := workspace.NewFileSystemWorkspaceContext(pfePath)
	registry, err := lsp.NewRegistryWithDefaults()
	require.NoError(t, err, "Failed to create registry")

	err = registry.LoadFromWorkspace(wsCtx)
	require.NoError(t, err, "Failed to load from workspace")

	// Check pf-label definition
	if def, exists := registry.ElementDefinition("pf-label"); exists {
		t.Logf("pf-label definition:")
		t.Logf("  Package name: '%s'", def.PackageName())
		t.Logf("  Module path: '%s'", def.ModulePath())
		t.Logf("  Source href: '%s'", def.SourceHref())
	} else {
		t.Logf("pf-label definition NOT found")
	}

	// Check pf-button definition
	if def, exists := registry.ElementDefinition("pf-button"); exists {
		t.Logf("pf-button definition:")
		t.Logf("  Package name: '%s'", def.PackageName())
		t.Logf("  Module path: '%s'", def.ModulePath())
		t.Logf("  Source href: '%s'", def.SourceHref())
	} else {
		t.Logf("pf-button definition NOT found")
	}

	// Check pf-label-group definition
	if def, exists := registry.ElementDefinition("pf-label-group"); exists {
		t.Logf("pf-label-group definition:")
		t.Logf("  Package name: '%s'", def.PackageName())
		t.Logf("  Module path: '%s'", def.ModulePath())
		t.Logf("  Source href: '%s'", def.SourceHref())
	} else {
		t.Logf("pf-label-group definition NOT found")
	}
}
