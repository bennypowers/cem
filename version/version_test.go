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
package version_test

import (
	"testing"

	"bennypowers.dev/cem/version"
)

// setBuildState replaces the package-level build metadata for the duration of
// a test and restores the previous values when the test finishes.
func setBuildState(t *testing.T, v, commit, tag, buildTime, dirty string) {
	t.Helper()

	oldVersion := version.Version
	oldCommit := version.GitCommit
	oldTag := version.GitTag
	oldBuildTime := version.BuildTime
	oldDirty := version.GitDirty

	version.Version = v
	version.GitCommit = commit
	version.GitTag = tag
	version.BuildTime = buildTime
	version.GitDirty = dirty

	t.Cleanup(func() {
		version.Version = oldVersion
		version.GitCommit = oldCommit
		version.GitTag = oldTag
		version.BuildTime = oldBuildTime
		version.GitDirty = oldDirty
	})
}

// Inline assertions: this test compares a scalar version string, where a
// golden file would only duplicate the expected value.
func TestGetDefaultsToDev(t *testing.T) {
	setBuildState(t, "dev", "unknown", "unknown", "unknown", "")

	if got := version.Get(); got != "dev" {
		t.Errorf("Get() = %q, want %q", got, "dev")
	}
}

// Inline assertions: this test compares a scalar version string, where a
// golden file would only duplicate the expected value.
func TestGetUsesVersionOverride(t *testing.T) {
	setBuildState(t, "v1.2.3", "abcdef1", "v1.2.3", "2026-01-02T03:04:05Z", "")

	if got := version.Get(); got != "v1.2.3" {
		t.Errorf("Get() = %q, want %q", got, "v1.2.3")
	}
}

// Inline assertions: this table compares scalar version strings, where a
// golden file would only duplicate the expected values.
func TestGetCombinesGitMetadata(t *testing.T) {
	tests := []struct {
		name   string
		commit string
		tag    string
		dirty  string
		want   string
	}{
		{
			name:   "tag and commit",
			commit: "abcdef123456",
			tag:    "v1.2.3",
			want:   "v1.2.3-abcdef1",
		},
		{
			name:   "tag already contains short commit",
			commit: "abcdef123456",
			tag:    "v1.2.3-abcdef1",
			want:   "v1.2.3-abcdef1",
		},
		{
			name:   "dirty working tree",
			commit: "abcdef123456",
			tag:    "v1.2.3",
			dirty:  "dirty",
			want:   "v1.2.3-abcdef1-dirty",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			setBuildState(t, "dev", tc.commit, tc.tag, "unknown", tc.dirty)

			if got := version.Get(); got != tc.want {
				t.Errorf("Get() = %q, want %q", got, tc.want)
			}
		})
	}
}

// Inline assertions: this test compares a scalar version string, where a
// golden file would only duplicate the expected value.
func TestFullIncludesCommit(t *testing.T) {
	setBuildState(t, "v1.2.3", "abcdef123456", "v1.2.3", "unknown", "")

	if got, want := version.Full(), "v1.2.3 (commit: abcdef123456)"; got != want {
		t.Errorf("Full() = %q, want %q", got, want)
	}
}

// Inline assertions: this test compares a scalar version string, where a
// golden file would only duplicate the expected value.
func TestFullWithoutCommit(t *testing.T) {
	setBuildState(t, "v1.2.3", "unknown", "v1.2.3", "unknown", "")

	if got, want := version.Full(), "v1.2.3"; got != want {
		t.Errorf("Full() = %q, want %q", got, want)
	}
}

// Inline assertions: this test compares a small map of scalar metadata, where
// a golden file would only duplicate the expected values.
func TestInfoReportsBuildMetadata(t *testing.T) {
	setBuildState(t, "v1.2.3", "abcdef123456", "v1.2.3", "2026-01-02T03:04:05Z", "")

	info := version.Info()

	want := map[string]string{
		"version":   "v1.2.3",
		"gitCommit": "abcdef123456",
		"gitTag":    "v1.2.3",
		"buildTime": "2026-01-02T03:04:05Z",
		"gitDirty":  "",
	}
	for key, wantValue := range want {
		if got := info[key]; got != wantValue {
			t.Errorf("Info()[%q] = %q, want %q", key, got, wantValue)
		}
	}
	if len(info) != len(want) {
		t.Errorf("Info() has %d entries, want %d: %v", len(info), len(want), info)
	}
}
