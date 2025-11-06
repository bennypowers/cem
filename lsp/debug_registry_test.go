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

// TestDebugRegistry helps debug what's actually being loaded
func TestDebugRegistry(t *testing.T) {
	pfePath := "/home/bennyp/Developer/patternfly/patternfly-elements"

	wsCtx := workspace.NewFileSystemWorkspaceContext(pfePath)
	registry, err := lsp.NewRegistryWithDefaults()
	require.NoError(t, err, "Failed to create registry")

	err = registry.LoadFromWorkspace(wsCtx)
	require.NoError(t, err, "Failed to load from workspace")

	// Check what was loaded
	t.Logf("Total elements loaded: %d", len(registry.AllTagNames()))

	// Check for specific elements
	if element, exists := registry.Element("pf-label"); exists {
		t.Logf("✓ pf-label found with %d attributes", len(element.Attributes))
	} else {
		t.Logf("✗ pf-label NOT found")
	}

	if element, exists := registry.Element("pf-button"); exists {
		t.Logf("✓ pf-button found with %d attributes", len(element.Attributes))
	} else {
		t.Logf("✗ pf-button NOT found")
	}

	if element, exists := registry.Element("pf-label-group"); exists {
		t.Logf("✓ pf-label-group found with %d attributes", len(element.Attributes))
	} else {
		t.Logf("✗ pf-label-group NOT found")
	}

	// List first 20 elements
	allTags := registry.AllTagNames()
	t.Logf("First 20 elements:")
	for i := 0; i < 20 && i < len(allTags); i++ {
		t.Logf("  - %s", allTags[i])
	}
}
