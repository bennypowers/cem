# Development Testing Scripts

This directory contains manual testing and debugging scripts for the LSP benchmarks. These are not part of the automated benchmark suite but can be useful for development and debugging.

## Files

- **`test_wc_capabilities.js`** - Manual script to test wc-toolkit LSP server initialization
- **`test_wc_server_direct.js`** - Manual script to test wc-toolkit LSP with proper framing
- **`wc.config.js`** - Example wc-toolkit configuration file

## Usage

These scripts are for manual testing only. Run them with Node.js:

```bash
node test_wc_capabilities.js
node test_wc_server_direct.js
```

**Note**: These scripts assume the wc-language-server is set up via `setup_wc_toolkit.sh`.
