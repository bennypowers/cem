import { LitElement } from 'lit';
import { customElement, property } from 'lit/decorators.js';

/**
 * Test class for properties with @property({ type: ... }) decorator options
 * that don't include attribute or reflect options.
 *
 * This is a regression test for issue #191 where properties with
 * @property({ type: Number }) were not being detected.
 */
@customElement('class-property-decorator-with-type-option')
export class ClassPropertyDecoratorWithTypeOption extends LitElement {
  /**
   * Property with type option only (no initializer)
   * This was previously not detected
   */
  @property({ type: Number }) size?: number;

  /**
   * Property with type option and initializer
   * This should also be detected
   */
  @property({ type: Boolean }) enabled = false;

  /**
   * Property with type option only (required)
   */
  @property({ type: String }) label!: string;

  /**
   * Property with bare decorator (for comparison, this works)
   */
  @property() name = 'default';
}
