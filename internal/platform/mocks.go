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

package platform

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// MockTimeProvider provides controllable time for testing.
// Time advances only when explicitly advanced, eliminating real delays.
type MockTimeProvider struct {
	mu          sync.Mutex
	currentTime time.Time
	sleepCalls  []time.Duration // Track sleep calls for testing
}

// NewMockTimeProvider creates a new mock time provider starting at the given time.
func NewMockTimeProvider(startTime time.Time) *MockTimeProvider {
	return &MockTimeProvider{
		currentTime: startTime,
		sleepCalls:  make([]time.Duration, 0),
	}
}

func (m *MockTimeProvider) Sleep(d time.Duration) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.sleepCalls = append(m.sleepCalls, d)
	m.currentTime = m.currentTime.Add(d)
	// No actual sleeping - time advances instantly
}

func (m *MockTimeProvider) Now() time.Time {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.currentTime
}

func (m *MockTimeProvider) After(d time.Duration) <-chan time.Time {
	// For testing, advance time instantly and return the result immediately
	m.Sleep(d) // Advance time instantly

	ch := make(chan time.Time, 1)
	ch <- m.Now()
	close(ch)
	return ch
}

// AdvanceTime manually advances the mock time by the given duration.
func (m *MockTimeProvider) AdvanceTime(d time.Duration) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.currentTime = m.currentTime.Add(d)
}

// GetSleepCalls returns all Sleep() calls made to this provider.
func (m *MockTimeProvider) GetSleepCalls() []time.Duration {
	m.mu.Lock()
	defer m.mu.Unlock()
	return append([]time.Duration(nil), m.sleepCalls...)
}

// MockFileWatcher provides controllable file watching for testing.
// Events are triggered manually, eliminating real file system dependencies.
// With Go 1.25's synctest, channels work reliably in concurrent tests.
type MockFileWatcher struct {
	mu           sync.RWMutex
	watchedPaths map[string]bool
	closed       bool
	events       chan FileWatchEvent
	errors       chan error
}

// NewMockFileWatcher creates a new mock file watcher.
func NewMockFileWatcher() *MockFileWatcher {
	return &MockFileWatcher{
		watchedPaths: make(map[string]bool),
		events:       make(chan FileWatchEvent, 100), // Buffered to prevent blocking
		errors:       make(chan error, 10),
	}
}

func (m *MockFileWatcher) Add(name string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.closed {
		return fmt.Errorf("file watcher is closed")
	}

	m.watchedPaths[name] = true
	return nil
}

func (m *MockFileWatcher) Remove(name string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.closed {
		return fmt.Errorf("file watcher is closed")
	}

	delete(m.watchedPaths, name)
	return nil
}

func (m *MockFileWatcher) Close() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.closed {
		return nil
	}

	m.closed = true
	close(m.events)
	close(m.errors)
	return nil
}

func (m *MockFileWatcher) Events() <-chan FileWatchEvent {
	return m.events
}

func (m *MockFileWatcher) Errors() <-chan error {
	return m.errors
}

// TriggerEvent manually triggers a file watch event.
// This allows tests to simulate file changes instantly.
// With Go 1.25's synctest, events are now safely delivered via channels.
func (m *MockFileWatcher) TriggerEvent(name string, op WatchOp) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if m.closed {
		return
	}

	// Check if path is being watched
	watched := false
	for watchedPath := range m.watchedPaths {
		if name == watchedPath || strings.HasPrefix(name, watchedPath+string(filepath.Separator)) {
			watched = true
			break
		}
	}

	if watched {
		// Send event via channel - synctest makes this safe
		select {
		case m.events <- FileWatchEvent{Name: name, Op: op}:
		default:
			// Channel full, drop event (shouldn't happen with buffered channel)
		}
	}
}

// TriggerError manually triggers a file watch error.
// With Go 1.25's synctest, errors are now safely delivered via channels.
func (m *MockFileWatcher) TriggerError(err error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if !m.closed {
		// Send error via channel - synctest makes this safe
		select {
		case m.errors <- err:
		default:
			// Channel full, drop error (shouldn't happen with buffered channel)
		}
	}
}

// GetWatchedPaths returns the currently watched paths.
func (m *MockFileWatcher) GetWatchedPaths() []string {
	m.mu.RLock()
	defer m.mu.RUnlock()

	paths := make([]string, 0, len(m.watchedPaths))
	for path := range m.watchedPaths {
		paths = append(paths, path)
	}
	return paths
}

// TempDirFileSystem wraps the OS filesystem but uses a temporary directory
// for all operations. This provides isolation for tests while still using
// real filesystem operations (useful for testing filesystem-dependent code
// without mocking the entire filesystem).
type TempDirFileSystem struct {
	*OSFileSystem
	tempDir string
}

// NewTempDirFileSystem creates a filesystem that operates within a temporary directory.
// This provides test isolation while maintaining filesystem realism.
func NewTempDirFileSystem() (*TempDirFileSystem, error) {
	tempDir, err := os.MkdirTemp("", "cem-test-*")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp directory: %w", err)
	}

	return &TempDirFileSystem{
		OSFileSystem: NewOSFileSystem(),
		tempDir:      tempDir,
	}, nil
}

// resolvePath converts relative paths to absolute paths within the temp directory
func (fs *TempDirFileSystem) resolvePath(name string) string {
	if filepath.IsAbs(name) {
		// For absolute paths, place them relative to temp dir
		rel, err := filepath.Rel("/", name)
		if err != nil {
			rel = name[1:] // fallback: remove leading slash
		}
		return filepath.Join(fs.tempDir, rel)
	}
	return filepath.Join(fs.tempDir, name)
}

func (fs *TempDirFileSystem) WriteFile(name string, data []byte, perm fs.FileMode) error {
	path := fs.resolvePath(name)

	// Ensure parent directory exists
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}

	return fs.OSFileSystem.WriteFile(path, data, perm)
}

func (fs *TempDirFileSystem) ReadFile(name string) ([]byte, error) {
	return fs.OSFileSystem.ReadFile(fs.resolvePath(name))
}

func (fs *TempDirFileSystem) Remove(name string) error {
	return fs.OSFileSystem.Remove(fs.resolvePath(name))
}

func (fs *TempDirFileSystem) MkdirAll(path string, perm fs.FileMode) error {
	return fs.OSFileSystem.MkdirAll(fs.resolvePath(path), perm)
}

func (fs *TempDirFileSystem) TempDir() string {
	return fs.tempDir
}

func (fs *TempDirFileSystem) Stat(name string) (fs.FileInfo, error) {
	return fs.OSFileSystem.Stat(fs.resolvePath(name))
}

func (fs *TempDirFileSystem) Exists(path string) bool {
	return fs.OSFileSystem.Exists(fs.resolvePath(path))
}

// Cleanup removes the temporary directory and all its contents.
// Should be called when done with the filesystem (typically in test cleanup).
func (fs *TempDirFileSystem) Cleanup() error {
	return os.RemoveAll(fs.tempDir)
}

// RealPath returns the real filesystem path for a given logical path.
// Useful for debugging or when you need to interact with the temp directory directly.
func (fs *TempDirFileSystem) RealPath(name string) string {
	return fs.resolvePath(name)
}
