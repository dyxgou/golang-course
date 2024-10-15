package maps

import (
	"errors"
	"fmt"
	"log"
)

type Set[v comparable] interface {
	insertElement(value v) error
	deleteElement(value v) error
}

type set[v comparable] struct {
	elements map[v]struct{}
}

func (s *set[v]) insertElement(value v) error {
	if _, ok := s.elements[value]; ok {
		return errors.New("The value is already in the set")
	}

	s.elements[value] = struct{}{}

	return nil
}

func (s *set[v]) deleteElement(value v) error {
	if _, ok := s.elements[value]; !ok {
		return errors.New("The element doesnt exists in the set")
	}

	delete(s.elements, value)

	return nil
}

func NewSet[v comparable]() *set[v] {
	return &set[v]{
		elements: make(map[v]struct{}),
	}
}

func Sets() {
	intSet := NewSet[int]()

	intSet.insertElement(13)
	intSet.insertElement(14)
	intSet.insertElement(15)
	intSet.insertElement(16)

	log.Fatal(intSet.insertElement(16)) // Should return an error

	intSet.insertElement(17)
	intSet.insertElement(18)
	intSet.insertElement(19)

	fmt.Printf("intSet.elements: %v\n", intSet.elements)
}
