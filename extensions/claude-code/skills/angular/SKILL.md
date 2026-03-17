---
name: angular
description: Show how to use custom elements in Angular, including required CUSTOM_ELEMENTS_SCHEMA configuration. Use when the user asks to "use in Angular", "Angular custom elements", "custom elements in Angular", "CUSTOM_ELEMENTS_SCHEMA", or mentions "Angular module", "standalone component", or "angular integration".
tools: Read, Glob, Grep
---

# Custom Elements in Angular

Generate framework-specific guidance for using custom elements in Angular, including required schema configuration.

## Workflow

### Phase 1: Detect Angular Version and Style

Search the project for Angular version and patterns:

- Check `package.json` for `@angular/core` version
- Determine if using standalone components (Angular 14+) or NgModules
- Check for existing `CUSTOM_ELEMENTS_SCHEMA` usage
- Check for TypeScript strict mode settings

### Phase 2: Gather Element Data

Read the target element(s):

```
cem://element/{tagName}
cem://element/{tagName}/attributes
cem://element/{tagName}/events
cem://element/{tagName}/slots
```

Understand which features need framework-specific handling:
- **Properties**: Angular's property binding works well with custom elements
- **Events**: Angular's event binding works with custom events
- **Two-way binding**: Needs manual setup
- **Forms**: Custom elements need `ControlValueAccessor` for Angular forms

### Phase 3: Required Configuration

#### Standalone Components (Angular 14+, recommended)

```ts
// my.component.ts
import { Component, CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';

@Component({
  selector: 'app-my-page',
  standalone: true,
  schemas: [CUSTOM_ELEMENTS_SCHEMA],
  template: `
    <my-element variant="primary" (my-event)="handleEvent($event)">
      Content
    </my-element>
  `,
})
export class MyPageComponent {
  handleEvent(e: Event) {
    console.log((e as CustomEvent).detail);
  }
}
```

Each standalone component that uses custom elements needs `CUSTOM_ELEMENTS_SCHEMA` in its `schemas` array.

#### NgModule-based (Angular < 14 or legacy projects)

```ts
// app.module.ts
import { NgModule, CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';

@NgModule({
  declarations: [AppComponent],
  schemas: [CUSTOM_ELEMENTS_SCHEMA],
  // ...
})
export class AppModule {}
```

The schema applies to all components declared in that module. For larger apps, add it only to modules that use custom elements.

### Phase 4: Generate Integration Code

#### Basic Usage

```ts
@Component({
  standalone: true,
  schemas: [CUSTOM_ELEMENTS_SCHEMA],
  template: `
    <my-element
      variant="primary"
      [disabled]="isDisabled"
      (my-event)="onMyEvent($event)"
    >
      Default slot content
      <span slot="header">Header content</span>
    </my-element>
  `,
})
export class MyComponent {
  isDisabled = false;

  onMyEvent(event: Event) {
    const detail = (event as CustomEvent).detail;
  }
}
```

#### Property Binding

Angular's `[property]` syntax sets properties directly on the element, which works correctly with custom elements:

```html
<!-- Attribute (string) -->
<my-element variant="primary"></my-element>

<!-- Property binding (passes JS values) -->
<my-element [complexData]="myObject"></my-element>

<!-- Boolean -->
<my-element [disabled]="isDisabled"></my-element>
```

Angular maps `[camelCase]` property bindings to the element's camelCase properties. For kebab-case attributes, use `[attr.kebab-case]`:

```html
<!-- Property: sets el.complexData -->
<my-element [complexData]="value"></my-element>

<!-- Attribute: sets el.setAttribute('aria-label', value) -->
<my-element [attr.aria-label]="label"></my-element>
```

#### Event Handling

Angular's `(event)` syntax works with custom element events:

```html
<!-- Listen to custom events -->
<my-element (my-event)="handleEvent($event)"></my-element>

<!-- Access detail inline -->
<my-element (change)="value = $any($event).detail.value"></my-element>
```

For TypeScript, cast the event to the correct type of the event based on the manifest.

```ts
handleCustomEvent(event: CustomEvent<{ value: string }>) {
  this.value = event.detail.value;
}

// manifest has "type": { "text": "TypedInManifestEvent", "source": ... }
handleTypedInManifestEvent(event: TypedInManifestEvent) {
  this.value = event.detail.value;
}
```

#### Slots

Use the native HTML `slot` attribute:

```html
<my-element>
  <!-- Default slot -->
  <span>Default content</span>

  <!-- Named slots -->
  <div slot="header">Header content</div>
  <div slot="footer">Footer content</div>
</my-element>
```

Angular content projection (`<ng-content>`) and web component slots are different mechanisms. Use `slot="name"` for web component slots.

#### Angular Forms Integration

Custom elements don't work with Angular's `ngModel` or reactive forms out of the box. Create a `ControlValueAccessor`:

```ts
import { Directive, ElementRef, forwardRef } from '@angular/core';
import { ControlValueAccessor, NG_VALUE_ACCESSOR } from '@angular/forms';

@Directive({
  selector: 'my-input[formControlName], my-input[ngModel], my-input[formControl]',
  standalone: true,
  host: {
    '(input)': 'onInput($event)',
    '(blur)': 'onTouched()',
  },
  providers: [
    {
      provide: NG_VALUE_ACCESSOR,
      useExisting: forwardRef(() => MyInputValueAccessor),
      multi: true,
    },
  ],
})
export class MyInputValueAccessor implements ControlValueAccessor {
  private onChange: (value: any) => void = () => {};
  onTouched: () => void = () => {};

  constructor(private el: ElementRef<HTMLElement>) {}

  writeValue(value: any): void {
    (this.el.nativeElement as any).value = value;
  }

  registerOnChange(fn: (value: any) => void): void {
    this.onChange = fn;
  }

  registerOnTouched(fn: () => void): void {
    this.onTouched = fn;
  }

  setDisabledState(isDisabled: boolean): void {
    (this.el.nativeElement as any).disabled = isDisabled;
  }

  onInput(event: Event): void {
    const value = (event as CustomEvent).detail?.value
      ?? (event.target as any).value;
    this.onChange(value);
  }
}
```

Then import the directive alongside the component:

```ts
@Component({
  standalone: true,
  schemas: [CUSTOM_ELEMENTS_SCHEMA],
  imports: [ReactiveFormsModule, MyInputValueAccessor],
  template: `<my-input [formControl]="nameControl"></my-input>`,
})
export class MyFormComponent {
  nameControl = new FormControl('');
}
```

Be sure to check the element descriptions for indication if it is a Form-Associated Custom Element (FACE).

#### TypeScript Declarations

Extend the `HTMLElementTagNameMap` so Angular's template type checker understands the elements:

```ts
// custom-elements.d.ts
declare global {
  interface HTMLElementTagNameMap {
    'my-element': import('my-element-library').MyElement;
  }
}
```

For stricter typing in templates, combine with `CUSTOM_ELEMENTS_SCHEMA` — Angular won't type-check custom element bindings by default when the schema is present.

### Phase 5: Output

Present the integration with:

1. **Schema configuration** appropriate to standalone vs NgModule style
2. **Usage examples** for each element
3. **Property binding** showing `[property]` vs `[attr.attribute]`
4. **Event handling** with proper `CustomEvent` typing
5. **Slot usage** with native `slot` attribute
6. **Forms integration** if the element is form-associated (has value/input patterns)

## Guidelines

- **Always include `CUSTOM_ELEMENTS_SCHEMA`**: Without it, Angular throws template errors for unknown elements
- **Use `[property]` binding, not `attr.`**: Angular's property binding sets properties directly, which custom elements handle correctly
- **Cast events to `CustomEvent`**: Angular's template type system doesn't know about `CustomEvent.detail`
- **Don't confuse `ng-content` with `slot`**: Angular content projection and web component slots are separate systems
- **ControlValueAccessor is required for forms**: Without it, `ngModel` and reactive forms won't work with custom elements
- **Check for `CUSTOM_ELEMENTS_SCHEMA` in tests**: Test modules/components need the schema too
