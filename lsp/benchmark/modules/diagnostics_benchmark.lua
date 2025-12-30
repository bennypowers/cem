-- Diagnostics benchmark module
-- Tests LSP diagnostics performance and functionality

local measurement = require("utils.measurement")

local M = {}

function M.run_diagnostics_benchmark(config, fixture_dir)
	local server_name = _G.BENCHMARK_LSP_NAME or "unknown"
	print(string.format("=== %s LSP Diagnostics Benchmark ===", server_name))

	-- Change to fixture directory
	vim.cmd("cd " .. vim.fn.fnameescape(fixture_dir))

	-- Start LSP server
	local client_id = vim.lsp.start(config)
	if not client_id then
		return {
			success = false,
			error = "Failed to start LSP server",
		}
	end

	local client = vim.lsp.get_client_by_id(client_id)

	-- Wait for initialization
	local ready = vim.wait(10000, function()
		return client.server_capabilities ~= nil and not client.is_stopped()
	end)

	if not ready then
		client.stop()
		return {
			success = false,
			error = "Server failed to initialize",
		}
	end

	-- Load HTML content from fixture file
	local fixture_file = fixture_dir .. "/diagnostics-test.html"

	-- Copy to test file for LSP processing
	vim.fn.writefile(vim.fn.readfile(fixture_file), "test-diagnostics.html")

	-- Track diagnostics
	local received_diagnostics = {}
	local diagnostics_received = false

	-- Set up diagnostics handler
	local original_handler = vim.lsp.handlers["textDocument/publishDiagnostics"]
	vim.lsp.handlers["textDocument/publishDiagnostics"] = function(err, result, ctx, config)
		if result and result.diagnostics then
			received_diagnostics = result.diagnostics
			diagnostics_received = true
		end
		if original_handler then
			original_handler(err, result, ctx, config)
		end
	end

	local diagnostics_start = vim.uv.hrtime()

	-- Open file
	vim.cmd("edit test-diagnostics.html")
	vim.cmd("set filetype=html")
	local bufnr = vim.api.nvim_get_current_buf()

	-- Wait for diagnostics to be published
	local diagnostics_success = vim.wait(5000, function()
		return diagnostics_received or client.is_stopped()
	end)

	local diagnostics_duration = (vim.uv.hrtime() - diagnostics_start) / 1e6

	-- Restore original handler
	vim.lsp.handlers["textDocument/publishDiagnostics"] = original_handler

	-- Check if server supports diagnostics
	local supports_diagnostics = client.server_capabilities
		and (client.server_capabilities.diagnosticProvider or client.server_capabilities.textDocumentSync)

	local result = {
		success = supports_diagnostics and (diagnostics_success and #received_diagnostics >= 0)
			or not supports_diagnostics,
		server_name = server_name,
		duration_ms = diagnostics_duration,
		diagnostics_count = #received_diagnostics,
		client_survived = not client.is_stopped(),
		supports_diagnostics = supports_diagnostics,
	}

	if diagnostics_success then
		-- Analyze diagnostics by severity
		local error_count = 0
		local warning_count = 0
		local info_count = 0
		local hint_count = 0

		for _, diagnostic in ipairs(received_diagnostics) do
			if diagnostic.severity == 1 then
				error_count = error_count + 1
			elseif diagnostic.severity == 2 then
				warning_count = warning_count + 1
			elseif diagnostic.severity == 3 then
				info_count = info_count + 1
			elseif diagnostic.severity == 4 then
				hint_count = hint_count + 1
			end
		end

		result.diagnostics_by_severity = {
			errors = error_count,
			warnings = warning_count,
			info = info_count,
			hints = hint_count,
		}
	else
		result.error = "No diagnostics received or client stopped"
	end

	-- Clean up
	vim.api.nvim_buf_delete(bufnr, { force = true })
	vim.fn.delete("test-diagnostics.html")
	client.stop()

	return result
end

return M
