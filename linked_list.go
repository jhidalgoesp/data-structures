package data_structures

type LinkedListNode[T any] struct {
	Data T
	Next *LinkedListNode[T]
}

func NewLinkedListNode[T any](data T) LinkedListNode[T] {
	return LinkedListNode[T]{
		Data: data,
		Next: nil,
	}
}

type LinkedList[T any] struct {
	head *LinkedListNode[T]
	tail *LinkedListNode[T]
	size int
}

func NewLinkedList[T any]() LinkedList[T] {
	return LinkedList[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (l *LinkedList[T]) AddToTheFront(data T) {
	node := NewLinkedListNode(data)

	node.Next = l.head

	l.head = &node

	if l.size == 0 {
		l.tail = &node
	}

	l.size += 1
}

func (l *LinkedList[T]) AddToTheBack(data T) {
	node := NewLinkedListNode(data)

	if l.size == 0 {
		l.head = &node
		l.tail = &node

		l.size += 1

		return
	}

	l.tail.Next = &node

	l.tail = &node

	l.size += 1
}

func (l *LinkedList[T]) RemoveFromTheFront() *T {
	if l.size == 0 {
		return nil
	}

	temp := l.head.Data

	if l.size == 1 {
		l.head = nil
		l.tail = nil
		l.size -= 1

		return &temp
	}

	l.head = l.head.Next
	l.size -= 1

	return &temp
}

func (l *LinkedList[T]) RemoveFromTheBack() *T {
	if l.size == 0 {
		return nil
	}

	temp := l.tail.Data

	if l.size == 1 {
		l.head = nil
		l.tail = nil
		l.size -= 1

		return &temp
	}

	current := l.head

	for current.Next.Next != nil {
		current = current.Next
	}

	current.Next = nil
	l.tail = current
	l.size -= 1

	return &temp
}

func (l *LinkedList[T]) CreateIterator() Iterator[T] {
	return &LinkedListIterator[T]{
		head:  l.head,
		index: 0,
		size:  l.size,
	}
}

type LinkedListIterator[T any] struct {
	head  *LinkedListNode[T]
	index int
	size  int
}

func (l *LinkedListIterator[T]) HasNext() bool {
	return l.index < l.size
}

func (l *LinkedListIterator[T]) GetNext() *T {
	if l.HasNext() {
		current := l.head

		for i := 0; i < l.index; i++ {
			current = current.Next
		}

		l.index += 1
		return &current.Data
	}

	return nil
}
