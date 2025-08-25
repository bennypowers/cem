-- Hover benchmark module
-- Tests LSP hover request performance and functionality

local measurement = require('utils.measurement')

local M = {}

function M.run_hover_benchmark(config, fixture_dir)
  local server_name = _G.BENCHMARK_LSP_NAME or "unknown"
  print(string.format("=== %s LSP Hover Benchmark ===", server_name))
  
  -- Change to fixture directory
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
  
  -- Create test HTML file with custom elements
  local html_content = [[<!DOCTYPE html>
<html>
<head><title>Hover Test</title></head>
<body>
  <my-button type="primary">Click me</my-button>
  <my-card title="Test Card">
    <div slot="content">Card content</div>
    <div slot="footer">Footer content</div>
  </my-card>
  <my-icon name="star" size="large"></my-icon>
</body>
</html>]]
  
  vim.fn.writefile(vim.split(html_content, '\n'), 'test-hover.html')
  
  -- Open file
  vim.cmd('edit test-hover.html')
  vim.cmd('set filetype=html')
  local bufnr = vim.api.nvim_get_current_buf()
  
  -- Wait for document processing
  vim.wait(2000)
  
  if client.is_stopped() then
    vim.fn.delete('test-hover.html')
    return {
      success = false,
      error = "Client stopped during file processing"
    }
  end
  
  -- Test hover on various elements with statistical analysis
  local test_positions = {
    {element = "my-button", line = 4, character = 10},
    {element = "my-card", line = 5, character = 10}, 
    {element = "my-icon", line = 9, character = 10}
  }
  
  local hover_results = {}
  local iterations_per_element = 10 -- Multiple iterations per element for statistical confidence
  
  for i, test_pos in ipairs(test_positions) do
    print(string.format("Testing hover on %s (%d iterations)...", test_pos.element, iterations_per_element))
    
    local function single_hover()
      local hover_result = nil
      local hover_error = nil
      local hover_completed = false
      
      local hover_start = vim.uv.hrtime()
      
      client.request('textDocument/hover', {
        textDocument = {uri = vim.uri_from_bufnr(bufnr)},
        position = {line = test_pos.line, character = test_pos.character}
      }, function(err, result)
        hover_error = err
        hover_result = result
        hover_completed = true
      end)
      
      local hover_success = vim.wait(3000, function()
        return hover_completed
      end)
      
      local hover_duration = (vim.uv.hrtime() - hover_start) / 1e6
      
      if not hover_success or hover_error or not hover_result then
        return nil, hover_error and vim.inspect(hover_error) or "No result"
      end
      
      local content_length = 0
      if hover_result.contents then
        if type(hover_result.contents) == "string" then
          content_length = string.len(hover_result.contents)
        elseif hover_result.contents.value then
          content_length = string.len(hover_result.contents.value)
        end
      end
      
      return {hover_duration, content_length}
    end
    
    -- Run multiple iterations for this element
    local element_results = measurement.run_iterations(single_hover, iterations_per_element, 
      string.format("hover on %s", test_pos.element))
    
    table.insert(hover_results, {
      element = test_pos.element,
      iterations = iterations_per_element,
      successful_runs = element_results.successful_runs,
      failed_runs = element_results.failed_runs,
      success_rate = element_results.success_rate,
      statistics = element_results.statistics,
      all_times_ms = element_results.times_ms,
      errors = element_results.errors,
      -- Average content length from successful runs
      avg_content_length = element_results.results and #element_results.results > 0 and 
        (function()
          local total = 0
          for _, res in ipairs(element_results.results) do
            total = total + (res[2] or 0)
          end
          return total / #element_results.results
        end)() or 0
    })
    
    -- Check if client survived
    if client.is_stopped() then
      break
    end
    
    -- Small delay between different elements
    vim.wait(100)
  end
  
  -- Clean up
  vim.api.nvim_buf_delete(bufnr, {force = true})
  vim.fn.delete('test-hover.html')
  client.stop()
  
  -- Calculate overall statistics
  local total_successful = 0
  local total_attempts = 0
  local all_times = {}
  
  for _, result in ipairs(hover_results) do
    total_successful = total_successful + result.successful_runs
    total_attempts = total_attempts + result.iterations
    for _, time in ipairs(result.all_times_ms) do
      table.insert(all_times, time)
    end
  end
  
  local overall_stats = measurement.calculate_statistics(all_times)
  
  return {
    success = total_successful > 0,
    server_name = server_name,
    total_elements = #test_positions,
    total_attempts = total_attempts,
    total_successful = total_successful,
    success_rate = total_attempts > 0 and total_successful / total_attempts or 0,
    overall_statistics = overall_stats,
    hover_results = hover_results,
    client_survived = not client.is_stopped(),
    -- Backward compatibility
    successful_hovers = total_successful,
    average_duration_ms = overall_stats.mean
  }
end

return M