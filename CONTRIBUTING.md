# Contributing to cem

Thank you for your interest in contributing to `cem`! This guide will help you set up your environment for development, building, and testing. Please read carefully and reach out if you have questions.

## Getting Started

- Fork and clone the repo:  
  ```sh
  git clone https://github.com/bennypowers/cem.git
  cd cem
  ```
- Install [Go](https://golang.org/doc/install) (version 1.24 or newer recommended).
- Install [Node.js](https://nodejs.org/) (version 22 recommended) and [npm](https://www.npmjs.com/).

## Building

### Native Build (Linux/macOS)

To build the project for your local architecture:
```sh
make build
```
The binary will be output to `dist/cem`.

### Local Windows Build Using Podman

You can cross-compile Windows binaries from any OS using Podman (or Docker). This will output `cem-windows-x64.exe` and `cem-windows-arm64.exe` in your project root.

1. [Install Podman](https://podman.io/docs/installation) for your platform.
2. Run:
   ```sh
   make windows
   ```
   This will build both Windows x64 and arm64 executables using the parameterized Containerfile.

**Tip:**  
You can build a specific Windows architecture:
- x64: `make windows-x64`
- arm64: `make windows-arm64`

> Note: These builds are cross-compiled and cannot be run directly on macOS or Linux. Test on a Windows machine or VM if needed.

## Testing

Run Go tests:
```sh
make test
```

### Interpreting Test and Benchmark Results

- **Test Output:**  
  `make test` uses `go tool tparse` for readable, colorized test output.

- **Benchmark Output:**  
  `make bench` runs the main generator benchmark and produces a CPU profile (`cpu.out`).  
  `make benchmem` adds allocation stats (ns/op, B/op, allocs/op).

- **Profile Analysis:**  
  `make profile` and `make flamegraph` let you visualize CPU usage (`go tool pprof` at http://localhost:8080).

- **Coverage:**  
  `make coverage` generates a coverage profile;  
  `make coverage-html` opens a browsable code coverage report.

- **Race Detection:**  
  `make test-race` runs tests with the race detector enabled (highly recommended when working with concurrency and pooling).

**Tips:**  
- For benchmarks, look for improvements in `ns/op` (time per operation), `B/op` (bytes allocated), and `allocs/op` (number of allocations).
- For test coverage, aim for high statement coverage and add more tests for error/edge cases.
- Use `go tool pprof cpu.out` or the `flamegraph` target to spot performance hot spots.

## Continuous Integration (CI)

- All PRs are built and tested via GitHub Actions.
- The CI will cross-compile for Linux, macOS, and Windows (x64 and arm64), package npm binaries for each platform, and check that npm packaging works.
- You can see the exact build matrices in `.github/workflows/release.yml` and `.github/workflows/test-build.yml`.

## NPM Packaging

The project produces platform-specific npm packages. Platform detection and binary installation are managed by scripts in the `npm` directory. You can test npm packaging locally with:
```sh
node scripts/gen-platform-package-jsons.js
```
and validate the package with:
```sh
cd platforms/cem-<platform>-<arch>
npm pack --dry-run
```

## Code Formatting & Linting

- Go:  
  ```sh
  make format
  make lint
  ```

## Submitting Changes

- Open a pull request from your fork.
- Ensure all CI checks pass.
- Provide a clear description of your changes.

---

Thank you for contributing!
