-- Attribute hover benchmark module
-- Tests LSP hover accuracy for attributes in large manifests

local measurement = require("utils.measurement")
local benchmark = require("utils.benchmark")

local M = {}

function M.run_hover_attribute_benchmark(config, fixture_dir)
	local server_name = _G.BENCHMARK_LSP_NAME or "unknown"

	-- Setup LSP client
	local client, err = benchmark.setup_client(config, fixture_dir)
	if not client then
		return {
			success = false,
			error = err,
		}
	end

	-- Setup test file (use large project for comprehensive testing)
	local bufnr, test_file = benchmark.setup_test_file(fixture_dir, "hover-attribute", "html", 3000, client)
	if not bufnr then
		benchmark.cleanup_test(nil, nil, client)
		return {
			success = false,
			error = test_file, -- test_file contains error message on failure
		}
	end

	-- Dynamically find attribute positions using tree-sitter
	-- This replaces hardcoded positions and adapts automatically to fixture changes
	local position_finder = require("utils.position_finder")
	local test_file_path = fixture_dir .. "/hover-attribute-test.html"

	-- Define attributes to test (semantic specification without hardcoded line/character)
	local attribute_specs = {
		-- Button attributes
		{ element = "my-button", attr = "variant", expected_content = "Button style variant" },
		{ element = "my-button", attr = "size", expected_content = "Button size" },
		{ element = "my-button", attr = "disabled", expected_content = "Disabled state" },
		{ element = "my-button", attr = "loading", expected_content = "Loading state" },
		{ element = "my-button", attr = "icon", expected_content = "Icon name" },

		-- Form attributes
		{ element = "my-input-text", attr = "value", expected_content = "Input value" },
		{ element = "my-input-text", attr = "placeholder", expected_content = "Placeholder text" },
		{ element = "my-input-text", attr = "required", expected_content = "Required field" },
		{ element = "my-input-select", attr = "multiple", expected_content = "Multiple selection" },
		{ element = "my-input-select", attr = "searchable", expected_content = "searchable" },

		-- Layout attributes
		{ element = "my-container-flex", attr = "gap", expected_content = "Gap size" },
		{ element = "my-container-flex", attr = "justify", expected_content = "Justification" },
		{ element = "my-card-basic", attr = "variant", expected_content = "Card style" },
		{ element = "my-card-basic", attr = "clickable", expected_content = "Clickable card" },

		-- Chart attributes
		{ element = "my-chart-line", attr = "responsive", expected_content = "responsive" },
		{ element = "my-chart-line", attr = "animated", expected_content = "animated" },

		-- Table attributes
		{ element = "my-table-sortable", attr = "striped", expected_content = "striped" },
		{ element = "my-table-sortable", attr = "page-size", expected_content = "page-size" },

		-- Media attributes
		{ element = "my-image-responsive", attr = "loading", expected_content = "Loading strategy" },
		{ element = "my-image-responsive", attr = "fit", expected_content = "fit" },
		{ element = "my-image-avatar", attr = "size", expected_content = "size" },
	}

	local test_positions = position_finder.find_attributes_batch(test_file_path, attribute_specs)

	-- Validate all attribute positions found
	local errors = {}
	for _, pos in ipairs(test_positions) do
		if pos.error then
			table.insert(errors, string.format("%s.%s: %s", pos.element, pos.attr, pos.error))
		end
	end

	if #errors > 0 then
		return {
			success = false,
			error = string.format("Position finding failed:\n  %s", table.concat(errors, "\n  ")),
		}
	end

	if #test_positions ~= #attribute_specs then
		return {
			success = false,
			error = string.format("Expected %d attribute positions, found %d", #attribute_specs, #test_positions),
		}
	end

	local hover_results = {}
	local iterations_per_attribute = 3 -- Test each attribute multiple times

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

			local content = ""
			if hover_result.contents then
				if type(hover_result.contents) == "string" then
					content = hover_result.contents
				elseif hover_result.contents.value then
					content = hover_result.contents.value
				end
			end

			-- Check if content matches expected (basic substring check)
			local content_matches = test_pos.expected_content
				and string.find(content:lower(), test_pos.expected_content:lower()) ~= nil

			return { hover_duration, string.len(content), content_matches and 1 or 0, content }
		end

		-- Run multiple iterations for this attribute
		local attr_results = measurement.run_iterations(
			single_hover,
			iterations_per_attribute,
			string.format("hover on %s.%s", test_pos.element, test_pos.attr)
		)

		-- Calculate content accuracy
		local accurate_hovers = 0
		local total_content_length = 0

		for _, result in ipairs(attr_results.results or {}) do
			if result[3] == 1 then
				accurate_hovers = accurate_hovers + 1
			end
			total_content_length = total_content_length + (result[2] or 0)
		end

		table.insert(hover_results, {
			element = test_pos.element,
			attribute = test_pos.attr,
			iterations = iterations_per_attribute,
			successful_runs = attr_results.successful_runs,
			failed_runs = attr_results.failed_runs,
			success_rate = attr_results.success_rate,
			statistics = attr_results.statistics,
			accuracy_rate = attr_results.successful_runs > 0 and accurate_hovers / attr_results.successful_runs or 0,
			avg_content_length = attr_results.successful_runs > 0
					and total_content_length / attr_results.successful_runs
				or 0,
			expected_content = test_pos.expected_content,
			sample_content = attr_results.results and #attr_results.results > 0 and attr_results.results[1][4] or "",
			errors = attr_results.errors,
		})

		-- Check if client survived
		if client:is_stopped() then
			break
		end

		-- Small delay between different attributes
		vim.wait(50)
	end

	-- Capture client state before cleanup
	local client_survived = not client:is_stopped()

	-- Clean up
	benchmark.cleanup_test(bufnr, test_file, client)

	-- Calculate overall statistics (includes accuracy tracking)
	local aggregated = benchmark.aggregate_results(hover_results)
	local total_accurate = 0

	for _, result in ipairs(hover_results) do
		total_accurate = total_accurate + (result.accuracy_rate * result.successful_runs)
	end

	return {
		success = aggregated.total_successful > 0,
		server_name = server_name,
		total_attributes_tested = #test_positions,
		total_attempts = aggregated.total_attempts,
		total_successful = aggregated.total_successful,
		success_rate = aggregated.success_rate,
		accuracy_rate = aggregated.total_successful > 0 and total_accurate / aggregated.total_successful or 0,
		overall_statistics = aggregated.overall_statistics,
		attribute_results = hover_results,
		client_survived = client_survived,
	}
end

return M
