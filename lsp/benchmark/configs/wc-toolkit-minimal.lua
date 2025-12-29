-- Minimal Neovim configuration for wc-toolkit LSP benchmarking
-- Assumes wc-language-server installed via Mason: :MasonInstall wc-language-server
-- See: https://wc-toolkit.com/integrations/neovim/

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

-- wc-toolkit LSP configuration (using Mason-installed binary)
local wc_toolkit_config = {
  name = 'wc-language-server',
  cmd = { 'wc-language-server', '--stdio' },  -- Mason-installed binary
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
