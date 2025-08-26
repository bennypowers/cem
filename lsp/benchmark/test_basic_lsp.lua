#!/usr/bin/env nvim --headless --clean -l
-- Basic LSP functionality test
-- Usage: nvim --headless --clean -u configs/cem-minimal.lua -l test_basic_lsp.lua

-- Ensure the benchmark directory is in the path
local script_dir = vim.fn.getcwd()
package.path = script_dir .. '/?.lua;' .. package.path

local function test_basic_lsp_features()
  local server_name = _G.BENCHMARK_LSP_NAME or "unknown"
  local config = _G.BENCHMARK_LSP_CONFIG
  
  if not config then
    print("ERROR: No LSP configuration found.")
    return
  end
  
  print(string.format("Testing basic LSP features for %s LSP server", server_name))
  
  local fixture_dir = script_dir .. '/fixtures/large_project'
  local test_file = fixture_dir .. '/rapid-typing-test.html'
  
  if vim.fn.filereadable(test_file) == 0 then
    print("ERROR: Test file not found: " .. test_file)
    return
  end
  
  local results = {
    lsp_initialized = false,
    server_capabilities = {},
    hover_working = false,
    workspace_diagnostics = false,
    document_symbols = false,
    errors = {}
  }
  
  local success, error_msg = pcall(function()
    -- Open test document
    vim.cmd('edit ' .. test_file)
    vim.wait(3000, function() return false end) -- Wait for LSP initialization
    
    -- Check if LSP client is attached
    local clients = vim.lsp.get_active_clients()
    local cem_client = nil
    
    for _, client in ipairs(clients) do
      if client.name == server_name then
        cem_client = client
        results.lsp_initialized = true
        results.server_capabilities = client.server_capabilities or {}
        break
      end
    end
    
    if not cem_client then
      table.insert(results.errors, "LSP client not found or not attached")
      return
    end
    
    print(string.format("✅ LSP client '%s' successfully attached", server_name))
    
    -- Test basic capabilities
    local caps = results.server_capabilities
    print("Server capabilities:")
    print(string.format("  Hover: %s", caps.hoverProvider and "✅" or "❌"))
    print(string.format("  Completion: %s", caps.completionProvider and "✅" or "❌"))
    print(string.format("  Definition: %s", caps.definitionProvider and "✅" or "❌"))
    print(string.format("  References: %s", caps.referencesProvider and "✅" or "❌"))
    print(string.format("  Diagnostics: %s", caps.textDocumentSync and "✅" or "❌"))
    print(string.format("  Workspace symbols: %s", caps.workspaceSymbolProvider and "✅" or "❌"))
    
    -- Set up some test content
    vim.cmd('normal! ggdG')
    vim.api.nvim_buf_set_lines(0, 0, -1, false, {'<my-card variant="outlined">', '  <h3>Test</h3>', '</my-card>'})
    
    -- Test hover
    vim.api.nvim_win_set_cursor(0, {1, 5}) -- Position on 'my-card'
    
    local hover_result = nil
    vim.lsp.handlers['textDocument/hover'] = function(err, result, ctx, config)
      if err then
        table.insert(results.errors, "Hover error: " .. tostring(err))
      else
        hover_result = result
        if result and result.contents then
          results.hover_working = true
        end
      end
    end
    
    vim.lsp.buf.hover()
    vim.wait(3000, function() return hover_result ~= nil end, 100)
    
    if results.hover_working then
      print("✅ Hover is working")
    else
      print("❌ Hover not working or timed out")
    end
    
    -- Wait for any diagnostics to arrive
    vim.wait(2000, function() return false end)
    
    local diagnostics = vim.diagnostic.get(0)
    if #diagnostics > 0 then
      results.workspace_diagnostics = true
      print(string.format("✅ Diagnostics working (%d found)", #diagnostics))
    else
      print("⚠️  No diagnostics found (may be expected)")
    end
    
  end)
  
  if success then
    print("\n" .. string.rep("=", 50))
    print(string.format("BASIC LSP TEST SUMMARY FOR %s", string.upper(server_name)))
    print(string.rep("=", 50))
    print(string.format("LSP Initialized: %s", results.lsp_initialized and "✅" or "❌"))
    print(string.format("Hover Working: %s", results.hover_working and "✅" or "❌"))
    print(string.format("Diagnostics: %s", results.workspace_diagnostics and "✅" or "⚠️"))
    
    local capabilities_count = 0
    for cap_name, cap_value in pairs(results.server_capabilities) do
      if cap_value then
        capabilities_count = capabilities_count + 1
      end
    end
    
    print(string.format("Active Capabilities: %d", capabilities_count))
    
    if #results.errors > 0 then
      print("\nErrors:")
      for _, error in ipairs(results.errors) do
        print("  " .. error)
      end
    end
    
    local overall_health = results.lsp_initialized and results.hover_working
    if overall_health then
      print("🎉 LSP server is working correctly")
    else
      print("🚨 LSP server has issues")
    end
  else
    print(string.format("❌ Test failed: %s", error_msg))
  end
end

test_basic_lsp_features()