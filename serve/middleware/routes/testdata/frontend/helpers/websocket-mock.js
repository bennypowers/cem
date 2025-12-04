/**
 * Mock WebSocket implementation for testing
 * Simulates WebSocket behavior without needing a real server
 */
export class MockWebSocket {
  static CONNECTING = 0;
  static OPEN = 1;
  static CLOSING = 2;
  static CLOSED = 3;

  constructor(url) {
    this.url = url;
    this.readyState = MockWebSocket.CONNECTING;
    this.onopen = null;
    this.onmessage = null;
    this.onclose = null;
    this.onerror = null;
    this.sent = [];

    // Store for testing
    MockWebSocket.instances.push(this);
  }

  static instances = [];

  static reset() {
    this.instances = [];
  }

  send(data) {
    if (this.readyState !== MockWebSocket.OPEN) {
      throw new Error('WebSocket is not open');
    }
    this.sent.push(data);
  }

  close(code = 1000, reason = '') {
    this.readyState = MockWebSocket.CLOSING;
    setTimeout(() => {
      this.readyState = MockWebSocket.CLOSED;
      this.onclose?.({ code, reason });
    }, 0);
  }

  // Test helpers
  simulateOpen() {
    this.readyState = MockWebSocket.OPEN;
    setTimeout(() => {
      this.onopen?.({ target: this });
    }, 0);
  }

  simulateMessage(data) {
    if (this.readyState !== MockWebSocket.OPEN) {
      throw new Error('Cannot send message on closed WebSocket');
    }
    setTimeout(() => {
      this.onmessage?.({
        data: typeof data === 'string' ? data : JSON.stringify(data),
        target: this
      });
    }, 0);
  }

  simulateError(error = {}) {
    setTimeout(() => {
      this.onerror?.({ ...error, target: this });
    }, 0);
  }

  simulateClose(code = 1000, reason = '') {
    this.readyState = MockWebSocket.CLOSING;
    setTimeout(() => {
      this.readyState = MockWebSocket.CLOSED;
      this.onclose?.({ code, reason, target: this });
    }, 0);
  }
}

/**
 * Install WebSocket mock globally
 * @returns {function} Cleanup function to restore original WebSocket
 */
export function installWebSocketMock() {
  const original = globalThis.WebSocket;

  // Copy static properties
  MockWebSocket.CONNECTING = WebSocket.CONNECTING;
  MockWebSocket.OPEN = WebSocket.OPEN;
  MockWebSocket.CLOSING = WebSocket.CLOSING;
  MockWebSocket.CLOSED = WebSocket.CLOSED;

  globalThis.WebSocket = MockWebSocket;

  return () => {
    globalThis.WebSocket = original;
    MockWebSocket.reset();
  };
}
