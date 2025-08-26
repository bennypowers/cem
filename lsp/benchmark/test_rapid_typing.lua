#!/usr/bin/env nvim --headless --clean -l
-- Quick test of rapid typing functionality
-- Usage: nvim --headless --clean -u configs/cem-minimal.lua -l test_rapid_typing.lua

-- Ensure the benchmark directory is in the path
local script_dir = vim.fn.getcwd()
package.path = script_dir .. '/?.lua;' .. package.path

local rapid_typing_benchmark = require('modules.rapid_typing_benchmark')

local function quick_typing_test()
  local server_name = _G.BENCHMARK_LSP_NAME or "unknown"
  local config = _G.BENCHMARK_LSP_CONFIG
  
  if not config then
    print("ERROR: No LSP configuration found.")
    return
  end
  
  print(string.format("Quick rapid typing test for %s LSP server", server_name))
  print("Testing 30 seconds of rapid typing simulation...")
  
  local fixture_dir = script_dir .. '/fixtures/large_project'
  
  if vim.fn.isdirectory(fixture_dir) == 0 then
    print("ERROR: Fixture directory not found: " .. fixture_dir)
    return
  end
  
  -- Override the benchmark time limit to be shorter
  _G.BENCHMARK_TIME_LIMIT = 30
  
  local start_time = vim.fn.reltime()
  local success, result = pcall(function()
    return rapid_typing_benchmark.run_rapid_typing_benchmark(config, fixture_dir)
  end)
  local elapsed_time = vim.fn.reltime(start_time)[1]
  
  print(string.format("Test completed in %.1fs", elapsed_time))
  
  if success then
    print("✅ Test successful")
    if result.success then
      print(string.format("Characters typed: %d", result.total_characters_typed or 0))
      print(string.format("Completions: %d", result.completion_responses or 0))
      print(string.format("Duration: %.2fs", result.total_time or 0))
      if result.avg_typing_speed then
        print(string.format("Average speed: %.1f WPM", result.avg_typing_speed))
      end
    else
      print(string.format("❌ Test failed: %s", result.error or "unknown"))
    end
  else
    print(string.format("❌ Test crashed: %s", result))
  end
end

quick_typing_test()