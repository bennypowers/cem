import '/__cem/elements/pf-v6-button/pf-v6-button.js';
import '/__cem/elements/pf-v6-switch/pf-v6-switch.js';
import '/__cem/elements/pf-v6-text-input/pf-v6-text-input.js';
import '/__cem/elements/pf-v6-select/pf-v6-select.js';
import '/__cem/elements/pf-v6-tab/pf-v6-tab.js';
import '/__cem/elements/pf-v6-tabs/pf-v6-tabs.js';
import '/__cem/elements/pf-v6-navigation/pf-v6-navigation.js';
import '/__cem/elements/pf-v6-nav-list/pf-v6-nav-list.js';
import '/__cem/elements/pf-v6-nav-item/pf-v6-nav-item.js';
import '/__cem/elements/pf-v6-nav-link/pf-v6-nav-link.js';
import '/__cem/elements/pf-v6-nav-group/pf-v6-nav-group.js';
import '/__cem/elements/pf-v6-page/pf-v6-page.js';
import '/__cem/elements/pf-v6-page-main/pf-v6-page-main.js';
import '/__cem/elements/pf-v6-page-sidebar/pf-v6-page-sidebar.js';
import '/__cem/elements/pf-v6-masthead/pf-v6-masthead.js';
import '/__cem/elements/pf-v6-skip-to-content/pf-v6-skip-to-content.js';
import '/__cem/elements/pf-v6-toolbar/pf-v6-toolbar.js';
import '/__cem/elements/pf-v6-toolbar-group/pf-v6-toolbar-group.js';
import '/__cem/elements/pf-v6-toolbar-item/pf-v6-toolbar-item.js';
import '/__cem/elements/pf-v6-toggle-group/pf-v6-toggle-group.js';
import '/__cem/elements/pf-v6-toggle-group-item/pf-v6-toggle-group-item.js';
import '/__cem/elements/pf-v6-modal/pf-v6-modal.js';
import '/__cem/elements/pf-v6-card/pf-v6-card.js';
import '/__cem/elements/pf-v6-label/pf-v6-label.js';
import '/__cem/elements/cem-drawer/cem-drawer.js';
import '/__cem/elements/pf-v6-form/pf-v6-form.js';
import '/__cem/elements/pf-v6-form-group/pf-v6-form-group.js';
import '/__cem/elements/pf-v6-form-label/pf-v6-form-label.js';
import '/__cem/elements/pf-v6-popover/pf-v6-popover.js';
import '/__cem/elements/cem-serve-demo/cem-serve-demo.js';
import '/__cem/elements/cem-serve-knobs/cem-serve-knobs.js';
import '/__cem/elements/cem-serve-knob-group/cem-serve-knob-group.js';
import '/__cem/elements/cem-color-scheme-toggle/cem-color-scheme-toggle.js';
import '/__cem/elements/cem-connection-status/cem-connection-status.js';
import '/__cem/elements/cem-reconnection-content/cem-reconnection-content.js';
import '/__cem/elements/cem-transform-error-overlay/cem-transform-error-overlay.js';

import { CemElement } from '/__cem/cem-element.js';
import { CEMReloadClient } from '/__cem/websocket-client.js';

/**
 * CEM Serve Chrome - Main demo wrapper component
 *
 * @customElement cem-serve-chrome
 */
export class CemServeChrome extends CemElement {
  static is = 'cem-serve-chrome';
  static observedAttributes = ['knobs', 'tag-name'];

  #$ = (selector) => this.shadowRoot.querySelector(selector);
  #$$ = (selector) => this.shadowRoot.querySelectorAll(selector);

  #logContainer = null;
  #drawerOpen = false;
  #initialLogsFetched = false;
  #isInitialLoad = true;
  #debugData = null;
  // Track ws connection state
  #hasConnected = false;


  #wsClient = new CEMReloadClient({
    jitterMax: 1000,
    overlayThreshold: 15,
    badgeFadeDelay: 2000,
    callbacks: {
      onOpen: () => {
        console.debug('[cem-serve] WebSocket connected');
        // Only show "connected" toast if this is a reconnection
        if (this.#hasConnected) {
          this.#$('#connection-status')?.show('connected', 'Connected', { fadeDelay: 2000 });
        }
        this.#hasConnected = true;
        this.#$('#reconnection-modal')?.close();
      },
      onClose: () => {
        console.debug('[cem-serve] Connection closed');
      },
      onError: (errorData) => {
        // Handle both WebSocket errors and server error messages
        if (errorData?.title && errorData?.message) {
          // Server error message
          console.error('[cem-serve] Server error:', errorData);
          this.#$('#error-overlay')?.show(errorData.title, errorData.message, errorData.file);
        } else {
          // WebSocket connection error
          console.error('[cem-serve] WebSocket error:', errorData);
        }
      },
      onReconnecting: ({ attempt, delay }) => {
        console.debug(`[cem-serve] Reconnecting in ${Math.ceil(delay/1000)}s (attempt #${attempt})...`);
        this.#$('#connection-status')?.show('reconnecting', `Reconnecting (attempt #${attempt})...`);

        // Show modal after threshold
        if (attempt >= 15) {
          this.#$('#reconnection-modal')?.showModal();
          this.#$('#reconnection-content')?.updateRetryInfo(attempt, delay);
        }
      },
      onReload: (data) => {
        console.debug('[cem-serve] Reloading page:', data.reason, data.files);
        // Hide error overlay on reload (error was fixed)
        const errorOverlay = this.#$('#error-overlay');
        if (errorOverlay?.hasAttribute('open')) {
          errorOverlay.hide();
        }
        window.location.reload();
      },
      onShutdown: () => {
        console.log('[cem-serve] Server shutting down gracefully');
        this.#$('#connection-status')?.show('reconnecting', 'Server restarting...');
        this.#$('#reconnection-modal')?.showModal();
        this.#$('#reconnection-content')?.updateRetryInfo(30, 30000);
      },
      onLogs: (logs) => {
        // Dispatch logs event for log container
        window.dispatchEvent(new CustomEvent('cem:logs', {
          detail: { logs },
          bubbles: true,
          composed: true
        }));
      }
    }
  });

  get demo() { return this.querySelector('cem-serve-demo'); }

  get knobs() { return this.getAttribute('knobs') || ''; }

  get tagName() { return this.getAttribute('tag-name') || ''; }

  async afterTemplateLoaded() {
    // Add initializing class to prevent flash during state restoration
    this.classList.add('initializing');

    // Set up debug overlay
    this.#setupDebugOverlay();

    // Set up color scheme toggle
    this.#setupColorSchemeToggle();

    // Set up footer drawer and tabs
    this.#setupFooterDrawer();

    // Listen for server log messages from WebSocket
    this.#setupLogListener();

    // Set up reconnection modal button handlers
    this.#$('#reload-button')?.addEventListener('click', () => {
      window.location.reload();
    });

    this.#$('#retry-button')?.addEventListener('click', () => {
      this.#$('#reconnection-modal')?.close();
      this.#wsClient.retry();
    });

    // Initialize WebSocket connection
    this.#wsClient.init();

    // Remove initializing class after all state is restored (prevents flash)
    // Wait for browser to apply DOM changes before revealing tabs
    requestAnimationFrame(() => {
      this.classList.remove('initializing');
    });

    console.debug('[cem-serve-chrome] Demo chrome initialized for', this.tagName);
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
    const demoUrlsContainer = this.#$('#demo-urls-container');
    const importMapEl = this.#$('#debug-importmap');

    if (versionEl) versionEl.textContent = data.version || '-';
    if (osEl) osEl.textContent = data.os || '-';
    if (watchDirEl) watchDirEl.textContent = data.watchDir || '-';
    if (manifestSizeEl) manifestSizeEl.textContent = data.manifestSize || '-';
    if (demoCountEl) demoCountEl.textContent = data.demoCount || '0';

    // Populate demo URLs section
    if (demoUrlsContainer) {
      this.#populateDemoUrls(demoUrlsContainer, data.demos);
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

  #populateDemoUrls(container, demos) {
    if (!demos?.length) {
      container.innerHTML = '<p>No demos found in manifest</p>';
      return;
    }

    const currentTagName = this.getAttribute('tag-name') || '';
    const isOnDemoPage = !!currentTagName;

    if (isOnDemoPage) {
      // On demo page - show only current demo, not in details
      const currentDemo = demos.find(demo => demo.tagName === currentTagName);
      if (!currentDemo) {
        container.innerHTML = '<p>Current demo not found in manifest</p>';
        return;
      }

      container.innerHTML = `
        <h3>Demo Information</h3>
        <dl class="pf-v6-c-description-list pf-m-horizontal pf-m-compact">
          <div class="pf-v6-c-description-list__group">
            <dt class="pf-v6-c-description-list__term">Tag Name</dt>
            <dd class="pf-v6-c-description-list__description">${currentDemo.tagName}</dd>
          </div>
          ${currentDemo.description ? `
          <div class="pf-v6-c-description-list__group">
            <dt class="pf-v6-c-description-list__term">Description</dt>
            <dd class="pf-v6-c-description-list__description">${currentDemo.description}</dd>
          </div>
          ` : ''}
          ${currentDemo.filepath ? `
          <div class="pf-v6-c-description-list__group">
            <dt class="pf-v6-c-description-list__term">File Path</dt>
            <dd class="pf-v6-c-description-list__description">${currentDemo.filepath}</dd>
          </div>
          ` : ''}
          <div class="pf-v6-c-description-list__group">
            <dt class="pf-v6-c-description-list__term">Canonical URL</dt>
            <dd class="pf-v6-c-description-list__description">${currentDemo.canonicalURL}</dd>
          </div>
          <div class="pf-v6-c-description-list__group">
            <dt class="pf-v6-c-description-list__term">Local Route</dt>
            <dd class="pf-v6-c-description-list__description">${currentDemo.localRoute}</dd>
          </div>
        </dl>
      `;
    } else {
      // On index page - show all demos in details with description list
      const demoGroups = demos.map(demo => `
        <div class="pf-v6-c-description-list__group">
          <dt class="pf-v6-c-description-list__term">${demo.tagName}</dt>
          <dd class="pf-v6-c-description-list__description">
            ${demo.description || '(no description)'}<br>
            ${demo.filepath ? `<small>File: ${demo.filepath}</small><br>` : ''}
            <small>Canonical: ${demo.canonicalURL}</small><br>
            <small>Local: ${demo.localRoute}</small>
          </dd>
        </div>
      `).join('');

      container.innerHTML = `
        <details id="debug-demos-details">
          <summary>Show Demo URLs from Manifest (${demos.length})</summary>
          <dl class="pf-v6-c-description-list pf-m-horizontal pf-m-compact">
            ${demoGroups}
          </dl>
        </details>
      `;
    }
  }

  #setupLogListener() {
    // Set up log container
    this.#logContainer = this.#$('#log-container');

    // Listen for server log messages from the WebSocket client
    // The websocket-client.js dispatches 'cem:logs' events with server logs
    window.addEventListener('cem:logs', (event) => {
      const logs = event.detail?.logs || event.logs;
      if (logs) {
        this.#renderLogs(logs);
      }
    });
  }

  #setupDebugOverlay() {
    // Set up debug modal
    const debugButton = this.#$('#debug-info');
    const debugModal = this.#$('#debug-modal');
    const debugClose = this.#$('.debug-close');
    const debugCopy = this.#$('.debug-copy');

    if (debugButton && debugModal) {
      debugButton.addEventListener('click', () => {
        this.#fetchDebugInfo();
        debugModal.showModal();
      });

      debugClose?.addEventListener('click', () => {
        debugModal.close();
      });

      debugCopy?.addEventListener('click', () => {
        this.#copyDebugInfo();
      });
    }
  }

  #setupFooterDrawer() {
    // Get the drawer and tabs components
    const drawer = this.#$('cem-drawer');
    const tabs = this.#$('pf-v6-tabs');

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

    // Restore tabs state from localStorage (default to first tab)
    const savedTab = localStorage.getItem('cem-serve-active-tab');
    if (savedTab) {
      tabs.value = savedTab;
    }

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
    this.#$$('#debug-modal dl dt').forEach(dt => {
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
      const labelText = this.#getLogBadge(log.type);
      const labelAttrs = this.#getLogLabelAttrs(log.type);

      return `<div class="log-entry ${log.type}" data-log-id="${log.date}">
        <pf-v6-label ${labelAttrs} compact>${labelText}</pf-v6-label>
        <time class="log-time" datetime="${log.date}">${time}</time>
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
      case 'info': return 'Info';
      case 'warning': return 'Warn';
      case 'error': return 'Error';
      case 'debug': return 'Debug';
      default: return type.toUpperCase();
    }
  }

  #getLogLabelAttrs(type) {
    switch (type) {
      case 'info': return 'status="info"';
      case 'warning': return 'status="warning"';
      case 'error': return 'status="danger"';
      case 'debug': return 'color="purple"';
      default: return 'color="grey"';
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

  #setupColorSchemeToggle() {
    const toggleGroup = this.#$('.color-scheme-toggle');
    if (!toggleGroup) return;

    // Storage access gatekeeper - localStorage can throw in Safari private mode
    const getStorageItem = (key, defaultValue) => {
      try {
        return localStorage.getItem(key) ?? defaultValue;
      } catch (e) {
        return defaultValue;
      }
    };

    const setStorageItem = (key, value) => {
      try {
        localStorage.setItem(key, value);
      } catch (e) {
        // Storage unavailable (private mode), silently continue
      }
    };

    // Restore saved color scheme preference and mark item as selected
    const saved = getStorageItem('cem-serve-color-scheme', 'system');
    const items = toggleGroup.querySelectorAll('pf-v6-toggle-group-item');
    items.forEach(item => {
      if (item.getAttribute('value') === saved) {
        item.setAttribute('selected', '');
      }
    });
    this.#applyColorScheme(saved);

    // Listen for toggle group changes
    toggleGroup.addEventListener('pf-v6-toggle-group-change', (e) => {
      const scheme = e.value;
      this.#applyColorScheme(scheme);
      setStorageItem('cem-serve-color-scheme', scheme);
    });
  }

  #applyColorScheme(scheme) {
    const body = document.body;

    switch (scheme) {
      case 'light':
        body.style.colorScheme = 'light';
        break;
      case 'dark':
        body.style.colorScheme = 'dark';
        break;
      case 'system':
      default:
        body.style.colorScheme = 'light dark';
        break;
    }
  }

  static {
    customElements.define(this.is, this);
  }
}
