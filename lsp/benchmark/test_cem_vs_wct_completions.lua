-- Compare actual completion output
local config = _G.BENCHMARK_LSP_CONFIG
vim.cmd('cd fixtures/large_project')

local client_id = vim.lsp.start(config)
local client = vim.lsp.get_client_by_id(client_id)
vim.wait(5000, function() return client.server_capabilities ~= nil end)

-- Test in simple HTML
local html = [[<!DOCTYPE html>
<html><body>
  <my-button >Click</my-button>
</body></html>]]

vim.fn.writefile(vim.split(html, '\n'), 'test.html')
vim.cmd('edit test.html')
vim.cmd('set filetype=html')
local bufnr = vim.api.nvim_get_current_buf()
vim.wait(2000)

print("=== CEM Attribute Completions (HTML file) ===")
local result = nil
local done = false

-- Position after space in <my-button >
client.request('textDocument/completion', {
  textDocument = {uri = vim.uri_from_bufnr(bufnr)},
  position = {line = 2, character = 14}
}, function(err, res)
  result = res
  done = true
end)

vim.wait(5000, function() return done end)

if result then
  local items = result.items or result or {}
  print("CEM returned " .. #items .. " items:")
  for i = 1, #items do
    print(string.format("  %d. %s", i, items[i].label))
  end
else
  print("CEM: No results")
end

client.stop()
vim.fn.delete('test.html')

-- Now test wc-toolkit
print("\n=== wc-toolkit Attribute Completions (HTML file) ===")

-- Load wc-toolkit config
local wc_config = {
  name = 'wc-language-server',
  cmd = { 'wc-language-server', '--stdio' },
  filetypes = {'html'},
  single_file_support = true,
  capabilities = vim.lsp.protocol.make_client_capabilities(),
}

local wc_client_id = vim.lsp.start(wc_config)
local wc_client = vim.lsp.get_client_by_id(wc_client_id)
vim.wait(5000, function() return wc_client.server_capabilities ~= nil end)

vim.fn.writefile(vim.split(html, '\n'), 'test.html')
vim.cmd('edit test.html')
vim.cmd('set filetype=html')
bufnr = vim.api.nvim_get_current_buf()
vim.wait(2000)

local wc_result = nil
local wc_done = false

wc_client.request('textDocument/completion', {
  textDocument = {uri = vim.uri_from_bufnr(bufnr)},
  position = {line = 2, character = 14}
}, function(err, res)
  wc_result = res
  wc_done = true
end)

vim.wait(5000, function() return wc_done end)

if wc_result then
  local items = wc_result.items or wc_result or {}
  print("wc-toolkit returned " .. #items .. " items:")
  for i = 1, math.min(20, #items) do
    print(string.format("  %d. %s", i, items[i].label))
  end
  if #items > 20 then
    print("  ... and " .. (#items - 20) .. " more")
  end
else
  print("wc-toolkit: No results")
end

wc_client.stop()
vim.fn.delete('test.html')
vim.cmd('qa!')
