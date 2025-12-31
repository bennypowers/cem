# LSP Server Benchmark Suite

A comprehensive, statistically rigorous benchmarking framework for comparing Language Server Protocol (LSP) implementations. This benchmarks cem LSP vs wc-toolkit LSP with fair, reproducible testing methodology.

## 🎯 Features

- **Statistically Comprehensive**: Multiple iterations with P50/P95/P99 percentile analysis
- **Modular Architecture**: Clean separation of configuration and test logic
- **Memory Tracking**: Monitors memory usage during operations
- **Stress Testing**: Large document and bulk request scenarios
- **Fair Comparison**: Identical client configuration for both servers
- **Automated Reports**: Statistical analysis and comparison generation

## 📁 Architecture

```text
benchmark/
├── configs/                    # Minimal Neovim configurations per tool
│   ├── cem-minimal.lua        # cem LSP configuration
│   └── wc-toolkit-minimal.lua # wc-toolkit LSP configuration
├── modules/                    # Modular benchmark test modules
│   ├── startup_benchmark.lua            # Server initialization (20 iterations) ✅ ACTIVE
│   ├── hover_benchmark.lua              # Hover performance (10 iterations per element) ✅ ACTIVE
│   ├── completion_benchmark.lua         # Completion performance (15 iterations per context) ✅ ACTIVE
│   ├── diagnostics_benchmark.lua        # Diagnostics publishing ✅ ACTIVE
│   ├── references_performance_benchmark.lua  # Go-to-references workspace search ✅ ACTIVE
│   ├── attribute_hover_benchmark.lua    # Attribute hover correctness (21 attributes) ✅ ACTIVE
│   ├── lit_template_benchmark.lua       # Lit-html template support testing ✅ ACTIVE
│   ├── file_lifecycle_benchmark.lua     # ⚠️ DEPRECATED: Artificial delays dominate timing
│   ├── stress_test_benchmark.lua        # ⚠️ DEPRECATED: Bulk operations not realistic
│   ├── edit_cycles_benchmark.lua        # ⚠️ DEPRECATED: Synthetic edit patterns
│   ├── incremental_parsing_benchmark.lua # ⚠️ DEPRECATED: Neovim overhead obscures LSP timing
│   ├── multi_buffer_benchmark.lua       # ⚠️ DEPRECATED: Artificial wait periods
│   ├── neovim_workflow_benchmark.lua    # ⚠️ DEPRECATED: Editor operations, not LSP
│   └── large_project_benchmark.lua      # ⚠️ DEPRECATED: Small test fixture (5 files)
├── utils/                      # Measurement and reporting utilities
│   ├── measurement.lua        # Statistical analysis functions
│   ├── reporting.lua          # Output formatting
│   └── workspace.lua          # Workspace management
├── fixtures/                  # Test data
│   ├── medium_project/        # Standard test project
│   ├── large_project/         # Stress testing project
│   └── small_project/         # Minimal test project
├── results/                   # Generated benchmark results
│   ├── cem_modular_results_*.json      # cem LSP results
│   ├── wc-toolkit_modular_results_*.json # wc-toolkit LSP results
│   └── comparison_report_*.txt         # Statistical comparison reports
├── run_modular_benchmark.lua  # Main benchmark runner
├── run_comparison.sh          # Convenience script for both servers
└── generate_comparison_report.lua # Statistical analysis generator
```

## 🚀 Quick Start

### Prerequisites

- **cem LSP**: Install with `go install bennypowers.dev/cem@latest` or build from source (`make build`)
- **wc-toolkit LSP**: Run `./setup_wc_toolkit.sh` to create an isolated npx wrapper (requires Node.js/npx)
- **Neovim**: Version 0.9+ with LSP support

### Setup wc-toolkit

The setup script creates an isolated npx wrapper that won't affect your Mason configuration:

```bash
cd lsp/benchmark
./setup_wc_toolkit.sh
```

This creates `bin/wc-language-server` which uses `npx @wc-toolkit/language-server` on demand.

### Run Benchmarks

```bash
# Comparison benchmarks for both servers
./run_comparison.sh

# Or use make targets
make bench-lsp        # Both servers
make bench-lsp-cem    # CEM only
make bench-lsp-wc     # wc-toolkit only

# Individual server runs
nvim --headless --clean -u configs/cem-minimal.lua -l run_modular_benchmark.lua
nvim --headless --clean -u configs/wc-toolkit-minimal.lua -l run_modular_benchmark.lua

# Generate statistical comparison report
nvim --headless --clean -l generate_comparison_report.lua
```

## 📊 Statistical Analysis

### Metrics Collected

- **Timing Statistics**: Mean, Median, P95, P99, Standard Deviation
- **Success Rates**: Percentage of successful operations
- **Memory Usage**: Memory delta during operations
- **Reliability**: Overall benchmark completion rates
- **Error Analysis**: Detailed failure categorization

## 🎯 Benchmark Philosophy

This suite focuses on **pure LSP protocol timing** without editor overhead. The benchmarks measure actual LSP server performance by isolating protocol request/response cycles from editor operations.

### What We Measure

- **Protocol Operations**: Direct LSP method calls (hover, completion, diagnostics, etc.)
- **Statistical Rigor**: Multiple iterations with P50/P95/P99 percentile analysis
- **Memory Efficiency**: Server memory usage during operations
- **Reliability**: Success rates and error categorization

### What We Don't Measure

**Workflow benchmarks (deprecated)**: Tests like "incremental parsing", "edit cycles", and "file lifecycle" are **not used** because they measure artificial delays (50ms-2000ms waits) rather than actual LSP performance. These tests obscure real performance differences with Neovim buffer operation overhead and synthetic timing delays.

**Focus on signal, not noise**: Pure LSP protocol benchmarks reveal actual server performance without confounding variables.

## 🧪 Active Benchmarks

### 1. Startup Benchmark
- **Iterations**: 20
- **Measures**: Server initialization time, capabilities loading
- **Metrics**: Mean, P95, P99 latency, memory usage

### 2. Hover Benchmark
- **Elements**: my-button, my-card, my-icon
- **Iterations**: 10 per element (30 total)
- **Measures**: Pure LSP hover response time, content accuracy

### 3. Completion Benchmark
- **Contexts**: Tag names, attributes, nested elements
- **Iterations**: 15 per context (45 total)
- **Measures**: Pure LSP completion latency, item count, relevance

### 4. Diagnostics Benchmark
- **Content**: HTML with intentional errors
- **Measures**: Diagnostics publishing time, error categorization
- **Capability Detection**: Gracefully handles servers without diagnostics support

### 5. Attribute Hover Benchmark
- **Attributes**: 21 different attribute types across element categories
- **Coverage**: Buttons, forms, layout, charts, tables, media, overlays
- **Iterations**: 3 per attribute type
- **Measures**: Hover content accuracy, response completeness

### 6. Lit Template Benchmark
- **Templates**: Lit html`` template literals
- **Iterations**: 5 scenarios per template type
- **Measures**: Template parsing, completion, hover accuracy in TypeScript

### 7. References Performance Benchmark
- **Operations**: Go-to-references workspace search
- **Scales**: Small (15 refs), medium (42 refs) projects
- **Measures**: Reference search time, accuracy, context awareness

## 📈 Running Benchmarks

To generate current benchmark results:

```bash
# Run comparison benchmarks for both servers
make bench-lsp

# Run individual server benchmarks
make bench-lsp-cem    # CEM LSP only
make bench-lsp-wc     # wc-toolkit LSP only
```

Results are saved to `lsp/benchmark/results/` as JSON files with:
- Statistical analysis (mean, median, P95, P99, stddev)
- Success rates and error categorization
- Memory usage metrics
- Timestamp and server information

View results at: [CEM Docs](https://cem.run/docs/benchmarks/)

## 🛠 Development

### Adding New Benchmarks

1. Create module in `modules/new_benchmark.lua`:
```lua
local measurement = require('utils.measurement')
local M = {}

function M.run_new_benchmark(config, fixture_dir)
  -- Implementation
  return {
    success = true,
    server_name = _G.BENCHMARK_LSP_NAME,
    statistics = stats,
    -- Additional metrics
  }
end

return M
```

2. Add to `run_modular_benchmark.lua`:
```lua
local new_benchmark = require('modules.new_benchmark')
-- Add to benchmarks table
{name = "new", module = new_benchmark}
```

### Statistical Significance

- **20 iterations** for startup (captures initialization variance)
- **10-15 iterations** for operations (balances accuracy vs time)
- **P95/P99 percentiles** for latency analysis (identifies outliers)
- **Multiple test positions** per scenario (reduces positional bias)

## 📝 Contributing

When adding new benchmarks:
1. Follow the established module interface
2. Use `measurement.run_iterations()` for statistical rigor
3. Include both timing and functional validation
4. Add appropriate error handling and cleanup
5. Update this README with new scenarios

## 🎛 Configuration

Modify iterations for different trade-offs:
- **More iterations**: Higher statistical confidence, longer runtime
- **Fewer iterations**: Faster feedback, lower precision
- **Different document sizes**: Test various workload scenarios
- **Additional LSP methods**: Extend coverage (definitions, references, etc.)

The current configuration provides a good balance of statistical rigor and reasonable execution time (~5-10 minutes for full comparison).
