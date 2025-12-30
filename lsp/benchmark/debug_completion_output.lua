-- Debug script to see actual completion output
local measurement = require('utils.measurement')

local config = _G.BENCHMARK_LSP_CONFIG
vim.cmd('cd fixtures/large_project')

local client_id = vim.lsp.start(config)
local client = vim.lsp.get_client_by_id(client_id)

vim.wait(5000, function()
  return client.server_capabilities ~= nil and not client.is_stopped()
end)

-- Open the HTML test file
vim.cmd('edit test-hover.html')
vim.cmd('set filetype=html')
local bufnr = vim.api.nvim_get_current_buf()

-- Write simple HTML with custom element
local html_content = [[<!DOCTYPE html>
<html>
<body>
  <my-button type="primary">Click</my-button>
</body>
</html>]]

vim.fn.writefile(vim.split(html_content, '\n'), 'test-simple.html')
vim.cmd('edit test-simple.html')
vim.cmd('set filetype=html')
bufnr = vim.api.nvim_get_current_buf()

vim.wait(2000)

print("=== Testing attribute completion after 'type=' ===")
local result = nil
local completed = false

client.request('textDocument/completion', {
  textDocument = {uri = vim.uri_from_bufnr(bufnr)},
  position = {line = 3, character = 24}  -- After type="
}, function(err, res)
  result = res
  completed = true
  print("Request completed. Error:", err, "Result:", res ~= nil)
end)

vim.wait(5000, function() return completed end)

if result then
  local items = result.items or result or {}
  print("Got " .. #items .. " completions:")
  for i = 1, math.min(20, #items) do
    print(string.format("  %d. %s", i, items[i].label))
  end
else
  print("No results received")
end

client.stop()
vim.fn.delete('test-simple.html')
vim.cmd('qa!')
