export class Class extends Event {
  constructor(
    private priv: boolean,
    /** description */
    public pub?: boolean,
  ) {
    super('*', { bubbles: true, cancelable: true });
  }
}

