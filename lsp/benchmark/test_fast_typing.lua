#!/usr/bin/env nvim --headless --clean -l
-- Fast typing test for benchmarking LSP responsiveness
-- Usage: nvim --headless --clean -u configs/cem-minimal.lua -l test_fast_typing.lua

-- Ensure the benchmark directory is in the path
local script_dir = vim.fn.getcwd()
package.path = script_dir .. '/?.lua;' .. package.path

local function fast_typing_test()
  local server_name = _G.BENCHMARK_LSP_NAME or "unknown"
  local config = _G.BENCHMARK_LSP_CONFIG
  
  if not config then
    print("ERROR: No LSP configuration found.")
    return
  end
  
  print(string.format("Fast typing test for %s LSP server", server_name))
  print("Testing LSP responsiveness during rapid text input...")
  
  local fixture_dir = script_dir .. '/fixtures/large_project'
  local test_file = fixture_dir .. '/rapid-typing-test.html'
  
  if vim.fn.filereadable(test_file) == 0 then
    print("ERROR: Test file not found: " .. test_file)
    return
  end
  
  local start_time = vim.fn.reltime()
  local results = {
    success = false,
    characters_typed = 0,
    completion_requests = 0,
    completion_responses = 0,
    errors = {}
  }
  
  local success, error_msg = pcall(function()
    -- Open test document
    vim.cmd('edit ' .. test_file)
    vim.wait(2000, function() return false end) -- Wait for LSP initialization
    
    print("LSP initialized, starting typing test...")
    
    -- Clear document and start typing
    vim.cmd('normal! ggdG')
    
    -- Test rapid typing with completions
    local test_text = '<my-card variant="outlined" size="large">\n  <h3 slot="header">Title</h3>\n  <my-button variant="primary">Click</my-button>\n</my-card>'
    
    local completion_triggers = {10, 20, 35, 55, 75, 95} -- Trigger completions at these positions
    
    -- Type character by character with fast timing
    for i = 1, #test_text do
      local char = test_text:sub(i, i)
      
      -- Insert character
      vim.api.nvim_put({char}, 'c', false, true)
      results.characters_typed = results.characters_typed + 1
      
      -- Check if we should trigger completion
      local should_complete = false
      for _, pos in ipairs(completion_triggers) do
        if i == pos then
          should_complete = true
          break
        end
      end
      
      if should_complete then
        results.completion_requests = results.completion_requests + 1
        
        -- Trigger completion
        local completion_start = vim.fn.reltime()
        local completion_success = pcall(function()
          vim.lsp.buf.completion()
        end)
        
        if completion_success then
          results.completion_responses = results.completion_responses + 1
        else
          table.insert(results.errors, string.format("Completion failed at position %d", i))
        end
      end
      
      -- Small delay for LSP processing (much faster than realistic typing)
      vim.wait(10, function() return false end)
    end
    
    -- Wait for any pending LSP responses
    vim.wait(1000, function() return false end)
    
    results.success = true
  end)
  
  local elapsed_time = vim.fn.reltime(start_time)[1]
  
  if success then
    print("✅ Fast typing test completed successfully")
    print(string.format("Duration: %.2fs", elapsed_time))
    print(string.format("Characters typed: %d", results.characters_typed))
    print(string.format("Completion requests: %d", results.completion_requests))
    print(string.format("Completion responses: %d", results.completion_responses))
    
    if #results.errors > 0 then
      print("Errors:")
      for _, error in ipairs(results.errors) do
        print("  " .. error)
      end
    end
    
    local completion_rate = results.completion_responses / math.max(1, results.completion_requests)
    print(string.format("Completion success rate: %.1f%%", completion_rate * 100))
    
    if completion_rate >= 0.8 then
      print("🎉 LSP is very responsive during fast typing")
    elseif completion_rate >= 0.5 then
      print("⚠️  LSP responsiveness is moderate")
    else
      print("🚨 LSP may be too slow for fast typing")
    end
  else
    print(string.format("❌ Test failed: %s", error_msg))
  end
end

fast_typing_test()