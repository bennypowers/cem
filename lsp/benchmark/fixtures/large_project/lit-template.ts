import { LitElement, html, css } from 'lit';
import { customElement, property, state } from 'lit/decorators.js';

@customElement('app-dashboard')
export class AppDashboard extends LitElement {
  static styles = css`
    :host {
      display: block;
      padding: 1rem;
      background: #f8f9fa;
      min-height: 100vh;
    }

    .header {
      background: white;
      padding: 1.5rem;
      margin-bottom: 2rem;
      border-radius: 8px;
      box-shadow: 0 1px 3px rgba(0,0,0,0.1);
    }

    .grid {
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
      gap: 1.5rem;
      margin-bottom: 2rem;
    }

    .sidebar {
      grid-column: 1;
      background: white;
      padding: 1rem;
      border-radius: 8px;
      box-shadow: 0 1px 3px rgba(0,0,0,0.1);
    }

    .main-content {
      grid-column: 2 / -1;
    }
  `;

  @property({ type: String })
  accessor title = 'Dashboard';

  @property({ type: Array })
  accessor users = [
    { id: 1, name: 'John Doe', email: 'john@example.com', role: 'Admin', status: 'active' },
    { id: 2, name: 'Jane Smith', email: 'jane@example.com', role: 'User', status: 'pending' },
    { id: 3, name: 'Bob Johnson', email: 'bob@example.com', role: 'Editor', status: 'active' },
    { id: 4, name: 'Alice Wilson', email: 'alice@example.com', role: 'User', status: 'inactive' },
    { id: 5, name: 'Charlie Brown', email: 'charlie@example.com', role: 'Admin', status: 'active' }
  ];

  @state()
  private accessor selectedUser: any = null;

  @state()
  private accessor showModal = false;

  @state()
  private accessor formData = {
    name: '',
    email: '',
    role: 'user',
    preferences: {
      notifications: true,
      darkMode: false,
      language: 'en'
    }
  };

  render() {
    return html`
      <!-- Page Header -->
      <div class="header">
        <my-container-flex justify="space-between" align="center">
          <div>
            <h1>${this.title}</h1>
            <p>Manage your application dashboard with comprehensive controls</p>
          </div>

          <my-button-group>
            <my-button-primary @click=${this._handleCreateUser}>
              <my-image-icon slot="icon" name="plus"></my-image-icon>
              Add User
            </my-button-primary>
            <my-button-secondary @click=${this._handleExport}>
              <my-image-icon slot="icon" name="download"></my-image-icon>
              Export
            </my-button-secondary>
            <my-button-icon icon="settings" @click=${this._openSettings}></my-button-icon>
          </my-button-group>
        </my-container-flex>

        <!-- Navigation Tabs -->
        <my-nav-tabs>
          <my-nav-item href="#users" active>Users</my-nav-item>
          <my-nav-item href="#analytics">Analytics</my-nav-item>
          <my-nav-item href="#settings">Settings</my-nav-item>
          <my-nav-item href="#reports">Reports</my-nav-item>
        </my-nav-tabs>
      </div>

      <!-- Alert Messages -->
      ${this._renderAlerts()}

      <!-- Main Layout -->
      <div class="grid">
        <!-- Sidebar -->
        <aside class="sidebar">
          <my-container-stack gap="large">
            <!-- Quick Stats -->
            <my-card-stats>
              <h3 slot="header">Quick Stats</h3>
              <my-container-stack gap="medium">
                <my-container-flex justify="space-between">
                  <span>Total Users</span>
                  <my-badge-primary>${this.users.length}</my-badge-primary>
                </my-container-flex>
                <my-container-flex justify="space-between">
                  <span>Active Users</span>
                  <my-badge-success>${this.users.filter(u => u.status === 'active').length}</my-badge-success>
                </my-container-flex>
                <my-container-flex justify="space-between">
                  <span>Pending</span>
                  <my-badge-warning>${this.users.filter(u => u.status === 'pending').length}</my-badge-warning>
                </my-container-flex>
              </my-container-stack>
            </my-card-stats>

            <!-- Filters -->
            <my-card-basic>
              <h3 slot="header">Filters</h3>
              <my-container-stack gap="medium">
                <my-input-search
                  placeholder="Search users..."
                  @input=${this._handleSearch}>
                </my-input-search>

                <my-input-select label="Role Filter">
                  <option value="">All Roles</option>
                  <option value="admin">Admin</option>
                  <option value="user">User</option>
                  <option value="editor">Editor</option>
                </my-input-select>

                <my-input-select label="Status Filter">
                  <option value="">All Status</option>
                  <option value="active">Active</option>
                  <option value="pending">Pending</option>
                  <option value="inactive">Inactive</option>
                </my-input-select>

                <my-button-outline full-width @click=${this._clearFilters}>
                  Clear Filters
                </my-button-outline>
              </my-container-stack>
            </my-card-basic>

            <!-- Recent Activity -->
            <my-card-elevated>
              <h3 slot="header">Recent Activity</h3>
              <my-list-basic>
                <my-list-item>
                  <my-image-icon slot="leading" name="user-plus" color="success"></my-image-icon>
                  <div>
                    <strong>John Doe</strong> joined
                    <div>2 minutes ago</div>
                  </div>
                </my-list-item>
                <my-list-item>
                  <my-image-icon slot="leading" name="edit" color="primary"></my-image-icon>
                  <div>
                    <strong>Profile updated</strong>
                    <div>1 hour ago</div>
                  </div>
                </my-list-item>
                <my-list-item>
                  <my-image-icon slot="leading" name="download" color="info"></my-image-icon>
                  <div>
                    <strong>Report exported</strong>
                    <div>3 hours ago</div>
                  </div>
                </my-list-item>
              </my-list-basic>
            </my-card-elevated>
          </my-container-stack>
        </aside>

        <!-- Main Content -->
        <main class="main-content">
          <my-container-stack gap="large">
            <!-- Data Table -->
            <my-card-basic>
              <my-container-flex slot="header" justify="space-between" align="center">
                <h2>User Management</h2>
                <my-button-group size="small">
                  <my-button-icon icon="refresh" @click=${this._refreshData}></my-button-icon>
                  <my-button-icon icon="filter"></my-button-icon>
                  <my-button-icon icon="sort"></my-button-icon>
                </my-button-group>
              </my-container-flex>

              <my-table-sortable striped selectable>
                <thead>
                  <tr>
                    <th>
                      <my-input-checkbox @change=${this._selectAll}></my-input-checkbox>
                    </th>
                    <th>User</th>
                    <th>Email</th>
                    <th>Role</th>
                    <th>Status</th>
                    <th>Last Active</th>
                    <th>Actions</th>
                  </tr>
                </thead>
                <tbody>
                  ${this.users.map(user => html`
                    <tr>
                      <td>
                        <my-input-checkbox .checked=${user.id === this.selectedUser?.id}></my-input-checkbox>
                      </td>
                      <td>
                        <my-container-flex gap="small" align="center">
                          <my-image-avatar
                            src="/api/avatar/${user.id}"
                            alt="${user.name}"
                            size="small">
                          </my-image-avatar>
                          <div>
                            <strong>${user.name}</strong>
                            <div style="font-size: 0.875rem; color: #666;">ID: ${user.id}</div>
                          </div>
                        </my-container-flex>
                      </td>
                      <td>${user.email}</td>
                      <td>
                        <my-badge-outline variant=${this._getRoleVariant(user.role)}>
                          ${user.role}
                        </my-badge-outline>
                      </td>
                      <td>
                        <my-badge-${user.status === 'active' ? 'success' : user.status === 'pending' ? 'warning' : 'error'}>
                          ${user.status}
                        </my-badge-${user.status === 'active' ? 'success' : user.status === 'pending' ? 'warning' : 'error'}>
                      </td>
                      <td>
                        <my-tooltip content="Last seen 2 hours ago">
                          2h ago
                        </my-tooltip>
                      </td>
                      <td>
                        <my-button-group size="small">
                          <my-button-icon
                            icon="edit"
                            @click=${() => this._editUser(user)}
                            variant="ghost">
                          </my-button-icon>
                          <my-button-icon
                            icon="eye"
                            @click=${() => this._viewUser(user)}
                            variant="ghost">
                          </my-button-icon>
                          <my-dropdown-menu>
                            <my-button-icon slot="trigger" icon="more-vertical" variant="ghost"></my-button-icon>
                            <my-dropdown-item @click=${() => this._duplicateUser(user)}>
                              <my-image-icon slot="icon" name="copy"></my-image-icon>
                              Duplicate
                            </my-dropdown-item>
                            <my-dropdown-item @click=${() => this._resetPassword(user)}>
                              <my-image-icon slot="icon" name="key"></my-image-icon>
                              Reset Password
                            </my-dropdown-item>
                            <my-divider-horizontal></my-divider-horizontal>
                            <my-dropdown-item @click=${() => this._deleteUser(user)} variant="danger">
                              <my-image-icon slot="icon" name="trash"></my-image-icon>
                              Delete User
                            </my-dropdown-item>
                          </my-dropdown-menu>
                        </my-button-group>
                      </td>
                    </tr>
                  `)}
                </tbody>
              </my-table-sortable>

              <!-- Pagination -->
              <my-container-flex slot="footer" justify="space-between" align="center">
                <div>
                  Showing ${this.users.length} of ${this.users.length} users
                </div>
                <my-nav-pagination>
                  <my-button-previous disabled>Previous</my-button-previous>
                  <my-nav-item active>1</my-nav-item>
                  <my-nav-item>2</my-nav-item>
                  <my-nav-item>3</my-nav-item>
                  <my-button-next>Next</my-button-next>
                </my-nav-pagination>
              </my-container-flex>
            </my-card-basic>

            <!-- Charts Section -->
            <my-container-grid gap="medium">
              <my-card-elevated>
                <h3 slot="header">User Growth</h3>
                <my-chart-line
                  .data=${[
                    { x: 'Jan', y: 10 },
                    { x: 'Feb', y: 15 },
                    { x: 'Mar', y: 12 },
                    { x: 'Apr', y: 20 },
                    { x: 'May', y: 18 },
                    { x: 'Jun', y: 25 }
                  ]}
                  responsive
                  animated>
                </my-chart-line>
              </my-card-elevated>

              <my-card-elevated>
                <h3 slot="header">Role Distribution</h3>
                <my-chart-doughnut
                  .data=${[
                    { label: 'Admin', value: 2, color: '#ef4444' },
                    { label: 'User', value: 2, color: '#3b82f6' },
                    { label: 'Editor', value: 1, color: '#10b981' }
                  ]}
                  responsive>
                </my-chart-doughnut>
              </my-card-elevated>

              <my-card-elevated>
                <h3 slot="header">Activity Heatmap</h3>
                <my-chart-heatmap
                  .data=${this._generateHeatmapData()}
                  responsive>
                </my-chart-heatmap>
              </my-card-elevated>
            </my-container-grid>
          </my-container-stack>
        </main>
      </div>

      <!-- Modals and Overlays -->
      ${this._renderModals()}
    `;
  }

  private _renderAlerts() {
    return html`
      <my-container-stack gap="small">
        <my-alert-success dismissible icon="check">
          User data has been successfully updated.
        </my-alert-success>
        <my-alert-info icon="info">
          System maintenance scheduled for tonight at 2 AM UTC.
        </my-alert-info>
      </my-container-stack>
    `;
  }

  private _renderModals() {
    return html`
      <!-- User Edit Modal -->
      <my-modal-dialog .open=${this.showModal} size="large" @close=${this._closeModal}>
        <h2 slot="header">
          ${this.selectedUser ? 'Edit User' : 'Create User'}
        </h2>

        <my-container-stack gap="large">
          <!-- Personal Information -->
          <my-card-basic>
            <h3 slot="header">Personal Information</h3>
            <my-container-grid gap="medium">
              <my-input-text
                label="Full Name"
                .value=${this.formData.name}
                @input=${this._updateFormData('name')}
                required>
              </my-input-text>
              <my-input-email
                label="Email Address"
                .value=${this.formData.email}
                @input=${this._updateFormData('email')}
                required>
              </my-input-email>
              <my-input-select
                label="Role"
                .value=${this.formData.role}
                @change=${this._updateFormData('role')}>
                <option value="user">User</option>
                <option value="admin">Admin</option>
                <option value="editor">Editor</option>
              </my-input-select>

              <my-input-tel
                label="Phone Number"
                placeholder="+1 (555) 123-4567">
              </my-input-tel>
            </my-container-grid>
          </my-card-basic>

          <!-- Preferences -->
          <my-card-basic>
            <h3 slot="header">Preferences</h3>
            <my-container-stack gap="medium">
              <my-input-switch
                label="Email Notifications"
                .checked=${this.formData.preferences.notifications}
                @change=${this._updatePreference('notifications')}>
              </my-input-switch>

              <my-input-switch
                label="Dark Mode"
                .checked=${this.formData.preferences.darkMode}
                @change=${this._updatePreference('darkMode')}>
              </my-input-switch>

              <my-input-select
                label="Language"
                .value=${this.formData.preferences.language}
                @change=${this._updatePreference('language')}>
                <option value="en">English</option>
                <option value="es">Spanish</option>
                <option value="fr">French</option>
                <option value="de">German</option>
              </my-input-select>
            </my-container-stack>
          </my-card-basic>
        </my-container-stack>

        <my-container-flex slot="footer" gap="medium" justify="end">
          <my-button-secondary @click=${this._closeModal}>
            Cancel
          </my-button-secondary>
          <my-button-primary @click=${this._saveUser}>
            ${this.selectedUser ? 'Update User' : 'Create User'}
          </my-button-primary>
        </my-container-flex>
      </my-modal-dialog>

      <!-- Settings Drawer -->
      <my-modal-drawer placement="right" .open=${false}>
        <h2 slot="header">Dashboard Settings</h2>

        <my-container-stack gap="large">
          <my-input-switch label="Auto-refresh data"></my-input-switch>
          <my-input-slider label="Refresh interval (seconds)" min="5" max="60" value="30"></my-input-slider>
          <my-input-select label="Default view">
            <option value="table">Table View</option>
            <option value="grid">Grid View</option>
            <option value="list">List View</option>
          </my-input-select>
        </my-container-stack>
      </my-modal-drawer>

      <!-- Toast Notifications -->
      <my-toast-success .open=${false} timeout="3000">
        Operation completed successfully!
      </my-toast-success>
    `;
  }

  private _generateHeatmapData() {
    return Array.from({ length: 7 }, (_, day) =>
      Array.from({ length: 24 }, (_, hour) => ({
        day,
        hour,
        value: (day * 24 + hour) % 100  // Deterministic pattern for consistent benchmarking
      }))
    ).flat();
  }

  private _getRoleVariant(role: string) {
    const variants: Record<string, string> = {
      'admin': 'error',
      'editor': 'warning',
      'user': 'info'
    };
    return variants[role.toLowerCase()] || 'info';
  }

  private _handleCreateUser() {
    this.selectedUser = null;
    this.formData = {
      name: '',
      email: '',
      role: 'user',
      preferences: {
        notifications: true,
        darkMode: false,
        language: 'en'
      }
    };
    this.showModal = true;
  }

  private _handleExport() {
    // Simulate export functionality
    console.log('Exporting user data...');
  }

  private _openSettings() {
    // Open settings drawer
    console.log('Opening settings...');
  }

  private _handleSearch(e: CustomEvent) {
    const query = (e.target as HTMLInputElement).value;
    console.log('Search:', query);
  }

  private _clearFilters() {
    console.log('Clearing filters...');
  }

  private _refreshData() {
    console.log('Refreshing data...');
  }

  private _selectAll(e: CustomEvent) {
    const checked = (e.target as HTMLInputElement).checked;
    console.log('Select all:', checked);
  }

  private _editUser(user: any) {
    this.selectedUser = user;
    this.formData = {
      name: user.name,
      email: user.email,
      role: user.role,
      preferences: user.preferences || {
        notifications: true,
        darkMode: false,
        language: 'en'
      }
    };
    this.showModal = true;
  }

  private _viewUser(user: any) {
    console.log('Viewing user:', user);
  }

  private _duplicateUser(user: any) {
    console.log('Duplicating user:', user);
  }

  private _resetPassword(user: any) {
    console.log('Resetting password for:', user);
  }

  private _deleteUser(user: any) {
    console.log('Deleting user:', user);
  }

  private _closeModal() {
    this.showModal = false;
    this.selectedUser = null;
  }

  private _saveUser() {
    console.log('Saving user:', this.formData);
    this.showModal = false;
  }

  private _updateFormData(field: string) {
    return (e: CustomEvent) => {
      const value = (e.target as HTMLInputElement).value;
      this.formData = { ...this.formData, [field]: value };
    };
  }

  private _updatePreference(field: string) {
    return (e: CustomEvent) => {
      const value = (e.target as HTMLInputElement).checked;
      this.formData = {
        ...this.formData,
        preferences: { ...this.formData.preferences, [field]: value }
      };
    };
  }
}
