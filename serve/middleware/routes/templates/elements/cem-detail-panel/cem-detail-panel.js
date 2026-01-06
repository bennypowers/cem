import { CemElement } from '/__cem/cem-element.js';

/**
 * Detail panel that renders manifest item details on-demand
 * @customElement cem-detail-panel
 */
export class CemDetailPanel extends CemElement {
  static is = 'cem-detail-panel';

  #cache = new Map();

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
      case 'package':
        return this.#buildPackageDetails(item, manifest);
      case 'module':
        return this.#buildModuleDetails(item, manifest);
      case 'custom-element':
        return this.#buildCustomElementDetails(item, manifest);
      case 'category':
      case 'group':
        // Categories/groups are organizational only - show a simple message
        return `<div class="empty-state">Select an item to view details</div>`;
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
      case 'css-property':
        return this.#buildCSSPropertyDetails(item, manifest);
      case 'css-part':
        return this.#buildCSSPartDetails(item, manifest);
      case 'css-state':
        return this.#buildCSSStateDetails(item, manifest);
      case 'demo':
        return this.#buildDemoDetails(item, manifest);
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

  async #buildPackageDetails(item, manifest) {
    const pkg = manifest.packages?.find(p => p.name === item.packageName);
    if (!pkg) return '<div class="empty-state">Package not found</div>';

    return `
      <h3>${this.#escapeHtml(pkg.name)}</h3>
      <dl class="pf-v6-c-description-list pf-m-horizontal pf-m-compact">
        <div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Type</dt><dd class="pf-v6-c-description-list__description">Package</dd></div>
        <div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Modules</dt><dd class="pf-v6-c-description-list__description">${pkg.modules?.length || 0}</dd></div>
        ${pkg.schemaVersion ? `<div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Schema Version</dt><dd class="pf-v6-c-description-list__description">${this.#escapeHtml(pkg.schemaVersion)}</dd></div>` : ''}
      </dl>
    `;
  }

  async #buildModuleDetails(item, manifest) {
    const module = this.#findModule(manifest, item.modulePath);
    if (!module) return '<div class="empty-state">Module not found</div>';

    const escapedPath = item.modulePath.replace(/\\/g, '\\\\').replace(/"/g, '\\"');
    const summaryPath = module.summary ? `modules.#(path=="${escapedPath}").summary` : '';
    const descriptionPath = module.description ? `modules.#(path=="${escapedPath}").description` : '';

    const summary = await this.#renderMarkdown(summaryPath);
    const description = await this.#renderMarkdown(descriptionPath);

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
    // Defensive check - if no tagName, this isn't a valid custom element
    if (!item.tagName) {
      return '<div class="empty-state">Invalid custom element (missing tagName)</div>';
    }

    const ce = this.#findCustomElement(manifest, item.modulePath, item.tagName);
    if (!ce) return '<div class="empty-state">Custom element not found</div>';

    const escapedPath = item.modulePath.replace(/\\/g, '\\\\').replace(/"/g, '\\"');
    const escapedTagName = item.tagName.replace(/\\/g, '\\\\').replace(/"/g, '\\"');
    const summaryPath = ce.summary ? `modules.#(path=="${escapedPath}").declarations.#(tagName=="${escapedTagName}").summary` : '';
    const descriptionPath = ce.description ? `modules.#(path=="${escapedPath}").declarations.#(tagName=="${escapedTagName}").description` : '';

    const summary = await this.#renderMarkdown(summaryPath);
    const description = await this.#renderMarkdown(descriptionPath);

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

    const escapedPath = item.modulePath.replace(/\\/g, '\\\\').replace(/"/g, '\\"');
    const escapedTagName = item.tagName.replace(/\\/g, '\\\\').replace(/"/g, '\\"');
    const escapedName = item.name.replace(/\\/g, '\\\\').replace(/"/g, '\\"');
    const summaryPath = attr.summary ? `modules.#(path=="${escapedPath}").declarations.#(tagName=="${escapedTagName}").attributes.#(name=="${escapedName}").summary` : '';
    const descriptionPath = attr.description ? `modules.#(path=="${escapedPath}").declarations.#(tagName=="${escapedTagName}").attributes.#(name=="${escapedName}").description` : '';

    const summary = await this.#renderMarkdown(summaryPath);
    const description = await this.#renderMarkdown(descriptionPath);

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

    const escapedPath = item.modulePath.replace(/\\/g, '\\\\').replace(/"/g, '\\"');
    const escapedTagName = item.tagName.replace(/\\/g, '\\\\').replace(/"/g, '\\"');
    const escapedName = item.name.replace(/\\/g, '\\\\').replace(/"/g, '\\"');
    const summaryPath = prop.summary ? `modules.#(path=="${escapedPath}").declarations.#(tagName=="${escapedTagName}").members.#(name=="${escapedName}").summary` : '';
    const descriptionPath = prop.description ? `modules.#(path=="${escapedPath}").declarations.#(tagName=="${escapedTagName}").members.#(name=="${escapedName}").description` : '';

    const summary = await this.#renderMarkdown(summaryPath);
    const description = await this.#renderMarkdown(descriptionPath);

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

    const escapedPath = item.modulePath.replace(/\\/g, '\\\\').replace(/"/g, '\\"');
    const escapedTagName = item.tagName.replace(/\\/g, '\\\\').replace(/"/g, '\\"');
    const escapedName = item.name.replace(/\\/g, '\\\\').replace(/"/g, '\\"');
    const summaryPath = method.summary ? `modules.#(path=="${escapedPath}").declarations.#(tagName=="${escapedTagName}").members.#(name=="${escapedName}").summary` : '';
    const descriptionPath = method.description ? `modules.#(path=="${escapedPath}").declarations.#(tagName=="${escapedTagName}").members.#(name=="${escapedName}").description` : '';

    const summary = await this.#renderMarkdown(summaryPath);
    const description = await this.#renderMarkdown(descriptionPath);

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

    const escapedPath = item.modulePath.replace(/\\/g, '\\\\').replace(/"/g, '\\"');
    const escapedTagName = item.tagName.replace(/\\/g, '\\\\').replace(/"/g, '\\"');
    const escapedName = item.name.replace(/\\/g, '\\\\').replace(/"/g, '\\"');
    const summaryPath = event.summary ? `modules.#(path=="${escapedPath}").declarations.#(tagName=="${escapedTagName}").events.#(name=="${escapedName}").summary` : '';
    const descriptionPath = event.description ? `modules.#(path=="${escapedPath}").declarations.#(tagName=="${escapedTagName}").events.#(name=="${escapedName}").description` : '';

    const summary = await this.#renderMarkdown(summaryPath);
    const description = await this.#renderMarkdown(descriptionPath);

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

    const escapedPath = item.modulePath.replace(/\\/g, '\\\\').replace(/"/g, '\\"');
    const escapedTagName = item.tagName.replace(/\\/g, '\\\\').replace(/"/g, '\\"');
    const escapedName = item.name.replace(/\\/g, '\\\\').replace(/"/g, '\\"');
    const summaryPath = slot.summary ? `modules.#(path=="${escapedPath}").declarations.#(tagName=="${escapedTagName}").slots.#(name=="${escapedName}").summary` : '';
    const descriptionPath = slot.description ? `modules.#(path=="${escapedPath}").declarations.#(tagName=="${escapedTagName}").slots.#(name=="${escapedName}").description` : '';

    const summary = await this.#renderMarkdown(summaryPath);
    const description = await this.#renderMarkdown(descriptionPath);

    return `
      <h3>${this.#escapeHtml(slot.name) || '(default)'}</h3>
      <dl class="pf-v6-c-description-list pf-m-horizontal pf-m-compact">
        <div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Element</dt><dd class="pf-v6-c-description-list__description"><code>&lt;${this.#escapeHtml(ce.tagName)}&gt;</code></dd></div>
        ${summary ? `<div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Summary</dt><dd class="pf-v6-c-description-list__description">${summary}</dd></div>` : ''}
        ${description ? `<div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Description</dt><dd class="pf-v6-c-description-list__description">${description}</dd></div>` : ''}
      </dl>
    `;
  }

  async #buildCSSPropertyDetails(item, manifest) {
    const ce = this.#findCustomElement(manifest, item.modulePath, item.tagName);
    if (!ce) return '<div class="empty-state">Custom element not found</div>';

    const cssProp = ce.cssProperties?.find(p => p.name === item.name);
    if (!cssProp) return '<div class="empty-state">CSS property not found</div>';

    const escapedPath = item.modulePath.replace(/\\/g, '\\\\').replace(/"/g, '\\"');
    const escapedTagName = item.tagName.replace(/\\/g, '\\\\').replace(/"/g, '\\"');
    const escapedName = item.name.replace(/\\/g, '\\\\').replace(/"/g, '\\"');
    const summaryPath = cssProp.summary ? `modules.#(path=="${escapedPath}").declarations.#(tagName=="${escapedTagName}").cssProperties.#(name=="${escapedName}").summary` : '';
    const descriptionPath = cssProp.description ? `modules.#(path=="${escapedPath}").declarations.#(tagName=="${escapedTagName}").cssProperties.#(name=="${escapedName}").description` : '';

    const summary = await this.#renderMarkdown(summaryPath);
    const description = await this.#renderMarkdown(descriptionPath);

    return `
      <h3>${this.#escapeHtml(cssProp.name)}</h3>
      <dl class="pf-v6-c-description-list pf-m-horizontal pf-m-compact">
        <div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Element</dt><dd class="pf-v6-c-description-list__description"><code>&lt;${this.#escapeHtml(ce.tagName)}&gt;</code></dd></div>
        ${cssProp.syntax ? `<div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Syntax</dt><dd class="pf-v6-c-description-list__description"><code>${this.#escapeHtml(cssProp.syntax)}</code></dd></div>` : ''}
        ${summary ? `<div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Summary</dt><dd class="pf-v6-c-description-list__description">${summary}</dd></div>` : ''}
        ${description ? `<div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Description</dt><dd class="pf-v6-c-description-list__description">${description}</dd></div>` : ''}
        ${cssProp.default ? `<div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Default</dt><dd class="pf-v6-c-description-list__description"><code>${this.#escapeHtml(cssProp.default)}</code></dd></div>` : ''}
      </dl>
    `;
  }

  async #buildCSSPartDetails(item, manifest) {
    const ce = this.#findCustomElement(manifest, item.modulePath, item.tagName);
    if (!ce) return '<div class="empty-state">Custom element not found</div>';

    const cssPart = ce.cssParts?.find(p => p.name === item.name);
    if (!cssPart) return '<div class="empty-state">CSS part not found</div>';

    const escapedPath = item.modulePath.replace(/\\/g, '\\\\').replace(/"/g, '\\"');
    const escapedTagName = item.tagName.replace(/\\/g, '\\\\').replace(/"/g, '\\"');
    const escapedName = item.name.replace(/\\/g, '\\\\').replace(/"/g, '\\"');
    const summaryPath = cssPart.summary ? `modules.#(path=="${escapedPath}").declarations.#(tagName=="${escapedTagName}").cssParts.#(name=="${escapedName}").summary` : '';
    const descriptionPath = cssPart.description ? `modules.#(path=="${escapedPath}").declarations.#(tagName=="${escapedTagName}").cssParts.#(name=="${escapedName}").description` : '';

    const summary = await this.#renderMarkdown(summaryPath);
    const description = await this.#renderMarkdown(descriptionPath);

    return `
      <h3>${this.#escapeHtml(cssPart.name)}</h3>
      <dl class="pf-v6-c-description-list pf-m-horizontal pf-m-compact">
        <div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Element</dt><dd class="pf-v6-c-description-list__description"><code>&lt;${this.#escapeHtml(ce.tagName)}&gt;</code></dd></div>
        ${summary ? `<div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Summary</dt><dd class="pf-v6-c-description-list__description">${summary}</dd></div>` : ''}
        ${description ? `<div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Description</dt><dd class="pf-v6-c-description-list__description">${description}</dd></div>` : ''}
      </dl>
    `;
  }

  async #buildCSSStateDetails(item, manifest) {
    const ce = this.#findCustomElement(manifest, item.modulePath, item.tagName);
    if (!ce) return '<div class="empty-state">Custom element not found</div>';

    const cssState = ce.cssStates?.find(s => s.name === item.name);
    if (!cssState) return '<div class="empty-state">CSS state not found</div>';

    const escapedPath = item.modulePath.replace(/\\/g, '\\\\').replace(/"/g, '\\"');
    const escapedTagName = item.tagName.replace(/\\/g, '\\\\').replace(/"/g, '\\"');
    const escapedName = item.name.replace(/\\/g, '\\\\').replace(/"/g, '\\"');
    const summaryPath = cssState.summary ? `modules.#(path=="${escapedPath}").declarations.#(tagName=="${escapedTagName}").cssStates.#(name=="${escapedName}").summary` : '';
    const descriptionPath = cssState.description ? `modules.#(path=="${escapedPath}").declarations.#(tagName=="${escapedTagName}").cssStates.#(name=="${escapedName}").description` : '';

    const summary = await this.#renderMarkdown(summaryPath);
    const description = await this.#renderMarkdown(descriptionPath);

    return `
      <h3>:--${this.#escapeHtml(cssState.name)}</h3>
      <dl class="pf-v6-c-description-list pf-m-horizontal pf-m-compact">
        <div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Element</dt><dd class="pf-v6-c-description-list__description"><code>&lt;${this.#escapeHtml(ce.tagName)}&gt;</code></dd></div>
        ${summary ? `<div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Summary</dt><dd class="pf-v6-c-description-list__description">${summary}</dd></div>` : ''}
        ${description ? `<div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Description</dt><dd class="pf-v6-c-description-list__description">${description}</dd></div>` : ''}
      </dl>
    `;
  }

  async #buildDemoDetails(item, manifest) {
    const ce = this.#findCustomElement(manifest, item.modulePath, item.tagName);
    if (!ce) return '<div class="empty-state">Custom element not found</div>';

    const demo = ce.demos?.find(d => d.url === item.url);
    if (!demo) return '<div class="empty-state">Demo not found</div>';

    const escapedPath = item.modulePath.replace(/\\/g, '\\\\').replace(/"/g, '\\"');
    const escapedTagName = item.tagName.replace(/\\/g, '\\\\').replace(/"/g, '\\"');
    const escapedUrl = demo.url.replace(/\\/g, '\\\\').replace(/"/g, '\\"');
    const descriptionPath = demo.description ? `modules.#(path=="${escapedPath}").declarations.#(tagName=="${escapedTagName}").demos.#(url=="${escapedUrl}").description` : '';

    const description = await this.#renderMarkdown(descriptionPath);

    return `
      <h3>Demo</h3>
      <dl class="pf-v6-c-description-list pf-m-horizontal pf-m-compact">
        <div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Element</dt><dd class="pf-v6-c-description-list__description"><code>&lt;${this.#escapeHtml(ce.tagName)}&gt;</code></dd></div>
        ${demo.url ? `<div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">URL</dt><dd class="pf-v6-c-description-list__description"><a href="${this.#escapeHtml(demo.url)}" target="_blank" rel="noopener noreferrer">${this.#escapeHtml(demo.url)}</a></dd></div>` : ''}
        ${description ? `<div class="pf-v6-c-description-list__group"><dt class="pf-v6-c-description-list__term">Description</dt><dd class="pf-v6-c-description-list__description">${description}</dd></div>` : ''}
      </dl>
    `;
  }

  async #buildClassDetails(item, manifest) {
    const cls = this.#findDeclaration(manifest, item.modulePath, item.name, 'class');
    if (!cls) return '<div class="empty-state">Class not found</div>';

    const escapedPath = item.modulePath.replace(/\\/g, '\\\\').replace(/"/g, '\\"');
    const escapedName = item.name.replace(/\\/g, '\\\\').replace(/"/g, '\\"');
    const summaryPath = cls.summary ? `modules.#(path=="${escapedPath}").declarations.#(name=="${escapedName}").summary` : '';
    const descriptionPath = cls.description ? `modules.#(path=="${escapedPath}").declarations.#(name=="${escapedName}").description` : '';

    const summary = await this.#renderMarkdown(summaryPath);
    const description = await this.#renderMarkdown(descriptionPath);

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

    const escapedPath = item.modulePath.replace(/\\/g, '\\\\').replace(/"/g, '\\"');
    const escapedName = item.name.replace(/\\/g, '\\\\').replace(/"/g, '\\"');
    const summaryPath = func.summary ? `modules.#(path=="${escapedPath}").declarations.#(name=="${escapedName}").summary` : '';
    const descriptionPath = func.description ? `modules.#(path=="${escapedPath}").declarations.#(name=="${escapedName}").description` : '';

    const summary = await this.#renderMarkdown(summaryPath);
    const description = await this.#renderMarkdown(descriptionPath);

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

    const escapedPath = item.modulePath.replace(/\\/g, '\\\\').replace(/"/g, '\\"');
    const escapedName = item.name.replace(/\\/g, '\\\\').replace(/"/g, '\\"');
    const summaryPath = variable.summary ? `modules.#(path=="${escapedPath}").declarations.#(name=="${escapedName}").summary` : '';
    const descriptionPath = variable.description ? `modules.#(path=="${escapedPath}").declarations.#(name=="${escapedName}").description` : '';

    const summary = await this.#renderMarkdown(summaryPath);
    const description = await this.#renderMarkdown(descriptionPath);

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

    const escapedPath = item.modulePath.replace(/\\/g, '\\\\').replace(/"/g, '\\"');
    const escapedName = item.name.replace(/\\/g, '\\\\').replace(/"/g, '\\"');
    const summaryPath = mixin.summary ? `modules.#(path=="${escapedPath}").declarations.#(name=="${escapedName}").summary` : '';
    const descriptionPath = mixin.description ? `modules.#(path=="${escapedPath}").declarations.#(name=="${escapedName}").description` : '';

    const summary = await this.#renderMarkdown(summaryPath);
    const description = await this.#renderMarkdown(descriptionPath);

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
   * Call markdown API to render markdown from a manifest path
   */
  async #renderMarkdown(path) {
    if (!path) return '';

    try {
      const response = await fetch('/__cem/api/markdown', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ path })
      });

      if (!response.ok) {
        console.error('Markdown API returned error:', response.status);
        return '';
      }

      const data = await response.json();
      return data.html;
    } catch (error) {
      console.error('Failed to render markdown:', error);
      return '';
    }
  }

  /**
   * Find a module in the manifest
   * Handles both single-package format (manifest.modules) and
   * workspace format (manifest.packages[*].modules)
   */
  #findModule(manifest, modulePath) {
    // Try single-package format first
    if (manifest.modules) {
      const module = manifest.modules.find(m => m.path === modulePath);
      if (module) return module;
    }

    // Try workspace format with packages array
    if (manifest.packages) {
      for (const pkg of manifest.packages) {
        const module = pkg.modules?.find(m => m.path === modulePath);
        if (module) return module;
      }
    }

    return null;
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
