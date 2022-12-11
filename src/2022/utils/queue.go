package utils

import "fmt"

type Queue[T any] struct {
	head   *Node[T]
	tail   *Node[T]
	length int
}

func (q *Queue[T]) Enqueue(value T) {
	q.length++

	newNode := createNode(value)

	if q.head == nil {
		q.head, q.tail = newNode, newNode
		return
	}

	q.tail.prev = newNode
	q.tail = newNode
}

func (q *Queue[T]) Deque() (T, error) {
	if q.head == nil {
		var out T
		return out, fmt.Errorf("queue is empty")
	}

	q.length--

	node := q.head
	q.head = q.head.prev
	out := node.value

	// Free
	node.prev = nil
	node = nil

	if q.length == 0 {
		q.tail = nil
	}

	return out, nil
}

func (q *Queue[T]) Peek() (T, error) {
	if q.head == nil {
		var out T
		return out, fmt.Errorf("queue is empty")
	}

	return q.head.value, nil
}

func (q *Queue[T]) Length() int {
	return q.length
}
