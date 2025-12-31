--[[
Workspace utilities for LSP benchmarking
Handles test workspace setup, file creation, and cleanup
]]

local M = {}
local original_dir = nil

-- Setup a temporary workspace for testing
function M.setup_workspace(fixture_name)
	-- Store original directory for cleanup
	original_dir = vim.fn.getcwd()
	-- Get benchmark directory from lua script path
	local benchmark_dir = vim.fn.fnamemodify(debug.getinfo(1).source:sub(2), ":h:h")
	local fixture_path = benchmark_dir .. "/fixtures/" .. fixture_name

	-- Verify fixture exists
	if vim.fn.isdirectory(fixture_path) == 0 then
		error(string.format("Fixture directory not found: %s", fixture_path))
	end

	-- Change to fixture directory
	vim.cmd("cd " .. vim.fn.fnameescape(fixture_path))

	return fixture_path
end

-- Open a file in Neovim buffer
function M.open_file(filepath)
	vim.cmd("edit " .. vim.fn.fnameescape(filepath))
	local bufnr = vim.api.nvim_get_current_buf()

	-- Wait for buffer to be fully loaded (necessary for LSP attachment and initial processing)
	vim.wait(100)

	-- Ensure LSP clients are attached
	local attached_clients = vim.lsp.get_clients({ bufnr = bufnr })
	if #attached_clients > 0 then
		-- Manually trigger didOpen if needed
		local uri = vim.uri_from_bufnr(bufnr)
		local content = table.concat(vim.api.nvim_buf_get_lines(bufnr, 0, -1, false), "\n")
		local filetype = vim.bo[bufnr].filetype

		for _, client in ipairs(attached_clients) do
			if client.server_capabilities.textDocumentSync then
				local params = {
					textDocument = {
						uri = uri,
						languageId = filetype,
						version = 1,
						text = content,
					},
				}
				client.notify("textDocument/didOpen", params)

				-- Wait for document to be processed
				vim.wait(200)
			end
		end
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
function M.wait_for_diagnostics(bufnr, timeout_ms)
	bufnr = bufnr or vim.api.nvim_get_current_buf()
	timeout_ms = timeout_ms or 2000

	local start = vim.uv.hrtime()
	local diagnostics = nil

	-- Set up temporary diagnostic handler
	local original_handler = vim.lsp.handlers["textDocument/publishDiagnostics"]

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

	-- Restore original handler
	vim.lsp.handlers["textDocument/publishDiagnostics"] = original_handler

	local duration = (vim.uv.hrtime() - start) / 1e6

	if not success then
		diagnostics = {} -- Return empty if timeout
	end

	return diagnostics, duration
end

-- Clean up test files and buffers
function M.cleanup()
	-- Close all buffers except initial
	local current_buf = vim.api.nvim_get_current_buf()
	for _, bufnr in ipairs(vim.api.nvim_list_bufs()) do
		if bufnr ~= current_buf and vim.api.nvim_buf_is_loaded(bufnr) then
			vim.api.nvim_buf_delete(bufnr, { force = true })
		end
	end

	-- Return to original directory
	if original_dir then
		vim.cmd("cd " .. vim.fn.fnameescape(original_dir))
		original_dir = nil
	end
end

return M
