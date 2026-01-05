---
title: Installation
weight: 10
---

Install CEM and verify your installation.

## Installation Methods

### NPM Projects

For Node.js-based projects, install CEM as a development dependency:

```sh
npm install --save-dev @pwrs/cem
```

Or with pnpm:

```sh
pnpm add -D @pwrs/cem
```

### Go Binary

Install the latest version directly with Go:

```sh
go install bennypowers.dev/cem@latest
```

This installs the `cem` binary to your `$GOPATH/bin` directory.

### Build from Source

Clone the repository and build locally:

```sh
git clone https://github.com/bennypowers/cem.git
cd cem
make
```

The compiled binary will be available in the `dist/` directory. To install it in
`~/.local/bin`, run:

```sh
make install
```

## Verify Installation

After installation, verify CEM is working:

```sh
cem version
```

You should see the version number printed to your terminal.

## Shell Completion

CEM supports shell completion for Bash, Zsh, Fish, and PowerShell. The completion scripts provide tab completion for commands, flags, and file paths.

### Bash

To load completions in your current shell session:

```bash
source <(cem completion bash)
```

To load completions for every new session, execute once:

**Linux:**

```bash
cem completion bash > /etc/bash_completion.d/cem
```

**macOS:**

```bash
cem completion bash > $(brew --prefix)/etc/bash_completion.d/cem
```

### Zsh

To load completions in your current shell session:

```zsh
source <(cem completion zsh)
```

To load completions for every new session, execute once:

```zsh
cem completion zsh > "${fpath[1]}/_cem"
```

You will need to start a new shell for this setup to take effect.

### Fish

To load completions in your current shell session:

```fish
cem completion fish | source
```

To load completions for every new session, execute once:

```fish
cem completion fish > ~/.config/fish/completions/cem.fish
```

### PowerShell

To load completions in your current shell session:

```powershell
cem completion powershell | Out-String | Invoke-Expression
```

To load completions for every new session, add the output of the above command to your PowerShell profile.

## Next Steps

After installation, set up CEM for your development environment:

- **[LSP Integration](./lsp/)** - Configure language server features in your editor
- **[MCP Integration](./mcp/)** - Set up AI assistant integration
- **[Getting Started](/docs/usage/getting-started/)** - Create your first project