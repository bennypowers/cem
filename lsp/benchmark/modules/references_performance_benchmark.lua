local M = {}

-- Go-to-references performance benchmark: Tests real LSP reference finding across various project sizes
-- Time limit: 45 seconds total
-- Focus: Workspace traversal, reference accuracy, and scalability
-- Implementation: Uses actual vim.lsp.buf_request_sync() calls to measure real LSP performance

function M.run_references_performance_benchmark(_, fixture_dir)
	local start_time = vim.fn.reltime()
	local max_time_seconds = 45

	local results = {
		success = false,
		server_name = _G.BENCHMARK_LSP_NAME,
		error = nil,
		reference_tests = {},
		total_references_found = 0,
		total_search_time = 0,
		average_search_time = 0,
		project_scale_results = {},
		memory_usage = 0,
	}

	local success, error_msg = pcall(function()
		-- Test 1: Small Scale References (Current Project)
		local small_start = vim.fn.reltime()

		-- Open test document and find references to common elements
		vim.cmd("edit " .. fixture_dir .. "/large-page.html")
		vim.wait(1000, function()
			return false
		end) -- Wait for LSP initialization

		local small_scale_results = test_references_for_elements({
			{ element = "my-nav", expected_min = 1, line = 21, col = 8 },
			{ element = "my-button-primary", expected_min = 2, line = 53, col = 13 },
			{ element = "my-card-product", expected_min = 1, line = 134, col = 15 },
		}, "small_scale")

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
			vim.wait(200, function()
				return false
			end) -- Brief pause for file processing
		end

		-- Set active file back to large-page for reference testing
		vim.cmd("edit " .. fixture_dir .. "/large-page.html")

		local medium_scale_results = test_references_for_elements({
			{ element = "my-nav", expected_min = 2, line = 21, col = 8 },
			{ element = "my-button-primary", expected_min = 3, line = 53, col = 13 },
			{ element = "my-divider", expected_min = 2, line = 32, col = 8 },
		}, "medium_scale")

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
		local editing_results = test_references_during_editing()

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
		vim.wait(500, function()
			return false
		end)

		local template_results = test_template_literal_references()

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
	end)

	if not success then
		results.error = error_msg
	end

	local total_time = vim.fn.reltime(start_time)
	results.total_time = tonumber(vim.fn.reltimestr(total_time))

	return results
end

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
local function test_references_during_editing()
	local results = {
		tests = {},
		total_references = 0,
		avg_search_time = 0,
		editing_interference = {},
	}

	-- Start editing document while triggering references
	local test_elements = {
		{ element = "my-nav", line = 21, col = 8 },
		{ element = "my-button-primary", line = 53, col = 13 },
	}

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
local function test_template_literal_references()
	local results = {
		tests = {},
		total_references = 0,
		avg_search_time = 0,
	}

	-- Test references within template literals (if the file has them)
	local template_elements = {
		{ element = "my-container-flex", line = 176, col = 11 }, -- Approximate position in template
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

return M
