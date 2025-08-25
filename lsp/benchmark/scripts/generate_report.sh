#!/usr/bin/env bash

set -euo pipefail

# Generate Report Script
# Post-processes benchmark results and generates additional analysis

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
BENCHMARK_DIR="$(dirname "$SCRIPT_DIR")"
RESULTS_DIR="$BENCHMARK_DIR/results"

echo "LSP Benchmark Report Generator"
echo "============================="

# Color output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

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

# Check if results exist
check_results() {
    if [[ ! -d "$RESULTS_DIR" ]]; then
        log_error "Results directory not found: $RESULTS_DIR"
        log_error "Please run the benchmark suite first: ./run_benchmarks.sh"
        exit 1
    fi
    
    if [[ ! -f "$RESULTS_DIR/comparison_report.md" ]]; then
        log_error "Comparison report not found. Benchmark may not have completed successfully."
        exit 1
    fi
    
    log_success "Results directory found"
}

# Generate HTML report from markdown
generate_html_report() {
    log_info "Generating HTML report..."
    
    if command -v pandoc &> /dev/null; then
        pandoc "$RESULTS_DIR/comparison_report.md" \
            -f markdown \
            -t html \
            --standalone \
            --css https://cdn.jsdelivr.net/npm/github-markdown-css@5/github-markdown.css \
            --metadata title="LSP Benchmark Comparison" \
            -o "$RESULTS_DIR/comparison_report.html"
        
        log_success "HTML report generated: $RESULTS_DIR/comparison_report.html"
    else
        log_warn "pandoc not found, skipping HTML report generation"
        log_info "Install pandoc to generate HTML reports: apt-get install pandoc"
    fi
}

# Generate performance charts (if gnuplot is available)
generate_charts() {
    log_info "Generating performance charts..."
    
    if ! command -v gnuplot &> /dev/null; then
        log_warn "gnuplot not found, skipping chart generation"
        log_info "Install gnuplot to generate charts: apt-get install gnuplot"
        return 0
    fi
    
    # Create gnuplot script for performance comparison
    cat > "$RESULTS_DIR/performance_chart.gnuplot" << 'EOF'
set terminal png size 800,600
set output "performance_comparison.png"
set title "LSP Performance Comparison (Lower is Better)"
set ylabel "Response Time (ms)"
set xlabel "Scenario"
set grid
set style data histograms
set style histogram cluster gap 1
set style fill solid border -1
set boxwidth 0.9
set xtic rotate by -45 scale 0

# Read data from CSV file
plot "benchmark_data.csv" using 3:xtic(1) title "cem LSP" linecolor rgb "#2E86AB", \
     "" using 4 title "wc-toolkit LSP" linecolor rgb "#A23B72"
EOF
    
    # Generate chart if CSV data exists
    if [[ -f "$RESULTS_DIR/benchmark_data.csv" ]]; then
        cd "$RESULTS_DIR"
        
        # Process CSV to create gnuplot-friendly format
        tail -n +2 benchmark_data.csv | \
        grep "_score_ms" | \
        awk -F, '{print $1 " " $3 " " $4}' > plot_data.txt
        
        # Create simplified gnuplot script
        cat > simple_plot.gnuplot << 'EOF'
set terminal png size 1000,600
set output "performance_comparison.png"
set title "LSP Response Time Comparison (Lower is Better)"
set ylabel "Response Time (ms)"
set xlabel "Test Scenario"
set grid
set style data histograms
set style histogram cluster gap 1
set style fill solid border -1
set boxwidth 0.8
set xtic rotate by -45
set key outside right top

plot "plot_data.txt" using 2:xtic(1) title "cem LSP" linecolor rgb "#0066CC", \
     "" using 3 title "wc-toolkit LSP" linecolor rgb "#CC6600"
EOF
        
        if gnuplot simple_plot.gnuplot; then
            log_success "Performance chart generated: $RESULTS_DIR/performance_comparison.png"
        else
            log_warn "Failed to generate performance chart"
        fi
        
        # Clean up temporary files
        rm -f plot_data.txt simple_plot.gnuplot
    else
        log_warn "CSV data not found, cannot generate charts"
    fi
}

# Generate JSON summary for CI systems
generate_ci_summary() {
    log_info "Generating CI-friendly JSON summary..."
    
    if [[ ! -f "$RESULTS_DIR/summary.txt" ]]; then
        log_warn "Summary text file not found, cannot generate CI JSON"
        return 0
    fi
    
    # Parse summary.txt and convert to JSON
    local cem_wins=0
    local wc_wins=0
    local ties=0
    local overall_winner="unknown"
    
    if [[ -f "$RESULTS_DIR/summary.txt" ]]; then
        cem_wins=$(grep "cem: " "$RESULTS_DIR/summary.txt" | grep -oE '[0-9]+' || echo "0")
        wc_wins=$(grep "wc-toolkit: " "$RESULTS_DIR/summary.txt" | grep -oE '[0-9]+' || echo "0")
        ties=$(grep "ties: " "$RESULTS_DIR/summary.txt" | grep -oE '[0-9]+' || echo "0")
        
        if grep -q "OVERALL WINNER: cem" "$RESULTS_DIR/summary.txt"; then
            overall_winner="cem"
        elif grep -q "OVERALL WINNER: wc-toolkit" "$RESULTS_DIR/summary.txt"; then
            overall_winner="wc-toolkit"
        elif grep -q "RESULT: Performance tie" "$RESULTS_DIR/summary.txt"; then
            overall_winner="tie"
        fi
    fi
    
    # Generate JSON
    cat > "$RESULTS_DIR/ci_summary.json" << EOF
{
  "timestamp": "$(date -u +"%Y-%m-%dT%H:%M:%SZ")",
  "benchmark_version": "1.0.0",
  "results": {
    "cem_wins": $cem_wins,
    "wc_toolkit_wins": $wc_wins,
    "ties": $ties,
    "total_scenarios": $((cem_wins + wc_wins + ties)),
    "overall_winner": "$overall_winner"
  },
  "files_generated": [
    "comparison_report.md",
    "benchmark_data.csv",
    "summary.txt",
    "cem_results.json",
    "wc_toolkit_results.json"
  ]
}
EOF
    
    log_success "CI summary generated: $RESULTS_DIR/ci_summary.json"
}

# Display summary
display_summary() {
    echo ""
    log_info "Benchmark Results Summary"
    echo "========================="
    
    if [[ -f "$RESULTS_DIR/summary.txt" ]]; then
        cat "$RESULTS_DIR/summary.txt"
    else
        log_warn "Summary file not found"
    fi
    
    echo ""
    log_info "Generated Reports:"
    echo ""
    
    cd "$RESULTS_DIR"
    for file in *.md *.html *.png *.json *.csv *.txt; do
        if [[ -f "$file" ]]; then
            local size
            size=$(du -h "$file" | cut -f1)
            echo "  📄 $file ($size)"
        fi
    done
    
    echo ""
    log_info "View reports:"
    
    if [[ -f "comparison_report.html" ]]; then
        echo "  🌐 HTML: file://$RESULTS_DIR/comparison_report.html"
    fi
    
    if [[ -f "comparison_report.md" ]]; then
        echo "  📝 Markdown: cat $RESULTS_DIR/comparison_report.md"
    fi
    
    if [[ -f "performance_comparison.png" ]]; then
        echo "  📊 Chart: $RESULTS_DIR/performance_comparison.png"
    fi
}

# Archive results with timestamp
archive_results() {
    log_info "Archiving results..."
    
    local timestamp
    timestamp=$(date +"%Y%m%d_%H%M%S")
    local archive_name="lsp_benchmark_$timestamp.tar.gz"
    
    cd "$BENCHMARK_DIR"
    
    tar -czf "$archive_name" results/
    
    log_success "Results archived: $BENCHMARK_DIR/$archive_name"
    
    # Keep only the last 5 archives
    ls -t lsp_benchmark_*.tar.gz 2>/dev/null | tail -n +6 | xargs rm -f || true
}

# Main execution
main() {
    check_results
    generate_html_report
    generate_charts
    generate_ci_summary
    display_summary
    
    if [[ "${ARCHIVE:-false}" == "true" ]]; then
        archive_results
    fi
    
    log_success "Report generation completed!"
}

# Handle command line arguments
case "${1:-}" in
    --help|-h)
        echo "Usage: $0 [options]"
        echo ""
        echo "Options:"
        echo "  --help, -h    Show this help message"
        echo "  --archive     Archive results with timestamp"
        echo "  --html-only   Generate only HTML report"
        echo "  --charts-only Generate only performance charts"
        echo ""
        echo "Environment variables:"
        echo "  ARCHIVE=true  Archive results after generation"
        exit 0
        ;;
    --archive)
        ARCHIVE=true main
        ;;
    --html-only)
        check_results
        generate_html_report
        ;;
    --charts-only)
        check_results
        generate_charts
        ;;
    *)
        main
        ;;
esac