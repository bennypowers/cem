-- Startup benchmark module
-- Tests LSP server initialization and readiness time

local measurement = require('utils.measurement')

local M = {}

function M.run_startup_benchmark(config, fixture_dir)
  local server_name = _G.BENCHMARK_LSP_NAME or "unknown"
  print(string.format("=== %s LSP Startup Benchmark ===", server_name))
  
  -- Change to fixture directory
  vim.cmd('cd ' .. vim.fn.fnameescape(fixture_dir))
  print("Working directory:", vim.fn.getcwd())
  
  -- Run multiple iterations for statistical analysis
  local iterations = 20 -- Increased for better statistical confidence
  
  local function single_startup()
    local startup_start = vim.uv.hrtime()
    
    -- Start LSP server
    local client_id = vim.lsp.start(config)
    if not client_id then
      return nil, "Failed to start LSP server"
    end
    
    local client = vim.lsp.get_client_by_id(client_id)
    
    -- Wait for initialization
    local ready = vim.wait(10000, function()
      return client.server_capabilities ~= nil and not client.is_stopped()
    end)
    
    local startup_duration = (vim.uv.hrtime() - startup_start) / 1e6
    
    if not ready then
      client.stop()
      return nil, "Server failed to initialize"
    end
    
    local capabilities_count = vim.tbl_count(client.server_capabilities or {})
    
    -- Clean shutdown
    client.stop()
    vim.wait(100) -- Allow cleanup
    
    return {startup_duration, capabilities_count}
  end
  
  -- Collect multiple measurements
  local results = measurement.run_iterations(single_startup, iterations, "startup")
  
  if results.successful_runs == 0 then
    return {
      success = false,
      error = "All startup attempts failed",
      iterations = iterations,
      failed_runs = results.failed_runs,
      errors = results.errors
    }
  end
  
  return {
    success = true,
    server_name = server_name,
    iterations = iterations,
    successful_runs = results.successful_runs,
    failed_runs = results.failed_runs,
    success_rate = results.success_rate,
    statistics = results.statistics,
    memory_usage_bytes = results.memory_usage_bytes,
    all_times_ms = results.times_ms,
    errors = results.errors,
    -- Keep backward compatibility
    duration_ms = results.statistics.mean,
    capabilities_count = results.results[1] and results.results[1][2] or 0
  }
end

return M