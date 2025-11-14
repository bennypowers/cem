import { CemServeKnobBase } from '/__cem/cem-serve-knob-base.js';
import { KnobAttributeChangeEvent } from '/__cem/knob-events.js';

export class CemServeKnobAttribute extends CemServeKnobBase {
  static is = 'cem-serve-knob-attribute';
  static { customElements.define(this.is, this); }

  createChangeEvent(name, value) {
    return new KnobAttributeChangeEvent(name, value);
  }
}
