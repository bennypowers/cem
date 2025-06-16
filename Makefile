.PHONY: build test update watch bench profile flamegraph coverage clean lint format prepare-npm

PLATFORMS = linux/amd64 linux/arm64 darwin/amd64 darwin/arm64 windows/amd64 windows/arm64
BIN_DIR = npm/bin

build: $(wildcard *.{go,scm})
	go build

prepare-npm:
	@echo "Preparing npm package"
	@mkdir -p $(BIN_DIR)
	@cp artifacts/**/cem-* $(BIN_DIR) 2>/dev/null || true

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
