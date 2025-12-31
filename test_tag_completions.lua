-- Test tag name completions with the fix
local config = _G.BENCHMARK_LSP_CONFIG
vim.cmd('cd fixtures/large_project')

local client_id = vim.lsp.start(config)
local client = vim.lsp.get_client_by_id(client_id)
vim.wait(5000, function() return client.server_capabilities ~= nil end)

local html = [[<!DOCTYPE html>
<html><body>
  <my-
</body></html>]]

vim.fn.writefile(vim.split(html, '\n'), 'test.html')
vim.cmd('edit test.html')
vim.cmd('set filetype=html')
local bufnr = vim.api.nvim_get_current_buf()
vim.wait(2000)

print("=== Tag name completions after '<my-' ===")
local result = nil
local done = false

client.request('textDocument/completion', {
  textDocument = {uri = vim.uri_from_bufnr(bufnr)},
  position = {line = 2, character = 6}
}, function(err, res)
  result = res
  done = true
end)

vim.wait(5000, function() return done end)

if result then
  local items = result.items or result or {}
  print("Got " .. #items .. " custom element completions:")
  for i = 1, math.min(20, #items) do
    print("  " .. items[i].label)
  end
else
  print("No results")
end

client.stop()
vim.fn.delete('test.html')
vim.cmd('qa!')
