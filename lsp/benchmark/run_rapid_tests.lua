#!/usr/bin/env nvim --headless --clean -l
-- Rapid Typing and References Benchmark Runner
-- Usage: nvim --headless --clean -u configs/cem-minimal.lua -l run_rapid_tests.lua
-- Usage: nvim --headless --clean -u configs/wc-toolkit-minimal.lua -l run_rapid_tests.lua

-- Ensure the benchmark directory is in the path
local script_dir = vim.fn.getcwd()
package.path = script_dir .. '/?.lua;' .. package.path

-- Import the new benchmark modules
local rapid_typing_benchmark = require('modules.rapid_typing_benchmark')
local references_performance_benchmark = require('modules.references_performance_benchmark')
local integrated_workflow_benchmark = require('modules.integrated_workflow_benchmark')

local function run_rapid_tests()
  local overall_start_time = vim.fn.reltime()
  local max_total_time_seconds = 60 -- 1 minute total
  
  local server_name = _G.BENCHMARK_LSP_NAME or "unknown"
  local config = _G.BENCHMARK_LSP_CONFIG
  
  if not config then
    print("ERROR: No LSP configuration found. Make sure to use the appropriate config file:")
    print("  nvim --headless --clean -u configs/cem-minimal.lua -l run_rapid_tests.lua")
    print("  nvim --headless --clean -u configs/wc-toolkit-minimal.lua -l run_rapid_tests.lua")
    return
  end
  
  print(string.format("Running RAPID TYPING tests for %s LSP server (max %ds)", server_name, max_total_time_seconds))
  print("🚀 Testing realistic typing speeds at 60-80 WPM with go-to-references")
  print("=" .. string.rep("=", 70))
  
  -- Use large project fixture
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
    benchmark_type = "rapid_typing_performance",
    benchmarks = {}
  }
  
  -- Rapid typing focused benchmarks
  local benchmarks = {
    {name = "rapid_typing", module = rapid_typing_benchmark, time_limit = 25},
    {name = "references_performance", module = references_performance_benchmark, time_limit = 20},
    {name = "integrated_workflow", module = integrated_workflow_benchmark, time_limit = 15}
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
        
        -- Display rapid typing specific metrics
        if result.total_characters_typed then
          print(string.format("   Rapid Typing: %d chars, %d completions", 
            result.total_characters_typed, result.completion_responses or 0))
          if result.avg_typing_speed then
            print(string.format("   Average Speed: %.1f WPM", result.avg_typing_speed))
          end
        end
        
        if result.total_references_searched then
          print(string.format("   References: %d searches, %d refs found", 
            result.total_references_searched, result.total_references_found or 0))
          if result.avg_references_per_element then
            print(string.format("   Average refs per element: %.1f", result.avg_references_per_element))
          end
        end
        
        if result.workflow_iterations then
          print(string.format("   Integrated: %d workflows, %d typing sessions", 
            result.workflow_iterations, result.typing_sessions or 0))
        end
        
        if result.memory_usage and tonumber(result.memory_usage) and tonumber(result.memory_usage) > 0 then
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
  print(string.format("RAPID TYPING BENCHMARK SUMMARY FOR %s", string.upper(server_name)))
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
  print(string.format("Rapid Typing Success Rate: %.1f%% (%d/%d)", 
    overall_success_rate * 100, successful_benchmarks, total_benchmarks))
  
  -- Performance assessment
  if overall_success_rate >= 1.0 then
    print("🎉 Rapid typing performance is EXCELLENT")
    print("   All typing speeds and references working smoothly")
  elseif overall_success_rate >= 0.67 then
    print("⚠️  Rapid typing performance is GOOD")
    print("   Most features working but some may need optimization")
  else
    print("🚨 Rapid typing performance needs IMPROVEMENT")
    print("   Typing or references features may be too slow for real-world use")
  end
  
  -- Save results to file
  local results_file = string.format('results/%s_rapid_typing_%s.json', 
    server_name, os.date("%Y%m%d_%H%M%S"))
  
  -- Ensure results directory exists
  vim.fn.mkdir(script_dir .. '/results', 'p')
  
  local json_content = vim.fn.json_encode(all_results)
  vim.fn.writefile({json_content}, script_dir .. '/' .. results_file)
  
  print(string.format("\nRapid typing results saved to: %s", results_file))
  
  -- Final timing report
  local total_elapsed_time = vim.fn.reltime(overall_start_time)[1]
  print(string.format("Total rapid typing test time: %.1fs / %.0fs limit", total_elapsed_time, max_total_time_seconds))
  
  print(string.rep("=", 70))
end

-- Run rapid typing tests
run_rapid_tests()