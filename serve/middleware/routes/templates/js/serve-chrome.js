// CEM Serve Chrome - Main demo wrapper component

import {
  KnobAttributeChangeEvent,
  KnobPropertyChangeEvent,
  KnobCSSPropertyChangeEvent
} from '/__cem/knob-events.js';

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

  get demo() { return this.querySelector('cem-serve-demo'); }

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
    // Listen for knob change events from knob custom elements
    // Phase 5b: Events bubble from knob elements through their <details> container
    // The container has data-instance-index and data-tag-name to identify which element instance
    this.addEventListener('knob:attribute-change', (e) => {
      console.log('[cem-serve-chrome] knob:attribute-change event:', {
        target: e.target,
        targetTagName: e.target.tagName,
        composedPath: e.composedPath(),
        knobName: e.name
      });
      const { tagName, instanceIndex } = this.#getKnobTarget(e.target, e);
      this.#handleAttributeKnobChange(e.name, e.value, tagName, instanceIndex);
    });

    this.addEventListener('knob:property-change', (e) => {
      const { tagName, instanceIndex } = this.#getKnobTarget(e.target, e);
      this.#handlePropertyKnobChange(e.name, e.value, tagName, instanceIndex);
    });

    this.addEventListener('knob:css-property-change', (e) => {
      const { tagName, instanceIndex } = this.#getKnobTarget(e.target, e);
      this.#handleCSSPropertyKnobChange(e.name, e.value, tagName, instanceIndex);
    });
  }

  /**
   * Get the target element tag name and instance index from a knob element
   * by finding its parent <details> container.
   * Uses the event's composedPath to traverse through shadow boundaries.
   * Returns {tagName, instanceIndex}
   */
  #getKnobTarget(knobElement, event) {
    // If we have the event, use composedPath to traverse shadow boundaries
    if (event && event.composedPath) {
      const path = event.composedPath();
      console.log('[cem-serve-chrome] Searching composedPath for .knob-group-instance:', path);

      for (const element of path) {
        if (element.classList && element.classList.contains('knob-group-instance')) {
          const tagName = element.dataset.tagName;
          const instanceIndex = parseInt(element.dataset.instanceIndex, 10);
          console.log('[cem-serve-chrome] Found target:', { tagName, instanceIndex }, 'from element:', element);
          return { tagName, instanceIndex };
        }
      }
    }

    // Fallback to closest() for non-shadow cases
    const details = knobElement.closest('.knob-group-instance');
    if (details && details.dataset.instanceIndex !== undefined) {
      const tagName = details.dataset.tagName;
      const instanceIndex = parseInt(details.dataset.instanceIndex, 10);
      console.log('[cem-serve-chrome] Found target via closest():', { tagName, instanceIndex });
      return { tagName, instanceIndex };
    }

    // Fallback - log warning and use chrome's tag name
    console.warn('[cem-serve-chrome] Could not find knob target, falling back to chrome tagName');
    return { tagName: this.tagName, instanceIndex: 0 };
  }

  #handleAttributeKnobChange(name, value, tagName, instanceIndex = 0) {
    if (!this.demo) {
      console.warn('[cem-serve-chrome] Demo wrapper not found');
      return;
    }

    // Phase 5b: Find the Nth element instance of the specified tag
    const element = this.#getElementInstance(tagName, instanceIndex);
    if (!element) {
      console.warn('[cem-serve-chrome] Demo element not found:', tagName, 'at index', instanceIndex);
      return;
    }

    // Directly manipulate the element instead of using selector-based approach
    if (typeof value === 'boolean') {
      if (value) {
        element.setAttribute(name, '');
      } else {
        element.removeAttribute(name);
      }
    } else if (value === '' || value === null || value === undefined) {
      element.removeAttribute(name);
    } else {
      element.setAttribute(name, value);
    }

    console.log(`[cem-serve-chrome] Attribute changed on ${tagName} instance ${instanceIndex}: ${name} = ${value}`);
  }

  #handlePropertyKnobChange(name, value, tagName, instanceIndex = 0) {
    if (!this.demo) {
      console.warn('[cem-serve-chrome] Demo wrapper not found');
      return;
    }

    const element = this.#getElementInstance(tagName, instanceIndex);
    if (!element) {
      console.warn('[cem-serve-chrome] Demo element not found:', tagName, 'at index', instanceIndex);
      return;
    }

    element[name] = value;
    console.log(`[cem-serve-chrome] Property changed on ${tagName} instance ${instanceIndex}: ${name} = ${value}`);
  }

  #handleCSSPropertyKnobChange(name, value, tagName, instanceIndex = 0) {
    if (!this.demo) {
      console.warn('[cem-serve-chrome] Demo wrapper not found');
      return;
    }

    const element = this.#getElementInstance(tagName, instanceIndex);
    if (!element) {
      console.warn('[cem-serve-chrome] Demo element not found:', tagName, 'at index', instanceIndex);
      return;
    }

    const propertyName = name.startsWith('--') ? name : `--${name}`;
    element.style.setProperty(propertyName, value);
    console.log(`[cem-serve-chrome] CSS property changed on ${tagName} instance ${instanceIndex}: ${propertyName} = ${value}`);
  }

  /**
   * Get the Nth instance of the specified element from the demo
   */
  #getElementInstance(tagName, index) {
    if (!this.demo || !tagName) return null;

    const root = this.demo.shadowRoot ?? this.demo;
    const elements = root.querySelectorAll(tagName);

    console.log(`[cem-serve-chrome] Getting element instance ${index} of ${tagName}:`, {
      totalElements: elements.length,
      requestedIndex: index,
      allElements: Array.from(elements).map(el => ({
        id: el.id,
        variant: el.getAttribute('variant'),
        className: el.className,
        element: el
      })),
      selectedElement: elements[index]
    });

    return elements[index] || null;
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
      importMapSection = `\n${'='.repeat(40)}\nImport Map:\n${'='.repeat(40)}\n${this.#debugData.importMapJSON}\n`;
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
