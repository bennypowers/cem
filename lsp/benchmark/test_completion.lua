#!/usr/bin/env nvim --headless --clean -l
-- Quick test of completion functionality
-- Usage: nvim --headless --clean -u configs/cem-minimal.lua -l test_completion.lua

-- Ensure the benchmark directory is in the path
local script_dir = vim.fn.getcwd()
package.path = script_dir .. '/?.lua;' .. package.path

local function test_completion()
  local server_name = _G.BENCHMARK_LSP_NAME or "unknown"
  local config = _G.BENCHMARK_LSP_CONFIG
  
  if not config then
    print("ERROR: No LSP configuration found.")
    return
  end
  
  print(string.format("Testing completion for %s LSP server", server_name))
  
  local fixture_dir = script_dir .. '/fixtures/large_project'
  local test_file = fixture_dir .. '/rapid-typing-test.html'
  
  if vim.fn.filereadable(test_file) == 0 then
    print("ERROR: Test file not found: " .. test_file)
    return
  end
  
  local results = {
    success = false,
    completion_tests = 0,
    successful_completions = 0,
    avg_response_time = 0,
    response_times = {},
    errors = {}
  }
  
  local success, error_msg = pcall(function()
    -- Open test document
    vim.cmd('edit ' .. test_file)
    vim.wait(3000, function() return false end) -- Wait for LSP initialization
    
    print("LSP initialized, testing completion scenarios...")
    
    -- Clear document 
    vim.cmd('normal! ggdG')
    
    -- Test 1: Tag name completion
    print("Test 1: Tag name completion")
    vim.api.nvim_buf_set_lines(0, 0, -1, false, {'<my-'})
    vim.api.nvim_win_set_cursor(0, {1, 4}) -- Position after 'my-'
    
    results.completion_tests = results.completion_tests + 1
    local completion_start = vim.fn.reltime()
    
    -- Use vim.lsp.buf.completion with completion callback
    local completion_received = false
    vim.lsp.handlers['textDocument/completion'] = function(err, result, ctx, config)
      completion_received = true
      if err then
        table.insert(results.errors, "Completion error: " .. tostring(err))
      elseif result and (result.items or result) then
        local items = result.items or result
        if #items > 0 then
          results.successful_completions = results.successful_completions + 1
          print(string.format("  ✅ Found %d completion items", #items))
        else
          print("  ⚠️  No completion items returned")
        end
      else
        print("  ❌ No completion result")
      end
    end
    
    vim.lsp.buf.completion()
    
    -- Wait for completion response
    local timeout = 5000 -- 5 second timeout
    local waited = vim.wait(timeout, function() return completion_received end, 100)
    
    local completion_time = tonumber(vim.fn.reltimestr(vim.fn.reltime(completion_start)))
    table.insert(results.response_times, completion_time)
    
    if not waited then
      table.insert(results.errors, "Completion timed out after 5s")
      print("  ❌ Completion timed out")
    end
    
    -- Test 2: Attribute completion  
    print("Test 2: Attribute completion")
    vim.cmd('normal! ggdG')
    vim.api.nvim_buf_set_lines(0, 0, -1, false, {'<my-button variant="'})
    vim.api.nvim_win_set_cursor(0, {1, 19}) -- Position after 'variant="'
    
    results.completion_tests = results.completion_tests + 1
    completion_received = false
    
    completion_start = vim.fn.reltime()
    vim.lsp.buf.completion()
    
    waited = vim.wait(timeout, function() return completion_received end, 100)
    completion_time = tonumber(vim.fn.reltimestr(vim.fn.reltime(completion_start)))
    table.insert(results.response_times, completion_time)
    
    if not waited then
      table.insert(results.errors, "Attribute completion timed out")
      print("  ❌ Attribute completion timed out")
    end
    
    -- Calculate average response time
    if #results.response_times > 0 then
      local total_time = 0
      for _, time in ipairs(results.response_times) do
        total_time = total_time + time
      end
      results.avg_response_time = total_time / #results.response_times
    end
    
    results.success = true
  end)
  
  if success then
    print("\n" .. string.rep("=", 50))
    print("COMPLETION TEST SUMMARY")
    print(string.rep("=", 50))
    print(string.format("Completion tests: %d", results.completion_tests))
    print(string.format("Successful completions: %d", results.successful_completions))
    print(string.format("Average response time: %.3fs", results.avg_response_time))
    
    if #results.errors > 0 then
      print("\nErrors:")
      for _, error in ipairs(results.errors) do
        print("  " .. error)
      end
    end
    
    local success_rate = results.successful_completions / math.max(1, results.completion_tests)
    print(string.format("Completion success rate: %.1f%%", success_rate * 100))
    
    if success_rate >= 0.8 and results.avg_response_time < 1.0 then
      print("🎉 Completion performance is EXCELLENT")
    elseif success_rate >= 0.5 and results.avg_response_time < 2.0 then
      print("⚠️  Completion performance is GOOD")
    else
      print("🚨 Completion performance needs improvement")
    end
  else
    print(string.format("❌ Test failed: %s", error_msg))
  end
end

test_completion()