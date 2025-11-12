// Knob Change Events - Custom events for interactive knob controls

/**
 * Event fired when an attribute knob value changes
 */
export class KnobAttributeChangeEvent extends Event {
  /**
   * @param {string} name - Attribute name
   * @param {string} value - New attribute value
   */
  constructor(name, value) {
    super('knob:attribute-change', { bubbles: true, composed: true });
    this.name = name;
    this.value = value;
  }
}

/**
 * Event fired when a property knob value changes
 */
export class KnobPropertyChangeEvent extends Event {
  /**
   * @param {string} name - Property name
   * @param {*} value - New property value (typed based on property type)
   */
  constructor(name, value) {
    super('knob:property-change', { bubbles: true, composed: true });
    this.name = name;
    this.value = value;
  }
}

/**
 * Event fired when a CSS custom property knob value changes
 */
export class KnobCSSPropertyChangeEvent extends Event {
  /**
   * @param {string} name - CSS custom property name (including --)
   * @param {string} value - New CSS custom property value
   */
  constructor(name, value) {
    super('knob:css-property-change', { bubbles: true, composed: true });
    this.name = name;
    this.value = value;
  }
}
