package list

import (
	"testing"
)

// Mock implementation for testing
type mockMember struct {
	name       string
	memberType string
	deprecated bool
}

func (m mockMember) ToTableRow() []string {
	return []string{m.name, "mockType", "mockDetails"}
}

func (m mockMember) IsDeprecated() bool {
	return m.deprecated
}

func (m mockMember) GetMemberType() string {
	return m.memberType
}

func TestFilterByDeprecated(t *testing.T) {
	members := []mockMember{
		{name: "active", memberType: "method", deprecated: false},
		{name: "deprecated", memberType: "method", deprecated: true},
		{name: "another-active", memberType: "field", deprecated: false},
	}

	// Test filtering for deprecated only
	deprecated := FilterByDeprecated(members, true)
	if len(deprecated) != 1 || deprecated[0].name != "deprecated" {
		t.Errorf("Expected 1 deprecated member, got %d", len(deprecated))
	}

	// Test filtering for non-deprecated only
	active := FilterByDeprecated(members, false)
	if len(active) != 2 {
		t.Errorf("Expected 2 active members, got %d", len(active))
	}
}

func TestRenderOutput(t *testing.T) {
	members := []mockMember{
		{name: "testMethod", memberType: "method", deprecated: false},
		{name: "testField", memberType: "field", deprecated: true},
	}

	headers := []string{"Name", "Type", "Details"}
	title := "Test Members"

	// Test table format
	err := RenderOutput(title, headers, members, []string{}, "table", nil)
	if err != nil {
		t.Errorf("Table format should not error: %v", err)
	}

	// Test tree format
	err = RenderOutput(title, headers, members, []string{}, "tree", nil)
	if err != nil {
		t.Errorf("Tree format should not error: %v", err)
	}

	// Test invalid format
	err = RenderOutput(title, headers, members, []string{}, "invalid", nil)
	if err == nil {
		t.Error("Invalid format should return error")
	}
}

func TestRenderOutputWithDeprecatedFilter(t *testing.T) {
	members := []mockMember{
		{name: "testMethod", memberType: "method", deprecated: false},
		{name: "testField", memberType: "field", deprecated: true},
	}

	headers := []string{"Name", "Type", "Details"}
	title := "Test Members"

	// Test with deprecated only filter
	showDeprecated := true
	err := RenderOutput(title, headers, members, []string{}, "tree", &showDeprecated)
	if err != nil {
		t.Errorf("Tree format with deprecated filter should not error: %v", err)
	}

	// Test with non-deprecated only filter
	showDeprecated = false
	err = RenderOutput(title, headers, members, []string{}, "tree", &showDeprecated)
	if err != nil {
		t.Errorf("Tree format with non-deprecated filter should not error: %v", err)
	}
}