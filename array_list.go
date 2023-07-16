package data_structures

const (
	initialSize = 9
)

type ArrayList[T any] struct {
	backingArray []*T
	size         int
}

func NewArrayList[T any]() ArrayList[T] {
	return ArrayList[T]{
		backingArray: make([]*T, initialSize),
		size:         0,
	}
}

func (a *ArrayList[T]) AddToFront(data T) {
	if a.size == len(a.backingArray) {
		a.resize()
	}

	for i := a.size; i > 0; i-- {
		if i-1 < 0 {
			break
		}

		a.backingArray[i] = a.backingArray[i-1]
	}

	a.backingArray[0] = &data
	a.size += 1
}

func (a *ArrayList[T]) AddToBack(data T) {
	if a.size == len(a.backingArray) {
		a.resize()
	}

	a.backingArray[a.size] = &data
	a.size += 1
}

func (a *ArrayList[T]) RemoveFromFront() {
	for i := 0; i < a.size; i++ {
		a.backingArray[i] = a.backingArray[i+1]
	}

	a.size -= 1
}

func (a *ArrayList[T]) RemoveFromBack() {
	a.size -= 1
	a.backingArray[a.size] = nil
}

func (a *ArrayList[T]) resize() {
	temp := make([]*T, a.size*2)

	copy(temp, a.backingArray)

	a.backingArray = temp
}
