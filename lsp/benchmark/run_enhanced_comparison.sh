#!/bin/bash
# Enhanced benchmark comparison script  
# Runs ENHANCED benchmarks for both cem and wc-toolkit LSP servers
# Focus: Real-world performance differences (multi-buffer, workflows, incremental parsing)
# Time limit: 120 seconds total per server (CI-friendly)

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR"

echo "Enhanced LSP Server Performance Comparison"
echo "=========================================="
echo "🚀 Real-world performance testing: multi-buffer, Neovim workflows, incremental parsing"
echo "⏱️  Time limit: 120 seconds per server (CI-friendly)"
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

if [[ ! -f "run_enhanced_benchmarks.lua" ]]; then
    echo "ERROR: run_enhanced_benchmarks.lua not found"
    exit 1
fi

# Check if cem is available
echo "Checking cem LSP availability..."
if ! command -v cem &> /dev/null; then
    echo "ERROR: cem command not found. Please install cem LSP server."
    exit 1
fi

# Check if wc-toolkit server exists (using the corrected path from investigation)
WC_SERVER_PATH="servers/wc-language-server/packages/vscode/dist/server.js"
echo "Checking wc-toolkit LSP availability..."
if [[ ! -f "$WC_SERVER_PATH" ]]; then
    echo "ERROR: wc-toolkit server not found at $WC_SERVER_PATH"
    echo "Please ensure VS Code extension is built: cd $WC_SERVER_PATH/../.. && npm run build"
    exit 1
fi

# Create results directory
mkdir -p results

# Track total time
TOTAL_START_TIME=$(date +%s)

# Run cem LSP enhanced benchmark
echo "1/2: Running cem LSP enhanced benchmark..."
echo "=========================================="
CEM_START_TIME=$(date +%s)
if timeout 150s nvim --headless --clean -u configs/cem-minimal.lua -l run_enhanced_benchmarks.lua; then
    CEM_END_TIME=$(date +%s)
    CEM_DURATION=$((CEM_END_TIME - CEM_START_TIME))
    echo "✅ cem enhanced benchmark completed successfully in ${CEM_DURATION}s"
else
    CEM_END_TIME=$(date +%s)
    CEM_DURATION=$((CEM_END_TIME - CEM_START_TIME))
    echo "❌ cem enhanced benchmark failed or timed out after ${CEM_DURATION}s"
fi

echo
echo "Waiting 3 seconds between benchmarks..."
sleep 3
echo

# Run wc-toolkit LSP enhanced benchmark  
echo "2/2: Running wc-toolkit LSP enhanced benchmark..."
echo "=================================================="
WCT_START_TIME=$(date +%s)
if timeout 150s nvim --headless --clean -u configs/wc-toolkit-minimal.lua -l run_enhanced_benchmarks.lua; then
    WCT_END_TIME=$(date +%s)
    WCT_DURATION=$((WCT_END_TIME - WCT_START_TIME))
    echo "✅ wc-toolkit enhanced benchmark completed successfully in ${WCT_DURATION}s"
else
    WCT_END_TIME=$(date +%s)
    WCT_DURATION=$((WCT_END_TIME - WCT_START_TIME))
    echo "❌ wc-toolkit enhanced benchmark failed or timed out after ${WCT_DURATION}s"
fi

echo
echo "Enhanced Comparison Results"
echo "==========================="

# Find the latest enhanced results files
CEM_RESULT=$(find results -name "cem_enhanced_results_*.json" -type f -printf '%T@ %p\n' | sort -k 1nr | head -1 | cut -d' ' -f2- 2>/dev/null || echo "")
WCT_RESULT=$(find results -name "wc-toolkit_enhanced_results_*.json" -type f -printf '%T@ %p\n' | sort -k 1nr | head -1 | cut -d' ' -f2- 2>/dev/null || echo "")

if [[ -n "$CEM_RESULT" && -f "$CEM_RESULT" ]]; then
    echo "✅ cem enhanced results: $CEM_RESULT (${CEM_DURATION}s)"
else
    echo "❌ No cem enhanced results found"
fi

if [[ -n "$WCT_RESULT" && -f "$WCT_RESULT" ]]; then
    echo "✅ wc-toolkit enhanced results: $WCT_RESULT (${WCT_DURATION}s)"
else
    echo "❌ No wc-toolkit enhanced results found"
fi

# Calculate total time
TOTAL_END_TIME=$(date +%s)
TOTAL_DURATION=$((TOTAL_END_TIME - TOTAL_START_TIME))

echo
echo "Performance Summary"
echo "==================="

# Quick comparison from JSON files if available
if [[ -n "$CEM_RESULT" && -f "$CEM_RESULT" && -n "$WCT_RESULT" && -f "$WCT_RESULT" ]]; then
    echo "📊 Quick Comparison:"
    
    # Extract success rates using jq if available, otherwise skip detailed comparison
    if command -v jq &> /dev/null; then
        CEM_SUCCESS_COUNT=$(jq -r '.benchmarks | to_entries | map(select(.value.success == true)) | length' "$CEM_RESULT" 2>/dev/null || echo "N/A")
        CEM_TOTAL_COUNT=$(jq -r '.benchmarks | length' "$CEM_RESULT" 2>/dev/null || echo "N/A")
        WCT_SUCCESS_COUNT=$(jq -r '.benchmarks | to_entries | map(select(.value.success == true)) | length' "$WCT_RESULT" 2>/dev/null || echo "N/A")
        WCT_TOTAL_COUNT=$(jq -r '.benchmarks | length' "$WCT_RESULT" 2>/dev/null || echo "N/A")
        
        echo "   cem:        ${CEM_SUCCESS_COUNT}/${CEM_TOTAL_COUNT} benchmarks passed"
        echo "   wc-toolkit: ${WCT_SUCCESS_COUNT}/${WCT_TOTAL_COUNT} benchmarks passed"
    else
        echo "   (Install jq for detailed comparison metrics)"
    fi
    
    echo "   Execution time: cem ${CEM_DURATION}s, wc-toolkit ${WCT_DURATION}s"
else
    echo "❌ Cannot generate comparison - missing results files"
fi

echo
echo "🎯 Architecture Performance Insights:"
echo "   • Multi-buffer handling tests Go's concurrency vs TypeScript single-threading"
echo "   • Incremental parsing tests tree-sitter benefits vs traditional parsing"  
echo "   • Neovim workflows test real user experience vs synthetic benchmarks"
echo "   • Large project simulation tests workspace handling at scale"

echo
echo "Total enhanced benchmark time: ${TOTAL_DURATION}s"
echo
echo "Enhanced benchmark comparison completed!"
echo "Check the results/ directory for detailed JSON output with real-world performance metrics."