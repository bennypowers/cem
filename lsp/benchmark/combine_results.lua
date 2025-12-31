#!/usr/bin/env nvim --headless --clean -l
-- Combine CEM and wc-toolkit benchmark results into a single timestamped file

local cem_file = '/tmp/cem-benchmark-cem.json'
local wct_file = '/tmp/cem-benchmark-wc-toolkit.json'

-- Read both result files
local function read_json_file(filepath)
  local lines = vim.fn.readfile(filepath)
  if #lines == 0 then
    return nil
  end
  local content = table.concat(lines, '\n')
  return vim.fn.json_decode(content)
end

-- Check if both files exist
if vim.fn.filereadable(cem_file) == 0 then
  print("ERROR: CEM results not found at " .. cem_file)
  vim.cmd('quit')
  return
end

if vim.fn.filereadable(wct_file) == 0 then
  print("ERROR: wc-toolkit results not found at " .. wct_file)
  vim.cmd('quit')
  return
end

local cem_results = read_json_file(cem_file)
local wct_results = read_json_file(wct_file)

if not cem_results or not wct_results then
  print("ERROR: Failed to parse JSON results")
  vim.cmd('quit')
  return
end

-- Create combined result structure
local combined = {
  timestamp = os.date("%Y-%m-%d %H:%M:%S"),
  servers = {
    cem = cem_results,
    ["wc-toolkit"] = wct_results
  }
}

-- Get script directory
local script_dir = vim.fn.getcwd()

-- Save to docs/data for documentation site
local docs_data_dir = script_dir .. '/../../docs/data'
vim.fn.mkdir(docs_data_dir, 'p')
local docs_file = docs_data_dir .. '/lsp-benchmark-results.json'
local json_content = vim.fn.json_encode(combined)
vim.fn.writefile({json_content}, docs_file)
print(string.format("✅ Combined results saved to: docs/data/lsp-benchmark-results.json"))

-- Also save timestamped copy to results/ directory for history
local results_dir = script_dir .. '/results'
vim.fn.mkdir(results_dir, 'p')
local timestamped_file = string.format('%s/benchmark-results-%s.json', results_dir, os.date("%Y%m%d_%H%M%S"))
vim.fn.writefile({json_content}, timestamped_file)
print(string.format("📊 Timestamped copy saved to: %s", timestamped_file))

-- Clean up temp files
vim.fn.delete(cem_file)
vim.fn.delete(wct_file)
print("🧹 Cleaned up temporary files")

vim.cmd('quit')
