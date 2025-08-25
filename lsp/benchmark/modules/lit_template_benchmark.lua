-- Lit-html template benchmark module
-- Tests LSP functionality within TypeScript template literals

local measurement = require('utils.measurement')

local M = {}

function M.run_lit_template_benchmark(config, fixture_dir)
  local server_name = _G.BENCHMARK_LSP_NAME or "unknown"
  print(string.format("=== %s LSP Lit-html Template Benchmark ===", server_name))
  
  -- Change to fixture directory (use large project for realistic testing)
  vim.cmd('cd ' .. vim.fn.fnameescape(fixture_dir))
  
  -- Start LSP server
  local client_id = vim.lsp.start(config)
  if not client_id then
    return {
      success = false,
      error = "Failed to start LSP server"
    }
  end
  
  local client = vim.lsp.get_client_by_id(client_id)
  
  -- Wait for initialization
  local ready = vim.wait(10000, function()
    return client.server_capabilities ~= nil and not client.is_stopped()
  end)
  
  if not ready then
    client.stop()
    return {
      success = false,
      error = "Server failed to initialize"
    }
  end
  
  -- Use the existing lit-template.ts file for realistic testing
  local ts_file = 'lit-template.ts'
  
  -- Open file
  vim.cmd('edit ' .. ts_file)
  vim.cmd('set filetype=typescript')
  local bufnr = vim.api.nvim_get_current_buf()
  
  -- Wait for initial document processing
  vim.wait(3000)
  
  if client.is_stopped() then
    return {
      success = false,
      error = "Client stopped during initial processing"
    }
  end
  
  -- Lit-html template test scenarios
  local template_scenarios = {
    {
      name = "element_hover",
      description = "Hover information for custom elements in templates",
      tests = {
        {line = 76, character = 12, element = "my-container-flex", expected_content = "layout"},
        {line = 83, character = 15, element = "my-button-primary", expected_content = "button"},
        {line = 96, character = 12, element = "my-nav-tabs", expected_content = "navigation"},
        {line = 113, character = 15, element = "my-card-stats", expected_content = "card"},
        {line = 194, character = 15, element = "my-table-sortable", expected_content = "table"}
      }
    },
    {
      name = "attribute_completion",
      description = "Attribute completions within template elements",
      tests = {
        {line = 76, character = 45, element = "my-container-flex", attr_context = "gap", expected_min = 3},
        {line = 83, character = 35, element = "my-button-primary", attr_context = "variant", expected_min = 2},
        {line = 118, character = 35, element = "my-badge-primary", attr_context = "size", expected_min = 3},
        {line = 140, character = 30, element = "my-input-search", attr_context = "placeholder", expected_min = 1},
        {line = 204, character = 40, element = "my-table-sortable", attr_context = "striped", expected_min = 1}
      }
    },
    {
      name = "tag_completion",
      description = "Tag name completions in template contexts",
      tests = {
        {line = 100, character = 15, context = "nav", expected_min = 5, expected_contains = "my-nav"},
        {line = 150, character = 20, context = "button", expected_min = 8, expected_contains = "my-button"},
        {line = 200, character = 18, context = "input", expected_min = 10, expected_contains = "my-input"},
        {line = 250, character = 16, context = "card", expected_min = 7, expected_contains = "my-card"},
        {line = 300, character = 20, context = "image", expected_min = 5, expected_contains = "my-image"}
      }
    },
    {
      name = "template_interpolation",
      description = "LSP behavior in template interpolation contexts",
      tests = {
        {line = 78, character = 25, context = "${this.title}", expect_no_completions = true},
        {line = 118, character = 45, context = "${this.users.length}", expect_no_completions = true},
        {line = 219, character = 35, context = "${this.users.map(user => html`", expect_no_completions = false},
        {line = 232, character = 40, context = "${user.name}", expect_no_completions = true},
        {line = 244, character = 30, context = "${user.status === 'active'", expect_no_completions = true}
      }
    },
    {
      name = "event_binding",
      description = "Event binding attribute completions and validation",
      tests = {
        {line = 83, character = 50, element = "my-button-primary", event_context = "@click", expected_min = 1},
        {line = 87, character = 45, element = "my-button-secondary", event_context = "@click", expected_min = 1},
        {line = 137, character = 40, element = "my-input-search", event_context = "@input", expected_min = 1},
        {line = 154, character = 35, element = "my-button-outline", event_context = "@click", expected_min = 1},
        {line = 198, character = 30, element = "my-button-icon", event_context = "@click", expected_min = 1}
      }
    }
  }
  
  local scenario_results = {}
  local iterations_per_scenario = 5
  
  for _, scenario in ipairs(template_scenarios) do
    print(string.format("Testing %s scenario (%d iterations)...", scenario.name, iterations_per_scenario))
    
    local function single_scenario_test()
      local scenario_start = vim.uv.hrtime()
      local test_results = {}
      local successful_tests = 0
      local total_tests = #scenario.tests
      
      for i, test in ipairs(scenario.tests) do
        local test_start = vim.uv.hrtime()
        local test_success = false
        local test_data = {}
        
        -- Execute test based on scenario type
        if scenario.name == "element_hover" then
          -- Test hover functionality
          local hover_result = nil
          local hover_error = nil
          local hover_completed = false
          
          client.request('textDocument/hover', {
            textDocument = {uri = vim.uri_from_bufnr(bufnr)},
            position = {line = test.line - 1, character = test.character}
          }, function(err, result)
            hover_error = err
            hover_result = result
            hover_completed = true
          end)
          
          local hover_success = vim.wait(2000, function()
            return hover_completed
          end)
          
          if hover_success and not hover_error and hover_result then
            local content = ""
            if hover_result.contents then
              if type(hover_result.contents) == "string" then
                content = hover_result.contents
              elseif hover_result.contents.value then
                content = hover_result.contents.value
              end
            end
            
            test_success = test.expected_content and 
              string.find(content:lower(), test.expected_content:lower()) ~= nil
            test_data = {content_length = string.len(content), has_expected = test_success and 1 or 0}
          end
          
        elseif scenario.name == "attribute_completion" or scenario.name == "tag_completion" then
          -- Test completion functionality
          local comp_result = nil
          local comp_error = nil
          local comp_completed = false
          
          client.request('textDocument/completion', {
            textDocument = {uri = vim.uri_from_bufnr(bufnr)},
            position = {line = test.line - 1, character = test.character}
          }, function(err, result)
            comp_error = err
            comp_result = result
            comp_completed = true
          end)
          
          local comp_success = vim.wait(2000, function()
            return comp_completed
          end)
          
          if comp_success and not comp_error and comp_result then
            local items = comp_result.items or comp_result or {}
            local item_count = #items
            
            test_success = item_count >= (test.expected_min or 1)
            
            -- Check for specific content if specified
            if test.expected_contains then
              local contains_expected = false
              for _, item in ipairs(items) do
                if item.label and string.find(item.label, test.expected_contains) then
                  contains_expected = true
                  break
                end
              end
              test_success = test_success and contains_expected
            end
            
            test_data = {completion_count = item_count, meets_minimum = test_success and 1 or 0}
          end
          
        elseif scenario.name == "template_interpolation" then
          -- Test that completions behave correctly in interpolation contexts
          local comp_result = nil
          local comp_error = nil
          local comp_completed = false
          
          client.request('textDocument/completion', {
            textDocument = {uri = vim.uri_from_bufnr(bufnr)},
            position = {line = test.line - 1, character = test.character}
          }, function(err, result)
            comp_error = err
            comp_result = result
            comp_completed = true
          end)
          
          local comp_success = vim.wait(2000, function()
            return comp_completed
          end)
          
          if comp_success and not comp_error then
            local items = comp_result and (comp_result.items or comp_result) or {}
            local item_count = #items
            
            if test.expect_no_completions then
              test_success = item_count == 0
            else
              test_success = item_count > 0
            end
            
            test_data = {completion_count = item_count, expectation_met = test_success and 1 or 0}
          end
          
        elseif scenario.name == "event_binding" then
          -- Test event binding completions
          local comp_result = nil
          local comp_error = nil
          local comp_completed = false
          
          client.request('textDocument/completion', {
            textDocument = {uri = vim.uri_from_bufnr(bufnr)},
            position = {line = test.line - 1, character = test.character}
          }, function(err, result)
            comp_error = err
            comp_result = result
            comp_completed = true
          end)
          
          local comp_success = vim.wait(2000, function()
            return comp_completed
          end)
          
          if comp_success and not comp_error and comp_result then
            local items = comp_result.items or comp_result or {}
            local item_count = #items
            
            test_success = item_count >= (test.expected_min or 1)
            test_data = {completion_count = item_count, meets_minimum = test_success and 1 or 0}
          end
        end
        
        local test_duration = (vim.uv.hrtime() - test_start) / 1e6
        
        if test_success then
          successful_tests = successful_tests + 1
        end
        
        table.insert(test_results, {
          test_duration,
          test_success and 1 or 0,
          test_data
        })
        
        -- Check if client survived
        if client.is_stopped() then
          break
        end
        
        -- Small delay between tests
        vim.wait(100)
      end
      
      local scenario_duration = (vim.uv.hrtime() - scenario_start) / 1e6
      
      return {
        scenario_duration,
        total_tests,
        successful_tests,
        test_results
      }
    end
    
    -- Run multiple iterations for this scenario
    local scenario_iter_results = measurement.run_iterations(single_scenario_test, iterations_per_scenario, 
      string.format("lit template %s", scenario.name))
    
    -- Calculate scenario-specific metrics
    local total_tests = 0
    local successful_tests = 0
    local avg_test_duration = 0
    local test_count = 0
    
    for _, result in ipairs(scenario_iter_results.results or {}) do
      if result then
        total_tests = total_tests + (result[2] or 0)
        successful_tests = successful_tests + (result[3] or 0)
        
        -- Calculate average test duration from individual test results
        if result[4] then
          for _, test_result in ipairs(result[4]) do
            avg_test_duration = avg_test_duration + (test_result[1] or 0)
            test_count = test_count + 1
          end
        end
      end
    end
    
    table.insert(scenario_results, {
      scenario = scenario.name,
      description = scenario.description,
      iterations = iterations_per_scenario,
      successful_runs = scenario_iter_results.successful_runs,
      failed_runs = scenario_iter_results.failed_runs,
      success_rate = scenario_iter_results.success_rate,
      statistics = scenario_iter_results.statistics,
      total_tests = total_tests,
      successful_tests = successful_tests,
      test_success_rate = total_tests > 0 and successful_tests / total_tests or 0,
      avg_test_duration = test_count > 0 and avg_test_duration / test_count or 0,
      errors = scenario_iter_results.errors
    })
    
    -- Check if client survived
    if client.is_stopped() then
      break
    end
    
    -- Small delay between scenarios
    vim.wait(200)
  end
  
  -- Clean up
  vim.api.nvim_buf_delete(bufnr, {force = true})
  client.stop()
  
  -- Calculate overall statistics
  local total_scenarios = #template_scenarios
  local total_iterations = 0
  local total_successful = 0
  local all_scenario_times = {}
  
  for _, result in ipairs(scenario_results) do
    total_iterations = total_iterations + result.iterations
    total_successful = total_successful + result.successful_runs
    for _, time in ipairs(result.statistics and {result.statistics.mean} or {}) do
      table.insert(all_scenario_times, time)
    end
  end
  
  local overall_stats = measurement.calculate_statistics(all_scenario_times)
  
  return {
    success = total_successful > 0,
    server_name = server_name,
    total_scenarios = total_scenarios,
    total_iterations = total_iterations,
    total_successful = total_successful,
    success_rate = total_iterations > 0 and total_successful / total_iterations or 0,
    overall_statistics = overall_stats,
    scenario_results = scenario_results,
    client_survived = not client.is_stopped()
  }
end

return M