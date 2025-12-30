-- Test if LSP loads the manifest
local config = _G.BENCHMARK_LSP_CONFIG
vim.cmd('cd fixtures/large_project')

print("Working directory:", vim.fn.getcwd())
print("Manifest exists:", vim.fn.filereadable('custom-elements.json'))

local client_id = vim.lsp.start(config)
local client = vim.lsp.get_client_by_id(client_id)

print("Client started:", client_id)

vim.wait(5000, function()
  return client.server_capabilities ~= nil
end)

print("Server capabilities:", vim.inspect(client.server_capabilities))

-- Test hover instead of completion
local html_content = [[<!DOCTYPE html>
<html><body>
  <my-button>Click</my-button>
</body></html>]]

vim.fn.writefile(vim.split(html_content, '\n'), 'test.html')
vim.cmd('edit test.html')
vim.cmd('set filetype=html')
local bufnr = vim.api.nvim_get_current_buf()
vim.wait(2000)

print("\n=== Testing hover on my-button ===")
local result = nil
local completed = false

client.request('textDocument/hover', {
  textDocument = {uri = vim.uri_from_bufnr(bufnr)},
  position = {line = 2, character = 5}
}, function(err, res)
  result = res
  completed = true
end)

vim.wait(5000, function() return completed end)

if result and result.contents then
  print("Hover SUCCESS! Content preview:")
  if type(result.contents) == "table" and result.contents.value then
    print(result.contents.value:sub(1, 200))
  end
else
  print("Hover FAILED - no content")
end

client.stop()
vim.fn.delete('test.html')
vim.cmd('qa!')
