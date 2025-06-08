@customElement('class-accessors')
class ClassAccessors extends LitElement {
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

  /** property pair with description */
  @property({ attribute: 'property-pair-described' })
  get propertyPairDescribed(): string {}
  set propertyPairDescribed(v: string) {}

  /** @summary property pair with summary */
  @property({ reflect: true })
  get propertyPairSummarized(): string {}
  set propertyPairSummarized(v: string) {}
}
