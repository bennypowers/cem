local M = {}

-- Go-to-references performance benchmark: Tests real LSP reference finding across various project sizes
-- Time limit: 45 seconds total
-- Focus: Workspace traversal, reference accuracy, and scalability
-- Implementation: Uses actual vim.lsp.buf_request_sync() calls to measure real LSP performance

-- Function to test references for specific elements
local function test_references_for_elements(elements, _)
	local results = {
		tests = {},
		total_references = 0,
		avg_search_time = 0,
	}

	local total_time = 0

	for _, element_test in ipairs(elements) do
		local search_start = vim.fn.reltime()

		-- Position cursor on element
		pcall(vim.api.nvim_win_set_cursor, 0, { element_test.line, element_test.col })

		-- Make actual LSP textDocument/references request
		local references_found = 0
		local bufnr = vim.api.nvim_get_current_buf()
		local params = vim.lsp.util.make_position_params(0, "utf-8")
		params.context = { includeDeclaration = true }

		-- Request references from LSP server
		local lsp_results, err = vim.lsp.buf_request_sync(bufnr, "textDocument/references", params, 10000)

		if lsp_results then
			-- Count references from all LSP clients
			for _, client_result in pairs(lsp_results) do
				if client_result.result then
					references_found = references_found + #client_result.result
				elseif client_result.err then
					err = client_result.err
				end
			end
		end

		local search_time = tonumber(vim.fn.reltimestr(vim.fn.reltime(search_start))) * 1000 -- Convert to ms

		local test_result = {
			element = element_test.element,
			search_time_ms = search_time,
			references_found = references_found,
			expected_min = element_test.expected_min,
			success = references_found >= element_test.expected_min,
			lsp_error = err,
		}

		table.insert(results.tests, test_result)
		results.total_references = results.total_references + references_found
		total_time = total_time + search_time
	end

	results.avg_search_time = #elements > 0 and (total_time / #elements) or 0

	return results
end

-- Function to test references during active editing
-- NOTE: Uses positions_map passed from parent function
local function test_references_during_editing(positions_map)
	local results = {
		tests = {},
		total_references = 0,
		avg_search_time = 0,
		editing_interference = {},
	}

	-- Start editing document while triggering references
	-- Positions are dynamically found via tree-sitter
	local test_elements = {}
	if positions_map["my-nav"] then
		table.insert(test_elements, {
			element = "my-nav",
			line = positions_map["my-nav"].line,
			col = positions_map["my-nav"].character,
		})
	end
	if positions_map["my-button-primary"] then
		table.insert(test_elements, {
			element = "my-button-primary",
			line = positions_map["my-button-primary"].line,
			col = positions_map["my-button-primary"].character,
		})
	end

	for _, element_test in ipairs(test_elements) do
		-- Start typing to simulate active editing
		local typing_start = vim.fn.reltime()

		-- Add some content to simulate active editing
		vim.api.nvim_win_set_cursor(0, { element_test.line + 2, 0 })
		vim.api.nvim_put({ "  <!-- Added during reference test -->" }, "l", true, true)

		local typing_time = tonumber(vim.fn.reltimestr(vim.fn.reltime(typing_start)))

		-- Now trigger references while document is "dirty"
		local search_start = vim.fn.reltime()
		vim.api.nvim_win_set_cursor(0, { element_test.line, element_test.col })

		-- Make actual LSP textDocument/references request
		local references_found = 0
		local bufnr = vim.api.nvim_get_current_buf()
		local params = vim.lsp.util.make_position_params(0, "utf-8")
		params.context = { includeDeclaration = true }

		-- Request references from LSP server
		local lsp_results, err = vim.lsp.buf_request_sync(bufnr, "textDocument/references", params, 10000)

		if lsp_results then
			-- Count references from all LSP clients
			for _, client_result in pairs(lsp_results) do
				if client_result.result then
					references_found = references_found + #client_result.result
				end
			end
		end

		local search_time = tonumber(vim.fn.reltimestr(vim.fn.reltime(search_start))) * 1000 -- Convert to ms

		local test_result = {
			element = element_test.element,
			search_time_ms = search_time,
			references_found = references_found,
			editing_time_ms = typing_time,
			success = references_found >= 2,
			lsp_error = err,
		}

		table.insert(results.tests, test_result)
		table.insert(results.editing_interference, {
			element = element_test.element,
			performance_impact = search_time / 1000, -- Convert to seconds for impact assessment
		})

		results.total_references = results.total_references + references_found
	end

	local total_time = 0
	for _, test in ipairs(results.tests) do
		total_time = total_time + test.search_time_ms
	end
	results.avg_search_time = #results.tests > 0 and (total_time / #results.tests) or 0

	return results
end

-- Function to test references in TypeScript template literals
-- NOTE: Uses positions_map passed from parent function
local function test_template_literal_references(positions_map)
	local results = {
		tests = {},
		total_references = 0,
		avg_search_time = 0,
	}

	-- Test references within template literals
	-- NOTE: For TypeScript files, positions should be computed from lit-template.ts
	-- For now, using approximate hardcoded positions as tree-sitter TypeScript template support
	-- may require additional work. This is a known limitation.
	local template_elements = {
		{ element = "my-container-flex", line = 176, col = 11 },
		{ element = "my-button-primary", line = 182, col = 13 },
	}

	for _, element_test in ipairs(template_elements) do
		local search_start = vim.fn.reltime()

		pcall(vim.api.nvim_win_set_cursor, 0, { element_test.line, element_test.col })

		-- Make actual LSP textDocument/references request
		local references_found = 0
		local bufnr = vim.api.nvim_get_current_buf()
		local params = vim.lsp.util.make_position_params(0, "utf-8")
		params.context = { includeDeclaration = true }

		-- Request references from LSP server
		local lsp_results, err = vim.lsp.buf_request_sync(bufnr, "textDocument/references", params, 10000)

		if lsp_results then
			-- Count references from all LSP clients
			for _, client_result in pairs(lsp_results) do
				if client_result.result then
					references_found = references_found + #client_result.result
				end
			end
		end

		local search_time = tonumber(vim.fn.reltimestr(vim.fn.reltime(search_start))) * 1000 -- Convert to ms

		local test_result = {
			element = element_test.element,
			search_time_ms = search_time,
			references_found = references_found,
			context = "template_literal",
			success = references_found >= 1,
			lsp_error = err,
		}

		table.insert(results.tests, test_result)
		results.total_references = results.total_references + references_found
	end

	local total_time = 0
	for _, test in ipairs(results.tests) do
		total_time = total_time + test.search_time_ms
	end
	results.avg_search_time = #results.tests > 0 and (total_time / #results.tests) or 0

	return results
end

function M.run_references_benchmark(config, fixture_dir)
	local benchmark = require("utils.benchmark")
	local start_time = vim.fn.reltime()
	local max_time_seconds = 45

	-- Setup LSP client for references testing
	local client, err = benchmark.setup_client(config, fixture_dir)
	if not client then
		return {
			success = false,
			success_rate = 0.0,
			server_name = _G.BENCHMARK_LSP_NAME or "unknown",
			error = err,
		}
	end

	-- Pre-compute element positions using tree-sitter (before any LSP timing)
	local position_finder = require("utils.position_finder")
	local test_file_path = fixture_dir .. "/large-page.html"

	local element_specs = {
		{ element = "my-nav", occurrence = 1 },
		{ element = "my-button-primary", occurrence = 1 },
		{ element = "my-card-product", occurrence = 1 },
		{ element = "my-divider", occurrence = 1 },
		{ element = "my-container-flex", occurrence = 1 },
	}

	local positions_map = {}
	for _, spec in ipairs(element_specs) do
		local pos, err = position_finder.find_element_position(test_file_path, spec.element, spec.occurrence)
		if pos then
			positions_map[spec.element] = pos
		else
			-- Log warning but continue (some elements may not exist in fixture)
			print(string.format("Warning: %s", err))
		end
	end

	local results = {
		success = false,
		success_rate = 0.0,
		server_name = _G.BENCHMARK_LSP_NAME or "unknown",
		error = nil,
		reference_tests = {},
		total_references_found = 0,
		total_search_time = 0,
		average_search_time = 0,
		project_scale_results = {},
		memory_usage = 0,
		references_supported = nil,
		not_supported = false,
	}

	local success, error_msg = pcall(function()
		-- Test 1: Small Scale References (Current Project)
		local small_start = vim.fn.reltime()

		-- Open test document and find references to common elements
		vim.cmd("edit " .. fixture_dir .. "/large-page.html")
		local bufnr = vim.api.nvim_get_current_buf()
		vim.bo[bufnr].filetype = "html"

		-- Wait for LSP client to attach to buffer
		vim.wait(1000, function()
			local clients = vim.lsp.get_clients({ bufnr = bufnr })
			return #clients > 0
		end, 100)

		-- Wait longer for workspace indexing (critical for references to work)
		vim.wait(2000, function()
			return false
		end)

		-- Check if client supports references
		local supports_references = client.server_capabilities
			and client.server_capabilities.referencesProvider
		results.references_supported = supports_references
		results.not_supported = not supports_references

		-- Build small scale test list using dynamic positions
		local small_scale_tests = {}
		if positions_map["my-nav"] then
			table.insert(small_scale_tests, {
				element = "my-nav",
				expected_min = 1,
				line = positions_map["my-nav"].line,
				col = positions_map["my-nav"].character,
			})
		end
		if positions_map["my-button-primary"] then
			table.insert(small_scale_tests, {
				element = "my-button-primary",
				expected_min = 2,
				line = positions_map["my-button-primary"].line,
				col = positions_map["my-button-primary"].character,
			})
		end
		if positions_map["my-card-product"] then
			table.insert(small_scale_tests, {
				element = "my-card-product",
				expected_min = 1,
				line = positions_map["my-card-product"].line,
				col = positions_map["my-card-product"].character,
			})
		end

		local small_scale_results = test_references_for_elements(small_scale_tests, "small_scale")

		results.project_scale_results.small = {
			duration = tonumber(vim.fn.reltimestr(vim.fn.reltime(small_start))),
			tests = small_scale_results.tests,
			total_references = small_scale_results.total_references,
			avg_search_time = small_scale_results.avg_search_time,
		}

		-- Check time limit
		if vim.fn.reltime(start_time)[1] >= max_time_seconds then
			return
		end

		-- Test 2: Medium Scale References (Multiple Files)
		local medium_start = vim.fn.reltime()

		-- Open multiple files to create larger reference scope
		local files_to_open = {
			fixture_dir .. "/large-page.html",
			fixture_dir .. "/medium-page.html",
			fixture_dir .. "/complex-page.html",
			fixture_dir .. "/nested-elements.html",
			fixture_dir .. "/small-page.html",
		}

		for _, file in ipairs(files_to_open) do
			vim.cmd("edit " .. file)
			local file_bufnr = vim.api.nvim_get_current_buf()
			vim.bo[file_bufnr].filetype = "html"
			vim.wait(200, function()
				return false
			end) -- Brief pause for file processing
		end

		-- Set active file back to large-page for reference testing
		vim.cmd("edit " .. fixture_dir .. "/large-page.html")
		local active_bufnr = vim.api.nvim_get_current_buf()
		vim.bo[active_bufnr].filetype = "html"

		-- Build medium scale test list using dynamic positions
		local medium_scale_tests = {}
		if positions_map["my-nav"] then
			table.insert(medium_scale_tests, {
				element = "my-nav",
				expected_min = 2,
				line = positions_map["my-nav"].line,
				col = positions_map["my-nav"].character,
			})
		end
		if positions_map["my-button-primary"] then
			table.insert(medium_scale_tests, {
				element = "my-button-primary",
				expected_min = 3,
				line = positions_map["my-button-primary"].line,
				col = positions_map["my-button-primary"].character,
			})
		end
		if positions_map["my-divider"] then
			table.insert(medium_scale_tests, {
				element = "my-divider",
				expected_min = 2,
				line = positions_map["my-divider"].line,
				col = positions_map["my-divider"].character,
			})
		end

		local medium_scale_results = test_references_for_elements(medium_scale_tests, "medium_scale")

		results.project_scale_results.medium = {
			duration = tonumber(vim.fn.reltimestr(vim.fn.reltime(medium_start))),
			tests = medium_scale_results.tests,
			total_references = medium_scale_results.total_references,
			avg_search_time = medium_scale_results.avg_search_time,
		}

		-- Check time limit
		if vim.fn.reltime(start_time)[1] >= max_time_seconds then
			return
		end

		-- Test 3: Reference Performance During Active Editing
		local editing_start = vim.fn.reltime()

		-- Create a scenario where we trigger references while document is being modified
		local editing_results = test_references_during_editing(positions_map)

		results.project_scale_results.during_editing = {
			duration = tonumber(vim.fn.reltimestr(vim.fn.reltime(editing_start))),
			tests = editing_results.tests,
			total_references = editing_results.total_references,
			avg_search_time = editing_results.avg_search_time,
			editing_interference = editing_results.editing_interference,
		}

		-- Check time limit
		if vim.fn.reltime(start_time)[1] >= max_time_seconds then
			return
		end

		-- Test 4: Template Literal References (TypeScript)
		local template_start = vim.fn.reltime()

		-- Test references within TypeScript template literals
		vim.cmd("edit " .. fixture_dir .. "/lit-template.ts")
		local ts_bufnr = vim.api.nvim_get_current_buf()
		vim.bo[ts_bufnr].filetype = "typescript"
		vim.wait(500, function()
			return false
		end)

		local template_results = test_template_literal_references(positions_map)

		results.project_scale_results.template_literals = {
			duration = tonumber(vim.fn.reltimestr(vim.fn.reltime(template_start))),
			tests = template_results.tests,
			total_references = template_results.total_references,
			avg_search_time = template_results.avg_search_time,
		}

		-- Calculate overall statistics
		local total_tests = 0
		local total_search_time = 0
		local total_references = 0

		for _, scale_result in pairs(results.project_scale_results) do
			total_tests = total_tests + #scale_result.tests
			total_search_time = total_search_time + (scale_result.avg_search_time * #scale_result.tests)
			total_references = total_references + scale_result.total_references
		end

		results.total_references_found = total_references
		results.total_search_time = total_search_time
		results.average_search_time = total_tests > 0 and (total_search_time / total_tests) or 0

		-- Memory usage (cross-platform using Lua garbage collector)
		collectgarbage("collect")
		results.memory_usage = collectgarbage("count") * 1024 -- Convert KB to bytes

		-- Success requires both completing tests AND finding references
		results.success = total_references >= 10 and total_tests >= 8
		results.success_rate = results.success and 1.0 or 0.0

		-- Set error message if criteria not met
		if not results.success then
			if results.not_supported then
				results.error = "Feature not available in server capabilities"
			else
				local reasons = {}
				if total_references < 10 then
					-- Check if there were any LSP errors
					local had_errors = false
					for _, scale in pairs(results.project_scale_results) do
						if scale.tests then
							for _, test in ipairs(scale.tests) do
								if test.lsp_error then
									had_errors = true
									break
								end
							end
						end
					end

					if had_errors then
						table.insert(
							reasons,
							string.format(
								"insufficient references found (%d < 10) - LSP errors occurred",
								total_references
							)
						)
					else
						table.insert(
							reasons,
							string.format(
								"insufficient references found (%d < 10) - server returned empty results",
								total_references
							)
						)
					end
				end
				if total_tests < 8 then
					table.insert(reasons, string.format("insufficient tests completed (%d < 8)", total_tests))
				end
				results.error = table.concat(reasons, "; ")
			end
		end
	end)

	if not success then
		results.error = error_msg
		results.success = false
		results.success_rate = 0.0
	end

	-- Clean up LSP client
	if client and not client:is_stopped() then
		client:stop()
	end

	local total_time = vim.fn.reltime(start_time)
	results.total_time = tonumber(vim.fn.reltimestr(total_time))

	return results
end

return M
