SHELL := /bin/bash

.PHONY: build test update watch bench profile flamegraph coverage clean lint format prepare-npm install-bindings windows windows-x64 windows-arm64

build:
	go build -o dist/cem .

windows-x64:
	mkdir -p dist/windows-x64
	podman build -t cem-windows-x64 --build-arg OUTPUT=dist/windows-x64 .

windows-arm64:
	echo "Note: windows-arm64 builds are not yet supported. Generating a stub script instead"
	echo '#!/usr/bin/env node\nconsole.error("cem does not yet support Windows on ARM64. Please use x64 or another operating system.");process.exit(1);' > cem.exe
	chmod +x cem.exe
	mkdir -p dist/cem-win32-arm64
	mv cem.exe dist/cem-win32-arm64/

# Convenience target to build both Windows variants
windows: windows-x64 windows-arm64

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
	go test -cpuprofile=cpu.out -bench=BenchmarkGenerate ./generate/...

profile:
	go test -bench=... -cpuprofile=cpu.out ./generate/...

flamegraph: profile
	go tool pprof -http=:8080 cpu.out # Visual analysis

coverage:
	go test -coverprofile=cover.out

clean:
	rm -rf dist/ cpu.out cover.out npm/bin artifacts platforms
