# Modular LSP Benchmark Architecture

This document describes the refactored benchmark architecture that separates Neovim configuration from test code, as requested.

## Architecture Overview

```
benchmark/
├── configs/                    # Minimal Neovim configurations per tool
│   ├── cem-minimal.lua        # Minimal config for cem LSP
│   └── wc-toolkit-minimal.lua # Minimal config for wc-toolkit LSP
├── modules/                    # Modular benchmark test modules
│   ├── startup_benchmark.lua            # Server initialization (20 iterations)
│   ├── hover_benchmark.lua              # Hover performance (10 iterations per element)  
│   ├── completion_benchmark.lua         # Completion performance (15 iterations per context)
│   ├── diagnostics_benchmark.lua        # Diagnostics publishing (with capability detection)
│   ├── file_lifecycle_benchmark.lua     # File operations (5 iterations)
│   ├── stress_test_benchmark.lua        # Large workload scenarios
│   ├── attribute_hover_benchmark.lua    # Attribute hover correctness (21 attributes)
│   ├── edit_cycles_benchmark.lua        # Edit responsiveness (rapid HTML editing)
│   └── lit_template_benchmark.lua       # Lit-html template support testing
├── run_modular_benchmark.lua   # Main benchmark runner
├── run_comparison.sh          # Convenience script for both servers
└── fixtures/                  # Test data
    └── medium_project/        # Test project with custom-elements.json
```

## Design Principles

### 1. Separation of Concerns
- **Configs**: Minimal Neovim configurations with only LSP setup
- **Modules**: Pure benchmark logic with no configuration dependencies  
- **Runner**: Orchestrates benchmarks using imported modules

### 2. Dependency Injection
- Configs export `_G.BENCHMARK_LSP_CONFIG` and `_G.BENCHMARK_LSP_NAME`
- Modules receive config and fixture directory as parameters
- No hard-coded paths or server-specific logic in modules

### 3. Modular Testing
- Each benchmark scenario is its own module
- Modules can be run individually or as a suite
- Clean interfaces with standardized return formats

## Usage

### Run Individual Server
```bash
# cem LSP benchmarks
nvim --headless --clean -u configs/cem-minimal.lua -l run_modular_benchmark.lua

# wc-toolkit LSP benchmarks  
nvim --headless --clean -u configs/wc-toolkit-minimal.lua -l run_modular_benchmark.lua
```

### Run Comparison
```bash
# Run both servers and compare results
./run_comparison.sh
```

### Run Individual Module
```bash
# Test just hover performance for cem
nvim --headless --clean -u configs/cem-minimal.lua -c "
  local hover = require('modules.hover_benchmark')
  local result = hover.run_hover_benchmark(_G.BENCHMARK_LSP_CONFIG, './fixtures/medium_project')
  print(vim.inspect(result))
  vim.cmd('qa!')
"
```

## Configuration Files

### configs/cem-minimal.lua
- Minimal LSP setup for cem server
- Auto-starts cem LSP for HTML/TS/JS files
- Exports `_G.BENCHMARK_LSP_CONFIG` and `_G.BENCHMARK_LSP_NAME = "cem"`

### configs/wc-toolkit-minimal.lua  
- Minimal LSP setup for wc-toolkit server
- Finds wc-toolkit server path automatically
- Exports `_G.BENCHMARK_LSP_CONFIG` and `_G.BENCHMARK_LSP_NAME = "wc-toolkit"`

## Benchmark Modules

Each module follows the same interface pattern:

```lua
local M = {}

function M.run_<scenario>_benchmark(config, fixture_dir)
  return {
    success = true|false,
    server_name = "cem"|"wc-toolkit", 
    duration_ms = number,
    -- scenario-specific data
  }
end

return M
```

### Available Modules

1. **startup_benchmark.lua**: Tests server initialization time (20 iterations)
2. **hover_benchmark.lua**: Tests hover request performance and functionality (30 total tests)
3. **completion_benchmark.lua**: Tests completion request performance (45 total tests)
4. **diagnostics_benchmark.lua**: Tests diagnostics publishing with capability detection
5. **file_lifecycle_benchmark.lua**: Tests file open/change/close operations (5 iterations)
6. **stress_test_benchmark.lua**: Tests high-load scenarios with large documents (800+ lines)
7. **attribute_hover_benchmark.lua**: Tests attribute hover correctness across 21 attribute types
8. **edit_cycles_benchmark.lua**: Tests LSP responsiveness during rapid HTML editing (3 scenarios)
9. **lit_template_benchmark.lua**: Tests LSP functionality within TypeScript template literals (5 scenarios)

## Results Format

Results are saved as JSON files in `results/` directory:

```json
{
  "server_name": "cem",
  "timestamp": "2025-08-24 10:30:45", 
  "fixture_dir": "/path/to/fixtures/medium_project",
  "benchmarks": {
    "startup": {
      "success": true,
      "duration_ms": 150.5,
      "capabilities_count": 8
    },
    "hover": {
      "success": true,
      "successful_hovers": 3,
      "success_rate": 1.0,
      "average_duration_ms": 45.2
    }
    // ... other benchmarks
  }
}
```

## Benefits of New Architecture

### 1. Clean Separation
- Config files contain only Neovim/LSP setup
- Test modules contain only benchmark logic
- No mixing of configuration and testing concerns

### 2. Better Debugging
- Can test individual scenarios easily
- Clear error isolation between config and test issues
- Minimal configs reduce variables in debugging

### 3. Extensibility
- Easy to add new LSP servers (just add new config file)
- Easy to add new benchmark scenarios (just add new module)
- Module interfaces are standardized and reusable

### 4. Maintainability  
- Single responsibility principle applied throughout
- Dependencies are explicit and injected
- No duplicate configuration across different test files

### 5. CI/CD Friendly
- Scripts can run individual servers or full comparisons
- JSON output format for automated analysis
- Exit codes indicate overall success/failure

## Migration from Previous Architecture

The previous monolithic files like `benchmark.lua`, `benchmark_enhanced.lua` etc. mixed configuration and testing logic. The new architecture provides:

1. **Better Organization**: Clear file structure with logical separation
2. **Reusability**: Modules can be used in different contexts
3. **Testability**: Individual components can be tested in isolation
4. **Maintainability**: Changes to one aspect don't affect others

## Future Enhancements

The modular architecture enables easy addition of:

- New LSP servers (just add config file)
- New benchmark scenarios (just add module) 
- New output formats (modify runner)
- Advanced analysis tools (consume JSON results)
- CI/CD integration (use run_comparison.sh)