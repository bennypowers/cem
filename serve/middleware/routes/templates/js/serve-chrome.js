// CEM Serve Chrome - Main demo wrapper component

import {
  KnobAttributeChangeEvent,
  KnobPropertyChangeEvent,
  KnobCSSPropertyChangeEvent
} from './knob-events.js';

class CemServeChrome extends HTMLElement {
  static is = 'cem-serve-chrome'
  static { customElements.define(this.is, this); }

  #logContainer = null;
  #drawerOpen = false;
  #initialLogsFetched = false;
  #isInitialLoad = true;
  #debugData = null;

  #$ = (selector) => this.shadowRoot.querySelector(selector);
  #$$ = (selector) => this.shadowRoot.querySelectorAll(selector);

  get demo() { return this.#$('cem-serve-demo'); }

  get knobs() { return this.getAttribute('knobs') || ''; }

  get tagName() { return this.getAttribute('tag-name') || ''; }

  connectedCallback() {
    if (!this.shadowRoot) return;

    // Add initializing class to prevent flash during state restoration
    this.classList.add('initializing');

    // Set up debug overlay
    this.#setupDebugOverlay();

    // Set up footer drawer and tabs
    this.#setupFooterDrawer();

    // Set up knobs event listeners
    this.#setupKnobs();

    // Listen for server log messages from WebSocket
    this.#setupLogListener();

    // Remove initializing class after all state is restored (prevents flash)
    // Wait for browser to apply DOM changes before revealing tabs
    requestAnimationFrame(() => {
      this.classList.remove('initializing');
    });

    console.log('[cem-serve-chrome] Demo chrome initialized for', this.tagName);
  }

  async #fetchDebugInfo() {
    // Populate browser info immediately
    const browserEl = this.#$('#debug-browser');
    const uaEl = this.#$('#debug-ua');
    if (browserEl) {
      const browser = this.#detectBrowser();
      browserEl.textContent = browser;
    }
    if (uaEl) {
      uaEl.textContent = navigator.userAgent;
    }

    // Fetch server debug info
    const data = await fetch('/__cem/debug')
      .then(res => res.json())
      .catch(err => {
        console.error('[cem-serve-chrome] Failed to fetch debug info:', err);
      });

    if (!data) return;
    const versionEl = this.#$('#debug-version');
    const osEl = this.#$('#debug-os');
    const watchDirEl = this.#$('#debug-watch-dir');
    const manifestSizeEl = this.#$('#debug-manifest-size');
    const demoCountEl = this.#$('#debug-demo-count');
    const demosListEl = this.#$('#debug-demos-list');
    const importMapEl = this.#$('#debug-importmap');

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
    this.#debugData = data;
  }

  #setupLogListener() {
    // Set up log container
    this.#logContainer = this.#$('#log-container');

    // Listen for server log messages from the WebSocket client
    // The websocket-client.js dispatches 'cem:logs' events with server logs
    window.addEventListener('cem:logs', (event) => {
      if (event.logs) {
        this.#renderLogs(event.logs);
      }
    });
  }

  #setupDebugOverlay() {
    // Set up debug overlay
    const debugButton = this.#$('.debug-button');
    const debugOverlay = this.#$('.debug-overlay');
    const debugClose = this.#$('.debug-close');
    const debugCopy = this.#$('.debug-copy');

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
  }

  #setupFooterDrawer() {
    // Get the drawer and tabs components
    const drawer = this.#$('cem-drawer');
    const tabs = this.#$('cem-tabs');

    if (!drawer || !tabs) {
      // Components not found
      return;
    }

    // Restore drawer state from localStorage
    const savedDrawerOpen = localStorage.getItem('cem-serve-drawer-open');
    if (savedDrawerOpen === 'true') {
      drawer.open = true;
    }

    // Restore drawer height from localStorage
    const savedDrawerHeight = localStorage.getItem('cem-serve-drawer-height');
    if (savedDrawerHeight && drawer.open) {
      const content = drawer.shadowRoot.getElementById('content');
      if (content) {
        content.style.height = `${savedDrawerHeight}px`;
      }
    }

    // Restore tabs state from localStorage (default to 'panel-knobs')
    const savedTab = localStorage.getItem('cem-serve-active-tab') || 'panel-knobs';
    tabs.value = savedTab;

    // Listen for drawer changes and persist to localStorage
    drawer.addEventListener('change', (e) => {
      localStorage.setItem('cem-serve-drawer-open', String(e.open));
      this.#drawerOpen = e.open;

      // Scroll logs when drawer opens
      if (e.open) {
        this.#scrollLogsToBottom();
      }
    });

    // Listen for drawer resize and persist to localStorage
    drawer.addEventListener('resize', (e) => {
      localStorage.setItem('cem-serve-drawer-height', String(e.height));
    });

    // Listen for tab changes and persist to localStorage
    tabs.addEventListener('change', (e) => {
      localStorage.setItem('cem-serve-active-tab', e.value);

      // Scroll logs if switching to logs panel and drawer is open
      if (e.value === 'panel-logs' && drawer.open) {
        this.#scrollLogsToBottom();
      }
    });
  }

  #setupKnobs() {
    // Get the knobs container
    const knobsContainer = this.#$('#knobs-container');
    if (!knobsContainer) return;

    // Get all knob inputs
    const knobInputs = knobsContainer.querySelectorAll('[data-knob-category]');

    knobInputs.forEach(knobControl => {
      const category = knobControl.dataset.knobCategory;
      const input = knobControl.querySelector('input, select');
      if (!input) return;

      const knobName = input.dataset.knobName;
      if (!knobName) return;

      // Listen for input changes
      input.addEventListener('input', () => {
        this.#handleKnobChange(category, knobName, input);
      });

      // For select/checkbox, also listen to change event
      if (input.tagName === 'SELECT' || input.type === 'checkbox') {
        input.addEventListener('change', () => {
          this.#handleKnobChange(category, knobName, input);
        });
      }
    });
  }

  #handleKnobChange(category, name, input) {
    const demoElement = this.demo?.querySelector(this.tagName);
    if (!demoElement) {
      console.warn('[cem-serve-chrome] Demo element not found:', this.tagName);
      return;
    }

    let value;
    let event;

    // Get the value based on input type
    if (input.type === 'checkbox') {
      value = input.checked;
    } else if (input.type === 'number') {
      value = input.valueAsNumber;
    } else {
      value = input.value;
    }

    // Handle different knob categories
    switch (category) {
      case 'attributes':
        // Set or remove attribute based on value
        if (input.type === 'checkbox') {
          if (value) {
            demoElement.setAttribute(name, '');
          } else {
            demoElement.removeAttribute(name);
          }
        } else if (value === '') {
          demoElement.removeAttribute(name);
        } else {
          demoElement.setAttribute(name, value);
        }
        event = new KnobAttributeChangeEvent(name, value);
        break;

      case 'properties':
        // Set property value
        demoElement[name] = value;
        event = new KnobPropertyChangeEvent(name, value);
        break;

      case 'css-properties':
        // Set CSS custom property
        demoElement.style.setProperty(name, value);
        event = new KnobCSSPropertyChangeEvent(name, value);
        break;

      default:
        console.warn('[cem-serve-chrome] Unknown knob category:', category);
        return;
    }

    // Dispatch the custom event
    if (event) {
      this.dispatchEvent(event);
    }

    console.log(`[cem-serve-chrome] Knob changed: ${category}/${name} = ${value}`);
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
    // Collect all debug info
    const info = [];
    this.#$$('.debug-panel dl dt').forEach(dt => {
      const dd = dt.nextElementSibling;
      if (dd && dd.tagName === 'DD') {
        info.push(`${dt.textContent}: ${dd.textContent}`);
      }
    });

    // Include import map if available
    let importMapSection = '';
    if (this.#debugData?.importMapJSON) {
      importMapSection = `\n${'='.repeat(40)}\nImport Map:\n${'='.repeat(40)}\n${this._debugData.importMapJSON}\n`;
    }

    const debugText = `CEM Serve Debug Information
${'='.repeat(40)}
${info.join('\n')}${importMapSection}
${'='.repeat(40)}
Generated: ${new Date().toISOString()}`;

    navigator.clipboard.writeText(debugText)
      .then(() => {
        const copyButton = this.#$('.debug-copy');
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

  #escapeHtml(text) {
    const div = document.createElement('div');
    div.textContent = text;
    return div.innerHTML;
  }
}
