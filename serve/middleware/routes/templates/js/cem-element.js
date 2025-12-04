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
 * - static is = 'custom-name'
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

  /** Cache for complete component templates (HTML + stylesheet) by tag name and attrs */
  static #componentCache = new Map();

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
  static is = null;

  /**
   * Fetch element HTML or CSS from the server
   * @param {string} name - Component name (e.g., 'pf-v6-button')
   * @param {'html'|'css'} type - File type
   * @param {URLSearchParams} [params] - Request Params
   * @returns {Promise<string>} The content
   */
  static async #fetchText(name, type, params) {
    const prettyType = type.toUpperCase();
    try {
      const url = new URL(`/__cem/elements/${name}/${name}.${type}`, window.location.origin);
      if (params) {
        for (const [key, value] of params) {
          url.searchParams.append(key, value);
        }
      }
      const response = await fetch(url);
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
   */
  static async #loadCSS(name) {
    if (this.#stylesheetCache.has(name)) {
      return this.#stylesheetCache.get(name);
    }

    // Cache the Promise immediately to prevent race conditions
    const promise = (async () => {
      const css = await this.#fetchText(name, 'css');
      try {
        const sheet = new CSSStyleSheet();
        await sheet.replace(css);
        return sheet;
      } catch (error) {
        console.error(`Failed to construct stylesheet for ${name}:`, error);
        throw error;
      }
    })();

    this.#stylesheetCache.set(name, promise);
    return promise;
  }

  /**
   * Gets an element template from client-side cache or from the server
   * @param {string} name - Component name (e.g., 'pf-v6-button')
   * @param {HTMLElement} element - The element instance with attributes
   * @returns {Promise<string>} The cached or fetched template string
   */
  static async #loadHTML(name, element) {
    const cacheKey = `${name}:${this.#getAttrCacheKey(element)}`;

    if (this.#templateCache.has(cacheKey)) {
      return this.#templateCache.get(cacheKey);
    }

    // Cache the Promise immediately to prevent race conditions
    const promise = (async () => {
      // Build URL with query parameters
      const params = new URLSearchParams();
      params.set('content', 'shadow');

      // Add element attributes to query
      if (element) {
        for (const attr of element.attributes) {
          params.set(`attrs[${attr.name}]`, attr.value);
        }
      }

      return await this.#fetchText(name, 'html', params);
    })();

    this.#templateCache.set(cacheKey, promise);
    return promise;
  }

  /**
   * Generate cache key from element attributes
   * @param {HTMLElement} element - The element instance
   * @returns {string} Cache key based on sorted attributes
   */
  static #getAttrCacheKey(element) {
    if (!element) return '';
    return Array.from(element.attributes)
      .map(a => `${a.name}=${a.value}`)
      .sort()
      .join('&');
  }

  /**
   * Loads a complete component template (HTML and CSS)
   * @param {string} name - Component name (e.g., 'pf-v6-button')
   * @param {HTMLElement} element - The element instance with attributes
   * @returns {Promise<{html: string, stylesheet: CSSStyleSheet}>} The template HTML and stylesheet
   */
  static async #loadComponentTemplate(name, element) {
    const cacheKey = `${name}:${this.#getAttrCacheKey(element)}`;

    // Check unified component cache first
    if (this.#componentCache.has(cacheKey)) {
      return this.#componentCache.get(cacheKey);
    }

    // Cache the Promise immediately to prevent race conditions
    const promise = (async () => {
      // Load HTML and CSS in parallel
      const [html, stylesheet] = await Promise.all([
        this.#loadHTML(name, element),
        this.#loadCSS(name),
      ]);
      return { html, stylesheet };
    })();

    // Cache the complete component template
    this.#componentCache.set(cacheKey, promise);

    return promise;
  }

  /** @type {Promise<void>} Promise that resolves after first render completes */
  rendered;
  #resolveRendered;

  constructor() {
    super();
    const options = this.constructor.shadowRootOptions ?? { mode: 'open' };

    // Create rendered promise for testing
    this.rendered = new Promise(resolve => {
      this.#resolveRendered = resolve;
    });

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

    // Call lifecycle hook for subclasses (works for both SSR and CSR)
    await this.afterTemplateLoaded?.();

    // Resolve rendered promise for testing (works for both SSR and CSR)
    this.#resolveRendered?.();
  }

  /**
   * Load and apply the component template.
   * Uses the element name from `static is` or falls back to this.localName.
   */
  async #populateShadowRoot() {
    const elementName = this.constructor.is || this.localName;

    try {
      // Use CemElement explicitly since private static methods aren't inherited
      // Pass 'this' to load template with current element's attributes
      const { html, stylesheet } = await CemElement.#loadComponentTemplate(elementName, this);
      this.shadowRoot.adoptedStyleSheets = [stylesheet];
      this.shadowRoot.innerHTML = html;
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
