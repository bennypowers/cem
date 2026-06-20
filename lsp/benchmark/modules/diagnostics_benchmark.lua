-- Diagnostics benchmark module
-- Tests LSP diagnostics performance and functionality

local benchmark = require("utils.benchmark")

local M = {}

function M.run_diagnostics_benchmark(config, fixture_dir)
	local server_name = _G.BENCHMARK_LSP_NAME or "unknown"
	local DIAGNOSTICS_TIMEOUT_MS = 3000

	local client, err = benchmark.setup_client(config, fixture_dir)
	if not client then
		return {
			success = false,
			success_rate = 0.0,
			error = err,
		}
	end

	local diagnostics_start = vim.uv.hrtime()

	local bufnr, test_file = benchmark.setup_test_file(fixture_dir, "diagnostics", "html", 2000, client)
	if not bufnr then
		benchmark.cleanup_test(nil, nil, client)
		return {
			success = false,
			success_rate = 0.0,
			error = test_file,
		}
	end

	-- Neovim 0.11+ routes publishDiagnostics through vim.diagnostic internally;
	-- custom client handlers for this method are bypassed. Poll vim.diagnostic.get()
	-- instead of waiting for a handler callback.
	local diagnostics_success = #vim.diagnostic.get(bufnr) > 0
		or vim.wait(DIAGNOSTICS_TIMEOUT_MS, function()
			return #vim.diagnostic.get(bufnr) > 0 or client:is_stopped()
		end)

	local diagnostics_duration = (vim.uv.hrtime() - diagnostics_start) / 1e6
	local client_stopped = client:is_stopped()
	local received_diagnostics = vim.diagnostic.get(bufnr)

	local supports_diagnostics = client.server_capabilities
		and (client.server_capabilities.diagnosticProvider or client.server_capabilities.textDocumentSync)

	local result = {
		success = false,
		success_rate = 0.0,
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
		local has_valid_diagnostics = false
		local error_count = 0
		local warning_count = 0
		local info_count = 0
		local hint_count = 0

		for _, diagnostic in ipairs(received_diagnostics) do
			if diagnostic.lnum ~= nil and diagnostic.message and #diagnostic.message > 0 then
				has_valid_diagnostics = true

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
		end

		result.diagnostics_by_severity = {
			errors = error_count,
			warnings = warning_count,
			info = info_count,
			hints = hint_count,
		}

		if has_valid_diagnostics then
			result.success = true
			result.success_rate = 1.0
		else
			result.error = "No valid diagnostics received (missing lnum or message)"
			result.success_rate = 0.0
		end
	end

	benchmark.cleanup_test(bufnr, test_file, client)

	return result
end

return M
