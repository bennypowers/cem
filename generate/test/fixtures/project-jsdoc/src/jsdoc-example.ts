/**
 * Class with inline caption example
 * @example Basic usage
 *          ```html
 *          <example-inline></example-inline>
 *          ```
 */
@customElement('example-inline')
class ExampleInline extends LitElement { }

/**
 * Class with explicit caption tag
 * @example
 * <caption>Advanced example with caption tag</caption>
 * ```html
 * <example-explicit></example-explicit>
 * ```
 */
@customElement('example-explicit')
class ExampleExplicit extends LitElement { }

/**
 * Class with no caption (code-only)
 * @example
 * ```html
 * <example-no-caption></example-no-caption>
 * ```
 */
@customElement('example-no-caption')
class ExampleNoCaption extends LitElement { }

/**
 * Class with multiple examples
 * @example Simple case
 *          ```html
 *          <example-multiple></example-multiple>
 *          ```
 * @example Complex case with attributes
 *          ```html
 *          <example-multiple foo="bar" baz="qux"></example-multiple>
 *          ```
 */
@customElement('example-multiple')
class ExampleMultiple extends LitElement { }

/**
 * Class with multi-line caption
 * @example
 * This example demonstrates
 * advanced usage patterns
 * with multiple lines
 * ```html
 * <example-multiline></example-multiline>
 * ```
 */
@customElement('example-multiline')
class ExampleMultiline extends LitElement { }

/**
 * Class with description and example
 */
@customElement('example-with-methods')
class ExampleWithMethods extends LitElement {
  /**
   * Method with example
   * @example Basic method call
   *          ```typescript
   *          element.doSomething();
   *          ```
   */
  doSomething() { }

  /**
   * Property with example
   * @example Setting the value
   *          ```typescript
   *          element.value = 'hello';
   *          ```
   */
  @property() value: string;
}

/**
 * Class with example without language specifier
 * @example
 * ```
 * <example-no-lang></example-no-lang>
 * ```
 */
@customElement('example-no-lang')
class ExampleNoLang extends LitElement { }
