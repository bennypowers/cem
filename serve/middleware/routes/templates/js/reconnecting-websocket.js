// WebSocket client with automatic reconnection logic

export class ReconnectingWebSocket {
  #url;
  #jitterMax;
  #retryCount = 0;
  #retryTimeout = null;
  #isConnected = false;
  #listeners = {
    open: [],
    message: [],
    close: [],
    error: [],
    reconnecting: []
  };

  // Public property - accessed externally for debugging
  ws = null;

  constructor(url, options = {}) {
    this.#url = url;
    this.#jitterMax = options.jitterMax ?? 1000;
  }

  connect() {
    this.ws = new WebSocket(this.#url);

    this.ws.onopen = (event) => {
      this.#isConnected = true;
      this.#retryCount = 0;
      clearTimeout(this.#retryTimeout);
      this.#retryTimeout = null;
      this.#dispatch('open', event);
    };

    this.ws.onmessage = (event) => {
      this.#dispatch('message', event);
    };

    this.ws.onclose = (event) => {
      this.#isConnected = false;
      this.ws = null;
      this.#dispatch('close', event);
      this.#reconnect();
    };

    this.ws.onerror = (event) => {
      this.#isConnected = false;
      this.#dispatch('error', event);
    };

    return this;
  }

  send(data) {
    if (this.ws && this.#isConnected) {
      this.ws.send(data);
    }
  }

  close() {
    clearTimeout(this.#retryTimeout);
    this.#retryTimeout = null;
    if (this.ws) {
      this.ws.close();
      this.ws = null;
    }
  }

  on(event, callback) {
    if (this.#listeners[event]) {
      this.#listeners[event].push(callback);
    }
    return this;
  }

  off(event, callback) {
    if (this.#listeners[event]) {
      this.#listeners[event] = this.#listeners[event].filter(cb => cb !== callback);
    }
    return this;
  }

  #dispatch(event, data) {
    if (this.#listeners[event]) {
      this.#listeners[event].forEach(callback => callback(data));
    }
  }

  #calculateBackoff() {
    // Progressive backoff strategy that stays aggressive early, backs off gradually
    // Attempts 1-20: 1 second (20 seconds total)
    // Attempts 21-40: 2 seconds (40 seconds more = 1 minute total)
    // Attempts 41-60: 5 seconds (100 seconds more = 2:40 total)
    // Attempts 61-80: 10 seconds (200 seconds more = 6:00 total)
    // Attempts 81+: 30 seconds

    let delay;
    if (this.#retryCount <= 20) {
      delay = 1000; // 1 second
    } else if (this.#retryCount <= 40) {
      delay = 2000; // 2 seconds
    } else if (this.#retryCount <= 60) {
      delay = 5000; // 5 seconds
    } else if (this.#retryCount <= 80) {
      delay = 10000; // 10 seconds
    } else {
      delay = 30000; // 30 seconds
    }

    // Add jitter to avoid thundering herd
    const jitter = Math.random() * this.#jitterMax;
    return delay + jitter;
  }

  #reconnect() {
    if (this.#retryTimeout) return; // Already scheduled

    this.#retryCount++;
    const delay = this.#calculateBackoff();

    this.#dispatch('reconnecting', {
      attempt: this.#retryCount,
      delay
    });

    this.#retryTimeout = setTimeout(() => {
      this.#retryTimeout = null;
      this.connect();
    }, delay);
  }

  retry() {
    clearTimeout(this.#retryTimeout);
    this.#retryTimeout = null;
    this.#retryCount = 0;
    this.connect();
  }
}
