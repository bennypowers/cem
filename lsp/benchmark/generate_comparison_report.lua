#!/usr/bin/env nvim --headless --clean -l
-- Generate statistical comparison report from benchmark results
-- Usage: nvim --headless --clean -l generate_comparison_report.lua

local function read_json_file(filepath)
  local file = io.open(filepath, 'r')
  if not file then
    return nil, "Could not open file: " .. filepath
  end
  
  local content = file:read('*all')
  file:close()
  
  local success, result = pcall(vim.fn.json_decode, content)
  if not success then
    return nil, "Invalid JSON in file: " .. filepath
  end
  
  return result
end

local function find_latest_results()
  local results = {}
  
  -- Find cem results
  local cem_pattern = "results/cem_modular_results_*.json"
  local cem_files = vim.fn.glob(cem_pattern, false, true)
  
  if #cem_files > 0 then
    -- Sort by filename (which includes timestamp) and take the latest
    table.sort(cem_files)
    local cem_data, err = read_json_file(cem_files[#cem_files])
    if cem_data then
      results.cem = {
        filepath = cem_files[#cem_files],
        data = cem_data
      }
    else
      print("Error reading cem results: " .. (err or "unknown"))
    end
  end
  
  -- Find wc-toolkit results
  local wct_pattern = "results/wc-toolkit_modular_results_*.json"
  local wct_files = vim.fn.glob(wct_pattern, false, true)
  
  if #wct_files > 0 then
    -- Sort by filename and take the latest
    table.sort(wct_files)
    local wct_data, err = read_json_file(wct_files[#wct_files])
    if wct_data then
      results.wc_toolkit = {
        filepath = wct_files[#wct_files],
        data = wct_data
      }
    else
      print("Error reading wc-toolkit results: " .. (err or "unknown"))
    end
  end
  
  return results
end

local function format_statistics(stats, name)
  if not stats or stats.count == 0 then
    return string.format("%-12s: No data", name)
  end
  
  return string.format("%-12s: Mean=%.2fms, P95=%.2fms, P99=%.2fms, StdDev=%.2fms", 
    name, stats.mean, stats.p95, stats.p99, stats.stddev)
end

local function compare_benchmark(cem_result, wct_result, benchmark_name)
  print(string.format("\n%s BENCHMARK COMPARISON", string.upper(benchmark_name)))
  print(string.rep("=", 60))
  
  if not cem_result and not wct_result then
    print("No results available for either server")
    return
  end
  
  if not cem_result then
    print("cem: No results available")
  else
    print(string.format("cem: Success=%s, Success Rate=%.1f%%", 
      cem_result.success and "✅" or "❌",
      (cem_result.success_rate or 0) * 100))
    
    if cem_result.statistics then
      print("    " .. format_statistics(cem_result.statistics, "cem"))
    elseif cem_result.duration_ms then
      print(string.format("    Duration: %.2fms", cem_result.duration_ms))
    end
  end
  
  if not wct_result then
    print("wc-toolkit: No results available")
  else
    print(string.format("wc-toolkit: Success=%s, Success Rate=%.1f%%", 
      wct_result.success and "✅" or "❌",
      (wct_result.success_rate or 0) * 100))
    
    if wct_result.statistics then
      print("    " .. format_statistics(wct_result.statistics, "wc-toolkit"))
    elseif wct_result.duration_ms then
      print(string.format("    Duration: %.2fms", wct_result.duration_ms))
    end
  end
  
  -- Performance comparison
  if cem_result and wct_result and cem_result.statistics and wct_result.statistics then
    local cem_mean = cem_result.statistics.mean
    local wct_mean = wct_result.statistics.mean
    
    if cem_mean < wct_mean then
      local improvement = ((wct_mean - cem_mean) / wct_mean) * 100
      print(string.format("🏆 cem is %.1f%% faster (%.2fms vs %.2fms)", 
        improvement, cem_mean, wct_mean))
    elseif wct_mean < cem_mean then
      local improvement = ((cem_mean - wct_mean) / cem_mean) * 100
      print(string.format("🏆 wc-toolkit is %.1f%% faster (%.2fms vs %.2fms)", 
        improvement, wct_mean, cem_mean))
    else
      print("⚖️  Performance is equivalent")
    end
  elseif cem_result and wct_result and cem_result.duration_ms and wct_result.duration_ms then
    local cem_time = cem_result.duration_ms
    local wct_time = wct_result.duration_ms
    
    if cem_time < wct_time then
      local improvement = ((wct_time - cem_time) / wct_time) * 100
      print(string.format("🏆 cem is %.1f%% faster (%.2fms vs %.2fms)", 
        improvement, cem_time, wct_time))
    elseif wct_time < cem_time then
      local improvement = ((cem_time - wct_time) / cem_time) * 100
      print(string.format("🏆 wc-toolkit is %.1f%% faster (%.2fms vs %.2fms)", 
        improvement, wct_time, cem_time))
    else
      print("⚖️  Performance is equivalent")
    end
  end
end

local function generate_summary_report(results)
  print("\n" .. string.rep("=", 80))
  print("COMPREHENSIVE LSP SERVER COMPARISON REPORT")
  print(string.rep("=", 80))
  
  if results.cem then
    print(string.format("cem results from: %s (%s)", 
      results.cem.filepath, results.cem.data.timestamp))
  end
  
  if results.wc_toolkit then
    print(string.format("wc-toolkit results from: %s (%s)", 
      results.wc_toolkit.filepath, results.wc_toolkit.data.timestamp))
  end
  
  local benchmark_names = {"startup", "hover", "completion", "diagnostics", "file_lifecycle", "stress_test"}
  
  for _, benchmark_name in ipairs(benchmark_names) do
    local cem_result = results.cem and results.cem.data.benchmarks[benchmark_name]
    local wct_result = results.wc_toolkit and results.wc_toolkit.data.benchmarks[benchmark_name]
    
    if cem_result or wct_result then
      compare_benchmark(cem_result, wct_result, benchmark_name)
    end
  end
  
  -- Overall summary
  print(string.format("\n%s", string.rep("=", 80)))
  print("OVERALL PERFORMANCE SUMMARY")
  print(string.rep("=", 80))
  
  local cem_wins = 0
  local wct_wins = 0
  local ties = 0
  
  for _, benchmark_name in ipairs(benchmark_names) do
    local cem_result = results.cem and results.cem.data.benchmarks[benchmark_name]
    local wct_result = results.wc_toolkit and results.wc_toolkit.data.benchmarks[benchmark_name]
    
    if cem_result and wct_result and cem_result.success and wct_result.success then
      local cem_time = cem_result.statistics and cem_result.statistics.mean or cem_result.duration_ms
      local wct_time = wct_result.statistics and wct_result.statistics.mean or wct_result.duration_ms
      
      if cem_time and wct_time then
        if cem_time < wct_time * 0.95 then -- 5% threshold
          cem_wins = cem_wins + 1
        elseif wct_time < cem_time * 0.95 then
          wct_wins = wct_wins + 1
        else
          ties = ties + 1
        end
      end
    end
  end
  
  print(string.format("Performance Comparison: cem wins %d, wc-toolkit wins %d, ties %d", 
    cem_wins, wct_wins, ties))
  
  if cem_wins > wct_wins then
    print("🏆 Overall Winner: cem LSP")
  elseif wct_wins > cem_wins then
    print("🏆 Overall Winner: wc-toolkit LSP")
  else
    print("⚖️  Overall Result: Tie")
  end
  
  -- Reliability comparison
  local cem_reliability = 0
  local wct_reliability = 0
  local total_benchmarks = 0
  
  for _, benchmark_name in ipairs(benchmark_names) do
    local cem_result = results.cem and results.cem.data.benchmarks[benchmark_name]
    local wct_result = results.wc_toolkit and results.wc_toolkit.data.benchmarks[benchmark_name]
    
    if cem_result then
      total_benchmarks = total_benchmarks + 1
      if cem_result.success then
        cem_reliability = cem_reliability + 1
      end
    end
    
    if wct_result then
      if wct_result.success then
        wct_reliability = wct_reliability + 1
      end
    end
  end
  
  if total_benchmarks > 0 then
    local cem_reliability_pct = (cem_reliability / total_benchmarks) * 100
    local wct_reliability_pct = (wct_reliability / total_benchmarks) * 100
    
    print(string.format("Reliability: cem %.1f%%, wc-toolkit %.1f%%", 
      cem_reliability_pct, wct_reliability_pct))
  end
  
  print(string.rep("=", 80))
end

-- Main execution
local results = find_latest_results()

if not results.cem and not results.wc_toolkit then
  print("ERROR: No benchmark results found!")
  print("Please run ./run_comparison.sh first to generate benchmark data.")
  return
end

generate_summary_report(results)

-- Save report to file
local timestamp = os.date("%Y%m%d_%H%M%S")
local report_file = string.format("results/comparison_report_%s.txt", timestamp)

-- Redirect output to file (simplified approach)
print(string.format("\nComparison report generated at: %s", os.date("%Y-%m-%d %H:%M:%S")))
print("To save this report, run:")
print(string.format("nvim --headless --clean -l generate_comparison_report.lua > %s", report_file))