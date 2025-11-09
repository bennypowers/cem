// Configuration
const config = {
  baseDelay: 1000,           // 1 second
  maxDelay: 30000,           // 30 seconds max backoff
  jitterMax: 1000,           // Up to 1 second of jitter
  overlayThreshold: 5,       // Show dialog after 5 failed attempts
  badgeFadeDelay: 2000       // Fade badge after 2s when connected
};

// State
let ws = null;
let retryCount = 0;
let retryTimeout = null;
let statusBadge = null;
let errorDialog = null;
let isConnected = false;

// Create connection status badge
function createStatusBadge() {
    if (statusBadge) return;

    statusBadge = document.createElement('div');
    statusBadge.id = '__cem-status-badge';
    statusBadge.style.cssText = `
      position: fixed;
      bottom: 20px;
      right: 20px;
      padding: 8px 12px;
      border-radius: 6px;
      font-family: system-ui, -apple-system, sans-serif;
      font-size: 12px;
      font-weight: 500;
      z-index: 999999;
      display: flex;
      align-items: center;
      gap: 8px;
      box-shadow: 0 4px 6px rgba(0,0,0,0.3);
      transition: opacity 0.3s ease;
    `;
    document.body.appendChild(statusBadge);
  }

  // Update status badge
  function updateStatus(state, message) {
    if (!statusBadge) createStatusBadge();

    let color, bgColor, text, dot;

    switch(state) {
      case 'connected':
        color = '#10b981';
        bgColor = '#064e3b';
        text = message || 'Connected';
        dot = '🟢';
        statusBadge.style.opacity = '1';
        // Fade out after delay
        setTimeout(() => {
          if (isConnected) statusBadge.style.opacity = '0.3';
        }, config.badgeFadeDelay);
        break;
      case 'reconnecting':
        color = '#fbbf24';
        bgColor = '#78350f';
        text = message || `Reconnecting...`;
        dot = '🟡';
        statusBadge.style.opacity = '1';
        break;
      case 'disconnected':
        color = '#ef4444';
        bgColor = '#7f1d1d';
        text = message || 'Disconnected';
        dot = '🔴';
        statusBadge.style.opacity = '1';
        break;
    }

    statusBadge.style.color = color;
    statusBadge.style.backgroundColor = bgColor;
    statusBadge.style.border = `1px solid ${color}`;
    statusBadge.innerHTML = `<span>${dot}</span><span>${text}</span>`;
  }

  // Create error dialog using native <dialog>
  function createErrorDialog() {
    if (errorDialog) return;

    errorDialog = document.createElement('dialog');
    errorDialog.id = '__cem-error-dialog';
    errorDialog.style.cssText = `
      border: none;
      border-radius: 8px;
      padding: 0;
      max-width: 500px;
      width: 90%;
      box-shadow: 0 20px 25px -5px rgba(0,0,0,0.5);
    `;

    errorDialog.innerHTML = `
      <div style="
        background: #1e293b;
        color: #e2e8f0;
        padding: 24px;
        font-family: system-ui, -apple-system, sans-serif;
      ">
        <h2 style="
          margin: 0 0 16px 0;
          color: #60a5fa;
          font-size: 20px;
          font-weight: 600;
        ">Development Server Disconnected</h2>

        <p style="
          margin: 0 0 16px 0;
          line-height: 1.6;
          color: #cbd5e1;
        ">
          The connection to the development server was lost.
          Automatically retrying connection...
        </p>

        <div id="__cem-retry-info" style="
          background: #0f172a;
          padding: 12px;
          border-radius: 4px;
          margin-bottom: 16px;
          font-size: 14px;
          color: #94a3b8;
        "></div>

        <div style="
          display: flex;
          gap: 12px;
          justify-content: flex-end;
        ">
          <button id="__cem-reload-btn" style="
            padding: 8px 16px;
            background: #334155;
            color: #e2e8f0;
            border: 1px solid #475569;
            border-radius: 4px;
            cursor: pointer;
            font-size: 14px;
            font-weight: 500;
          ">Reload Page</button>

          <button id="__cem-retry-btn" style="
            padding: 8px 16px;
            background: #60a5fa;
            color: #0f172a;
            border: none;
            border-radius: 4px;
            cursor: pointer;
            font-size: 14px;
            font-weight: 600;
          ">Retry Now</button>
        </div>
      </div>
    `;

    // Button event listeners
    errorDialog.querySelector('#__cem-retry-btn').addEventListener('click', function() {
      hideOverlay();
      clearTimeout(retryTimeout);
      retryTimeout = null;
      retryCount = 0; // Reset to try immediately
      reconnect();
    });

    errorDialog.querySelector('#__cem-reload-btn').addEventListener('click', function() {
      window.location.reload();
    });

    // Handle ESC key
    errorDialog.addEventListener('close', function() {
      // Dialog closed with ESC - that's fine, keep retrying in background
    });

    document.body.appendChild(errorDialog);
  }

  // Show error overlay
  function showOverlay() {
    if (!errorDialog) createErrorDialog();
    if (!errorDialog.open) {
      errorDialog.showModal();
    }
  }

  // Hide error overlay
  function hideOverlay() {
    if (errorDialog && errorDialog.open) {
      errorDialog.close();
    }
  }

  // Update retry info in dialog
  function updateRetryInfo(nextRetryMs) {
    const retryInfo = document.getElementById('__cem-retry-info');
    if (retryInfo) {
      const seconds = Math.ceil(nextRetryMs / 1000);
      retryInfo.textContent = `Attempt #${retryCount}. Next retry in ${seconds}s...`;
    }
  }

  // Calculate exponential backoff with jitter
  function calculateBackoff() {
    const exponential = Math.pow(2, retryCount) * config.baseDelay;
    const jitter = Math.random() * config.jitterMax;
    return Math.min(exponential + jitter, config.maxDelay);
  }

  // Connect to WebSocket
  function connect() {
    const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:';
    ws = new WebSocket(protocol + '//' + window.location.host + '/__cem-reload');

    // Expose socket on window for debugging
    window.__cemReloadSocket = ws;

    ws.onopen = function() {
      console.log('[cem-serve] WebSocket connected');
      isConnected = true;
      retryCount = 0;
      clearTimeout(retryTimeout);
      retryTimeout = null;
      updateStatus('connected', 'Connected');
      hideOverlay();
    };

    ws.onmessage = function(event) {
      const data = JSON.parse(event.data);
      console.log('[cem-serve] Received reload payload:', data);
      if (data.type === 'reload') {
        console.log('[cem-serve] Reloading page:', data.reason, data.files);
        window.location.reload();
      }
    };

    ws.onclose = function() {
      console.log('[cem-serve] Connection closed');
      isConnected = false;
      ws = null;
      reconnect();
    };

    ws.onerror = function(error) {
      console.error('[cem-serve] WebSocket error:', error);
      isConnected = false;
    };
  }

  // Reconnect with exponential backoff
  function reconnect() {
    if (retryTimeout) return; // Already scheduled

    retryCount++;
    const delay = calculateBackoff();

    console.log(`[cem-serve] Reconnecting in ${Math.ceil(delay/1000)}s (attempt #${retryCount})...`);

    updateStatus('reconnecting', `Reconnecting (attempt #${retryCount})...`);

    // Show overlay after threshold
    if (retryCount >= config.overlayThreshold) {
      showOverlay();
      updateRetryInfo(delay);
    }

    retryTimeout = setTimeout(function() {
      retryTimeout = null;
      connect();
    }, delay);
  }

// Initialize
createStatusBadge();
connect();

// Cleanup on page unload
window.addEventListener('beforeunload', function() {
  if (ws) {
    ws.close();
  }
  clearTimeout(retryTimeout);
});
