/**
 * @customElement my-element
 * @attr {string} name - The name attribute
 */
export class MyElement extends HTMLElement {
  static observedAttributes = ['name'];
}
