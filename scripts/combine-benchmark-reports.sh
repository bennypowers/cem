#!/usr/bin/env bash

set -euo pipefail

# Combine LSP benchmark report and generate benchmark report into a single file
# LSP benchmarks come first to ensure they're visible even if GitHub truncates the comment
cat lsp_bench_report.md > combined_bench_report.md
echo -e "\n---\n" >> combined_bench_report.md
cat bench_report.md >> combined_bench_report.md

echo "✅ Combined benchmark reports created: combined_bench_report.md"
