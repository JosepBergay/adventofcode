package utils

import "fmt"

type Node[T any] struct {
	value T
	prev  *Node[T]
}

type Stack[T any] struct {
	tail   *Node[T]
	length int
}

func createNode[T any](value T) *Node[T] {
	return &Node[T]{value: value}
}

func (s *Stack[T]) Push(item T) {
	s.length++

	newNode := createNode(item)

	if s.tail == nil {
		s.tail = newNode
		return
	}

	newNode.prev = s.tail
	s.tail = newNode
}

func (s *Stack[T]) Pop() (T, error) {
	if s.tail == nil {
		var out T
		return out, fmt.Errorf("stack is empty")
	}

	s.length--

	node := s.tail
	s.tail = s.tail.prev
	out := node.value

	// Free
	node.prev = nil
	node = nil

	return out, nil

}

func (s *Stack[T]) Length() int {
	return s.length
}

// Reverse will make items on top of the stack appear on bottom. This is not why stacks are made!
func (s *Stack[T]) Reverse() {
	items := make([]T, s.length)

	for s.tail != nil {
		v, err := s.Pop()
		if err != nil {
			panic("SOMETHING's WRONG SHERLOCK")
		}

		items = append(items, v)
	}

	for _, item := range items {
		s.Push(item)
	}
}
