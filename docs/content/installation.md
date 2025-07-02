---
title: Installation
---

For go binaries:
```sh
go install bennypowers.dev/cem@latest
```

For NPM projects:

```sh
npm install --save-dev @pwrs/cem
```

Or clone this repository and build from source:

```sh
git clone https://github.com/bennypowers/cem.git
cd cem
make
```

## Shell Completion
`cem` supports shell completion for Bash, Zsh, Fish, and PowerShell. The completion scripts provide tab completion for commands, flags, and file paths.

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

