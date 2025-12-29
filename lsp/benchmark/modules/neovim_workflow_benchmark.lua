local measurement = require('utils.measurement')

local M = {}

-- Neovim workflow benchmark: Tests realistic editor-driven workflows
-- Time limit: 40 seconds total  
-- Focus: End-to-end user experience latency from command to LSP response

function M.run_neovim_workflow_benchmark(config, fixture_dir)
  local start_time = vim.fn.reltime()
  local max_time_seconds = 40
  
  print("[NEOVIM_WORKFLOW] Starting Neovim workflow benchmark (40s max)")
  
  local results = {
    success = false,
    server_name = _G.BENCHMARK_LSP_NAME,
    error = nil,
    buffer_open_times = {},
    buffer_switch_times = {},
    completion_times = {},
    navigation_times = {},
    total_operations = 0,
    successful_completions = 0,
    memory_usage = 0
  }
  
  local success, error_msg = pcall(function()
    -- Phase 1: Realistic buffer opening workflow (10s)
    print("[NEOVIM_WORKFLOW] Phase 1: Opening buffers with :e commands")
    local files = {
      fixture_dir .. '/large-page.html',
      fixture_dir .. '/medium-page.html', 
      fixture_dir .. '/complex-page.html'
    }
    
    for i, file in ipairs(files) do
      if vim.fn.reltime(start_time)[1] >= max_time_seconds then
        break
      end
      
      local open_start = vim.fn.reltime()
      
      -- Simulate user typing :e command
      vim.cmd('edit ' .. file)
      vim.wait(500, function() return false end) -- Wait for LSP to process
      
      local open_time = tonumber(vim.fn.reltimestr(vim.fn.reltime(open_start)))
      table.insert(results.buffer_open_times, open_time)
      
      print(string.format("[NEOVIM_WORKFLOW] Opened %s in %.2fms", vim.fn.fnamemodify(file, ':t'), open_time))
      results.total_operations = results.total_operations + 1
    end
    
    -- Phase 2: Buffer navigation workflow (10s) 
    print("[NEOVIM_WORKFLOW] Phase 2: Buffer navigation with :b commands")
    local nav_start = vim.fn.reltime()
    local navigations = 0
    
    -- Simulate rapid buffer switching like developers do
    local buffer_commands = {':b1', ':b2', ':b3', ':b1', ':b3', ':b2'}
    
    while vim.fn.reltime(nav_start)[1] < 10 and vim.fn.reltime(start_time)[1] < max_time_seconds do
      for _, cmd in ipairs(buffer_commands) do
        if vim.fn.reltime(start_time)[1] >= max_time_seconds then
          break
        end
        
        local switch_start = vim.fn.reltime()
        
        -- Execute buffer switch command
        pcall(vim.cmd, cmd)
        vim.wait(200, function() return false end) -- Brief pause for buffer switch
        
        local switch_time = tonumber(vim.fn.reltimestr(vim.fn.reltime(switch_start)))
        table.insert(results.buffer_switch_times, switch_time)
        
        navigations = navigations + 1
        results.total_operations = results.total_operations + 1
        
        if navigations >= 12 then -- Limit to stay within time budget
          break
        end
      end
    end
    
    -- Phase 3: Insert mode completion workflow (15s)
    print("[NEOVIM_WORKFLOW] Phase 3: Insert mode completion with <C-x><C-o>")
    local completion_start = vim.fn.reltime()
    
    -- Test completion in various contexts
    local completion_scenarios = {
      {line = 45, col = 10, context = "inside element"},
      {line = 72, col = 15, context = "attribute position"},
      {line = 89, col = 8, context = "element start"},
      {line = 102, col = 20, context = "nested element"},
      {line = 156, col = 12, context = "template content"}
    }
    
    local completions = 0
    while vim.fn.reltime(completion_start)[1] < 15 and vim.fn.reltime(start_time)[1] < max_time_seconds do
      for _, scenario in ipairs(completion_scenarios) do
        if vim.fn.reltime(start_time)[1] >= max_time_seconds then
          break
        end
        
        -- Set cursor position
        pcall(vim.api.nvim_win_set_cursor, 0, {scenario.line, scenario.col})

        local comp_start = vim.fn.reltime()

        -- Enter insert mode and trigger completion like a user would
        vim.cmd('startinsert')
        vim.wait(50, function() return false end)

        -- Trigger LSP completion with Ctrl-X Ctrl-O
        local completion_received = false
        vim.schedule(function()
          pcall(vim.lsp.buf.completion)
          completion_received = true
        end)

        -- Wait for completion response
        vim.wait(500, function() return completion_received end)

        -- Additional wait to let async completion handler finish before exiting insert mode
        vim.wait(200, function() return false end)

        -- Exit insert mode
        vim.cmd('stopinsert')

        if completion_received then
          local comp_time = tonumber(vim.fn.reltimestr(vim.fn.reltime(comp_start)))
          table.insert(results.completion_times, comp_time)
          results.successful_completions = results.successful_completions + 1

          print(string.format("[NEOVIM_WORKFLOW] Completion in %s: %.2fms", scenario.context, comp_time))
        end
        
        completions = completions + 1
        results.total_operations = results.total_operations + 1
        
        if completions >= 10 then -- Limit to stay within time budget
          break
        end
      end
    end
    
    -- Phase 4: Navigation and hover workflow (5s)
    print("[NEOVIM_WORKFLOW] Phase 4: Navigation and hover actions")
    local nav_hover_start = vim.fn.reltime()
    
    while vim.fn.reltime(nav_hover_start)[1] < 5 and vim.fn.reltime(start_time)[1] < max_time_seconds do
      -- Move cursor and trigger hover like user would
      local positions = {{50, 12}, {75, 18}, {95, 6}}
      
      for _, pos in ipairs(positions) do
        if vim.fn.reltime(start_time)[1] >= max_time_seconds then
          break
        end
        
        local hover_start = vim.fn.reltime()
        
        -- Move cursor and trigger hover
        pcall(vim.api.nvim_win_set_cursor, 0, pos)
        vim.lsp.buf.hover()
        vim.wait(300, function() return false end)
        
        local nav_time = tonumber(vim.fn.reltimestr(vim.fn.reltime(hover_start)))
        table.insert(results.navigation_times, nav_time)
        results.total_operations = results.total_operations + 1
      end
    end
    
    -- Calculate statistics
    if #results.buffer_open_times > 0 then
      results.buffer_open_stats = measurement.calculate_statistics(results.buffer_open_times)
    end
    
    if #results.buffer_switch_times > 0 then
      results.buffer_switch_stats = measurement.calculate_statistics(results.buffer_switch_times)
    end
    
    if #results.completion_times > 0 then
      results.completion_stats = measurement.calculate_statistics(results.completion_times)
    end
    
    if #results.navigation_times > 0 then
      results.navigation_stats = measurement.calculate_statistics(results.navigation_times)
    end
    
    -- Memory usage
    results.memory_usage = vim.fn.system('ps -o rss= -p ' .. vim.fn.getpid()):gsub('%s+', '') or 0
    
    -- Success criteria: completed operations and some successful completions
    results.success = results.total_operations >= 15 and results.successful_completions >= 3
  end)
  
  if not success then
    results.error = error_msg
    print("[NEOVIM_WORKFLOW] Error: " .. tostring(error_msg))
  end
  
  local total_time = vim.fn.reltime(start_time)
  results.total_time = tonumber(vim.fn.reltimestr(total_time))
  
  print(string.format("[NEOVIM_WORKFLOW] Completed in %.2fs - Operations: %d, Completions: %d, Success: %s", 
    results.total_time, results.total_operations, results.successful_completions, tostring(results.success)))
    
  return results
end

return M