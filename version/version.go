/*
Copyright © 2026 Benny Powers <web@bennypowers.com>

SPDX-License-Identifier: GPL-3.0-or-later

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

// Package version reports the version and build metadata of the cem binary
// and of programs that embed cem packages.
//
// It is intended for external consumption. The package is part of the
// supported public Go API, so it follows the compatibility policy described
// in CONTRIBUTING.md.
//
// Release builds inject metadata with -ldflags -X, for example:
//
//	go build -ldflags "\
//	  -X bennypowers.dev/cem/version.Version=v1.2.3 \
//	  -X bennypowers.dev/cem/version.GitCommit=abcdef1 \
//	  -X bennypowers.dev/cem/version.GitTag=v1.2.3" .
//
// When the metadata is absent, Get falls back to the version recorded by the
// Go toolchain in the binary's build information. If that is unavailable too,
// Get returns "dev".
//
// Usage:
//
//	fmt.Println("cem", version.Get())
//	for key, value := range version.Info() {
//		fmt.Printf("%s: %s\n", key, value)
//	}
package version

import (
	"fmt"
	"runtime/debug"
	"strings"
)

var (
	// Version is the release version, for example "v1.2.3". It is normally
	// set at build time with -ldflags -X and defaults to "dev".
	Version = "dev"

	// GitCommit is the full or abbreviated git commit hash. It is normally
	// set at build time with -ldflags -X and defaults to "unknown".
	GitCommit = "unknown"

	// GitTag is the git tag at build time. It is normally set at build time
	// with -ldflags -X and defaults to "unknown".
	GitTag = "unknown"

	// BuildTime is the RFC 3339 build timestamp. It is normally set at build
	// time with -ldflags -X and defaults to "unknown".
	BuildTime = "unknown"

	// GitDirty is "dirty" when the working tree contained uncommitted changes
	// at build time. It is normally set at build time with -ldflags -X and
	// defaults to the empty string.
	GitDirty = ""
)

// Get returns the version string for the application.
//
// It prefers the value injected at build time, then the main module version
// recorded by the Go toolchain, then a string assembled from git metadata,
// and finally "dev".
func Get() string {
	// If Version was set via ldflags, use it.
	if Version != "dev" {
		return Version
	}

	// Fallback: try to get version from build info.
	if info, ok := debug.ReadBuildInfo(); ok {
		if info.Main.Version != "(devel)" && info.Main.Version != "" {
			return info.Main.Version
		}
	}

	// Final fallback: construct from git information if available.
	if GitTag != "unknown" && GitCommit != "unknown" {
		version := GitTag
		if GitCommit != "" {
			// Safe substring that handles short commits.
			commitSuffix := GitCommit
			if len(GitCommit) > 7 {
				commitSuffix = GitCommit[:7]
			}
			if !strings.HasSuffix(GitTag, commitSuffix) {
				// Add commit hash if tag does not contain it.
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

// Full returns the version string with the commit hash appended when it is
// known.
func Full() string {
	version := Get()
	if GitCommit != "unknown" {
		return fmt.Sprintf("%s (commit: %s)", version, GitCommit)
	}
	return version
}

// Info returns the version and build metadata as a map suitable for JSON
// encoding. Keys are "version", "gitCommit", "gitTag", "buildTime", and
// "gitDirty".
func Info() map[string]string {
	return map[string]string{
		"version":   Get(),
		"gitCommit": GitCommit,
		"gitTag":    GitTag,
		"buildTime": BuildTime,
		"gitDirty":  GitDirty,
	}
}
