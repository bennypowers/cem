// WebSocket client with automatic reconnection logic

export class ReconnectingWebSocket {
  constructor(url, options = {}) {
    this.url = url;
    this.config = {
      baseDelay: options.baseDelay || 1000,
      maxDelay: options.maxDelay || 30000,
      jitterMax: options.jitterMax || 1000,
      ...options
    };

    this.ws = null;
    this.retryCount = 0;
    this.retryTimeout = null;
    this.isConnected = false;
    this.listeners = {
      open: [],
      message: [],
      close: [],
      error: [],
      reconnecting: []
    };
  }

  connect() {
    this.ws = new WebSocket(this.url);

    this.ws.onopen = (event) => {
      this.isConnected = true;
      this.retryCount = 0;
      clearTimeout(this.retryTimeout);
      this.retryTimeout = null;
      this._dispatch('open', event);
    };

    this.ws.onmessage = (event) => {
      this._dispatch('message', event);
    };

    this.ws.onclose = (event) => {
      this.isConnected = false;
      this.ws = null;
      this._dispatch('close', event);
      this._reconnect();
    };

    this.ws.onerror = (event) => {
      this.isConnected = false;
      this._dispatch('error', event);
    };

    return this;
  }

  send(data) {
    if (this.ws && this.isConnected) {
      this.ws.send(data);
    }
  }

  close() {
    clearTimeout(this.retryTimeout);
    this.retryTimeout = null;
    if (this.ws) {
      this.ws.close();
      this.ws = null;
    }
  }

  on(event, callback) {
    if (this.listeners[event]) {
      this.listeners[event].push(callback);
    }
    return this;
  }

  off(event, callback) {
    if (this.listeners[event]) {
      this.listeners[event] = this.listeners[event].filter(cb => cb !== callback);
    }
    return this;
  }

  _dispatch(event, data) {
    if (this.listeners[event]) {
      this.listeners[event].forEach(callback => callback(data));
    }
  }

  _calculateBackoff() {
    const exponential = Math.pow(2, this.retryCount) * this.config.baseDelay;
    const jitter = Math.random() * this.config.jitterMax;
    return Math.min(exponential + jitter, this.config.maxDelay);
  }

  _reconnect() {
    if (this.retryTimeout) return; // Already scheduled

    this.retryCount++;
    const delay = this._calculateBackoff();

    this._dispatch('reconnecting', {
      attempt: this.retryCount,
      delay
    });

    this.retryTimeout = setTimeout(() => {
      this.retryTimeout = null;
      this.connect();
    }, delay);
  }

  retry() {
    clearTimeout(this.retryTimeout);
    this.retryTimeout = null;
    this.retryCount = 0;
    this.connect();
  }
}
