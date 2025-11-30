// CEM Dev Server Live Reload Client
// WebSocket client with automatic reconnection and callback-based UI updates

/**
 * WebSocket client for CEM dev server live reload
 * Uses callbacks for UI updates - DOM management is handled by the caller
 * Includes automatic reconnection logic with progressive backoff
 */
export class CEMReloadClient {
  #url = null;
  #jitterMax;
  #retryCount = 0;
  #retryTimeout = null;
  #isConnected = false;

  constructor(options = {}) {
    this.config = {
      jitterMax: options.jitterMax ?? 1000,
      overlayThreshold: options.overlayThreshold ?? 15,
      badgeFadeDelay: options.badgeFadeDelay ?? 2000,
      protocol: options.protocol // Optional protocol override for testing
    };

    this.#jitterMax = this.config.jitterMax;
    this.callbacks = options.callbacks ?? {};
    this.ws = null;
  }

  async init() {
    // Fetch initial log history on page load
    try {
      const response = await fetch('/__cem/logs');
      if (response.ok) {
        const logs = await response.json();
        if (logs && logs.length > 0) {
          this.callbacks.onLogs?.(logs);
        }
      }
    } catch (err) {
      console.warn('[cem-serve] Failed to fetch initial logs:', err);
    }

    // Build WebSocket URL with current page as query parameter
    // Browsers don't send Referer header for WebSocket connections, so we pass it explicitly
    const httpProtocol = this.config.protocol ?? window.location.protocol;
    const protocol = httpProtocol === 'https:' ? 'wss:' : 'ws:';
    const pageURL = encodeURIComponent(window.location.pathname);
    this.#url = protocol + '//' + window.location.host + '/__cem/reload?page=' + pageURL;
    console.debug('[cem-serve] Connecting to WebSocket:', this.#url, 'from page:', window.location.pathname);

    // Start connection
    this.#connect();
  }

  #connect() {
    this.ws = new WebSocket(this.#url);

    // Expose for debugging
    window.__cemReloadSocket = this.ws;

    this.ws.onopen = () => {
      this.#isConnected = true;
      this.#retryCount = 0;
      clearTimeout(this.#retryTimeout);
      this.#retryTimeout = null;
      this.handleOpen();
    };

    this.ws.onmessage = (event) => {
      this.handleMessage(event);
    };

    this.ws.onclose = () => {
      this.#isConnected = false;
      this.ws = null;
      this.handleClose();
      this.#reconnect();
    };

    this.ws.onerror = (event) => {
      this.#isConnected = false;
      this.handleError(event);
    };
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

    this.handleReconnecting({ attempt: this.#retryCount, delay });

    this.#retryTimeout = setTimeout(() => {
      this.#retryTimeout = null;
      this.#connect();
    }, delay);
  }

  handleOpen() {
    this.callbacks.onOpen?.();
  }

  handleMessage(event) {
    const data = JSON.parse(event.data);
    console.debug('[cem-serve] Received message:', data);

    if (data.type === 'reload') {
      this.callbacks.onReload?.(data);
    } else if (data.type === 'shutdown') {
      this.callbacks.onShutdown?.();
    } else if (data.type === 'logs') {
      this.callbacks.onLogs?.(data.logs);
    } else if (data.type === 'error') {
      this.callbacks.onError?.(data);
    }
  }

  handleClose() {
    this.callbacks.onClose?.();
  }

  handleError(error) {
    this.callbacks.onError?.(error);
  }

  handleReconnecting({ attempt, delay }) {
    this.callbacks.onReconnecting?.({ attempt, delay });
  }

  retry() {
    clearTimeout(this.#retryTimeout);
    this.#retryTimeout = null;
    this.#retryCount = 0;
    this.#connect();
  }

  send(data) {
    if (this.ws && this.#isConnected) {
      this.ws.send(data);
    }
  }

  destroy() {
    clearTimeout(this.#retryTimeout);
    this.#retryTimeout = null;
    if (this.ws) {
      this.ws.close();
      this.ws = null;
    }
  }
}
