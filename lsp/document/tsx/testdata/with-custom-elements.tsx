import React from 'react';

export function App() {
  return (
    <div>
      <my-element disabled variant="primary">
        <span slot="label">Hello</span>
      </my-element>
      <other-element name="test"></other-element>
    </div>
  );
}
