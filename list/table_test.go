package list

import (
	"strings"
	"testing"
)

func TestBuildTableData_AllColumns(t *testing.T) {
	headers := []string{"Tag", "Class", "Module", "Summary"}
	rows := [][]string{
		{"foo", "Foo", "src/foo.js", "Foo summary"},
		{"bar", "Bar", "src/bar.js", "Bar summary"},
	}
	finalHeaders, finalRows, err := BuildTableData(headers, rows, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(finalHeaders) != 4 || len(finalRows) != 2 {
		t.Errorf("unexpected output: headers=%v rows=%v", finalHeaders, finalRows)
	}
}

func TestBuildTableData_SelectedColumns(t *testing.T) {
	headers := []string{"Tag", "Class", "Module"}
	rows := [][]string{
		{"foo", "Foo", "src/foo.js"},
	}
	finalHeaders, finalRows, err := BuildTableData(headers, rows, []string{"Module"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if want := []string{"Tag", "Module"}; !equalStrings(finalHeaders, want) {
		t.Errorf("headers = %v, want %v", finalHeaders, want)
	}
	if want := []string{"foo", "src/foo.js"}; !equalStrings(finalRows[0], want) {
		t.Errorf("row = %v, want %v", finalRows[0], want)
	}
}

func TestBuildTableData_UnknownColumn(t *testing.T) {
	headers := []string{"Tag", "Class", "Module"}
	rows := [][]string{{"foo", "Foo", "src/foo.js"}}
	_, _, err := BuildTableData(headers, rows, []string{"NotAColumn"})
	if err == nil || !strings.Contains(err.Error(), "unknown column") {
		t.Errorf("expected error, got %v", err)
	}
}

func TestBuildTableData_EmptyRows(t *testing.T) {
	headers := []string{"Tag", "Class", "Module"}
	rows := [][]string{}
	_, finalRows, err := BuildTableData(headers, rows, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(finalRows) != 0 {
		t.Errorf("expected 0 rows, got %d", len(finalRows))
	}
}

// Utility
func equalStrings(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
