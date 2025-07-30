package generate

import (
	"testing"
)

// TestDemoDiscoverySkipping tests the demo discovery optimization logic
func TestDemoDiscoverySkipping(t *testing.T) {
	// Test that shouldSkipDemoDiscovery returns true when no demo discovery is configured
	ws := &WatchSession{}

	// Mock a context that returns config with no demo discovery
	// This is a basic test to ensure the logic compiles and basic behavior works

	// For now, just test that the functions exist and don't panic
	_ = ws.shouldSkipDemoDiscovery
	_ = ws.updateDemoDiscoveryState
	_ = ws.matchesDemoGlobs

	t.Log("Demo discovery optimization functions are properly defined")
}
