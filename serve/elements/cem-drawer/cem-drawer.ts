import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';
import { property } from 'lit/decorators/property.js';

import styles from './cem-drawer.css' with { type: 'css' };

import '../pf-v6-button/pf-v6-button.js';

/**
 * Custom event fired when the drawer opens or closes
 */
export class CemDrawerChangeEvent extends Event {
  open: boolean;
  constructor(open: boolean) {
    super('change', { bubbles: true });
    this.open = open;
  }
}

/**
 * Custom event fired when the drawer is resized
 */
export class CemDrawerResizeEvent extends Event {
  height: number;
  constructor(height: number) {
    super('resize', { bubbles: true });
    this.height = height;
  }
}

/**
 * CEM Serve Drawer - Collapsible, resizable drawer component
 *
 * @slot - Default slot for drawer content
 *
 * @fires change - Fires when the drawer opens or closes
 * @fires resize - Fires when the drawer is resized
 */
@customElement('cem-drawer')
export class CemServeDrawer extends LitElement {
  static styles = styles;

  @property({ type: Boolean, reflect: true })
  accessor open = false;

  @property({ type: Number, reflect: true, attribute: 'drawer-height' })
  accessor drawerHeight = 400;

  #isDragging = false;
  #startY = 0;
  #startHeight = 0;
  #maxHeight: number | null = null;
  #rafId: number | null = null;
  #pendingHeight: number | null = null;

  #$(id: string) {
    return this.shadowRoot?.getElementById(id);
  }

  /**
   * Returns the maximum safe height for the drawer content in pixels.
   *
   * The toggle button must always remain visible below the masthead so the
   * user can close or resize the drawer at any time.
   */
  #getMaxHeight(): number {
    const toggle = this.#$('toggle');
    const handle = this.#$('resize-handle');
    const mastheadHeight = 56;
    const toggleHeight = toggle?.offsetHeight ?? 32;
    const handleHeight = handle?.offsetHeight ?? 4;
    return Math.max(100, window.innerHeight - mastheadHeight - toggleHeight - handleHeight);
  }

  connectedCallback() {
    super.connectedCallback();
    window.addEventListener('resize', this.#handleWindowResize);
  }

  disconnectedCallback() {
    super.disconnectedCallback();
    window.removeEventListener('resize', this.#handleWindowResize);
  }

  render() {
    return html`
      <div id="resize-handle"
           role="separator"
           aria-orientation="horizontal"
           aria-label="Resize drawer"
           tabindex="0"
           aria-controls="content"
           aria-valuemin="100"
           aria-valuemax="1000"
           aria-valuenow="400"
           @mousedown=${this.#startResize}
           @keydown=${this.#handleKeydown}></div>
      <pf-v6-button id="toggle"
                     variant="plain"
                     aria-label="Toggle drawer"
                     aria-expanded="${this.open}"
                     aria-controls="content"
                     @click=${this.#onToggleClick}>
        <svg width="16"
             height="16"
             viewBox="0 0 16 16"
             fill="currentColor"
             role="presentation">
          <path fill-rule="evenodd"
                d="M1.646 4.646a.5.5 0 0 1 .708 0L8 10.293l5.646-5.647a.5.5 0 0 1 .708.708l-6 6a.5.5 0 0 1-.708 0l-6-6a.5.5 0 0 1 0-.708z"/>
        </svg>
      </pf-v6-button>
      <div id="content"
           style="height: ${this.open ? this.drawerHeight : 0}px;">
        <slot></slot>
      </div>
    `;
  }

  updated(changed: Map<string, unknown>) {
    if (changed.has('open')) {
      const content = this.#$('content');
      if (content && this.open) {
        let height = this.drawerHeight;
        let needsPersist = false;
        if (height <= 0 || isNaN(height)) {
          height = 400;
          needsPersist = true;
        }
        const maxHeight = this.#getMaxHeight();
        if (height > maxHeight) {
          height = maxHeight;
          needsPersist = true;
        }
        if (needsPersist) {
          this.drawerHeight = height;
          this.dispatchEvent(new CemDrawerResizeEvent(height));
        }
      }
      this.dispatchEvent(new CemDrawerChangeEvent(this.open));
    }
  }

  #onToggleClick = () => {
    const content = this.#$('content');
    if (content) {
      content.classList.add('transitions-enabled');
    }
    this.toggle();
  };

  toggle() {
    this.open = !this.open;
  }

  openDrawer() {
    this.open = true;
  }

  close() {
    this.open = false;
  }

  #startResize = (e: MouseEvent) => {
    this.#isDragging = true;
    this.#startY = e.clientY;
    const content = this.#$('content');
    if (!content) return;
    this.#startHeight = content.offsetHeight;
    this.#maxHeight = this.#getMaxHeight();

    content.classList.remove('transitions-enabled');
    content.classList.add('resizing');

    document.addEventListener('mousemove', this.#handleResize, { passive: true });
    document.addEventListener('mouseup', this.#stopResize);

    e.preventDefault();
  };

  #handleResize = (e: MouseEvent) => {
    if (!this.#isDragging) return;

    const deltaY = this.#startY - e.clientY;
    const newHeight = Math.max(100, Math.min(this.#maxHeight!, this.#startHeight + deltaY));

    this.#pendingHeight = newHeight;

    if (!this.#rafId) {
      this.#rafId = requestAnimationFrame(this.#applyResize);
    }
  };

  #applyResize = () => {
    if (this.#pendingHeight === null) return;

    const content = this.#$('content');
    if (content) {
      content.style.height = `${this.#pendingHeight}px`;
    }

    this.#pendingHeight = null;
    this.#rafId = null;
  };

  #handleKeydown = (e: KeyboardEvent) => {
    const content = this.#$('content');
    if (!content) return;

    const step = e.shiftKey ? 50 : 10;
    let currentHeight = parseInt(content.style.height, 10) || 400;
    let newHeight = currentHeight;

    switch (e.key) {
      case 'ArrowUp':
        e.preventDefault();
        newHeight = currentHeight + step;
        break;
      case 'ArrowDown':
        e.preventDefault();
        newHeight = currentHeight - step;
        break;
      case 'Home':
        e.preventDefault();
        newHeight = 100;
        break;
      case 'End':
        e.preventDefault();
        newHeight = this.#getMaxHeight();
        break;
      default:
        return;
    }

    newHeight = Math.max(100, Math.min(this.#getMaxHeight(), newHeight));

    content.style.height = `${newHeight}px`;
    this.#updateAriaValueNow(newHeight);
    this.dispatchEvent(new CemDrawerResizeEvent(newHeight));
  };

  #updateAriaValueNow(height: number) {
    const resizeHandle = this.#$('resize-handle');
    if (resizeHandle) {
      resizeHandle.setAttribute('aria-valuenow', String(Math.round(height)));
      resizeHandle.setAttribute('aria-valuemax', String(Math.round(this.#getMaxHeight())));
    }
  }

  #stopResize = () => {
    this.#isDragging = false;

    if (this.#rafId) {
      cancelAnimationFrame(this.#rafId);
      this.#rafId = null;
    }

    const content = this.#$('content');
    if (!content) return;
    content.classList.add('transitions-enabled');
    content.classList.remove('resizing');

    const height = parseInt(content.style.height, 10);
    this.#updateAriaValueNow(height);

    this.dispatchEvent(new CemDrawerResizeEvent(height));

    document.removeEventListener('mousemove', this.#handleResize);
    document.removeEventListener('mouseup', this.#stopResize);
  };

  #handleWindowResize = () => {
    if (!this.open) return;
    const content = this.#$('content');
    if (!content) return;
    const currentHeight = parseInt(content.style.height) || 0;
    const maxHeight = this.#getMaxHeight();
    if (currentHeight > maxHeight) {
      content.style.height = `${maxHeight}px`;
      this.drawerHeight = maxHeight;
      this.dispatchEvent(new CemDrawerResizeEvent(maxHeight));
    }
  };
}

declare global {
  interface HTMLElementTagNameMap {
    'cem-drawer': CemServeDrawer;
  }
}
