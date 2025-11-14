import { CemServeKnobBase } from '/__cem/cem-serve-knob-base.js';
import { KnobCSSPropertyChangeEvent } from '/__cem/knob-events.js';

export class CemServeKnobCSSProperty extends CemServeKnobBase {
  static is = 'cem-serve-knob-css-property';
  static { customElements.define(this.is, this); }

  createChangeEvent(name, value) {
    return new KnobCSSPropertyChangeEvent(name, value);
  }
}
