#!/usr/bin/env nvim --headless --clean -l
-- Enhanced LSP Benchmark Runner (Real-world Performance Testing)
-- Usage: nvim --headless --clean -u configs/cem-minimal.lua -l run_enhanced_benchmarks.lua
-- Usage: nvim --headless --clean -u configs/wc-toolkit-minimal.lua -l run_enhanced_benchmarks.lua
-- 
-- This runner focuses on the new enhanced benchmarks that reveal architectural performance differences
-- Time limit: 120 seconds total (CI-friendly)

-- Ensure the benchmark directory is in the path
local script_dir = vim.fn.getcwd()
package.path = script_dir .. '/?.lua;' .. package.path

-- Import enhanced benchmark modules only
local multi_buffer_benchmark = require('modules.multi_buffer_benchmark')
local neovim_workflow_benchmark = require('modules.neovim_workflow_benchmark')
local incremental_parsing_benchmark = require('modules.incremental_parsing_benchmark')
local large_project_benchmark = require('modules.large_project_benchmark')

local function run_enhanced_benchmarks()
  local overall_start_time = vim.fn.reltime()
  local max_total_time_seconds = 120 -- 2 minutes total for CI
  
  local server_name = _G.BENCHMARK_LSP_NAME or "unknown"
  local config = _G.BENCHMARK_LSP_CONFIG
  
  if not config then
    print("ERROR: No LSP configuration found. Make sure to use the appropriate config file:")
    print("  nvim --headless --clean -u configs/cem-minimal.lua -l run_enhanced_benchmarks.lua")
    print("  nvim --headless --clean -u configs/wc-toolkit-minimal.lua -l run_enhanced_benchmarks.lua")
    return
  end
  
  print(string.format("Running ENHANCED benchmarks for %s LSP server (max %ds)", server_name, max_total_time_seconds))
  print("🚀 Real-world performance testing: multi-buffer, Neovim workflows, incremental parsing")
  print("=" .. string.rep("=", 70))
  
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
    benchmark_type = "enhanced_performance",
    benchmarks = {}
  }
  
  -- Enhanced benchmarks with time allocations
  local benchmarks = {
    {name = "multi_buffer", module = multi_buffer_benchmark, time_limit = 30},
    {name = "neovim_workflow", module = neovim_workflow_benchmark, time_limit = 40},
    {name = "incremental_parsing", module = incremental_parsing_benchmark, time_limit = 30},
    {name = "large_project", module = large_project_benchmark, time_limit = 20}
  }
  
  for _, benchmark in ipairs(benchmarks) do
    -- Check total time limit
    local elapsed_time = vim.fn.reltime(overall_start_time)[1]
    if elapsed_time >= max_total_time_seconds then
      print(string.format("\n⏰ Time limit reached (%ds), stopping remaining benchmarks", max_total_time_seconds))
      break
    end
    
    local remaining_time = max_total_time_seconds - elapsed_time
    print(string.format("\n--- Running %s benchmark (%.0fs elapsed, %.0fs remaining, %ds allocated) ---", 
      benchmark.name, elapsed_time, remaining_time, benchmark.time_limit))
    
    local success, result = pcall(function()
      return benchmark.module['run_' .. benchmark.name .. '_benchmark'](config, fixture_dir)
    end)
    
    if success then
      all_results.benchmarks[benchmark.name] = result
      if result.success then
        print(string.format("✅ %s benchmark completed successfully", benchmark.name))
        
        -- Display timing information
        if result.total_time then
          print(string.format("   Duration: %.2fs / %ds limit", result.total_time, benchmark.time_limit))
        end
        
        -- Display enhanced benchmark specific metrics
        if result.total_buffers_opened then
          print(string.format("   Multi-buffer: %d buffers, %d concurrent requests", 
            result.total_buffers_opened, result.concurrent_requests_completed or 0))
        end
        
        if result.total_operations and result.successful_completions then
          print(string.format("   Workflow: %d operations, %d completions", 
            result.total_operations, result.successful_completions))
        end
        
        if result.total_edits then
          print(string.format("   Incremental: %d edits, %d parsing responses", 
            result.total_edits, result.parsing_responses or 0))
        end
        
        if result.files_opened then
          print(string.format("   Project: %d files, %d symbols found", 
            result.files_opened, result.symbols_found or 0))
        end
        
        if result.memory_usage and result.memory_usage > 0 then
          print(string.format("   Memory: %s KB", result.memory_usage))
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
    vim.wait(500)
  end
  
  -- Generate summary report
  print("\n" .. string.rep("=", 70))
  print(string.format("ENHANCED BENCHMARK SUMMARY FOR %s", string.upper(server_name)))
  print(string.rep("=", 70))
  
  local successful_benchmarks = 0
  local total_benchmarks = 0
  
  for benchmark_name, result in pairs(all_results.benchmarks) do
    total_benchmarks = total_benchmarks + 1
    if result.success then
      successful_benchmarks = successful_benchmarks + 1
      local status = "✅ PASS"
      local details = ""
      
      if result.total_time then
        details = details .. string.format(" (%.1fs)", result.total_time)
      end
      
      print(string.format("%-20s %s%s", benchmark_name, status, details))
    else
      print(string.format("%-20s ❌ FAIL - %s", benchmark_name, result.error or "unknown"))
    end
  end
  
  local overall_success_rate = successful_benchmarks / total_benchmarks
  print(string.rep("-", 70))
  print(string.format("Enhanced Benchmark Success Rate: %.1f%% (%d/%d)", 
    overall_success_rate * 100, successful_benchmarks, total_benchmarks))
  
  -- Performance assessment
  if overall_success_rate >= 0.75 then
    print("🎉 Real-world performance is EXCELLENT")
    print("   Multi-buffer handling, workflows, and incremental parsing all working well")
  elseif overall_success_rate >= 0.5 then
    print("⚠️  Real-world performance is GOOD")
    print("   Some advanced features may need optimization")
  else
    print("🚨 Real-world performance needs IMPROVEMENT")
    print("   Consider architectural optimizations for better user experience")
  end
  
  -- Save results to file
  local results_file = string.format('results/%s_enhanced_results_%s.json', 
    server_name, os.date("%Y%m%d_%H%M%S"))
  
  -- Ensure results directory exists
  vim.fn.mkdir(script_dir .. '/results', 'p')
  
  local json_content = vim.fn.json_encode(all_results)
  vim.fn.writefile({json_content}, script_dir .. '/' .. results_file)
  
  print(string.format("\nEnhanced results saved to: %s", results_file))
  
  -- Final timing report
  local total_elapsed_time = vim.fn.reltime(overall_start_time)[1]
  print(string.format("Total enhanced benchmark time: %.1fs / %.0fs limit", total_elapsed_time, max_total_time_seconds))
  
  if total_elapsed_time > max_total_time_seconds then
    print("⚠️  Warning: Exceeded time limit - some benchmarks may have been skipped")
  end
  
  print(string.rep("=", 70))
end

-- Run enhanced benchmarks
run_enhanced_benchmarks()