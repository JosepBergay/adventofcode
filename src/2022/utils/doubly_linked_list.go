package utils

import "fmt"

type DoublyLinkedList[T comparable] struct {
	head   *Node[T]
	tail   *Node[T]
	length int
}

func (l *DoublyLinkedList[T]) Append(item T) {
	l.length++

	node := createNode(item)

	if l.head == nil {
		l.head = node
		l.tail = l.head
		return
	}

	l.tail.next = node
	node.prev = l.tail
	l.tail = node
}

func (l *DoublyLinkedList[T]) Get(index int) (T, error) {
	if index >= l.length {
		return createNotFoundError[T]()
	}

	curr := l.head
	for i := 0; i < index; i++ {
		curr = curr.next
	}

	return curr.value, nil
}

func (l *DoublyLinkedList[T]) GetIndex(value T) (int, error) {
	curr := l.head
	for i := 0; i < l.length; i++ {
		if curr.value == value {
			return i, nil
		}
		curr = curr.next
	}

	return createNotFoundError[int]()
}

func (l *DoublyLinkedList[T]) InsertAt(item T, idx int) error {
	if idx > l.length || idx < 0 {
		return createOutOfBoundsError(l.length, idx)
	}

	l.length++

	newNode := createNode(item)

	if l.head == nil {
		l.head = newNode
		l.tail = l.head
		return nil
	}

	i := 0
	curr := l.head
	for i < idx-1 {
		curr = curr.next
		i++
	}

	if curr == nil {
		fmt.Println(item, idx, l.length-1, i, curr)
	}

	newNode.next = curr.next
	if curr.next != nil {
		curr.next.prev = newNode
	}

	newNode.prev = curr
	curr.next = newNode

	switch idx {
	case 0:
		l.head = newNode
	case l.length - 1:
		l.tail = newNode
	}

	return nil
}

func (l *DoublyLinkedList[T]) Length() int {
	return l.length
}

func (l *DoublyLinkedList[T]) Prepend(item T) {
	l.length++

	node := createNode(item)

	if l.head == nil {
		l.head = node
		l.tail = l.head
		return
	}

	l.head.prev = node
	node.next = l.head
	l.head = node
}

func (l *DoublyLinkedList[T]) Remove(item T) (int, error) {
	// Starting from head, could also start from tail
	curr := l.head
	i := 0
	for i < l.length {
		if curr.value == item {
			break
		}
		curr = curr.next
		i++
	}

	if curr == nil {
		// We got to the end of the list but didn't find item.
		_, err := createNotFoundError[T]()
		return -1, err
	}

	l.breakLinksAndFreeNode(curr)

	l.length--

	return i, nil
}

func (l *DoublyLinkedList[T]) RemoveAt(index int) (T, error) {
	if index >= l.length {
		return createNotFoundError[T]()
	}

	curr := l.head
	for i := 0; i < index; i++ {
		curr = curr.next
	}

	out := l.breakLinksAndFreeNode(curr)

	l.length--

	return out, nil
}

// BreakLinksAndFreeNode will break links and point node to nil.
func (l *DoublyLinkedList[T]) breakLinksAndFreeNode(curr *Node[T]) T {
	if curr.prev != nil {
		curr.prev.next = curr.next
	} else {
		// We are deleting head so point head to next
		l.head = l.head.next
	}

	if curr.next != nil {
		curr.next.prev = curr.prev
	} else {
		// We are deleting tail so point tail to prev
		l.tail = l.tail.prev
	}

	out := curr.value

	// Free
	curr.next = nil
	curr.prev = nil
	curr = nil

	return out
}

func createNotFoundError[T any]() (T, error) {
	var out T
	return out, fmt.Errorf("not found")
}

func createOutOfBoundsError(len, idx int) error {
	return fmt.Errorf("index %v is out of bounds. Current length: %v", idx, len)
}
