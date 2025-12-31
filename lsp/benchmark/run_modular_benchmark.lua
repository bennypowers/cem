#!/usr/bin/env nvim --headless --clean -l
-- Modular LSP Benchmark Runner
-- Usage: nvim --headless --clean -u configs/cem-minimal.lua -l run_modular_benchmark.lua
-- Usage: nvim --headless --clean -u configs/wc-toolkit-minimal.lua -l run_modular_benchmark.lua

-- Ensure the benchmark directory is in the path
local script_dir = vim.fn.getcwd()
package.path = script_dir .. "/?.lua;" .. package.path

-- Core protocol benchmark modules (pure LSP timing without editor overhead)
local startup_benchmark = require("modules.startup_benchmark")
local hover_benchmark = require("modules.hover_benchmark")
local completion_benchmark = require("modules.completion_benchmark")
local diagnostics_benchmark = require("modules.diagnostics_benchmark")
local attribute_hover_benchmark = require("modules.attribute_hover_benchmark")
local references_performance_benchmark = require("modules.references_performance_benchmark")

local function run_all_benchmarks()
	local overall_start_time = vim.fn.reltime()
	local max_total_time_seconds = 300 -- 5 minutes total limit for CI safety

	local server_name = _G.BENCHMARK_LSP_NAME or "unknown"
	local config = _G.BENCHMARK_LSP_CONFIG

	if not config then
		print("ERROR: No LSP configuration found. Make sure to use the appropriate config file:")
		print("  nvim --headless --clean -u configs/cem-minimal.lua -l run_modular_benchmark.lua")
		print("  nvim --headless --clean -u configs/wc-toolkit-minimal.lua -l run_modular_benchmark.lua")
		return
	end

	print(string.format("Running modular benchmarks for %s LSP server (max %ds)", server_name, max_total_time_seconds))
	print("=" .. string.rep("=", 50))

	-- Use large project fixture for expanded features
	local fixture_dir = script_dir .. "/fixtures/large_project"

	-- Check if fixture exists
	if vim.fn.isdirectory(fixture_dir) == 0 then
		print("ERROR: Fixture directory not found: " .. fixture_dir)
		return
	end

	local all_results = {
		server_name = server_name,
		timestamp = os.date("%Y-%m-%d %H:%M:%S"),
		fixture_dir = fixture_dir,
		benchmarks = {},
	}

	-- Core protocol benchmarks (pure LSP timing without editor overhead)
	local benchmarks = {
		{ name = "startup", module = startup_benchmark },
		{ name = "hover", module = hover_benchmark },
		{ name = "completion", module = completion_benchmark },
		{ name = "diagnostics", module = diagnostics_benchmark },
		{ name = "attribute_hover", module = attribute_hover_benchmark },
		{ name = "references_performance", module = references_performance_benchmark },
	}

	for _, benchmark in ipairs(benchmarks) do
		-- Check total time limit
		local elapsed_time = vim.fn.reltime(overall_start_time)[1]
		if elapsed_time >= max_total_time_seconds then
			print(
				string.format("\n⏰ Time limit reached (%ds), stopping remaining benchmarks", max_total_time_seconds)
			)
			break
		end

		io.write(string.format("Running %s... ", benchmark.name))
		io.flush()

		local success, result = pcall(function()
			return benchmark.module["run_" .. benchmark.name .. "_benchmark"](config, fixture_dir)
		end)

		if success then
			all_results.benchmarks[benchmark.name] = result
			if result.success then
				print("✅")
			else
				print(string.format("❌ %s", result.error or "unknown error"))
			end
		else
			print(string.format("❌ crashed: %s", result))
			all_results.benchmarks[benchmark.name] = {
				success = false,
				error = "benchmark_crashed: " .. tostring(result),
			}
		end

		-- Small delay between benchmarks
		vim.wait(1000)
	end

	-- Generate summary report
	print("\n" .. string.rep("=", 50))
	print(string.format("BENCHMARK SUMMARY FOR %s", string.upper(server_name)))
	print(string.rep("=", 50))

	local successful_benchmarks = 0
	local total_benchmarks = 0

	for benchmark_name, result in pairs(all_results.benchmarks) do
		total_benchmarks = total_benchmarks + 1
		if result.success then
			successful_benchmarks = successful_benchmarks + 1
			local status = "✅ PASS"
			local details = ""

			if result.duration_ms then
				details = details .. string.format(" (%.2fms)", result.duration_ms)
			end
			if result.success_rate then
				details = details .. string.format(" [%.0f%% success]", result.success_rate * 100)
			end

			-- Add correctness information to summary
			local correctness = ""
			if result.scenario_results then
				-- Lit template benchmark with test correctness
				local total_tests = 0
				local successful_tests = 0
				for _, scenario in ipairs(result.scenario_results) do
					if scenario.total_tests and scenario.successful_tests then
						total_tests = total_tests + scenario.total_tests
						successful_tests = successful_tests + scenario.successful_tests
					end
				end
				if total_tests > 0 then
					correctness = string.format(" [%.0f%% correct]", (successful_tests / total_tests) * 100)
				end
			elseif result.completion_results then
				-- Completion benchmark - show avg items
				local total_items = 0
				local count = 0
				for _, comp_result in ipairs(result.completion_results) do
					if comp_result.avg_item_count then
						total_items = total_items + comp_result.avg_item_count
						count = count + 1
					end
				end
				if count > 0 then
					correctness = string.format(" [%.1f items]", total_items / count)
				end
			elseif result.hover_results then
				-- Hover benchmark - show avg content length
				local total_content = 0
				local count = 0
				for _, hover_result in ipairs(result.hover_results) do
					if hover_result.avg_content_length then
						total_content = total_content + hover_result.avg_content_length
						count = count + 1
					end
				end
				if count > 0 then
					correctness = string.format(" [%.0f chars]", total_content / count)
				end
			elseif result.total_references_found ~= nil then
				-- References benchmark - show total references found (0 = feature not implemented)
				correctness = string.format(" [%d refs]", result.total_references_found)
			end

			print(string.format("%-20s %s%s%s", benchmark_name, status, details, correctness))
		else
			print(string.format("%-20s ❌ FAIL - %s", benchmark_name, result.error or "unknown"))
		end
	end

	local overall_success_rate = total_benchmarks > 0 and (successful_benchmarks / total_benchmarks) or 0
	print(string.rep("-", 50))
	print(
		string.format(
			"Overall Success Rate: %.1f%% (%d/%d)",
			overall_success_rate * 100,
			successful_benchmarks,
			total_benchmarks
		)
	)

	if overall_success_rate >= 0.8 then
		print("🎉 Server performance is GOOD")
	elseif overall_success_rate >= 0.6 then
		print("⚠️  Server performance is FAIR")
	else
		print("🚨 Server performance is POOR")
	end

	-- Save individual server results to temporary file for combining later
	local temp_results_file = string.format("/tmp/cem-benchmark-%s.json", server_name)
	local json_content = vim.fn.json_encode(all_results)
	vim.fn.writefile({ json_content }, temp_results_file)
	print(string.format("\nResults saved to: %s", temp_results_file))

	-- Final timing report
	local total_elapsed_time = vim.fn.reltime(overall_start_time)[1]
	print(string.format("Total benchmark time: %.1fs / %.0fs limit", total_elapsed_time, max_total_time_seconds))

	print(string.rep("=", 50))
end

-- Run benchmarks
run_all_benchmarks()
