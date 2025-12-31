#!/usr/bin/env node

const { spawn } = require('child_process');
const path = require('path');

// Start the proper wc-toolkit server
const serverPath = path.join(__dirname, 'servers/wc-language-server/packages/vscode/dist/server.js');
const fixtureDir = path.join(__dirname, 'fixtures/large_project');

const server = spawn('node', [serverPath, '--stdio'], {
  cwd: fixtureDir,
  stdio: ['pipe', 'pipe', 'pipe']
});

// Initialize request with proper TypeScript SDK
const tsdkPath = path.join(fixtureDir, 'node_modules/typescript/lib');
const initRequest = {
  jsonrpc: "2.0",
  id: 1,
  method: "initialize",
  params: {
    processId: process.pid,
    rootUri: `file://${fixtureDir}`,
    capabilities: {},
    initializationOptions: {
      typescript: {
        tsdk: tsdkPath
      }
    }
  }
};

let responseData = '';

server.stdout.on('data', (data) => {
  responseData += data.toString();
  console.log('Server Response:', data.toString());
});

server.stderr.on('data', (data) => {
  console.error('Server Error:', data.toString());
});

server.on('close', (code) => {
  console.log('Server closed with code:', code);
  
  // Try to parse the response
  try {
    const lines = responseData.split('\n').filter(line => line.trim());
    for (const line of lines) {
      if (line.includes('"result"') && line.includes('capabilities')) {
        const response = JSON.parse(line);
        console.log('\n=== SERVER CAPABILITIES ===');
        console.log(JSON.stringify(response.result.capabilities, null, 2));
        break;
      }
    }
  } catch (e) {
    console.log('Could not parse server response:', e.message);
    console.log('Raw response:', responseData);
  }
  
  process.exit(0);
});

// Send initialize request with LSP framing
const requestJson = JSON.stringify(initRequest);
const requestLength = Buffer.byteLength(requestJson, 'utf8');
const requestMessage = `Content-Length: ${requestLength}\r\n\r\n${requestJson}`;
console.log('Sending request:', requestMessage);
server.stdin.write(requestMessage);

// Send initialized notification with LSP framing
setTimeout(() => {
  const initializedNotification = {
    jsonrpc: "2.0",
    method: "initialized",
    params: {}
  };
  const notificationJson = JSON.stringify(initializedNotification);
  const notificationLength = Buffer.byteLength(notificationJson, 'utf8');
  const notificationMessage = `Content-Length: ${notificationLength}\r\n\r\n${notificationJson}`;
  server.stdin.write(notificationMessage);
}, 1000);

// Exit after 5 seconds
setTimeout(() => {
  server.kill();
}, 5000);
