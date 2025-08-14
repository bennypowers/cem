# Platform Abstractions

This package provides platform abstractions for filesystem, time, and file watching operations. These interfaces enable testing with controllable dependencies and future portability to WASM, cloud functions, and embedded systems.

## Interfaces

### FileSystem
Abstracts filesystem operations for testing and portability.

**Production**: `OSFileSystem` - uses standard `os` package
**Testing**: `TempDirFileSystem` - isolated temporary directory operations
**Future**: Could support in-memory, WASM, or remote filesystems

```go
// Production usage
fs := platform.NewOSFileSystem()

// Test usage with isolation
fs, err := platform.NewTempDirFileSystem()
defer fs.Cleanup()

// All operations work the same
err = fs.WriteFile("config.json", data, 0644)
data, err := fs.ReadFile("config.json")
```

### TimeProvider
Abstracts time operations for deterministic testing.

**Production**: `RealTimeProvider` - uses standard `time` package
**Testing**: `MockTimeProvider` - controllable, instant time advancement

```go
// Production usage
timeProvider := platform.NewRealTimeProvider()

// Test usage with controllable time
startTime := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
timeProvider := platform.NewMockTimeProvider(startTime)

// All operations work the same
timeProvider.Sleep(5 * time.Second) // Instant in tests, real delay in production
now := timeProvider.Now()           // Controllable in tests
```

### FileWatcher
Abstracts file watching for testable event handling.

**Production**: `FSNotifyFileWatcher` - uses `fsnotify` package
**Testing**: `MockFileWatcher` - manually triggered events

```go
// Production usage
watcher, err := platform.NewFSNotifyFileWatcher()
defer watcher.Close()

// Test usage with manual control
watcher := platform.NewMockFileWatcher()
defer watcher.Close()

// All operations work the same
err = watcher.Add("/path/to/watch")
events := watcher.Events()

// In tests, manually trigger events
mockWatcher.TriggerEvent("/path/to/file.txt", platform.Write)
```

## Testing Benefits

### Before (Direct Dependencies)
```go
func TestSomething(t *testing.T) {
    // Real file I/O
    tempDir := t.TempDir()
    os.WriteFile(filepath.Join(tempDir, "test.txt"), data, 0644)
    
    // Real time delays
    time.Sleep(5 * time.Second) // Slow!
    
    // Real file watching (complex setup)
    watcher, _ := fsnotify.NewWatcher()
    watcher.Add(tempDir)
    // ... complex event handling
}
```

### After (Dependency Injection)
```go
func TestSomething(t *testing.T) {
    // Isolated filesystem
    fs, _ := platform.NewTempDirFileSystem()
    defer fs.Cleanup()
    
    // Instant time
    timeProvider := platform.NewMockTimeProvider(time.Now())
    
    // Controllable file watching
    watcher := platform.NewMockFileWatcher()
    defer watcher.Close()
    
    // Test the actual logic without environmental dependencies
    component := NewComponent(fs, timeProvider, watcher)
    
    // Trigger events instantly
    watcher.TriggerEvent("/test/file.txt", platform.Write)
    
    // Verify behavior without real delays
    if timeProvider.GetSleepCalls()[0] != 5*time.Second {
        t.Error("Expected 5 second delay")
    }
}
```

## Portability Benefits

These interfaces prepare the codebase for future deployment scenarios:

- **WASM**: Custom filesystem implementations for browser environments
- **Cloud Functions**: Remote storage backends (S3, GCS, etc.)
- **Embedded Systems**: Custom storage and timing implementations
- **Testing**: Hermetic, fast, deterministic test environments

## Migration Strategy

1. **Add interface parameters** to components that use filesystem, time, or file watching
2. **Use production implementations** by default (no behavior change)
3. **Switch to test implementations** in test files for speed and isolation
4. **Future-proof** for other deployment targets

Example component refactoring:
```go
// Before
type Registry struct {
    // direct dependencies
}

func (r *Registry) LoadManifest(path string) error {
    data, err := os.ReadFile(path) // Direct dependency
    // ...
}

// After  
type Registry struct {
    fs platform.FileSystem // Injected dependency
}

func NewRegistry(fs platform.FileSystem) *Registry {
    return &Registry{fs: fs}
}

func (r *Registry) LoadManifest(path string) error {
    data, err := r.fs.ReadFile(path) // Abstracted
    // ...
}
```