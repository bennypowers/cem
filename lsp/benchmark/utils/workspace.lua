--[[
Workspace utilities for LSP benchmarking
Handles test workspace setup, file creation, and cleanup

IMPORTANT SIDE EFFECTS:
- setup_workspace() changes the GLOBAL working directory (vim.cmd("cd ..."))
- This affects all code in the same Neovim instance
- Benchmarks MUST be run sequentially, NOT in parallel
- Always call cleanup() to restore the original working directory
]]

local M = {}
-- Module-level original_dir for backward compatibility (prefer using return value)
local original_dir = nil

--- Setup a temporary workspace for testing
---
--- WARNING: This function changes the GLOBAL working directory using vim.cmd("cd ...").
--- This affects all code running in the same Neovim instance and can interfere with
--- concurrent or subsequent tests. Benchmarks should be run sequentially, not in parallel.
--- Always call M.cleanup() with the returned orig_dir to restore the working directory.
---
--- @param fixture_name string Name of the fixture directory
--- @param benchmark_dir string|nil Optional benchmark directory (defaults to auto-detection via debug.getinfo)
--- @return string fixture_path Path to the fixture directory
--- @return string original_dir Original working directory (for cleanup)
function M.setup_workspace(fixture_name, benchmark_dir)
	-- Store original directory for cleanup
	local orig_dir = vim.fn.getcwd()
	original_dir = orig_dir -- Module-level for backward compatibility

	-- Get benchmark directory: use parameter if provided, otherwise auto-detect
	if not benchmark_dir then
		-- Fallback: auto-detect from lua script path (fragile, may break if called from different contexts)
		benchmark_dir = vim.fn.fnamemodify(debug.getinfo(1).source:sub(2), ":h:h")
	end

	local fixture_path = benchmark_dir .. "/fixtures/" .. fixture_name

	-- Verify fixture exists
	if vim.fn.isdirectory(fixture_path) == 0 then
		error(string.format("Fixture directory not found: %s", fixture_path))
	end

	-- Change to fixture directory
	vim.cmd("cd " .. vim.fn.fnameescape(fixture_path))

	return fixture_path, orig_dir
end

--- Open a file in Neovim buffer and wait for LSP to attach
--- NOTE: Neovim's LSP client automatically sends textDocument/didOpen when a buffer
--- is opened with an attached LSP client. This function does NOT manually trigger didOpen
--- to avoid duplicate notifications.
--- @param filepath string Path to file to open
--- @param timeout_ms number|nil Timeout for LSP attachment (default: 2000ms)
--- @return number bufnr Buffer number
function M.open_file(filepath, timeout_ms)
	timeout_ms = timeout_ms or 2000

	vim.cmd("edit " .. vim.fn.fnameescape(filepath))
	local bufnr = vim.api.nvim_get_current_buf()

	-- Wait for buffer to be valid and loaded using condition-based polling
	local buffer_loaded = vim.wait(timeout_ms, function()
		return vim.api.nvim_buf_is_valid(bufnr) and vim.api.nvim_buf_is_loaded(bufnr)
	end)

	if not buffer_loaded then
		error(string.format("Buffer failed to load within %dms", timeout_ms))
	end

	-- Wait for at least one LSP client to attach to this buffer
	-- This ensures LSP is ready before returning, but doesn't manually trigger didOpen
	-- since Neovim's LSP client handles that automatically
	local lsp_attached = vim.wait(timeout_ms, function()
		local clients = vim.lsp.get_clients({ bufnr = bufnr })
		return #clients > 0
	end)

	if not lsp_attached then
		-- LSP attachment is optional - warn but don't error
		-- Some files may not have LSP support
		vim.notify(
			string.format("No LSP clients attached to buffer %d within %dms", bufnr, timeout_ms),
			vim.log.levels.WARN
		)
	end

	return bufnr
end

-- Get buffer content as string
function M.get_buffer_content(bufnr)
	bufnr = bufnr or vim.api.nvim_get_current_buf()
	local lines = vim.api.nvim_buf_get_lines(bufnr, 0, -1, false)
	return table.concat(lines, "\n")
end

-- Find text position in buffer (1-based line, 0-based character)
function M.find_text_position(text, bufnr)
	bufnr = bufnr or vim.api.nvim_get_current_buf()
	local lines = vim.api.nvim_buf_get_lines(bufnr, 0, -1, false)

	for line_num, line in ipairs(lines) do
		local col = string.find(line, text, 1, true) -- Plain text search
		if col then
			return line_num, col - 1 -- Convert to 0-based character position
		end
	end

	error(string.format("Text '%s' not found in buffer", text))
end

-- Wait for LSP diagnostics to be published
-- NOTE: This function modifies the global diagnostic handler temporarily.
-- For production use, prefer passing handlers in the client config via vim.lsp.start.
-- @param bufnr number Buffer number to wait for diagnostics
-- @param timeout_ms number Timeout in milliseconds
-- @return table diagnostics Diagnostics array
-- @return number duration Duration in milliseconds
function M.wait_for_diagnostics(bufnr, timeout_ms)
	bufnr = bufnr or vim.api.nvim_get_current_buf()
	timeout_ms = timeout_ms or 2000

	local start = vim.uv.hrtime()
	local diagnostics = nil

	-- Save original handler for safe restoration
	local original_handler = vim.lsp.handlers["textDocument/publishDiagnostics"]

	-- Use pcall to guarantee handler restoration even on error
	local ok, result = pcall(function()
		-- Set up temporary diagnostic handler
		vim.lsp.handlers["textDocument/publishDiagnostics"] = function(err, result, ctx, config)
			if result and result.uri == vim.uri_from_bufnr(bufnr) then
				diagnostics = result.diagnostics
			end
			-- Call original handler
			if original_handler then
				original_handler(err, result, ctx, config)
			end
		end

		-- Wait for diagnostics
		local success = vim.wait(timeout_ms, function()
			return diagnostics ~= nil
		end)

		return success
	end)

	-- Always restore original handler
	vim.lsp.handlers["textDocument/publishDiagnostics"] = original_handler

	local duration = (vim.uv.hrtime() - start) / 1e6

	if not ok or not result then
		diagnostics = {} -- Return empty if error or timeout
	end

	return diagnostics, duration
end

--- Clean up test files and buffers
---
--- WARNING: This function restores the GLOBAL working directory using vim.cmd("cd ...").
--- Should be called after setup_workspace() to restore the original working directory.
---
--- @param orig_dir string|nil Optional original directory to restore (uses module-level if not provided)
function M.cleanup(orig_dir)
	-- Close all buffers except initial
	local current_buf = vim.api.nvim_get_current_buf()
	for _, bufnr in ipairs(vim.api.nvim_list_bufs()) do
		if bufnr ~= current_buf and vim.api.nvim_buf_is_loaded(bufnr) then
			vim.api.nvim_buf_delete(bufnr, { force = true })
		end
	end

	-- Return to original directory: use parameter if provided, otherwise module-level
	local dir_to_restore = orig_dir or original_dir
	if dir_to_restore then
		vim.cmd("cd " .. vim.fn.fnameescape(dir_to_restore))
		if not orig_dir then
			-- Only clear module-level if we used it
			original_dir = nil
		end
	end
end

return M
