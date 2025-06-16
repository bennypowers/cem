SHELL := /bin/bash

.PHONY: build build-all test update watch bench profile flamegraph coverage clean lint format

PLATFORMS = linux/amd64 linux/arm64 darwin/amd64 darwin/arm64 windows/amd64 windows/arm64

build: $(wildcard *.{go,scm})
	go build

build-all:
	mkdir -p dist/npm/bin
	set -euo pipefail; \
	for platform in $(PLATFORMS); do \
		GOOS=$${platform%/*}; \
		GOARCH=$${platform#*/}; \
		out="dist/npm/bin/cem-$${GOOS}-$${GOARCH}"; \
		if [ "$${GOOS}" = "windows" ]; then out="$${out}.exe"; fi; \
		GOOS=$${GOOS} GOARCH=$${GOARCH} CGO_ENABLED=0 go build -o "$${out}" ./cmd/cem; \
		echo "Built $$out"; \
	done

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
	go test -cpuprofile=cpu.out -bench=BenchmarkGenerate ./generate/...

profile:
	go test -bench=... -cpuprofile=cpu.out ./generate/...

flamegraph: profile
	go tool pprof -http=:8080 cpu.out # Visual analysis

coverage:
	go test -coverprofile=cover.out

clean:
	rm -rf dist/ cpu.out cover.out
