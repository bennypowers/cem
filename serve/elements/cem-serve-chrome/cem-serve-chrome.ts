import { LitElement, html, nothing } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';
import { property } from 'lit/decorators/property.js';
import { state } from 'lit/decorators/state.js';

import styles from './cem-serve-chrome.css' with { type: 'css' };

import '../cem-color-scheme-toggle/cem-color-scheme-toggle.js';
import '../cem-drawer/cem-drawer.js';
import '../cem-health-panel/cem-health-panel.js';
import '../cem-manifest-browser/cem-manifest-browser.js';
import '../cem-reconnection-content/cem-reconnection-content.js';
import '../cem-serve-demo/cem-serve-demo.js';
import '../cem-serve-knob-group/cem-serve-knob-group.js';
import '../cem-serve-knobs/cem-serve-knobs.js';
import '../cem-transform-error-overlay/cem-transform-error-overlay.js';
import '../pf-v6-alert/pf-v6-alert.js';
import '../pf-v6-alert-group/pf-v6-alert-group.js';
import '../pf-v6-button/pf-v6-button.js';
import '../pf-v6-card/pf-v6-card.js';
import '../pf-v6-badge/pf-v6-badge.js';
import '../pf-v6-dropdown/pf-v6-dropdown.js';
import '../pf-v6-expandable-section/pf-v6-expandable-section.js';
import '../pf-v6-label/pf-v6-label.js';
import '../pf-v6-masthead/pf-v6-masthead.js';
import '../pf-v6-modal/pf-v6-modal.js';
import '../pf-v6-nav-group/pf-v6-nav-group.js';
import '../pf-v6-nav-item/pf-v6-nav-item.js';
import '../pf-v6-nav-link/pf-v6-nav-link.js';
import '../pf-v6-nav-list/pf-v6-nav-list.js';
import '../pf-v6-navigation/pf-v6-navigation.js';
import '../pf-v6-page-main/pf-v6-page-main.js';
import '../pf-v6-page-sidebar/pf-v6-page-sidebar.js';
import '../pf-v6-page/pf-v6-page.js';
import '../pf-v6-popover/pf-v6-popover.js';
import '../pf-v6-select/pf-v6-select.js';
import '../pf-v6-skip-to-content/pf-v6-skip-to-content.js';
import '../pf-v6-switch/pf-v6-switch.js';
import '../pf-v6-tab/pf-v6-tab.js';
import '../pf-v6-tabs/pf-v6-tabs.js';
import '../pf-v6-text-input-group/pf-v6-text-input-group.js';
import '../pf-v6-text-input/pf-v6-text-input.js';
import '../pf-v6-toggle-group-item/pf-v6-toggle-group-item.js';
import '../pf-v6-toggle-group/pf-v6-toggle-group.js';
import '../pf-v6-toolbar-group/pf-v6-toolbar-group.js';
import '../pf-v6-toolbar-item/pf-v6-toolbar-item.js';
import '../pf-v6-toolbar/pf-v6-toolbar.js';
import '../pf-v6-tree-item/pf-v6-tree-item.js';
import '../pf-v6-tree-view/pf-v6-tree-view.js';

// @ts-ignore -- plain JS module served at runtime by Go server
import { CEMReloadClient } from '/__cem/websocket-client.js';
// @ts-ignore -- plain JS module served at runtime by Go server
import { StatePersistence } from '/__cem/state-persistence.js';
// @ts-ignore -- plain JS module served at runtime by Go server
import '/__cem/health-badges.js';

interface EventInfo {
  eventNames: Set<string>;
  events: Array<{ name: string; type?: { text: string }; summary?: string; description?: string }>;
}

interface EventRecord {
  id: string;
  timestamp: Date;
  eventName: string;
  tagName: string;
  elementId: string | null;
  elementClass: string | null;
  customProperties: Record<string, unknown>;
  manifestType: string | null;
  summary: string | null;
  description: string | null;
  bubbles: boolean;
  composed: boolean;
  cancelable: boolean;
  defaultPrevented: boolean;
}

interface DebugData {
  version?: string;
  os?: string;
  watchDir?: string;
  manifestSize?: string;
  demoCount?: number;
  demos?: Array<{
    tagName: string;
    description?: string;
    filepath?: string;
    canonicalURL: string;
    localRoute: string;
  }>;
  importMap?: Record<string, unknown>;
  importMapJSON?: string;
}

interface Manifest {
  modules?: Array<{
    declarations?: Array<{
      customElement?: boolean;
      tagName?: string;
      name?: string;
      kind?: string;
      events?: Array<{ name: string; type?: { text: string }; summary?: string; description?: string }>;
    }>;
  }>;
}

/**
 * Custom event fired when logs are received
 */
export class CemLogsEvent extends Event {
  logs: Array<{ type: string; date: string; message: string }>;
  constructor(logs: Array<{ type: string; date: string; message: string }>) {
    super('cem:logs', { bubbles: true });
    this.logs = logs;
  }
}

/**
 * CEM Serve Chrome - Main demo wrapper component
 *
 * @slot - Default slot for demo content
 * @slot navigation - Navigation sidebar content
 * @slot knobs - Knob controls
 * @slot description - Demo description
 * @slot manifest-tree - Manifest tree view
 * @slot manifest-name - Manifest name display
 * @slot manifest-details - Manifest details display
 *
 * @customElement cem-serve-chrome
 */
@customElement('cem-serve-chrome')
export class CemServeChrome extends LitElement {
  static styles = styles;

  // Static templates for DOM cloning (logs, events, debug info)
  static #demoInfoTemplate = (() => {
    const t = document.createElement('template');
    t.innerHTML = `
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
    return t;
  })();

  static #demoGroupTemplate = (() => {
    const t = document.createElement('template');
    t.innerHTML = `
      <div class="pf-v6-c-description-list__group">
        <dt class="pf-v6-c-description-list__term" data-field="tagName"></dt>
        <dd class="pf-v6-c-description-list__description">
          <span data-field="description"></span><br>
          <small data-field-group="filepath">File: <span data-field="filepath"></span></small>
          <small>Canonical: <span data-field="canonicalURL"></span></small><br>
          <small>Local: <span data-field="localRoute"></span></small>
        </dd>
      </div>`;
    return t;
  })();

  static #demoListTemplate = (() => {
    const t = document.createElement('template');
    t.innerHTML = `
      <pf-v6-expandable-section id="debug-demos-section"
                                toggle-text="Show Demos Info">
        <dl class="pf-v6-c-description-list pf-m-horizontal pf-m-compact" data-container="groups"></dl>
      </pf-v6-expandable-section>`;
    return t;
  })();

  static #logEntryTemplate = (() => {
    const t = document.createElement('template');
    t.innerHTML = `
      <div class="log-entry" data-field="container">
        <pf-v6-label compact data-field="label"></pf-v6-label>
        <time class="log-time" data-field="time"></time>
        <span class="log-message" data-field="message"></span>
      </div>`;
    return t;
  })();

  static #eventEntryTemplate = (() => {
    const t = document.createElement('template');
    t.innerHTML = `
      <button class="event-list-item" data-field="container">
        <pf-v6-label compact data-field="label"></pf-v6-label>
        <time class="event-time" data-field="time"></time>
        <span class="event-element" data-field="element"></span>
      </button>`;
    return t;
  })();

  @property({ attribute: 'primary-tag-name' })
  accessor primaryTagName = '';

  @property({ attribute: 'demo-title' })
  accessor demoTitle = '';

  @property({ attribute: 'package-name' })
  accessor packageName = '';

  @property({ attribute: 'canonical-url' })
  accessor canonicalURL = '';

  @property({ attribute: 'source-url' })
  accessor sourceURL = '';

  @property()
  accessor knobs = '';

  @property()
  accessor drawer: 'expanded' | 'collapsed' | '' = '';

  @property({ attribute: 'drawer-height' })
  accessor drawerHeight = '';

  @property({ attribute: 'tabs-selected' })
  accessor tabsSelected = '';

  @property()
  accessor sidebar: 'expanded' | 'collapsed' | '' = '';

  @property({ type: Boolean, attribute: 'has-description' })
  accessor hasDescription = false;

  #$(id: string) {
    return this.shadowRoot?.getElementById(id);
  }

  #$$(selector: string) {
    return this.shadowRoot?.querySelectorAll(selector) ?? [];
  }

  #logContainer: HTMLElement | null = null;
  #drawerOpen = false;
  #initialLogsFetched = false;
  #isInitialLoad = true;
  #debugData: DebugData | null = null;

  // Element event tracking
  #elementEventMap: Map<string, EventInfo> | null = null;
  #manifest: Manifest | null = null;
  #capturedEvents: EventRecord[] = [];
  #maxCapturedEvents = 1000;
  #eventList: HTMLElement | null = null;
  #eventDetailHeader: HTMLElement | null = null;
  #eventDetailBody: HTMLElement | null = null;
  #selectedEventId: string | null = null;
  #eventsFilterValue = '';
  #eventsFilterDebounceTimer: ReturnType<typeof setTimeout> | null = null;
  #eventTypeFilters = new Set<string>();
  #elementFilters = new Set<string>();
  #discoveredElements = new Set<string>();

  // Event listener references for cleanup
  #handleLogsEvent: ((event: Event) => void) | null = null;
  #handleTreeExpand: ((event: Event) => void) | null = null;
  #handleTreeCollapse: ((event: Event) => void) | null = null;
  #handleTreeSelect: ((event: Event) => void) | null = null;

  // Timeout IDs for cleanup
  #copyLogsFeedbackTimeout: ReturnType<typeof setTimeout> | null = null;
  #copyDebugFeedbackTimeout: ReturnType<typeof setTimeout> | null = null;
  #copyEventsFeedbackTimeout: ReturnType<typeof setTimeout> | null = null;

  // Log filter state
  #logsFilterValue = '';
  #logsFilterDebounceTimer: ReturnType<typeof setTimeout> | null = null;
  #logLevelFilters = new Set(['info', 'warn', 'error', 'debug']);
  #logLevelDropdown: Element | null = null;

  // Watch for dynamically added elements
  /* c8 ignore start - MutationObserver callback tested via integration */
  #observer = new MutationObserver((mutations) => {
    let needsUpdate = false;

    for (const mutation of mutations) {
      for (const node of mutation.addedNodes) {
        if (node instanceof HTMLElement) {
          const tagName = node.tagName.toLowerCase();
          if (this.#elementEventMap?.has(tagName) && !node.dataset.cemEventsAttached) {
            const eventInfo = this.#elementEventMap.get(tagName)!;
            for (const eventName of eventInfo.eventNames) {
              node.addEventListener(eventName, this.#handleElementEvent, { capture: true });
            }
            node.dataset.cemEventsAttached = 'true';
            this.#discoveredElements.add(tagName);
            needsUpdate = true;
          }
        }
      }
    }

    if (needsUpdate) {
      this.#updateEventFilters();
    }
  });
  /* c8 ignore stop */

  #wsClient = new CEMReloadClient({
    jitterMax: 1000,
    overlayThreshold: 15,
    badgeFadeDelay: 2000,
    /* c8 ignore start - WebSocket callbacks tested via integration */
    callbacks: {
      onOpen: () => {
        this.#$('reconnection-modal')?.close();
      },
      onError: (errorData: { title?: string; message?: string; file?: string }) => {
        if (errorData?.title && errorData?.message) {
          console.error('[cem-serve] Server error:', errorData);
          (this.#$('error-overlay') as any)?.show(errorData.title, errorData.message, errorData.file);
        } else {
          console.error('[cem-serve] WebSocket error:', errorData);
        }
      },
      onReconnecting: ({ attempt, delay }: { attempt: number; delay: number }) => {
        if (attempt >= 15) {
          (this.#$('reconnection-modal') as any)?.showModal();
          (this.#$('reconnection-content') as any)?.updateRetryInfo(attempt, delay);
        }
      },
      onReload: () => {
        const errorOverlay = this.#$('error-overlay');
        if (errorOverlay?.hasAttribute('open')) {
          (errorOverlay as any).hide();
        }
        window.location.reload();
      },
      onShutdown: () => {
        (this.#$('reconnection-modal') as any)?.showModal();
        (this.#$('reconnection-content') as any)?.updateRetryInfo(30, 30000);
      },
      onLogs: (logs: Array<{ type: string; date: string; message: string }>) => {
        window.dispatchEvent(new CemLogsEvent(logs));
      }
    }
    /* c8 ignore stop */
  });

  get demo(): Element | null {
    return this.querySelector('cem-serve-demo');
  }

  render() {
    return html`
      <link rel="stylesheet" href="/__cem/pf-v6-c-description-list.css">
      <link rel="stylesheet" href="/__cem/pf-lightdom.css">

      <pf-v6-page ?sidebar-collapsed=${this.sidebar === 'collapsed'}>
        <pf-v6-skip-to-content slot="skip-to-content"
                               href="#main-content">
          Skip to content
        </pf-v6-skip-to-content>

        <pf-v6-masthead slot="masthead">
          <a class="masthead-logo"
             href="/"
             slot="logo">
            <img alt="CEM Dev Server"
                 src="/__cem/logo.svg">
            ${this.packageName ? html`<h1>${this.packageName}</h1>` : nothing}
          </a>
          <pf-v6-toolbar slot="toolbar">
            <pf-v6-toolbar-group variant="action-group">
              ${this.#renderSourceButton()}
              <pf-v6-toolbar-item>
                <pf-v6-button id="debug-info"
                              variant="plain"
                              aria-label="Debug info">
                  <svg width="16"
                       height="16"
                       viewBox="0 0 16 16"
                       fill="currentColor"
                       role="presentation">
                    <path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z"/>
                    <path d="m8.93 6.588-2.29.287-.082.38.45.083c.294.07.352.176.288.469l-.738 3.468c-.194.897.105 1.319.808 1.319.545 0 1.178-.252 1.465-.598l.088-.416c-.2.176-.492.246-.686.246-.275 0-.375-.193-.304-.533L8.93 6.588zM9 4.5a1 1 0 1 1-2 0 1 1 0 0 1 2 0z"/>
                  </svg>
                </pf-v6-button>
              </pf-v6-toolbar-item>
              <pf-v6-toolbar-item>
                <pf-v6-toggle-group class="color-scheme-toggle"
                                    aria-label="Color scheme">
                  <pf-v6-toggle-group-item value="light">
                    <svg width="16" height="16" viewBox="0 0 16 16" fill="currentColor" aria-label="Light mode">
                      <path d="M8 11a3 3 0 1 1 0-6 3 3 0 0 1 0 6zm0 1a4 4 0 1 0 0-8 4 4 0 0 0 0 8zM8 0a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-1 0v-2A.5.5 0 0 1 8 0zm0 13a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-1 0v-2A.5.5 0 0 1 8 13zm8-5a.5.5 0 0 1-.5.5h-2a.5.5 0 0 1 0-1h2a.5.5 0 0 1 .5.5zM3 8a.5.5 0 0 1-.5.5h-2a.5.5 0 0 1 0-1h2A.5.5 0 0 1 3 8zm10.657-5.657a.5.5 0 0 1 0 .707l-1.414 1.415a.5.5 0 1 1-.707-.708l1.414-1.414a.5.5 0 0 1 .707 0zm-9.193 9.193a.5.5 0 0 1 0 .707L3.05 13.657a.5.5 0 0 1-.707-.707l1.414-1.414a.5.5 0 0 1 .707 0zm9.193 2.121a.5.5 0 0 1-.707 0l-1.414-1.414a.5.5 0 0 1 .707-.707l1.414 1.414a.5.5 0 0 1 0 .707zM4.464 4.465a.5.5 0 0 1-.707 0L2.343 3.05a.5.5 0 1 1 .707-.707l1.414 1.414a.5.5 0 0 1 0 .708z"/>
                    </svg>
                  </pf-v6-toggle-group-item>
                  <pf-v6-toggle-group-item value="dark">
                    <svg width="16" height="16" viewBox="0 0 16 16" fill="currentColor" aria-label="Dark mode">
                      <path d="M6 .278a.768.768 0 0 1 .08.858 7.208 7.208 0 0 0-.878 3.46c0 4.021 3.278 7.277 7.318 7.277.527 0 1.04-.055 1.533-.16a.787.787 0 0 1 .81.316.733.733 0 0 1-.031.893A8.349 8.349 0 0 1 8.344 16C3.734 16 0 12.286 0 7.71 0 4.266 2.114 1.312 5.124.06A.752.752 0 0 1 6 .278zM4.858 1.311A7.269 7.269 0 0 0 1.025 7.71c0 4.02 3.279 7.276 7.319 7.276a7.316 7.316 0 0 0 5.205-2.162c-.337.042-.68.063-1.029.063-4.61 0-8.343-3.714-8.343-8.29 0-1.167.242-2.278.681-3.286z"/>
                    </svg>
                  </pf-v6-toggle-group-item>
                  <pf-v6-toggle-group-item value="system">
                    <svg width="16" height="16" viewBox="0 0 16 16" fill="currentColor" aria-label="System preference">
                      <path d="M0 1.5A1.5 1.5 0 0 1 1.5 0h13A1.5 1.5 0 0 1 16 1.5v8a1.5 1.5 0 0 1-1.5 1.5h-13A1.5 1.5 0 0 1 0 9.5v-8zM1.5 1a.5.5 0 0 0-.5.5v8a.5.5 0 0 0 .5.5h13a.5.5 0 0 0 .5-.5v-8a.5.5 0 0 0-.5-.5h-13z"/>
                      <path d="M2.5 12h11a.5.5 0 0 1 0 1h-11a.5.5 0 0 1 0-1zm0 2h11a.5.5 0 0 1 0 1h-11a.5.5 0 0 1 0-1z"/>
                    </svg>
                  </pf-v6-toggle-group-item>
                </pf-v6-toggle-group>
              </pf-v6-toolbar-item>
            </pf-v6-toolbar-group>
          </pf-v6-toolbar>
        </pf-v6-masthead>

        <pf-v6-page-sidebar slot="sidebar"
                            ?expanded=${this.sidebar === 'expanded'}
                            ?collapsed=${this.sidebar !== 'expanded'}>
          <slot name="navigation"></slot>
        </pf-v6-page-sidebar>

        <pf-v6-page-main slot="main" id="main-content">
          <slot></slot>
          <footer class="pf-m-sticky-bottom">
            <div class="footer-description${this.hasDescription ? '' : ' empty'}">
              <slot name="description"></slot>
            </div>
            <cem-drawer ?open=${this.drawer === 'expanded'}
                        drawer-height="${this.drawerHeight || '400'}">
              <pf-v6-tabs selected="${this.tabsSelected || '0'}">
                <pf-v6-tab title="Knobs">
                  <div id="knobs-container">
                    <slot name="knobs">
                      <p class="knobs-empty">No knobs available for this element.</p>
                    </slot>
                  </div>
                </pf-v6-tab>
                <pf-v6-tab title="Manifest Browser">
                  <cem-manifest-browser>
                    <slot name="manifest-tree" slot="manifest-tree"></slot>
                    <slot name="manifest-name" slot="manifest-name"></slot>
                    <slot name="manifest-details" slot="manifest-details"></slot>
                  </cem-manifest-browser>
                </pf-v6-tab>
                <pf-v6-tab title="Server Logs">
                  <div class="logs-wrapper">
                    <pf-v6-toolbar sticky>
                      <pf-v6-toolbar-group>
                        <pf-v6-toolbar-item>
                          <pf-v6-text-input-group id="logs-filter"
                                                  placeholder="Filter logs..."
                                                  icon>
                            <svg slot="icon"
                                 role="presentation"
                                 fill="currentColor"
                                 height="1em"
                                 width="1em"
                                 viewBox="0 0 512 512">
                              <path d="M505 442.7L405.3 343c-4.5-4.5-10.6-7-17-7H372c27.6-35.3 44-79.7 44-128C416 93.1 322.9 0 208 0S0 93.1 0 208s93.1 208 208 208c48.3 0 92.7-16.4 128-44v16.3c0 6.4 2.5 12.5 7 17l99.7 99.7c9.4 9.4 24.6 9.4 33.9 0l28.3-28.3c9.4-9.4 9.4-24.6.1-34zM208 336c-70.7 0-128-57.2-128-128 0-70.7 57.2-128 128-128 70.7 0 128 57.2 128 128 0 70.7-57.2 128-128 128z"></path>
                            </svg>
                          </pf-v6-text-input-group>
                        </pf-v6-toolbar-item>
                        <pf-v6-toolbar-item>
                          <pf-v6-dropdown id="log-level-filter"
                                          label="Filter log levels">
                            <span slot="toggle-text">Log Levels</span>
                            <pf-v6-menu-item variant="checkbox" value="info" checked>Info</pf-v6-menu-item>
                            <pf-v6-menu-item variant="checkbox" value="warn" checked>Warnings</pf-v6-menu-item>
                            <pf-v6-menu-item variant="checkbox" value="error" checked>Errors</pf-v6-menu-item>
                            <pf-v6-menu-item variant="checkbox" value="debug" checked>Debug</pf-v6-menu-item>
                          </pf-v6-dropdown>
                        </pf-v6-toolbar-item>
                      </pf-v6-toolbar-group>
                      <pf-v6-toolbar-group variant="action-group">
                        <pf-v6-toolbar-item>
                          <pf-v6-button id="copy-logs"
                                        variant="tertiary"
                                        size="small">
                            <svg slot="icon" width="16" height="16" viewBox="0 0 16 16" fill="currentColor">
                              <path d="M13 0H6a2 2 0 0 0-2 2 2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h7a2 2 0 0 0 2-2 2 2 0 0 0 2-2V2a2 2 0 0 0-2-2zm0 13V4a2 2 0 0 0-2-2H5a1 1 0 0 1 1-1h7a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1zM3 13V4a1 1 0 0 1 1-1h7a1 1 0 0 1 1 1v9a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1z"/>
                            </svg>
                            Copy Logs
                          </pf-v6-button>
                        </pf-v6-toolbar-item>
                      </pf-v6-toolbar-group>
                    </pf-v6-toolbar>
                    <div id="log-container"></div>
                  </div>
                </pf-v6-tab>
                <pf-v6-tab title="Events">
                  <div class="events-wrapper">
                    <pf-v6-toolbar sticky>
                      <pf-v6-toolbar-group>
                        <pf-v6-toolbar-item>
                          <pf-v6-text-input-group id="events-filter"
                                                  placeholder="Filter events..."
                                                  icon>
                            <svg slot="icon" width="16" height="16" viewBox="0 0 16 16" fill="currentColor">
                              <path d="M11.742 10.344a6.5 6.5 0 1 0-1.397 1.398h-.001c.03.04.062.078.098.115l3.85 3.85a1 1 0 0 0 1.415-1.414l-3.85-3.85a1.007 1.007 0 0 0-.115-.1zM12 6.5a5.5 5.5 0 1 1-11 0 5.5 5.5 0 0 1 11 0z"/>
                            </svg>
                          </pf-v6-text-input-group>
                        </pf-v6-toolbar-item>
                        <pf-v6-toolbar-item>
                          <pf-v6-dropdown id="event-type-filter"
                                          label="Filter event types">
                            <span slot="toggle-text">Event Types</span>
                          </pf-v6-dropdown>
                        </pf-v6-toolbar-item>
                        <pf-v6-toolbar-item>
                          <pf-v6-dropdown id="element-filter"
                                          label="Filter elements">
                            <span slot="toggle-text">Elements</span>
                          </pf-v6-dropdown>
                        </pf-v6-toolbar-item>
                      </pf-v6-toolbar-group>
                      <pf-v6-toolbar-group variant="action-group">
                        <pf-v6-toolbar-item>
                          <pf-v6-button id="clear-events"
                                        variant="tertiary"
                                        size="small">
                            <svg slot="icon" width="16" height="16" viewBox="0 0 16 16" fill="currentColor">
                              <path d="M5.5 5.5A.5.5 0 0 1 6 6v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm2.5 0a.5.5 0 0 1 .5.5v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm3 .5a.5.5 0 0 0-1 0v6a.5.5 0 0 0 1 0V6z"/>
                              <path fill-rule="evenodd" d="M14.5 3a1 1 0 0 1-1 1H13v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V4h-.5a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1H6a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1h3.5a1 1 0 0 1 1 1v1zM4.118 4 4 4.059V13a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V4.059L11.882 4H4.118zM2.5 3V2h11v1h-11z"/>
                            </svg>
                            Clear Events
                          </pf-v6-button>
                        </pf-v6-toolbar-item>
                        <pf-v6-toolbar-item>
                          <pf-v6-button id="copy-events"
                                        variant="tertiary"
                                        size="small">
                            <svg slot="icon" width="16" height="16" viewBox="0 0 16 16" fill="currentColor">
                              <path d="M13 0H6a2 2 0 0 0-2 2 2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h7a2 2 0 0 0 2-2 2 2 0 0 0 2-2V2a2 2 0 0 0-2-2zm0 13V4a2 2 0 0 0-2-2H5a1 1 0 0 1 1-1h7a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1zM3 13V4a1 1 0 0 1 1-1h7a1 1 0 0 1 1 1v9a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1z"/>
                            </svg>
                            Copy Events
                          </pf-v6-button>
                        </pf-v6-toolbar-item>
                      </pf-v6-toolbar-group>
                    </pf-v6-toolbar>
                    <pf-v6-drawer id="event-drawer" expanded>
                      <div id="event-list"></div>
                      <div id="event-detail-header" slot="panel-header"></div>
                      <div id="event-detail-body" slot="panel-body"></div>
                    </pf-v6-drawer>
                  </div>
                </pf-v6-tab>
                <pf-v6-tab title="Health">
                  <cem-health-panel ${this.primaryTagName ? html`component="${this.primaryTagName}"` : nothing}>
                  </cem-health-panel>
                </pf-v6-tab>
              </pf-v6-tabs>
            </cem-drawer>
          </footer>
        </pf-v6-page-main>
      </pf-v6-page>

      <pf-v6-modal id="debug-modal" variant="large">
        <h2 slot="header">Debug Information</h2>
        <dl class="pf-v6-c-description-list pf-m-horizontal pf-m-compact">
          <div class="pf-v6-c-description-list__group">
            <dt class="pf-v6-c-description-list__term">Server Version</dt>
            <dd class="pf-v6-c-description-list__description" id="debug-version">-</dd>
          </div>
          <div class="pf-v6-c-description-list__group">
            <dt class="pf-v6-c-description-list__term">Server OS</dt>
            <dd class="pf-v6-c-description-list__description" id="debug-os">-</dd>
          </div>
          <div class="pf-v6-c-description-list__group">
            <dt class="pf-v6-c-description-list__term">Watch Directory</dt>
            <dd class="pf-v6-c-description-list__description" id="debug-watch-dir">-</dd>
          </div>
          <div class="pf-v6-c-description-list__group">
            <dt class="pf-v6-c-description-list__term">Manifest Size</dt>
            <dd class="pf-v6-c-description-list__description" id="debug-manifest-size">-</dd>
          </div>
          <div class="pf-v6-c-description-list__group">
            <dt class="pf-v6-c-description-list__term">Demos Found</dt>
            <dd class="pf-v6-c-description-list__description" id="debug-demo-count">-</dd>
          </div>
          <div class="pf-v6-c-description-list__group">
            <dt class="pf-v6-c-description-list__term">Tag Name</dt>
            <dd class="pf-v6-c-description-list__description">${this.primaryTagName || '-'}</dd>
          </div>
          <div class="pf-v6-c-description-list__group">
            <dt class="pf-v6-c-description-list__term">Demo Title</dt>
            <dd class="pf-v6-c-description-list__description">${this.demoTitle || '-'}</dd>
          </div>
          <div class="pf-v6-c-description-list__group">
            <dt class="pf-v6-c-description-list__term">Browser</dt>
            <dd class="pf-v6-c-description-list__description" id="debug-browser">-</dd>
          </div>
          <div class="pf-v6-c-description-list__group">
            <dt class="pf-v6-c-description-list__term">User Agent</dt>
            <dd class="pf-v6-c-description-list__description" id="debug-ua">-</dd>
          </div>
        </dl>
        <div id="demo-urls-container"></div>
        <pf-v6-expandable-section id="debug-importmap-details"
                                  toggle-text="Show Import Map">
          <pre id="debug-importmap">-</pre>
        </pf-v6-expandable-section>
        <div slot="footer" class="button-container">
          <pf-v6-button class="debug-copy" variant="primary">
            Copy Debug Info
          </pf-v6-button>
          <pf-v6-button class="debug-close" variant="secondary">
            Close
          </pf-v6-button>
        </div>
      </pf-v6-modal>

      <!-- Reconnection modal -->
      <pf-v6-modal id="reconnection-modal" variant="large">
        <h2 slot="header">Development Server Disconnected</h2>
        <cem-reconnection-content id="reconnection-content"></cem-reconnection-content>
        <pf-v6-button id="reload-button"
                      slot="footer"
                      variant="primary">Reload Page</pf-v6-button>
        <pf-v6-button id="retry-button"
                      slot="footer"
                      variant="secondary">Retry Now</pf-v6-button>
      </pf-v6-modal>

      <!-- Transform error overlay -->
      <cem-transform-error-overlay id="error-overlay">
      </cem-transform-error-overlay>
    `;
  }

  #renderSourceButton() {
    if (!this.sourceURL) return nothing;

    let label = 'Version Control';
    let path = 'M5.854 4.854a.5.5 0 1 0-.708-.708l-3.5 3.5a.5.5 0 0 0 0 .708l3.5 3.5a.5.5 0 0 0 .708-.708L2.707 8l3.147-3.146zm4.292 0a.5.5 0 0 1 .708-.708l3.5 3.5a.5.5 0 0 1 0 .708l-3.5 3.5a.5.5 0 0 1-.708-.708L13.293 8l-3.147-3.146z';

    if (this.sourceURL.includes('github.com')) {
      label = 'GitHub.com';
      path = 'M8 0C3.58 0 0 3.58 0 8c0 3.54 2.29 6.53 5.47 7.59.4.07.55-.17.55-.38 0-.19-.01-.82-.01-1.49-2.01.37-2.53-.49-2.69-.94-.09-.23-.48-.94-.82-1.13-.28-.15-.68-.52-.01-.53.63-.01 1.08.58 1.23.82.72 1.21 1.87.87 2.33.66.07-.52.28-.87.51-1.07-1.78-.2-3.64-.89-3.64-3.95 0-.87.31-1.59.82-2.15-.08-.2-.36-1.02.08-2.12 0 0 .67-.21 2.2.82.64-.18 1.32-.27 2-.27.68 0 1.36.09 2 .27 1.53-1.04 2.2-.82 2.2-.82.44 1.1.16 1.92.08 2.12.51.56.82 1.27.82 2.15 0 3.07-1.87 3.75-3.65 3.95.29.25.54.73.54 1.48 0 1.07-.01 1.93-.01 2.2 0 .21.15.46.55.38A8.013 8.013 0 0016 8c0-4.42-3.58-8-8-8z';
    } else if (this.sourceURL.includes('gitlab.com')) {
      label = 'GitLab';
      path = 'm15.734 6.1-.022-.058L13.534.358a.568.568 0 0 0-.563-.356.583.583 0 0 0-.328.122.582.582 0 0 0-.193.294l-1.47 4.499H5.025l-1.47-4.5A.572.572 0 0 0 3.360.174a.572.572 0 0 0-.328-.172.582.582 0 0 0-.563.357L.29 6.04l-.022.057A4.044 4.044 0 0 0 1.61 10.77l.007.006.02.014 3.318 2.485 1.64 1.242 1 .755a.673.673 0 0 0 .814 0l1-.755 1.64-1.242 3.338-2.5.009-.007a4.05 4.05 0 0 0 1.34-4.668Z';
    } else if (this.sourceURL.includes('bitbucket.org')) {
      label = 'Bitbucket';
      path = 'M0 1.5A1.5 1.5 0 0 1 1.5 0h13A1.5 1.5 0 0 1 16 1.5v13a1.5 1.5 0 0 1-1.5 1.5h-13A1.5 1.5 0 0 1 0 14.5v-13zM2.5 3a.5.5 0 0 0-.5.5v9a.5.5 0 0 0 .5.5h11a.5.5 0 0 0 .5-.5v-9a.5.5 0 0 0-.5-.5h-11zm5.038 1.435a.5.5 0 0 1 .924 0l1.42 3.37H8.78l-.243.608-.243-.608H5.082l1.42-3.37zM8 9.143l-.743 1.857H4.743L6.076 7.608 8 9.143z';
    }

    return html`
      <pf-v6-toolbar-item>
        <pf-v6-button href="${this.sourceURL}"
                      target="_blank"
                      variant="plain"
                      aria-label="View source file">
          <svg aria-label="${label}"
               width="1rem"
               height="1rem"
               fill="currentColor"
               viewBox="0 0 16 16">
            <path d="${path}"/>
          </svg>
        </pf-v6-button>
      </pf-v6-toolbar-item>
    `;
  }

  connectedCallback() {
    super.connectedCallback();

    // Check if we need to migrate from localStorage
    this.#migrateFromLocalStorageIfNeeded();
  }

  firstUpdated() {
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

    // Set up sidebar state persistence
    this.#setupSidebarStatePersistence();

    // Set up element event capture
    this.#setupEventCapture().then(() => {
      this.#setupEventListeners();
    });

    // Set up reconnection modal button handlers
    /* c8 ignore start - window.location.reload would reload test page */
    this.#$('reload-button')?.addEventListener('click', () => {
      window.location.reload();
    });
    /* c8 ignore stop */

    this.#$('retry-button')?.addEventListener('click', () => {
      (this.#$('reconnection-modal') as any)?.close();
      this.#wsClient.retry();
    });

    // Initialize WebSocket connection
    this.#wsClient.init();
  }

  async #fetchDebugInfo() {
    // Populate browser info immediately
    const browserEl = this.#$('debug-browser');
    const uaEl = this.#$('debug-ua');
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
      .catch((err: Error) => {
        console.error('[cem-serve-chrome] Failed to fetch debug info:', err);
      });

    if (!data) return;
    const versionEl = this.#$('debug-version');
    const osEl = this.#$('debug-os');
    const watchDirEl = this.#$('debug-watch-dir');
    const manifestSizeEl = this.#$('debug-manifest-size');
    const demoCountEl = this.#$('debug-demo-count');
    const demoUrlsContainer = this.#$('demo-urls-container');
    const importMapEl = this.#$('debug-importmap');

    if (versionEl) versionEl.textContent = data.version || '-';
    if (osEl) osEl.textContent = data.os || '-';
    if (watchDirEl) watchDirEl.textContent = data.watchDir || '-';
    if (manifestSizeEl) manifestSizeEl.textContent = data.manifestSize || '-';
    if (demoCountEl) demoCountEl.textContent = data.demoCount || '0';

    if (demoUrlsContainer) {
      this.#populateDemoUrls(demoUrlsContainer, data.demos);
    }

    if (importMapEl && data.importMap) {
      const importMapJSON = JSON.stringify(data.importMap, null, 2);
      importMapEl.textContent = importMapJSON;
      data.importMapJSON = importMapJSON;
    } else if (importMapEl) {
      importMapEl.textContent = 'No import map generated';
    }

    this.#debugData = data;
  }

  #populateDemoUrls(container: HTMLElement, demos: DebugData['demos']) {
    if (!demos?.length) {
      container.textContent = 'No demos found in manifest';
      return;
    }

    const currentTagName = this.primaryTagName || '';
    const isOnDemoPage = !!currentTagName;

    if (isOnDemoPage) {
      const currentDemo = demos.find(demo => demo.tagName === currentTagName);
      if (!currentDemo) {
        container.textContent = 'Current demo not found in manifest';
        return;
      }

      const fragment = CemServeChrome.#demoInfoTemplate.content.cloneNode(true) as DocumentFragment;

      fragment.querySelector('[data-field="tagName"]')!.textContent = currentDemo.tagName;
      fragment.querySelector('[data-field="canonicalURL"]')!.textContent = currentDemo.canonicalURL;
      fragment.querySelector('[data-field="localRoute"]')!.textContent = currentDemo.localRoute;

      const descriptionGroup = fragment.querySelector('[data-field-group="description"]');
      if (currentDemo.description) {
        fragment.querySelector('[data-field="description"]')!.textContent = currentDemo.description;
      } else {
        descriptionGroup?.remove();
      }

      const filepathGroup = fragment.querySelector('[data-field-group="filepath"]');
      if (currentDemo.filepath) {
        fragment.querySelector('[data-field="filepath"]')!.textContent = currentDemo.filepath;
      } else {
        filepathGroup?.remove();
      }

      container.replaceChildren(fragment);
    } else {
      const listFragment = CemServeChrome.#demoListTemplate.content.cloneNode(true) as DocumentFragment;

      const groupsContainer = listFragment.querySelector('[data-container="groups"]')!;

      for (const demo of demos) {
        const groupFragment = CemServeChrome.#demoGroupTemplate.content.cloneNode(true) as DocumentFragment;

        groupFragment.querySelector('[data-field="tagName"]')!.textContent = demo.tagName;
        groupFragment.querySelector('[data-field="description"]')!.textContent =
          demo.description || '(no description)';
        groupFragment.querySelector('[data-field="canonicalURL"]')!.textContent = demo.canonicalURL;
        groupFragment.querySelector('[data-field="localRoute"]')!.textContent = demo.localRoute;

        const filepathGroup = groupFragment.querySelector('[data-field-group="filepath"]');
        if (demo.filepath) {
          groupFragment.querySelector('[data-field="filepath"]')!.textContent = demo.filepath;
        } else {
          filepathGroup?.remove();
        }

        groupsContainer.appendChild(groupFragment);
      }

      container.replaceChildren(listFragment);
    }
  }

  #setupLogListener() {
    this.#logContainer = this.#$('log-container');

    const logsFilter = this.#$('logs-filter') as HTMLInputElement | null;
    if (logsFilter) {
      logsFilter.addEventListener('input', () => {
        const { value = '' } = logsFilter;
        clearTimeout(this.#logsFilterDebounceTimer!);
        this.#logsFilterDebounceTimer = setTimeout(() => {
          this.#filterLogs(value);
        }, 300);
      });
    }

    this.#logLevelDropdown = this.shadowRoot?.querySelector('#log-level-filter') ?? null;
    if (this.#logLevelDropdown) {
      requestAnimationFrame(() => {
        this.#loadLogFilterState();
      });
      this.#logLevelDropdown.addEventListener('select', this.#handleLogFilterChange as EventListener);
    }

    this.#$('copy-logs')?.addEventListener('click', () => {
      this.#copyLogs();
    });

    this.#handleLogsEvent = ((event: Event) => {
      const logs = (event as CemLogsEvent).logs;
      if (logs) {
        this.#renderLogs(logs);
      }
    });
    window.addEventListener('cem:logs', this.#handleLogsEvent);
  }

  #filterLogs(query: string) {
    this.#logsFilterValue = query.toLowerCase();

    if (!this.#logContainer) return;

    for (const entry of this.#logContainer.children) {
      const text = entry.textContent?.toLowerCase() ?? '';
      const textMatch = !this.#logsFilterValue || text.includes(this.#logsFilterValue);

      const logType = this.#getLogTypeFromEntry(entry);
      const levelMatch = this.#logLevelFilters.has(logType);

      (entry as HTMLElement).hidden = !(textMatch && levelMatch);
    }
  }

  #getLogTypeFromEntry(entry: Element): string {
    for (const cls of entry.classList) {
      if (['info', 'warning', 'error', 'debug'].includes(cls)) {
        return cls === 'warning' ? 'warn' : cls;
      }
    }
    return 'info';
  }

  #loadLogFilterState() {
    try {
      const saved = localStorage.getItem('cem-serve-log-filters');
      if (saved) {
        this.#logLevelFilters = new Set(JSON.parse(saved));
      }
    } catch (e) {
      console.debug('[cem-serve-chrome] localStorage unavailable, using default log filters');
    }
    this.#syncCheckboxStates();
  }

  #syncCheckboxStates() {
    if (!this.#logLevelDropdown) return;
    const menuItems = this.#logLevelDropdown.querySelectorAll('pf-v6-menu-item');
    menuItems.forEach(item => {
      const value = (item as any).value;
      (item as any).checked = this.#logLevelFilters.has(value);
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

  #handleLogFilterChange = (event: Event) => {
    const { value, checked } = event as Event & { value: string; checked: boolean };

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

    const logs = Array.from(this.#logContainer.children)
      .filter(entry => !(entry as HTMLElement).hidden)
      .map(entry => {
        const type = entry.querySelector('[data-field="label"]')?.textContent?.trim() || 'INFO';
        const time = entry.querySelector('[data-field="time"]')?.textContent?.trim() || '';
        const message = entry.querySelector('[data-field="message"]')?.textContent?.trim() || '';
        return `[${type}] ${time} ${message}`;
      }).join('\n');

    if (!logs) return;

    try {
      await navigator.clipboard.writeText(logs);
      const btn = this.#$('copy-logs');
      if (btn) {
        const textNode = Array.from(btn.childNodes).find(
          n => n.nodeType === Node.TEXT_NODE && (n.textContent?.trim().length ?? 0) > 0
        );
        if (textNode) {
          const original = textNode.textContent;
          textNode.textContent = 'Copied!';

          if (this.#copyLogsFeedbackTimeout) {
            clearTimeout(this.#copyLogsFeedbackTimeout);
          }

          this.#copyLogsFeedbackTimeout = setTimeout(() => {
            if (this.isConnected && textNode.parentNode) {
              textNode.textContent = original;
            }
            this.#copyLogsFeedbackTimeout = null;
          }, 2000);
        }
      }
    } catch (err) {
      console.error('[cem-serve-chrome] Failed to copy logs:', err);
    }
  }

  #setupDebugOverlay() {
    const debugButton = this.#$('debug-info');
    const debugModal = this.#$('debug-modal');
    const debugClose = this.shadowRoot?.querySelector('.debug-close');
    const debugCopy = this.shadowRoot?.querySelector('.debug-copy');

    if (debugButton && debugModal) {
      debugButton.addEventListener('click', () => {
        this.#fetchDebugInfo();
        (debugModal as any).showModal();
      });

      debugClose?.addEventListener('click', () => (debugModal as any).close());

      debugCopy?.addEventListener('click', () => {
        this.#copyDebugInfo();
      });
    }
  }

  #setupFooterDrawer() {
    const drawer = this.shadowRoot?.querySelector('cem-drawer');
    const tabs = this.shadowRoot?.querySelector('pf-v6-tabs');

    if (!drawer || !tabs) return;

    this.#drawerOpen = (drawer as any).open;

    drawer.addEventListener('change', (e: Event) => {
      this.#drawerOpen = (e as any).open;

      StatePersistence.updateState({
        drawer: { open: (e as any).open }
      });

      if ((e as any).open) {
        this.#scrollLogsToBottom();
      }
    });

    drawer.addEventListener('resize', (e: Event) => {
      (drawer as any).setAttribute('drawer-height', (e as any).height);

      StatePersistence.updateState({
        drawer: { height: (e as any).height }
      });
    });

    tabs.addEventListener('change', (e: Event) => {
      StatePersistence.updateState({
        tabs: { selectedIndex: (e as any).selectedIndex }
      });

      if ((e as any).selectedIndex === 2 && (drawer as any).open) {
        this.#scrollLogsToBottom();
      }

      if ((e as any).selectedIndex === 3 && (drawer as any).open) {
        this.#scrollEventsToBottom();
      }
    });
  }

  #detectBrowser(): string {
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
    const info = Array.from(this.#$$('#debug-modal dl dt')).map(dt => {
      const dd = dt.nextElementSibling;
      if (dd && dd.tagName === 'DD') {
        return `${dt.textContent}: ${dd.textContent}`;
      }
      return '';
    }).join('\n');

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
      const copyButton = this.shadowRoot?.querySelector('.debug-copy');
      if (copyButton) {
        const originalText = copyButton.textContent;
        copyButton.textContent = 'Copied!';

        if (this.#copyDebugFeedbackTimeout) {
          clearTimeout(this.#copyDebugFeedbackTimeout);
        }

        this.#copyDebugFeedbackTimeout = setTimeout(() => {
          if (this.isConnected && copyButton.parentNode) {
            copyButton.textContent = originalText;
          }
          this.#copyDebugFeedbackTimeout = null;
        }, 2000);
      }
    } catch (err) {
      console.error('[cem-serve-chrome] Failed to copy debug info:', err);
    }
  }

  #renderLogs(logs: Array<{ type: string; date: string; message: string }>) {
    if (!this.#logContainer) return;

    const logElements = logs.map(log => {
      const fragment = CemServeChrome.#logEntryTemplate.content.cloneNode(true) as DocumentFragment;

      const date = new Date(log.date);
      const time = date.toLocaleTimeString();

      const container = fragment.querySelector('[data-field="container"]') as HTMLElement;
      container.classList.add(log.type);
      container.setAttribute('data-log-id', log.date);

      const typeLabel = this.#getLogBadge(log.type);
      const searchText = `${typeLabel} ${time} ${log.message}`.toLowerCase();
      const textMatch = !this.#logsFilterValue || searchText.includes(this.#logsFilterValue);

      const logTypeForFilter = log.type === 'warning' ? 'warn' : log.type;
      const levelMatch = this.#logLevelFilters.has(logTypeForFilter);

      if (!(textMatch && levelMatch)) {
        container.setAttribute('hidden', '');
      }

      const label = fragment.querySelector('[data-field="label"]') as HTMLElement;
      label.textContent = this.#getLogBadge(log.type);
      this.#applyLogLabelAttrs(label, log.type);

      const timeEl = fragment.querySelector('[data-field="time"]') as HTMLElement;
      timeEl.setAttribute('datetime', log.date);
      timeEl.textContent = time;

      (fragment.querySelector('[data-field="message"]') as HTMLElement).textContent = log.message;

      return fragment;
    });

    if (!this.#initialLogsFetched) {
      this.#logContainer.replaceChildren(...logElements);
      this.#initialLogsFetched = true;

      if (this.#drawerOpen) {
        this.#scrollLatestIntoView();
      }
    } else {
      this.#logContainer.append(...logElements);

      if (this.#drawerOpen) {
        this.#scrollLatestIntoView();
      }
    }
  }

  #getLogBadge(type: string): string {
    switch (type) {
      case 'info': return 'Info';
      case 'warning': return 'Warn';
      case 'error': return 'Error';
      case 'debug': return 'Debug';
      default: return type.toUpperCase();
    }
  }

  #applyLogLabelAttrs(label: HTMLElement, type: string) {
    switch (type) {
      case 'info':
        return label.setAttribute('status', 'info');
      case 'warning':
        return label.setAttribute('status', 'warning');
      case 'error':
        return label.setAttribute('status', 'danger');
      case 'debug':
        return label.setAttribute('color', 'purple');
      default:
        label.setAttribute('color', 'grey');
    }
  }

  #scrollLatestIntoView() {
    if (!this.#logContainer) return;

    requestAnimationFrame(() => {
      const lastLog = this.#logContainer!.lastElementChild;
      if (lastLog) {
        lastLog.scrollIntoView({ behavior: 'auto', block: 'end' });
      }
    });
  }

  #scrollLogsToBottom() {
    if (!this.#logContainer) return;

    if (this.#isInitialLoad) {
      this.#scrollLatestIntoView();
    } else {
      setTimeout(() => {
        this.#scrollLatestIntoView();
      }, 350);
    }
  }

  #migrateFromLocalStorageIfNeeded() {
    try {
      const hasLocalStorage =
        localStorage.getItem('cem-serve-color-scheme') !== null ||
        localStorage.getItem('cem-serve-drawer-open') !== null ||
        localStorage.getItem('cem-serve-drawer-height') !== null ||
        localStorage.getItem('cem-serve-active-tab') !== null;

      if (hasLocalStorage) {
        const migrated = localStorage.getItem('cem-serve-migrated-to-cookies');
        if (!migrated) {
          StatePersistence.migrateFromLocalStorage();
          localStorage.setItem('cem-serve-migrated-to-cookies', 'true');
          setTimeout(() => window.location.reload(), 100);
        }
      }
    } catch (e) {
      // localStorage not available, skip migration
    }
  }

  #setupColorSchemeToggle() {
    const toggleGroup = this.shadowRoot?.querySelector('.color-scheme-toggle');
    if (!toggleGroup) return;

    const state = StatePersistence.getState();

    this.#applyColorScheme(state.colorScheme);

    const items = toggleGroup.querySelectorAll('pf-v6-toggle-group-item');
    items.forEach(item => {
      if ((item as any).value === state.colorScheme) {
        item.setAttribute('selected', '');
      }
    });

    toggleGroup.addEventListener('pf-v6-toggle-group-change', (e: Event) => {
      const scheme = (e as any).value;
      this.#applyColorScheme(scheme);
      StatePersistence.updateState({ colorScheme: scheme });
    });
  }

  #applyColorScheme(scheme: string) {
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
    this.addEventListener('knob:attribute-change', this.#onKnobChange);
    this.addEventListener('knob:property-change', this.#onKnobChange);
    this.addEventListener('knob:css-property-change', this.#onKnobChange);
    this.addEventListener('knob:attribute-clear', this.#onKnobClear);
    this.addEventListener('knob:property-clear', this.#onKnobClear);
    this.addEventListener('knob:css-property-clear', this.#onKnobClear);
  }

  #onKnobChange = (event: Event) => {
    const target = this.#getKnobTarget(event);
    if (!target) {
      console.warn('[cem-serve-chrome] Could not find knob target info in event path');
      return;
    }

    const { tagName, instanceIndex } = target;

    const demo = this.demo;
    if (!demo) return;

    const knobType = this.#getKnobTypeFromEvent(event);

    const success = (demo as any).applyKnobChange(
      knobType,
      (event as any).name,
      (event as any).value,
      tagName,
      instanceIndex
    );

    if (!success) {
      console.warn('[cem-serve-chrome] Failed to apply knob change:', {
        type: knobType,
        name: (event as any).name,
        tagName,
        instanceIndex
      });
    }
  };

  #onKnobClear = (event: Event) => {
    const target = this.#getKnobTarget(event);
    if (!target) {
      console.warn('[cem-serve-chrome] Could not find knob target info in event path');
      return;
    }

    const { tagName, instanceIndex } = target;

    const demo = this.demo;
    if (!demo) return;

    const knobType = this.#getKnobTypeFromClearEvent(event);
    const clearValue = knobType === 'property' ? undefined : '';

    const success = (demo as any).applyKnobChange(
      knobType,
      (event as any).name,
      clearValue,
      tagName,
      instanceIndex
    );

    if (!success) {
      console.warn('[cem-serve-chrome] Failed to clear knob:', {
        type: knobType,
        name: (event as any).name,
        tagName,
        instanceIndex
      });
    }
  };

  #getKnobTarget(event: Event): { tagName: string; instanceIndex: number } | null {
    const defaultTagName = this.primaryTagName || '';

    if (event.composedPath) {
      for (const element of event.composedPath()) {
        if (!(element instanceof Element)) continue;

        if ((element as HTMLElement).dataset?.isElementKnob === 'true') {
          const tagName = (element as HTMLElement).dataset.tagName || defaultTagName;
          let instanceIndex = Number.parseInt((element as HTMLElement).dataset.instanceIndex ?? '', 10);
          if (Number.isNaN(instanceIndex)) instanceIndex = 0;
          return { tagName, instanceIndex };
        }
      }
    }

    return { tagName: defaultTagName, instanceIndex: 0 };
  }

  #getKnobTypeFromEvent(event: Event): string {
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

  #getKnobTypeFromClearEvent(event: Event): string {
    switch (event.type) {
      case 'knob:attribute-clear':
        return 'attribute';
      case 'knob:property-clear':
        return 'property';
      case 'knob:css-property-clear':
        return 'css-property';
      default:
        return 'unknown';
    }
  }

  #setupTreeStatePersistence() {
    this.#handleTreeExpand = (e: Event) => {
      if ((e.target as Element)?.tagName !== 'PF-V6-TREE-ITEM') return;

      const nodeId = this.#getTreeNodeId(e.target as Element);
      const treeState = StatePersistence.getTreeState();
      if (!treeState.expanded.includes(nodeId)) {
        treeState.expanded.push(nodeId);
        StatePersistence.setTreeState(treeState);
      }
    };
    this.addEventListener('expand', this.#handleTreeExpand);

    this.#handleTreeCollapse = (e: Event) => {
      if ((e.target as Element)?.tagName !== 'PF-V6-TREE-ITEM') return;

      const nodeId = this.#getTreeNodeId(e.target as Element);
      const treeState = StatePersistence.getTreeState();
      const index = treeState.expanded.indexOf(nodeId);
      if (index > -1) {
        treeState.expanded.splice(index, 1);
        StatePersistence.setTreeState(treeState);
      }
    };
    this.addEventListener('collapse', this.#handleTreeCollapse);

    this.#handleTreeSelect = (e: Event) => {
      if ((e.target as Element)?.tagName !== 'PF-V6-TREE-ITEM') return;

      const nodeId = this.#getTreeNodeId(e.target as Element);
      StatePersistence.updateTreeState({ selected: nodeId });
    };
    this.addEventListener('select', this.#handleTreeSelect);

    this.#applyTreeState();
  }

  #applyTreeState() {
    const treeState = StatePersistence.getTreeState();

    for (const nodeId of treeState.expanded) {
      const treeItem = this.#findTreeItemById(nodeId);
      if (treeItem && !treeItem.hasAttribute('expanded')) {
        treeItem.setAttribute('expanded', '');
      }
    }

    if (treeState.selected) {
      const treeItem = this.#findTreeItemById(treeState.selected);
      if (treeItem && !treeItem.hasAttribute('current')) {
        treeItem.setAttribute('current', '');
      }
    }
  }

  #setupSidebarStatePersistence() {
    const page = this.shadowRoot?.querySelector('pf-v6-page');

    if (!page) return;

    page.addEventListener('sidebar-toggle', (event: Event) => {
      const collapsed = !(event as any).expanded;

      StatePersistence.updateState({
        sidebar: { collapsed }
      });
    });
  }

  #findTreeItemById(nodeId: string): Element | null {
    const parts = nodeId.split(':');
    const [type, modulePath, tagName, name] = parts;

    let attrSuffix = '';
    if (tagName) {
      attrSuffix += `[data-tag-name="${CSS.escape(tagName)}"]`;
    }
    if (name) {
      attrSuffix += `[data-name="${CSS.escape(name)}"]`;
    }

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

  #getTreeNodeId(treeItem: Element): string {
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

  // Event Discovery & Capture Methods

  async #discoverElementEvents(): Promise<Map<string, EventInfo>> {
    try {
      const response = await fetch('/custom-elements.json');
      if (!response.ok) {
        console.warn('[cem-serve-chrome] No manifest available for event discovery');
        return new Map();
      }

      const manifest = await response.json() as Manifest;
      this.#manifest = manifest;

      const eventMap = new Map<string, EventInfo>();

      for (const module of manifest.modules || []) {
        for (const declaration of module.declarations || []) {
          if (declaration.customElement && declaration.tagName) {
            const tagName = declaration.tagName;
            const events = declaration.events || [];

            if (events.length > 0) {
              const eventNames = new Set(events.map(e => e.name));
              eventMap.set(tagName, {
                eventNames,
                events: events
              });
            }
          }
        }
      }

      return eventMap;
    } catch (error) {
      console.warn('[cem-serve-chrome] Error loading manifest for event discovery:', error);
      return new Map();
    }
  }

  async #setupEventCapture() {
    this.#elementEventMap = await this.#discoverElementEvents();

    if (this.#elementEventMap.size === 0) return;

    this.#attachEventListeners();
    this.#updateEventFilters();
    this.#observeDemoMutations();
  }

  #attachEventListeners() {
    const demo = this.demo;
    if (!demo) return;

    const root = demo.shadowRoot ?? demo;

    for (const [tagName, eventInfo] of this.#elementEventMap!) {
      const elements = root.querySelectorAll(tagName);

      for (const element of elements) {
        for (const eventName of eventInfo.eventNames) {
          element.addEventListener(eventName, this.#handleElementEvent, { capture: true });
        }
        (element as HTMLElement).dataset.cemEventsAttached = 'true';
        this.#discoveredElements.add(tagName);
      }
    }
  }

  #observeDemoMutations() {
    const demo = this.demo;
    if (!demo) return;

    const root = demo.shadowRoot ?? demo;

    this.#observer.observe(root, {
      childList: true,
      subtree: true
    });
  }

  #handleElementEvent = (event: Event) => {
    const element = event.currentTarget;
    if (!(element instanceof HTMLElement)) return;

    const tagName = element.tagName.toLowerCase();
    const eventInfo = this.#elementEventMap?.get(tagName);

    if (!eventInfo || !eventInfo.eventNames.has(event.type)) return;

    this.#discoveredElements.add(tagName);
    this.#captureEvent(event, element, tagName, eventInfo);
  };

  #getEventDocumentation(manifestEvent: EventInfo['events'][0] | undefined) {
    if (!manifestEvent) {
      return { summary: null, description: null };
    }

    let summary = manifestEvent.summary || null;
    let description = manifestEvent.description || null;

    if (manifestEvent.type?.text && this.#manifest) {
      const typeName = manifestEvent.type.text;
      const typeDeclaration = this.#findTypeDeclaration(typeName);

      if (typeDeclaration) {
        if (!summary && typeDeclaration.summary) {
          summary = typeDeclaration.summary;
        } else if (typeDeclaration.summary && typeDeclaration.summary !== summary) {
          summary = summary ? `${summary}\n\nFrom ${typeName}: ${typeDeclaration.summary}` : typeDeclaration.summary;
        }

        if (!description && typeDeclaration.description) {
          description = typeDeclaration.description;
        } else if (typeDeclaration.description && typeDeclaration.description !== description) {
          description = description ? `${description}\n\n${typeDeclaration.description}` : typeDeclaration.description;
        }
      }
    }

    return { summary, description };
  }

  #findTypeDeclaration(typeName: string) {
    if (!this.#manifest) return null;

    for (const module of this.#manifest.modules || []) {
      for (const declaration of module.declarations || []) {
        if (declaration.name === typeName &&
            (declaration.kind === 'class' || declaration.kind === 'interface')) {
          return declaration as { summary?: string; description?: string };
        }
      }
    }

    return null;
  }

  #captureEvent(event: Event, target: HTMLElement, tagName: string, eventInfo: EventInfo) {
    const manifestEvent = eventInfo.events.find(e => e.name === event.type);

    const eventDocs = this.#getEventDocumentation(manifestEvent);

    const customProperties = this.#extractEventProperties(event);

    const eventRecord: EventRecord = {
      id: `${Date.now()}-${Math.random()}`,
      timestamp: new Date(),
      eventName: event.type,
      tagName: tagName,
      elementId: target.id || null,
      elementClass: target.className || null,
      customProperties: customProperties,
      manifestType: manifestEvent?.type?.text || null,
      summary: eventDocs.summary,
      description: eventDocs.description,
      bubbles: event.bubbles,
      composed: event.composed,
      cancelable: event.cancelable,
      defaultPrevented: event.defaultPrevented
    };

    this.#capturedEvents.push(eventRecord);

    if (this.#capturedEvents.length > this.#maxCapturedEvents) {
      this.#capturedEvents.shift();
    }

    this.#renderEvent(eventRecord);
  }

  #extractEventProperties(event: Event): Record<string, unknown> {
    const properties: Record<string, unknown> = {};
    const eventPrototypeKeys = new Set(Object.getOwnPropertyNames(Event.prototype));

    const serializeValue = (value: unknown): unknown => {
      try {
        return JSON.parse(JSON.stringify(value));
      } catch (e) {
        try {
          return String(value);
        } catch (stringErr) {
          return '[Not serializable]';
        }
      }
    };

    if (event instanceof CustomEvent && event.detail !== undefined) {
      properties.detail = serializeValue(event.detail);
    }

    for (const key of Object.getOwnPropertyNames(event)) {
      if (!eventPrototypeKeys.has(key) && !key.startsWith('_') && !properties.hasOwnProperty(key)) {
        properties[key] = serializeValue((event as any)[key]);
      }
    }

    return properties;
  }

  #renderEvent(eventRecord: EventRecord) {
    if (!this.#eventList) return;

    const fragment = CemServeChrome.#eventEntryTemplate.content.cloneNode(true) as DocumentFragment;

    const time = eventRecord.timestamp.toLocaleTimeString();

    const container = fragment.querySelector('[data-field="container"]') as HTMLElement;
    container.dataset.eventId = eventRecord.id;
    container.dataset.eventType = eventRecord.eventName;
    container.dataset.elementType = eventRecord.tagName;

    const textMatch = this.#eventMatchesTextFilter(eventRecord);
    const typeMatch = this.#eventTypeFilters.size === 0 || this.#eventTypeFilters.has(eventRecord.eventName);
    const elementMatch = this.#elementFilters.size === 0 || this.#elementFilters.has(eventRecord.tagName);

    if (!(textMatch && typeMatch && elementMatch)) {
      container.setAttribute('hidden', '');
    }

    const label = fragment.querySelector('[data-field="label"]') as HTMLElement;
    label.textContent = eventRecord.eventName;
    label.setAttribute('status', 'info');

    const timeEl = fragment.querySelector('[data-field="time"]') as HTMLElement;
    timeEl.setAttribute('datetime', eventRecord.timestamp.toISOString());
    timeEl.textContent = time;

    const elementEl = fragment.querySelector('[data-field="element"]') as HTMLElement;
    let elementText = `<${eventRecord.tagName}>`;
    if (eventRecord.elementId) {
      elementText += `#${eventRecord.elementId}`;
    }
    elementEl.textContent = elementText;

    this.#eventList.append(fragment);

    if (!this.#selectedEventId) {
      this.#selectEvent(eventRecord.id);
    }

    if (this.#drawerOpen && this.#isEventsTabActive()) {
      this.#scrollEventsToBottom();
    }
  }

  #selectEvent(eventId: string) {
    const eventRecord = this.#getEventRecordById(eventId);
    if (!eventRecord) return;

    this.#selectedEventId = eventId;

    const allItems = this.#eventList?.querySelectorAll('.event-list-item');
    allItems?.forEach(item => {
      if ((item as HTMLElement).dataset.eventId === eventId) {
        item.classList.add('selected');
        item.setAttribute('aria-selected', 'true');
      } else {
        item.classList.remove('selected');
        item.setAttribute('aria-selected', 'false');
      }
    });

    if (this.#eventDetailHeader) {
      this.#eventDetailHeader.innerHTML = '';

      const headerContent = document.createElement('div');
      headerContent.className = 'event-detail-header-content';

      const eventName = document.createElement('h3');
      eventName.textContent = eventRecord.eventName;
      eventName.className = 'event-detail-name';
      headerContent.appendChild(eventName);

      if (eventRecord.summary) {
        const summary = document.createElement('p');
        summary.textContent = eventRecord.summary;
        summary.className = 'event-detail-summary';
        headerContent.appendChild(summary);
      }

      if (eventRecord.description) {
        const description = document.createElement('p');
        description.textContent = eventRecord.description;
        description.className = 'event-detail-description';
        headerContent.appendChild(description);
      }

      const meta = document.createElement('div');
      meta.className = 'event-detail-meta';

      const timeEl = document.createElement('time');
      timeEl.setAttribute('datetime', eventRecord.timestamp.toISOString());
      timeEl.textContent = eventRecord.timestamp.toLocaleTimeString();
      timeEl.className = 'event-detail-time';

      const element = document.createElement('span');
      let elementText = `<${eventRecord.tagName}>`;
      if (eventRecord.elementId) {
        elementText += `#${eventRecord.elementId}`;
      }
      element.textContent = elementText;
      element.className = 'event-detail-element';

      meta.appendChild(timeEl);
      meta.appendChild(element);

      headerContent.appendChild(meta);

      this.#eventDetailHeader.appendChild(headerContent);
    }

    if (this.#eventDetailBody) {
      this.#eventDetailBody.innerHTML = '';

      const propertiesHeading = document.createElement('h4');
      propertiesHeading.textContent = 'Properties';
      propertiesHeading.className = 'event-detail-properties-heading';

      const propertiesContainer = document.createElement('div');
      propertiesContainer.className = 'event-detail-properties';

      const eventProperties = this.#buildPropertiesForDisplay(eventRecord);
      if (Object.keys(eventProperties).length > 0) {
        propertiesContainer.appendChild(this.#buildPropertyTree(eventProperties));
      } else {
        propertiesContainer.textContent = 'No properties to display';
      }

      this.#eventDetailBody.appendChild(propertiesHeading);
      this.#eventDetailBody.appendChild(propertiesContainer);
    }
  }

  #buildPropertiesForDisplay(eventRecord: EventRecord): Record<string, unknown> {
    const properties: Record<string, unknown> = {};

    if (eventRecord.customProperties) {
      Object.assign(properties, eventRecord.customProperties);
    }

    properties.bubbles = eventRecord.bubbles;
    properties.cancelable = eventRecord.cancelable;
    properties.defaultPrevented = eventRecord.defaultPrevented;
    properties.composed = eventRecord.composed;

    if (eventRecord.manifestType) {
      properties.type = eventRecord.manifestType;
    }

    return properties;
  }

  #buildPropertyTree(obj: Record<string, unknown>, depth = 0): HTMLUListElement {
    const ul = document.createElement('ul');
    ul.className = 'event-property-tree';
    if (depth > 0) {
      ul.classList.add('nested');
    }

    for (const [key, value] of Object.entries(obj)) {
      const li = document.createElement('li');
      li.className = 'property-item';

      const keySpan = document.createElement('span');
      keySpan.className = 'property-key';
      keySpan.textContent = key;

      const colonSpan = document.createElement('span');
      colonSpan.className = 'property-colon';
      colonSpan.textContent = ': ';

      li.appendChild(keySpan);
      li.appendChild(colonSpan);

      if (value === null || value === undefined) {
        const valueSpan = document.createElement('span');
        valueSpan.className = 'property-value null';
        valueSpan.textContent = String(value);
        li.appendChild(valueSpan);
      } else if (typeof value === 'boolean') {
        const valueSpan = document.createElement('span');
        valueSpan.className = 'property-value boolean';
        valueSpan.textContent = String(value);
        li.appendChild(valueSpan);
      } else if (typeof value === 'number') {
        const valueSpan = document.createElement('span');
        valueSpan.className = 'property-value number';
        valueSpan.textContent = String(value);
        li.appendChild(valueSpan);
      } else if (typeof value === 'string') {
        const valueSpan = document.createElement('span');
        valueSpan.className = 'property-value string';
        valueSpan.textContent = `"${value}"`;
        li.appendChild(valueSpan);
      } else if (Array.isArray(value)) {
        const valueSpan = document.createElement('span');
        valueSpan.className = 'property-value array';
        valueSpan.textContent = `Array(${value.length})`;
        li.appendChild(valueSpan);

        if (value.length > 0 && depth < 3) {
          const nestedObj: Record<string, unknown> = {};
          value.forEach((item, index) => {
            nestedObj[index] = item;
          });
          li.appendChild(this.#buildPropertyTree(nestedObj, depth + 1));
        }
      } else if (typeof value === 'object') {
        const valueSpan = document.createElement('span');
        valueSpan.className = 'property-value object';
        const keys = Object.keys(value as Record<string, unknown>);
        valueSpan.textContent = keys.length > 0 ? `Object` : `{}`;
        li.appendChild(valueSpan);

        if (keys.length > 0 && depth < 3) {
          li.appendChild(this.#buildPropertyTree(value as Record<string, unknown>, depth + 1));
        }
      } else {
        const valueSpan = document.createElement('span');
        valueSpan.className = 'property-value';
        valueSpan.textContent = String(value);
        li.appendChild(valueSpan);
      }

      ul.appendChild(li);
    }

    return ul;
  }

  #scrollEventsToBottom() {
    if (!this.#eventList) return;

    requestAnimationFrame(() => {
      const lastEvent = this.#eventList!.lastElementChild;
      if (lastEvent) {
        lastEvent.scrollIntoView({ behavior: 'auto', block: 'end' });
      }
    });
  }

  #isEventsTabActive(): boolean {
    const tabs = this.shadowRoot?.querySelector('pf-v6-tabs');
    if (!tabs) return false;

    const selectedIndex = parseInt(tabs.getAttribute('selected') || '0', 10);
    return selectedIndex === 3;
  }

  #filterEvents(query: string) {
    this.#eventsFilterValue = query.toLowerCase();

    if (!this.#eventList) return;

    for (const entry of this.#eventList.children) {
      const eventRecord = this.#getEventRecordById((entry as HTMLElement).dataset.eventId!);

      if (!eventRecord) continue;

      const textMatch = this.#eventMatchesTextFilter(eventRecord);
      const typeMatch = this.#eventTypeFilters.size === 0 || this.#eventTypeFilters.has(eventRecord.eventName);
      const elementMatch = this.#elementFilters.size === 0 || this.#elementFilters.has(eventRecord.tagName);

      (entry as HTMLElement).hidden = !(textMatch && typeMatch && elementMatch);
    }
  }

  #eventMatchesTextFilter(eventRecord: EventRecord): boolean {
    if (!this.#eventsFilterValue) return true;

    const searchText = [
      eventRecord.tagName,
      eventRecord.eventName,
      eventRecord.elementId || '',
      JSON.stringify(eventRecord.customProperties || {})
    ].join(' ').toLowerCase();

    return searchText.includes(this.#eventsFilterValue);
  }

  #getEventRecordById(id: string): EventRecord | undefined {
    return this.#capturedEvents.find(e => e.id === id);
  }

  #updateEventFilters() {
    const savedPreferences = this.#loadEventFiltersFromStorage();

    const eventTypeFilter = this.#$('event-type-filter');
    if (eventTypeFilter && this.#elementEventMap) {
      let menu = eventTypeFilter.querySelector('pf-v6-menu');
      if (!menu) {
        menu = document.createElement('pf-v6-menu');
        eventTypeFilter.appendChild(menu);
      }

      const existingItems = menu.querySelectorAll('pf-v6-menu-item');
      existingItems.forEach(item => item.remove());

      const allEventTypes = new Set<string>();
      for (const [tagName, eventInfo] of this.#elementEventMap) {
        if (this.#discoveredElements.has(tagName)) {
          for (const eventName of eventInfo.eventNames) {
            allEventTypes.add(eventName);
          }
        }
      }

      if (savedPreferences.eventTypes) {
        this.#eventTypeFilters = (savedPreferences.eventTypes as Set<string> & { intersection: (other: Set<string>) => Set<string> }).intersection(allEventTypes);
      } else {
        this.#eventTypeFilters = new Set(allEventTypes);
      }

      for (const eventName of allEventTypes) {
        const item = document.createElement('pf-v6-menu-item');
        item.setAttribute('variant', 'checkbox');
        item.setAttribute('value', eventName);
        if (this.#eventTypeFilters.has(eventName)) {
          item.setAttribute('checked', '');
        }
        item.textContent = eventName;
        menu.appendChild(item);
      }
    }

    const elementFilter = this.#$('element-filter');
    if (elementFilter && this.#elementEventMap) {
      let menu = elementFilter.querySelector('pf-v6-menu');
      if (!menu) {
        menu = document.createElement('pf-v6-menu');
        elementFilter.appendChild(menu);
      }

      const existingItems = menu.querySelectorAll('pf-v6-menu-item');
      existingItems.forEach(item => item.remove());

      const allElements = new Set<string>();
      for (const tagName of this.#elementEventMap.keys()) {
        if (this.#discoveredElements.has(tagName)) {
          allElements.add(tagName);
        }
      }

      if (savedPreferences.elements) {
        this.#elementFilters = (savedPreferences.elements as Set<string> & { intersection: (other: Set<string>) => Set<string> }).intersection(allElements);
      } else {
        this.#elementFilters = new Set(allElements);
      }

      for (const tagName of allElements) {
        const item = document.createElement('pf-v6-menu-item');
        item.setAttribute('variant', 'checkbox');
        item.setAttribute('value', tagName);
        if (this.#elementFilters.has(tagName)) {
          item.setAttribute('checked', '');
        }
        item.textContent = `<${tagName}>`;
        menu.appendChild(item);
      }
    }
  }

  #handleEventTypeFilterChange = (event: Event) => {
    const { value, checked } = event as Event & { value: string; checked: boolean };

    if (!value) return;

    if (checked) {
      this.#eventTypeFilters.add(value);
    } else {
      this.#eventTypeFilters.delete(value);
    }

    this.#saveEventFilters();
    this.#filterEvents(this.#eventsFilterValue);
  };

  #handleElementFilterChange = (event: Event) => {
    const { value, checked } = event as Event & { value: string; checked: boolean };

    if (!value) return;

    if (checked) {
      this.#elementFilters.add(value);
    } else {
      this.#elementFilters.delete(value);
    }

    this.#saveEventFilters();
    this.#filterEvents(this.#eventsFilterValue);
  };

  #loadEventFiltersFromStorage(): { eventTypes: Set<string> | null; elements: Set<string> | null } {
    const preferences: { eventTypes: Set<string> | null; elements: Set<string> | null } = {
      eventTypes: null,
      elements: null
    };

    try {
      const savedEventTypes = localStorage.getItem('cem-serve-event-type-filters');
      if (savedEventTypes) {
        preferences.eventTypes = new Set(JSON.parse(savedEventTypes));
      }

      const savedElements = localStorage.getItem('cem-serve-element-filters');
      if (savedElements) {
        preferences.elements = new Set(JSON.parse(savedElements));
      }
    } catch (e) {
      console.debug('[cem-serve-chrome] localStorage unavailable for event filters');
    }

    return preferences;
  }

  #saveEventFilters() {
    try {
      localStorage.setItem('cem-serve-event-type-filters',
        JSON.stringify([...this.#eventTypeFilters]));
      localStorage.setItem('cem-serve-element-filters',
        JSON.stringify([...this.#elementFilters]));
    } catch (e) {
      // localStorage unavailable (private mode), silently continue
    }
  }

  #clearEvents() {
    this.#capturedEvents = [];
    this.#selectedEventId = null;
    if (this.#eventList) {
      this.#eventList.replaceChildren();
    }
    if (this.#eventDetailHeader) {
      this.#eventDetailHeader.innerHTML = '';
    }
    if (this.#eventDetailBody) {
      this.#eventDetailBody.innerHTML = '';
    }
  }

  async #copyEvents() {
    if (!this.#eventList) return;

    const visibleEvents = Array.from(this.#eventList.children)
      .filter(entry => !(entry as HTMLElement).hidden)
      .map(entry => {
        const id = (entry as HTMLElement).dataset.eventId!;
        return this.#getEventRecordById(id);
      })
      .filter((event): event is EventRecord => !!event)
      .map(event => {
        const time = event.timestamp.toLocaleTimeString();
        const element = event.elementId ? `#${event.elementId}` : event.tagName;
        const props = event.customProperties && Object.keys(event.customProperties).length > 0
          ? ` - ${JSON.stringify(event.customProperties)}`
          : '';
        return `[${time}] <${event.tagName}> ${element} \u2192 ${event.eventName}${props}`;
      })
      .join('\n');

    if (!visibleEvents) return;

    try {
      await navigator.clipboard.writeText(visibleEvents);
      const btn = this.#$('copy-events');
      if (btn) {
        const textNode = Array.from(btn.childNodes).find(
          n => n.nodeType === Node.TEXT_NODE && (n.textContent?.trim().length ?? 0) > 0
        );
        if (textNode) {
          const original = textNode.textContent;
          textNode.textContent = 'Copied!';

          if (this.#copyEventsFeedbackTimeout) {
            clearTimeout(this.#copyEventsFeedbackTimeout);
          }

          this.#copyEventsFeedbackTimeout = setTimeout(() => {
            if (this.isConnected && textNode.parentNode) {
              textNode.textContent = original;
            }
            this.#copyEventsFeedbackTimeout = null;
          }, 2000);
        }
      }
    } catch (err) {
      console.error('[cem-serve-chrome] Failed to copy events:', err);
    }
  }

  #setupEventListeners() {
    this.#eventList = this.#$('event-list');
    this.#eventDetailHeader = this.#$('event-detail-header');
    this.#eventDetailBody = this.#$('event-detail-body');

    if (this.#eventList) {
      this.#eventList.addEventListener('click', (e: Event) => {
        const listItem = (e.target as Element).closest('.event-list-item') as HTMLElement;
        if (listItem) {
          const eventId = listItem.dataset.eventId;
          if (eventId) {
            this.#selectEvent(eventId);
          }
        }
      });
    }

    const eventsFilter = this.#$('events-filter');
    if (eventsFilter) {
      eventsFilter.addEventListener('input', (e: Event) => {
        const { value = '' } = e.target as HTMLInputElement;
        clearTimeout(this.#eventsFilterDebounceTimer!);
        this.#eventsFilterDebounceTimer = setTimeout(() => {
          this.#filterEvents(value);
        }, 300);
      });
    }

    const eventTypeFilter = this.#$('event-type-filter');
    if (eventTypeFilter) {
      eventTypeFilter.addEventListener('select', this.#handleEventTypeFilterChange as EventListener);
    }

    const elementFilter = this.#$('element-filter');
    if (elementFilter) {
      elementFilter.addEventListener('select', this.#handleElementFilterChange as EventListener);
    }

    this.#$('clear-events')?.addEventListener('click', () => {
      this.#clearEvents();
    });

    this.#$('copy-events')?.addEventListener('click', () => {
      this.#copyEvents();
    });
  }

  disconnectedCallback() {
    super.disconnectedCallback();

    // Clean up knob listeners
    this.removeEventListener('knob:attribute-change', this.#onKnobChange);
    this.removeEventListener('knob:property-change', this.#onKnobChange);
    this.removeEventListener('knob:css-property-change', this.#onKnobChange);
    this.removeEventListener('knob:attribute-clear', this.#onKnobClear);
    this.removeEventListener('knob:property-clear', this.#onKnobClear);
    this.removeEventListener('knob:css-property-clear', this.#onKnobClear);

    // Clean up tree state listeners
    if (this.#handleTreeExpand) {
      this.removeEventListener('expand', this.#handleTreeExpand);
    }
    if (this.#handleTreeCollapse) {
      this.removeEventListener('collapse', this.#handleTreeCollapse);
    }
    if (this.#handleTreeSelect) {
      this.removeEventListener('select', this.#handleTreeSelect);
    }

    // Clean up window listener
    if (this.#handleLogsEvent) {
      window.removeEventListener('cem:logs', this.#handleLogsEvent);
    }

    // Clear pending feedback timeouts
    if (this.#copyLogsFeedbackTimeout) {
      clearTimeout(this.#copyLogsFeedbackTimeout);
      this.#copyLogsFeedbackTimeout = null;
    }
    if (this.#copyDebugFeedbackTimeout) {
      clearTimeout(this.#copyDebugFeedbackTimeout);
      this.#copyDebugFeedbackTimeout = null;
    }
    if (this.#copyEventsFeedbackTimeout) {
      clearTimeout(this.#copyEventsFeedbackTimeout);
      this.#copyEventsFeedbackTimeout = null;
    }

    // Disconnect mutation observer
    this.#observer.disconnect();

    // Close WebSocket connection
    if (this.#wsClient) {
      this.#wsClient.destroy();
    }
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'cem-serve-chrome': CemServeChrome;
  }
}
