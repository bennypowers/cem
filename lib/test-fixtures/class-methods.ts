import { LitElement } from 'lit';
import { customElement, property } from 'lit/decorators.js';

@customElement('class-methods')
export class ClassMethods extends LitElement {
  method() {}

  static $static() {}

  public $public() {}

  protected $protected() {}

  private $private() {}

  unary(unary: string): string {}

  binary(binary0: string, binary1: number): string {}

  optional(optional0: string, optional1?: number): string {}

  spread(...spreads: number[]): string {}

  /**
   * description
   * @param name name param
   * @param [optional] optional param
   * @param [default='default'] default param
   * @param {Type} typed typed param
   * @param {Type} [typedOptional] typed optional param
   * @param {Type} [typedOptionalDefault='default'] typed optional param with default
   * @param obj object param
   * @param obj.prop
   * @param [obj.optional]
   * @param [obj.default='default']
   * @param {Type} obj.typed
   * @param {Type} [obj.typedOptional]
   * @param {Type} [obj.typedOptionalDefault='default']
   * @return foo
   */
  jsdoc(
    name: never,
    optional: never,
    default: never,
    typed: never,
    typedOptional: never,
    typedOptionalDefault: never,
    obj: never,
  ) {}
}
