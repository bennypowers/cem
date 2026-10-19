@customElement('method-overloads-annotations')
class MethodOverloadsAnnotations extends LitElement {
  // method signatures inside inline object types belong to the object type,
  // not to the class: no phantom onClick member may be emitted
  handlers: { onClick(e: Event): void };

  // object-literal methods inside method bodies belong to the object
  // literal, not to the class: no phantom dispose member may be emitted
  register(cb: { onReady(): void }): { dispose(): void } {
    void cb;
    return { dispose() {} };
  }

  // an annotation-nested signature sharing a real method's name must not
  // displace the body-carrying method when overload sets are resolved
  options: { format(value: string): string };

  format(value: string | number): string {
    return String(value);
  }

  plain(): void {}
}
