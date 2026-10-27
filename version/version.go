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
// A build can inject metadata with -ldflags -X, for example:
//
//	go build -ldflags "\
//	  -X bennypowers.dev/cem/version.Version=v1.2.3 \
//	  -X bennypowers.dev/cem/version.GitCommit=abcdef1 \
//	  -X bennypowers.dev/cem/version.GitTag=v1.2.3" .
//
// When the metadata is absent, Get reads the Go build information. If cem is
// the main module, it uses that module's version. If cem is embedded as a
// dependency, it reports the cem dependency version, honoring any Replace
// directive, rather than the host program's version. If no module version is
// available, Get assembles a version from GitTag, GitCommit, and GitDirty when
// they are set, and finally returns "dev".
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

// modulePath is the import path recorded in Go build information.
const modulePath = "bennypowers.dev/cem"

// readBuildInfo is indirected so tests can inject build information.
var readBuildInfo = debug.ReadBuildInfo

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
// It prefers the value injected at build time, then a usable version from Go
// build information, then a string assembled from git metadata, and finally
// "dev". For build information, it uses the main module version when cem is
// the main module; otherwise it scans dependencies for the cem module, so
// embedders report cem's version rather than the host program's version.
func Get() string {
	// Fast path: if Version was set via ldflags, use it without reading build
	// information.
	if Version != "dev" {
		return Version
	}

	info, _ := readBuildInfo()
	return get(Version, GitCommit, GitTag, GitDirty, info)
}

// get implements Get with its inputs supplied explicitly so the version
// resolution logic can be tested without build-time injection.
func get(version, gitCommit, gitTag, gitDirty string, info *debug.BuildInfo) string {
	if version != "dev" {
		return version
	}

	if v := moduleVersion(info); v != "" {
		return v
	}

	// Fallback: construct from git information if available.
	if gitTag != "unknown" && gitCommit != "unknown" {
		v := gitTag
		if gitCommit != "" {
			// Safe substring that handles short commits.
			commitSuffix := gitCommit
			if len(gitCommit) > 7 {
				commitSuffix = gitCommit[:7]
			}
			if !strings.HasSuffix(gitTag, commitSuffix) {
				// Add commit hash if tag does not contain it.
				v = fmt.Sprintf("%s-%s", gitTag, commitSuffix)
			}
		}
		if gitDirty == "dirty" {
			v += "-dirty"
		}
		return v
	}

	return "dev"
}

// moduleVersion returns the cem module version recorded in build information.
// It reports the main module version only when cem is the main module;
// otherwise it scans dependencies so embedders report the cem version, not the
// host program's version. Replacement versions are honored for any replacement
// path; local directory replacements have no version and are skipped.
func moduleVersion(info *debug.BuildInfo) string {
	if info == nil {
		return ""
	}
	if info.Main.Path == modulePath {
		if info.Main.Replace != nil {
			if v := usableVersion(info.Main.Replace.Version); v != "" {
				return v
			}
			// Local directory replacements have no version to report.
			return ""
		}
		return usableVersion(info.Main.Version)
	}
	for _, dep := range info.Deps {
		if dep == nil || dep.Path != modulePath {
			continue
		}
		if dep.Replace != nil {
			if v := usableVersion(dep.Replace.Version); v != "" {
				return v
			}
			// Local directory replacements have no version to report.
			continue
		}
		if v := usableVersion(dep.Version); v != "" {
			return v
		}
	}
	return ""
}

// usableVersion filters out versions that do not identify a release.
func usableVersion(v string) string {
	if v == "" || v == "(devel)" {
		return ""
	}
	return v
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
