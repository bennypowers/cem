import '/__cem/elements/cem-color-scheme-toggle/cem-color-scheme-toggle.js';
import '/__cem/elements/cem-drawer/cem-drawer.js';
import '/__cem/elements/cem-manifest-browser/cem-manifest-browser.js';
import '/__cem/elements/cem-reconnection-content/cem-reconnection-content.js';
import '/__cem/elements/cem-serve-demo/cem-serve-demo.js';
import '/__cem/elements/cem-serve-knob-group/cem-serve-knob-group.js';
import '/__cem/elements/cem-serve-knobs/cem-serve-knobs.js';
import '/__cem/elements/cem-transform-error-overlay/cem-transform-error-overlay.js';
import '/__cem/elements/pf-v6-alert/pf-v6-alert.js';
import '/__cem/elements/pf-v6-alert-group/pf-v6-alert-group.js';
import '/__cem/elements/pf-v6-button/pf-v6-button.js';
import '/__cem/elements/pf-v6-card/pf-v6-card.js';
import '/__cem/elements/pf-v6-badge/pf-v6-badge.js';
import '/__cem/elements/pf-v6-dropdown/pf-v6-dropdown.js';
import '/__cem/elements/pf-v6-expandable-section/pf-v6-expandable-section.js';
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
import { StatePersistence } from '/__cem/state-persistence.js';

/**
 * Custom event fired when logs are received
 */
export class CemLogsEvent extends Event {
  constructor(logs) {
    super('cem:logs', { bubbles: true });
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
    'drawer-open',
    'drawer-height',
    'tabs-selected',
  ];

  // Static templates for demo URL display
  static #demoInfoTemplate = document.createElement('template');
  static #demoGroupTemplate = document.createElement('template');
  static #demoListTemplate = document.createElement('template');
  static #logEntryTemplate = document.createElement('template');
  static #connectionAlertTemplate = document.createElement('template');
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
      <pf-v6-expandable-section id="debug-demos-section"
                                toggle-text="Show Demos Info">
        <dl class="pf-v6-c-description-list pf-m-horizontal pf-m-compact" data-container="groups"></dl>
      </pf-v6-expandable-section>
    `;
    this.#logEntryTemplate.innerHTML = `
      <div class="log-entry" data-field="container">
        <pf-v6-label compact data-field="label"></pf-v6-label>
        <time class="log-time" data-field="time"></time>
        <span class="log-message" data-field="message"></span>
      </div>
    `;
    this.#connectionAlertTemplate.innerHTML = `
      <pf-v6-alert dismissable data-field="alert"></pf-v6-alert>
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
        // Clear any reconnecting/restarting alerts
        this.#clearConnectionAlerts();

        // Only show "connected" toast if this is a reconnection
        if (this.#hasConnected) {
          this.#showConnectionAlert('success', 'Connected');
        }
        this.#hasConnected = true;
        this.#$('#reconnection-modal')?.close();
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
        this.#showConnectionAlert('warning', 'Reconnecting');
        // Show modal after threshold
        if (attempt >= 15) {
          this.#$('#reconnection-modal')?.showModal();
          this.#$('#reconnection-content')?.updateRetryInfo(attempt, delay);
        }
      },
      onReload: () => {
        // Hide error overlay on reload (error was fixed)
        const errorOverlay = this.#$('#error-overlay');
        if (errorOverlay?.hasAttribute('open')) {
          errorOverlay.hide();
        }
        window.location.reload();
      },
      onShutdown: () => {
        this.#showConnectionAlert('info', 'Server Restarting');
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
    // Check if we need to migrate from localStorage
    this.#migrateFromLocalStorageIfNeeded();

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

    // Set up tree state persistence
    this.#setupTreeStatePersistence();

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

    // Pre-cache connection alert template in idle time
    requestIdleCallback(() => {
      CemServeChrome.#connectionAlertTemplate.content.cloneNode(true);
    });
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

      listFragment.querySelector('[data-field="summary"]').toggleText =
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
  #logsFilterDebounceTimer = null;
  #logLevelFilters = new Set(['info', 'warn', 'error', 'debug']);
  #logLevelDropdown = null;

  #setupLogListener() {
    // Set up log container
    this.#logContainer = this.#$('#log-container');

    // Set up filter input with debouncing
    const logsFilter = this.#$('#logs-filter');
    if (logsFilter) {
      logsFilter.addEventListener('input', () => {
        const value = logsFilter.getAttribute('value') || '';

        // Debounce filter - wait 300ms after user stops typing
        clearTimeout(this.#logsFilterDebounceTimer);
        this.#logsFilterDebounceTimer = setTimeout(() => {
          this.#filterLogs(value);
        }, 300);
      });
    }

    // Set up log level filter dropdown
    this.#logLevelDropdown = this.#$('#log-level-filter');
    if (this.#logLevelDropdown) {
      // Load saved filter state and sync checkboxes
      // Use requestAnimationFrame to ensure menu items are fully rendered
      requestAnimationFrame(() => {
        this.#loadLogFilterState();
      });

      // Listen for filter changes
      this.#logLevelDropdown.addEventListener('select', this.#handleLogFilterChange);
    }

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
      // Check text filter
      const text = entry.textContent.toLowerCase();
      const textMatch = !this.#logsFilterValue || text.includes(this.#logsFilterValue);

      // Check log level filter
      const logType = this.#getLogTypeFromEntry(entry);
      const levelMatch = this.#logLevelFilters.has(logType);

      entry.hidden = !(textMatch && levelMatch);
    }
  }

  #getLogTypeFromEntry(entry) {
    // Extract log type from classList (info, warning, error, debug)
    for (const cls of entry.classList) {
      if (['info', 'warning', 'error', 'debug'].includes(cls)) {
        // Map 'warning' to 'warn' for consistency with log.type
        return cls === 'warning' ? 'warn' : cls;
      }
    }
    return 'info'; // Default
  }

  #loadLogFilterState() {
    try {
      const saved = localStorage.getItem('cem-serve-log-filters');
      if (saved) {
        this.#logLevelFilters = new Set(JSON.parse(saved));
      }
    } catch (e) {
      // localStorage unavailable (private mode), use defaults
      console.debug('[cem-serve-chrome] localStorage unavailable, using default log filters');
    }

    // Always sync checkbox states with current filter Set
    this.#syncCheckboxStates();
  }

  #syncCheckboxStates() {
    if (!this.#logLevelDropdown) return;

    const menuItems = this.#logLevelDropdown.querySelectorAll('pf-v6-menu-item');
    menuItems.forEach(item => {
      const value = item.getAttribute('value');
      item.checked = this.#logLevelFilters.has(value);
    });
  }

  #saveLogFilterState() {
    try {
      localStorage.setItem('cem-serve-log-filters',
        JSON.stringify([...this.#logLevelFilters]));
    } catch (e) {
      // localStorage unavailable (private mode), silently continue
    }
  }

  #handleLogFilterChange = (event) => {
    // Event is PfMenuItemSelectEvent with value and checked properties
    const { value, checked } = event;

    if (checked) {
      this.#logLevelFilters.add(value);
    } else {
      this.#logLevelFilters.delete(value);
    }

    this.#saveLogFilterState();
    this.#filterLogs(this.#logsFilterValue);
  };

  async #copyLogs() {
    if (!this.#logContainer) return;

    // Filter visible logs if there is a filter
    const logs = Array.from(this.#logContainer.children)
      .filter(entry => !entry.hidden)
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

    // Track drawer open state for logs scrolling (state already restored in afterTemplateLoaded)
    this.#drawerOpen = drawer.open;

    // Listen for drawer changes and update tracked state
    drawer.addEventListener('change', (e) => {
      this.#drawerOpen = e.open;

      // Persist drawer open state
      StatePersistence.updateState({
        drawer: { open: e.open }
      });

      // Scroll logs when drawer opens
      if (e.open) {
        this.#scrollLogsToBottom();
      }
    });

    // Listen for drawer resize events
    drawer.addEventListener('resize', (e) => {
      // Update drawer-height attribute
      drawer.setAttribute('drawer-height', e.height);

      // Persist drawer height
      StatePersistence.updateState({
        drawer: { height: e.height }
      });
    });

    // Listen for tab changes and persist to cookie
    tabs.addEventListener('change', (e) => {
      StatePersistence.updateState({
        tabs: { selectedIndex: e.selectedIndex }
      });

      // Scroll logs if switching to logs panel (index 2) and drawer is open
      if (e.selectedIndex === 2 && drawer.open) {
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

      // Apply current filters
      // Check text filter
      const typeLabel = this.#getLogBadge(log.type);
      const searchText = `${typeLabel} ${time} ${log.message}`.toLowerCase();
      const textMatch = !this.#logsFilterValue || searchText.includes(this.#logsFilterValue);

      // Check log level filter (map 'warning' to 'warn')
      const logTypeForFilter = log.type === 'warning' ? 'warn' : log.type;
      const levelMatch = this.#logLevelFilters.has(logTypeForFilter);

      if (!(textMatch && levelMatch)) {
        container.setAttribute('hidden', '');
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

  #migrateFromLocalStorageIfNeeded() {
    // Check if localStorage has any values we should migrate
    try {
      const hasLocalStorage =
        localStorage.getItem('cem-serve-color-scheme') !== null ||
        localStorage.getItem('cem-serve-drawer-open') !== null ||
        localStorage.getItem('cem-serve-drawer-height') !== null ||
        localStorage.getItem('cem-serve-active-tab') !== null;

      if (hasLocalStorage) {
        // Check if we've already migrated (look for a migration marker)
        const migrated = localStorage.getItem('cem-serve-migrated-to-cookies');
        if (!migrated) {
          StatePersistence.migrateFromLocalStorage();
          localStorage.setItem('cem-serve-migrated-to-cookies', 'true');

          // Reload to apply migrated state via SSR
          // Use setTimeout to avoid blocking
          setTimeout(() => window.location.reload(), 100);
        }
      }
    } catch (e) {
      // localStorage not available, skip migration
    }
  }

  #setupColorSchemeToggle() {
    const toggleGroup = this.#$('.color-scheme-toggle');
    if (!toggleGroup) return;

    // Get state from cookie (already SSR'd, this is for client-side updates)
    const state = StatePersistence.getState();

    // Apply current scheme
    this.#applyColorScheme(state.colorScheme);

    // Mark correct toggle as selected (SSR should have done this via body style, but ensure it)
    const items = toggleGroup.querySelectorAll('pf-v6-toggle-group-item');
    items.forEach(item => {
      if (item.getAttribute('value') === state.colorScheme) {
        item.setAttribute('selected', '');
      }
    });

    // Listen for toggle group changes and update cookie
    toggleGroup.addEventListener('pf-v6-toggle-group-change', (e) => {
      const scheme = e.value;
      this.#applyColorScheme(scheme);
      StatePersistence.updateState({ colorScheme: scheme });
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

  #setupTreeStatePersistence() {
    // Listen for tree item events and persist state to localStorage
    this.addEventListener('expand', (e) => {
      if (e.target.tagName !== 'PF-V6-TREE-ITEM') return;

      const nodeId = this.#getTreeNodeId(e.target);
      const treeState = StatePersistence.getTreeState();
      if (!treeState.expanded.includes(nodeId)) {
        treeState.expanded.push(nodeId);
        StatePersistence.setTreeState(treeState);
      }
    });

    this.addEventListener('collapse', (e) => {
      if (e.target.tagName !== 'PF-V6-TREE-ITEM') return;

      const nodeId = this.#getTreeNodeId(e.target);
      const treeState = StatePersistence.getTreeState();
      const index = treeState.expanded.indexOf(nodeId);
      if (index > -1) {
        treeState.expanded.splice(index, 1);
        StatePersistence.setTreeState(treeState);
      }
    });

    this.addEventListener('select', (e) => {
      if (e.target.tagName !== 'PF-V6-TREE-ITEM') return;

      const nodeId = this.#getTreeNodeId(e.target);
      StatePersistence.updateTreeState({ selected: nodeId });
    });

    // Apply tree state from localStorage on load
    this.#applyTreeState();
  }

  #applyTreeState() {
    const treeState = StatePersistence.getTreeState();

    // Expand nodes
    for (const nodeId of treeState.expanded) {
      const treeItem = this.#findTreeItemById(nodeId);
      if (treeItem && !treeItem.hasAttribute('expanded')) {
        treeItem.setAttribute('expanded', '');
      }
    }

    // Select node
    if (treeState.selected) {
      const treeItem = this.#findTreeItemById(treeState.selected);
      if (treeItem && !treeItem.hasAttribute('current')) {
        treeItem.setAttribute('current', '');
      }
    }
  }

  #findTreeItemById(nodeId) {
    const parts = nodeId.split(':');
    const [type, modulePath, tagName, name] = parts;

    // Build attribute suffix for tag name and name
    let attrSuffix = '';
    if (tagName) {
      attrSuffix += `[data-tag-name="${CSS.escape(tagName)}"]`;
    }
    if (name) {
      attrSuffix += `[data-name="${CSS.escape(name)}"]`;
    }

    // Build complete selectors with all attributes
    let selector = `pf-v6-tree-item[data-type="${CSS.escape(type)}"]`;
    if (modulePath) {
      const escapedModulePath = CSS.escape(modulePath);
      const escapedType = CSS.escape(type);
      const selector1 = `pf-v6-tree-item[data-type="${escapedType}"][data-module-path="${escapedModulePath}"]${attrSuffix}`;
      const selector2 = `pf-v6-tree-item[data-type="${escapedType}"][data-path="${escapedModulePath}"]${attrSuffix}`;
      selector = `${selector1}, ${selector2}`;
    } else {
      selector += attrSuffix;
    }

    return this.querySelector(selector);
  }

  #getTreeNodeId(treeItem) {
    const type = treeItem.getAttribute('data-type');
    const modulePath = treeItem.getAttribute('data-module-path') || treeItem.getAttribute('data-path');
    const tagName = treeItem.getAttribute('data-tag-name');
    const name = treeItem.getAttribute('data-name');
    const category = treeItem.getAttribute('data-category');

    const parts = [type];
    if (modulePath) parts.push(modulePath);
    if (tagName) parts.push(tagName);
    if (category) {
      parts.push(category);
    } else if (name) {
      parts.push(name);
    }

    return parts.join(':');
  }

  #showConnectionAlert(variant, title) {
    const alertGroup = this.#$('#connection-alerts');
    if (!alertGroup) return;

    // Clone template
    const fragment = CemServeChrome.#connectionAlertTemplate.content.cloneNode(true);
    const alert = fragment.querySelector('[data-field="alert"]');

    // Set variant and title
    alert.setAttribute('variant', variant);
    alert.dataset.connectionAlert = variant;
    alert.textContent = title;

    // Auto-remove timeouts per PatternFly guidelines
    let timeout;
    let isHovered = false;
    let timeoutExpired = false;

    const remove = () => {
      alert.remove();
    };

    const scheduleTimeout = (duration) => {
      timeout = setTimeout(() => {
        timeoutExpired = true;
        // If user is hovering, delay removal
        if (!isHovered) {
          remove();
        }
      }, duration);
    };

    // Hover handling - pause removal while hovering
    alert.addEventListener('mouseenter', () => {
      isHovered = true;
    });

    alert.addEventListener('mouseleave', () => {
      isHovered = false;
      // If timeout already expired while hovering, remove after hover delay
      if (timeoutExpired) {
        setTimeout(remove, 3000); // 3s hover delay per PF
      }
    });

    // Set timeout for all variants
    if (variant === 'success') {
      scheduleTimeout(8000); // 8s for success
    } else if (variant === 'warning') {
      scheduleTimeout(15000); // 15s for warning
    } else if (variant === 'info') {
      scheduleTimeout(15000); // 15s for info
    }

    // Clear any existing alerts of the same type
    const existingAlerts = alertGroup.querySelectorAll(`[data-connection-alert="${variant}"]`);
    existingAlerts.forEach(existing => {
      existing.remove();
    });

    alertGroup.addAlert(alert);
  }

  #clearConnectionAlerts() {
    const alertGroup = this.#$('#connection-alerts');
    if (!alertGroup) return;

    // Clear warning and info alerts (not success - let those auto-remove)
    const alerts = alertGroup.querySelectorAll('[data-connection-alert="warning"], [data-connection-alert="info"]');
    alerts.forEach(alert => {
      alert.remove();
    });
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
