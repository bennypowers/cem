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
	"sync"
	"testing"
	"testing/fstest"
	"time"

	"bennypowers.dev/cem/internal/platform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFileDependencyTracker_UpdateFileHash_ModTimeBasedCaching(t *testing.T) {
	// Inline: scalar hash equality checks don't benefit from golden files.
	t0 := time.Date(2025, 6, 1, 0, 0, 0, 0, time.UTC)
	t1 := t0.Add(time.Second)

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

	// Initial scan of both files.
	hashA1, err := tracker.UpdateFileHash("src/a.ts")
	require.NoError(t, err)
	hashB1, err := tracker.UpdateFileHash("src/b.ts")
	require.NoError(t, err)

	// Modify B with new modtime.
	mapFS.MapFS["src/b.ts"] = &fstest.MapFile{
		Data:    []byte("content-b-v2"),
		Mode:    0644,
		ModTime: t1,
	}

	// Re-scan A (unchanged modtime) -- cached hash returned.
	hashA2, err := tracker.UpdateFileHash("src/a.ts")
	require.NoError(t, err)
	assert.Equal(t, hashA1, hashA2, "same modtime = cached hash")

	// Re-scan B -- new modtime triggers rehash.
	hashB2, err := tracker.UpdateFileHash("src/b.ts")
	require.NoError(t, err)
	assert.NotEqual(t, hashB1, hashB2, "different modtime = rehash")
}

func TestFileDependencyTracker_HasFileChanged_PerFileTracking(t *testing.T) {
	// Inline: boolean changed/unchanged checks don't benefit from golden files.
	t0 := time.Date(2025, 6, 1, 0, 0, 0, 0, time.UTC)
	t1 := t0.Add(time.Second)

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

	// New file (no stored hash) reports changed.
	changed, err := tracker.HasFileChanged("src/x.ts")
	require.NoError(t, err)
	assert.True(t, changed, "new file should be reported as changed")

	// Seed hash.
	_, err = tracker.UpdateFileHash("src/x.ts")
	require.NoError(t, err)

	// Same modtime = unchanged.
	changed, err = tracker.HasFileChanged("src/x.ts")
	require.NoError(t, err)
	assert.False(t, changed, "unmodified file should not be reported as changed")

	// New modtime + new content = changed.
	mapFS.MapFS["src/x.ts"] = &fstest.MapFile{
		Data:    []byte("modified"),
		Mode:    0644,
		ModTime: t1,
	}

	changed, err = tracker.HasFileChanged("src/x.ts")
	require.NoError(t, err)
	assert.True(t, changed, "modified file should be reported as changed")
}

func TestFileDependencyTracker_UpdateFileHash_SameModTimeCachesHash(t *testing.T) {
	// Inline: scalar hash equality documents design trade-off, not golden-file material.
	t0 := time.Date(2025, 6, 1, 0, 0, 0, 0, time.UTC)

	mapFS := &platform.MapFS{
		MapFS: fstest.MapFS{
			"file.ts": &fstest.MapFile{
				Data:    []byte("original"),
				Mode:    0644,
				ModTime: t0,
			},
		},
	}

	tracker := NewFileDependencyTracker(nil, mapFS)

	hash1, err := tracker.UpdateFileHash("file.ts")
	require.NoError(t, err)

	// Change content but keep same modtime.
	mapFS.MapFS["file.ts"] = &fstest.MapFile{
		Data:    []byte("modified"),
		Mode:    0644,
		ModTime: t0,
	}

	hash2, err := tracker.UpdateFileHash("file.ts")
	require.NoError(t, err)
	assert.Equal(t, hash1, hash2,
		"same modtime returns cached hash; content-only changes with preserved mtime are not detected")
}

func TestFileDependencyTracker_UpdateFileHash_Concurrent(t *testing.T) {
	// Inline: verifying all goroutines return identical hashes is a scalar check.
	t0 := time.Date(2025, 6, 1, 0, 0, 0, 0, time.UTC)

	mapFS := &platform.MapFS{
		MapFS: fstest.MapFS{
			"file.ts": &fstest.MapFile{
				Data:    []byte("content"),
				Mode:    0644,
				ModTime: t0,
			},
		},
	}

	tracker := NewFileDependencyTracker(nil, mapFS)

	var wg sync.WaitGroup
	hashes := make([][32]byte, 10)
	errs := make([]error, 10)

	for i := range 10 {
		wg.Go(func() {
			hashes[i], errs[i] = tracker.UpdateFileHash("file.ts")
		})
	}
	wg.Wait()

	for i := range 10 {
		require.NoError(t, errs[i])
		assert.Equal(t, hashes[0], hashes[i], "concurrent calls should return same hash")
	}
}
