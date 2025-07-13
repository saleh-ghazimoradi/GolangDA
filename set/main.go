package main

import "fmt"

type set struct {
	elements []int
	lookup   map[int]struct{}
}

func (s *set) add(value int) bool {
	if _, exists := s.lookup[value]; exists {
		return false
	}
	s.lookup[value] = struct{}{}
	s.elements = append(s.elements, value)
	return true
}

func (s *set) remove(value int) bool {
	if _, exists := s.lookup[value]; !exists {
		return false
	}

	delete(s.lookup, value)

	for i, v := range s.elements {
		if v == value {
			s.elements = append(s.elements[:i], s.elements[i+1:]...)
			return true
		}
	}
	return false
}

func (s *set) size() int {
	return len(s.elements)
}

func (s *set) contains(value int) bool {
	_, exists := s.lookup[value]
	return exists
}

func (s *set) element() []int {
	return append([]int{}, s.elements...)
}

func (s *set) clear() {
	s.elements = make([]int, 0)
	s.lookup = make(map[int]struct{})
}

func NewSet() *set {
	return &set{
		elements: make([]int, 0),
		lookup:   make(map[int]struct{}),
	}
}

func main() {
	s := NewSet()
	s.add(1)
	s.add(2)
	s.add(1)
	fmt.Println("set elements:", s.element())
	fmt.Println("contain 2:", s.contains(2))
	fmt.Println("size:", s.size())
	s.remove(1)
	fmt.Println("after removing 1:", s.element())
	s.clear()
	fmt.Println("after clearing:", s.element())
}
