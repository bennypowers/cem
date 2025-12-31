#!/usr/bin/env node

const { spawn } = require('child_process');
const path = require('path');

// Start the wc-toolkit LSP server
const serverPath = path.join(__dirname, 'servers/wc-language-server/packages/language-server/bin/wc-language-server.js');
const server = spawn('node', [serverPath, '--stdio'], {
  cwd: path.join(__dirname, 'fixtures/large_project'),
  stdio: ['pipe', 'pipe', 'pipe']
});

// Initialize request
const initRequest = {
  jsonrpc: "2.0",
  id: 1,
  method: "initialize",
  params: {
    processId: process.pid,
    rootUri: `file://${path.join(__dirname, 'fixtures/large_project')}`,
    capabilities: {}
  }
};

let responseData = '';

server.stdout.on('data', (data) => {
  responseData += data.toString();
  console.log('Response:', data.toString());
});

server.stderr.on('data', (data) => {
  console.error('Server Error:', data.toString());
});

server.on('close', (code) => {
  console.log('Server closed with code:', code);
  process.exit(0);
});

// Send initialize request
try {
  server.stdin.write(JSON.stringify(initRequest) + '\n');
} catch (err) {
  console.error('Failed to send initialize request:', err);
  process.exit(1);
}

// Exit after 5 seconds
setTimeout(() => {
  server.kill();
  console.log('\nFinal response data:', responseData);
  if (!responseData) {
    console.warn('Warning: No response received from server');
  }
}, 5000);