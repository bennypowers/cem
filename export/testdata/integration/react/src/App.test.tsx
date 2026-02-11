import { describe, it, expect } from 'vitest';
import { render } from '@testing-library/react';
import { DemoButton, DemoCard, DemoSpinner } from '../generated';

describe('React wrappers render custom elements', () => {
  it('renders demo-button with properties', () => {
    const { container } = render(<DemoButton variant="brand" disabled>Click me</DemoButton>);
    const el = container.querySelector('demo-button') as any;
    expect(el).toBeTruthy();
    expect(el.variant).toBe('brand');
    expect(el.disabled).toBe(true);
  });

  it('renders demo-card with slots', () => {
    const { container } = render(
      <DemoCard>
        <span slot="header">Title</span>
        <p>Body content</p>
      </DemoCard>
    );
    const el = container.querySelector('demo-card');
    expect(el).toBeTruthy();
    expect(el?.querySelector('[slot="header"]')?.textContent).toBe('Title');
  });

  it('renders demo-spinner', () => {
    const { container } = render(<DemoSpinner />);
    expect(container.querySelector('demo-spinner')).toBeTruthy();
  });
});
