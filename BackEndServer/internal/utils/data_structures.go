package utils

// Set data structure. Uses a map to check for duplicate entries
type Set[T comparable] struct {
	Elements map[T]bool
	Count    int
}

func (s *Set[T]) Contains(val T) bool {
	return s.Elements[val]
}

// Guards against double add
func (s *Set[T]) Add(val T) {
	if s.Contains(val) {
		return
	}
	s.Elements[val] = true
	s.Count = s.Count + 1
}

// Guards against double remove
func (s *Set[T]) Remove(val T) {
	if !s.Contains(val) {
		return
	}
	delete(s.Elements, val)
	s.Count = s.Count - 1
}
