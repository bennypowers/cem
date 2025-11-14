import { CemServeKnobBase } from '/__cem/cem-serve-knob-base.js';
import { KnobPropertyChangeEvent } from '/__cem/knob-events.js';

export class CemServeKnobProperty extends CemServeKnobBase {
  static is = 'cem-serve-knob-property';
  static { customElements.define(this.is, this); }

  createChangeEvent(name, value) {
    return new KnobPropertyChangeEvent(name, value);
  }
}
