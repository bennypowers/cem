-- Minimal Neovim configuration for wc-toolkit LSP benchmarking
-- Assumes wc-language-server installed via setup_wc_toolkit.sh
-- Run: make setup-wc-toolkit

local benchmark_dir = vim.fn.getcwd()
vim.opt.runtimepath:prepend(benchmark_dir)

-- Load common configuration
local common = require("configs.common")

-- Setup common settings and tree-sitter
common.setup_common_settings()
common.setup_treesitter()

-- wc-toolkit LSP configuration (using local npx wrapper)
local wc_server_path = benchmark_dir .. '/bin/wc-language-server'

local wc_toolkit_config = {
  name = 'wc-language-server',
  cmd = { wc_server_path, '--stdio' },  -- Local npx wrapper
  root_markers = {
    'custom-elements.json',
    'wc.config.js',
    'package.json',
    '.git',
  },
  filetypes = {
    'html',
    'typescript',
    'javascript',
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

-- Auto-start wc-toolkit LSP for supported filetypes
vim.api.nvim_create_autocmd("FileType", {
  pattern = wc_toolkit_config.filetypes,
  callback = function()
    vim.lsp.start(wc_toolkit_config)
  end,
})

-- Export config for benchmark modules
_G.BENCHMARK_LSP_CONFIG = wc_toolkit_config
_G.BENCHMARK_LSP_NAME = "wc-toolkit"
