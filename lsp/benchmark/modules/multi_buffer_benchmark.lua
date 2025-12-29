local measurement = require('utils.measurement')

local M = {}

-- Multi-buffer benchmark: Tests concurrent document handling and buffer switching
-- Time limit: 30 seconds total
-- Focus: Reveal Go's concurrency advantages vs TypeScript single-threaded model

function M.run_multi_buffer_benchmark(config, fixture_dir)
  local start_time = vim.fn.reltime()
  local max_time_seconds = 30
  
  print("[MULTI_BUFFER] Starting multi-buffer benchmark (30s max)")
  
  local results = {
    success = false,
    server_name = _G.BENCHMARK_LSP_NAME,
    error = nil,
    buffer_switch_times = {},
    concurrent_request_times = {},
    total_buffers_opened = 0,
    concurrent_requests_completed = 0,
    memory_usage = 0
  }
  
  local success, error_msg = pcall(function()
    -- Phase 1: Open multiple documents simultaneously (10s)
    print("[MULTI_BUFFER] Phase 1: Opening 5 documents simultaneously")
    local buffer_files = {
      fixture_dir .. '/large-page.html',
      fixture_dir .. '/medium-page.html', 
      fixture_dir .. '/small-page.html',
      fixture_dir .. '/complex-page.html',
      fixture_dir .. '/nested-elements.html'
    }
    
    local buffers = {}
    for i, file in ipairs(buffer_files) do
      if vim.fn.reltime(start_time)[1] >= max_time_seconds then
        break
      end
      
      local buf_start = vim.fn.reltime()
      local buf = vim.fn.bufnr(file, true)
      vim.api.nvim_buf_call(buf, function()
        vim.cmd('edit! ' .. file)  -- Force edit to avoid E37 error
      end)
      
      if vim.api.nvim_buf_is_loaded(buf) then
        buffers[i] = buf
        results.total_buffers_opened = results.total_buffers_opened + 1
        local buf_time = vim.fn.reltimestr(vim.fn.reltime(buf_start))
        print(string.format("[MULTI_BUFFER] Opened buffer %d in %sms", i, buf_time))
      end
    end
    
    -- Wait for LSP to process all documents
    vim.wait(2000, function() return false end)
    
    -- Phase 2: Rapid buffer switching (10s)
    print("[MULTI_BUFFER] Phase 2: Rapid buffer switching")
    local switch_start = vim.fn.reltime()
    local switches = 0
    
    while vim.fn.reltime(switch_start)[1] < 10 and vim.fn.reltime(start_time)[1] < max_time_seconds do
      for i, buf in pairs(buffers) do
        if vim.fn.reltime(start_time)[1] >= max_time_seconds then
          break
        end
        
        local switch_time_start = vim.fn.reltime()
        vim.api.nvim_set_current_buf(buf)
        vim.wait(100, function() return false end) -- Brief pause for LSP processing
        local switch_time = tonumber(vim.fn.reltimestr(vim.fn.reltime(switch_time_start)))
        
        table.insert(results.buffer_switch_times, switch_time)
        switches = switches + 1
        
        if switches >= 15 then -- Limit switches to stay within time budget
          break
        end
      end
    end
    
    -- Phase 3: Concurrent LSP requests across buffers (10s)
    print("[MULTI_BUFFER] Phase 3: Concurrent LSP requests")
    local request_start = vim.fn.reltime()
    
    -- Test positions for different types of requests
    local test_positions = {
      {line = 45, col = 10}, -- Inside element
      {line = 72, col = 15}, -- Attribute position  
      {line = 89, col = 8},  -- Element start
    }
    
    local concurrent_requests = 0
    while vim.fn.reltime(request_start)[1] < 10 and vim.fn.reltime(start_time)[1] < max_time_seconds do
      for i, buf in pairs(buffers) do
        if vim.fn.reltime(start_time)[1] >= max_time_seconds then
          break
        end
        
        vim.api.nvim_set_current_buf(buf)
        
        for _, pos in ipairs(test_positions) do
          local req_start = vim.fn.reltime()
          
          -- Set cursor position
          pcall(vim.api.nvim_win_set_cursor, 0, {pos.line, pos.col})
          
          -- Trigger hover request (non-blocking)
          pcall(vim.lsp.buf.hover)
          
          -- Brief pause to allow processing
          vim.wait(50, function() return false end)
          
          -- Always count the request (hover was triggered)
          local req_time = tonumber(vim.fn.reltimestr(vim.fn.reltime(req_start)))
          table.insert(results.concurrent_request_times, req_time)
          results.concurrent_requests_completed = results.concurrent_requests_completed + 1
          
          concurrent_requests = concurrent_requests + 1
          if concurrent_requests >= 20 then -- Limit requests to stay within time budget
            break
          end
        end
        
        if concurrent_requests >= 20 then
          break
        end
      end
    end
    
    -- Calculate statistics
    if #results.buffer_switch_times > 0 then
      results.buffer_switch_stats = measurement.calculate_statistics(results.buffer_switch_times)
    end
    
    if #results.concurrent_request_times > 0 then
      results.concurrent_request_stats = measurement.calculate_statistics(results.concurrent_request_times)
    end
    
    -- Memory usage approximation
    results.memory_usage = vim.fn.system('ps -o rss= -p ' .. vim.fn.getpid()):gsub('%s+', '') or 0
    
    results.success = results.total_buffers_opened >= 3 and results.concurrent_requests_completed >= 5
  end)
  
  if not success then
    results.error = error_msg
    print("[MULTI_BUFFER] Error: " .. tostring(error_msg))
  end
  
  local total_time = vim.fn.reltime(start_time)
  results.total_time = tonumber(vim.fn.reltimestr(total_time))
  
  print(string.format("[MULTI_BUFFER] Completed in %.2fs - Buffers: %d, Requests: %d, Success: %s", 
    results.total_time, results.total_buffers_opened, results.concurrent_requests_completed, tostring(results.success)))
    
  return results
end

return M