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

package serve

import (
	"runtime"
	"testing"
)

func TestGetResourceLimits(t *testing.T) {
	limits := getResourceLimits()

	// Verify limits are reasonable
	if limits.MaxWorkers < 1 {
		t.Errorf("MaxWorkers should be at least 1, got %d", limits.MaxWorkers)
	}

	if limits.QueueDepth < 1 {
		t.Errorf("QueueDepth should be at least 1, got %d", limits.QueueDepth)
	}

	if limits.MaxConnections < 1 {
		t.Errorf("MaxConnections should be at least 1, got %d", limits.MaxConnections)
	}

	if limits.MaxConcurrentStreams < 1 {
		t.Errorf("MaxConcurrentStreams should be at least 1, got %d", limits.MaxConcurrentStreams)
	}

	// Log limits for visibility
	inContainer := isRunningInContainer()
	t.Logf("Environment: container=%v, NumCPU=%d", inContainer, runtime.NumCPU())
	t.Logf("Resource limits: workers=%d, queue=%d, connections=%d, streams=%d",
		limits.MaxWorkers, limits.QueueDepth, limits.MaxConnections, limits.MaxConcurrentStreams)
}

func TestContainerDetection(t *testing.T) {
	inContainer := isRunningInContainer()
	t.Logf("Running in container: %v", inContainer)

	// This test just logs the detection result for manual verification
	// In real CI/container environments, this should detect correctly
}
