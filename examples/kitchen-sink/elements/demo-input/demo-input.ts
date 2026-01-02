import { LitElement, html } from 'lit';
import { customElement, property, query } from 'lit/decorators.js';
import { classMap } from 'lit/directives/class-map.js';
import { ifDefined } from 'lit/directives/if-defined.js';
import styles from './demo-input.css' with { type: 'css' };

@customElement('demo-input')
export class DemoInput extends LitElement {
  static styles = [styles];

  @query('.input') input!: HTMLInputElement;

  @property() label = '';
  @property() value = '';
  @property() placeholder = '';
  @property() type: 'text' | 'email' | 'password' | 'number' | 'tel' | 'url' = 'text';
  @property({ type: Boolean, reflect: true }) disabled = false;
  @property({ type: Boolean, reflect: true }) required = false;
  @property({ type: Boolean, reflect: true }) readonly = false;
  @property() name = '';
  @property() autocomplete = '';
  @property({ attribute: 'helper-text' }) helperText = '';
  @property({ attribute: 'error-text' }) errorText = '';
  @property({ type: Boolean, reflect: true }) error = false;

  private handleInput(e: Event) {
    this.value = (e.target as HTMLInputElement).value;
    this.dispatchEvent(new CustomEvent('input', {
      detail: { value: this.value },
      bubbles: true,
      composed: true
    }));
  }

  private handleChange(e: Event) {
    this.value = (e.target as HTMLInputElement).value;
    this.dispatchEvent(new CustomEvent('change', {
      detail: { value: this.value },
      bubbles: true,
      composed: true
    }));
  }

  focus(options?: FocusOptions) {
    this.input?.focus(options);
  }

  blur() {
    this.input?.blur();
  }

  select() {
    this.input?.select();
  }

  render() {
    const displayHelperText = this.error && this.errorText ? this.errorText : this.helperText;

    return html`
      <div id="wrapper">
        ${this.label ? html`
          <label id="label" part="label">
            ${this.label}
            ${this.required ? html`<span aria-label="required">*</span>` : ''}
          </label>
        ` : ''}
        <input id="input"
               part="input"
               class=${classMap({ error: this.error })}
               type=${this.type}
               .value=${this.value}
               placeholder=${ifDefined(this.placeholder || undefined)}
               ?disabled=${this.disabled}
               ?required=${this.required}
               ?readonly=${this.readonly}
               name=${ifDefined(this.name || undefined)}
               autocomplete=${ifDefined(this.autocomplete || undefined)}
               @input=${this.handleInput}
               @change=${this.handleChange}
        />
        ${displayHelperText ? html`
          <div id="helper-text"
               part="helper-text"
               class=${classMap({ error: this.error })}>
            ${displayHelperText}
          </div>
        ` : ''}
      </div>
    `;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'demo-input': DemoInput;
  }
}
