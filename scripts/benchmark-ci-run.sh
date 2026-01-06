#!/usr/bin/env bash
mkdir -p docs/data
bench_stdout="${1}_bench.txt"
# prewarm the benchmarks
make bench-generate
# run for realsies
start=$(date +%s.%N)
if ! make bench-generate > $bench_stdout 2>&1; then
  cat $bench_stdout
  echo "::error::Benchmarking (make bench-generate) failed on $1 branch"
  exit 1
fi
cat $bench_stdout
end=$(date +%s.%N)
awk "BEGIN {print $end - $start}" > ${1}_bench_time.txt
cp docs/data/lastBenchmark.json ${1}_bench.json
