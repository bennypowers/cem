--[[
Benchmark utilities for LSP testing
Provides common setup, teardown, and aggregation functions

IMPORTANT SIDE EFFECTS:
- setup_client() changes the GLOBAL working directory (vim.cmd("cd ..."))
- This affects all code in the same Neovim instance
- Benchmarks MUST be run sequentially, NOT in parallel
- Working directory is NOT automatically restored by cleanup_test()
]]

local measurement = require("utils.measurement")

local M = {}

--- Setup and initialize LSP client
---
--- WARNING: This function changes the GLOBAL working directory to fixture_dir.
--- The working directory is NOT automatically restored by cleanup_test().
---
--- @param config table LSP client configuration
--- @param fixture_dir string Directory to use as working directory
--- @return table|nil client LSP client if successful
--- @return string|nil error Error message if failed
function M.setup_client(config, fixture_dir)
	-- Change to fixture directory
	vim.cmd("cd " .. vim.fn.fnameescape(fixture_dir))

	-- Start LSP server
	local client_id = vim.lsp.start(config)
	if not client_id then
		return nil, "Failed to start LSP server"
	end

	local client = vim.lsp.get_client_by_id(client_id)
	if not client then
		return nil, "Failed to get client by ID"
	end

	-- Wait for initialization
	local ready = vim.wait(10000, function()
		return client.server_capabilities ~= nil and not client:is_stopped()
	end)

	if not ready then
		client:stop()
		return nil, "Server failed to initialize"
	end

	return client
end

--- Setup test file from fixture
--- @param fixture_dir string Directory containing fixture files
--- @param test_name string Name of the test file (without extension)
--- @param filetype string File type (e.g., "html", "typescript")
--- @param wait_ms number|nil Time to wait for document processing (default: 2000)
--- @param client table LSP client to check for stopped state
--- @return number|nil bufnr Buffer number if successful
--- @return string|nil error Error message if failed
function M.setup_test_file(fixture_dir, test_name, filetype, wait_ms, client)
	wait_ms = wait_ms or 2000

	local fixture_file = fixture_dir .. "/" .. test_name .. "-test." .. filetype
	local test_file = "test-" .. test_name .. "." .. filetype

	-- Copy fixture to test file for LSP processing
	if vim.fn.filereadable(fixture_file) == 0 then
		return nil, "Fixture file not found: " .. fixture_file
	end

	local fixture_lines = vim.fn.readfile(fixture_file)
	if #fixture_lines == 0 then
		return nil, "Fixture file is empty: " .. fixture_file
	end

	vim.fn.writefile(fixture_lines, test_file)

	-- Open file
	vim.cmd("edit " .. vim.fn.fnameescape(test_file))
	vim.cmd("set filetype=" .. filetype)
	local bufnr = vim.api.nvim_get_current_buf()

	-- Wait for document processing
	vim.wait(wait_ms)

	if client and client:is_stopped() then
		vim.fn.delete(test_file)
		return nil, "Client stopped during file processing"
	end

	return bufnr, test_file
end

--- Cleanup test resources
--- @param bufnr number|nil Buffer to delete
--- @param filename string|nil Test file to delete
--- @param client table LSP client to stop
function M.cleanup_test(bufnr, filename, client)
	if bufnr and vim.api.nvim_buf_is_valid(bufnr) then
		vim.api.nvim_buf_delete(bufnr, { force = true })
	end

	if filename then
		vim.fn.delete(filename)
	end

	if client and not client:is_stopped() then
		client:stop()
	end
end

--- Aggregate statistics from multiple benchmark results
--- @param results table Array of benchmark results with successful_runs, iterations, and all_times_ms
--- @return table aggregate Aggregated statistics
function M.aggregate_results(results)
	local total_successful = 0
	local total_attempts = 0
	local all_times = {}

	for _, result in ipairs(results) do
		total_successful = total_successful + (result.successful_runs or 0)
		total_attempts = total_attempts + (result.iterations or 0)

		-- Collect all times
		if result.all_times_ms then
			for _, time in ipairs(result.all_times_ms) do
				table.insert(all_times, time)
			end
		end
	end

	local overall_stats = measurement.calculate_statistics(all_times)

	return {
		total_successful = total_successful,
		total_attempts = total_attempts,
		success_rate = total_attempts > 0 and total_successful / total_attempts or 0,
		overall_statistics = overall_stats,
	}
end

return M
