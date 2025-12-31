-- Hover benchmark module
-- Tests LSP hover request performance and functionality

local measurement = require("utils.measurement")
local benchmark = require("utils.benchmark")

local M = {}

function M.run_hover_benchmark(config, fixture_dir)
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
	local bufnr, test_file = benchmark.setup_test_file(fixture_dir, "hover", "html", 2000, client)
	if not bufnr then
		benchmark.cleanup_test(nil, nil, client)
		return {
			success = false,
			error = test_file, -- test_file contains error message on failure
		}
	end

	-- Dynamically find test positions using tree-sitter
	-- This replaces hardcoded positions and adapts automatically to fixture changes
	local position_finder = require("utils.position_finder")
	local test_file_path = fixture_dir .. "/hover-test.html"

	local element_specs = {
		{ element = "my-button", occurrence = 1 },
		{ element = "my-card", occurrence = 1 },
		{ element = "my-icon", occurrence = 1 },
	}

	local positions = position_finder.find_elements_batch(test_file_path, element_specs)

	-- Validate all positions found
	local test_positions = {}
	for _, pos in ipairs(positions) do
		if pos.error then
			return {
				success = false,
				error = string.format("Position finding failed: %s", pos.error),
			}
		end
		table.insert(test_positions, pos)
	end

	if #test_positions ~= 3 then
		return {
			success = false,
			error = string.format("Expected 3 positions, found %d", #test_positions),
		}
	end

	local hover_results = {}
	local iterations_per_element = 10 -- Multiple iterations per element for statistical confidence

	for _, test_pos in ipairs(test_positions) do
		local function single_hover()
			local hover_result = nil
			local hover_error = nil
			local hover_completed = false

			local hover_start = vim.uv.hrtime()

			client:request("textDocument/hover", {
				textDocument = { uri = vim.uri_from_bufnr(bufnr) },
				position = { line = test_pos.line, character = test_pos.character },
			}, function(hover_err, result)
				hover_error = hover_err
				hover_result = result
				hover_completed = true
			end)

			local hover_success = vim.wait(3000, function()
				return hover_completed
			end)

			local hover_duration = (vim.uv.hrtime() - hover_start) / 1e6

			if not hover_success or hover_error or not hover_result then
				return nil, hover_error and vim.inspect(hover_error) or "No result"
			end

			local content_length = 0
			if hover_result.contents then
				if type(hover_result.contents) == "string" then
					content_length = string.len(hover_result.contents)
				elseif hover_result.contents.value then
					content_length = string.len(hover_result.contents.value)
				end
			end

			return { hover_duration, content_length }
		end

		-- Run multiple iterations for this element
		local element_results = measurement.run_iterations(
			single_hover,
			iterations_per_element,
			string.format("hover on %s", test_pos.element)
		)

		table.insert(hover_results, {
			element = test_pos.element,
			iterations = iterations_per_element,
			successful_runs = element_results.successful_runs,
			failed_runs = element_results.failed_runs,
			success_rate = element_results.success_rate,
			statistics = element_results.statistics,
			all_times_ms = element_results.times_ms,
			errors = element_results.errors,
			-- Average content length from successful runs
			avg_content_length = element_results.results and #element_results.results > 0 and (function()
				local total = 0
				for _, res in ipairs(element_results.results) do
					total = total + (res[2] or 0)
				end
				return total / #element_results.results
			end)() or 0,
		})

		-- Check if client survived
		if client:is_stopped() then
			break
		end

		-- No artificial delay - real-world usage is rapid successive requests
	end

	-- Capture client state before cleanup
	local client_survived = not client:is_stopped()

	-- Clean up
	benchmark.cleanup_test(bufnr, test_file, client)

	-- Calculate overall statistics
	local aggregated = benchmark.aggregate_results(hover_results)

	return {
		success = aggregated.total_successful > 0,
		server_name = server_name,
		total_elements = #test_positions,
		total_attempts = aggregated.total_attempts,
		total_successful = aggregated.total_successful,
		success_rate = aggregated.success_rate,
		overall_statistics = aggregated.overall_statistics,
		hover_results = hover_results,
		client_survived = client_survived,
		-- Backward compatibility
		successful_hovers = aggregated.total_successful,
		average_duration_ms = aggregated.overall_statistics.mean,
	}
end

return M
