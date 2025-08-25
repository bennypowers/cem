-- Attribute hover benchmark module
-- Tests LSP hover accuracy for attributes in large manifests

local measurement = require('utils.measurement')

local M = {}

function M.run_attribute_hover_benchmark(config, fixture_dir)
  local server_name = _G.BENCHMARK_LSP_NAME or "unknown"
  print(string.format("=== %s LSP Attribute Hover Benchmark ===", server_name))
  
  -- Change to fixture directory (use large project for comprehensive testing)
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
  
  -- Create HTML content with various attribute types to test hover accuracy
  local html_content = [[<!DOCTYPE html>
<html>
<head><title>Attribute Hover Test</title></head>
<body>
  <!-- Button attributes -->
  <my-button variant="primary" size="large" disabled loading icon="star">Primary Button</my-button>
  <my-button-secondary variant="outline" size="medium">Secondary</my-button-secondary>
  <my-button-icon icon="save" size="small" disabled></my-button-icon>
  
  <!-- Form attributes -->
  <my-input-text value="test" placeholder="Enter text" required readonly></my-input-text>
  <my-input-email placeholder="user@example.com" disabled></my-input-email>
  <my-input-select multiple searchable size="large">
    <option value="1">Option 1</option>
  </my-input-select>
  
  <!-- Layout attributes -->
  <my-container-flex gap="large" padding="medium" justify="space-between" align="center">
    <my-card-basic variant="elevated" clickable selected>
      <h3 slot="header">Card Title</h3>
      <p>Card content</p>
    </my-card-basic>
  </my-container-flex>
  
  <!-- Complex attributes -->
  <my-chart-line data='[{"x": 1, "y": 10}]' responsive animated></my-chart-line>
  <my-table-sortable striped bordered paginated page-size="20"></my-table-sortable>
  
  <!-- Media attributes -->
  <my-image-responsive src="/test.jpg" alt="Test image" loading="lazy" fit="cover" ratio="16:9"></my-image-responsive>
  <my-image-avatar src="/avatar.jpg" size="large"></my-image-avatar>
  
  <!-- Overlay attributes -->
  <my-modal-dialog open size="medium" closable backdrop-close="true">
    <h2 slot="header">Modal Title</h2>
    <p>Modal content</p>
  </my-modal-dialog>
  
  <!-- Feedback attributes -->
  <my-alert-info dismissible icon="info" timeout="5000">Alert message</my-alert-info>
  <my-progress-linear value="75" variant="success"></my-progress-linear>
  <my-spinner-primary size="medium" color="blue"></my-spinner-primary>
</body>
</html>]]
  
  vim.fn.writefile(vim.split(html_content, '\n'), 'test-attribute-hover.html')
  
  -- Open file
  vim.cmd('edit test-attribute-hover.html')
  vim.cmd('set filetype=html')
  local bufnr = vim.api.nvim_get_current_buf()
  
  -- Wait for document processing
  vim.wait(3000)
  
  if client.is_stopped() then
    vim.fn.delete('test-attribute-hover.html')
    return {
      success = false,
      error = "Client stopped during file processing"
    }
  end
  
  -- Test positions for different attribute types
  local test_positions = {
    -- Button attributes
    {element = "my-button", attr = "variant", line = 5, character = 18, expected_content = "Button style variant"},
    {element = "my-button", attr = "size", line = 5, character = 35, expected_content = "Button size"},
    {element = "my-button", attr = "disabled", line = 5, character = 48, expected_content = "Disabled state"},
    {element = "my-button", attr = "loading", line = 5, character = 57, expected_content = "Loading state"},
    {element = "my-button", attr = "icon", line = 5, character = 67, expected_content = "Icon name"},
    
    -- Form attributes
    {element = "my-input-text", attr = "value", line = 10, character = 22, expected_content = "Input value"},
    {element = "my-input-text", attr = "placeholder", line = 10, character = 35, expected_content = "Placeholder text"},
    {element = "my-input-text", attr = "required", line = 10, character = 60, expected_content = "Required field"},
    {element = "my-input-select", attr = "multiple", line = 12, character = 25, expected_content = "Multiple selection"},
    {element = "my-input-select", attr = "searchable", line = 12, character = 35, expected_content = "searchable"},
    
    -- Layout attributes
    {element = "my-container-flex", attr = "gap", line = 17, character = 28, expected_content = "Gap size"},
    {element = "my-container-flex", attr = "justify", line = 17, character = 55, expected_content = "Justification"},
    {element = "my-card-basic", attr = "variant", line = 18, character = 30, expected_content = "Card style"},
    {element = "my-card-basic", attr = "clickable", line = 18, character = 50, expected_content = "Clickable card"},
    
    -- Chart attributes
    {element = "my-chart-line", attr = "responsive", line = 25, character = 50, expected_content = "responsive"},
    {element = "my-chart-line", attr = "animated", line = 25, character = 62, expected_content = "animated"},
    
    -- Table attributes
    {element = "my-table-sortable", attr = "striped", line = 26, character = 28, expected_content = "striped"},
    {element = "my-table-sortable", attr = "page-size", line = 26, character = 60, expected_content = "page-size"},
    
    -- Media attributes
    {element = "my-image-responsive", attr = "loading", line = 29, character = 65, expected_content = "Loading strategy"},
    {element = "my-image-responsive", attr = "fit", line = 29, character = 82, expected_content = "fit"},
    {element = "my-image-avatar", attr = "size", line = 30, character = 40, expected_content = "size"}
  }
  
  local hover_results = {}
  local iterations_per_attribute = 3 -- Test each attribute multiple times
  
  for i, test_pos in ipairs(test_positions) do
    print(string.format("Testing hover on %s.%s (%d iterations)...", test_pos.element, test_pos.attr, iterations_per_attribute))
    
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
      
      local content = ""
      if hover_result.contents then
        if type(hover_result.contents) == "string" then
          content = hover_result.contents
        elseif hover_result.contents.value then
          content = hover_result.contents.value
        end
      end
      
      -- Check if content matches expected (basic substring check)
      local content_matches = test_pos.expected_content and 
        string.find(content:lower(), test_pos.expected_content:lower()) ~= nil
      
      return {hover_duration, string.len(content), content_matches and 1 or 0, content}
    end
    
    -- Run multiple iterations for this attribute
    local attr_results = measurement.run_iterations(single_hover, iterations_per_attribute, 
      string.format("hover on %s.%s", test_pos.element, test_pos.attr))
    
    -- Calculate content accuracy
    local accurate_hovers = 0
    local total_content_length = 0
    
    for _, result in ipairs(attr_results.results or {}) do
      if result[3] == 1 then accurate_hovers = accurate_hovers + 1 end
      total_content_length = total_content_length + (result[2] or 0)
    end
    
    table.insert(hover_results, {
      element = test_pos.element,
      attribute = test_pos.attr,
      iterations = iterations_per_attribute,
      successful_runs = attr_results.successful_runs,
      failed_runs = attr_results.failed_runs,
      success_rate = attr_results.success_rate,
      statistics = attr_results.statistics,
      accuracy_rate = attr_results.successful_runs > 0 and accurate_hovers / attr_results.successful_runs or 0,
      avg_content_length = attr_results.successful_runs > 0 and total_content_length / attr_results.successful_runs or 0,
      expected_content = test_pos.expected_content,
      sample_content = attr_results.results and #attr_results.results > 0 and attr_results.results[1][4] or "",
      errors = attr_results.errors
    })
    
    -- Check if client survived
    if client.is_stopped() then
      break
    end
    
    -- Small delay between different attributes
    vim.wait(50)
  end
  
  -- Clean up
  vim.api.nvim_buf_delete(bufnr, {force = true})
  vim.fn.delete('test-attribute-hover.html')
  client.stop()
  
  -- Calculate overall statistics
  local total_successful = 0
  local total_attempts = 0
  local total_accurate = 0
  local all_times = {}
  
  for _, result in ipairs(hover_results) do
    total_successful = total_successful + result.successful_runs
    total_attempts = total_attempts + result.iterations
    total_accurate = total_accurate + (result.accuracy_rate * result.successful_runs)
    for _, time in ipairs(result.statistics and {result.statistics.mean} or {}) do
      table.insert(all_times, time)
    end
  end
  
  local overall_stats = measurement.calculate_statistics(all_times)
  
  return {
    success = total_successful > 0,
    server_name = server_name,
    total_attributes_tested = #test_positions,
    total_attempts = total_attempts,
    total_successful = total_successful,
    success_rate = total_attempts > 0 and total_successful / total_attempts or 0,
    accuracy_rate = total_successful > 0 and total_accurate / total_successful or 0,
    overall_statistics = overall_stats,
    attribute_results = hover_results,
    client_survived = not client.is_stopped()
  }
end

return M