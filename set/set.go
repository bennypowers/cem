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
