#!/usr/bin/env bash

set -euo pipefail

# LSP Benchmark Runner
# Runs comprehensive benchmarks comparing cem LSP vs wc-toolkit LSP

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
BENCHMARK_DIR="$(dirname "$SCRIPT_DIR")"
PROJECT_ROOT="$(dirname "$(dirname "$(dirname "$SCRIPT_DIR")")")"

echo "========================================"
echo "LSP Benchmark Suite"
echo "========================================"
echo "Script dir: $SCRIPT_DIR"
echo "Benchmark dir: $BENCHMARK_DIR"
echo "Project root: $PROJECT_ROOT"
echo ""

# Color output for better readability
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

log_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

log_warn() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

log_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

# Check prerequisites
check_prerequisites() {
    log_info "Checking prerequisites..."
    
    # Check Neovim
    if ! command -v nvim &> /dev/null; then
        log_error "Neovim not found. Please install Neovim 0.9+ with LSP support."
        exit 1
    fi
    
    local nvim_version
    nvim_version=$(nvim --version | head -n1 | grep -oE 'v[0-9]+\.[0-9]+')
    log_info "Found Neovim $nvim_version"
    
    # Check if we're in the correct directory
    if [[ ! -f "$PROJECT_ROOT/go.mod" ]]; then
        log_error "Not in cem project root. Expected to find go.mod at $PROJECT_ROOT"
        exit 1
    fi
    
    # Change to project root for builds
    cd "$PROJECT_ROOT"
    
    log_success "Prerequisites check passed"
}

# Build cem LSP server
build_cem_lsp() {
    log_info "Building cem LSP server..."
    
    if [[ -f "cem" ]]; then
        log_info "Found existing cem binary, checking if rebuild needed..."
        if [[ "cem" -ot "$(find . -name '*.go' -type f | head -1)" ]]; then
            log_info "Source files newer than binary, rebuilding..."
            rm -f cem
        else
            log_info "Binary is up to date"
            return 0
        fi
    fi
    
    go build -o cem .
    
    if [[ ! -x "cem" ]]; then
        log_error "Failed to build cem binary"
        exit 1
    fi
    
    log_success "cem LSP server built successfully"
}

# Setup wc-toolkit LSP server
setup_wc_toolkit() {
    log_info "Setting up wc-toolkit LSP server..."
    
    local repo_dir="$BENCHMARK_DIR/servers/wc-language-server"
    local server_path="$repo_dir/packages/language-server/bin/wc-language-server.js"
    
    # Check if already built
    if [[ -f "$server_path" ]]; then
        log_info "wc-toolkit LSP already built at $server_path"
        return 0
    fi
    
    # Check prerequisites
    if ! command -v pnpm &> /dev/null; then
        log_error "pnpm not found. Please install pnpm: npm install -g pnpm"
        exit 1
    fi
    
    if ! command -v git &> /dev/null; then
        log_error "git not found. Please install git"
        exit 1
    fi
    
    # Create servers directory
    mkdir -p "$BENCHMARK_DIR/servers"
    cd "$BENCHMARK_DIR/servers"
    
    # Clone the repository if it doesn't exist
    if [[ ! -d "wc-language-server" ]]; then
        log_info "Cloning wc-toolkit language server repository..."
        git clone https://github.com/wc-toolkit/wc-language-server.git
    fi
    
    # Install dependencies (build may have some errors but the existing JS files should work)
    log_info "Installing wc-toolkit LSP server dependencies..."
    cd wc-language-server
    
    # Use pnpm as required by the project
    PNPM_HOME=$HOME/.local/share/pnpm pnpm install --no-optional
    
    # Try to build, but don't fail if TypeScript compilation has errors since JS files exist
    log_info "Attempting to build (may show TypeScript errors but should still work)..."
    PNPM_HOME=$HOME/.local/share/pnpm pnpm run build || {
        log_warn "Build failed with TypeScript errors, but using existing compiled JavaScript files"
    }
    
    if [[ ! -f "$server_path" ]]; then
        log_error "wc-toolkit LSP server not found after build: $server_path"
        exit 1
    fi
    
    # Return to project root
    cd "$PROJECT_ROOT"
    
    log_success "wc-toolkit LSP setup complete"
}

# Run the benchmark suite
run_benchmarks() {
    log_info "Running LSP benchmark suite..."
    
    cd "$BENCHMARK_DIR"
    
    # Create results directory
    mkdir -p results
    
    # Run benchmarks with Neovim headless mode
    local benchmark_start
    benchmark_start=$(date +%s)
    
    echo ""
    log_info "Starting benchmark execution..."
    echo "This may take several minutes depending on your system..."
    echo ""
    
    if nvim --headless -c "luafile benchmark.lua" 2>&1; then
        local benchmark_end
        benchmark_end=$(date +%s)
        local duration=$((benchmark_end - benchmark_start))
        
        log_success "Benchmark suite completed in ${duration} seconds"
        
        # Check if results were generated
        if [[ -f "results/comparison_report.md" ]]; then
            log_success "Comparison report generated: results/comparison_report.md"
        else
            log_warn "Comparison report not found, but benchmark may have completed partially"
        fi
        
        if [[ -f "results/summary.txt" ]]; then
            echo ""
            log_info "Benchmark Summary:"
            echo "=================="
            cat results/summary.txt
        fi
        
    else
        log_error "Benchmark execution failed"
        exit 1
    fi
}

# Display results
display_results() {
    log_info "Generated files:"
    
    cd "$BENCHMARK_DIR/results"
    
    for file in *.md *.json *.csv *.txt; do
        if [[ -f "$file" ]]; then
            local size
            size=$(du -h "$file" | cut -f1)
            echo "  - $file ($size)"
        fi
    done
    
    echo ""
    log_info "To view the detailed comparison report:"
    echo "  cat $BENCHMARK_DIR/results/comparison_report.md"
    echo ""
    log_info "To view raw data:"
    echo "  cat $BENCHMARK_DIR/results/benchmark_data.csv"
}

# Cleanup function
cleanup() {
    log_info "Cleaning up temporary files..."
    # Add any cleanup operations here
}

# Main execution
main() {
    trap cleanup EXIT
    
    check_prerequisites
    build_cem_lsp
    setup_wc_toolkit
    run_benchmarks
    display_results
    
    log_success "LSP benchmark suite completed successfully!"
}

# Handle command line arguments
case "${1:-}" in
    --help|-h)
        echo "Usage: $0 [options]"
        echo ""
        echo "Options:"
        echo "  --help, -h    Show this help message"
        echo "  --check       Check prerequisites only"
        echo "  --build       Build cem LSP only"
        echo "  --setup       Setup wc-toolkit LSP only"
        echo ""
        echo "This script runs comprehensive benchmarks comparing cem LSP vs wc-toolkit LSP"
        echo "using Neovim as a consistent client for fair comparison."
        exit 0
        ;;
    --check)
        check_prerequisites
        log_success "Prerequisites check completed"
        exit 0
        ;;
    --build)
        check_prerequisites
        build_cem_lsp
        exit 0
        ;;
    --setup)
        setup_wc_toolkit
        exit 0
        ;;
    *)
        main
        ;;
esac