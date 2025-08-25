-- File lifecycle benchmark module
-- Tests file open/change/close operations and LSP responsiveness

local measurement = require('utils.measurement')

local M = {}

function M.run_file_lifecycle_benchmark(config, fixture_dir)
  local server_name = _G.BENCHMARK_LSP_NAME or "unknown"
  print(string.format("=== %s LSP File Lifecycle Benchmark ===", server_name))
  
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
  
  local lifecycle_results = {}
  local iterations = 5 -- Multiple iterations for lifecycle operations
  
  -- Test 1: File open performance with statistical analysis
  print(string.format("Testing file open operations (%d iterations)...", iterations))
  
  local html_content = [[<!DOCTYPE html>
<html>
<head><title>Lifecycle Test</title></head>
<body>
  <my-button type="primary">Test</my-button>
</body>
</html>]]
  
  local function single_file_open()
    local filename = string.format('test-lifecycle-%d.html', math.random(10000))
    
    local open_start = vim.uv.hrtime()
    
    vim.fn.writefile(vim.split(html_content, '\n'), filename)
    vim.cmd('edit ' .. filename)
    vim.cmd('set filetype=html')
    local bufnr = vim.api.nvim_get_current_buf()
    
    vim.wait(500) -- Allow document processing
    
    local open_duration = (vim.uv.hrtime() - open_start) / 1e6
    
    -- Clean up
    vim.api.nvim_buf_delete(bufnr, {force = true})
    vim.fn.delete(filename)
    
    if client.is_stopped() then
      return nil, "Client stopped during file open"
    end
    
    return open_duration
  end
  
  local open_results = measurement.run_iterations(single_file_open, iterations, "file open")
  table.insert(lifecycle_results, {
    operation = "file_open",
    iterations = iterations,
    successful_runs = open_results.successful_runs,
    failed_runs = open_results.failed_runs,
    success_rate = open_results.success_rate,
    statistics = open_results.statistics,
    all_times_ms = open_results.times_ms,
    errors = open_results.errors
  })
  
  if client.is_stopped() then
    return {
      success = false,
      error = "Client stopped during file open",
      lifecycle_results = lifecycle_results
    }
  end
  
  -- Test 2: Simple file operations test
  print("Testing basic file operations...")
  local basic_operations_start = vim.uv.hrtime()
  
  vim.fn.writefile(vim.split(html_content, '\n'), 'test-lifecycle-main.html')
  vim.cmd('edit test-lifecycle-main.html')
  vim.cmd('set filetype=html')
  local bufnr = vim.api.nvim_get_current_buf()
  
  vim.wait(1000) -- Allow processing
  
  -- Add new content
  local new_lines = {
    '  <my-card title="Added dynamically">',
    '    <div slot="content">New content</div>',
    '  </my-card>'
  }
  
  vim.api.nvim_buf_set_lines(bufnr, 4, 4, false, new_lines)
  vim.wait(1000) -- Allow change processing
  
  local change_duration = (vim.uv.hrtime() - basic_operations_start) / 1e6
  table.insert(lifecycle_results, {
    operation = "file_change",
    duration_ms = change_duration,
    success = not client.is_stopped()
  })
  
  if client.is_stopped() then
    vim.api.nvim_buf_delete(bufnr, {force = true})
    vim.fn.delete('test-lifecycle-main.html')
    return {
      success = false,
      error = "Client stopped during file change",
      lifecycle_results = lifecycle_results
    }
  end
  
  -- Test 3: Multiple rapid changes
  local rapid_start = vim.uv.hrtime()
  
  for i = 1, 5 do
    vim.api.nvim_buf_set_lines(bufnr, -1, -1, false, {
      string.format('  <my-element-%d></my-element-%d>', i, i)
    })
    vim.wait(100) -- Small delay between changes
  end
  
  vim.wait(1000) -- Allow final processing
  
  local rapid_duration = (vim.uv.hrtime() - rapid_start) / 1e6
  table.insert(lifecycle_results, {
    operation = "rapid_changes",
    duration_ms = rapid_duration,
    change_count = 5,
    success = not client.is_stopped()
  })
  
  -- Test 4: Hover after changes (functionality test)
  local hover_result = nil
  local hover_error = nil
  local hover_completed = false
  
  local hover_start = vim.uv.hrtime()
  
  client.request('textDocument/hover', {
    textDocument = {uri = vim.uri_from_bufnr(bufnr)},
    position = {line = 4, character = 10} -- On my-button
  }, function(err, result)
    hover_error = err
    hover_result = result
    hover_completed = true
  end)
  
  local hover_success = vim.wait(3000, function()
    return hover_completed
  end)
  
  local hover_duration = (vim.uv.hrtime() - hover_start) / 1e6
  table.insert(lifecycle_results, {
    operation = "hover_after_changes",
    duration_ms = hover_duration,
    success = hover_success and not hover_error and hover_result ~= nil
  })
  
  -- Test 5: File close performance
  local close_start = vim.uv.hrtime()
  vim.api.nvim_buf_delete(bufnr, {force = true})
  vim.wait(500) -- Allow cleanup
  local close_duration = (vim.uv.hrtime() - close_start) / 1e6
  
  table.insert(lifecycle_results, {
    operation = "file_close",
    duration_ms = close_duration,
    success = true
  })
  
  -- Clean up
  vim.fn.delete('test-lifecycle-main.html')
  client.stop()
  
  -- Calculate summary stats
  local total_duration = 0
  local successful_operations = 0
  
  for _, result in ipairs(lifecycle_results) do
    local duration = result.duration_ms or (result.statistics and result.statistics.mean) or 0
    total_duration = total_duration + duration
    if result.success or (result.success_rate and result.success_rate > 0) then
      successful_operations = successful_operations + 1
    end
  end
  
  return {
    success = successful_operations == #lifecycle_results,
    server_name = server_name,
    total_operations = #lifecycle_results,
    successful_operations = successful_operations,
    success_rate = successful_operations / #lifecycle_results,
    total_duration_ms = total_duration,
    average_duration_ms = total_duration / #lifecycle_results,
    lifecycle_results = lifecycle_results,
    client_survived = not client.is_stopped()
  }
end

return M