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

// Buffer for accumulating LSP messages
let buffer = Buffer.alloc(0);

// Parse LSP messages from buffer
function parseLSPMessages(buffer) {
  const messages = [];
  let offset = 0;

  while (offset < buffer.length) {
    // Find Content-Length header
    const headerEnd = buffer.indexOf('\r\n\r\n', offset);
    if (headerEnd === -1) break; // Incomplete header

    const headerText = buffer.slice(offset, headerEnd).toString('utf8');
    const contentLengthMatch = headerText.match(/Content-Length: (\d+)/i);

    if (!contentLengthMatch) {
      console.error('No Content-Length header found');
      break;
    }

    const contentLength = parseInt(contentLengthMatch[1], 10);
    const messageStart = headerEnd + 4; // Skip \r\n\r\n
    const messageEnd = messageStart + contentLength;

    if (messageEnd > buffer.length) break; // Incomplete message body

    const messageBody = buffer.slice(messageStart, messageEnd).toString('utf8');

    try {
      messages.push(JSON.parse(messageBody));
    } catch (e) {
      console.error('Failed to parse message body:', e.message);
    }

    offset = messageEnd;
  }

  // Return remaining buffer (incomplete message)
  return { messages, remaining: buffer.slice(offset) };
}

server.stdout.on('data', (data) => {
  // Accumulate data into buffer
  buffer = Buffer.concat([buffer, data]);

  // Parse complete messages
  const { messages, remaining } = parseLSPMessages(buffer);
  buffer = remaining;

  // Process each complete message
  for (const message of messages) {
    console.log('Received LSP message:', JSON.stringify(message, null, 2));

    // Check for initialize response
    if (message.id === 1 && message.result && message.result.capabilities) {
      console.log('\n=== SERVER CAPABILITIES ===');
      console.log(JSON.stringify(message.result.capabilities, null, 2));
    }
  }
});

server.stderr.on('data', (data) => {
  console.error('Server Error:', data.toString());
});

server.on('close', (code) => {
  console.log('Server closed with code:', code);
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
