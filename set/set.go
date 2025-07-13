package set

type Set struct {
	Elements []int
	Lookup   map[int]struct{}
}

func (s *Set) Add(value int) bool {
	if _, exists := s.Lookup[value]; exists {
		return false
	}
	s.Lookup[value] = struct{}{}
	s.Elements = append(s.Elements, value)
	return true
}

func (s *Set) Remove(value int) bool {
	if _, exists := s.Lookup[value]; !exists {
		return false
	}

	delete(s.Lookup, value)

	for i, v := range s.Elements {
		if v == value {
			s.Elements = append(s.Elements[:i], s.Elements[i+1:]...)
			return true
		}
	}
	return false
}

func (s *Set) Size() int {
	return len(s.Elements)
}

func (s *Set) Contains(value int) bool {
	_, exists := s.Lookup[value]
	return exists
}

func (s *Set) Element() []int {
	return append([]int{}, s.Elements...)
}

func (s *Set) Clear() {
	s.Elements = make([]int, 0)
	s.Lookup = make(map[int]struct{})
}

func NewSet() *Set {
	return &Set{
		Elements: make([]int, 0),
		Lookup:   make(map[int]struct{}),
	}
}
