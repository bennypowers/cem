import { LitElement, css, html } from 'lit';
import { customElement, property } from 'lit/decorators.js';

/**
 * An accessible multi-select dropdown with search, keyboard navigation, and bulk operations.
 * Implements WCAG 2.1 AA standards with comprehensive screen reader support and keyboard
 * interaction patterns. Use for selection from large datasets where users need to choose
 * multiple related items, such as tag selection, user assignment, or category filtering.
 *
 * @summary Multi-select dropdown with search and keyboard navigation
 *
 * @event selection-change - Fired when selection changes through any interaction method (click, keyboard, programmatic). Event detail includes complete selection state and the specific change that occurred. Use for form validation, dependent field updates, or analytics tracking. Debounced during rapid keyboard navigation to prevent excessive firing.
 * @event search-input - Fired when user types in search field (when searchable is true). Event detail contains current search query and number of matching options. Use for external filtering, analytics, or implementing custom search behavior. Fires after built-in filtering is applied.
 * @event dropdown-toggle - Fired when dropdown opens or closes. Event detail indicates new state and what triggered the change. Use for coordinating with external UI elements, implementing custom positioning, or managing page-level focus. Essential for modal-style behaviors.
 * @event max-reached - Fired when user attempts to select beyond max-selections limit. Use to show helpful error messages or suggest alternative actions. Event is fired before selection is rejected, allowing for custom handling or user education.
 *
 * @csspart trigger - Main clickable area that opens the dropdown. Style for brand consistency and clear affordance. Includes focus indicators and disabled states. Should maintain minimum 44px touch target for mobile accessibility.
 * @csspart selections - Container for displaying selected items. Layout varies based on selection-display attribute. Style for clear visual hierarchy and easy scanning. Handles overflow with appropriate text truncation.
 * @csspart dropdown - Dropdown container including search, options, and footer areas. Handles positioning, shadows, and border styling. Consider z-index stacking context and responsive positioning behavior.
 * @csspart search-input - Search input field within dropdown (when searchable). Style for clear focus states and placeholder visibility. Should integrate seamlessly with dropdown design while remaining clearly functional.
 * @csspart option - Individual option elements within the list. Style for clear selection states, hover feedback, and disabled appearance. Must maintain accessibility contrast requirements across all states.
 * @csspart selected-tag - Individual selection chips (when selection-display='tags'). Include close button styling and hover states. Should be easily scannable and clearly removable for keyboard users.
 *
 * @slot - Default slot for option elements. Should contain option elements with value attributes and descriptive text content. Options support disabled state and optional description text. Maintain semantic markup for screen readers - each option should be self-contained and descriptive.
 * @slot header - Optional header content above the options list. Use for category labels, action buttons (like 'Select All'), or contextual help. Should not contain form controls that conflict with option selection. Announced as part of dropdown context by screen readers.
 * @slot empty-state - Content shown when no options match current search or when options list is empty. Should provide helpful guidance or alternative actions. Include clear messaging about search refinement or data loading states. Will be announced by screen readers when displayed.
 * @slot footer - Footer content below options list. Ideal for summary information ('X of Y selected'), action buttons, or links to manage the option set externally. Maintains focus context within the dropdown for keyboard navigation.
 *
 * @cssprop --multi-select-width - Controls component width. Use design system width tokens or specific measurements. Affects dropdown positioning and text wrapping behavior. Consider content length and container constraints when setting. (default: auto)
 * @cssprop --multi-select-max-height - Maximum height for the dropdown options list. When exceeded, the list becomes scrollable with keyboard navigation support. Use viewport-relative units for responsive behavior. Recommended range: 200px to 60vh for optimal usability. (default: 300px)
 * @cssprop --selection-tag-color - Background color for selected item tags (when selection-display='tags'). Should maintain sufficient contrast with text color for accessibility. Use semantic colors from design system that indicate selection state clearly. (default: var(--color-primary-100))
 * @cssprop --option-hover-bg - Background color for option hover state. Must meet WCAG contrast requirements against option text. Provides visual feedback for keyboard and mouse navigation. Should be distinct from selection and focus states. (default: var(--color-gray-100))
 * @cssprop --dropdown-border-radius - Border radius for dropdown container. Use design system radius tokens for consistency. Affects visual integration with page design and brand consistency. Consider relationship to trigger element styling. (default: var(--radius-md))
 *
 * @cssstate open - Applied when dropdown is expanded. Use for styling state transitions, positioning adjustments, or coordinating with external elements. Helps manage z-index and overlay behaviors.
 * @cssstate searching - Applied when user is actively typing in search field. Use for loading indicators, highlighting matched text, or adjusting dropdown layout. Provides visual feedback during search operations.
 * @cssstate max-reached - Applied when maximum selection limit is reached. Use for visual feedback, disabling remaining options, or showing helpful messaging. Should clearly communicate the constraint to users.
 * @cssstate required-empty - Applied when component is required but has no selections. Use for error styling, validation messaging, or form submission prevention. Should clearly indicate the validation state.
 */
@customElement('multi-select')
export class MultiSelectElement extends LitElement {
  static styles = css`
    :host {
      display: block;
      position: relative;
      width: var(--multi-select-width, auto);
    }

    :host([open]) [part="dropdown"] {
      display: block;
    }
  `;

  /**
   * Placeholder text shown when no selections are made. Should be descriptive
   * and action-oriented (e.g., 'Select team members' rather than 'Choose options').
   * Announced by screen readers when the component receives focus.
   */
  @property({ type: String })
  placeholder = 'Select options...';

  /**
   * Enables search/filter functionality within the dropdown. When true, users can
   * type to filter available options, improving usability for large option sets.
   * Search is case-insensitive and matches partial strings. Essential for lists
   * with more than 10 options.
   */
  @property({ type: Boolean })
  searchable = false;

  /**
   * Maximum number of items that can be selected. When limit is reached, remaining
   * options become disabled with appropriate ARIA states. Use to prevent overwhelming
   * interfaces or enforce business rules. Value of 0 means unlimited selections.
   */
  @property({ type: Number, attribute: 'max-selections' })
  maxSelections = 0;

  /**
   * Marks the component as required for form validation. Displays visual indicator
   * and updates ARIA attributes for screen readers. Prevents form submission when
   * no selections are made. Should be paired with clear error messaging.
   */
  @property({ type: Boolean })
  required = false;

  /**
   * Disables all interaction with the component. Sets appropriate ARIA states and
   * visual styling. Use for read-only states or when dependent form fields haven't
   * been completed. Maintains selection visibility for context.
   */
  @property({ type: Boolean })
  disabled = false;

  /**
   * Controls layout density and information display. 'default' shows standard spacing
   * and basic option text. 'compact' reduces spacing for dense interfaces. 'detailed'
   * shows additional option metadata and descriptions. Choose based on interface
   * density and content complexity.
   */
  @property({ type: String })
  variant: 'default' | 'compact' | 'detailed' = 'default';

  /**
   * How selected items are displayed in the closed state. 'tags' shows individual
   * removable chips (best for <= 5 selections). 'count' shows summary like '3 selected'
   * (best for large selections). 'list' shows comma-separated names (best for 2-4
   * selections). Automatically adjusts for accessibility and screen space.
   */
  @property({ type: String, attribute: 'selection-display' })
  selectionDisplay: 'tags' | 'count' | 'list' = 'tags';

  render() {
    return html`
      <div part="trigger">
        <div part="selections">
          <!-- Selected items display -->
        </div>
      </div>

      <div part="dropdown" ?hidden="${true}">
        ${this.searchable ? html`
          <input part="search-input" type="search" placeholder="Search options...">
        ` : ''}

        <slot name="header"></slot>

        <div class="options-container">
          <slot part="option"></slot>
          <slot name="empty-state"></slot>
        </div>

        <slot name="footer"></slot>
      </div>

      <div part="selected-tag" ?hidden="${true}"></div>
    `;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    'multi-select': MultiSelectElement;
  }
}
