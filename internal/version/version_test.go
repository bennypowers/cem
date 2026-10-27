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

	internalversion "bennypowers.dev/cem/internal/version"
	publicversion "bennypowers.dev/cem/version"
)

func TestInternalShimDelegatesToPublicPackage(t *testing.T) {
	if got, want := internalversion.GetVersion(), publicversion.Get(); got != want {
		t.Errorf("GetVersion() = %q, want %q", got, want)
	}
	if got, want := internalversion.GetFullVersion(), publicversion.Full(); got != want {
		t.Errorf("GetFullVersion() = %q, want %q", got, want)
	}

	internalInfo := internalversion.GetBuildInfo()
	publicInfo := publicversion.Info()
	if len(internalInfo) != len(publicInfo) {
		t.Fatalf("GetBuildInfo() has %d entries, want %d", len(internalInfo), len(publicInfo))
	}
	for key, want := range publicInfo {
		if got := internalInfo[key]; got != want {
			t.Errorf("GetBuildInfo()[%q] = %q, want %q", key, got, want)
		}
	}
}
