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
package version

import (
	"fmt"
	"runtime/debug"
	"strings"
)

var (
	// Version information, set at build time via ldflags
	Version   = "dev"     // Version string (e.g., "v0.3.0")
	GitCommit = "unknown" // Git commit hash
	GitTag    = "unknown" // Git tag
	BuildTime = "unknown" // Build timestamp
	GitDirty  = ""        // "dirty" if working directory has uncommitted changes
)

// GetVersion returns the version string for the application
func GetVersion() string {
	// If Version was set via ldflags, use it
	if Version != "dev" {
		return Version
	}

	// Fallback: try to get version from build info
	if info, ok := debug.ReadBuildInfo(); ok {
		if info.Main.Version != "(devel)" && info.Main.Version != "" {
			return info.Main.Version
		}
	}

	// Final fallback: construct from git information if available
	if GitTag != "unknown" && GitCommit != "unknown" {
		version := GitTag
		if GitCommit != "" {
			// Safe substring that handles short commits
			commitSuffix := GitCommit
			if len(GitCommit) > 7 {
				commitSuffix = GitCommit[:7]
			}
			if !strings.HasSuffix(GitTag, commitSuffix) {
				// Add commit hash if tag doesn't contain it
				version = fmt.Sprintf("%s-%s", GitTag, commitSuffix)
			}
		}
		if GitDirty == "dirty" {
			version += "-dirty"
		}
		return version
	}

	return "dev"
}

// GetFullVersion returns detailed version information
func GetFullVersion() string {
	version := GetVersion()
	if GitCommit != "unknown" {
		return fmt.Sprintf("%s (commit: %s)", version, GitCommit)
	}
	return version
}

// GetBuildInfo returns detailed build information
func GetBuildInfo() map[string]string {
	return map[string]string{
		"version":   GetVersion(),
		"gitCommit": GitCommit,
		"gitTag":    GitTag,
		"buildTime": BuildTime,
		"gitDirty":  GitDirty,
	}
}
