#!/usr/bin/env bash

set -euo pipefail

# Combine generate benchmark report and LSP benchmark report into a single file
cat bench_report.md > combined_bench_report.md
echo -e "\n---\n" >> combined_bench_report.md
cat lsp_bench_report.md >> combined_bench_report.md

echo "✅ Combined benchmark reports created: combined_bench_report.md"
