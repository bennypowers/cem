import { describe, it, expect } from 'vitest';
import { mount } from '@vue/test-utils';
import { DemoButton, DemoCard, DemoSpinner } from '../generated';

describe('Vue wrappers render custom elements', () => {
  it('renders demo-button with properties', () => {
    const wrapper = mount(DemoButton, {
      props: { variant: 'brand', disabled: true },
      slots: { default: 'Click me' },
    });
    const el = wrapper.find('demo-button');
    expect(el.exists()).toBe(true);
  });

  it('renders demo-card with slots', () => {
    const wrapper = mount(DemoCard, {
      slots: {
        header: '<span>Title</span>',
        default: '<p>Body content</p>',
      },
    });
    const el = wrapper.find('demo-card');
    expect(el.exists()).toBe(true);
  });

  it('renders demo-spinner', () => {
    const wrapper = mount(DemoSpinner);
    const el = wrapper.find('demo-spinner');
    expect(el.exists()).toBe(true);
  });
});
