"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.activate = activate;
exports.deactivate = deactivate;
const vscode = require("vscode");
const path = require("path");
const fs = require("fs");
const os = require("os");
const node_1 = require("vscode-languageclient/node");
let client;
function getBinaryName() {
    const platform = os.platform();
    const arch = os.arch();
    // For now, use 'cem' since we're bundling the npm package
    // In future, could support platform-specific binaries like design tokens does
    return 'cem';
}
function findCemExecutable(context) {
    // First, try the bundled version from client dependencies
    const bundledCem = path.join(context.extensionPath, 'client', 'node_modules', '@pwrs', 'cem', 'bin', 'cem');
    if (fs.existsSync(bundledCem)) {
        return bundledCem;
    }
    // Fallback to user configuration or PATH
    const config = vscode.workspace.getConfiguration('cem.lsp');
    const configuredExecutable = config.get('executable', '');
    if (configuredExecutable) {
        return configuredExecutable;
    }
    // Default to binary name in PATH as last resort
    return getBinaryName();
}
function activate(context) {
    try {
        // Get configuration
        const config = vscode.workspace.getConfiguration('cem.lsp');
        const debugLogging = config.get('debugLogging', false);
        const trace = config.get('trace.server', 'off');
        // Find the CEM executable
        const executable = findCemExecutable(context);
        // Create server options
        const serverOptions = {
            command: executable,
            args: ['lsp'],
            transport: node_1.TransportKind.stdio,
        };
        // Create client options
        const clientOptions = {
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
        client = new node_1.LanguageClient('cem-language-server', 'CEM Language Server', serverOptions, clientOptions);
        // Start the client and server
        client.start().then(() => {
            console.log(`CEM Language Server started successfully using: ${executable}`);
            if (debugLogging) {
                vscode.window.showInformationMessage(`CEM LSP started with executable: ${executable}`);
            }
        }).catch((error) => {
            console.error('Failed to start CEM Language Server:', error);
            vscode.window.showErrorMessage(`Failed to start CEM Language Server using "${executable}". ` +
                `Please check that CEM is installed or configure a custom path in settings. ` +
                `Error: ${error.message || error}`);
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
                vscode.window.showInformationMessage('CEM Language Server configuration changed. Restart the server to apply changes.', 'Restart').then((selection) => {
                    if (selection === 'Restart') {
                        vscode.commands.executeCommand('cem.restartServer');
                    }
                });
            }
        });
        context.subscriptions.push(configWatcher);
    }
    catch (error) {
        console.error('Failed to activate CEM Language Server extension:', error);
        vscode.window.showErrorMessage(`Failed to activate CEM Language Server: ${error}`);
    }
}
function deactivate() {
    if (!client) {
        return undefined;
    }
    return client.stop();
}
//# sourceMappingURL=extension.js.map