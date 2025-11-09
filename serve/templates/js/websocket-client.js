// CEM Dev Server Live Reload Client
// Integrates WebSocket reconnection with UI components

import { ReconnectingWebSocket } from '/__cem/reconnecting-websocket.js';
import { CEMConnectionStatus } from '/__cem/connection-status.js';
import { CEMErrorDialog } from '/__cem/error-dialog.js';
import { CEMReconnectionContent } from '/__cem/reconnection-content.js';

class CEMReloadClient {
  constructor() {
    this.config = {
      baseDelay: 1000,
      maxDelay: 30000,
      jitterMax: 1000,
      overlayThreshold: 5,
      badgeFadeDelay: 2000
    };

    this.status = null;
    this.dialog = null;
    this.content = null;
    this.ws = null;
  }

  init() {
    // Create UI components
    this.createStatus();
    this.createDialog();

    // Create WebSocket client
    const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:';
    const url = protocol + '//' + window.location.host + '/__cem-reload';

    this.ws = new ReconnectingWebSocket(url, {
      baseDelay: this.config.baseDelay,
      maxDelay: this.config.maxDelay,
      jitterMax: this.config.jitterMax
    });

    // Expose for debugging
    window.__cemReloadSocket = this.ws.ws;

    // Set up event handlers
    this.ws
      .on('open', () => this.handleOpen())
      .on('message', (event) => this.handleMessage(event))
      .on('close', () => this.handleClose())
      .on('error', (error) => this.handleError(error))
      .on('reconnecting', (data) => this.handleReconnecting(data));

    // Start connection
    this.ws.connect();
  }

  createStatus() {
    this.status = document.createElement('cem-connection-status');
    document.body.appendChild(this.status);
  }

  createDialog() {
    this.dialog = document.createElement('cem-error-dialog');
    this.dialog.title = 'Development Server Disconnected';

    // Create reconnection content component
    this.content = document.createElement('cem-reconnection-content');

    // Handle events from content component
    this.content.addEventListener('reload', () => {
      window.location.reload();
    });

    this.content.addEventListener('retry', () => {
      this.dialog.hide();
      this.ws.retry();
    });

    this.dialog.appendChild(this.content);
    document.body.appendChild(this.dialog);
  }

  handleOpen() {
    console.log('[cem-serve] WebSocket connected');
    this.status.show('connected', 'Connected', {
      fadeDelay: this.config.badgeFadeDelay
    });
    this.dialog.hide();
  }

  handleMessage(event) {
    const data = JSON.parse(event.data);
    console.log('[cem-serve] Received message:', data);

    if (data.type === 'reload') {
      console.log('[cem-serve] Reloading page:', data.reason, data.files);
      window.location.reload();
    } else if (data.type === 'shutdown') {
      console.log('[cem-serve] Server shutting down gracefully');
      // Immediately show dialog and start countdown from attempt 30 (30s delay)
      this.status.show('reconnecting', 'Server restarting...');
      this.dialog.show();
      this.content.updateRetryInfo(30, 30000);
    } else if (data.type === 'logs') {
      // Dispatch custom event for log updates
      window.dispatchEvent(new CustomEvent('cem:logs', {
        detail: { logs: data.logs }
      }));
    }
  }

  handleClose() {
    console.log('[cem-serve] Connection closed');
  }

  handleError(error) {
    console.error('[cem-serve] WebSocket error:', error);
  }

  handleReconnecting({ attempt, delay }) {
    console.log(`[cem-serve] Reconnecting in ${Math.ceil(delay/1000)}s (attempt #${attempt})...`);

    this.status.show('reconnecting', `Reconnecting (attempt #${attempt})...`);

    // Show dialog after threshold
    if (attempt >= this.config.overlayThreshold) {
      this.dialog.show();
      this.content.updateRetryInfo(attempt, delay);
    }
  }

  destroy() {
    if (this.ws) {
      this.ws.close();
    }
    if (this.status) {
      this.status.remove();
    }
    if (this.dialog) {
      this.dialog.remove();
    }
  }
}

// Initialize on load
const client = new CEMReloadClient();
client.init();

// Cleanup on unload
window.addEventListener('beforeunload', () => client.destroy());
