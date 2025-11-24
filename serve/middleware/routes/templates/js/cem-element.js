/**
 * Base class for CEM web components that use template loading.
 *
 * Handles the common pattern of:
 * - SSR-aware shadow root creation
 * - Automatic template loading with caching
 * - Applying stylesheets via adoptedStyleSheets
 * - CSS Anchor Positioning polyfill for cross-browser support
 * - Error handling with console warnings
 *
 * Subclasses can override:
 * - static shadowRootOptions = { mode: 'open', delegatesFocus: false }
 * - static elementName = 'custom-name' (defaults to this.localName)
 * - async afterTemplateLoaded() - called after template is loaded and applied
 *
 * @example
 * // Simple element with no additional logic
 * export class MyElement extends CemElement {}
 *
 * @example
 * // Element with setup logic
 * export class MyElement extends CemElement {
 *   async afterTemplateLoaded() {
 *     this.button = this.shadowRoot.getElementById('button');
 *     this.button.addEventListener('click', this.#handleClick);
 *   }
 * }
 *
 * @example
 * // Element with custom shadow root options
 * export class MyElement extends CemElement {
 *   static shadowRootOptions = { mode: 'open', delegatesFocus: true };
 * }
 */
export class CemElement extends HTMLElement {
  /** Cache for Constructable Stylesheet objects shared across all instances */
  static #stylesheetCache = new Map();

  /** Cache for HTML templates shared across all instances */
  static #templateCache = new Map();

  /** Promise for loading the anchor positioning polyfill */
  static #anchorPolyfillPromise = null;

  /** Whether anchor positioning is supported natively */
  static #anchorPositioningSupported = null;

  /**
   * Shadow root options. Override in subclass to customize.
   * @type {{ mode: 'open' | 'closed', delegatesFocus?: boolean }}
   */
  static shadowRootOptions = { mode: 'open' };

  /**
   * Element name used for template loading.
   * Defaults to the element's tag name (this.localName).
   * Override if the template name differs from the tag name.
   * @type {string | null}
   */
  static elementName = null;

  /**
   * Check if CSS Anchor Positioning is supported natively
   * @returns {boolean} Whether anchor positioning is supported
   * @private
   */
  static #supportsAnchorPositioning() {
    if (this.#anchorPositioningSupported !== null) {
      return this.#anchorPositioningSupported;
    }

    // Check for CSS.supports() API
    if (!('CSS' in window) || typeof CSS.supports !== 'function') {
      this.#anchorPositioningSupported = false;
      return false;
    }

    // Check for anchor-name property support
    this.#anchorPositioningSupported = CSS.supports('anchor-name', '--test');
    return this.#anchorPositioningSupported;
  }

  /**
   * Load the CSS Anchor Positioning polyfill if needed
   * @returns {Promise<Function|null>} Resolves with polyfill function or null if not needed
   * @private
   */
  static async #loadAnchorPolyfill() {
    // If native support exists, no need for polyfill
    if (this.#supportsAnchorPositioning()) {
      return null;
    }

    // If polyfill is already loading/loaded, return that promise
    if (this.#anchorPolyfillPromise) {
      return this.#anchorPolyfillPromise;
    }

    // Load the polyfill script (function variant)
    this.#anchorPolyfillPromise = (async () => {
      try {
        const module = await import('/__cem/css-anchor-positioning-fn.js');
        // The -fn variant exports the polyfill function as default
        return module.default || null;
      } catch (error) {
        console.error('Failed to load CSS Anchor Positioning polyfill:', error);
        throw error;
      }
    })();

    return this.#anchorPolyfillPromise;
  }

  /**
   * Apply the anchor positioning polyfill to a shadow root
   * @param {ShadowRoot} shadowRoot - The shadow root to polyfill
   * @param {HTMLElement} element - The element for debugging context
   * @private
   */
  static async #applyAnchorPolyfill(shadowRoot, element) {
    // Only apply if polyfill is needed
    if (this.#supportsAnchorPositioning()) {
      return;
    }

    const polyfill = await this.#loadAnchorPolyfill();

    // Apply polyfill to the shadow root if the function is available
    if (typeof polyfill === 'function') {
      try {
        await polyfill({ roots: [shadowRoot] });
      } catch (error) {
        console.error('Failed to apply CSS Anchor Positioning polyfill:', error);
      }
    }
  }

  /**
   * Fetch element HTML or CSS from the server
   * @param {string} name - Component name (e.g., 'pf-v6-button')
   * @param {'html'|'css'} type - File type
   * @returns {Promise<string>} The content
   * @private
   */
  static async #fetchText(name, type) {
    const prettyType = type.toUpperCase();
    try {
      const response = await fetch(`/__cem/elements/${name}/${name}.${type}`);
      if (!response.ok) {
        throw new Error(`Failed to fetch ${prettyType}: ${response.status} ${response.statusText}`);
      }
      return await response.text();
    } catch (error) {
      console.error(`Failed to load ${prettyType} for ${name}:`, error);
      throw error;
    }
  }

  /**
   * Gets an existing stylesheet from cache or creates a new one
   * @param {string} name - Component name (e.g., 'pf-v6-button')
   * @returns {Promise<CSSStyleSheet>} The cached or newly created stylesheet
   * @private
   */
  static async #loadCSS(name) {
    if (this.#stylesheetCache.has(name)) {
      return this.#stylesheetCache.get(name);
    }

    const css = await this.#fetchText(name, 'css');
    try {
      const sheet = new CSSStyleSheet();
      await sheet.replace(css);
      this.#stylesheetCache.set(name, sheet);
      return sheet;
    } catch (error) {
      console.error(`Failed to construct stylesheet for ${name}:`, error);
      throw error;
    }
  }

  /**
   * Gets an element template from client-side cache or from the server
   * @param {string} name - Component name (e.g., 'pf-v6-button')
   * @returns {Promise<string>} The cached or fetched template string
   * @private
   */
  static async #loadHTML(name) {
    if (this.#templateCache.has(name)) {
      return this.#templateCache.get(name);
    }
    const html = await this.#fetchText(name, 'html');
    this.#templateCache.set(name, html);
    return html;
  }

  /**
   * Loads a complete component template (HTML and CSS)
   * @param {string} name - Component name (e.g., 'pf-v6-button')
   * @returns {Promise<{html: string, stylesheet: CSSStyleSheet}>} The template HTML and stylesheet
   * @private
   */
  static async #loadComponentTemplate(name) {
    const [html, stylesheet] = await Promise.all([this.#loadHTML(name), this.#loadCSS(name)]);
    return { html, stylesheet };
  }

  constructor() {
    super();
    const options = this.constructor.shadowRootOptions ?? { mode: 'open' };

    // SSR-aware: only create shadow root if it doesn't exist
    if (!this.shadowRoot) {
      this.attachShadow(options);
    }
  }

  async connectedCallback() {
    // Only populate if shadow root is empty
    if (!this.shadowRoot.firstChild) {
      await this.#populateShadowRoot();
    }

    // Apply anchor positioning polyfill to this shadow root
    await CemElement.#applyAnchorPolyfill(this.shadowRoot, this);

    // Call lifecycle hook for subclasses
    await this.afterTemplateLoaded?.();
  }

  /**
   * Load and apply the component template.
   * Uses the element name from static elementName or falls back to this.localName.
   */
  async #populateShadowRoot() {
    const elementName = this.constructor.elementName || this.localName;

    try {
      // Use CemElement explicitly since private static methods aren't inherited
      const { html, stylesheet } = await CemElement.#loadComponentTemplate(elementName);

      // If anchor positioning polyfill is needed, use <style> tag instead of adoptedStyleSheets
      // because the polyfill doesn't support constructed stylesheets
      if (!CemElement.#supportsAnchorPositioning()) {
        // Get the CSS text from the stylesheet
        const cssText = await CemElement.#fetchText(elementName, 'css');
        this.shadowRoot.innerHTML = `<style>${cssText}</style>${html}`;
      } else {
        // Use adoptedStyleSheets for better performance when native support exists
        this.shadowRoot.adoptedStyleSheets = [stylesheet];
        this.shadowRoot.innerHTML = html;
      }
    } catch (error) {
      console.error(`Failed to load ${elementName} template:`, error);
    }
  }

  /**
   * Lifecycle hook called after template is loaded and applied.
   * Override this in subclasses to set up event listeners, DOM references, etc.
   *
   * @example
   * async afterTemplateLoaded() {
   *   this.button = this.shadowRoot.getElementById('button');
   *   this.button.addEventListener('click', () => console.log('clicked'));
   * }
   *
   * @protected
   */
  async afterTemplateLoaded() {
    // Subclasses can override this
  }
}
