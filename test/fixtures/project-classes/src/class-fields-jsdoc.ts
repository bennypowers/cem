export class ClassFieldsJsdoc {
  /** string description */
  description: string;
  /** @summary summary */
  summary: string;
  /** @type {Foo} */
  type: string;
  /**
   * initialized description
   * @summary initialized summary
   * @type {Foo}
   */
  initialized = 'string';
}
