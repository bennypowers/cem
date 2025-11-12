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
  #initialLogsFetched = false;
  #drawerContent = null;
  #isInitialLoad = true;

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

  get drawerOpen() {
    return this.#drawerOpen;
  }

  set drawerOpen(value) {
    const newValue = Boolean(value);
    if (this.#drawerOpen === newValue) return;

    this.#drawerOpen = newValue;

    // Update DOM
    if (this.#drawerContent) {
      if (this.#drawerOpen) {
        this.#drawerContent.setAttribute('open', '');
        // Scroll to bottom when opening drawer
        this.#scrollLogsToBottom();
      } else {
        this.#drawerContent.removeAttribute('open');
      }
    }

    // Persist to localStorage
    localStorage.setItem('cem-serve-drawer-open', String(this.#drawerOpen));
  }

  connectedCallback() {
    const shadow = this.shadowRoot;
    if (!shadow) return;

    // Set up navigation drawer
    this.#setupNavigationDrawer(shadow);

    // Set up footer drawer tabs
    this.#setupTabs(shadow);

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

    // Set up drawer toggle with localStorage persistence
    const drawerToggle = shadow.querySelector('.drawer-toggle');
    this.#drawerContent = shadow.querySelector('.drawer-content');

    if (drawerToggle && this.#drawerContent) {
      // Restore drawer state from localStorage without animation
      const savedState = localStorage.getItem('cem-serve-drawer-open');
      if (savedState === 'true') {
        // Set state directly without triggering setter (no transitions class yet)
        this.#drawerOpen = true;
        this.#drawerContent.setAttribute('open', '');
      }

      drawerToggle.addEventListener('click', () => {
        // Enable transitions on first click (and keep enabled)
        this.#drawerContent.classList.add('transitions-enabled');
        this.#isInitialLoad = false;
        this.drawerOpen = !this.drawerOpen;
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

    const logEntries = logs.map(log => {
      // Log is now structured: { type: 'info'|'warning'|'error'|'debug', date: ISO8601, message: string }
      const date = new Date(log.date);
      const time = date.toLocaleTimeString();
      const badge = this.#getLogBadge(log.type);

      return `<div class="log-entry ${log.type}" data-log-id="${log.date}">
        <span class="log-badge ${log.type}">${badge}</span>
        <span class="log-time">${time}</span>
        <span class="log-message">${this.#escapeHtml(log.message)}</span>
      </div>`;
    }).join('');

    // Initial load: replace all logs (from fetch on page load - container is empty)
    // Individual logs: append to existing logs (from WebSocket stream - container has content)
    if (!this.#initialLogsFetched) {
      // First batch of logs from initial fetch - replace
      this.#logContainer.innerHTML = logEntries;
      this.#initialLogsFetched = true;

      // Scroll latest log into view after initial logs are rendered (if drawer is open)
      if (this.#drawerOpen) {
        this.#scrollLatestIntoView();
      }
    } else {
      // Individual log from WebSocket or subsequent updates - append
      this.#logContainer.insertAdjacentHTML('beforeend', logEntries);

      // Auto-scroll latest log into view if drawer is open (for streaming logs)
      if (this.#drawerOpen) {
        this.#scrollLatestIntoView();
      }
    }
  }

  #getLogBadge(type) {
    switch (type) {
      case 'info': return 'INFO';
      case 'warning': return 'WARN';
      case 'error': return 'ERROR';
      case 'debug': return 'DEBUG';
      default: return type.toUpperCase();
    }
  }

  #scrollLatestIntoView() {
    if (!this.#logContainer) return;

    requestAnimationFrame(() => {
      const lastLog = this.#logContainer.lastElementChild;
      if (lastLog) {
        lastLog.scrollIntoView({ behavior: 'auto', block: 'end' });
      }
    });
  }

  #scrollLogsToBottom() {
    if (!this.#logContainer) return;

    if (this.#isInitialLoad) {
      // Initial load - scroll immediately (no transition)
      this.#scrollLatestIntoView();
    } else {
      // User toggle - wait for drawer transition (300ms) plus a small buffer
      setTimeout(() => {
        this.#scrollLatestIntoView();
      }, 350);
    }
  }

  #setupNavigationDrawer(shadow) {
    const navDrawer = shadow.querySelector('.nav-drawer');
    const navDrawerToggle = shadow.querySelector('.nav-drawer-toggle');
    const navDrawerClose = shadow.querySelector('.nav-drawer-close');
    const navDrawerOverlay = shadow.querySelector('.nav-drawer-overlay');

    if (!navDrawer || !navDrawerToggle) {
      // No navigation drawer on this page
      return;
    }

    // Get all focusable elements in drawer for focus trap
    const getFocusableElements = () => {
      return navDrawer.querySelectorAll(
        'a[href], button:not([disabled]), details, [tabindex]:not([tabindex="-1"])'
      );
    };

    // Toggle drawer (no localStorage - sidebar closes on navigation)
    const toggleDrawer = () => {
      const isOpen = navDrawer.hasAttribute('open');
      if (isOpen) {
        navDrawer.removeAttribute('open');
        navDrawer.setAttribute('aria-hidden', 'true');
        navDrawerOverlay?.setAttribute('aria-hidden', 'true');
        navDrawerToggle.setAttribute('aria-expanded', 'false');
      } else {
        navDrawer.setAttribute('open', '');
        navDrawer.setAttribute('aria-hidden', 'false');
        navDrawerOverlay?.setAttribute('aria-hidden', 'false');
        navDrawerToggle.setAttribute('aria-expanded', 'true');
        // Focus first focusable element (close button)
        navDrawerClose?.focus();
      }
    };

    navDrawerToggle.addEventListener('click', toggleDrawer);
    navDrawerClose?.addEventListener('click', toggleDrawer);
    navDrawerOverlay?.addEventListener('click', toggleDrawer);

    // Escape key to close drawer
    shadow.addEventListener('keydown', (e) => {
      if (e.key === 'Escape' && navDrawer.hasAttribute('open')) {
        toggleDrawer();
        navDrawerToggle.focus(); // Return focus to toggle button
      }
    });

    // Basic focus trap - keep focus within drawer when open
    navDrawer.addEventListener('keydown', (e) => {
      if (e.key !== 'Tab' || !navDrawer.hasAttribute('open')) {
        return;
      }

      const focusableElements = Array.from(getFocusableElements());
      if (focusableElements.length === 0) return;

      const firstElement = focusableElements[0];
      const lastElement = focusableElements[focusableElements.length - 1];

      if (e.shiftKey && document.activeElement === firstElement) {
        // Shift+Tab on first element - go to last
        e.preventDefault();
        lastElement.focus();
      } else if (!e.shiftKey && document.activeElement === lastElement) {
        // Tab on last element - go to first
        e.preventDefault();
        firstElement.focus();
      }
    });

    // Mark current page in navigation
    const currentPath = window.location.pathname;
    const navLinks = shadow.querySelectorAll('.nav-demo-link');
    navLinks.forEach(link => {
      const href = link.getAttribute('href');
      if (href === currentPath) {
        link.setAttribute('aria-current', 'page');
        // Open the parent details element
        const details = link.closest('details');
        if (details) {
          details.setAttribute('open', '');
        }
      }
    });
  }

  #setupTabs(shadow) {
    const tabs = shadow.querySelectorAll('.tab');
    const panels = shadow.querySelectorAll('.tab-panel');

    if (tabs.length === 0 || panels.length === 0) {
      // No tabs on this page
      return;
    }

    // Restore active tab from localStorage (default to 'panel-knobs')
    const savedTab = localStorage.getItem('cem-serve-active-tab') || 'panel-knobs';

    // Set initial state based on saved tab
    tabs.forEach(tab => {
      const controlsId = tab.getAttribute('aria-controls');
      if (controlsId === savedTab) {
        tab.setAttribute('aria-selected', 'true');
        shadow.getElementById(controlsId)?.setAttribute('active', '');
      } else {
        tab.setAttribute('aria-selected', 'false');
        shadow.getElementById(controlsId)?.removeAttribute('active');
      }
    });

    // Add click handlers for tab switching
    tabs.forEach(tab => {
      tab.addEventListener('click', () => {
        const targetId = tab.getAttribute('aria-controls');

        // Update all tabs
        tabs.forEach(t => t.setAttribute('aria-selected', 'false'));
        tab.setAttribute('aria-selected', 'true');

        // Update all panels
        panels.forEach(p => p.removeAttribute('active'));
        const targetPanel = shadow.getElementById(targetId);
        targetPanel?.setAttribute('active', '');

        // Save to localStorage
        localStorage.setItem('cem-serve-active-tab', targetId);

        // If switching to logs panel and drawer is open, scroll to bottom
        if (targetId === 'panel-logs' && this.#drawerOpen) {
          this.#scrollLogsToBottom();
        }
      });
    });
  }

  #escapeHtml(text) {
    const div = document.createElement('div');
    div.textContent = text;
    return div.innerHTML;
  }
}
