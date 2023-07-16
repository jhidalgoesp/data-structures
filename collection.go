package data_structures

type Collection[T any] interface {
	CreateIterator() Iterator[T]
}
