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

// Package set provides a generic set data structure implementation.
package set

import "fmt"

// Set is a generic set data structure that stores unique values of any comparable type.
type Set[T comparable] map[T]struct{}

// NewSet creates and returns a new Set containing the provided values.
// Duplicate values in the input are automatically deduplicated.
func NewSet[T comparable](vs ...T) Set[T] {
	s := Set[T]{}
	s.Add(vs...)
	return s
}

// Add adds one or more values to the set.
// If a value already exists in the set, it is not added again.
func (s Set[T]) Add(vs ...T) {
	for _, v := range vs {
		s[v] = struct{}{}
	}
}

// Has returns true if the set contains the specified value.
func (s Set[T]) Has(v T) bool {
	_, ok := s[v]
	return ok
}

// Members returns a slice containing all values in the set.
// The returned slice is independent of the set and can be modified without affecting the set.
// The order of elements is not guaranteed to be consistent across calls.
func (s Set[T]) Members() []T {
	r := make([]T, 0, len(s))
	for v := range s {
		r = append(r, v)
	}
	return r
}

// String returns a string representation of the set in the format "[value1 value2 ...]".
func (s Set[T]) String() string {
	return fmt.Sprintf("%v", s.Members())
}
