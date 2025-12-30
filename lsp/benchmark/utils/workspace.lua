--[[
Workspace utilities for LSP benchmarking
Handles test workspace setup, file creation, and cleanup
]]

local M = {}

-- Setup a temporary workspace for testing
function M.setup_workspace(fixture_name)
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

-- Create a test file with specified content
function M.create_test_file(filename, content)
	local filepath = vim.fn.expand(filename)

	-- Ensure directory exists
	local dir = vim.fn.fnamemodify(filepath, ":h")
	vim.fn.mkdir(dir, "p")

	-- Write content to file
	local lines = vim.split(content, "\n")
	vim.fn.writefile(lines, filepath)

	return filepath
end

-- Open a file in Neovim buffer
function M.open_file(filepath)
	vim.cmd("edit " .. vim.fn.fnameescape(filepath))
	local bufnr = vim.api.nvim_get_current_buf()

	-- Wait for buffer to be fully loaded
	vim.wait(100)

	-- Ensure LSP clients are attached
	local attached_clients = vim.lsp.get_active_clients({ bufnr = bufnr })
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

-- Set cursor position in buffer
function M.set_cursor(line, col)
	vim.api.nvim_win_set_cursor(0, { line, col })
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

-- Update file content and trigger LSP change events
function M.update_file_content(filepath, new_content)
	local bufnr = vim.fn.bufnr(filepath)

	if bufnr == -1 then
		-- File not open, create/update directly
		local lines = vim.split(new_content, "\n")
		vim.fn.writefile(lines, filepath)
	else
		-- File is open in buffer, update buffer content
		local lines = vim.split(new_content, "\n")
		vim.api.nvim_buf_set_lines(bufnr, 0, -1, false, lines)

		-- Save buffer to trigger file system events
		vim.api.nvim_buf_call(bufnr, function()
			vim.cmd("write")
		end)
	end
end

-- Wait for LSP diagnostics to be published
function M.wait_for_diagnostics(bufnr, timeout_ms)
	bufnr = bufnr or vim.api.nvim_get_current_buf()
	timeout_ms = timeout_ms or 2000

	local start = vim.uv.hrtime()
	local diagnostics = nil

	-- Set up temporary diagnostic handler
	local handler_set = false
	local original_handler = vim.lsp.handlers["textDocument/publishDiagnostics"]

	vim.lsp.handlers["textDocument/publishDiagnostics"] = function(err, result, ctx, config)
		if result and result.uri == vim.uri_from_bufnr(bufnr) then
			diagnostics = result.diagnostics
			handler_set = true
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
	vim.cmd("cd -")
end

-- Generate test HTML content with custom elements
function M.generate_html_content(elements)
	elements = elements or { "my-element" }

	local content = [[<!DOCTYPE html>
<html>
<head>
  <title>Test Document</title>
</head>
<body>
]]

	for _, element in ipairs(elements) do
		content = content .. string.format("  <%s></%s>\n", element, element)
	end

	content = content .. [[</body>
</html>]]

	return content
end

-- Generate test TypeScript content with Lit templates
function M.generate_typescript_content(elements)
	elements = elements or { "my-element" }

	local content = [[import { html, LitElement } from 'lit';

export class TestComponent extends LitElement {
  render() {
    return html`
]]

	for _, element in ipairs(elements) do
		content = content .. string.format("      <%s></%s>\n", element, element)
	end

	content = content .. [[    `;
  }
}]]

	return content
end

return M
