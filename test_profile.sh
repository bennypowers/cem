#!/bin/bash
# Quick profile test
cd lsp/benchmark
timeout 30s nvim --headless --clean -u configs/cem-minimal.lua << 'VIMEOF'
lua << LUA
local config = _G.BENCHMARK_LSP_CONFIG
vim.cmd('cd fixtures/large_project')
local client = vim.lsp.start(config)
vim.wait(5000, function() return vim.lsp.get_client_by_id(client).server_capabilities ~= nil end)

vim.cmd('edit lit-template.ts')
vim.cmd('set filetype=typescript')
local bufnr = vim.api.nvim_get_current_buf()
vim.wait(2000)

print("=== Running 20 completion requests at line 76, char 45 ===")
local total = 0
for i = 1, 20 do
  local start = vim.uv.hrtime()
  local done = false
  
  vim.lsp.get_client_by_id(client).request('textDocument/completion', {
    textDocument = {uri = vim.uri_from_bufnr(bufnr)},
    position = {line = 75, character = 45}
  }, function(err, res)
    done = true
  end)
  
  vim.wait(2000, function() return done end)
  local duration = (vim.uv.hrtime() - start) / 1e6
  total = total + duration
end

print(string.format("Average: %.2fms per completion", total / 20))
vim.lsp.get_client_by_id(client).stop()
vim.cmd('qa!')
LUA
VIMEOF
