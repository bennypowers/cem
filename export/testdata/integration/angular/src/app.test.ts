import { describe, it, expect, afterEach, beforeAll } from 'vitest';
import { TestBed } from '@angular/core/testing';
import { BrowserDynamicTestingModule, platformBrowserDynamicTesting } from '@angular/platform-browser-dynamic/testing';
import { Component, CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';
import { DemoButtonComponent, DemoCardComponent, DemoSpinnerComponent } from '../generated';

@Component({
  standalone: true,
  imports: [DemoButtonComponent],
  schemas: [CUSTOM_ELEMENTS_SCHEMA],
  template: `<demo-button [variant]="'brand'" [disabled]="true">Click me</demo-button>`,
})
class DemoButtonHost {}

@Component({
  standalone: true,
  imports: [DemoCardComponent],
  schemas: [CUSTOM_ELEMENTS_SCHEMA],
  template: `<demo-card><span slot="header">Title</span><p>Body content</p></demo-card>`,
})
class DemoCardHost {}

@Component({
  standalone: true,
  imports: [DemoSpinnerComponent],
  schemas: [CUSTOM_ELEMENTS_SCHEMA],
  template: `<demo-spinner></demo-spinner>`,
})
class DemoSpinnerHost {}

beforeAll(() => {
  TestBed.initTestEnvironment(
    BrowserDynamicTestingModule,
    platformBrowserDynamicTesting(),
  );
});

afterEach(() => {
  TestBed.resetTestingModule();
});

describe('Angular wrappers render custom elements', () => {
  it('renders demo-button with properties', async () => {
    await TestBed.configureTestingModule({
      imports: [DemoButtonHost],
    }).compileComponents();
    const fixture = TestBed.createComponent(DemoButtonHost);
    fixture.detectChanges();
    const el = fixture.nativeElement.querySelector('demo-button');
    expect(el).toBeTruthy();
  });

  it('renders demo-card with slots', async () => {
    await TestBed.configureTestingModule({
      imports: [DemoCardHost],
    }).compileComponents();
    const fixture = TestBed.createComponent(DemoCardHost);
    fixture.detectChanges();
    const el = fixture.nativeElement.querySelector('demo-card');
    expect(el).toBeTruthy();
  });

  it('renders demo-spinner', async () => {
    await TestBed.configureTestingModule({
      imports: [DemoSpinnerHost],
    }).compileComponents();
    const fixture = TestBed.createComponent(DemoSpinnerHost);
    fixture.detectChanges();
    const el = fixture.nativeElement.querySelector('demo-spinner');
    expect(el).toBeTruthy();
  });
});
