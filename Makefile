test:
	go test -json ./... "${@:2}" | go tool tparse -all

update:
	go test -json ./... "${@:2}" --update | go tool tparse -all

# watch:
# 	-while true; do ls ./**/*.{go,scm,ts} | entr -dcr sh -c "make test"; done

bench:
	go test -cpuprofile=cpu.out -bench=BenchmarkGenerate ./generate/...

# Benchmarks	go test -bench=. ./...	Find slow tests

# Capture CPU hotspots
profile:
	go test -bench=... -cpuprofile=cpu.out

flamegraph:
	go tool pprof -http=:8080 cpu.out # Visual analysis

coverage:
	go test -coverprofile=cover.out
