-- Minimal Neovim configuration for cem LSP benchmarking
-- This config starts headless Neovim with only cem LSP support

-- Set up runtime path to include benchmark directory
local benchmark_dir = vim.fn.getcwd()
vim.opt.runtimepath:prepend(benchmark_dir)

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

-- Minimal LSP configuration for cem
local cem_config = {
	name = "cem-lsp",
	cmd = { "cem", "lsp" },
	root_markers = {
		"custom-elements.json",
		"package.json",
		".git",
	},
	filetypes = {
		"html",
		"typescript",
		"javascript",
	},
	single_file_support = true,
	capabilities = vim.lsp.protocol.make_client_capabilities(),
	on_attach = function(client, bufnr)
		-- Minimal attach - no keybindings or fancy features
	end,
	on_init = function(client, initialize_result)
		-- Minimal init
	end,
}

-- Auto-start cem LSP for supported filetypes
vim.api.nvim_create_autocmd("FileType", {
	pattern = cem_config.filetypes,
	callback = function()
		vim.lsp.start(cem_config)
	end,
})

-- Export config for benchmark modules
_G.BENCHMARK_LSP_CONFIG = cem_config
_G.BENCHMARK_LSP_NAME = "cem"
