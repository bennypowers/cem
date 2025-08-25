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

```
benchmark/
├── configs/                    # Minimal Neovim configurations per tool
│   ├── cem-minimal.lua        # cem LSP configuration
│   └── wc-toolkit-minimal.lua # wc-toolkit LSP configuration
├── modules/                    # Modular benchmark test modules
│   ├── startup_benchmark.lua            # Server initialization (20 iterations)
│   ├── hover_benchmark.lua              # Hover performance (10 iterations per element)
│   ├── completion_benchmark.lua         # Completion performance (15 iterations per context)
│   ├── diagnostics_benchmark.lua        # Diagnostics publishing
│   ├── file_lifecycle_benchmark.lua     # File operations (5 iterations)
│   ├── stress_test_benchmark.lua        # Large workload scenarios
│   ├── attribute_hover_benchmark.lua    # Attribute hover correctness (21 attributes)
│   ├── edit_cycles_benchmark.lua        # Edit responsiveness (rapid HTML editing)
│   └── lit_template_benchmark.lua       # Lit-html template support testing
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

- **cem LSP**: Install with `go install bennypowers.dev/cem@latest`
- **wc-toolkit LSP**: Run `./scripts/setup_wc_toolkit.sh` to extract from VS Code extension
- **Neovim**: Recent version with LSP support

### Run Benchmarks

```bash
# Enhanced real-world performance comparison (recommended for CI/development)
./run_enhanced_comparison.sh

# Full comprehensive comparison (all benchmarks)
./run_comparison.sh

# Individual servers - enhanced benchmarks only (120s limit)
nvim --headless --clean -u configs/cem-minimal.lua -l run_enhanced_benchmarks.lua
nvim --headless --clean -u configs/wc-toolkit-minimal.lua -l run_enhanced_benchmarks.lua

# Individual servers - all benchmarks
nvim --headless --clean -u configs/cem-minimal.lua -l run_modular_benchmark.lua
nvim --headless --clean -u configs/wc-toolkit-minimal.lua -l run_modular_benchmark.lua

# Generate statistical comparison
nvim --headless --clean -l generate_comparison_report.lua
```

## 📊 Statistical Analysis

### Metrics Collected

- **Timing Statistics**: Mean, Median, P95, P99, Standard Deviation
- **Success Rates**: Percentage of successful operations
- **Memory Usage**: Memory delta during operations
- **Reliability**: Overall benchmark completion rates
- **Error Analysis**: Detailed failure categorization

### Example Output

```
BENCHMARK SUMMARY FOR CEM
==================================================
startup              ✅ PASS (1.98ms) [100% success]
  Stats: Mean=1.98ms, P95=2.65ms, P99=2.96ms, StdDev=0.40ms
  Success rate: 100.0% (20/20)
  Memory usage: 375.2KB

hover                ✅ PASS [67% success]
  Success rate: 66.7% (20/30)

completion           ✅ PASS [100% success]
  Success rate: 100.0% (45/45)

Overall Success Rate: 100.0% (6/6)
🎉 Server performance is GOOD
```

## 🚀 Enhanced Real-World Benchmarks (NEW)

The enhanced benchmarks focus on revealing architectural performance differences between Go/tree-sitter (cem) and TypeScript/Volar (wc-toolkit) through realistic usage scenarios.

### Enhanced Benchmark Modules (120s total time limit)

#### 1. Multi-Buffer Benchmark (30s)
- **Purpose**: Test Go's concurrency advantages vs TypeScript single-threading
- **Scenarios**: Opens 5 documents simultaneously, rapid buffer switching, concurrent LSP requests
- **Metrics**: Buffer switch times, concurrent request latency, throughput measurements
- **Expected Difference**: Go should handle multiple documents more efficiently

#### 2. Neovim Workflow Benchmark (40s)  
- **Purpose**: Test real user experience with editor-driven commands vs direct LSP calls
- **Scenarios**: `:e` file opening, `:b` buffer navigation, `<C-x><C-o>` completion, hover actions
- **Metrics**: End-to-end command latency, user experience timing
- **Expected Difference**: Reveals real-world responsiveness differences

#### 3. Incremental Parsing Benchmark (30s)
- **Purpose**: Stress tree-sitter's incremental parsing advantages during active editing
- **Scenarios**: Character-by-character typing, element insertion/removal, attribute edits, large paste operations
- **Metrics**: Parsing response time after each edit, edit session performance
- **Expected Difference**: Tree-sitter should excel with frequent document changes

#### 4. Large Project Benchmark (20s)
- **Purpose**: Test workspace navigation and symbol search at scale
- **Scenarios**: Multi-file opening, workspace symbol search, navigation operations
- **Metrics**: File opening speed, symbol search performance, workspace operations
- **Expected Difference**: Go's memory management and concurrency should scale better

### Why Enhanced Benchmarks Matter

The original benchmarks showed both servers at 100% reliability with similar performance (1.98ms vs 2.01ms startup). Enhanced benchmarks reveal architectural differences by testing:

- **Concurrent operations** (Go's goroutines vs Node.js single-threading)  
- **Incremental parsing** (tree-sitter's advantages vs traditional parsing)
- **Real user workflows** (editor integration vs synthetic LSP calls)
- **Large project handling** (memory efficiency and workspace operations)

## 🧪 Original Test Scenarios

### 1. Startup Benchmark
- **Iterations**: 20
- **Measures**: Server initialization time, capabilities loading
- **Metrics**: P95/P99 latency, memory usage

### 2. Hover Benchmark
- **Elements**: my-button, my-card, my-icon
- **Iterations**: 10 per element (30 total)
- **Measures**: Hover response time, content length

### 3. Completion Benchmark
- **Contexts**: Tag names, attributes, nested elements
- **Iterations**: 15 per context (45 total)
- **Measures**: Completion latency, item count

### 4. Diagnostics Benchmark
- **Content**: HTML with intentional errors
- **Measures**: Diagnostics publishing time, error categorization
- **Capability Detection**: Gracefully handles servers without diagnostics support

### 5. File Lifecycle Benchmark
- **Operations**: File open, modification, rapid changes
- **Iterations**: 5 file operations
- **Measures**: Document processing performance

### 6. Stress Test Benchmark
- **Document Size**: 800+ lines with 100 custom elements
- **Load Tests**: 50 bulk hovers, 30 rapid completions
- **Measures**: High-load performance, stability

### 7. Attribute Hover Benchmark
- **Attributes**: 21 different attribute types across element categories
- **Coverage**: Buttons, forms, layout, charts, tables, media, overlays
- **Iterations**: 3 per attribute type
- **Measures**: Hover content accuracy, response completeness

### 8. Edit Cycles Benchmark
- **Scenarios**: Attribute addition, element insertion, content modification
- **Edit Types**: 5 rapid edits per scenario across 3 scenarios
- **Measures**: Diagnostics/completion response times during active editing
- **Real-world Simulation**: Mimics developer typing patterns

### 9. Lit-html Template Benchmark
- **Test Areas**: Element hover, attribute completion, tag completion, template interpolation, event binding
- **Template Coverage**: 280+ line TypeScript dashboard application
- **Measures**: LSP functionality within template literals, interpolation awareness

## 📈 Latest Results Summary

Based on the most recent comprehensive benchmarks with bug fixes applied:

### Performance Comparison
- **Startup Time**: cem LSP slightly faster (1.98ms vs 2.01ms)
- **Hover Success**: Both servers 66.7% (my-icon element not found in manifests)
- **Completion**: Both servers 100% success rate
- **Diagnostics**: Both servers support diagnostics (wc-toolkit configuration issue resolved)
- **Edit Cycles**: Both servers 100% success rate after benchmark fixes
- **Stress Testing**: Both servers handle large workloads well

### Reliability Analysis
- **cem LSP**: 100% reliability (9/9 benchmarks pass)
- **wc-toolkit LSP**: 100% reliability (9/9 benchmarks pass, diagnostics working)

### Key Findings
1. **Performance**: Very similar startup and operation times
2. **Functionality**: Both servers provide excellent LSP feature support including diagnostics
3. **Stability**: Both servers handle stress testing and edit cycles excellently  
4. **Completeness**: Both provide excellent completion and hover support
5. **Feature Coverage**: Both servers provide comprehensive LSP support with different architectural approaches

### Recent Bug Fixes Applied ✅
- **Edit Cycles Benchmark**: Fixed invalid column positions causing "out of range" errors
- **Diagnostics Request Protocol**: Fixed incorrect `publishDiagnostics` usage (server-to-client notification vs client-to-server request)
- **Server Capability Detection**: Enhanced diagnostics benchmark to gracefully handle servers without diagnostics support
- **Realistic Edit Positions**: Updated edit scenarios to use valid positions in large-page.html test fixture
- **wc-toolkit Configuration**: Fixed server binary path and TypeScript SDK initialization for proper diagnostics support

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