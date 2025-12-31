-- Completion benchmark module
-- Tests LSP completion request performance and functionality

local measurement = require("utils.measurement")
local benchmark = require("utils.benchmark")

local M = {}

function M.run_completion_benchmark(config, fixture_dir)
	local server_name = _G.BENCHMARK_LSP_NAME or "unknown"

	-- Setup LSP client
	local client, err = benchmark.setup_client(config, fixture_dir)
	if not client then
		return {
			success = false,
			error = err,
		}
	end

	-- Setup test file
	local bufnr, test_file = benchmark.setup_test_file(fixture_dir, "completion", "html", 2000, client)
	if not bufnr then
		benchmark.cleanup_test(nil, nil, client)
		return {
			success = false,
			error = test_file, -- test_file contains error message on failure
		}
	end

	-- Test completion at various positions with statistical analysis
	-- NOTE: These positions are hardcoded for fixtures/medium_project/completion-test.html
	-- If the fixture file is modified, these positions must be updated accordingly
	local test_positions = {
		{ context = "tag-name-completion", line = 4, character = 6 }, -- After "my-"
		{ context = "attribute-completion", line = 5, character = 13 }, -- Inside my-button tag
		{ context = "nested-completion", line = 6, character = 14 }, -- Inside my-card tag
	}

	local completion_results = {}
	-- 15 iterations per context (vs 10 for hover) to capture completion's higher variance
	-- due to manifest traversal and attribute enumeration complexity
	local iterations_per_context = 15

	for _, test_pos in ipairs(test_positions) do
		local function single_completion()
			local comp_result = nil
			local comp_error = nil
			local comp_completed = false

			local comp_start = vim.uv.hrtime()

			client:request("textDocument/completion", {
				textDocument = { uri = vim.uri_from_bufnr(bufnr) },
				position = { line = test_pos.line, character = test_pos.character },
			}, function(comp_err, result)
				comp_error = comp_err
				comp_result = result
				comp_completed = true
			end)

			local comp_success = vim.wait(3000, function()
				return comp_completed
			end)

			local comp_duration = (vim.uv.hrtime() - comp_start) / 1e6

			if not comp_success or comp_error or not comp_result then
				return nil, comp_error and vim.inspect(comp_error) or "No result"
			end

			local item_count = 0
			if comp_result.items then
				item_count = #comp_result.items
			elseif type(comp_result) == "table" then
				item_count = #comp_result
			end

			return { comp_duration, item_count }
		end

		-- Run multiple iterations for this context
		local context_results = measurement.run_iterations(
			single_completion,
			iterations_per_context,
			string.format("completion in %s", test_pos.context)
		)

		table.insert(completion_results, {
			context = test_pos.context,
			iterations = iterations_per_context,
			successful_runs = context_results.successful_runs,
			failed_runs = context_results.failed_runs,
			success_rate = context_results.success_rate,
			statistics = context_results.statistics,
			all_times_ms = context_results.times_ms,
			errors = context_results.errors,
			-- Average item count from successful runs
			avg_item_count = context_results.results and #context_results.results > 0 and (function()
				local total = 0
				for _, res in ipairs(context_results.results) do
					total = total + (res[2] or 0)
				end
				return total / #context_results.results
			end)() or 0,
		})

		-- Check if client survived
		if client:is_stopped() then
			break
		end

		-- No artificial delay - real-world usage is rapid successive requests
	end

	-- Clean up
	benchmark.cleanup_test(bufnr, test_file, client)

	-- Calculate overall statistics
	local aggregated = benchmark.aggregate_results(completion_results)

	return {
		success = aggregated.total_successful > 0,
		server_name = server_name,
		total_contexts = #test_positions,
		total_attempts = aggregated.total_attempts,
		total_successful = aggregated.total_successful,
		success_rate = aggregated.success_rate,
		overall_statistics = aggregated.overall_statistics,
		completion_results = completion_results,
		client_survived = not client:is_stopped(),
		-- Backward compatibility
		successful_completions = aggregated.total_successful,
		average_duration_ms = aggregated.overall_statistics.mean,
	}
end

return M
