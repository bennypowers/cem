RUNS ?= 4
SHELL := /bin/bash
CONTRIBUTING_PATH = docs/content/docs/contributing.md
WINDOWS_CC_IMAGE := cem-windows-cc-image

.PHONY: build test test-unit test-e2e update watch bench profile flamegraph coverage show-coverage clean lint format prepare-npm install-bindings windows windows-x64 windows-arm64 build-windows-cc-image rebuild-windows-cc-image install-git-hooks update-html-attributes vscode-build vscode-package

# NOTE: this is a non-traditional install target, which installs to ~/.local/bin/
# It's mostly intended for local development, not for distribution
install: build
	mkdir -p ~/.local/bin/
	cp dist/cem ~/.local/bin/

all: windows

clean:
	rm -rf dist/ cpu.out cover.out artifacts platforms

build:
	@mkdir -p dist
	go build -ldflags="$(shell ./scripts/ldflags.sh)" -o dist/cem .

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

## Test, lint, etc
install-bindings:
	go generate ./...

test-unit:
	# Run race tests excluding internal/platform due to Go race detector "hole in findfunctab" issue
	# The issue occurs with channel operations in FSNotifyFileWatcher and mock implementations
	# Race-unsafe tests can be run manually: go test -tags=race_unsafe ./internal/platform/
	gotestsum -- -race $$(go list ./... | grep -v internal/platform) && go test -tags="!race_unsafe" ./internal/platform/

test-e2e:
	gotestsum -- -race -tags=e2e ./cmd/

test: test-unit test-e2e

update:
	go test -race -json ./... --update | go tool tparse -all

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

profile:
	go test -bench=... -run=^$ -cpuprofile=cpu.out ./generate/

flamegraph: profile
	go tool pprof -http=:8080 cpu.out # Visual analysis

coverage:
	@echo "Running unit tests with coverage..."
	go test -coverprofile=cover.unit.out -covermode=atomic ./...
	@echo "Running e2e tests with coverage..."
	go test -count=1 -tags=e2e ./cmd/
	@echo "Merging coverage profiles..."
	@echo "mode: atomic" > cover.out
	@grep -h -v "^mode:" cover.unit.out cmd/coverage.e2e.out >> cover.out
	@rm cover.unit.out cmd/coverage.e2e.out
	@echo "Coverage report:"
	@go tool cover -func=cover.out
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

vscode-package:
	@echo "Packaging VSCode extension..."
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
