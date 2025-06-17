SHELL := /bin/bash

.PHONY: build test update watch bench profile flamegraph coverage clean lint format prepare-npm install-bindings

build:
	go build -o dist/cem .

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
