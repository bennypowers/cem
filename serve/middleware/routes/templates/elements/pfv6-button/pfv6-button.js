/**
 * PatternFly v6 inspired button component
 *
 * @attr {string} variant - Button variant: primary, secondary, tertiary (default), danger, plain, link
 * @attr {string} size - Button size: sm, md (default), lg
 * @attr {boolean} block - Make button full width
 * @attr {boolean} disabled - Disable the button
 * @attr {string} aria-label - Accessible label for the button
 *
 * @slot - Default slot for button text/content
 * @slot icon-start - Slot for icon before text
 * @slot icon-end - Slot for icon after text
 *
 * @fires click - Bubbles click events from internal button
 */
class Pfv6Button extends HTMLElement {
  constructor() {
    super();

    // Don't recreate shadow root if it already exists (from SSR)
    if (!this.shadowRoot) {
      this.attachShadow({ mode: 'open' });
    }
  }

  async connectedCallback() {
    // If shadow root is empty (client-side rendering), populate it
    let button = this.shadowRoot?.querySelector('button');
    if (!button && this.shadowRoot) {
      await this.#populateShadowRoot();
      button = this.shadowRoot.querySelector('button');
    }

    if (!button) return;

    // Forward attributes to internal button
    this.#syncAttributes(button);

    // Delegate button events
    button.addEventListener('click', (e) => {
      if (this.disabled) {
        e.preventDefault();
        e.stopPropagation();
        return;
      }
      // Event bubbles naturally through shadow DOM
    });

    // Observe attribute changes
    this.#observer = new MutationObserver(() => {
      this.#syncAttributes(button);
    });

    this.#observer.observe(this, {
      attributes: true,
      attributeFilter: ['disabled', 'aria-label', 'aria-expanded', 'aria-controls']
    });
  }

  disconnectedCallback() {
    this.#observer?.disconnect();
  }

  async #populateShadowRoot() {
    // Fetch HTML template and CSS for client-side rendering
    const [htmlResponse, cssResponse] = await Promise.all([
      fetch('/__cem/elements/pfv6-button/pfv6-button.html'),
      fetch('/__cem/elements/pfv6-button/pfv6-button.css')
    ]);

    const [html, css] = await Promise.all([
      htmlResponse.text(),
      cssResponse.text()
    ]);

    // Populate shadow root
    this.shadowRoot.innerHTML = `
      <style>${css}</style>
      ${html}
    `;
  }

  #syncAttributes(button) {
    // Sync disabled
    if (this.hasAttribute('disabled')) {
      button.disabled = true;
    } else {
      button.disabled = false;
    }

    // Sync aria attributes
    const ariaAttrs = ['aria-label', 'aria-expanded', 'aria-controls', 'aria-haspopup'];
    ariaAttrs.forEach(attr => {
      if (this.hasAttribute(attr)) {
        button.setAttribute(attr, this.getAttribute(attr));
      } else {
        button.removeAttribute(attr);
      }
    });

    // Sync type attribute
    if (this.hasAttribute('type')) {
      button.setAttribute('type', this.getAttribute('type'));
    }
  }

  get disabled() {
    return this.hasAttribute('disabled');
  }

  set disabled(value) {
    if (value) {
      this.setAttribute('disabled', '');
    } else {
      this.removeAttribute('disabled');
    }
  }

  #observer;
}

customElements.define('pfv6-button', Pfv6Button);
