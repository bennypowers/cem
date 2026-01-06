-- Diagnostics benchmark module
-- Tests LSP diagnostics performance and functionality

local benchmark = require("utils.benchmark")

local M = {}

function M.run_diagnostics_benchmark(config, fixture_dir)
	local server_name = _G.BENCHMARK_LSP_NAME or "unknown"
	local DIAGNOSTICS_TIMEOUT_MS = 3000

	-- Track diagnostics
	local received_diagnostics = {}
	local diagnostics_received = false

	-- Add client-specific diagnostics handler to config
	-- This avoids modifying global handlers and prevents collisions with other LSP usage
	local client_config = vim.tbl_deep_extend("force", config, {
		handlers = {
			["textDocument/publishDiagnostics"] = function(diag_err, result, ctx, handler_config)
				if result and result.diagnostics then
					received_diagnostics = result.diagnostics
					diagnostics_received = true
					-- Update diagnostics state using modern API
					-- Create namespace from client ID (vim.diagnostic.set expects namespace, not client_id)
					local namespace = vim.api.nvim_create_namespace(("lsp_%d"):format(ctx.client_id))
					vim.diagnostic.set(namespace, ctx.bufnr, result.diagnostics)
				end
			end,
		},
	})

	-- Setup LSP client with custom handlers
	local client, err = benchmark.setup_client(client_config, fixture_dir)
	if not client then
		return {
			success = false,
			error = err,
		}
	end

	local diagnostics_start = vim.uv.hrtime()

	-- Now open the file (which will trigger didOpen and diagnostics)
	local bufnr, test_file = benchmark.setup_test_file(fixture_dir, "diagnostics", "html", 2000, client)
	if not bufnr then
		benchmark.cleanup_test(nil, nil, client)
		return {
			success = false,
			error = test_file, -- test_file contains error message on failure
		}
	end

	-- Wait for diagnostics to be published (if not already received)
	local diagnostics_success = diagnostics_received
		or vim.wait(DIAGNOSTICS_TIMEOUT_MS, function()
			return diagnostics_received or client:is_stopped()
		end)

	local diagnostics_duration = (vim.uv.hrtime() - diagnostics_start) / 1e6
	local client_stopped = client:is_stopped()

	-- Check if server supports diagnostics
	local supports_diagnostics = client.server_capabilities
		and (client.server_capabilities.diagnosticProvider or client.server_capabilities.textDocumentSync)

	local result = {
		success = false, -- Will be set to true only if diagnostics work correctly
		success_rate = 0.0, -- Will be set to 1.0 if diagnostics work correctly
		server_name = server_name,
		duration_ms = diagnostics_duration,
		diagnostics_count = #received_diagnostics,
		client_survived = not client_stopped,
		supports_diagnostics = supports_diagnostics,
		not_supported = not supports_diagnostics,
	}

	if not supports_diagnostics then
		result.error = "Feature not available in server capabilities"
	elseif client_stopped then
		result.error = "Client stopped during diagnostics test"
	elseif not diagnostics_success then
		result.error = string.format("Timeout waiting for diagnostics (%gs)", DIAGNOSTICS_TIMEOUT_MS / 1000)
	else
		-- Analyze diagnostics by severity
		local error_count = 0
		local warning_count = 0
		local info_count = 0
		local hint_count = 0

		for _, diagnostic in ipairs(received_diagnostics) do
			if diagnostic.severity == vim.diagnostic.severity.ERROR then
				error_count = error_count + 1
			elseif diagnostic.severity == vim.diagnostic.severity.WARN then
				warning_count = warning_count + 1
			elseif diagnostic.severity == vim.diagnostic.severity.INFO then
				info_count = info_count + 1
			elseif diagnostic.severity == vim.diagnostic.severity.HINT then
				hint_count = hint_count + 1
			end
		end

		result.diagnostics_by_severity = {
			errors = error_count,
			warnings = warning_count,
			info = info_count,
			hints = hint_count,
		}
		result.success = true
		result.success_rate = 1.0
	end

	-- Clean up
	benchmark.cleanup_test(bufnr, test_file, client)

	return result
end

return M
