// CEM Serve Chrome - Demo wrapper components

class CemServeDemo extends HTMLElement {
  static {
    customElements.define('cem-serve-demo', this);
  }

  #getShadowRoot() {
    // Check if demo content is in shadow root (shadow mode)
    return this.shadowRoot;
  }

  /**
   * Set an attribute on an element in the demo
   * @param {string} selector - CSS selector for target element
   * @param {string} attribute - Attribute name
   * @param {string} value - Attribute value
   * @returns {boolean} - Whether the operation succeeded
   */
  setDemoAttribute(selector, attribute, value) {
    const root = this.#getShadowRoot() || this;
    const target = root.querySelector(selector);
    if (target) {
      target.setAttribute(attribute, value);
      return true;
    }
    return false;
  }

  /**
   * Set a property on an element in the demo
   * @param {string} selector - CSS selector for target element
   * @param {string} property - Property name
   * @param {*} value - Property value
   * @returns {boolean} - Whether the operation succeeded
   */
  setDemoProperty(selector, property, value) {
    const root = this.#getShadowRoot() || this;
    const target = root.querySelector(selector);
    if (target) {
      target[property] = value;
      return true;
    }
    return false;
  }

  /**
   * Set a CSS custom property on an element in the demo
   * @param {string} selector - CSS selector for target element
   * @param {string} cssProperty - CSS custom property name (with or without --)
   * @param {string} value - CSS property value
   * @returns {boolean} - Whether the operation succeeded
   */
  setDemoCssCustomProperty(selector, cssProperty, value) {
    const root = this.#getShadowRoot() || this;
    const target = root.querySelector(selector);
    if (target) {
      const propertyName = cssProperty.startsWith('--') ? cssProperty : `--${cssProperty}`;
      target.style.setProperty(propertyName, value);
      return true;
    }
    return false;
  }
}

class CemServeChrome extends HTMLElement {
  #logContainer = null;
  #drawerOpen = false;

  static {
    customElements.define('cem-serve-chrome', this);
  }

  get demo() {
    return this.querySelector('cem-serve-demo');
  }

  get knobs() {
    return this.getAttribute('knobs') || '';
  }

  get tagName() {
    return this.getAttribute('tag-name') || '';
  }

  connectedCallback() {
    const shadow = this.shadowRoot;
    if (!shadow) return;

    // Set up debug overlay
    const debugButton = shadow.querySelector('.debug-button');
    const debugOverlay = shadow.querySelector('.debug-overlay');
    const debugClose = shadow.querySelector('.debug-close');
    const debugCopy = shadow.querySelector('.debug-copy');

    if (debugButton && debugOverlay) {
      debugButton.addEventListener('click', () => {
        this.#fetchDebugInfo();
        debugOverlay.setAttribute('open', '');
      });

      debugClose?.addEventListener('click', () => {
        debugOverlay.removeAttribute('open');
      });

      debugCopy?.addEventListener('click', () => {
        this.#copyDebugInfo();
      });

      debugOverlay.addEventListener('click', (e) => {
        if (e.target === debugOverlay) {
          debugOverlay.removeAttribute('open');
        }
      });
    }

    // Set up drawer toggle
    const drawerToggle = shadow.querySelector('.drawer-toggle');
    const drawerContent = shadow.querySelector('.drawer-content');

    if (drawerToggle && drawerContent) {
      drawerToggle.addEventListener('click', () => {
        this.#drawerOpen = !this.#drawerOpen;
        if (this.#drawerOpen) {
          drawerContent.setAttribute('open', '');
        } else {
          drawerContent.removeAttribute('open');
        }
      });
    }

    // Set up log container
    this.#logContainer = shadow.querySelector('#log-container');

    // Listen for server log messages from WebSocket
    this.#setupLogListener();

    console.log('[cem-serve-chrome] Demo chrome initialized for', this.tagName);
  }

  #fetchDebugInfo() {
    const shadow = this.shadowRoot;
    if (!shadow) return;

    // Populate browser info immediately
    const browserEl = shadow.querySelector('#debug-browser');
    const uaEl = shadow.querySelector('#debug-ua');
    if (browserEl) {
      const browser = this.#detectBrowser();
      browserEl.textContent = browser;
    }
    if (uaEl) {
      uaEl.textContent = navigator.userAgent;
    }

    // Fetch server debug info
    fetch('/__cem-debug')
      .then(res => res.json())
      .then(data => {
        const versionEl = shadow.querySelector('#debug-version');
        const osEl = shadow.querySelector('#debug-os');
        const watchDirEl = shadow.querySelector('#debug-watch-dir');
        const manifestSizeEl = shadow.querySelector('#debug-manifest-size');
        const demoCountEl = shadow.querySelector('#debug-demo-count');
        const demosListEl = shadow.querySelector('#debug-demos-list');
        const importMapEl = shadow.querySelector('#debug-importmap');

        if (versionEl) versionEl.textContent = data.version || '-';
        if (osEl) osEl.textContent = data.os || '-';
        if (watchDirEl) watchDirEl.textContent = data.watchDir || '-';
        if (manifestSizeEl) manifestSizeEl.textContent = data.manifestSize || '-';
        if (demoCountEl) demoCountEl.textContent = data.demoCount || '0';

        if (demosListEl && data.demos && data.demos.length > 0) {
          const demosList = data.demos.map(demo =>
            `${demo.tagName}: ${demo.description || '(no description)'}\n  Canonical: ${demo.canonicalURL}\n  Local Route: ${demo.localRoute}`
          ).join('\n\n');
          demosListEl.textContent = demosList;
        } else if (demosListEl) {
          demosListEl.textContent = 'No demos found in manifest';
        }

        if (importMapEl && data.importMap) {
          // Import map comes as an object, stringify it for display
          const importMapJSON = JSON.stringify(data.importMap, null, 2);
          importMapEl.textContent = importMapJSON;
          // Store for copying
          data.importMapJSON = importMapJSON;
        } else if (importMapEl) {
          importMapEl.textContent = 'No import map generated';
        }

        // Store data for copy function
        this._debugData = data;
      })
      .catch(err => {
        console.error('[cem-serve-chrome] Failed to fetch debug info:', err);
      });
  }

  #detectBrowser() {
    const ua = navigator.userAgent;
    if (ua.includes('Firefox/')) {
      const match = ua.match(/Firefox\/(\d+)/);
      return match ? `Firefox ${match[1]}` : 'Firefox';
    } else if (ua.includes('Edg/')) {
      const match = ua.match(/Edg\/(\d+)/);
      return match ? `Edge ${match[1]}` : 'Edge';
    } else if (ua.includes('Chrome/')) {
      const match = ua.match(/Chrome\/(\d+)/);
      return match ? `Chrome ${match[1]}` : 'Chrome';
    } else if (ua.includes('Safari/') && !ua.includes('Chrome')) {
      const match = ua.match(/Version\/(\d+)/);
      return match ? `Safari ${match[1]}` : 'Safari';
    }
    return 'Unknown';
  }

  #copyDebugInfo() {
    const shadow = this.shadowRoot;
    if (!shadow) return;

    // Collect all debug info
    const info = [];
    shadow.querySelectorAll('.debug-panel dl dt').forEach(dt => {
      const dd = dt.nextElementSibling;
      if (dd && dd.tagName === 'DD') {
        info.push(`${dt.textContent}: ${dd.textContent}`);
      }
    });

    // Include import map if available
    let importMapSection = '';
    if (this._debugData && this._debugData.importMapJSON) {
      importMapSection = `\n${'='.repeat(40)}\nImport Map:\n${'='.repeat(40)}\n${this._debugData.importMapJSON}\n`;
    }

    const debugText = `CEM Serve Debug Information
${'='.repeat(40)}
${info.join('\n')}${importMapSection}
${'='.repeat(40)}
Generated: ${new Date().toISOString()}`;

    navigator.clipboard.writeText(debugText)
      .then(() => {
        const copyButton = shadow.querySelector('.debug-copy');
        if (copyButton) {
          const originalText = copyButton.textContent;
          copyButton.textContent = 'Copied!';
          setTimeout(() => {
            copyButton.textContent = originalText;
          }, 2000);
        }
      })
      .catch(err => {
        console.error('[cem-serve-chrome] Failed to copy debug info:', err);
        alert('Failed to copy to clipboard');
      });
  }

  #setupLogListener() {
    // Listen for server log messages from the WebSocket client
    // The websocket-client.js dispatches 'cem:logs' events with server logs
    window.addEventListener('cem:logs', (event) => {
      if (event.logs) {
        this.#renderLogs(event.logs);
      }
    });
  }

  #renderLogs(logs) {
    if (!this.#logContainer) return;

    // Server logs come formatted from the logger
    this.#logContainer.innerHTML = logs.map(log => {
      const level = log.includes('[ERROR]') ? 'error' :
                    log.includes('[WARN]') ? 'warn' :
                    log.includes('[INFO]') ? 'info' :
                    log.includes('[DEBUG]') ? 'debug' : '';
      return `<div class="log-entry ${level}">${this.#escapeHtml(log)}</div>`;
    }).join('');

    // Auto-scroll to bottom
    this.#logContainer.scrollTop = this.#logContainer.scrollHeight;
  }

  #escapeHtml(text) {
    const div = document.createElement('div');
    div.textContent = text;
    return div.innerHTML;
  }
}
