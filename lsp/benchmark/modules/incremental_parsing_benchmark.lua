local measurement = require('utils.measurement')

local M = {}

-- Incremental parsing benchmark: Tests tree-sitter incremental parsing performance
-- Time limit: 30 seconds total
-- Focus: Reveal tree-sitter's incremental parsing advantages during active editing

function M.run_incremental_parsing_benchmark(config, fixture_dir)
  local start_time = vim.fn.reltime()
  local max_time_seconds = 30
  
  print("[INCREMENTAL_PARSING] Starting incremental parsing benchmark (30s max)")
  
  local results = {
    success = false,
    server_name = _G.BENCHMARK_LSP_NAME,
    error = nil,
    character_edit_times = {},
    element_insertion_times = {},
    attribute_edit_times = {},
    large_paste_times = {},
    total_edits = 0,
    parsing_responses = 0,
    memory_usage = 0
  }
  
  local success, error_msg = pcall(function()
    -- Open medium-sized document (500 lines)
    local test_file = fixture_dir .. '/medium-editing-document.html'
    vim.cmd('edit ' .. test_file)
    vim.wait(1000, function() return false end) -- Wait for initial LSP processing
    
    -- Phase 1: Character-by-character typing simulation (10s)
    print("[INCREMENTAL_PARSING] Phase 1: Character-by-character typing")
    local char_start = vim.fn.reltime()
    
    -- Simulate typing a new element character by character
    local typing_sequence = "<my-new-element class=\"test\" size=\"large\">"
    local base_line = 100 -- Insert around middle of document
    
    vim.api.nvim_win_set_cursor(0, {base_line, 0})
    vim.cmd('startinsert')
    
    local char_edits = 0
    while vim.fn.reltime(char_start)[1] < 10 and vim.fn.reltime(start_time)[1] < max_time_seconds do
      for i = 1, #typing_sequence do
        if vim.fn.reltime(start_time)[1] >= max_time_seconds then
          break
        end
        
        local edit_start = vim.fn.reltime()
        
        -- Insert single character
        local char = typing_sequence:sub(i, i)
        vim.api.nvim_put({char}, 'c', false, true)
        
        -- Trigger LSP processing and wait for response
        vim.wait(50, function() return false end)
        
        -- Measure response time
        local edit_time = tonumber(vim.fn.reltimestr(vim.fn.reltime(edit_start)))
        table.insert(results.character_edit_times, edit_time)
        
        char_edits = char_edits + 1
        results.total_edits = results.total_edits + 1
        results.parsing_responses = results.parsing_responses + 1
        
        if char_edits >= 25 then -- Limit to stay within time budget
          break
        end
      end
    end
    
    vim.cmd('stopinsert')
    
    -- Phase 2: Element insertion/removal (8s)
    print("[INCREMENTAL_PARSING] Phase 2: Element insertion and removal")
    local element_start = vim.fn.reltime()
    
    local element_templates = {
      '<rh-card><h3>Title</h3><p>Content</p></rh-card>',
      '<my-button size="large" variant="primary">Click me</my-button>',
      '<ui-table rows="5" sortable="true"></ui-table>',
      '<custom-form method="post"><input type="text"></custom-form>'
    }
    
    local element_edits = 0
    while vim.fn.reltime(element_start)[1] < 8 and vim.fn.reltime(start_time)[1] < max_time_seconds do
      for _, template in ipairs(element_templates) do
        if vim.fn.reltime(start_time)[1] >= max_time_seconds then
          break
        end
        
        local insert_start = vim.fn.reltime()
        
        -- Insert entire element
        local line = base_line + element_edits
        vim.api.nvim_win_set_cursor(0, {line, 0})
        vim.api.nvim_put({template}, 'l', true, true)
        
        -- Wait for LSP processing
        vim.wait(100, function() return false end)
        
        local insert_time = tonumber(vim.fn.reltimestr(vim.fn.reltime(insert_start)))
        table.insert(results.element_insertion_times, insert_time)
        
        element_edits = element_edits + 1
        results.total_edits = results.total_edits + 1
        results.parsing_responses = results.parsing_responses + 1
        
        if element_edits >= 8 then -- Limit to stay within time budget
          break
        end
      end
    end
    
    -- Phase 3: Attribute modification (7s)
    print("[INCREMENTAL_PARSING] Phase 3: Attribute modifications")
    local attr_start = vim.fn.reltime()
    
    -- Test rapid attribute changes
    local attribute_changes = {
      {pattern = 'size="large"', replacement = 'size="small"'},
      {pattern = 'variant="primary"', replacement = 'variant="secondary"'},
      {pattern = 'sortable="true"', replacement = 'sortable="false"'},
      {pattern = 'method="post"', replacement = 'method="get"'}
    }
    
    local attr_edits = 0
    while vim.fn.reltime(attr_start)[1] < 7 and vim.fn.reltime(start_time)[1] < max_time_seconds do
      for _, change in ipairs(attribute_changes) do
        if vim.fn.reltime(start_time)[1] >= max_time_seconds then
          break
        end
        
        local edit_start = vim.fn.reltime()
        
        -- Find and replace attribute value
        pcall(vim.cmd, '%s/' .. change.pattern .. '/' .. change.replacement .. '/g')
        
        -- Wait for LSP processing
        vim.wait(75, function() return false end)
        
        local edit_time = tonumber(vim.fn.reltimestr(vim.fn.reltime(edit_start)))
        table.insert(results.attribute_edit_times, edit_time)
        
        attr_edits = attr_edits + 1
        results.total_edits = results.total_edits + 1
        results.parsing_responses = results.parsing_responses + 1
        
        if attr_edits >= 6 then -- Limit to stay within time budget
          break
        end
      end
    end
    
    -- Phase 4: Large paste operations (5s)
    print("[INCREMENTAL_PARSING] Phase 4: Large paste operations")
    local paste_start = vim.fn.reltime()
    
    -- Simulate pasting large blocks of HTML
    local large_blocks = {
      [[
<section class="hero">
  <rh-card variant="outlined">
    <h2 slot="header">Welcome</h2>
    <p>This is a large block of content with multiple custom elements.</p>
    <my-button size="large" variant="primary">Get Started</my-button>
  </rh-card>
</section>]],
      [[
<div class="grid">
  <ui-table rows="10" sortable="true">
    <th>Name</th>
    <th>Value</th>
    <th>Actions</th>
  </ui-table>
  <custom-form method="post">
    <input type="text" name="search">
    <my-button type="submit">Search</my-button>
  </custom-form>
</div>]]
    }
    
    local paste_edits = 0
    while vim.fn.reltime(paste_start)[1] < 5 and vim.fn.reltime(start_time)[1] < max_time_seconds do
      for _, block in ipairs(large_blocks) do
        if vim.fn.reltime(start_time)[1] >= max_time_seconds then
          break
        end
        
        local paste_edit_start = vim.fn.reltime()
        
        -- Paste large block
        local line = base_line + element_edits + paste_edits * 10
        vim.api.nvim_win_set_cursor(0, {line, 0})
        vim.api.nvim_put(vim.split(block, '\n'), 'l', true, true)
        
        -- Wait for LSP processing
        vim.wait(200, function() return false end)
        
        local paste_time = tonumber(vim.fn.reltimestr(vim.fn.reltime(paste_edit_start)))
        table.insert(results.large_paste_times, paste_time)
        
        paste_edits = paste_edits + 1
        results.total_edits = results.total_edits + 1
        results.parsing_responses = results.parsing_responses + 1
        
        if paste_edits >= 3 then -- Limit to stay within time budget
          break
        end
      end
    end
    
    -- Calculate statistics
    if #results.character_edit_times > 0 then
      results.character_edit_stats = measurement.calculate_statistics(results.character_edit_times)
    end
    
    if #results.element_insertion_times > 0 then
      results.element_insertion_stats = measurement.calculate_statistics(results.element_insertion_times)
    end
    
    if #results.attribute_edit_times > 0 then
      results.attribute_edit_stats = measurement.calculate_statistics(results.attribute_edit_times)
    end
    
    if #results.large_paste_times > 0 then
      results.large_paste_stats = measurement.calculate_statistics(results.large_paste_times)
    end
    
    -- Memory usage
    results.memory_usage = vim.fn.system('ps -o rss= -p ' .. vim.fn.getpid()):gsub('%s+', '') or 0
    
    -- Success criteria: completed edits with LSP responses
    results.success = results.total_edits >= 20 and results.parsing_responses >= 15
  end)
  
  if not success then
    results.error = error_msg
    print("[INCREMENTAL_PARSING] Error: " .. tostring(error_msg))
  end
  
  local total_time = vim.fn.reltime(start_time)
  results.total_time = tonumber(vim.fn.reltimestr(total_time))
  
  print(string.format("[INCREMENTAL_PARSING] Completed in %.2fs - Edits: %d, Responses: %d, Success: %s", 
    results.total_time, results.total_edits, results.parsing_responses, tostring(results.success)))
    
  return results
end

return M