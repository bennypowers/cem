#!/bin/bash
# Modular benchmark comparison script
# Runs benchmarks for both cem and wc-toolkit LSP servers using separate configs

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR"

echo "LSP Server Comparison Benchmark"
echo "================================"
echo "Using modular architecture with separate configs per tool"
echo

# Check if benchmark files exist
if [[ ! -f "configs/cem-minimal.lua" ]]; then
    echo "ERROR: cem-minimal.lua config not found"
    exit 1
fi

if [[ ! -f "configs/wc-toolkit-minimal.lua" ]]; then
    echo "ERROR: wc-toolkit-minimal.lua config not found"
    exit 1
fi

if [[ ! -f "run_modular_benchmark.lua" ]]; then
    echo "ERROR: run_modular_benchmark.lua not found"
    exit 1
fi

# Check if cem is available
echo "Checking cem LSP availability..."
if ! command -v cem &> /dev/null; then
    echo "ERROR: cem command not found. Please install cem LSP server."
    exit 1
fi

# Check if wc-toolkit server exists
WC_SERVER_PATH="servers/wc-language-server/packages/language-server/bin/wc-language-server.js"
echo "Checking wc-toolkit LSP availability..."
if [[ ! -f "$WC_SERVER_PATH" ]]; then
    echo "ERROR: wc-toolkit server not found at $WC_SERVER_PATH"
    echo "Please run: ./scripts/setup_wc_toolkit.sh"
    exit 1
fi

# Create results directory
mkdir -p results

# Run cem LSP benchmark
echo "1/2: Running cem LSP benchmark..."
echo "=================================="
if nvim --headless --clean -u configs/cem-minimal.lua -l run_modular_benchmark.lua; then
    echo "✅ cem benchmark completed successfully"
else
    echo "❌ cem benchmark failed"
fi

echo
echo "Waiting 2 seconds between benchmarks..."
sleep 2
echo

# Run wc-toolkit LSP benchmark  
echo "2/2: Running wc-toolkit LSP benchmark..."
echo "========================================="
if nvim --headless --clean -u configs/wc-toolkit-minimal.lua -l run_modular_benchmark.lua; then
    echo "✅ wc-toolkit benchmark completed successfully"
else
    echo "❌ wc-toolkit benchmark failed"
fi

echo
echo "Comparison Results"
echo "=================="

# Find the latest results files
CEM_RESULT=$(find results -name "cem_modular_results_*.json" -type f -printf '%T@ %p\n' | sort -k 1nr | head -1 | cut -d' ' -f2- || echo "")
WCT_RESULT=$(find results -name "wc-toolkit_modular_results_*.json" -type f -printf '%T@ %p\n' | sort -k 1nr | head -1 | cut -d' ' -f2- || echo "")

if [[ -n "$CEM_RESULT" && -f "$CEM_RESULT" ]]; then
    echo "✅ cem results: $CEM_RESULT"
else
    echo "❌ No cem results found"
fi

if [[ -n "$WCT_RESULT" && -f "$WCT_RESULT" ]]; then
    echo "✅ wc-toolkit results: $WCT_RESULT"
else
    echo "❌ No wc-toolkit results found"
fi

echo
echo "Benchmark comparison completed!"
echo "Check the results/ directory for detailed JSON output."