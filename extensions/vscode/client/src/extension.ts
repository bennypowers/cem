import * as vscode from 'vscode';
import * as path from 'path';
import * as fs from 'fs';
import * as os from 'os';
import {
  LanguageClient,
  LanguageClientOptions,
  ServerOptions,
  TransportKind, } from 'vscode-languageclient/node';

let client: LanguageClient;

function getBinaryName(): string {
  const platform = os.platform();
  const arch = os.arch();

  // Determine the OS-specific binary name
  const archMapping: Record<string, string> = {
    arm64: "aarch64",
    x64: "x86_64",
  };

  const osMapping: Record<string, string> = {
    darwin: "apple-darwin",
    linux: "unknown-linux-gnu",
    win32: "pc-windows-msvc.exe",
  };

  const architecture = archMapping[arch];
  const operatingSystem = osMapping[platform];

  if (!architecture || !operatingSystem) {
    throw new Error(
      `Unsupported platform or architecture: ${platform}-${arch}`,
    );
  }

  return `cem-${architecture}-${operatingSystem}`;
}

function findCemExecutable(context: vscode.ExtensionContext): string {
  // First, try the bundled platform-specific binary
  const binaryName = getBinaryName();
  const bundledCem = path.join(context.extensionPath, 'dist', 'bin', binaryName);
  if (fs.existsSync(bundledCem)) {
    return bundledCem;
  }

  // Fallback to npm package version
  const npmCem = path.join(context.extensionPath, 'client', 'node_modules', '@pwrs', 'cem', 'bin', 'cem');
  if (fs.existsSync(npmCem)) {
    return npmCem;
  }

  // Fallback to user configuration or PATH
  const config = vscode.workspace.getConfiguration('cem.lsp');
  const configuredExecutable = config.get<string>('executable', '');

  if (configuredExecutable) {
    return configuredExecutable;
  }

  // Default to 'cem' in PATH as last resort
  return 'cem';
}

export function activate(context: vscode.ExtensionContext) {
  try {
    // Get configuration
    const config = vscode.workspace.getConfiguration('cem.lsp');
    const debugLogging = config.get<boolean>('debugLogging', false);
    const trace = config.get<string>('trace.server', 'off');

    // Find the CEM executable
    const executable = findCemExecutable(context);

    // Create server options
    const serverOptions: ServerOptions = {
      command: executable,
      args: ['lsp'],
      transport: TransportKind.stdio,
    };

    // Create client options
    const clientOptions: LanguageClientOptions = {
      documentSelector: [
        { scheme: 'file', language: 'html' },
        { scheme: 'file', language: 'typescript' },
        { scheme: 'file', language: 'javascript' },
      ],
      initializationOptions: {
        debugLogging,
      },
      outputChannelName: 'CEM Language Server',
      traceOutputChannel: trace !== 'off' ? vscode.window.createOutputChannel('CEM LSP Trace') : undefined,
    };

    // Create and start the language client
    client = new LanguageClient(
      'cem-language-server',
      'CEM Language Server',
      serverOptions,
      clientOptions
    );

    // Start the client and server
    client.start().then(() => {
      console.log(`CEM Language Server started successfully using: ${executable}`);
      if (debugLogging) {
        vscode.window.showInformationMessage(`CEM LSP started with executable: ${executable}`);
      }
    }).catch((error) => {
      console.error('Failed to start CEM Language Server:', error);
      vscode.window.showErrorMessage(
        `Failed to start CEM Language Server using "${executable}". ` +
        `Please check that CEM is installed or configure a custom path in settings. ` +
        `Error: ${error.message || error}`
      );
    });

    // Register commands
    const restartCommand = vscode.commands.registerCommand('cem.restartServer', async () => {
      if (client) {
        await client.stop();
        await client.start();
        vscode.window.showInformationMessage('CEM Language Server restarted');
      }
    });

    const showOutputCommand = vscode.commands.registerCommand('cem.showOutputChannel', () => {
      client?.outputChannel?.show();
    });

    // Add to context subscriptions for cleanup
    context.subscriptions.push(restartCommand, showOutputCommand);

    // Watch for configuration changes
    const configWatcher = vscode.workspace.onDidChangeConfiguration((event) => {
      if (event.affectsConfiguration('cem.lsp')) {
        vscode.window.showInformationMessage(
          'CEM Language Server configuration changed. Restart the server to apply changes.',
          'Restart'
        ).then((selection) => {
          if (selection === 'Restart') {
            vscode.commands.executeCommand('cem.restartServer');
          }
        });
      }
    });

    context.subscriptions.push(configWatcher);
  } catch (error) {
    console.error('Failed to activate CEM Language Server extension:', error);
    vscode.window.showErrorMessage(`Failed to activate CEM Language Server: ${error}`);
  }
}

export function deactivate(): Thenable<void> | undefined {
  if (!client) {
    return undefined;
  }
  return client.stop();
}
