-- Stress test benchmark module
-- Tests LSP server performance under heavy load with large documents and many requests

local measurement = require('utils.measurement')

local M = {}

function M.run_stress_test_benchmark(config, fixture_dir)
  local server_name = _G.BENCHMARK_LSP_NAME or "unknown"
  print(string.format("=== %s LSP Stress Test Benchmark ===", server_name))
  
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
  
  -- Create a large HTML document with many custom elements
  local large_html_content = {'<!DOCTYPE html>', '<html>', '<head><title>Stress Test</title></head>', '<body>'}
  
  -- Generate 100 different custom element instances
  for i = 1, 100 do
    table.insert(large_html_content, string.format('  <my-element-%d class="test-%d">', i, i))
    table.insert(large_html_content, string.format('    <my-button type="primary" size="large">Button %d</my-button>', i))
    table.insert(large_html_content, string.format('    <my-card title="Card %d" variant="elevated">', i))
    table.insert(large_html_content, '      <div slot="content">Content for card</div>')
    table.insert(large_html_content, '      <div slot="footer">Footer content</div>')
    table.insert(large_html_content, '    </my-card>')
    table.insert(large_html_content, string.format('    <my-icon name="star" size="medium" color="primary-%d"></my-icon>', i % 5))
    table.insert(large_html_content, string.format('  </my-element-%d>', i))
  end
  
  table.insert(large_html_content, '</body>')
  table.insert(large_html_content, '</html>')
  
  vim.fn.writefile(large_html_content, 'test-stress.html')
  
  -- Open the large file
  vim.cmd('edit test-stress.html')
  vim.cmd('set filetype=html')
  local bufnr = vim.api.nvim_get_current_buf()
  
  print(string.format("Created large document with %d lines", #large_html_content))
  
  -- Wait for document processing
  vim.wait(3000)
  
  if client.is_stopped() then
    vim.fn.delete('test-stress.html')
    return {
      success = false,
      error = "Client stopped during large document processing"
    }
  end
  
  local stress_results = {}
  
  -- Test 1: Bulk hover requests
  print("Running bulk hover stress test (50 requests)...")
  local hover_positions = {}
  for i = 1, 50 do
    local line = 4 + (i - 1) * 8 + 1 -- Position on my-button elements
    table.insert(hover_positions, {line = line, character = 10})
  end
  
  local function bulk_hover_test()
    local start_time = vim.uv.hrtime()
    local completed_requests = 0
    local total_requests = #hover_positions
    
    for _, pos in ipairs(hover_positions) do
      client.request('textDocument/hover', {
        textDocument = {uri = vim.uri_from_bufnr(bufnr)},
        position = pos
      }, function(err, result)
        completed_requests = completed_requests + 1
      end)
    end
    
    -- Wait for all requests to complete
    local success = vim.wait(10000, function()
      return completed_requests >= total_requests
    end)
    
    local duration = (vim.uv.hrtime() - start_time) / 1e6
    
    if not success then
      return nil, string.format("Only %d/%d hover requests completed", completed_requests, total_requests)
    end
    
    return {duration, completed_requests}
  end
  
  local bulk_hover_results = measurement.run_iterations(bulk_hover_test, 3, "bulk hover requests")
  table.insert(stress_results, {
    test = "bulk_hover",
    request_count = #hover_positions,
    iterations = 3,
    statistics = bulk_hover_results.statistics,
    success_rate = bulk_hover_results.success_rate,
    errors = bulk_hover_results.errors
  })
  
  -- Test 2: Rapid completion requests
  print("Running rapid completion stress test (30 requests)...")
  local completion_positions = {}
  for i = 1, 30 do
    local line = 4 + (i - 1) * 8 + 2 -- Position in my-card elements
    table.insert(completion_positions, {line = line, character = 15})
  end
  
  local function rapid_completion_test()
    local start_time = vim.uv.hrtime()
    local completed_requests = 0
    local total_requests = #completion_positions
    
    for _, pos in ipairs(completion_positions) do
      client.request('textDocument/completion', {
        textDocument = {uri = vim.uri_from_bufnr(bufnr)},
        position = pos
      }, function(err, result)
        completed_requests = completed_requests + 1
      end)
      
      -- Small delay to simulate rapid typing
      vim.wait(10)
    end
    
    -- Wait for all requests to complete
    local success = vim.wait(15000, function()
      return completed_requests >= total_requests
    end)
    
    local duration = (vim.uv.hrtime() - start_time) / 1e6
    
    if not success then
      return nil, string.format("Only %d/%d completion requests completed", completed_requests, total_requests)
    end
    
    return {duration, completed_requests}
  end
  
  local rapid_completion_results = measurement.run_iterations(rapid_completion_test, 3, "rapid completion requests")
  table.insert(stress_results, {
    test = "rapid_completion", 
    request_count = #completion_positions,
    iterations = 3,
    statistics = rapid_completion_results.statistics,
    success_rate = rapid_completion_results.success_rate,
    errors = rapid_completion_results.errors
  })
  
  -- Test 3: Document modification stress test
  print("Running document modification stress test...")
  
  local function document_modification_test()
    local start_time = vim.uv.hrtime()
    
    -- Make 20 rapid modifications
    for i = 1, 20 do
      local new_line = string.format('  <dynamic-element-%d>Added at runtime</dynamic-element-%d>', i, i)
      vim.api.nvim_buf_set_lines(bufnr, -2, -2, false, {new_line})
      vim.wait(50) -- Brief pause between modifications
    end
    
    -- Wait for processing
    vim.wait(2000)
    
    local duration = (vim.uv.hrtime() - start_time) / 1e6
    
    if client.is_stopped() then
      return nil, "Client stopped during document modifications"
    end
    
    return {duration, 20}
  end
  
  local doc_mod_results = measurement.run_iterations(document_modification_test, 3, "document modifications")
  table.insert(stress_results, {
    test = "document_modifications",
    modification_count = 20,
    iterations = 3,
    statistics = doc_mod_results.statistics,
    success_rate = doc_mod_results.success_rate,
    errors = doc_mod_results.errors
  })
  
  -- Clean up
  vim.api.nvim_buf_delete(bufnr, {force = true})
  vim.fn.delete('test-stress.html')
  client.stop()
  
  -- Calculate overall stress test success
  local total_tests = #stress_results
  local successful_tests = 0
  
  for _, result in ipairs(stress_results) do
    if result.success_rate and result.success_rate > 0.8 then -- 80% success threshold
      successful_tests = successful_tests + 1
    end
  end
  
  return {
    success = successful_tests > 0,
    server_name = server_name,
    total_stress_tests = total_tests,
    successful_stress_tests = successful_tests,
    overall_stress_success_rate = successful_tests / total_tests,
    stress_results = stress_results,
    document_size_lines = #large_html_content,
    client_survived = not client.is_stopped()
  }
end

return M