import React from 'react';

export function TestComponent() {
  return (
    <my-card title="Test" variant="primary" disabled>
      <ui-button size="large" loading={true}>
        Button Text
      </ui-button>
    </my-card>
  );
}