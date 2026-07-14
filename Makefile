RUNS ?= 4
SHELL := /bin/bash
CONTRIBUTING_PATH = docs/content/docs/contributing.md
WINDOWS_CC_IMAGE := cem-windows-cc-image

# Extract version from goals if present (e.g., "make release v0.6.6" or "make release patch")
VERSION ?= $(filter v% patch minor major,$(MAKECMDGOALS))

# Use Go 1.25 toolchain automatically with JSON v2 experiment
export GOEXPERIMENT := jsonv2

# Workaround for Gentoo Linux "hole in findfunctab" error with race detector
# See: https://bugs.gentoo.org/961618
# Gentoo's Go build has issues with the race detector and internal linker.
# Using external linker resolves the issue.
ifeq ($(shell test -f /etc/gentoo-release && echo yes),yes)
    RACE_LDFLAGS := -ldflags="-linkmode=external"
else
    RACE_LDFLAGS :=
endif

.DEFAULT_GOAL := build

LDFLAGS := $(shell ./scripts/ldflags.sh)

# Source file lists for dependency tracking
GO_SOURCES := $(shell find . -name '*.go' -not -path '*/node_modules/*' -not -name '*_test.go' -not -name 'generate_elements.go') go.mod go.sum
ELEMENT_SOURCES := $(shell find serve/elements -name '*.ts' -o -name '*.css' 2>/dev/null)

## npm dependency targets (file-based, skip if up-to-date)

serve/elements/node_modules/.package-lock.json: serve/elements/package-lock.json
	cd serve/elements && npm ci
	@touch $@

serve/node_modules/.package-lock.json: serve/package-lock.json
	cd serve && npm ci
	@touch $@

.PHONY: install-elements install-frontend
install-elements: serve/elements/node_modules/.package-lock.json
install-frontend: serve/node_modules/.package-lock.json

## Code generation (stamp file tracks go generate output)

.generate.stamp: serve/elements/node_modules/.package-lock.json $(ELEMENT_SOURCES) serve/generate_elements.go
	go generate ./...
	@touch $@

.PHONY: generate install-bindings
generate: .generate.stamp
install-bindings: .generate.stamp

## Build targets (file-based)

dist/cem: .generate.stamp $(GO_SOURCES)
	@mkdir -p dist
	go build -ldflags="$$(./scripts/ldflags.sh)" -o $@ .

.PHONY: build
build: dist/cem

dev-serve: generate
	@mkdir -p dist
	go build -tags cemdev -ldflags="$(shell ./scripts/ldflags.sh)" -o dist/cem .

# NOTE: this is a non-traditional install target, which installs to ~/.local/bin/
# It's mostly intended for local development, not for distribution
.PHONY: install
install: dist/cem
	mkdir -p ~/.local/bin/
	cp dist/cem ~/.local/bin/

## Cross-platform build targets
## Output to dist/bin/<binary>-<platform>[.exe] per bennypowers/go-release-workflows contract

dist/bin/cem-linux-x64: .generate.stamp $(GO_SOURCES)
	@mkdir -p dist/bin
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC=gcc \
		go build -ldflags="$(LDFLAGS)" -o $@ .

dist/bin/cem-linux-arm64: .generate.stamp $(GO_SOURCES)
	@mkdir -p dist/bin
	CGO_ENABLED=1 GOOS=linux GOARCH=arm64 CC=aarch64-linux-gnu-gcc \
		go build -ldflags="$(LDFLAGS)" -o $@ .

# Darwin targets (must run on macOS)
# Explicit -arch flags ensure correct architecture when cross-compiling on macOS
dist/bin/cem-darwin-x64: .generate.stamp $(GO_SOURCES)
	@mkdir -p dist/bin
	CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 \
		CC="clang -arch x86_64" \
		CGO_CFLAGS="-arch x86_64" CGO_LDFLAGS="-arch x86_64" \
		go build -ldflags="$(LDFLAGS)" -o $@ .

dist/bin/cem-darwin-arm64: .generate.stamp $(GO_SOURCES)
	@mkdir -p dist/bin
	CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 \
		CC="clang -arch arm64" \
		CGO_CFLAGS="-arch arm64" CGO_LDFLAGS="-arch arm64" \
		go build -ldflags="$(LDFLAGS)" -o $@ .

# Windows targets (use shared Containerfile.windows via podman)
SHARED_WINDOWS_CC_IMAGE := cem-shared-windows-cc

.PHONY: build-windows-cc-image rebuild-windows-cc-image build-shared-windows-image
build-windows-cc-image:
	@if ! podman image exists $(WINDOWS_CC_IMAGE); then \
		echo "Building image..."; \
		podman build -t $(WINDOWS_CC_IMAGE) . ; \
	else \
		echo "Image $(WINDOWS_CC_IMAGE) already exists, skipping build."; \
	fi

rebuild-windows-cc-image:
	podman build --no-cache -t $(WINDOWS_CC_IMAGE) .

build-shared-windows-image:
	@if [ ! -f Containerfile.windows ]; then \
		echo "Error: Containerfile.windows not found. Run setup-windows-build action first."; \
		exit 1; \
	fi
	@if ! podman image exists $(SHARED_WINDOWS_CC_IMAGE); then \
		echo "Building shared Windows cross-compile image..."; \
		podman build -t $(SHARED_WINDOWS_CC_IMAGE) -f Containerfile.windows . ; \
	else \
		echo "Image $(SHARED_WINDOWS_CC_IMAGE) already exists, skipping build."; \
	fi

dist/bin/cem-win32-x64.exe: .generate.stamp $(GO_SOURCES) | build-shared-windows-image
	@mkdir -p dist/bin
	podman run --rm -v $(PWD):/app:Z -w /app \
		-e GOARCH=amd64 -e BINARY_NAME=cem -e GOEXPERIMENT=$(GOEXPERIMENT) $(SHARED_WINDOWS_CC_IMAGE)
	@mv dist/bin/cem-windows-amd64.exe $@

dist/bin/cem-win32-arm64.exe: .generate.stamp $(GO_SOURCES) | build-shared-windows-image
	@mkdir -p dist/bin
	podman run --rm -v $(PWD):/app:Z -w /app \
		-e GOARCH=arm64 -e BINARY_NAME=cem -e GOEXPERIMENT=$(GOEXPERIMENT) $(SHARED_WINDOWS_CC_IMAGE)
	@mv dist/bin/cem-windows-arm64.exe $@

# Platform aliases (satisfy go-release-workflows Makefile contract)
.PHONY: linux-x64 linux-arm64 darwin-x64 darwin-arm64 win32-x64 win32-arm64
linux-x64: dist/bin/cem-linux-x64
linux-arm64: dist/bin/cem-linux-arm64
darwin-x64: dist/bin/cem-darwin-x64
darwin-arm64: dist/bin/cem-darwin-arm64
win32-x64: dist/bin/cem-win32-x64.exe
win32-arm64: dist/bin/cem-win32-arm64.exe

# Legacy Windows targets (local podman-based cross-compile)
.PHONY: windows windows-x64 windows-arm64
windows: windows-x64 windows-arm64

windows-x64: build-windows-cc-image
	mkdir -p dist
	podman run --rm -v $(PWD):/app:Z -w /app -e GOARCH=amd64 -e GOEXPERIMENT=$(GOEXPERIMENT) $(WINDOWS_CC_IMAGE)

windows-arm64: build-windows-cc-image
	mkdir -p dist
	podman run --rm -v $(PWD):/app:Z -w /app -e GOARCH=arm64 -e GOEXPERIMENT=$(GOEXPERIMENT) $(WINDOWS_CC_IMAGE)

.PHONY: all
all: windows

## Testing

.PHONY: test test-unit test-e2e test-frontend test-frontend-watch test-frontend-update test-pkg update
test-unit: .generate.stamp
	gotestsum --rerun-fails --rerun-fails-max-failures=3 --packages=./... -- -race $(RACE_LDFLAGS)

test-e2e: .generate.stamp
	gotestsum --rerun-fails --rerun-fails-max-failures=3 --packages=./cmd/ -- -race $(RACE_LDFLAGS) -tags=e2e

test-frontend: serve/node_modules/.package-lock.json dist/cem
	@set -e; \
	PIDFILE=$$(mktemp); \
	LOGFILE=$$(mktemp); \
	cleanup() { \
		if [ -f "$$PIDFILE" ] && [ -s "$$PIDFILE" ]; then \
			PID=$$(cat "$$PIDFILE"); \
			if [ -n "$$PID" ] && kill -0 "$$PID" 2>/dev/null; then \
				if ps -p "$$PID" -o comm= 2>/dev/null | grep -q '^cem$$'; then \
					echo "Stopping cem serve (PID $$PID)..."; \
					kill "$$PID" 2>/dev/null || true; \
					sleep 0.5; \
					kill -0 "$$PID" 2>/dev/null && kill -9 "$$PID" 2>/dev/null || true; \
				fi; \
			fi; \
		fi; \
		rm -f "$$PIDFILE" "$$LOGFILE"; \
	}; \
	trap cleanup EXIT INT TERM; \
	echo "Starting cem serve on dynamic port for tests..."; \
	(cd serve/testdata/demo-routing && exec ../../../dist/cem serve --port 0) > "$$LOGFILE" 2>&1 & \
	echo $$! > "$$PIDFILE"; \
	echo "Waiting for server to be ready..."; \
	PID=$$(cat "$$PIDFILE"); \
	TIMEOUT=30; \
	ELAPSED=0; \
	PORT=""; \
	while [ $$ELAPSED -lt $$TIMEOUT ]; do \
		if ! kill -0 "$$PID" 2>/dev/null; then \
			echo "ERROR: cem serve process (PID $$PID) exited prematurely"; \
			echo "Server log:"; \
			cat "$$LOGFILE"; \
			exit 1; \
		fi; \
		if [ -z "$$PORT" ]; then \
			if command -v ss >/dev/null 2>&1; then \
				PORT=$$(ss -tlnp 2>/dev/null | awk "/pid=$$PID/"' {match($$4, /:([0-9]+)$$/, a); if (a[1]) print a[1]}' | head -1); \
			elif command -v lsof >/dev/null 2>&1; then \
				PORT=$$(lsof -iTCP -sTCP:LISTEN -nP -a -p "$$PID" 2>/dev/null | awk 'NR>1 {match($$9, /:([0-9]+)$$/, a); if (a[1]) print a[1]}' | head -1); \
			fi; \
		fi; \
		if [ -n "$$PORT" ] && \
		   curl -sf "http://localhost:$$PORT/__cem/api/health" >/dev/null 2>&1; then \
			echo "Server is ready on port $$PORT."; \
			break; \
		fi; \
		sleep 1; \
		ELAPSED=$$((ELAPSED + 1)); \
	done; \
	if [ $$ELAPSED -ge $$TIMEOUT ]; then \
		echo "ERROR: Server failed to start within $${TIMEOUT}s"; \
		echo "Server log:"; \
		cat "$$LOGFILE"; \
		exit 1; \
	fi; \
	echo "Running frontend tests..."; \
	cd serve && CEM_TEST_PORT=$$PORT npm test; \
	TEST_EXIT=$$?; \
	exit $$TEST_EXIT

test-frontend-watch: serve/node_modules/.package-lock.json
	cd serve && npm run test:watch

test-frontend-update: serve/node_modules/.package-lock.json
	cd serve && npm run test:update

test: test-unit test-e2e test-frontend

# Flexible test target that accepts TEST_ARGS for filtering
# Usage: make test-pkg TEST_ARGS="-v ./lsp/methods/textDocument/definition/ -run TestDefinition"
test-pkg:
	@if [ -z "$(TEST_ARGS)" ]; then \
		echo "Usage: make test-pkg TEST_ARGS=\"-v ./path/to/package/ -run TestName\""; \
		exit 1; \
	fi
	go test -race $(RACE_LDFLAGS) $(TEST_ARGS)

update:
	go test -race $(RACE_LDFLAGS) -json ./... --update | go tool tparse -all
	make test-frontend-update

## Code quality

.PHONY: lint format
lint:
	golangci-lint run

format:
	gofmt -s -w .
	goimports -w .

## Watch

.PHONY: watch
watch:
	while true; do \
		find . -type f \( -name "*.go" -o -name "*.scm" -o -name "*.ts" \) | entr -d sh -c 'make test || true'; \
	done

## Benchmarks

.PHONY: bench bench-generate bench-lookup bench-lsp bench-lsp-cem bench-lsp-wc setup-wc-toolkit
bench: bench-generate bench-lsp

bench-generate: .generate.stamp
	go test -v -cpuprofile=cpu.out -bench=BenchmarkGenerate -run=^$$ ./generate/

bench-lookup:
	@echo "=== Running Attribute Lookup Benchmarks ==="
	go test -bench=BenchmarkAttributeLookup -benchmem -count=5 -run=^$$ ./manifest/
	@echo ""
	@echo "=== Running Export Lookup Benchmarks ==="
	go test -bench=BenchmarkExportLookup -benchmem -count=5 -run=^$$ ./manifest/
	@echo ""
	@echo "=== Running Renderable Creation Benchmarks ==="
	go test -bench=BenchmarkRenderableCreation -benchmem -count=5 -run=^$$ ./manifest/

setup-wc-toolkit:
	@cd lsp/benchmark && ./setup_wc_toolkit.sh

bench-lsp: dist/cem setup-wc-toolkit
	@cd lsp/benchmark && \
	BENCHMARK_COMPARISON_MODE=1 nvim --headless --clean -u configs/cem-minimal.lua -l run_modular_benchmark.lua && \
	if [ -f bin/wc-language-server ]; then \
		BENCHMARK_COMPARISON_MODE=1 nvim --headless --clean -u configs/wc-toolkit-minimal.lua -l run_modular_benchmark.lua; \
	else \
		echo "⚠️  Skipping wc-toolkit benchmarks: server not found at lsp/benchmark/bin/wc-language-server" && \
		echo "   Run 'make setup-wc-toolkit' or './lsp/benchmark/setup_wc_toolkit.sh' to install."; \
	fi && \
	nvim --headless --clean -l compare_results.lua

bench-lsp-cem: dist/cem
	@echo "Running benchmarks for cem LSP server only..."
	@cd lsp/benchmark && nvim --headless --clean -u configs/cem-minimal.lua -l run_modular_benchmark.lua

bench-lsp-wc: dist/cem setup-wc-toolkit
	@echo "Running benchmarks for wc-toolkit LSP server only..."
	@if [ ! -f lsp/benchmark/bin/wc-language-server ]; then \
		echo "✗ Error: wc-toolkit server not found at lsp/benchmark/bin/wc-language-server" && \
		echo "  Run 'make setup-wc-toolkit' or './lsp/benchmark/setup_wc_toolkit.sh' to install." && \
		exit 1; \
	fi
	@cd lsp/benchmark && nvim --headless --clean -u configs/wc-toolkit-minimal.lua -l run_modular_benchmark.lua

## Profiling

cpu.out: .generate.stamp
	go test -bench=... -run=^$$ -cpuprofile=$@ ./generate/

.PHONY: profile flamegraph
profile: cpu.out

flamegraph: cpu.out
	go tool pprof -http=:8080 cpu.out

## Coverage

.PHONY: coverage show-coverage
coverage:
	@echo "Running unit tests with coverage..."
	@mkdir -p coverage/unit
	go test ./... -cover -args -test.gocoverdir="$(PWD)/coverage/unit"
	@echo "Running e2e tests with coverage..."
	go test -count=1 -tags=e2e ./cmd/
	@echo "Merging coverage data..."
	@mkdir -p coverage/merged
	@COVDIRS=./coverage/unit; \
	if [ -d ./cmd/coverage.e2e ]; then \
		COVDIRS="$$COVDIRS,./cmd/coverage.e2e"; \
	fi; \
	go tool covdata merge -i=$$COVDIRS -o coverage/merged
	@echo "Converting to text format..."
	go tool covdata textfmt -i=coverage/merged -o cover.out
	@echo "Coverage report:"
	@go tool cover -func=cover.out | tail -20
	@echo ""
	@echo "Total coverage:"
	@go tool cover -func=cover.out | grep total
	@echo "To view the full report, run 'make show-coverage'"

show-coverage:
	go tool cover -html=cover.out

## Clean

.PHONY: clean
clean:
	rm -rf dist/ cpu.out cover.out coverage/ cmd/coverage.e2e/ artifacts platforms existing-binaries .generate.stamp

## Git hooks

.PHONY: install-git-hooks
install-git-hooks:
	@echo "Installing git hooks..."
	@git config core.hooksPath .githooks
	@echo "Git hooks installed successfully!"
	@echo "  - pre-push: checks formatting and runs linting before push"
	@echo ""
	@echo "To bypass hooks, use --no-verify flag (e.g., 'git push --no-verify')"

## Data updates

.PHONY: update-html-attributes
update-html-attributes:
	@echo "Updating HTML global attributes from MDN browser-compat-data..."
	@mkdir -p lsp/methods/textDocument/publishDiagnostics/data
	@curl -s "https://raw.githubusercontent.com/mdn/browser-compat-data/main/html/global_attributes.json" -o lsp/methods/textDocument/publishDiagnostics/data/global_attributes.json
	@echo "HTML global attributes updated successfully!"

## Docs CI

.PHONY: docs-ci
docs-ci: update-html-attributes dist/cem
	@echo "Running generate benchmarks with $(RUNS) runs"
	./scripts/benchmark.sh $(RUNS)
	@echo "Running LSP benchmarks..."
	make bench-lsp
	@echo "Combining LSP benchmark results..."
	@cd lsp/benchmark && nvim --headless --clean -l combine_results.lua
	cp -f "$(CONTRIBUTING_PATH)" /tmp/cem-contributing.md
	cat CONTRIBUTING.md >> "$(CONTRIBUTING_PATH)"
	hugo mod clean
	hugo --gc --minify --source docs
	mv /tmp/cem-contributing.md "$(CONTRIBUTING_PATH)"

## VSCode extension

extensions/vscode/out/client/extension.js: extensions/vscode/client/package.json extensions/vscode/build.js
	@echo "Building VSCode extension..."
	@cd extensions/vscode/client && npm install
	@cd extensions/vscode && node build.js
	@echo "VSCode extension built successfully"

.PHONY: vscode-build vscode-package vscode-publish
vscode-build: extensions/vscode/out/client/extension.js

vscode-package: dist/cem
	@echo "Packaging VSCode extension..."
	@echo "Copying CEM binary for current platform..."
	@mkdir -p extensions/vscode/dist/bin
	@if [ "$$(uname -s)" = "Linux" ] && [ "$$(uname -m)" = "x86_64" ]; then \
		cp dist/cem extensions/vscode/dist/bin/cem-x86_64-unknown-linux-gnu; \
	elif [ "$$(uname -s)" = "Linux" ] && [ "$$(uname -m)" = "aarch64" ]; then \
		cp dist/cem extensions/vscode/dist/bin/cem-aarch64-unknown-linux-gnu; \
	elif [ "$$(uname -s)" = "Darwin" ] && [ "$$(uname -m)" = "x86_64" ]; then \
		cp dist/cem extensions/vscode/dist/bin/cem-x86_64-apple-darwin; \
	elif [ "$$(uname -s)" = "Darwin" ] && [ "$$(uname -m)" = "arm64" ]; then \
		cp dist/cem extensions/vscode/dist/bin/cem-aarch64-apple-darwin; \
	else \
		echo "Warning: Unsupported platform $$(uname -s)-$$(uname -m), copying as generic binary"; \
		cp dist/cem extensions/vscode/dist/bin/cem-$$(uname -m)-unknown-$$(uname -s | tr '[:upper:]' '[:lower:]')-gnu; \
	fi
	@chmod +x extensions/vscode/dist/bin/cem-*
	@cd extensions/vscode && npm run build
	@echo "VSCode extension packaged successfully"

vscode-publish:
	@echo "Manually building and publishing VSCode extension for latest tag..."
	@LATEST_TAG=$$(git describe --tags --abbrev=0) && \
	echo "Using tag: $$LATEST_TAG" && \
	mkdir -p extensions/vscode/dist/bin && \
	echo "Downloading binaries from GitHub release..." && \
	gh release download $$LATEST_TAG \
		--pattern "cem-linux-x64" \
		--pattern "cem-linux-arm64" \
		--pattern "cem-darwin-x64" \
		--pattern "cem-darwin-arm64" \
		--pattern "cem-win32-x64.exe" \
		--pattern "cem-win32-arm64.exe" \
		--dir temp-binaries && \
	echo "Renaming binaries..." && \
	mv temp-binaries/cem-linux-x64 extensions/vscode/dist/bin/cem-x86_64-unknown-linux-gnu && \
	mv temp-binaries/cem-linux-arm64 extensions/vscode/dist/bin/cem-aarch64-unknown-linux-gnu && \
	mv temp-binaries/cem-darwin-x64 extensions/vscode/dist/bin/cem-x86_64-apple-darwin && \
	mv temp-binaries/cem-darwin-arm64 extensions/vscode/dist/bin/cem-aarch64-apple-darwin && \
	mv temp-binaries/cem-win32-x64.exe extensions/vscode/dist/bin/cem-x86_64-pc-windows-msvc.exe && \
	mv temp-binaries/cem-win32-arm64.exe extensions/vscode/dist/bin/cem-aarch64-pc-windows-msvc.exe && \
	chmod +x extensions/vscode/dist/bin/cem-* && \
	rm -rf temp-binaries && \
	cd extensions/vscode && \
	VERSION=$$(echo $$LATEST_TAG | sed 's/^v//') && \
	npm version $$VERSION --no-git-tag-version --allow-same-version && \
	npm run build && \
	npm run publish

## Verify npm packaging structure before release

.PHONY: verify-npm
verify-npm: dist/cem
	@./scripts/verify-npm-packaging.sh

## Make version targets (v*) and bump types no-ops for "make release" syntax
v%:
	@:

.PHONY: patch minor major release
patch minor major:
	@:

## Release (creates version commit, pushes it, then uses gh to tag and create release)
release:
	@if [ -z "$(VERSION)" ]; then \
		echo "Error: VERSION or bump type is required"; \
		echo "Usage: make release <version|patch|minor|major>"; \
		echo "  make release v0.6.6   - Release explicit version"; \
		echo "  make release patch    - Bump patch version (0.0.x)"; \
		echo "  make release minor    - Bump minor version (0.x.0)"; \
		echo "  make release major    - Bump major version (x.0.0)"; \
		exit 1; \
	fi
	@./scripts/release.sh $(VERSION)

## Examples

.PHONY: examples-analyze examples-verify examples-clean
examples-analyze: dist/cem
	@echo "Generating manifests for all examples..."
	@for dir in examples/*/; do \
		echo "  Generating $$dir"; \
		(cd "$$dir" && ../../dist/cem generate) || exit 1; \
	done
	@echo "✓ All example manifests generated"

examples-verify: examples-analyze
	@echo "Verifying example manifests are up-to-date..."
	@if git diff --exit-code examples/*/custom-elements.json; then \
		echo "✓ All example manifests are up-to-date"; \
	else \
		echo "✗ Example manifests are out of date. Run 'make examples-analyze' and commit the changes."; \
		exit 1; \
	fi

examples-clean:
	@echo "Cleaning example outputs..."
	@rm -f examples/*/custom-elements.json
	@echo "✓ Example outputs cleaned"
