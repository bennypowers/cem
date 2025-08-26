// TypeScript template literal examples for large project testing
// Contains multiple html template literals with custom elements

import { html, css, LitElement } from 'lit';
import { customElement, property } from 'lit/decorators.js';

@customElement('app-layout')
export class AppLayout extends LitElement {
  @property() title = 'Application';
  @property() showSidebar = true;

  static styles = css`
    :host {
      display: block;
      min-height: 100vh;
    }
    .layout {
      display: grid;
      grid-template-columns: auto 1fr;
      min-height: 100vh;
    }
  `;

  render() {
    return html`
      <div class="layout">
        ${this.showSidebar ? this.renderSidebar() : ''}
        <main>
          <my-header variant="primary" size="large">
            <h1>${this.title}</h1>
            <my-button-group slot="actions">
              <my-button variant="ghost" icon="search">Search</my-button>
              <my-button variant="primary" icon="user">Profile</my-button>
            </my-button-group>
          </my-header>
          <slot></slot>
        </main>
      </div>
    `;
  }

  private renderSidebar() {
    return html`
      <my-nav variant="sidebar" orientation="vertical">
        <my-nav-item href="/dashboard" active>
          <my-icon slot="icon" name="home"></my-icon>
          Dashboard
        </my-nav-item>
        <my-nav-item href="/components">
          <my-icon slot="icon" name="components"></my-icon>
          Components
        </my-nav-item>
        <my-nav-item href="/settings">
          <my-icon slot="icon" name="settings"></my-icon>
          Settings
        </my-nav-item>
      </my-nav>
    `;
  }
}

@customElement('dashboard-overview')
export class DashboardOverview extends LitElement {
  @property() metrics = [];

  render() {
    return html`
      <section class="dashboard">
        <my-container layout="grid" gap="large">
          ${this.renderMetrics()}
          ${this.renderCharts()}
          ${this.renderDataTable()}
        </my-container>
      </section>
    `;
  }

  private renderMetrics() {
    return html`
      <div class="metrics-grid">
        <my-card variant="filled" size="medium">
          <my-icon slot="icon" name="users" size="large"></my-icon>
          <h3 slot="header">Active Users</h3>
          <div class="metric-value">12,847</div>
          <my-progress value="85" max="100" size="small"></my-progress>
          <my-button slot="footer" variant="ghost" size="small">View Details</my-button>
        </my-card>

        <my-card variant="filled" size="medium">
          <my-icon slot="icon" name="revenue" size="large"></my-icon>
          <h3 slot="header">Total Revenue</h3>
          <div class="metric-value">$847,293</div>
          <my-progress value="72" max="100" size="small"></my-progress>
          <my-button slot="footer" variant="ghost" size="small">View Report</my-button>
        </my-card>

        <my-card variant="filled" size="medium">
          <my-icon slot="icon" name="growth" size="large"></my-icon>
          <h3 slot="header">Growth Rate</h3>
          <div class="metric-value">+23.5%</div>
          <my-progress value="90" max="100" size="small"></my-progress>
          <my-button slot="footer" variant="ghost" size="small">Analyze</my-button>
        </my-card>
      </div>
    `;
  }

  private renderCharts() {
    return html`
      <my-card variant="outlined" size="large">
        <my-header slot="header">
          <h2>Performance Analytics</h2>
          <my-button-group>
            <my-button variant="ghost" size="small">Day</my-button>
            <my-button variant="ghost" size="small">Week</my-button>
            <my-button variant="primary" size="small">Month</my-button>
          </my-button-group>
        </my-header>

        <my-chart type="line" data-source="api/analytics" responsive>
          <my-chart-legend position="bottom">
            <my-chart-legend-item color="blue" label="Users"></my-chart-legend-item>
            <my-chart-legend-item color="green" label="Revenue"></my-chart-legend-item>
            <my-chart-legend-item color="orange" label="Conversion"></my-chart-legend-item>
          </my-chart-legend>
        </my-chart>

        <my-button-group slot="footer">
          <my-button variant="primary">Export Chart</my-button>
          <my-button variant="secondary">Configure</my-button>
        </my-button-group>
      </my-card>
    `;
  }

  private renderDataTable() {
    return html`
      <my-card variant="outlined" size="large">
        <h2 slot="header">Recent Activity</h2>

        <my-table size="medium" sortable filterable>
          <my-table-header>
            <my-table-cell type="header">
              <my-button variant="ghost" size="small">
                User
                <my-icon slot="suffix" name="sort"></my-icon>
              </my-button>
            </my-table-cell>
            <my-table-cell type="header">Action</my-table-cell>
            <my-table-cell type="header">Time</my-table-cell>
            <my-table-cell type="header">Status</my-table-cell>
          </my-table-header>

          ${this.renderTableRows()}
        </my-table>

        <my-pagination slot="footer" 
          page="1" 
          total-pages="10" 
          size="medium">
        </my-pagination>
      </my-card>
    `;
  }

  private renderTableRows() {
    return html`
      <my-table-row>
        <my-table-cell>
          <my-avatar size="small" src="user1.jpg"></my-avatar>
          John Doe
        </my-table-cell>
        <my-table-cell>Updated profile</my-table-cell>
        <my-table-cell>2 minutes ago</my-table-cell>
        <my-table-cell>
          <my-badge variant="success" size="small">Complete</my-badge>
        </my-table-cell>
      </my-table-row>

      <my-table-row>
        <my-table-cell>
          <my-avatar size="small" src="user2.jpg"></my-avatar>
          Jane Smith
        </my-table-cell>
        <my-table-cell>Created report</my-table-cell>
        <my-table-cell>5 minutes ago</my-table-cell>
        <my-table-cell>
          <my-badge variant="warning" size="small">Processing</my-badge>
        </my-table-cell>
      </my-table-row>

      <my-table-row>
        <my-table-cell>
          <my-avatar size="small" src="user3.jpg"></my-avatar>
          Mike Johnson
        </my-table-cell>
        <my-table-cell>Exported data</my-table-cell>
        <my-table-cell>10 minutes ago</my-table-cell>
        <my-table-cell>
          <my-badge variant="info" size="small">Ready</my-badge>
        </my-table-cell>
      </my-table-row>
    `;
  }
}

@customElement('form-builder')
export class FormBuilder extends LitElement {
  @property() formData = {};

  render() {
    return html`
      <my-card variant="outlined" size="large">
        <h2 slot="header">Form Configuration</h2>
        
        <my-form>
          <my-form-section>
            <h3>Basic Information</h3>
            
            <my-form-group>
              <my-label for="title">Form Title</my-label>
              <my-input id="title" type="text" size="large" required>
                <my-validation-message slot="error">
                  Title is required
                </my-validation-message>
              </my-input>
            </my-form-group>

            <my-form-group>
              <my-label for="description">Description</my-label>
              <my-textarea id="description" rows="3" size="large">
                <my-character-counter slot="footer" max="500"></my-character-counter>
              </my-textarea>
            </my-form-group>
          </my-form-section>

          <my-form-section>
            <h3>Form Fields</h3>
            
            ${this.renderFieldBuilder()}
          </my-form-section>

          <my-form-actions>
            <my-button-group>
              <my-button type="submit" variant="primary" size="large">
                <my-icon slot="prefix" name="save"></my-icon>
                Save Form
              </my-button>
              <my-button type="button" variant="secondary" size="large">
                <my-icon slot="prefix" name="preview"></my-icon>
                Preview
              </my-button>
              <my-button type="reset" variant="tertiary" size="large">
                <my-icon slot="prefix" name="reset"></my-icon>
                Reset
              </my-button>
            </my-button-group>
          </my-form-actions>
        </my-form>
      </my-card>
    `;
  }

  private renderFieldBuilder() {
    return html`
      <my-container layout="grid" gap="medium">
        <my-card variant="filled" size="small">
          <h4 slot="header">Text Field</h4>
          <my-input type="text" placeholder="Sample text input"></my-input>
          <my-button slot="footer" variant="ghost" size="small">Configure</my-button>
        </my-card>

        <my-card variant="filled" size="small">
          <h4 slot="header">Select Field</h4>
          <my-select>
            <option value="">Choose option</option>
            <option value="opt1">Option 1</option>
            <option value="opt2">Option 2</option>
          </my-select>
          <my-button slot="footer" variant="ghost" size="small">Configure</my-button>
        </my-card>

        <my-card variant="filled" size="small">
          <h4 slot="header">Checkbox Group</h4>
          <my-checkbox-group>
            <my-checkbox value="cb1">Checkbox 1</my-checkbox>
            <my-checkbox value="cb2">Checkbox 2</my-checkbox>
          </my-checkbox-group>
          <my-button slot="footer" variant="ghost" size="small">Configure</my-button>
        </my-card>
      </my-container>
    `;
  }
}