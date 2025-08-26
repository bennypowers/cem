local measurement = require('utils.measurement')

local M = {}

-- Integrated workflow benchmark: Tests realistic developer workflows combining typing + references
-- Time limit: 75 seconds total  
-- Focus: Real-world development patterns with multiple concurrent LSP operations

function M.run_integrated_workflow_benchmark(config, fixture_dir)
  local start_time = vim.fn.reltime()
  local max_time_seconds = 75
  
  print("[INTEGRATED] Starting integrated workflow benchmark (75s max)")
  print("[INTEGRATED] Testing: Rapid typing + references + completion + navigation")
  
  local results = {
    success = false,
    server_name = _G.BENCHMARK_LSP_NAME,
    error = nil,
    workflows = {},
    total_operations = 0,
    typing_performance = {},
    reference_performance = {},
    multi_operation_performance = {},
    memory_usage = 0
  }
  
  local success, error_msg = pcall(function()
    -- Workflow 1: Code Exploration Pattern (25s)
    print("[INTEGRATED] Workflow 1: Code exploration with typing interruptions")
    local exploration_start = vim.fn.reltime()
    
    local exploration_result = run_code_exploration_workflow(fixture_dir)
    
    results.workflows.code_exploration = {
      duration = tonumber(vim.fn.reltimestr(vim.fn.reltime(exploration_start))),
      operations = exploration_result.operations,
      references_triggered = exploration_result.references_triggered,
      typing_responsiveness = exploration_result.typing_responsiveness,
      navigation_speed = exploration_result.navigation_speed
    }
    
    -- Check time limit
    if vim.fn.reltime(start_time)[1] >= max_time_seconds then
      return
    end
    
    -- Workflow 2: Refactoring with References (25s)
    print("[INTEGRATED] Workflow 2: Refactoring with continuous reference checking")
    local refactoring_start = vim.fn.reltime()
    
    local refactoring_result = run_refactoring_workflow(fixture_dir)
    
    results.workflows.refactoring = {
      duration = tonumber(vim.fn.reltimestr(vim.fn.reltime(refactoring_start))),
      operations = refactoring_result.operations,
      references_accuracy = refactoring_result.references_accuracy,
      edit_consistency = refactoring_result.edit_consistency,
      background_performance = refactoring_result.background_performance
    }
    
    -- Check time limit
    if vim.fn.reltime(start_time)[1] >= max_time_seconds then
      return
    end
    
    -- Workflow 3: Rapid Development Pattern (25s)
    print("[INTEGRATED] Workflow 3: Rapid development with mixed operations")
    local development_start = vim.fn.reltime()
    
    local development_result = run_rapid_development_workflow(fixture_dir)
    
    results.workflows.rapid_development = {
      duration = tonumber(vim.fn.reltimestr(vim.fn.reltime(development_start))),
      operations = development_result.operations,
      completion_during_typing = development_result.completion_during_typing,
      reference_during_typing = development_result.reference_during_typing,
      overall_responsiveness = development_result.overall_responsiveness
    }
    
    -- Calculate overall statistics
    local total_ops = 0
    local total_typing_ops = 0
    local total_reference_ops = 0
    
    for _, workflow in pairs(results.workflows) do
      total_ops = total_ops + workflow.operations
      if workflow.typing_responsiveness then
        total_typing_ops = total_typing_ops + (workflow.typing_responsiveness.operations or 0)
      end
      if workflow.references_triggered then
        total_reference_ops = total_reference_ops + workflow.references_triggered
      end
    end
    
    results.total_operations = total_ops
    results.typing_performance.total_operations = total_typing_ops
    results.reference_performance.total_operations = total_reference_ops
    
    -- Memory usage
    results.memory_usage = vim.fn.system('ps -o rss= -p ' .. vim.fn.getpid()):gsub('%s+', '') or 0
    
    results.success = total_ops >= 20 and total_typing_ops >= 100 and total_reference_ops >= 5
  end)
  
  if not success then
    results.error = error_msg
    print("[INTEGRATED] Error: " .. tostring(error_msg))
  end
  
  local total_time = vim.fn.reltime(start_time)
  results.total_time = tonumber(vim.fn.reltimestr(total_time))
  
  print(string.format("[INTEGRATED] Completed in %.2fs - Ops: %d, Typing: %d, Refs: %d, Success: %s", 
    results.total_time, results.total_operations, results.typing_performance.total_operations or 0, 
    results.reference_performance.total_operations or 0, tostring(results.success)))
    
  return results
end

-- Workflow 1: Code Exploration Pattern
function run_code_exploration_workflow(fixture_dir)
  local results = {
    operations = 0,
    references_triggered = 0,
    typing_responsiveness = {},
    navigation_speed = {}
  }
  
  -- Open main file
  vim.cmd('edit ' .. fixture_dir .. '/large-page.html')
  vim.wait(1000, function() return false end)
  
  -- Pattern: Explore → Find element → Get references → Type notes → Repeat
  local exploration_elements = {
    {element = "my-button", line = 21, col = 8, note = "<!-- Primary button usage -->"},
    {element = "my-card", line = 24, col = 8, note = "<!-- Card container pattern -->"},
    {element = "my-nav", line = 21, col = 8, note = "<!-- Navigation component -->"}
  }
  
  for _, element in ipairs(exploration_elements) do
    -- Navigate to element
    local nav_start = vim.fn.reltime()
    vim.api.nvim_win_set_cursor(0, {element.line, element.col})
    local nav_time = tonumber(vim.fn.reltimestr(vim.fn.reltime(nav_start)))
    table.insert(results.navigation_speed, nav_time)
    
    -- Trigger references
    local ref_start = vim.fn.reltime()
    pcall(vim.lsp.buf.references)
    vim.wait(800, function() return false end) -- Wait for references
    results.references_triggered = results.references_triggered + 1
    
    -- Type notes while references are processing
    local typing_start = vim.fn.reltime()
    vim.api.nvim_win_set_cursor(0, {element.line + 1, 0})
    
    -- Type note with realistic character timing
    type_text_with_timing(element.note, 160) -- 160ms per character
    
    local typing_time = tonumber(vim.fn.reltimestr(vim.fn.reltime(typing_start)))
    table.insert(results.typing_responsiveness, typing_time)
    
    results.operations = results.operations + 3 -- Nav + Ref + Type
  end
  
  -- Calculate typing responsiveness metrics
  if #results.typing_responsiveness > 0 then
    local total_typing_time = 0
    for _, time in ipairs(results.typing_responsiveness) do
      total_typing_time = total_typing_time + time
    end
    results.typing_responsiveness = {
      operations = #results.typing_responsiveness,
      avg_time = total_typing_time / #results.typing_responsiveness,
      times = results.typing_responsiveness
    }
  end
  
  return results
end

-- Workflow 2: Refactoring with References
function run_refactoring_workflow(fixture_dir)
  local results = {
    operations = 0,
    references_accuracy = {},
    edit_consistency = {},
    background_performance = {}
  }
  
  -- Open file for refactoring
  vim.cmd('edit ' .. fixture_dir .. '/medium-page.html')
  vim.wait(500, function() return false end)
  
  -- Pattern: Find references → Make edit → Check references again → Verify consistency
  local refactoring_targets = {
    {element = "my-button", old_attr = 'size="medium"', new_attr = 'size="large"', line = 15},
    {element = "my-card", old_attr = 'variant="outlined"', new_attr = 'variant="filled"', line = 10}
  }
  
  for _, target in ipairs(refactoring_targets) do
    -- Get initial references
    local initial_ref_start = vim.fn.reltime()
    vim.api.nvim_win_set_cursor(0, {target.line, 5})
    pcall(vim.lsp.buf.references)
    vim.wait(1000, function() return false end)
    local initial_ref_time = tonumber(vim.fn.reltimestr(vim.fn.reltime(initial_ref_start)))
    
    -- Make refactoring edit
    local edit_start = vim.fn.reltime()
    pcall(vim.cmd, '%s/' .. target.old_attr .. '/' .. target.new_attr .. '/g')
    local edit_time = tonumber(vim.fn.reltimestr(vim.fn.reltime(edit_start)))
    
    -- Check references again after edit
    local post_ref_start = vim.fn.reltime()
    pcall(vim.lsp.buf.references)
    vim.wait(1000, function() return false end)
    local post_ref_time = tonumber(vim.fn.reltimestr(vim.fn.reltime(post_ref_start)))
    
    -- Record performance data
    table.insert(results.references_accuracy, {
      element = target.element,
      initial_ref_time = initial_ref_time,
      post_ref_time = post_ref_time,
      edit_time = edit_time
    })
    
    table.insert(results.background_performance, {
      operation = "refactor_" .. target.element,
      total_time = initial_ref_time + edit_time + post_ref_time
    })
    
    results.operations = results.operations + 3 -- Initial ref + Edit + Post ref
  end
  
  return results
end

-- Workflow 3: Rapid Development Pattern
function run_rapid_development_workflow(fixture_dir)
  local results = {
    operations = 0,
    completion_during_typing = {},
    reference_during_typing = {},
    overall_responsiveness = {}
  }
  
  -- Create new document for rapid development
  vim.cmd('edit ' .. fixture_dir .. '/rapid-development-test.html')
  vim.wait(500, function() return false end)
  
  -- Insert basic structure
  vim.api.nvim_put({'<!DOCTYPE html>', '<html>', '<body>', '  <!-- Development area -->', '</body>', '</html>'}, 'l', false, true)
  
  -- Pattern: Type rapidly → Trigger completion → Get references → Continue typing
  local development_scenarios = {
    {
      typing_text = '<my-card variant="outlined"',
      completion_trigger_pos = 15, -- After "my-card"
      reference_element = "my-card",
      continuation_text = ' size="large">Content</my-card>'
    },
    {
      typing_text = '<my-button-group><my-button variant="primary"',
      completion_trigger_pos = 25, -- After "my-button"
      reference_element = "my-button",
      continuation_text = '>Click</my-button></my-button-group>'
    }
  }
  
  for i, scenario in ipairs(development_scenarios) do
    vim.api.nvim_win_set_cursor(0, {3 + i, 2}) -- Position in development area
    
    -- Start typing
    local typing_start = vim.fn.reltime()
    
    -- Type initial text with character timing
    local chars_before_completion = scenario.typing_text:sub(1, scenario.completion_trigger_pos)
    type_text_with_timing(chars_before_completion, 150) -- Fast typing
    
    -- Trigger completion mid-typing
    local completion_start = vim.fn.reltime()
    pcall(vim.lsp.buf.completion)
    vim.wait(200, function() return false end)
    local completion_time = tonumber(vim.fn.reltimestr(vim.fn.reltime(completion_start)))
    
    -- Continue typing
    local remaining_text = scenario.typing_text:sub(scenario.completion_trigger_pos + 1)
    type_text_with_timing(remaining_text, 150)
    
    -- Trigger references while typing is fresh
    local ref_start = vim.fn.reltime()
    pcall(vim.lsp.buf.references)
    local ref_time = tonumber(vim.fn.reltimestr(vim.fn.reltime(ref_start)))
    
    -- Finish typing
    type_text_with_timing(scenario.continuation_text, 150)
    
    local total_typing_time = tonumber(vim.fn.reltimestr(vim.fn.reltime(typing_start)))
    
    -- Record performance
    table.insert(results.completion_during_typing, {
      scenario = i,
      completion_time = completion_time,
      total_typing_time = total_typing_time,
      chars_typed = #scenario.typing_text + #scenario.continuation_text
    })
    
    table.insert(results.reference_during_typing, {
      scenario = i,
      reference_time = ref_time,
      element = scenario.reference_element
    })
    
    results.operations = results.operations + 4 -- Type + Complete + Reference + Continue
  end
  
  -- Calculate overall responsiveness
  local total_completion_time = 0
  local total_reference_time = 0
  
  for _, comp in ipairs(results.completion_during_typing) do
    total_completion_time = total_completion_time + comp.completion_time
  end
  
  for _, ref in ipairs(results.reference_during_typing) do
    total_reference_time = total_reference_time + ref.reference_time
  end
  
  results.overall_responsiveness = {
    avg_completion_time = #results.completion_during_typing > 0 and (total_completion_time / #results.completion_during_typing) or 0,
    avg_reference_time = #results.reference_during_typing > 0 and (total_reference_time / #results.reference_during_typing) or 0,
    total_operations = results.operations
  }
  
  return results
end

-- Helper function to type text with specific character timing
function type_text_with_timing(text, delay_ms)
  local chars = {}
  for i = 1, #text do
    chars[i] = text:sub(i, i)
  end
  
  for _, char in ipairs(chars) do
    vim.api.nvim_put({char}, 'c', false, true)
    vim.wait(delay_ms, function() return false end)
  end
end

return M