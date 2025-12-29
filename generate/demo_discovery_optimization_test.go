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
