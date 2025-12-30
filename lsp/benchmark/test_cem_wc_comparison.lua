-- Direct attribute completion comparison
print("=== CEM vs wc-toolkit Attribute Completion Comparison ===\n")

-- Position: line 76 (0-indexed: 75), after "align="
-- <my-container-flex justify="space-between" align="center">
local test_position = {line = 75, character = 56}

-- CEM test
local cem_config = _G.BENCHMARK_LSP_CONFIG or {
  name = 'cem-lsp',
  cmd = {'../../dist/cem', 'lsp'},
  filetypes = {'typescript'},
  single_file_support = true,
}

vim.cmd('edit fixtures/large_project/lit-template.ts')
vim.cmd('set filetype=typescript')
local bufnr = vim.api.nvim_get_current_buf()

local cem_client_id = vim.lsp.start(cem_config)
local cem_client = vim.lsp.get_client_by_id(cem_client_id)
vim.wait(3000, function() return cem_client.server_capabilities ~= nil end)
vim.wait(1000)

local cem_done = false
local cem_result = nil
cem_client.request('textDocument/completion', {
  textDocument = {uri = vim.uri_from_bufnr(bufnr)},
  position = test_position
}, function(err, result)
  cem_result = result
  cem_done = true
end)

vim.wait(3000, function() return cem_done end)

print("CEM Results:")
if cem_result then
  local items = cem_result.items or cem_result or {}
  print("  Count: " .. #items)
  for i = 1, math.min(15, #items) do
    print(string.format("  %2d. %-20s (kind: %s)", i, items[i].label, items[i].kind or "?"))
  end
else
  print("  No results")
end

cem_client.stop()
vim.api.nvim_buf_delete(bufnr, {force = true})

vim.cmd('qa!')
