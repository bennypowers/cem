type SubmitType = 'submit';
type ResetType = 'reset';
type ButtonType = SubmitType | ResetType | 'button';
type NullableButtonType = ButtonType | null;
type OptionalButtonType = ButtonType | undefined;
type ComplexUnion = SubmitType | 'custom' | null;

export class MyButton extends HTMLElement {
  // Simple union with type aliases
  type: ButtonType = 'submit';
  
  // Union with null
  nullableType: NullableButtonType = null;
  
  // Union with undefined
  optionalType: OptionalButtonType = undefined;
  
  // Mixed union (aliases, literals, primitives)
  complexType: ComplexUnion = 'submit';
}

customElements.define('my-button', MyButton);
