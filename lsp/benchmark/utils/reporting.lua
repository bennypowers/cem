--[[
Reporting utilities for LSP benchmark results
Generates comparison reports between cem and wc-toolkit LSP servers
]]

local M = {}

-- Generate a comprehensive comparison report
function M.generate_comparison(results)
	print("Generating benchmark comparison report...")

	local cem_results = results.cem
	local wc_results = results.wc_toolkit

	-- Create results directory
	local results_dir = "results"
	vim.fn.mkdir(results_dir, "p")

	-- Save individual results as JSON
	M.save_json_results(cem_results, results_dir .. "/cem_results.json")
	M.save_json_results(wc_results, results_dir .. "/wc_toolkit_results.json")

	-- Generate markdown comparison report
	local markdown_report = M.generate_markdown_comparison(cem_results, wc_results)
	M.save_text_file(markdown_report, results_dir .. "/comparison_report.md")

	-- Generate CSV data for external analysis
	local csv_data = M.generate_csv_comparison(cem_results, wc_results)
	M.save_text_file(csv_data, results_dir .. "/benchmark_data.csv")

	-- Generate summary table for CI
	local summary = M.generate_summary_table(cem_results, wc_results)
	M.save_text_file(summary, results_dir .. "/summary.txt")

	print("Reports generated in results/ directory:")
	print("  - comparison_report.md (detailed comparison)")
	print("  - benchmark_data.csv (raw data)")
	print("  - summary.txt (CI-friendly summary)")
	print("  - cem_results.json (cem detailed results)")
	print("  - wc_toolkit_results.json (wc-toolkit detailed results)")
end

-- Save results as JSON
function M.save_json_results(results, filepath)
	local json_str = vim.fn.json_encode(results)
	M.save_text_file(json_str, filepath)
end

-- Save text content to file
function M.save_text_file(content, filepath)
	local file = io.open(filepath, "w")
	if file then
		file:write(content)
		file:close()
	else
		error("Failed to write file: " .. filepath)
	end
end

-- Generate markdown comparison report
function M.generate_markdown_comparison(cem_results, wc_results)
	local report = {}

	table.insert(report, "# LSP Performance Comparison: cem vs wc-toolkit")
	table.insert(report, "")
	table.insert(report, string.format("**Report Generated:** %s", os.date("%Y-%m-%d %H:%M:%S")))
	table.insert(report, "")

	-- Executive Summary
	table.insert(report, "## Executive Summary")
	table.insert(report, "")

	local summary_table = M.create_summary_comparison_table(cem_results, wc_results)
	for _, line in ipairs(summary_table) do
		table.insert(report, line)
	end

	table.insert(report, "")

	-- Detailed Results by Scenario
	table.insert(report, "## Detailed Results by Scenario")
	table.insert(report, "")

	local scenarios = { "startup", "completion", "hover", "diagnostics", "code-actions", "add-files", "file-updates" }

	for _, scenario in ipairs(scenarios) do
		local cem_scenario = cem_results.scenarios[scenario]
		local wc_scenario = wc_results.scenarios[scenario]

		if cem_scenario or wc_scenario then
			table.insert(report, string.format("### %s", scenario:gsub("-", " "):gsub("^%l", string.upper)))
			table.insert(report, "")

			local scenario_comparison = M.create_scenario_comparison(cem_scenario, wc_scenario, scenario)
			for _, line in ipairs(scenario_comparison) do
				table.insert(report, line)
			end

			table.insert(report, "")
		end
	end

	-- Performance Analysis
	table.insert(report, "## Performance Analysis")
	table.insert(report, "")
	table.insert(report, M.generate_performance_analysis(cem_results, wc_results))
	table.insert(report, "")

	-- Recommendations
	table.insert(report, "## Recommendations")
	table.insert(report, "")
	table.insert(report, M.generate_recommendations(cem_results, wc_results))
	table.insert(report, "")

	-- Technical Details
	table.insert(report, "## Technical Details")
	table.insert(report, "")
	table.insert(report, "### Server Capabilities")
	table.insert(report, "")
	table.insert(report, M.generate_capabilities_comparison(cem_results, wc_results))

	return table.concat(report, "\n")
end

-- Create summary comparison table
function M.create_summary_comparison_table(cem_results, wc_results)
	local table_lines = {}

	table.insert(table_lines, "| Scenario | cem LSP | wc-toolkit LSP | Winner | Performance Ratio |")
	table.insert(table_lines, "|----------|---------|----------------|--------|-------------------|")

	local scenarios = {
		{ key = "startup", name = "Startup", field = "overall_startup_score_ms" },
		{ key = "completion", name = "Completion", field = "overall_completion_score_ms" },
		{ key = "hover", name = "Hover", field = "overall_hover_score_ms" },
		{ key = "diagnostics", name = "Diagnostics", field = "overall_diagnostics_score_ms" },
		{ key = "code-actions", name = "Code Actions", field = "overall_code_actions_score_ms" },
		{ key = "add-files", name = "Add Files", field = "overall_add_files_score_ms" },
		{ key = "file-updates", name = "File Updates", field = "overall_file_updates_score_ms" },
	}

	for _, scenario in ipairs(scenarios) do
		local cem_scenario = cem_results.scenarios[scenario.key]
		local wc_scenario = wc_results.scenarios[scenario.key]

		if cem_scenario and wc_scenario then
			local cem_time = cem_scenario[scenario.field] or 0
			local wc_time = wc_scenario[scenario.field] or 0

			local winner = "Tie"
			local ratio = "1.0x"

			if cem_time > 0 and wc_time > 0 then
				if cem_time < wc_time then
					winner = "**cem**"
					ratio = string.format("%.1fx faster", wc_time / cem_time)
				elseif wc_time < cem_time then
					winner = "**wc-toolkit**"
					ratio = string.format("%.1fx faster", cem_time / wc_time)
				end
			elseif cem_time == 0 and wc_time > 0 then
				winner = "**cem**"
				ratio = "N/A"
			elseif wc_time == 0 and cem_time > 0 then
				winner = "**wc-toolkit**"
				ratio = "N/A"
			end

			table.insert(
				table_lines,
				string.format("| %s | %.2fms | %.2fms | %s | %s |", scenario.name, cem_time, wc_time, winner, ratio)
			)
		end
	end

	return table_lines
end

-- Create detailed scenario comparison
function M.create_scenario_comparison(cem_scenario, wc_scenario, scenario_name)
	local lines = {}

	if not cem_scenario and not wc_scenario then
		table.insert(lines, "No data available for this scenario.")
		return lines
	end

	-- Performance metrics
	table.insert(lines, "| Metric | cem LSP | wc-toolkit LSP | Better |")
	table.insert(lines, "|--------|---------|----------------|--------|")

	local metrics = {
		{ field = "overall_" .. scenario_name:gsub("-", "_") .. "_score_ms", name = "Overall Score", unit = "ms" },
		{ field = "memory_usage_bytes", name = "Memory Usage", unit = "bytes" },
		{ field = "average_completion_items", name = "Avg Completion Items", unit = "" },
		{ field = "content_success_rate", name = "Content Success Rate", unit = "%" },
		{ field = "average_diagnostic_count", name = "Avg Diagnostic Count", unit = "" },
		{ field = "average_action_count", name = "Avg Action Count", unit = "" },
	}

	for _, metric in ipairs(metrics) do
		local cem_value = cem_scenario and cem_scenario[metric.field]
		local wc_value = wc_scenario and wc_scenario[metric.field]

		if cem_value or wc_value then
			local cem_display = cem_value and M.format_metric_value(cem_value, metric.unit) or "N/A"
			local wc_display = wc_value and M.format_metric_value(wc_value, metric.unit) or "N/A"

			local better = "Tie"
			if cem_value and wc_value then
				if metric.field:match("score_ms") or metric.field == "memory_usage_bytes" then
					-- Lower is better
					better = cem_value < wc_value and "cem" or (wc_value < cem_value and "wc-toolkit" or "Tie")
				else
					-- Higher is better
					better = cem_value > wc_value and "cem" or (wc_value > cem_value and "wc-toolkit" or "Tie")
				end
			end

			table.insert(lines, string.format("| %s | %s | %s | %s |", metric.name, cem_display, wc_display, better))
		end
	end

	-- Test breakdown
	if cem_scenario and cem_scenario.tests then
		table.insert(lines, "")
		table.insert(lines, "#### Test Breakdown")
		table.insert(lines, "")
		table.insert(
			lines,
			"| Test | cem Mean (ms) | cem Success Rate | wc-toolkit Mean (ms) | wc-toolkit Success Rate |"
		)
		table.insert(
			lines,
			"|------|---------------|------------------|----------------------|-------------------------|"
		)

		for test_name, test_result in pairs(cem_scenario.tests) do
			local cem_mean = test_result.statistics and test_result.statistics.mean or 0
			local cem_success = test_result.success_rate or 1

			local wc_test = wc_scenario and wc_scenario.tests and wc_scenario.tests[test_name]
			local wc_mean = wc_test and wc_test.statistics and wc_test.statistics.mean or 0
			local wc_success = wc_test and wc_test.success_rate or (wc_test and 1 or 0)

			table.insert(
				lines,
				string.format(
					"| %s | %.2f | %.1f%% | %.2f | %.1f%% |",
					test_name:gsub("_", " "),
					cem_mean,
					cem_success * 100,
					wc_mean,
					wc_success * 100
				)
			)
		end
	end

	return lines
end

-- Format metric values with appropriate units
function M.format_metric_value(value, unit)
	if unit == "ms" then
		return string.format("%.2f%s", value, unit)
	elseif unit == "bytes" then
		if value > 1024 * 1024 then
			return string.format("%.1fMB", value / (1024 * 1024))
		elseif value > 1024 then
			return string.format("%.1fKB", value / 1024)
		else
			return string.format("%d%s", value, unit)
		end
	elseif unit == "%" then
		return string.format("%.1f%%", value * 100)
	else
		return string.format("%.1f%s", value, unit)
	end
end

-- Generate performance analysis
function M.generate_performance_analysis(cem_results, wc_results)
	local analysis = {}

	table.insert(analysis, "### Key Findings")
	table.insert(analysis, "")

	-- Compare overall performance
	local cem_total = M.calculate_total_score(cem_results)
	local wc_total = M.calculate_total_score(wc_results)

	if cem_total < wc_total then
		table.insert(
			analysis,
			string.format(
				"- **cem LSP** shows overall better performance with %.2fms average response time vs %.2fms for wc-toolkit (%.1fx faster)",
				cem_total,
				wc_total,
				wc_total / cem_total
			)
		)
	elseif wc_total < cem_total then
		table.insert(
			analysis,
			string.format(
				"- **wc-toolkit LSP** shows overall better performance with %.2fms average response time vs %.2fms for cem (%.1fx faster)",
				wc_total,
				cem_total,
				cem_total / wc_total
			)
		)
	else
		table.insert(analysis, "- Both LSP servers show similar overall performance")
	end

	table.insert(analysis, "")

	-- Memory usage comparison
	local cem_memory = M.calculate_average_memory(cem_results)
	local wc_memory = M.calculate_average_memory(wc_results)

	if cem_memory and wc_memory then
		table.insert(
			analysis,
			string.format(
				"- **Memory Usage**: cem (%.1fMB) vs wc-toolkit (%.1fMB)",
				cem_memory / (1024 * 1024),
				wc_memory / (1024 * 1024)
			)
		)
	end

	table.insert(analysis, "")

	return table.concat(analysis, "\n")
end

-- Generate recommendations
function M.generate_recommendations(cem_results, wc_results)
	local recommendations = {}

	table.insert(recommendations, "### Use Case Recommendations")
	table.insert(recommendations, "")

	-- Analyze strengths of each server
	local cem_strengths = {}
	local wc_strengths = {}

	local scenarios = { "startup", "completion", "hover", "diagnostics", "code-actions", "add-files", "file-updates" }
	for _, scenario in ipairs(scenarios) do
		local cem_scenario = cem_results.scenarios[scenario]
		local wc_scenario = wc_results.scenarios[scenario]

		if cem_scenario and wc_scenario then
			local field = "overall_" .. scenario:gsub("-", "_") .. "_score_ms"
			local cem_time = cem_scenario[field] or 999999
			local wc_time = wc_scenario[field] or 999999

			if cem_time < wc_time then
				table.insert(cem_strengths, (scenario:gsub("-", " ")))
			elseif wc_time < cem_time then
				table.insert(wc_strengths, (scenario:gsub("-", " ")))
			end
		end
	end

	if #cem_strengths > 0 then
		table.insert(recommendations, "**Choose cem LSP for:**")
		for _, strength in ipairs(cem_strengths) do
			table.insert(recommendations, string.format("- %s performance", strength))
		end
		table.insert(recommendations, "")
	end

	if #wc_strengths > 0 then
		table.insert(recommendations, "**Choose wc-toolkit LSP for:**")
		for _, strength in ipairs(wc_strengths) do
			table.insert(recommendations, string.format("- %s performance", strength))
		end
		table.insert(recommendations, "")
	end

	return table.concat(recommendations, "\n")
end

-- Generate capabilities comparison
function M.generate_capabilities_comparison(cem_results, wc_results)
	local comparison = {}

	table.insert(comparison, "| Capability | cem LSP | wc-toolkit LSP |")
	table.insert(comparison, "|------------|---------|----------------|")

	-- Get supported methods from startup scenario
	local cem_startup = cem_results.scenarios.startup
	local wc_startup = wc_results.scenarios.startup

	if cem_startup and cem_startup.supported_methods and wc_startup and wc_startup.supported_methods then
		local all_methods = {}

		-- Collect all methods from both servers
		for method, _ in pairs(cem_startup.supported_methods) do
			all_methods[method] = true
		end
		for method, _ in pairs(wc_startup.supported_methods) do
			all_methods[method] = true
		end

		-- Compare support
		for method, _ in pairs(all_methods) do
			local cem_support = cem_startup.supported_methods[method] and "✅" or "❌"
			local wc_support = wc_startup.supported_methods[method] and "✅" or "❌"

			table.insert(comparison, string.format("| %s | %s | %s |", method, cem_support, wc_support))
		end
	end

	return table.concat(comparison, "\n")
end

-- Calculate total performance score
function M.calculate_total_score(results)
	local total = 0
	local count = 0

	for _, scenario in pairs(results.scenarios) do
		for field, value in pairs(scenario) do
			if field:match("overall_.*_score_ms") and type(value) == "number" and value > 0 then
				total = total + value
				count = count + 1
			end
		end
	end

	return count > 0 and (total / count) or 0
end

-- Calculate average memory usage
function M.calculate_average_memory(results)
	local total = 0
	local count = 0

	for _, scenario in pairs(results.scenarios) do
		if scenario.memory_usage_bytes and scenario.memory_usage_bytes > 0 then
			total = total + scenario.memory_usage_bytes
			count = count + 1
		end
	end

	return count > 0 and (total / count) or nil
end

-- Generate CSV data for external analysis
function M.generate_csv_comparison(cem_results, wc_results)
	local csv_lines = {}

	table.insert(csv_lines, "scenario,metric,cem_value,wc_toolkit_value,unit,winner")

	local scenarios = { "startup", "completion", "hover", "diagnostics", "code-actions", "add-files", "file-updates" }

	for _, scenario in ipairs(scenarios) do
		local cem_scenario = cem_results.scenarios[scenario:gsub("_", "-")]
		local wc_scenario = wc_results.scenarios[scenario:gsub("_", "-")]

		if cem_scenario and wc_scenario then
			local metrics = {
				{ field = "overall_" .. scenario:gsub("-", "_") .. "_score_ms", unit = "ms" },
				{ field = "memory_usage_bytes", unit = "bytes" },
			}

			for _, metric in ipairs(metrics) do
				local cem_value = cem_scenario[metric.field] or ""
				local wc_value = wc_scenario[metric.field] or ""

				local winner = ""
				if cem_value ~= "" and wc_value ~= "" then
					if cem_value < wc_value then
						winner = "cem"
					elseif wc_value < cem_value then
						winner = "wc_toolkit"
					else
						winner = "tie"
					end
				end

				table.insert(
					csv_lines,
					string.format("%s,%s,%s,%s,%s,%s", scenario, metric.field, cem_value, wc_value, metric.unit, winner)
				)
			end
		end
	end

	return table.concat(csv_lines, "\n")
end

-- Generate CI-friendly summary
function M.generate_summary_table(cem_results, wc_results)
	local summary = {}

	table.insert(summary, "LSP BENCHMARK SUMMARY")
	table.insert(summary, "===================")
	table.insert(summary, "")

	local cem_wins = 0
	local wc_wins = 0
	local ties = 0

	local scenarios = { "startup", "completion", "hover", "diagnostics", "code-actions", "add-files", "file-updates" }

	for _, scenario in ipairs(scenarios) do
		local cem_scenario = cem_results.scenarios[scenario:gsub("_", "-")]
		local wc_scenario = wc_results.scenarios[scenario:gsub("_", "-")]

		if cem_scenario and wc_scenario then
			local field = "overall_" .. scenario:gsub("-", "_") .. "_score_ms"
			local cem_time = cem_scenario[field] or 999999
			local wc_time = wc_scenario[field] or 999999

			if cem_time < wc_time then
				cem_wins = cem_wins + 1
				table.insert(summary, string.format("%s: cem wins (%.2fms vs %.2fms)", scenario, cem_time, wc_time))
			elseif wc_time < cem_time then
				wc_wins = wc_wins + 1
				table.insert(
					summary,
					string.format("%s: wc-toolkit wins (%.2fms vs %.2fms)", scenario, wc_time, cem_time)
				)
			else
				ties = ties + 1
				table.insert(summary, string.format("%s: tie (%.2fms)", scenario, cem_time))
			end
		end
	end

	table.insert(summary, "")
	table.insert(summary, "FINAL SCORE:")
	table.insert(summary, string.format("cem: %d wins", cem_wins))
	table.insert(summary, string.format("wc-toolkit: %d wins", wc_wins))
	table.insert(summary, string.format("ties: %d", ties))
	table.insert(summary, "")

	if cem_wins > wc_wins then
		table.insert(summary, "OVERALL WINNER: cem LSP")
	elseif wc_wins > cem_wins then
		table.insert(summary, "OVERALL WINNER: wc-toolkit LSP")
	else
		table.insert(summary, "RESULT: Performance tie")
	end

	return table.concat(summary, "\n")
end

return M
