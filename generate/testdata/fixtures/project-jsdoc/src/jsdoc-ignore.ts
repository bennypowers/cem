@customElement('jsdoc-ignore-field')
class JsdocIgnoreField extends LitElement {
  /**
   * @ignore
   */
  @property() ignoredProp: string;

  @property() visibleProp: number;
}

@customElement('jsdoc-internal-field')
class JsdocInternalField extends LitElement {
  /**
   * @internal
   */
  @property() internalProp: boolean = false;

  @property() visibleProp: number;
}

@customElement('jsdoc-ignore-method')
class JsdocIgnoreMethod extends LitElement {
  /**
   * @ignore
   */
  ignoredMethod(): void {}

  visibleMethod(): void {}
}

@customElement('jsdoc-internal-method')
class JsdocInternalMethod extends LitElement {
  /**
   * @internal
   */
  internalMethod(): void {}

  visibleMethod(): void {}
}
