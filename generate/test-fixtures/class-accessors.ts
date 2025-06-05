import { LitElement } from 'lit';
import { customElement, property } from 'lit/decorators.js';

@customElement('class-accessors')
export class ClassAccessors extends LitElement {
  get getter(): string { }

  /** getter description */
  get getterDescribed(): string { }

  /** @summary getter summary */
  get getterSummarized(): string { }

  get pair(): string {}
  set pair(v: string) {}

  /** pair with description */
  get pairDescribed(): string {}
  set pairDescribed(v: string) {}

  /** @summary pair with summary */
  get pairSummarized(): string {}
  set pairSummarized(v: string) {}

  @property()
  get propertyPair(): string {}
  set propertyPair(v: string) {}

  @property({ attribute: 'property-pair-described' })
  get propertyPairDescribed(): string {}
  /** property pair with description */
  set propertyPairDescribed(v: string) {}

  @property({ reflects: true })
  /** @summary property pair with summary */
  get propertyPairSummarized(): string {}
  set propertyPairSummarized(v: string) {}
}
