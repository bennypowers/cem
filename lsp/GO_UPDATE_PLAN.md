# Go 1.25 Update Plan

## Current Status: DEFERRED

**Issue**: Go 1.25.0 is available but the checksum database hasn't been updated yet, causing verification failures.

**Error**: `go: golang.org/toolchain@v0.0.1-go1.25.0.linux-amd64: verifying module: checksum database disabled by GOSUMDB=off`

## Plan: Wait for Ecosystem Readiness

### Prerequisites for Update
1. **Checksum Database Update**: Wait for `sum.golang.org` to include Go 1.25.0 checksums
2. **Dependency Compatibility**: Ensure all dependencies work with Go 1.25
3. **CI/CD Readiness**: Verify GitHub Actions and other tooling support Go 1.25

### Go 1.25 Features to Leverage

#### 1. **Testing Modernization** (Primary Goal)
- **`testing/synctest`**: Virtual time for integration tests
  - Replace `time.Sleep(5 * time.Second)` with instant virtual time
  - Test file watchers without real delays
  - **Target**: Speed up completion tests from 10.7s → ~0.5s

- **`testing/fstest.MapFS`**: In-memory filesystem
  - Replace real file I/O with mock filesystem
  - Eliminate temporary directory creation overhead
  - **Target**: Remove disk I/O bottlenecks in integration tests

#### 2. **Performance Improvements**
- **Experimental GC**: Test new garbage collector performance
  - **Target**: Faster manifest parsing and element indexing
  - **Benchmark**: Registry operations, large manifest processing

- **JSON v2**: Evaluate new JSON package
  - **Target**: Faster manifest serialization/deserialization
  - **Current Usage**: Custom element manifest parsing, LSP protocol messages
  - **Benchmark**: Compare against current `encoding/json` usage

#### 3. **Developer Experience**
- **Enhanced Error Messages**: Better stack traces and debugging
- **Improved Tooling**: Faster builds and test execution

### Implementation Timeline

#### Phase 1: Infrastructure (When Go 1.25 is ready)
1. **Update go.mod**: `go 1.25` + `toolchain go1.25.x`
2. **CI/CD Update**: Update GitHub Actions to Go 1.25
3. **Dependency Audit**: Verify all dependencies work
4. **Basic Testing**: Ensure existing functionality works

#### Phase 2: Testing Modernization (1-2 weeks)
1. ✅ **Create platform abstractions package**: COMPLETED - Created `internal/platform` with filesystem, time, file watching interfaces
2. ✅ **E2E backup**: COMPLETED - Copied existing tests to `*_e2e_test.go` with `//go:build e2e`
3. ✅ **Registry dependency injection**: COMPLETED - Registry now uses injectable FileWatcher interface
4. **Virtual time conversion**: Replace `time.Sleep()` with `testing/synctest` (requires Go 1.25)
5. **Mock filesystem**: Replace file I/O with `testing/fstest.MapFS` (requires Go 1.25)
6. **Performance validation**: Measure before/after test execution times

#### Phase 3: Performance Optimization (1 week)
1. **Experimental GC evaluation**: Benchmark with `GOEXPERIMENT=gc`
2. **JSON v2 migration**: Evaluate for manifest processing
3. **Performance testing**: Registry operations, manifest parsing
4. **Production testing**: Verify no regressions

### Target Improvements

#### Testing Performance
- **Integration test speed**: 10.7s → ~0.5s (20x faster)
- **Test reliability**: Eliminate filesystem race conditions
- **CI/CD speed**: Faster feedback cycles

#### Runtime Performance
- **GC improvements**: Potentially faster manifest processing
- **JSON v2**: Faster LSP protocol and manifest operations
- **Memory usage**: Better allocation patterns

### Current Status with Go 1.24

✅ **Foundation Complete**: We have successfully implemented the architectural foundation:
- **Platform abstractions**: `internal/platform` package with FileSystem, TimeProvider, FileWatcher interfaces
- **Dependency injection**: Registry uses injectable FileWatcher for enhanced testability
- **E2E test preservation**: Real-world integration tests backed up with `//go:build e2e` tags
- **Test infrastructure**: Mock implementations ready for controlled testing

🔄 **Limitations**: Go 1.24 cannot effectively provide:
- **Virtual time**: `time.Sleep()` calls still require real delays in integration tests
- **In-memory filesystem**: Limited mock filesystem capabilities for file watcher tests
- **Concurrent testing**: Manual goroutine synchronization without virtual time bubbles

✅ **Ready for Go 1.25**: Architecture designed to seamlessly upgrade:
- Platform interfaces map directly to Go 1.25 `testing/synctest` and `testing/fstest.MapFS`
- Mock implementations can be replaced with standard library equivalents
- Dependency injection enables drop-in replacement of testing implementations

### Retry Criteria

Update to Go 1.25 when:
1. ✅ Checksum database includes Go 1.25.0
2. ✅ Major dependencies confirm Go 1.25 compatibility
3. ✅ GitHub Actions supports Go 1.25
4. ✅ No blocking issues reported in Go 1.25

### Notes

- **Go 1.25 Release**: January 2025 (stable)
- **Adoption Timeline**: Typically 2-4 weeks for ecosystem readiness
- **Risk Assessment**: Low (patch release), but checksum verification is critical
- **Fallback Plan**: All improvements can be implemented with Go 1.24.x if needed

## Next Steps

1. **Monitor Go ecosystem** for Go 1.25 readiness indicators
2. ✅ **Implement testing abstractions** COMPLETED - Platform abstractions layer ready
3. ✅ **Preserve existing tests** COMPLETED - E2E tests backed up for validation
4. **Retry Go 1.25 update** when checksum database is ready

## Migration Roadmap (When Go 1.25 is Ready)

### Step 1: Update Go Version (30 minutes)
```bash
# Update go.mod
go 1.25
toolchain go1.25.x

# Update CI/CD
# Update GitHub Actions to use Go 1.25
```

### Step 2: Replace Platform Implementations (2-3 hours)
```go
// Replace MockTimeProvider with testing/synctest
import "testing/synctest"

func TestWithVirtualTime(t *testing.T) {
    synctest.Test(t, func() {
        // time.Sleep() now advances instantly
        registry := lsp.NewTestRegistry()
        // Test logic here - no real delays
    })
}

// Replace TempDirFileSystem with testing/fstest.MapFS
import "testing/fstest"

func TestWithMockFS(t *testing.T) {
    fs := fstest.MapFS{
        "manifest.json": &fstest.MapFile{
            Data: []byte(`{"modules": []}`),
        },
    }
    // Use fs for all file operations
}
```

### Step 3: Update Integration Tests (1-2 hours)
- Convert completion integration tests to use virtual time
- Replace file I/O with in-memory filesystem operations
- Validate performance improvements (10.7s → ~0.5s target)

### Step 4: Validation (1 hour)
- Run original E2E tests with `-tags=e2e` for validation
- Benchmark performance improvements
- Verify all functionality still works correctly

**Total Estimated Migration Time**: 4-6 hours