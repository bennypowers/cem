SHELL := /bin/bash

.PHONY: build test update watch bench profile flamegraph coverage clean lint format prepare-npm install-bindings windows windows-x64 windows-arm64

build:
	go build -o dist/cem .

windows-x64:
	podman build --build-arg GOARCH=amd64 --build-arg CC=x86_64-w64-mingw32-gcc -t cem-windows-x64 .
	podman run --rm -v $(pwd):/output:Z cem-windows-x64 cp cem.exe /output/cem-windows-x64.exe

windows-arm64:
	podman build --build-arg GOARCH=arm64 --build-arg CC=aarch64-w64-mingw32-gcc -t cem-windows-arm64 .
	podman run --rm -v $(pwd):/output:Z cem-windows-arm64 cp cem.exe /output/cem-windows-arm64.exe

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
	rm -rf dist/ cpu.out cover.out npm/bin artifacts
