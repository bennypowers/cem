-- Minimal Neovim configuration for cem LSP benchmarking
-- This config starts headless Neovim with only cem LSP support

-- Set up runtime path to include benchmark directory
local benchmark_dir = vim.fn.getcwd()
vim.opt.runtimepath:prepend(benchmark_dir)

-- Load common configuration
local common = require("configs.common")

-- Setup common settings and tree-sitter
common.setup_common_settings()
common.setup_treesitter()

-- Minimal LSP configuration for cem
-- Use built binary from dist/ directory (absolute path from benchmark directory)
local cem_bin = benchmark_dir .. "/../../dist/cem"
local cem_config = {
	name = "cem-lsp",
	cmd = { cem_bin, "lsp" },
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
	on_attach = function(_, _) end,
	on_init = function(_, _) end,
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
