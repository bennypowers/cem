/*
Copyright © 2025 Benny Powers

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
	"testing/fstest"
	"time"

	"bennypowers.dev/cem/internal/platform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Inline: unit tests for per-file scan time tracking; scalar hash comparisons
// don't benefit from golden files.

func TestFileDependencyTracker_UpdateFileHash_SharedTimestampBug(t *testing.T) {
	// Regression: a shared lastScanTime field caused one file's scan to push the
	// timestamp forward, making a different file's modification appear stale.
	//
	// Timeline (wall clock order):
	//   1. Scan A, scan B              -- both hashed
	//   2. Modify B on disk            -- new content, ModTime = time of write
	//   3. Re-scan A (A unchanged)     -- with the bug, this bumps the shared
	//                                     lastScanTime past B's ModTime
	//   4. Re-scan B                   -- BUG: B's ModTime < shared lastScanTime
	//                                     so the tracker returns the stale hash
	//
	// Fix: each file tracks its own scannedAt, so scanning A cannot affect B.

	t0 := time.Date(2025, 6, 1, 0, 0, 0, 0, time.UTC)

	mapFS := &platform.MapFS{
		MapFS: fstest.MapFS{
			"src/a.ts": &fstest.MapFile{
				Data:    []byte("content-a-v1"),
				Mode:    0644,
				ModTime: t0,
			},
			"src/b.ts": &fstest.MapFile{
				Data:    []byte("content-b-v1"),
				Mode:    0644,
				ModTime: t0,
			},
		},
	}

	tracker := NewFileDependencyTracker(nil, mapFS)

	// Step 1: Initial scan of both files.
	hashA1, err := tracker.UpdateFileHash("src/a.ts")
	require.NoError(t, err)
	hashB1, err := tracker.UpdateFileHash("src/b.ts")
	require.NoError(t, err)

	// Step 2: Modify file B. Give it a ModTime slightly in the future so the
	// filesystem reports a newer timestamp, but still before time.Now() when
	// we re-scan A in step 3.
	time.Sleep(10 * time.Millisecond) // ensure ModTime ordering
	tMod := time.Now()
	mapFS.MapFS["src/b.ts"] = &fstest.MapFile{
		Data:    []byte("content-b-v2"),
		Mode:    0644,
		ModTime: tMod,
	}

	// Step 3: Touch file A so it needs a rescan (ModTime must be newer than A's
	// scannedAt). This re-scan sets lastScanTime = time.Now() in the buggy code.
	time.Sleep(10 * time.Millisecond)
	mapFS.MapFS["src/a.ts"] = &fstest.MapFile{
		Data:    []byte("content-a-v1"), // same content
		Mode:    0644,
		ModTime: time.Now(),
	}
	hashA2, err := tracker.UpdateFileHash("src/a.ts")
	require.NoError(t, err)
	assert.Equal(t, hashA1, hashA2, "file A content unchanged, hash should match")

	// Step 4: Re-scan file B. With the shared-timestamp bug, B's ModTime (tMod)
	// is before the lastScanTime that A's re-scan just set, so the tracker
	// incorrectly returns the stale hash.
	time.Sleep(10 * time.Millisecond)
	hashB2, err := tracker.UpdateFileHash("src/b.ts")
	require.NoError(t, err)
	assert.NotEqual(t, hashB1, hashB2,
		"file B hash should differ after content change; shared scan timestamp leaked from file A")
}

func TestFileDependencyTracker_HasFileChanged_PerFileTracking(t *testing.T) {
	t0 := time.Date(2025, 6, 1, 0, 0, 0, 0, time.UTC)

	mapFS := &platform.MapFS{
		MapFS: fstest.MapFS{
			"src/x.ts": &fstest.MapFile{
				Data:    []byte("original"),
				Mode:    0644,
				ModTime: t0,
			},
		},
	}

	tracker := NewFileDependencyTracker(nil, mapFS)

	// First check: file is new (no stored hash), should be reported as changed.
	changed, err := tracker.HasFileChanged("src/x.ts")
	require.NoError(t, err)
	assert.True(t, changed, "new file should be reported as changed")

	// Seed the hash via UpdateFileHash (as production callers do after processing).
	_, err = tracker.UpdateFileHash("src/x.ts")
	require.NoError(t, err)

	// Second check without modification: should not be changed.
	changed, err = tracker.HasFileChanged("src/x.ts")
	require.NoError(t, err)
	assert.False(t, changed, "unmodified file should not be reported as changed")

	// Modify the file with a ModTime after the scan time.
	time.Sleep(10 * time.Millisecond)
	mapFS.MapFS["src/x.ts"] = &fstest.MapFile{
		Data:    []byte("modified"),
		Mode:    0644,
		ModTime: time.Now(),
	}

	changed, err = tracker.HasFileChanged("src/x.ts")
	require.NoError(t, err)
	assert.True(t, changed, "modified file should be reported as changed")
}
