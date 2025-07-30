export class ClassConstructorProperties {
  declare field: boolean;
  constructor(
    public pub: boolean,
    private priv: string,
    protected prot: number,
    /** JSDoc description */
    public documented: string,
    /** @deprecated Use something else */
    public deprecated: boolean,
  ) {}
}