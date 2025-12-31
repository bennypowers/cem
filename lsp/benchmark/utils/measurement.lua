--[[
Measurement utilities for LSP benchmarking
Provides timing, memory usage, and async measurement functions
]]

local M = {}

-- High precision timing using libuv
function M.time_sync(func)
	local start = vim.uv.hrtime()
	local result = func()
	local duration = (vim.uv.hrtime() - start) / 1e6 -- Convert nanoseconds to milliseconds
	return duration, result
end

-- Asynchronous operation timing
function M.time_async(func, timeout_ms)
	timeout_ms = timeout_ms or 5000

	local start = vim.uv.hrtime()
	local done = false
	local result = nil
	local error_msg = nil

	func(function(err, res)
		if err then
			error_msg = err
		else
			result = res
		end
		done = true
	end)

	-- Wait for completion with timeout
	local success = vim.wait(timeout_ms, function()
		return done
	end)

	local duration = (vim.uv.hrtime() - start) / 1e6

	if not success then
		error_msg = string.format("Operation timed out after %dms", timeout_ms)
	end

	return duration, result, error_msg
end

-- Memory usage measurement
function M.memory_usage()
	collectgarbage("collect") -- Force garbage collection for accurate measurement
	return collectgarbage("count") * 1024 -- Convert KB to bytes
end

-- Run multiple iterations and collect statistics
function M.run_iterations(func, iterations, description)
	iterations = iterations or 10
	description = description or "operation"

	print(string.format("Running %d iterations of %s...", iterations, description))

	local times = {}
	local results = {}
	local errors = {}
	local start_memory = M.memory_usage()

	for i = 1, iterations do
		local result, error_msg = func()

		if error_msg then
			table.insert(errors, { iteration = i, error = error_msg })
		elseif result then
			-- Handle both single values and table results
			if type(result) == "table" then
				-- First element is duration, rest is extra data
				table.insert(times, result[1])
				table.insert(results, result)
			else
				-- Single value is duration
				table.insert(times, result)
				table.insert(results, result)
			end
		end

		-- Small delay between iterations
		vim.wait(10)
	end

	local end_memory = M.memory_usage()
	local memory_delta = end_memory - start_memory

	-- Calculate statistics
	local stats = M.calculate_statistics(times)

	return {
		iterations = iterations,
		successful_runs = #times,
		failed_runs = #errors,
		times_ms = times,
		results = results,
		errors = errors,
		statistics = stats,
		memory_usage_bytes = memory_delta,
		success_rate = #times / iterations,
	}
end

-- Calculate statistical measures from timing data
function M.calculate_statistics(times)
	if #times == 0 then
		return {
			count = 0,
			min = 0,
			max = 0,
			mean = 0,
			median = 0,
			p95 = 0,
			p99 = 0,
			stddev = 0,
		}
	end

	-- Sort times for percentile calculations
	local sorted_times = {}
	for _, time in ipairs(times) do
		table.insert(sorted_times, time)
	end
	table.sort(sorted_times)

	local count = #sorted_times
	local min = sorted_times[1]
	local max = sorted_times[count]

	-- Calculate mean
	local sum = 0
	for _, time in ipairs(sorted_times) do
		sum = sum + time
	end
	local mean = sum / count

	-- Calculate median
	local median
	if count % 2 == 0 then
		median = (sorted_times[count / 2] + sorted_times[count / 2 + 1]) / 2
	else
		median = sorted_times[math.ceil(count / 2)]
	end

	-- Calculate percentiles
	local function percentile(p)
		local index = math.ceil(count * p / 100)
		return sorted_times[math.min(index, count)]
	end

	local p95 = percentile(95)
	local p99 = percentile(99)

	-- Calculate sample standard deviation (using count - 1 for better estimate with small samples)
	local variance_sum = 0
	for _, time in ipairs(sorted_times) do
		variance_sum = variance_sum + (time - mean) ^ 2
	end
	local stddev = count > 1 and math.sqrt(variance_sum / (count - 1)) or 0

	return {
		count = count,
		min = min,
		max = max,
		mean = mean,
		median = median,
		p95 = p95,
		p99 = p99,
		stddev = stddev,
	}
end

-- Format timing statistics for display
function M.format_statistics(stats)
	if stats.count == 0 then
		return "No successful measurements"
	end

	return string.format(
		"Count: %d, Mean: %.2fms, Median: %.2fms, P95: %.2fms, P99: %.2fms, StdDev: %.2fms",
		stats.count,
		stats.mean,
		stats.median,
		stats.p95,
		stats.p99,
		stats.stddev
	)
end

-- Wait for LSP client to be ready
function M.wait_for_client_ready(client, timeout_ms)
	timeout_ms = timeout_ms or 5000

	local start = vim.uv.hrtime()

	local ready = vim.wait(timeout_ms, function()
		return client and client.server_capabilities
	end)

	local duration = (vim.uv.hrtime() - start) / 1e6

	if not ready then
		error(string.format("LSP client not ready after %dms", timeout_ms))
	end

	return duration
end

-- Create LSP request parameters for current cursor position
--- @param bufnr number|nil Buffer number (defaults to current)
--- @param line number|nil Line number (1-based). If provided with col, cursor is moved as side effect.
--- @param col number|nil Column (0-based). If provided with line, cursor is moved as side effect.
--- @return table LSP position parameters
function M.make_position_params(bufnr, line, col)
	bufnr = bufnr or vim.api.nvim_get_current_buf()
	if line and col then
		vim.api.nvim_win_set_cursor(0, { line, col })
	end
	return vim.lsp.util.make_position_params(0, "utf-8")
end

-- Create LSP range parameters
function M.make_range_params(bufnr, start_line, start_col, end_line, end_col)
	bufnr = bufnr or vim.api.nvim_get_current_buf()

	return {
		textDocument = vim.lsp.util.make_text_document_params(bufnr),
		range = {
			start = { line = start_line - 1, character = start_col }, -- LSP uses 0-based indexing
			["end"] = { line = end_line - 1, character = end_col },
		},
	}
end

return M
