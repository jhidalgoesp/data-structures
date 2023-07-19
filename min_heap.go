package data_structures

import "fmt"

// ComparatorFunc should return true if the first element is lesser than the second element
type ComparatorFunc[T any] func(a, b T) bool

type MinHeap[T any] struct {
	backingArray []*T
	size         int
	less         ComparatorFunc[T]
}

func NewMinHeap[T any](comparatorFunc ComparatorFunc[T]) (*MinHeap[T], error) {
	if comparatorFunc == nil {
		return nil, fmt.Errorf("comparatorFunc is requirad")
	}

	return &MinHeap[T]{
		backingArray: make([]*T, initialSize),
		less:         comparatorFunc,
		size:         0,
	}, nil
}

func (m *MinHeap[T]) Add(data T) {
	m.backingArray[m.size+1] = &data

	m.size += 1

	m.upHeap(m.size)
}

func (m *MinHeap[T]) Remove() *T {
	if m.size == 0 {
		return nil
	}

	temp := *m.backingArray[1]

	m.backingArray[1] = m.backingArray[m.size]

	m.backingArray[m.size] = nil

	m.downHeap(1)

	m.size -= 1

	return &temp
}

func (m *MinHeap[T]) downHeap(index int) {
	var leftChild, rightChild *T

	parent := m.backingArray[index]

	if index*2 <= m.size {
		leftChild = m.backingArray[index*2]
	}

	if index*2+1 <= m.size {
		rightChild = m.backingArray[index*2+1]
	}

	if leftChild != nil && rightChild != nil {
		if m.less(*leftChild, *rightChild) {
			m.backingArray[index], m.backingArray[index*2] = m.backingArray[index*2], m.backingArray[index]
			m.downHeap(index * 2)
		}

		if m.less(*rightChild, *leftChild) {
			m.backingArray[index], m.backingArray[index*2+1] = m.backingArray[index*2+1], m.backingArray[index]
			m.downHeap(index*2 + 1)
		}

		return
	}

	if leftChild != nil && m.less(*leftChild, *parent) {
		m.backingArray[index], m.backingArray[index*2] = m.backingArray[index*2], m.backingArray[index]
		m.downHeap(index * 2)
	}

	if rightChild != nil && m.less(*rightChild, *parent) {
		m.backingArray[index], m.backingArray[index*2+1] = m.backingArray[index*2+1], m.backingArray[index]
		m.downHeap(index*2 + 1)
	}
}

func (m *MinHeap[T]) upHeap(index int) {
	if index == 1 {
		return
	}

	node := *m.backingArray[index]
	parent := *m.backingArray[index/2]

	if m.less(node, parent) {
		m.backingArray[index/2], m.backingArray[index] = &node, &parent
	}

	m.upHeap(index / 2)
}
