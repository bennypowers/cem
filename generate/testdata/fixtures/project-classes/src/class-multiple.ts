export class ClassA extends Event {
  aField: string;
  aMethod(): string {}
}

@customElement('class-b')
class ClassB extends LitElement {
  bField: string;
  @property() bProp: string;
  bMethod(): string {}
}
