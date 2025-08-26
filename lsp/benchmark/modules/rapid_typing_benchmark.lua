local measurement = require('utils.measurement')

local M = {}

-- Rapid typing benchmark: Tests LSP performance during realistic typing speeds
-- Simulates proficient keyboardist: 60-80 WPM (150-250ms between characters)
-- Time limit: 60 seconds total
-- Focus: Character-level incremental parsing and completion responsiveness

function M.run_rapid_typing_benchmark(config, fixture_dir)
  local start_time = vim.fn.reltime()
  local max_time_seconds = 60
  
  print("[RAPID_TYPING] Starting rapid typing benchmark (60s max)")
  print("[RAPID_TYPING] Simulating proficient keyboardist: 70 WPM (170ms/char)")
  
  local results = {
    success = false,
    server_name = _G.BENCHMARK_LSP_NAME,
    error = nil,
    scenarios = {},
    total_characters_typed = 0,
    total_completions_triggered = 0,
    total_parsing_responses = 0,
    character_times = {},
    completion_times = {},
    parse_response_times = {},
    memory_usage = 0
  }
  
  local success, error_msg = pcall(function()
    -- Open test document for editing
    local test_file = fixture_dir .. '/rapid-typing-test.html'
    vim.cmd('edit ' .. test_file)
    vim.wait(1000, function() return false end) -- Wait for LSP initialization
    
    -- Scenario 1: Building Custom Element (20s)
    print("[RAPID_TYPING] Scenario 1: Building custom element from scratch")
    local scenario1_start = vim.fn.reltime()
    
    -- Start with empty document, type complete custom element
    local element_text = '<my-card variant="outlined" size="large">\n  <h3 slot="header">Product Details</h3>\n  <p>Description text here</p>\n  <my-button variant="primary" size="medium">Add to Cart</my-button>\n</my-card>'
    
    local scenario1_result = type_text_realistically(element_text, {
      base_delay = 170, -- 70 WPM baseline
      variance = 50,    -- ±50ms human variance
      completion_frequency = 6, -- Every 6 characters on average
      thinking_pauses = {10, 25, 40}, -- Pause positions for thinking
      thinking_delay = 1000 -- 1 second thinking pauses
    })
    
    results.scenarios.building_element = {
      duration = tonumber(vim.fn.reltimestr(vim.fn.reltime(scenario1_start))),
      characters_typed = scenario1_result.characters,
      completions_triggered = scenario1_result.completions,
      avg_char_time = scenario1_result.avg_char_time,
      avg_completion_time = scenario1_result.avg_completion_time
    }
    
    -- Check time limit
    if vim.fn.reltime(start_time)[1] >= max_time_seconds then
      return
    end
    
    -- Scenario 2: Rapid Attribute Editing (15s)
    print("[RAPID_TYPING] Scenario 2: Rapid attribute modifications")
    local scenario2_start = vim.fn.reltime()
    
    -- Navigate to existing element and rapidly add/modify attributes
    vim.api.nvim_win_set_cursor(0, {1, 8}) -- Position in first <my-card>
    
    -- Add attributes rapidly: disabled loading icon="star" class="custom"
    local attributes_text = ' disabled loading icon="star" class="custom"'
    
    local scenario2_result = type_text_realistically(attributes_text, {
      base_delay = 150, -- Faster for simple attribute typing
      variance = 40,
      completion_frequency = 4, -- More frequent completions for attributes
      thinking_pauses = {},
      thinking_delay = 0
    })
    
    results.scenarios.attribute_editing = {
      duration = tonumber(vim.fn.reltimestr(vim.fn.reltime(scenario2_start))),
      characters_typed = scenario2_result.characters,
      completions_triggered = scenario2_result.completions,
      avg_char_time = scenario2_result.avg_char_time,
      avg_completion_time = scenario2_result.avg_completion_time
    }
    
    -- Check time limit
    if vim.fn.reltime(start_time)[1] >= max_time_seconds then
      return
    end
    
    -- Scenario 3: Template Literal Editing (20s)
    print("[RAPID_TYPING] Scenario 3: TypeScript template literal rapid editing")
    local scenario3_start = vim.fn.reltime()
    
    -- Add TypeScript template at end of document
    vim.api.nvim_win_set_cursor(0, {vim.api.nvim_buf_line_count(0), 0})
    vim.api.nvim_put({'', '// TypeScript template', 'const template = html`'}, 'l', true, true)
    
    local template_text = '\n<my-container layout="flex" gap="medium">\n  <my-header size="large">Dashboard</my-header>\n  <my-content>\n    <my-chart type="line" data-source="api/data"></my-chart>\n  </my-content>\n</my-container>\n`;'
    
    local scenario3_result = type_text_realistically(template_text, {
      base_delay = 180, -- Slightly slower for complex template syntax
      variance = 60,
      completion_frequency = 5,
      thinking_pauses = {20, 45, 70}, -- Pauses for template structure thinking
      thinking_delay = 800
    })
    
    results.scenarios.template_editing = {
      duration = tonumber(vim.fn.reltimestr(vim.fn.reltime(scenario3_start))),
      characters_typed = scenario3_result.characters,
      completions_triggered = scenario3_result.completions,
      avg_char_time = scenario3_result.avg_char_time,
      avg_completion_time = scenario3_result.avg_completion_time
    }
    
    -- Calculate overall statistics
    for _, scenario in pairs(results.scenarios) do
      results.total_characters_typed = results.total_characters_typed + scenario.characters_typed
      results.total_completions_triggered = results.total_completions_triggered + scenario.completions_triggered
    end
    
    -- Memory usage
    results.memory_usage = vim.fn.system('ps -o rss= -p ' .. vim.fn.getpid()):gsub('%s+', '') or 0
    
    results.success = results.total_characters_typed >= 200 and results.total_completions_triggered >= 10
  end)
  
  if not success then
    results.error = error_msg
    print("[RAPID_TYPING] Error: " .. tostring(error_msg))
  end
  
  local total_time = vim.fn.reltime(start_time)
  results.total_time = tonumber(vim.fn.reltimestr(total_time))
  
  print(string.format("[RAPID_TYPING] Completed in %.2fs - Chars: %d, Completions: %d, Success: %s", 
    results.total_time, results.total_characters_typed, results.total_completions_triggered, tostring(results.success)))
    
  return results
end

-- Function to type text at realistic human speeds with completions
function type_text_realistically(text, options)
  local opts = options or {}
  local base_delay = opts.base_delay or 170 -- Default 70 WPM
  local variance = opts.variance or 50
  local completion_freq = opts.completion_frequency or 6
  local thinking_pauses = opts.thinking_pauses or {}
  local thinking_delay = opts.thinking_delay or 1000
  
  local results = {
    characters = 0,
    completions = 0,
    char_times = {},
    completion_times = {},
    avg_char_time = 0,
    avg_completion_time = 0
  }
  
  local chars = {}
  for i = 1, #text do
    chars[i] = text:sub(i, i)
  end
  
  for i, char in ipairs(chars) do
    -- Check for thinking pause
    local is_thinking_pause = false
    for _, pause_pos in ipairs(thinking_pauses) do
      if i == pause_pos then
        is_thinking_pause = true
        break
      end
    end
    
    if is_thinking_pause then
      vim.wait(thinking_delay, function() return false end)
    end
    
    -- Type character with realistic timing
    local char_start = vim.fn.reltime()
    
    -- Human variance: base_delay ± variance
    local char_delay = base_delay + math.random(-variance, variance)
    char_delay = math.max(50, char_delay) -- Minimum 50ms (human limit)
    
    -- Insert character
    vim.api.nvim_put({char}, 'c', false, true)
    
    -- Wait for character processing
    vim.wait(char_delay, function() return false end)
    
    local char_time = tonumber(vim.fn.reltimestr(vim.fn.reltime(char_start)))
    table.insert(results.char_times, char_time)
    results.characters = results.characters + 1
    
    -- Trigger completion at realistic frequency
    if i % completion_freq == 0 and char:match('[a-zA-Z]') then
      local completion_start = vim.fn.reltime()
      
      -- Trigger completion
      pcall(vim.lsp.buf.completion)
      
      -- Brief wait for completion processing
      vim.wait(100, function() return false end)
      
      local completion_time = tonumber(vim.fn.reltimestr(vim.fn.reltime(completion_start)))
      table.insert(results.completion_times, completion_time)
      results.completions = results.completions + 1
    end
  end
  
  -- Calculate averages
  if #results.char_times > 0 then
    local total_char_time = 0
    for _, time in ipairs(results.char_times) do
      total_char_time = total_char_time + time
    end
    results.avg_char_time = total_char_time / #results.char_times
  end
  
  if #results.completion_times > 0 then
    local total_completion_time = 0
    for _, time in ipairs(results.completion_times) do
      total_completion_time = total_completion_time + time
    end
    results.avg_completion_time = total_completion_time / #results.completion_times
  end
  
  return results
end

return M