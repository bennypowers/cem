/*
Copyright © 2025 Benny Powers <web@bennypowers.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package set

import (
	"testing"
)

func TestNewSet(t *testing.T) {
	t.Run("empty set", func(t *testing.T) {
		s := NewSet[string]()
		if len(s) != 0 {
			t.Errorf("expected empty set, got %d members", len(s))
		}
	})

	t.Run("set with initial values", func(t *testing.T) {
		s := NewSet("a", "b", "c")
		if len(s) != 3 {
			t.Errorf("expected 3 members, got %d", len(s))
		}
		if !s.Has("a") || !s.Has("b") || !s.Has("c") {
			t.Error("set missing expected initial values")
		}
	})

	t.Run("set with duplicate initial values", func(t *testing.T) {
		s := NewSet("a", "b", "a", "c", "b")
		if len(s) != 3 {
			t.Errorf("expected 3 unique members, got %d", len(s))
		}
	})

	t.Run("set with integers", func(t *testing.T) {
		s := NewSet(1, 2, 3)
		if len(s) != 3 {
			t.Errorf("expected 3 members, got %d", len(s))
		}
		if !s.Has(1) || !s.Has(2) || !s.Has(3) {
			t.Error("set missing expected initial values")
		}
	})
}

func TestAdd(t *testing.T) {
	t.Run("add to empty set", func(t *testing.T) {
		s := NewSet[string]()
		s.Add("a")
		if !s.Has("a") {
			t.Error("value not added to set")
		}
		if len(s) != 1 {
			t.Errorf("expected 1 member, got %d", len(s))
		}
	})

	t.Run("add multiple values", func(t *testing.T) {
		s := NewSet[string]()
		s.Add("a", "b", "c")
		if len(s) != 3 {
			t.Errorf("expected 3 members, got %d", len(s))
		}
	})

	t.Run("add duplicate values", func(t *testing.T) {
		s := NewSet("a")
		s.Add("a")
		if len(s) != 1 {
			t.Errorf("expected 1 member after duplicate add, got %d", len(s))
		}
	})

	t.Run("add to existing set", func(t *testing.T) {
		s := NewSet("a", "b")
		s.Add("c", "d")
		if len(s) != 4 {
			t.Errorf("expected 4 members, got %d", len(s))
		}
		for _, v := range []string{"a", "b", "c", "d"} {
			if !s.Has(v) {
				t.Errorf("missing expected value: %s", v)
			}
		}
	})

	t.Run("add no values", func(t *testing.T) {
		s := NewSet("a")
		s.Add()
		if len(s) != 1 {
			t.Errorf("expected 1 member after adding nothing, got %d", len(s))
		}
	})
}

func TestHas(t *testing.T) {
	t.Run("has existing value", func(t *testing.T) {
		s := NewSet("a", "b", "c")
		if !s.Has("a") {
			t.Error("Has returned false for existing value")
		}
	})

	t.Run("has non-existing value", func(t *testing.T) {
		s := NewSet("a", "b", "c")
		if s.Has("d") {
			t.Error("Has returned true for non-existing value")
		}
	})

	t.Run("has on empty set", func(t *testing.T) {
		s := NewSet[string]()
		if s.Has("a") {
			t.Error("Has returned true for empty set")
		}
	})

	t.Run("has with different types", func(t *testing.T) {
		s := NewSet(1, 2, 3)
		if !s.Has(2) {
			t.Error("Has returned false for existing integer")
		}
		if s.Has(4) {
			t.Error("Has returned true for non-existing integer")
		}
	})
}

func TestMembers(t *testing.T) {
	t.Run("empty set members", func(t *testing.T) {
		s := NewSet[string]()
		members := s.Members()
		if len(members) != 0 {
			t.Errorf("expected empty slice, got %d members", len(members))
		}
	})

	t.Run("set with members", func(t *testing.T) {
		s := NewSet("a", "b", "c")
		members := s.Members()
		if len(members) != 3 {
			t.Errorf("expected 3 members, got %d", len(members))
		}
		// Check all expected members are present
		memberMap := make(map[string]bool)
		for _, m := range members {
			memberMap[m] = true
		}
		for _, expected := range []string{"a", "b", "c"} {
			if !memberMap[expected] {
				t.Errorf("missing expected member: %s", expected)
			}
		}
	})

	t.Run("members are independent of original set", func(t *testing.T) {
		s := NewSet("a", "b")
		members1 := s.Members()
		s.Add("c")
		members2 := s.Members()
		if len(members1) != 2 {
			t.Errorf("first Members call affected by later Add, got %d", len(members1))
		}
		if len(members2) != 3 {
			t.Errorf("expected 3 members in second call, got %d", len(members2))
		}
	})
}

func TestString(t *testing.T) {
	t.Run("empty set string", func(t *testing.T) {
		s := NewSet[string]()
		str := s.String()
		if str != "[]" {
			t.Errorf("expected '[]', got '%s'", str)
		}
	})

	t.Run("set with single member string", func(t *testing.T) {
		s := NewSet("a")
		str := s.String()
		if str != "[a]" {
			t.Errorf("expected '[a]', got '%s'", str)
		}
	})

	t.Run("set with multiple members string", func(t *testing.T) {
		s := NewSet(1, 2, 3)
		str := s.String()
		// Note: order is not guaranteed in maps, so we just check it's a valid format
		if len(str) < 5 { // At minimum "[1 2 3]" or similar
			t.Errorf("expected string representation, got '%s'", str)
		}
		if str[0] != '[' || str[len(str)-1] != ']' {
			t.Errorf("expected bracketed format, got '%s'", str)
		}
	})
}
