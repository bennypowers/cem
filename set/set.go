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

import "fmt"

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](vs ...T) Set[T] {
	s := Set[T]{}
	s.Add(vs...)
	return s
}

func (s Set[T]) Add(vs ...T) {
	for _, v := range vs {
		s[v] = struct{}{}
	}
}

func (s Set[T]) Has(v T) bool {
	_, ok := s[v]
	return ok
}

func (s Set[T]) Members() []T {
	r := make([]T, 0, len(s))
	for v := range s {
		r = append(r, v)
	}
	return r
}

func (s Set[T]) String() string {
	return fmt.Sprintf("%v", s.Members())
}
