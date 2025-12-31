-- Common configuration for LSP benchmarks
-- Handles tree-sitter parser setup for dynamic position finding

local M = {}

--- Bootstrap nvim-treesitter and ensure HTML/TypeScript parsers are installed
--- This is required for dynamic position finding in benchmarks
function M.setup_treesitter()
	local data_path = vim.fn.stdpath("data")

	-- Check for nvim-treesitter in two possible locations:
	-- 1. CI: ~/.local/share/nvim/site/pack/plugins/start/nvim-treesitter (installed by workflow)
	-- 2. Local: ~/.local/share/nvim/lazy/nvim-treesitter (installed manually with lazy.nvim)
	local treesitter_path = nil
	local ci_path = data_path .. "/site/pack/plugins/start/nvim-treesitter"
	local local_path = data_path .. "/lazy/nvim-treesitter"

	if vim.fn.isdirectory(ci_path) == 1 then
		treesitter_path = ci_path
		print("Using nvim-treesitter from CI location: " .. ci_path)
	elseif vim.fn.isdirectory(local_path) == 1 then
		treesitter_path = local_path
		print("Using nvim-treesitter from local location: " .. local_path)
	else
		error(
			"nvim-treesitter not found!\n"
				.. "Expected at either:\n"
				.. "  - "
				.. ci_path
				.. " (CI)\n"
				.. "  - "
				.. local_path
				.. " (local)\n"
				.. "Install with: nvim --headless -c 'TSInstallSync html typescript' -c 'quit'"
		)
	end

	-- IMPORTANT: Add nvim-treesitter to runtime path so it can be loaded with --clean flag
	vim.opt.rtp:prepend(treesitter_path)

	-- Verify HTML and TypeScript parsers are available
	-- NOTE: In CI, parsers are pre-installed by the workflow before running benchmarks
	-- This just verifies they're accessible, doesn't try to install them
	local missing_parsers = {}
	for _, lang in ipairs({ "html", "typescript" }) do
		-- Check if parser library exists and can be loaded
		local has_parser = pcall(function()
			-- Try to create a simple string parser - this will fail if parser isn't available
			vim.treesitter.get_string_parser("<test />", lang)
		end)

		if not has_parser then
			table.insert(missing_parsers, lang)
		end
	end

	if #missing_parsers > 0 then
		print("ERROR: Missing tree-sitter parsers: " .. table.concat(missing_parsers, ", "))
		print("")
		print("In CI, parsers should be pre-installed by the 'Install nvim-treesitter and parsers' workflow step.")
		print("Locally, install parsers by running:")
		print("  nvim --headless -c 'TSInstallSync html typescript' -c 'quit'")
		print("")
		error("Required tree-sitter parsers not available")
	end
end

--- Setup common benchmark settings (disable visual features for performance)
function M.setup_common_settings()
	-- Basic vim settings for LSP
	vim.opt.compatible = false
	vim.opt.hidden = true
	vim.opt.backup = false
	vim.opt.writebackup = false
	vim.opt.swapfile = false

	-- Disable all unnecessary features for performance
	vim.opt.number = false
	vim.opt.relativenumber = false
	vim.opt.cursorline = false
	vim.opt.cursorcolumn = false
	vim.opt.showmode = false
	vim.opt.showcmd = false
	vim.opt.ruler = false
	vim.opt.laststatus = 0
	vim.opt.signcolumn = "no"

	-- Suppress LSP info messages (prevents "No information available" spam during benchmarks)
	local original_notify = vim.notify
	vim.notify = function(msg, level, opts)
		-- Only suppress INFO level messages, keep warnings and errors visible
		if level ~= vim.log.levels.INFO then
			original_notify(msg, level, opts)
		end
	end
end

return M
