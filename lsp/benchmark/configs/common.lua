-- Common configuration for LSP benchmarks
-- Handles tree-sitter parser setup for dynamic position finding

local M = {}

--- Bootstrap nvim-treesitter and ensure HTML/TypeScript parsers are installed
--- This is required for dynamic position finding in benchmarks
function M.setup_treesitter()
	local data_path = vim.fn.stdpath("data")
	local lazypath = data_path .. "/lazy/lazy.nvim"
	local treesitter_path = data_path .. "/lazy/nvim-treesitter"

	-- Install lazy.nvim if not present
	if vim.fn.isdirectory(lazypath) == 0 then
		print("Installing lazy.nvim plugin manager...")
		vim.fn.system({
			"git",
			"clone",
			"--filter=blob:none",
			"https://github.com/folke/lazy.nvim.git",
			"--branch=stable",
			lazypath,
		})
	end
	vim.opt.rtp:prepend(lazypath)

	-- Install nvim-treesitter if not present
	if vim.fn.isdirectory(treesitter_path) == 0 then
		print("Installing nvim-treesitter...")
		require("lazy").setup({
			{
				"nvim-treesitter/nvim-treesitter",
				build = ":TSUpdate",
			},
		})
	else
		vim.opt.rtp:prepend(treesitter_path)
	end

	-- Ensure HTML and TypeScript parsers are installed
	local parsers_to_install = {}
	for _, lang in ipairs({ "html", "typescript" }) do
		-- Check if parser can actually parse content (not just if it's registered)
		local test_ok = pcall(function()
			vim.treesitter.get_string_parser("<div></div>", lang)
		end)

		if not test_ok then
			table.insert(parsers_to_install, lang)
		end
	end

	if #parsers_to_install > 0 then
		print(string.format("Installing tree-sitter parsers: %s", table.concat(parsers_to_install, ", ")))

		-- Check if tree-sitter CLI is available
		local has_tree_sitter_cli = vim.fn.executable("tree-sitter") == 1
		if not has_tree_sitter_cli then
			print("ERROR: tree-sitter CLI not found!")
			print("Install it with: npm install -g tree-sitter-cli")
			print("Or on Arch Linux: sudo pacman -S tree-sitter")
			print("Or on macOS: brew install tree-sitter")
			error("tree-sitter CLI required for parser installation")
		end

		-- Set up nvim-treesitter configs to enable installation commands
		require("nvim-treesitter.configs").setup({
			ensure_installed = parsers_to_install,
			sync_install = true,
		})
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
