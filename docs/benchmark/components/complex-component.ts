import { LitElement, html, css, property, customElement } from 'lit-element';

/**
 * A complex component with many fields and methods.
 * Used to test validation behavior with larger APIs.
 */
@customElement('complex-component')
export class ComplexComponent extends LitElement {
  // Field 1
  @property({ type: String }) 
  title = '';

  // Field 2
  @property({ type: String }) 
  description = '';

  // Field 3
  @property({ type: Number }) 
  count = 0;

  // Field 4
  @property({ type: Boolean }) 
  enabled = true;

  // Field 5
  @property({ type: Array }) 
  items: string[] = [];

  // Field 6 (bonus)
  @property({ type: String }) 
  theme = 'light';

  // Field 7 (bonus)
  @property({ type: Object }) 
  config = {};

  static styles = css`
    :host {
      display: block;
      border: 1px solid #ccc;
      padding: 1rem;
    }
    .header { font-weight: bold; }
    .content { margin: 0.5rem 0; }
    .actions { margin-top: 1rem; }
    button { margin-right: 0.5rem; }
  `;

  render() {
    return html`
      <div class="header">${this.title}</div>
      <div class="content">${this.description}</div>
      <div class="content">Count: ${this.count}</div>
      <div class="content">Items: ${this.items.length}</div>
      <div class="actions">
        <button @click=${this.increment} ?disabled=${!this.enabled}>
          Increment
        </button>
        <button @click=${this.decrement} ?disabled=${!this.enabled}>
          Decrement
        </button>
        <button @click=${this.reset}>Reset</button>
        <button @click=${this.addItem}>Add Item</button>
        <button @click=${this.clearItems}>Clear Items</button>
      </div>
    `;
  }

  // Method 1
  /**
   * Increments the count by 1.
   */
  increment() {
    this.count++;
    this.notifyChange();
  }

  // Method 2
  /**
   * Decrements the count by 1.
   */
  decrement() {
    if (this.count > 0) {
      this.count--;
      this.notifyChange();
    }
  }

  // Method 3
  /**
   * Resets the count to 0.
   */
  reset() {
    this.count = 0;
    this.items = [];
    this.notifyChange();
  }

  // Method 4
  /**
   * Adds a new item to the items array.
   */
  addItem() {
    const newItem = `Item ${this.items.length + 1}`;
    this.items = [...this.items, newItem];
    this.notifyChange();
  }

  // Method 5
  /**
   * Clears all items from the array.
   */
  clearItems() {
    this.items = [];
    this.notifyChange();
  }

  // Method 6 (bonus)
  /**
   * Toggles the enabled state.
   */
  toggle() {
    this.enabled = !this.enabled;
    this.notifyChange();
  }

  // Method 7 (bonus)
  /**
   * Sets the theme for the component.
   * @param theme - The theme name ('light' or 'dark')
   */
  setTheme(theme: string) {
    this.theme = theme;
    this.notifyChange();
  }

  // Method 8 (bonus)
  /**
   * Updates the configuration object.
   * @param newConfig - The new configuration to merge
   */
  updateConfig(newConfig: any) {
    this.config = { ...this.config, ...newConfig };
    this.notifyChange();
  }

  // Private helper method
  /**
   * Notifies listeners of state changes.
   * @private
   */
  private notifyChange() {
    this.dispatchEvent(new CustomEvent('complex-change', {
      detail: {
        title: this.title,
        description: this.description,
        count: this.count,
        enabled: this.enabled,
        items: this.items,
        theme: this.theme,
        config: this.config
      }
    }));
  }
}