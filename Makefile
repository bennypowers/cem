SHELL := /bin/bash
WINDOWS_CC_IMAGE := cem-windows-cc-image

.PHONY: build test update watch bench profile flamegraph coverage clean lint format prepare-npm install-bindings windows windows-x64 windows-arm64 build-windows-cc-image rebuild-windows-cc-image

all: windows

clean:
	rm -rf dist/ cpu.out cover.out artifacts platforms

build:
	go build -o dist/cem .

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

test:
	go test -json ./... | go tool tparse -all

update:
	go test -json ./... --update | go tool tparse -all

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
	go test -coverprofile=cover.out
