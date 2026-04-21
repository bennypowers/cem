@customElement('jsdoc-default-bare')
class JsdocDefaultBare extends LitElement {
  /**
   * @default hello
   */
  @property() value: string;
}

@customElement('jsdoc-default-number')
class JsdocDefaultNumber extends LitElement {
  /**
   * @default 42
   */
  @property() count: number;
}

@customElement('jsdoc-default-quoted')
class JsdocDefaultQuoted extends LitElement {
  /**
   * @default 'red'
   */
  @property() color: string;
}

@customElement('jsdoc-default-override')
class JsdocDefaultOverride extends LitElement {
  /**
   * @default 'blue'
   */
  @property() color: string = 'red';
}

