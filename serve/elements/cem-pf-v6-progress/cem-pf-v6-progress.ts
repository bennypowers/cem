import type { TemplateResult, PropertyValues } from 'lit';
import { LitElement, html, nothing } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';
import { property } from 'lit/decorators/property.js';
import { classMap } from 'lit/directives/class-map.js';
import { styleMap } from 'lit/directives/style-map.js';

import styles from './cem-pf-v6-progress.css' with { type: 'css' };

export type ProgressSize = 'sm' | 'lg';
export type ProgressMeasureLocation = 'outside' | 'inside' | 'none';
export type ProgressVariant = 'success' | 'danger' | 'warning';

const checkCircleIcon = html`<svg id="status-icon" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" fill="currentColor" aria-hidden="true"><path d="M256 512A256 256 0 1 0 256 0a256 256 0 1 0 0 512zM369 209L241 337c-9.4 9.4-24.6 9.4-33.9 0l-64-64c-9.4-9.4-9.4-24.6 0-33.9s24.6-9.4 33.9 0l47 47L335 175c9.4-9.4 24.6-9.4 33.9 0s9.4 24.6 0 33.9z"/></svg>`;
const triangleExclamationIcon = html`<svg id="status-icon" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" fill="currentColor" aria-hidden="true"><path d="M256 32c14.2 0 27.3 7.5 34.5 19.8l216 368c7.3 12.4 7.3 27.7 .2 40.1S486.3 480 472 480H40c-14.3 0-27.6-7.7-34.7-20.1s-7-27.8 .2-40.1l216-368C228.7 39.5 241.8 32 256 32zm0 128c-13.3 0-24 10.7-24 24V296c0 13.3 10.7 24 24 24s24-10.7 24-24V184c0-13.3-10.7-24-24-24zm32 224a32 32 0 1 0 -64 0 32 32 0 1 0 64 0z"/></svg>`;
const circleExclamationIcon = html`<svg id="status-icon" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" fill="currentColor" aria-hidden="true"><path d="M256 512A256 256 0 1 0 256 0a256 256 0 1 0 0 512zm0-384c13.3 0 24 10.7 24 24V264c0 13.3-10.7 24-24 24s-24-10.7-24-24V152c0-13.3 10.7-24 24-24zM224 352a32 32 0 1 1 64 0 32 32 0 1 1 -64 0z"/></svg>`;

const VARIANT_ICONS = new Map<ProgressVariant, TemplateResult>([
  ['success', checkCircleIcon],
  ['warning', triangleExclamationIcon],
  ['danger', circleExclamationIcon],
]);

/**
 * A progress bar providing visual representation of completion status.
 *
 * @summary Displays completion status of an ongoing process or task.
 * @slot helper-text - Supplementary text below the progress bar.
 */
@customElement('cem-pf-v6-progress')
export class CemPfV6Progress extends LitElement {
  static styles = styles;

  #internals = this.attachInternals();

  @property({ type: Number }) accessor value = 0;
  @property() accessor description: string | undefined;
  @property({ type: Boolean }) accessor truncated = false;
  @property({ attribute: 'accessible-label' }) accessor accessibleLabel: string | undefined;
  @property({ type: Number }) accessor max = 100;
  @property({ type: Number }) accessor min = 0;
  @property() accessor size: ProgressSize | undefined;
  @property({ attribute: 'measure-location' }) accessor measureLocation: ProgressMeasureLocation | undefined;
  @property() accessor variant: ProgressVariant | undefined;
  @property({ type: Boolean, attribute: 'hide-status-icon' }) accessor hideStatusIcon = false;
  @property({ attribute: 'value-text' }) accessor valueText: string | undefined;

  #hasHelperText = false;

  get #calculatedPercentage(): number {
    const { value, min, max } = this;
    const percentage = Math.round((value - min) / (max - min) * 100);
    if (Number.isNaN(percentage) || percentage < 0) {
      return 0;
    }
    return Math.min(percentage, 100);
  }

  get #displayText(): string {
    return this.valueText ?? `${this.#calculatedPercentage}%`;
  }

  get #icon(): TemplateResult | typeof nothing {
    if (this.hideStatusIcon) {
      return nothing;
    }
    return VARIANT_ICONS.get(this.variant!) ?? nothing;
  }

  override updated(changed: PropertyValues): void {
    if (changed.has('value') || changed.has('min') || changed.has('max')) {
      this.#internals.ariaValueNow = this.#calculatedPercentage.toString();
      this.#internals.ariaValueMin = '0';
      this.#internals.ariaValueMax = '100';
    }
    if (changed.has('valueText')) {
      this.#internals.ariaValueText = this.valueText ?? null;
    }
    if (changed.has('accessibleLabel') || changed.has('description')) {
      this.#internals.ariaLabel = this.accessibleLabel ?? this.description ?? 'Progress status';
    }
  }

  override connectedCallback(): void {
    super.connectedCallback();
    this.#internals.role = 'progressbar';
    this.#internals.ariaValueNow = this.#calculatedPercentage.toString();
    this.#internals.ariaValueMin = '0';
    this.#internals.ariaValueMax = '100';
    this.#internals.ariaLabel = this.accessibleLabel ?? this.description ?? 'Progress status';
  }

  #onHelperTextSlotchange(event: Event) {
    const slot = event.target as HTMLSlotElement;
    const elements = slot?.assignedElements() ?? [];
    this.#hasHelperText = elements.length > 0;
    this.#internals.ariaDescribedByElements = elements.length ? elements : null;
    this.requestUpdate();
  }

  override render(): TemplateResult<1> {
    const pct = this.#calculatedPercentage;
    const displayText = this.#displayText;
    const icon = this.#icon;
    const noMeasure = this.measureLocation === 'none';
    const inside = this.measureLocation === 'inside';
    const hasDescription = this.description != null;
    const hasIcon = this.variant != null && !this.hideStatusIcon;
    const singleline = !hasDescription;

    const classes = {
      [this.size ?? '']: !!this.size,
      [this.measureLocation ?? '']: !!this.measureLocation && this.measureLocation !== 'none',
      [this.variant ?? '']: !!this.variant,
      singleline,
      truncated: this.truncated,
    };

    return html`
      <div id="container" class="${classMap(classes)}">
        <div id="description"
             ?hidden="${!hasDescription}"
>${this.description ?? ''}</div>

        <div id="status"
             aria-hidden="true"
             ?hidden="${noMeasure && !hasIcon}">
          ${!inside && !noMeasure ? html`<span id="measure">${displayText}</span>` : nothing}
          ${icon}
        </div>

        <div id="bar">
          <div id="indicator"
               style="${styleMap({ width: `${pct}%` })}">
            ${inside && !noMeasure ? html`<span id="measure">${displayText}</span>` : nothing}
          </div>
        </div>

        <div id="helper-text" ?hidden="${!this.#hasHelperText}">
          <slot id="helper-text-slot"
                name="helper-text"
                @slotchange="${this.#onHelperTextSlotchange}"></slot>
        </div>
      </div>
    `;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'cem-pf-v6-progress': CemPfV6Progress;
  }
}
