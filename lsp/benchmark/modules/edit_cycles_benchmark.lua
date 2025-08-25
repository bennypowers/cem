-- Edit cycles benchmark module
-- Tests LSP response times for diagnostics/completions during rapid HTML editing

local measurement = require('utils.measurement')

local M = {}

function M.run_edit_cycles_benchmark(config, fixture_dir)
  local server_name = _G.BENCHMARK_LSP_NAME or "unknown"
  print(string.format("=== %s LSP Edit Cycles Benchmark ===", server_name))
  
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
  
  -- Use the existing large HTML page for realistic editing
  local html_file = 'large-page.html'
  
  -- Open file
  vim.cmd('edit ' .. html_file)
  vim.cmd('set filetype=html')
  local bufnr = vim.api.nvim_get_current_buf()
  
  -- Wait for initial document processing
  vim.wait(3000)
  
  if client.is_stopped() then
    return {
      success = false,
      error = "Client stopped during initial processing"
    }
  end
  
  -- Edit cycles test scenarios - use safe positions that exist in large-page.html
  local edit_scenarios = {
    {
      name = "attribute_addition",
      description = "Adding attributes to existing elements",
      edits = {
        {line = 22, col = 50, insert = ' data-test="1"'},
        {line = 72, col = 30, insert = ' variant="primary"'},
        {line = 89, col = 45, insert = ' size="large"'},
        {line = 111, col = 50, insert = ' disabled'},
        {line = 150, col = 40, insert = ' required'}
      }
    },
    {
      name = "element_insertion",
      description = "Inserting new elements",
      edits = {
        {line = 50, col = 0, insert = '  <my-button-new>New Button</my-button-new>\n'},
        {line = 100, col = 0, insert = '  <my-input-search placeholder="Search here"></my-input-search>\n'},
        {line = 150, col = 0, insert = '  <my-card-highlight selected>Highlight</my-card-highlight>\n'},
        {line = 200, col = 0, insert = '  <my-alert-notification icon="bell">New message</my-alert-notification>\n'},
        {line = 250, col = 0, insert = '  <my-image-placeholder src="/test.jpg" alt="Test"></my-image-placeholder>\n'}
      }
    },
    {
      name = "content_modification",
      description = "Modifying element content and structure",
      edits = {
        {line = 60, col = 10, replace = "Updated Header", length = 20},
        {line = 110, col = 15, replace = "Modified content with more text", length = 12},
        {line = 160, col = 8, replace = "Enhanced description", length = 15},
        {line = 210, col = 12, replace = "Improved functionality", length = 18},
        {line = 260, col = 5, replace = "Better user experience", length = 10}
      }
    }
  }
  
  local scenario_results = {}
  local iterations_per_scenario = 3
  
  for _, scenario in ipairs(edit_scenarios) do
    print(string.format("Testing %s scenario (%d iterations)...", scenario.name, iterations_per_scenario))
    
    local function single_edit_cycle()
      local cycle_start = vim.uv.hrtime()
      local diagnostics_times = {}
      local completion_times = {}
      local total_edits = 0
      
      -- Perform rapid edits
      for i, edit in ipairs(scenario.edits) do
        local edit_start = vim.uv.hrtime()
        
        -- Apply the edit
        if edit.insert then
          vim.api.nvim_buf_set_text(bufnr, edit.line - 1, edit.col, edit.line - 1, edit.col, 
            vim.split(edit.insert, '\n'))
        elseif edit.replace then
          vim.api.nvim_buf_set_text(bufnr, edit.line - 1, edit.col, edit.line - 1, 
            edit.col + edit.length, {edit.replace})
        end
        
        total_edits = total_edits + 1
        
        -- Small delay to simulate realistic typing
        vim.wait(50)
        
        -- Check for diagnostics support and request if available
        local diag_start = vim.uv.hrtime()
        local diag_received = false
        local diag_error = nil
        
        -- Check if server supports diagnostics (LSP 3.17+ textDocument/diagnostic)
        local server_capabilities = client.server_capabilities
        if server_capabilities and server_capabilities.diagnosticProvider then
          client.request('textDocument/diagnostic', {
            textDocument = {uri = vim.uri_from_bufnr(bufnr)},
            identifier = nil,
            previousResultId = nil
          }, function(err, result)
            diag_error = err
            diag_received = true
          end)
          
          -- Wait for diagnostics (shorter timeout for rapid cycles)
          local diag_success = vim.wait(1000, function()
            return diag_received
          end)
          
          if diag_success and not diag_error then
            table.insert(diagnostics_times, (vim.uv.hrtime() - diag_start) / 1e6)
          end
        else
          -- Server doesn't support diagnostics - this is fine, skip timing
          diag_received = true
        end
        
        -- Request completion at edit position
        local comp_start = vim.uv.hrtime()
        local comp_received = false
        local comp_error = nil
        local comp_result = nil
        
        client.request('textDocument/completion', {
          textDocument = {uri = vim.uri_from_bufnr(bufnr)},
          position = {line = edit.line - 1, character = edit.col + (edit.insert and #edit.insert or 0)}
        }, function(err, result)
          comp_error = err
          comp_result = result
          comp_received = true
        end)
        
        -- Wait for completion
        local comp_success = vim.wait(1000, function()
          return comp_received
        end)
        
        if comp_success and not comp_error and comp_result then
          table.insert(completion_times, (vim.uv.hrtime() - comp_start) / 1e6)
        end
        
        -- Check if server survived
        if client.is_stopped() then
          return nil, "Server stopped during edit cycle"
        end
      end
      
      local cycle_duration = (vim.uv.hrtime() - cycle_start) / 1e6
      
      -- Calculate statistics
      local avg_diag_time = #diagnostics_times > 0 and 
        (table.unpack and table.unpack(diagnostics_times) and 
         vim.tbl_reduce(function(sum, time) return sum + time end, diagnostics_times, 0) / #diagnostics_times) or 0
      
      local avg_comp_time = #completion_times > 0 and 
        (table.unpack and table.unpack(completion_times) and
         vim.tbl_reduce(function(sum, time) return sum + time end, completion_times, 0) / #completion_times) or 0
      
      return {
        cycle_duration,
        total_edits,
        #diagnostics_times,
        #completion_times,
        avg_diag_time,
        avg_comp_time
      }
    end
    
    -- Run multiple iterations for this scenario
    local scenario_iter_results = measurement.run_iterations(single_edit_cycle, iterations_per_scenario, 
      string.format("edit cycle %s", scenario.name))
    
    -- Calculate scenario-specific metrics
    local total_edits = 0
    local successful_diagnostics = 0
    local successful_completions = 0
    local total_diag_time = 0
    local total_comp_time = 0
    
    for _, result in ipairs(scenario_iter_results.results or {}) do
      if result then
        total_edits = total_edits + (result[2] or 0)
        successful_diagnostics = successful_diagnostics + (result[3] or 0)
        successful_completions = successful_completions + (result[4] or 0)
        total_diag_time = total_diag_time + (result[5] or 0)
        total_comp_time = total_comp_time + (result[6] or 0)
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
      total_edits = total_edits,
      diagnostic_success_rate = total_edits > 0 and successful_diagnostics / total_edits or 0,
      completion_success_rate = total_edits > 0 and successful_completions / total_edits or 0,
      avg_diagnostic_time = successful_diagnostics > 0 and total_diag_time / successful_diagnostics or 0,
      avg_completion_time = successful_completions > 0 and total_comp_time / successful_completions or 0,
      errors = scenario_iter_results.errors
    })
    
    -- Reset file state between scenarios
    vim.cmd('edit!')
    vim.wait(500)
    
    -- Check if client survived
    if client.is_stopped() then
      break
    end
  end
  
  -- Clean up
  vim.api.nvim_buf_delete(bufnr, {force = true})
  client.stop()
  
  -- Calculate overall statistics
  local total_cycles = 0
  local total_successful = 0
  local all_cycle_times = {}
  
  for _, result in ipairs(scenario_results) do
    total_cycles = total_cycles + result.iterations
    total_successful = total_successful + result.successful_runs
    for _, time in ipairs(result.statistics and {result.statistics.mean} or {}) do
      table.insert(all_cycle_times, time)
    end
  end
  
  local overall_stats = measurement.calculate_statistics(all_cycle_times)
  
  return {
    success = total_successful > 0,
    server_name = server_name,
    total_scenarios = #edit_scenarios,
    total_cycles = total_cycles,
    total_successful = total_successful,
    success_rate = total_cycles > 0 and total_successful / total_cycles or 0,
    overall_statistics = overall_stats,
    scenario_results = scenario_results,
    client_survived = not client.is_stopped()
  }
end

return M