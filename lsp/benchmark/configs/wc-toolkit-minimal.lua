-- Minimal Neovim configuration for wc-toolkit LSP benchmarking
-- This config starts headless Neovim with only wc-toolkit LSP support

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

-- Find wc-toolkit server path
local function get_wc_toolkit_server_path()
  local current_dir = vim.fn.getcwd()
  local server_path = current_dir .. '/servers/wc-language-server/packages/language-server/bin/wc-language-server.js'
  if vim.fn.filereadable(server_path) == 1 then
    return {'node', server_path, '--stdio'}
  end
  error("wc-toolkit LSP server not found at " .. server_path .. ". Please run the setup script first.")
end

-- Minimal LSP configuration for wc-toolkit
local wc_toolkit_config = {
  name = 'wc-toolkit-lsp',
  cmd = get_wc_toolkit_server_path(),
  root_markers = {
    'custom-elements.json',
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