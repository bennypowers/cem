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
package version

import (
	"runtime/debug"
	"testing"
)

func buildInfo(main debug.Module, deps ...*debug.Module) *debug.BuildInfo {
	return &debug.BuildInfo{Main: main, Deps: deps}
}

// Inline assertions: this test compares a scalar version string, where a
// golden file would only duplicate the expected value.
func TestGetHandlesNilBuildInfo(t *testing.T) {
	got := get("dev", "unknown", "unknown", "", nil)
	if got != "dev" {
		t.Fatalf("get() = %q, want %q", got, "dev")
	}
}

// Inline assertions: this test compares a scalar version string, where a
// golden file would only duplicate the expected value.
func TestGetFallsThroughDevelMainModuleToGitMetadata(t *testing.T) {
	got := get("dev", "abcdef1234567890", "v0.3.3", "", buildInfo(debug.Module{
		Path:    modulePath,
		Version: "(devel)",
	}))
	if got != "v0.3.3-abcdef1" {
		t.Fatalf("get() = %q, want %q", got, "v0.3.3-abcdef1")
	}
}

// Inline assertions: this test compares a scalar version string, where a
// golden file would only duplicate the expected value.
func TestGetUsesDifferentPathReplacementVersion(t *testing.T) {
	got := get("dev", "unknown", "unknown", "", buildInfo(
		debug.Module{Path: "example.com/host", Version: "v9.9.9"},
		&debug.Module{
			Path:    modulePath,
			Version: "v0.0.0",
			Replace: &debug.Module{Path: "example.com/cem-fork", Version: "v0.4.0"},
		},
	))
	if got != "v0.4.0" {
		t.Fatalf("get() = %q, want different-path replacement version %q", got, "v0.4.0")
	}
}

// Inline assertions: this test compares a scalar version string, where a
// golden file would only duplicate the expected value.
//
// Go records a local directory replacement with version "(devel)", which has
// no release version to report, so (devel) is the realistic fixture value.
func TestGetSkipsLocalReplaceWithoutUsableVersion(t *testing.T) {
	got := get("dev", "abcdef1234567890", "v0.3.3", "", buildInfo(
		debug.Module{Path: "example.com/host", Version: "v9.9.9"},
		&debug.Module{
			Path:    modulePath,
			Version: "v0.3.3",
			Replace: &debug.Module{Path: "../cem", Version: "(devel)"},
		},
	))
	if got != "v0.3.3-abcdef1" {
		t.Fatalf("get() = %q, want git fallback %q", got, "v0.3.3-abcdef1")
	}
}

// Inline assertions: this test compares a scalar version string, where a
// golden file would only duplicate the expected value.
func TestGetPrefersInjectedVersionWithoutReadingBuildInfo(t *testing.T) {
	oldVersion := Version
	oldReadBuildInfo := readBuildInfo
	Version = "v2.0.0"
	readBuildInfo = func() (*debug.BuildInfo, bool) {
		t.Fatal("build info should not be read when Version is injected")
		return nil, false
	}
	t.Cleanup(func() {
		Version = oldVersion
		readBuildInfo = oldReadBuildInfo
	})

	if got := Get(); got != "v2.0.0" {
		t.Fatalf("Get() = %q, want %q", got, "v2.0.0")
	}
}

// Inline assertions: this test compares a scalar version string, where a
// golden file would only duplicate the expected value.
func TestGetUsesMainModuleVersionForCLI(t *testing.T) {
	got := get("dev", "unknown", "unknown", "", buildInfo(debug.Module{
		Path:    modulePath,
		Version: "v1.2.3",
	}))
	if got != "v1.2.3" {
		t.Fatalf("get() = %q, want %q", got, "v1.2.3")
	}
}

// Inline assertions: this test compares a scalar version string, where a
// golden file would only duplicate the expected value.
func TestGetUsesCemDependencyVersionForEmbedder(t *testing.T) {
	got := get("dev", "unknown", "unknown", "", buildInfo(
		debug.Module{Path: "example.com/host", Version: "v9.9.9"},
		&debug.Module{Path: modulePath, Version: "v0.3.3"},
	))
	if got != "v0.3.3" {
		t.Fatalf("get() = %q, want the cem dependency version %q", got, "v0.3.3")
	}
}

// Inline assertions: this test compares a scalar version string, where a
// golden file would only duplicate the expected value.
func TestGetUsesReplacementVersionForEmbedder(t *testing.T) {
	got := get("dev", "unknown", "unknown", "", buildInfo(
		debug.Module{Path: "example.com/host", Version: "v9.9.9"},
		&debug.Module{
			Path:    modulePath,
			Version: "v0.0.0",
			Replace: &debug.Module{Path: modulePath, Version: "v0.4.0"},
		},
	))
	if got != "v0.4.0" {
		t.Fatalf("get() = %q, want replacement version %q", got, "v0.4.0")
	}
}

// Inline assertions: this test compares a scalar version string, where a
// golden file would only duplicate the expected value.
func TestGetFallsBackWhenDependencyMissing(t *testing.T) {
	got := get("dev", "unknown", "unknown", "", buildInfo(
		debug.Module{Path: "example.com/host", Version: "v9.9.9"},
	))
	if got != "dev" {
		t.Fatalf("get() = %q, want %q", got, "dev")
	}
}

// Inline assertions: this test compares a scalar version string, where a
// golden file would only duplicate the expected value.
func TestGetFallsBackToGitMetadataWhenDependencyMissing(t *testing.T) {
	got := get("dev", "abcdef1234567890", "v0.3.3", "", buildInfo(
		debug.Module{Path: "example.com/host", Version: "v9.9.9"},
	))
	if got != "v0.3.3-abcdef1" {
		t.Fatalf("get() = %q, want %q", got, "v0.3.3-abcdef1")
	}
}

// Inline assertions: this test compares a scalar version string, where a
// golden file would only duplicate the expected value.
func TestGetPrefersInjectedVersion(t *testing.T) {
	got := get("v2.0.0", "unknown", "unknown", "", buildInfo(
		debug.Module{Path: "example.com/host", Version: "v9.9.9"},
		&debug.Module{Path: modulePath, Version: "v0.3.3"},
	))
	if got != "v2.0.0" {
		t.Fatalf("get() = %q, want %q", got, "v2.0.0")
	}
}
