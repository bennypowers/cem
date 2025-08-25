#!/usr/bin/env nvim --headless --clean -l
-- Modular LSP Benchmark Runner
-- Usage: nvim --headless --clean -u configs/cem-minimal.lua -l run_modular_benchmark.lua
-- Usage: nvim --headless --clean -u configs/wc-toolkit-minimal.lua -l run_modular_benchmark.lua

-- Ensure the benchmark directory is in the path
local script_dir = vim.fn.getcwd()
package.path = script_dir .. '/?.lua;' .. package.path

-- Import benchmark modules
local startup_benchmark = require('modules.startup_benchmark')
local hover_benchmark = require('modules.hover_benchmark')
local completion_benchmark = require('modules.completion_benchmark')
local diagnostics_benchmark = require('modules.diagnostics_benchmark')
local file_lifecycle_benchmark = require('modules.file_lifecycle_benchmark')
local stress_test_benchmark = require('modules.stress_test_benchmark')
local attribute_hover_benchmark = require('modules.attribute_hover_benchmark')
local edit_cycles_benchmark = require('modules.edit_cycles_benchmark')
local lit_template_benchmark = require('modules.lit_template_benchmark')

-- Import enhanced benchmark modules
local multi_buffer_benchmark = require('modules.multi_buffer_benchmark')
local neovim_workflow_benchmark = require('modules.neovim_workflow_benchmark')
local incremental_parsing_benchmark = require('modules.incremental_parsing_benchmark')
local large_project_benchmark = require('modules.large_project_benchmark')

local function run_all_benchmarks()
  local overall_start_time = vim.fn.reltime()
  local max_total_time_seconds = 300 -- 5 minutes total limit for CI safety
  
  local server_name = _G.BENCHMARK_LSP_NAME or "unknown"
  local config = _G.BENCHMARK_LSP_CONFIG
  
  if not config then
    print("ERROR: No LSP configuration found. Make sure to use the appropriate config file:")
    print("  nvim --headless --clean -u configs/cem-minimal.lua -l run_modular_benchmark.lua")
    print("  nvim --headless --clean -u configs/wc-toolkit-minimal.lua -l run_modular_benchmark.lua")
    return
  end
  
  print(string.format("Running modular benchmarks for %s LSP server (max %ds)", server_name, max_total_time_seconds))
  print("=" .. string.rep("=", 50))
  
  -- Use large project fixture for expanded features
  local fixture_dir = script_dir .. '/fixtures/large_project'
  
  -- Check if fixture exists
  if vim.fn.isdirectory(fixture_dir) == 0 then
    print("ERROR: Fixture directory not found: " .. fixture_dir)
    return
  end
  
  local all_results = {
    server_name = server_name,
    timestamp = os.date("%Y-%m-%d %H:%M:%S"),
    fixture_dir = fixture_dir,
    benchmarks = {}
  }
  
  -- Run each benchmark module
  local benchmarks = {
    {name = "startup", module = startup_benchmark},
    {name = "hover", module = hover_benchmark},
    {name = "completion", module = completion_benchmark},
    {name = "diagnostics", module = diagnostics_benchmark},
    {name = "file_lifecycle", module = file_lifecycle_benchmark},
    {name = "stress_test", module = stress_test_benchmark},
    {name = "attribute_hover", module = attribute_hover_benchmark},
    {name = "edit_cycles", module = edit_cycles_benchmark},
    {name = "lit_template", module = lit_template_benchmark},
    -- Enhanced benchmarks for real-world performance testing (120s total time limit)
    {name = "multi_buffer", module = multi_buffer_benchmark},
    {name = "neovim_workflow", module = neovim_workflow_benchmark},
    {name = "incremental_parsing", module = incremental_parsing_benchmark},
    {name = "large_project", module = large_project_benchmark}
  }
  
  for _, benchmark in ipairs(benchmarks) do
    -- Check total time limit
    local elapsed_time = vim.fn.reltime(overall_start_time)[1]
    if elapsed_time >= max_total_time_seconds then
      print(string.format("\n⏰ Time limit reached (%ds), stopping remaining benchmarks", max_total_time_seconds))
      break
    end
    
    print(string.format("\n--- Running %s benchmark (%.0fs elapsed) ---", benchmark.name, elapsed_time))
    
    local success, result = pcall(function()
      return benchmark.module['run_' .. benchmark.name .. '_benchmark'](config, fixture_dir)
    end)
    
    if success then
      all_results.benchmarks[benchmark.name] = result
      if result.success then
        print(string.format("✅ %s benchmark completed successfully", benchmark.name))
        
        -- Display statistical information
        if result.statistics then
          local stats = result.statistics
          print(string.format("   Stats: Mean=%.2fms, P95=%.2fms, P99=%.2fms, StdDev=%.2fms", 
            stats.mean, stats.p95, stats.p99, stats.stddev))
        elseif result.duration_ms then
          print(string.format("   Duration: %.2fms", result.duration_ms))
        end
        
        if result.success_rate then
          print(string.format("   Success rate: %.1f%% (%d/%d)", 
            result.success_rate * 100, 
            result.successful_runs or result.successful_hovers or result.successful_completions or 0,
            result.iterations or result.total_attempts or 1))
        end
        
        -- Enhanced benchmark specific reporting
        if result.total_buffers_opened then
          print(string.format("   Buffers opened: %d, Concurrent requests: %d", 
            result.total_buffers_opened, result.concurrent_requests_completed or 0))
        end
        
        if result.total_operations then
          print(string.format("   Total operations: %d", result.total_operations))
        end
        
        if result.total_edits then
          print(string.format("   Total edits: %d, Parsing responses: %d", 
            result.total_edits, result.parsing_responses or 0))
        end
        
        if result.files_opened then
          print(string.format("   Files opened: %d, Symbols found: %d", 
            result.files_opened, result.symbols_found or 0))
        end
        
        if result.memory_usage_bytes and result.memory_usage_bytes > 0 then
          print(string.format("   Memory usage: %.1fKB", result.memory_usage_bytes / 1024))
        end
      else
        print(string.format("❌ %s benchmark failed: %s", benchmark.name, result.error or "unknown error"))
      end
    else
      print(string.format("❌ %s benchmark crashed: %s", benchmark.name, result))
      all_results.benchmarks[benchmark.name] = {
        success = false,
        error = "benchmark_crashed: " .. tostring(result)
      }
    end
    
    -- Small delay between benchmarks
    vim.wait(1000)
  end
  
  -- Generate summary report
  print("\n" .. string.rep("=", 50))
  print(string.format("BENCHMARK SUMMARY FOR %s", string.upper(server_name)))
  print(string.rep("=", 50))
  
  local successful_benchmarks = 0
  local total_benchmarks = 0
  
  for benchmark_name, result in pairs(all_results.benchmarks) do
    total_benchmarks = total_benchmarks + 1
    if result.success then
      successful_benchmarks = successful_benchmarks + 1
      local status = "✅ PASS"
      local details = ""
      
      if result.duration_ms then
        details = details .. string.format(" (%.2fms)", result.duration_ms)
      end
      if result.success_rate then
        details = details .. string.format(" [%.0f%% success]", result.success_rate * 100)
      end
      
      print(string.format("%-20s %s%s", benchmark_name, status, details))
    else
      print(string.format("%-20s ❌ FAIL - %s", benchmark_name, result.error or "unknown"))
    end
  end
  
  local overall_success_rate = successful_benchmarks / total_benchmarks
  print(string.rep("-", 50))
  print(string.format("Overall Success Rate: %.1f%% (%d/%d)", 
    overall_success_rate * 100, successful_benchmarks, total_benchmarks))
  
  if overall_success_rate >= 0.8 then
    print("🎉 Server performance is GOOD")
  elseif overall_success_rate >= 0.6 then
    print("⚠️  Server performance is FAIR")
  else
    print("🚨 Server performance is POOR")
  end
  
  -- Save results to file
  local results_file = string.format('results/%s_modular_results_%s.json', 
    server_name, os.date("%Y%m%d_%H%M%S"))
  
  -- Ensure results directory exists
  vim.fn.mkdir(script_dir .. '/results', 'p')
  
  local json_content = vim.fn.json_encode(all_results)
  vim.fn.writefile({json_content}, script_dir .. '/' .. results_file)
  
  print(string.format("\nResults saved to: %s", results_file))
  
  -- Final timing report
  local total_elapsed_time = vim.fn.reltime(overall_start_time)[1]
  print(string.format("Total benchmark time: %.1fs / %.0fs limit", total_elapsed_time, max_total_time_seconds))
  
  print(string.rep("=", 50))
end

-- Run benchmarks
run_all_benchmarks()