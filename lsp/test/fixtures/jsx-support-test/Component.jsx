import React, { useState } from 'react';

// JSX file with custom elements for testing
export function Component() {
  const [cardTitle, setCardTitle] = useState('Dynamic Title');
  const [buttonSize, setButtonSize] = useState('medium');

  return (
    <div>
      {/* Custom elements in JSX */}
      <my-card title={cardTitle} variant="danger">
        <h3 slot="header">Dynamic Header</h3>
        <p>Content with dynamic title: {cardTitle}</p>
        <ui-button
          slot="footer"
          size={buttonSize}
          onClick={() => setCardTitle('Updated!')}
        >
          Update Title
        </ui-button>
      </my-card>

      {/* Self-closing custom element */}
      <ui-button size="small" loading />

      {/* Nested custom elements */}
      <my-card variant="primary">
        <my-card variant="secondary" title="Nested Card">
          <ui-button>Nested Button</ui-button>
        </my-card>
      </my-card>

      {/* Conditional rendering */}
      {cardTitle === 'Updated!' && (
        <my-card title="Conditional Card">
          <p>This card appears conditionally</p>
        </my-card>
      )}

      {/* Map over elements */}
      {['small', 'medium', 'large'].map(size => (
        <ui-button key={size} size={size}>
          {size.charAt(0).toUpperCase() + size.slice(1)} Button
        </ui-button>
      ))}
    </div>
  );
}