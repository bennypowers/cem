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
	"testing/fstest"
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
		// Handle root path specially
		if watchedPath == "/" {
			watched = true
			break
		}
		// Check exact match or subdirectory
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

// MapFileSystem wraps Go's testing/fstest.MapFS to implement the platform.FileSystem interface.
// This provides a standards-compliant in-memory filesystem for testing using Go 1.25's
// improved fstest package features.
//
// Features from Go 1.25:
// - Enhanced MapFS with Lstat() for symlink handling
// - ReadLink() support for symbolic links
// - Comprehensive TestFS validation
// - Integration with file watchers for event simulation
type MapFileSystem struct {
	mu           sync.RWMutex
	mapFS        fstest.MapFS
	tempDir      string
	watcher      *MockFileWatcher
	timeProvider TimeProvider
}

// NewMapFileSystem creates a new filesystem based on testing/fstest.MapFS.
// If timeProvider is nil, a default mock time provider is created.
func NewMapFileSystem(timeProvider TimeProvider) *MapFileSystem {
	if timeProvider == nil {
		timeProvider = NewMockTimeProvider(time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC))
	}

	return &MapFileSystem{
		mapFS:        make(fstest.MapFS),
		tempDir:      "/tmp",
		watcher:      NewMockFileWatcher(),
		timeProvider: timeProvider,
	}
}

// SetWatcher allows injection of a custom file watcher for testing integration.
func (mfs *MapFileSystem) SetWatcher(watcher *MockFileWatcher) {
	mfs.mu.Lock()
	defer mfs.mu.Unlock()
	mfs.watcher = watcher
}

// GetWatcher returns the associated file watcher for testing integration.
func (mfs *MapFileSystem) GetWatcher() *MockFileWatcher {
	mfs.mu.RLock()
	defer mfs.mu.RUnlock()
	return mfs.watcher
}

// GetMapFS returns the underlying fstest.MapFS for direct manipulation in tests.
func (mfs *MapFileSystem) GetMapFS() fstest.MapFS {
	mfs.mu.RLock()
	defer mfs.mu.RUnlock()
	return mfs.mapFS
}

func (mfs *MapFileSystem) WriteFile(name string, data []byte, perm fs.FileMode) error {
	mfs.mu.Lock()
	defer mfs.mu.Unlock()

	name = mfs.cleanPath(name)

	// Ensure parent directory structure exists in MapFS
	if err := mfs.ensureParentDirLocked(name); err != nil {
		return err
	}

	// Create MapFile
	mfs.mapFS[name] = &fstest.MapFile{
		Data:    append([]byte(nil), data...), // copy data
		Mode:    perm,
		ModTime: mfs.timeProvider.Now(),
	}

	// Trigger file watcher event
	if mfs.watcher != nil {
		mfs.watcher.TriggerEvent("/"+name, Write)
	}

	return nil
}

func (mfs *MapFileSystem) ReadFile(name string) ([]byte, error) {
	mfs.mu.RLock()
	defer mfs.mu.RUnlock()

	name = mfs.cleanPath(name)

	// Use MapFS's ReadFile method
	return fs.ReadFile(mfs.mapFS, name)
}

func (mfs *MapFileSystem) Remove(name string) error {
	mfs.mu.Lock()
	defer mfs.mu.Unlock()

	name = mfs.cleanPath(name)

	// Check if file exists
	if _, exists := mfs.mapFS[name]; !exists {
		return &fs.PathError{Op: "remove", Path: name, Err: fs.ErrNotExist}
	}

	delete(mfs.mapFS, name)

	// Trigger file watcher event
	if mfs.watcher != nil {
		mfs.watcher.TriggerEvent("/"+name, Remove)
	}

	return nil
}

func (mfs *MapFileSystem) MkdirAll(path string, perm fs.FileMode) error {
	mfs.mu.Lock()
	defer mfs.mu.Unlock()

	path = mfs.cleanPath(path)

	// MapFS represents directories implicitly through file paths
	// We'll create a .keep file to represent empty directories
	keepFile := path + "/.keep"

	// Check if path conflicts with existing file
	if file, exists := mfs.mapFS[path]; exists && !file.Mode.IsDir() {
		return &fs.PathError{Op: "mkdir", Path: path, Err: fmt.Errorf("not a directory")}
	}

	// Create the .keep file to represent the directory
	mfs.mapFS[keepFile] = &fstest.MapFile{
		Data:    []byte(""),
		Mode:    0644,
		ModTime: mfs.timeProvider.Now(),
	}

	// Trigger file watcher event
	if mfs.watcher != nil {
		mfs.watcher.TriggerEvent("/"+path, Create)
	}

	return nil
}

func (mfs *MapFileSystem) TempDir() string {
	mfs.mu.RLock()
	defer mfs.mu.RUnlock()
	return mfs.tempDir
}

// SetTempDir allows customizing the temp directory path for testing.
func (mfs *MapFileSystem) SetTempDir(dir string) {
	mfs.mu.Lock()
	defer mfs.mu.Unlock()
	mfs.tempDir = dir
}

func (mfs *MapFileSystem) Stat(name string) (fs.FileInfo, error) {
	mfs.mu.RLock()
	defer mfs.mu.RUnlock()

	name = mfs.cleanPath(name)

	// Use MapFS's Stat method (available via fs.Stat)
	return fs.Stat(mfs.mapFS, name)
}

func (mfs *MapFileSystem) Exists(path string) bool {
	mfs.mu.RLock()
	defer mfs.mu.RUnlock()

	path = mfs.cleanPath(path)

	// Check if file exists directly
	if _, exists := mfs.mapFS[path]; exists {
		return true
	}

	// Check if it's a directory by looking for any files that start with path + "/"
	prefix := path + "/"
	for filePath := range mfs.mapFS {
		if strings.HasPrefix(filePath, prefix) {
			return true
		}
	}

	return false
}

// Helper methods

func (mfs *MapFileSystem) cleanPath(path string) string {
	cleaned := filepath.Clean(path)
	if !filepath.IsAbs(cleaned) {
		cleaned = "/" + cleaned
	}
	// Remove leading slash for MapFS compatibility (MapFS uses relative paths)
	return strings.TrimPrefix(cleaned, "/")
}

func (mfs *MapFileSystem) ensureParentDirLocked(filePath string) error {
	dir := filepath.Dir(filePath)
	if dir == "." || dir == "/" || dir == "" {
		return nil
	}

	// For MapFS, we don't need to pre-create directories
	// They are implicitly created when files are added
	// Just check if a file exists with the same name as the parent directory
	if file, exists := mfs.mapFS[dir]; exists && !file.Mode.IsDir() {
		return &fs.PathError{Op: "open", Path: filePath, Err: fmt.Errorf("not a directory")}
	}

	return nil
}

func (mfs *MapFileSystem) existsLocked(path string) bool {
	// Check if file exists directly
	if _, exists := mfs.mapFS[path]; exists {
		return true
	}

	// Check if it's a directory by looking for any files that start with path + "/"
	prefix := path + "/"
	for filePath := range mfs.mapFS {
		if strings.HasPrefix(filePath, prefix) {
			return true
		}
	}

	return false
}

// TestFS validates the filesystem using Go's fstest.TestFS.
// This is particularly useful for ensuring filesystem compliance.
func (mfs *MapFileSystem) TestFS(expectedFiles ...string) error {
	mfs.mu.RLock()
	defer mfs.mu.RUnlock()

	return fstest.TestFS(mfs.mapFS, expectedFiles...)
}

// Debug utilities

// ListFiles returns all files in the MapFS for debugging.
func (mfs *MapFileSystem) ListFiles() map[string]interface{} {
	mfs.mu.RLock()
	defer mfs.mu.RUnlock()

	result := make(map[string]interface{})

	for path, file := range mfs.mapFS {
		if file.Mode.IsDir() {
			result[path] = "directory"
		} else {
			result[path] = fmt.Sprintf("file (%d bytes)", len(file.Data))
		}
	}

	return result
}

// AddFile adds a file directly to the MapFS for test setup.
// This is a convenience method for test initialization.
func (mfs *MapFileSystem) AddFile(path string, content string, mode fs.FileMode) {
	mfs.mu.Lock()
	defer mfs.mu.Unlock()

	path = mfs.cleanPath(path)
	mfs.mapFS[path] = &fstest.MapFile{
		Data:    []byte(content),
		Mode:    mode,
		ModTime: mfs.timeProvider.Now(),
	}
}

// AddDir adds a directory directly to the MapFS for test setup.
// Note: MapFS represents directories implicitly through file paths.
// Empty directories need a placeholder file to exist.
func (mfs *MapFileSystem) AddDir(path string, mode fs.FileMode) {
	mfs.mu.Lock()
	defer mfs.mu.Unlock()

	path = mfs.cleanPath(path)

	// For empty directories, create a special .keep file
	keepFile := path + "/.keep"
	mfs.mapFS[keepFile] = &fstest.MapFile{
		Data:    []byte(""),
		Mode:    0644,
		ModTime: mfs.timeProvider.Now(),
	}
}
