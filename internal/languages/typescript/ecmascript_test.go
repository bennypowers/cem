package typescript

import (
	"testing"

	_ "bennypowers.dev/cem/internal/languages/html"
	Q "bennypowers.dev/cem/internal/treesitter"
)

func TestCamelToKebab(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"myAttr", "my-attr"},
		{"fooBarBaz", "foo-bar-baz"},
		{"already", "already"},
		{"URL", "u-r-l"},
		{"", ""},
		{"A", "a"},
		{"onClick", "on-click"},
		{"borderTopColor", "border-top-color"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := camelToKebab(tt.input)
			if got != tt.want {
				t.Errorf("camelToKebab(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestMatchesAttribute(t *testing.T) {
	tests := []struct {
		name          string
		memberName    string
		decoratorAttr string
		targetAttr    string
		want          bool
	}{
		{"decorator name match", "foo", "my-attr", "my-attr", true},
		{"decorator name mismatch", "foo", "my-attr", "other", false},
		{"member exact match", "disabled", "", "disabled", true},
		{"camelCase to kebab", "myAttr", "", "my-attr", true},
		{"camelCase mismatch", "myAttr", "", "other", false},
		{"decorator takes priority over member", "myAttr", "custom-name", "custom-name", true},
		{"decorator takes priority, member ignored", "myAttr", "custom-name", "my-attr", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := matchesAttribute(tt.memberName, tt.decoratorAttr, tt.targetAttr)
			if got != tt.want {
				t.Errorf("matchesAttribute(%q, %q, %q) = %v, want %v",
					tt.memberName, tt.decoratorAttr, tt.targetAttr, got, tt.want)
			}
		})
	}
}

func newTestQueryManager(t *testing.T) *Q.QueryManager {
	t.Helper()
	qm, err := Q.NewQueryManager(Q.LSPQueries())
	if err != nil {
		t.Fatalf("NewQueryManager: %v", err)
	}
	t.Cleanup(qm.Close)
	return qm
}

func newDefinitionQueryManager(t *testing.T) *Q.QueryManager {
	t.Helper()
	qm, err := Q.NewQueryManager(Q.QuerySelector{
		"html":       {"slotsAndParts"},
		"typescript": {"classes", "classMemberDeclaration"},
	})
	if err != nil {
		t.Fatalf("NewQueryManager: %v", err)
	}
	t.Cleanup(qm.Close)
	return qm
}

func TestFindAttributeDeclarationInSource(t *testing.T) {
	qm := newDefinitionQueryManager(t)

	src := []byte(`import { LitElement, html } from 'lit';
import { customElement, property } from 'lit/decorators.js';

@customElement('test-element')
export class TestElement extends LitElement {
  @property({ type: String })
  myGreeting = 'hello';

  @property({ type: Number, attribute: 'item-count' })
  itemCount = 0;
}
`)

	t.Run("camelCase property via kebab attribute", func(t *testing.T) {
		r, err := FindAttributeDeclarationInSource(src, "my-greeting", qm)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if r == nil {
			t.Fatal("expected range, got nil")
		}
		if r.Start.Line != 6 {
			t.Errorf("expected line 6, got %d", r.Start.Line)
		}
	})

	t.Run("explicit attribute name in decorator", func(t *testing.T) {
		r, err := FindAttributeDeclarationInSource(src, "item-count", qm)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if r == nil {
			t.Fatal("expected range, got nil")
		}
		if r.Start.Line != 9 {
			t.Errorf("expected line 9, got %d", r.Start.Line)
		}
	})

	t.Run("no match", func(t *testing.T) {
		r, err := FindAttributeDeclarationInSource(src, "nonexistent", qm)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if r != nil {
			t.Errorf("expected nil, got range at line %d", r.Start.Line)
		}
	})
}

func TestFindSlotDefinitionInSource(t *testing.T) {
	qm := newDefinitionQueryManager(t)

	src := []byte(`import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators.js';

@customElement('test-element')
export class TestElement extends LitElement {
  render() {
    return html` + "`" + `
      <div>
        <slot name="header"></slot>
        <slot></slot>
      </div>
    ` + "`" + `;
  }
}
`)

	t.Run("named slot", func(t *testing.T) {
		r, err := FindSlotDefinitionInSource(src, "header", qm)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if r == nil {
			t.Fatal("expected range for slot 'header', got nil")
		}
	})

	t.Run("no match", func(t *testing.T) {
		r, err := FindSlotDefinitionInSource(src, "nonexistent", qm)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if r != nil {
			t.Errorf("expected nil, got range at line %d", r.Start.Line)
		}
	})
}

func TestFindDefinedElementTags(t *testing.T) {
	qm := newTestQueryManager(t)

	t.Run("customElement decorator", func(t *testing.T) {
		src := []byte(`import { LitElement } from 'lit';
import { customElement } from 'lit/decorators.js';

@customElement('my-element')
export class MyElement extends LitElement {}
`)
		tags := FindDefinedElementTags(src, qm)
		if len(tags) != 1 || tags[0] != "my-element" {
			t.Errorf("expected [my-element], got %v", tags)
		}
	})

	t.Run("customElements.define", func(t *testing.T) {
		src := []byte(`class MyElement extends HTMLElement {}
customElements.define('my-element', MyElement);
`)
		tags := FindDefinedElementTags(src, qm)
		if len(tags) != 1 || tags[0] != "my-element" {
			t.Errorf("expected [my-element], got %v", tags)
		}
	})

	t.Run("multiple elements", func(t *testing.T) {
		src := []byte(`import { LitElement } from 'lit';
import { customElement } from 'lit/decorators.js';

@customElement('element-one')
export class ElementOne extends LitElement {}

@customElement('element-two')
export class ElementTwo extends LitElement {}
`)
		tags := FindDefinedElementTags(src, qm)
		if len(tags) != 2 {
			t.Fatalf("expected 2 tags, got %d: %v", len(tags), tags)
		}
		found := map[string]bool{}
		for _, tag := range tags {
			found[tag] = true
		}
		if !found["element-one"] || !found["element-two"] {
			t.Errorf("expected element-one and element-two, got %v", tags)
		}
	})

	t.Run("no definitions", func(t *testing.T) {
		src := []byte(`export class PlainClass {}`)
		tags := FindDefinedElementTags(src, qm)
		if len(tags) != 0 {
			t.Errorf("expected no tags, got %v", tags)
		}
	})

	t.Run("nil input", func(t *testing.T) {
		tags := FindDefinedElementTags(nil, qm)
		if tags != nil {
			t.Errorf("expected nil, got %v", tags)
		}
	})
}
