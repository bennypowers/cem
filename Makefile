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

.PHONY: all build test test-unit test-e2e test-frontend test-frontend-watch test-frontend-update install-frontend update watch bench bench-lookup profile flamegraph coverage show-coverage clean lint format prepare-npm generate install-bindings windows windows-x64 windows-arm64 build-windows-cc-image rebuild-windows-cc-image install-git-hooks update-html-attributes vscode-build vscode-package release patch minor major

build: generate
	@mkdir -p dist
	go build -ldflags="$(shell ./scripts/ldflags.sh)" -o dist/cem .

# NOTE: this is a non-traditional install target, which installs to ~/.local/bin/
# It's mostly intended for local development, not for distribution
install: build
	mkdir -p ~/.local/bin/
	cp dist/cem ~/.local/bin/

all: windows

clean:
	rm -rf dist/ cpu.out cover.out coverage/ cmd/coverage.e2e/ artifacts platforms

# Convenience target to build both Windows variants
windows: windows-x64 windows-arm64

# Build the Podman image only if it doesn't exist
build-windows-cc-image:
	@if ! podman image exists $(WINDOWS_CC_IMAGE); then \
		echo "Building image..."; \
		podman build -t $(WINDOWS_CC_IMAGE) . ; \
	else \
		echo "Image $(WINDOWS_CC_IMAGE) already exists, skipping build."; \
	fi

# Force rebuild of the image
rebuild-windows-cc-image:
	podman build --no-cache -t $(WINDOWS_CC_IMAGE) .

windows-x64: build-windows-cc-image
	mkdir -p dist
	podman run --rm -v $(PWD):/app:Z -w /app -e GOARCH=amd64 $(WINDOWS_CC_IMAGE)

windows-arm64: build-windows-cc-image
	mkdir -p dist
	podman run --rm -v $(PWD):/app:Z -w /app -e GOARCH=arm64 $(WINDOWS_CC_IMAGE)

## Code generation and dependencies
generate:
	go generate ./...

install-bindings: generate

test-unit: generate
	gotestsum -- -race $(RACE_LDFLAGS) ./...

test-e2e: generate
	gotestsum -- -race $(RACE_LDFLAGS) -tags=e2e ./cmd/

install-frontend:
	cd serve && npm ci

test-frontend: install-frontend build
	@set -e; \
	WORKDIR=$$(pwd); \
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
	echo "Starting cem serve on port 9876 for tests..."; \
	cd serve/testdata/demo-routing && ../../../dist/cem serve --port 9876 > "$$LOGFILE" 2>&1 & echo $$! > "$$PIDFILE"; \
	cd "$$WORKDIR"; \
	echo "Waiting for server to be ready..."; \
	TIMEOUT=30; \
	ELAPSED=0; \
	while [ $$ELAPSED -lt $$TIMEOUT ]; do \
		if nc -z localhost 9876 2>/dev/null || \
		   (command -v curl >/dev/null 2>&1 && curl -s http://localhost:9876 >/dev/null 2>&1); then \
			echo "Server is ready."; \
			break; \
		fi; \
		sleep 0.5; \
		ELAPSED=$$((ELAPSED + 1)); \
	done; \
	if [ $$ELAPSED -ge $$TIMEOUT ]; then \
		echo "ERROR: Server failed to start within $${TIMEOUT}s"; \
		echo "Server log:"; \
		cat "$$LOGFILE"; \
		exit 1; \
	fi; \
	echo "Running frontend tests..."; \
	cd serve && npm test; \
	TEST_EXIT=$$?; \
	exit $$TEST_EXIT

test-frontend-watch: install-frontend
	cd serve && npm run test:watch

test-frontend-update: install-frontend
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

lint:
	golangci-lint run

format:
	gofmt -s -w .
	goimports -w .

watch:
	while true; do \
		find . -type f \( -name "*.go" -o -name "*.scm" -o -name "*.ts" \) | entr -d sh -c 'make test || true'; \
	done

bench:
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

profile:
	go test -bench=... -run=^$ -cpuprofile=cpu.out ./generate/

flamegraph: profile
	go tool pprof -http=:8080 cpu.out # Visual analysis

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

install-git-hooks:
	@echo "Installing git hooks..."
	@mkdir -p .git/hooks
	@cp scripts/pre-commit .git/hooks/pre-commit
	@chmod +x .git/hooks/pre-commit
	@echo "Git hooks installed successfully!"
	@echo "The pre-commit hook will run 'go fmt' on staged .go files."

update-html-attributes:
	@echo "Updating HTML global attributes from MDN browser-compat-data..."
	@mkdir -p lsp/methods/textDocument/publishDiagnostics/data
	@curl -s "https://raw.githubusercontent.com/mdn/browser-compat-data/main/html/global_attributes.json" -o lsp/methods/textDocument/publishDiagnostics/data/global_attributes.json
	@echo "HTML global attributes updated successfully!"

docs-ci: update-html-attributes
	make build
	@echo "Running benchmarks with $(RUNS) runs"
	./scripts/benchmark.sh $(RUNS)
	cp -f "$(CONTRIBUTING_PATH)" /tmp/cem-contributing.md
	cat CONTRIBUTING.md >> "$(CONTRIBUTING_PATH)"
	hugo mod clean
	hugo --gc --minify --source docs
	mv /tmp/cem-contributing.md "$(CONTRIBUTING_PATH)"

vscode-build:
	@echo "Building VSCode extension..."
	@cd extensions/vscode/client && npm install
	@cd extensions/vscode && node build.js
	@echo "VSCode extension built successfully"

vscode-package: build
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
		--pattern "cem-linux-amd64" \
		--pattern "cem-linux-arm64" \
		--pattern "cem-darwin-amd64" \
		--pattern "cem-darwin-arm64" \
		--pattern "cem-windows-amd64" \
		--pattern "cem-windows-arm64" \
		--dir temp-binaries && \
	echo "Renaming binaries..." && \
	mv temp-binaries/cem-linux-amd64 extensions/vscode/dist/bin/cem-x86_64-unknown-linux-gnu && \
	mv temp-binaries/cem-linux-arm64 extensions/vscode/dist/bin/cem-aarch64-unknown-linux-gnu && \
	mv temp-binaries/cem-darwin-amd64 extensions/vscode/dist/bin/cem-x86_64-apple-darwin && \
	mv temp-binaries/cem-darwin-arm64 extensions/vscode/dist/bin/cem-aarch64-apple-darwin && \
	mv temp-binaries/cem-windows-amd64 extensions/vscode/dist/bin/cem-x86_64-pc-windows-msvc.exe && \
	mv temp-binaries/cem-windows-arm64 extensions/vscode/dist/bin/cem-aarch64-pc-windows-msvc.exe && \
	chmod +x extensions/vscode/dist/bin/cem-* && \
	rm -rf temp-binaries && \
	cd extensions/vscode && \
	VERSION=$$(echo $$LATEST_TAG | sed 's/^v//') && \
	npm version $$VERSION --no-git-tag-version --allow-same-version && \
	npm run build && \
	npm run publish

## Make version targets (v*) and bump types no-ops for "make release" syntax
v%:
	@:

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
