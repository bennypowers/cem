export class Class extends Event {
  declare field: boolean;
  constructor(
    public thing: boolean,
    private priv: boolean,
    /** description */
    public pub?: boolean,
  ) {
    super('*', { bubbles: true, cancelable: true });
  }
}

