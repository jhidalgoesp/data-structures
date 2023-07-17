package data_structures

type DoubleLinkedListNode[T any] struct {
	data     T
	previous *DoubleLinkedListNode[T]
	next     *DoubleLinkedListNode[T]
}

func NewDoubleLinkedListNode[T any](data T) DoubleLinkedListNode[T] {
	return DoubleLinkedListNode[T]{
		data:     data,
		previous: nil,
		next:     nil,
	}
}

type DoubleLinkedList[T any] struct {
	head *DoubleLinkedListNode[T]
	tail *DoubleLinkedListNode[T]
	size int
}

func NewDoubleLinkedList[T any]() DoubleLinkedList[T] {
	return DoubleLinkedList[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (d *DoubleLinkedList[T]) AddToTheFront(data T) {
	node := NewDoubleLinkedListNode(data)

	if d.size == 0 {
		d.head = &node
		d.tail = &node
		d.size += 1

		return
	}

	d.head.previous = &node
	node.next = d.head
	d.head = &node
	d.size += 1
}

func (d *DoubleLinkedList[T]) AddToTheBack(data T) {
	node := NewDoubleLinkedListNode(data)

	if d.size == 0 {
		d.head = &node
		d.tail = &node
		d.size += 1

		return
	}

	d.tail.next = &node
	node.previous = d.tail
	d.tail = &node
	d.size += 1
}

func (d *DoubleLinkedList[T]) RemoveFromTheFront() *T {
	if d.size == 0 {
		return nil
	}

	var temp *T

	if d.size == 1 {
		temp = &d.head.data
		d.head = nil
		d.tail = nil
		d.size -= 1

		return temp
	}

	temp = &d.head.data
	d.head.next.previous = nil
	d.head = d.head.next
	d.size -= 1

	return temp
}

func (d *DoubleLinkedList[T]) RemoveFromTheBack() *T {
	if d.size == 0 {
		return nil
	}

	var temp *T

	if d.size == 1 {
		temp = &d.head.data
		d.head = nil
		d.tail = nil
		d.size -= 1

		return temp
	}

	temp = &d.tail.data
	d.tail.previous.next = nil
	d.tail = d.tail.previous
	d.size -= 1

	return temp
}

func (l *DoubleLinkedList[T]) CreateIterator() Iterator[T] {
	return &DoubleLinkedListIterator[T]{
		head:  l.head,
		index: 0,
		size:  l.size,
	}
}

type DoubleLinkedListIterator[T any] struct {
	head  *DoubleLinkedListNode[T]
	index int
	size  int
}

func (l *DoubleLinkedListIterator[T]) HasNext() bool {
	return l.index < l.size
}

func (l *DoubleLinkedListIterator[T]) GetNext() *T {
	if l.HasNext() {
		current := l.head

		for i := 0; i < l.index; i++ {
			current = current.next
		}

		l.index += 1
		return &current.data
	}

	return nil
}
