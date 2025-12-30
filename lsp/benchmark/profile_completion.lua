-- Profile completion vs hover to see the difference
local config = _G.BENCHMARK_LSP_CONFIG
vim.cmd('cd fixtures/large_project')

local client_id = vim.lsp.start(config)
local client = vim.lsp.get_client_by_id(client_id)
vim.wait(5000, function() return client.server_capabilities ~= nil end)

-- Open TypeScript template file
vim.cmd('edit lit-template.ts')
vim.cmd('set filetype=typescript')
local bufnr = vim.api.nvim_get_current_buf()
vim.wait(2000)

print("=== Timing comparison: hover vs completion ===\n")

-- Test hover (fast)
local hover_times = {}
for i = 1, 10 do
  local start = vim.uv.hrtime()
  local done = false
  
  client.request('textDocument/hover', {
    textDocument = {uri = vim.uri_from_bufnr(bufnr)},
    position = {line = 75, character = 12}
  }, function(err, res)
    done = true
  end)
  
  vim.wait(2000, function() return done end)
  local duration = (vim.uv.hrtime() - start) / 1e6
  table.insert(hover_times, duration)
end

local hover_avg = 0
for _, t in ipairs(hover_times) do hover_avg = hover_avg + t end
hover_avg = hover_avg / #hover_times

print(string.format("Hover: %.2fms average (10 iterations)", hover_avg))

-- Test completion (slow)
local comp_times = {}
for i = 1, 10 do
  local start = vim.uv.hrtime()
  local done = false
  
  client.request('textDocument/completion', {
    textDocument = {uri = vim.uri_from_bufnr(bufnr)},
    position = {line = 75, character = 45}
  }, function(err, res)
    done = true
  end)
  
  vim.wait(2000, function() return done end)
  local duration = (vim.uv.hrtime() - start) / 1e6
  table.insert(comp_times, duration)
end

local comp_avg = 0
for _, t in ipairs(comp_times) do comp_avg = comp_avg + t end
comp_avg = comp_avg / #comp_times

print(string.format("Completion: %.2fms average (10 iterations)", comp_avg))
print(string.format("\nCompletion is %.1fx slower than hover", comp_avg / hover_avg))

client.stop()
vim.cmd('qa!')
