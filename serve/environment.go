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
	"os"
	"runtime"
	"strings"
)

// isRunningInContainer detects if we're running in a containerized environment
// This uses multiple heuristics to detect Docker, Podman, Kubernetes, etc.
func isRunningInContainer() bool {
	// Check for Docker-specific file
	if _, err := os.Stat("/.dockerenv"); err == nil {
		return true
	}

	// Check for container-specific environment variables
	containerEnvVars := []string{
		"KUBERNETES_SERVICE_HOST", // Kubernetes
		"container",               // Podman/systemd
		"DOCKER_CONTAINER",        // Some Docker setups
	}
	for _, envVar := range containerEnvVars {
		if os.Getenv(envVar) != "" {
			return true
		}
	}

	// Check cgroup for container indicators
	if data, err := os.ReadFile("/proc/1/cgroup"); err == nil {
		content := string(data)
		containerIndicators := []string{
			"/docker/",
			"/kubepods/",
			"/podman/",
			"/containerd/",
		}
		for _, indicator := range containerIndicators {
			if strings.Contains(content, indicator) {
				return true
			}
		}
	}

	return false
}

// getResourceLimits returns appropriate concurrency limits based on the environment
type ResourceLimits struct {
	MaxWorkers           int // Transform pool workers
	QueueDepth           int // Transform pool queue depth
	MaxConnections       int // HTTP connection limit
	MaxConcurrentStreams uint32 // HTTP/2 stream limit per connection
}

// getResourceLimits calculates appropriate limits based on environment detection
func getResourceLimits() ResourceLimits {
	inContainer := isRunningInContainer()
	numCPU := runtime.NumCPU()

	if inContainer {
		// Conservative limits for containers (strict pthread limits)
		// Even if container sees many CPUs, thread limits are often very low
		maxWorkers := max(numCPU/4, 1)
		if maxWorkers > 4 {
			maxWorkers = 4
		}

		return ResourceLimits{
			MaxWorkers:           maxWorkers,
			QueueDepth:           maxWorkers * 24, // Generous queue for burst traffic
			MaxConnections:       25,              // Conservative connection limit
			MaxConcurrentStreams: 50,              // Conservative HTTP/2 stream limit
		}
	}

	// Native/localhost - optimize for performance
	maxWorkers := max(numCPU/2, 2)
	if maxWorkers > 8 {
		maxWorkers = 8
	}

	return ResourceLimits{
		MaxWorkers:           maxWorkers,
		QueueDepth:           maxWorkers * 12, // Balanced queue
		MaxConnections:       100,             // Generous for development
		MaxConcurrentStreams: 250,             // Standard HTTP/2 limit
	}
}
