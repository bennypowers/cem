-- Replicate exact benchmark scenario
local config = _G.BENCHMARK_LSP_CONFIG
vim.cmd('cd fixtures/large_project')

local client_id = vim.lsp.start(config)
local client = vim.lsp.get_client_by_id(client_id)
vim.wait(10000, function() return client.server_capabilities ~= nil end)

local html_content = [[<!DOCTYPE html>
<html>
<head><title>Completion Test</title></head>
<body>
  <my-
  <my-button >
    <my-card >
</body>
</html>]]

vim.fn.writefile(vim.split(html_content, '\n'), 'test-completion.html')
vim.cmd('edit test-completion.html')
vim.cmd('set filetype=html')
local bufnr = vim.api.nvim_get_current_buf()
vim.wait(2000)

print("=== Test 1: Tag name completion (after 'my-') ===")
local result1 = nil
local done1 = false
client.request('textDocument/completion', {
  textDocument = {uri = vim.uri_from_bufnr(bufnr)},
  position = {line = 4, character = 6}
}, function(err, res)
  result1 = res
  done1 = true
end)
vim.wait(3000, function() return done1 end)

if result1 then
  local items = result1.items or result1 or {}
  print("  Got " .. #items .. " completions")
  for i = 1, math.min(10, #items) do
    print("    " .. items[i].label)
  end
else
  print("  No results")
end

print("\n=== Test 2: Attribute completion (inside my-button tag) ===")
local result2 = nil
local done2 = false
client.request('textDocument/completion', {
  textDocument = {uri = vim.uri_from_bufnr(bufnr)},
  position = {line = 5, character = 13}
}, function(err, res)
  result2 = res
  done2 = true
end)
vim.wait(3000, function() return done2 end)

if result2 then
  local items = result2.items or result2 or {}
  print("  Got " .. #items .. " completions:")
  for i = 1, #items do
    print("    " .. items[i].label)
  end
else
  print("  No results")
end

client.stop()
vim.fn.delete('test-completion.html')
vim.cmd('qa!')
