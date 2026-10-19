@customElement('method-overloads-exclusions')
class MethodOverloadsExclusions extends LitElement {
  // constructor overload sets are not emitted as members; only parameter
  // properties become fields, and these signatures carry none
  constructor(x: string);
  constructor(x: number);
  constructor(x: string | number) {
    super();
  }

  // body-less accessor signatures are skipped; accessor pairs with bodies
  // merge into a single field
  get label(): string;
  set label(v: string);

  get realLabel(): string {
    return 'label';
  }
  set realLabel(v: string) {
    void v;
  }

  // invalid TypeScript, parsed through error recovery: the grammar does
  // not support `declare` methods, and the misparse must not fabricate a
  // phantom member named "declare"
  declare foo(): void;

  // legitimate method names that could be mistaken for keywords keep
  // working
  type(): void {}
  global(): void {}
}
