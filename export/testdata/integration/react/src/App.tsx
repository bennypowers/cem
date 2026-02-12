import { DemoButton, DemoCard, DemoSpinner } from '../generated';

export function App() {
  return (
    <div style={{ padding: '2rem', fontFamily: 'system-ui' }}>
      <h1>React Export Demo (Library Author)</h1>

      <section>
        <h2>DemoButton</h2>
        <DemoButton variant="brand">Brand Button</DemoButton>
        <DemoButton variant="danger" disabled>Disabled Danger</DemoButton>
      </section>

      <section>
        <h2>DemoCard</h2>
        <DemoCard>
          <span slot="header">Card Title</span>
          <p>Card body content goes here.</p>
        </DemoCard>
      </section>

      <section>
        <h2>DemoSpinner</h2>
        <DemoSpinner />
      </section>
    </div>
  );
}
