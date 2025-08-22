package utils

type Set[T comparable] struct {
	Elements map[T]bool
}

func (s *Set[T]) Contains(val T) bool {
	return s.Elements[val]
}

func (s *Set[T]) Add(val T) {
	s.Elements[val] = true
}

func (s *Set[T]) Remove(val T) {
	delete(s.Elements, val)
}
