package data_structures

const (
	initialSize = 9
)

type Collection[T any] interface {
	CreateIterator() Iterator[T]
}
