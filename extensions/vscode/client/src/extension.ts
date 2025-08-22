import * as vscode from 'vscode';
import * as path from 'path';
import * as fs from 'fs';
import * as os from 'os';
import {
  LanguageClient,
  LanguageClientOptions,
  ServerOptions,
} from 'vscode-languageclient/node';

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
  console.log('CEM: Finding executable...');

  // First, try the bundled platform-specific binary
  const binaryName = getBinaryName();
  console.log(`CEM: Looking for binary: ${binaryName}`);

  const bundledCem = path.join(context.extensionPath, 'dist', 'bin', binaryName);
  console.log(`CEM: Checking bundled path: ${bundledCem}`);

  if (fs.existsSync(bundledCem)) {
    console.log('CEM: Found bundled binary');
    return bundledCem;
  }

  // Fallback to npm package version
  const npmCem = path.join(context.extensionPath, 'client', 'node_modules', '@pwrs', 'cem', 'bin', 'cem');
  console.log(`CEM: Checking npm path: ${npmCem}`);

  if (fs.existsSync(npmCem)) {
    console.log('CEM: Found npm binary');
    return npmCem;
  }

  // Fallback to user configuration or PATH
  const config = vscode.workspace.getConfiguration('cem.lsp');
  const configuredExecutable = config.get<string>('executable', '');
  console.log(`CEM: Configured executable: ${configuredExecutable || 'none'}`);

  if (configuredExecutable) {
    console.log('CEM: Using configured executable');
    return configuredExecutable;
  }

  // Default to 'cem' in PATH as last resort
  console.log('CEM: Falling back to PATH lookup');
  return 'cem';
}

export function activate(context: vscode.ExtensionContext) {
  try {
    console.log('CEM: Extension activating...');

    // Get configuration
    const config = vscode.workspace.getConfiguration('cem.lsp');
    const debugLogging = config.get<boolean>('debugLogging', false);
    const trace = config.get<string>('trace.server', 'off');

    console.log(`CEM: Debug logging: ${debugLogging}, Trace: ${trace}`);

    // Find the CEM executable
    const executable = findCemExecutable(context);
    console.log(`CEM: Using executable: ${executable}`);

    // Create server options
    const serverOptions: ServerOptions = {
      command: executable,
      args: ['lsp', '--stdio'],
      options: {
        stdio: 'pipe'
      }
    };

    console.log('CEM: Server options:', serverOptions);

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

    console.log('CEM: Client options:', clientOptions);

    // Create and start the language client
    client = new LanguageClient(
      'cem-language-server',
      'CEM Language Server',
      serverOptions,
      clientOptions
    );

    console.log('CEM: Language client created, starting...');

    // Start the client and server
    client.start().then(() => {
      console.log(`CEM: Language Server started successfully using: ${executable}`);
      if (debugLogging) {
        vscode.window.showInformationMessage(`CEM LSP started with executable: ${executable}`);
      }
    }).catch((error) => {
      console.error('CEM: Failed to start Language Server:', error);
      console.error('CEM: Error details:', {
        message: error.message,
        stack: error.stack,
        executable,
        workingDir: process.cwd()
      });
      vscode.window.showErrorMessage(
        `Failed to start CEM Language Server using "${executable}". ` +
        `Please check that CEM is installed or configure a custom path in settings. ` +
        `Error: ${error.message || error}`
      );
    });

    // Register commands
    const restartCommand = vscode.commands.registerCommand('cem.restartServer', async () => {
      console.log('CEM: Restart command triggered');
      if (client) {
        console.log('CEM: Stopping client...');
        await client.stop();
        console.log('CEM: Starting client...');
        await client.start();
        console.log('CEM: Client restarted');
        vscode.window.showInformationMessage('CEM Language Server restarted');
      }
    });

    const showOutputCommand = vscode.commands.registerCommand('cem.showOutputChannel', () => {
      console.log('CEM: Show output command triggered');
      client?.outputChannel?.show();
    });

    // Add to context subscriptions for cleanup
    context.subscriptions.push(restartCommand, showOutputCommand);
    console.log('CEM: Commands registered');

    // Watch for configuration changes
    const configWatcher = vscode.workspace.onDidChangeConfiguration((event) => {
      if (event.affectsConfiguration('cem.lsp')) {
        console.log('CEM: Configuration changed');
        vscode.window.showInformationMessage(
          'CEM Language Server configuration changed. Restart the server to apply changes.',
          'Restart'
        ).then((selection) => {
          if (selection === 'Restart') {
            console.log('CEM: User chose to restart after config change');
            vscode.commands.executeCommand('cem.restartServer');
          }
        });
      }
    });

    context.subscriptions.push(configWatcher);
    console.log('CEM: Extension activation complete');
  } catch (error) {
    console.error('CEM: Failed to activate extension:', error);
    vscode.window.showErrorMessage(`Failed to activate CEM Language Server: ${error}`);
  }
}

export function deactivate(): Thenable<void> | undefined {
  console.log('CEM: Extension deactivating...');
  if (!client) {
    console.log('CEM: No client to stop');
    return undefined;
  }
  console.log('CEM: Stopping client...');
  return client.stop();
}
