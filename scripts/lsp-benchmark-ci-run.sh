#!/usr/bin/env bash

set -euo pipefail

branch_name="${1}" # "pr" or "base"

echo "Running LSP benchmarks for ${branch_name} branch..."

# Verify prerequisites
if [ ! -f "dist/cem" ]; then
  echo "::error::cem binary not found. Run 'make build' first."
  exit 1
fi

if ! command -v nvim &> /dev/null; then
  echo "::error::nvim not found. LSP benchmarks require Neovim."
  exit 1
fi

# Run LSP benchmark (CEM only, not wc-toolkit)
cd lsp/benchmark

if ! nvim --headless --clean -u configs/cem-minimal.lua -l run_modular_benchmark.lua 2>&1; then
  echo "::error::LSP benchmarking failed on ${branch_name} branch"
  echo "::warning::Continuing without LSP benchmarks - see logs for details"
  # Create placeholder to avoid breaking downstream scripts
  echo '{"benchmarks":{}, "error": "benchmark_failed"}' > "../../${branch_name}_lsp_bench.json"
  exit 0  # Don't fail workflow, allow graceful degradation
fi

# Verify output was created
if [ ! -f "/tmp/cem-benchmark-cem.json" ]; then
  echo "::error::LSP benchmark did not produce expected output file"
  echo '{"benchmarks":{}, "error": "no_output"}' > "../../${branch_name}_lsp_bench.json"
  exit 0
fi

# Copy result to branch-specific file
cp /tmp/cem-benchmark-cem.json "../../${branch_name}_lsp_bench.json"
echo "✅ LSP benchmark results saved to ${branch_name}_lsp_bench.json"
