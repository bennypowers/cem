import '/__cem/elements/cem-color-scheme-toggle/cem-color-scheme-toggle.js';
import '/__cem/elements/cem-connection-status/cem-connection-status.js';
import '/__cem/elements/cem-drawer/cem-drawer.js';
import '/__cem/elements/cem-manifest-browser/cem-manifest-browser.js';
import '/__cem/elements/cem-reconnection-content/cem-reconnection-content.js';
import '/__cem/elements/cem-serve-demo/cem-serve-demo.js';
import '/__cem/elements/cem-serve-knob-group/cem-serve-knob-group.js';
import '/__cem/elements/cem-serve-knobs/cem-serve-knobs.js';
import '/__cem/elements/cem-transform-error-overlay/cem-transform-error-overlay.js';
import '/__cem/elements/pf-v6-button/pf-v6-button.js';
import '/__cem/elements/pf-v6-card/pf-v6-card.js';
import '/__cem/elements/pf-v6-badge/pf-v6-badge.js';
import '/__cem/elements/pf-v6-label/pf-v6-label.js';
import '/__cem/elements/pf-v6-masthead/pf-v6-masthead.js';
import '/__cem/elements/pf-v6-modal/pf-v6-modal.js';
import '/__cem/elements/pf-v6-nav-group/pf-v6-nav-group.js';
import '/__cem/elements/pf-v6-nav-item/pf-v6-nav-item.js';
import '/__cem/elements/pf-v6-nav-link/pf-v6-nav-link.js';
import '/__cem/elements/pf-v6-nav-list/pf-v6-nav-list.js';
import '/__cem/elements/pf-v6-navigation/pf-v6-navigation.js';
import '/__cem/elements/pf-v6-page-main/pf-v6-page-main.js';
import '/__cem/elements/pf-v6-page-sidebar/pf-v6-page-sidebar.js';
import '/__cem/elements/pf-v6-page/pf-v6-page.js';
import '/__cem/elements/pf-v6-popover/pf-v6-popover.js';
import '/__cem/elements/pf-v6-select/pf-v6-select.js';
import '/__cem/elements/pf-v6-skip-to-content/pf-v6-skip-to-content.js';
import '/__cem/elements/pf-v6-switch/pf-v6-switch.js';
import '/__cem/elements/pf-v6-tab/pf-v6-tab.js';
import '/__cem/elements/pf-v6-tabs/pf-v6-tabs.js';
import '/__cem/elements/pf-v6-text-input-group/pf-v6-text-input-group.js';
import '/__cem/elements/pf-v6-text-input/pf-v6-text-input.js';
import '/__cem/elements/pf-v6-toggle-group-item/pf-v6-toggle-group-item.js';
import '/__cem/elements/pf-v6-toggle-group/pf-v6-toggle-group.js';
import '/__cem/elements/pf-v6-toolbar-group/pf-v6-toolbar-group.js';
import '/__cem/elements/pf-v6-toolbar-item/pf-v6-toolbar-item.js';
import '/__cem/elements/pf-v6-toolbar/pf-v6-toolbar.js';
import '/__cem/elements/pf-v6-tree-item/pf-v6-tree-item.js';
import '/__cem/elements/pf-v6-tree-view/pf-v6-tree-view.js';

import { CemElement } from '/__cem/cem-element.js';
import { CEMReloadClient } from '/__cem/websocket-client.js';

/**
 * Custom event fired when logs are received
 */
export class CemLogsEvent extends Event {
  constructor(logs) {
    super('cem:logs', { bubbles: true, composed: true });
    this.logs = logs;
  }
}

/**
 * CEM Serve Chrome - Main demo wrapper component
 *
 * @customElement cem-serve-chrome
 */
export class CemServeChrome extends CemElement {
  static is = 'cem-serve-chrome';

  static observedAttributes = [
    'knobs',
    'primary-tag-name',
    'demo-title',
    'package-name',
    'canonical-url',
    'source-url',
  ];

  // Static templates for demo URL display
  static #demoInfoTemplate = document.createElement('template');
  static #demoGroupTemplate = document.createElement('template');
  static #demoListTemplate = document.createElement('template');
  static #logEntryTemplate = document.createElement('template');
  static {
    this.#demoInfoTemplate.innerHTML = `
      <h3>Demo Information</h3>
      <dl class="pf-v6-c-description-list pf-m-horizontal pf-m-compact">
        <div class="pf-v6-c-description-list__group">
          <dt class="pf-v6-c-description-list__term">Tag Name</dt>
          <dd class="pf-v6-c-description-list__description" data-field="tagName"></dd>
        </div>
        <div class="pf-v6-c-description-list__group" data-field-group="description">
          <dt class="pf-v6-c-description-list__term">Description</dt>
          <dd class="pf-v6-c-description-list__description" data-field="description"></dd>
        </div>
        <div class="pf-v6-c-description-list__group" data-field-group="filepath">
          <dt class="pf-v6-c-description-list__term">File Path</dt>
          <dd class="pf-v6-c-description-list__description" data-field="filepath"></dd>
        </div>
        <div class="pf-v6-c-description-list__group">
          <dt class="pf-v6-c-description-list__term">Canonical URL</dt>
          <dd class="pf-v6-c-description-list__description" data-field="canonicalURL"></dd>
        </div>
        <div class="pf-v6-c-description-list__group">
          <dt class="pf-v6-c-description-list__term">Local Route</dt>
          <dd class="pf-v6-c-description-list__description" data-field="localRoute"></dd>
        </div>
      </dl>`;
    this.#demoGroupTemplate.innerHTML = `
      <div class="pf-v6-c-description-list__group">
        <dt class="pf-v6-c-description-list__term" data-field="tagName"></dt>
        <dd class="pf-v6-c-description-list__description">
          <span data-field="description"></span><br>
          <small data-field-group="filepath">File: <span data-field="filepath"></span></small>
          <small>Canonical: <span data-field="canonicalURL"></span></small><br>
          <small>Local: <span data-field="localRoute"></span></small>
        </dd>
      </div>
    `;
    this.#demoListTemplate.innerHTML = `
      <details id="debug-demos-details">
        <summary data-field="summary"></summary>
        <dl class="pf-v6-c-description-list pf-m-horizontal pf-m-compact" data-container="groups">
        </dl>
      </details>
    `;
    this.#logEntryTemplate.innerHTML = `
      <div class="log-entry" data-field="container">
        <pf-v6-label compact data-field="label"></pf-v6-label>
        <time class="log-time" data-field="time"></time>
        <span class="log-message" data-field="message"></span>
      </div>
    `;
  }

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
        window.dispatchEvent(new CemLogsEvent(logs));
      }
    }
  });

  get demo() { return this.querySelector('cem-serve-demo'); }

  get knobs() { return this.getAttribute('knobs') || ''; }

  get primaryTagName() { return this.getAttribute('primary-tag-name') || ''; }

  get demoTitle() { return this.getAttribute('demo-title') || ''; }

  get packageName() { return this.getAttribute('package-name') || ''; }

  get canonicalURL() { return this.getAttribute('canonical-url') || ''; }

  get sourceURL() { return this.getAttribute('source-url') || ''; }

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

    // Set up knob event coordination
    this.#setupKnobCoordination();

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

    console.debug('[cem-serve-chrome] Demo chrome initialized for', this.primaryTagName);
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
      container.textContent = 'No demos found in manifest';
      return;
    }

    const currentTagName = this.primaryTagName || '';
    const isOnDemoPage = !!currentTagName;

    if (isOnDemoPage) {
      // On demo page - show only current demo, not in details
      const currentDemo = demos.find(demo => demo.tagName === currentTagName);
      if (!currentDemo) {
        container.textContent = 'Current demo not found in manifest';
        return;
      }

      const fragment = CemServeChrome.#demoInfoTemplate.content.cloneNode(true);

      // Fill in required fields
      fragment.querySelector('[data-field="tagName"]').textContent = currentDemo.tagName;
      fragment.querySelector('[data-field="canonicalURL"]').textContent = currentDemo.canonicalURL;
      fragment.querySelector('[data-field="localRoute"]').textContent = currentDemo.localRoute;

      // Conditionally show/hide optional fields
      const descriptionGroup = fragment.querySelector('[data-field-group="description"]');
      if (currentDemo.description) {
        fragment.querySelector('[data-field="description"]').textContent = currentDemo.description;
      } else {
        descriptionGroup.remove();
      }

      const filepathGroup = fragment.querySelector('[data-field-group="filepath"]');
      if (currentDemo.filepath) {
        fragment.querySelector('[data-field="filepath"]').textContent = currentDemo.filepath;
      } else {
        filepathGroup.remove();
      }

      container.replaceChildren(fragment);
    } else {
      // On index page - show all demos in details with description list
      const listFragment = CemServeChrome.#demoListTemplate.content.cloneNode(true);

      listFragment.querySelector('[data-field="summary"]').textContent =
        `Show Demo URLs from Manifest (${demos.length})`;

      const groupsContainer = listFragment.querySelector('[data-container="groups"]');

      for (const demo of demos) {
        const groupFragment = CemServeChrome.#demoGroupTemplate.content.cloneNode(true);

        groupFragment.querySelector('[data-field="tagName"]').textContent = demo.tagName;
        groupFragment.querySelector('[data-field="description"]').textContent =
          demo.description || '(no description)';
        groupFragment.querySelector('[data-field="canonicalURL"]').textContent = demo.canonicalURL;
        groupFragment.querySelector('[data-field="localRoute"]').textContent = demo.localRoute;

        const filepathGroup = groupFragment.querySelector('[data-field-group="filepath"]');
        if (demo.filepath) {
          groupFragment.querySelector('[data-field="filepath"]').textContent = demo.filepath;
        } else {
          filepathGroup.remove();
        }

        groupsContainer.appendChild(groupFragment);
      }

      container.replaceChildren(listFragment);
    }
  }

  #logsFilterValue = '';

  #setupLogListener() {
    // Set up log container
    this.#logContainer = this.#$('#log-container');

    // Set up filter input
    this.#$('#logs-filter')?.addEventListener('input', (e) => {
      this.#filterLogs(e.target.value);
    });

    // Set up copy logs button
    this.#$('#copy-logs')?.addEventListener('click', () => {
      this.#copyLogs();
    });

    // Listen for server log messages from the WebSocket client
    // The websocket-client.js dispatches 'cem:logs' events with server logs
    window.addEventListener('cem:logs', ({ logs }) => {
      if (logs) {
        this.#renderLogs(logs);
      }
    });
  }

  #filterLogs(query) {
    this.#logsFilterValue = query.toLowerCase();
    
    if (!this.#logContainer) return;

    for (const entry of this.#logContainer.children) {
      const text = entry.textContent.toLowerCase();
      if (this.#logsFilterValue && !text.includes(this.#logsFilterValue)) {
        entry.setAttribute('hidden', '');
      } else {
        entry.removeAttribute('hidden');
      }
    }
  }

  async #copyLogs() {
    if (!this.#logContainer) return;

    // Filter visible logs if there is a filter
    const logs = Array.from(this.#logContainer.children)
      .filter(entry => !entry.hasAttribute('hidden'))
      .map(entry => {
      const type = entry.querySelector('[data-field="label"]')?.textContent?.trim() || 'INFO';
      const time = entry.querySelector('[data-field="time"]')?.textContent?.trim() || '';
      const message = entry.querySelector('[data-field="message"]')?.textContent?.trim() || '';
      return `[${type}] ${time} ${message}`;
    }).join('\n');

    if (!logs) return;

    try {
      await navigator.clipboard.writeText(logs);
      const btn = this.#$('#copy-logs');
      if (btn) {
        const textNode = Array.from(btn.childNodes).find(n => n.nodeType === Node.TEXT_NODE && n.textContent.trim().length > 0);
        if (textNode) {
          const original = textNode.textContent;
          textNode.textContent = '\n              Copied!\n            ';
          setTimeout(() => {
            textNode.textContent = original;
          }, 2000);
        }
      }
    } catch (err) {
      console.error('[cem-serve-chrome] Failed to copy logs:', err);
    }
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

      debugClose?.addEventListener('click', () => debugModal.close());

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

    // Restore drawer state from localStorage
    const savedDrawerOpen = getStorageItem('cem-serve-drawer-open', null);
    if (savedDrawerOpen === 'true') {
      drawer.open = true;
    }

    // Restore drawer height from localStorage
    const savedDrawerHeight = getStorageItem('cem-serve-drawer-height', null);
    if (savedDrawerHeight && drawer.open) {
      const content = drawer.shadowRoot.getElementById('content');
      if (content) {
        content.style.height = `${savedDrawerHeight}px`;
      }
    }

    // Restore tabs state from localStorage (default to first tab)
    const savedTab = getStorageItem('cem-serve-active-tab', null);
    if (savedTab) {
      tabs.value = savedTab;
    }

    // Listen for drawer changes and persist to localStorage
    drawer.addEventListener('change', (e) => {
      setStorageItem('cem-serve-drawer-open', String(e.open));
      this.#drawerOpen = e.open;

      // Scroll logs when drawer opens
      if (e.open) {
        this.#scrollLogsToBottom();
      }
    });

    // Listen for drawer resize and persist to localStorage
    drawer.addEventListener('resize', (e) => {
      setStorageItem('cem-serve-drawer-height', String(e.height));
    });

    // Listen for tab changes and persist to localStorage
    tabs.addEventListener('change', (e) => {
      setStorageItem('cem-serve-active-tab', e.value);

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

  async #copyDebugInfo() {
    // Collect all debug info
    const info = Array.from(this.#$$('#debug-modal dl dt'), dt => {
      const dd = dt.nextElementSibling;
      if (dd && dd.tagName === 'DD') {
        return `${dt.textContent}: ${dd.textContent}`;
      }
    }).join('\n');

    // Include import map if available
    let importMapSection = '';
    if (this.#debugData?.importMapJSON) {
      importMapSection = `\n${'='.repeat(40)}\nImport Map:\n${'='.repeat(40)}\n${this.#debugData.importMapJSON}\n`;
    }

    const debugText = `CEM Serve Debug Information
${'='.repeat(40)}
${info}${importMapSection}
${'='.repeat(40)}
Generated: ${new Date().toISOString()}`;

    try {
      await navigator.clipboard.writeText(debugText);
      const copyButton = this.#$('.debug-copy');
      if (copyButton) {
        const originalText = copyButton.textContent;
        copyButton.textContent = 'Copied!';
        setTimeout(() => {
          copyButton.textContent = originalText;
        }, 2000);
      }
    } catch (err) {
      console.error('[cem-serve-chrome] Failed to copy debug info:', err);
    }
  }

  #renderLogs(logs) {
    if (!this.#logContainer) return;

    const logElements = logs.map(log => {
      // Log is now structured: { type: 'info'|'warning'|'error'|'debug', date: ISO8601, message: string }
      const fragment = CemServeChrome.#logEntryTemplate.content.cloneNode(true);

      const date = new Date(log.date);
      const time = date.toLocaleTimeString();

      // Get container and add type class and data attribute
      const container = fragment.querySelector('[data-field="container"]');
      container.classList.add(log.type);
      container.setAttribute('data-log-id', log.date);

      // Apply current filter
      if (this.#logsFilterValue) {
        // Construct text content for checking
        // Note: we're checking against raw message here, which matches what's displayed
        // We also check against type label implicitly in #filterLogs, but here we can just check message + type
        const typeLabel = this.#getLogBadge(log.type);
        const searchText = `${typeLabel} ${time} ${log.message}`.toLowerCase();
        
        if (!searchText.includes(this.#logsFilterValue)) {
          container.setAttribute('hidden', '');
        }
      }

      // Set label text and attributes based on type
      const label = fragment.querySelector('[data-field="label"]');
      label.textContent = this.#getLogBadge(log.type);
      this.#applyLogLabelAttrs(label, log.type);

      // Set time
      const timeEl = fragment.querySelector('[data-field="time"]');
      timeEl.setAttribute('datetime', log.date);
      timeEl.textContent = time;

      // Set message (textContent automatically escapes)
      fragment.querySelector('[data-field="message"]').textContent = log.message;

      return fragment;
    });

    // Initial load: replace all logs (from fetch on page load - container is empty)
    // Individual logs: append to existing logs (from WebSocket stream - container has content)
    if (!this.#initialLogsFetched) {
      // First batch of logs from initial fetch - replace
      this.#logContainer.replaceChildren(...logElements);
      this.#initialLogsFetched = true;

      // Scroll latest log into view after initial logs are rendered (if drawer is open)
      if (this.#drawerOpen) {
        this.#scrollLatestIntoView();
      }
    } else {
      // Individual log from WebSocket or subsequent updates - append
      this.#logContainer.append(...logElements);

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

  #applyLogLabelAttrs(label, type) {
    switch (type) {
      case 'info':
        label.setAttribute('status', 'info');
        break;
      case 'warning':
        label.setAttribute('status', 'warning');
        break;
      case 'error':
        label.setAttribute('status', 'danger');
        break;
      case 'debug':
        label.setAttribute('color', 'purple');
        break;
      default:
        label.setAttribute('color', 'grey');
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

  #setupKnobCoordination() {
    // Listen for knob events from cem-serve-knob-group elements
    this.addEventListener('knob:attribute-change', this.#handleKnobChange);
    this.addEventListener('knob:property-change', this.#handleKnobChange);
    this.addEventListener('knob:css-property-change', this.#handleKnobChange);
  }

  #handleKnobChange = (event) => {
    // Extract targeting info from event path
    const target = this.#getKnobTarget(event);
    if (!target) {
      console.warn('[cem-serve-chrome] Could not find knob target info in event path');
      return;
    }

    const { tagName, instanceIndex } = target;

    // Find the demo element
    const demo = this.demo;
    if (!demo) {
      console.warn('[cem-serve-chrome] No demo element found');
      return;
    }

    // Determine knob type from event type
    const knobType = this.#getKnobTypeFromEvent(event);

    // Delegate to demo
    const success = demo.applyKnobChange(
      knobType,
      event.name,
      event.value,
      tagName,
      instanceIndex
    );

    if (!success) {
      console.warn('[cem-serve-chrome] Failed to apply knob change:', {
        type: knobType,
        name: event.name,
        tagName,
        instanceIndex
      });
    }
  };

  /**
   * Extract target element info from knob event by traversing composed path
   * to find the pf-v6-card with data-tag-name and data-instance-index
   */
  #getKnobTarget(event) {
    const defaultTagName = this.primaryTagName || '';

    if (event.composedPath) {
      for (const element of event.composedPath()) {
        if (!(element instanceof Element)) continue;

        // Look for element with data-is-element-knob marker
        if (element.dataset?.isElementKnob === 'true') {
          const tagName = element.dataset.tagName || defaultTagName;
          let instanceIndex = Number.parseInt(element.dataset.instanceIndex ?? '', 10);
          if (Number.isNaN(instanceIndex)) instanceIndex = 0;
          return { tagName, instanceIndex };
        }
      }
    }

    return { tagName: defaultTagName, instanceIndex: 0 };
  }

  #getKnobTypeFromEvent(event) {
    switch (event.type) {
      case 'knob:attribute-change':
        return 'attribute';
      case 'knob:property-change':
        return 'property';
      case 'knob:css-property-change':
        return 'css-property';
      default:
        return 'unknown';
    }
  }

  disconnectedCallback() {
    // Clean up knob listeners
    this.removeEventListener('knob:attribute-change', this.#handleKnobChange);
    this.removeEventListener('knob:property-change', this.#handleKnobChange);
    this.removeEventListener('knob:css-property-change', this.#handleKnobChange);
  }

  static {
    customElements.define(this.is, this);
  }
}
