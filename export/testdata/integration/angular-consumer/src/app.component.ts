import { Component, CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';
import { DemoButtonComponent, DemoCardComponent, DemoSpinnerComponent } from '../generated';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [DemoButtonComponent, DemoCardComponent, DemoSpinnerComponent],
  schemas: [CUSTOM_ELEMENTS_SCHEMA],
  template: `
    <div style="padding: 2rem; font-family: system-ui">
      <h1>Angular Export Demo (Consumer)</h1>

      <section>
        <h2>DemoButton</h2>
        <demo-button [variant]="'brand'">Brand Button</demo-button>
        <demo-button [variant]="'danger'" [disabled]="true">Disabled Danger</demo-button>
      </section>

      <section>
        <h2>DemoCard</h2>
        <demo-card>
          <span slot="header">Card Title</span>
          <p>Card body content goes here.</p>
        </demo-card>
      </section>

      <section>
        <h2>DemoSpinner</h2>
        <demo-spinner></demo-spinner>
      </section>
    </div>
  `,
})
export class AppComponent {}
