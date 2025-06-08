export class DupeEvent extends Event {
  public dupe: boolean
}

@customElement('class-fields-duplicate')
@themable
export class ClassFieldsMultiple extends LitElement {
  static dupe = true;
  #dupe = true;
  /** Even though `DupeEvent` has an `dupe` field, this is still tracked */
  @property({ reflect: true }) dupe = false;
}

