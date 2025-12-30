-- Compare CEM vs wc-toolkit completion responses
local benchmark_dir = vim.fn.getcwd()
vim.opt.runtimepath:prepend(benchmark_dir)

-- Load CEM config
local cem_config = {
  name = 'cem-lsp',
  cmd = {'dist/cem', 'lsp'},
  root_dir = vim.fn.getcwd(),
  filetypes = {'html', 'typescript', 'javascript'},
  single_file_support = true,
  capabilities = vim.lsp.protocol.make_client_capabilities(),
}

local fixture_dir = 'fixtures/large_project'
vim.cmd('cd ' .. vim.fn.fnameescape(fixture_dir))

print("=== Testing CEM LSP ===")
local cem_client = vim.lsp.start(cem_config)
if cem_client then
  local client = vim.lsp.get_client_by_id(cem_client)
  vim.wait(5000, function() return client.server_capabilities ~= nil end)
  
  vim.cmd('edit lit-template.ts')
  vim.cmd('set filetype=typescript')
  local bufnr = vim.api.nvim_get_current_buf()
  vim.wait(2000)
  
  -- Test attribute completion at line 76, character 45 (after "align=")
  local cem_result = nil
  local cem_completed = false
  
  client.request('textDocument/completion', {
    textDocument = {uri = vim.uri_from_bufnr(bufnr)},
    position = {line = 75, character = 50}
  }, function(err, result)
    cem_result = result
    cem_completed = true
  end)
  
  vim.wait(5000, function() return cem_completed end)
  
  if cem_result then
    local items = cem_result.items or cem_result or {}
    print("CEM returned " .. #items .. " completions:")
    for i = 1, math.min(10, #items) do
      local item = items[i]
      print(string.format("  %d. %s (kind=%s)", i, item.label, item.kind or "nil"))
    end
  else
    print("CEM returned no results")
  end
  
  client.stop()
  vim.api.nvim_buf_delete(bufnr, {force = true})
end

vim.cmd('qa!')
