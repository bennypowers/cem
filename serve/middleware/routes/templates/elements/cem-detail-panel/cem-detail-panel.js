import { CemElement } from '/__cem/cem-element.js';

/**
 * Detail panel that renders manifest item details on-demand
 * @customElement cem-detail-panel
 */
export class CemDetailPanel extends CemElement {
  static is = 'cem-detail-panel';

  #cache = new Map();
  #currentItemId = null;

  afterTemplateLoaded() {
    // Template is loaded, ready to render items
  }

  /**
   * Render details for a manifest item
   * @param {Object} item - The manifest item to render
   * @param {Object} manifest - The full manifest for context
   */
  async renderItem(item, manifest) {
    if (!item) {
      this.#clearContent();
      return;
    }

    const itemId = this.#getItemId(item);
    this.#currentItemId = itemId;

    // Check cache first
    if (this.#cache.has(itemId)) {
      this.#setContent(this.#cache.get(itemId));
      return;
    }

    // Build the detail HTML
    const html = await this.#buildDetailHTML(item, manifest);
    this.#cache.set(itemId, html);
    this.#setContent(html);
  }

  /**
   * Generate a unique ID for a manifest item
   */
  #getItemId(item) {
    const parts = [item.type];
    if (item.modulePath) parts.push(item.modulePath);
    if (item.tagName) parts.push(item.tagName);
    if (item.name) parts.push(item.name);
    return parts.join(':');
  }

  /**
   * Build HTML for detail panel based on item type
   */
  async #buildDetailHTML(item, manifest) {
    const type = item.type;

    switch (type) {
      case 'module':
        return this.#buildModuleDetails(item, manifest);
      case 'custom-element':
        return this.#buildCustomElementDetails(item, manifest);
      case 'attribute':
        return this.#buildAttributeDetails(item, manifest);
      case 'property':
        return this.#buildPropertyDetails(item, manifest);
      case 'method':
        return this.#buildMethodDetails(item, manifest);
      case 'event':
        return this.#buildEventDetails(item, manifest);
      case 'slot':
        return this.#buildSlotDetails(item, manifest);
      case 'class':
        return this.#buildClassDetails(item, manifest);
      case 'function':
        return this.#buildFunctionDetails(item, manifest);
      case 'variable':
        return this.#buildVariableDetails(item, manifest);
      case 'mixin':
        return this.#buildMixinDetails(item, manifest);
      default:
        return `<div class="empty-state">No details available for ${type}</div>`;
    }
  }

  async #buildModuleDetails(item, manifest) {
    const module = this.#findModule(manifest, item.modulePath);
    if (!module) return '<div class="empty-state">Module not found</div>';

    const summary = module.summary ? await this.#renderMarkdown(module.summary) : '';
    const description = module.description ? await this.#renderMarkdown(module.description) : '';

    return `
      <h3>${this.#escapeHtml(module.path)}</h3>
      <dl class="pf-v6-c-description-list pf-m-horizontal pf-m-compact">
        ${summary ? `<div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Summary</dt><dd class="pf-v6-c-description-list__description">${summary}</dd></div>` : ''}
        ${description ? `<div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Description</dt><dd class="pf-v6-c-description-list__description">${description}</dd></div>` : ''}
        <div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Declarations</dt><dd class="pf-v6-c-description-list__description">${module.declarations?.length || 0}</dd></div>
      </dl>
    `;
  }

  async #buildCustomElementDetails(item, manifest) {
    const ce = this.#findCustomElement(manifest, item.modulePath, item.tagName);
    if (!ce) return '<div class="empty-state">Custom element not found</div>';

    const summary = ce.summary ? await this.#renderMarkdown(ce.summary) : '';
    const description = ce.description ? await this.#renderMarkdown(ce.description) : '';

    return `
      <h3>&lt;${this.#escapeHtml(ce.tagName)}&gt;</h3>
      <dl class="pf-v6-c-description-list pf-m-horizontal pf-m-compact">
        <div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Name</dt><dd class="pf-v6-c-description-list__description">${this.#escapeHtml(ce.name)}</dd></div>
        ${summary ? `<div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Summary</dt><dd class="pf-v6-c-description-list__description">${summary}</dd></div>` : ''}
        ${description ? `<div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Description</dt><dd class="pf-v6-c-description-list__description">${description}</dd></div>` : ''}
        ${ce.superclass?.name ? `<div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Extends</dt><dd class="pf-v6-c-description-list__description"><code>${this.#escapeHtml(ce.superclass.name)}</code></dd></div>` : ''}
        <div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Attributes</dt><dd class="pf-v6-c-description-list__description">${ce.attributes?.length || 0}</dd></div>
        <div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Events</dt><dd class="pf-v6-c-description-list__description">${ce.events?.length || 0}</dd></div>
        <div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Slots</dt><dd class="pf-v6-c-description-list__description">${ce.slots?.length || 0}</dd></div>
      </dl>
    `;
  }

  async #buildAttributeDetails(item, manifest) {
    const ce = this.#findCustomElement(manifest, item.modulePath, item.tagName);
    if (!ce) return '<div class="empty-state">Custom element not found</div>';

    const attr = ce.attributes?.find(a => a.name === item.name);
    if (!attr) return '<div class="empty-state">Attribute not found</div>';

    const summary = attr.summary ? await this.#renderMarkdown(attr.summary) : '';
    const description = attr.description ? await this.#renderMarkdown(attr.description) : '';

    return `
      <h3>${this.#escapeHtml(attr.name)}</h3>
      <dl class="pf-v6-c-description-list pf-m-horizontal pf-m-compact">
        <div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Element</dt><dd class="pf-v6-c-description-list__description"><code>&lt;${this.#escapeHtml(ce.tagName)}&gt;</code></dd></div>
        ${attr.type?.text ? `<div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Type</dt><dd class="pf-v6-c-description-list__description"><code>${this.#escapeHtml(attr.type.text)}</code></dd></div>` : ''}
        ${summary ? `<div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Summary</dt><dd class="pf-v6-c-description-list__description">${summary}</dd></div>` : ''}
        ${description ? `<div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Description</dt><dd class="pf-v6-c-description-list__description">${description}</dd></div>` : ''}
        ${attr.default ? `<div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Default</dt><dd class="pf-v6-c-description-list__description"><code>${this.#escapeHtml(attr.default)}</code></dd></div>` : ''}
        ${attr.fieldName ? `<div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Reflects to</dt><dd class="pf-v6-c-description-list__description"><code>${this.#escapeHtml(attr.fieldName)}</code></dd></div>` : ''}
      </dl>
    `;
  }

  async #buildPropertyDetails(item, manifest) {
    const ce = this.#findCustomElement(manifest, item.modulePath, item.tagName);
    if (!ce) return '<div class="empty-state">Custom element not found</div>';

    const prop = ce.members?.find(m => m.kind === 'field' && m.name === item.name);
    if (!prop) return '<div class="empty-state">Property not found</div>';

    const summary = prop.summary ? await this.#renderMarkdown(prop.summary) : '';
    const description = prop.description ? await this.#renderMarkdown(prop.description) : '';

    return `
      <h3>${this.#escapeHtml(prop.name)}</h3>
      <dl class="pf-v6-c-description-list pf-m-horizontal pf-m-compact">
        <div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Element</dt><dd class="pf-v6-c-description-list__description"><code>&lt;${this.#escapeHtml(ce.tagName)}&gt;</code></dd></div>
        ${prop.type?.text ? `<div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Type</dt><dd class="pf-v6-c-description-list__description"><code>${this.#escapeHtml(prop.type.text)}</code></dd></div>` : ''}
        ${summary ? `<div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Summary</dt><dd class="pf-v6-c-description-list__description">${summary}</dd></div>` : ''}
        ${description ? `<div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Description</dt><dd class="pf-v6-c-description-list__description">${description}</dd></div>` : ''}
        ${prop.default ? `<div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Default</dt><dd class="pf-v6-c-description-list__description"><code>${this.#escapeHtml(prop.default)}</code></dd></div>` : ''}
        ${prop.privacy ? `<div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Privacy</dt><dd class="pf-v6-c-description-list__description">${this.#escapeHtml(prop.privacy)}</dd></div>` : ''}
      </dl>
    `;
  }

  async #buildMethodDetails(item, manifest) {
    const ce = this.#findCustomElement(manifest, item.modulePath, item.tagName);
    if (!ce) return '<div class="empty-state">Custom element not found</div>';

    const method = ce.members?.find(m => m.kind === 'method' && m.name === item.name);
    if (!method) return '<div class="empty-state">Method not found</div>';

    const summary = method.summary ? await this.#renderMarkdown(method.summary) : '';
    const description = method.description ? await this.#renderMarkdown(method.description) : '';

    return `
      <h3>${this.#escapeHtml(method.name)}()</h3>
      <dl class="pf-v6-c-description-list pf-m-horizontal pf-m-compact">
        <div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Element</dt><dd class="pf-v6-c-description-list__description"><code>&lt;${this.#escapeHtml(ce.tagName)}&gt;</code></dd></div>
        ${summary ? `<div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Summary</dt><dd class="pf-v6-c-description-list__description">${summary}</dd></div>` : ''}
        ${description ? `<div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Description</dt><dd class="pf-v6-c-description-list__description">${description}</dd></div>` : ''}
        ${method.return?.type?.text ? `<div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Returns</dt><dd class="pf-v6-c-description-list__description"><code>${this.#escapeHtml(method.return.type.text)}</code></dd></div>` : ''}
        ${method.privacy ? `<div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Privacy</dt><dd class="pf-v6-c-description-list__description">${this.#escapeHtml(method.privacy)}</dd></div>` : ''}
      </dl>
    `;
  }

  async #buildEventDetails(item, manifest) {
    const ce = this.#findCustomElement(manifest, item.modulePath, item.tagName);
    if (!ce) return '<div class="empty-state">Custom element not found</div>';

    const event = ce.events?.find(e => e.name === item.name);
    if (!event) return '<div class="empty-state">Event not found</div>';

    const summary = event.summary ? await this.#renderMarkdown(event.summary) : '';
    const description = event.description ? await this.#renderMarkdown(event.description) : '';

    return `
      <h3>${this.#escapeHtml(event.name)}</h3>
      <dl class="pf-v6-c-description-list pf-m-horizontal pf-m-compact">
        <div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Element</dt><dd class="pf-v6-c-description-list__description"><code>&lt;${this.#escapeHtml(ce.tagName)}&gt;</code></dd></div>
        ${event.type?.text ? `<div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Detail Type</dt><dd class="pf-v6-c-description-list__description"><code>${this.#escapeHtml(event.type.text)}</code></dd></div>` : ''}
        ${summary ? `<div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Summary</dt><dd class="pf-v6-c-description-list__description">${summary}</dd></div>` : ''}
        ${description ? `<div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Description</dt><dd class="pf-v6-c-description-list__description">${description}</dd></div>` : ''}
      </dl>
    `;
  }

  async #buildSlotDetails(item, manifest) {
    const ce = this.#findCustomElement(manifest, item.modulePath, item.tagName);
    if (!ce) return '<div class="empty-state">Custom element not found</div>';

    const slot = ce.slots?.find(s => s.name === item.name);
    if (!slot) return '<div class="empty-state">Slot not found</div>';

    const summary = slot.summary ? await this.#renderMarkdown(slot.summary) : '';
    const description = slot.description ? await this.#renderMarkdown(slot.description) : '';

    return `
      <h3>${slot.name || '(default)'}</h3>
      <dl class="pf-v6-c-description-list pf-m-horizontal pf-m-compact">
        <div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Element</dt><dd class="pf-v6-c-description-list__description"><code>&lt;${this.#escapeHtml(ce.tagName)}&gt;</code></dd></div>
        ${summary ? `<div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Summary</dt><dd class="pf-v6-c-description-list__description">${summary}</dd></div>` : ''}
        ${description ? `<div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Description</dt><dd class="pf-v6-c-description-list__description">${description}</dd></div>` : ''}
      </dl>
    `;
  }

  async #buildClassDetails(item, manifest) {
    const cls = this.#findDeclaration(manifest, item.modulePath, item.name, 'class');
    if (!cls) return '<div class="empty-state">Class not found</div>';

    const summary = cls.summary ? await this.#renderMarkdown(cls.summary) : '';
    const description = cls.description ? await this.#renderMarkdown(cls.description) : '';

    return `
      <h3>${this.#escapeHtml(cls.name)}</h3>
      <dl class="pf-v6-c-description-list pf-m-horizontal pf-m-compact">
        <div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Type</dt><dd class="pf-v6-c-description-list__description">Class</dd></div>
        ${summary ? `<div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Summary</dt><dd class="pf-v6-c-description-list__description">${summary}</dd></div>` : ''}
        ${description ? `<div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Description</dt><dd class="pf-v6-c-description-list__description">${description}</dd></div>` : ''}
        ${cls.superclass?.name ? `<div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Extends</dt><dd class="pf-v6-c-description-list__description"><code>${this.#escapeHtml(cls.superclass.name)}</code></dd></div>` : ''}
      </dl>
    `;
  }

  async #buildFunctionDetails(item, manifest) {
    const func = this.#findDeclaration(manifest, item.modulePath, item.name, 'function');
    if (!func) return '<div class="empty-state">Function not found</div>';

    const summary = func.summary ? await this.#renderMarkdown(func.summary) : '';
    const description = func.description ? await this.#renderMarkdown(func.description) : '';

    return `
      <h3>${this.#escapeHtml(func.name)}()</h3>
      <dl class="pf-v6-c-description-list pf-m-horizontal pf-m-compact">
        <div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Type</dt><dd class="pf-v6-c-description-list__description">Function</dd></div>
        ${summary ? `<div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Summary</dt><dd class="pf-v6-c-description-list__description">${summary}</dd></div>` : ''}
        ${description ? `<div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Description</dt><dd class="pf-v6-c-description-list__description">${description}</dd></div>` : ''}
        ${func.return?.type?.text ? `<div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Returns</dt><dd class="pf-v6-c-description-list__description"><code>${this.#escapeHtml(func.return.type.text)}</code></dd></div>` : ''}
      </dl>
    `;
  }

  async #buildVariableDetails(item, manifest) {
    const variable = this.#findDeclaration(manifest, item.modulePath, item.name, 'variable');
    if (!variable) return '<div class="empty-state">Variable not found</div>';

    const summary = variable.summary ? await this.#renderMarkdown(variable.summary) : '';
    const description = variable.description ? await this.#renderMarkdown(variable.description) : '';

    return `
      <h3>${this.#escapeHtml(variable.name)}</h3>
      <dl class="pf-v6-c-description-list pf-m-horizontal pf-m-compact">
        <div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Type</dt><dd class="pf-v6-c-description-list__description">Variable</dd></div>
        ${variable.type?.text ? `<div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Value Type</dt><dd class="pf-v6-c-description-list__description"><code>${this.#escapeHtml(variable.type.text)}</code></dd></div>` : ''}
        ${summary ? `<div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Summary</dt><dd class="pf-v6-c-description-list__description">${summary}</dd></div>` : ''}
        ${description ? `<div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Description</dt><dd class="pf-v6-c-description-list__description">${description}</dd></div>` : ''}
        ${variable.default ? `<div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Default</dt><dd class="pf-v6-c-description-list__description"><code>${this.#escapeHtml(variable.default)}</code></dd></div>` : ''}
      </dl>
    `;
  }

  async #buildMixinDetails(item, manifest) {
    const mixin = this.#findDeclaration(manifest, item.modulePath, item.name, 'mixin');
    if (!mixin) return '<div class="empty-state">Mixin not found</div>';

    const summary = mixin.summary ? await this.#renderMarkdown(mixin.summary) : '';
    const description = mixin.description ? await this.#renderMarkdown(mixin.description) : '';

    return `
      <h3>${this.#escapeHtml(mixin.name)}()</h3>
      <dl class="pf-v6-c-description-list pf-m-horizontal pf-m-compact">
        <div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Type</dt><dd class="pf-v6-c-description-list__description">Mixin</dd></div>
        ${summary ? `<div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Summary</dt><dd class="pf-v6-c-description-list__description">${summary}</dd></div>` : ''}
        ${description ? `<div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Description</dt><dd class="pf-v6-c-description-list__description">${description}</dd></div>` : ''}
      </dl>
    `;
  }

  /**
   * Call markdown API to render markdown to HTML
   */
  async #renderMarkdown(text) {
    if (!text) return '';

    try {
      const response = await fetch('/__cem/api/markdown', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ text })
      });

      if (!response.ok) {
        console.error('Markdown API returned error:', response.status);
        return this.#escapeHtml(text);
      }

      const data = await response.json();
      return data.html;
    } catch (error) {
      console.error('Failed to render markdown:', error);
      return this.#escapeHtml(text);
    }
  }

  /**
   * Find a module in the manifest
   */
  #findModule(manifest, modulePath) {
    return manifest.modules?.find(m => m.path === modulePath);
  }

  /**
   * Find a custom element in the manifest
   */
  #findCustomElement(manifest, modulePath, tagName) {
    const module = this.#findModule(manifest, modulePath);
    if (!module) return null;

    return module.declarations?.find(d => d.kind === 'class' && d.customElement && d.tagName === tagName);
  }

  /**
   * Find a declaration (class, function, variable, mixin) in the manifest
   */
  #findDeclaration(manifest, modulePath, name, kind) {
    const module = this.#findModule(manifest, modulePath);
    if (!module) return null;

    return module.declarations?.find(d => d.kind === kind && d.name === name);
  }

  /**
   * Escape HTML to prevent XSS
   */
  #escapeHtml(str) {
    if (!str) return '';
    return String(str)
      .replace(/&/g, '&amp;')
      .replace(/</g, '&lt;')
      .replace(/>/g, '&gt;')
      .replace(/"/g, '&quot;')
      .replace(/'/g, '&#039;');
  }

  /**
   * Set content in the details panel
   */
  #setContent(html) {
    const container = this.shadowRoot.getElementById('details-content');
    if (container) {
      container.innerHTML = html;
    }
  }

  /**
   * Clear content from the details panel
   */
  #clearContent() {
    const container = this.shadowRoot.getElementById('details-content');
    if (container) {
      container.innerHTML = '<div class="empty-state">Select an item to view details</div>';
    }
  }

  static {
    customElements.define(this.is, this);
  }
}
