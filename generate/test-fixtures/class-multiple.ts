import { LitElement } from 'lit';
import { customElement, property } from 'lit/decorators.js';

export class ClassA extends Event {
  aField: string;
  aMethod(): string {}
}

@customElement('class-b')
export class ClassB extends LitElement {
  bField: string;
  @property() bProp: string;
  bMethod(): string {}
}
