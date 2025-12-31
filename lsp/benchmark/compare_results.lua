#!/usr/bin/env nvim --headless --clean -l
-- Compare benchmark results from multiple servers

local function read_json_file(filepath)
	local file = io.open(filepath, "r")
	if not file then
		return nil
	end
	local content = file:read("*all")
	file:close()
	return vim.fn.json_decode(content)
end

local function format_percentage(rate)
	if rate == nil then
		return "N/A"
	end
	return string.format("%.0f%%", rate * 100)
end

local function format_duration(ms)
	if ms == nil then
		return "N/A"
	end
	return string.format("%.2fms", ms)
end

local function calculate_attribute_hover_timing(result)
	-- Calculate average timing from individual attribute results
	if not result.attribute_results or #result.attribute_results == 0 then
		return nil
	end

	local total_time = 0
	local count = 0
	for _, attr_result in ipairs(result.attribute_results) do
		if attr_result.statistics and attr_result.statistics.mean and attr_result.statistics.mean > 0 then
			total_time = total_time + attr_result.statistics.mean
			count = count + 1
		end
	end

	if count == 0 then
		return nil
	end

	return total_time / count
end

local function calculate_references_success_rate(result)
	-- Calculate success rate from project_scale_results tests
	if not result.project_scale_results then
		return result.success and 1 or 0
	end

	local total_tests = 0
	local passed_tests = 0

	for _, scale_result in pairs(result.project_scale_results) do
		if scale_result.tests then
			for _, test in ipairs(scale_result.tests) do
				total_tests = total_tests + 1
				if test.success then
					passed_tests = passed_tests + 1
				end
			end
		end
	end

	if total_tests == 0 then
		return result.success and 1 or 0
	end

	return passed_tests / total_tests
end

local function get_benchmark_value(benchmark_data, benchmark_name)
	if not benchmark_data then
		return nil, nil, false
	end

	local result = benchmark_data[benchmark_name]
	if not result then
		return nil, nil, false
	end

	-- Extract timing and success rate for all benchmark types
	local timing_ms = nil
	local success_rate = nil

	if benchmark_name == "startup" then
		timing_ms = result.duration_ms or (result.statistics and result.statistics.mean)
		success_rate = result.success_rate
	elseif benchmark_name == "diagnostics" then
		timing_ms = result.duration_ms
		success_rate = result.success and 1 or 0
	elseif benchmark_name == "hover" or benchmark_name == "completion" then
		timing_ms = result.overall_statistics and result.overall_statistics.mean
		success_rate = result.success_rate
	elseif benchmark_name == "hover_attribute" then
		timing_ms = calculate_attribute_hover_timing(result)
		success_rate = result.success_rate
	elseif benchmark_name == "references" then
		timing_ms = result.average_search_time
		success_rate = calculate_references_success_rate(result)
	end

	-- Format as "Xms (Y%)"
	local value = nil
	if timing_ms and success_rate then
		value = string.format("%s (%s)", format_duration(timing_ms), format_percentage(success_rate))
	elseif timing_ms then
		value = format_duration(timing_ms)
	elseif success_rate then
		value = format_percentage(success_rate)
	end

	local success = result.success or false
	local not_supported = result.not_supported or false

	return value, success, not_supported
end

local function main()
	local cem_file = "/tmp/cem-benchmark-cem.json"
	local wc_file = "/tmp/cem-benchmark-wc-toolkit.json"

	local cem_data = read_json_file(cem_file)
	local wc_data = read_json_file(wc_file)

	if not cem_data and not wc_data then
		print("\nNo benchmark results found.")
		print("Expected files:")
		print("  " .. cem_file)
		print("  " .. wc_file)
		return
	end

	print("\n" .. string.rep("=", 70))
	print("LSP BENCHMARK COMPARISON")
	print(string.rep("=", 70))

	-- Benchmark order
	local benchmarks = {
		{ name = "startup", display = "Startup" },
		{ name = "hover", display = "Hover Tag Name" },
		{ name = "hover_attribute", display = "Hover Attribute" },
		{ name = "completion", display = "Completion" },
		{ name = "diagnostics", display = "Diagnostics" },
		{ name = "references", display = "References" },
	}

	-- Print table header
	print(string.format("%-20s | %-20s | %-20s", "Benchmark", "CEM", "WC-Toolkit"))
	print(string.rep("-", 70))

	-- Print each benchmark
	for _, bench in ipairs(benchmarks) do
		local cem_val, cem_success, cem_not_supported =
			get_benchmark_value(cem_data and cem_data.benchmarks, bench.name)
		local wc_val, wc_success, wc_not_supported = get_benchmark_value(wc_data and wc_data.benchmarks, bench.name)

		local function format_cell(val, success, not_supported, missing)
			if missing then
				return "---"
			elseif not_supported then
				return "⊘ NOT SUPPORTED"
			elseif success then
				return "✅ " .. (val or "PASS")
			else
				return "❌ " .. (val or "FAIL")
			end
		end

		local cem_cell = format_cell(cem_val, cem_success, cem_not_supported, not cem_data)
		local wc_cell = format_cell(wc_val, wc_success, wc_not_supported, not wc_data)

		print(string.format("%-20s | %-20s | %-20s", bench.display, cem_cell, wc_cell))
	end

	print(string.rep("-", 70))

	-- Overall success rates
	if cem_data and cem_data.benchmarks then
		local cem_passed = 0
		local cem_total = 0
		for _, bench in ipairs(benchmarks) do
			local result = cem_data.benchmarks[bench.name]
			if result then
				cem_total = cem_total + 1
				if result.success then
					cem_passed = cem_passed + 1
				end
			end
		end
		print(string.format("CEM Overall: %.0f%% (%d/%d)", (cem_passed / cem_total) * 100, cem_passed, cem_total))
	end

	if wc_data and wc_data.benchmarks then
		local wc_passed = 0
		local wc_total = 0
		for _, bench in ipairs(benchmarks) do
			local result = wc_data.benchmarks[bench.name]
			if result then
				wc_total = wc_total + 1
				if result.success then
					wc_passed = wc_passed + 1
				end
			end
		end
		print(string.format("WC-Toolkit Overall: %.0f%% (%d/%d)", (wc_passed / wc_total) * 100, wc_passed, wc_total))
	end

	print(string.rep("=", 70) .. "\n")
end

main()
