import React from 'react';
import './components/my-card';
import './components/ui-button';

interface AppProps {
  title: string;
}

export default function App({ title }: AppProps) {
  const handleCardClick = (event: CustomEvent) => {
    console.log('Card clicked:', event.detail);
  };

  return (
    <div className="app">
      <h1>{title}</h1>

      {/* Custom element: my-card with various attributes */}
      <my-card
        title="Welcome Card"
        variant="primary"
        onClick={handleCardClick}
      >
        <div slot="header">
          <h2>Card Header</h2>
        </div>

        <p>This is the main content of the card.</p>

        <div slot="footer">
          <ui-button size="small" loading={false}>
            Click me
          </ui-button>
        </div>
      </my-card>

      {/* Another custom element with different props */}
      <my-card variant="secondary" disabled>
        <p>A disabled secondary card</p>
      </my-card>

      {/* Standalone button component */}
      <ui-button size="large">
        Large Button
      </ui-button>

      {/* Button with event handler */}
      <ui-button
        size="medium"
        loading={true}
        onClick={() => console.log('Button clicked')}
      >
        Loading Button
      </ui-button>
    </div>
  );
}